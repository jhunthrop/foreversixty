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
    parse_rotation(document.rotation)
    return document


def load_all(apl_dir: Path = Path("curated/apl")) -> dict[str, AplDocument]:
    if not apl_dir.is_dir():
        raise AplError(f"missing {apl_dir}")
    return {path.stem: load_apl(path) for path in sorted(apl_dir.glob("*.json"))}
