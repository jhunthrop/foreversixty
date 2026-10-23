// web/src/lib/guides/race-table.ts
// The guide's Races table: one row per race the class can be, with the racial that matters
// for this kind of spec and what that racial does. The race list comes from the combos, the
// racial effects from the build's own race data (`races.json` `forever_changes`); the one
// editorial input is which racial matters to which kind of spec, kept here as a small table
// that a test holds to the race data, so a renamed racial fails a test instead of the page.
import { classRows, raceRows, racesForClass } from '../planner/reference';
import type { RaceRow } from '../planner/types';

export type SpecKind = 'caster' | 'melee' | 'ranged' | 'healer' | 'tank';

/** Every racial the data names for a race. */
export const RACIAL_NAMES: Readonly<Record<string, readonly string[]>> = {
  human: ['Will to Survive', 'Perception', 'Sword Specialization', 'The Human Spirit'],
  orc: ['Blood Fury', 'Shatter Curse', 'Axe Specialization', 'Hardiness'],
  dwarf: ['Stoneform', 'Find Treasure', 'Mace Specialization', 'Big Game Hunter'],
  'night-elf': ["Elune's Light", 'Shadowmeld', 'Quickness', 'Wisp Spirit'],
  undead: ['Will of the Forsaken', 'Cannibalize', 'Touch of the Grave'],
  tauren: ['War Stomp', 'Cultivation', 'Plainsrunning', 'Endurance'],
  gnome: ['Escape Artist', 'Eureka!', 'Expansive Mind', 'Engineering Specialization'],
  troll: ['Berserking', 'Rapid Regeneration', 'Beast Slaying', 'Regeneration'],
  'high-order-skyborne': ['Walk on Air', 'Read Ley Line', 'Wind Blessed', 'Elemental Insight'],
  'windshaper-skyborne': ['Walk on Air', 'Read Ley Line', 'Wind Blessed', 'Elemental Insight'],
};

type Picks = { default: string } & Partial<Record<SpecKind, string>>;

/** The racial that matters, by kind of spec. */
export const RACIAL_PICKS: Readonly<Record<string, Picks>> = {
  human: {
    default: 'Sword Specialization',
    caster: 'The Human Spirit',
    healer: 'The Human Spirit',
    ranged: 'Will to Survive',
  },
  orc: { default: 'Blood Fury', tank: 'Shatter Curse' },
  dwarf: { default: 'Mace Specialization', healer: 'Stoneform', caster: 'Stoneform', tank: 'Stoneform' },
  'night-elf': { default: "Elune's Light", tank: 'Quickness' },
  undead: {
    default: 'Will of the Forsaken',
    melee: 'Touch of the Grave',
    ranged: 'Touch of the Grave',
  },
  tauren: { default: 'Endurance' },
  gnome: { default: 'Eureka!', healer: 'Expansive Mind', tank: 'Escape Artist' },
  troll: { default: 'Berserking', tank: 'Rapid Regeneration' },
  'high-order-skyborne': { default: 'Wind Blessed', healer: 'Read Ley Line' },
  'windshaper-skyborne': { default: 'Wind Blessed', healer: 'Read Ley Line' },
};

const CASTER_CLASSES = new Set(['mage', 'warlock']);
const CASTER_SPECS = new Set(['priest-shadow', 'shaman-elemental', 'druid-balance']);

/** Which kind of spec this is, for picking its racial. `spec` is the full id
 *  (`mage-frost`); `role` is the guide's own frontmatter role. */
export function specKindOf(classSlug: string, spec: string, role: 'dps' | 'healer' | 'tank'): SpecKind {
  if (role === 'healer') return 'healer';
  if (role === 'tank') return 'tank';
  if (classSlug === 'hunter' && spec !== 'hunter-survival') return 'ranged';
  if (CASTER_CLASSES.has(classSlug) || CASTER_SPECS.has(spec)) return 'caster';
  return 'melee';
}

function racialText(race: RaceRow): string {
  return race.forever_changes.map((change) => change.text).join(' ');
}

function trimEffect(raw: string): string {
  let text = raw.trim();
  if (text.startsWith('(')) {
    const close = text.indexOf(')');
    text = close === -1 ? text.slice(1) : text.slice(1, close);
  }
  const stops = [text.indexOf('. '), text.indexOf(';')].filter((i) => i !== -1);
  if (stops.length > 0) text = text.slice(0, Math.min(...stops));
  text = text.replace(/[,\s]+(and\s*)?$/i, '').replace(/[.\s]+$/, '');
  text = text
    .replace(/^,\s*/, '')
    .replace(/^and\s+/i, '')
    .replace(/^now\s+/i, '');
  return text.trim();
}

/** What each racial of `race` does, in the data's own words, by racial name. A racial the
 *  data does not describe is absent. */
export function racialEffects(race: RaceRow): Map<string, string> {
  const text = racialText(race);
  const names = RACIAL_NAMES[race.slug] ?? [];
  // A racial whose name sits inside a longer one ("Regeneration" in "Rapid Regeneration") is
  // looked for outside the longer name's span.
  const spans = names.flatMap((name) => {
    const at = text.indexOf(name);
    return at === -1 ? [] : [{ name, start: at, end: at + name.length }];
  });
  const found = names
    .map((name) => {
      let at = text.indexOf(name);
      while (at !== -1) {
        const covered = spans.some(
          (span) => span.name !== name && span.name.includes(name) && at >= span.start && at < span.end,
        );
        if (!covered) break;
        at = text.indexOf(name, at + 1);
      }
      return { name, at };
    })
    .filter((entry) => entry.at !== -1)
    .sort((a, b) => a.at - b.at);
  const effects = new Map<string, string>();
  found.forEach((entry, index) => {
    const start = entry.at + entry.name.length;
    const end = found[index + 1]?.at ?? text.length;
    const effect = trimEffect(text.slice(start, end));
    if (effect !== '') effects.set(entry.name, effect);
  });
  return effects;
}

export interface GuideRaceRow {
  slug: string;
  name: string;
  faction: string;
  recommended: boolean;
  /** The racial that matters for this spec, absent when the data does not describe it. */
  racial: string | undefined;
  /** What it does, sentence case, no final stop. */
  effect: string | undefined;
}

function upperFirst(text: string): string {
  return text === '' ? text : text[0]!.toUpperCase() + text.slice(1);
}

export function raceTableFor(
  classSlug: string,
  recommendedRaces: readonly string[],
  kind: SpecKind,
): GuideRaceRow[] {
  const classRow = classRows.find((row) => row.slug === classSlug);
  if (classRow === undefined) return [];
  return racesForClass(classRow.id).map((race): GuideRaceRow => {
    const picks = RACIAL_PICKS[race.slug];
    const racial = picks === undefined ? undefined : (picks[kind] ?? picks.default);
    const effect = racial === undefined ? undefined : racialEffects(race).get(racial);
    return {
      slug: race.slug,
      name: race.name,
      faction: race.faction,
      recommended: recommendedRaces.includes(race.slug),
      racial: effect === undefined ? undefined : racial,
      effect: effect === undefined ? undefined : upperFirst(effect),
    };
  });
}

/** A race's display name from its slug (a BiS band names its race by slug); the slug itself
 *  when the race is unknown. */
export function raceNameFor(slug: string): string {
  return raceRows.find((race) => race.slug === slug)?.name ?? slug;
}
