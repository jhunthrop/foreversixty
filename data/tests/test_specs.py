import json
from pathlib import Path

import pytest

from pipeline.manifest import newest_build
from pipeline.models import SpecRecord
from pipeline.simdb.statmap import STAT_IDS
from pipeline.specs import (
    ROLES,
    SpecError,
    active_build_dir,
    check_specs,
    load_specs,
    render_go,
    render_json,
    render_ts,
    spec_icons,
    write_specs,
)

CURATED = Path("curated")
#: The newest client build, resolved from its manifest rather than named: this
#: file is not a single-build conformance test, and a build id here would be an
#: edit waiting to be forgotten the next time a client build lands.
BUILD_DIR = Path("builds") / newest_build()

#: The one spec name that deliberately differs from the client's tree name:
#: the client calls the tree "Feral Combat", design section 2.3 calls the spec
#: Feral, and the whole site keys on `druid-feral`. The bear tank is the same tree.
TREE_NAME_EXCEPTIONS = {"druid-feral": "Feral Combat", "druid-feral-bear": "Feral Combat"}


#: The build the site serves: the spec-icon map is read off its client data.
ACTIVE_DIR = active_build_dir()


def specs():
    return load_specs(CURATED)


def rendered():
    records = specs()
    icons = spec_icons(records, ACTIVE_DIR)
    return render_go(records, icons, ACTIVE_DIR.name), render_ts(records, icons, ACTIVE_DIR.name)


def test_there_is_one_spec_per_class_tree():
    records = specs()
    assert len(records) == 28
    per_class: dict[str, list[int]] = {}
    for record in records:
        per_class.setdefault(record.class_slug, []).append(record.tree_index)
    assert len(per_class) == 9
    # One tree can serve two specs (the Feral tree is the cat and the bear), so a class
    # covers each of its three trees at least once and no tree outside 0..2.
    assert all(set(indexes) == {0, 1, 2} for indexes in per_class.values())


def test_the_key_is_the_class_and_spec_slug():
    assert all(record.spec == f"{record.class_slug}-{record.spec_slug}" for record in specs())
    keys = [record.spec for record in specs()]
    assert len(set(keys)) == len(keys)


def test_every_slug_is_lower_kebab():
    for record in specs():
        for slug in (record.spec, record.class_slug, record.spec_slug):
            assert slug == slug.lower()
            assert " " not in slug and "_" not in slug


def test_every_role_is_one_of_the_three():
    assert {record.role for record in specs()} <= ROLES


def test_two_specs_are_the_launch_pair():
    keys = {record.spec for record in specs()}
    assert {"warrior-fury", "mage-frost"} <= keys


def test_every_class_slug_is_a_client_class():
    client = {row["slug"] for row in json.loads((BUILD_DIR / "classes.json").read_text())}
    assert {record.class_slug for record in specs()} == client


def test_every_spec_names_the_tree_at_its_index():
    """tree_index is the client's own tree position, and the name matches the
    client's tree name except for the one documented exception."""
    for record in specs():
        trees = json.loads((BUILD_DIR / "talents" / f"{record.class_slug}.json").read_text())[
            "trees"
        ]
        by_position = {tree["position"]: tree["name"] for tree in trees}
        assert record.tree_index in by_position, record.spec
        expected = TREE_NAME_EXCEPTIONS.get(record.spec, record.name)
        assert by_position[record.tree_index] == expected, record.spec


def test_a_duplicate_spec_key_is_rejected(tmp_path: Path):
    entry = {
        "spec": "warrior-fury",
        "class_slug": "warrior",
        "spec_slug": "fury",
        "name": "Fury",
        "role": "dps",
        "tree_index": 1,
        "reference_stat": "attack_power",
        "weight_stats": ["attack_power"],
    }
    (tmp_path / "specs.json").write_text(json.dumps([entry, entry]))
    with pytest.raises(SpecError, match="warrior-fury"):
        load_specs(tmp_path)


def test_a_key_that_is_not_its_two_slugs_is_rejected(tmp_path: Path):
    (tmp_path / "specs.json").write_text(
        json.dumps(
            [
                {
                    "spec": "fury",
                    "class_slug": "warrior",
                    "spec_slug": "fury",
                    "name": "Fury",
                    "role": "dps",
                    "tree_index": 1,
                    "reference_stat": "attack_power",
                    "weight_stats": ["attack_power"],
                }
            ]
        )
    )
    with pytest.raises(SpecError, match="fury"):
        load_specs(tmp_path)


