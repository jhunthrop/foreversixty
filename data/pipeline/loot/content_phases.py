"""The content-phase gate: which sources are not open at launch, by game id.

`curated/content-phases.json` is the one statement of it, read here and by
`sim/cmd/leveling-bis/contentphases.go`, which applies the same gates to
whatever loot.json it ranks from. The house rule (pending owner
confirmation) and the key vocabulary are in that file's own notes.

Two properties the earlier gates lacked:

* An id this build's sources do not name is a warning, never an error.
  The source-id overlay treats a stale instruction as an error, which is
  why commit 3592e40a5 deleted the Anachronos gate the night the vendor
  dropped out of one fetch -- and the vendor came back ungated.
* Duplicate quest entries are merged before gating. One ungated way to an
  item is enough for the ranker, so a bare copy of a quest used to defeat
  the gate its twin carried (Treasure of the Timeless One, The Eye of
  Zuldazar).
"""

from __future__ import annotations

import json
import logging
from collections import defaultdict
from dataclasses import dataclass, field
from datetime import datetime
from pathlib import Path
from typing import Literal

from pydantic import BaseModel

from pipeline.curated import parse_sources
from pipeline.forkdb import REP_LEVELS
from pipeline.loot.sources import KIND_ORDER, rep_gate_by_item
from pipeline.models import LootBoss, LootFile, LootSource, QuestSource, Source
from pipeline.phases import OPENS_LATER, load_phases

logger = logging.getLogger(__name__)

CONTENT_PHASES = "content-phases.json"
#: The phase every gate is measured against: open at launch or not.
LAUNCH_PHASE = "launch"
UNANNOUNCED = "unannounced"
#: Reputation standings, lowest first, for a faction gate's `from_standing`.
STANDING_ORDER = [REP_LEVELS[level] for level in sorted(REP_LEVELS)]
#: Source kinds whose bosses carry an npc id a gate can split out.
INSTANCE_KINDS = frozenset({"dungeon", "raid"})
#: Key types in the order they are listed and warned about.
_KEY_TYPES = ("npcs", "quests", "factions", "zones", "kinds", "items")
_KEY_LABELS = {
    "npcs": "npc", "quests": "quest", "factions": "faction",
    "zones": "zone", "kinds": "kind", "items": "item",
}  # fmt: skip


class ContentPhaseError(SystemExit):
    """curated/content-phases.json says something the pipeline will not apply."""


class GateId(BaseModel):
    id: int
    name: str
    #: Factions only: the lowest standing gated; every standing when unset.
    from_standing: str | None = None

    model_config = {"extra": "forbid"}


class ContentGate(BaseModel):
    key: str
    label: str
    original_patch: str
    phase: str
    date: str
    confidence: Literal["confirmed", "inferred", "uncertain"]
    why: str
    npcs: list[GateId] = []
    quests: list[GateId] = []
    factions: list[GateId] = []
    zones: list[GateId] = []
    kinds: list[str] = []
    items: list[GateId] = []

    model_config = {"extra": "forbid"}


class ContentPhasesFile(BaseModel):
    sources: list[Source]
    notes: str
    gates: list[ContentGate]

    model_config = {"extra": "forbid"}


@dataclass(frozen=True)
class GateRef:
    """What one listed id gates to, and which gate says so."""

    opens: str
    key: str
    from_standing: str | None = None


@dataclass(frozen=True)
class GateIndex:
    """Every CLOSED gate (its phase opens after launch), by key type."""

    npcs: dict[int, GateRef] = field(default_factory=dict)
    npc_names: dict[str, GateRef] = field(default_factory=dict)
    quests: dict[int, GateRef] = field(default_factory=dict)
    factions: dict[int, GateRef] = field(default_factory=dict)
    zones: dict[int, GateRef] = field(default_factory=dict)
    kinds: dict[str, GateRef] = field(default_factory=dict)
    items: dict[int, GateRef] = field(default_factory=dict)


@dataclass(frozen=True)
class ContentPhases:
    gates: list[ContentGate]
    #: phases.json's names in order, then `later`: a gate's strictness.
    phase_order: list[str]

    def rank(self, opens: str | None) -> int:
        """-1 for open; otherwise how late the phase is."""
        if not opens:
            return -1
        return self.phase_order.index(opens) if opens in self.phase_order else len(self.phase_order)

    def closed(self, phase: str) -> bool:
        return self.rank(phase) > self.rank(LAUNCH_PHASE)

    def strictest(self, *values: str | None) -> str | None:
        named = [value for value in values if value]
        return max(named, key=self.rank) if named else None

    def index(self) -> GateIndex:
        index = GateIndex()
        for gate in self.gates:
            if not self.closed(gate.phase):
                continue
            for entry in gate.npcs:
                ref = GateRef(gate.phase, gate.key)
                index.npcs[entry.id] = ref
                index.npc_names[entry.name] = ref
            for kind in gate.kinds:
                index.kinds[kind] = GateRef(gate.phase, gate.key)
            for entries, target in (
                (gate.quests, index.quests), (gate.zones, index.zones), (gate.items, index.items),
            ):  # fmt: skip
                for entry in entries:
                    target[entry.id] = GateRef(gate.phase, gate.key)
            for entry in gate.factions:
                index.factions[entry.id] = GateRef(gate.phase, gate.key, entry.from_standing)
        return index


