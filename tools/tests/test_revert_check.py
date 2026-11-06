"""revert_check.py fails a merge that brings back a card note or guide sentence we changed.

Each test builds a throwaway repository: our side changes a note, the other side does not,
and the merge is committed the wrong way (the other side's file wins).

Run: python3 -m unittest discover -s tools/tests -p 'test_revert*.py'
"""

import contextlib
import io
import json
import pathlib
import subprocess
import sys
import tempfile
import unittest

TOOLS = pathlib.Path(__file__).resolve().parent.parent
sys.path.insert(0, str(TOOLS))

import revert_check  # noqa: E402

APL = "data/curated/apl/mage-fire.json"
GUIDE = "web/src/content/guides/mage/fire.md"
OLD_NOTE = "Scorch to build Improved Scorch to 5 stacks."
NEW_NOTE = "Scorch to build Improved Scorch to 5 stacks, or to refresh it with 5 seconds left."
OLD_SENTENCE = "Fireball is the filler once you have Arcane Blast at 57."
NEW_SENTENCE = "Scorch is the filler from Heating Up at 30 until Arcane Blast at 57."


class RevertCheckTest(unittest.TestCase):
    def setUp(self):
        self.tmp = tempfile.TemporaryDirectory()
        self.repo = pathlib.Path(self.tmp.name)
        self.git("init", "-q", "-b", "main")
        self.git("config", "user.name", "test")
        self.git("config", "user.email", "test@example.com")
        self.write(OLD_NOTE, OLD_SENTENCE)
        self.commit("base")

    def tearDown(self):
        self.tmp.cleanup()

    def git(self, *args):
        return subprocess.run(
            ["git", "-C", str(self.repo), *args], capture_output=True, text=True, check=True
        ).stdout.strip()

    def write(self, note, sentence, other="Fire Blast on cooldown."):
        rotation = {"rotation": {"priorityList": [{"notes": note}, {"notes": other}]}}
        for path, text in (
            (APL, json.dumps(rotation, indent=2)),
            (GUIDE, f"# Fire\n\n{sentence}\n"),
        ):
            target = self.repo / path
            target.parent.mkdir(parents=True, exist_ok=True)
            target.write_text(text, encoding="utf-8")

    def commit(self, message):
        self.git("add", "-A")
        self.git("commit", "-q", "-m", message)

    def wrong_way_merge(self, message="Merge main", theirs_note=OLD_NOTE):
        """Our side fixes both texts; main edits another note; the merge takes main's files."""
        self.git("checkout", "-q", "-b", "land")
        self.write(NEW_NOTE, NEW_SENTENCE)
        self.commit("fix the Scorch note and the filler sentence")
        self.git("checkout", "-q", "main")
        self.write(theirs_note, OLD_SENTENCE, other="Fire Blast whenever it is off cooldown.")
        self.commit("main edits another note")
        self.git("checkout", "-q", "land")
        self.git("merge", "-q", "--no-commit", "-s", "ours", "main")
        self.git("checkout", "main", "--", APL, GUIDE)
        self.git("commit", "-q", "-m", message)

    def check(self):
        out = io.StringIO()
        with contextlib.redirect_stdout(out):
            status = revert_check.main(["--repo", str(self.repo)])
        return status, out.getvalue()

    def test_a_non_merge_commit_is_a_no_op(self):
        self.write(NEW_NOTE, NEW_SENTENCE)
        self.commit("fix")
        self.write(OLD_NOTE, OLD_SENTENCE)
        self.commit("plain revert, not a merge")
        status, out = self.check()
        self.assertEqual(status, 0)
        self.assertIn("not a merge", out)

    def test_a_merge_that_brings_back_old_text_fails(self):
        self.wrong_way_merge()
        status, out = self.check()
        self.assertEqual(status, 1)
        self.assertIn(f"REVERT {revert_check.text_hash(OLD_NOTE)} {APL}", out)
        self.assertIn(f"REVERT {revert_check.text_hash(OLD_SENTENCE)} {GUIDE}", out)

    def test_a_tagged_deliberate_return_passes(self):
        tags = "\n".join(
            f"Deliberate-return: {revert_check.text_hash(text)} the owner wants it back"
            for text in (OLD_NOTE, OLD_SENTENCE)
        )
        self.wrong_way_merge(message=f"Merge main\n\n{tags}\n")
        status, out = self.check()
        self.assertEqual(status, 0, out)
        self.assertIn("returned (tagged)", out)

    def test_one_tag_does_not_cover_the_other_text(self):
        tag = f"Deliberate-return: {revert_check.text_hash(OLD_NOTE)} back on purpose"
        self.wrong_way_merge(message=f"Merge main\n\n{tag}\n")
        status, out = self.check()
        self.assertEqual(status, 1)
        self.assertIn(f"REVERT {revert_check.text_hash(OLD_SENTENCE)} {GUIDE}", out)

    def test_a_merge_that_keeps_our_fix_passes(self):
        self.git("checkout", "-q", "-b", "land")
        self.write(NEW_NOTE, NEW_SENTENCE)
        self.commit("fix")
        self.git("checkout", "-q", "main")
        self.write(OLD_NOTE, OLD_SENTENCE, other="Fire Blast whenever it is off cooldown.")
        self.commit("main edits another note")
        self.git("checkout", "-q", "land")
        self.git("merge", "-q", "--no-edit", "main")
        status, out = self.check()
        self.assertEqual(status, 0, out)

    def test_text_the_other_side_wrote_itself_passes(self):
        """Main rewrote the note and then chose the old wording again in a commit of its own."""
        self.git("checkout", "-q", "-b", "land")
        self.write(NEW_NOTE, OLD_SENTENCE)
        self.commit("fix the note")
        self.git("checkout", "-q", "main")
        self.write("Scorch first.", OLD_SENTENCE)
        self.commit("main rewrites the note")
        self.write(OLD_NOTE, OLD_SENTENCE)
        self.commit("main returns the old wording")
        self.git("checkout", "-q", "land")
        self.git("merge", "-q", "--no-commit", "-s", "ours", "main")
        self.git("checkout", "main", "--", APL)
        self.git("commit", "-q", "-m", "Merge main")
        status, out = self.check()
        self.assertEqual(status, 0, out)


if __name__ == "__main__":
    unittest.main()
