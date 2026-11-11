// web/src/lib/guides/race-table.ts
// The guide's Races table: one row per race the class can be, with the racial that matters
// for this kind of spec and what that racial does. The race list comes from the build's own
// combos, the racial effects from its race data (`races.json` `forever_changes`), both read
// off data/builds/<build>/ (build-races.ts) so the table is the same under either
// FOREVER_DATA; the one editorial input is which racial matters to which kind of spec, kept
// here as a small table that a test holds to the race data, so a renamed racial fails a test
// instead of the page.
import type { RaceRow } from '../planner/types';
import { buildIconExists } from '../data/build-files';
import { buildRace, buildRacesForClass } from './build-races';
import { readBuildJson } from './build-files';
import { upperFirst } from '../text';

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
  'high-order-skyborne': ['Walk on Air', 'Read Leyline', 'Wind Blessed', 'Elemental Insight'],
  'windshaper-skyborne': ['Walk on Air', 'Skysight', 'Wind Blessed', 'Elemental Insight'],
};

type Picks = { default: string[] } & Partial<Record<SpecKind, string[]>>;

/** The racials that matter, by kind of spec. The rule: the racial the sim credits with damage,
 *  healing or threat for that kind leads (Gnome Eureka!, Orc Blood Fury, Troll Berserking,
 *  Night Elf Elune's Light, Human Ruthless, Dwarf Artisan's Precision, Undead Touch of the Grave,
 *  Skyborne Wind Blessed), then the one that carries the kind's resource or survival (Expansive
 *  Mind, The Human Spirit, Read Leyline, Stoneform); a utility racial (Escape Artist, Will of
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
    caster: ['Wind Blessed', 'Read Leyline'],
    healer: ['Wind Blessed', 'Read Leyline'],
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

/** The stat words a racial's tooltip spells out, in the short form a player says. */
const STAT_SHORT: readonly (readonly [RegExp, string])[] = [
  [/^critical strike chance(?: with all spells and attacks)?$/i, 'crit'],
  [/^(?:your )?chance to hit with spells and attacks$|^hit chance with spells and attacks$/i, 'hit'],
  [/^attack power, ranged attack power and spell power$/i, 'attack and spell power'],
  [/^casting and attack speed$/i, 'attack and casting speed'],
  [/^the damage dealt by hunter and warlock pets$/i, 'pet damage'],
  [/^health$/i, 'health'],
  [/^mana, rage and energy$/i, 'mana, rage and energy'],
  [/^spirit$/i, 'spirit'],
];

/** Eureka!'s sentence: "Makes your [non-periodic] damaging [and healing] abilities cost 10% less
 *  Rage and deal 10% more damage [or healing]". */
const CHEAPER_AND_HARDER =
  /^Makes your (non-periodic )?damaging(?: and healing)? abilities cost (\d+%) less (\w+) and deal (\d+%) more (damage(?: or healing)?)$/i;
/** Touch of the Grave's: "… drain Health from the target, up to 5% of your maximum Health, at most
 *  once every 15 seconds". */
const DRAIN =
  /^Gives .+ a chance to drain Health from the target, up to (\d+%) of your maximum Health, at most once every (\d+) seconds$/i;

/** A racial's tooltip sentence as a player says it: "Grants 10% critical strike chance for 15
 *  seconds on a 3 minute cooldown" is "10% crit, 15 s / 3 min", "Increases hit chance with spells
 *  and attacks by 1%" is "1% hit" (h-audit 2026-11-09: the tooltip pasted beside the name was
 *  slop). A sentence in any other words stays as it is. */
