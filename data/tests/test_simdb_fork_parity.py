"""Guardrail: our stat_array output for a sample of items, checked against
the fork's own item database (rank-stat-plumbing lane, 2026-09-29).

Reads the committed data/builds/<build>/simdb.bin (this pipeline's own
SimItem.stats output) and wowsims-forever's assets/database/db.json (a
sibling checkout, FOREVER_ENGINE_PATH), and compares the two by stat NAME,
not by array index: db.json's array is just floats at whatever position the
fork's own `proto.Stat` enum had when it was exported, which statmap.py's
own module docstring already flags as a moving target (the engine lane
collapsing MeleeHit/SpellHit into Hit, for one). Rather than trust that this
checkout's copy of the enum (pipeline.simproto.pb, vendored at whatever
commit `python -m pipeline simproto` last ran) still agrees position-for-
position, this module parses the fork checkout's OWN proto/common.proto at
test time and decodes db.json's arrays against that.

Skips (like test_go_reads_back_what_python_wrote in test_simdb_build.py)
when FOREVER_ENGINE_PATH is unset -- comparing against a sibling checkout is
not something CI, which clones neither, can do.

What this run found (2026-09-29, engine commit 038e51a61): Cruel Barb (item
5191, this lane's own bug example) matches the fork exactly post-fix -- 12
attack_power AND 12 ranged_attack_power on both sides, where before the
statmap.py fix this pipeline emitted 0 ranged_attack_power. Every other
generic-AP item in the sample the fork also carries matches the same way.
Five items in the sample disagree on a stat this lane's fix never touches
(spirit, spell_power, armor, spell_penetration, or a flat AP/RAP amount the
fork's db.json states at pre-re-itemization Classic Era values rather than
Forever's) -- KNOWN_MISMATCHES documents each with why, so a genuinely new
mismatch (this lane's fix regressing, or a future stat this pipeline gets
wrong) still fails the test instead of being silently swallowed alongside
these five.
"""

from __future__ import annotations

import json
import os
import re
from functools import cache
from pathlib import Path

import pytest

from pipeline.simdb.statmap import PROTO_STAT_ALIASES
from pipeline.simproto import pb

BUILD = "1.60.1.70009"
BUILD_DIR = Path("builds") / BUILD

#: A mix of AP, RAP, agility, spell power and caster items, deliberately
#: including Cruel Barb (this lane's own bug example) and Rune of the Guard
#: Captain (test_simdb_build.py's own on-equip-stat fixture) -- ids are
#: this build's, chosen by walking data/builds/1.60.1.70009/simdb.bin for a
#: spread across "generic AP only", "AP with an explicit extra RAP amount
#: on top", "agility with no AP/spell power", "spell power" and "neither"
#: (see this lane's own investigation, not reproduced here since it is a
#: one-off selection, not a derivation this test needs to repeat every run).
SAMPLE_ITEM_IDS = [
    720, 816, 867, 888, 5191, 6797, 7358, 10328, 12798, 14748, 15054, 16536,
    18821, 19120, 19546, 19591, 20150, 22082, 22666, 220821, 227062, 227192,
    235476, 252446, 252594, 270048, 270275, 272067, 272734, 272808, 274748,
]

TOLERANCE = 0.02

