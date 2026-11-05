// web/src/lib/guild/week-one.ts
// The Before raids open card (design/specs/2026-11-05-guild-week-one.md §3.1): the guild's level
// spread, the raids-open line, who has which profession, and the viewer's one next step. Pure, so
// every branch is pinned without a page; GuildWeekOneCard.svelte draws it.
import phases from '../../data/phases.json';
import { levelingGuideHref, levelingSpecKey, specNameOf } from '../leveling/leveling-plan';
import { SPECS } from '../sim/specs';
import type { GuildRosterEntry, NextRaidView } from './api';
import { isPrimaryProfession, professionName } from './professions';
import { countdown, whenPhrase } from './raid-copy';
import { dateLabel } from './raid-time';
import { RAID_LEVEL } from './readiness-view';
import { sameSpecName } from './spec-name';

/** Characters under this are bank alts and fresh rolls, not raiders: the spread leaves them out. */
const LOWEST_COUNTED_LEVEL = 10;
const FIFTIES = 50;
const FORTIES = 40;
/** A profession lists this many people, then `+{n}`. */
const PROFESSION_NAMES_SHOWN = 3;
const UTC = 'UTC';

const RAIDS_ONE = phases.find((phase) => phase.name === 'raids-1')?.start;
if (RAIDS_ONE === undefined) throw new Error('phases.json has no raids-1 phase');

/** When the tier's raids open (`raids-1` in phases.json), the one source of that date. */
export const RAIDS_OPEN_AT = new Date(RAIDS_ONE);

export const weekOneCopy = {
  titleBefore: 'Before raids open',
  titleAfter: 'Your first raid',
  raidsOpen: 'Raids are open. Your first logged night starts the progression board.',
  noLevel: (n: number): string => `${n} with no level on file yet.`,
  empty: "Nobody's level is on file yet. Levels show once members sync the addon or import from Battle.net.",
  error: "The guild's levels did not load.",
  professions: 'Professions',
  invite: 'Copy the invite link',
} as const;

/** The tier's raids open on this date, written the way every player knows it: "Wed 9 Dec". */
export function raidsOpenDay(locale?: string): string {
  return dateLabel(RAIDS_OPEN_AT, UTC, locale);
}

export function raidsAreOpen(now: Date): boolean {
  return now.getTime() >= RAIDS_OPEN_AT.getTime();
}

/** The card's title: before the raids open, then the first raid it is still waiting for. */
export function weekOneTitle(now: Date): string {
  return raidsAreOpen(now) ? weekOneCopy.titleAfter : weekOneCopy.titleBefore;
}

/** `Barrow Deeps, Hyjal Summit and Onyxia's Lair open Wed 9 Dec, in 35 days.` Or the open line. */
export function raidsOpenLine(now: Date, locale?: string): string {
  if (raidsAreOpen(now)) return weekOneCopy.raidsOpen;
  const left = countdown(RAIDS_OPEN_AT, RAIDS_OPEN_AT, now);
  return `Barrow Deeps, Hyjal Summit and Onyxia's Lair open ${raidsOpenDay(locale)}, ${left}.`;
}

/** The four tiles and the characters with no level on file. */
export interface LevelSpread {
  at60: number;
  fifties: number;
  forties: number;
  under40: number;
  noLevel: number;
}

/** True when the guild has no level on file for anyone counted: the tiles give way to one line. */
export function spreadIsEmpty(spread: LevelSpread): boolean {
  return spread.at60 + spread.fifties + spread.forties + spread.under40 === 0;
}

/** Every roster character by level: an in-game name no account has claimed is not on the guild's
 *  roster yet, and a character under 10 is not a raider. */
export function levelSpread(entries: readonly GuildRosterEntry[]): LevelSpread {
  const spread: LevelSpread = { at60: 0, fifties: 0, forties: 0, under40: 0, noLevel: 0 };
  return entries.reduce((total, entry) => {
    if (entry.state === 'no_account') return total;
    const level = entry.level ?? null;
    if (level === null) return { ...total, noLevel: total.noLevel + 1 };
    if (level < LOWEST_COUNTED_LEVEL) return total;
    if (level >= RAID_LEVEL) return { ...total, at60: total.at60 + 1 };
    if (level >= FIFTIES) return { ...total, fifties: total.fifties + 1 };
    if (level >= FORTIES) return { ...total, forties: total.forties + 1 };
    return { ...total, under40: total.under40 + 1 };
  }, spread);
}

/** The part of a roster row the professions line reads. */
export interface ProfessionHolder {
  name: string;
  level?: number | null;
  professions?: readonly string[];
}

/** `Alchemy: Rootbound, Cindergale`, one fact per primary profession, alphabetical, from the counted
 *  rows that carry any. A profession with more than three people reads its first three then `+{n}`.
 *  The export carries slugs only, so names sort alphabetically until it sends the skill. */
export function professionFacts(rows: readonly ProfessionHolder[]): string[] {
  const holders = new Map<string, string[]>();
  for (const row of rows) {
    if ((row.level ?? 0) < LOWEST_COUNTED_LEVEL) continue;
    for (const slug of row.professions ?? []) {
      if (!isPrimaryProfession(slug)) continue;
      const name = professionName(slug);
      holders.set(name, [...(holders.get(name) ?? []), row.name]);
    }
  }
  return [...holders.keys()]
    .sort((a, b) => a.localeCompare(b))
    .map((profession) => {
      const names = [...(holders.get(profession) ?? [])].sort((a, b) => a.localeCompare(b));
      const shown = names.slice(0, PROFESSION_NAMES_SHOWN).join(', ');
      const more = names.length - PROFESSION_NAMES_SHOWN;
      return `${profession}: ${shown}${more > 0 ? ` +${more}` : ''}`;
    });
}

