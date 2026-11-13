import json
import re
from pathlib import Path

import pytest

from pipeline.addonrotation import (
    CONDITION_MAX_CHARS,
    LEVEL_BANDS,
    AddonRotationError,
    _cast_lines,
    _one_line,
    apl_cast_spell_ids,
    build_rotation,
    build_rotations,
    coach_persona_spell_ids,
    coach_spell_ids,
    compute_spell_icons,
    load_talent_spells,
    load_trainable_ids,
    raid_cover_spell_ids,
    rank_siblings,
)

BUILD = "1.60.1.70374"

SPELLRANKS = {
    "build": BUILD,
    "classes": {
        "warrior": {
            "Battle Shout": [
                {"id": 6673, "level": 1, "rank": 1},
                {"id": 5242, "level": 12, "rank": 2},
                {"id": 25289, "level": 60, "rank": 7},
            ],
            "Bloodrage": [{"id": 2687, "level": 1, "rank": 0}],
            "Bloodthirst": [
                {"id": 23881, "level": 40, "rank": 1},
                {"id": 23894, "level": 60, "rank": 4},
            ],
            "Whirlwind": [{"id": 1680, "level": 36, "rank": 0}],
        }
    },
}


def _write_apl(curated_dir: Path, spec: str, priority_list: list[dict]) -> None:
    apl_dir = curated_dir / "apl"
    apl_dir.mkdir(parents=True, exist_ok=True)
    (apl_dir / f"{spec}.json").write_text(
        json.dumps({"spec": spec, "rotation": {"priorityList": priority_list}}),
        encoding="utf-8",
    )


def _bands(spec: str, curated_dir: Path) -> dict:
    """spec's rotation, keyed by level, against the test module's SPELLRANKS."""
    return {band.level: band.lines for band in build_rotation(spec, curated_dir, SPELLRANKS)}


def _cast(spell_id: int, rank: int | None = None, notes: str = "") -> dict:
    spell = {"spellId": spell_id} if rank is None else {"spellId": spell_id, "rank": rank}
    return {"notes": notes, "action": {"castSpell": {"spellId": spell}}}


def test_level_bands_match_the_ladders_own_rungs():
    """sim/request/ladder.go's ladderLevels is the one source; this module's
    LEVEL_BANDS is a copy (Go and Python do not import each other), and this
    test is what keeps the copy honest."""
    source = Path("../sim/request/ladder.go").read_text(encoding="utf-8")
    match = re.search(r"var ladderLevels = \[\]int\{([^}]*)\}", source)
    assert match, "ladderLevels literal not found in ladder.go"
    levels = [int(part.strip()) for part in match.group(1).split(",")]
    assert LEVEL_BANDS == levels


def test_a_line_resolves_to_the_highest_rank_learned_at_each_band(tmp_path):
    _write_apl(
        tmp_path,
        "warrior-fury",
        [
            _cast(2687, notes="Bloodrage on cooldown."),
            _cast(25289, rank=7, notes="Battle Shout up."),
        ],
    )
    bands = _bands("warrior-fury", tmp_path)
    assert [line.spell_id for line in bands[10]] == [2687, 6673]
    assert [line.spell_id for line in bands[20]] == [2687, 5242]
    assert [line.spell_id for line in bands[60]] == [2687, 25289]


def test_an_ability_not_learned_yet_is_left_out_of_that_band_only(tmp_path):
    _write_apl(tmp_path, "warrior-fury", [_cast(23894, rank=4, notes="Bloodthirst on cooldown.")])
    bands = _bands("warrior-fury", tmp_path)
    assert bands[30] == []
    assert [line.spell_id for line in bands[40]] == [23881]
    assert [line.spell_id for line in bands[60]] == [23894]


#: Recklessness as the 70707 client lists it: the trainable 1719 at 50, and a
#: copy (13847) at 22 with no SkillLineAbility row.
_RECKLESSNESS_RANKS = {
    "build": BUILD,
    "classes": {
        "warrior": {
            "Recklessness": [{"id": 1719, "level": 50}, {"id": 13847, "level": 22}],
        }
    },
}


def test_an_unranked_ability_joins_a_band_at_its_trainable_ids_level(tmp_path):
    """Round 5 integration: the addon told a level-30 Fury warrior to press
    Recklessness, which is trained at 50, because a copy with no trainer row
    sits at 22. The class's trainables (SkillLineAbility) decide which id of
    an unranked family the player learns, and when."""
    _write_apl(tmp_path, "warrior-fury", [_cast(1719, notes="Recklessness.")])
    bands = {
        band.level: band.lines
        for band in build_rotation(
            "warrior-fury", tmp_path, _RECKLESSNESS_RANKS, trainable_ids=frozenset({1719})
        )
    }
    assert bands[40] == []
    assert [line.spell_id for line in bands[50]] == [1719]


def test_an_unranked_family_with_no_trainable_id_keeps_its_rows(tmp_path):
    """A rune or talent spell the trainables never list keeps resolving as before."""
    _write_apl(tmp_path, "warrior-fury", [_cast(1719, notes="Recklessness.")])
    bands = {
        band.level: band.lines
        for band in build_rotation("warrior-fury", tmp_path, _RECKLESSNESS_RANKS)
    }
    assert [line.spell_id for line in bands[30]] == [13847]


#: A Season of Discovery rune the client still carries (AcquireMethod 3): no trainer row, no talent.
_RUNE_RANKS = {
    "build": BUILD,
    "classes": {"warrior": {"Rune Strike": [{"id": 402927, "level": 20, "rank": 0}]}},
}


def test_a_family_neither_trainable_nor_a_talent_never_reaches_a_band(tmp_path):
    """The addon panel's trainables gate (design/specs/2026-11-08-addon-ux.md 4.3): once the
    class has a trainables file, a spell no trainer teaches and no talent gives is a rune a
    Forever player never learns, and no band names it."""
    _write_apl(tmp_path, "warrior-fury", [_cast(402927, notes="Rune Strike.")])
    bands = build_rotation("warrior-fury", tmp_path, _RUNE_RANKS, trainable_ids=frozenset({1719}))
    assert all(line.spell_id != 402927 for band in bands for line in band.lines)


def test_a_talent_spell_outside_the_trainables_keeps_its_rows(tmp_path):
    """Water Shield (408510) is a Restoration talent on 1.60.1.70790 that the trainables file does
    not list: the talent keeps it, from the level its tier opens."""
    _write_apl(tmp_path, "warrior-fury", [_cast(402927, notes="Rune Strike.")])
    bands = {
        band.level: band.lines
        for band in build_rotation(
            "warrior-fury",
            tmp_path,
            _RUNE_RANKS,
            trainable_ids=frozenset({1719}),
            talent_spells={"Rune Strike": [(20, 402927)]},
        )
    }
    assert [line.spell_id for line in bands[20]] == [402927]


def test_the_committed_rotations_name_only_trainable_or_talent_spells():
    """Every rotation line on the active build is a spell the class learns from a trainer or a
    talent of its tree: never a rune."""
    build = json.loads(Path("../web/src/data/active-build.json").read_text())["build"]
    for spec, bands in build_rotations(Path("builds"), build).items():
        class_slug = spec.split("-", 1)[0]
        known = load_trainable_ids(Path("builds") / build, class_slug) | {
            spell_id
            for pairs in load_talent_spells(Path("builds") / build, class_slug).values()
            for _, spell_id in pairs
        }
        for band in bands:
            for line in band.lines:
                assert line.spell_id in known or line.name in load_talent_spells(
                    Path("builds") / build, class_slug
                ), (spec, band.level, line.name, line.spell_id)