def _phase_dates(curated_dir: Path) -> dict[str, str]:
    """phase name -> its start as a UTC date, in phases.json order."""
    return {
        phase.name: datetime.fromisoformat(phase.start.replace("Z", "+00:00")).date().isoformat()
        for phase in load_phases(curated_dir)
    }


def _check_gate(gate: ContentGate, dates: dict[str, str], where: str) -> None:
    if gate.phase == OPENS_LATER:
        if gate.date != UNANNOUNCED:
            raise ContentPhaseError(
                f"{where} gate {gate.key!r} opens {OPENS_LATER!r}, so its date must be "
                f"{UNANNOUNCED!r}, not {gate.date!r}"
            )
    elif gate.phase not in dates:
        raise ContentPhaseError(
            f"{where} gate {gate.key!r} names phase {gate.phase!r}; use one of "
            f"{[*dates, OPENS_LATER]}"
        )
    elif gate.date != dates[gate.phase]:
        raise ContentPhaseError(
            f"{where} gate {gate.key!r} dates phase {gate.phase!r} {gate.date!r}; "
            f"phases.json starts it {dates[gate.phase]}"
        )
    if not any(getattr(gate, key_type) for key_type in _KEY_TYPES):
        raise ContentPhaseError(f"{where} gate {gate.key!r} names no id")
    for entry in [*gate.npcs, *gate.quests, *gate.zones, *gate.items]:
        if entry.from_standing is not None:
            raise ContentPhaseError(
                f"{where} gate {gate.key!r}: from_standing is for factions only"
            )
    for entry in gate.factions:
        if entry.from_standing is not None and entry.from_standing not in STANDING_ORDER:
            raise ContentPhaseError(
                f"{where} gate {gate.key!r} faction {entry.id} from_standing "
                f"{entry.from_standing!r}; use one of {STANDING_ORDER}"
            )


def _check_unique(gates: list[ContentGate], where: str) -> None:
    """One id in two gates would leave its phase to file order."""
    seen: dict[tuple[str, object], str] = {}
    keys: set[str] = set()
    for gate in gates:
        if gate.key in keys:
            raise ContentPhaseError(f"{where} names gate {gate.key!r} twice")
        keys.add(gate.key)
        for key_type in _KEY_TYPES:
            for entry in getattr(gate, key_type):
                value = entry if isinstance(entry, str) else entry.id
                where_else = seen.setdefault((key_type, value), gate.key)
                if where_else != gate.key:
                    raise ContentPhaseError(
                        f"{where} lists {_KEY_LABELS[key_type]} {value} in both "
                        f"{where_else!r} and {gate.key!r}"
                    )


def load_content_phases(curated_dir: Path = Path("curated")) -> ContentPhases:
    path = curated_dir / CONTENT_PHASES
    if not path.exists():
        raise ContentPhaseError(f"missing {path}")
    document = ContentPhasesFile(**json.loads(path.read_text(encoding="utf-8")))
    parse_sources([source.model_dump() for source in document.sources], str(path))
    if not document.notes.strip():
        raise ContentPhaseError(f"{path} has no notes; say what the gates are for")
    dates = _phase_dates(curated_dir)
    if LAUNCH_PHASE not in dates:
        raise ContentPhaseError(f"phases.json names no {LAUNCH_PHASE!r} phase to gate against")
    for gate in document.gates:
        _check_gate(gate, dates, str(path))
    _check_unique(document.gates, str(path))
    return ContentPhases(gates=document.gates, phase_order=[*dates, OPENS_LATER])


# -- duplicate quests ------------------------------------------------------


def _merge_entries(entries: list[QuestSource], phases: ContentPhases | None) -> QuestSource:
    first = entries[0]
    rep = next((entry for entry in entries if entry.required_rep_faction is not None), first)
    opens_values = [entry.opens for entry in entries]
    opens = (
        phases.strictest(*opens_values)
        if phases is not None
        else next((value for value in opens_values if value), None)
    )
    classes = next((entry.classes for entry in entries if entry.classes), first.classes)
    return first.model_copy(
        update={
            "required_rep_faction": rep.required_rep_faction,
            "required_rep_standing": rep.required_rep_standing,
            "opens": opens,
            "classes": classes,
        }
    )


