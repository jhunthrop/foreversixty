"""The curated default rotation per spec, checked against the engine's schema.

Design section 2.4: each spec ships one default APL, stored as data from day
one so sharing and editing are later features rather than a redesign. The
contract puts them in `data/curated/apl/<spec_slug>.json` as an APLRotation in
its JSON form plus sources and notes, like every other curated fact.

Validation is the engine's own schema. `json_format.Parse` into the generated
`APLRotation` rejects an unknown field, so there is no second schema to keep in
step -- when the engine adds an APL value, regenerating the bindings is the
whole change.

Spell ids and ranks are checked against the build too, in tests/test_apl.py.
That matters more than it sounds: the engine's checked-in preset rotations are
stale on ranks (their Frostbolt is rank 10 where 25304 is Rank 11 on both the
Era and the beta client), so a preset copied without checking names a spell the
engine cannot resolve.
"""

from __future__ import annotations

import json
from collections.abc import Iterator
from pathlib import Path
from typing import Any

from google.protobuf import json_format

from pipeline.curated import SOURCE_KINDS
from pipeline.models import AplDocument
from pipeline.simproto import apl

APL_STATES = frozenset({"written", "unwritten"})

#: What a spec whose rotation nobody has written yet carries. An empty priority
#: list is a spec the sim reports as unsupported; a wrong one would be a spec
#: it silently sims badly.
EMPTY_ROTATION: dict[str, Any] = {"type": "TypeAPL", "priorityList": []}


class AplError(SystemExit):
    """A curated rotation the pipeline will not publish."""


def parse_rotation(rotation: dict[str, Any]) -> apl.APLRotation:
    """Parse through the engine's own APLRotation, rejecting unknown fields."""
    try:
        return json_format.Parse(json.dumps(rotation), apl.APLRotation())
    except json_format.ParseError as error:
        raise AplError(f"rotation is not a valid APLRotation: {error}") from error


def action_ids(node: Any) -> Iterator[tuple[int, int]]:
    """Every (spell id, rank) an ActionID in the document names.

    An ActionID is the only object with an integer `spellId`; the outer
    `spellId` and `auraId` keys hold the ActionID itself. Item and other
    actions carry no spell id and are skipped. A reference with no `rank` is
    rank 0, which is what the constants file records for a rankless spell.
    """
    if isinstance(node, dict):
        if isinstance(node.get("spellId"), int):
            yield node["spellId"], int(node.get("rank", 0))
        for value in node.values():
            yield from action_ids(value)
    elif isinstance(node, list):
        for value in node:
            yield from action_ids(value)


def aura_reference_ids(node: Any) -> Iterator[tuple[int, int]]:
    """Every (spell id, rank) referenced as an aura -- `auraIsActive`,
    `auraNumStacks`, `auraRemainingTime`, and the engine's other aura-value
    shapes all carry their ActionID under the `auraId` key, never `spellId`.
    That is what lets `ENGINE_AURA_IDS` tell an aura reference apart from a
    `castSpell` action naming the same id: only the former key means "check
    whether this aura is active/stacked/expiring", never "cast this".
    """
    if isinstance(node, dict):
        aura_id = node.get("auraId")
        if isinstance(aura_id, dict) and isinstance(aura_id.get("spellId"), int):
            yield aura_id["spellId"], int(aura_id.get("rank", 0))
        for value in node.values():
            yield from aura_reference_ids(value)
    elif isinstance(node, list):
        for value in node:
            yield from aura_reference_ids(value)


#: Aura ids the engine itself registers under a different number than the
#: client's own copy of the same ability, confirmed by reading the engine's
#: Go source (not guessed, and not something spellconst -- client data --
#: can tell us on its own). `test_every_spell_the_rotations_name_exists_with_that_rank`
#: accepts an id from this table only where a rotation references it as an
#: aura (see `aura_reference_ids`); a `castSpell` action must still name an
#: id spellconst has, exception table or not, because that is the id the
#: engine is actually asked to cast. Delete an entry here the moment the
#: engine registers the client's id instead, or spellconst grows an entry
#: for the engine's id -- whichever the engine/data lanes land first.
ENGINE_AURA_IDS: dict[int, str] = {}