def test_the_committed_rotations_never_show_an_ability_before_its_trainer_level():
    """Every rotation line on the active build is an id the class learns from a
    trainer by that band, whenever its family has a trainable id at all."""
    build = json.loads(Path("../web/src/data/active-build.json").read_text())["build"]
    rotations = build_rotations(Path("builds"), build)
    trainable_level: dict[str, dict[int, int]] = {}
    for spec, bands in rotations.items():
        class_slug = spec.split("-", 1)[0]
        if class_slug not in trainable_level:
            path = Path("builds") / build / "trainables" / f"{class_slug}.json"
            entries = json.loads(path.read_text(encoding="utf-8"))["trainables"]
            trainable_level[class_slug] = {
                rank["id"]: rank["level"] for entry in entries for rank in entry["ranks"]
            }
        levels = trainable_level[class_slug]
        for band in bands:
            for line in band.lines:
                if line.spell_id in levels:
                    assert levels[line.spell_id] <= band.level, (spec, band.level, line)
    fury = {band.level: [line.name for line in band.lines] for band in rotations["warrior-fury"]}
    assert "Recklessness" not in fury[40]


def test_a_strict_sequence_unfolds_into_one_line_per_step_sharing_the_notes(tmp_path):
    _write_apl(
        tmp_path,
        "warrior-fury",
        [
            {
                "notes": "Opener.",
                "action": {
                    "strictSequence": {
                        "actions": [
                            {"castSpell": {"spellId": {"spellId": 2687}}},
                            {"castSpell": {"spellId": {"spellId": 1680}}},
                        ]
                    }
                },
            },
        ],
    )
    bands = _bands("warrior-fury", tmp_path)
    assert [line.spell_id for line in bands[38]] == [2687, 1680]
    assert bands[38][1].condition == "Opener."
    assert bands[30] == [] or [line.spell_id for line in bands[30]] == [2687]


def test_an_id_spellranks_does_not_track_is_skipped_rather_than_erroring(tmp_path):
    _write_apl(tmp_path, "warrior-fury", [_cast(999999, notes="A talent proc, not in spellranks.")])
    bands = _bands("warrior-fury", tmp_path)
    assert all(band == [] for band in bands.values())


def test_actions_with_no_spell_id_name_no_line(tmp_path):
    _write_apl(
        tmp_path,
        "warrior-fury",
        [
            {"action": {"autocastOtherCooldowns": {}}},
            {"notes": "Seal.", "action": {"condition": {"cmp": {}}, "castPaladinPrimarySeal": {}}},
        ],
    )
    bands = build_rotation("warrior-fury", tmp_path, SPELLRANKS)
    assert all(band.lines == [] for band in bands)


def test_every_band_from_level_bands_is_present_in_order(tmp_path):
    _write_apl(tmp_path, "warrior-fury", [])
    bands = build_rotation("warrior-fury", tmp_path, SPELLRANKS)
    assert [band.level for band in bands] == LEVEL_BANDS


def test_missing_apl_refuses(tmp_path):
    with pytest.raises(AddonRotationError, match="warrior-fury"):
        build_rotation("warrior-fury", tmp_path, SPELLRANKS)


def test_build_rotations_covers_every_curated_spec(tmp_path):
    (tmp_path / BUILD).mkdir()
    (tmp_path / BUILD / "spellranks.json").write_text(json.dumps(SPELLRANKS), encoding="utf-8")
    rotations = build_rotations(tmp_path, BUILD, curated_dir=Path("curated"))
    assert "warrior-fury" in rotations
    assert len(rotations) == len(list(Path("curated/apl").glob("*.json")))


def test_missing_spellranks_refuses(tmp_path):
    (tmp_path / BUILD).mkdir()
    with pytest.raises(AddonRotationError, match="spellranks.json"):
        build_rotations(tmp_path, BUILD, curated_dir=Path("curated"))


def test_one_line_keeps_a_single_sentence_whole():
    assert _one_line("Frostbolt is the whole rotation.") == "Frostbolt is the whole rotation."


def test_one_line_cuts_at_the_first_sentence_boundary():
    assert _one_line("First bit. Second bit that a player never needs to read.") == "First bit."


def test_one_line_caps_a_run_on_sentence_with_no_boundary():
    long_sentence = "a" * (CONDITION_MAX_CHARS + 50)
    result = _one_line(long_sentence)
    assert len(result) == CONDITION_MAX_CHARS
    assert result.endswith("…")


def test_one_line_never_cuts_a_word_in_half():
    """data-followups-10 lane, 2026-09-30, item 8: the live repro
    (hunter-marksmanship/hunter-beast-mastery's Arcane Shot line)
    published "...are the sh…", the cap landing inside "shots" - a
    last-resort cut must back off to the last whole word instead. One
    long, real-shaped run-on sentence (no early ". " boundary) built
    from ordinary words: the text before the ellipsis must be an exact
    prefix of the original, cut only at a space the original text
    itself had."""
    words = [f"word{i}" for i in range(400)]
    sentence = " ".join(words) + "."  # one giant run-on, single trailing period
    result = _one_line(sentence)
    assert result.endswith("…")
    before_ellipsis = result[:-1]
    assert sentence.startswith(before_ellipsis), (
        f"{before_ellipsis!r} is not a clean prefix of the original text"
    )
    assert sentence[len(before_ellipsis)] == " ", (
        "the cut did not land on a word boundary in the original text"
    )


def test_one_line_real_curated_notes_never_hit_the_fallback_cap():
    """data-followups-10 lane, 2026-09-30, item 8: every curated
    apl/*.json note's own first sentence, across every spec on this
    build, must fit under CONDITION_MAX_CHARS without the last-resort
    cap ever firing - the regression this pins (Arcane Shot's real
    first sentence is 534 characters with no early ". " boundary, and
    the longest across the whole corpus, druid-feral's Claw line, is
    672) is the cap itself sitting well under that real maximum."""
    apl_dir = Path(__file__).resolve().parent.parent / "curated" / "apl"
    cut_lines = []

    def walk(obj):
        if isinstance(obj, dict):
            notes = obj.get("notes")
            if isinstance(notes, str) and notes.strip():
                if _one_line(notes).endswith("…"):
                    cut_lines.append(notes[:80])
            for value in obj.values():
                walk(value)
        elif isinstance(obj, list):
            for value in obj:
                walk(value)

    for path in sorted(apl_dir.glob("*.json")):
        walk(json.loads(path.read_text(encoding="utf-8")))

    assert not cut_lines, (
        f"{len(cut_lines)} curated note(s) still hit the fallback cap: {cut_lines}"
    )


def test_a_line_carries_only_the_first_sentence_of_a_long_curated_note(tmp_path):
    long_note = (
        "Short reason. A paragraph of engine-debugging narration that a "
        "player opening the rotation card has no use for and should never see."
    )
    _write_apl(tmp_path, "warrior-fury", [_cast(2687, notes=long_note)])
    bands = _bands("warrior-fury", tmp_path)
    assert bands[10][0].condition == "Short reason."


def test_the_real_committed_data_produces_a_rotation_for_every_spec():
    """End-to-end against the real curated files and the real build's
    spellranks.json, the same inputs `make` regenerates Data.lua from."""
    rotations = build_rotations(Path("builds"), BUILD)
    fury = {band.level: band.lines for band in rotations["warrior-fury"]}
    names_at_60 = {line.name for line in fury[60]}
    assert "Bloodthirst" in names_at_60
    assert "Heroic Strike" in names_at_60
    # Bloodthirst is not learned until 40; a real leveling player's rotation
    # card must not show it earlier.
    names_at_30 = {line.name for line in fury[30]}
    assert "Bloodthirst" not in names_at_30


