"""Turn Wowhead's Forever talent payload into our per-class talent files.

Before the beta client ships on 2026-09-17 this is the only real Forever talent data
that exists: Blizzard's own tooltip text for every rank of all 470 talents, served by
Wowhead's `classicplus` data environment. Every community dataset was instead built by
reading BlizzCon stream frames, which only ever showed rank 1, so their higher ranks are
arithmetic extrapolated from the Classic talent of the same name. See
`data/raw-forever/README.md` for the snapshot and how it was checked.

The payload's shape maps onto ours exactly, verified against the 2026-09-14 snapshot:

* its 27 tree ids are our 27 tree ids, so `WarriorArms` is 161 on both sides;
* every talent's `ranks` array is the same length as its `descriptions` map;
* all 70 prerequisites name a talent in the same tree;
* no talent has more than one prerequisite, which is what `prereq_talent_id` plus
  `prereq_rank` can express.

`normalize_forever_talents` asserts each of those rather than assuming them, because the
payload is a live endpoint that can change under us before the beta.
"""

from __future__ import annotations

from pipeline.models import ClassTalents, TalentEntry, TalentRank, TalentTree


class ForeverTalentError(ValueError):
    """The payload does not have the shape our talent model can carry."""


def _rank_count(talent: dict) -> int:
    ranks = talent.get("ranks") or []
    descriptions = talent.get("descriptions") or {}
    if len(ranks) != len(descriptions):
        raise ForeverTalentError(
            f"talent {talent.get('id')} {talent.get('name')!r} has {len(ranks)} ranks "
            f"but {len(descriptions)} descriptions"
        )
    if not ranks:
        raise ForeverTalentError(f"talent {talent.get('id')} {talent.get('name')!r} has no ranks")
    return len(ranks)


def _prereq(talent: dict, tree_ids: set[int]) -> tuple[int | None, int | None]:
    requires = talent.get("requires") or []
    if not requires:
        return None, None
    if len(requires) > 1:
        raise ForeverTalentError(
            f"talent {talent.get('id')} {talent.get('name')!r} has {len(requires)} "
            "prerequisites; the planner's rules carry one"
        )
    req = requires[0]
    prereq_id = int(req["id"])
    if prereq_id not in tree_ids:
        raise ForeverTalentError(
            f"talent {talent.get('id')} {talent.get('name')!r} requires {prereq_id}, "
            "which is not in the same tree"
        )
    return prereq_id, int(req["qty"])


def normalize_forever_talents(
    payload: dict,
    *,
    build: str,
    classes: list[dict],
    tree_class: dict[int, int],
    tree_names: dict[int, str],
    tree_backgrounds: dict[int, str],
) -> list[ClassTalents]:
    """Build one `ClassTalents` per class from Wowhead's payload.

    `classes` is our own class table (id, name, slug), `tree_class` maps a tree id to a
    class id, `tree_names` gives each tree its display name and `tree_backgrounds` its
    background art name. All four come from the build we already normalized, so nothing
    about the class, tree list or art is inferred from Wowhead. The tree names matter:
    Wowhead's `description` glues the class onto the tree with no separator
    ("WarriorArms", "HunterBeastMastery"), and un-gluing it would have to guess where the
    words break. The tree ids are identical on both sides, so a lookup is exact where a
    split would be a guess. Wowhead's payload carries no background art at all, so that
    field is always the already-normalized build's own.
    """
    trees_meta = payload.get("trees") or {}
    talents_by_tree = payload.get("talents") or {}
    if not trees_meta or not talents_by_tree:
        raise ForeverTalentError("payload has no trees or no talents")

    by_class: dict[int, list[TalentTree]] = {}
    for tree_id_str in trees_meta:
        tree_id = int(tree_id_str)
        class_id = tree_class.get(tree_id)
        if class_id is None:
            raise ForeverTalentError(f"tree {tree_id} is not one of ours")
        raw = talents_by_tree.get(tree_id_str)
        if not raw:
            raise ForeverTalentError(f"tree {tree_id} has no talents")

        ids = {int(t["id"]) for t in raw.values()}
        entries: list[TalentEntry] = []
        for talent in raw.values():
            count = _rank_count(talent)
            prereq_id, prereq_rank = _prereq(talent, ids)
            descriptions = talent["descriptions"]
            entries.append(
                TalentEntry(
                    id=int(talent["id"]),
                    name=talent["name"],
                    icon=talent.get("icon", ""),
                    max_rank=count,
                    tier=int(talent["row"]),
                    column=int(talent["col"]),
                    prereq_talent_id=prereq_id,
                    prereq_rank=prereq_rank,
                    # Wowhead publishes one spell id per talent, not one per rank, so a
                    # rank carries the talent's id and its own text. The per-rank spell
                    # ids arrive with the beta client.
                    ranks=[
                        TalentRank(
                            spell_id=int(talent["id"]),
                            description=descriptions[str(n)],
                        )
                        for n in range(1, count + 1)
                    ],
                    # Same story as the ranks above: Wowhead has no real spell id for
                    # the talent itself, only the talent's own id.
                    spell_id=int(talent["id"]),
                )
            )
        entries.sort(key=lambda e: (e.tier, e.column))

        name = tree_names.get(tree_id)
        if not name:
            raise ForeverTalentError(f"tree {tree_id} has no name in our tables")
        background = tree_backgrounds.get(tree_id)
        if not background:
            raise ForeverTalentError(f"tree {tree_id} has no background in our tables")
        by_class.setdefault(class_id, []).append(
            TalentTree(
                id=tree_id,
                name=name,
                position=0,
                talents=entries,
                background=background,
            )
        )

    out: list[ClassTalents] = []
    for klass in classes:
        trees = by_class.get(klass["id"])
        if not trees:
            continue
        trees.sort(key=lambda t: t.id)
        for position, tree in enumerate(trees):
            tree.position = position
        out.append(
            ClassTalents(
                build=build,
                class_id=klass["id"],
                class_slug=klass["slug"],
                trees=trees,
            )
        )
    return out
