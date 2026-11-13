// web/src/lib/tiers/tier-copy.ts
// Every string the tier list shows (design/specs/2026-11-06-tiers.md section 6, over the rules of
// 2026-10-09-tier-list.md). The
// strings describe the rules; a figure that moves nightly is passed in from the published
// files, never written here (the owner's "prose never quotes nightly numbers" rule).
import { answerCopy } from '../account/answer-copy';
import { SPEC_BY_KEY } from '../sim/specs';
import type { BisRole, Faction } from '../bis/types';
import { TIE_MARGIN_PERCENT, TIER_BANDS, type TierLetter } from './tier-rules';
import type { TierRow } from './tier-list';

import { grouped, oneDecimal } from '../format/number';
import { classDisplayName } from '../sim/spec-label';
import { isClassSlug } from '../ui/primitives';
/** One piece of a note: plain text or a link. */
export interface NoteSegment {
  text: string;
  href?: string;
}

/** The sim's single-target damage fight length; the BiS ranker's `DPS` runs share it. */
const DPS_FIGHT_SECONDS = 180;
/** The stand-in tank boss: the target the character already fights, three levels above a 60. */
const TANK_BOSS_LEVEL = 63;
const LEVELING_BANDS_HREF = '/bis';
/** The log rankings: how specs do in real raid logs, the counterpart to this page's sim. */
const LOG_RANKINGS_HREF = '/rankings';

export const ROLE_LABELS: Readonly<Record<BisRole, string>> = {
  dps: 'DPS',
  tank: 'Tank',
  healer: 'Healer',
};

/** The role as a word in a sentence, for the meta descriptions. */
const ROLE_NOUNS: Readonly<Record<BisRole, string>> = {
  dps: 'damage',
  tank: 'tank',
  healer: 'healer',
};

/** The role's path: DPS is the root, the other two are sub-paths. */
export const ROLE_PATHS: Readonly<Record<BisRole, string>> = {
  dps: '/tiers',
  tank: '/tiers/tank',
  healer: '/tiers/healer',
};

export const FACTION_LABELS: Readonly<Record<Faction, string>> = {
  alliance: 'Alliance',
  horde: 'Horde',
};

/** What the notes and the caption read from the published files, never typed in. */
export interface NoteContext {
  role: BisRole;
  faction: Faction;
  presetLabel: string;
  /** `heal_profile.duration_sec`; healer copy only. */
  healSeconds?: number;
  /** `tank-encounter.json` `boss.swing_speed_sec.value`; tank copy only. */
  bossSwingSeconds?: number;
  /** Whether the healer list's order is measured (`TierData.healerOrderMeasured`); healer copy only. */
  healerOrderMeasured?: boolean;
  /** Whether any row on the list carries a tie: the healer caption explains ≈ tie only then. */
  hasTie?: boolean;
}

/** One column of "How to read this": a caps label over its sentences. */
export interface ReadColumn {
  label: string;
  body: NoteSegment[];
}

function requireContext(value: number | undefined, name: string, role: BisRole): number {
  if (value === undefined) throw new Error(`tier list: the ${role} copy needs ${name}`);
  return value;
}

/**
 * How a row's race was chosen, in one place so the wording can change in one edit: the ranker
 * reads it from each guide's `recommendedRaces` frontmatter and does not sim the races (a ranker
 * race sweep is a later lane; when it lands this reads "best" again).
 */
export function raceBasis(faction: Faction): string {
  return `the guide's recommended ${FACTION_LABELS[faction]} race`;
}

function gearSentence(ctx: NoteContext): string {
  return `Every spec is simmed in its best ${ctx.presetLabel} gear, with enchants a fresh level 60 can buy, raid buffs and consumables, as ${raceBasis(ctx.faction)}, named under it.`;
}

function freshSixty(): NoteSegment[] {
  return [
    { text: ' At a fresh 60 the order differs: ' },
    { text: 'see the leveling bands', href: LEVELING_BANDS_HREF },
    { text: '.' },
  ];
}

/** "S is within 5% of the top, A within 10%, ..., D beyond." */
function tierLetterSentence(): string {
  const last = TIER_BANDS[TIER_BANDS.length - 1]!;
  const clauses = TIER_BANDS.slice(0, -1).map((band, index) => {
    const reach = `within ${TIER_BANDS[index + 1]!.from}%`;
    return index === 0 ? `${band.letter} is ${reach} of the top` : `${band.letter} ${reach}`;
  });
  return `${[...clauses, `${last.letter} beyond`].join(', ')}.`;
}