#: spec -> {ability name -> known-good icon}, each pinned against TWO independent primary
#: sources (player-review sweep 16, 2026-09-30): this build's own raw SpellMisc.csv ->
#: ManifestInterfaceData.csv join (re-fetched from wago.tools for both 1.60.1.70009 and
#: Classic Era 1.15.9.70003 -- identical on both) and Wowhead Classic's own rendered
#: icondb id for the same spell ids (78, 284, 285, 1608, 11564-11567, 25286 for Heroic
#: Strike; 772, 6546-6548, 11572-11574 for Rend; 19434/27632 for Aimed Shot; 2643 for
#: Multi-Shot). Heroic Strike and Rend look wrong at a glance -- they carry Rogue
#: Ambush's and Gouge's own icon files -- but that is a real, longstanding Classic client
#: quirk (both ranks' SpellMisc.SpellIconFileDataID rows have carried those file ids since
#: at least 1.12, on both Era and this Forever beta), not a pipeline bug: sweep 15 flagged
#: it as one, sweep 16 re-verified it against the client tables and Wowhead and found the
#: join correct. Aimed Shot/Multi-Shot are the "looks right, still checked" control pair.
#: A change to any of these values fails this test until a person re-verifies it
#: against the same two sources and explains why the new answer is right (tenet 8).
#: Heroic Strike is pinned on Fury only: the 2026-10-06 rotation search removed it from
#: the Arms list (curated/apl/warrior-arms.json notes), so Arms band 60 no longer has the
#: line to check; Rend still covers the Arms side of the Ambush/Gouge quirk.
#: Aimed Shot is pinned nowhere since 2026-10-31: it shares Multi-Shot's cooldown in
#: Forever and is only the pre-pull opener now (curated/apl/hunter-*.json notes), so no
#: band 60 priority line carries it; Multi-Shot still covers the control pair's side.
KNOWN_GOOD_ROTATION_ICONS = {
    "warrior-arms": {"Rend": "ability_gouge"},
    "warrior-fury": {"Heroic Strike": "ability_rogue_ambush"},
    "hunter-marksmanship": {"Multi-Shot": "ability_upgrademoonglaive"},
    "hunter-beast-mastery": {"Multi-Shot": "ability_upgrademoonglaive"},
}


def test_known_spells_pin_to_their_verified_client_icon():
    """Regression pin for player-review sweep 15's "wrong icon" report: re-verified in
    sweep 16 against a fresh wago.tools SpellMisc/ManifestInterfaceData fetch and
    Wowhead Classic's own icondb id, both of which agree with what this build already
    committed. See KNOWN_GOOD_ROTATION_ICONS's own doc. It reads the active
    build: an older build's spellicons.json holds only the ids its own
    rotations named then (1.60.1.70374's has Heroic Strike's book rank 25286,
    which the addon no longer shows, and not the trainer's rank 8, 11567)."""
    active_json = Path("../web/src/data/active-build.json").read_text(encoding="utf-8")
    active = json.loads(active_json)["build"]
    rotations = build_rotations(Path("builds"), active)
    for spec, expected_by_name in KNOWN_GOOD_ROTATION_ICONS.items():
        band_60 = next(band for band in rotations[spec] if band.level == 60)
        icon_by_name = {line.name: line.icon for line in band_60.lines}
        for name, expected_icon in expected_by_name.items():
            assert icon_by_name.get(name) == expected_icon, (
                f"{spec} band 60 {name!r}: expected verified icon {expected_icon!r}, "
                f"got {icon_by_name.get(name)!r}"
            )


def _write_coach_rules(curated_dir: Path, spec: str, rules: dict) -> None:
    coach_dir = curated_dir / "coach"
    coach_dir.mkdir(parents=True, exist_ok=True)
    (coach_dir / f"{spec}.json").write_text(json.dumps(rules), encoding="utf-8")


def test_coach_spell_ids_reads_every_rank_wherever_a_rule_file_holds_one(tmp_path):
    _write_coach_rules(
        tmp_path,
        "priest-holy",
        {
            "core": [{"family": "Greater Heal", "spell_ids": [2060, 10963, 10964, 10965]}],
            "auras": [{"name": "Inner Fire", "spell_ids": [588, 7128]}],
            "efficiency": {"downranks": [{"spell_id": 2054, "talent": False}]},
            "note": "a spell id in prose is not one: 9999",
        },
    )
    _write_coach_rules(
        tmp_path, "warrior-fury", {"core": [{"family": "Bloodrage", "spell_ids": [2687]}]}
    )
    (tmp_path / "coach" / "priest-holy.md").write_text("spell_ids: [1]", encoding="utf-8")
    assert coach_spell_ids(tmp_path) == {2060, 10963, 10964, 10965, 588, 7128, 2054, 2687}


def test_coach_spell_ids_is_empty_without_rule_files(tmp_path):
    assert coach_spell_ids(tmp_path) == set()


def test_compute_spell_icons_adds_the_coach_spells_no_rotation_line_names(tmp_path):
    _write_apl(tmp_path, "warrior-fury", [_cast(23894, rank=4, notes="Bloodthirst on cooldown.")])
    _write_coach_rules(
        tmp_path, "priest-holy", {"core": [{"family": "Greater Heal", "spell_ids": [2060, 10965]}]}
    )
    (tmp_path / BUILD).mkdir()
    (tmp_path / BUILD / "spellranks.json").write_text(json.dumps(SPELLRANKS), encoding="utf-8")
    icons = compute_spell_icons(tmp_path, BUILD, curated_dir=tmp_path)
    assert {2060, 10965} <= set(icons)
    assert 23894 in icons


def test_compute_spell_icons_adds_the_spec_form_spells(tmp_path):
    """Cat Form and Dire Bear Form draw Feral and Feral Bear (pipeline.specs), so their
    icons are written to spellicons.json beside the rotation spells."""
    import shutil

    _write_apl(tmp_path, "warrior-fury", [_cast(23894, rank=4, notes="Bloodthirst on cooldown.")])
    shutil.copy(Path("curated") / "specs.json", tmp_path / "specs.json")
    (tmp_path / BUILD).mkdir()
    (tmp_path / BUILD / "spellranks.json").write_text(json.dumps(SPELLRANKS), encoding="utf-8")
    icons = compute_spell_icons(tmp_path, BUILD, curated_dir=tmp_path)
    assert {768, 9634} <= set(icons)


def test_the_committed_coach_rule_files_name_the_reviews_spells():
    ids = coach_spell_ids(Path("curated"))
    for spell_id in (2687, 10952):  # Bloodrage; Inner Fire, top rank
        assert spell_id in ids


def test_coach_persona_spell_ids_reads_every_rank_of_a_named_family(tmp_path):
    (tmp_path / "coach").mkdir()
    (tmp_path / "coach" / "warrior-fury.md").write_text(
        "Bloodthirst on cooldown, and Bloodrage for rage. Not Whirlwinds, not battle shout.",
        encoding="utf-8",
    )
    assert coach_persona_spell_ids(tmp_path, SPELLRANKS) == {23881, 23894, 2687}


def test_the_committed_personas_name_the_heals_the_rank_finding_reads():
    spellranks = json.loads(Path("builds/1.60.1.70630/spellranks.json").read_text(encoding="utf-8"))
    ids = coach_persona_spell_ids(Path("curated"), spellranks)
    assert 10965 in ids  # Greater Heal, Rank 4


