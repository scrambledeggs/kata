import os
import shutil
import subprocess

project_name = '{{ cookiecutter.project_name }}'
print(f"Current working directory: {os.getcwd()}")

subprocess.run(['go', 'mod', 'init', project_name], check=True)
subprocess.run(['go', 'mod', 'tidy'], check=True)

REMOVE_PATHS = [
    '{% if cookiecutter.with_own_db == "n" %} db/migrations/ {% endif %}',
    '{% if cookiecutter.use_sqlc == "n" %} db/queries/ {% endif %}',
    '{% if cookiecutter.use_sqlc == "n" %} sqlc.yaml {% endif %}',
]

for path in REMOVE_PATHS:
    path = path.strip()
    if path and os.path.exists(path):
        os.unlink(path) if os.path.isfile(path) else shutil.rmtree(path)

print(f"\nSAM initialized for service {project_name}.\n")

print("Next steps:")
print(f"- Run `cd {project_name}` to switch to your new project directory.")
print("- Fill out the `.secrets.local.json` file with your environment variables.")
if "{{ cookiecutter.db_migration }}" == "goose":
    print("- Install goose (https://github.com/pressly/goose?tab=readme-ov-file#install)")
    print("- Add 'DatabaseURL' to your secrets.local.json and add your database URL")
elif "{{ cookiecutter.db_migration }}" == "dbmate":
    print("- Install dbmate (https://github.com/amacneil/dbmate?tab=readme-ov-file#installation)")
    print("- Add 'DatabaseURL' to your secrets.local.json and add your database URL")
if "{{ cookiecutter.use_sqlc }}":
    print("- Install sqlc (https://docs.sqlc.dev/en/latest/overview/install.html)")
print("- Run `make dev` to start the local development environment.")
print("  - Optionally, use `make dev-watch` for live reloading.")

print("\nAfter...")
print("1. Delete `handlers/HelloWorldV1` and any instances in template.yaml.")
print("2. Update docs/api_contract.yaml, remove HelloWorld instances too.\n")