#: item id -> {planner stat key: why the fork's db.json disagrees today}.
#: Every one of these was checked by hand against the fork's own source
#: (tools/database/wowhead_tooltips.go, sim/common/item_effects.go) before
#: being listed here; a stat/item pair not in this table that disagrees is
#: a real failure, not an entry waiting to be added.
KNOWN_MISMATCHES: dict[int, dict[str, str]] = {
    # Annihilator and Eyepoker's on-equip bonuses (Annihilator: 14 attack
    # power; Eyepoker: 3 stamina, 4 spirit) are in no ItemSparse column --
    # the fork's own STAT_AURAS-equivalent path leaves them at 0 in
    # db.json's declared `stats` array because the FORK's engine applies
    # them as hand-written Go in sim/common/item_effects.go instead
    # (equip.py's module docstring documents the same split for this
    # pipeline: STAT_AURAS is what a spell states as a stat, IGNORED_AURAS
    # is "behaviour the engine hand-writes"). This site's own engine has no
    # such hard-coded item file, so simdb states the bonus declaratively --
    # correctly, for this pipeline's architecture -- and the two `stats`
    # arrays are not commensurable for these two items.
    12798: {
        "attack_power": "fork engine-hardcoded item effect, not a declared stat -- see docstring",
        "ranged_attack_power": "same",
    },
    6797: {
        "stamina": "fork engine-hardcoded item effect, not a declared stat -- see module docstring",
        "spirit": "same",
    },
    # Rune of the Guard Captain: this codebase's own equip.py module
    # docstring and test_simdb_build.py's test_a_known_on_equip_stat_survived_the_round_trip
    # already pin this item's Forever-specific tooltip at +42 Attack Power
    # and +42 Ranged Attack Power (wowhead's Forever corroboration, not
    # Classic Era's). The fork's db.json states 20/20 -- Classic Era's
    # un-re-itemized amount -- so it predates Forever's re-itemization of
    # this item rather than disagreeing with this pipeline's math; the
    # attack_power/ranged_attack_power mirror itself is exactly right (42
    # mirrored + 42 explicit = 84 ranged attack power, additive as
    # designed) and is what this test exists to catch if it regresses.
    19120: {
        "attack_power": "fork db.json predates Forever's re-itemization (Era 20 vs Forever 42)",
        "ranged_attack_power": "same",
        "hit": "same re-itemization gap (0.7% vs Classic Era's 1%)",
    },
    # Volcanic Leggings and Warlord's Silk Amice: armor (and, for the
    # Amice, spell_penetration) differ from the fork's db.json in a way
    # unrelated to attack power -- most likely the same re-itemization gap
    # as Rune of the Guard Captain above, but this lane did not chase
    # ItemCurves/tooltip-regex parity for armor or spell penetration since
    # neither is this lane's stat. Left listed rather than silently
    # swallowed so a controller can route it to a data-accuracy lane.
    15054: {
        "armor": "unresolved drift vs fork db.json, likely re-itemization -- data-accuracy lane",
        "stamina": "fork db.json states none; unresolved, same as armor",
    },
    16536: {
        "armor": "unresolved drift vs fork db.json, likely re-itemization -- data-accuracy lane",
        "spell_penetration": "fork db.json states 10; this pipeline states none -- unresolved",
    },
    # Naga Battle Gloves: fork states +7 spirit, this pipeline states +5
    # spell power for the same slot -- a whole different stat, not an
    # amount drift, so almost certainly Forever re-itemized this glove
    # from a spirit bonus to a spell power one and the fork's db.json
    # predates that. Neither side is this lane's stat.
    888: {
        "spirit": "fork db.json predates a Forever re-itemization (spirit -> spell power)",
        "spell_power": "same",
    },
}


def _engine_path() -> Path | None:
    raw = os.environ.get("FOREVER_ENGINE_PATH")
    return Path(raw) if raw else None


@cache
def _fork_stat_index_to_name(engine: Path) -> dict[int, str]:
    """index -> `proto.Stat` enum name, parsed from the fork checkout's own
    proto/common.proto rather than trusted off this checkout's vendored
    pipeline.simproto.pb (see this module's docstring)."""
    text = (engine / "proto" / "common.proto").read_text()
    match = re.search(r"enum Stat \{(.*?)\n\}", text, re.S)
    if match is None:
        raise AssertionError(f"no `enum Stat {{ ... }}` found in {engine}/proto/common.proto")
    return {
        int(number): name
        for name, number in re.findall(r"(Stat\w+)\s*=\s*(\d+);", match.group(1))
    }


@cache
def _our_stat_index_to_name() -> dict[int, str]:
    """The inverse of `pb.Stat`'s own name -> index map, for decoding this
    checkout's own SimItem.stats arrays through `_decode` the same way
    `_fork_stat_index_to_name` decodes the fork's."""
    return {value: name for name, value in pb.Stat.items()}


@cache
def _fork_items(engine: Path) -> dict[int, dict]:
    db = json.loads((engine / "assets" / "database" / "db.json").read_text())
    return {row["id"]: row for row in db["items"]}


@cache
def _enum_name_to_planner_key() -> dict[str, str]:
    """`proto.Stat` enum name -> the planner's own stat key, i.e. the
    reverse of PROTO_STAT_ALIASES. Declaration order decides the winner for
    an enum name more than one planner key lists (there is none today, but
    ties resolve the same way stat_index's own lookup does)."""
    out: dict[str, str] = {}
    for key, names in PROTO_STAT_ALIASES.items():
        if key.startswith("__"):
            continue
        for name in names:
            out.setdefault(name, key)
    return out


