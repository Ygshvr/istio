// Code generated for package validation by go-bindata DO NOT EDIT. (@generated)
// sources:
// dataset/authentication-v1alpha1-MeshPolicy-invalid.yaml
// dataset/authentication-v1alpha1-MeshPolicy-valid.yaml
// dataset/authentication-v1alpha1-Policy-invalid.yaml
// dataset/authentication-v1alpha1-Policy-valid.yaml
// dataset/config-v1alpha2-HTTPAPISpec-invalid.yaml
// dataset/config-v1alpha2-HTTPAPISpec-valid.yaml
// dataset/config-v1alpha2-HTTPAPISpecBinding-invalid.yaml
// dataset/config-v1alpha2-HTTPAPISpecBinding-valid.yaml
// dataset/config-v1alpha2-QuotaSpec-invalid.yaml
// dataset/config-v1alpha2-QuotaSpec-valid.yaml
// dataset/config-v1alpha2-QuotaSpecBinding-invalid.yaml
// dataset/config-v1alpha2-QuotaSpecBinding-valid.yaml
// dataset/config-v1alpha2-adapter-invalid.yaml
// dataset/config-v1alpha2-adapter-valid.yaml
// dataset/config-v1alpha2-attributemanifest-invalid.yaml
// dataset/config-v1alpha2-attributemanifest-valid.yaml
// dataset/config-v1alpha2-handler-invalid.yaml
// dataset/config-v1alpha2-handler-valid.yaml
// dataset/config-v1alpha2-instance-invalid.yaml
// dataset/config-v1alpha2-instance-valid.yaml
// dataset/config-v1alpha2-rule-invalid.yaml
// dataset/config-v1alpha2-rule-valid.yaml
// dataset/config-v1alpha2-template-invalid.yaml
// dataset/config-v1alpha2-template-valid.yaml
// dataset/networking-v1alpha3-DestinationRule-invalid.yaml
// dataset/networking-v1alpha3-DestinationRule-valid.yaml
// dataset/networking-v1alpha3-EnvoyFilter-invalid.yaml
// dataset/networking-v1alpha3-EnvoyFilter-valid.yaml
// dataset/networking-v1alpha3-Gateway-invalid.yaml
// dataset/networking-v1alpha3-Gateway-valid.yaml
// dataset/networking-v1alpha3-ServiceEntry-invalid.yaml
// dataset/networking-v1alpha3-ServiceEntry-valid.yaml
// dataset/networking-v1alpha3-Sidecar-invalid.yaml
// dataset/networking-v1alpha3-Sidecar-valid.yaml
// dataset/networking-v1alpha3-VirtualService-invalid.yaml
// dataset/networking-v1alpha3-VirtualService-valid.yaml
// dataset/networking-v1beta-DestinationRule-invalid.yaml
// dataset/networking-v1beta-DestinationRule-valid.yaml
// dataset/networking-v1beta-Gateway-invalid.yaml
// dataset/networking-v1beta-Gateway-valid.yaml
// dataset/networking-v1beta-Sidecar-invalid.yaml
// dataset/networking-v1beta-Sidecar-valid.yaml
// dataset/networking-v1beta-VirtualService-invalid.yaml
// dataset/networking-v1beta-VirtualService-valid.yaml
// dataset/rbac-v1alpha1-ClusterRbacConfig-invalid.yaml
// dataset/rbac-v1alpha1-ClusterRbacConfig-valid.yaml
// dataset/rbac-v1alpha1-RBacConfig-invalid.yaml
// dataset/rbac-v1alpha1-RBacConfig-valid.yaml
// dataset/rbac-v1alpha1-ServiceRole-invalid.yaml
// dataset/rbac-v1alpha1-ServiceRole-valid.yaml
// dataset/rbac-v1alpha1-ServiceRoleBinding-invalid.yaml
// dataset/rbac-v1alpha1-ServiceRoleBinding-valid.yaml
// dataset/security-v1beta1-AuthorizationPolicy-invalid.yaml
// dataset/security-v1beta1-AuthorizationPolicy-valid.yaml
// dataset/security-v1beta1-PeerAuthentication-invalid.yaml
// dataset/security-v1beta1-PeerAuthentication-valid.yaml
// dataset/security-v1beta1-RequestAuthentication-invalid.yaml
// dataset/security-v1beta1-RequestAuthentication-valid.yaml
package validation

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _datasetAuthenticationV1alpha1MeshpolicyInvalidYaml = []byte(`apiVersion: authentication.istio.io/v1alpha1
kind: MeshPolicy
metadata:
  name: "boo"
spec:
  peers:
    - mtls:
        mode: PERMISSIVE
`)

func datasetAuthenticationV1alpha1MeshpolicyInvalidYamlBytes() ([]byte, error) {
	return _datasetAuthenticationV1alpha1MeshpolicyInvalidYaml, nil
}

func datasetAuthenticationV1alpha1MeshpolicyInvalidYaml() (*asset, error) {
	bytes, err := datasetAuthenticationV1alpha1MeshpolicyInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/authentication-v1alpha1-MeshPolicy-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetAuthenticationV1alpha1MeshpolicyValidYaml = []byte(`apiVersion: authentication.istio.io/v1alpha1
kind: MeshPolicy
metadata:
  name: "default"
spec:
  peers:
    - mtls:
        mode: PERMISSIVE
`)

func datasetAuthenticationV1alpha1MeshpolicyValidYamlBytes() ([]byte, error) {
	return _datasetAuthenticationV1alpha1MeshpolicyValidYaml, nil
}

func datasetAuthenticationV1alpha1MeshpolicyValidYaml() (*asset, error) {
	bytes, err := datasetAuthenticationV1alpha1MeshpolicyValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/authentication-v1alpha1-MeshPolicy-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetAuthenticationV1alpha1PolicyInvalidYaml = []byte(`apiVersion: "authentication.istio.io/v1alpha1"
kind: "Policy"
metadata:
  name: invalid-authentication-policy
spec:
  targets:
  - name: "bad.target"
`)

func datasetAuthenticationV1alpha1PolicyInvalidYamlBytes() ([]byte, error) {
	return _datasetAuthenticationV1alpha1PolicyInvalidYaml, nil
}

func datasetAuthenticationV1alpha1PolicyInvalidYaml() (*asset, error) {
	bytes, err := datasetAuthenticationV1alpha1PolicyInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/authentication-v1alpha1-Policy-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetAuthenticationV1alpha1PolicyValidYaml = []byte(`apiVersion: "authentication.istio.io/v1alpha1"
kind: "Policy"
metadata:
  name: valid-authentication-policy
spec:
  targets:
  - name: good-target
`)

func datasetAuthenticationV1alpha1PolicyValidYamlBytes() ([]byte, error) {
	return _datasetAuthenticationV1alpha1PolicyValidYaml, nil
}

func datasetAuthenticationV1alpha1PolicyValidYaml() (*asset, error) {
	bytes, err := datasetAuthenticationV1alpha1PolicyValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/authentication-v1alpha1-Policy-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2HttpapispecInvalidYaml = []byte(`apiVersion: config.istio.io/v1alpha2
kind: HTTPAPISpec
metadata:
  name: invalid-http-api-spec
spec:
  apiKeys:
  - query: key
  attributes:
    attributes:
      api.service:
        stringValue: bookinfo.endpoints.istio-manlinl.cloud.goog
      api.version:
        stringValue: v1
`)

func datasetConfigV1alpha2HttpapispecInvalidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2HttpapispecInvalidYaml, nil
}

func datasetConfigV1alpha2HttpapispecInvalidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2HttpapispecInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-HTTPAPISpec-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2HttpapispecValidYaml = []byte(`apiVersion: config.istio.io/v1alpha2
kind: HTTPAPISpec
metadata:
  name: valid-http-api-spec
spec:
  apiKeys:
  - query: key
  attributes:
    attributes:
      api.service:
        stringValue: bookinfo.endpoints.istio-manlinl.cloud.goog
      api.version:
        stringValue: v1
  patterns:
  - attributes:
      attributes:
        api.operation:
          stringValue: getProducts
    httpMethod: GET
    uriTemplate: /productpage
`)

func datasetConfigV1alpha2HttpapispecValidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2HttpapispecValidYaml, nil
}

func datasetConfigV1alpha2HttpapispecValidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2HttpapispecValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-HTTPAPISpec-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2HttpapispecbindingInvalidYaml = []byte(`apiVersion: config.istio.io/v1alpha2
kind: HTTPAPISpecBinding
metadata:
  name: invalid-http-api-spec-binding
spec:
  services:
`)

func datasetConfigV1alpha2HttpapispecbindingInvalidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2HttpapispecbindingInvalidYaml, nil
}

func datasetConfigV1alpha2HttpapispecbindingInvalidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2HttpapispecbindingInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-HTTPAPISpecBinding-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2HttpapispecbindingValidYaml = []byte(`apiVersion: config.istio.io/v1alpha2
kind: HTTPAPISpecBinding
metadata:
  name: valid-http-api-spec-binding
spec:
  apiSpecs:
  - name: productpage
    namespace: default
  services:
  - name: productpage
    namespace: default
`)

func datasetConfigV1alpha2HttpapispecbindingValidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2HttpapispecbindingValidYaml, nil
}

func datasetConfigV1alpha2HttpapispecbindingValidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2HttpapispecbindingValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-HTTPAPISpecBinding-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2QuotaspecInvalidYaml = []byte(`apiVersion: config.istio.io/v1alpha2
kind: QuotaSpec
metadata:
  name: invalid-quota-spec
spec:
  quota:
`)

func datasetConfigV1alpha2QuotaspecInvalidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2QuotaspecInvalidYaml, nil
}

func datasetConfigV1alpha2QuotaspecInvalidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2QuotaspecInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-QuotaSpec-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2QuotaspecValidYaml = []byte(`apiVersion: config.istio.io/v1alpha2
kind: QuotaSpec
metadata:
  name: valid-quota-spec
spec:
  rules:
  - match:
    - clause:
        api.operation:
          exact: getProducts
    quotas:
    - charge: 1
      quota: read-requests
`)

func datasetConfigV1alpha2QuotaspecValidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2QuotaspecValidYaml, nil
}

func datasetConfigV1alpha2QuotaspecValidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2QuotaspecValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-QuotaSpec-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2QuotaspecbindingInvalidYaml = []byte(`apiVersion: config.istio.io/v1alpha2
kind: QuotaSpecBinding
metadata:
  name: valid-quota-spec-binding
spec:
  services:`)

func datasetConfigV1alpha2QuotaspecbindingInvalidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2QuotaspecbindingInvalidYaml, nil
}

func datasetConfigV1alpha2QuotaspecbindingInvalidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2QuotaspecbindingInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-QuotaSpecBinding-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2QuotaspecbindingValidYaml = []byte(`apiVersion: config.istio.io/v1alpha2
kind: QuotaSpecBinding
metadata:
  name: valid-quota-spec-binding
spec:
  quotaSpecs:
  - name: bookinfo
    namespace: default
  services:
  - name: bookinfo
    namespace: default
`)

func datasetConfigV1alpha2QuotaspecbindingValidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2QuotaspecbindingValidYaml, nil
}

func datasetConfigV1alpha2QuotaspecbindingValidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2QuotaspecbindingValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-QuotaSpecBinding-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2AdapterInvalidYaml = []byte(`apiVersion: "config.istio.io/v1alpha2"
kind: adapter
metadata:
  name: keyval
spec:
  description: 0
  session_based: false
  templates:
  - keyval
  config: CsgCCh5taXhlci90ZXN0L2tleXZhbC9jb25maWcucHJvdG8SBmtleXZhbCJzCgZQYXJhbXMSLwoFdGFibGUYASADKAsyGS5rZXl2YWwuUGFyYW1zLlRhYmxlRW50cnlSBXRhYmxlGjgKClRhYmxlRW50cnkSEAoDa2V5GAEgASgJUgNrZXkSFAoFdmFsdWUYAiABKAlSBXZhbHVlOgI4AUqgAQoGEgQAAAgBCggKAQwSAwAAEgoICgECEgMCCA4KIAoCBAASBAUACAEaFCBBZGFwdGVyIHBhcmFtZXRlcnMKCgoKAwQAARIDBQgOChsKBAQAAgASAwcCIBoOIExvb2t1cCB0YWJsZQoKDQoFBAACAAQSBAcCBRAKDAoFBAACAAYSAwcCFQoMCgUEAAIAARIDBxYbCgwKBQQAAgADEgMHHh9iBnByb3RvMw==
---
`)

func datasetConfigV1alpha2AdapterInvalidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2AdapterInvalidYaml, nil
}

func datasetConfigV1alpha2AdapterInvalidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2AdapterInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-adapter-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2AdapterValidYaml = []byte(`apiVersion: "config.istio.io/v1alpha2"
kind: adapter
metadata:
  name: keyval
spec:
  description: 
  session_based: false
  templates:
  - keyval
  config: CsgCCh5taXhlci90ZXN0L2tleXZhbC9jb25maWcucHJvdG8SBmtleXZhbCJzCgZQYXJhbXMSLwoFdGFibGUYASADKAsyGS5rZXl2YWwuUGFyYW1zLlRhYmxlRW50cnlSBXRhYmxlGjgKClRhYmxlRW50cnkSEAoDa2V5GAEgASgJUgNrZXkSFAoFdmFsdWUYAiABKAlSBXZhbHVlOgI4AUqgAQoGEgQAAAgBCggKAQwSAwAAEgoICgECEgMCCA4KIAoCBAASBAUACAEaFCBBZGFwdGVyIHBhcmFtZXRlcnMKCgoKAwQAARIDBQgOChsKBAQAAgASAwcCIBoOIExvb2t1cCB0YWJsZQoKDQoFBAACAAQSBAcCBRAKDAoFBAACAAYSAwcCFQoMCgUEAAIAARIDBxYbCgwKBQQAAgADEgMHHh9iBnByb3RvMw==
---
`)

func datasetConfigV1alpha2AdapterValidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2AdapterValidYaml, nil
}

func datasetConfigV1alpha2AdapterValidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2AdapterValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-adapter-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2AttributemanifestInvalidYaml = []byte(`apiVersion: "config.istio.io/v1alpha2"
kind: attributemanifest
metadata:
  name: valid-attribute-manifest
spec:
  attributes:
    origin.ip:
      valueType: IP_ADDRESS
    origin.uid:
      valueType: BAD_TYPE
    origin.user:
      valueType: STRING
    request.headers:
      valueType: STRING_MAP
    request.id:
      valueType: STRING
`)

func datasetConfigV1alpha2AttributemanifestInvalidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2AttributemanifestInvalidYaml, nil
}

func datasetConfigV1alpha2AttributemanifestInvalidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2AttributemanifestInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-attributemanifest-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2AttributemanifestValidYaml = []byte(`apiVersion: "config.istio.io/v1alpha2"
kind: attributemanifest
metadata:
  name: valid-attribute-manifest
spec:
  attributes:
    origin.ip:
      valueType: IP_ADDRESS
    origin.uid:
      valueType: STRING
    origin.user:
      valueType: STRING
    request.headers:
      valueType: STRING_MAP
    request.id:
      valueType: STRING
    request.host:
      valueType: STRING
    request.method:
      valueType: STRING
    request.path:
      valueType: STRING
    request.reason:
      valueType: STRING
    request.referer:
      valueType: STRING
    request.scheme:
      valueType: STRING
    request.total_size:
      valueType: INT64
    request.size:
      valueType: INT64
    request.time:
      valueType: TIMESTAMP
    request.useragent:
      valueType: STRING
    response.code:
      valueType: INT64
    response.duration:
      valueType: DURATION
    response.headers:
      valueType: STRING_MAP
    response.total_size:
      valueType: INT64
    response.size:
      valueType: INT64
    response.time:
      valueType: TIMESTAMP
    source.uid:
      valueType: STRING
    source.user:
      valueType: STRING
    destination.uid:
      valueType: STRING
    connection.id:
      valueType: STRING
    connection.received.bytes:
      valueType: INT64
    connection.received.bytes_total:
      valueType: INT64
    connection.sent.bytes:
      valueType: INT64
    connection.sent.bytes_total:
      valueType: INT64
    connection.duration:
      valueType: DURATION
    connection.mtls:
      valueType: BOOL
    context.protocol:
      valueType: STRING
    context.timestamp:
      valueType: TIMESTAMP
    context.time:
      valueType: TIMESTAMP
    api.service:
      valueType: STRING
    api.version:
      valueType: STRING
    api.operation:
      valueType: STRING
    api.protocol:
      valueType: STRING
    request.auth.principal:
      valueType: STRING
    request.auth.audiences:
      valueType: STRING
    request.auth.presenter:
      valueType: STRING
    request.api_key:
      valueType: STRING
    source.ip:
      valueType: IP_ADDRESS
    source.labels:
      valueType: STRING_MAP
    source.name:
      valueType: STRING
    source.namespace:
      valueType: STRING
    source.service:
      valueType: STRING
    source.serviceAccount:
      valueType: STRING
    destination.ip:
      valueType: IP_ADDRESS
    destination.labels:
      valueType: STRING_MAP
    destination.name:
      valueType: STRING
    destination.namespace:
      valueType: STRING
    destination.service:
      valueType: STRING
    destination.serviceAccount:
      valueType: STRING
`)

func datasetConfigV1alpha2AttributemanifestValidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2AttributemanifestValidYaml, nil
}

func datasetConfigV1alpha2AttributemanifestValidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2AttributemanifestValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-attributemanifest-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2HandlerInvalidYaml = []byte(`apiVersion: "config.istio.io/v1alpha2"
kind: handler
metadata:
  name: prometheusdouble
spec:
  compiledAdapter: 0
  params:
    metrics:
    - name: double_request_count # Prometheus metric name
      instance_name: doublerequestcount.metric.istio-system # Mixer instance name (fully-qualified)
      kind: COUNTER
      label_names:
      - reporter
      - source
      - destination
      - message
`)

func datasetConfigV1alpha2HandlerInvalidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2HandlerInvalidYaml, nil
}

func datasetConfigV1alpha2HandlerInvalidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2HandlerInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-handler-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2HandlerValidYaml = []byte(`apiVersion: "config.istio.io/v1alpha2"
kind: handler
metadata:
  name: prometheusdouble
spec:
  compiledAdapter: prometheus
  params:
    metrics:
    - name: double_request_count # Prometheus metric name
      instance_name: doublerequestcount.metric.istio-system # Mixer instance name (fully-qualified)
      kind: COUNTER
      label_names:
      - reporter
      - source
      - destination
      - message
`)

func datasetConfigV1alpha2HandlerValidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2HandlerValidYaml, nil
}

func datasetConfigV1alpha2HandlerValidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2HandlerValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-handler-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2InstanceInvalidYaml = []byte(`apiVersion: "config.istio.io/v1alpha2"
kind: instance
metadata:
  name: source-info
spec:
  template: listentry
  zarams:
    value: source.name | "unknown"
`)

func datasetConfigV1alpha2InstanceInvalidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2InstanceInvalidYaml, nil
}

func datasetConfigV1alpha2InstanceInvalidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2InstanceInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-instance-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2InstanceValidYaml = []byte(`apiVersion: "config.istio.io/v1alpha2"
kind: instance
metadata:
  name: source-info
spec:
  template: listentry
  params:
    value: source.name | "unknown"
`)

func datasetConfigV1alpha2InstanceValidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2InstanceValidYaml, nil
}

func datasetConfigV1alpha2InstanceValidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2InstanceValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-instance-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2RuleInvalidYaml = []byte(`apiVersion: "config.istio.io/v1alpha2"
kind: handler
metadata:
  name: handler-for-invalid-rule
spec:
  compiledAdapter: denier
  params:
    status:
      code: 7
      message: Not allowed
---
apiVersion: "config.istio.io/v1alpha2"
kind: instance
metadata:
  name: instance-for-invalid-rule
spec:
  compiledTemplate: checknothing
---
apiVersion: "config.istio.io/v1alpha2"
kind: rule
metadata:
  name: invalid-rule
spec:
  badField: foo
  match: request.headers["clnt"] == "abc"
  actions:
  - handler: handler-for-invalid-rule
    instances:
    - instance-for-invalid-rule
`)

func datasetConfigV1alpha2RuleInvalidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2RuleInvalidYaml, nil
}

func datasetConfigV1alpha2RuleInvalidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2RuleInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-rule-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2RuleValidYaml = []byte(`apiVersion: "config.istio.io/v1alpha2"
kind: handler
metadata:
  name: handler-for-valid-rule
spec:
  compiledAdapter: denier
  params:
    status:
      code: 7
      message: Not allowed
---
apiVersion: "config.istio.io/v1alpha2"
kind: instance
metadata:
  name: instance-for-valid-rule
spec:
  compiledTemplate: checknothing
---
apiVersion: "config.istio.io/v1alpha2"
kind: rule
metadata:
  name: valid-rule
spec:
  match: request.headers["clnt"] == "abc"
  actions:
  - handler: handler-for-valid-rule
    instances:
    - instance-for-valid-rule
`)

func datasetConfigV1alpha2RuleValidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2RuleValidYaml, nil
}

func datasetConfigV1alpha2RuleValidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2RuleValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-rule-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2TemplateInvalidYaml = []byte(`apiVersion: "config.istio.io/v1alpha2"
kind: template
metadata:
  name: keyval
spec:
  prescriptor: 0
`)

func datasetConfigV1alpha2TemplateInvalidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2TemplateInvalidYaml, nil
}

