import json
from pathlib import Path

import pytest

from pipeline.manifest import newest_build
from pipeline.simdb.statmap import STAT_IDS
from pipeline.specs import (
    ROLES,
    SpecError,
    check_specs,
    load_specs,
    render_go,
    render_ts,
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


def specs():
    return load_specs(CURATED)


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
        "icon": "icon",
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
                    "icon": "icon",
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
                    "icon": "icon",
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
    source = render_go(specs())
    assert source.startswith("// Code generated by `python -m pipeline specs`. DO NOT EDIT.")
    assert "package specs" in source
    assert 'Spec: "warrior-fury", ClassSlug: "warrior", SpecSlug: "fury"' in source
    assert source.count("\t{Spec:") == 28
    assert source.endswith("\n")


def test_the_generated_typescript_carries_the_same_rows():
    source = render_ts(specs())
    assert source.startswith("// Code generated by `python -m pipeline specs`. DO NOT EDIT.")
    assert "export const SPECS" in source
    assert "spec: 'warrior-fury'" in source
    assert source.count("  {\n") == 28
    assert source.endswith("\n")


def test_the_two_generated_files_on_disk_are_up_to_date():
    """`python -m pipeline specs` is not run by normalize, so a hand edit to
    specs.json -- or to a generated file -- that was never regenerated would
    ship a Go and a TypeScript list that disagree with the curated source.
    This is the same check `python -m pipeline specs --check` runs in CI."""
    assert check_specs(CURATED) == []


def test_check_specs_names_the_file_that_drifted(tmp_path: Path):
    go_path = tmp_path / "specs.go"
    ts_path = tmp_path / "specs.ts"
    write_specs(CURATED, go_path, ts_path)
    assert check_specs(CURATED, go_path, ts_path) == []
    ts_path.write_text(render_ts(specs()).replace("warrior-fury", "warrior-furry"))
    assert check_specs(CURATED, go_path, ts_path) == [ts_path]
    go_path.unlink()
    assert check_specs(CURATED, go_path, ts_path) == [go_path, ts_path]


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
    assert len(STAT_IDS) == 41
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
            "icon": "icon",
        }
    ]
    (tmp_path / "specs.json").write_text(json.dumps(rows), encoding="utf-8")
    with pytest.raises(SpecError, match="reference_stat"):
        load_specs(tmp_path)


def test_both_generated_files_carry_the_reference_stat():
    records = specs()
    go, ts = render_go(records), render_ts(records)
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
        carried = forbidden & set(record.weight_stats)
        assert not carried, f"{record.spec} carries {carried}"


def test_feral_attack_power_belongs_to_druid_feral_and_nowhere_else():
    carriers = {record.spec for record in specs() if "feral_attack_power" in record.weight_stats}
    assert carriers == {"druid-feral"}


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
            "icon": "icon",
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
            "icon": "icon",
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
            "icon": "icon",
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
            "icon": "icon",
        }
    ]
    (tmp_path / "specs.json").write_text(json.dumps(rows), encoding="utf-8")
    with pytest.raises(SpecError, match="weight_stats"):
        load_specs(tmp_path)


def test_both_generated_files_carry_the_weight_stats():
    records = specs()
    go, ts = render_go(records), render_ts(records)
    assert 'WeightStats []string `json:"weight_stats"`' in go
    assert 'WeightStats: []string{"attack_power"' in go
    # The TypeScript list carries the same field, so the web's own test that
    # SPECS equals data/curated/specs.json keeps the data lane the only author.
    assert "weight_stats: readonly string[];" in ts
    assert "weight_stats: ['attack_power'" in ts


#: day3 data-followups-11 lane: every spec gets its own talent-tab icon
#: (tenet 3, "an ability is never just a name" -- the same rule for a
#: spec's own tab).


def test_every_spec_has_a_nonempty_icon():
    for record in specs():
        assert record.icon, record.spec


def test_an_empty_icon_is_rejected(tmp_path: Path):
    rows = [
        {
            "spec": "warrior-arms",
            "class_slug": "warrior",
            "spec_slug": "arms",
            "name": "Arms",
            "role": "dps",
            "tree_index": 0,
            "reference_stat": "attack_power",
            "weight_stats": ["attack_power"],
            "icon": "",
        }
    ]
    (tmp_path / "specs.json").write_text(json.dumps(rows), encoding="utf-8")
    with pytest.raises(SpecError, match="icon"):
        load_specs(tmp_path)


def test_both_generated_files_carry_the_icon():
    records = specs()
    go, ts = render_go(records), render_ts(records)
    assert 'Icon string `json:"icon"`' in go
    warrior_fury = next(r for r in records if r.spec == "warrior-fury")
    assert f'Icon: "{warrior_fury.icon}"' in go
    assert "icon: string;" in ts
    assert f"icon: '{warrior_fury.icon}'," in ts


def test_every_specs_icon_matches_its_own_builds_tree_icon():
    """The curated icon is hand-duplicated from the build's own talent
    tree (SpecRecord.icon's own doc) precisely so the site never has to
    join specs.json against a build to draw a spec picker -- this test
    is what keeps the duplicate honest.

    A floor, not an unconditional pass: data.yml's test job runs against
    the COMMITTED build files before the same run regenerates them (day3
    data-followups-11 lane's own report), so a strict assertion here
    would block the very regen that adds `icon` to talents/<class>.json's
    trees. Skipped while the committed build still predates that field
    (no tree anywhere carries one); once a `python -m pipeline normalize`
    regen lands it, this asserts the match strictly, same as before."""
    class_slugs = {record.class_slug for record in specs()}
    all_trees = {
        slug: json.loads((BUILD_DIR / "talents" / f"{slug}.json").read_text())["trees"]
        for slug in class_slugs
    }
    if not any("icon" in tree for trees in all_trees.values() for tree in trees):
        pytest.skip(
            f"builds/{BUILD_DIR.name}/talents/*.json predates tree icons; "
            "regen with `python -m pipeline normalize` to enable this check"
        )
    for record in specs():
        by_position = {tree["position"]: tree["icon"] for tree in all_trees[record.class_slug]}
        assert by_position[record.tree_index] == record.icon, record.spec