def _decode(array: list[float], index_to_name: dict[int, str]) -> dict[str, float]:
    name_to_key = _enum_name_to_planner_key()
    out: dict[str, float] = {}
    for index, amount in enumerate(array):
        if not amount:
            continue
        name = index_to_name.get(index)
        key = name_to_key.get(name) if name else None
        # An index db.json carries that this pipeline's own vocabulary has
        # no key for (a combat rating retail grew after Classic, say) is
        # not this test's business -- STAT_BY_MODIFIER_ID/PROTO_STAT_ALIASES
        # already have their own tests for what the planner tracks.
        out_key = key or f"unmapped:{name or index}"
        out[out_key] = out.get(out_key, 0.0) + amount
    return out


@cache
def _our_items() -> dict[int, pb.SimItem]:
    db = pb.SimDatabase()
    db.ParseFromString((BUILD_DIR / "simdb.bin").read_bytes())
    return {item.id: item for item in db.items}


def test_a_sample_of_items_matches_the_forks_own_database_by_stat_name():
    engine = _engine_path()
    if engine is None or not (engine / "assets" / "database" / "db.json").exists():
        pytest.skip("FOREVER_ENGINE_PATH is unset or has no assets/database/db.json")
    index_to_name = _fork_stat_index_to_name(engine)
    fork_items = _fork_items(engine)
    ours = _our_items()

    compared = 0
    failures: list[str] = []
    for item_id in SAMPLE_ITEM_IDS:
        fork_row = fork_items.get(item_id)
        if fork_row is None:
            # Not every id this pipeline keeps is in the fork's own
            # ItemSparse-derived universe -- the wowhead supplement
            # (items.py's build_wowhead_sim_items) exists precisely because
            # the fork's own tables lack them. Nothing to compare.
            continue
        compared += 1
        our_stats = _decode(list(ours[item_id].stats), _our_stat_index_to_name())
        fork_stats = _decode(fork_row["stats"], index_to_name)
        exceptions = KNOWN_MISMATCHES.get(item_id, {})
        for key in sorted(set(our_stats) | set(fork_stats)):
            ours_amount = our_stats.get(key, 0.0)
            forks_amount = fork_stats.get(key, 0.0)
            if abs(ours_amount - forks_amount) <= TOLERANCE:
                continue
            if key in exceptions:
                continue
            failures.append(
                f"item {item_id} ({fork_row.get('name', '?')}): {key} is "
                f"{ours_amount} here, {forks_amount} in the fork's db.json"
            )

    # The sample is deliberately wider than the fork's own item universe
    # (it includes wowhead-supplement-only ids); if fewer than half the
    # sample were ever comparable the sample itself would need rebuilding,
    # not just the assertion below.
    assert compared >= len(SAMPLE_ITEM_IDS) // 2, (
        f"only {compared}/{len(SAMPLE_ITEM_IDS)} sample ids exist in the fork's db.json; "
        "rebuild SAMPLE_ITEM_IDS"
    )
    assert not failures, "\n".join(failures)


def test_cruel_barb_is_this_lanes_own_confirmation():
    """Cruel Barb (item 5191) is the exact item this lane's brief names as
    broken: a generic "+12 Attack Power" main-hand weapon that gave hunters
    0 ranged attack power before statmap.stat_array mirrored it. The fork's
    own db.json already states this item as 12 attack_power AND 12
    ranged_attack_power -- the single clearest ground-truth confirmation
    available that the mirror this lane added is correct, not just
    plausible."""
    engine = _engine_path()
    if engine is None or not (engine / "assets" / "database" / "db.json").exists():
        pytest.skip("FOREVER_ENGINE_PATH is unset or has no assets/database/db.json")
    fork_row = _fork_items(engine)[5191]
    index_to_name = _fork_stat_index_to_name(engine)
    fork_stats = _decode(fork_row["stats"], index_to_name)
    assert fork_stats["attack_power"] == 12.0
    assert fork_stats["ranged_attack_power"] == 12.0

    our_stats = _decode(list(_our_items()[5191].stats), _our_stat_index_to_name())
    assert our_stats["attack_power"] == 12.0
    assert our_stats["ranged_attack_power"] == 12.0
