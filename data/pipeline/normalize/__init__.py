import json
import logging
import shutil
from collections.abc import Sequence
from dataclasses import dataclass
from pathlib import Path
from typing import Any

from pydantic import BaseModel

logger = logging.getLogger(__name__)


@dataclass(frozen=True)
class NormalizeResult:
    """What one normalize run produced, and what it could not.

    `skipped` is empty on a complete run and otherwise holds one human-readable
    reason per output the run could not emit. A caller that wants the soft
    behaviour (keep every other file, do not raise) just uses `build_dir` and
    ignores it; the CLI turns a non-empty `skipped` into a non-zero exit, so an
    incomplete build directory can never be committed by a green CI run.
    """

    build_dir: Path
    skipped: tuple[str, ...] = ()


#: Minimum fraction of a committed `items/<class>.json`'s previous item count
#: a regenerated file may keep. `csvio.check_item_sparse_completeness` guards
#: the raw ItemSparse export itself; this is the second half of the same
#: incident's fix (night-fetch-guard, 2026-09-29) -- a per-build loss (a
#: curated override, a filter regression) that does not trip the raw-table
#: ratio would still shrink one class's file, and this catches it right
#: before that file would be overwritten.
CLASS_ITEMS_SHRINK_RATIO = 0.9


def _check_class_items_not_shrunk(
    path: Path, previous_count: int, new_count: int, *, allow_shrink: bool
) -> None:
    """Refuse to overwrite `path` with a class item list that shrank too much.

    `previous_count` is 0 for a class with no committed file yet (a new
    class, or the first build ever normalized), which is never a shrink and
    always passes. `allow_shrink=True` (the CLI's `--allow-shrink`) skips
    the check for a deliberate re-baseline and logs loudly that it did.
    """
    if previous_count == 0:
        return
    if allow_shrink:
        logger.warning("items shrink gate skipped (--allow-shrink) for %s", path)
        return
    if new_count < previous_count * CLASS_ITEMS_SHRINK_RATIO:
        raise SystemExit(
            f"{path} would shrink from {previous_count} to {new_count} items "
            f"({new_count / previous_count:.0%}, need {CLASS_ITEMS_SHRINK_RATIO:.0%}); "
            f"refusing to overwrite the committed file -- re-fetch the build, or pass "
            f"--allow-shrink to force a deliberate re-baseline"
        )


def _write(payload: Any, path: Path, sort_keys: bool = False) -> None:
    """Write one JSON payload in the pipeline's only serialization format.

    indent=2, ensure_ascii=False and a single trailing newline are the
    determinism contract, so they are defined here once and nowhere else.
    `sort_keys` is for a caller whose payload's dict order is not already
    deterministic on its own (`pipeline.spellranks`, which builds its
    `classes`/spell-name keys by iterating a `dict[str, SpellConstant]` in
    whatever order Python handed it back).
    """
    path.parent.mkdir(parents=True, exist_ok=True)
    text = json.dumps(payload, indent=2, ensure_ascii=False, sort_keys=sort_keys)
    path.write_text(text + "\n", encoding="utf-8")


def write_json(records: Sequence[BaseModel], path: Path) -> None:
    ordered = sorted(records, key=lambda r: r.id)  # type: ignore[attr-defined]
    _write([r.model_dump() for r in ordered], path)


def write_records(records: Sequence[BaseModel], path: Path) -> None:
    """Write a list of records in the order given, without sorting by id."""
    _write([r.model_dump() for r in records], path)


def write_model(record: BaseModel, path: Path) -> None:
    """Write a single record as a JSON object."""
    _write(record.model_dump(), path)


def write_document(record: BaseModel, path: Path) -> None:
    """Write a single record as a JSON object, dropping the keys it leaves unset.

    `write_model` emits every field, which is right for a record whose shape
    is fixed. `loot.json`'s sources are a union -- a crafted source has no
    `bosses` and a raid has no `profession` -- and a file of nulls is a file
    every consumer has to filter, so `None` means "not part of this kind"
    and is left out.
    """
    _write(record.model_dump(exclude_none=True), path)


