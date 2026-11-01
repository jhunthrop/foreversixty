"""Each spell generator reads the layered push (sim-accuracy round 6, pipeline lane).

The fixture is the real shipped 1.60.1.70707 rows of Ice Barrier (11426),
Frigid Power (1326891) and Touch of the Grave (1260189), with the real
entries wago.tools recorded for them in the 70707 push (push 112753 still
says 3; push 113486 says 5 and 15). Round 5 found these values by hand.

The engine fork's relicmods, periodiccrit and eureka generators read the
site's `raw/<Table>.csv` directly, so the in-place layer reaches them the way
it reaches `read_csv` here; test_hotfix_push covers that the layer is written
to `raw/` itself.
"""

from __future__ import annotations

import csv
from pathlib import Path

import pytest

from pipeline import __main__ as cli
from pipeline.csvio import read_csv
from pipeline.hotfix_push import STATUS_REMOVED, STATUS_VALID, PushEntry, layer_hotfix_push
from pipeline.simconst import build_spell_constants
from pipeline.spelltext import ExtraRows, load_spell_text
from pipeline.trainables import build_trainables
from tests.test_hotfix_push import BUILD, build_root, fixture_push  # noqa: F401

ICE_BARRIER, FRIGID_POWER, TOUCH_OF_THE_GRAVE = "11426", "1326891", 1260189


def layered(root: Path) -> Path:
    layer_hotfix_push(BUILD, root=root, fetch=lambda b, t: fixture_push())
    return root / BUILD / "raw"


def mage_amounts(raw: Path) -> dict[str, list[int]]:
    mage = next(r for r in build_spell_constants(BUILD, raw) if r.class_slug == "mage")
    return {
        spell_id: [e.amount for e in mage.spells[spell_id].effects if e.aura == 108]
        for spell_id in (ICE_BARRIER, FRIGID_POWER)
    }


def test_simconst_reads_ice_barrier_and_frigid_power_from_the_push(build_root):  # noqa: F811
    assert mage_amounts(build_root / BUILD / "raw") == {ICE_BARRIER: [3, 3], FRIGID_POWER: [3, 3]}
    assert mage_amounts(layered(build_root)) == {ICE_BARRIER: [5, 5], FRIGID_POWER: [15, 15]}


def proc_cooldown(raw: Path) -> str:
    """The `$proccooldown` token through `describe`: the shipped description of
    Touch of the Grave does not print it, so the test row asks for it."""
    spells = [
        {**row, "Description_lang": "once every $proccooldown sec"}
        if int(row["ID"]) == TOUCH_OF_THE_GRAVE
        else row
        for row in read_csv(raw / "Spell.csv")
    ]
    text = load_spell_text(
        spells,
        read_csv(raw / "SpellMisc.csv"),
        read_csv(raw / "SpellEffect.csv"),
        read_csv(raw / "SpellDuration.csv"),
        ExtraRows(aura_options=read_csv(raw / "SpellAuraOptions.csv")),
    )
    return text.describe(TOUCH_OF_THE_GRAVE)


def test_spell_text_reads_touch_of_the_graves_cooldown_from_the_push(build_root):  # noqa: F811
    assert proc_cooldown(build_root / BUILD / "raw") == "once every 25 sec"
    assert proc_cooldown(layered(build_root)) == "once every 15 sec"


# --- trainables: a skill line row the push removes or adds -------------------

FROSTBOLT_1, FROSTBOLT_2, FROSTBOLT_3 = 116, 205, 837
FROST_LINE, MAGE_MASK, MAGE_FAMILY = 6, 128, 3


def _write(path: Path, header: list[str], rows: list[list[object]]) -> None:
    with path.open("w", newline="", encoding="utf-8") as handle:
        writer = csv.writer(handle)
        writer.writerow(header)
        writer.writerows(rows)