func datasetConfigV1alpha2TemplateInvalidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2TemplateInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-template-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetConfigV1alpha2TemplateValidYaml = []byte(`apiVersion: "config.istio.io/v1alpha2"
kind: template
metadata:
  name: keyval
spec:
  descriptor: "CvD6AgogZ29vZ2xlL3Byb3RvYnVmL2Rlc2NyaXB0b3IucHJvdG8SD2dvb2dsZS5wcm90b2J1ZiJNChFGaWxlRGVzY3JpcHRvclNldBI4CgRmaWxlGAEgAygLMiQuZ29vZ2xlLnByb3RvYnVmLkZpbGVEZXNjcmlwdG9yUHJvdG9SBGZpbGUi5AQKE0ZpbGVEZXNjcmlwdG9yUHJvdG8SEgoEbmFtZRgBIAEoCVIEbmFtZRIYCgdwYWNrYWdlGAIgASgJUgdwYWNrYWdlEh4KCmRlcGVuZGVuY3kYAyADKAlSCmRlcGVuZGVuY3kSKwoRcHVibGljX2RlcGVuZGVuY3kYCiADKAVSEHB1YmxpY0RlcGVuZGVuY3kSJwoPd2Vha19kZXBlbmRlbmN5GAsgAygFUg53ZWFrRGVwZW5kZW5jeRJDCgxtZXNzYWdlX3R5cGUYBCADKAsyIC5nb29nbGUucHJvdG9idWYuRGVzY3JpcHRvclByb3RvUgttZXNzYWdlVHlwZRJBCgllbnVtX3R5cGUYBSADKAsyJC5nb29nbGUucHJvdG9idWYuRW51bURlc2NyaXB0b3JQcm90b1IIZW51bVR5cGUSQQoHc2VydmljZRgGIAMoCzInLmdvb2dsZS5wcm90b2J1Zi5TZXJ2aWNlRGVzY3JpcHRvclByb3RvUgdzZXJ2aWNlEkMKCWV4dGVuc2lvbhgHIAMoCzIlLmdvb2dsZS5wcm90b2J1Zi5GaWVsZERlc2NyaXB0b3JQcm90b1IJZXh0ZW5zaW9uEjYKB29wdGlvbnMYCCABKAsyHC5nb29nbGUucHJvdG9idWYuRmlsZU9wdGlvbnNSB29wdGlvbnMSSQoQc291cmNlX2NvZGVfaW5mbxgJIAEoCzIfLmdvb2dsZS5wcm90b2J1Zi5Tb3VyY2VDb2RlSW5mb1IOc291cmNlQ29kZUluZm8SFgoGc3ludGF4GAwgASgJUgZzeW50YXgiuQYKD0Rlc2NyaXB0b3JQcm90bxISCgRuYW1lGAEgASgJUgRuYW1lEjsKBWZpZWxkGAIgAygLMiUuZ29vZ2xlLnByb3RvYnVmLkZpZWxkRGVzY3JpcHRvclByb3RvUgVmaWVsZBJDCglleHRlbnNpb24YBiADKAsyJS5nb29nbGUucHJvdG9idWYuRmllbGREZXNjcmlwdG9yUHJvdG9SCWV4dGVuc2lvbhJBCgtuZXN0ZWRfdHlwZRgDIAMoCzIgLmdvb2dsZS5wcm90b2J1Zi5EZXNjcmlwdG9yUHJvdG9SCm5lc3RlZFR5cGUSQQoJZW51bV90eXBlGAQgAygLMiQuZ29vZ2xlLnByb3RvYnVmLkVudW1EZXNjcmlwdG9yUHJvdG9SCGVudW1UeXBlElgKD2V4dGVuc2lvbl9yYW5nZRgFIAMoCzIvLmdvb2dsZS5wcm90b2J1Zi5EZXNjcmlwdG9yUHJvdG8uRXh0ZW5zaW9uUmFuZ2VSDmV4dGVuc2lvblJhbmdlEkQKCm9uZW9mX2RlY2wYCCADKAsyJS5nb29nbGUucHJvdG9idWYuT25lb2ZEZXNjcmlwdG9yUHJvdG9SCW9uZW9mRGVjbBI5CgdvcHRpb25zGAcgASgLMh8uZ29vZ2xlLnByb3RvYnVmLk1lc3NhZ2VPcHRpb25zUgdvcHRpb25zElUKDnJlc2VydmVkX3JhbmdlGAkgAygLMi4uZ29vZ2xlLnByb3RvYnVmLkRlc2NyaXB0b3JQcm90by5SZXNlcnZlZFJhbmdlUg1yZXNlcnZlZFJhbmdlEiMKDXJlc2VydmVkX25hbWUYCiADKAlSDHJlc2VydmVkTmFtZRp6Cg5FeHRlbnNpb25SYW5nZRIUCgVzdGFydBgBIAEoBVIFc3RhcnQSEAoDZW5kGAIgASgFUgNlbmQSQAoHb3B0aW9ucxgDIAEoCzImLmdvb2dsZS5wcm90b2J1Zi5FeHRlbnNpb25SYW5nZU9wdGlvbnNSB29wdGlvbnMaNwoNUmVzZXJ2ZWRSYW5nZRIUCgVzdGFydBgBIAEoBVIFc3RhcnQSEAoDZW5kGAIgASgFUgNlbmQifAoVRXh0ZW5zaW9uUmFuZ2VPcHRpb25zElgKFHVuaW50ZXJwcmV0ZWRfb3B0aW9uGOcHIAMoCzIkLmdvb2dsZS5wcm90b2J1Zi5VbmludGVycHJldGVkT3B0aW9uUhN1bmludGVycHJldGVkT3B0aW9uKgkI6AcQgICAgAIimAYKFEZpZWxkRGVzY3JpcHRvclByb3RvEhIKBG5hbWUYASABKAlSBG5hbWUSFgoGbnVtYmVyGAMgASgFUgZudW1iZXISQQoFbGFiZWwYBCABKA4yKy5nb29nbGUucHJvdG9idWYuRmllbGREZXNjcmlwdG9yUHJvdG8uTGFiZWxSBWxhYmVsEj4KBHR5cGUYBSABKA4yKi5nb29nbGUucHJvdG9idWYuRmllbGREZXNjcmlwdG9yUHJvdG8uVHlwZVIEdHlwZRIbCgl0eXBlX25hbWUYBiABKAlSCHR5cGVOYW1lEhoKCGV4dGVuZGVlGAIgASgJUghleHRlbmRlZRIjCg1kZWZhdWx0X3ZhbHVlGAcgASgJUgxkZWZhdWx0VmFsdWUSHwoLb25lb2ZfaW5kZXgYCSABKAVSCm9uZW9mSW5kZXgSGwoJanNvbl9uYW1lGAogASgJUghqc29uTmFtZRI3CgdvcHRpb25zGAggASgLMh0uZ29vZ2xlLnByb3RvYnVmLkZpZWxkT3B0aW9uc1IHb3B0aW9ucyK2AgoEVHlwZRIPCgtUWVBFX0RPVUJMRRABEg4KClRZUEVfRkxPQVQQAhIOCgpUWVBFX0lOVDY0EAMSDwoLVFlQRV9VSU5UNjQQBBIOCgpUWVBFX0lOVDMyEAUSEAoMVFlQRV9GSVhFRDY0EAYSEAoMVFlQRV9GSVhFRDMyEAcSDQoJVFlQRV9CT09MEAgSDwoLVFlQRV9TVFJJTkcQCRIOCgpUWVBFX0dST1VQEAoSEAoMVFlQRV9NRVNTQUdFEAsSDgoKVFlQRV9CWVRFUxAMEg8KC1RZUEVfVUlOVDMyEA0SDQoJVFlQRV9FTlVNEA4SEQoNVFlQRV9TRklYRUQzMhAPEhEKDVRZUEVfU0ZJWEVENjQQEBIPCgtUWVBFX1NJTlQzMhAREg8KC1RZUEVfU0lOVDY0EBIiQwoFTGFiZWwSEgoOTEFCRUxfT1BUSU9OQUwQARISCg5MQUJFTF9SRVFVSVJFRBACEhIKDkxBQkVMX1JFUEVBVEVEEAMiYwoUT25lb2ZEZXNjcmlwdG9yUHJvdG8SEgoEbmFtZRgBIAEoCVIEbmFtZRI3CgdvcHRpb25zGAIgASgLMh0uZ29vZ2xlLnByb3RvYnVmLk9uZW9mT3B0aW9uc1IHb3B0aW9ucyLjAgoTRW51bURlc2NyaXB0b3JQcm90bxISCgRuYW1lGAEgASgJUgRuYW1lEj8KBXZhbHVlGAIgAygLMikuZ29vZ2xlLnByb3RvYnVmLkVudW1WYWx1ZURlc2NyaXB0b3JQcm90b1IFdmFsdWUSNgoHb3B0aW9ucxgDIAEoCzIcLmdvb2dsZS5wcm90b2J1Zi5FbnVtT3B0aW9uc1IHb3B0aW9ucxJdCg5yZXNlcnZlZF9yYW5nZRgEIAMoCzI2Lmdvb2dsZS5wcm90b2J1Zi5FbnVtRGVzY3JpcHRvclByb3RvLkVudW1SZXNlcnZlZFJhbmdlUg1yZXNlcnZlZFJhbmdlEiMKDXJlc2VydmVkX25hbWUYBSADKAlSDHJlc2VydmVkTmFtZRo7ChFFbnVtUmVzZXJ2ZWRSYW5nZRIUCgVzdGFydBgBIAEoBVIFc3RhcnQSEAoDZW5kGAIgASgFUgNlbmQigwEKGEVudW1WYWx1ZURlc2NyaXB0b3JQcm90bxISCgRuYW1lGAEgASgJUgRuYW1lEhYKBm51bWJlchgCIAEoBVIGbnVtYmVyEjsKB29wdGlvbnMYAyABKAsyIS5nb29nbGUucHJvdG9idWYuRW51bVZhbHVlT3B0aW9uc1IHb3B0aW9ucyKnAQoWU2VydmljZURlc2NyaXB0b3JQcm90bxISCgRuYW1lGAEgASgJUgRuYW1lEj4KBm1ldGhvZBgCIAMoCzImLmdvb2dsZS5wcm90b2J1Zi5NZXRob2REZXNjcmlwdG9yUHJvdG9SBm1ldGhvZBI5CgdvcHRpb25zGAMgASgLMh8uZ29vZ2xlLnByb3RvYnVmLlNlcnZpY2VPcHRpb25zUgdvcHRpb25zIokCChVNZXRob2REZXNjcmlwdG9yUHJvdG8SEgoEbmFtZRgBIAEoCVIEbmFtZRIdCgppbnB1dF90eXBlGAIgASgJUglpbnB1dFR5cGUSHwoLb3V0cHV0X3R5cGUYAyABKAlSCm91dHB1dFR5cGUSOAoHb3B0aW9ucxgEIAEoCzIeLmdvb2dsZS5wcm90b2J1Zi5NZXRob2RPcHRpb25zUgdvcHRpb25zEjAKEGNsaWVudF9zdHJlYW1pbmcYBSABKAg6BWZhbHNlUg9jbGllbnRTdHJlYW1pbmcSMAoQc2VydmVyX3N0cmVhbWluZxgGIAEoCDoFZmFsc2VSD3NlcnZlclN0cmVhbWluZyK5CAoLRmlsZU9wdGlvbnMSIQoMamF2YV9wYWNrYWdlGAEgASgJUgtqYXZhUGFja2FnZRIwChRqYXZhX291dGVyX2NsYXNzbmFtZRgIIAEoCVISamF2YU91dGVyQ2xhc3NuYW1lEjUKE2phdmFfbXVsdGlwbGVfZmlsZXMYCiABKAg6BWZhbHNlUhFqYXZhTXVsdGlwbGVGaWxlcxJECh1qYXZhX2dlbmVyYXRlX2VxdWFsc19hbmRfaGFzaBgUIAEoCEICGAFSGWphdmFHZW5lcmF0ZUVxdWFsc0FuZEhhc2gSOgoWamF2YV9zdHJpbmdfY2hlY2tfdXRmOBgbIAEoCDoFZmFsc2VSE2phdmFTdHJpbmdDaGVja1V0ZjgSUwoMb3B0aW1pemVfZm9yGAkgASgOMikuZ29vZ2xlLnByb3RvYnVmLkZpbGVPcHRpb25zLk9wdGltaXplTW9kZToFU1BFRURSC29wdGltaXplRm9yEh0KCmdvX3BhY2thZ2UYCyABKAlSCWdvUGFja2FnZRI1ChNjY19nZW5lcmljX3NlcnZpY2VzGBAgASgIOgVmYWxzZVIRY2NHZW5lcmljU2VydmljZXMSOQoVamF2YV9nZW5lcmljX3NlcnZpY2VzGBEgASgIOgVmYWxzZVITamF2YUdlbmVyaWNTZXJ2aWNlcxI1ChNweV9nZW5lcmljX3NlcnZpY2VzGBIgASgIOgVmYWxzZVIRcHlHZW5lcmljU2VydmljZXMSNwoUcGhwX2dlbmVyaWNfc2VydmljZXMYKiABKAg6BWZhbHNlUhJwaHBHZW5lcmljU2VydmljZXMSJQoKZGVwcmVjYXRlZBgXIAEoCDoFZmFsc2VSCmRlcHJlY2F0ZWQSLwoQY2NfZW5hYmxlX2FyZW5hcxgfIAEoCDoFZmFsc2VSDmNjRW5hYmxlQXJlbmFzEioKEW9iamNfY2xhc3NfcHJlZml4GCQgASgJUg9vYmpjQ2xhc3NQcmVmaXgSKQoQY3NoYXJwX25hbWVzcGFjZRglIAEoCVIPY3NoYXJwTmFtZXNwYWNlEiEKDHN3aWZ0X3ByZWZpeBgnIAEoCVILc3dpZnRQcmVmaXgSKAoQcGhwX2NsYXNzX3ByZWZpeBgoIAEoCVIOcGhwQ2xhc3NQcmVmaXgSIwoNcGhwX25hbWVzcGFjZRgpIAEoCVIMcGhwTmFtZXNwYWNlElgKFHVuaW50ZXJwcmV0ZWRfb3B0aW9uGOcHIAMoCzIkLmdvb2dsZS5wcm90b2J1Zi5VbmludGVycHJldGVkT3B0aW9uUhN1bmludGVycHJldGVkT3B0aW9uIjoKDE9wdGltaXplTW9kZRIJCgVTUEVFRBABEg0KCUNPREVfU0laRRACEhAKDExJVEVfUlVOVElNRRADKgkI6AcQgICAgAJKBAgmECci0QIKDk1lc3NhZ2VPcHRpb25zEjwKF21lc3NhZ2Vfc2V0X3dpcmVfZm9ybWF0GAEgASgIOgVmYWxzZVIUbWVzc2FnZVNldFdpcmVGb3JtYXQSTAofbm9fc3RhbmRhcmRfZGVzY3JpcHRvcl9hY2Nlc3NvchgCIAEoCDoFZmFsc2VSHG5vU3RhbmRhcmREZXNjcmlwdG9yQWNjZXNzb3ISJQoKZGVwcmVjYXRlZBgDIAEoCDoFZmFsc2VSCmRlcHJlY2F0ZWQSGwoJbWFwX2VudHJ5GAcgASgIUghtYXBFbnRyeRJYChR1bmludGVycHJldGVkX29wdGlvbhjnByADKAsyJC5nb29nbGUucHJvdG9idWYuVW5pbnRlcnByZXRlZE9wdGlvblITdW5pbnRlcnByZXRlZE9wdGlvbioJCOgHEICAgIACSgQICBAJSgQICRAKIuIDCgxGaWVsZE9wdGlvbnMSQQoFY3R5cGUYASABKA4yIy5nb29nbGUucHJvdG9idWYuRmllbGRPcHRpb25zLkNUeXBlOgZTVFJJTkdSBWN0eXBlEhYKBnBhY2tlZBgCIAEoCFIGcGFja2VkEkcKBmpzdHlwZRgGIAEoDjIkLmdvb2dsZS5wcm90b2J1Zi5GaWVsZE9wdGlvbnMuSlNUeXBlOglKU19OT1JNQUxSBmpzdHlwZRIZCgRsYXp5GAUgASgIOgVmYWxzZVIEbGF6eRIlCgpkZXByZWNhdGVkGAMgASgIOgVmYWxzZVIKZGVwcmVjYXRlZBIZCgR3ZWFrGAogASgIOgVmYWxzZVIEd2VhaxJYChR1bmludGVycHJldGVkX29wdGlvbhjnByADKAsyJC5nb29nbGUucHJvdG9idWYuVW5pbnRlcnByZXRlZE9wdGlvblITdW5pbnRlcnByZXRlZE9wdGlvbiIvCgVDVHlwZRIKCgZTVFJJTkcQABIICgRDT1JEEAESEAoMU1RSSU5HX1BJRUNFEAIiNQoGSlNUeXBlEg0KCUpTX05PUk1BTBAAEg0KCUpTX1NUUklORxABEg0KCUpTX05VTUJFUhACKgkI6AcQgICAgAJKBAgEEAUicwoMT25lb2ZPcHRpb25zElgKFHVuaW50ZXJwcmV0ZWRfb3B0aW9uGOcHIAMoCzIkLmdvb2dsZS5wcm90b2J1Zi5VbmludGVycHJldGVkT3B0aW9uUhN1bmludGVycHJldGVkT3B0aW9uKgkI6AcQgICAgAIiwAEKC0VudW1PcHRpb25zEh8KC2FsbG93X2FsaWFzGAIgASgIUgphbGxvd0FsaWFzEiUKCmRlcHJlY2F0ZWQYAyABKAg6BWZhbHNlUgpkZXByZWNhdGVkElgKFHVuaW50ZXJwcmV0ZWRfb3B0aW9uGOcHIAMoCzIkLmdvb2dsZS5wcm90b2J1Zi5VbmludGVycHJldGVkT3B0aW9uUhN1bmludGVycHJldGVkT3B0aW9uKgkI6AcQgICAgAJKBAgFEAYingEKEEVudW1WYWx1ZU9wdGlvbnMSJQoKZGVwcmVjYXRlZBgBIAEoCDoFZmFsc2VSCmRlcHJlY2F0ZWQSWAoUdW5pbnRlcnByZXRlZF9vcHRpb24Y5wcgAygLMiQuZ29vZ2xlLnByb3RvYnVmLlVuaW50ZXJwcmV0ZWRPcHRpb25SE3VuaW50ZXJwcmV0ZWRPcHRpb24qCQjoBxCAgICAAiKcAQoOU2VydmljZU9wdGlvbnMSJQoKZGVwcmVjYXRlZBghIAEoCDoFZmFsc2VSCmRlcHJlY2F0ZWQSWAoUdW5pbnRlcnByZXRlZF9vcHRpb24Y5wcgAygLMiQuZ29vZ2xlLnByb3RvYnVmLlVuaW50ZXJwcmV0ZWRPcHRpb25SE3VuaW50ZXJwcmV0ZWRPcHRpb24qCQjoBxCAgICAAiLgAgoNTWV0aG9kT3B0aW9ucxIlCgpkZXByZWNhdGVkGCEgASgIOgVmYWxzZVIKZGVwcmVjYXRlZBJxChFpZGVtcG90ZW5jeV9sZXZlbBgiIAEoDjIvLmdvb2dsZS5wcm90b2J1Zi5NZXRob2RPcHRpb25zLklkZW1wb3RlbmN5TGV2ZWw6E0lERU1QT1RFTkNZX1VOS05PV05SEGlkZW1wb3RlbmN5TGV2ZWwSWAoUdW5pbnRlcnByZXRlZF9vcHRpb24Y5wcgAygLMiQuZ29vZ2xlLnByb3RvYnVmLlVuaW50ZXJwcmV0ZWRPcHRpb25SE3VuaW50ZXJwcmV0ZWRPcHRpb24iUAoQSWRlbXBvdGVuY3lMZXZlbBIXChNJREVNUE9URU5DWV9VTktOT1dOEAASEwoPTk9fU0lERV9FRkZFQ1RTEAESDgoKSURFTVBPVEVOVBACKgkI6AcQgICAgAIimgMKE1VuaW50ZXJwcmV0ZWRPcHRpb24SQQoEbmFtZRgCIAMoCzItLmdvb2dsZS5wcm90b2J1Zi5VbmludGVycHJldGVkT3B0aW9uLk5hbWVQYXJ0UgRuYW1lEikKEGlkZW50aWZpZXJfdmFsdWUYAyABKAlSD2lkZW50aWZpZXJWYWx1ZRIsChJwb3NpdGl2ZV9pbnRfdmFsdWUYBCABKARSEHBvc2l0aXZlSW50VmFsdWUSLAoSbmVnYXRpdmVfaW50X3ZhbHVlGAUgASgDUhBuZWdhdGl2ZUludFZhbHVlEiEKDGRvdWJsZV92YWx1ZRgGIAEoAVILZG91YmxlVmFsdWUSIQoMc3RyaW5nX3ZhbHVlGAcgASgMUgtzdHJpbmdWYWx1ZRInCg9hZ2dyZWdhdGVfdmFsdWUYCCABKAlSDmFnZ3JlZ2F0ZVZhbHVlGkoKCE5hbWVQYXJ0EhsKCW5hbWVfcGFydBgBIAIoCVIIbmFtZVBhcnQSIQoMaXNfZXh0ZW5zaW9uGAIgAigIUgtpc0V4dGVuc2lvbiKnAgoOU291cmNlQ29kZUluZm8SRAoIbG9jYXRpb24YASADKAsyKC5nb29nbGUucHJvdG9idWYuU291cmNlQ29kZUluZm8uTG9jYXRpb25SCGxvY2F0aW9uGs4BCghMb2NhdGlvbhIWCgRwYXRoGAEgAygFQgIQAVIEcGF0aBIWCgRzcGFuGAIgAygFQgIQAVIEc3BhbhIpChBsZWFkaW5nX2NvbW1lbnRzGAMgASgJUg9sZWFkaW5nQ29tbWVudHMSKwoRdHJhaWxpbmdfY29tbWVudHMYBCABKAlSEHRyYWlsaW5nQ29tbWVudHMSOgoZbGVhZGluZ19kZXRhY2hlZF9jb21tZW50cxgGIAMoCVIXbGVhZGluZ0RldGFjaGVkQ29tbWVudHMi0QEKEUdlbmVyYXRlZENvZGVJbmZvEk0KCmFubm90YXRpb24YASADKAsyLS5nb29nbGUucHJvdG9idWYuR2VuZXJhdGVkQ29kZUluZm8uQW5ub3RhdGlvblIKYW5ub3RhdGlvbhptCgpBbm5vdGF0aW9uEhYKBHBhdGgYASADKAVCAhABUgRwYXRoEh8KC3NvdXJjZV9maWxlGAIgASgJUgpzb3VyY2VGaWxlEhQKBWJlZ2luGAMgASgFUgViZWdpbhIQCgNlbmQYBCABKAVSA2VuZEKPAQoTY29tLmdvb2dsZS5wcm90b2J1ZkIQRGVzY3JpcHRvclByb3Rvc0gBWj5naXRodWIuY29tL2dvbGFuZy9wcm90b2J1Zi9wcm90b2MtZ2VuLWdvL2Rlc2NyaXB0b3I7ZGVzY3JpcHRvcvgBAaICA0dQQqoCGkdvb2dsZS5Qcm90b2J1Zi5SZWZsZWN0aW9uSqrAAgoHEgUnAOcGAQqqDwoBDBIDJwASMsEMIFByb3RvY29sIEJ1ZmZlcnMgLSBHb29nbGUncyBkYXRhIGludGVyY2hhbmdlIGZvcm1hdAogQ29weXJpZ2h0IDIwMDggR29vZ2xlIEluYy4gIEFsbCByaWdodHMgcmVzZXJ2ZWQuCiBodHRwczovL2RldmVsb3BlcnMuZ29vZ2xlLmNvbS9wcm90b2NvbC1idWZmZXJzLwoKIFJlZGlzdHJpYnV0aW9uIGFuZCB1c2UgaW4gc291cmNlIGFuZCBiaW5hcnkgZm9ybXMsIHdpdGggb3Igd2l0aG91dAogbW9kaWZpY2F0aW9uLCBhcmUgcGVybWl0dGVkIHByb3ZpZGVkIHRoYXQgdGhlIGZvbGxvd2luZyBjb25kaXRpb25zIGFyZQogbWV0OgoKICAgICAqIFJlZGlzdHJpYnV0aW9ucyBvZiBzb3VyY2UgY29kZSBtdXN0IHJldGFpbiB0aGUgYWJvdmUgY29weXJpZ2h0CiBub3RpY2UsIHRoaXMgbGlzdCBvZiBjb25kaXRpb25zIGFuZCB0aGUgZm9sbG93aW5nIGRpc2NsYWltZXIuCiAgICAgKiBSZWRpc3RyaWJ1dGlvbnMgaW4gYmluYXJ5IGZvcm0gbXVzdCByZXByb2R1Y2UgdGhlIGFib3ZlCiBjb3B5cmlnaHQgbm90aWNlLCB0aGlzIGxpc3Qgb2YgY29uZGl0aW9ucyBhbmQgdGhlIGZvbGxvd2luZyBkaXNjbGFpbWVyCiBpbiB0aGUgZG9jdW1lbnRhdGlvbiBhbmQvb3Igb3RoZXIgbWF0ZXJpYWxzIHByb3ZpZGVkIHdpdGggdGhlCiBkaXN0cmlidXRpb24uCiAgICAgKiBOZWl0aGVyIHRoZSBuYW1lIG9mIEdvb2dsZSBJbmMuIG5vciB0aGUgbmFtZXMgb2YgaXRzCiBjb250cmlidXRvcnMgbWF5IGJlIHVzZWQgdG8gZW5kb3JzZSBvciBwcm9tb3RlIHByb2R1Y3RzIGRlcml2ZWQgZnJvbQogdGhpcyBzb2Z0d2FyZSB3aXRob3V0IHNwZWNpZmljIHByaW9yIHdyaXR0ZW4gcGVybWlzc2lvbi4KCiBUSElTIFNPRlRXQVJFIElTIFBST1ZJREVEIEJZIFRIRSBDT1BZUklHSFQgSE9MREVSUyBBTkQgQ09OVFJJQlVUT1JTCiAiQVMgSVMiIEFORCBBTlkgRVhQUkVTUyBPUiBJTVBMSUVEIFdBUlJBTlRJRVMsIElOQ0xVRElORywgQlVUIE5PVAogTElNSVRFRCBUTywgVEhFIElNUExJRUQgV0FSUkFOVElFUyBPRiBNRVJDSEFOVEFCSUxJVFkgQU5EIEZJVE5FU1MgRk9SCiBBIFBBUlRJQ1VMQVIgUFVSUE9TRSBBUkUgRElTQ0xBSU1FRC4gSU4gTk8gRVZFTlQgU0hBTEwgVEhFIENPUFlSSUdIVAogT1dORVIgT1IgQ09OVFJJQlVUT1JTIEJFIExJQUJMRSBGT1IgQU5ZIERJUkVDVCwgSU5ESVJFQ1QsIElOQ0lERU5UQUwsCiBTUEVDSUFMLCBFWEVNUExBUlksIE9SIENPTlNFUVVFTlRJQUwgREFNQUdFUyAoSU5DTFVESU5HLCBCVVQgTk9UCiBMSU1JVEVEIFRPLCBQUk9DVVJFTUVOVCBPRiBTVUJTVElUVVRFIEdPT0RTIE9SIFNFUlZJQ0VTOyBMT1NTIE9GIFVTRSwKIERBVEEsIE9SIFBST0ZJVFM7IE9SIEJVU0lORVNTIElOVEVSUlVQVElPTikgSE9XRVZFUiBDQVVTRUQgQU5EIE9OIEFOWQogVEhFT1JZIE9GIExJQUJJTElUWSwgV0hFVEhFUiBJTiBDT05UUkFDVCwgU1RSSUNUIExJQUJJTElUWSwgT1IgVE9SVAogKElOQ0xVRElORyBORUdMSUdF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bm5vdAogc3RvcmUgdGhlIHZhbHVlcyBpbiB0aGVtLiAgSW5zdGVhZCBpdCBzdG9yZXMgdGhlbSBpbiBhIGZpZWxkIGluIHRoZSAqT3B0aW9ucwogbWVzc2FnZSBjYWxsZWQgdW5pbnRlcnByZXRlZF9vcHRpb24uIFRoaXMgZmllbGQgbXVzdCBoYXZlIHRoZSBzYW1lIG5hbWUKIGFjcm9zcyBhbGwgKk9wdGlvbnMgbWVzc2FnZXMuIFdlIHRoZW4gdXNlIHRoaXMgZmllbGQgdG8gcG9wdWxhdGUgdGhlCiBleHRlbnNpb25zIHdoZW4gd2UgYnVpbGQgYSBkZXNjcmlwdG9yLCBhdCB3aGljaCBwb2ludCBhbGwgcHJvdG9zIGhhdmUgYmVlbgogcGFyc2VkIGFuZCBzbyBhbGwgZXh0ZW5zaW9ucyBhcmUga25vd24uCgogRXh0ZW5zaW9uIG51bWJlcnMgZm9yIGN1c3RvbSBvcHRpb25zIG1heSBiZSBjaG9zZW4gYXMgZm9sbG93czoKICogRm9yIG9wdGlvbnMgd2hpY2ggd2lsbCBvbmx5IGJlIHVzZWQgd2l0aGluIGEgc2luZ2xlIGFwcGxpY2F0aW9uIG9yCiAgIG9yZ2FuaXphdGlvbiwgb3IgZm9yIGV4cGVyaW1lbnRhbCBvcHRpb25zLCB1c2UgZmllbGQgbnVtYmVycyA1MDAwMAogICB0aHJvdWdoIDk5OTk5LiAgSXQgaXMgdXAgdG8geW91IHRvIGVuc3VyZSB0aGF0IHlvdSBkbyBub3QgdXNlIHRoZQogICBzYW1lIG51bWJlciBmb3IgbXVsdGlwbGUgb3B0aW9ucy4KICogRm9yIG9wdGlvbnMgd2hpY2ggd2lsbCBiZSBwdWJsaXNoZWQgYW5kIHVzZWQgcHVibGljbHkgYnkgbXVsdGlwbGUKICAgaW5kZXBlbmRlbnQgZW50aXRpZXMsIGUtbWFpbCBwcm90b2J1Zi1nbG9iYWwtZXh0ZW5zaW9uLXJlZ2lzdHJ5QGdvb2dsZS5jb20KICAgdG8gcmVzZXJ2ZSBleHRlbnNpb24gbnVtYmVycy4gU2ltcGx5IHByb3ZpZGUgeW91ciBwcm9qZWN0IG5hbWUgKGUuZy4KICAgT2JqZWN0aXZlLUMgcGx1Z2luKSBhbmQgeW91ciBwcm9qZWN0IHdlYnNpdGUgKGlmIGF2YWlsYWJsZSkgLS0gdGhlcmUncyBubwogICBuZWVkIHRvIGV4cGxhaW4gaG93IHlvdSBpbnRlbmQgdG8gdXNlIHRoZW0uIFVzdWFsbHkgeW91IG9ubHkgbmVlZCBvbmUKICAgZXh0ZW5zaW9uIG51bWJlci4gWW91IGNhbiBkZWNsYXJlIG11bHRpcGxlIG9wdGlvbnMgd2l0aCBvbmx5IG9uZSBleHRlbnNpb24KICAgbnVtYmVyIGJ5IHB1dHRpbmcgdGhlbSBpbiBhIHN1Yi1tZXNzYWdlLiBTZWUgdGhlIEN1c3RvbSBPcHRpb25zIHNlY3Rpb24gb2YKICAgdGhlIGRvY3MgZm9yIGV4YW1wbGVzOgogICBodHRwczovL2RldmVsb3BlcnMuZ29vZ2xlLmNvbS9wcm90b2NvbC1idWZmZXJzL2RvY3MvcHJvdG8jb3B0aW9ucwogICBJZiB0aGlzIHR1cm5zIG91dCB0byBiZSBwb3B1bGFyLCBhIHdlYiBzZXJ2aWNlIHdpbGwgYmUgc2V0IHVwCiAgIHRvIGF1dG9tYXRpY2FsbHkgYXNzaWduIG9wdGlvbiBudW1iZXJzLgoKCwoDBAoBEgS9AggTCvQBCgQECgIAEgTDAgIjGuUBIFNldHMgdGhlIEphdmEgcGFja2FnZSB3aGVyZSBjbGFzc2VzIGdlbmVyYXRlZCBmcm9tIHRoaXMgLnByb3RvIHdpbGwgYmUKIHBsYWNlZC4gIEJ5IGRlZmF1bHQsIHRoZSBwcm90byBwYWNrYWdlIGlzIHVzZWQsIGJ1dCB0aGlzIGlzIG9mdGVuCiBpbmFwcHJvcHJpYXRlIGJlY2F1c2UgcHJvdG8gcGFja2FnZXMgZG8gbm90IG5vcm1hbGx5IHN0YXJ0IHdpdGggYmFja3dhcmRzCiBkb21haW4gbmFtZXMuCgoNCgUECgIABBIEwwICCgoNCgUECgIABRIEwwILEQoNCgUECgIAARIEwwISHgoNCgUECgIAAxIEwwIhIgq/AgoEBAoCARIEywICKxqwAiBJZiBzZXQsIGFsbCB0aGUgY2xhc3NlcyBmcm9tIHRoZSAucHJvdG8gZmlsZSBhcmUgd3JhcHBlZCBpbiBhIHNpbmdsZQogb3V0ZXIgY2xh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CgUAEgSpAw0YCg0KBQQKBQABEgSpAw0RCg0KBQQKBQACEgSpAxUYCgsKAwQKCRIEqwMLDgoMCgQECgkAEgSrAwsNCg0KBQQKCQABEgSrAwsNCg0KBQQKCQACEgSrAwsNCgwKAgQLEgauAwDtAwEKCwoDBAsBEgSuAwgWCtgFCgQECwIAEgTBAwI8GskFIFNldCB0cnVlIHRvIHVzZSB0aGUgb2xkIHByb3RvMSBNZXNzYWdlU2V0IHdpcmUgZm9ybWF0IGZvciBleHRlbnNpb25zLgogVGhpcyBpcyBwcm92aWRlZCBmb3IgYmFja3dhcmRzLWNvbXBhdGliaWxpdHkgd2l0aCB0aGUgTWVzc2FnZVNldCB3aXJlCiBmb3JtYXQuICBZb3Ugc2hvdWxkIG5vdCB1c2UgdGhpcyBmb3IgYW55IG90aGVyIHJlYXNvbjogIEl0J3MgbGVzcwogZWZmaWNpZW50LCBoYXMgZmV3ZXIgZmVhdHVyZXMsIGFuZCBpcyBtb3JlIGNvbXBsaWNhdGVkLgoKIFRoZSBtZXNzYWdlIG11c3QgYmUgZGVmaW5lZCBleGFjdGx5IGFzIGZvbGxvd3M6CiAgIG1lc3NhZ2UgRm9vIHsKICAgICBvcHRpb24gbWVzc2FnZV9zZXRfd2lyZV9mb3JtYXQgPSB0cnVlOwogICAgIGV4dGVuc2lvbnMgNCB0byBtYXg7CiAgIH0KIE5vdGUgdGhhdCB0aGUgbWVzc2FnZSBjYW5ub3QgaGF2ZSBhbnkgZGVmaW5lZCBmaWVsZHM7IE1lc3NhZ2VTZXRzIG9ubHkKIGhhdmUgZXh0ZW5zaW9ucy4KCiBBbGwgZXh0ZW5zaW9ucyBvZiB5b3VyIHR5cGUgbXVzdCBiZSBzaW5ndWxhciBtZXNzYWdlczsgZS5nLiB0aGV5IGNhbm5vdAogYmUgaW50MzJzLCBlbnVtcywgb3IgcmVwZWF0ZWQgbWVzc2FnZXMuCgogQmVjYXVzZSB0aGlzIGlzIGFuIG9wdGlvbiwgdGhlIGFib3ZlIHR3byByZXN0cmljdGlvbnMgYXJlIG5vdCBlbmZvcmNlZCBieQogdGhlIHByb3RvY29sIGNvbXBpbGVyLgoKDQoFBAsCAAQSBMEDAgoKDQoFBAsCAAUSBMEDCw8KDQoFBAsCAAESBMEDECcK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CiBhIG1vcmUgZWZmaWNpZW50IHJlcHJlc2VudGF0aW9uIG9uIHRoZSB3aXJlLiBSYXRoZXIgdGhhbiByZXBlYXRlZGx5CiB3cml0aW5nIHRoZSB0YWcgYW5kIHR5cGUgZm9yIGVhY2ggZWxlbWVudCwgdGhlIGVudGlyZSBhcnJheSBpcyBlbmNvZGVkIGFzCiBhIHNpbmdsZSBsZW5ndGgtZGVsaW1pdGVkIGJsb2IuIEluIHByb3RvMywgb25seSBleHBsaWNpdCBzZXR0aW5nIGl0IHRvCiBmYWxzZSB3aWxsIGF2b2lkIHVzaW5nIHBhY2tlZCBlbmNvZGluZy4KCg0KBQQMAgEEEgSCBAIKCg0KBQQMAgEFEgSCBAsPCg0KBQQMAgEBEgSCBBAWCg0KBQQMAgEDEgSCBBkaCpoFCgQEDAICEgSPBAIzGosFIFRoZSBqc3R5cGUgb3B0aW9uIGRldGVybWluZXMgdGhlIEphdmFTY3JpcHQgdHlwZSB1c2VkIGZvciB2YWx1ZXMgb2YgdGhlCiBmaWVsZC4gIFRoZSBvcHRpb24gaXMgcGVybWl0dGVkIG9ubHkgZm9yIDY0IGJpdCBpbnRlZ3JhbCBhbmQgZml4ZWQgdHlwZXMKIChpbnQ2NCwgdWludDY0LCBzaW50NjQsIGZpeGVkNjQsIHNmaXhlZDY0KS4gIEEgZmllbGQgd2l0aCBqc3R5cGUgSlNfU1RSSU5HCiBpcyByZXByZXNlbnRlZCBhcyBKYXZhU2NyaXB0IHN0cmluZywgd2hpY2ggYXZvaWRzIGxvc3Mgb2YgcHJlY2lzaW9uIHRoYXQKIGNhbiBoYXBwZW4gd2hlbiBhIGxhcmdlIHZhbHVlIGlzIGNvbnZlcnRlZCB0byBhIGZsb2F0aW5nIHBvaW50IEphdmFTY3JpcHQuCiBTcGVjaWZ5aW5nIEpTX05VTUJFUiBmb3IgdGhlIGpzdHlwZSBjYXVzZXMgdGhlIGdlbmVyYXRlZCBKYXZhU2NyaXB0IGNvZGUgdG8KIHVzZSB0aGUgSmF2YVNjcmlwdCAibnVtYmVyIiB0eXBlLiAgVGhlIGJlaGF2aW9yIG9mIHRoZSBkZWZhdWx0IG9wdGlvbgogSlNfTk9STUFMIGlzIGltcGxlbWVudGF0aW9uIGRlcGVuZGVudC4KCiBUaGlzIG9wdGlvbiBpcyBhbiBlbnVtIHRvIHBlcm1pdCBhZGRpdGlvbmFsIHR5cGVzIHRvIGJlIGFkZGVkLCBlLmcuCiBnb29nLm1hdGguSW50ZWdlci4KCg0KBQQMAgIEEgSPBAIKCg0KBQQMAgIGEgSPBAsRCg0KBQQMAgIBEgSPBBIYCg0KBQQMAgIDEgSPBBscCg0KBQQMAgIIEgSPBB0yCg0KBQQMAgIHEgSPBCgxCg4KBAQMBAESBpAEApkEAwoNCgUEDAQBARIEkAQHDQonCgYEDAQBAgASBJIEBBIaFyBVc2UgdGhlIGRlZmF1bHQgdHlwZS4KCg8KBwQMBAECAAESBJIEBA0KDwoHBAwEAQIAAhIEkgQQEQopCgYEDAQBAgESBJUEBBIaGSBVc2UgSmF2YVNjcmlwdCBzdHJpbmdzLgoKDwoHBAwEAQIBARIElQQEDQoPCgcEDAQBAgECEgSVBBARCikKBgQMBAECAhIEmAQEEhoZIFVzZSBKYXZhU2NyaXB0IG51bWJlcnMuCgoPCgcEDAQBAgIBEgSYBAQNCg8KBwQMBAECAgISBJgEEBEK7wwKBAQMAgMSBLcEAika4AwgU2hvdWxkIHRoaXMgZmllbGQgYmUgcGFyc2VkIGxhemlseT8gIExhenkgYXBwbGllcyBvbmx5IHRvIG1lc3NhZ2UtdHlwZQogZmllbGRzLiAgSXQgbWVhbnMgdGhhdCB3aGVuIHRoZSBvdXRlciBtZXNzYWdlIGlzIGluaXRpYWxseSBwYXJzZWQsIHRoZQogaW5uZXIgbWVzc2FnZSdzIGNvbnRlbnRzIHdpbGwgbm90IGJlIHBhcnNlZCBidXQgaW5zdGVhZCBzdG9yZWQgaW4gZW5jb2RlZAogZm9ybS4gIFRoZSBpbm5lciBtZXNzYWdlIHdpbGwgYWN0dWFsbHkgYmUgcGFyc2VkIHdoZW4gaXQgaXMgZmlyc3QgYWNjZXNzZWQuCgogVGhpcyBpcyBvbmx5IGEgaGludC4gIEltcGxlbWVudGF0aW9ucyBhcmUgZnJlZSB0byBjaG9vc2Ugd2hldGhlciB0byB1c2UKIGVhZ2VyIG9yIGxhenkgcGFyc2luZyByZWdhcmRsZXNzIG9mIHRoZSB2YWx1ZSBvZiB0aGlzIG9wdGlvbi4gIEhvd2V2ZXIsCiBzZXR0aW5nIHRoaXMgb3B0aW9uIHRydWUgc3VnZ2VzdHMgdGhhdCB0aGUgcHJvdG9jb2wgYXV0aG9yIGJlbGlldmVzIHRoYXQKIHVzaW5nIGxhenkgcGFyc2luZyBvbiB0aGlzIGZpZWxkIGlzIHdvcnRoIHRoZSBhZGRpdGlvbmFsIGJvb2trZWVwaW5nCiBvdmVyaGVhZCB0eXBpY2FsbHkgbmVlZGVkIHRvIGltcGxlbWVudCBpdC4KCiBUaGlzIG9wdGlvbiBkb2VzIG5vdCBhZmZlY3QgdGhlIHB1YmxpYyBpbnRlcmZhY2Ugb2YgYW55IGdlbmVyYXRlZCBjb2RlOwogYWxsIG1ldGhvZCBzaWduYXR1cmVzIHJlbWFpbiB0aGUgc2FtZS4gIEZ1cnRoZXJtb3JlLCB0aHJlYWQtc2FmZXR5IG9mIHRoZQogaW50ZXJmYWNlIGlzIG5vdCBhZmZlY3RlZCBieSB0aGlzIG9wdGlvbjsgY29uc3QgbWV0aG9kcyByZW1haW4gc2FmZSB0bwogY2FsbCBmcm9tIG11bHRpcGxlIHRocmVhZHMgY29uY3VycmVudGx5LCB3aGlsZSBub24tY29uc3QgbWV0aG9kcyBjb250aW51ZQogdG8gcmVxdWlyZSBleGNsdXNpdmUgYWNjZXNzLgoKCiBOb3RlIHRoYXQgaW1wbGVtZW50YXRpb25zIG1heSBjaG9vc2Ugbm90IHRvIGNoZWNrIHJlcXVpcmVkIGZpZWxkcyB3aXRoaW4KIGEgbGF6eSBzdWItbWVzc2FnZS4gIFRoYXQgaXMsIGNhbGxpbmcgSXNJbml0aWFsaXplZCgpIG9uIHRoZSBvdXRlciBtZXNzYWdlCiBtYXkgcmV0dXJuIHRydWUgZXZlbiBpZiB0aGUgaW5uZXIgbWVzc2FnZSBoYXMgbWlzc2luZyByZXF1aXJlZCBmaWVsZHMuCiBUaGlzIGlzIG5lY2Vzc2FyeSBiZWNhdXNlIG90aGVyd2lzZSB0aGUgaW5uZXIgbWVzc2FnZSB3b3VsZCBoYXZlIHRvIGJlCiBwYXJzZWQgaW4gb3JkZXIgdG8gcGVyZm9ybSB0aGUgY2hlY2ssIGRlZmVhdGluZyB0aGUgcHVycG9zZSBvZiBsYXp5CiBwYXJzaW5nLiAgQW4gaW1wbGVtZW50YXRpb24gd2hpY2ggY2hvb3NlcyBub3QgdG8gY2hlY2sgcmVxdWlyZWQgZmllbGRzCiBtdXN0IGJlIGNvbnNpc3RlbnQgYWJvdXQgaXQuICBUaGF0IGlzLCBmb3IgYW55IHBhcnRpY3VsYXIgc3ViLW1lc3NhZ2UsIHRoZQogaW1wbGVtZW50YXRpb24gbXVzdCBlaXRoZXIgKmFsd2F5cyogY2hlY2sgaXRzIHJlcXVpcmVkIGZpZWxkcywgb3IgKm5ldmVyKgogY2hlY2sgaXRzIHJlcXVpcmVkIGZpZWxkcywgcmVnYXJkbGVzcyBvZiB3aGV0aGVyIG9yIG5vdCB0aGUgbWVzc2FnZSBoYXMKIGJlZW4gcGFyc2VkLgoKDQoFBAwCAwQSBLcEAgoKDQoFBAwCAwUSBLcECw8KDQoFBAwCAwESBLcEEBQKDQoFBAwCAwMSBLcEFxgKDQoFBAwCAwgSBLcEGSgKDQoFBAwCAwcSBLcEIicK6AEKBAQMAgQSBL0EAi8a2QEgSXMgdGhpcyBmaWVsZCBkZXByZWNhdGVkPwogRGVwZW5kaW5nIG9uIHRoZSB0YXJnZXQgcGxhdGZvcm0sIHRoaXMgY2FuIGVtaXQgRGVwcmVjYXRlZCBhbm5vdGF0aW9ucwogZm9yIGFjY2Vzc29ycywgb3IgaXQgd2lsbCBiZSBjb21wbGV0ZWx5IGlnbm9yZWQ7IGluIHRoZSB2ZXJ5IGxlYXN0LCB0aGlzCiBpcyBhIGZvcm1hbGl6YXRpb24gZm9yIGRlcHJlY2F0aW5nIGZpZWxkcy4KCg0KBQQMAgQEEgS9BAIKCg0KBQQMAgQFEgS9BAsPCg0KBQQMAgQBEgS9BBAaCg0KBQQMAgQDEgS9BB0eCg0KBQQMAgQIEgS9BB8uCg0KBQQMAgQHEgS9BCgtCj8KBAQMAgUSBMAEAioaMSBGb3IgR29vZ2xlLWludGVybmFsIG1pZ3JhdGlvbiBvbmx5LiBEbyBub3QgdXNlLgoKDQoFBAwCBQQSBMAEAgoKDQoFBAwCBQUSBMAECw8KDQoFBAwCBQESBMAEEBQKDQoFBAwCBQMSBMAEFxkKDQoFBAwCBQgSBMAEGikKDQoFBAwCBQcSBMAEIygKTwoEBAwCBhIExAQCOhpBIFRoZSBwYXJzZXIgc3RvcmVzIG9wdGlvbnMgaXQgZG9lc24ndCByZWNvZ25pemUgaGVyZS4gU2VlIGFib3ZlLgoKDQoFBAwCBgQSBMQEAgoKDQoFBAwCBgYSBMQECx4KDQoFBAwCBgESBMQEHzMKDQoFBAwCBgMSBMQENjkKWgoDBAwFEgTHBAIZGk0gQ2xpZW50cyBjYW4gZGVmaW5lIGN1c3RvbSBvcHRpb25zIGluIGV4dGVuc2lvbnMgb2YgdGhpcyBtZXNzYWdlLiBTZWUgYWJvdmUuCgoMCgQEDAUAEgTHBA0YCg0KBQQMBQABEgTHBA0RCg0KBQQMBQACEgTHBBUYChwKAwQMCRIEyQQLDSIPIHJlbW92ZWQganR5cGUKCgwKBAQMCQASBMkECwwKDQoFBAwJAAESBMkECwwKDQoFBAwJAAISBMkECwwKDAoCBA0SBswEANIEAQoLCgMEDQESBMwECBQKTwoEBA0CABIEzgQCOhpBIFRoZSBwYXJzZXIgc3RvcmVzIG9wdGlvbnMgaXQgZG9lc24ndCByZWNvZ25pemUgaGVyZS4gU2VlIGFib3ZlLgoKDQoFBA0CAAQSBM4EAgoKDQoFBA0CAAYSBM4ECx4KDQoFBA0CAAESBM4EHzMKDQoFBA0CAAMSBM4ENjkKWgoDBA0FEgTRBAIZGk0gQ2xpZW50cyBjYW4gZGVmaW5lIGN1c3RvbSBvcHRpb25zIGluIGV4dGVuc2lvbnMgb2YgdGhpcyBtZXNzYWdlLiBTZWUgYWJvdmUuCgoMCgQEDQUAEgTRBA0YCg0KBQQNBQABEgTRBA0RCg0KBQQNBQACEgTRBBUYCgwKAgQOEgbUBADnBAEKCwoDBA4BEgTUBAgTCmAKBAQOAgASBNgEAiAaUiBTZXQgdGhpcyBvcHRpb24gdG8gdHJ1ZSB0byBhbGxvdyBtYXBwaW5nIGRpZmZlcmVudCB0YWcgbmFtZXMgdG8gdGhlIHNhbWUKIHZhbHVlLgoKDQoFBA4CAAQSBNgEAgoKDQoFBA4CAAUSBNgECw8KDQoFBA4CAAESBNgEEBsKDQoFBA4CAAMSBNgEHh8K5QEKBAQOAgESBN4EAi8a1gEgSXMgdGhpcyBlbnVtIGRlcHJlY2F0ZWQ/CiBEZXBlbmRpbmcgb24gdGhlIHRhcmdldCBwbGF0Zm9ybSwgdGhpcyBjYW4gZW1pdCBEZXByZWNhdGVkIGFubm90YXRpb25zCiBmb3IgdGhlIGVudW0sIG9yIGl0IHdpbGwgYmUgY29tcGxldGVseSBpZ25vcmVkOyBpbiB0aGUgdmVyeSBsZWFzdCwgdGhpcwogaXMgYSBmb3JtYWxpemF0aW9uIGZvciBkZXByZWNhdGluZyBlbnVtcy4KCg0KBQQOAgEEEgTeBAIKCg0KBQQOAgEFEgTeBAsPCg0KBQQOAgEBEgTeBBAaCg0KBQQOAgEDEgTeBB0eCg0KBQQOAgEIEgTeBB8uCg0KBQQOAgEHEgTeBCgtCh8KAwQOCRIE4AQLDSISIGphdmFuYW5vX2FzX2xpdGUKCgwKBAQOCQASBOAECwwKDQoFBA4JAAESBOAECwwKDQoFBA4JAAISBOAECwwKTwoEBA4CAhIE4wQCOhpBIFRoZSBwYXJzZXIgc3RvcmVzIG9wdGlvbnMgaXQgZG9lc24ndCByZWNvZ25pemUgaGVyZS4gU2VlIGFib3ZlLgoKDQoFBA4CAgQSBOMEAgoKDQoFBA4CAgYSBOMECx4KDQoFBA4CAgESBOMEHzMKDQoFBA4CAgMSBOMENjkKWgoDBA4FEgTmBAIZGk0gQ2xpZW50cyBjYW4gZGVmaW5lIGN1c3RvbSBvcHRpb25zIGluIGV4dGVuc2lvbnMgb2YgdGhpcyBtZXNzYWdlLiBTZWUgYWJvdmUuCgoMCgQEDgUAEgTmBA0YCg0KBQQOBQABEgTmBA0RCg0KBQQOBQACEgTmBBUYCgwKAgQPEgbpBAD1BAEKCwoDBA8BEgTpBAgYCvcBCgQEDwIAEgTuBAIvGugBIElzIHRoaXMgZW51bSB2YWx1ZSBkZXByZWNhdGVkPwogRGVwZW5kaW5nIG9uIHRoZSB0YXJnZXQgcGxhdGZvcm0sIHRoaXMgY2FuIGVtaXQgRGVwcmVjYXRlZCBhbm5vdGF0aW9ucwogZm9yIHRoZSBlbnVtIHZhbHVlLCBvciBpdCB3aWxsIGJlIGNvbXBsZXRlbHkgaWdub3JlZDsgaW4gdGhlIHZlcnkgbGVhc3QsCiB0aGlzIGlzIGEgZm9ybWFsaXphdGlvbiBmb3IgZGVwcmVjYXRpbmcgZW51bSB2YWx1ZXMuCgoNCgUEDwIABBIE7gQCCgoNCgUEDwIABRIE7gQLDwoNCgUEDwIAARIE7gQQGgoNCgUEDwIAAxIE7gQdHgoNCgUEDwIACBIE7gQfLgoNCgUEDwIABxIE7gQoLQpPCgQEDwIBEgTxBAI6GkEgVGhlIHBhcnNlciBzdG9yZXMgb3B0aW9ucyBpdCBkb2Vzbid0IHJlY29nbml6ZSBoZXJlLiBTZWUgYWJvdmUuCgoNCgUEDwIBBBIE8QQCCgoNCgUEDwIBBhIE8QQLHgoNCgUEDwIBARIE8QQfMwoNCgUEDwIBAxIE8QQ2OQpaCgMEDwUSBPQEAhkaTSBDbGllbnRzIGNhbiBkZWZpbmUgY3VzdG9tIG9wdGlvbnMgaW4gZXh0ZW5zaW9ucyBvZiB0aGlzIG1lc3NhZ2UuIFNlZSBhYm92ZS4KCgwKBAQPBQASBPQEDRgKDQoFBA8FAAESBPQEDREKDQoFBA8FAAISBPQEFRgKDAoCBBASBvcEAIkFAQoLCgMEEAESBPcECBYK2QMKBAQQAgASBIIFAjAa3wEgSXMgdGhpcyBzZXJ2aWNlIGRlcHJlY2F0ZWQ/CiBEZXBlbmRpbmcgb24gdGhlIHRhcmdldCBwbGF0Zm9ybSwgdGhpcyBjYW4gZW1pdCBEZXByZWNhdGVkIGFubm90YXRpb25zCiBmb3IgdGhlIHNlcnZpY2UsIG9yIGl0IHdpbGwgYmUgY29tcGxldGVseSBpZ25vcmVkOyBpbiB0aGUgdmVyeSBsZWFzdCwKIHRoaXMgaXMgYSBmb3JtYWxpemF0aW9uIGZvciBkZXByZWNhdGluZyBzZXJ2aWNlcy4KMugBIE5vdGU6ICBGaWVsZCBudW1iZXJzIDEgdGhyb3VnaCAzMiBhcmUgcmVzZXJ2ZWQgZm9yIEdvb2dsZSdzIGludGVybmFsIFJQQwogICBmcmFtZXdvcmsuICBXZSBhcG9sb2dpemUgZm9yIGhvYXJkaW5nIHRoZXNlIG51bWJlcnMgdG8gb3Vyc2VsdmVzLCBidXQKICAgd2Ugd2VyZSBhbHJlYWR5IHVzaW5nIHRoZW0gbG9uZyBiZWZvcmUgd2UgZGVjaWRlZCB0byByZWxlYXNlIFByb3RvY29sCiAgIEJ1ZmZlcnMuCgoNCgUEEAIABBIEggUCCgoNCgUEEAIABRIEggULDwoNCgUEEAIAARIEggUQGgoNCgUEEAIAAxIEggUdHwoNCgUEEAIACBIEggUgLwoNCgUEEAIABxIEggUpLgpPCgQEEAIBEgSFBQI6GkEgVGhlIHBhcnNlciBzdG9yZXMgb3B0aW9ucyBpdCBkb2Vzbid0IHJlY29nbml6ZSBoZXJlLiBTZWUgYWJvdmUuCgoNCgUEEAIBBBIEhQUCCgoNCgUEEAIBBhIEhQULHgoNCgUEEAIBARIEhQUfMwoNCgUEEAIBAxIEhQU2OQpaCgMEEAUSBIgFAhkaTSBDbGllbnRzIGNhbiBkZWZpbmUgY3VzdG9tIG9wdGlvbnMgaW4gZXh0ZW5zaW9ucyBvZiB0aGlzIG1lc3NhZ2UuIFNlZSBhYm92ZS4KCgwKBAQQBQASBIgFDRgKDQoFBBAFAAESBIgFDREKDQoFBBAFAAISBIgFFRgKDAoCBBESBosFAKgFAQoLCgMEEQESBIsFCBUK1gMKBAQRAgASBJYFAjAa3AEgSXMgdGhpcyBtZXRob2QgZGVwcmVjYXRlZD8KIERlcGVuZGluZyBvbiB0aGUgdGFyZ2V0IHBsYXRmb3JtLCB0aGlzIGNhbiBlbWl0IERlcHJlY2F0ZWQgYW5ub3RhdGlvbnMKIGZvciB0aGUgbWV0aG9kLCBvciBpdCB3aWxsIGJlIGNvbXBsZXRlbHkgaWdub3JlZDsgaW4gdGhlIHZlcnkgbGVhc3QsCiB0aGlzIGlzIGEgZm9ybWFsaXphdGlvbiBmb3IgZGVwcmVjYXRpbmcgbWV0aG9kcy4KMugBIE5vdGU6ICBGaWVsZCBudW1iZXJzIDEgdGhyb3VnaCAzMiBhcmUgcmVzZXJ2ZWQgZm9yIEdvb2dsZSdzIGludGVybmFsIFJQQwogICBmcmFtZXdvcmsuICBXZSBhcG9sb2dpemUgZm9yIGhvYXJkaW5nIHRoZXNlIG51bWJlcnMgdG8gb3Vyc2VsdmVzLCBidXQKICAgd2Ugd2VyZSBhbHJlYWR5IHVzaW5nIHRoZW0gbG9uZyBiZWZvcmUgd2UgZGVjaWRlZCB0byByZWxlYXNlIFByb3RvY29sCiAgIEJ1ZmZlcnMuCgoNCgUEEQIABBIElgUCCgoNCgUEEQIABRIElgULDwoNCgUEEQIAARIElgUQGgoNCgUEEQIAAxIElgUdHwoNCgUEEQIACBIElgUgLwoNCgUEEQIABxIElgUpLgrwAQoEBBEEABIGmwUCnwUDGt8BIElzIHRoaXMgbWV0aG9kIHNpZGUtZWZmZWN0LWZyZWUgKG9yIHNhZmUgaW4gSFRUUCBwYXJsYW5jZSksIG9yIGlkZW1wb3RlbnQsCiBvciBuZWl0aGVyPyBIVFRQIGJhc2VkIFJQQyBpbXBsZW1lbnRhdGlvbiBtYXkgY2hvb3NlIEdFVCB2ZXJiIGZvciBzYWZlCiBtZXRob2RzLCBhbmQgUFVUIHZlcmIgZm9yIGlkZW1wb3RlbnQgbWV0aG9kcyBpbnN0ZWFkIG9mIHRoZSBkZWZhdWx0IFBPU1QuCgoNCgUEEQQAARIEmwUHFwoOCgYEEQQAAgASBJwFBBwKDwoHBBEEAAIAARIEnAUEFwoPCgcEEQQAAgACEgScBRobCiQKBgQRBAACARIEnQUEHCIUIGltcGxpZXMgaWRlbXBvdGVudAoKDwoHBBEEAAIBARIEnQUEEwoPCgcEEQQAAgECEgSdBRobCjcKBgQRBAACAhIEngUEHCInIGlkZW1wb3RlbnQsIGJ1dCBtYXkgaGF2ZSBzaWRlIGVmZmVjdHMKCg8KBwQRBAACAgESBJ4FBA4KDwoHBBEEAAICAhIEngUaGwoOCgQEEQIBEgagBQKhBScKDQoFBBECAQQSBKAFAgoKDQoFBBECAQYSBKAFCxsKDQoFBBECAQESBKAFHC0KDQoFBBECAQMSBKEFBggKDQoFBBECAQgSBKEFCSYKDQoFBBECAQcSBKEFEiUKTwoEBBECAhIEpAUCOhpBIFRoZSBwYXJzZXIgc3RvcmVzIG9wdGlvbnMgaXQgZG9lc24ndCByZWNvZ25pemUgaGVyZS4gU2VlIGFib3ZlLgoKDQoFBBECAgQSBKQFAgoKDQoFBBECAgYSBKQFCx4KDQoFBBECAgESBKQFHzMKDQoFBBECAgMSBKQFNjkKWgoDBBEFEgSnBQIZGk0gQ2xpZW50cyBjYW4gZGVmaW5lIGN1c3RvbSBvcHRpb25zIGluIGV4dGVuc2lvbnMgb2YgdGhpcyBtZXNzYWdlLiBTZWUgYWJvdmUuCgoMCgQEEQUAEgSnBQ0YCg0KBQQRBQABEgSnBQ0RCg0KBQQRBQACEgSnBRUYCosDCgIEEhIGsQUAxQUBGvwCIEEgbWVzc2FnZSByZXByZXNlbnRpbmcgYSBvcHRpb24gdGhlIHBhcnNlciBkb2VzIG5vdCByZWNvZ25pemUuIFRoaXMgb25seQogYXBwZWFycyBpbiBvcHRpb25zIHByb3RvcyBjcmVhdGVkIGJ5IHRoZSBjb21waWxlcjo6UGFyc2VyIGNsYXNzLgogRGVzY3JpcHRvclBvb2wgcmVzb2x2ZXMgdGhlc2Ugd2hlbiBidWlsZGluZyBEZXNjcmlwdG9yIG9iamVjdHMuIFRoZXJlZm9yZSwKIG9wdGlvbnMgcHJvdG9zIGluIGRlc2NyaXB0b3Igb2JqZWN0cyAoZS5nLiByZXR1cm5lZCBieSBEZXNjcmlwdG9yOjpvcHRpb25zKCksCiBvciBwcm9kdWNlZCBieSBEZXNjcmlwdG9yOjpDb3B5VG8oKSkgd2lsbCBuZXZlciBoYXZlIFVuaW50ZXJwcmV0ZWRPcHRpb25zCiBpbiB0aGVtLgoKCwoDBBIBEgSxBQgbCssCCgQEEgMAEga3BQK6BQMaugIgVGhlIG5hbWUgb2YgdGhlIHVuaW50ZXJwcmV0ZWQgb3B0aW9uLiAgRWFjaCBzdHJpbmcgcmVwcmVzZW50cyBhIHNlZ21lbnQgaW4KIGEgZG90LXNlcGFyYXRlZCBuYW1lLiAgaXNfZXh0ZW5zaW9uIGlzIHRydWUgaWZmIGEgc2VnbWVudCByZXByZXNlbnRzIGFuCiBleHRlbnNpb24gKGRlbm90ZWQgd2l0aCBwYXJlbnRoZXNlcyBpbiBvcHRpb25zIHNwZWNzIGluIC5wcm90byBmaWxlcykuCiBFLmcuLHsgWyJmb28iLCBmYWxzZV0sIFsiYmFyLmJheiIsIHRydWVdLCBbInF1eCIsIGZhbHNlXSB9IHJlcHJlc2VudHMKICJmb28uKGJhci5iYXopLnF1eCIuCgoNCgUEEgMAARIEtwUKEgoOCgYEEgMAAgASBLgFBCIKDwoHBBIDAAIABBIEuAUEDAoPCgcEEgMAAgAFEgS4BQ0TCg8KBwQSAwACAAESBLgFFB0KDwoHBBIDAAIAAxIEuAUgIQoOCgYEEgMAAgESBLkFBCMKDwoHBBIDAAIBBBIEuQUEDAoPCgcEEgMAAgEFEgS5BQ0RCg8KBwQSAwACAQESBLkFEh4KDwoHBBIDAAIBAxIEuQUhIgoMCgQEEgIAEgS7BQIdCg0KBQQSAgAEEgS7BQIKCg0KBQQSAgAGEgS7BQsTCg0KBQQSAgABEgS7BRQYCg0KBQQSAgADEgS7BRscCpwBCgQEEgIBEgS/BQInGo0BIFRoZSB2YWx1ZSBvZiB0aGUgdW5pbnRlcnByZXRlZCBvcHRpb24sIGluIHdoYXRldmVyIHR5cGUgdGhlIHRva2VuaXplcgogaWRlbnRpZmllZCBpdCBhcyBkdXJpbmcgcGFyc2luZy4gRXhhY3RseSBvbmUgb2YgdGhlc2Ugc2hvdWxkIGJlIHNldC4KCg0KBQQSAgEEEgS/BQIKCg0KBQQSAgEFEgS/BQsRCg0KBQQSAgEBEgS/BRIiCg0KBQQSAgEDEgS/BSUmCgwKBAQSAgISBMAFAikKDQoFBBICAgQSBMAFAgoKDQoFBBICAgUSBMAFCxEKDQoFBBICAgESBMAFEiQKDQoFBBICAgMSBMAFJygKDAoEBBICAxIEwQUCKAoNCgUEEgIDBBIEwQUCCgoNCgUEEgIDBRIEwQULEAoNCgUEEgIDARIEwQURIwoNCgUEEgIDAxIEwQUmJwoMCgQEEgIEEgTCBQIjCg0KBQQSAgQEEgTCBQIKCg0KBQQSAgQFEgTCBQsRCg0KBQQSAgQBEgTCBRIeCg0KBQQSAgQDEgTCBSEiCgwKBAQSAgUSBMMFAiIKDQoFBBICBQQSBMMFAgoKDQoFBBICBQUSBMMFCxAKDQoFBBICBQESBMMFER0KDQoFBBICBQMSBMMFICEKDAoEBBICBhIExAUCJgoNCgUEEgIGBBIExAUCCgoNCgUEEgIGBRIExAULEQoNCgUEEgIGARIExAUSIQoNCgUEEgIGAxIExAUkJQraAQoCBBMSBswFAM0GARpqIEVuY2Fwc3VsYXRlcyBpbmZvcm1hdGlvbiBhYm91dCB0aGUgb3JpZ2luYWwgc291cmNlIGZpbGUgZnJvbSB3aGljaCBhCiBGaWxlRGVzY3JpcHRvclByb3RvIHdhcyBnZW5lcmF0ZWQuCjJgID09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT09PT0KIE9wdGlvbmFsIHNvdXJjZSBjb2RlIGluZm8KCgsKAwQTARIEzAUIFgqCEQoEBBMCABIE+AUCIRrzECBBIExvY2F0aW9uIGlkZW50aWZpZXMgYSBwaWVjZSBvZiBzb3VyY2UgY29kZSBpbiBhIC5wcm90byBmaWxlIHdoaWNoCiBjb3JyZXNwb25kcyB0byBhIHBhcnRpY3VsYXIgZGVmaW5pdGlvbi4gIFRoaXMgaW5mb3JtYXRpb24gaXMgaW50ZW5kZWQKIHRvIGJlIHVzZWZ1bCB0byBJREVzLCBjb2RlIGluZGV4ZXJzLCBkb2N1bWVudGF0aW9uIGdlbmVyYXRvcnMsIGFuZCBzaW1pbGFyCiB0b29scy4KCiBGb3IgZXhhbXBsZSwgc2F5IHdlIGhhdmUgYSBmaWxlIGxpa2U6CiAgIG1lc3NhZ2UgRm9vIHsKICAgICBvcHRpb25hbCBzdHJpbmcgZm9vID0gMTsKICAgfQogTGV0J3MgbG9vayBhdCBqdXN0IHRoZSBmaWVsZCBkZWZpbml0aW9uOgogICBvcHRpb25hbCBzdHJpbmcgZm9vID0gMTsKICAgXiAgICAgICBeXiAgICAgXl4gIF4gIF5eXgogICBhICAgICAgIGJjICAgICBkZSAgZiAgZ2hpCiBXZSBoYXZlIHRoZSBmb2xsb3dpbmcgbG9jYXRpb25zOgogICBzcGFuICAgcGF0aCAgICAgICAgICAgICAgIHJlcHJlc2VudHMKICAgW2EsaSkgIFsgNCwgMCwgMiwgMCBdICAgICBUaGUgd2hvbGUgZmllbGQgZGVmaW5pdGlvbi4KICAgW2EsYikgIFsgNCwgMCwgMiwgMCwgNCBdICBUaGUgbGFiZWwgKG9wdGlvbmFsKS4KICAgW2MsZCkgIFsgNCwgMCwgMiwgMCwgNSBdICBUaGUgdHlwZSAoc3RyaW5nKS4KICAgW2UsZikgIFsgNCwgMCwgMiwgMCwgMSBdICBUaGUgbmFtZSAoZm9vKS4KICAgW2csaCkgIFsgNCwgMCwgMiwgMCwgMyBdICBUaGUgbnVtYmVyICgxKS4KCiBOb3RlczoKIC0gQSBsb2NhdGlvbiBtYXkgcmVmZXIgdG8gYSByZXBlYXRlZCBmaWVsZCBpdHNlbGYgKGkuZS4gbm90IHRvIGFueQogICBwYXJ0aWN1bGFyIGluZGV4IHdpdGhpbiBpdCkuICBUaGlzIGlzIHVzZWQgd2hlbmV2ZXIgYSBzZXQgb2YgZWxlbWVudHMgYXJlCiAgIGxvZ2ljYWxseSBlbmNsb3NlZCBpbiBhIHNpbmdsZSBjb2RlIHNlZ21lbnQuICBGb3IgZXhhbXBsZSwgYW4gZW50aXJlCiAgIGV4dGVuZCBibG9jayAocG9zc2libHkgY29udGFpbmluZyBtdWx0aXBsZSBleHRlbnNpb24gZGVmaW5pdGlvbnMpIHdpbGwKICAgaGF2ZSBhbiBvdXRlciBsb2NhdGlvbiB3aG9zZSBwYXRoIHJlZmVycyB0byB0aGUgImV4dGVuc2lvbnMiIHJlcGVhdGVkCiAgIGZpZWxkIHdpdGhvdXQgYW4gaW5kZXguCiAtIE11bHRpcGxlIGxvY2F0aW9ucyBtYXkgaGF2ZSB0aGUgc2FtZSBwYXRoLiAgVGhpcyBoYXBwZW5zIHdoZW4gYSBzaW5nbGUKICAgbG9naWNhbCBkZWNsYXJhdGlvbiBpcyBzcHJlYWQgb3V0IGFjcm9zcyBtdWx0aXBsZSBwbGFjZXMuICBUaGUgbW9zdAogICBvYnZpb3VzIGV4YW1wbGUgaXMgdGhlICJleHRlbmQiIGJsb2NrIGFnYWluIC0tIHRoZXJlIG1heSBiZSBtdWx0aXBsZQogICBleHRlbmQgYmxvY2tzIGluIHRoZSBzYW1lIHNjb3BlLCBlYWNoIG9mIHdoaWNoIHdpbGwgaGF2ZSB0aGUgc2FtZSBwYXRoLgogLSBBIGxvY2F0aW9uJ3Mgc3BhbiBpcyBub3QgYWx3YXlzIGEgc3Vic2V0IG9mIGl0cyBwYXJlbnQncyBzcGFuLiAgRm9yCiAgIGV4YW1wbGUsIHRoZSAiZXh0ZW5kZWUiIG9mIGFuIGV4dGVuc2lvbiBkZWNsYXJhdGlvbiBhcHBlYXJzIGF0IHRoZQogICBiZWdpbm5pbmcgb2YgdGhlICJleHRlbmQiIGJsb2NrIGFuZCBpcyBzaGFyZWQgYnkgYWxsIGV4dGVuc2lvbnMgd2l0aGluCiAgIHRoZSBibG9jay4KIC0gSnVzdCBiZWNhdXNlIGEgbG9jYXRpb24ncyBzcGFuIGlzIGEgc3Vic2V0IG9mIHNvbWUgb3RoZXIgbG9jYXRpb24ncyBzcGFuCiAgIGRvZXMgbm90IG1lYW4gdGhhdCBpdCBpcyBhIGRlc2NlbmRlbnQuICBGb3IgZXhhbXBsZSwgYSAiZ3JvdXAiIGRlZmluZXMKICAgYm90aCBhIHR5cGUgYW5kIGEgZmllbGQgaW4gYSBzaW5nbGUgZGVjbGFyYXRpb24uICBUaHVzLCB0aGUgbG9jYXRpb25zCiAgIGNvcnJlc3BvbmRpbmcgdG8gdGhlIHR5cGUgYW5kIGZpZWxkIGFuZCB0aGVpciBjb21wb25lbnRzIHdpbGwgb3ZlcmxhcC4KIC0gQ29kZSB3aGljaCB0cmllcyB0byBpbnRlcnByZXQgbG9jYXRpb25zIHNob3VsZCBwcm9iYWJseSBiZSBkZXNpZ25lZCB0bwogICBpZ25vcmUgdGhvc2UgdGhhdCBpdCBkb2Vzbid0IHVuZGVyc3RhbmQsIGFzIG1vcmUgdHlwZXMgb2YgbG9jYXRpb25zIGNvdWxkCiAgIGJlIHJlY29yZGVkIGluIHRoZSBmdXR1cmUuCgoNCgUEEwIABBIE+AUCCgoNCgUEEwIABhIE+AULEwoNCgUEEwIAARIE+AUUHAoNCgUEEwIAAxIE+AUfIAoOCgQEEwMAEgb5BQLMBgMKDQoFBBMDAAESBPkFChIKgwcKBgQTAwACABIEkQYEKhryBiBJZGVudGlmaWVzIHdoaWNoIHBhcnQgb2YgdGhlIEZpbGVEZXNjcmlwdG9yUHJvdG8gd2FzIGRlZmluZWQgYXQgdGhpcwogbG9jYXRpb24uCgogRWFjaCBlbGVtZW50IGlzIGEgZmllbGQgbnVtYmVyIG9yIGFuIGluZGV4LiAgVGhleSBmb3JtIGEgcGF0aCBmcm9tCiB0aGUgcm9vdCBGaWxlRGVzY3JpcHRvclByb3RvIHRvIHRoZSBwbGFjZSB3aGVyZSB0aGUgZGVmaW5pdGlvbi4gIEZvcgogZXhhbXBsZSwgdGhpcyBwYXRoOgogICBbIDQsIDMsIDIsIDcsIDEgXQogcmVmZXJzIHRvOgogICBmaWxlLm1lc3NhZ2VfdHlwZSgzKSAgLy8gNCwgMwogICAgICAgLmZpZWxkKDcpICAgICAgICAgLy8gMiwgNwogICAgICAgLm5hbWUoKSAgICAgICAgICAgLy8gMQogVGhpcyBpcyBiZWNhdXNlIEZpbGVEZXNjcmlwdG9yUHJvdG8ubWVzc2FnZV90eXBlIGhhcyBmaWVsZCBudW1iZXIgNDoKICAgcmVwZWF0ZWQgRGVzY3JpcHRvclByb3RvIG1lc3NhZ2VfdHlwZSA9IDQ7CiBhbmQgRGVzY3JpcHRvclByb3RvLmZpZWxkIGhhcyBmaWVsZCBudW1iZXIgMjoKICAgcmVwZWF0ZWQgRmllbGREZXNjcmlwdG9yUHJvdG8gZmllbGQgPSAyOwogYW5kIEZpZWxkRGVzY3JpcHRvclByb3RvLm5hbWUgaGFzIGZpZWxkIG51bWJlciAxOgogICBvcHRpb25hbCBzdHJpbmcgbmFtZSA9IDE7CgogVGh1cywgdGhlIGFib3ZlIHBhdGggZ2l2ZXMgdGhlIGxvY2F0aW9uIG9mIGEgZmllbGQgbmFtZS4gIElmIHdlIHJlbW92ZWQKIHRoZSBsYXN0IGVsZW1lbnQ6CiAgIFsgNCwgMywgMiwgNyBdCiB0aGlzIHBhdGggcmVmZXJzIHRvIHRoZSB3aG9sZSBmaWVsZCBkZWNsYXJhdGlvbiAoZnJvbSB0aGUgYmVnaW5uaW5nCiBvZiB0aGUgbGFiZWwgdG8gdGhlIHRlcm1pbmF0aW5nIHNlbWljb2xvbikuCgoPCgcEEwMAAgAEEgSRBgQMCg8KBwQTAwACAAUSBJEGDRIKDwoHBBMDAAIAARIEkQYTFwoPCgcEEwMAAgADEgSRBhobCg8KBwQTAwACAAgSBJEGHCkKEgoKBBMDAAIACOcHABIEkQYdKAoTCgsEEwMAAgAI5wcAAhIEkQYdIwoUCgwEEwMAAgAI5wcAAgASBJEGHSMKFQoNBBMDAAIACOcHAAIAARIEkQYdIwoTCgsEEwMAAgAI5wcAAxIEkQYkKArSAgoGBBMDAAIBEgSYBgQqGsECIEFsd2F5cyBoYXMgZXhhY3RseSB0aHJlZSBvciBmb3VyIGVsZW1lbnRzOiBzdGFydCBsaW5lLCBzdGFydCBjb2x1bW4sCiBlbmQgbGluZSAob3B0aW9uYWwsIG90aGVyd2lzZSBhc3N1bWVkIHNhbWUgYXMgc3RhcnQgbGluZSksIGVuZCBjb2x1bW4uCiBUaGVzZSBhcmUgcGFja2VkIGludG8gYSBzaW5nbGUgZmllbGQgZm9yIGVmZmljaWVuY3kuICBOb3RlIHRoYXQgbGluZQogYW5kIGNvbHVtbiBudW1iZXJzIGFyZSB6ZXJvLWJhc2VkIC0tIHR5cGljYWxseSB5b3Ugd2lsbCB3YW50IHRvIGFkZAogMSB0byBlYWNoIGJlZm9yZSBkaXNwbGF5aW5nIHRvIGEgdXNlci4KCg8KBwQTAwACAQQSBJgGBAwKDwoHBBMDAAIBBRIEmAYNEgoPCgcEEwMAAgEBEgSYBhMXCg8KBwQTAwACAQMSBJgGGhsKDwoHBBMDAAIBCBIEmAYcKQoSCgoEEwMAAgEI5wcAEgSYBh0oChMKCwQTAwACAQjnBwACEgSYBh0jChQKDAQTAwACAQjnBwACABIEmAYdIwoVCg0EEwMAAgEI5wcAAgABEgSYBh0jChMKCwQTAwACAQjnBwADEgSYBiQoCqUMCgYEEwMAAgISBMkGBCkalAwgSWYgdGhpcyBTb3VyY2VDb2RlSW5mbyByZXByZXNlbnRzIGEgY29tcGxldGUgZGVjbGFyYXRpb24sIHRoZXNlIGFyZSBhbnkKIGNvbW1lbnRzIGFwcGVhcmluZyBiZWZvcmUgYW5kIGFmdGVyIHRoZSBkZWNsYXJhdGlvbiB3aGljaCBhcHBlYXIgdG8gYmUKIGF0dGFjaGVkIHRvIHRoZSBkZWNsYXJhdGlvbi4KCiBBIHNlcmllcyBvZiBsaW5lIGNvbW1lbnRzIGFwcGVhcmluZyBvbiBjb25zZWN1dGl2ZSBsaW5lcywgd2l0aCBubyBvdGhlcgogdG9rZW5zIGFwcGVhcmluZyBvbiB0aG9zZSBsaW5lcywgd2lsbCBiZSB0cmVhdGVkIGFzIGEgc2luZ2xlIGNvbW1lbnQuCgogbGVhZGluZ19kZXRhY2hlZF9jb21tZW50cyB3aWxsIGtlZXAgcGFyYWdyYXBocyBvZiBjb21tZW50cyB0aGF0IGFwcGVhcgogYmVmb3JlIChidXQgbm90IGNvbm5lY3RlZCB0bykgdGhlIGN1cnJlbnQgZWxlbWVudC4gRWFjaCBwYXJhZ3JhcGgsCiBzZXBhcmF0ZWQgYnkgZW1wdHkgbGluZXMsIHdpbGwgYmUgb25lIGNvbW1lbnQgZWxlbWVudCBpbiB0aGUgcmVwZWF0ZWQKIGZpZWxkLgoKIE9ubHkgdGhlIGNvbW1lbnQgY29udGVudCBpcyBwcm92aWRlZDsgY29tbWVudCBtYXJrZXJzIChlLmcuIC8vKSBhcmUKIHN0cmlwcGVkIG91dC4gIEZvciBibG9jayBjb21tZW50cywgbGVhZGluZyB3aGl0ZXNwYWNlIGFuZCBhbiBhc3Rlcmlzawogd2lsbCBiZSBzdHJpcHBlZCBmcm9tIHRoZSBiZWdpbm5pbmcgb2YgZWFjaCBsaW5lIG90aGVyIHRoYW4gdGhlIGZpcnN0LgogTmV3bGluZXMgYXJlIGluY2x1ZGVkIGluIHRoZSBvdXRwdXQuCgogRXhhbXBsZXM6CgogICBvcHRpb25hbCBpbnQzMiBmb28gPSAxOyAgLy8gQ29tbWVudCBhdHRhY2hlZCB0byBmb28uCiAgIC8vIENvbW1lbnQgYXR0YWNoZWQgdG8gYmFyLgogICBvcHRpb25hbCBpbnQzMiBiYXIgPSAyOwoKICAgb3B0aW9uYWwgc3RyaW5nIGJheiA9IDM7CiAgIC8vIENvbW1lbnQgYXR0YWNoZWQgdG8gYmF6LgogICAvLyBBbm90aGVyIGxpbmUgYXR0YWNoZWQgdG8gYmF6LgoKICAgLy8gQ29tbWVudCBhdHRhY2hlZCB0byBxdXguCiAgIC8vCiAgIC8vIEFub3RoZXIgbGluZSBhdHRhY2hlZCB0byBxdXguCiAgIG9wdGlvbmFsIGRvdWJsZSBxdXggPSA0OwoKICAgLy8gRGV0YWNoZWQgY29tbWVudCBmb3IgY29yZ2UuIFRoaXMgaXMgbm90IGxlYWRpbmcgb3IgdHJhaWxpbmcgY29tbWVudHMKICAgLy8gdG8gcXV4IG9yIGNvcmdlIGJlY2F1c2UgdGhlcmUgYXJlIGJsYW5rIGxpbmVzIHNlcGFyYXRpbmcgaXQgZnJvbQogICAvLyBib3RoLgoKICAgLy8gRGV0YWNoZWQgY29tbWVudCBmb3IgY29yZ2UgcGFyYWdyYXBoIDIuCgogICBvcHRpb25hbCBzdHJpbmcgY29yZ2UgPSA1OwogICAvKiBCbG9jayBjb21tZW50IGF0dGFjaGVkCiAgICAqIHRvIGNvcmdlLiAgTGVhZGluZyBhc3Rlcmlza3MKICAgICogd2lsbCBiZSByZW1vdmVkLiAqLwogICAvKiBCbG9jayBjb21tZW50IGF0dGFjaGVkIHRvCiAgICAqIGdyYXVsdC4gKi8KICAgb3B0aW9uYWwgaW50MzIgZ3JhdWx0ID0gNjsKCiAg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bGUucHJvdG9idWYuRW51bU9wdGlvbnMYseQDIAEoCFIRZ29wcm90b0VudW1QcmVmaXg6UgoVZ29wcm90b19lbnVtX3N0cmluZ2VyEhwuZ29vZ2xlLnByb3RvYnVmLkVudW1PcHRpb25zGMXkAyABKAhSE2dvcHJvdG9FbnVtU3RyaW5nZXI6QwoNZW51bV9zdHJpbmdlchIcLmdvb2dsZS5wcm90b2J1Zi5FbnVtT3B0aW9ucxjG5AMgASgIUgxlbnVtU3RyaW5nZXI6RwoPZW51bV9jdXN0b21uYW1lEhwuZ29vZ2xlLnByb3RvYnVmLkVudW1PcHRpb25zGMfkAyABKAlSDmVudW1DdXN0b21uYW1lOjoKCGVudW1kZWNsEhwuZ29vZ2xlLnByb3RvYnVmLkVudW1PcHRpb25zGMjkAyABKAhSCGVudW1kZWNsOlYKFGVudW12YWx1ZV9jdXN0b21uYW1lEiEuZ29vZ2xlLnByb3RvYnVmLkVudW1WYWx1ZU9wdGlvbnMY0YMEIAEoCVITZW51bXZhbHVlQ3VzdG9tbmFtZTpOChNnb3Byb3RvX2dldHRlcnNfYWxsEhwuZ29vZ2xlLnByb3RvYnVmLkZpbGVPcHRpb25zGJnsAyABKAhSEWdvcHJvdG9HZXR0ZXJzQWxsOlUKF2dvcHJvdG9fZW51bV9wcmVmaXhfYWxsEhwuZ29vZ2xlLnByb3RvYnVmLkZpbGVPcHRpb25zGJrsAyABKAhSFGdvcHJvdG9FbnVtUHJlZml4QWxsOlAKFGdvcHJvdG9fc3RyaW5nZXJfYWxsEhwuZ29vZ2xlLnByb3RvYnVmLkZpbGVPcHRpb25zGJvsAyABKAhSEmdvcHJvdG9TdHJpbmdlckFsbDpKChF2ZXJib3NlX2VxdWFsX2FsbBIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxic7AMgASgIUg92ZXJib3NlRXF1YWxBbGw6OQoIZmFjZV9hbGwSHC5nb29nbGUucHJvdG9idWYuRmlsZU9wdGlvbnMYnewDIAEoCFIHZmFjZUFsbDpBCgxnb3N0cmluZ19hbGwSHC5nb29nbGUucHJvdG9idWYuRmlsZU9wdGlvbnMYnuwDIAEoCFILZ29zdHJpbmdBbGw6QQoMcG9wdWxhdGVfYWxsEhwuZ29vZ2xlLnByb3RvYnVmLkZpbGVPcHRpb25zGJ/sAyABKAhSC3BvcHVsYXRlQWxsOkEKDHN0cmluZ2VyX2FsbBIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxig7AMgASgIUgtzdHJpbmdlckFsbDo/Cgtvbmx5b25lX2FsbBIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxih7AMgASgIUgpvbmx5b25lQWxsOjsKCWVxdWFsX2FsbBIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxil7AMgASgIUghlcXVhbEFsbDpHCg9kZXNjcmlwdGlvbl9hbGwSHC5nb29nbGUucHJvdG9idWYuRmlsZU9wdGlvbnMYpuwDIAEoCFIOZGVzY3JpcHRpb25BbGw6PwoLdGVzdGdlbl9hbGwSHC5nb29nbGUucHJvdG9idWYuRmlsZU9wdGlvbnMYp+wDIAEoCFIKdGVzdGdlbkFsbDpBCgxiZW5jaGdlbl9hbGwSHC5nb29nbGUucHJvdG9idWYuRmlsZU9wdGlvbnMYqOwDIAEoCFILYmVuY2hnZW5BbGw6QwoNbWFyc2hhbGVyX2FsbBIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxip7AMgASgIUgxtYXJzaGFsZXJBbGw6RwoPdW5tYXJzaGFsZXJfYWxsEhwuZ29vZ2xlLnByb3RvYnVmLkZpbGVPcHRpb25zGKrsAyABKAhSDnVubWFyc2hhbGVyQWxsOlAKFHN0YWJsZV9tYXJzaGFsZXJfYWxsEhwuZ29vZ2xlLnByb3RvYnVmLkZpbGVPcHRpb25zGKvsAyABKAhSEnN0YWJsZU1hcnNoYWxlckFsbDo7CglzaXplcl9hbGwSHC5nb29nbGUucHJvdG9idWYuRmlsZU9wdGlvbnMYrOwDIAEoCFIIc2l6ZXJBbGw6WQoZZ29wcm90b19lbnVtX3N0cmluZ2VyX2FsbBIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxit7AMgASgIUhZnb3Byb3RvRW51bVN0cmluZ2VyQWxsOkoKEWVudW1fc3RyaW5nZXJfYWxsEhwuZ29vZ2xlLnByb3RvYnVmLkZpbGVPcHRpb25zGK7sAyABKAhSD2VudW1TdHJpbmdlckFsbDpQChR1bnNhZmVfbWFyc2hhbGVyX2FsbBIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxiv7AMgASgIUhJ1bnNhZmVNYXJzaGFsZXJBbGw6VAoWdW5zYWZlX3VubWFyc2hhbGVyX2FsbBIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxiw7AMgASgIUhR1bnNhZmVVbm1hcnNoYWxlckFsbDpbChpnb3Byb3RvX2V4dGVuc2lvbnNfbWFwX2FsbBIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxix7AMgASgIUhdnb3Byb3RvRXh0ZW5zaW9uc01hcEFsbDpYChhnb3Byb3RvX3VucmVjb2duaXplZF9hbGwSHC5nb29nbGUucHJvdG9idWYuRmlsZU9wdGlvbnMYsuwDIAEoCFIWZ29wcm90b1VucmVjb2duaXplZEFsbDpJChBnb2dvcHJvdG9faW1wb3J0EhwuZ29vZ2xlLnByb3RvYnVmLkZpbGVPcHRpb25zGLPsAyABKAhSD2dvZ29wcm90b0ltcG9ydDpFCg5wcm90b3NpemVyX2FsbBIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxi07AMgASgIUg1wcm90b3NpemVyQWxsOj8KC2NvbXBhcmVfYWxsEhwuZ29vZ2xlLnByb3RvYnVmLkZpbGVPcHRpb25zGLXsAyABKAhSCmNvbXBhcmVBbGw6QQoMdHlwZWRlY2xfYWxsEhwuZ29vZ2xlLnByb3RvYnVmLkZpbGVPcHRpb25zGLbsAyABKAhSC3R5cGVkZWNsQWxsOkEKDGVudW1kZWNsX2FsbBIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxi37AMgASgIUgtlbnVtZGVjbEFsbDpRChRnb3Byb3RvX3JlZ2lzdHJhdGlvbhIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxi47AMgASgIUhNnb3Byb3RvUmVnaXN0cmF0aW9uOkcKD21lc3NhZ2VuYW1lX2FsbBIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxi57AMgASgIUg5tZXNzYWdlbmFtZUFsbDpKCg9nb3Byb3RvX2dldHRlcnMSHy5nb29nbGUucHJvdG9idWYuTWVzc2FnZU9wdGlvbnMYgfQDIAEoCFIOZ29wcm90b0dldHRlcnM6TAoQZ29wcm90b19zdHJpbmdlchIfLmdvb2dsZS5wcm90b2J1Zi5NZXNzYWdlT3B0aW9ucxiD9AMgASgIUg9nb3Byb3RvU3RyaW5nZXI6RgoNdmVyYm9zZV9lcXVhbBIfLmdvb2dsZS5wcm90b2J1Zi5NZXNzYWdlT3B0aW9ucxiE9AMgASgIUgx2ZXJib3NlRXF1YWw6NQoEZmFjZRIfLmdvb2dsZS5wcm90b2J1Zi5NZXNzYWdlT3B0aW9ucxiF9AMgASgIUgRmYWNlOj0KCGdvc3RyaW5nEh8uZ29vZ2xlLnByb3RvYnVmLk1lc3NhZ2VPcHRpb25zGIb0AyABKAhSCGdvc3RyaW5nOj0KCHBvcHVsYXRlEh8uZ29vZ2xlLnByb3RvYnVmLk1lc3NhZ2VPcHRpb25zGIf0AyABKAhSCHBvcHVsYXRlOj0KCHN0cmluZ2VyEh8uZ29vZ2xlLnByb3RvYnVmLk1lc3NhZ2VPcHRpb25zGMCLBCABKAhSCHN0cmluZ2VyOjsKB29ubHlvbmUSHy5nb29nbGUucHJvdG9idWYuTWVzc2FnZU9wdGlvbnMYifQDIAEoCFIHb25seW9uZTo3CgVlcXVhbBIfLmdvb2dsZS5wcm90b2J1Zi5NZXNzYWdlT3B0aW9ucxiN9AMgASgIUgVlcXVhbDpDCgtkZXNjcmlwdGlvbhIfLmdvb2dsZS5wcm90b2J1Zi5NZXNzYWdlT3B0aW9ucxiO9AMgASgIUgtkZXNjcmlwdGlvbjo7Cgd0ZXN0Z2VuEh8uZ29vZ2xlLnByb3RvYnVmLk1lc3NhZ2VPcHRpb25zGI/0AyABKAhSB3Rlc3RnZW46PQoIYmVuY2hnZW4SHy5nb29nbGUucHJvdG9idWYuTWVzc2FnZU9wdGlvbnMYkPQDIAEoCFIIYmVuY2hnZW46PwoJbWFyc2hhbGVyEh8uZ29vZ2xlLnByb3RvYnVmLk1lc3NhZ2VPcHRpb25zGJH0AyABKAhSCW1hcnNoYWxlcjpDCgt1bm1hcnNoYWxlchIfLmdvb2dsZS5wcm90b2J1Zi5NZXNzYWdlT3B0aW9ucxiS9AMgASgIUgt1bm1hcnNoYWxlcjpMChBzdGFibGVfbWFyc2hhbGVyEh8uZ29vZ2xlLnByb3RvYnVmLk1lc3NhZ2VPcHRpb25zGJP0AyABKAhSD3N0YWJsZU1hcnNoYWxlcjo3CgVzaXplchIfLmdvb2dsZS5wcm90b2J1Zi5NZXNzYWdlT3B0aW9ucxiU9AMgASgIUgVzaXplcjpMChB1bnNhZmVfbWFyc2hhbGVyEh8uZ29vZ2xlLnByb3RvYnVmLk1lc3NhZ2VPcHRpb25zGJf0AyABKAhSD3Vuc2FmZU1hcnNoYWxlcjpQChJ1bnNhZmVfdW5tYXJzaGFsZXISHy5nb29nbGUucHJvdG9idWYuTWVzc2FnZU9wdGlvbnMYmPQDIAEoCFIRdW5zYWZlVW5tYXJzaGFsZXI6VwoWZ29wcm90b19leHRlbnNpb25zX21hcBIfLmdvb2dsZS5wcm90b2J1Zi5NZXNzYWdlT3B0aW9ucxiZ9AMgASgIUhRnb3Byb3RvRXh0ZW5zaW9uc01hcDpUChRnb3Byb3RvX3VucmVjb2duaXplZBIfLmdvb2dsZS5wcm90b2J1Zi5NZXNzYWdlT3B0aW9ucxia9AMgASgIUhNnb3Byb3RvVW5yZWNvZ25pemVkOkEKCnByb3Rvc2l6ZXISHy5nb29nbGUucHJvdG9idWYuTWVzc2FnZU9wdGlvbnMYnPQDIAEoCFIKcHJvdG9zaXplcjo7Cgdjb21wYXJlEh8uZ29vZ2xlLnByb3RvYnVmLk1lc3NhZ2VPcHRpb25zGJ30AyABKAhSB2NvbXBhcmU6PQoIdHlwZWRlY2wSHy5nb29nbGUucHJvdG9idWYuTWVzc2FnZU9wdGlvbnMYnvQDIAEoCFIIdHlwZWRlY2w6QwoLbWVzc2FnZW5hbWUSHy5nb29nbGUucHJvdG9idWYuTWVzc2FnZU9wdGlvbnMYofQDIAEoCFILbWVzc2FnZW5hbWU6OwoIbnVsbGFibGUSHS5nb29nbGUucHJvdG9idWYuRmllbGRPcHRpb25zGOn7AyABKAhSCG51bGxhYmxlOjUKBWVtYmVkEh0uZ29vZ2xlLnByb3RvYnVmLkZpZWxkT3B0aW9ucxjq+wMgASgIUgVlbWJlZDo/CgpjdXN0b210eXBlEh0uZ29vZ2xlLnByb3RvYnVmLkZpZWxkT3B0aW9ucxjr+wMgASgJUgpjdXN0b210eXBlOj8KCmN1c3RvbW5hbWUSHS5nb29nbGUucHJvdG9idWYuRmllbGRPcHRpb25zGOz7AyABKAlSCmN1c3RvbW5hbWU6OQoHanNvbnRhZxIdLmdvb2dsZS5wcm90b2J1Zi5GaWVsZE9wdGlvbnMY7fsDIAEoCVIHanNvbnRhZzo7Cghtb3JldGFncxIdLmdvb2dsZS5wcm90b2J1Zi5GaWVsZE9wdGlvbnMY7vsDIAEoCVIIbW9yZXRhZ3M6OwoIY2FzdHR5cGUSHS5nb29nbGUucHJvdG9idWYuRmllbGRPcHRpb25zGO/7AyABKAlSCGNhc3R0eXBlOjkKB2Nhc3RrZXkSHS5nb29nbGUucHJvdG9idWYuRmllbGRPcHRpb25zGPD7AyABKAlSB2Nhc3RrZXk6PQoJY2FzdHZhbHVlEh0uZ29vZ2xlLnByb3RvYnVmLkZpZWxkT3B0aW9ucxjx+wMgASgJUgljYXN0dmFsdWU6OQoHc3RkdGltZRIdLmdvb2dsZS5wcm90b2J1Zi5GaWVsZE9wdGlvbnMY8vsDIAEoCFIHc3RkdGltZTpBCgtzdGRkdXJhdGlvbhIdLmdvb2dsZS5wcm90b2J1Zi5GaWVsZE9wdGlvbnMY8/sDIAEoCFILc3RkZHVyYXRpb25CRQoTY29tLmdvb2dsZS5wcm90b2J1ZkIKR29Hb1Byb3Rvc1oiZ2l0aHViLmNvbS9nb2dvL3Byb3RvYnVmL2dvZ29wcm90b0qaNQoHEgUcAIcBAQr8CgoBDBIDHAASMvEKIFByb3RvY29sIEJ1ZmZlcnMgZm9yIEdvIHdpdGggR2FkZ2V0cwoKIENvcHlyaWdodCAoYykgMjAxMywgVGhlIEdvR28gQXV0aG9ycy4gQWxsIHJpZ2h0cyByZXNlcnZlZC4KIGh0dHA6Ly9naXRodWIuY29tL2dvZ28vcHJvdG9idWYKCiBSZWRpc3RyaWJ1dGlvbiBhbmQgdXNlIGluIHNvdXJjZSBhbmQgYmluYXJ5IGZvcm1zLCB3aXRoIG9yIHdpdGhvdXQKIG1vZGlmaWNhdGlvbiwgYXJlIHBlcm1pdHRlZCBwcm92aWRlZCB0aGF0IHRoZSBmb2xsb3dpbmcgY29uZGl0aW9ucyBhcmUKIG1ldDoKCiAgICAgKiBSZWRpc3RyaWJ1dGlvbnMgb2Ygc291cmNlIGNvZGUgbXVzdCByZXRhaW4gdGhlIGFib3ZlIGNvcHlyaWdodAogbm90aWNlLCB0aGlzIGxpc3Qgb2YgY29uZGl0aW9ucyBhbmQgdGhlIGZvbGxvd2luZyBkaXNjbGFpbWVyLgogICAgICogUmVkaXN0cmlidXRpb25zIGluIGJpbmFyeSBmb3JtIG11c3QgcmVwcm9kdWNlIHRoZSBhYm92ZQogY29weXJpZ2h0IG5vdGljZSwgdGhpcyBsaXN0IG9mIGNvbmRpdGlvbnMgYW5kIHRoZSBmb2xsb3dpbmcgZGlzY2xhaW1lcgogaW4gdGhlIGRvY3VtZW50YXRpb24gYW5kL29yIG90aGVyIG1hdGVyaWFscyBwcm92aWRlZCB3aXRoIHRoZQogZGlzdHJpYnV0aW9uLgoKIFRISVMgU09GVFdBUkUgSVMgUFJPVklERUQgQlkgVEhFIENPUFlSSUdIVCBIT0xERVJTIEFORCBDT05UUklCVVRPUlMKICJBUyBJUyIgQU5EIEFOWSBFWFBSRVNTIE9SIElNUExJRUQgV0FSUkFOVElFUywgSU5DTFVESU5HLCBCVVQgTk9UCiBMSU1JVEVEIFRPLCBUSEUgSU1QTElFRCBXQVJSQU5USUVTIE9GIE1FUkNIQU5UQUJJTElUWSBBTkQgRklUTkVTUyBGT1IKIEEgUEFSVElDVUxBUiBQVVJQT1NFIEFSRSBESVNDTEFJTUVELiBJTiBOTyBFVkVOVCBTSEFMTCBUSEUgQ09QWVJJR0hUCiBPV05FUiBPUiBDT05UUklCVVRPUlMgQkUgTElBQkxFIEZPUiBBTlkgRElSRUNULCBJTkRJUkVDVCwgSU5DSURFTlRBTCwKIFNQRUNJQUwsIEVYRU1QTEFSWSwgT1IgQ09OU0VRVUVOVElBTCBEQU1BR0VTIChJTkNMVURJTkcsIEJVVCBOT1QKIExJTUlURUQgVE8sIFBST0NVUkVNRU5UIE9GIFNVQlNUSVRVVEUgR09PRFMgT1IgU0VSVklDRVM7IExPU1MgT0YgVVNFLAogREFUQSwgT1IgUFJPRklUUzsgT1IgQlVTSU5FU1MgSU5URVJSVVBUSU9OKSBIT1dFVkVSIENBVVNFRCBBTkQgT04gQU5ZCiBUSEVPUlkgT0YgTElBQklMSVRZLCBXSEVUSEVSIElOIENPTlRSQUNULCBTVFJJQ1QgTElBQklM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B0USBIUBCCYKCgoDB0UCEgN6ByMKCwoDB0UEEgSFAQgQCgsKAwdFBRIEhQERFQoLCgMHRQESBIUBFh0KCwoDB0UDEgSFASAlCgoKAgdGEgSGAQgqCgoKAwdGAhIDegcjCgsKAwdGBBIEhgEIEAoLCgMHRgUSBIYBERUKCwoDB0YBEgSGARYhCgsKAwdGAxIEhgEkKQqMFwosbWl4ZXIvYWRhcHRlci9tb2RlbC92MWJldGExL2V4dGVuc2lvbnMucHJvdG8SIWlzdGlvLm1peGVyLmFkYXB0ZXIubW9kZWwudjFiZXRhMRogZ29vZ2xlL3Byb3RvYnVmL2Rlc2NyaXB0b3IucHJvdG8quAEKD1RlbXBsYXRlVmFyaWV0eRIaChZURU1QTEFURV9WQVJJRVRZX0NIRUNLEAASGwoXVEVNUExBVEVfVkFSSUVUWV9SRVBPUlQQARIaChZURU1QTEFURV9WQVJJRVRZX1FVT1RBEAISKAokVEVNUExBVEVfVkFSSUVUWV9BVFRSSUJVVEVfR0VORVJBVE9SEAMSJgoiVEVNUExBVEVfVkFSSUVUWV9DSEVDS19XSVRIX09VVFBVVBAEOn4KEHRlbXBsYXRlX3ZhcmlldHkSHC5nb29nbGUucHJvdG9idWYuRmlsZU9wdGlvbnMYr8q8IiABKA4yMi5pc3Rpby5taXhlci5hZGFwdGVyLm1vZGVsLnYxYmV0YTEuVGVtcGxhdGVWYXJpZXR5Ug90ZW1wbGF0ZVZhcmlldHk6RAoNdGVtcGxhdGVfbmFtZRIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxjQy7wiIAEoCVIMdGVtcGxhdGVOYW1lQipaKGlzdGlvLmlvL2FwaS9taXhlci9hZGFwdGVyL21vZGVsL3YxYmV0YTFK4RIKBhIEDgAyAQq/BAoBDBIDDgASMrQEIENvcHlyaWdodCAyMDE4IElzdGlvIEF1dGhvcnMKCiBMaWNlbnNlZCB1bmRlciB0aGUgQXBhY2hlIExpY2Vuc2UsIFZlcnNpb24gMi4wICh0aGUgIkxpY2Vuc2UiKTsKIHlvdSBtYXkgbm90IHVzZSB0aGlzIGZpbGUgZXhjZXB0IGluIGNvbXBsaWFuY2Ugd2l0aCB0aGUgTGljZW5zZS4KIFlvdSBtYXkgb2J0YWluIGEgY29weSBvZiB0aGUgTGljZW5zZSBhdAoKICAgICBodHRwOi8vd3d3LmFwYWNoZS5vcmcvbGljZW5zZXMvTElDRU5TRS0yLjAKCiBVbmxlc3MgcmVxdWlyZWQgYnkgYXBwbGljYWJsZSBsYXcgb3IgYWdyZWVkIHRvIGluIHdyaXRpbmcsIHNvZnR3YXJlCiBkaXN0cmlidXRlZCB1bmRlciB0aGUgTGljZW5zZSBpcyBkaXN0cmlidXRlZCBvbiBhbiAiQVMgSVMiIEJBU0lTLAogV0lUSE9VVCBXQVJSQU5USUVTIE9SIENPTkRJVElPTlMgT0YgQU5ZIEtJTkQsIGVpdGhlciBleHByZXNzIG9yIGltcGxpZWQuCiBTZWUgdGhlIExpY2Vuc2UgZm9yIHRoZSBzcGVjaWZpYyBsYW5ndWFnZSBnb3Zlcm5pbmcgcGVybWlzc2lvbnMgYW5kCiBsaW1pdGF0aW9ucyB1bmRlciB0aGUgTGljZW5zZS4KCggKAQISAxAIKQoICgEIEgMSAD0KCwoECOcHABIDEgA9CgwKBQjnBwACEgMSBxEKDQoGCOcHAAIAEgMSBxEKDgoHCOcHAAIAARIDEgcRCgwKBQjnBwAHEgMSEjwKCQoCAwASAxQHKQp5CgIFABIEGAAoARptIFRoZSBhdmFpbGFibGUgdmFyaWV0aWVzIG9mIHRlbXBsYXRlcywgY29udHJvbGxpbmcgdGhlIHNlbWFudGljcyBvZiB3aGF0IGFuIGFkYXB0ZXIgZG9lcyB3aXRoIGVhY2ggaW5zdGFuY2UuCgoKCgMFAAESAxgFFArHAQoEBQACABIDGwQfGrkBIE1ha2VzIHRoZSB0ZW1wbGF0ZSBhcHBsaWNhYmxlIGZvciBNaXhlcidzIGNoZWNrIGNhbGxzLiBJbnN0YW5jZXMgb2Ygc3VjaCB0ZW1wbGF0ZSBhcmUgY3JlYXRlZCBkdXJpbmcKIGNoZWNrIGNhbGxzIGluIE1peGVyIGFuZCBwYXNzZWQgdG8gdGhlIGhhbmRsZXJzIGJhc2VkIG9uIHRoZSBydWxlIGNvbmZpZ3VyYXRpb25zLgoKDAoFBQACAAESAxsEGgoMCgUFAAIAAhIDGx0eCskBCgQFAAIBEgMeBCAauwEgTWFrZXMgdGhlIHRlbXBsYXRlIGFwcGxpY2FibGUgZm9yIE1peGVyJ3MgcmVwb3J0IGNhbGxzLiBJbnN0YW5jZXMgb2Ygc3VjaCB0ZW1wbGF0ZSBhcmUgY3JlYXRlZCBkdXJpbmcKIHJlcG9ydCBjYWxscyBpbiBNaXhlciBhbmQgcGFzc2VkIHRvIHRoZSBoYW5kbGVycyBiYXNlZCBvbiB0aGUgcnVsZSBjb25maWd1cmF0aW9ucy4KCgwKBQUAAgEBEgMeBBsKDAoFBQACAQISAx4eHwrNAQoEBQACAhIDIQQfGr8BIE1ha2VzIHRoZSB0ZW1wbGF0ZSBhcHBsaWNhYmxlIGZvciBNaXhlcidzIHF1b3RhIGNhbGxzLiBJbnN0YW5jZXMgb2Ygc3VjaCB0ZW1wbGF0ZSBhcmUgY3JlYXRlZCBkdXJpbmcKIHF1b3RhIGNoZWNrIGNhbGxzIGluIE1peGVyIGFuZCBwYXNzZWQgdG8gdGhlIGhhbmRsZXJzIGJhc2VkIG9uIHRoZSBydWxlIGNvbmZpZ3VyYXRpb25zLgoKDAoFBQACAgESAyEEGgoMCgUFAAICAhIDIR0eCusBCgQFAAIDEgMkBC0a3QEgTWFrZXMgdGhlIHRlbXBsYXRlIGFwcGxpY2FibGUgZm9yIE1peGVyJ3MgYXR0cmlidXRlIGdlbmVyYXRpb24gcGhhc2UuIEluc3RhbmNlcyBvZiBzdWNoIHRlbXBsYXRlIGFyZSBjcmVhdGVkIGR1cmluZwogcHJlLXByb2Nlc3NpbmcgYXR0cmlidXRlIGdlbmVyYXRpb24gcGhhc2UgYW5kIHBhc3NlZCB0byB0aGUgaGFuZGxlcnMgYmFzZWQgb24gdGhlIHJ1bGUgY29uZmlndXJhdGlvbnMuCgoMCgUFAAIDARIDJAQoCgwKBQUAAgMCEgMkKywKugEKBAUAAgQSAycEKxqsASBNYWtlcyB0aGUgdGVtcGxhdGUgYXBwbGljYWJsZSBmb3IgTWl4ZXIncyBjaGVjayBjYWxscy4gSW5zdGFuY2VzIG9mIHN1Y2ggdGVtcGxhdGUgYXJlIGNyZWF0ZWQgZHVyaW5nCiBjaGVjayBjYWxscyBpbiBNaXhlciBhbmQgcGFzc2VkIHRvIHRoZSBoYW5kbGVycyB0aGF0IHByb2R1Y2UgdmFsdWVzLgoKDAoFBQACBAESAycEJgoMCgUFAAIEAhIDJykqCjEKAQcSBCsAMgEaJiBGaWxlIGxldmVsIG9wdGlvbnMgZm9yIHRoZSB0ZW1wbGF0ZS4KCjYKAgcAEgMtBDAaKyBSZXF1aXJlZDogb3B0aW9uIGZvciB0aGUgVGVtcGxhdGVWYXJpZXR5LgoKCgoDBwACEgMrByIKCwoDBwAEEgQtBCskCgoKAwcABhIDLQQTCgoKAwcAARIDLRQkCgoKAwcAAxIDLScvCqQBCgIHARIDMQQkGpgBIE9wdGlvbmFsOiBvcHRpb24gZm9yIHRoZSB0ZW1wbGF0ZSBuYW1lLgogSWYgbm90IHNwZWNpZmllZCwgdGhlIGxhc3Qgc2VnbWVudCBvZiB0aGUgdGVtcGxhdGUgcHJvdG8ncyBwYWNrYWdlIG5hbWUgaXMgdXNlZCB0bwogZGVyaXZlIHRoZSB0ZW1wbGF0ZSBuYW1lLgoKCgoDBwECEgMrByIKCwoDBwEEEgQxBC0wCgoKAwcBBRIDMQQKCgoKAwcBARIDMQsYCgoKAwcBAxIDMRsjYgZwcm90bzMK3CwKGWdvb2dsZS9wcm90b2J1Zi9hbnkucHJvdG8SD2dvb2dsZS5wcm90b2J1ZiI2CgNBbnkSGQoIdHlwZV91cmwYASABKAlSB3R5cGVVcmwSFAoFdmFsdWUYAiABKAxSBXZhbHVlQm8KE2NvbS5nb29nbGUucHJvdG9idWZCCEFueVByb3RvUAFaJWdpdGh1Yi5jb20vZ29sYW5nL3Byb3RvYnVmL3B0eXBlcy9hbnmiAgNHUEKqAh5Hb29nbGUuUHJvdG9idWYuV2VsbEtub3duVHlwZXNK/CoKBxIFHgCUAQEKzAwKAQwSAx4AEjLBDCBQcm90b2NvbCBCdWZmZXJzIC0gR29vZ2xlJ3MgZGF0YSBpbnRlcmNoYW5nZSBmb3JtYXQKIENvcHlyaWdodCAyMDA4IEdvb2dsZSBJbmMuICBBbGwgcmlnaHRzIHJlc2VydmVkLgogaHR0cHM6Ly9kZXZlbG9wZXJzLmdvb2dsZS5jb20vcHJvdG9jb2wtYnVmZmVycy8KCiBSZWRpc3RyaWJ1dGlvbiBhbmQgdXNlIGluIHNvdXJjZSBhbmQgYmluYXJ5IGZvcm1zLCB3aXRoIG9yIHdpdGhvdXQKIG1vZGlmaWNhdGlvbiwgYXJlIHBlcm1pdHRlZCBwcm92aWRlZCB0aGF0IHRoZSBmb2xsb3dpbmcgY29uZGl0aW9ucyBhcmUKIG1ldDoKCiAgICAgKiBSZWRpc3RyaWJ1dGlvbnMgb2Ygc291cmNlIGNvZGUgbXVzdCByZXRhaW4gdGhlIGFib3ZlIGNvcHlyaWdodAogbm90aWNlLCB0aGlzIGxpc3Qgb2YgY29uZGl0aW9ucyBhbmQgdGhlIGZvbGxvd2luZyBkaXNjbGFpbWVyLgogICAgICogUmVkaXN0cmlidXRpb25zIGluIGJpbmFyeSBmb3JtIG11c3QgcmVwcm9kdWNlIHRoZSBhYm92ZQogY29weXJpZ2h0IG5vdGljZSwgdGhpcyBsaXN0IG9mIGNvbmRpdGlvbnMgYW5kIHRoZSBmb2xsb3dpbmcgZGlzY2xhaW1lcgogaW4gdGhlIGRvY3VtZW50YXRpb24gYW5kL29yIG90aGVyIG1hdGVyaWFscyBwcm92aWRlZCB3aXRoIHRoZQogZGlzdHJpYnV0aW9uLgogICAgICogTmVpdGhlciB0aGUgbmFtZSBvZiBHb29nbGUgSW5jLiBub3IgdGhlIG5hbWVzIG9mIGl0cwogY29udHJpYnV0b3JzIG1heSBiZSB1c2VkIHRvIGVuZG9yc2Ugb3IgcHJvbW90ZSBwcm9kdWN0cyBkZXJpdmVkIGZyb20KIHRoaXMgc29mdHdhcmUgd2l0aG91dCBzcGVjaWZpYyBwcmlvciB3cml0dGVuIHBlcm1pc3Npb24uCgogVEhJUyBTT0ZUV0FSRSBJUyBQUk9WSURFRCBCWSBUSEUgQ09QWVJJR0hUIEhPTERFUlMgQU5EIENPTlRSSUJVVE9SUwogIkFTIElTIiBBTkQgQU5ZIEVYUFJFU1MgT1IgSU1QTElFRCBXQVJSQU5USUVTLCBJTkNMVURJTkcsIEJVVCBOT1QKIExJTUlURUQgVE8sIFRIRSBJTVBMSUVEIFdBUlJBTlRJRVMgT0YgTUVSQ0hBTlRBQklMSVRZIEFORCBGSVRORVNTIEZPUgogQSBQQVJUSUNVTEFSIFBVUlBPU0UgQVJFIERJU0NMQUlNRUQuIElOIE5PIEVWRU5UIFNIQUxMIFRIRSBDT1BZUklHSFQKIE9XTkVSIE9SIENPTlRSSUJVVE9SUyBCRSBMSUFCTEUgRk9SIEFO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cGFjayBhIG1lc3NhZ2UgaW4gR28KCiAgICAgIGZvbyA6PSAmcGIuRm9vey4uLn0KICAgICAgYW55LCBlcnIgOj0gcHR5cGVzLk1hcnNoYWxBbnkoZm9vKQogICAgICAuLi4KICAgICAgZm9vIDo9ICZwYi5Gb297fQogICAgICBpZiBlcnIgOj0gcHR5cGVzLlVubWFyc2hhbEFueShhbnksIGZvbyk7IGVyciAhPSBuaWwgewogICAgICAgIC4uLgogICAgICB9CgogVGhlIHBhY2sgbWV0aG9kcyBwcm92aWRlZCBieSBwcm90b2J1ZiBsaWJyYXJ5IHdpbGwgYnkgZGVmYXVsdCB1c2UKICd0eXBlLmdvb2dsZWFwaXMuY29tL2Z1bGwudHlwZS5uYW1lJyBhcyB0aGUgdHlwZSBVUkwgYW5kIHRoZSB1bnBhY2sKIG1ldGhvZHMgb25seSB1c2UgdGhlIGZ1bGx5IHF1YWxpZmllZCB0eXBlIG5hbWUgYWZ0ZXIgdGhlIGxhc3QgJy8nCiBpbiB0aGUgdHlwZSBVUkwsIGZvciBleGFtcGxlICJmb28uYmFyLmNvbS94L3kueiIgd2lsbCB5aWVsZCB0eXBlCiBuYW1lICJ5LnoiLgoKCiBKU09OCiA9PT09CiBUaGUgSlNPTiByZXByZXNlbnRhdGlvbiBvZiBhbiBgQW55YCB2YWx1ZSB1c2VzIHRoZSByZWd1bGFyCiByZXByZXNlbnRhdGlvbiBvZiB0aGUgZGVzZXJpYWxpemVkLCBlbWJlZGRlZCBtZXNzYWdlLCB3aXRoIGFuCiBhZGRpdGlvbmFsIGZpZWxkIGBAdHlwZWAgd2hpY2ggY29udGFpbnMgdGhlIHR5cGUgVVJMLiBFeGFtcGxlOgoKICAgICBwYWNrYWdlIGdvb2dsZS5wcm9maWxlOwogICAgIG1lc3NhZ2UgUGVyc29uIHsKICAgICAgIHN0cmluZyBmaXJzdF9uYW1lID0gMTsKICAgICAgIHN0cmluZyBsYXN0X25hbWUgPSAyOwogICAgIH0KCiAgICAgewogICAgICAgIkB0eXBlIjogInR5cGUuZ29vZ2xlYXBpcy5jb20vZ29vZ2xlLnByb2ZpbGUuUGVyc29uIiwKICAgICAgICJmaXJzdE5hbWUiOiA8c3RyaW5nPiwKICAgICAgICJsYXN0TmFtZSI6IDxzdHJpbmc+CiAgICAgfQoKIElmIHRoZSBlbWJlZGRlZCBtZXNzYWdlIHR5cGUgaXMgd2VsbC1rbm93biBhbmQgaGFzIGEgY3VzdG9tIEpTT04KIHJlcHJlc2VudGF0aW9uLCB0aGF0IHJlcHJlc2VudGF0aW9uIHdpbGwgYmUgZW1iZWRkZWQgYWRkaW5nIGEgZmllbGQKIGB2YWx1ZWAgd2hpY2ggaG9sZHMgdGhlIGN1c3RvbSBKU09OIGluIGFkZGl0aW9uIHRvIHRoZSBgQHR5cGVgCiBmaWVsZC4gRXhhbXBsZSAoZm9yIG1lc3NhZ2UgW2dvb2dsZS5wcm90b2J1Zi5EdXJhdGlvbl1bXSk6CgogICAgIHsKICAgICAgICJAdHlwZSI6ICJ0eXBlLmdvb2dsZWFwaXMuY29tL2dvb2dsZS5wcm90b2J1Zi5EdXJhdGlvbiIsCiAgICAgICAidmFsdWUiOiAiMS4yMTJzIgogICAgIH0KCgoKCgMEAAESA3kICwrkBwoEBAACABIEkAECFhrVByBBIFVSTC9yZXNvdXJjZSBuYW1lIHdob3NlIGNvbnRlbnQgZGVzY3JpYmVzIHRoZSB0eXBlIG9mIHRoZQogc2VyaWFsaXplZCBwcm90b2NvbCBidWZmZXIgbWVzc2FnZS4KCiBGb3IgVVJMcyB3aGljaCB1c2UgdGhlIHNjaGVtZSBgaHR0cGAsIGBodHRwc2AsIG9yIG5vIHNjaGVtZSwgdGhlCiBmb2xsb3dpbmcgcmVzdHJpY3Rpb25zIGFuZCBpbnRlcnByZXRhdGlvbnMgYXBwbHk6CgogKiBJZiBubyBzY2hlbWUgaXMgcHJvdmlkZWQsIGBodHRwc2AgaXMgYXNzdW1lZC4KICogVGhlIGxhc3Qgc2VnbWVudCBvZiB0aGUgVVJMJ3MgcGF0aCBtdXN0IHJlcHJlc2VudCB0aGUgZnVsbHkKICAgcXVhbGlmaWVkIG5hbWUgb2YgdGhlIHR5cGUgKGFzIGluIGBwYXRoL2dvb2dsZS5wcm90b2J1Zi5EdXJhdGlvbmApLgogICBUaGUgbmFtZSBzaG91bGQgYmUgaW4gYSBjYW5vbmljYWwgZm9ybSAoZS5nLiwgbGVhZGluZyAiLiIgaXMKICAgbm90IGFjY2VwdGVkKS4KICogQW4gSFRUUCBHRVQgb24gdGhlIFVSTCBtdXN0IHlpZWxkIGEgW2dvb2dsZS5wcm90b2J1Zi5UeXBlXVtdCiAgIHZhbHVlIGluIGJpbmFyeSBmb3JtYXQsIG9yIHByb2R1Y2UgYW4gZXJyb3IuCiAqIEFwcGxpY2F0aW9ucyBhcmUgYWxsb3dlZCB0byBjYWNoZSBsb29rdXAgcmVzdWx0cyBiYXNlZCBvbiB0aGUKICAgVVJMLCBvciBoYXZlIHRoZW0gcHJlY29tcGlsZWQgaW50byBhIGJpbmFyeSB0byBhdm9pZCBhbnkKICAgbG9va3VwLiBUaGVyZWZvcmUsIGJpbmFyeSBjb21wYXRpYmlsaXR5IG5lZWRzIHRvIGJlIHByZXNlcnZlZAogICBvbiBjaGFuZ2VzIHRvIHR5cGVzLiAoVXNlIHZlcnNpb25lZCB0eXBlIG5hbWVzIHRvIG1hbmFnZQogICBicmVha2luZyBjaGFuZ2VzLikKCiBTY2hlbWVzIG90aGVyIHRoYW4gYGh0dHBgLCBgaHR0cHNgIChvciB0aGUgZW1wdHkgc2NoZW1lKSBtaWdodCBiZQogdXNlZCB3aXRoIGltcGxlbWVudGF0aW9uIHNwZWNpZmljIHNlbWFudGljcy4KCgoOCgUEAAIABBIFkAECeQ0KDQoFBAACAAUSBJABAggKDQoFBAACAAESBJABCREKDQoFBAACAAMSBJABFBUKVwoEBAACARIEkwECEhpJIE11c3QgYmUgYSB2YWxpZCBzZXJpYWxpemVkIHByb3RvY29sIGJ1ZmZlciBvZiB0aGUgYWJvdmUgc3BlY2lmaWVkIHR5cGUuCgoPCgUEAAIBBBIGkwECkAEWCg0KBQQAAgEFEgSTAQIHCg0KBQQAAgEBEgSTAQgNCg0KBQQAAgEDEgSTARARYgZwcm90bzMKmikKHmdvb2dsZS9wcm90b2J1Zi9kdXJhdGlvbi5wcm90bxIPZ29vZ2xlLnByb3RvYnVmIjoKCER1cmF0aW9uEhgKB3NlY29uZHMYASABKANSB3NlY29uZHMSFAoFbmFub3MYAiABKAVSBW5hbm9zQnwKE2NvbS5nb29nbGUucHJvdG9idWZCDUR1cmF0aW9uUHJvdG9QAVoqZ2l0aHViLmNvbS9nb2xhbmcvcHJvdG9idWYvcHR5cGVzL2R1cmF0aW9u+AEBogIDR1BCqgIeR29vZ2xlLlByb3RvYnVmLldlbGxLbm93blR5cGVzSqQnCgYSBB4AdAEKzAwKAQwSAx4AEjLBDCBQcm90b2NvbCBCdWZmZXJzIC0gR29vZ2xlJ3MgZGF0YSBpbnRlcmNoYW5nZSBmb3JtYXQKIENvcHlyaWdodCAyMDA4IEdvb2dsZSBJbmMuICBBbGwgcmlnaHRzIHJlc2VydmVkLgogaHR0cHM6Ly9kZXZlbG9wZXJzLmdvb2dsZS5jb20vcHJvdG9jb2wtYnVmZmVycy8KCiBSZWRpc3RyaWJ1dGlvbiBhbmQgdXNlIGluIHNvdXJjZSBhbmQgYmluYXJ5IGZvcm1zLCB3aXRoIG9yIHdpdGhvdXQKIG1vZGlmaWNhdGlvbiwgYXJlIHBlcm1pdHRlZCBwcm92aWRlZCB0aGF0IHRoZSBmb2xsb3dpbmcgY29uZGl0aW9ucyBhcmUKIG1ldDoKCiAgICAgKiBSZWRpc3RyaWJ1dGlvbnMgb2Ygc291cmNlIGNvZGUgbXVzdCByZXRhaW4gdGhlIGFib3ZlIGNvcHlyaWdodAogbm90aWNlLCB0aGlzIGxpc3Qgb2YgY29uZGl0aW9ucyBhbmQgdGhlIGZvbGxvd2luZyBkaXNjbGFpbWVyLgogICAgICogUmVkaXN0cmlidXRpb25zIGluIGJpbmFyeSBmb3JtIG11c3QgcmVwcm9kdWNlIHRoZSBhYm92ZQogY29weXJpZ2h0IG5vdGljZSwgdGhpcyBsaXN0IG9mIGNvbmRpdGlvbnMgYW5kIHRoZSBmb2xsb3dpbmcgZGlzY2xhaW1lcgogaW4gdGhlIGRvY3VtZW50YXRpb24gYW5kL29yIG90aGVyIG1hdGVyaWFscyBwcm92aWRlZCB3aXRoIHRoZQogZGlzdHJpYnV0aW9uLgogICAgICogTmVpdGhlciB0aGUgbmFtZSBvZiBHb29nbGUgSW5jLiBub3IgdGhlIG5hbWVzIG9mIGl0cwogY29udHJpYnV0b3JzIG1heSBiZSB1c2VkIHRvIGVuZG9yc2Ugb3IgcHJvbW90ZSBwcm9kdWN0cyBkZXJpdmVkIGZyb20KIHRoaXMgc29mdHdhcmUgd2l0aG91dCBzcGVjaWZpYyBwcmlvciB3cml0dGVuIHBlcm1pc3Npb24uCgogVEhJUyBTT0ZUV0FSRSBJUyBQUk9WSURFRCBCWSBUSEUgQ09QWVJJR0hUIEhPTERFUlMgQU5EIENPTlRSSUJVVE9SUwogIkFTIElTIiBBTkQgQU5ZIEVYUFJFU1MgT1IgSU1QTElFRCBXQVJSQU5USUVTLCBJTkNMVURJTkcsIEJVVCBOT1QKIExJTUlURUQgVE8sIFRIRSBJTVBMSUVEIFdBUlJBTlRJRVMgT0YgTUVSQ0hBTlRBQklMSVRZIEFORCBGSVRORVNTIEZPUgogQSBQQVJUSUNVTEFSIFBVUlBPU0UgQVJFIERJU0NMQUlNRUQuIElOIE5PIEVWRU5UIFNIQUxMIFRIRSBDT1BZUklHSFQKIE9XTkVSIE9SIENPTlRSSUJVVE9SUyBCRSBMSUFCTEUgRk9SIEFOWSBESVJFQ1QsIElORElSRUNULCBJTkNJREVOVEFMLAogU1BFQ0lBTCwgRVhFTVBMQVJZLCBPUiBDT05TRVFVRU5USUFMIERBTUFHRVMgKElOQ0xVRElORywgQlVUIE5PVAogTElNSVRFRCBUTywgUFJPQ1VSRU1FTlQgT0YgU1VCU1RJVFVURSBHT09EUyBPUiBTRVJWSUNFUzsgTE9TUyBPRiBVU0UsCiBEQVRBLCBPUiBQUk9GSVRTOyBPUiBCVVNJTkVTUyBJTlRFUlJVUFRJT04pIEhPV0VWRVIgQ0FVU0VEIEFORCBPTiBBTlkKIFRIRU9SWSBPRiBMSUFCSUxJVFksIFdIRVRIRVIgSU4gQ09OVFJBQ1QsIFNUUklDVCBMSUFCSUxJVFksIE9SIFRPUlQKIChJTkNMVURJTkcgTkVHTElHRU5DRSBPUiBPVEhFUldJU0UpIEFSSVNJTkcgSU4gQU5ZIFdBWSBPVVQgT0YgVEhFIFVTRQogT0YgVEhJUyBTT0ZUV0FSRSwgRVZFTiBJRiBBRFZJU0VEIE9GIFRIRSBQT1NTSUJJTElUWSBPRiBTVUNIIERBTUFHRS4KCggKAQISAyAIFwoICgEIEgMiADsKCwoECOcHABIDIgA7CgwKBQjnBwACEgMiBxcKDQoGCOcHAAIAEgMiBxcKDgoHCOcHAAIAARIDIgcXCgwKBQjnBwAHEgMiGjoKCAoBCBIDIwAfCgsKBAjnBwESAyMAHwoMCgUI5wcBAhIDIwcXCg0KBgjnBwECABIDIwcXCg4KBwjnBwECAAESAyMHFwoMCgUI5wcBAxIDIxoeCggKAQgSAyQAQQoLCgQI5wcCEgMkAEEKDAoFCOcHAgISAyQHEQoNCgYI5wcCAgASAyQHEQoOCgcI5wcCAgABEgMkBxEKDAoFCOcHAgcSAyQUQAoICgEIEgMlACwKCwoECOcHAxIDJQAsCgwKBQjnBwMCEgMlBxMKDQoGCOcHAwIAEgMlBxMKDgoHCOcHAwIAARIDJQcTCgwKBQjnBwMHEgMlFisKCAoBCBIDJgAuCgsKBAjnBwQSAyYALgoMCgUI5wcEAhIDJgcbCg0KBgjnBwQCABIDJgcbCg4KBwjnBwQCAAESAyYHGwoMCgUI5wcEBxIDJh4tCggKAQgSAycAIgoLCgQI5wcFEgMnACIKDAoFCOcHBQISAycHGgoNCgYI5wcFAgASAycHGgoOCgcI5wcFAgABEgMnBxoKDAoFCOcHBQMSAycdIQoICgEIEgMoACEKCwoECOcHBhIDKAAhCgwKBQjnBwYCEgMoBxgKDQoGCOcHBgIAEgMoBxgKDgoHCOcHBgIAARIDKAcYCgwKBQjnBwYHEgMoGyAKnxAKAgQAEgRmAHQBGpIQIEEgRHVyYXRpb24gcmVwcmVzZW50cyBhIHNpZ25lZCwgZml4ZWQtbGVuZ3RoIHNwYW4gb2YgdGltZSByZXByZXNlbnRlZAogYXMgYSBjb3VudCBvZiBzZWNvbmRzIGFuZCBmcmFjdGlvbnMgb2Ygc2Vjb25kcyBhdCBuYW5vc2Vjb25kCiByZXNvbHV0aW9uLiBJdCBpcyBpbmRlcGVuZGVudCBvZiBhbnkgY2FsZW5kYXIgYW5kIGNvbmNlcHRzIGxpa2UgImRheSIKIG9yICJtb250aCIuIEl0IGlzIHJlbGF0ZWQgdG8gVGltZXN0YW1wIGluIHRoYXQgdGhlIGRpZmZlcmVuY2UgYmV0d2VlbgogdHdvIFRpbWVzdGFtcCB2YWx1ZXMgaXMgYSBEdXJhdGlvbiBhbmQgaXQgY2FuIGJlIGFkZGVkIG9yIHN1YnRyYWN0ZWQKIGZyb20gYSBUaW1lc3RhbXAuIFJhbmdlIGlzIGFwcHJveGltYXRlbHkgKy0xMCwwMDAgeWVhcnMuCgogIyBFeGFtcGxlcwoKIEV4YW1wbGUgMTogQ29tcHV0ZSBEdXJhdGlvbiBmcm9tIHR3byBUaW1lc3RhbXBzIGluIHBzZXVkbyBjb2RlLgoKICAgICBUaW1lc3RhbXAgc3RhcnQgPSAuLi47CiAgICAgVGltZXN0YW1wIGVuZCA9IC4uLjsKICAgICBEdXJhdGlvbiBkdXJhdGlvbiA9IC4uLjsKCiAgICAgZHVyYXRpb24uc2Vjb25kcyA9IGVuZC5zZWNvbmRzIC0gc3RhcnQuc2Vjb25kczsKICAgICBkdXJhdGlvbi5uYW5vcyA9IGVuZC5uYW5vcyAtIHN0YXJ0Lm5hbm9zOwoKICAgICBpZiAoZHVyYXRpb24uc2Vjb25kcyA8IDAgJiYgZHVyYXRpb24ubmFub3MgPiAwKSB7CiAgICAgICBkdXJhdGlvbi5zZWNvbmRzICs9IDE7CiAgICAgICBkdXJhdGlvbi5uYW5vcyAtPSAxMDAwMDAwMDAwOwogICAgIH0gZWxzZSBpZiAoZHVyYXRpb25zLnNlY29uZHMgPiAwICYmIGR1cmF0aW9uLm5hbm9zIDwgMCkgewogICAgICAgZHVyYXRpb24uc2Vjb25kcyAtPSAxOwogICAgICAgZHVyYXRpb24ubmFub3MgKz0gMTAwMDAwMDAwMDsKICAgICB9CgogRXhhbXBsZSAyOiBDb21wdXRlIFRpbWVzdGFtcCBmcm9tIFRpbWVzdGFtcCArIER1cmF0aW9uIGluIHBzZXVkbyBjb2RlLgoKICAgICBUaW1lc3RhbXAgc3RhcnQgPSAuLi47CiAgICAgRHVyYXRpb24gZHVyYXRpb24gPSAuLi47CiAgICAgVGltZXN0YW1wIGVuZCA9IC4uLjsKCiAgICAgZW5kLnNlY29uZHMgPSBzdGFydC5zZWNvbmRzICsgZHVyYXRpb24uc2Vjb25kczsKICAgICBlbmQubmFub3MgPSBzdGFydC5uYW5vcyArIGR1cmF0aW9uLm5hbm9zOwoKICAgICBpZiAoZW5kLm5hbm9zIDwgMCkgewogICAgICAgZW5kLnNlY29uZHMgLT0gMTsKICAgICAgIGVuZC5uYW5vcyArPSAxMDAwMDAwMDAwOwogICAgIH0gZWxzZSBpZiAoZW5kLm5hbm9zID49IDEwMDAwMDAwMDApIHsKICAgICAgIGVuZC5zZWNvbmRzICs9IDE7CiAgICAgICBlbmQubmFub3MgLT0gMTAwMDAwMDAwMDsKICAgICB9CgogRXhhbXBsZSAzOiBDb21wdXRlIER1cmF0aW9uIGZyb20gZGF0ZXRpbWUudGltZWRlbHRhIGluIFB5dGhvbi4KCiAgICAgdGQgPSBkYXRldGltZS50aW1lZGVsdGEoZGF5cz0zLCBtaW51dGVzPTEwKQogICAgIGR1cmF0aW9uID0gRHVyYXRpb24oKQogICAgIGR1cmF0aW9uLkZyb21UaW1lZGVsdGEodGQpCgogIyBKU09OIE1hcHBpbmcKCiBJbiBKU09OIGZvcm1hdCwgdGhlIER1cmF0aW9uIHR5cGUgaXMgZW5jb2RlZCBhcyBhIHN0cmluZyByYXRoZXIgdGhhbiBhbgogb2JqZWN0LCB3aGVyZSB0aGUgc3RyaW5nIGVuZHMgaW4gdGhlIHN1ZmZpeCAicyIgKGluZGljYXRpbmcgc2Vjb25kcykgYW5kCiBpcyBwcmVjZWRlZCBieSB0aGUgbnVtYmVyIG9mIHNlY29uZHMsIHdpdGggbmFub3NlY29uZHMgZXhwcmVzc2VkIGFzCiBmcmFjdGlvbmFsIHNlY29uZHMuIEZvciBleGFtcGxlLCAzIHNlY29uZHMgd2l0aCAwIG5hbm9zZWNvbmRzIHNob3VsZCBiZQogZW5jb2RlZCBpbiBKU09OIGZvcm1hdCBhcyAiM3MiLCB3aGlsZSAzIHNlY29uZHMgYW5kIDEgbmFub3NlY29uZCBzaG91bGQKIGJlIGV4cHJlc3NlZCBpbiBKU09OIGZvcm1hdCBhcyAiMy4wMDAwMDAwMDFzIiwgYW5kIDMgc2Vjb25kcyBhbmQgMQogbWljcm9zZWNvbmQgc2hvdWxkIGJlIGV4cHJlc3NlZCBpbiBKU09OIGZvcm1hdCBhcyAiMy4wMDAwMDFzIi4KCgoKCgoDBAABEgNmCBAK3AEKBAQAAgASA2sCFBrOASBTaWduZWQgc2Vjb25kcyBvZiB0aGUgc3BhbiBvZiB0aW1lLiBNdXN0IGJlIGZyb20gLTMxNSw1NzYsMDAwLDAwMAogdG8gKzMxNSw1NzYsMDAwLDAwMCBpbmNsdXNpdmUuIE5vdGU6IHRoZXNlIGJvdW5kcyBhcmUgY29tcHV0ZWQgZnJvbToKIDYwIHNlYy9taW4gKiA2MCBtaW4vaHIgKiAyNCBoci9kYXkgKiAzNjUuMjUgZGF5cy95ZWFyICogMTAwMDAgeWVhcnMKCg0KBQQAAgAEEgRrAmYSCgwKBQQAAgAFEgNrAgcKDAoFBAACAAESA2sIDwoMCgUEAAIAAxIDaxITCoMDCgQEAAIBEgNzAhIa9QIgU2lnbmVkIGZyYWN0aW9ucyBvZiBhIHNlY29uZCBhdCBuYW5vc2Vjb25kIHJlc29sdXRpb24gb2YgdGhlIHNwYW4KIG9mIHRpbWUuIER1cmF0aW9ucyBsZXNzIHRoYW4gb25lIHNlY29uZCBhcmUgcmVwcmVzZW50ZWQgd2l0aCBhIDAKIGBzZWNvbmRzYCBmaWVsZCBhbmQgYSBwb3NpdGl2ZSBvciBuZWdhdGl2ZSBgbmFub3NgIGZpZWxkLiBGb3IgZHVyYXRpb25zCiBvZiBvbmUgc2Vjb25kIG9yIG1vcmUsIGEgbm9uLXplcm8gdmFsdWUgZm9yIHRoZSBgbmFub3NgIGZpZWxkIG11c3QgYmUKIG9mIHRoZSBzYW1lIHNpZ24gYXMgdGhlIGBzZWNvbmRzYCBmaWVsZC4gTXVzdCBiZSBmcm9tIC05OTksOTk5LDk5OQogdG8gKzk5OSw5OTksOTk5IGluY2x1c2l2ZS4KCg0KBQQAAgEEEgRzAmsUCgwKBQQAAgEFEgNzAgcKDAoFBAACAQESA3MIDQoMCgUEAAIBAxIDcxARYgZwcm90bzMKqyIKF2dvb2dsZS9ycGMvc3RhdHVzLnByb3RvEgpnb29nbGUucnBjGhlnb29nbGUvcHJvdG9idWYvYW55LnByb3RvImYKBlN0YXR1cxISCgRjb2RlGAEgASgFUgRjb2RlEhgKB21lc3NhZ2UYAiABKAlSB21lc3NhZ2USLgoHZGV0YWlscxgDIAMoCzIULmdvb2dsZS5wcm90b2J1Zi5BbnlSB2RldGFpbHNCKgoOY29tLmdvb2dsZS5ycGNCC1N0YXR1c1Byb3RvUAFaA3JwY6ICA1JQQ0rMIAoGEgQOAFsBCr0ECgEMEgMOABIysgQgQ29weXJpZ2h0IDIwMTcgR29vZ2xlIEluYy4KCiBMaWNlbnNlZCB1bmRlciB0aGUgQXBhY2hlIExpY2Vuc2UsIFZlcnNpb24gMi4wICh0aGUgIkxpY2Vuc2UiKTsKIHlvdSBtYXkgbm90IHVzZSB0aGlzIGZpbGUgZXhjZXB0IGluIGNvbXBsaWFuY2Ugd2l0aCB0aGUgTGljZW5zZS4KIFlvdSBtYXkgb2J0YWluIGEgY29weSBvZiB0aGUgTGljZW5zZSBhdAoKICAgICBodHRwOi8vd3d3LmFwYWNoZS5vcmcvbGljZW5zZXMvTElDRU5TRS0yLjAKCiBVbmxlc3MgcmVxdWlyZWQgYnkgYXBwbGljYWJsZSBsYXcgb3IgYWdyZWVkIHRvIGluIHdyaXRpbmcsIHNvZnR3YXJlCiBkaXN0cmlidXRlZCB1bmRlciB0aGUgTGljZW5zZSBpcyBkaXN0cmlidXRlZCBvbiBhbiAiQVMgSVMiIEJBU0lTLAogV0lUSE9VVCBXQVJSQU5USUVTIE9SIENPTkRJVElPTlMgT0YgQU5ZIEtJTkQsIGVpdGhlciBleHByZXNzIG9yIGltcGxpZWQuCiBTZWUgdGhlIExpY2Vuc2UgZm9yIHRoZSBzcGVjaWZpYyBsYW5ndWFnZSBnb3Zlcm5pbmcgcGVybWlzc2lvbnMgYW5kCiBsaW1pdGF0aW9ucyB1bmRlciB0aGUgTGljZW5zZS4KCggKAQISAxAIEgoJCgIDABIDEgciCggKAQgSAxQAGgoLCgQI5wcAEgMUABoKDAoFCOcHAAISAxQHEQoNCgYI5wcAAgASAxQHEQoOCgcI5wcAAgABEgMUBxEKDAoFCOcHAAcSAxQUGQoICgEIEgMVACIKCwoECOcHARIDFQAiCgwKBQjnBwECEgMVBxoKDQoGCOcHAQIAEgMVBxoKDgoHCOcHAQIAARIDFQcaCgwKBQjnBwEDEgMVHSEKCAoBCBIDFgAsCgsKBAjnBwISAxYALAoMCgUI5wcCAhIDFgcbCg0KBgjnBwICABIDFgcbCg4KBwjnBwICAAESAxYHGwoMCgUI5wcCBxIDFh4rCggKAQgSAxcAJwoLCgQI5wcDEgMXACcKDAoFCOcHAwISAxcHEwoNCgYI5wcDAgASAxcHEwoOCgcI5wcDAgABEgMXBxMKDAoFCOcHAwcSAxcWJgoICgEIEgMYACEKCwoECOcHBBIDGAAhCgwKBQjnBwQCEgMYBxgKDQoGCOcHBAIAEgMYBxgKDgoHCOcHBAIAARIDGAcYCgwKBQjnBwQHEgMYGyAKzRMKAgQAEgRPAFsBGsATIFRoZSBgU3RhdHVzYCB0eXBlIGRlZmluZXMgYSBsb2dpY2FsIGVycm9yIG1vZGVsIHRoYXQgaXMgc3VpdGFibGUgZm9yIGRpZmZlcmVudAogcHJvZ3JhbW1pbmcgZW52aXJvbm1lbnRzLCBpbmNsdWRpbmcgUkVTVCBBUElzIGFuZCBSUEMgQVBJcy4gSXQgaXMgdXNlZCBieQogW2dSUENdKGh0dHBzOi8vZ2l0aHViLmNvbS9ncnBjKS4gVGhlIGVycm9yIG1vZGVsIGlzIGRlc2lnbmVkIHRvIGJlOgoKIC0gU2ltcGxlIHRvIHVzZSBhbmQgdW5kZXJzdGFuZCBmb3IgbW9zdCB1c2VycwogLSBGbGV4aWJsZSBlbm91Z2ggdG8gbWVldCB1bmV4cGVjdGVkIG5lZWRzCgogIyBPdmVydmlldwoKIFRoZSBgU3RhdHVzYCBtZXNzYWdlIGNvbnRhaW5zIHRocmVlIHBpZWNlcyBvZiBkYXRhOiBlcnJvciBjb2RlLCBlcnJvciBtZXNzYWdlLAogYW5kIGVycm9yIGRldGFpbHMuIFRoZSBlcnJvciBjb2RlIHNob3VsZCBiZSBhbiBlbnVtIHZhbHVlIG9mCiBbZ29vZ2xlLnJwYy5Db2RlXVtnb29nbGUucnBjLkNvZGVdLCBidXQgaXQgbWF5IGFjY2VwdCBhZGRpdGlvbmFsIGVycm9yIGNvZGVzIGlmIG5lZWRlZC4gIFRoZQogZXJyb3IgbWVzc2FnZSBzaG91bGQgYmUgYSBkZXZlbG9wZXItZmFjaW5nIEVuZ2xpc2ggbWVzc2FnZSB0aGF0IGhlbHBzCiBkZXZlbG9wZXJzICp1bmRlcnN0YW5kKiBhbmQgKnJlc29sdmUqIHRoZSBlcnJvci4gSWYgYSBsb2NhbGl6ZWQgdXNlci1mYWNpbmcKIGVycm9yIG1lc3NhZ2UgaXMgbmVlZGVkLCBwdXQgdGhlIGxvY2FsaXplZCBtZXNzYWdlIGluIHRoZSBlcnJvciBkZXRhaWxzIG9yCiBsb2NhbGl6ZSBpdCBpbiB0aGUgY2xpZW50LiBUaGUgb3B0aW9uYWwgZXJyb3IgZGV0YWlscyBtYXkgY29udGFpbiBhcmJpdHJhcnkKIGluZm9ybWF0aW9uIGFib3V0IHRoZSBlcnJvci4gVGhlcmUgaXMgYSBwcmVkZWZpbmVkIHNldCBvZiBlcnJvciBkZXRhaWwgdHlwZXMKIGluIHRoZSBwYWNrYWdlIGBnb29nbGUucnBjYCB0aGF0IGNhbiBiZSB1c2VkIGZvciBjb21tb24gZXJyb3IgY29uZGl0aW9ucy4KCiAjIExhbmd1YWdlIG1hcHBpbmcKCiBUaGUgYFN0YXR1c2AgbWVzc2FnZSBpcyB0aGUgbG9naWNhbCByZXByZXNlbnRhdGlvbiBvZiB0aGUgZXJyb3IgbW9kZWwsIGJ1dCBpdAogaXMgbm90IG5lY2Vzc2FyaWx5IHRoZSBhY3R1YWwgd2lyZSBmb3JtYXQuIFdoZW4gdGhlIGBTdGF0dXNgIG1lc3NhZ2UgaXMKIGV4cG9zZWQgaW4gZGlmZmVyZW50IGNsaWVudCBsaWJyYXJpZXMgYW5kIGRpZmZlcmVudCB3aXJlIHByb3RvY29scywgaXQgY2FuIGJlCiBtYXBwZWQgZGlmZmVyZW50bHkuIEZvciBleGFtcGxlLCBpdCB3aWxsIGxpa2VseSBiZSBtYXBwZWQgdG8gc29tZSBleGNlcHRpb25zCiBpbiBKYXZhLCBidXQgbW9yZSBsaWtlbHkgbWFwcGVkIHRvIHNvbWUgZXJyb3IgY29kZXMgaW4gQy4KCiAjIE90aGVyIHVzZXMKCiBUaGUgZXJyb3IgbW9kZWwgYW5kIHRoZSBgU3RhdHVzYCBtZXNzYWdlIGNhbiBiZSB1c2VkIGluIGEgdmFyaWV0eSBvZgogZW52aXJvbm1lbnRzLCBlaXRoZXIgd2l0aCBvciB3aXRob3V0IEFQSXMsIHRvIHByb3ZpZGUgYQogY29uc2lzdGVudCBkZXZlbG9wZXIgZXhwZXJpZW5jZSBhY3Jvc3MgZGlmZmVyZW50IGVudmlyb25tZW50cy4KCiBFeGFtcGxlIHVzZXMgb2YgdGhpcyBlcnJvciBtb2RlbCBpbmNsdWRlOgoKIC0gUGFydGlhbCBlcnJvcnMuIElmIGEgc2VydmljZSBuZWVkcyB0byByZXR1cm4gcGFydGlhbCBlcnJvcnMgdG8gdGhlIGNsaWVudCwKICAgICBpdCBtYXkgZW1iZWQgdGhlIGBTdGF0dXNgIGluIHRoZSBub3JtYWwgcmVzcG9uc2UgdG8gaW5kaWNhdGUgdGhlIHBhcnRpYWwKICAgICBlcnJvcnMuCgogLSBXb3JrZmxvdyBlcnJvcnMuIEEgdHlwaWNhbCB3b3JrZmxvdyBoYXMgbXVsdGlwbGUgc3RlcHMuIEVhY2ggc3RlcCBtYXkKICAgICBoYXZlIGEgYFN0YXR1c2AgbWVzc2FnZSBmb3IgZXJyb3IgcmVwb3J0aW5nLgoKIC0gQmF0Y2ggb3BlcmF0aW9ucy4gSWYgYSBjbGllbnQgdXNlcyBiYXRjaCByZXF1ZXN0IGFuZCBiYXRjaCByZXNwb25zZSwgdGhlCiAgICAgYFN0YXR1c2AgbWVzc2FnZSBzaG91bGQgYmUgdXNlZCBkaXJlY3RseSBpbnNpZGUgYmF0Y2ggcmVzcG9uc2UsIG9uZSBmb3IKICAgICBlYWNoIGVycm9yIHN1Yi1yZXNwb25zZS4KCiAtIEFzeW5jaHJvbm91cyBvcGVyYXRpb25zLiBJZiBhbiBBUEkgY2FsbCBlbWJlZHMgYXN5bmNocm9ub3VzIG9wZXJhdGlvbgogICAgIHJlc3VsdHMgaW4gaXRzIHJlc3BvbnNlLCB0aGUgc3RhdHVzIG9mIHRob3NlIG9wZXJhdGlvbnMgc2hvdWxkIGJlCiAgICAgcmVwcmVzZW50ZWQgZGlyZWN0bHkgdXNpbmcgdGhlIGBTdGF0dXNgIG1lc3NhZ2UuCgogLSBMb2dnaW5nLiBJZiBzb21lIEFQSSBlcnJvcnMgYXJlIHN0b3JlZCBpbiBsb2dzLCB0aGUgbWVzc2FnZSBgU3RhdHVzYCBjb3VsZAogICAgIGJlIHVzZWQgZGlyZWN0bHkgYWZ0ZXIgYW55IHN0cmlwcGluZyBuZWVkZWQgZm9yIHNlY3VyaXR5L3ByaXZhY3kgcmVhc29ucy4KCgoKAwQAARIDTwgOCmQKBAQAAgASA1ECERpXIFRoZSBzdGF0dXMgY29kZSwgd2hpY2ggc2hvdWxkIGJlIGFuIGVudW0gdmFsdWUgb2YgW2dvb2dsZS5ycGMuQ29kZV1bZ29vZ2xlLnJwYy5Db2RlXS4KCg0KBQQAAgAEEgRRAk8QCgwKBQQAAgAFEgNRAgcKDAoFBAACAAESA1EIDAoMCgUEAAIAAxIDUQ8QCusBCgQEAAIBEgNWAhUa3QEgQSBkZXZlbG9wZXItZmFjaW5nIGVycm9yIG1lc3NhZ2UsIHdoaWNoIHNob3VsZCBiZSBpbiBFbmdsaXNoLiBBbnkKIHVzZXItZmFjaW5nIGVycm9yIG1lc3NhZ2Ugc2hvdWxkIGJlIGxvY2FsaXplZCBhbmQgc2VudCBpbiB0aGUKIFtnb29nbGUucnBjLlN0YXR1cy5kZXRhaWxzXVtnb29nbGUucnBjLlN0YXR1cy5kZXRhaWxzXSBmaWVsZCwgb3IgbG9jYWxpemVkIGJ5IHRoZSBjbGllbnQuCgoNCgUEAAIBBBIEVgJREQoMCgUEAAIBBRIDVgIICgwKBQQAAgEBEgNWCRAKDAoFBAACAQMSA1YTFAp5CgQEAAICEgNaAisabCBBIGxpc3Qgb2YgbWVzc2FnZXMgdGhhdCBjYXJyeSB0aGUgZXJyb3IgZGV0YWlscy4gIFRoZXJlIGlzIGEgY29tbW9uIHNldCBvZgogbWVzc2FnZSB0eXBlcyBmb3IgQVBJcyB0byB1c2UuCgoMCgUEAAICBBIDWgIKCgwKBQQAAgIGEgNaCx4KDAoFBAACAgESA1ofJgoMCgUEAAICAxIDWikqYgZwcm90bzMKvBEKJ21peGVyL2FkYXB0ZXIvbW9kZWwvdjFiZXRhMS9jaGVjay5wcm90bxIhaXN0aW8ubWl4ZXIuYWRhcHRlci5tb2RlbC52MWJldGExGhRnb2dvcHJvdG8vZ29nby5wcm90bxoeZ29vZ2xlL3Byb3RvYnVmL2R1cmF0aW9uLnByb3RvGhdnb29nbGUvcnBjL3N0YXR1cy5wcm90byKzAQoLQ2hlY2tSZXN1bHQSMAoGc3RhdHVzGAEgASgLMhIuZ29vZ2xlLnJwYy5TdGF0dXNCBMjeHwBSBnN0YXR1cxJKCg52YWxpZF9kdXJhdGlvbhgCIAEoCzIZLmdvb2dsZS5wcm90b2J1Zi5EdXJhdGlvbkIIyN4fAJjfHwFSDXZhbGlkRHVyYXRpb24SJgoPdmFsaWRfdXNlX2NvdW50GAMgASgFUg12YWxpZFVzZUNvdW50QjZaKGlzdGlvLmlvL2FwaS9taXhlci9hZGFwdGVyL21vZGVsL3YxYmV0YTHI4R4AqOIe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CgUEAAICBRIDJAIHCgwKBQQAAgIBEgMkCBcKDAoFBAACAgMSAyQaG2IGcHJvdG8zCoEhCjBtaXhlci90ZXN0L2tleXZhbC90ZW1wbGF0ZV9oYW5kbGVyX3NlcnZpY2UucHJvdG8SBmtleXZhbBoUZ29nb3Byb3RvL2dvZ28ucHJvdG8aLG1peGVyL2FkYXB0ZXIvbW9kZWwvdjFiZXRhMS9leHRlbnNpb25zLnByb3RvGhlnb29nbGUvcHJvdG9idWYvYW55LnByb3RvGidtaXhlci9hZGFwdGVyL21vZGVsL3YxYmV0YTEvY2hlY2sucHJvdG8ingEKE0hhbmRsZUtleXZhbFJlcXVlc3QSLwoIaW5zdGFuY2UYASABKAsyEy5rZXl2YWwuSW5zdGFuY2VNc2dSCGluc3RhbmNlEjsKDmFkYXB0ZXJfY29uZmlnGAIgASgLMhQuZ29vZ2xlLnByb3RvYnVmLkFueVINYWRhcHRlckNvbmZpZxIZCghkZWR1cF9pZBgDIAEoCVIHZGVkdXBJZCKJAQoUSGFuZGxlS2V5dmFsUmVzcG9uc2USRgoGcmVzdWx0GAEgASgLMi4uaXN0aW8ubWl4ZXIuYWRhcHRlci5tb2RlbC52MWJldGExLkNoZWNrUmVzdWx0UgZyZXN1bHQSKQoGb3V0cHV0GAIgASgLMhEua2V5dmFsLk91dHB1dE1zZ1IGb3V0cHV0IiEKCU91dHB1dE1zZxIUCgV2YWx1ZRgCIAEoCVIFdmFsdWUiNgoLSW5zdGFuY2VNc2cSFQoEbmFtZRivyrwiIAEoCVIEbmFtZRIQCgNrZXkYASABKAlSA2tleSIGCgRUeXBlIiEKDUluc3RhbmNlUGFyYW0SEAoDa2V5GAEgASgJUgNrZXkyYAoTSGFuZGxlS2V5dmFsU2VydmljZRJJCgxIYW5kbGVLZXl2YWwSGy5rZXl2YWwuSGFuZGxlS2V5dmFsUmVxdWVzdBocLmtleXZhbC5IYW5kbGVLZXl2YWxSZXNwb25zZUIe+NLkkwIEgt3kkwIGa2V5dmFsyOEeAKjiHgDw4R4ASv8aCgYSBBAAZQEK8gQKAQwSAxAAEjK0BCBDb3B5cmlnaHQgMjAxNyBJc3RpbyBBdXRob3JzCgogTGljZW5zZWQgdW5kZXIgdGhlIEFwYWNoZSBMaWNlbnNlLCBWZXJzaW9uIDIuMCAodGhlICJMaWNlbnNlIik7CiB5b3UgbWF5IG5vdCB1c2UgdGhpcyBmaWxlIGV4Y2VwdCBpbiBjb21wbGlhbmNlIHdpdGggdGhlIExpY2Vuc2UuCiBZb3UgbWF5IG9idGFpbiBhIGNvcHkgb2YgdGhlIExpY2Vuc2UgYXQKCiAgICAgaHR0cDovL3d3dy5hcGFjaGUub3JnL2xpY2Vuc2VzL0xJQ0VOU0UtMi4wCgogVW5sZXNzIHJlcXVpcmVkIGJ5IGFwcGxpY2FibGUgbGF3IG9yIGFncmVlZCB0byBpbiB3cml0aW5nLCBzb2Z0d2FyZQogZGlzdHJpYnV0ZWQgdW5kZXIgdGhlIExpY2Vuc2UgaXMgZGlzdHJpYnV0ZWQgb24gYW4gIkFTIElTIiBCQVNJUywKIFdJVEhPVVQgV0FSUkFOVElFUyBPUiBDT05ESVRJT05TIE9GIEFOWSBLSU5ELCBlaXRoZXIgZXhwcmVzcyBvciBpbXBsaWVkLgogU2VlIHRoZSBMaWNlbnNlIGZvciB0aGUgc3BlY2lmaWMgbGFuZ3VhZ2UgZ292ZXJuaW5nIHBlcm1pc3Npb25zIGFuZAogbGltaXRhdGlvbnMgdW5kZXIgdGhlIExpY2Vuc2UuCjIxIFRISVMgRklMRSBJUyBBVVRPTUFUSUNBTExZIEdFTkVSQVRFRCBCWSBNSVhHRU4uCgobCgECEgMWCA4aEQogSW5wdXQgdGVtcGxhdGUKCgkKAgMAEgMZBx0KCQoCAwESAxoHNQoJCgIDAhIDGwciCgkKAgMDEgMcBzAKCAoBCBIDHwBhCgsKBAjnBwASAx8AYQoMCgUI5wcAAhIDHwc7Cg0KBgjnBwACABIDHwc7Cg4KBwjnBwACAAESAx8IOgoMCgUI5wcAAxIDHz5gCggKAQgSAyAARAoLCgQI5wcBEgMgAEQKDAoFCOcHAQISAyAHOAoNCgYI5wcBAgASAyAHOAoOCgcI5wcBAgABEgMgCDcKDAoFCOcHAQcSAyA7QwoICgEIEgMiAC8KCwoECOcHAhIDIgAvCgwKBQjnBwICEgMiByYKDQoGCOcHAgIAEgMiByYKDgoHCOcHAgIAARIDIgglCgwKBQjnBwIDEgMiKS4KCAoBCBIDIwAlCgsKBAjnBwMSAyMAJQoMCgUI5wcDAhIDIwccCg0KBgjnBwMCABIDIwccCg4KBwjnBwMCAAESAyMIGwoMCgUI5wcDAxIDIx8kCggKAQgSAyQAKAoLCgQI5wcEEgMkACgKDAoFCOcHBAISAyQHHwoNCgYI5wcEAgASAyQHHwoOCgcI5wcEAgABEgMkCB4KDAoFCOcHBAMSAyQiJwpyCgIGABIEJwArARpmIEhhbmRsZUtleXZhbFNlcnZpY2UgaXMgaW1wbGVtZW50ZWQgYnkgYmFja2VuZHMgdGhhdCB3YW50cyB0byBoYW5kbGUgcmVxdWVzdC10aW1lICdrZXl2YWwnIGluc3RhbmNlcy4KCgoKAwYAARIDJwgbCmwKBAYAAgASAykESRpfIEhhbmRsZUtleXZhbCBpcyBjYWxsZWQgYnkgTWl4ZXIgYXQgcmVxdWVzdC10aW1lIHRvIGRlbGl2ZXIgJ2tleXZhbCcgaW5zdGFuY2VzIHRvIHRoZSBiYWNrZW5kLgoKDAoFBgACAAESAykIFAoMCgUGAAIAAhIDKRUoCgwKBQYAAgADEgMpM0cKNgoCBAASBC4APQEaKiBSZXF1ZXN0IG1lc3NhZ2UgZm9yIEhhbmRsZUtleXZhbCBtZXRob2QuCgoKCgMEAAESAy4IGwohCgQEAAIAEgMxBB0aFCAna2V5dmFsJyBpbnN0YW5jZS4KCg0KBQQAAgAEEgQxBC4dCgwKBQQAAgAGEgMxBA8KDAoFBAACAAESAzEQGAoMCgUEAAIAAxIDMRscCrkECgQEAAIBEgM5BCsaqwQgQWRhcHRlciBzcGVjaWZpYyBoYW5kbGVyIGNvbmZpZ3VyYXRpb24uCgogTm90ZTogQmFja2VuZHMgY2FuIGFsc28gaW1wbGVtZW50IFtJbmZyYXN0cnVjdHVyZUJhY2tlbmRdW2h0dHBzOi8vaXN0aW8uaW8vZG9jcy9yZWZlcmVuY2UvY29uZmlnL21peGVyL2lzdGlvLm1peGVyLmFkYXB0ZXIubW9kZWwudjFiZXRhMS5odG1sI0luZnJhc3RydWN0dXJlQmFja2VuZF0KIHNlcnZpY2UgYW5kIHRoZXJlZm9yZSBvcHQgdG8gcmVjZWl2ZSBoYW5kbGVyIGNvbmZpZ3VyYXRpb24gZHVyaW5nIHNlc3Npb24gY3JlYXRpb24gdGhyb3VnaCBbSW5mcmFzdHJ1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"
`)

