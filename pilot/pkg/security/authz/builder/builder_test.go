// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package builder

import (
	"strings"
	"testing"

	http_config "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/rbac/v2"
	tcp_config "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/rbac/v2"
	"github.com/envoyproxy/go-control-plane/pkg/util"

	istio_rbac "istio.io/api/rbac/v1alpha1"
	"istio.io/istio/pilot/pkg/model"
	authz_model "istio.io/istio/pilot/pkg/security/authz/model"
	"istio.io/istio/pilot/pkg/security/authz/policy"
	"istio.io/istio/pkg/config"
)

func newAuthzPolicyWithRbacConfig(mode istio_rbac.RbacConfig_Mode, include *istio_rbac.RbacConfig_Target,
	exclude *istio_rbac.RbacConfig_Target) *model.AuthorizationPolicies {
	return &model.AuthorizationPolicies{
		RbacConfig: &istio_rbac.RbacConfig{
			Mode:      mode,
			Inclusion: include,
			Exclusion: exclude,
		},
	}
}

func newService(hostname string, labels map[string]string, t *testing.T) *model.ServiceInstance {
	t.Helper()
	splits := strings.Split(hostname, ".")
	if len(splits) < 2 {
		t.Fatalf("failed to initialize service instance: invalid hostname")
	}
	name := splits[0]
	namespace := splits[1]
	return &model.ServiceInstance{
		Service: &model.Service{
			Attributes: model.ServiceAttributes{
				Name:      name,
				Namespace: namespace,
			},
			Hostname: config.Hostname(hostname),
		},
		Labels: labels,
	}
}

func simpleGlobalPermissiveMode() *model.Config {
	cfg := &model.Config{
		ConfigMeta: model.ConfigMeta{
			Type:      model.ClusterRbacConfig.Type,
			Name:      "default",
			Namespace: "default",
		},
		Spec: &istio_rbac.RbacConfig{
			Mode:            istio_rbac.RbacConfig_ON,
			EnforcementMode: istio_rbac.EnforcementMode_PERMISSIVE,
		},
	}
	return cfg
}

func TestIsRbacEnabled(t *testing.T) {
	target := &istio_rbac.RbacConfig_Target{
		Services:   []string{"review.default.svc", "product.default.svc"},
		Namespaces: []string{"special"},
	}
	cfg1 := newAuthzPolicyWithRbacConfig(istio_rbac.RbacConfig_ON, nil, nil)
	cfg2 := newAuthzPolicyWithRbacConfig(istio_rbac.RbacConfig_OFF, nil, nil)
	cfg3 := newAuthzPolicyWithRbacConfig(istio_rbac.RbacConfig_ON_WITH_INCLUSION, target, nil)
	cfg4 := newAuthzPolicyWithRbacConfig(istio_rbac.RbacConfig_ON_WITH_EXCLUSION, nil, target)
	cfg5 := newAuthzPolicyWithRbacConfig(istio_rbac.RbacConfig_ON, nil, nil)

	testCases := []struct {
		Name          string
		AuthzPolicies *model.AuthorizationPolicies
		Service       string
		Namespace     string
		want          bool
	}{
		{
			Name:          "rbac plugin enabled",
			AuthzPolicies: cfg1,
			want:          true,
		},
		{
			Name:          "rbac plugin disabled",
			AuthzPolicies: cfg2,
		},
		{
			Name:          "rbac plugin enabled by inclusion.service",
			AuthzPolicies: cfg3,
			Service:       "product.default.svc",
			Namespace:     "default",
			want:          true,
		},
		{
			Name:          "rbac plugin enabled by inclusion.namespace",
			AuthzPolicies: cfg3,
			Service:       "other.special.svc",
			Namespace:     "special",
			want:          true,
		},
		{
			Name:          "rbac plugin disabled by exclusion.service",
			AuthzPolicies: cfg4,
			Service:       "product.default.svc",
			Namespace:     "default",
		},
		{
			Name:          "rbac plugin disabled by exclusion.namespace",
			AuthzPolicies: cfg4,
			Service:       "other.special.svc",
			Namespace:     "special",
		},
		{
			Name:          "rbac plugin enabled with permissive",
			AuthzPolicies: cfg5,
			want:          true,
		},
	}

	for _, tc := range testCases {
		got := isRbacEnabled(tc.Service, tc.Namespace, tc.AuthzPolicies)
		if tc.want != got {
			t.Errorf("%s: expecting %v but got %v", tc.Name, tc.want, got)
		}
	}
}

