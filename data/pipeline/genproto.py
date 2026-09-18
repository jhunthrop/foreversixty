"""Vendor the engine's .proto files and generate the Python bindings from them.

The simulator's data outputs are the engine's own protobuf messages, so the
schema must be the engine's and nothing else. Rather than hand-write an
encoder, this module copies the engine's unmodified `.proto` files into
`data/proto/` and runs `grpc_tools.protoc` over them. Both the copies and the
generated modules are committed, so running the pipeline needs neither protoc
nor an engine checkout; only regeneration does.

protoc emits `import common_pb2 as common__pb2` for a sibling file, which only
resolves when the output directory is on sys.path. `fix_imports` rewrites those
lines to package-absolute imports so the generated modules live inside
`pipeline.simproto` like any other module.
"""

from __future__ import annotations

import re
import shutil
import subprocess
from pathlib import Path

PROTO_FILES = ("common.proto", "apl.proto", "shaman.proto")
PACKAGE = "pipeline.simproto"

_SIBLING_IMPORT = re.compile(r"^import (\w+_pb2) as (\w+__pb2)$", re.MULTILINE)


class ProtoError(SystemExit):
    """Regeneration could not run, or produced something unusable."""


def fix_imports(source: str, package: str) -> str:
    """Rewrite protoc's sibling imports to package-absolute ones."""
    return _SIBLING_IMPORT.sub(rf"from {package} import \1 as \2", source)


def engine_sha(engine_path: Path) -> str:
    result = subprocess.run(
        ["git", "-C", str(engine_path), "rev-parse", "--short", "HEAD"],
        capture_output=True,
        text=True,
        check=False,
    )
    if result.returncode != 0:
        raise ProtoError(f"{engine_path} is not a git checkout: {result.stderr.strip()}")
    return result.stdout.strip()


def vendor_protos(engine_path: Path, proto_dir: Path) -> str:
    """Copy the engine's .proto files into proto_dir and record the sha."""
    source_dir = engine_path / "proto"
    if not source_dir.is_dir():
        raise ProtoError(f"no proto directory at {source_dir}")
    proto_dir.mkdir(parents=True, exist_ok=True)
    for name in PROTO_FILES:
        source = source_dir / name
        if not source.is_file():
            raise ProtoError(f"the engine has no {name}; its proto API changed shape")
        shutil.copyfile(source, proto_dir / name)
    sha = engine_sha(engine_path)
    (proto_dir / "ENGINE_SHA").write_text(sha + "\n", encoding="utf-8")
    return sha


def generate(proto_dir: Path, out_dir: Path) -> list[Path]:
    """Generate the Python bindings from the vendored protos."""
    from grpc_tools import protoc

    out_dir.mkdir(parents=True, exist_ok=True)
    args = ["protoc", f"-I{proto_dir}", f"--python_out={out_dir}"]
    args += [str(proto_dir / name) for name in PROTO_FILES]
    if protoc.main(args) != 0:
        raise ProtoError(f"protoc failed over {proto_dir}")
    written = []
    for name in PROTO_FILES:
        path = out_dir / (name.removesuffix(".proto") + "_pb2.py")
        path.write_text(fix_imports(path.read_text(encoding="utf-8"), PACKAGE), encoding="utf-8")
        written.append(path)
    return sorted(written)


def refresh(engine_path: Path, proto_dir: Path, out_dir: Path) -> str:
    sha = vendor_protos(engine_path, proto_dir)
    generate(proto_dir, out_dir)
    return sha
