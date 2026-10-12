// web/src/lib/race-faction.ts
// The faction a race plays (L8.4 fix round 3 item 15), from one fixed table: Classic's eight
// races, which addon/ForeverSixty/Rules.lua's PAIRS lists. Forever's six new pairs (gnome
// priest, human hunter, dwarf shaman, orc mage, troll warlock, undead paladin) recombine those
// races with classes and add no race, so no faction changes. The profile's tooltip mock takes
// its faction line from here, so a visitor and the owner see the same mock; an unknown race
// gives no faction (the line is left out).
import type { Faction } from './faction-mark';

export const RACE_FACTION: Readonly<Record<string, Faction>> = {
  human: 'alliance',
  dwarf: 'alliance',
  nightelf: 'alliance',
  gnome: 'alliance',
  orc: 'horde',
  undead: 'horde',
  tauren: 'horde',
  troll: 'horde',
};

/** The client's own token for undead (Rules.lua's RACE_ALIAS). */
const ALIAS: Readonly<Record<string, string>> = { scourge: 'undead' };

/** The race's faction, keyed the way Rules.lua keys it (case and everything but letters
 *  stripped), so "Night Elf", "NightElf" and "night-elf" read alike; null when unknown. */
export function factionOfRace(race: string | null | undefined): Faction | null {
  if (typeof race !== 'string') return null;
  const key = race.toLowerCase().replace(/[^a-z]/g, '');
  return RACE_FACTION[ALIAS[key] ?? key] ?? null;
}