def test_an_unknown_role_is_rejected(tmp_path: Path):
    (tmp_path / "specs.json").write_text(
        json.dumps(
            [
                {
                    "spec": "warrior-fury",
                    "class_slug": "warrior",
                    "spec_slug": "fury",
                    "name": "Fury",
                    "role": "dancer",
                    "tree_index": 1,
                    "reference_stat": "attack_power",
                    "weight_stats": ["attack_power"],
                }
            ]
        )
    )
    with pytest.raises(SpecError, match="dancer"):
        load_specs(tmp_path)


def test_a_missing_curated_file_is_a_clear_error(tmp_path: Path):
    with pytest.raises(SpecError, match="specs.json"):
        load_specs(tmp_path)


def test_the_generated_go_is_a_compilable_package():
    source = rendered()[0]
    assert source.startswith("// Code generated by `python -m pipeline specs`. DO NOT EDIT.")
    assert "package specs" in source
    assert 'Spec: "warrior-fury", ClassSlug: "warrior", SpecSlug: "fury"' in source
    assert source.count("\t{Spec:") == 28
    assert source.endswith("\n")


def test_the_generated_typescript_carries_the_same_rows():
    source = rendered()[1]
    assert source.startswith("// Code generated by `python -m pipeline specs`. DO NOT EDIT.")
    assert "export const SPECS" in source
    assert "spec: 'warrior-fury'" in source
    assert source.count("  {\n") == 28
    assert source.endswith("\n")


def test_the_generated_files_on_disk_are_up_to_date():
    """`python -m pipeline specs` is not run by normalize, so a hand edit to
    specs.json -- or to a generated file, or a client move that changed a tab
    icon -- that was never regenerated would ship lists that disagree with the
    curated source and the client. This is the check `specs --check` runs in CI."""
    assert check_specs(CURATED, ACTIVE_DIR) == []


def test_check_specs_names_the_file_that_drifted(tmp_path: Path):
    outputs = (tmp_path / "specs.go", tmp_path / "specs.ts", tmp_path / "specs.json")
    go_path, ts_path, json_path = outputs
    write_specs(CURATED, ACTIVE_DIR, *outputs)
    assert check_specs(CURATED, ACTIVE_DIR, *outputs) == []
    ts_path.write_text(rendered()[1].replace("warrior-fury", "warrior-furry"))
    assert check_specs(CURATED, ACTIVE_DIR, *outputs) == [ts_path]
    go_path.unlink()
    json_path.unlink()
    assert check_specs(CURATED, ACTIVE_DIR, *outputs) == [go_path, ts_path, json_path]


#: The parity contract's section 1.4 default per spec: attack power for
#: melee and hunters, spell power for casters. Transcribed here from the
#: plan's table rather than imported, so the curated file is checked
#: against the decision and not against itself.
#:
#: hunter-survival reads attack_power, not ranged_attack_power, since the
#: rotation-accuracy program's 2026-09-28 rewrite (owner: "survival will
#: be a melee hunter spec") -- the other two hunter specs are unchanged.
REFERENCE_BY_SPEC = {
    "druid-balance": "spell_power",
    "druid-feral": "attack_power",
    "druid-feral-bear": "stamina",
    "druid-restoration": "healing_power",
    "hunter-beast-mastery": "ranged_attack_power",
    "hunter-marksmanship": "ranged_attack_power",
    "hunter-survival": "attack_power",
    "mage-arcane": "spell_power",
    "mage-fire": "spell_power",
    "mage-frost": "spell_power",
    "paladin-holy": "healing_power",
    "paladin-protection": "stamina",
    "paladin-retribution": "attack_power",
    "priest-discipline": "healing_power",
    "priest-holy": "healing_power",
    "priest-shadow": "spell_power",
    "rogue-assassination": "attack_power",
    "rogue-combat": "attack_power",
    "rogue-subtlety": "attack_power",
    "shaman-elemental": "spell_power",
    "shaman-enhancement": "attack_power",
    "shaman-restoration": "healing_power",
    "warlock-affliction": "spell_power",
    "warlock-demonology": "spell_power",
    "warlock-destruction": "spell_power",
    "warrior-arms": "attack_power",
    "warrior-fury": "attack_power",
    "warrior-protection": "stamina",
}