#: Restoration shaman and Holy priest rows as the 70707 spellranks.json lists
#: them: Healing Wave's rank 10 (25357) is an Ahn'Qiraj book rank (item 21291
#: "Tablet of Healing Wave X" teaches it); Lesser Healing Wave's rank 6 has a
#: Burning Crusade copy (27624) at the same level; Water Shield's buff
#: (408510) and its proc (408511) share level 20; Renew is cast at rank 3 on
#: purpose at 60.
_HEALER_RANKS = {
    "build": BUILD,
    "classes": {
        "shaman": {
            "Healing Wave": [
                {"id": 8005, "level": 40, "rank": 7},
                {"id": 10396, "level": 56, "rank": 9},
                {"book": True, "id": 25357, "level": 60, "rank": 10},
            ],
            "Lesser Healing Wave": [
                {"id": 10466, "level": 44, "rank": 4},
                {"id": 10468, "level": 60, "rank": 6},
                {"id": 27624, "level": 60, "rank": 6},
            ],
            "Water Shield": [
                {"id": 408510, "level": 20, "rank": 0},
                {"id": 408511, "level": 20, "rank": 0},
            ],
        },
        "priest": {
            "Renew": [
                {"id": 6074, "level": 14, "rank": 2},
                {"id": 6075, "level": 20, "rank": 3},
                {"id": 10928, "level": 50, "rank": 8},
                {"id": 10929, "level": 56, "rank": 9},
            ],
        },
    },
}


def _healer_bands(spec: str, curated_dir: Path) -> dict:
    return {band.level: band.lines for band in build_rotation(spec, curated_dir, _HEALER_RANKS)}


def test_a_book_rank_is_never_the_rank_a_band_casts(tmp_path):
    _write_apl(tmp_path, "shaman-restoration", [_cast(10396, rank=9, notes="Healing Wave.")])
    assert [line.spell_id for line in _healer_bands("shaman-restoration", tmp_path)[60]] == [10396]


def test_a_rank_shared_by_two_ids_keeps_the_one_the_list_names(tmp_path):
    _write_apl(
        tmp_path,
        "shaman-restoration",
        [
            _cast(10468, rank=6, notes="Lesser Healing Wave."),
            _cast(408510, notes="Water Shield."),
        ],
    )
    bands = _healer_bands("shaman-restoration", tmp_path)
    assert [line.spell_id for line in bands[60]] == [10468, 408510]
    assert [line.spell_id for line in bands[50]] == [10466, 408510]


def test_the_level_sixty_band_casts_the_list_as_written(tmp_path):
    """The site runs a band-60 list as written (sim/cmd/spec-breakdown and
    the request path rewrite ranks only below 60), so a deliberate lower
    rank there is what the player is told to cast."""
    _write_apl(tmp_path, "priest-holy", [_cast(6075, rank=3, notes="Renew on the tank.")])
    bands = _healer_bands("priest-holy", tmp_path)
    assert [line.spell_id for line in bands[60]] == [6075]


def test_a_deliberate_lower_rank_is_never_lifted_below_sixty(tmp_path):
    """Owner ruling 3 (round 9): below 60 the sim casts the highest rank
    learned that is not above the list's own (sim/internal/spellranks
    LearnedSpellID), so Renew rank 3 stays rank 3, and below its level the
    rank under it stands in."""
    _write_apl(tmp_path, "priest-holy", [_cast(6075, rank=3, notes="Renew on the tank.")])
    bands = _healer_bands("priest-holy", tmp_path)
    assert [line.spell_id for line in bands[50]] == [6075]
    assert [line.spell_id for line in bands[10]] == []
    assert [line.spell_id for line in _healer_bands_at("priest-holy", tmp_path, 16)] == [6074]


def test_a_band_casts_the_rank_choice_the_curated_table_names(tmp_path):
    """data/curated/rank-choices.json picks a lower rank over a band below 60,
    as the sim's rewrite does (sim/request rankResolver)."""
    _write_apl(tmp_path, "priest-holy", [_cast(10929, rank=9, notes="Renew.")])
    (tmp_path / "rank-choices.json").write_text(
        json.dumps(
            {
                "specs": {
                    "priest-holy": [
                        {
                            "spell": "Renew",
                            "authored": 10929,
                            "bands": [
                                {"from_level": 40, "to_level": 49, "rank": 2, "preset": "raid"},
                                {"from_level": 50, "to_level": 59, "rank": 3},
                            ],
                        }
                    ]
                }
            }
        ),
        encoding="utf-8",
    )
    bands = _healer_bands("priest-holy", tmp_path)
    assert [line.spell_id for line in bands[40]] == [6075]
    assert [line.spell_id for line in bands[50]] == [6075]
    assert [line.spell_id for line in bands[60]] == [10929]


def test_level_sixty_casts_the_raid_presets_rank_row(tmp_path):
    """Round 22 (Holy paladin card): level 60 is published under the raid preset
    (sim/request PublishedPreset), so a row for the raid preset applies at 60 and a
    row for the bare preset does not. The addon and the card read the base row
    instead: party Flash of Light printed rank 4 where the raid sim casts rank 5."""
    _write_apl(tmp_path, "priest-holy", [_cast(10929, rank=9, notes="Renew.")])
    (tmp_path / "rank-choices.json").write_text(
        json.dumps(
            {
                "specs": {
                    "priest-holy": [
                        {
                            "spell": "Renew",
                            "authored": 10929,
                            "bands": [
                                {"from_level": 60, "to_level": 60, "rank": 3, "preset": "raid"},
                                {"from_level": 60, "to_level": 60, "rank": 2, "preset": "bare"},
                            ],
                        }
                    ]
                }
            }
        ),
        encoding="utf-8",
    )
    assert [line.spell_id for line in _healer_bands("priest-holy", tmp_path)[60]] == [6075]


def test_a_rank_row_or_a_rank_learned_mid_band_starts_its_own_band(tmp_path):
    """Round 16 (shaman lane): rank rows start at any level now (Lightning Bolt
    rank 4 at 34-38, rank 5 at 39-44), and a rank is learned at any level. The
    addon reads the band whose level is the highest at or below the player's
    (Rotation.lua bandFor), so a band starts wherever a line changes; a
    player at 52 sees the row that starts at 52, not band 50's."""
    _write_apl(tmp_path, "priest-holy", [_cast(10929, rank=9, notes="Renew.")])
    (tmp_path / "rank-choices.json").write_text(
        json.dumps(
            {
                "specs": {
                    "priest-holy": [
                        {
                            "spell": "Renew",
                            "authored": 10929,
                            "bands": [{"from_level": 52, "to_level": 55, "rank": 3}],
                        }
                    ]
                }
            }
        ),
        encoding="utf-8",
    )
    rotation = build_rotation("priest-holy", tmp_path, _HEALER_RANKS)
    levels = [band.level for band in rotation]
    assert levels == [10, 14, 20, 30, 38, 40, 50, 52, 56, 60]
    bands = {band.level: [line.spell_id for line in band.lines] for band in rotation}
    assert bands[14] == [6074]
    assert bands[50] == [10928]
    assert bands[52] == [6075]
    assert bands[56] == [10929]


def test_a_rank_choice_scoped_to_lines_rewrites_only_those_lines(tmp_path):
    """A spell's rows that name "lines" (0-based priority-list entries) apply
    only there, as sim/request's rewrite does: Restoration's Nature's
    Swiftness save keeps its top rank while the filler on the same id downranks."""
    _write_apl(
        tmp_path,
        "priest-holy",
        [_cast(10929, rank=9, notes="Save."), _cast(10929, rank=9, notes="Filler.")],
    )
    (tmp_path / "rank-choices.json").write_text(
        json.dumps(
            {
                "specs": {
                    "priest-holy": [
                        {
                            "spell": "Renew",
                            "authored": 10929,
                            "lines": [1],
                            "bands": [{"from_level": 50, "to_level": 59, "rank": 3}],
                        }
                    ]
                }
            }
        ),
        encoding="utf-8",
    )
    bands = _healer_bands("priest-holy", tmp_path)
    assert [line.spell_id for line in bands[50]][1] == 6075
    assert [line.spell_id for line in bands[50]][0] != 6075


