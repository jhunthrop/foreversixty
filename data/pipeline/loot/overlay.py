"""Forever's own loot facts, laid over the two databases' Era ones.

Parity contract 6.1 as corrected by 10.4, and the design's risk list: the
fork's item database covers Classic Era, and Forever re-itemises. Only
2,811 of the fork's 7,553 item ids exist in the 1.60 client at all, so a
great deal of what a Forever player will actually loot has no source in
either database. This is where a sourced statement about that goes -- and
where the phase a raid opens in goes, since neither database states a date.

Every file is a `LootOverlay`: curated `sources` and `notes` like every
other hand-maintained fact here, plus `add`, `replace` and `remove`.
Files apply in filename order; within a file, add, then replace, then
remove. A stale instruction -- adding a source that exists, patching or
removing one that does not -- is an error, because the alternative is an
overlay that quietly stops doing anything.
"""

from __future__ import annotations

import json
from pathlib import Path

from pipeline.curated import parse_sources
from pipeline.loot.sources import KIND_ORDER
from pipeline.models import LootFile, LootOverlay


class OverlayError(SystemExit):
    """A curated loot overlay names something the generated file does not."""


def load_overlays(overlay_dir: Path) -> list[tuple[Path, LootOverlay]]:
    """Every overlay in filename order, with its provenance validated.

    A missing directory is not an error: a build with nothing to say about
    its loot is a normal state, and the day one exists the file appears.
    """
    if not overlay_dir.is_dir():
        return []
    loaded: list[tuple[Path, LootOverlay]] = []
    for path in sorted(overlay_dir.glob("*.json")):
        document = LootOverlay(**json.loads(path.read_text(encoding="utf-8")))
        parse_sources([source.model_dump() for source in document.sources], str(path))
        if not document.notes.strip():
            raise OverlayError(f"{path} has no notes; say why the overlay exists")
        loaded.append((path, document))
    return loaded


def apply_overlays(
    document: LootFile, overlays: list[tuple[Path, LootOverlay]]
) -> LootFile:
    by_id = {source.id: source for source in document.sources}
    for path, overlay in overlays:
        for source in overlay.add:
            if source.id in by_id:
                raise OverlayError(
                    f"{path} adds source {source.id!r}, which the generated file "
                    f"already has; use `replace` to change it"
                )
            by_id[source.id] = source
        for patch in overlay.replace:
            if patch.id not in by_id:
                raise OverlayError(
                    f"{path} replaces keys on source {patch.id!r}, which no source "
                    f"has; use `add`, or delete the stale instruction"
                )
            existing = by_id[patch.id]
            changes = patch.model_dump(exclude_unset=True)
            changes.pop("id")
            kind = changes.pop("kind", existing.kind)
            if kind != existing.kind:
                raise OverlayError(
                    f"{path} would change source {patch.id!r} from kind "
                    f"{existing.kind!r} to {kind!r}; a source's kind is part of "
                    f"its id and of every `drop:` origin that names it"
                )
            # `model_copy(update=...)` assigns verbatim, without
            # re-validating -- a patch that touches a nested-model field
            # like `bosses` would then store plain dicts where `LootBoss`
            # instances belong. Merging through `model_validate` instead
            # re-parses the whole source, so a patched `bosses` list comes
            # back out as real `LootBoss` instances, not raw dicts.
            by_id[patch.id] = type(existing).model_validate(
                {**existing.model_dump(), **changes}
            )
        for source_id in overlay.remove:
            if source_id not in by_id:
                raise OverlayError(
                    f"{path} removes source {source_id!r}, which is not there"
                )
            del by_id[source_id]
    return LootFile(
        sources=sorted(
            by_id.values(), key=lambda source: (KIND_ORDER.index(source.kind), source.id)
        )
    )
