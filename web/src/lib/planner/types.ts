// web/src/lib/planner/types.ts
// Shapes are copied verbatim from docs/superpowers/specs/2026-09-13-phase-1-interfaces.md.
// The API validates the same rules server-side; anything renamed here has to be renamed there.
import type { SourceKind } from '../sources';
import { COSMETIC_SLOTS, FRAME_SLOTS, type CosmeticSlot, type FrameSlot } from '../gear/slots';

/** A build spends at most this many talent points. */
export const MAX_POINTS = 51;
/** Tier `t` of a tree unlocks once `POINTS_PER_TIER * t` points sit in that tree. */
export const POINTS_PER_TIER = 5;
/** The first talent point is spent at this character level. */
export const FIRST_POINT_LEVEL = 10;
/** The level of a character that has spent no points. */
export const BASE_LEVEL = FIRST_POINT_LEVEL - 1;

export interface TalentRank {
  spell_id: number;
  description: string;
}

export interface Talent {
  id: number;
  name: string;
  icon: string;
  max_rank: number;
  /** 0-based row. */
  tier: number;
  /** 0-based column. */
  column: number;
  prereq_talent_id: number | null;
  prereq_rank: number | null;
  /** Exactly `max_rank` entries. */
  ranks: TalentRank[];
  /** The spell the client writes when the talent is learned. */
  spell_id: number;
}

export interface TalentTree {
  id: number;
  name: string;
  /** Left-to-right order as in the client. */
  position: number;
  talents: Talent[];
  /** Names the tree's art at `/data/<build>/trees/<background>.webp`. */
  background: string;
}

export interface TalentFile {
  build: string;
  class_id: number;
  class_slug: string;
  trees: TalentTree[];
}

/** One of the 17 planner gear slots: every frame slot but the cosmetic shirt and tabard. */
export type Slot = Exclude<FrameSlot, CosmeticSlot>;

/** The 17 planner gear slots, in the client's paperdoll order (lib/gear/slots.ts). */
export const SLOTS: readonly Slot[] = FRAME_SLOTS.filter((slot): slot is Slot => !COSMETIC_SLOTS.has(slot));

/** Item files say `finger` and `trinket`; the planner maps each to both numbered slots. */
export const SLOT_ALIASES: Record<string, Slot[]> = {
  finger: ['finger1', 'finger2'],
  trinket: ['trinket1', 'trinket2'],
};

export const STAT_KEYS = [
  'strength',
  'agility',
  'stamina',
  'intellect',
  'spirit',
  'armor',
  'crit',
  'hit',
  'expertise',
  'spell_power',
  'healing',
  'attack_power',
  'ranged_attack_power',
  'feral_attack_power',
  'defense',
  'dodge',
  'parry',
  'block',
  'block_value',
  'mp5',
  'spell_penetration',
  'fire_power',
  'frost_power',
  'shadow_power',
  'arcane_power',
  'nature_power',
  'holy_power',
  'fire_res',
  'frost_res',
  'nature_res',
  'shadow_res',
  'arcane_res',
] as const;
export type StatKey = (typeof STAT_KEYS)[number];

/** Human labels for the stat rows in the gear panel. */
export const STAT_LABELS: Record<StatKey, string> = {
  strength: 'Strength',
  agility: 'Agility',
  stamina: 'Stamina',
  intellect: 'Intellect',
  spirit: 'Spirit',
  armor: 'Armor',
  crit: 'Crit',
  hit: 'Hit',
  expertise: 'Expertise',
  spell_power: 'Spell power',
  healing: 'Healing',
  attack_power: 'Attack power',
  ranged_attack_power: 'Ranged attack power',
  feral_attack_power: 'Feral attack power',
  defense: 'Defense',
  dodge: 'Dodge',
  parry: 'Parry',
  block: 'Block',
  block_value: 'Block value',
  mp5: 'Mana per 5',
  spell_penetration: 'Spell penetration',
  fire_power: 'Fire spell damage',
  frost_power: 'Frost spell damage',
  shadow_power: 'Shadow spell damage',
  arcane_power: 'Arcane spell damage',
  nature_power: 'Nature spell damage',
  holy_power: 'Holy spell damage',
  fire_res: 'Fire resistance',
  frost_res: 'Frost resistance',
  nature_res: 'Nature resistance',
  shadow_res: 'Shadow resistance',
  arcane_res: 'Arcane resistance',
};

