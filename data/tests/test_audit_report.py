# data/tests/test_audit_report.py
"""`pipeline.audit.report` -- the summary line's exact format (the brief's
own contract: `python -m pipeline audit` prints this to stdout) and that
`write_report` produces both files."""

import json

from pipeline.audit.findings import CategoryResult
from pipeline.audit.report import render_markdown, summary_line, write_report


def _results():
    a = CategoryResult(code="A", label="Items vs client", primary_source="raw CSVs")
    a.add("blocker", 1, "one blocker")
    a.add("minor", 2, "one minor")
    b = CategoryResult(code="B", label="Quests vs classic-db", primary_source="classic-db")
    b.add("major", 3, "one major")
    return [a, b]


def test_summary_line_format():
    line = summary_line(_results())
    assert line == "audit: 3 findings (1/1/1) across 2 categories"


def test_summary_line_with_no_findings():
    a = CategoryResult(code="A", label="Items vs client", primary_source="raw CSVs")
    assert summary_line([a]) == "audit: 0 findings (0/0/0) across 1 categories"


def test_markdown_has_a_section_per_category():
    md = render_markdown(_results(), "testbuild")
    assert "## A. Items vs client" in md
    assert "## B. Quests vs classic-db" in md
    assert "one blocker" in md


def test_write_report_writes_both_files(tmp_path):
    json_path, md_path = write_report(_results(), tmp_path, "testbuild")
    assert json_path.exists() and md_path.exists()
    document = json.loads(json_path.read_text(encoding="utf-8"))
    assert document["build"] == "testbuild"
    assert len(document["categories"]) == 2
    assert document["categories"][0]["counts_by_severity"]["blocker"] == 1
