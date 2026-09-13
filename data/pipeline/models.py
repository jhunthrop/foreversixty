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


class TalentTree(BaseModel):
    id: int
    name: str
    position: int
    talents: list[TalentEntry]


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