export const SLOT_LABELS: Record<Slot, string> = {
  head: 'Head',
  neck: 'Neck',
  shoulder: 'Shoulder',
  back: 'Back',
  chest: 'Chest',
  wrist: 'Wrist',
  hands: 'Hands',
  waist: 'Waist',
  legs: 'Legs',
  feet: 'Feet',
  finger1: 'Finger 1',
  finger2: 'Finger 2',
  trinket1: 'Trinket 1',
  trinket2: 'Trinket 2',
  main_hand: 'Main hand',
  off_hand: 'Off hand',
  ranged: 'Ranged',
};

export interface Item {
  id: number;
  name: string;
  icon: string;
  /** One of SLOTS, or the `finger`/`trinket` aliases. */
  slot: string;
  /** Client value: 0 poor … 5 legendary. */
  quality: number;
  required_level: number;
  item_level: number;
  armor: number;
  stats: Partial<Record<StatKey, number>>;
  set_id: number | null;
  unique: boolean;
  /**
   * The random-suffix ids this item can roll (contract 6.3: "items.json rows gain
   * `suffixes: [id, ...]` where the item rolls one"). Absent on a build the data lane has
   * not regenerated, and on every item that rolls none.
   */
  suffixes?: number[];
  /**
   * True for a classic-db row the client's own Item and item-effect tables do not name in
   * classic-db's shape (`data/pipeline/client_evidence.py`): missing, reworked or another
   * kind of item. Never a BiS pick. Absent (never fabricated `false`) on a build the data
   * lane has not regenerated with this field, and on every item the client or wowhead
   * already names.
   */
  client_unconfirmed?: boolean;
  /**
   * Where the row's numbers come from: `classic-db` is 1.12's own, `wowhead` the Forever
   * gear planner's, and absent or null the client's own.
   */
  stats_source?: string | null;
  /**
   * Weapon fields ([data] Tasks 5 and 6): min and max weapon damage, speed in seconds, and
   * the derived DPS. Optional the same way `suffixes` is above -- absent on a build the
   * data lane has not regenerated, and on every item that is not a weapon.
   */
  damage_min?: number;
  damage_max?: number;
  speed?: number;
  dps?: number;
  /** True for a two-handed weapon; rule 6's client mirror refuses an off-hand beside one. */
  two_hand?: boolean;
  /** The weapon's on-hit or on-use effect text, when it has one. */
  effect_text?: string;
  /** The rating-family stats an on-equip spell states as a literal percent (`hit: 2` is "2%"),
   *  a subset of `stats`; the tooltip words these as the client does. */
  percent_stats?: Partial<Record<StatKey, number>>;
  /** ItemSparse `Bonding` as a word (api/openapi.yaml `Item.bonding`, spec C6). Absent or null
   *  when the data states none (a classic-db or wowhead row, or a build not yet regenerated). */
  bonding?: 'pickup' | 'equip' | 'use' | 'quest' | null;
  /** The armour word from the item's class and subclass (`Item.armor_type`; a cloak is
   *  `cloth`); null for jewellery, a weapon or unknown; absent on a build not yet regenerated. */
  armor_type?: 'cloth' | 'leather' | 'mail' | 'plate' | 'shield' | null;
  /** The client's ItemClass, ItemSubClass and InventoryType, for an enchant's
   *  SpellEquippedItems rule (lib/gear/enchant-fit.ts); absent on a build not yet regenerated. */
  class_id?: number | null;
  subclass_id?: number | null;
  inventory_type?: number | null;
}

export interface ItemFile {
  build: string;
  class_slug: string;
  items: Item[];
}

export interface SetBonus {
  pieces: number;
  description: string;
}

export interface ItemSet {
  id: number;
  name: string;
  item_ids: number[];
  bonuses: SetBonus[];
}

export interface SourceRef {
  label: string;
  url: string;
  kind: SourceKind;
}

export interface ForeverChange {
  text: string;
  sources: SourceRef[];
}

export interface ClassRow {
  id: number;
  name: string;
  slug: string;
  color: string;
  forever_changes: ForeverChange[];
}

export interface RaceRow {
  id: number;
  name: string;
  slug: string;
  faction: string;
  forever_changes: ForeverChange[];
  /** True while the row comes from data/curated/ rather than the client tables. */
  placeholder?: boolean;
}

export interface Combo {
  race_id: number;
  class_id: number;
  new_in_forever: boolean;
}

export type Gear = Partial<Record<Slot, number>>;

/** What POST /v1/builds accepts. */
export interface BuildDraft {
  class_id: number;
  race_id: number;
  tree_version: string;
  point_order: number[];
  gear?: Gear;
  title?: string;
}

/** What GET /v1/builds/{id} returns and what the API inlines into `data-build`. */
export interface BuildRecord extends BuildDraft {
  id: string;
  created_at: string;
  views: number;
}