def test_a_rank_choice_scoped_to_pre_pull_actions_leaves_every_line(tmp_path):
    """Rows scoped only to pre-pull actions change no priority-list line."""
    _write_apl(tmp_path, "priest-holy", [_cast(10929, rank=9, notes="Renew.")])
    (tmp_path / "rank-choices.json").write_text(
        json.dumps(
            {
                "specs": {
                    "priest-holy": [
                        {
                            "spell": "Renew",
                            "authored": 10929,
                            "prepull": [0],
                            "bands": [{"from_level": 50, "to_level": 59, "rank": 3}],
                        }
                    ]
                }
            }
        ),
        encoding="utf-8",
    )
    assert [line.spell_id for line in _healer_bands("priest-holy", tmp_path)[50]] != [6075]


def test_a_band_scoped_to_its_own_lines_rewrites_only_those_lines(tmp_path):
    """A band may carry its own "lines" (round 17), as sim/request's parse
    reads it: rank 3 at 50-59 on line 1 only, while a band with none keeps the
    spell's scope (every line here) at 40-49."""
    _write_apl(
        tmp_path,
        "priest-holy",
        [_cast(10929, rank=9, notes="Save."), _cast(10929, rank=9, notes="Filler.")],
    )
    (tmp_path / "rank-choices.json").write_text(
        json.dumps(
            {
                "specs": {
                    "priest-holy": [
                        {
                            "spell": "Renew",
                            "authored": 10929,
                            "bands": [
                                {"from_level": 40, "to_level": 49, "rank": 3},
                                {"from_level": 50, "to_level": 59, "rank": 3, "lines": [1]},
                            ],
                        }
                    ]
                }
            }
        ),
        encoding="utf-8",
    )
    bands = _healer_bands("priest-holy", tmp_path)
    assert [line.spell_id for line in bands[40]] == [6075, 6075]
    assert [line.spell_id for line in bands[50]][1] == 6075
    assert [line.spell_id for line in bands[50]][0] != 6075


def test_a_band_scoped_to_its_own_pre_pull_actions_leaves_every_line(tmp_path):
    """A band naming only "prepull" changes no priority-list line, even when its
    spell's rows otherwise apply everywhere."""
    _write_apl(tmp_path, "priest-holy", [_cast(10929, rank=9, notes="Renew.")])
    (tmp_path / "rank-choices.json").write_text(
        json.dumps(
            {
                "specs": {
                    "priest-holy": [
                        {
                            "spell": "Renew",
                            "authored": 10929,
                            "bands": [
                                {"from_level": 50, "to_level": 59, "rank": 3, "prepull": [0]}
                            ],
                        }
                    ]
                }
            }
        ),
        encoding="utf-8",
    )
    assert [line.spell_id for line in _healer_bands("priest-holy", tmp_path)[50]] != [6075]


def test_a_race_rank_choice_below_sixty_names_each_races_rank(tmp_path):
    """A rank-choice row may name races (round 13). The addon carries one
    rotation per band for every race, so a line whose rank a race row changes
    keeps the every-race id as `spell_id` and names each keyed race's own id
    in `race_spell_ids`; the addon picks by the player's race. A race no row
    names casts `spell_id`."""
    _write_apl(tmp_path, "priest-holy", [_cast(10929, rank=9, notes="Renew.")])
    (tmp_path / "rank-choices.json").write_text(
        json.dumps(
            {
                "specs": {
                    "priest-holy": [
                        {
                            "spell": "Renew",
                            "authored": 10929,
                            "bands": [
                                {"from_level": 40, "to_level": 49, "rank": 2},
                                {
                                    "from_level": 50,
                                    "to_level": 59,
                                    "rank": 3,
                                    "races": ["undead", "night-elf"],
                                },
                                {"from_level": 50, "to_level": 59, "rank": 2, "races": ["gnome"]},
                            ],
                        }
                    ]
                }
            }
        ),
        encoding="utf-8",
    )
    bands = _healer_bands("priest-holy", tmp_path)
    assert [(line.spell_id, line.race_spell_ids) for line in bands[40]] == [(6074, {})]
    [line] = bands[50]
    assert line.spell_id == 10928
    assert line.race_spell_ids == {"gnome": 6074, "night-elf": 6075, "undead": 6075}


def test_a_race_row_that_casts_the_every_race_id_names_no_race(tmp_path):
    """A race row whose rank resolves to the line's own id adds nothing."""
    _write_apl(tmp_path, "priest-holy", [_cast(10929, rank=9, notes="Renew.")])
    (tmp_path / "rank-choices.json").write_text(
        json.dumps(
            {
                "specs": {
                    "priest-holy": [
                        {
                            "spell": "Renew",
                            "authored": 10929,
                            "bands": [
                                {"from_level": 40, "to_level": 49, "rank": 2, "races": ["troll"]},
                                {"from_level": 40, "to_level": 49, "rank": 2, "races": ["human"]},
                            ],
                        }
                    ]
                }
            }
        ),
        encoding="utf-8",
    )
    bands = _healer_bands("priest-holy", tmp_path)
    [line] = bands[40]
    assert line.spell_id != 6074
    assert line.race_spell_ids == {"human": 6074, "troll": 6074}


def _healer_bands_at(spec: str, curated_dir: Path, level: int) -> list:
    from pipeline import addonrotation

    original = addonrotation.LEVEL_BANDS
    addonrotation.LEVEL_BANDS = [level]
    try:
        return build_rotation(spec, curated_dir, _HEALER_RANKS)[0].lines
    finally:
        addonrotation.LEVEL_BANDS = original


def test_a_band_names_the_id_the_sim_casts_when_a_rank_has_a_forever_copy(tmp_path):
    """A rank the client carries twice (Exorcism rank 1: the trainer's 879 and
    the Forever copy 415068) resolves within the list's own id band, the rule
    the sim's rewrite applies (sim/internal/spellranks sameBandOrFirst): the
    band names the spell the sim casts, not the copy no trainer teaches."""
    ranks = {
        "build": BUILD,
        "classes": {
            "paladin": {
                "Exorcism": [
                    {"id": 879, "level": 20, "rank": 1},
                    {"id": 415068, "level": 20, "rank": 1},
                    {"id": 5614, "level": 28, "rank": 2},
                    {"id": 10314, "level": 60, "rank": 6},
                ]
            }
        },
    }
    _write_apl(tmp_path, "paladin-retribution", [_cast(10314, rank=6, notes="Exorcism.")])
    rotation = build_rotation("paladin-retribution", tmp_path, ranks)
    bands = {band.level: band.lines for band in rotation}
    assert [line.spell_id for line in bands[20]] == [879]
    assert [line.spell_id for line in bands[30]] == [5614]


def test_two_ranks_learned_at_one_level_cast_the_higher(tmp_path):
    """Counterattack ranks 1 and 2 are both learned at 30 in Forever; the sim
    casts rank 2 there, so the band names rank 2."""
    ranks = {
        "build": BUILD,
        "classes": {
            "hunter": {
                "Counterattack": [
                    {"id": 19306, "level": 30, "rank": 1},
                    {"id": 1242634, "level": 30, "rank": 2},
                    {"id": 20910, "level": 54, "rank": 4},
                ]
            }
        },
    }
    _write_apl(tmp_path, "hunter-survival", [_cast(20910, rank=4, notes="Counterattack.")])
    bands = {band.level: band.lines for band in build_rotation("hunter-survival", tmp_path, ranks)}
    assert [line.spell_id for line in bands[30]] == [1242634]


