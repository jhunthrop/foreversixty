"""The one finding shape every check module (A-H) emits, and the small
helpers `report.py` and the check modules share to build up a per-category
result. Tenet 8 (docs/tenets.md): every number, pick, source or weight this
audit reports is compared to a primary source and labelled -- a `Finding`
always names `source`, and a check that cannot reach a primary source for
part of its own brief says so in `CategoryResult.skipped` rather than
fabricating a verdict.
"""

from __future__ import annotations

from dataclasses import dataclass, field

#: Ordered worst-first -- `report.severity_counts` and the stdout summary
#: line both rely on this order.
SEVERITIES: tuple[str, ...] = ("blocker", "major", "minor")


@dataclass(frozen=True)
class Finding:
    """One published fact that disagreed with its primary source.

    `subject` is whichever id the brief's own `item_id|quest_id|npc_id`
    names for this finding -- always a string so a category can mix id
    kinds (a quest check may name either a quest id or the item id its
    reward attaches to) without the report needing to know which.
    """

    category: str
    severity: str
    subject: str
    message: str
    ours: str = ""
    theirs: str = ""
    source: str = ""

    def __post_init__(self) -> None:
        if self.severity not in SEVERITIES:
            raise ValueError(f"finding severity {self.severity!r} must be one of {SEVERITIES}")


@dataclass
class CategoryResult:
    """One check module's (A-H) own output: its findings, how many
    subjects it actually compared against a primary source, and the
    primary source(s) it used -- `report.py`'s per-category table reads
    every field here.
    """

    code: str
    label: str
    primary_source: str
    findings: list[Finding] = field(default_factory=list)
    #: How many subjects (items, quests, bosses, picks...) this category
    #: actually compared against a primary source -- the denominator the
    #: report's counts are read against.
    checked: int = 0
    #: Non-empty when part of this category's own brief could not be
    #: verified against a primary source this run (a raw table absent
    #: locally, a dump not passed) -- named, not silently dropped.
    skipped: str = ""

    def add(
        self,
        severity: str,
        subject: str,
        message: str,
        *,
        ours: str = "",
        theirs: str = "",
        source: str = "",
    ) -> None:
        self.findings.append(
            Finding(
                category=self.code,
                severity=severity,
                subject=str(subject),
                message=message,
                ours=ours,
                theirs=theirs,
                source=source or self.primary_source,
            )
        )

    def counts_by_severity(self) -> dict[str, int]:
        counts = dict.fromkeys(SEVERITIES, 0)
        for finding in self.findings:
            counts[finding.severity] += 1
        return counts
