import hashlib
import json
from pathlib import Path

MANIFEST = "manifest.json"


def _sha256(path: Path) -> str:
    return hashlib.sha256(path.read_bytes()).hexdigest()


def write_manifest(build_dir: Path, build: str, product: str, fetched_at: str) -> dict:
    files: dict[str, str] = {}
    for path in sorted(build_dir.rglob("*")):
        if not path.is_file():
            continue
        name = path.relative_to(build_dir).as_posix()
        if name == MANIFEST or name.startswith("raw/"):
            continue
        files[name] = _sha256(path)
    m = {"build": build, "product": product, "fetched_at": fetched_at, "files": files}
    (build_dir / MANIFEST).write_text(json.dumps(m, indent=2) + "\n", encoding="utf-8")
    return m


def read_manifest(build_dir: Path) -> dict:
    return json.loads((build_dir / MANIFEST).read_text(encoding="utf-8"))


def verify(build_dir: Path) -> list[str]:
    m = read_manifest(build_dir)
    return [
        name
        for name, digest in m["files"].items()
        if not (build_dir / name).exists() or _sha256(build_dir / name) != digest
    ]