export function shortEffect(text: string): string {
  const drain = DRAIN.exec(text);
  if (drain !== null) return `drain up to ${drain[1]} health, once per ${drain[2]} s`;
  const timing = /\s+for (\d+) seconds(?: on a (\d+) minute cooldown)?$/i.exec(text);
  const body = timing === null ? text : text.slice(0, timing.index);
  const when =
    timing === null ? '' : `, ${timing[1]} s${timing[2] === undefined ? '' : ` / ${timing[2]} min`}`;
  const eureka = CHEAPER_AND_HARDER.exec(body);
  if (eureka !== null) {
    const [, periodic, cost, resource, more, what] = eureka;
    return `${cost} less ${resource}, ${more} more ${periodic ?? ''}${what!.toLowerCase()}${when}`;
  }
  const grants = /^Grants (\d+%) (.+)$/i.exec(body);
  const increases = /^Increases (.+) by (\d+%)$/i.exec(body);
  const [pct, stat] =
    grants !== null ? [grants[1]!, grants[2]!] : increases !== null ? [increases[2]!, increases[1]!] : [];
  const short = stat === undefined ? undefined : STAT_SHORT.find(([pattern]) => pattern.test(stat))?.[1];
  // Already short ("50% health and mana regeneration"): only its timing is said the short way.
  if (pct === undefined && timing !== null && !/^(?:Grants|Increases|Makes|Gives)\b/i.test(body))
    return `${body}${when}`;
  if (pct === undefined || short === undefined) return text;
  return `${pct} ${short}${when}`;
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

/** Where `name` first appears in `text` at or after `from`, read the way `racialKey` reads a
 *  name: past its first letter, a space may stand between two letters and the case may differ
 *  ("Read Leyline" finds "Read Ley Line"). The first letter keeps its case, so a racial named
 *  "Regeneration" never matches the word "regeneration" in another racial's effect. */
function findRacialName(
  text: string,
  name: string,
  from: number,
): { at: number; length: number } | undefined {
  const escape = (ch: string) => ch.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
  const [first = '', ...rest] = [...name.replace(/\s+/g, '')];
  const tail = rest.map((ch) =>
    ch.toLowerCase() === ch.toUpperCase()
      ? `\\s?${escape(ch)}`
      : `\\s?[${ch.toLowerCase()}${ch.toUpperCase()}]`,
  );
  const match = new RegExp(escape(first) + tail.join(''), 'g');
  match.lastIndex = from;
  const hit = match.exec(text);
  return hit === null ? undefined : { at: hit.index, length: hit[0].length };
}

/** What each racial of `race` does, in the data's own words, by racial name. A racial the
 *  data does not describe is absent. */
export function racialEffects(race: RaceRow): Map<string, string> {
  const text = racialText(race);
  const names = RACIAL_NAMES[race.slug] ?? [];
  // A racial whose name sits inside a longer one ("Regeneration" in "Rapid Regeneration") is
  // looked for outside the longer name's span.
  const spans = names.flatMap((name) => {
    const hit = findRacialName(text, name, 0);
    return hit === undefined ? [] : [{ name, start: hit.at, end: hit.at + hit.length }];
  });
  const found = names
    .map((name) => {
      let hit = findRacialName(text, name, 0);
      while (hit !== undefined) {
        const at = hit.at;
        const covered = spans.some(
          (span) => span.name !== name && span.name.includes(name) && at >= span.start && at < span.end,
        );
        if (!covered) break;
        hit = findRacialName(text, name, at + 1);
      }
      return { name, hit };
    })
    .flatMap(({ name, hit }) => (hit === undefined ? [] : [{ name, at: hit.at, length: hit.length }]))
    .sort((a, b) => a.at - b.at);
  const effects = new Map<string, string>();
  found.forEach((entry, index) => {
    const start = entry.at + entry.length;
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
  /** The racial's own client icon (`racialicons.json`), absent when the build ships none. */
  icon?: string;
}

/** A racial name as the client and the race data both spell it ("Read Ley Line", "Read Leyline"). */
function racialKey(name: string): string {
  return name.toLowerCase().replace(/[^a-z0-9]+/g, '');
}

/** The build's racial icons (`data/builds/<build>/racialicons.json`), by `racialKey`. */
export function racialIconsFor(build: string): Map<string, string> {
  const file = readBuildJson<Record<string, string>>(build, 'racialicons.json') ?? {};
  return new Map(
    Object.entries(file)
      .filter(([, icon]) => buildIconExists(build, icon))
      .map(([name, icon]) => [racialKey(name), icon]),
  );
}

export interface GuideRaceRow {
  slug: string;
  name: string;
  faction: string;
  recommended: boolean;
  /** The racials that matter for this spec; the ones the data does not describe are left out. */
  racials: GuideRacial[];
}

export function raceTableFor(
  build: string,
  classSlug: string,
  recommendedRaces: readonly string[],
  kind: SpecKind,
): GuideRaceRow[] {
  const icons = racialIconsFor(build);
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
        const icon = icons.get(racialKey(name));
        return [
          {
            name,
            effect: upperFirst(shortEffect(forClass ?? upperFirst(effect))),
            ...(icon === undefined ? {} : { icon }),
          },
        ];
      }),
    };
  });
}

/** The recommended races as the one line under the table: "Gnome is the Alliance pick: Eureka!
 *  and Expansive Mind." ("a" when a faction has more than one). A race the guide's own Races
 *  `prose` already names is left out: the prose owns the why, and the table already tags it. */
export function recommendedRaceLines(
  rows: readonly GuideRaceRow[],
  prose = '',
): { slug: string; faction: string; names: string[] }[] {
  return rows
    .filter((row) => row.recommended && row.racials.length > 0 && !proseNamesRace(prose, row.name))
    .map((row) => ({
      slug: row.slug,
      faction: row.faction,
      names: row.racials.map((racial) => racial.name),
    }));
}

/** Whether the guide's own Races prose already names `race`: the prose then says why it is
 *  the pick, and the line under the table would only say it again (one job per page). */
export function proseNamesRace(prose: string, race: string): boolean {
  const escaped = race.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
  return new RegExp(`\\b${escaped}\\b`).test(prose);
}

/** A race's display name from its slug (a BiS band names its race by slug); the slug itself
 *  when the build has no such race. */
export function raceNameFor(build: string, slug: string): string {
  return buildRace(build, slug)?.name ?? slug;
}
