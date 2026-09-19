from pydantic import BaseModel


class Zone(BaseModel):
    id: int
    name: str
    map_id: int
    map_name: str
    parent_id: int | None
    level_min: int | None
    level_max: int | None


class Dungeon(BaseModel):
    id: int
    name: str
    map_id: int
    description: str
    order: int


class Item(BaseModel):
    id: int
    name: str
    quality: int
    item_level: int
    required_level: int
    class_id: int
    subclass_id: int
    inventory_type: int
    #: The random suffixes this item rolls, from the engine fork's
    #: `randomSuffixOptions` (parity contract 6.3). Empty for an item that
    #: rolls none, and for every row until `python -m pipeline loot` has run
    #: for the build -- `normalize` has no fork database to read.
    suffixes: list[int] = []
    #: "alliance_only", "horde_only", or "" for no restriction, from the
    #: fork's `factionRestriction` (contract 10.3's `SimItem` field). The
    #: client states none of this on build 1.60.1.69893, so like `suffixes`
    #: it is empty until `loot` has run.
    faction_restriction: str = ""


class Spell(BaseModel):
    id: int
    name: str


class Source(BaseModel):
    label: str
    url: str
    kind: str


class ForeverChange(BaseModel):
    text: str
    sources: list[Source]


class PlayableClass(BaseModel):
    id: int
    name: str
    slug: str
    color: str
    forever_changes: list[ForeverChange] = []


class PlayableRace(BaseModel):
    id: int
    name: str
    slug: str
    faction: str
    placeholder: bool = False
    forever_changes: list[ForeverChange] = []


class Combo(BaseModel):
    race_id: int
    class_id: int
    new_in_forever: bool


class TalentNode(BaseModel):
    id: int
    tab_id: int
    tab_name: str
    class_id: int
    tier: int
    column: int
    spell_ids: list[int]
    prereq_talent_id: int | None


class TalentRank(BaseModel):
    spell_id: int
    description: str


class TalentEntry(BaseModel):
    id: int
    name: str
    icon: str
    max_rank: int
    tier: int
    column: int
    prereq_talent_id: int | None
    prereq_rank: int | None
    ranks: list[TalentRank]
    #: The client's own spell for the talent, which is what every consumer
    #: keyed by spell id (the sim, the report's planner link) will see written
    #: by the game. Appended last, like `background` below: the emitted key
    #: order is the schema's only compatibility surface, so new fields go on
    #: the end and existing ones never move.
    spell_id: int


class TalentTree(BaseModel):
    id: int
    name: str
    position: int
    talents: list[TalentEntry]
    #: The tab's `BackgroundFile`, lowercased. The site fetches the processed
    #: image at `/data/<build>/trees/<background>.webp`.
    background: str


class ClassTalents(BaseModel):
    build: str
    class_id: int
    class_slug: str
    trees: list[TalentTree]


class GearItem(BaseModel):
    id: int
    name: str
    icon: str
    slot: str
    quality: int
    required_level: int
    item_level: int
    armor: int
    stats: dict[str, int]
    set_id: int | None
    unique: bool


class ClassItems(BaseModel):
    build: str
    class_slug: str
    items: list[GearItem]


class ItemSetBonus(BaseModel):
    pieces: int
    description: str


class ItemSetRecord(BaseModel):
    id: int
    name: str
    item_ids: list[int]
    bonuses: list[ItemSetBonus]


class SuffixRecord(BaseModel):
    id: int
    name: str
    stats: dict[str, float]


class EnchantRecord(BaseModel):
    #: The effect id. NOT unique: the fork's own table repeats an effect
    #: across the slots it can go in (19 of its 150 ids), which is why
    #: `spell_id` and `item_id` are here too.
    id: int
    name: str
    icon: str
    #: The planner slot names this enchant can be applied in.
    slots: list[str]
    #: The shape of item it needs: normal, two_hand, shield, kit or staff.
    item_types: list[str]
    #: Class slugs allowed to use it. Empty means no restriction.
    classes: list[str]
    stats: dict[str, float]
    phase: int
    #: Appended after the contract's keys, for telling two rows with the
    #: same effect id apart. 0 where the fork states neither.
    spell_id: int
    item_id: int


class LootBoss(BaseModel):
    id: str
    #: Empty where the fork database names no NPC for the id. 35 of the 67
    #: raid bosses and 39 of the 230 dungeon bosses on the pinned fork are
    #: in that state; an invented name would be worse than a blank one.
    name: str
    npc_id: int
    items: list[int]


