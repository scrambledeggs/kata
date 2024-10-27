import os
import subprocess

project_name = '{{ cookiecutter.project_name }}'
print(f"Current working directory: {os.getcwd()}")

subprocess.run(['go', 'mod', 'init', project_name], check=True)
subprocess.run(['go', 'mod', 'tidy'], check=True)

print(f"\nSAM initialized for service {project_name}.\n")

print("Next steps:")
print(f"1. Run `cd {project_name}` to switch to your new project directory.")
print("2. Fill out the `.secrets.local.json` file with your environment variables.")
print("3. Run `make dev` to start the local development environment.")
print("  3.1. Optionally, use `make dev-watch` for live reloading.")

print("\nAfter...")
print("1. Delete `handlers/HelloWorldV1` and any instances in template.yaml.")
print("2. Update docs/api_contract.yaml, remove HelloWorld instances too.\n")