function dpsColumns(ctx: NoteContext): ReadColumn[] {
  return [
    {
      label: 'What the number is',
      body: [
        {
          text: `Damage per second from our own sim against one target for ${DPS_FIGHT_SECONDS} seconds. ${gearSentence(ctx)}`,
        },
      ],
    },
    {
      label: 'What it leaves out',
      body: [
        {
          text: 'Cleave, add fights, movement and what a spec brings the raid are not counted, so a spec can sit lower here than it plays in your raid. For how specs do in real raid logs, see ',
        },
        { text: 'the log rankings', href: LOG_RANKINGS_HREF },
        { text: '.' },
        ...freshSixty(),
      ],
    },
    {
      label: 'Letters and ties',
      body: [
        {
          text: `${tierLetterSentence()} ≈ tie names the specs within ${TIE_MARGIN_PERCENT}% of a spec; a tie across a line names the other letter, and the two are a tie.`,
        },
      ],
    },
  ];
}

/** The healer list's rule in plain words: why a pair ties and how a shared rank counts, or,
 *  with no measured error published yet, that the ranks are still to come. */
function healerOrderNote(measured: boolean): string {
  if (!measured) return 'Ranks appear once the next ranking publishes the measured error.';
  return '≈ tie means the sim cannot order the pair: their order changes with the heal profile’s tuning, or their gap is within the measured error. Tied healers share a rank: =2, =2, then 4.';
}

function healerColumns(ctx: NoteContext): ReadColumn[] {
  const seconds = requireContext(ctx.healSeconds, 'the heal profile duration', ctx.role);
  return [
    {
      label: 'What the number is',
      body: [
        {
          text: `Effective healing per second over ${seconds} seconds against a stand-in Phase 1 fight: tank hits and raid-wide pulses, not a named boss. Overhealing is not counted. ${gearSentence(ctx)}`,
        },
      ],
    },
    {
      label: 'What it leaves out',
      body: [
        {
          text: 'Mana and raid utility are not ranked here. No level-60 log is behind this order yet: it rests on the sim’s heal model until launch logs exist.',
        },
        ...freshSixty(),
      ],
    },
    {
      label: 'Ties and ranks',
      body: [
        {
          text: healerOrderNote(ctx.healerOrderMeasured !== false),
        },
      ],
    },
  ];
}

function tankColumns(ctx: NoteContext): ReadColumn[] {
  const swing = requireContext(ctx.bossSwingSeconds, 'the boss swing speed', ctx.role);
  return [
    {
      label: 'What the number is',
      body: [
        {
          text: 'Taken / s is the damage the tank takes a second from the boss, lower is better, and the list is sorted by it. Health is maximum health with this gear and raid buffs. Threat / s is threat a second while tanking. vs #1 is how much more damage a second than the top spec, and the bar is the top spec’s damage taken as a share of this spec’s. Specs within 1% of each other share a rank. The best of each column is green: lowest taken, highest health, highest threat.',
        },
      ],
    },
    {
      label: 'The boss',
      body: [
        {
          text: `A stand-in level ${TANK_BOSS_LEVEL} target with one melee swing every ${swing} seconds, not a named boss.`,
        },
      ],
    },
    { label: 'The gear', body: [{ text: gearSentence(ctx) }, ...freshSixty()] },
  ];
}

/** "How to read this": three columns per role, in the order a raider asks. */
export function readColumns(ctx: NoteContext): ReadColumn[] {
  if (ctx.role === 'tank') return tankColumns(ctx);
  return ctx.role === 'healer' ? healerColumns(ctx) : dpsColumns(ctx);
}

/** "Warrior": a class slug as the game names it (lib/sim/spec-label.ts). Fails fast on a slug
 *  that is not a playable class. */
export function classNameOf(classSlug: string): string {
  if (!isClassSlug(classSlug)) throw new Error(`tier list: no class named ${classSlug}`);
  return classDisplayName(classSlug);
}

/** "Fury Warrior", "Feral Druid (Bear)": the catalogue's full name. Fails fast on a spec the
 *  catalogue lacks. */
export function specFullName(spec: Pick<TierRow, 'classSlug' | 'specSlug'>): string {
  const key = `${spec.classSlug}-${spec.specSlug}`;
  const found = SPEC_BY_KEY[key];
  if (found === undefined) throw new Error(`tier list: no spec ${key} in the catalogue`);
  return found.full_name;
}

