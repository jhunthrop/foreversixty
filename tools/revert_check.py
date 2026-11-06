#!/usr/bin/env python3
"""revert_check.py: a merge must not quietly bring back card notes or guide text we fixed.

Given a merge commit (HEAD by default), list every card note (a `notes` string in
data/curated/apl/*.json, `level_notes` entries included) and every guide or coach sentence
(the .md files under web/src/content/guides and data/curated/coach) that the merge brought
in relative to its first parent, and fail when one is a verbatim copy of text the first
parent's own history had and later changed: a merge that resolved a conflict the wrong way,
or a stale side, reverted a fix. Text the other side wrote again itself, in an ordinary
(non-merge) commit only it has, is its own edit and passes.

The history read is the last --depth (80) first-parent commits that touched each file.
A non-merge commit is a no-op (exit 0).

A deliberate return passes when the merge commit message carries one line per returned
text, naming the hash this tool prints for it and a reason:

    Deliberate-return: 3f9a0c1b2d4e the owner wants the old Scorch wording back

Run: python3 tools/revert_check.py [--repo DIR] [--rev REV] [--depth N]
Exit status: 0 clean, 1 a revert was found, 2 usage or git error.
Tests: python3 -m unittest discover -s tools/tests -p 'test_revert*.py'
"""

from __future__ import annotations

import argparse
import hashlib
import json
import re
import subprocess
import sys
from collections.abc import Callable, Iterable
from dataclasses import dataclass

NOTE_ROOT = "data/curated/apl"
PROSE_ROOTS = ("web/src/content/guides", "data/curated/coach")
#: Longer `notes` strings are top-level engineering commentary, not card text.
NOTE_MAX_CHARS = 900
#: Shorter fragments (headings, list stubs) repeat by accident.
SENTENCE_MIN_CHARS = 30
DEFAULT_DEPTH = 80
HASH_CHARS = 12
TAG_PATTERN = re.compile(rf"^Deliberate-return:\s*([0-9a-f]{{{HASH_CHARS}}})\b", re.MULTILINE)
SENTENCE_SPLIT = re.compile(r"(?<=[.!?])\s+|\n+")


@dataclass(frozen=True)
class Revert:
    path: str
    text: str

    @property
    def digest(self) -> str:
        return text_hash(self.text)


def text_hash(text: str) -> str:
    return hashlib.sha256(text.encode("utf-8")).hexdigest()[:HASH_CHARS]


class Git:
    """The few git reads the check needs, with blob contents cached by object id."""

    def __init__(self, repo: str) -> None:
        self.repo = repo
        self._blobs: dict[str, str] = {}

    def run(self, *args: str, check: bool = True) -> str:
        return subprocess.run(
            ["git", "-C", self.repo, *args], capture_output=True, text=True, check=check
        ).stdout

    def parents(self, rev: str) -> list[str]:
        return self.run("rev-list", "--parents", "-n", "1", rev).split()[1:]

    def message(self, rev: str) -> str:
        return self.run("log", "-1", "--format=%B", rev)

    def merge_base(self, ours: str, theirs: str) -> str:
        """The common ancestor, or "" for unrelated histories."""
        try:
            return self.run("merge-base", ours, theirs).strip()
        except subprocess.CalledProcessError:
            return ""

    def changed(self, base: str, rev: str, roots: Iterable[str]) -> list[str]:
        return self.run("diff", "--name-only", base, rev, "--", *roots).split()

    def history(self, rev: str, path: str, depth: int) -> list[str]:
        """The first-parent commits before `rev`'s own version that touched `path`."""
        revs = self.run("log", "--first-parent", f"-{depth}", "--format=%H", rev, "--", path)
        return revs.split()[1:]

    def blob_ids(self, pairs: Iterable[tuple[str, str]]) -> dict[tuple[str, str], str]:
        """Each (rev, path)'s blob id, from trees alone: a blobless clone (CI) fetches nothing."""
        by_rev: dict[str, set[str]] = {}
        for rev, path in pairs:
            by_rev.setdefault(rev, set()).add(path)
        found: dict[tuple[str, str], str] = {}
        for rev, paths in by_rev.items():
            try:
                listing = self.run("ls-tree", "-z", "--full-tree", rev, "--", *sorted(paths))
            except subprocess.CalledProcessError:
                continue  # a root commit's "<sha>^"
            for entry in filter(None, listing.split("\0")):
                meta, path = entry.split("\t", 1)
                _mode, kind, oid = meta.split()
                if kind == "blob":
                    found[(rev, path)] = oid
        return found

    def prefetch(self, oids: Iterable[str]) -> None:
        """In a blobless clone (CI), fetch the needed blobs in one round trip, not one each."""
        promisors = self.run("config", "--get-regexp", r"^remote\..*\.promisor$", check=False)
        remotes = [
            line.split()[0][len("remote.") : -len(".promisor")] for line in promisors.splitlines()
        ]
        missing = [oid for oid in sorted(set(oids)) if oid not in self._blobs]
        if not remotes or not missing:
            return
        self.run(
            "-c",
            "fetch.negotiationAlgorithm=noop",
            "fetch",
            "-q",
            "--no-tags",
            "--no-write-fetch-head",
            "--filter=blob:none",
            remotes[0],
            *missing,
        )

    def blobs(self, oids: Iterable[str]) -> dict[str, str]:
        missing = sorted(set(oids) - self._blobs.keys())
        if missing:
            raw = subprocess.run(
                ["git", "-C", self.repo, "cat-file", "--batch"],
                input="\n".join(missing).encode() + b"\n",
                capture_output=True,
                check=True,
            ).stdout
            at = 0
            for _ in missing:
                header_end = raw.index(b"\n", at)
                oid, _kind, size = raw[at:header_end].decode().split()
                start = header_end + 1
                self._blobs[oid] = raw[start : start + int(size)].decode("utf-8", "replace")
                at = start + int(size) + 1
        return self._blobs