func datasetConfigV1alpha2TemplateValidYamlBytes() ([]byte, error) {
	return _datasetConfigV1alpha2TemplateValidYaml, nil
}

func datasetConfigV1alpha2TemplateValidYaml() (*asset, error) {
	bytes, err := datasetConfigV1alpha2TemplateValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/config-v1alpha2-template-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1alpha3DestinationruleInvalidYaml = []byte(`apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: invalid-destination-rule
spec:
  subsets:
    - name: v1
      labels:
        version: v1
    - name: v2
      labels:
        version: v2
`)

func datasetNetworkingV1alpha3DestinationruleInvalidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1alpha3DestinationruleInvalidYaml, nil
}

func datasetNetworkingV1alpha3DestinationruleInvalidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1alpha3DestinationruleInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1alpha3-DestinationRule-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1alpha3DestinationruleValidYaml = []byte(`apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: valid-destination-rule
spec:
  host: c
  subsets:
    - name: v1
      labels:
        version: v1
    - name: v2
      labels:
        version: v2
`)

func datasetNetworkingV1alpha3DestinationruleValidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1alpha3DestinationruleValidYaml, nil
}

func datasetNetworkingV1alpha3DestinationruleValidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1alpha3DestinationruleValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1alpha3-DestinationRule-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1alpha3EnvoyfilterInvalidYaml = []byte(`apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: simple-envoy-filter
spec:
  workloadSelector:
    labels:
`)