/** The role tabs on the band: the list each one opens, in the mock's words. */
export const ROLE_TABS: Readonly<Record<BisRole, string>> = {
  dps: 'Damage',
  tank: 'Tanks',
  healer: 'Healers',
};

const SECONDS_PER_MINUTE = 60;

/** "3 min", or "90 s" for a fight that is not whole minutes. */
function fightLength(seconds: number): string {
  return seconds % SECONDS_PER_MINUTE === 0 ? `${seconds / SECONDS_PER_MINUTE} min` : `${seconds} s`;
}

/** What every spec on the list wears and fights with (the ranker's raid preset). */
const RANKED_ON = 'Best in slot, raid buffs';

/** The title band's one fact line: the gear, then the fight (the target's level, its length). */
function metaLine(ctx: Pick<NoteContext, 'role' | 'healSeconds'>): string {
  const fight =
    ctx.role === 'healer'
      ? [fightLength(requireContext(ctx.healSeconds, 'the heal profile duration', ctx.role))]
      : ctx.role === 'tank'
        ? [`vs a level ${TANK_BOSS_LEVEL} boss`]
        : [`vs a level ${TANK_BOSS_LEVEL} boss`, fightLength(DPS_FIGHT_SECONDS)];
  return [RANKED_ON, ...fight].join(' · ');
}

/** "Within 5% of #1", "5 to 10% behind #1", "30% or more behind #1": a letter's hover title. */
function tierBand(letter: TierLetter): string {
  const index = TIER_BANDS.findIndex((band) => band.letter === letter);
  const from = TIER_BANDS[index]!.from;
  const to = TIER_BANDS[index + 1]?.from;
  if (from === 0) return `Within ${to}% of #1`;
  return to === undefined ? `${from}% or more behind #1` : `${from} to ${to}% behind #1`;
}

/** Whether a row shares its rank with a tied neighbour: healers within error, tanks within the
 *  margin. DPS ties keep their own ranks and letters. */
function sharesRank(row: Pick<TierRow, 'role' | 'tie'>): boolean {
  return row.role !== 'dps' && row.tie !== null;
}

const SHARED_RANK_REASON: Record<Exclude<BisRole, 'dps'>, (joined: string) => string> = {
  healer: (joined) =>
    `The sim cannot order this healer and ${joined}: the order changes with the heal profile’s tuning, or the gap is within the measured error`,
  tank: (joined) =>
    `This tank and ${joined} take damage within 1% of each other: too close for the sim to order`,
};

/** A shared rank's partners in full ("≈ tie with Restoration Shaman"), so the rank reads as the
 *  sim's answer and not a slip; `null` on a DPS row or an untied one. */
function rowTie(row: Pick<TierRow, 'role' | 'tiedWith'>): { text: string; title: string } | null {
  if (row.role === 'dps' || row.tiedWith.length === 0) return null;
  const names = row.tiedWith.map(specFullName);
  const joined =
    names.length === 1 ? names[0]! : `${names.slice(0, -1).join(', ')} and ${names[names.length - 1]}`;
  return { text: `≈ tie with ${joined}`, title: SHARED_RANK_REASON[row.role](joined) };
}