def test_every_spec_names_the_reference_stat_the_contract_gives_it():
    assert {record.spec: record.reference_stat for record in specs()} == REFERENCE_BY_SPEC


def test_the_stat_vocabulary_is_the_proto_enum_in_snake_case():
    """Contract 10.1 A7. Spot values, plus the shape of the whole set, so a
    renamed or added engine stat shows up here and not in a sim that
    silently normalises to nothing."""
    assert "spell_power" in STAT_IDS
    assert "attack_power" in STAT_IDS
    assert "spell_haste" in STAT_IDS and "melee_haste" in STAT_IDS
    assert "haste" not in STAT_IDS
    assert "mp5" in STAT_IDS
    # Round 9: crit on every table plus one stat per table.
    assert {"crit", "melee_crit", "spell_crit"} <= STAT_IDS
    assert len(STAT_IDS) == 43
    assert all(name == name.lower() for name in STAT_IDS)


def test_a_reference_stat_the_engine_has_no_stat_for_is_refused(tmp_path):
    rows = [
        {
            "spec": "warrior-arms",
            "class_slug": "warrior",
            "spec_slug": "arms",
            "name": "Arms",
            "role": "dps",
            "tree_index": 0,
            "reference_stat": "swagger",
            "weight_stats": ["swagger"],
        }
    ]
    (tmp_path / "specs.json").write_text(json.dumps(rows), encoding="utf-8")
    with pytest.raises(SpecError, match="reference_stat"):
        load_specs(tmp_path)


def test_both_generated_files_carry_the_reference_stat():
    go, ts = rendered()
    assert 'ReferenceStat string `json:"reference_stat"`' in go
    assert 'ReferenceStat: "attack_power"' in go
    assert "reference_stat: string;" in ts
    assert "reference_stat: 'spell_power'," in ts


#: Task 5(c): specs gain weight_stats, the closed list /sim/weights offers
#: for a spec. A physical spec (reference_stat attack_power, stamina for a tank, or, for a
#: ranged-primary hunter spec, ranged_attack_power) is never asked about a
#: caster stat; a caster spec (reference_stat spell_power) is never asked
#: about a melee-only stat. Named as a table so the test states the rule
#: rather than re-deriving it.
PHYSICAL_REFERENCE_STATS = {"attack_power", "ranged_attack_power", "stamina"}
PHYSICAL_FORBIDDEN = {
    "spirit",
    "mp5",
    "intellect",
    "spell_power",
    "spell_haste",
    "spell_penetration",
    "arcane_power",
    "fire_power",
    "frost_power",
    "holy_power",
    "nature_power",
    "shadow_power",
}
CASTER_FORBIDDEN = {"strength", "expertise", "armor_penetration", "feral_attack_power"}
#: Forever's exceptions to the physical rule (round 5): its talents turn
#: Intellect into attack power and spell power (Mental Dexterity and Mental
#: Quickness for Enhancement, Careful Aim for hunters, Champion of the Light
#: for Retribution), and the engine measures spell power on Retribution's
#: seals and the hunters' rotations. Each spec is allowed exactly these.
FOREVER_PHYSICAL_CASTER_STATS = {
    "shaman-enhancement": {"intellect", "spell_power"},
    "paladin-retribution": {"intellect", "spell_power"},
    "hunter-marksmanship": {"intellect", "spell_power"},
    "hunter-beast-mastery": {"intellect", "spell_power"},
    "hunter-survival": {"intellect"},
}


def test_every_spec_has_a_nonempty_weight_stats():
    for record in specs():
        assert record.weight_stats, record.spec


def test_every_weight_stat_is_a_known_engine_stat():
    for record in specs():
        for stat in record.weight_stats:
            assert stat in STAT_IDS, f"{record.spec}: {stat}"


def test_the_reference_stat_is_always_in_its_own_weight_stats():
    """WeightsSpec.validate refuses a reference that is not among the
    stats being weighed, so a curated list that violated this would make
    the spec's own default weights request illegal."""
    for record in specs():
        assert record.reference_stat in record.weight_stats, record.spec


