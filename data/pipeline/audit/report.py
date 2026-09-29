"""Renders a run's `CategoryResult`s to `audit.json` (machine-readable,
every finding), `audit.md` (one table per category: counts by severity,
top-10 examples, the primary source used) and the one-line stdout summary
`python -m pipeline audit` prints.
"""

from __future__ import annotations

import json
from dataclasses import asdict
from datetime import UTC, datetime
from pathlib import Path

from pipeline.audit.findings import SEVERITIES, CategoryResult

#: How many example findings `audit.md`'s own table shows per category --
#: the brief's own number. `audit.json` always carries every finding.
MAX_EXAMPLES = 10


def summary_line(results: list[CategoryResult]) -> str:
    total = 0
    by_severity = dict.fromkeys(SEVERITIES, 0)
    for result in results:
        for finding in result.findings:
            total += 1
            by_severity[finding.severity] += 1
    counts = "/".join(str(by_severity[s]) for s in SEVERITIES)
    return f"audit: {total} findings ({counts}) across {len(results)} categories"


def _as_json(results: list[CategoryResult], build: str) -> dict:
    return {
        "build": build,
        "generated_at": datetime.now(UTC).isoformat(),
        "summary": summary_line(results),
        "categories": [
            {
                "code": result.code,
                "label": result.label,
                "primary_source": result.primary_source,
                "checked": result.checked,
                "skipped": result.skipped,
                "counts_by_severity": result.counts_by_severity(),
                "findings": [asdict(f) for f in result.findings],
            }
            for result in results
        ],
    }


def _render_category_md(result: CategoryResult) -> str:
    counts = result.counts_by_severity()
    lines = [
        f"## {result.code}. {result.label}",
        "",
        f"- Checked: {result.checked}",
        f"- Primary source: {result.primary_source}",
        f"- Findings: {sum(counts.values())} "
        f"(blocker {counts['blocker']}, major {counts['major']}, minor {counts['minor']})",
    ]
    if result.skipped:
        lines.append(f"- **Skipped part of this check:** {result.skipped}")
    lines.append("")
    if result.findings:
        lines.append("| severity | subject | message | ours | theirs | source |")
        lines.append("|---|---|---|---|---|---|")
        # Worst-first, then as encountered, so the top 10 are the most
        # actionable ones a fix lane would triage first.
        ordered = sorted(result.findings, key=lambda f: SEVERITIES.index(f.severity))
        for finding in ordered[:MAX_EXAMPLES]:
            lines.append(
                f"| {finding.severity} | {finding.subject} | {finding.message} | "
                f"{finding.ours} | {finding.theirs} | {finding.source} |"
            )
        if len(result.findings) > MAX_EXAMPLES:
            lines.append(f"| ... | {len(result.findings) - MAX_EXAMPLES} more | | | | |")
    lines.append("")
    return "\n".join(lines)


def render_markdown(results: list[CategoryResult], build: str) -> str:
    parts = [f"# Accuracy audit: {build}", "", summary_line(results), ""]
    parts.extend(_render_category_md(result) for result in results)
    return "\n".join(parts)


def write_report(results: list[CategoryResult], out_dir: Path, build: str) -> tuple[Path, Path]:
    out_dir.mkdir(parents=True, exist_ok=True)
    json_path = out_dir / "audit.json"
    md_path = out_dir / "audit.md"
    json_path.write_text(json.dumps(_as_json(results, build), indent=1) + "\n", encoding="utf-8")
    md_path.write_text(render_markdown(results, build), encoding="utf-8")
    return json_path, md_path