func TestBuilder_BuildHTTPFilter(t *testing.T) {
	service := newService("bar.a.svc.cluster.local", nil, t)

	testCases := []struct {
		name                        string
		policies                    []*model.Config
		isXDSMarshalingToAnyEnabled bool
		wantRuleWithPolicies        bool
	}{
		{
			name:                        "XDSMarshalingToAnyEnabled",
			isXDSMarshalingToAnyEnabled: true,
		},
		{
			name: "HTTP rule",
			policies: []*model.Config{
				policy.SimpleRole("role-1", "a", "bar"),
				policy.SimpleBinding("binding-1", "a", "role-1"),
			},
			wantRuleWithPolicies: true,
		},
	}

	for _, tc := range testCases {
		p := policy.NewAuthzPolicies(tc.policies, t)
		b := NewBuilder(service, p, tc.isXDSMarshalingToAnyEnabled)

		got := b.BuildHTTPFilter()
		if got.Name != authz_model.RBACHTTPFilterName {
			t.Errorf("%s: got filter name %q but want %q", tc.name, got.Name, authz_model.RBACHTTPFilterName)
		}

		if tc.isXDSMarshalingToAnyEnabled {
			if got.GetTypedConfig() == nil {
				t.Errorf("%s: want typed config when isXDSMarshalingToAnyEnabled is true", tc.name)
			}
		} else {
			rbacConfig := &http_config.RBAC{}
			if got.GetConfig() == nil {
				t.Errorf("%s: want struct config when isXDSMarshalingToAnyEnabled is false", tc.name)
			} else if err := util.StructToMessage(got.GetConfig(), rbacConfig); err != nil {
				t.Errorf("%s: failed to convert struct to message: %s", tc.name, err)
			} else if len(rbacConfig.GetRules().GetPolicies()) > 0 != tc.wantRuleWithPolicies {
				t.Errorf("%s: got rules with policies %v but want %v",
					tc.name, len(rbacConfig.GetRules().GetPolicies()) > 0, tc.wantRuleWithPolicies)
			}
		}
	}
}

func TestBuilder_BuildTCPFilter(t *testing.T) {
	service := newService("foo.a.svc.cluster.local", nil, t)

	testCases := []struct {
		name                        string
		policies                    []*model.Config
		isXDSMarshalingToAnyEnabled bool
		wantRules                   bool
		wantRuleWithPolicies        bool
		wantShadowRules             bool
	}{
		{
			name:                        "XDSMarshalingToAnyEnabled",
			isXDSMarshalingToAnyEnabled: true,
		},
		{
			name: "HTTP rule",
			policies: []*model.Config{
				policy.SimpleRole("role-1", "a", "foo"),
				policy.SimpleBinding("binding-1", "a", "role-1"),
			},
			wantRules:            true,
			wantRuleWithPolicies: false,
		},
		{
			name:      "normal rule",
			wantRules: true,
		},
		{
			name: "normal shadow rule",
			policies: []*model.Config{
				simpleGlobalPermissiveMode(),
			},
			wantShadowRules: true,
		},
	}

	for _, tc := range testCases {
		p := policy.NewAuthzPolicies(tc.policies, t)
		b := NewBuilder(service, p, tc.isXDSMarshalingToAnyEnabled)

		got := b.BuildTCPFilter()
		if got.Name != authz_model.RBACTCPFilterName {
			t.Errorf("%s: got filter name %q but want %q", tc.name, got.Name, authz_model.RBACTCPFilterName)
		}

		if tc.isXDSMarshalingToAnyEnabled {
			if got.GetTypedConfig() == nil {
				t.Errorf("%s: want typed config when isXDSMarshalingToAnyEnabled is true", tc.name)
			}
		} else {
			rbacConfig := &tcp_config.RBAC{}
			if got.GetConfig() == nil {
				t.Errorf("%s: want struct config when isXDSMarshalingToAnyEnabled is false", tc.name)
			} else if err := util.StructToMessage(got.GetConfig(), rbacConfig); err != nil {
				t.Errorf("%s: failed to convert struct to message: %s", tc.name, err)
			} else {
				if rbacConfig.StatPrefix != authz_model.RBACTCPFilterStatPrefix {
					t.Errorf("%s: got filter stat prefix %q but want %q",
						tc.name, rbacConfig.StatPrefix, authz_model.RBACTCPFilterStatPrefix)
				}

				if len(rbacConfig.GetRules().GetPolicies()) > 0 != tc.wantRuleWithPolicies {
					t.Errorf("%s: got rules with policies %v but want %v",
						tc.name, len(rbacConfig.GetRules().GetPolicies()) > 0, tc.wantRuleWithPolicies)
				}
				if (rbacConfig.GetRules().GetPolicies() != nil) != tc.wantRules {
					t.Errorf("%s: got rules %v but want %v",
						tc.name, rbacConfig.GetRules().GetPolicies() != nil, tc.wantRules)
				}
				if (rbacConfig.GetShadowRules().GetPolicies() != nil) != tc.wantShadowRules {
					t.Errorf("%s: got shadow rules %v but want %v",
						tc.name, rbacConfig.GetShadowRules().GetPolicies() != nil, tc.wantShadowRules)
				}
			}
		}
	}
}
