# data/tests/test_audit_addon.py
"""Check G (addon parity) -- against a fresh regeneration of the active build
(the `regenerated_addon_build` fixture in conftest.py), plus a synthetic
drifted `addon-data.json` to prove the check module actually notices.

The committed addon-data.json and Data.lua trail their curated inputs until
data.yml's `addon-data` job (on a push) or the nightly bis.yml regenerates
them, so this suite never asserts the committed copies are in step."""

import json
from pathlib import Path

from pipeline.addondata import check_addon_data
from pipeline.addonlua import lua_has_drifted
from pipeline.audit import check_addon
from pipeline.audit.context import AuditContext


def test_clean_build_has_no_drift_findings(regenerated_addon_build):
    fresh = regenerated_addon_build
    ctx = AuditContext(
        fresh.build, root=fresh.root, curated_dir=Path("curated"), lua_path=fresh.lua_path
    )
    result = check_addon.check(ctx)
    assert result.checked >= 2
    # This assertion mirrors the exact functions CI's own `addon-data --check`
    # runs; on a clean build the audit must say so too.
    assert not check_addon_data(fresh.build, root=fresh.root)
    assert not lua_has_drifted(fresh.build, root=fresh.root, lua_path=fresh.lua_path)
    assert not any("drifted" in f.message for f in result.findings)


def test_drifted_addon_data_is_a_blocker(regenerated_addon_build, tmp_path):
    import shutil

    fresh = regenerated_addon_build
    root = tmp_path / "builds"
    shutil.copytree(fresh.root / fresh.build, root / fresh.build)
    # Corrupt the regenerated copy so it disagrees with what build_addon_data
    # regenerates from the same talents/bis inputs; the weights are the ONLY
    # difference.
    data_path = root / fresh.build / "addon-data.json"
    data = json.loads(data_path.read_text(encoding="utf-8"))
    data["weights"] = {}
    data_path.write_text(json.dumps(data), encoding="utf-8")

    ctx = AuditContext(fresh.build, root=root, curated_dir=Path("curated"), lua_path=fresh.lua_path)
    result = check_addon.check(ctx)
    messages = [f.message for f in result.findings if f.severity == "blocker"]
    assert any(m.startswith("addon-data.json has drifted from") for m in messages), messages
    # The same inputs regenerate the scratch Data.lua exactly.
    assert not any("Data.lua" in m for m in messages), messages
