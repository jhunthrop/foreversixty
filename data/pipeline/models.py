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


class PlayableClass(BaseModel):
    id: int
    name: str
    slug: str
    color: str


class PlayableRace(BaseModel):
    id: int
    name: str
    slug: str
    faction: str


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
