// web/src/lib/character/page-view.ts
// The character page's words and rows (design/specs/2026-11-06-character.md, rulings C2 to C11), as
// plain functions over the GET /v1/characters/... answer, so the page island only lays them out:
// the head's fact line, the upgrade total in the band's unit, the owner's one Next line, the five
// best fights, the three builds seen and the not-found copy.
import type { MeCharacter } from '../account/api';
import {
  characterHeaderCopy,
  profileHeader,
  resolvedProfileSpec,
  seenBuildSpec,
  specClassLabel,
} from '../character-profile';
import type { CharacterPath } from '../characters';
import { rulesetLabel } from '../characters';
import { isTie } from '../bis/doll-view';
import type { PaperdollSlot } from '../gear/paperdoll';
import { homeHrefs } from '../home/home-view';
import { PREMIUM_LINKS } from '../billing/premium-links';
import { percentileToken } from '../report/format';
import { executionLabel } from '../sim/execution';
import { signedFigure } from '../ui/primitives';
import type { BuildSeen, CharacterFight, CharacterPage } from '../rankings/api';
import { pointsEarnedAt } from '../leveling/leveling-plan';

/** C6: the fights shown before the link to all of them. */
export const FIGHTS_SHOWN = 5;
/** C7: the builds shown. */
export const BUILDS_SHOWN = 3;

export const characterPageCopy = {
  /** The band's answer after its figure: "+128.9 HPS from upgrades". */
  answerUnit: (unit: string): string => `${unit} from upgrades`,
  /** The head's sim link: the character on screen, in the simulator. */
  simThis: 'Sim this character',
  upgrades: 'Upgrades',
  allUpgrades: 'All upgrades',
  /** "8 more, +23.5 HPS": the upgrades past the frame's centre and what they add. */
  moreUpgrades: (count: number, gain: string): string => `${count} more, ${gain}`,
  openBis: 'Open in best in slot',
  gear: 'Gear',
  gearAction: 'Best in slot',
  score: 'Forever Score',
  scoreFigure: 'Score',
  rating: 'Rating',
  ratingValue: (overall: number): string => `${overall} of 100`,
  /** The score trend's label: "Since 20 Oct". */
  since: (date: string): string => `Since ${dayMonth(`${date}T00:00:00Z`)}`,
  noLogsOwner: 'No logs yet. Upload a log to get a rating.',
  noLogs: 'No logs yet.',
  noScoreYet: 'No Forever Score yet.',
  best: 'Best per encounter',
  /** A best fight's execution against the sim, in its sub: "Execution 90%". */
  execution: (score: number): string => `Execution ${executionLabel(score)}`,
  /** A best fight's change against the kill of that boss before it: "+82 vs 27 Oct". */
  sinceKill: (change: string, date: string): string => `${change} vs ${date}`,
  allFights: (count: number): string => `All fights (${count})`,
  builds: 'Builds seen',
  /** A build seen row's sub: the spec and the date it was first seen ("Holy · 3 Nov"). */
  buildSub: (spec: string, date: string): string => (spec === '' ? date : `${spec} · ${date}`),
  history: 'History',
  /** The Premium entries under the best fights (gold link rows that open in place). */
  coach: 'Coach',
  /** The History disclosure's meta: the fights behind its charts. */
  fightCount: (count: number): string => `${count} ${count === 1 ? 'fight' : 'fights'}`,
  historyLockTitle: 'Score over time and parse per boss',
  historyLockSub: 'Premium keeps every rated fight.',
  seePlans: 'See plans',
  /** The paperdoll's title bar meta: "Level 60". */
  levelMeta: (level: number): string => `Level ${level}`,
  failed: 'We could not load this character.',
  notFound: (name: string, region: string, ruleset: string): string =>
    `No character named ${name} on ${region.toUpperCase()} ${rulesetLabel(ruleset)}.`,
  elsewhere: 'Same name elsewhere',
  lookup: 'Look up a character',
  lookupName: 'Name',
  lookupSubmit: 'Look up',
  add: 'Add this character',
  signIn: 'Sign in to see your characters',
  signInSub: 'Battle.net lists every character on your account.',
  addonSub: 'Export with the addon and the companion to add this character.',
  findTitle: 'Find a character',
} as const;

