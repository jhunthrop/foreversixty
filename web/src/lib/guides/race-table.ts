// web/src/lib/guides/race-table.ts
// The guide's Races table: one row per race the class can be, with the racial that matters
// for this kind of spec and what that racial does. The race list comes from the build's own
// combos, the racial effects from its race data (`races.json` `forever_changes`), both read
// off data/builds/<build>/ (build-races.ts) so the table is the same under either
// FOREVER_DATA; the one editorial input is which racial matters to which kind of spec, kept
// here as a small table that a test holds to the race data, so a renamed racial fails a test
// instead of the page.
import type { RaceRow } from '../planner/types';
import { buildRace, buildRacesForClass } from './build-races';

export type SpecKind = 'caster' | 'melee' | 'ranged' | 'healer' | 'tank';

/** Every racial the data names for a race. */
export const RACIAL_NAMES: Readonly<Record<string, readonly string[]>> = {
  human: ['Will to Survive', 'Perception', 'Ruthless', 'The Human Spirit'],
  orc: ['Blood Fury', 'Shatter Curse', 'Command', 'Hardiness'],
  dwarf: ['Stoneform', 'Find Treasure', "Artisan's Precision", 'Big Game Hunter'],
  'night-elf': ["Elune's Light", 'Shadowmeld', 'Quickness', 'Wisp Spirit'],
  undead: ['Will of the Forsaken', 'Cannibalize', 'Touch of the Grave'],
  tauren: ['War Stomp', 'Cultivation', 'Plainsrunning', 'Endurance'],
  gnome: ['Escape Artist', 'Eureka!', 'Expansive Mind', 'Engineering Specialization'],
  troll: ['Berserking', 'Rapid Regeneration', 'Beast Slaying', 'Regeneration'],
  'high-order-skyborne': ['Walk on Air', 'Read Ley Line', 'Wind Blessed', 'Elemental Insight'],
  'windshaper-skyborne': ['Walk on Air', 'Skysight', 'Wind Blessed', 'Elemental Insight'],
};

type Picks = { default: string[] } & Partial<Record<SpecKind, string[]>>;

/** The racials that matter, by kind of spec. The rule: the racial the sim credits with damage,
 *  healing or threat for that kind leads (Gnome Eureka!, Orc Blood Fury, Troll Berserking,
 *  Night Elf Elune's Light, Human Ruthless, Dwarf Artisan's Precision, Undead Touch of the Grave,
 *  Skyborne Wind Blessed), then the one that carries the kind's resource or survival (Expansive
 *  Mind, The Human Spirit, Read Ley Line, Stoneform); a utility racial (Escape Artist, Will of
 *  the Forsaken, Shatter Curse, Quickness) shows only where nothing else helps that kind more.
 *  Names are the race data's own. */
export const RACIAL_PICKS: Readonly<Record<string, Picks>> = {
  human: {
    default: ['Ruthless', 'The Human Spirit'],
    melee: ['Ruthless'],
    tank: ['Ruthless'],
    ranged: ['Ruthless'],
  },
  orc: { default: ['Blood Fury'], tank: ['Blood Fury', 'Shatter Curse'] },
  dwarf: {
    default: ["Artisan's Precision"],
    healer: ['Stoneform'],
    tank: ['Stoneform', "Artisan's Precision"],
  },
  'night-elf': { default: ["Elune's Light"], tank: ["Elune's Light", 'Quickness'] },
  undead: {
    default: ['Will of the Forsaken'],
    melee: ['Touch of the Grave'],
    ranged: ['Touch of the Grave'],
    caster: ['Touch of the Grave'],
    tank: ['Touch of the Grave', 'Will of the Forsaken'],
  },
  tauren: { default: ['Endurance'] },
  gnome: { default: ['Eureka!', 'Expansive Mind'], tank: ['Eureka!', 'Escape Artist'] },
  troll: { default: ['Berserking'] },
  'high-order-skyborne': {
    default: ['Wind Blessed'],
    caster: ['Wind Blessed', 'Read Ley Line'],
    healer: ['Wind Blessed', 'Read Ley Line'],
  },
  'windshaper-skyborne': { default: ['Wind Blessed'] },
};

/** Picks for one class of a race, ahead of its kind's: Orc Command raises only Hunter and
 *  Warlock pets' damage (the client's 20576 and 20575), so only those two classes show it. */
export const RACIAL_PICKS_BY_CLASS: Readonly<Record<string, Readonly<Record<string, string[]>>>> = {
  orc: { hunter: ['Blood Fury', 'Command'], warlock: ['Blood Fury', 'Command'] },
};

/** A racial whose client row differs by class, in that class's own tooltip words (the race data
 *  carries one sentence for the race). Gnome's Eureka! is one row per class, 1259813 Warrior,
 *  1259812 Rogue, 1259817 Mage, 1259821 Warlock and 1259823 Priest: the resource it discounts,
 *  and whether periodic damage and healing count, follow the class. A test holds each wording
 *  to the race data's own duration and cooldown. */