def test_a_talents_own_spell_joins_the_rotation_from_its_tiers_level(tmp_path):
    """Stormstrike is the Enhancement talent the rotation leads with; no
    trainer teaches it, so spellranks.json does not track it, and the addon
    used to drop the line at every band. It now joins from the level its
    tier opens (tier 3: level 25), as the sim casts it once talented."""
    from pipeline.addonrotation import load_talent_spells

    build_dir = tmp_path / "builds" / BUILD
    (build_dir / "talents").mkdir(parents=True)
    (build_dir / "talents" / "shaman.json").write_text(
        json.dumps(
            {
                "trees": [
                    {
                        "talents": [
                            {"name": "Stormstrike", "tier": 3, "spell_id": 17364},
                            {"name": "Thundering Strikes", "tier": 0, "spell_id": 16255},
                        ]
                    }
                ]
            }
        ),
        encoding="utf-8",
    )
    ranks = {
        "build": BUILD,
        "classes": {
            "shaman": {
                "Earth Shock": [
                    {"id": 8042, "level": 4, "rank": 1},
                    {"id": 10414, "level": 60, "rank": 7},
                ]
            }
        },
    }
    _write_apl(
        tmp_path,
        "shaman-enhancement",
        [
            _cast(17364, notes="Stormstrike."),
            _cast(10414, rank=7, notes="Earth Shock."),
        ],
    )
    talents = load_talent_spells(build_dir, "shaman")
    rotation = build_rotation("shaman-enhancement", tmp_path, ranks, talent_spells=talents)
    bands = {band.level: band.lines for band in rotation}
    assert [line.name for line in bands[20]] == ["Earth Shock"]
    assert [(line.spell_id, line.name) for line in bands[30]] == [
        (17364, "Stormstrike"),
        (8042, "Earth Shock"),
    ]
    assert [line.spell_id for line in bands[60]] == [17364, 10414]


def test_apl_cast_spell_ids_reads_the_prepull_and_every_nested_cast():
    prepull = {"castSpell": {"spellId": {"spellId": 688}}}
    rotation = {
        "prepullActions": [{"action": prepull, "doAtValue": {"const": {"val": "-3s"}}}],
        "priorityList": [
            {"action": {"castSpell": {"spellId": {"spellId": 11661, "rank": 9}}}},
            {
                "action": {
                    "strictSequence": {"actions": [{"castSpell": {"spellId": {"spellId": 5019}}}]}
                }
            },
            {"action": {"autocastOtherCooldowns": {}}},
        ],
    }
    assert apl_cast_spell_ids(rotation) == {688, 11661, 5019}


def test_compute_spell_icons_covers_the_prepull_casts_no_rank_family_resolves(tmp_path):
    _write_apl(tmp_path, "warrior-fury", [_cast(23894, rank=4, notes="Bloodthirst on cooldown.")])
    apl_path = tmp_path / "apl" / "warrior-fury.json"
    apl = json.loads(apl_path.read_text(encoding="utf-8"))
    prepull = {"castSpell": {"spellId": {"spellId": 688}}}
    apl["rotation"]["prepullActions"] = [
        {"action": prepull, "doAtValue": {"const": {"val": "-3s"}}}
    ]
    apl_path.write_text(json.dumps(apl), encoding="utf-8")
    (tmp_path / BUILD).mkdir()
    (tmp_path / BUILD / "spellranks.json").write_text(json.dumps(SPELLRANKS), encoding="utf-8")
    icons = compute_spell_icons(tmp_path, BUILD, curated_dir=tmp_path)
    assert 688 in icons
    assert 23894 in icons


def test_the_lua_names_each_races_id_only_where_they_differ():
    """Data.lua carries `raceSpellId` on a line only where a race row changes
    its rank, keyed by the site's race slug (quoted: `night-elf` is no bare Lua
    name); the addon reads it by the player's race."""
    from pipeline.addonlua import _rotations_lua
    from pipeline.models import AddonRotationBand, AddonRotationLine

    every = AddonRotationLine(spell_id=8092, name="Mind Blast", condition="", icon="x")
    split = AddonRotationLine(
        spell_id=10945,
        name="Mind Blast",
        condition="",
        icon="x",
        race_spell_ids={"undead": 8103, "night-elf": 8102},
    )
    lua = "\n".join(
        _rotations_lua(
            {
                "priest-shadow": [
                    AddonRotationBand(level=40, lines=[every]),
                    AddonRotationBand(level=50, lines=[split]),
                ]
            }
        )
    )
    assert '{ spellId = 8092, name = "Mind Blast", condition = "" },' in lua
    assert (
        '{ spellId = 10945, name = "Mind Blast", condition = "", '
        'raceSpellId = { ["night-elf"] = 8102, ["undead"] = 8103 } },' in lua
    )
    assert "faction" not in lua


def test_the_holy_paladin_band_40_flash_of_light_pair_exports_the_ranks_it_casts():
    """Sim-accuracy round 16 wave 2 (per-level rows on the MaxLevel + 12
    penalty): at level 40 the curated rows cast Flash of Light rank 1 on the
    tank (rank 2 from 44) and on the group. Round 19: Holy Light is rank 4 at
    40 to 42 (the top rank again at 43 to 46, rank 1 from 47) and Holy Shock
    rank 1 at 40 to 46. The addon's list names those ranks' own ids: Flash of
    Light rank 1 is 19750 (level 20), Holy Light rank 4 is 1026 (level 22),
    Holy Shock rank 1 is 1311606 (level 30)."""
    spellranks = json.loads(Path("builds/1.60.1.70753/spellranks.json").read_text(encoding="utf-8"))
    learned = {r["id"]: r["level"] for r in spellranks["classes"]["paladin"]["Flash of Light"]}
    assert learned[19750] <= 40
    band = next(
        b for b in build_rotation("paladin-holy", Path("curated"), spellranks) if b.level == 40
    )
    flash = [line.spell_id for line in band.lines if line.name == "Flash of Light"]
    # Round 24: the group's Flash of Light is rank 2 (19939) at 40 to 42. Round 29: the tank's is
    # rank 2 from 40 too. Sim-accuracy 2026-11-12 (Holy paladin lane), with Concentration Aura
    # applied and Illumination's casting regen gone: rank 1 (19750) on both at 40 and 41, so band
    # 40 casts rank 1 only.
    assert set(flash) == {19750}
    assert {line.spell_id for line in band.lines if line.name == "Holy Light"} == {1026}
    assert {line.spell_id for line in band.lines if line.name == "Holy Shock"} == {1311606}


def test_rank_siblings_adds_every_rank_of_a_named_family() -> None:
    # Lesser Heal rank 2 is named, so rank 1 (2050) and rank 3 come with it; Smite is untouched.
    spellranks = {
        "classes": {
            "priest": {
                "Lesser Heal": [
                    {"id": 2050, "level": 1, "rank": 1},
                    {"id": 2052, "level": 4, "rank": 2},
                    {"id": 2053, "level": 10, "rank": 3},
                ],
                "Smite": [{"id": 585, "level": 1, "rank": 1}],
            }
        }
    }
    assert rank_siblings(spellranks, {2052}) == {2050, 2053}
    assert rank_siblings(spellranks, {9999}) == set()