func datasetNetworkingV1alpha3EnvoyfilterInvalidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1alpha3EnvoyfilterInvalidYaml, nil
}

func datasetNetworkingV1alpha3EnvoyfilterInvalidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1alpha3EnvoyfilterInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1alpha3-EnvoyFilter-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1alpha3EnvoyfilterValidYaml = []byte(`apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: simple-envoy-filter
spec:
  workloadLabels:
    app: c
  filters:
    - listenerMatch:
        listenerType: SIDECAR_INBOUND
        listenerProtocol: HTTP
      insertPosition:
        index: FIRST
      filterType: HTTP
      filterName: envoy.fault
      filterConfig:
        abort:
          percentage:
            numerator: 100
            denominator: HUNDRED
          httpStatus: 444
        headers:
          name: envoyfilter-test
          exactMatch: foobar123
`)

func datasetNetworkingV1alpha3EnvoyfilterValidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1alpha3EnvoyfilterValidYaml, nil
}

func datasetNetworkingV1alpha3EnvoyfilterValidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1alpha3EnvoyfilterValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1alpha3-EnvoyFilter-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1alpha3GatewayInvalidYaml = []byte(`# Routes TCP traffic through the ingressgateway Gateway to service A.
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: invalid-gateway
spec:
  selector:
    # DO NOT CHANGE THESE LABELS
    # The ingressgateway is defined in install/kubernetes/helm/istio/values.yaml
    # with these labels
    istio: ingressgateway
`)