export const characterHrefs = {
  logs: '/logs',
  login: '/login',
  setup: '/setup',
  plans: PREMIUM_LINKS.history,
  lookup: (name: string): string => `/characters?q=${encodeURIComponent(name)}#lookup`,
  character: (region: string, ruleset: string, name: string): string =>
    `/character/${region}/${ruleset}/${encodeURIComponent(name.toLowerCase().replace(/\s+/g, '-'))}`,
  encounter: (slug: string): string => `/rankings/${slug}`,
} as const;

/** The page's spec: the one profile rule (talents, builds seen, fights), else the spec the gear is
 *  priced for (`gear_spec`, "hunter-beast-mastery" is Beast Mastery), so a leveler with no logs
 *  still reads as a spec. Null when nothing names one. */
export function pageSpec(page: CharacterPage, own: MeCharacter | null): string | null {
  const resolved = resolvedProfileSpec(page, own);
  if (resolved !== null) return resolved;
  const key = page.gear_spec ?? '';
  const words = key
    .slice(key.indexOf('-') + 1)
    .split('-')
    .filter(Boolean);
  if (key === '' || words.length === 0) return null;
  return words.map((word) => word.charAt(0).toUpperCase() + word.slice(1)).join(' ');
}

/** One piece of the head's fact line: words, or a link. */
export interface HeadFact {
  key: string;
  text: string;
  href?: string;
}

/** C3: level and race, spec and class, the talent split (to the planner), the guild, the realm,
 *  the ruleset and region; the owner's sync line is added by the caller. Unknown pieces are left out. */
export function headFacts(page: CharacterPage, path: CharacterPath, own: MeCharacter | null): HeadFact[] {
  const header = profileHeader(page, path, own ?? undefined);
  const facts: HeadFact[] = [];
  if (header.levelRace !== null) facts.push({ key: 'level', text: header.levelRace });
  for (const piece of header.pieces) {
    if (piece.key === 'spec') {
      facts.push({ key: 'spec', text: specClassLabel(page.character.class, pageSpec(page, own)) });
      if (piece.split !== undefined) facts.push({ key: 'split', text: piece.split, href: plannerHref(own) });
    } else facts.push({ key: piece.key, text: piece.text, href: piece.href });
  }
  facts.push({ key: 'chip', text: characterHeaderCopy.chip(path.ruleset, path.region) });
  return facts;
}

/** The planner: the owner's own character, else the planner's landing. */
export function plannerHref(own: MeCharacter | null): string {
  return own === null ? '/planner' : homeHrefs.planner(own.key);
}

/** C2: the priced gains summed (ties and unpriced slots add nothing, never a fabricated figure). */
export function upgradeTotal(slots: readonly PaperdollSlot[]): number {
  return slots.reduce(
    (sum, slot) =>
      slot.upgrade?.kind === 'gain' && slot.upgrade.value !== null && !isTie(slot.upgrade.value)
        ? sum + slot.upgrade.value
        : sum,
    0,
  );
}

/** The slot with the largest priced gain, or null. */
export function topUpgrade(slots: readonly PaperdollSlot[]): PaperdollSlot | null {
  let top: PaperdollSlot | null = null;
  let best = 0;
  for (const slot of slots) {
    const up = slot.upgrade;
    if (up?.kind === 'gain' && up.value !== null && !isTie(up.value) && up.value > best) {
      best = up.value;
      top = slot;
    }
  }
  return top;
}

/** One fact waiting on the owner, and where to act on it. */
export interface WaitingFact {
  key: 'talents' | 'upgrade';
  /** The words before the fact that a phone drops ("Upgrade: "). */
  lead?: string;
  text: string;
  href: string | null;
  /** The upgrade's item, drawn as its icon in its quality ring. */
  item?: { icon: string | null; quality: number };
}

/** A source's boss without its zone ("Stratholme: Baron Rivendare" reads "Baron Rivendare"), so the
 *  upgrade line fits one line on a phone. */
export function bossOf(source: string): string {
  const at = source.lastIndexOf(': ');
  return at === -1 ? source : source.slice(at + 2);
}

/** C11: what waits on the owner, as plain facts (owner 2026-11-09: no "Next:" nudge): unspent
 *  talent points, then the top upgrade. Empty with neither. */