class LootSource(BaseModel):
    """One place loot comes from (parity contract 6.1).

    Every key after `name` is optional and omitted when it does not apply
    to the kind, which is what `write_document`'s `exclude_none` is for: a
    crafted source carries `profession` and `items`, a raid carries
    `zone_id`, `bosses` and `trash`.
    """

    id: str
    kind: str
    name: str
    zone_id: int | None = None
    #: A phase name from api/internal/phase. None means open from launch.
    #: Never set by the generator -- only by a curated overlay, because the
    #: databases state no dates.
    opens: str | None = None
    profession: str | None = None
    faction_id: int | None = None
    standing: str | None = None
    rank: int | None = None
    bosses: list[LootBoss] | None = None
    trash: list[int] | None = None
    items: list[int] | None = None


class LootFile(BaseModel):
    sources: list[LootSource]


class LootSourcePatch(BaseModel):
    """Contract 10.4's `replace` entry: a loot source with every key but
    `id` optional, so gating a raid is one line instead of a restatement
    of its eleven bosses. Only the keys the file actually writes are
    applied -- `model_dump(exclude_unset=True)` is what tells "absent"
    from "explicitly null".

    `kind` is accepted but may not change: a source's kind decides the
    shape of its id, which candidates carry as `drop:<source id>`, so
    changing it would orphan every reference rather than edit one.
    """

    id: str
    kind: str | None = None
    name: str | None = None
    zone_id: int | None = None
    opens: str | None = None
    profession: str | None = None
    faction_id: int | None = None
    standing: str | None = None
    rank: int | None = None
    bosses: list[LootBoss] | None = None
    trash: list[int] | None = None
    items: list[int] | None = None

    model_config = {"extra": "forbid"}


class LootOverlay(BaseModel):
    """One `data/curated/loot/*.json`, in contract 10.4's shape.

    `sources` and `notes` are the curated provenance every hand-maintained
    fact here carries; `add`, `replace` and `remove` are the three
    operations, as flat arrays.
    """

    sources: list[Source] = []
    notes: str = ""
    add: list[LootSource] = []
    replace: list[LootSourcePatch] = []
    remove: list[str] = []

    model_config = {"extra": "forbid"}


class SpecRecord(BaseModel):
    spec: str
    class_slug: str
    spec_slug: str
    name: str
    role: str
    tree_index: int
    #: The stat `StatWeights` normalises to 1.0 for this spec (parity
    #: contract 1.4). Appended last, like `TalentEntry.spell_id`: the
    #: emitted key order is the generated files' only compatibility
    #: surface, so new fields go on the end and existing ones never move.
    reference_stat: str


class PhaseBoundary(BaseModel):
    """One content phase and the instant it opens (parity contract 10.4)."""

    name: str
    start: str


class AplDocument(BaseModel):
    spec: str
    state: str
    rotation: dict
    sources: list[Source] = []
    notes: str = ""
    #: Spell ids this rotation names on purpose that the pinned engine build
    #: cannot act on yet -- a talent-granted ability with no spell file, or an
    #: aura nothing registers. The notes say so in prose; this is the same
    #: statement in a form a test can read, and the engine lane's rotation
    #: smoke test holds the engine's own unknown-action warnings to exactly
    #: this set. Empty is the normal case.
    inert: list[int] = []


class SpellEffectConstant(BaseModel):
    index: int
    effect: int
    aura: int
    #: The amount the client shows, via pipeline.spelltext.effect_amount: the
    #: 1.60 client states it outright, Classic Era states it one short with the
    #: last point in EffectDieSides. One number, not two columns to recombine.
    amount: int
    sp_coefficient: float
    ap_coefficient: float
    period_ms: int
    misc_value: int
    trigger_spell: int


class SpellConstant(BaseModel):
    name: str
    rank: int
    school_mask: int
    cast_time_ms: int
    gcd_ms: int
    cooldown_ms: int
    category_cooldown_ms: int
    duration_ms: int
    cost: int
    cost_type: int
    spell_level: int
    family_mask: list[int]
    effects: list[SpellEffectConstant]


class ClassSpellConstants(BaseModel):
    build: str
    class_slug: str
    family: int
    spells: dict[str, SpellConstant]


class ConsumableRecord(BaseModel):
    id: int
    name: str
    quality: int
    required_level: int
    spell_ids: list[int]
