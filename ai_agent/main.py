import os
import sys
from dotenv import load_dotenv
from google import genai
from google.genai import types


load_dotenv()
api_key = os.environ.get("GEMINI_API_KEY")


client = genai.Client(api_key=api_key)


system_prompt = """
You are a helpful AI coding agent.

When a user asks a question or makes a request, make a function call plan. You can perform the following operations:

- List files and directories

All paths you provide should be relative to the working directory. You do not need to specify the working directory in your function calls as it is automatically injected for security reasons.
"""
model_name = "gemini-2.0-flash-001"


schema_get_files_info = types.FunctionDeclaration(
    name="get_files_info",
    description="Lists files in the specified directory along with their sizes, constrained to the working directory.",
    parameters=types.Schema(
        type=types.Type.OBJECT,
        properties={
            "directory": types.Schema(
                type=types.Type.STRING,
                description="The directory to list files from, relative to the working directory. If not provided, lists files in the working directory itself.",
            ),
        },
    ),
)

available_functions = types.Tool(
    function_declarations=[
        schema_get_files_info,
    ]
)

user_prompt = sys.argv[1]
messages = [
    types.Content(role="user", parts=[types.Part(text=user_prompt)]),
]
response = client.models.generate_content(model=model_name,
                                          contents=messages,
                                          config=types.GenerateContentConfig(
                                              tools=[available_functions],
                                              system_instruction=system_prompt),
                                          )

verbose = False
for arg in sys.argv[1:]:
    if arg == "--verbose":
        verbose = True

prompt_tokens = response.usage_metadata.prompt_token_count
response_tokens = response.usage_metadata.candidates_token_count 

if verbose:
    print(f"User prompt: {user_prompt}")
    print(f"Prompt tokens: {prompt_tokens}")
    print(f"Response tokens: {response_tokens}")
else:
    if response.function_calls:
        print(f"Calling function: {response.function_calls[0].name}({response.function_calls[0].args})")
    else:
        print(response.text)