export function characterWaiting(input: {
  own: MeCharacter;
  top: { name: string; source: string; icon: string | null; quality: number } | null;
  bisHref: string | null;
}): WaitingFact[] {
  const { own, top, bisHref } = input;
  const points = own.build?.talents?.points;
  const level = own.level;
  const toSpend =
    points === undefined || level === undefined
      ? 0
      : Math.max(0, pointsEarnedAt(level) - points.reduce((sum, count) => sum + count, 0));
  const talents: WaitingFact[] =
    toSpend > 0
      ? [
          {
            key: 'talents',
            text: `${toSpend} talent ${toSpend === 1 ? 'point' : 'points'} unspent`,
            href: homeHrefs.planner(own.key),
          },
        ]
      : [];
  const upgrade: WaitingFact[] =
    top === null
      ? []
      : [
          {
            key: 'upgrade',
            lead: 'Upgrade: ',
            text: `${top.name}, ${bossOf(top.source)}`,
            href: bisHref,
            item: { icon: top.icon, quality: top.quality },
          },
        ];
  return [...talents, ...upgrade];
}

/** `3 Nov`, or `3 Nov 2026` with the year. */
export function dayMonth(iso: string, year = false): string {
  return new Date(iso).toLocaleDateString('en-GB', {
    day: 'numeric',
    month: 'short',
    ...(year ? { year: 'numeric' } : {}),
    timeZone: 'UTC',
  });
}

export interface FightRowView {
  key: string;
  boss: string;
  sub: string;
  percentile: string | null;
  color: string | null;
  value: string;
  /** The fight's own rate word: HPS for a heal, DPS otherwise. */
  unit: string;
  encounter: string;
  /** The boss's parses kill by kill, oldest first; empty with fewer than two (h-audit 2026-11-09). */
  trend: number[];
  /** "+82 vs 27 Oct" on the parse line; '' with no kill before. */
  since: string;
}

/** The same boss, difficulty and metric as a best fight. */
const sameBoss = (a: CharacterFight, b: CharacterFight): boolean =>
  a.encounter_id === b.encounter_id && a.difficulty === b.difficulty && a.metric === b.metric;

/** A best fight's boss kill by kill from the fights on the page, oldest first (a wipe is left out). */
function bossKills(row: CharacterFight, history: readonly CharacterFight[]): CharacterFight[] {
  return history
    .filter((fight) => sameBoss(fight, row) && fight.kill !== false)
    .sort((a, b) => a.fought_at.localeCompare(b.fought_at));
}

/** "+82 since 27 Oct": the best fight against the kill of that boss before it; '' with none. */
function sinceLast(row: CharacterFight, kills: readonly CharacterFight[]): string {
  const before = kills.filter((fight) => fight.fought_at < row.fought_at).at(-1);
  return before === undefined
    ? ''
    : characterPageCopy.sinceKill(signedFigure(row.value - before.value, 0), dayMonth(before.fought_at));
}

/** C6: the best fight per encounter, five at most, each with its boss's trend from `history`. */
export function bestRows(
  best: readonly CharacterFight[],
  history: readonly CharacterFight[] = [],
  pageSpecName: string | null = null,
): FightRowView[] {
  return best.slice(0, FIGHTS_SHOWN).map((row) => ({
    key: `${row.encounter_id}-${row.difficulty}-${row.metric}`,
    boss: row.encounter,
    sub: [
      dayMonth(row.fought_at),
      // The page's own spec goes without saying; another spec's fight names it.
      row.spec === pageSpecName ? '' : row.spec,
      typeof row.execution_score === 'number' ? characterPageCopy.execution(row.execution_score) : '',
    ]
      .filter(Boolean)
      .join(' · '),
    percentile: row.percentile === undefined ? null : String(Math.round(row.percentile)),
    color: row.percentile === undefined ? null : percentileToken(row.percentile),
    value: Math.round(row.value).toLocaleString('en-US'),
    unit: row.metric === 'hps' ? 'HPS' : 'DPS',
    encounter: row.encounter,
    trend: trendOf(bossKills(row, history)),
    since: sinceLast(row, bossKills(row, history)),
  }));
}

function trendOf(kills: readonly CharacterFight[]): number[] {
  const parses = kills.flatMap((fight) => (fight.percentile === undefined ? [] : [fight.percentile]));
  return parses.length > 1 ? parses : [];
}

export interface BuildRowView {
  key: string;
  split: string;
  sub: string;
}

/** C7: the three newest builds seen at pull. */
export function buildRows(className: string | undefined, builds: readonly BuildSeen[]): BuildRowView[] {
  return [...builds]
    .sort((a, b) => b.first_seen.localeCompare(a.first_seen))
    .slice(0, BUILDS_SHOWN)
    .map((build) => ({
      key: `${build.talent_split}-${build.first_seen}`,
      split: build.talent_split,
      sub: characterPageCopy.buildSub(seenBuildSpec(className, build), dayMonth(build.first_seen)),
    }));
}
