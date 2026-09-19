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


class SpecRecord(BaseModel):
    spec: str
    class_slug: str
    spec_slug: str
    name: str
    role: str
    tree_index: int


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
