# Empty conftest.py at the project root so pytest adds this directory to
# sys.path, making the local `pipeline` package importable in tests without
# installing it (this is a non-packaged/virtual uv project).
