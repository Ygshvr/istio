import os
import filecmp
import os.path
import shutil

def are_dir_trees_equal(fixtures_dir, generated_dir):
    """
    Compare two directories recursively. Files in each directory are
    assumed to be equal if their names and contents are equal.

    @param dir1: First directory path
    @param dir2: Second directory path

    @return: True if the directory trees are the same and 
        there were no errors while accessing the directories or files, 
        False otherwise.
   """

    dirs_cmp = filecmp.dircmp(fixtures_dir, generated_dir)
    if len(dirs_cmp.left_only)>0 or len(dirs_cmp.right_only)>0 or \
        len(dirs_cmp.funny_files)>0:
        print('Mismatch in number of files.')
        return False
    (_, mismatch, errors) =  filecmp.cmpfiles(
        fixtures_dir, generated_dir, dirs_cmp.common_files, shallow=False)
    if len(mismatch)>0 or len(errors)>0:
        print('Mismatch while comparing the files.')
        dirs_cmp.report()
        return False
    for common_dir in dirs_cmp.common_dirs:
        new_dir1 = os.path.join(fixtures_dir, common_dir)
        new_dir2 = os.path.join(generated_dir, common_dir)
        if not are_dir_trees_equal(new_dir1, new_dir2):
            return False
    return True


def main():
    
    if(os.path.isdir('./generated')):
        shutil.rmtree('./generated')
    
    # testing tsb_util.py
    os.system('python tsb_util.py --config ./fixtures/general-config.yml')
    assert are_dir_trees_equal('./fixtures/general_generated', './generated') == True, 'tsb_utils.py test failed.'
    print('>> 1. tsb_utils test completed successfully.')
    # Doing clean up
    shutil.rmtree('./generated')
    
    # testing single_ns.py
    os.system('python single_ns.py --config ./fixtures/httpbin-config.yml')
    assert are_dir_trees_equal('./fixtures/httpbin_generated', './generated') == True, 'single_ns.py test failed.'
    print('>> 2. single_ns.py test completed successfully.')
     # Doing clean up
    shutil.rmtree('./generated')

    # testing bookinfo-single-gw.py
    os.system('python bookinfo-single-gw.py --config ./fixtures/bookinfo-single-direct.yml')
    assert are_dir_trees_equal('./fixtures/bookinfo_single_direct_generated', './generated') == True, 'bookinfo-single.yaml test failed in direct mode.'
    print('>> 3.1 bookinfo-single.yaml direct mode test completed successfully.')
     # Doing clean up
    shutil.rmtree('./generated')
    os.system('python bookinfo-single-gw.py --config ./fixtures/bookinfo-single-bridged.yml')
    assert are_dir_trees_equal('./fixtures/bookinfo_single_bridged_generated', './generated') == True, 'bookinfo-single.yaml test failed in bridged mode.'
    print('>> 3.2 bookinfo-single.yaml bridged mode test completed successfully.')
     # Doing clean up
    shutil.rmtree('./generated')
    
if __name__ == "__main__":
    main()