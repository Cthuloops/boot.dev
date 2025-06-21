import os
import subprocess


def run_python_file(working_directory, file_path):
    if not file_path.endswith(".py"):
        return f'Error: "{file_path}" is not a Python file.'

    if not os.path.exists(os.path.join(working_directory, file_path)):
        return f'Error: File "{file_path}" not found.'

    if file_path.split("/")[0] not in os.listdir(working_directory):
        return f'Error: Cannot execute "{file_path}" as it is outside the permitted working directory'

    file_path = os.path.join(working_directory, file_path)

    try:
        result = subprocess.run(
            ["python", file_path], # .split("/")[len(file_path.split("/")) - 1]],
            capture_output=True,
            text=True,
            timeout=30.0
        )
    except Exception as e:
        return f"Error: executing Python file: {e}"

    if result.stdout or result.stderr:
        print(f"STDOUT: {result.stdout}")
        print(f"STDERR: {result.stderr}")
    if result.returncode != 0:
        print(f"Process exited with code {result.returncode}")

    if not result.stdout and not result.stderr:
        return "No output produced."