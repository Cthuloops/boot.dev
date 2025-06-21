import os


def get_file_content(working_directory, file_path):
    # if the parent directory of the file at filepath, (e.g. pkg/calculator.py) pkg,
    # is a subdirectory of the working directory, access should be granted.
    if file_path.split("/")[0] not in os.listdir(working_directory):
        return f'Error: Cannot read "{file_path}" as it is outside the permitted working directory'

    file_path = os.path.join(working_directory, file_path)

    if not os.path.isfile(file_path):
        return f'Error: File not found or is not a regular file: "{file_path}"'

    MAX_CHARS = 10000

    with open(file_path, mode="rt", encoding="utf-8", newline='') as f:
        file_content = f.read(MAX_CHARS)

    if os.path.getsize(filename=file_path) > MAX_CHARS:
        file_content += f"[...File \"{file_path}\" truncated at {MAX_CHARS} characters]"


    return file_content