"""Regenerate sim/cmd/leveling-bis/effectids_generated.go from the pinned
engine fork's own source.

Two passes over every .go file under <fork>/sim whose basename does not
start with "_" (Go does not compile those) or end in "_test.go" (a test
registering a made-up item id is not an engine effect): first collect every
`Name = <int>` constant, then every registration call --
`core.NewItemEffect(X, ...)` or an `itemhelpers.CreateWeaponProc*(X, ...)`
helper -- and resolve X (a constant name or a literal) to its item id. The
heuristic is the one the generated file's own header documents.

Usage: python3 sim/scripts/effectids.py <fork checkout> [--check]
"""
from __future__ import annotations

import re
import subprocess
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
OUT = HERE.parent / "cmd" / "leveling-bis" / "effectids_generated.go"
CONST = re.compile(r"^\s*([A-Za-z_][A-Za-z0-9_]*)\s*=\s*(\d+)\s*$", re.M)
CALL = re.compile(r"\b(?:core\.)?New(?!Enchant)\w*Effect\(\s*([A-Za-z_][A-Za-z0-9_]*|\d+)\s*,|\b(?:itemhelpers\.)?CreateWeapon\w*\(\s*([A-Za-z_][A-Za-z0-9_]*|\d+)\s*,")


def implemented_ids(fork: Path) -> set[int]:
    files = [p for p in (fork / "sim").rglob("*.go") if not p.name.startswith("_") and not p.name.endswith("_test.go")]
    consts: dict[str, int] = {}
    for path in files:
        for name, value in CONST.findall(path.read_text(encoding="utf-8")):
            consts.setdefault(name, int(value))
    ids: set[int] = set()
    for path in files:
        for a, b in CALL.findall(path.read_text(encoding="utf-8")):
            token = a or b
            if token.isdigit():
                ids.add(int(token))
            elif token in consts:
                ids.add(consts[token])
    return ids


def render(ids: set[int], header: str, sha: str) -> str:
    rows = sorted(ids)
    lines = []
    for i in range(0, len(rows), 10):
        lines.append("\t" + " ".join(f"{n}: true," for n in rows[i : i + 10]))
    body = "\n".join(lines)
    return (
        f"{header}// Generated from wowsims-forever {sha} by sim/scripts/effectids.py.\n"
        f"var engineImplementedEffectItemIDs = map[int]bool{{\n{body}\n}}\n\n"
        "func effectImplemented(itemID int) bool {\n\treturn engineImplementedEffectItemIDs[itemID]\n}\n"
    )


def main() -> int:
    fork = Path(sys.argv[1]).resolve()
    check = "--check" in sys.argv
    sha = subprocess.check_output(["git", "-C", str(fork), "rev-parse", "--short=9", "HEAD"], text=True).strip()
    current = OUT.read_text(encoding="utf-8")
    header = current.split("var engineImplementedEffectItemIDs")[0]
    header = re.sub(r"// Generated from wowsims-forever [0-9a-f]+ by sim/scripts/effectids.py\.\n", "", header)
    rendered = render(implemented_ids(fork), header, sha)
    if check:
        if rendered != current:
            print("effectids_generated.go is stale; run sim/scripts/effectids.py", file=sys.stderr)
            return 1
        return 0
    OUT.write_text(rendered, encoding="utf-8")
    print(f"wrote {OUT} ({len(implemented_ids(fork))} ids) from {sha}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
