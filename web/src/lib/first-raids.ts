// web/src/lib/first-raids.ts
// The first raid tier as the site draws it (dates.json role `raids`, phases.json `raids-1`): each
// raid's name and the icon every page shows for it (home's countdowns, the guild's next raid).

export const FIRST_RAIDS = [
  { name: 'Barrow Deeps', icon: 'spell_shadow_raisedead' },
  { name: 'Hyjal Summit', icon: 'spell_nature_timestop' },
  { name: "Onyxia's Lair", icon: 'inv_misc_head_dragon_black' },
] as const;

/** A first-tier raid's icon by its name; undefined for anything else (an event, a dungeon). */
export function raidZoneIcon(zoneName: string): string | undefined {
  return FIRST_RAIDS.find((raid) => raid.name === zoneName)?.icon;
}