def unchecked_engine_aura_ids(rotation: dict[str, Any]) -> set[int]:
    """`ENGINE_AURA_IDS` entries this rotation may cite without a spellconst
    entry: every occurrence of the id in this rotation is an aura reference,
    and none is a `castSpell` target."""
    cast_ids = {spell_id for spell_id, _rank, _tag in cast_spell_action_ids(rotation)}
    aura_ids = {spell_id for spell_id, _rank in aura_reference_ids(rotation)}
    return {
        spell_id
        for spell_id in ENGINE_AURA_IDS
        if spell_id in aura_ids and spell_id not in cast_ids
    }


def cast_spell_action_ids(node: Any) -> Iterator[tuple[int, int, int]]:
    """Every (spell id, rank, tag) a `castSpell` action names.

    Unlike `action_ids`, this only looks at the ActionID a cast actually
    resolves against -- not every id an aura or condition happens to
    reference -- because tag only matters for what gets cast. The fork
    registers some abilities twice under one spell id: a direct spell at tag
    0 (`SpellFlagNoOnCastComplete`, no GCD) and, for an on-next-swing
    ability, an APL-castable queue action at a nonzero tag
    (`core.ActionID.WithTag` in sim/core/agent.go). `core.ProtoToActionID`
    ignores rank, so a cast that omits tag resolves to tag 0 -- the direct
    spell -- and is cast every APL iteration as a free swing.
    """
    if isinstance(node, dict):
        cast = node.get("castSpell")
        if isinstance(cast, dict):
            spell_id_ref = cast.get("spellId")
            if isinstance(spell_id_ref, dict) and isinstance(spell_id_ref.get("spellId"), int):
                yield (
                    spell_id_ref["spellId"],
                    int(spell_id_ref.get("rank", 0)),
                    int(spell_id_ref.get("tag", 0)),
                )
        for value in node.values():
            yield from cast_spell_action_ids(value)
    elif isinstance(node, list):
        for value in node:
            yield from cast_spell_action_ids(value)


def load_apl(path: Path) -> AplDocument:
    document = AplDocument(**json.loads(path.read_text(encoding="utf-8")))
    if document.spec != path.stem:
        raise AplError(f"{path} declares spec {document.spec!r}; the filename says {path.stem!r}")
    if document.state not in APL_STATES:
        raise AplError(f"{path} has state {document.state!r}; use one of {sorted(APL_STATES)}")
    for source in document.sources:
        if source.kind not in SOURCE_KINDS:
            raise AplError(
                f"{path} has source kind {source.kind!r}; use one of {sorted(SOURCE_KINDS)}"
            )
    if document.state == "written" and not document.sources:
        raise AplError(f"{path} is written but cites nothing; a rotation states a fact")
    check_inert(path, document)
    parse_rotation(document.rotation)
    return document


def check_inert(path: Path, document: AplDocument) -> None:
    """Hold `inert` to the rotation it describes.

    `inert` is the machine-readable half of what a written rotation's notes
    already say in prose: the spell ids whose lines the pinned engine build
    warns about and skips rather than casts. The engine lane's rotation smoke
    test asserts the engine's own unknown-action warnings equal this set
    exactly, so an id the rotation never names could never be warned about and
    would make that assertion unsatisfiable -- a drifted declaration rather
    than a finding about the engine.
    """
    if not document.inert:
        return
    if document.state != "written":
        raise AplError(
            f"{path} is unwritten but declares inert lines {document.inert}; "
            f"an empty priority list names no spell"
        )
    seen = set()
    for spell_id in document.inert:
        if spell_id in seen:
            raise AplError(f"{path} lists inert spell {spell_id} twice")
        seen.add(spell_id)
    named = {spell_id for spell_id, _rank in action_ids(document.rotation)}
    named |= {spell_id for spell_id, _rank in aura_reference_ids(document.rotation)}
    for spell_id in document.inert:
        if spell_id not in named:
            raise AplError(
                f"{path} declares spell {spell_id} inert, but the rotation never names it"
            )


def load_all(apl_dir: Path = Path("curated/apl")) -> dict[str, AplDocument]:
    if not apl_dir.is_dir():
        raise AplError(f"missing {apl_dir}")
    return {path.stem: load_apl(path) for path in sorted(apl_dir.glob("*.json"))}
