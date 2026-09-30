# data/tests/test_audit_curated.py
"""Check H (curated specs vs engine) -- the real committed curated/specs.json
(already valid; `pipeline.specs.load_specs` enforces that at load time) plus
a synthetic specs.json with a bad weight_stats id and a missing apl file."""

import json

from pipeline.audit import check_curated
from pipeline.audit.context import AuditContext


def test_real_curated_specs_pass_and_have_apl_files():
    ctx = AuditContext("1.60.1.70009")
    result = check_curated.check(ctx)
    assert result.checked > 1
    blockers = [f for f in result.findings if f.severity == "blocker"]
    # simdb-supplement, 2026-09-30: the pick-coverage blocker reads the
    # committed simitems.json, which only carries the classic-db rows once
    # data.yml regenerates it (and data.yml tests BEFORE it regenerates), so
    # a pre-regen file (no `sim_source` map yet) is exempt from that one.
    simitems_path = ctx.build_dir / "simitems.json"
    pre_regen = "sim_source" not in json.loads(simitems_path.read_text(encoding="utf-8"))
    if pre_regen:
        blockers = [f for f in blockers if "not in simitems.json" not in f.message]
    assert blockers == []


def test_unknown_stat_id_is_a_blocker(tmp_path):
    curated_dir = tmp_path / "curated"
    curated_dir.mkdir()
    (curated_dir / "specs.json").write_text(
        json.dumps(
            [
                {
                    "spec": "warrior-fury", "class_slug": "warrior", "spec_slug": "fury",
                    "name": "Fury", "role": "dps", "tree_index": 2,
                    "reference_stat": "not_a_real_stat", "weight_stats": ["not_a_real_stat"],
                    "icon": "icon",
                }
            ]
        )
    )
    ctx = AuditContext("testbuild", curated_dir=curated_dir)
    result = check_curated.check(ctx)
    finding = next(f for f in result.findings)
    assert finding.severity == "blocker"
    assert "not_a_real_stat" in finding.message


def test_missing_apl_file_is_major(tmp_path):
    curated_dir = tmp_path / "curated"
    (curated_dir / "apl").mkdir(parents=True)
    (curated_dir / "specs.json").write_text(
        json.dumps(
            [
                {
                    "spec": "warrior-fury", "class_slug": "warrior", "spec_slug": "fury",
                    "name": "Fury", "role": "dps", "tree_index": 2,
                    "reference_stat": "attack_power", "weight_stats": ["attack_power"],
                    "icon": "icon",
                }
            ]
        )
    )
    ctx = AuditContext("testbuild", curated_dir=curated_dir)
    result = check_curated.check(ctx)
    finding = next(f for f in result.findings if f.subject == "warrior-fury")
    assert finding.severity == "major"
    assert "no apl file" in finding.message
