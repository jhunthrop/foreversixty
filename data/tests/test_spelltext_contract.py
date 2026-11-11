"""Every description the active build ships reads like the game client's.

Reads what is committed under builds/<active build>/ (talent ranks, spell and consumable
tooltips, item effects, set bonuses) through `pipeline.spelltext_lint`, the same reader
`pipeline spelltext-check` gates data.yml with after it re-renders a build.

A floor and a shape rule, never a pinned count (data contracts are floors): at least
`MIN_DESCRIPTIONS` descriptions, and none with an unrendered token or an empty
substitution. data.yml runs this suite BEFORE it re-renders, against the committed files,
so a build rendered by an older grammar (its manifest's `spelltext_renderer` is behind
`SPELLTEXT_RENDERER`) is skipped rather than failed: failing it would block the very run
that re-renders it.
"""

from pathlib import Path

import pytest

from conftest import ACTIVE_BUILD
from pipeline.manifest import read_manifest
from pipeline.spelltext import SPELLTEXT_MANIFEST_KEY, SPELLTEXT_RENDERER
from pipeline.spelltext_lint import (
    CLIENT_VERBATIM,
    defects_in,
    scan_build,
    shipped_texts,
    summarize,
)

BUILD_DIR = Path("builds") / ACTIVE_BUILD

#: 10,670 on 1.60.1.70790; a floor so an empty reader cannot pass.
MIN_DESCRIPTIONS = 5_000


def test_the_active_build_ships_descriptions():
    assert sum(1 for _ in shipped_texts(BUILD_DIR)) >= MIN_DESCRIPTIONS


def test_every_shipped_description_is_fully_rendered():
    rendered_with = read_manifest(BUILD_DIR).get(SPELLTEXT_MANIFEST_KEY, 0)
    if rendered_with < SPELLTEXT_RENDERER:
        pytest.skip(
            f"{ACTIVE_BUILD} was rendered with spelltext grammar {rendered_with}, before "
            f"{SPELLTEXT_RENDERER}; data.yml re-renders it and `spelltext-check` gates that run"
        )
    _, findings = scan_build(BUILD_DIR)
    shown = [f"{finding.where}: {finding.text[:120]}" for finding in findings[:20]]
    assert not findings, (summarize(findings), shown)


@pytest.mark.parametrize(
    ("text", "defects"),
    [
        ("Increases your Strength by $s1%.", ["unrendered $ token"]),
        ("dealing ${32+$rap} damage", ["unrendered $ token", "stray brace"]),
        (
            "Engrave your pants:\r\n\r\n$@spelldesc407624",
            ["unrendered $ token", "unrendered @ token"],
        ),
        ("|CFFFF2020Requires Cat Form|R\r\nPounce.", ["client colour code"]),
        ("reduces the Rage cost by of your Shouts", ["empty substitution"]),
        ("heal the target for .", ["empty substitution"]),
        ("over 3 sec  and slowing", ["double space mid-sentence"]),
        ("a 0% chance to immobilize", ["zero amount"]),
        ("reduced by -0.", ["negative zero"]),
        ("lasts 0 sec.", ["zero duration"]),
    ],
)
def test_each_defect_class_is_caught(text, defects):
    assert defects_in(text) == defects


@pytest.mark.parametrize(
    "text",
    [
        # The client's own two spaces after a full stop, and the combo-point tables' indent.
        "Absorbs 308 physical damage.  Lasts 15 sec.",
        "Lasts longer per combo point:\r\n   1 point  : 14 seconds",
        "increases the damage it deals by 2%.",
        "a 10% chance",
        "Reduces the cast time by 0.5 sec.",
        "(32 + Ranged Attack Power * 0.05) Physical damage",
    ],
)
def test_clean_client_text_is_not_a_defect(text):
    assert defects_in(text) == []


def test_a_client_typo_is_not_counted():
    """Booming Voice 12321's "by of" is the client's own text: Spell.Description_lang on
    1.60.1.70790 reads "reduces the Rage cost by of your Shouts and Howls by $s2%." and
    Wowhead's Forever tooltip prints the same (nether.wowhead.com/tooltip/spell/12321
    ?dataEnv=16, 2026-11-10)."""
    text = (
        "Increases the area of effect of your Shouts by 10% and reduces the Rage cost by of "
        "your Shouts and Howls by 5%."
    )
    assert defects_in(text) == []
    assert all(source for source in CLIENT_VERBATIM.values())
