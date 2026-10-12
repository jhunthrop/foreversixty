// web/src/fixtures/character-gear.ts
// A GET /v1/characters/... gear panel (api/openapi.yaml `CharacterGear`, Task L6.7) covering
// every row state the profile's Gear section draws (Task L8.6), for vitest and for the e2e
// specs (tests/e2e/support/character-fixture.ts re-exports it):
//
//   head, wrist, waist, main_hand  an upgrade: worn -> pick, a drop chance, a gap
//   neck                           an empty slot with a pick
//   shoulder, hands, feet,         the worn item is the pick (gap 0): "Best for your band"
//   trinket1
//   back                           an upgrade whose drop chance is under 1% ("<1%")
//   chest                          an upgrade from a source with no rate (crafted: no chance)
//   legs                           a worn item the data build does not have: no name, null gap
//   finger1, finger2               the rings worn the other way round: gap 0, ids differ
//   trinket2                       empty, and no pick
//   off_hand                       empty under a worn two-hander, with a pick: gap 0
//   ranged                         worn, and no pick
//
// The seven non-zero gaps add up to `gear_gap_total`, 70 (18.4 + 12.6 + 9.1 + 11.3 + 4.2 + 7.6
// + 6.8): a different number from the score card's "88 below its top" on purpose, since the two
// measure different things (L8.6 ruling 3).
import type { CharacterGear, CharacterGearSlot } from '../lib/gear/types';

export const GEAR_FIXTURE: CharacterGearSlot[] = [
  {
    slot: 'head',
    worn: { id: 10504, name: 'Green Lens', item_level: 58 },
    pick: { id: 16921, name: 'Halo of Transcendence', source: 'Onyxia’s Lair: Onyxia', chance: 0.169 },
    gap: 18.4,
  },
  {
    slot: 'neck',
    worn: null,
    pick: { id: 18814, name: 'Choker of the Fire Lord', source: 'Molten Core: Ragnaros', chance: 0.12 },
    gap: 12.6,
  },
  {
    slot: 'shoulder',
    worn: { id: 16924, name: 'Pauldrons of Transcendence', item_level: 76 },
    pick: {
      id: 16924,
      name: 'Pauldrons of Transcendence',
      source: 'Blackwing Lair: Chromaggus',
      chance: 0.18,
    },
    gap: 0,
  },
  {
    slot: 'back',
    worn: { id: 13386, name: 'Archivist Cape', item_level: 61 },
    pick: { id: 18389, name: 'Cloak of the Cosmos', source: 'Dire Maul: Tendris Warpwood', chance: 0.008 },
    gap: 9.1,
  },
  {
    slot: 'chest',
    worn: { id: 13346, name: 'Robes of the Exalted', item_level: 58 },
    pick: { id: 14152, name: 'Robe of the Archmage', source: 'Tailoring', chance: null },
    gap: 11.3,
  },
  {
    slot: 'wrist',
    worn: { id: 11766, name: 'Flameweave Cuffs', item_level: 56 },
    pick: { id: 16819, name: 'Vambraces of Prophecy', source: 'Molten Core: trash', chance: 0.02 },
    gap: 4.2,
  },
  {
    slot: 'hands',
    worn: { id: 16812, name: 'Gloves of Prophecy', item_level: 66 },
    pick: { id: 16812, name: 'Gloves of Prophecy', source: 'Molten Core: Gehennas', chance: 0.16 },
    gap: 0,
  },
  {
    slot: 'waist',
    worn: { id: 11662, name: 'Ban’thok Sash', item_level: 52 },
    pick: { id: 16817, name: 'Girdle of Prophecy', source: 'Molten Core: trash', chance: 0.02 },
    gap: 7.6,
  },
  {
    slot: 'legs',
    worn: { id: 99001, name: null, item_level: null },
    pick: { id: 16814, name: 'Pants of Prophecy', source: 'Molten Core: Magmadar', chance: 0.18 },
    gap: null,
  },
  {
    slot: 'feet',
    worn: { id: 16811, name: 'Boots of Prophecy', item_level: 66 },
    pick: { id: 16811, name: 'Boots of Prophecy', source: 'Molten Core: Shazzrah', chance: 0.17 },
    gap: 0,
  },
  {
    slot: 'finger1',
    worn: { id: 17713, name: 'Blackstone Ring', item_level: 63 },
    pick: { id: 19140, name: 'Cauterizing Band', source: 'Molten Core: trash', chance: 0.01 },
    gap: 0,
  },
  {
    slot: 'finger2',
    worn: { id: 19140, name: 'Cauterizing Band', item_level: 71 },
    pick: {
      id: 17713,
      name: 'Blackstone Ring',
      source: 'Blackrock Depths: Lord Incendius',
      chance: 0.2,
    },
    gap: 0,
  },
  {
    slot: 'trinket1',
    worn: { id: 18469, name: 'Royal Seal of Eldre’Thalas', item_level: 62 },
    pick: { id: 18469, name: 'Royal Seal of Eldre’Thalas', source: 'Dire Maul: Tribute', chance: null },
    gap: 0,
  },
  { slot: 'trinket2', worn: null, pick: null, gap: null },
  {
    slot: 'main_hand',
    worn: { id: 13000, name: 'Staff of Hale Magefire', item_level: 60 },
    pick: {
      id: 11923,
      name: 'The Hammer of Grace',
      source: 'Blackrock Spire: Overlord Wyrmthalak',
      chance: 0.11,
    },
    gap: 6.8,
  },
  {
    slot: 'off_hand',
    worn: null,
    pick: { id: 11904, name: 'Tome of Shadow Force', source: 'Blackrock Spire: Gyth', chance: 0.09 },
    gap: 0,
  },
  {
    slot: 'ranged',
    worn: { id: 13396, name: 'Skul’s Ghastly Touch', item_level: 60 },
    pick: null,
    gap: null,
  },
];

/** Both fields as the API sends them: the gaps' plain sum, one decimal. */
export const GEAR_PANEL_FIXTURE: CharacterGear = { gear: GEAR_FIXTURE, gear_gap_total: 70 };

/** Each fixture item's quality (0 poor to 5 legendary), as items/priest.json would carry it, so
 *  the e2e stub of that file colours the names; 99001 is not in the build. */
export const GEAR_FIXTURE_QUALITIES: Readonly<Record<number, number>> = {
  10504: 2,
  16921: 4,
  18814: 4,
  16924: 4,
  13386: 3,
  18389: 3,
  13346: 3,
  14152: 4,
  11766: 2,
  16819: 4,
  16812: 4,
  16817: 4,
  11662: 3,
  16814: 4,
  16811: 4,
  17713: 3,
  19140: 4,
  18469: 3,
  13000: 3,
  11923: 3,
  11904: 3,
  13396: 3,
};