export const tiersCopy = {
  title: 'Tier list at 60',
  /** Each role's <title>: three pages, three titles, so none of them reads as a duplicate. */
  pageTitle: (role: BisRole): string => `${ROLE_LABELS[role]} tier list`,
  /** Each role's meta description, unique per page for the same reason. */
  metaDescription: (role: BisRole): string =>
    `Where every ${ROLE_NOUNS[role]} spec stands at level 60 in WoW Forever, ranked by the Forever Sixty sim with the gap to the top.`,
  metaLine,
  controlsLabel: 'Tier list choices',
  roleGroupLabel: 'Role',
  factionGroupLabel: 'Faction',
  lowConfidence: 'Stat weights less certain',
  lowConfidenceTitle:
    "This spec's stat weights did not settle cleanly, so its gear pick and number are less firm.",
  /** The rank column: a shared rank (tied healer or tank) reads "=2"; every other row its own. */
  rankLabel: (row: Pick<TierRow, 'rank' | 'role' | 'tie' | 'ranked'>): string => {
    if (!row.ranked) return '';
    return sharesRank(row) ? `=${row.rank}` : String(row.rank);
  },
  /** A DPS row's figure opens the ranking sim behind it (the words of /bis's "Open in the sim"). */
  simLinkTitle: (spec: string): string => `Open in the sim: ${spec}`,
  simLinkLabel: (spec: string, value: string): string => `${value} DPS, open ${spec} in the sim`,
  sharesRank,
  topGap: 'Top',
  leastGap: 'Least',
  bestTitle: (count: number, role: BisRole): string =>
    `Best of the ${count} ${ROLE_LABELS[role].toLowerCase()}s`,
  tierLabel: (letter: TierLetter): string => `Tier ${letter}`,
  tierBand,
  rowTie,
  listLabel: (role: BisRole, faction: Faction): string =>
    `${ROLE_LABELS[role]} tier list, ${FACTION_LABELS[faction]}`,
  emptyRole: (role: BisRole): string => `No ${ROLE_LABELS[role]} specs are simmed yet.`,
  noData: 'No tier list is published for this build yet.',
  /** The unit after each figure: the sorted one, then a tank's health and threat. */
  units: {
    dps: 'DPS',
    healer: 'HPS',
    tank: 'taken / s',
    health: 'HP',
    threat: 'threat / s',
  },
  /** What assistive tech hears before each figure, and the phone label of a tank's extras. */
  figureLabels: {
    dps: 'Damage per second',
    healer: 'Healing per second',
    tank: 'Damage taken per second',
    health: 'Health',
    threat: 'Threat',
  },
  /** The Health figure's hover: it is the pool, not a figure made of the damage taken. */
  healthTitle: 'Maximum health with this gear and raid buffs',
  gapLabel: {
    dps: 'Gap to the top',
    healer: 'Gap to the top',
    tank: 'More damage taken than #1',
  } as const satisfies Record<BisRole, string>,
  /** The column heads: rank, crest, spec, bar, the figures (each head names its unit), the gap
   *  and the empty link column. */
  heads: {
    dps: ['#', '', 'Spec', '', 'DPS', 'vs #1', ''],
    healer: ['#', '', 'Spec', '', 'HPS', 'vs #1', ''],
    tank: ['#', '', 'Spec', '', 'Taken / s', 'Health', 'Threat / s', 'vs #1', ''],
  } as const satisfies Record<BisRole, readonly string[]>,
  /** The tank heads of a list whose bands do not all publish health: no half-filled column. */
  tankHeadsWithoutHealth: ['#', '', 'Spec', '', 'Taken / s', 'Threat / s', 'vs #1', ''] as const,
  /** The vs #1 cell on the top spec's own row; a tank list's top takes the least damage. */
  leaderGap: 'best',
  tankLeaderGap: 'lowest',
  /** A tank's vs #1 in words, on hover. */
  tankGapTitle: (gapPercent: number, leaderName: string): string =>
    `Takes ${Math.round(gapPercent)}% more damage a second than ${leaderName}`,
  /** The spec panel beside the list (directions mock g-tiers). */
  card: {
    heading: 'Your spec',
    signedOut: 'Sign in to mark your spec',
    signIn: 'Sign in',
    signInHref: (path: string): string => `/login?next=${encodeURIComponent(path)}`,
    bis: 'Best in slot',
    guide: 'Guide',
    otherRole: (role: BisRole): string => `On the ${ROLE_TABS[role]} list`,
    none: 'No character chosen.',
    unknown: answerCopy.unknown,
    choose: answerCopy.unknownLink,
    error: answerCopy.error,
    retry: answerCopy.retry,
  },
} as const;

/** A tank figure is a per-second average over thousands of runs, so it is written whole; a
 *  damage or healing rate keeps its one decimal. */
export function formatMetric(role: BisRole, value: number): string {
  return role === 'tank' ? grouped(value) : oneDecimal(value);
}

/** The gap column's text for a row: "Top"/"Least" on the leader, "−4.8%"/"+15.2%" below. */
export function gapText(row: Pick<TierRow, 'rank' | 'role' | 'gapPercent' | 'ranked'>): string {
  const figure = row.gapPercent.toFixed(1);
  // An unranked list claims no leader: its top row reads its own gap, 0.0%.
  if (!row.ranked) return row.gapPercent === 0 ? `${figure}%` : `−${figure}%`;
  // A healer tied for the top shares rank 1 but still has a gap; only the leader reads Top.
  if (row.rank === 1 && row.gapPercent === 0)
    return row.role === 'tank' ? tiersCopy.leastGap : tiersCopy.topGap;
  // A tank's gap is whole percent: a tenth of a percent of damage taken is inside the sim's noise.
  return row.role === 'tank' ? `+${Math.round(row.gapPercent)}% taken` : `−${figure}%`;
}