func datasetNetworkingV1alpha3GatewayInvalidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1alpha3GatewayInvalidYaml, nil
}

func datasetNetworkingV1alpha3GatewayInvalidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1alpha3GatewayInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1alpha3-Gateway-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1alpha3GatewayValidYaml = []byte(`# Routes TCP traffic through the ingressgateway Gateway to service A.
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: valid-gateway
spec:
  selector:
    # DO NOT CHANGE THESE LABELS
    # The ingressgateway is defined in install/kubernetes/helm/istio/values.yaml
    # with these labels
    istio: ingressgateway
  servers:
  - port:
      number: 31400
      protocol: TCP
      name: tcp
    hosts:
    - a.istio-system.svc.cluster.local
`)

func datasetNetworkingV1alpha3GatewayValidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1alpha3GatewayValidYaml, nil
}

func datasetNetworkingV1alpha3GatewayValidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1alpha3GatewayValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1alpha3-Gateway-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1alpha3ServiceentryInvalidYaml = []byte(`apiVersion: networking.istio.io/v1alpha3
kind: ServiceEntry
metadata:
  name: invalid-service-entry
spec:
  ports:
  - number: 80
    name: http
    protocol: HTTP
  discovery: DNS
  endpoints:
  # Rather than relying on an external host that might become unreachable (causing test failures)
  # we can mock the external endpoint using service t which has no sidecar.
  - address: t.istio-system.svc.cluster.local # TODO: this is brittle
    ports:
      http: 8080 # TODO test https
`)