@pytest.fixture
def mage_root(tmp_path: Path) -> Path:
    raw = tmp_path / BUILD / "raw"
    raw.mkdir(parents=True)
    spells = {FROSTBOLT_1: ("Rank 1", 4), FROSTBOLT_2: ("Rank 2", 8), FROSTBOLT_3: ("Rank 3", 14)}
    _write(raw / "SpellName.csv", ["ID", "Name_lang"], [[i, "Frostbolt"] for i in spells])
    _write(raw / "Spell.csv", ["ID", "NameSubtext_lang"], [[i, s[0]] for i, s in spells.items()])
    _write(raw / "SpellMisc.csv", ["ID", "DifficultyID", "SpellID", "CastingTimeIndex"],
           [[i, 0, i, 1] for i in spells])  # fmt: skip
    _write(raw / "SpellCastTimes.csv", ["ID", "Base"], [[1, 2500]])
    _write(raw / "SpellLevels.csv", ["ID", "DifficultyID", "SpellID", "SpellLevel"],
           [[i, 0, i, s[1]] for i, s in spells.items()])  # fmt: skip
    _write(raw / "SpellPower.csv", ["ID", "SpellID", "ManaCost", "PowerType", "PowerCostPct"],
           [[i, i, 25, 0, 0] for i in spells])  # fmt: skip
    _write(raw / "SpellCooldowns.csv",
           ["ID", "DifficultyID", "SpellID", "RecoveryTime", "CategoryRecoveryTime"],
           [[i, 0, i, 0, 0] for i in spells])  # fmt: skip
    # Ranks 2 and 3 are in no class table: only a skill line row teaches them.
    _write(raw / "SpellClassOptions.csv", ["ID", "SpellID", "SpellClassSet"],
           [[0, FROSTBOLT_1, MAGE_FAMILY]])  # fmt: skip
    _write(
        raw / "SkillLine.csv", ["ID", "DisplayName_lang", "CategoryID"], [[FROST_LINE, "Frost", 7]]
    )
    _write(raw / "SkillLineAbility.csv",
           ["ID", "SkillLine", "Spell", "ClassMask", "SupercedesSpell", "AcquireMethod"],
           [[1, FROST_LINE, FROSTBOLT_1, MAGE_MASK, 0, 0],
            [2, FROST_LINE, FROSTBOLT_2, MAGE_MASK, FROSTBOLT_1, 0]])  # fmt: skip
    return tmp_path


def frostbolt_ranks(raw: Path) -> list[int]:
    mage = next(r for r in build_trainables(BUILD, raw) if r.class_slug == "mage")
    return [rank.id for t in mage.trainables if t.name == "Frostbolt" for rank in t.ranks]


def test_trainables_follow_a_skill_line_row_the_push_removes_or_adds(mage_root):
    push = {
        "SkillLineAbility": [
            PushEntry(1, 113000, 2, STATUS_REMOVED, None),
            PushEntry(2, 113000, 3, STATUS_VALID, [3, FROST_LINE, FROSTBOLT_3, MAGE_MASK, 0, 0]),
        ]
    }
    assert frostbolt_ranks(mage_root / BUILD / "raw") == [FROSTBOLT_1, FROSTBOLT_2]
    layer_hotfix_push(BUILD, root=mage_root, fetch=lambda b, t: push)
    assert frostbolt_ranks(mage_root / BUILD / "raw") == [FROSTBOLT_1, FROSTBOLT_3]


# --- the CLI: fetch never leaves a build shipped-only ------------------------


def test_fetch_layers_the_push_after_the_download(monkeypatch, tmp_path):
    calls: list[tuple] = []
    monkeypatch.setattr(
        "pipeline.wago.fetch_build",
        lambda p, b: calls.append(("fetch", p, b)) or tmp_path / BUILD,
    )
    monkeypatch.setattr(
        "pipeline.hotfix_push.layer_hotfix_push",
        lambda build, offline=False: calls.append(("layer", build, offline)) or {"tables": {}},
    )
    cli.main(["fetch", "--product", "wowf", "--build", BUILD])
    assert calls == [("fetch", "wowf", BUILD), ("layer", BUILD, False)]


def test_layer_hotfix_push_has_an_explicit_offline_mode(monkeypatch):
    calls: list[tuple] = []
    monkeypatch.setattr(
        "pipeline.hotfix_push.layer_hotfix_push",
        lambda build, offline=False: calls.append((build, offline)) or {"tables": {}},
    )
    cli.main(["layer-hotfix-push", "--build", BUILD, "--offline"])
    assert calls == [(BUILD, True)]