def card_notes(text: str) -> set[str]:
    try:
        data = json.loads(text)
    except json.JSONDecodeError:
        return set()
    found: set[str] = set()

    def walk(node: object) -> None:
        if isinstance(node, dict):
            for key, value in node.items():
                if key == "notes" and isinstance(value, str):
                    if value.strip() and len(value) < NOTE_MAX_CHARS:
                        found.add(value)
                else:
                    walk(value)
        elif isinstance(node, list):
            for value in node:
                walk(value)

    walk(data)
    return found


def prose_sentences(text: str) -> set[str]:
    parts = (part.strip() for part in SENTENCE_SPLIT.split(text))
    return {part for part in parts if len(part) > SENTENCE_MIN_CHARS}


class _Texts:
    """Card notes or guide sentences per (rev, path), each blob read once."""

    def __init__(self, git: Git, extract_for: Callable[[str], Callable[[str], set[str]]]) -> None:
        self.git = git
        self.extract_for = extract_for
        self.ids: dict[tuple[str, str], str] = {}

    def load(self, pairs: Iterable[tuple[str, str]]) -> None:
        pairs = [pair for pair in pairs if pair not in self.ids]
        found = self.git.blob_ids(pairs)
        self.git.prefetch(found.values())
        self.git.blobs(found.values())
        self.ids.update(found)

    def of(self, rev: str, path: str) -> set[str]:
        oid = self.ids.get((rev, path))
        return self.extract_for(path)(self.git.blobs([oid])[oid]) if oid else set()


def _extractor(path: str) -> Callable[[str], set[str]]:
    return card_notes if path.endswith(".json") else prose_sentences


def find_reverts(git: Git, merge: str, depth: int) -> tuple[int, list[Revert]]:
    """Every text the merge brought in over its first parent (counted) and the reverts among them.

    A revert is a brought-in text our first-parent history had and later changed. Text the
    other side wrote itself, in an ordinary (non-merge) commit of its own, is that side's
    deliberate edit and passes; text it only carried in through merges is a stale copy.
    """
    ours, theirs = git.parents(merge)[:2]
    paths = [p for p in git.changed(ours, merge, [NOTE_ROOT]) if p.endswith(".json")]
    paths += [p for p in git.changed(ours, merge, PROSE_ROOTS) if p.endswith(".md")]
    histories = {path: git.history(ours, path, depth) for path in paths}
    texts = _Texts(git, _extractor)
    texts.load((rev, path) for path, revs in histories.items() for rev in (merge, ours, *revs))

    brought = 0
    candidates: dict[str, set[str]] = {}
    for path in paths:
        new = texts.of(merge, path) - texts.of(ours, path)
        brought += len(new)
        older = set().union(*(texts.of(rev, path) for rev in histories[path]))
        if new & older:
            candidates[path] = new & older

    reverts: list[Revert] = []
    for path, found in candidates.items():
        authored = _authored_by_theirs(git, texts, ours, theirs, path)
        reverts.extend(Revert(path, text) for text in sorted(found - authored))
    return brought, reverts


def _authored_by_theirs(git: Git, texts: _Texts, ours: str, theirs: str, path: str) -> set[str]:
    """Texts a non-merge commit only the other side has added to `path`."""
    commits = git.run("rev-list", "--no-merges", theirs, f"^{ours}", "--", path).split()
    texts.load((rev, path) for commit in commits for rev in (commit, f"{commit}^"))
    return set().union(*(texts.of(c, path) - texts.of(f"{c}^", path) for c in commits))


def main(argv: list[str] | None = None) -> int:
    parser = argparse.ArgumentParser(description=__doc__.splitlines()[0])
    parser.add_argument("--repo", default=".")
    parser.add_argument("--rev", default="HEAD")
    parser.add_argument("--depth", type=int, default=DEFAULT_DEPTH)
    args = parser.parse_args(argv)
    git = Git(args.repo)
    try:
        if len(git.parents(args.rev)) < 2:
            print(f"revert-check: {args.rev} is not a merge; nothing to check")
            return 0
        brought, reverts = find_reverts(git, args.rev, args.depth)
        tagged = set(TAG_PATTERN.findall(git.message(args.rev)))
    except subprocess.CalledProcessError as err:
        print(f"revert-check: git failed: {' '.join(err.cmd)}\n{err.stderr}", file=sys.stderr)
        return 2
    failed = [r for r in reverts if r.digest not in tagged]
    for revert in reverts:
        status = "returned (tagged)" if revert.digest in tagged else "REVERT"
        print(f"{status} {revert.digest} {revert.path}: {revert.text[:160]}")
    print(
        f"revert-check: {brought} notes and sentences brought in by the merge, "
        f"{len(reverts)} back to older text, {len(failed)} untagged"
    )
    if failed:
        print(
            "revert-check: restore the newer text, or, for a deliberate return, add one line "
            "per text to the merge message: Deliberate-return: <hash> <reason>"
        )
    return 1 if failed else 0


if __name__ == "__main__":
    sys.exit(main())