def merge_duplicate_quests(document: LootFile, phases: ContentPhases | None = None) -> LootFile:
    """One entry per quest id per item, keeping the reputation fields and the
    strictest `opens` any copy states."""
    quests: dict[str, list[QuestSource]] = {}
    for item_id, entries in document.quests.items():
        groups: dict[int, list[QuestSource]] = defaultdict(list)
        for entry in entries:
            groups[entry.quest_id].append(entry)
        quests[item_id] = [_merge_entries(group, phases) for group in groups.values()]
    return document.model_copy(update={"quests": quests})


# -- applying the gates ------------------------------------------------------


@dataclass
class _Matched:
    """Which listed ids tonight's document actually names."""

    ids: dict[str, set[object]] = field(default_factory=lambda: defaultdict(set))

    def add(self, key_type: str, value: object) -> None:
        self.ids[key_type].add(value)


def _standing_gated(ref: GateRef, standing: str | None) -> bool:
    if ref.from_standing is None or standing is None:
        return True
    return STANDING_ORDER.index(standing) >= STANDING_ORDER.index(ref.from_standing)


def _source_gate(
    source: LootSource, phases: ContentPhases, index: GateIndex, matched: _Matched
) -> GateRef | None:
    """The gate a whole source falls under, by its own keys."""
    refs: list[GateRef] = []
    if source.kind == "vendor" and source.npc_id in index.npcs:
        matched.add("npcs", source.npc_id)
        refs.append(index.npcs[source.npc_id])
    if source.kind == "world" and source.name in index.npc_names:
        matched.add("npc_names", source.name)
        refs.append(index.npc_names[source.name])
    if source.zone_id is not None and source.zone_id in index.zones:
        matched.add("zones", source.zone_id)
        refs.append(index.zones[source.zone_id])
    if source.kind in index.kinds:
        matched.add("kinds", source.kind)
        refs.append(index.kinds[source.kind])
    if source.kind == "rep" and source.faction_id in index.factions:
        matched.add("factions", source.faction_id)
        ref = index.factions[source.faction_id]
        if _standing_gated(ref, source.standing):
            refs.append(ref)
    return max(refs, key=lambda ref: phases.rank(ref.opens)) if refs else None


def _split_bosses(
    source: LootSource, phases: ContentPhases, index: GateIndex, matched: _Matched
) -> tuple[LootSource, list[LootSource]]:
    """A gated boss inside an open instance gets its own sibling source."""
    kept: list[LootBoss] = []
    gated: dict[GateRef, list[LootBoss]] = defaultdict(list)
    for boss in source.bosses or []:
        ref = index.npcs.get(boss.npc_id)
        if ref is None:
            kept.append(boss)
            continue
        matched.add("npcs", boss.npc_id)
        gated[ref].append(boss)
    if not gated:
        return source, []
    siblings = [
        source.model_copy(
            update={
                "id": f"{source.id}:{ref.key}",
                "bosses": bosses,
                "trash": None,
                "opens": phases.strictest(source.opens, ref.opens),
            }
        )
        for ref, bosses in gated.items()
    ]
    return source.model_copy(update={"bosses": kept}), siblings


def _split_items(
    source: LootSource, phases: ContentPhases, index: GateIndex, matched: _Matched
) -> tuple[LootSource | None, list[LootSource]]:
    """An item-keyed gate moves the item out of a flat source."""
    if not source.items or source.kind == "quest":
        # The flat "quest" bucket carries no gate of its own; a gated
        # item's quest entries are gated in `_gate_quests` instead.
        return source, []
    kept: list[int] = []
    gated: dict[GateRef, list[int]] = defaultdict(list)
    for item_id in source.items:
        ref = index.items.get(item_id)
        if ref is None:
            kept.append(item_id)
        else:
            matched.add("items", item_id)
            gated[ref].append(item_id)
    if not gated:
        return source, []
    siblings = [
        source.model_copy(
            update={
                "id": f"{source.id}:{ref.key}",
                "items": items,
                "opens": phases.strictest(source.opens, ref.opens),
            }
        )
        for ref, items in gated.items()
    ]
    return (source.model_copy(update={"items": kept}) if kept else None), siblings