def test_no_physical_spec_carries_a_caster_stat_and_no_caster_spec_carries_a_melee_stat():
    for record in specs():
        forbidden = (
            PHYSICAL_FORBIDDEN
            if record.reference_stat in PHYSICAL_REFERENCE_STATS
            else CASTER_FORBIDDEN
        )
        allowed = FOREVER_PHYSICAL_CASTER_STATS.get(record.spec, set())
        carried = (forbidden - allowed) & set(record.weight_stats)
        assert not carried, f"{record.spec} carries {carried}"


def test_feral_attack_power_belongs_to_the_two_feral_specs_and_nowhere_else():
    """Cat and Bear form both turn feral attack power into attack power, and
    items/<class>.json states form-gated attack power under that key."""
    carriers = {record.spec for record in specs() if "feral_attack_power" in record.weight_stats}
    assert carriers == {"druid-feral", "druid-feral-bear"}


def test_an_empty_weight_stats_is_rejected(tmp_path: Path):
    rows = [
        {
            "spec": "warrior-arms",
            "class_slug": "warrior",
            "spec_slug": "arms",
            "name": "Arms",
            "role": "dps",
            "tree_index": 0,
            "reference_stat": "attack_power",
            "weight_stats": [],
        }
    ]
    (tmp_path / "specs.json").write_text(json.dumps(rows), encoding="utf-8")
    with pytest.raises(SpecError, match="weight_stats"):
        load_specs(tmp_path)


def test_an_unknown_weight_stat_is_rejected(tmp_path: Path):
    rows = [
        {
            "spec": "warrior-arms",
            "class_slug": "warrior",
            "spec_slug": "arms",
            "name": "Arms",
            "role": "dps",
            "tree_index": 0,
            "reference_stat": "attack_power",
            "weight_stats": ["attack_power", "swagger"],
        }
    ]
    (tmp_path / "specs.json").write_text(json.dumps(rows), encoding="utf-8")
    with pytest.raises(SpecError, match="weight_stats"):
        load_specs(tmp_path)


def test_a_duplicate_weight_stat_is_rejected(tmp_path: Path):
    rows = [
        {
            "spec": "warrior-arms",
            "class_slug": "warrior",
            "spec_slug": "arms",
            "name": "Arms",
            "role": "dps",
            "tree_index": 0,
            "reference_stat": "attack_power",
            "weight_stats": ["attack_power", "attack_power"],
        }
    ]
    (tmp_path / "specs.json").write_text(json.dumps(rows), encoding="utf-8")
    with pytest.raises(SpecError, match="twice"):
        load_specs(tmp_path)


def test_a_reference_stat_missing_from_weight_stats_is_rejected(tmp_path: Path):
    rows = [
        {
            "spec": "warrior-arms",
            "class_slug": "warrior",
            "spec_slug": "arms",
            "name": "Arms",
            "role": "dps",
            "tree_index": 0,
            "reference_stat": "attack_power",
            "weight_stats": ["strength", "agility"],
        }
    ]
    (tmp_path / "specs.json").write_text(json.dumps(rows), encoding="utf-8")
    with pytest.raises(SpecError, match="weight_stats"):
        load_specs(tmp_path)


def test_both_generated_files_carry_the_weight_stats():
    go, ts = rendered()
    assert 'WeightStats []string `json:"weight_stats"`' in go
    assert 'WeightStats: []string{"attack_power"' in go
    # The TypeScript list carries the same field, so the web's own test that
    # SPECS equals data/curated/specs.json keeps the data lane the only author.
    assert "weight_stats: readonly string[];" in ts
    assert "weight_stats: ['attack_power'" in ts


#: The spec-icon map: each spec's icon is read off the active client, never curated.


def _tree_icon(class_slug: str, tree_index: int) -> str:
    trees = json.loads((ACTIVE_DIR / "talents" / f"{class_slug}.json").read_text())["trees"]
    return next(tree["icon"] for tree in trees if tree["position"] == tree_index)


def test_every_spec_icon_is_its_trees_talenttab_icon_or_its_form_spells_icon():
    """The map matches the client: a spec with a tree of its own draws that tree's
    TalentTab icon; Feral and Feral Bear draw Cat Form's and Dire Bear Form's."""
    spell_icons = json.loads((ACTIVE_DIR / "spellicons.json").read_text())
    icons = spec_icons(specs(), ACTIVE_DIR)
    for record in specs():
        if record.form_spell is None:
            expected = _tree_icon(record.class_slug, record.tree_index)
        else:
            expected = spell_icons[str(record.form_spell)]
        assert icons[record.spec] == expected, record.spec


