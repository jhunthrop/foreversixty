# data/tests/test_audit_real_build.py
"""The one test that runs the whole audit against the real committed build
(the lane brief's own instruction) rather than a small fixture -- guarded
by `FOREVER_AUDIT_REAL=1` so the ordinary test run stays fast and
independent of whatever the real build currently looks like."""

import os

import pytest

from pipeline.audit import run_audit
from pipeline.audit.findings import SEVERITIES
from pipeline.audit.report import summary_line

BUILD = "1.60.1.70291"

pytestmark = pytest.mark.skipif(
    os.environ.get("FOREVER_AUDIT_REAL") != "1",
    reason="set FOREVER_AUDIT_REAL=1 to run the full audit against the real build",
)


def test_full_audit_runs_on_the_real_build_without_crashing():
    results = run_audit(BUILD)
    assert len(results) == 8
    for result in results:
        assert result.code in "ABCDEFGH"
        for finding in result.findings:
            assert finding.severity in SEVERITIES
    line = summary_line(results)
    assert line.startswith("audit: ")
    print(line)