def _gate_sources(
    sources: list[LootSource], phases: ContentPhases, index: GateIndex, matched: _Matched
) -> list[LootSource]:
    by_id: dict[str, LootSource] = {}

    def put(source: LootSource) -> None:
        existing = by_id.get(source.id)
        if existing is not None and source.items is not None:
            source = existing.model_copy(
                update={"items": sorted({*(existing.items or []), *source.items})}
            )
        by_id[source.id] = source

    for source in sources:
        ref = _source_gate(source, phases, index, matched)
        if ref is not None:
            source = source.model_copy(update={"opens": phases.strictest(source.opens, ref.opens)})
        siblings: list[LootSource] = []
        if source.kind in INSTANCE_KINDS:
            source, siblings = _split_bosses(source, phases, index, matched)
        kept, item_siblings = _split_items(source, phases, index, matched)
        for each in [kept, *siblings, *item_siblings]:
            if each is not None:
                put(each)
    return sorted(by_id.values(), key=lambda source: (KIND_ORDER.index(source.kind), source.id))


def _quest_gate(entry: QuestSource, item_id: int, index: GateIndex, matched: _Matched) -> list[str]:
    gates: list[str] = []
    if entry.quest_id in index.quests:
        matched.add("quests", entry.quest_id)
        gates.append(index.quests[entry.quest_id].opens)
    faction = entry.required_rep_faction
    if faction is not None and faction in index.factions:
        matched.add("factions", faction)
        ref = index.factions[faction]
        if _standing_gated(ref, entry.required_rep_standing):
            gates.append(ref.opens)
    if item_id in index.items:
        matched.add("items", item_id)
        gates.append(index.items[item_id].opens)
    return gates


def _fill_rep_from_item(document: LootFile) -> LootFile:
    """A quest reward that is also a reputation reward needs that reputation
    (Earthstrike: quest 8573 states none, the same item is Cenarion Circle
    exalted)."""
    by_item = rep_gate_by_item(document.sources)
    quests: dict[str, list[QuestSource]] = {}
    for item_id, entries in document.quests.items():
        fallback = by_item.get(int(item_id))
        quests[item_id] = [
            entry.model_copy(
                update={"required_rep_faction": fallback[0], "required_rep_standing": fallback[1]}
            )
            if fallback is not None and entry.required_rep_faction is None
            else entry
            for entry in entries
        ]
    return document.model_copy(update={"quests": quests})


def _gate_quests(
    quests: dict[str, list[QuestSource]], phases: ContentPhases, index: GateIndex, matched: _Matched
) -> dict[str, list[QuestSource]]:
    gated: dict[str, list[QuestSource]] = {}
    for item_id, entries in quests.items():
        out = []
        for entry in entries:
            gates = _quest_gate(entry, int(item_id), index, matched)
            out.append(
                entry.model_copy(update={"opens": phases.strictest(entry.opens, *gates)})
                if gates
                else entry
            )
        gated[item_id] = out
    return gated


def _missing(phases: ContentPhases, matched: _Matched) -> list[str]:
    """Every id a closed gate lists that tonight's document does not name."""
    missing: list[str] = []
    for gate in phases.gates:
        if not phases.closed(gate.phase):
            continue
        for key_type in _KEY_TYPES:
            for entry in getattr(gate, key_type):
                if key_type == "kinds":
                    found = entry in matched.ids["kinds"]
                    label = f"{gate.key}: kind {entry}"
                elif key_type == "npcs":
                    found = (
                        entry.id in matched.ids["npcs"] or entry.name in matched.ids["npc_names"]
                    )
                    label = f"{gate.key}: npc {entry.id} {entry.name}"
                else:
                    found = entry.id in matched.ids[key_type]
                    label = f"{gate.key}: {_KEY_LABELS[key_type]} {entry.id} {entry.name}"
                if not found:
                    missing.append(label)
    return missing


def apply_content_phases(document: LootFile, phases: ContentPhases) -> tuple[LootFile, list[str]]:
    """Writes every closed gate into `document` and returns it with the ids
    tonight's sources do not name. Those are logged and kept: the gate
    applies again the night the source comes back.

    Runs after `apply_overlays` and before `apply_quest_opens_gate`, so a
    quest whose turn-in item is now gated is gated in turn.
    """
    index = phases.index()
    matched = _Matched()
    document = _fill_rep_from_item(merge_duplicate_quests(document, phases))
    out = document.model_copy(
        update={
            "sources": _gate_sources(document.sources, phases, index, matched),
            "quests": _gate_quests(document.quests, phases, index, matched),
        }
    )
    missing = _missing(phases, matched)
    if missing:
        logger.warning(
            "content-phases: %d gated ids are not in this build's sources; the gates are kept "
            "and apply when the source returns: %s",
            len(missing),
            "; ".join(missing),
        )
    return out, missing
