"""land_workflows.py leaves out a workflow someone turned off on GitHub, loudly, and fails on any
other state that is not active, so a workflow switched off by mistake never lets a batch land
quietly (the refresh hold keeps data, bis and engine disabled_manually; 2026-11-08 an addon batch
could not land because data.yml's paths admit addon/**).

Run: python3 -m unittest discover -s tools/tests -p 'test_land_workflows*.py'
"""

import contextlib
import io
import json
import os
import pathlib
import sys
import tempfile
import unittest
from unittest import mock

TOOLS = pathlib.Path(__file__).resolve().parent.parent
sys.path.insert(0, str(TOOLS))

import land_workflows  # noqa: E402

FOUND = [("addon.yml", "addon"), ("data.yml", "data"), ("web.yml", "web")]


class GatingTest(unittest.TestCase):
    def test_active_workflows_gate(self):
        kept, skipped = land_workflows.gating(
            FOUND, {"addon.yml": "active", "data.yml": "active", "web.yml": "active"}
        )
        self.assertEqual(FOUND, kept)
        self.assertEqual([], skipped)

    def test_a_workflow_github_does_not_list_yet_counts_as_active(self):
        kept, _ = land_workflows.gating([("new.yml", "new")], {})
        self.assertEqual([("new.yml", "new")], kept)

    def test_disabled_manually_is_skipped_with_a_loud_line(self):
        kept, skipped = land_workflows.gating(FOUND, {"data.yml": "disabled_manually"})
        self.assertEqual([("addon.yml", "addon"), ("web.yml", "web")], kept)
        self.assertEqual(
            ["skipping data: disabled (disabled_manually); not gating this landing"], skipped
        )

    def test_any_other_state_fails(self):
        for state in ("disabled_inactivity", "disabled_fork", "unknown"):
            with (
                self.subTest(state=state),
                self.assertRaises(land_workflows.DisabledWorkflow) as caught,
            ):
                land_workflows.gating(FOUND, {"data.yml": state})
            self.assertIn(state, str(caught.exception))


class MainTest(unittest.TestCase):
    """main() reads the states from GH and prints only the gating workflows."""

    def setUp(self):
        self.tmp = tempfile.TemporaryDirectory()
        self.addCleanup(self.tmp.cleanup)
        self.fake_gh = pathlib.Path(self.tmp.name) / "gh"

    def write_gh(self, states):
        payload = json.dumps(
            [
                {"path": f".github/workflows/{file}", "state": state}
                for file, state in states.items()
            ]
        )
        self.fake_gh.write_text(f"#!/bin/sh\necho '{payload}'\n")
        self.fake_gh.chmod(0o755)

    def run_main(self):
        def fake_git(*args):
            if args[0] == "diff":
                return "addon/ForeverSixty/Locale.lua\n"
            if args[0] == "ls-tree":
                return ".github/workflows/addon.yml\n.github/workflows/data.yml\n"
            name = args[1].split("/")[-1].removesuffix(".yml")
            return (
                f"name: {name}\non:\n  push:\n    branches: ['land/**']\n    paths: ['addon/**']\n"
            )

        out, err = io.StringIO(), io.StringIO()
        with (
            mock.patch.object(land_workflows, "git", fake_git),
            mock.patch.dict(os.environ, {"GH": str(self.fake_gh)}),
        ):
            with contextlib.redirect_stdout(out), contextlib.redirect_stderr(err):
                status = land_workflows.main(["land_workflows.py", "abc", "land/1", "main"])
        return status, out.getvalue(), err.getvalue()

    def test_prints_the_active_one_and_says_which_it_skipped(self):
        self.write_gh({"addon.yml": "active", "data.yml": "disabled_manually"})
        status, out, err = self.run_main()
        self.assertEqual(0, status)
        self.assertEqual("addon.yml\taddon\n", out)
        self.assertIn("skipping data: disabled (disabled_manually); not gating this landing", err)

    def test_fails_on_a_workflow_disabled_for_inactivity(self):
        self.write_gh({"addon.yml": "active", "data.yml": "disabled_inactivity"})
        status, out, err = self.run_main()
        self.assertEqual(1, status)
        self.assertEqual("", out)
        self.assertIn("disabled_inactivity", err)


if __name__ == "__main__":
    unittest.main()