func datasetNetworkingV1alpha3ServiceentryInvalidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1alpha3ServiceentryInvalidYaml, nil
}

func datasetNetworkingV1alpha3ServiceentryInvalidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1alpha3ServiceentryInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1alpha3-ServiceEntry-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1alpha3ServiceentryValidYaml = []byte(`apiVersion: networking.istio.io/v1alpha3
kind: ServiceEntry
metadata:
  name: valid-service-entry
spec:
  hosts:
  - eu.bookinfo.com
  ports:
  - number: 80
    name: http
    protocol: HTTP
  resolution: DNS
  endpoints:
  # Rather than relying on an external host that might become unreachable (causing test failures)
  # we can mock the external endpoint using service t which has no sidecar.
  - address: t.istio-system.svc.cluster.local # TODO: this is brittle
    ports:
      http: 8080 # TODO test https
`)

func datasetNetworkingV1alpha3ServiceentryValidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1alpha3ServiceentryValidYaml, nil
}

func datasetNetworkingV1alpha3ServiceentryValidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1alpha3ServiceentryValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1alpha3-ServiceEntry-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1alpha3SidecarInvalidYaml = []byte(`apiVersion: networking.istio.io/v1alpha3
kind: Sidecar
metadata:
  name: invalid-sidecar-config
spec:
  egress:
`)

