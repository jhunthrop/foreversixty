// web/src/lib/planner/types.ts
// Shapes are copied verbatim from docs/superpowers/specs/2026-09-13-phase-1-interfaces.md.
// The API validates the same rules server-side; anything renamed here has to be renamed there.
import type { SourceKind } from '../sources';

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
}

export interface TalentTree {
  id: number;
  name: string;
  /** Left-to-right order as in the client. */
  position: number;
  talents: Talent[];
}

export interface TalentFile {
  build: string;
  class_id: number;
  class_slug: string;
  trees: TalentTree[];
}

/** The 17 planner gear slots, in the order the slot grid lays them out. */
export const SLOTS = [
  'head',
  'neck',
  'shoulder',
  'back',
  'chest',
  'wrist',
  'hands',
  'waist',
  'legs',
  'feet',
  'finger1',
  'finger2',
  'trinket1',
  'trinket2',
  'main_hand',
  'off_hand',
  'ranged',
] as const;
export type Slot = (typeof SLOTS)[number];

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
  'spell_power',
  'healing',
  'attack_power',
  'defense',
  'dodge',
  'parry',
  'block',
  'mp5',
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
  spell_power: 'Spell power',
  healing: 'Healing',
  attack_power: 'Attack power',
  defense: 'Defense',
  dodge: 'Dodge',
  parry: 'Parry',
  block: 'Block',
  mp5: 'Mana per 5',
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
  gear: Gear;
  title?: string;
}

/** What GET /v1/builds/{id} returns and what the API inlines into `data-build`. */
export interface BuildRecord extends BuildDraft {
  id: string;
  title?: string;
  created_at: string;
  views: number;
}
