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

# The active build, so this test follows every client move instead of pinning one.
BUILD = json.loads(
    (Path(__file__).resolve().parents[2] / "web" / "src" / "data" / "active-build.json").read_text()
)["build"]


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
    import shutil

    # Every input build_addon_data reads, so the corrupted weights below are
    # the ONLY difference (fix round 1, item 6: with spellranks.json missing
    # the regeneration failed instead, and the test passed only because the
    # temp directory's own name contains "drifted"). Copied as bytes, so
    # nothing goes through the platform's default text encoding.
    for name in (
        "classes.json",
        "addon-data.json",
        "spellranks.json",
        "items.json",
        "loot.json",
        "addon-builds.json",
        "dps-per-point.json",
    ):
        if (real / name).exists():
            shutil.copyfile(real / name, build_dir / name)
    shutil.copytree(real / "talents", build_dir / "talents")
    shutil.copytree(real / "items", build_dir / "items")
    if (real / "bis").exists():  # a build the nightly has not ranked yet has none
        shutil.copytree(real / "bis", build_dir / "bis")
    # Corrupt the committed copy so it disagrees with what build_addon_data
    # would regenerate from the same talents/bis inputs.
    data = json.loads((build_dir / "addon-data.json").read_text(encoding="utf-8"))
    data["weights"] = {}
    (build_dir / "addon-data.json").write_text(json.dumps(data), encoding="utf-8")

    ctx = AuditContext(BUILD, root=tmp_path / "builds", curated_dir=Path("curated"))
    result = check_addon.check(ctx)
    messages = [f.message for f in result.findings if f.severity == "blocker"]
    assert any(m.startswith("addon-data.json has drifted from") for m in messages), messages
    # The same inputs regenerate the committed Data.lua exactly.
    assert not any("Data.lua" in m for m in messages), messages
