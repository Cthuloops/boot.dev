import os


def write_file(working_directory, file_path, content):
    if file_path.split("/")[0] not in os.listdir(working_directory):
        return f'Error: Cannot write to "{file_path}" as it is outside the permitted working directory' 

    full_file_path = os.path.join(working_directory, file_path)
    with open(full_file_path, mode="wt", encoding="utf-8") as f:
        length = f.write(content)

    return f'Successfully wrote to "{file_path}" ({length} characters written)'