func datasetNetworkingV1alpha3SidecarInvalidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1alpha3SidecarInvalidYaml, nil
}

func datasetNetworkingV1alpha3SidecarInvalidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1alpha3SidecarInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1alpha3-Sidecar-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1alpha3SidecarValidYaml = []byte(`apiVersion: networking.istio.io/v1alpha3
kind: Sidecar
metadata:
  name: valid-sidecar-config
spec:
  egress:
  - hosts:
    - "abc/*"
`)

func datasetNetworkingV1alpha3SidecarValidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1alpha3SidecarValidYaml, nil
}

func datasetNetworkingV1alpha3SidecarValidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1alpha3SidecarValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1alpha3-Sidecar-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1alpha3VirtualserviceInvalidYaml = []byte(`apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: invalid-virtual-service
spec:
  http:
    - route:
      - destination:
          host: c
          subset: v1
        weight: 75
      - destination:
          host: c
          subset: v2
        weight: 25
`)

func datasetNetworkingV1alpha3VirtualserviceInvalidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1alpha3VirtualserviceInvalidYaml, nil
}

func datasetNetworkingV1alpha3VirtualserviceInvalidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1alpha3VirtualserviceInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1alpha3-VirtualService-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1alpha3VirtualserviceValidYaml = []byte(`apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: valid-virtual-service
spec:
  hosts:
    - c
  http:
    - route:
      - destination:
          host: c
          subset: v1
        weight: 75
      - destination:
          host: c
          subset: v2
        weight: 25
`)

func datasetNetworkingV1alpha3VirtualserviceValidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1alpha3VirtualserviceValidYaml, nil
}

func datasetNetworkingV1alpha3VirtualserviceValidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1alpha3VirtualserviceValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1alpha3-VirtualService-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1betaDestinationruleInvalidYaml = []byte(`apiVersion: networking.istio.io/v1beta1
kind: DestinationRule
metadata:
  name: invalid-destination-rule
spec:
  subsets:
    - name: v1
      labels:
        version: v1
    - name: v2
      labels:
        version: v2
`)

func datasetNetworkingV1betaDestinationruleInvalidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1betaDestinationruleInvalidYaml, nil
}

func datasetNetworkingV1betaDestinationruleInvalidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1betaDestinationruleInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1beta-DestinationRule-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1betaDestinationruleValidYaml = []byte(`apiVersion: networking.istio.io/v1beta1
kind: DestinationRule
metadata:
  name: valid-destination-rule
spec:
  host: c
  subsets:
    - name: v1
      labels:
        version: v1
    - name: v2
      labels:
        version: v2
`)

func datasetNetworkingV1betaDestinationruleValidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1betaDestinationruleValidYaml, nil
}

func datasetNetworkingV1betaDestinationruleValidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1betaDestinationruleValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1beta-DestinationRule-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1betaGatewayInvalidYaml = []byte(`# Routes TCP traffic through the ingressgateway Gateway to service A.
apiVersion: networking.istio.io/v1beta1
kind: Gateway
metadata:
  name: invalid-gateway
spec:
  selector:
    # DO NOT CHANGE THESE LABELS
    # The ingressgateway is defined in install/kubernetes/helm/istio/values.yaml
    # with these labels
    istio: ingressgateway
`)

func datasetNetworkingV1betaGatewayInvalidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1betaGatewayInvalidYaml, nil
}

func datasetNetworkingV1betaGatewayInvalidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1betaGatewayInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1beta-Gateway-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1betaGatewayValidYaml = []byte(`# Routes TCP traffic through the ingressgateway Gateway to service A.
apiVersion: networking.istio.io/v1beta1
kind: Gateway
metadata:
  name: valid-gateway
spec:
  selector:
    # DO NOT CHANGE THESE LABELS
    # The ingressgateway is defined in install/kubernetes/helm/istio/values.yaml
    # with these labels
    istio: ingressgateway
  servers:
  - port:
      number: 31400
      protocol: TCP
      name: tcp
    hosts:
    - a.istio-system.svc.cluster.local
`)

func datasetNetworkingV1betaGatewayValidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1betaGatewayValidYaml, nil
}

func datasetNetworkingV1betaGatewayValidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1betaGatewayValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1beta-Gateway-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1betaSidecarInvalidYaml = []byte(`apiVersion: networking.istio.io/v1beta1
kind: Sidecar
metadata:
  name: invalid-sidecar-config
spec:
  egress:
`)

func datasetNetworkingV1betaSidecarInvalidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1betaSidecarInvalidYaml, nil
}

func datasetNetworkingV1betaSidecarInvalidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1betaSidecarInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1beta-Sidecar-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1betaSidecarValidYaml = []byte(`apiVersion: networking.istio.io/v1beta1
kind: Sidecar
metadata:
  name: valid-sidecar-config
spec:
  egress:
  - hosts:
    - "abc/*"
`)

func datasetNetworkingV1betaSidecarValidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1betaSidecarValidYaml, nil
}

func datasetNetworkingV1betaSidecarValidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1betaSidecarValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1beta-Sidecar-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1betaVirtualserviceInvalidYaml = []byte(`apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: invalid-virtual-service
spec:
  http:
    - route:
      - destination:
          host: c
          subset: v1
        weight: 75
      - destination:
          host: c
          subset: v2
        weight: 25
`)

func datasetNetworkingV1betaVirtualserviceInvalidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1betaVirtualserviceInvalidYaml, nil
}

func datasetNetworkingV1betaVirtualserviceInvalidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1betaVirtualserviceInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1beta-VirtualService-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetNetworkingV1betaVirtualserviceValidYaml = []byte(`apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: valid-virtual-service
spec:
  hosts:
    - c
  http:
    - route:
      - destination:
          host: c
          subset: v1
        weight: 75
      - destination:
          host: c
          subset: v2
        weight: 25
`)

func datasetNetworkingV1betaVirtualserviceValidYamlBytes() ([]byte, error) {
	return _datasetNetworkingV1betaVirtualserviceValidYaml, nil
}

