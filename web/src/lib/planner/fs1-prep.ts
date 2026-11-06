// web/src/lib/planner/fs1-prep.ts
// The FS1 export's raid-prep sections (owner decision D14), written by the addon's
// RaidPrep.lua through Codec.encodeFS1 and read by the API's api/internal/fs1/prep.go:
//
//   school=<fire>:<nature>:<frost>:<shadow>:<arcane>   resistance totals
//   supplies=<item_id>:<count>,...                     Consumable-class stacks in the bags
//   locks=<item_id>,...                                the keyring's items
//   quests=<quest_id>:<0|1>,...                        attunement quests, 1 when done
//
// Read leniently, unlike the other known sections: a malformed entry is dropped and a
// malformed school= is left out, never refusing the code. They are the client's own
// reading, never something a player types, and a stale planner must still load the build.

export interface FS1Resistances {
  fire: number;
  nature: number;
  frost: number;
  shadow: number;
  arcane: number;
}

export interface FS1Prep {
  resistances?: FS1Resistances;
  supplies?: { itemId: number; count: number }[];
  keyring?: number[];
  quests?: { questId: number; done: boolean }[];
}

export const PREP_SECTIONS = ['school', 'supplies', 'locks', 'quests'] as const;
type PrepSection = (typeof PREP_SECTIONS)[number];

const DIGITS = /^\d+$/;

function entries(field: string): string[] {
  return field === '' ? [] : field.split(',').filter((entry) => entry !== '');
}

function resistances(field: string): FS1Resistances | undefined {
  const parts = field.split(':');
  if (parts.length !== 5 || !parts.every((part) => DIGITS.test(part))) return undefined;
  const [fire, nature, frost, shadow, arcane] = parts.map(Number) as [number, number, number, number, number];
  return { fire, nature, frost, shadow, arcane };
}

function pairs(field: string, second: RegExp): [number, string][] {
  return entries(field).flatMap((entry): [number, string][] => {
    const [id, value, ...rest] = entry.split(':');
    if (
      id === undefined ||
      value === undefined ||
      rest.length > 0 ||
      !DIGITS.test(id) ||
      !second.test(value)
    ) {
      return [];
    }
    return [[Number(id), value]];
  });
}

export function isPrepSection(name: string): name is PrepSection {
  return (PREP_SECTIONS as readonly string[]).includes(name);
}

/** One raid-prep section's reading, as the FS1Build fields it fills. */
export function decodePrepSection(name: PrepSection, field: string): FS1Prep {
  switch (name) {
    case 'school': {
      const read = resistances(field);
      return read === undefined ? {} : { resistances: read };
    }
    case 'supplies':
      return { supplies: pairs(field, DIGITS).map(([itemId, count]) => ({ itemId, count: Number(count) })) };
    case 'locks':
      return {
        keyring: entries(field)
          .filter((entry) => DIGITS.test(entry))
          .map(Number),
      };
    case 'quests':
      return { quests: pairs(field, /^[01]$/).map(([questId, done]) => ({ questId, done: done === '1' })) };
  }
}