def test_a_line_scoped_to_a_level_range_is_in_those_bands_only(tmp_path):
    # Round 19 (priest lane): Discipline's step-down applies from 40 to 59 and waits on an
    # owner ruling at 60; the line is scoped with currentLevel, which the addon reads too.
    def level(op: str, value: int) -> dict:
        rhs = {"const": {"val": str(value)}}
        return {"cmp": {"op": op, "lhs": {"currentLevel": {}}, "rhs": rhs}}

    scoped = _cast(2687, notes="Bloodrage from 40 to 59.")
    scoped["action"]["condition"] = {"and": {"vals": [level("OpGe", 40), level("OpLt", 60)]}}
    _write_apl(tmp_path, "warrior-fury", [scoped, _cast(25289, rank=7, notes="Battle Shout up.")])
    bands = _bands("warrior-fury", tmp_path)
    assert 2687 not in [line.spell_id for line in bands[30]]
    assert 2687 in [line.spell_id for line in bands[40]]
    assert 2687 in [line.spell_id for line in bands[50]]
    assert 2687 not in [line.spell_id for line in bands[60]]


def test_fury_death_wish_card_line_carries_the_whole_rule():
    """The card shows a note's first sentence only, so the Fury Death Wish rule (pull when
    a second one fits, otherwise the execute phase or the last 30 seconds) must sit in one
    sentence; round 19 found the card cut it to the pull clause alone."""
    fury = json.loads(Path("curated/apl/warrior-fury.json").read_text(encoding="utf-8"))
    death_wish = next(
        entry
        for entry in fury["rotation"]["priorityList"]
        if entry["action"].get("castSpell", {}).get("spellId", {}).get("spellId") == 12328
    )
    line = _one_line(death_wish["notes"])
    assert "second one" in line
    assert "execute phase" in line.lower()
    assert "last 30 seconds" in line


def test_fury_berserker_rage_starts_at_36_and_overpower_ends_there():
    """Guide-accuracy round 19 (WARRIOR-R19-1): a Fury warrior fights in Battle Stance until
    Whirlwind at 36, and Berserker Rage is a Berserker Stance button, so its card row starts
    at 36; Overpower, a Battle Stance button, leaves the card there."""
    rotations = build_rotations(Path("builds"), BUILD)
    fury = {band.level: {line.name for line in band.lines} for band in rotations["warrior-fury"]}
    for level in (30, 32):
        assert "Berserker Rage" not in fury[level], level
        assert "Overpower" in fury[level], level
    for level in (36, 40, 60):
        assert "Berserker Rage" in fury[level], level
        assert "Overpower" not in fury[level], level


# Sim round 22 platform (hunter F3 on the talent-level fix): a card row for a talent's spell
# starts at the level the spec's published build spends its first point in that talent, not
# when its tier opens. Survival's order takes Strider Kick at 30 (round 29 dropped
# Counterattack from the build), Enhancement takes Stormstrike at 25 and Fury takes Death
# Wish at 30.
@pytest.mark.parametrize(
    ("spec", "spell", "first"),
    [
        ("hunter-survival", "Strider Kick", 30),
        ("shaman-enhancement", "Stormstrike", 25),
        ("warrior-fury", "Death Wish", 30),
    ],
)
def test_a_talent_row_starts_where_the_published_build_takes_the_talent(spec, spell, first):
    levels = [
        band.level
        for band in _active_build_bands(spec)
        if any(line.name == spell for line in band.lines)
    ]
    assert levels, f"{spell} is on no {spec} band"
    assert min(levels) == first


# Sim round 23 shaman: a line whose condition needs a talent known (auraIsKnown or
# spellIsKnown of the talent's spell, alone or in the condition's top-level and) cannot fire
# before the build takes that talent, so its card row starts there too. Restoration's
# Healing Wave right after Nature's Swiftness printed from 10 (Nature's Swiftness at 44), and
# Enhancement's five-stack Lightning Bolt from 10 (Maelstrom Weapon's first point at 37).
@pytest.mark.parametrize(
    ("spec", "condition", "first"),
    [
        ("shaman-restoration", "right after Nature's Swiftness, on the tank below 25%", 55),
        # Since Blizzard's 10 November hotfix five stacks make the bolt instant and free at any
        # rank, so the gate names the talent untagged and the row starts at its first point.
        ("shaman-enhancement", "at 5 Maelstrom Weapon stacks", 35),
        # Round 24 (platform): the gate is on Clearcasting (16246), the aura Elemental Focus
        # (16164) gives; the leveling build never takes the talent, and since the shaman lane's
        # level 60 builds (12 Nov) the raid order spends its last point, at 60, on it.
        ("shaman-enhancement", "on a Clearcasting proc", 60),
        # The mage proc buffs (Missile Barrage 400589, Heating Up 400625) and the Improved
        # Scorch debuff (12873) printed from 10, 20 and 22, before the talents.
        ("mage-fire", "at 3 Heating Up stacks", 30),
        ("mage-fire", "to 5 Improved Scorch stacks", 27),
        ("mage-fire", "while Missile Barrage is up", 55),
        ("mage-frost", "while Missile Barrage is up", 51),
        ("mage-arcane", "when Missile Barrage procs", 21),
    ],
)
def test_a_row_gated_on_a_talent_starts_where_the_published_build_takes_it(spec, condition, first):
    levels = [
        band.level
        for band in _active_build_bands(spec)
        if any(line.condition.startswith(condition) for line in band.lines)
    ]
    assert levels, f"{condition!r} is on no {spec} band"
    assert min(levels) == first


# Round 24 (platform, second pass): a negated rank-tagged gate (not auraIsKnown of the
# talent's rank N) ends the row where the published build takes rank N: Arcane's Missiles
# "until you have both ranks of Missile Barrage" printed on every band to 60, and
# Destruction's line for a character without Shadow and Flame kept printing after it.
@pytest.mark.parametrize(
    ("spec", "condition", "last"),
    [
        ("mage-arcane", "until you have both ranks of Missile Barrage", 21),
    ],
)
def test_a_row_gated_on_lacking_a_talent_rank_ends_where_the_build_takes_it(spec, condition, last):
    levels = [
        band.level
        for band in _active_build_bands(spec)
        if any(line.condition.startswith(condition) for line in band.lines)
    ]
    assert levels, f"{condition!r} is on no {spec} band"
    assert max(levels) == last


def test_rank_levels_count_each_point_of_a_talent():
    from pipeline.talent_levels import build_rank_levels, first_point_levels

    trees = [{"talents": [{"tier": 0, "column": 0, "ranks": [{"spell_id": 7}, {"spell_id": 7}]}]}]
    levels = build_rank_levels([[1, 1, 1], [1, 1, 1]], trees)
    assert levels == {(7, 1): 10, (7, 2): 11}
    assert first_point_levels(levels) == {7: 10}


def _active_build_bands(spec):
    """`spec`'s card on the active build's spellranks as the talent-level fix writes them
    (built here, not read: the committed file waits for the data job's regeneration)."""
    from pipeline import spellranks as sr
    from pipeline.addonrotation import (
        _class_slug,
        build_rotation,
        load_talent_spells,
        load_trainable_ids,
    )
    from pipeline.manifest import newest_build
    from pipeline.talent_levels import (
        load_spec_rank_levels,
        load_talent_owned_auras,
        load_talent_spell_levels,
    )

    data = Path(__file__).resolve().parents[1]
    build_dir = data / "builds" / newest_build(data / "builds")
    records = sr.load_class_spell_constants(build_dir / "spellconst")
    talent_levels = {
        r.class_slug: load_talent_spell_levels(build_dir / "talents", r.class_slug) for r in records
    }
    spellranks = sr.build_spell_ranks(
        build_dir.name,
        records,
        sr.load_book_ids(data / "curated"),
        sr.load_inferior_ids(data / "curated"),
        talent_levels,
    ).model_dump(exclude_defaults=True)
    cls = _class_slug(spec)
    return build_rotation(
        spec,
        data / "curated",
        spellranks,
        None,
        load_trainable_ids(build_dir, cls),
        load_talent_spells(build_dir, cls),
        load_spec_rank_levels(build_dir, spec, cls, load_talent_owned_auras(data / "curated", cls)),
    )


