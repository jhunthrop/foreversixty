"""apl-check holds the copies to what the engine runs, never to the notes.

The notes in data/curated/apl/<spec>.json are display text the site alone
renders; the curated file is their only owner. A notes-only edit must pass
`make apl-check` with no `make apl-sync`, fork commit or engine pin, and
`make apl-sync` must leave the fork's copy untouched for one. An action
change must still fail.

Run: python3 -m unittest discover -s tools/tests -p 'test_apl*.py'
"""

import copy
import json
import os
import pathlib
import subprocess
import sys
import tempfile
import unittest

TOOLS = pathlib.Path(__file__).resolve().parent.parent

ROTATION = {
    "type": "TypeAPL",
    "prepullActions": [
        {
            "action": {"castSpell": {"spellId": {"spellId": 2687}}},
            "doAtValue": {"const": {"val": "-1s"}},
            "notes": "Bloodrage before the pull.",
        }
    ],
    "priorityList": [
        {
            "notes": "Mortal Strike on cooldown.",
            "action": {"castSpell": {"spellId": {"spellId": 21553, "rank": 4}}},
        },
        {
            "notes": "Overpower whenever it lights up.",
            "action": {"castSpell": {"spellId": {"spellId": 11585, "rank": 4}}},
        },
    ],
}


class AplCheckTest(unittest.TestCase):
    def setUp(self):
        self.tmp = tempfile.TemporaryDirectory()
        self.root = pathlib.Path(self.tmp.name)
        self.curated_dir = self.root / "data/curated/apl"
        self.curated_dir.mkdir(parents=True)
        self.specs_json = self.root / "data/curated/specs.json"
        self.specs_json.write_text(json.dumps([{"spec": "warrior-arms", "spec_slug": "arms"}]))
        self.engine = self.root / "engine"
        self.fork_copy = self.engine / "ui/warrior/apls/forever_arms.apl.json"
        self.fork_copy.parent.mkdir(parents=True)
        self.site_copy = self.root / "sim/request/apl/warrior-arms.apl.json"
        self.site_copy.parent.mkdir(parents=True)
        self.write_curated(ROTATION)
        self.assertEqual(self.run_tool("apl_sync.py").returncode, 0)

    def tearDown(self):
        self.tmp.cleanup()

    def write_curated(self, rotation):
        doc = {"spec": "warrior-arms", "state": "written", "notes": "top level", "rotation": rotation}
        (self.curated_dir / "warrior-arms.json").write_text(json.dumps(doc, indent=1) + "\n")

    def run_tool(self, script):
        env = {
            **os.environ,
            "ENGINE_DIR": str(self.engine),
            "CURATED_APL_DIR": str(self.curated_dir),
            "CURATED_SPECS_JSON": str(self.specs_json),
        }
        return subprocess.run(
            [sys.executable, str(TOOLS / script)], cwd=self.root, env=env, capture_output=True, text=True
        )

    def test_synced_copies_pass(self):
        self.assertEqual(self.run_tool("apl_check.py").returncode, 0)

    def test_notes_only_edit_passes_without_a_sync(self):
        reworded = copy.deepcopy(ROTATION)
        reworded["priorityList"][0]["notes"] = "Mortal Strike beats everything per rage."
        reworded["prepullActions"][0]["notes"] = "Bloodrage, shout, then charge."
        del reworded["priorityList"][1]["notes"]
        self.write_curated(reworded)
        result = self.run_tool("apl_check.py")
        self.assertEqual(result.returncode, 0, result.stderr)

    def test_notes_only_edit_never_touches_the_fork(self):
        before = self.fork_copy.read_text()
        reworded = copy.deepcopy(ROTATION)
        reworded["priorityList"][0]["notes"] = "Reworded."
        self.write_curated(reworded)
        self.assertEqual(self.run_tool("apl_sync.py").returncode, 0)
        self.assertEqual(self.fork_copy.read_text(), before)
        self.assertIn("Reworded.", self.site_copy.read_text())

    def test_action_change_fails(self):
        for change in (
            lambda r: r["priorityList"][0]["action"]["castSpell"]["spellId"].update(rank=3),
            lambda r: r["priorityList"].reverse(),
            lambda r: r["priorityList"][1]["action"].update(condition={"isExecutePhase": {}}),
            lambda r: r["prepullActions"][0]["doAtValue"]["const"].update(val="-2s"),
        ):
            changed = copy.deepcopy(ROTATION)
            change(changed)
            self.write_curated(changed)
            result = self.run_tool("apl_check.py")
            self.assertEqual(result.returncode, 1)
            self.assertIn("runs a different rotation", result.stderr)

    def test_action_change_syncs_into_the_fork(self):
        changed = copy.deepcopy(ROTATION)
        changed["priorityList"][0]["action"]["castSpell"]["spellId"]["rank"] = 3
        self.write_curated(changed)
        self.assertEqual(self.run_tool("apl_sync.py").returncode, 0)
        self.assertEqual(json.loads(self.fork_copy.read_text()), changed)
        self.assertEqual(self.run_tool("apl_check.py").returncode, 0)


if __name__ == "__main__":
    unittest.main()
