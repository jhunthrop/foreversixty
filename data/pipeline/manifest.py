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


def refresh_manifest(build_dir: Path) -> dict:
    """Re-hash a build directory in place, keeping the provenance it already has.

    `normalize` writes the manifest at the end of its run. A later command that
    adds a file to the same build directory -- `simdb`, `simconst` -- would
    otherwise leave a manifest that no longer covers the directory, and
    `verify` would still pass, because it only checks the files it lists.
    """
    path = build_dir / MANIFEST
    if not path.exists():
        raise SystemExit(f"no {path}; run `python -m pipeline normalize` for this build first")
    existing = read_manifest(build_dir)
    return write_manifest(
        build_dir,
        build=existing["build"],
        product=existing["product"],
        fetched_at=existing["fetched_at"],
    )