def test_compute_spell_icons_lets_an_engine_id_take_another_rows_icon(tmp_path):
    (tmp_path / "extra-spell-icons.json").write_text(
        json.dumps({"spell_ids": [], "icon_rows": {"12867": 12162}}), encoding="utf-8"
    )
    build_dir = tmp_path / BUILD
    build_dir.mkdir(exist_ok=True)
    (build_dir / "spellranks.json").write_text(json.dumps({"classes": {}}), encoding="utf-8")
    assert 12867 in compute_spell_icons(tmp_path, BUILD, curated_dir=tmp_path)


# The engine reports Deep Wounds under one id per talent rank
# (sim/warrior/talents.go deepWoundsSpellIDs): 12834, 12849, 12867. Round 26
# platform: only rank 3 (12867) had an icon, so a leveling warrior with one or
# two points showed its bleed with none. 12834 is a client row (Deep Wounds,
# icon 132090, ability_backstab); 12849 is not and takes the bleed's (12162).
DEEP_WOUNDS_ENGINE_IDS = (12834, 12849, 12867)


def test_every_deep_wounds_rank_the_engine_reports_has_an_icon_source():
    extra = json.loads(
        (Path(__file__).resolve().parents[1] / "curated" / "extra-spell-icons.json").read_text(
            encoding="utf-8"
        )
    )
    sources = set(extra["spell_ids"]) | {int(k) for k in extra["icon_rows"]}
    assert [i for i in DEEP_WOUNDS_ENGINE_IDS if i not in sources] == []


def test_raid_cover_spell_ids_reads_rows_and_provider_spells(tmp_path):
    (tmp_path / "raid-cover.json").write_text(
        json.dumps(
            {
                "entries": [
                    {"spell_id": 10938, "providers": [{"class": "priest"}]},
                    {"spell_id": 20355, "providers": [{"class": "paladin", "spell": 20357}]},
                ]
            }
        ),
        encoding="utf-8",
    )
    assert raid_cover_spell_ids(tmp_path) == {10938, 20355, 20357}
    assert raid_cover_spell_ids(tmp_path / "missing") == set()


def test_compute_spell_icons_adds_the_raid_cover_spells(tmp_path):
    import shutil

    shutil.copy(Path("curated") / "raid-cover.json", tmp_path / "raid-cover.json")
    build_dir = tmp_path / BUILD
    build_dir.mkdir(exist_ok=True)
    (build_dir / "spellranks.json").write_text(json.dumps({"classes": {}}), encoding="utf-8")
    icons = compute_spell_icons(tmp_path, BUILD, curated_dir=tmp_path)
    assert raid_cover_spell_ids(Path("curated")) <= set(icons)


def test_every_engine_only_spell_id_has_an_icon_source():
    """Every id `engine-spell-names.json` names (an id the sim reports that no class
    spellbook or client row names) has an icon on the result page: already in the
    build's spellicons.json, or listed in extra-spell-icons.json (`spell_ids` for a
    client row, `icon_rows` for an id the client carries no row for). Round 26 left
    54 without one."""
    curated = Path(__file__).resolve().parents[1] / "curated"
    names = json.loads((curated / "engine-spell-names.json").read_text(encoding="utf-8"))
    extra = json.loads((curated / "extra-spell-icons.json").read_text(encoding="utf-8"))
    from pipeline.specs import active_build_dir

    published = json.loads((active_build_dir() / "spellicons.json").read_text(encoding="utf-8"))
    sources = (
        {int(k) for k in published}
        | set(extra["spell_ids"])
        | {int(k) for k in extra["icon_rows"]}
    )
    assert sorted(int(k) for k in names["spells"] if int(k) not in sources) == []


def _write_apl_with_panel_when(
    curated_dir: Path, spec: str, priority_list: list[dict], panel_when: dict
) -> None:
    apl_dir = curated_dir / "apl"
    apl_dir.mkdir(parents=True, exist_ok=True)
    (apl_dir / f"{spec}.json").write_text(
        json.dumps(
            {"spec": spec, "rotation": {"priorityList": priority_list}, "panel_when": panel_when}
        ),
        encoding="utf-8",
    )


def test_a_panel_when_line_is_the_condition_verbatim(tmp_path):
    _write_apl_with_panel_when(
        tmp_path,
        "warrior-fury",
        [_cast(2687, notes="Whenever rage is under 80. More prose for the guide.")],
        {"0": "under 80 rage"},
    )
    assert _bands("warrior-fury", tmp_path)[10][0].condition == "under 80 rage"


def test_a_step_without_a_panel_when_line_falls_back_to_its_note(tmp_path):
    _write_apl_with_panel_when(
        tmp_path,
        "warrior-fury",
        [_cast(2687, notes="Whenever rage is under 80. More."), _cast(6673, notes="Keep it up.")],
        {"0": "under 80 rage"},
    )
    lines = _bands("warrior-fury", tmp_path)[10]
    assert [line.condition for line in lines] == ["under 80 rage", "Keep it up."]


def test_a_sequence_shares_its_steps_panel_when_line(tmp_path):
    sequence = {
        "notes": "Both, in order.",
        "action": {
            "strictSequence": {
                "actions": [
                    {"castSpell": {"spellId": {"spellId": 2687}}},
                    {"castSpell": {"spellId": {"spellId": 6673}}},
                ]
            }
        },
    }
    _write_apl_with_panel_when(tmp_path, "warrior-fury", [sequence], {"0": "on the pull"})
    lines = _bands("warrior-fury", tmp_path)[10]
    assert [line.condition for line in lines] == ["on the pull", "on the pull"]


@pytest.mark.parametrize(
    ("panel_when", "match"),
    [
        ({"0": ""}, "empty"),
        ({"0": "   "}, "empty"),
        ({"0": " under 80 rage"}, "leading or trailing space"),
        ({"0": "under 80 rage — always"}, "em dash"),
        ({"0": "x" * 61}, "61 characters"),
        ({"zero": "under 80 rage"}, "not a priorityList index"),
        ({"3": "under 80 rage"}, r"step\(s\) \[3\]"),
    ],
)
def test_a_bad_panel_when_line_is_refused(tmp_path, panel_when, match):
    _write_apl_with_panel_when(tmp_path, "warrior-fury", [_cast(2687, notes="N.")], panel_when)
    with pytest.raises(AddonRotationError, match=match):
        build_rotation("warrior-fury", tmp_path, SPELLRANKS)


def test_a_sixty_character_panel_when_line_is_accepted(tmp_path):
    line = "x" * 60
    _write_apl_with_panel_when(tmp_path, "warrior-fury", [_cast(2687, notes="N.")], {"0": line})
    assert _bands("warrior-fury", tmp_path)[10][0].condition == line


def test_every_spell_casting_step_of_every_curated_rotation_has_a_panel_when_line():
    """The panel never falls back to a guide note on the real data: every step that casts
    a spell carries its own short line."""
    missing = []
    for path in sorted(Path("curated/apl").glob("*.json")):
        curated = json.loads(path.read_text(encoding="utf-8"))
        cast_steps = {cast["line"] for cast in _cast_lines(curated.get("rotation", {}))}
        lines = curated.get("panel_when", {})
        missing += [f"{path.stem}[{step}]" for step in sorted(cast_steps) if str(step) not in lines]
    assert not missing, missing
