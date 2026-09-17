"""The committed build must satisfy the Phase 1 interface contract.

Every other test in this suite runs the pipeline over the fixtures. This one
reads what is actually committed under `builds/<build>/` and checks it against
`docs/superpowers/specs/2026-09-13-phase-1-interfaces.md`, because that is the
file the api and web plans were written against.

The vocabularies and key sets below are transcribed from that contract on
purpose: importing them from `pipeline` would only prove the pipeline agrees
with itself. If one of these assertions fails, the contract is the authority —
fix the pipeline, not the constant.

Each JSON file is parsed once for the whole module (the item files are ~22 MB
together) and the icon directory is read as one listing.
"""

import json
from functools import cache
from pathlib import Path

BUILD = "1.15.9.69722"
BUILD_DIR = Path("builds") / BUILD

TALENT_FILE_KEYS = {"build", "class_id", "class_slug", "trees"}
TREE_KEYS = {"id", "name", "position", "talents", "background"}
TALENT_KEYS = {
    "id",
    "name",
    "icon",
    "max_rank",
    "tier",
    "column",
    "prereq_talent_id",
    "prereq_rank",
    "ranks",
    "spell_id",
}
RANK_KEYS = {"spell_id", "description"}

ITEM_FILE_KEYS = {"build", "class_slug", "items"}
ITEM_KEYS = {
    "id",
    "name",
    "icon",
    "slot",
    "quality",
    "required_level",
    "item_level",
    "armor",
    "stats",
    "set_id",
    "unique",
}

#: The contract's planner slots, minus the four numbered ones: "Rings and
#: trinkets are emitted with slot `finger` and `trinket`; the planner maps them
#: to the two slots."
EMITTED_SLOTS = set(
    "head neck shoulder back chest wrist hands waist legs feet finger trinket "
    "main_hand off_hand ranged".split()
)
STAT_KEYS = set(
    "strength agility stamina intellect spirit armor crit hit spell_power healing "
    "attack_power defense dodge parry block mp5 fire_res frost_res nature_res "
    "shadow_res arcane_res".split()
)


@cache
def talent_files() -> dict[str, dict]:
    return {p.stem: json.loads(p.read_text(encoding="utf-8")) for p in files("talents")}


@cache
def item_files() -> dict[str, dict]:
    return {p.stem: json.loads(p.read_text(encoding="utf-8")) for p in files("items")}


@cache
def sets_by_id() -> dict[int, dict]:
    rows = json.loads((BUILD_DIR / "sets.json").read_text(encoding="utf-8"))
    return {row["id"]: row for row in rows}


def files(name: str) -> list[Path]:
    return sorted((BUILD_DIR / name).glob("*.json"))


def talents() -> list[tuple[str, dict]]:
    """Every talent in the build, paired with the class slug it came from."""
    return [
        (slug, talent)
        for slug, payload in talent_files().items()
        for tree in payload["trees"]
        for talent in tree["talents"]
    ]


def items() -> list[tuple[str, dict]]:
    return [(slug, item) for slug, payload in item_files().items() for item in payload["items"]]


def test_every_class_has_a_talent_file_and_an_item_file():
    assert talent_files()
    assert set(item_files()) >= set(talent_files())
    # Guard every loop below: an empty build would make them all vacuously true.
    assert talents() and items()


def test_talent_files_carry_exactly_the_contract_keys():
    for slug, payload in talent_files().items():
        assert set(payload) == TALENT_FILE_KEYS, slug
        assert payload["build"] == BUILD
        assert payload["class_slug"] == slug
        for tree in payload["trees"]:
            assert set(tree) == TREE_KEYS, (slug, tree.get("id"))
            for talent in tree["talents"]:
                assert set(talent) == TALENT_KEYS, (slug, talent.get("id"))
                for rank in talent["ranks"]:
                    assert set(rank) == RANK_KEYS, (slug, talent["id"])


def test_item_files_carry_exactly_the_contract_keys():
    for slug, payload in item_files().items():
        assert set(payload) == ITEM_FILE_KEYS, slug
        assert payload["build"] == BUILD
        assert payload["class_slug"] == slug
    for slug, item in items():
        assert set(item) == ITEM_KEYS, (slug, item.get("id"))


def test_every_talent_has_one_rank_entry_per_point():
    for slug, talent in talents():
        assert len(talent["ranks"]) == talent["max_rank"], (slug, talent["id"])


def test_every_prereq_resolves_inside_its_own_file_with_a_reachable_rank():
    for slug, payload in talent_files().items():
        by_id = {talent["id"]: talent for tree in payload["trees"] for talent in tree["talents"]}
        for talent in by_id.values():
            prereq_id = talent["prereq_talent_id"]
            if prereq_id is None:
                assert talent["prereq_rank"] is None, (slug, talent["id"])
                continue
            assert prereq_id in by_id, (slug, talent["id"], prereq_id)
            assert 1 <= talent["prereq_rank"] <= by_id[prereq_id]["max_rank"], (
                slug,
                talent["id"],
            )


def test_no_talent_or_item_carries_an_empty_name_or_icon():
    """An empty icon resolves to `icons/.webp` in the site; an empty name renders blank."""
    for slug, talent in talents():
        assert talent["name"] and talent["icon"], (slug, talent["id"])
    for slug, item in items():
        assert item["name"] and item["icon"], (slug, item["id"])


def test_every_referenced_icon_is_on_disk_and_no_icon_is_orphaned():
    on_disk = {p.stem for p in (BUILD_DIR / "icons").iterdir()}
    assert all(p.suffix == ".webp" for p in (BUILD_DIR / "icons").iterdir())
    referenced = {talent["icon"] for _slug, talent in talents()}
    referenced |= {item["icon"] for _slug, item in items()}
    assert sorted(referenced - on_disk) == []
    assert sorted(on_disk - referenced) == []


def test_every_slot_and_stat_key_is_in_the_contract_vocabulary():
    for slug, item in items():
        assert item["slot"] in EMITTED_SLOTS, (slug, item["id"], item["slot"])
        unknown = set(item["stats"]) - STAT_KEYS
        assert not unknown, (slug, item["id"], sorted(unknown))


def test_every_set_an_item_belongs_to_is_in_sets_json():
    for slug, item in items():
        if item["set_id"] is not None:
            assert item["set_id"] in sets_by_id(), (slug, item["id"], item["set_id"])


def test_the_manifest_covers_exactly_the_files_on_disk():
    manifest = json.loads((BUILD_DIR / "manifest.json").read_text(encoding="utf-8"))
    on_disk = {
        path.relative_to(BUILD_DIR).as_posix() for path in BUILD_DIR.rglob("*") if path.is_file()
    }
    on_disk = {name for name in on_disk if name != "manifest.json" and not name.startswith("raw/")}
    assert set(manifest["files"]) == on_disk
    assert manifest["build"] == BUILD