/** What the card's one primary button does. A link goes there; `approve` and `plan` open a tab the
 *  page holds. */
export type WeekOneAction =
  | { kind: 'link'; href: string; label: string }
  | { kind: 'approve'; label: string }
  | { kind: 'plan'; label: string };

export interface WeekOneStep {
  /** The current character, drawn in its class colour ahead of `text`; null for a sentence not about it. */
  name: string | null;
  text: string;
  action: WeekOneAction;
  /** Officers only: the invite link, as an ActionLink. */
  invite: boolean;
}

/** The viewer's current character as the card needs it. */
export interface WeekOneCharacter {
  name: string;
  class: string;
  /** Null when neither the export nor the addon roster carries one (no export, or gear not shared). */
  level: number | null;
  /** The spec as the logs or the site know it; null at zero logs with none picked. */
  spec: string | null;
}

export interface WeekOneInput {
  officer: boolean;
  waiting: number;
  nextRaid: NextRaidView | null;
  /** The next raid's page, null with none. */
  raidHref: string | null;
  character: WeekOneCharacter | null;
  now: Date;
  locale: string;
  viewerTz?: string;
}

/** The catalogue key of a character's spec ("hunter-beast-mastery"), from a slug or any spelling
 *  of its name. Null when the character has no spec yet or it is not this class's. */
export function characterSpecKey(classSlug: string, spec: string | null): string | null {
  if (spec === null || spec === '') return null;
  const found = SPECS.find(
    (entry) => entry.class_slug === classSlug && (entry.spec === spec || sameSpecName(entry.name, spec)),
  );
  return found?.spec ?? null;
}

function raidWhen(raid: NextRaidView, input: WeekOneInput): string {
  const viewerTz = input.viewerTz ?? new Intl.DateTimeFormat().resolvedOptions().timeZone;
  return whenPhrase(
    new Date(raid.starts_at),
    input.now,
    input.locale,
    raid.guild_timezone ?? viewerTz,
    viewerTz,
  );
}

/** Who is signed up for a raid: the In and the Late. */
function signedUpCount(raid: NextRaidView): number {
  return raid.counts.in + raid.counts.late;
}

function officerStep(input: WeekOneInput): WeekOneStep {
  const { nextRaid, raidHref, waiting } = input;
  if (waiting > 0) {
    return {
      name: null,
      text: `${waiting} waiting for approval. Approve them so they can sign up.`,
      action: { kind: 'approve', label: `Approve ${waiting}` },
      invite: true,
    };
  }
  if (nextRaid === null || raidHref === null) {
    return {
      name: null,
      text: 'Nothing on the calendar yet. Post a dungeon run or a guild event.',
      action: { kind: 'plan', label: 'Plan an event' },
      invite: true,
    };
  }
  return {
    name: null,
    text: `${nextRaid.title} ${raidWhen(nextRaid, input)}, ${signedUpCount(nextRaid)} signed up.`,
    action: { kind: 'link', href: raidHref, label: 'Open the sign-ups' },
    invite: true,
  };
}

/** The raider's own character: how far to 60 and how to get there, or at 60 what to check. */
function characterStep(character: WeekOneCharacter, level: number, locale: string): WeekOneStep {
  const specKey = characterSpecKey(character.class, character.spec);
  if (level >= RAID_LEVEL) {
    const slug = SPECS.find((entry) => entry.spec === specKey)?.spec_slug;
    return {
      name: character.name,
      text: ` is level ${RAID_LEVEL}. Check your gear before ${raidsOpenDay(locale)}.`,
      action: {
        kind: 'link',
        href: slug === undefined ? '/bis' : `/bis/${character.class}/${slug}`,
        label: 'Open your best in slot',
      },
      invite: false,
    };
  }
  // At zero logs a character has no spec yet: the class guide's own leveling spec stands in.
  const levelingKey = specKey ?? levelingSpecKey({ class: character.class }) ?? '';
  return {
    name: character.name,
    text: ` is level ${level}, ${RAID_LEVEL - level} levels to ${RAID_LEVEL}.`,
    action: {
      kind: 'link',
      href: levelingGuideHref(levelingKey) ?? '/guides',
      label: `Level as ${specNameOf(levelingKey) ?? character.class}`,
    },
    invite: false,
  };
}

function raiderStep(input: WeekOneInput): WeekOneStep {
  const { character, nextRaid, raidHref } = input;
  if (character === null || character.level === null) {
    return {
      name: null,
      text: 'Sync the addon or import from Battle.net so the guild sees your level and gear.',
      action: { kind: 'link', href: '/setup', label: 'Import your character' },
      invite: false,
    };
  }
  if (nextRaid !== null && raidHref !== null && nextRaid.own === null) {
    return {
      name: null,
      text: `${nextRaid.title} ${raidWhen(nextRaid, input)}. You have not signed up yet.`,
      action: { kind: 'link', href: raidHref, label: 'Sign up' },
      invite: false,
    };
  }
  return characterStep(character, character.level, input.locale);
}

/** The viewer's next step: the first branch that applies, one primary button (spec §3.1). */
export function weekOneStep(input: WeekOneInput): WeekOneStep {
  return input.officer ? officerStep(input) : raiderStep(input);
}
