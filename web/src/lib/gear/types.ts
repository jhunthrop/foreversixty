// web/src/lib/gear/types.ts
// The gear panel fields of GET /v1/characters/{region}/{ruleset}/{name}, field for field from
// api/openapi.yaml's `CharacterGear`, `CharacterGearSlot` and `GearSlotName` (Task L6.7): the 17
// slots of the character's newest addon export in the export's order, each with what it wears,
// the band's pick and the gap, and `gear_gap_total` beside them. Nothing is invented: a pick the
// data does not have is null, and a gap that cannot be scored is null.

/** `GearSlotName`: the planner's slot vocabulary, in the export's order. */
export type GearSlotName =
  | 'head'
  | 'neck'
  | 'shoulder'
  | 'back'
  | 'chest'
  | 'wrist'
  | 'hands'
  | 'waist'
  | 'legs'
  | 'feet'
  | 'finger1'
  | 'finger2'
  | 'trinket1'
  | 'trinket2'
  | 'main_hand'
  | 'off_hand'
  | 'ranged';

/** `CharacterGearSlot.worn`: the item worn in the slot, from the newest addon export. */
export interface GearWorn {
  id: number;
  /** Null when the export's data build does not have the item. */
  name: string | null;
  /** Null when the export's data build does not have the item. */
  item_level: number | null;
}

/** `CharacterGearSlot.pick`: the band's best-in-slot pick for the spec and faction. */
export interface GearPick {
  id: number;
  name: string;
  /** Where the item comes from, as a player knows it: "Wailing Caverns: Mutanus the Devourer". */
  source: string;
  /** The drop chance as a fraction (0.333333 is one in three); null when the source states no
   *  rate (every quest, vendor, crafted, reputation and PvP source). */
  chance: number | null;
}

/** `CharacterGearSlot`: one slot of the gear panel. */
export interface CharacterGearSlot {
  slot: GearSlotName;
  /** Null when the slot is empty in the export. */
  worn: GearWorn | null;
  /** Null when the band has no pick for the slot (or the spec, level or faction is unknown). */
  pick: GearPick | null;
  /** The Gear component points the pick would add over the worn item, one decimal; 0 when the
   *  worn item is the pick or scores at least as much; null when there is no pick or an item
   *  cannot be scored. */
  gap: number | null;
}

/** `CharacterGear`: both fields ride on the character answer. */
export interface CharacterGear {
  /** All 17 slots when the character has an export; `[]` when it has none (never null). */
  gear: CharacterGearSlot[];
  /** The plain sum of every non-null gap, one decimal; null when no slot has a gap. It is not
   *  the Gear component's distance to its top (that is measured against the band's best set). */
  gear_gap_total: number | null;
  /** The spec the picks are for, the site's slug ("warrior-fury"); null with no export. */
  gear_spec?: string | null;
  /** Where `gear_spec` came from: a logged fight, the export's talents, or the class default. */
  gear_spec_source?: 'fight' | 'talents' | 'default' | null;
}