def test_feral_is_cat_form_and_feral_bear_is_dire_bear_form():
    by_key = {record.spec: record for record in specs()}
    assert by_key["druid-feral"].form_spell == 768
    assert by_key["druid-feral-bear"].form_spell == 9634
    with_form = [record.spec for record in specs() if record.form_spell is not None]
    assert with_form == ["druid-feral", "druid-feral-bear"]


def test_no_two_specs_of_a_class_share_an_icon():
    icons = spec_icons(specs(), ACTIVE_DIR)
    for class_slug in {record.class_slug for record in specs()}:
        mine = [icons[r.spec] for r in specs() if r.class_slug == class_slug]
        assert len(set(mine)) == len(mine), class_slug


def test_curated_specs_json_names_no_icon():
    """The icon is generated, so a hand-written one in the curated list is a second source."""
    rows = json.loads((CURATED / "specs.json").read_text())
    assert all("icon" not in row for row in rows)


def _write_build(build_dir: Path, tree_icons: list[str], spell_icons: dict[str, str]) -> None:
    (build_dir / "talents").mkdir(parents=True)
    trees = [{"position": i, "icon": icon} for i, icon in enumerate(tree_icons)]
    (build_dir / "talents" / "druid.json").write_text(json.dumps({"trees": trees}))
    (build_dir / "spellicons.json").write_text(json.dumps(spell_icons))


def _druid(spec_slug: str, tree_index: int, form_spell: int | None = None) -> SpecRecord:
    return SpecRecord(
        spec=f"druid-{spec_slug}",
        class_slug="druid",
        spec_slug=spec_slug,
        name=spec_slug,
        role="dps",
        tree_index=tree_index,
        reference_stat="attack_power",
        weight_stats=["attack_power"],
        form_spell=form_spell,
    )


def test_a_form_spell_the_build_has_no_icon_for_is_refused(tmp_path: Path):
    _write_build(tmp_path, ["a", "b", "c"], {})
    with pytest.raises(SpecError, match="spell 768"):
        spec_icons([_druid("feral", 1, 768)], tmp_path)


def test_a_tree_without_an_icon_is_refused(tmp_path: Path):
    _write_build(tmp_path, ["a", "", "c"], {})
    with pytest.raises(SpecError, match="druid-feral has no icon"):
        spec_icons([_druid("feral", 1)], tmp_path)


def test_two_specs_on_one_tree_without_form_spells_are_refused(tmp_path: Path):
    _write_build(tmp_path, ["a", "b", "c"], {})
    with pytest.raises(SpecError, match="share an icon"):
        spec_icons([_druid("feral", 1), _druid("feral-bear", 1)], tmp_path)


def test_every_generated_file_carries_the_icon_and_its_build():
    go, ts = rendered()
    icons = spec_icons(specs(), ACTIVE_DIR)
    assert 'Icon string `json:"icon"`' in go
    assert f'Icon: "{icons["warrior-fury"]}"' in go
    assert f"icons from data/builds/{ACTIVE_DIR.name}" in go
    assert "icon: string;" in ts
    assert f"icon: '{icons['warrior-fury']}'," in ts
    assert f"export const SPEC_ICON_BUILD = '{ACTIVE_DIR.name}';" in ts
    rows = json.loads(render_json(specs(), icons))
    assert {row["spec"]: row["icon"] for row in rows} == icons
    assert all("form_spell" not in row for row in rows)


def test_the_typescript_list_carries_each_capstone_the_json_list_does():
    """The saved sim's "never pressed" warning reads `capstone_spell` from specs.ts; the
    two generated lists must agree on it (web/src/lib/sim/spec-label.test.ts)."""
    _, ts = rendered()
    with_capstone = [record for record in specs() if record.capstone_spell is not None]
    assert with_capstone, "no spec carries a capstone_spell"
    for record in with_capstone:
        assert f"capstone_spell: {record.capstone_spell}," in ts
    assert "capstone_spell?: number;" in ts
