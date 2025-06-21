import os


def get_files_info(working_directory, directory=None):

    # Handle special case when checking working directory
    if directory == ".":
        directory = working_directory
        working_directory = os.getcwd()

            
    # Check if directory is a subdirectory of working directory.
    if directory not in os.listdir(working_directory):
        return f'Error: Cannot list "{directory}" as it is outside the permitted working directory'

    # Check if passed argument is actually a directory.
    if not os.path.isdir(os.path.join(working_directory, directory)):
        return f'Error: "{directory}" is not a directory'
    
    # Build list of files/subdirectories in target directory.
    target_directory_list = os.listdir(os.path.join(working_directory, directory))
    # Generate an absolute path for each file/subdirectory.
    abs_path = lambda x: os.path.abspath(os.path.join(working_directory, directory, x))
    try:
        paths_dict_list = [{item: 
                                {"file_size": os.path.getsize(abs_path(item)),
                                "is_dir": os.path.isdir(abs_path(item))}
                            }for item in target_directory_list]
    except OSError as e:
        return f'Error: {e}'

    # Build the return string
    try:
        return '\n'.join(f"- {key}: file_size={val['file_size']} bytes, is_dir={val['is_dir']}" for item in paths_dict_list for key, val in item.items())
    except TypeError as e:
        return f'Error: {e}'