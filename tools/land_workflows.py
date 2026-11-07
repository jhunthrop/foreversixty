#!/usr/bin/env python3
"""tools/land_workflows.py <sha> <branch> <base>: the check workflows a push of <branch> at <sha>
should start, one "<file>\t<name>" line each (tools/land.sh's guard).

A workflow is expected when its `on.push.branches` matches <branch> and its `on.push.paths`
(or `paths-ignore`) admits a file the batch changes, `git diff --name-only <base>...<sha>`. A
workflow with no path filter is always expected. GitHub evaluates a push's paths against the
commits that push carried, so a land branch whose head sha the remote already held (pushed
first on a lane branch) carries none and starts no path-filtered workflow; land.sh dispatches
the ones this lists that never started, instead of landing on whatever did.

Filter patterns follow GitHub's: `*` matches within one path segment, `**` across segments,
`?` one character, and a leading `!` negates a pattern written after a match.
"""

import re
import subprocess
import sys

try:
    import yaml
except ImportError:
    sys.exit("tools/land_workflows.py needs PyYAML: python3 -m pip install pyyaml")

WORKFLOW_DIR = ".github/workflows/"


def git(*args: str) -> str:
    return subprocess.run(["git", *args], check=True, capture_output=True, text=True).stdout


def glob_regex(pattern: str) -> re.Pattern[str]:
    out, i = "", 0
    while i < len(pattern):
        if pattern.startswith("**", i):
            out += ".*"
            i += 2
        elif pattern[i] == "*":
            out += "[^/]*"
            i += 1
        elif pattern[i] == "?":
            out += "[^/]"
            i += 1
        else:
            out += re.escape(pattern[i])
            i += 1
    return re.compile(out + r"\Z")


def matches(patterns: list[str], value: str) -> bool:
    """GitHub's ordered filter: the last pattern that matches decides, `!` excluding."""
    hit = False
    for pattern in patterns:
        negated = pattern.startswith("!")
        if glob_regex(pattern[1:] if negated else pattern).match(value):
            hit = not negated
    return hit


def push_trigger(doc: dict) -> dict | None:
    # YAML 1.1 reads a bare `on:` key as the boolean True.
    triggers = doc.get("on", doc.get(True))
    if isinstance(triggers, str):
        triggers = {triggers: None}
    elif isinstance(triggers, list):
        triggers = {name: None for name in triggers}
    if not isinstance(triggers, dict) or "push" not in triggers:
        return None
    return triggers["push"] or {}


def expected(doc: dict, branch: str, changed: list[str]) -> bool:
    push = push_trigger(doc)
    if push is None:
        return False
    branches = push.get("branches")
    if branches is None or not matches(list(branches), branch):
        return False
    if "paths" in push:
        return any(matches(list(push["paths"]), path) for path in changed)
    if "paths-ignore" in push:
        return any(not matches(list(push["paths-ignore"]), path) for path in changed)
    return True


def main(argv: list[str]) -> int:
    if len(argv) != 4:
        print("usage: tools/land_workflows.py <sha> <branch> <base>", file=sys.stderr)
        return 64
    sha, branch, base = argv[1:]
    changed = git("diff", "--name-only", f"{base}...{sha}").split()
    paths = git("ls-tree", "--name-only", sha, WORKFLOW_DIR).split()
    for path in sorted(paths):
        if not path.endswith((".yml", ".yaml")):
            continue
        doc = yaml.safe_load(git("show", f"{sha}:{path}")) or {}
        if expected(doc, branch, changed):
            print(f"{path.removeprefix(WORKFLOW_DIR)}\t{doc.get('name', path)}")
    return 0


if __name__ == "__main__":
    sys.exit(main(sys.argv))
