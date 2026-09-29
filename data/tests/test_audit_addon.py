# data/tests/test_audit_addon.py
"""Check G (addon parity) -- against the real committed build (same
convention `tests/test_addondata.py` already uses: `build_addon_data`
reads `builds/<BUILD>` relative to the test run's own cwd, so no synthetic
fixture is needed for the "clean" case), plus a synthetic drifted
`addon-data.json` to prove the check module actually notices."""

import json
from pathlib import Path

from pipeline.addondata import check_addon_data
from pipeline.addonlua import lua_has_drifted
from pipeline.audit import check_addon
from pipeline.audit.context import AuditContext

BUILD = "1.60.1.70009"


def test_clean_build_has_no_drift_findings():
    ctx = AuditContext(BUILD)
    result = check_addon.check(ctx)
    assert result.checked >= 2
    # This assertion mirrors the exact functions CI's own `addon-data --check`
    # runs; if the committed build is clean, the audit must say so too.
    assert not check_addon_data(BUILD)
    assert not lua_has_drifted(BUILD)
    assert not any("drifted" in f.message for f in result.findings)


def test_drifted_addon_data_is_a_blocker(tmp_path):
    real = Path("builds") / BUILD
    build_dir = tmp_path / "builds" / BUILD
    build_dir.mkdir(parents=True)
    for name in ("classes.json", "addon-data.json"):
        (build_dir / name).write_text((real / name).read_text(encoding="utf-8"))
    import shutil

    shutil.copytree(real / "talents", build_dir / "talents")
    shutil.copytree(real / "bis", build_dir / "bis")
    # Corrupt the committed copy so it disagrees with what build_addon_data
    # would regenerate from the same talents/bis inputs.
    data = json.loads((build_dir / "addon-data.json").read_text(encoding="utf-8"))
    data["weights"] = {}
    (build_dir / "addon-data.json").write_text(json.dumps(data))

    ctx = AuditContext(BUILD, root=tmp_path / "builds", curated_dir=Path("curated"))
    result = check_addon.check(ctx)
    assert any(f.severity == "blocker" and "drifted" in f.message for f in result.findings)