func datasetNetworkingV1betaVirtualserviceValidYaml() (*asset, error) {
	bytes, err := datasetNetworkingV1betaVirtualserviceValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/networking-v1beta-VirtualService-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetRbacV1alpha1ClusterrbacconfigInvalidYaml = []byte(`apiVersion: "rbac.istio.io/v1alpha1"
kind: ClusterRbacConfig
metadata:
  name: default
spec:
  mode: 'ON_WITH_EXCLUSION'
`)

func datasetRbacV1alpha1ClusterrbacconfigInvalidYamlBytes() ([]byte, error) {
	return _datasetRbacV1alpha1ClusterrbacconfigInvalidYaml, nil
}

func datasetRbacV1alpha1ClusterrbacconfigInvalidYaml() (*asset, error) {
	bytes, err := datasetRbacV1alpha1ClusterrbacconfigInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/rbac-v1alpha1-ClusterRbacConfig-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetRbacV1alpha1ClusterrbacconfigValidYaml = []byte(`apiVersion: "rbac.istio.io/v1alpha1"
kind: ClusterRbacConfig
metadata:
  name: default
spec:
  mode: 'ON_WITH_INCLUSION'
  inclusion:
    services: ["mongodb.default.svc.cluster.local"]
`)

func datasetRbacV1alpha1ClusterrbacconfigValidYamlBytes() ([]byte, error) {
	return _datasetRbacV1alpha1ClusterrbacconfigValidYaml, nil
}

func datasetRbacV1alpha1ClusterrbacconfigValidYaml() (*asset, error) {
	bytes, err := datasetRbacV1alpha1ClusterrbacconfigValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/rbac-v1alpha1-ClusterRbacConfig-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetRbacV1alpha1RbacconfigInvalidYaml = []byte(`apiVersion: "rbac.istio.io/v1alpha1"
kind: RbacConfig
metadata:
  name: default
spec:
  mode: 'ON_WITH_EXCLUSION'
`)

func datasetRbacV1alpha1RbacconfigInvalidYamlBytes() ([]byte, error) {
	return _datasetRbacV1alpha1RbacconfigInvalidYaml, nil
}

func datasetRbacV1alpha1RbacconfigInvalidYaml() (*asset, error) {
	bytes, err := datasetRbacV1alpha1RbacconfigInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/rbac-v1alpha1-RBacConfig-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetRbacV1alpha1RbacconfigValidYaml = []byte(`apiVersion: "rbac.istio.io/v1alpha1"
kind: RbacConfig
metadata:
  name: default
spec:
  mode: 'ON_WITH_INCLUSION'
  inclusion:
    services: ["mongodb.default.svc.cluster.local"]`)

func datasetRbacV1alpha1RbacconfigValidYamlBytes() ([]byte, error) {
	return _datasetRbacV1alpha1RbacconfigValidYaml, nil
}

func datasetRbacV1alpha1RbacconfigValidYaml() (*asset, error) {
	bytes, err := datasetRbacV1alpha1RbacconfigValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/rbac-v1alpha1-RBacConfig-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetRbacV1alpha1ServiceroleInvalidYaml = []byte(`apiVersion: "rbac.istio.io/v1alpha1"
kind: ServiceRole
metadata:
  name: products-viewer
spec:
  rules:

`)

func datasetRbacV1alpha1ServiceroleInvalidYamlBytes() ([]byte, error) {
	return _datasetRbacV1alpha1ServiceroleInvalidYaml, nil
}

func datasetRbacV1alpha1ServiceroleInvalidYaml() (*asset, error) {
	bytes, err := datasetRbacV1alpha1ServiceroleInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/rbac-v1alpha1-ServiceRole-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetRbacV1alpha1ServiceroleValidYaml = []byte(`apiVersion: "rbac.istio.io/v1alpha1"
kind: ServiceRole
metadata:
  name: products-viewer
spec:
  rules:
  - services: ["products.svc.cluster.local"]
    methods: ["GET", "HEAD"]
    constraints:
    - key: "version"
      values: ["v1", "v2"]

`)

func datasetRbacV1alpha1ServiceroleValidYamlBytes() ([]byte, error) {
	return _datasetRbacV1alpha1ServiceroleValidYaml, nil
}

func datasetRbacV1alpha1ServiceroleValidYaml() (*asset, error) {
	bytes, err := datasetRbacV1alpha1ServiceroleValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/rbac-v1alpha1-ServiceRole-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetRbacV1alpha1ServicerolebindingInvalidYaml = []byte(`apiVersion: "rbac.istio.io/v1alpha1"
kind: ServiceRoleBinding
metadata:
  name: test-binding-products
spec:
  subjects:

`)

func datasetRbacV1alpha1ServicerolebindingInvalidYamlBytes() ([]byte, error) {
	return _datasetRbacV1alpha1ServicerolebindingInvalidYaml, nil
}

func datasetRbacV1alpha1ServicerolebindingInvalidYaml() (*asset, error) {
	bytes, err := datasetRbacV1alpha1ServicerolebindingInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/rbac-v1alpha1-ServiceRoleBinding-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetRbacV1alpha1ServicerolebindingValidYaml = []byte(`apiVersion: "rbac.istio.io/v1alpha1"
kind: ServiceRoleBinding
metadata:
  name: test-binding-products
spec:
  subjects:
  - user: "alice@yahoo.com"
  - properties:
      service: "reviews"
      namespace: "abc"
  roleRef:
    kind: ServiceRole
    name: "products-viewer"

`)

func datasetRbacV1alpha1ServicerolebindingValidYamlBytes() ([]byte, error) {
	return _datasetRbacV1alpha1ServicerolebindingValidYaml, nil
}

func datasetRbacV1alpha1ServicerolebindingValidYaml() (*asset, error) {
	bytes, err := datasetRbacV1alpha1ServicerolebindingValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/rbac-v1alpha1-ServiceRoleBinding-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetSecurityV1beta1AuthorizationpolicyInvalidYaml = []byte(`apiVersion: security.istio.io/v1beta1
kind: AuthorizationPolicy
metadata:
  name: invalid-authorization-policy
spec:
  selector:
    matchLabels:
      app: httpbin
      version: v1
  rules:
  - from:
    - source:
        principals: ["cluster.local/ns/default/sa/sleep"]
    to:
    - operation:
        methods: ["GET"]
    when:
    - values: ["key is missing"]
`)

func datasetSecurityV1beta1AuthorizationpolicyInvalidYamlBytes() ([]byte, error) {
	return _datasetSecurityV1beta1AuthorizationpolicyInvalidYaml, nil
}

func datasetSecurityV1beta1AuthorizationpolicyInvalidYaml() (*asset, error) {
	bytes, err := datasetSecurityV1beta1AuthorizationpolicyInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/security-v1beta1-AuthorizationPolicy-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetSecurityV1beta1AuthorizationpolicyValidYaml = []byte(`apiVersion: security.istio.io/v1beta1
kind: AuthorizationPolicy
metadata:
 name: authorization-policy
spec:
 selector:
   matchLabels:
     app: httpbin
     version: v1
 rules:
 - from:
   - source:
       principals: ["cluster.local/ns/default/sa/sleep"]
   - source:
       namespaces: ["test"]
   to:
   - operation:
       methods: ["GET"]
       paths: ["/info*"]
   - operation:
       methods: ["POST"]
       paths: ["/data"]
   when:
   - key: request.auth.claims[iss]
     values: ["https://accounts.google.com"]
`)

func datasetSecurityV1beta1AuthorizationpolicyValidYamlBytes() ([]byte, error) {
	return _datasetSecurityV1beta1AuthorizationpolicyValidYaml, nil
}

func datasetSecurityV1beta1AuthorizationpolicyValidYaml() (*asset, error) {
	bytes, err := datasetSecurityV1beta1AuthorizationpolicyValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/security-v1beta1-AuthorizationPolicy-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetSecurityV1beta1PeerauthenticationInvalidYaml = []byte(`apiVersion: "security.istio.io/v1beta1"
kind: "PeerAuthentication"
metadata:
  name: invalid-peer-authentication
spec:
  portLevelMtls: {}
`)

func datasetSecurityV1beta1PeerauthenticationInvalidYamlBytes() ([]byte, error) {
	return _datasetSecurityV1beta1PeerauthenticationInvalidYaml, nil
}

func datasetSecurityV1beta1PeerauthenticationInvalidYaml() (*asset, error) {
	bytes, err := datasetSecurityV1beta1PeerauthenticationInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/security-v1beta1-PeerAuthentication-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetSecurityV1beta1PeerauthenticationValidYaml = []byte(`apiVersion: "security.istio.io/v1beta1"
kind: "PeerAuthentication"
metadata:
  name: valid-peer-authentication
spec:
  selector:
    matchLabels:
      app: httpbin
      version: v1
  mtls:
    mode: PERMISSIVE
  peerLevelMtls:
    8080:
      mode: STRICT
`)

func datasetSecurityV1beta1PeerauthenticationValidYamlBytes() ([]byte, error) {
	return _datasetSecurityV1beta1PeerauthenticationValidYaml, nil
}

func datasetSecurityV1beta1PeerauthenticationValidYaml() (*asset, error) {
	bytes, err := datasetSecurityV1beta1PeerauthenticationValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/security-v1beta1-PeerAuthentication-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetSecurityV1beta1RequestauthenticationInvalidYaml = []byte(`apiVersion: "security.istio.io/v1beta1"
kind: "RequestAuthentication"
metadata:
  name: invalid-request-authentication
spec:
  selector:
    matchLabels:
      app: httpbin
      version: v1
  jwtRules:
  - issuer: ""
`)

func datasetSecurityV1beta1RequestauthenticationInvalidYamlBytes() ([]byte, error) {
	return _datasetSecurityV1beta1RequestauthenticationInvalidYaml, nil
}

func datasetSecurityV1beta1RequestauthenticationInvalidYaml() (*asset, error) {
	bytes, err := datasetSecurityV1beta1RequestauthenticationInvalidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/security-v1beta1-RequestAuthentication-invalid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _datasetSecurityV1beta1RequestauthenticationValidYaml = []byte(`apiVersion: "security.istio.io/v1beta1"
kind: "RequestAuthentication"
metadata:
  name: valid-request-authentication
spec:
  selector:
    matchLabels:
      app: httpbin
      version: v1
  jwtRules:
  - issuer: example.com 
`)

func datasetSecurityV1beta1RequestauthenticationValidYamlBytes() ([]byte, error) {
	return _datasetSecurityV1beta1RequestauthenticationValidYaml, nil
}

func datasetSecurityV1beta1RequestauthenticationValidYaml() (*asset, error) {
	bytes, err := datasetSecurityV1beta1RequestauthenticationValidYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dataset/security-v1beta1-RequestAuthentication-valid.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"dataset/authentication-v1alpha1-MeshPolicy-invalid.yaml":     datasetAuthenticationV1alpha1MeshpolicyInvalidYaml,
	"dataset/authentication-v1alpha1-MeshPolicy-valid.yaml":       datasetAuthenticationV1alpha1MeshpolicyValidYaml,
	"dataset/authentication-v1alpha1-Policy-invalid.yaml":         datasetAuthenticationV1alpha1PolicyInvalidYaml,
	"dataset/authentication-v1alpha1-Policy-valid.yaml":           datasetAuthenticationV1alpha1PolicyValidYaml,
	"dataset/config-v1alpha2-HTTPAPISpec-invalid.yaml":            datasetConfigV1alpha2HttpapispecInvalidYaml,
	"dataset/config-v1alpha2-HTTPAPISpec-valid.yaml":              datasetConfigV1alpha2HttpapispecValidYaml,
	"dataset/config-v1alpha2-HTTPAPISpecBinding-invalid.yaml":     datasetConfigV1alpha2HttpapispecbindingInvalidYaml,
	"dataset/config-v1alpha2-HTTPAPISpecBinding-valid.yaml":       datasetConfigV1alpha2HttpapispecbindingValidYaml,
	"dataset/config-v1alpha2-QuotaSpec-invalid.yaml":              datasetConfigV1alpha2QuotaspecInvalidYaml,
	"dataset/config-v1alpha2-QuotaSpec-valid.yaml":                datasetConfigV1alpha2QuotaspecValidYaml,
	"dataset/config-v1alpha2-QuotaSpecBinding-invalid.yaml":       datasetConfigV1alpha2QuotaspecbindingInvalidYaml,
	"dataset/config-v1alpha2-QuotaSpecBinding-valid.yaml":         datasetConfigV1alpha2QuotaspecbindingValidYaml,
	"dataset/config-v1alpha2-adapter-invalid.yaml":                datasetConfigV1alpha2AdapterInvalidYaml,
	"dataset/config-v1alpha2-adapter-valid.yaml":                  datasetConfigV1alpha2AdapterValidYaml,
	"dataset/config-v1alpha2-attributemanifest-invalid.yaml":      datasetConfigV1alpha2AttributemanifestInvalidYaml,
	"dataset/config-v1alpha2-attributemanifest-valid.yaml":        datasetConfigV1alpha2AttributemanifestValidYaml,
	"dataset/config-v1alpha2-handler-invalid.yaml":                datasetConfigV1alpha2HandlerInvalidYaml,
	"dataset/config-v1alpha2-handler-valid.yaml":                  datasetConfigV1alpha2HandlerValidYaml,
	"dataset/config-v1alpha2-instance-invalid.yaml":               datasetConfigV1alpha2InstanceInvalidYaml,
	"dataset/config-v1alpha2-instance-valid.yaml":                 datasetConfigV1alpha2InstanceValidYaml,
	"dataset/config-v1alpha2-rule-invalid.yaml":                   datasetConfigV1alpha2RuleInvalidYaml,
	"dataset/config-v1alpha2-rule-valid.yaml":                     datasetConfigV1alpha2RuleValidYaml,
	"dataset/config-v1alpha2-template-invalid.yaml":               datasetConfigV1alpha2TemplateInvalidYaml,
	"dataset/config-v1alpha2-template-valid.yaml":                 datasetConfigV1alpha2TemplateValidYaml,
	"dataset/networking-v1alpha3-DestinationRule-invalid.yaml":    datasetNetworkingV1alpha3DestinationruleInvalidYaml,
	"dataset/networking-v1alpha3-DestinationRule-valid.yaml":      datasetNetworkingV1alpha3DestinationruleValidYaml,
	"dataset/networking-v1alpha3-EnvoyFilter-invalid.yaml":        datasetNetworkingV1alpha3EnvoyfilterInvalidYaml,
	"dataset/networking-v1alpha3-EnvoyFilter-valid.yaml":          datasetNetworkingV1alpha3EnvoyfilterValidYaml,
	"dataset/networking-v1alpha3-Gateway-invalid.yaml":            datasetNetworkingV1alpha3GatewayInvalidYaml,
	"dataset/networking-v1alpha3-Gateway-valid.yaml":              datasetNetworkingV1alpha3GatewayValidYaml,
	"dataset/networking-v1alpha3-ServiceEntry-invalid.yaml":       datasetNetworkingV1alpha3ServiceentryInvalidYaml,
	"dataset/networking-v1alpha3-ServiceEntry-valid.yaml":         datasetNetworkingV1alpha3ServiceentryValidYaml,
	"dataset/networking-v1alpha3-Sidecar-invalid.yaml":            datasetNetworkingV1alpha3SidecarInvalidYaml,
	"dataset/networking-v1alpha3-Sidecar-valid.yaml":              datasetNetworkingV1alpha3SidecarValidYaml,
	"dataset/networking-v1alpha3-VirtualService-invalid.yaml":     datasetNetworkingV1alpha3VirtualserviceInvalidYaml,
	"dataset/networking-v1alpha3-VirtualService-valid.yaml":       datasetNetworkingV1alpha3VirtualserviceValidYaml,
	"dataset/networking-v1beta-DestinationRule-invalid.yaml":      datasetNetworkingV1betaDestinationruleInvalidYaml,
	"dataset/networking-v1beta-DestinationRule-valid.yaml":        datasetNetworkingV1betaDestinationruleValidYaml,
	"dataset/networking-v1beta-Gateway-invalid.yaml":              datasetNetworkingV1betaGatewayInvalidYaml,
	"dataset/networking-v1beta-Gateway-valid.yaml":                datasetNetworkingV1betaGatewayValidYaml,
	"dataset/networking-v1beta-Sidecar-invalid.yaml":              datasetNetworkingV1betaSidecarInvalidYaml,
	"dataset/networking-v1beta-Sidecar-valid.yaml":                datasetNetworkingV1betaSidecarValidYaml,
	"dataset/networking-v1beta-VirtualService-invalid.yaml":       datasetNetworkingV1betaVirtualserviceInvalidYaml,
	"dataset/networking-v1beta-VirtualService-valid.yaml":         datasetNetworkingV1betaVirtualserviceValidYaml,
	"dataset/rbac-v1alpha1-ClusterRbacConfig-invalid.yaml":        datasetRbacV1alpha1ClusterrbacconfigInvalidYaml,
	"dataset/rbac-v1alpha1-ClusterRbacConfig-valid.yaml":          datasetRbacV1alpha1ClusterrbacconfigValidYaml,
	"dataset/rbac-v1alpha1-RBacConfig-invalid.yaml":               datasetRbacV1alpha1RbacconfigInvalidYaml,
	"dataset/rbac-v1alpha1-RBacConfig-valid.yaml":                 datasetRbacV1alpha1RbacconfigValidYaml,
	"dataset/rbac-v1alpha1-ServiceRole-invalid.yaml":              datasetRbacV1alpha1ServiceroleInvalidYaml,
	"dataset/rbac-v1alpha1-ServiceRole-valid.yaml":                datasetRbacV1alpha1ServiceroleValidYaml,
	"dataset/rbac-v1alpha1-ServiceRoleBinding-invalid.yaml":       datasetRbacV1alpha1ServicerolebindingInvalidYaml,
	"dataset/rbac-v1alpha1-ServiceRoleBinding-valid.yaml":         datasetRbacV1alpha1ServicerolebindingValidYaml,
	"dataset/security-v1beta1-AuthorizationPolicy-invalid.yaml":   datasetSecurityV1beta1AuthorizationpolicyInvalidYaml,
	"dataset/security-v1beta1-AuthorizationPolicy-valid.yaml":     datasetSecurityV1beta1AuthorizationpolicyValidYaml,
	"dataset/security-v1beta1-PeerAuthentication-invalid.yaml":    datasetSecurityV1beta1PeerauthenticationInvalidYaml,
	"dataset/security-v1beta1-PeerAuthentication-valid.yaml":      datasetSecurityV1beta1PeerauthenticationValidYaml,
	"dataset/security-v1beta1-RequestAuthentication-invalid.yaml": datasetSecurityV1beta1RequestauthenticationInvalidYaml,
	"dataset/security-v1beta1-RequestAuthentication-valid.yaml":   datasetSecurityV1beta1RequestauthenticationValidYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"dataset": &bintree{nil, map[string]*bintree{
		"authentication-v1alpha1-MeshPolicy-invalid.yaml":     &bintree{datasetAuthenticationV1alpha1MeshpolicyInvalidYaml, map[string]*bintree{}},
		"authentication-v1alpha1-MeshPolicy-valid.yaml":       &bintree{datasetAuthenticationV1alpha1MeshpolicyValidYaml, map[string]*bintree{}},
		"authentication-v1alpha1-Policy-invalid.yaml":         &bintree{datasetAuthenticationV1alpha1PolicyInvalidYaml, map[string]*bintree{}},
		"authentication-v1alpha1-Policy-valid.yaml":           &bintree{datasetAuthenticationV1alpha1PolicyValidYaml, map[string]*bintree{}},
		"config-v1alpha2-HTTPAPISpec-invalid.yaml":            &bintree{datasetConfigV1alpha2HttpapispecInvalidYaml, map[string]*bintree{}},
		"config-v1alpha2-HTTPAPISpec-valid.yaml":              &bintree{datasetConfigV1alpha2HttpapispecValidYaml, map[string]*bintree{}},
		"config-v1alpha2-HTTPAPISpecBinding-invalid.yaml":     &bintree{datasetConfigV1alpha2HttpapispecbindingInvalidYaml, map[string]*bintree{}},
		"config-v1alpha2-HTTPAPISpecBinding-valid.yaml":       &bintree{datasetConfigV1alpha2HttpapispecbindingValidYaml, map[string]*bintree{}},
		"config-v1alpha2-QuotaSpec-invalid.yaml":              &bintree{datasetConfigV1alpha2QuotaspecInvalidYaml, map[string]*bintree{}},
		"config-v1alpha2-QuotaSpec-valid.yaml":                &bintree{datasetConfigV1alpha2QuotaspecValidYaml, map[string]*bintree{}},
		"config-v1alpha2-QuotaSpecBinding-invalid.yaml":       &bintree{datasetConfigV1alpha2QuotaspecbindingInvalidYaml, map[string]*bintree{}},
		"config-v1alpha2-QuotaSpecBinding-valid.yaml":         &bintree{datasetConfigV1alpha2QuotaspecbindingValidYaml, map[string]*bintree{}},
		"config-v1alpha2-adapter-invalid.yaml":                &bintree{datasetConfigV1alpha2AdapterInvalidYaml, map[string]*bintree{}},
		"config-v1alpha2-adapter-valid.yaml":                  &bintree{datasetConfigV1alpha2AdapterValidYaml, map[string]*bintree{}},
		"config-v1alpha2-attributemanifest-invalid.yaml":      &bintree{datasetConfigV1alpha2AttributemanifestInvalidYaml, map[string]*bintree{}},
		"config-v1alpha2-attributemanifest-valid.yaml":        &bintree{datasetConfigV1alpha2AttributemanifestValidYaml, map[string]*bintree{}},
		"config-v1alpha2-handler-invalid.yaml":                &bintree{datasetConfigV1alpha2HandlerInvalidYaml, map[string]*bintree{}},
		"config-v1alpha2-handler-valid.yaml":                  &bintree{datasetConfigV1alpha2HandlerValidYaml, map[string]*bintree{}},
		"config-v1alpha2-instance-invalid.yaml":               &bintree{datasetConfigV1alpha2InstanceInvalidYaml, map[string]*bintree{}},
		"config-v1alpha2-instance-valid.yaml":                 &bintree{datasetConfigV1alpha2InstanceValidYaml, map[string]*bintree{}},
		"config-v1alpha2-rule-invalid.yaml":                   &bintree{datasetConfigV1alpha2RuleInvalidYaml, map[string]*bintree{}},
		"config-v1alpha2-rule-valid.yaml":                     &bintree{datasetConfigV1alpha2RuleValidYaml, map[string]*bintree{}},
		"config-v1alpha2-template-invalid.yaml":               &bintree{datasetConfigV1alpha2TemplateInvalidYaml, map[string]*bintree{}},
		"config-v1alpha2-template-valid.yaml":                 &bintree{datasetConfigV1alpha2TemplateValidYaml, map[string]*bintree{}},
		"networking-v1alpha3-DestinationRule-invalid.yaml":    &bintree{datasetNetworkingV1alpha3DestinationruleInvalidYaml, map[string]*bintree{}},
		"networking-v1alpha3-DestinationRule-valid.yaml":      &bintree{datasetNetworkingV1alpha3DestinationruleValidYaml, map[string]*bintree{}},
		"networking-v1alpha3-EnvoyFilter-invalid.yaml":        &bintree{datasetNetworkingV1alpha3EnvoyfilterInvalidYaml, map[string]*bintree{}},
		"networking-v1alpha3-EnvoyFilter-valid.yaml":          &bintree{datasetNetworkingV1alpha3EnvoyfilterValidYaml, map[string]*bintree{}},
		"networking-v1alpha3-Gateway-invalid.yaml":            &bintree{datasetNetworkingV1alpha3GatewayInvalidYaml, map[string]*bintree{}},
		"networking-v1alpha3-Gateway-valid.yaml":              &bintree{datasetNetworkingV1alpha3GatewayValidYaml, map[string]*bintree{}},
		"networking-v1alpha3-ServiceEntry-invalid.yaml":       &bintree{datasetNetworkingV1alpha3ServiceentryInvalidYaml, map[string]*bintree{}},
		"networking-v1alpha3-ServiceEntry-valid.yaml":         &bintree{datasetNetworkingV1alpha3ServiceentryValidYaml, map[string]*bintree{}},
		"networking-v1alpha3-Sidecar-invalid.yaml":            &bintree{datasetNetworkingV1alpha3SidecarInvalidYaml, map[string]*bintree{}},
		"networking-v1alpha3-Sidecar-valid.yaml":              &bintree{datasetNetworkingV1alpha3SidecarValidYaml, map[string]*bintree{}},
		"networking-v1alpha3-VirtualService-invalid.yaml":     &bintree{datasetNetworkingV1alpha3VirtualserviceInvalidYaml, map[string]*bintree{}},
		"networking-v1alpha3-VirtualService-valid.yaml":       &bintree{datasetNetworkingV1alpha3VirtualserviceValidYaml, map[string]*bintree{}},
		"networking-v1beta-DestinationRule-invalid.yaml":      &bintree{datasetNetworkingV1betaDestinationruleInvalidYaml, map[string]*bintree{}},
		"networking-v1beta-DestinationRule-valid.yaml":        &bintree{datasetNetworkingV1betaDestinationruleValidYaml, map[string]*bintree{}},
		"networking-v1beta-Gateway-invalid.yaml":              &bintree{datasetNetworkingV1betaGatewayInvalidYaml, map[string]*bintree{}},
		"networking-v1beta-Gateway-valid.yaml":                &bintree{datasetNetworkingV1betaGatewayValidYaml, map[string]*bintree{}},
		"networking-v1beta-Sidecar-invalid.yaml":              &bintree{datasetNetworkingV1betaSidecarInvalidYaml, map[string]*bintree{}},
		"networking-v1beta-Sidecar-valid.yaml":                &bintree{datasetNetworkingV1betaSidecarValidYaml, map[string]*bintree{}},
		"networking-v1beta-VirtualService-invalid.yaml":       &bintree{datasetNetworkingV1betaVirtualserviceInvalidYaml, map[string]*bintree{}},
		"networking-v1beta-VirtualService-valid.yaml":         &bintree{datasetNetworkingV1betaVirtualserviceValidYaml, map[string]*bintree{}},
		"rbac-v1alpha1-ClusterRbacConfig-invalid.yaml":        &bintree{datasetRbacV1alpha1ClusterrbacconfigInvalidYaml, map[string]*bintree{}},
		"rbac-v1alpha1-ClusterRbacConfig-valid.yaml":          &bintree{datasetRbacV1alpha1ClusterrbacconfigValidYaml, map[string]*bintree{}},
		"rbac-v1alpha1-RBacConfig-invalid.yaml":               &bintree{datasetRbacV1alpha1RbacconfigInvalidYaml, map[string]*bintree{}},
		"rbac-v1alpha1-RBacConfig-valid.yaml":                 &bintree{datasetRbacV1alpha1RbacconfigValidYaml, map[string]*bintree{}},
		"rbac-v1alpha1-ServiceRole-invalid.yaml":              &bintree{datasetRbacV1alpha1ServiceroleInvalidYaml, map[string]*bintree{}},
		"rbac-v1alpha1-ServiceRole-valid.yaml":                &bintree{datasetRbacV1alpha1ServiceroleValidYaml, map[string]*bintree{}},
		"rbac-v1alpha1-ServiceRoleBinding-invalid.yaml":       &bintree{datasetRbacV1alpha1ServicerolebindingInvalidYaml, map[string]*bintree{}},
		"rbac-v1alpha1-ServiceRoleBinding-valid.yaml":         &bintree{datasetRbacV1alpha1ServicerolebindingValidYaml, map[string]*bintree{}},
		"security-v1beta1-AuthorizationPolicy-invalid.yaml":   &bintree{datasetSecurityV1beta1AuthorizationpolicyInvalidYaml, map[string]*bintree{}},
		"security-v1beta1-AuthorizationPolicy-valid.yaml":     &bintree{datasetSecurityV1beta1AuthorizationpolicyValidYaml, map[string]*bintree{}},
		"security-v1beta1-PeerAuthentication-invalid.yaml":    &bintree{datasetSecurityV1beta1PeerauthenticationInvalidYaml, map[string]*bintree{}},
		"security-v1beta1-PeerAuthentication-valid.yaml":      &bintree{datasetSecurityV1beta1PeerauthenticationValidYaml, map[string]*bintree{}},
		"security-v1beta1-RequestAuthentication-invalid.yaml": &bintree{datasetSecurityV1beta1RequestauthenticationInvalidYaml, map[string]*bintree{}},
		"security-v1beta1-RequestAuthentication-valid.yaml":   &bintree{datasetSecurityV1beta1RequestauthenticationValidYaml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