export const RACIAL_EFFECTS_BY_CLASS: Readonly<
  Record<string, Readonly<Record<string, Readonly<Record<string, string>>>>>
> = {
  gnome: {
    'Eureka!': {
      mage: 'Makes your damaging abilities cost 10% less Mana and deal 10% more damage for 8 seconds on a 2 minute cooldown',
      priest:
        'Makes your non-periodic damaging and healing abilities cost 10% less Mana and deal 10% more damage or healing for 8 seconds on a 2 minute cooldown',
      rogue:
        'Makes your non-periodic damaging abilities cost 10% less Energy and deal 10% more damage for 8 seconds on a 2 minute cooldown',
      warlock:
        'Makes your non-periodic damaging abilities cost 10% less Mana and deal 10% more damage for 8 seconds on a 2 minute cooldown',
      warrior:
        'Makes your non-periodic damaging abilities cost 10% less Rage and deal 10% more damage for 8 seconds on a 2 minute cooldown',
    },
  },
};

/** Each race's offensive racial cooldown, the one the sim casts on cooldown with the rest of the
 *  cooldown pass (`autocastOtherCooldowns`). */
export const RACIAL_COOLDOWNS: Readonly<Record<string, string>> = {
  orc: 'Blood Fury',
  'night-elf': "Elune's Light",
  gnome: 'Eureka!',
  troll: 'Berserking',
};

export interface RacialCooldown {
  /** The race's display name. */
  race: string;
  racial: string;
}

/** The racial cooldowns of the races `classSlug` can be, in the build's race order. */
export function racialCooldownsFor(build: string, classSlug: string): RacialCooldown[] {
  return buildRacesForClass(build, classSlug).flatMap((race) => {
    const racial = RACIAL_COOLDOWNS[race.slug];
    return racial === undefined ? [] : [{ race: race.name, racial }];
  });
}

function picksFor(slug: string, classSlug: string, kind: SpecKind): string[] {
  const byClass = RACIAL_PICKS_BY_CLASS[slug]?.[classSlug];
  if (byClass !== undefined) return byClass;
  const picks = RACIAL_PICKS[slug];
  return picks === undefined ? [] : (picks[kind] ?? picks.default);
}

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
  // The data's history of a racial ("but no longer ...", ", in place of Earth Fury's ...") is not
  // what it does.
  text = text.replace(/\s+but no longer.*$/i, '').replace(/,\s+in place of\s.*$/i, '');
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

export interface GuideRacial {
  name: string;
  /** What it does, sentence case, no final stop. */
  effect: string;
}

export interface GuideRaceRow {
  slug: string;
  name: string;
  faction: string;
  recommended: boolean;
  /** The racials that matter for this spec; the ones the data does not describe are left out. */
  racials: GuideRacial[];
}

function upperFirst(text: string): string {
  return text === '' ? text : text[0]!.toUpperCase() + text.slice(1);
}

export function raceTableFor(
  build: string,
  classSlug: string,
  recommendedRaces: readonly string[],
  kind: SpecKind,
): GuideRaceRow[] {
  return buildRacesForClass(build, classSlug).map((race): GuideRaceRow => {
    const wanted = picksFor(race.slug, classSlug, kind);
    const effects = racialEffects(race);
    // A racial the editorial table does not know (or the build's text does not describe) falls
    // back to the race's first racial the data does describe, with its own sentence; a race
    // whose data describes none shows no racial at all.
    const known = wanted.filter((name) => effects.has(name));
    const names = known.length > 0 ? known : effects.size > 0 ? [[...effects.keys()][0]!] : [];
    return {
      slug: race.slug,
      name: race.name,
      faction: race.faction,
      recommended: recommendedRaces.includes(race.slug),
      racials: names.flatMap((name) => {
        const effect = effects.get(name);
        if (effect === undefined) return [];
        const forClass = RACIAL_EFFECTS_BY_CLASS[race.slug]?.[name]?.[classSlug];
        return [{ name, effect: forClass ?? upperFirst(effect) }];
      }),
    };
  });
}

/** The recommended races as the one line under the table: "Gnome is the Alliance pick: Eureka!
 *  and Expansive Mind." ("a" when a faction has more than one). */
export function recommendedRaceLines(
  rows: readonly GuideRaceRow[],
): { slug: string; faction: string; names: string[] }[] {
  return rows
    .filter((row) => row.recommended && row.racials.length > 0)
    .map((row) => ({
      slug: row.slug,
      faction: row.faction,
      names: row.racials.map((racial) => racial.name),
    }));
}

/** A race's display name from its slug (a BiS band names its race by slug); the slug itself
 *  when the build has no such race. */
export function raceNameFor(build: string, slug: string): string {
  return buildRace(build, slug)?.name ?? slug;
}
