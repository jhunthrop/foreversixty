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

A workflow someone turned off on GitHub (state `disabled_manually`: the refresh hold keeps
data, bis and engine off) can start no run, so it cannot gate a landing: it is left out, with a
loud line on stderr for each one. Any other state that is not `active` (disabled for
inactivity, by a fork, or unknown) fails the whole listing, so a workflow switched off by
mistake never lets a batch through quietly. The states come from `gh workflow list` (GH, as in
tools/land.sh); a workflow GitHub does not list yet (a file new in this batch) counts as active.
"""

import json
import os
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


def workflow_states() -> dict[str, str]:
    """Each workflow's state on GitHub by its file name: `gh workflow list --all`."""
    out = subprocess.run(
        [
            os.environ.get("GH", "gh"),
            "workflow",
            "list",
            "--all",
            "--limit",
            "500",
            "--json",
            "path,state",
        ],
        check=True,
        capture_output=True,
        text=True,
    ).stdout
    return {
        entry["path"].removeprefix(WORKFLOW_DIR): entry["state"]
        for entry in json.loads(out or "[]")
    }


class DisabledWorkflow(Exception):
    """An expected workflow in a state that is neither active nor disabled_manually."""


def gating(
    found: list[tuple[str, str]], states: dict[str, str]
) -> tuple[list[tuple[str, str]], list[str]]:
    """The expected (file, name) pairs that gate the landing, and the loud lines for those skipped:
    `active` (or not listed yet) gates; `disabled_manually` is skipped; anything else raises."""
    kept, skipped = [], []
    for file, name in found:
        state = states.get(file, "active")
        if state == "active":
            kept.append((file, name))
        elif state == "disabled_manually":
            skipped.append(f"skipping {name}: disabled ({state}); not gating this landing")
        else:
            raise DisabledWorkflow(
                f"{name} ({file}) is {state}: turn it on or off on purpose before landing"
            )
    return kept, skipped


def main(argv: list[str]) -> int:
    if len(argv) != 4:
        print("usage: tools/land_workflows.py <sha> <branch> <base>", file=sys.stderr)
        return 64
    sha, branch, base = argv[1:]
    changed = git("diff", "--name-only", f"{base}...{sha}").split()
    paths = git("ls-tree", "--name-only", sha, WORKFLOW_DIR).split()
    found = []
    for path in sorted(paths):
        if not path.endswith((".yml", ".yaml")):
            continue
        doc = yaml.safe_load(git("show", f"{sha}:{path}")) or {}
        if expected(doc, branch, changed):
            found.append((path.removeprefix(WORKFLOW_DIR), str(doc.get("name", path))))
    try:
        kept, skipped = gating(found, workflow_states() if found else {})
    except DisabledWorkflow as error:
        print(f"land_workflows: {error}", file=sys.stderr)
        return 1
    for line in skipped:
        print(f"land_workflows: {line}", file=sys.stderr)
    for file, name in kept:
        print(f"{file}\t{name}")
    return 0


if __name__ == "__main__":
    sys.exit(main(sys.argv))
