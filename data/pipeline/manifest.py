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


def newest_build(root: Path = Path("builds")) -> str:
    """The build directory whose manifest records the newest client fetch.

    A test that means "the build the site runs on" names it here rather than
    embedding the id: a build string in a test is an edit waiting to be
    forgotten the next time a build lands, and the plan keeps build strings
    inside `builds/` and the handful of single-build conformance tests.

    A directory whose manifest has no `fetched_at` is not a client build --
    `forever-prebeta` is regenerated from a Wowhead snapshot so links shared
    against it keep opening -- and is skipped rather than sorted against a
    null.
    """
    candidates: list[tuple[str, str]] = []
    for path in sorted(root.glob(f"*/{MANIFEST}")):
        manifest = json.loads(path.read_text(encoding="utf-8"))
        fetched_at = manifest.get("fetched_at")
        if fetched_at:
            candidates.append((fetched_at, manifest["build"]))
    if not candidates:
        raise SystemExit(
            f"no fetched build under {root}; run `python -m pipeline fetch` first"
        )
    return max(candidates)[1]