def normalize_build(
    build: str,
    root: Path = Path("builds"),
    curated_dir: Path = Path("curated"),
    *,
    allow_shrink: bool = False,
    engine: Path | None = None,
) -> NormalizeResult:
    """`engine`, optional, points at a wowsims-forever checkout: when given,
    an item whose `IconFileDataID` the client states as 0 (or names no row
    in `ManifestInterfaceData`) falls back to the fork's own icon before the
    wowhead payload's -- see `pipeline.normalize.gear.build_class_items`'s
    own doc. Without it (the default), only the wowhead payload is tried,
    the same as the standalone `python -m pipeline icons` pass without
    `--engine`.
    """
    from pipeline.csvio import check_item_sparse_completeness, read_csv
    from pipeline.curated import merge_curated
    from pipeline.curves import load_rank_points
    from pipeline.hotfix_merge import merge_hotfix_table
    from pipeline.icons import icon_names
    from pipeline.icons_fix import load_fork_icons, load_wowhead_icons
    from pipeline.manifest import write_manifest
    from pipeline.normalize.classes import normalize_classes, normalize_races
    from pipeline.normalize.classicdb import load_classicdb_supplement
    from pipeline.normalize.classicdb import merge_class_items as merge_classicdb_class_items
    from pipeline.normalize.classicdb import merge_items as merge_classicdb_items
    from pipeline.normalize.classicdb import merge_sets as merge_classicdb_sets
    from pipeline.normalize.dungeons import normalize_dungeons
    from pipeline.normalize.effects import EffectIndex
    from pipeline.normalize.gear import ItemDataError, build_class_items, build_item_sets
    from pipeline.normalize.item_curves import load_item_curves
    from pipeline.normalize.itemnames import write_item_names
    from pipeline.normalize.items import normalize_items
    from pipeline.normalize.sockets import check_no_sockets
    from pipeline.normalize.spells import normalize_spells
    from pipeline.normalize.talent_trees import build_talent_trees
    from pipeline.normalize.talents import flat_talents, normalize_talents
    from pipeline.normalize.trait_trees import build_trait_talent_trees
    from pipeline.normalize.traits import TraitDataError, TraitRows, has_trait_trees
    from pipeline.normalize.weapon_curves import load_weapon_curves
    from pipeline.normalize.wowhead import (
        load_required_levels,
        load_supplement,
        merge_class_items,
        merge_items,
        merge_sets,
    )
    from pipeline.normalize.zones import normalize_zones
    from pipeline.spelltext import ExtraRows, load_spell_text

    build_dir = root / build
    raw = build_dir / "raw"
    if not raw.exists():
        raise SystemExit(f"no raw data at {raw}; run `python -m pipeline fetch` first")
    # ItemSparse/Item are merged with any `raw/hotfixes/*.csv`
    # (`python -m pipeline hotfixes`, hotfix-cache lane 2026-09-29) before
    # anything -- including the completeness gate -- reads them, so a row
    # the client only carries as a runtime hotfix counts the same as one the
    # shipped .db2 has outright.
    sparse_rows = merge_hotfix_table(raw, "ItemSparse")
    item_rows = merge_hotfix_table(raw, "Item")

    def t(name: str) -> list[dict[str, str]]:
        if name == "ItemSparse":
            return sparse_rows
        if name == "Item":
            return item_rows
        return read_csv(raw / f"{name}.csv")

    # A table a build's client simply does not have (see wago.OPTIONAL_TABLES)
    # is written as a header-only CSV by the fetch, but a build fetched before
    # that table was added to TABLES has no file at all. Both mean "no rows".
    def optional(name: str) -> list[dict[str, str]]:
        path = raw / f"{name}.csv"
        return read_csv(path) if path.exists() else []

    # Phase 0 flat entities.
    class_rows = t("ChrClasses")
    # Contract 6.4: a socketed item means the gem decision in the parity
    # design's section 4.4 no longer holds. Checked before anything is
    # written, so a build that trips it leaves no half-emitted directory.
    check_no_sockets(t("ItemSparse"), build)
    write_json(normalize_zones(t("AreaTable"), t("Map")), build_dir / "zones.json")
    write_json(normalize_dungeons(t("JournalInstance")), build_dir / "dungeons.json")
    client_items = normalize_items(t("ItemSparse"), t("Item"))
    client_ids = {item.id for item in client_items}
    wowhead_supplement = load_supplement(build_dir, client_ids)
    items = (
        merge_items(client_items, wowhead_supplement)
        if wowhead_supplement is not None
        else client_items
    )
    # classic-db's 1.12 item_template only ever fills the ids BOTH the client
    # and wowhead's own supplement lack (catalogue-universe lane, 2026-09-30;
    # see pipeline.classicdb_items' own doc) -- so it runs strictly after the
    # wowhead merge above, over the union of ids either one already placed.
    #
    # "the client already has this id" means the client's own ItemSparse/
    # hotfix row is real equippable data (InventoryType != 0), not merely
    # that a row with this id exists: Orb of Deception (1973, build
    # 1.60.1.70009) is a case in point -- its hotfix-merged ItemSparse row
    # carries the right quality/required_level/item_level (3/54/59, matching
    # classic-db's own item_template exactly) but a bare 0 for
    # InventoryType, so build_class_items silently drops it as unslotted;
    # classic-db's real InventoryType 12 (trinket) is what actually ships
    # it. Treating "any row" as "known" would leave this id in the same
    # unshipped state a client-only build already left it in, for a
    # different reason -- the exact defect this lane exists to close.
    client_equippable_ids = {item.id for item in client_items if item.inventory_type != 0}
    known_ids = client_equippable_ids | (
        {item.id for item in wowhead_supplement} if wowhead_supplement is not None else set()
    )
    classicdb_supplement = load_classicdb_supplement(build_dir, known_ids)
    items = (
        merge_classicdb_items(items, classicdb_supplement)
        if classicdb_supplement is not None
        else items
    )
    # items.json's own completeness gate and write are deferred to just after
    # the per-class items/ gate below (still before "Curated Forever facts"),
    # not run here: checking (and possibly writing) it this early would let a
    # narrower per-class regression (`_check_class_items_not_shrunk`) get
    # masked by this broader, less specific gate firing first whenever the
    # same raw change shrinks both -- almost always, since the per-class
    # files are themselves built from these same raw rows.
    write_json(normalize_spells(t("SpellName")), build_dir / "spells.json")

    # Phase 1 planner data. Both directories are rebuilt from scratch so a class
    # that disappears between builds does not leave a stale file behind.
    spell_effect_rows = t("SpellEffect")
    spell_text = load_spell_text(
        t("Spell"),
        t("SpellMisc"),
        spell_effect_rows,
        t("SpellDuration"),
        ExtraRows(
            aura_options=optional("SpellAuraOptions"),
            radius=optional("SpellRadius"),
            range=optional("SpellRange"),
            names=t("SpellName"),
        ),
    )
    icons = icon_names(t("ManifestInterfaceData"))
    spell_names = {int(r["ID"]): r["Name_lang"] for r in t("SpellName")}

    trait_rows = TraitRows(
        skill_line=optional("SkillLine"),
        skill_line_x_trait_tree=optional("SkillLineXTraitTree"),
        talent_tab=t("TalentTab"),
        chr_classes=class_rows,
        node=optional("TraitNode"),
        node_entry=optional("TraitNodeEntry"),
        node_x_entry=optional("TraitNodeXTraitNodeEntry"),
        definition=optional("TraitDefinition"),
        edge=optional("TraitEdge"),
        cond=optional("TraitCond"),
        currency=optional("TraitCurrency"),
        node_group=optional("TraitNodeGroup"),
        node_group_x_node=optional("TraitNodeGroupXTraitNode"),
    )

    # The 1.60 client keeps Forever's talents in the trait tables; the legacy
    # Talent table it still ships holds Classic Era's, so reading that here
    # would draw the wrong trees. Classic Era has no class trait trees at all.
    if has_trait_trees(trait_rows.skill_line_x_trait_tree):
        talent_records = build_trait_talent_trees(
            trait_rows,
            class_rows,
            spell_names,
            spell_text,
            load_rank_points(optional("TraitDefinitionEffectPoints"), optional("CurvePoint")),
            icons,
            build,
        )
        flat = flat_talents(talent_records)
        if not flat:
            # has_trait_trees only checked SkillLineXTraitTree; a build that
            # populates that one table but is missing (or 404s on) another
            # required trait table -- TraitNode, TraitNodeEntry,
            # TraitDefinition, ... -- would otherwise reach here and commit
            # an empty talents.json for a real class, silently, at exit 0.
            raise TraitDataError(
                f"build {build} names a class trait tree in SkillLineXTraitTree but the "
                "trait reader produced zero talents; a required trait table is likely "
                "missing or empty"
            )
    else:
        talent_rows, tab_rows = t("Talent"), t("TalentTab")
        talent_records = build_talent_trees(
            talent_rows, tab_rows, class_rows, spell_names, spell_text, icons, build
        )
        flat = normalize_talents(talent_rows, tab_rows)
    write_json(flat, build_dir / "talents.json")
    shutil.rmtree(build_dir / "talents", ignore_errors=True)
    for record in talent_records:
        write_model(record, build_dir / "talents" / f"{record.class_slug}.json")
    item_sets = build_item_sets(t("ItemSet"), t("ItemSetSpell"), spell_text)
    if wowhead_supplement is not None:
        item_sets = merge_sets(item_sets, wowhead_supplement)
    if classicdb_supplement is not None:
        item_sets = merge_classicdb_sets(item_sets, classicdb_supplement)
    write_json(item_sets, build_dir / "sets.json")
    # Captured now, before anything below can delete or overwrite the
    # committed items/ directory, so the shrink gate has the old counts to
    # compare a regenerated file against -- and so a build that trips it
    # never even reaches the rmtree that would otherwise take the good,
    # committed files down with it (checked before writing, not after).
    previous_item_counts = {
        path.stem: len(json.loads(path.read_text(encoding="utf-8")).get("items", []))
        for path in sorted((build_dir / "items").glob("*.json"))
    }
    skipped: list[str] = []
    curves = load_item_curves(
        t("ItemArmorTotal"),
        t("ItemArmorQuality"),
        t("ItemArmorShield"),
        t("ArmorLocation"),
        t("RandPropPoints"),
    )
    weapon_curves = load_weapon_curves(
        *(
            optional(f"ItemDamage{kind}")
            for kind in ("OneHand", "TwoHand", "Ranged", "Wand", "Thrown")
        )
    )
    effects = EffectIndex(
        t("ItemEffect"),
        optional("ItemXItemEffect"),
        spell_effect_rows,
        spell_text,
    )
    fork_icons = load_fork_icons(engine) if engine is not None else {}
    wowhead_icons = load_wowhead_icons(build_dir)
    wowhead_required_levels = load_required_levels(build_dir)
    try:
        class_items = build_class_items(
            t("ItemSparse"),
            t("Item"),
            class_rows,
            icons,
            build,
            curves,
            effects=effects,
            weapon_curves=weapon_curves,
            fork_icons=fork_icons,
            wowhead_icons=wowhead_icons,
            wowhead_required_levels=wowhead_required_levels,
        )
    except ItemDataError as error:
        logger.warning("items not emitted for build %s: %s", build, error)
        skipped.append(f"items/: {error}")
        shutil.rmtree(build_dir / "items", ignore_errors=True)
    else:
        if wowhead_supplement is not None:
            class_items = merge_class_items(class_items, wowhead_supplement, class_rows)
        if classicdb_supplement is not None:
            class_items = merge_classicdb_class_items(
                class_items,
                classicdb_supplement,
                class_rows,
                spell_text,
                fork_icons,
                wowhead_icons,
            )
        # Every class is checked against its own previous count before any of
        # them is written -- one class failing the gate must not leave a
        # directory that is half regenerated and half deleted.
        for record in class_items:
            _check_class_items_not_shrunk(
                build_dir / "items" / f"{record.class_slug}.json",
                previous_item_counts.get(record.class_slug, 0),
                len(record.items),
                allow_shrink=allow_shrink,
            )
        shutil.rmtree(build_dir / "items", ignore_errors=True)
        for record in class_items:
            write_model(record, build_dir / "items" / f"{record.class_slug}.json")
        write_item_names(build, items, class_items, build_dir)

    # The flat catalog's own gate, now that any narrower per-class regression
    # above has already had first refusal -- see the comment where `items`
    # was computed. Runs (and, on success, writes) whether the per-class
    # build above succeeded or was skipped for an unrelated ItemDataError:
    # this list comes from `normalize_items`/the wowhead supplement, neither
    # of which depends on `build_class_items` succeeding.
    check_item_sparse_completeness(build_dir, len(items), allow_shrink=allow_shrink)
    write_json(items, build_dir / "items.json")

    # Curated Forever facts.
    classes, races, combos = merge_curated(
        normalize_classes(class_rows), normalize_races(t("ChrRaces")), curated_dir
    )
    write_json(classes, build_dir / "classes.json")
    write_json(races, build_dir / "races.json")
    write_records(combos, build_dir / "combos.json")

    meta = json.loads((raw / "_meta.json").read_text(encoding="utf-8"))
    write_manifest(
        build_dir, build=meta["build"], product=meta["product"], fetched_at=meta["fetched_at"]
    )
    return NormalizeResult(build_dir=build_dir, skipped=tuple(skipped))
