// web/src/lib/report/raid-cooldowns.ts
// The buffs a raid leader looks for first: the ones that move a whole pull. Matched by
// name against the fight's aura tracks, since Forever's spell ids are not final until the
// beta client; a name here that the fight never saw simply does not appear.
export const RAID_COOLDOWNS: readonly string[] = [
  'Bloodlust',
  'Heroism',
  'Drums of Battle',
  'Drums of War',
  'Power Infusion',
  'Innervate',
  'Mana Tide Totem',
  'Tranquility',
  'Shield Wall',
  'Last Stand',
  'Divine Shield',
  'Lay on Hands',
  'Divine Protection',
  'Recklessness',
  'Death Wish',
  'Avenging Wrath',
  'Divine Favor',
  'Inner Focus',
  'Arcane Power',
  'Combustion',
  'Presence of Mind',
  'Icy Veins',
  'Evocation',
  'Adrenaline Rush',
  'Blade Flurry',
  'Cold Blood',
  'Rapid Fire',
  'Bestial Wrath',
  "Nature's Swiftness",
  'Elemental Mastery',
  'Barkskin',
  'Frenzied Regeneration',
  'Berserk',
  'Metamorphosis',
  'Fear Ward',
  'Battle Shout',
  'Ancestral Fortitude',
  'Ardent Defender',
  'Guardian of Ancient Kings',
  'Revival',
  'Life Cocoon',
  'Pain Suppression',
  'Guardian Spirit',
  'Spirit Link Totem',
  'Healing Tide Totem',
  'Ascendance',
  'Aspect of the Turtle',
  'Ice Block',
  'Vampiric Embrace',
  'Rallying Cry',
  'Demoralizing Shout',
  'Challenging Shout',
];

const NAMES = new Set(RAID_COOLDOWNS.map((name) => name.toLowerCase()));

export function isRaidCooldown(name: string): boolean {
  return NAMES.has(name.toLowerCase());
}
