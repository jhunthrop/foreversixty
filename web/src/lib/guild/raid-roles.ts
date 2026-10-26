// web/src/lib/guild/raid-roles.ts
// The raid role a spec plays (spec 3.3: the role select is prefilled from the spec). One map for
// the strip, the band and the mock guild fixture.
import type { GuildRaidRole } from './api';

export const SPEC_ROLE: Readonly<Record<string, GuildRaidRole>> = {
  Protection: 'tank',
  Fury: 'dps',
  Arms: 'dps',
  Marksmanship: 'dps',
  'Beast Mastery': 'dps',
  Survival: 'dps',
  Combat: 'dps',
  Assassination: 'dps',
  Subtlety: 'dps',
  Affliction: 'dps',
  Demonology: 'dps',
  Destruction: 'dps',
  Fire: 'dps',
  Frost: 'dps',
  Arcane: 'dps',
  Holy: 'healer',
  Discipline: 'healer',
  Shadow: 'dps',
  Restoration: 'healer',
  Feral: 'dps',
  Balance: 'dps',
  Elemental: 'dps',
  Enhancement: 'dps',
  Retribution: 'dps',
};

const BY_LOWER_NAME = new Map(Object.entries(SPEC_ROLE).map(([spec, role]) => [spec.toLowerCase(), role]));

/** The role for a spec name in any case; `null` for an empty or unknown spec. */
export function roleForSpec(spec: string): GuildRaidRole | null {
  return BY_LOWER_NAME.get(spec.trim().toLowerCase()) ?? null;
}
