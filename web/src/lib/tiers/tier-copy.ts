// web/src/lib/tiers/tier-copy.ts
// Every string the tier list shows (design/specs/2026-10-09-tier-list.md, section 6). The
// strings describe the rules; a figure that moves nightly is passed in from the published
// files, never written here (the owner's "prose never quotes nightly numbers" rule).
import type { BisRole, Faction } from '../bis/types';
import { TIE_MARGIN_PERCENT, type TierRow, type TierTie } from './tier-list';

/** One piece of a note: plain text, bold text or a link. */
export interface NoteSegment {
  text: string;
  strong?: boolean;
  href?: string;
}

/** The sim's single-target damage fight length; the BiS ranker's `DPS` runs share it. */
const DPS_FIGHT_SECONDS = 180;
/** The stand-in tank boss: the target the character already fights, three levels above a 60. */
const TANK_BOSS_LEVEL = 63;
const LEVELING_BANDS_HREF = '/bis';

export const ROLE_LABELS: Readonly<Record<BisRole, string>> = {
  dps: 'DPS',
  tank: 'Tank',
  healer: 'Healer',
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

export interface NoteContext {
  role: BisRole;
  faction: Faction;
  presetLabel: string;
  /** `heal_profile.duration_sec`; healer notes only. */
  healSeconds?: number;
  /** `tank-encounter.json` `boss.swing_speed_sec.value`; tank notes only. */
  bossSwingSeconds?: number;
}

function plain(text: string): NoteSegment[] {
  return [{ text }];
}

function requireContext(value: number | undefined, name: string, role: BisRole): number {
  if (value === undefined) throw new Error(`tier list: the ${role} notes need ${name}`);
  return value;
}

function freshSixtyNote(presetLabel: string): NoteSegment[] {
  return [
    { text: `This is full ${presetLabel} gear. At a fresh 60 the order differs: ` },
    { text: 'see the leveling bands', href: LEVELING_BANDS_HREF },
    { text: '.' },
  ];
}

function raceSentence(faction: Faction): string {
  return `Each spec is simmed as its best ${FACTION_LABELS[faction]} race, named under it.`;
}

function dpsNotes(ctx: NoteContext): NoteSegment[][] {
  return [
    plain(
      `Damage per second on one target for ${DPS_FIGHT_SECONDS} seconds, with raid buffs and consumables. Cleave, adds, movement and what a spec brings the raid are not counted.`,
    ),
    freshSixtyNote(ctx.presetLabel),
    [
      { text: `${raceSentence(ctx.faction)} ` },
      { text: '≈ tie', strong: true },
      { text: ` marks two specs within ${TIE_MARGIN_PERCENT}% of each other.` },
    ],
  ];
}

function healerNotes(ctx: NoteContext): NoteSegment[][] {
  const seconds = requireContext(ctx.healSeconds, 'the heal profile duration', ctx.role);
  return [
    plain(
      `Effective healing per second over ${seconds} seconds against a stand-in Phase 1 fight: tank hits and raid-wide pulses, not a named boss. Overhealing is not counted. Mana and raid utility are not ranked here.`,
    ),
    freshSixtyNote(ctx.presetLabel),
    [
      { text: `${raceSentence(ctx.faction)} ` },
      { text: 'Gap to the top', strong: true },
      { text: ' is less healing in this fight, not a verdict on the class.' },
    ],
  ];
}

function tankNotes(ctx: NoteContext): NoteSegment[][] {
  const swing = requireContext(ctx.bossSwingSeconds, 'the boss swing speed', ctx.role);
  return [
    plain(
      "Sorted by damage taken per second, lower is better: it is the sim's outcome against the boss, where effective health is only the hit-point pool going in. The best of each column is green.",
    ),
    plain(
      `The boss is a stand-in level ${TANK_BOSS_LEVEL} target with one melee swing every ${swing} seconds, not a named boss; the tank wears raid-ready gear with tank consumables.`,
    ),
    [...freshSixtyNote(ctx.presetLabel), { text: ` ${raceSentence(ctx.faction)}` }],
  ];
}

/** The three short lines above the list, one fact each, in the order a raider asks. */
export function tierNotes(ctx: NoteContext): NoteSegment[][] {
  if (ctx.role === 'tank') return tankNotes(ctx);
  return ctx.role === 'healer' ? healerNotes(ctx) : dpsNotes(ctx);
}

/** The long-form box below the list: only what the notes above do not already say. */
export function howToReadBullets(role: BisRole): string[] {
  return role === 'dps'
    ? ['A spec can sit lower here than it plays in your raid.']
    : ["Same boss profile as the BiS pages, so the number here is the number on the spec's page."];
}

const ORDINAL_SUFFIXES: Readonly<Record<number, string>> = { 1: 'st', 2: 'nd', 3: 'rd' };
const ORDINAL_TEEN_FLOOR = 10;
const ORDINAL_TEEN_CEILING = 20;
const ORDINAL_BASE = 100;
const ORDINAL_DIGIT_BASE = 10;

export function ordinal(n: number): string {
  const teen = n % ORDINAL_BASE >= ORDINAL_TEEN_FLOOR && n % ORDINAL_BASE <= ORDINAL_TEEN_CEILING;
  return `${n}${teen ? 'th' : (ORDINAL_SUFFIXES[n % ORDINAL_DIGIT_BASE] ?? 'th')}`;
}

function tieTitle(tie: TierTie): string {
  if (tie === 'both') return `Within ${TIE_MARGIN_PERCENT}% of the specs above and below: a tie`;
  return `Within ${TIE_MARGIN_PERCENT}% of the spec ${tie}: a tie`;
}

export const tiersCopy = {
  title: 'Tier list',
  description:
    'Where every spec stands at level 60, from the Forever Sixty sim: damage, tank and healer lists with the gap to the top.',
  eyebrow: (presetLabel: string): string => `Level 60 · ${presetLabel}`,
  jobLine:
    'Where every spec stands at level 60, from our own sim. Open yours for its best gear, guide and talents.',
  roleTabsLabel: 'Role',
  factionGroupLabel: 'Faction',
  calloutButton: 'Your best in slot →',
  yourSpec: 'Your spec',
  lowConfidence: 'Stat weights less certain',
  lowConfidenceTitle:
    "This spec's stat weights did not settle cleanly, so its gear pick and number are less firm.",
  tie: '≈ tie',
  tieTitle,
  topGap: 'Top',
  leastGap: 'Least',
  rowLinks: { bis: 'BiS', guide: 'Guide', planner: 'Planner' },
  plannerTitle: (specName: string): string => `Open the ${specName} level 60 build in the planner`,
  rulerLabel: (percent: number): string => `${percent}% or more behind the top`,
  howToReadHeading: 'How to read this',
  simChecksLink: 'How we check the sim →',
  simChecksHref: '/sim/specs',
  emptyRole: (role: BisRole): string => `No ${ROLE_LABELS[role]} specs are simmed yet.`,
  noData: 'No tier list is published for this build yet.',
  rowUnit: { dps: 'DPS', healer: 'HPS', tank: 'taken per sec' } as const satisfies Record<BisRole, string>,
  headers: {
    rank: '#',
    spec: 'Spec',
    bar: {
      dps: "Bar: share of the top spec's damage",
      healer: "Bar: share of the top healer's healing",
      tank: 'Bar: how close to the least damage taken',
    } as const satisfies Record<BisRole, string>,
    number: {
      dps: 'Damage per second',
      healer: 'Healing per second',
      tank: 'Damage taken per second, lower is better',
    } as const satisfies Record<BisRole, string>,
    gap: {
      dps: 'Gap to the top, %',
      healer: 'Gap to the top, %',
      tank: 'More taken than best',
    } as const satisfies Record<BisRole, string>,
    effectiveHealth: 'Effective health (HP)',
    threat: 'Threat per second',
  },
  effectiveHealthShort: 'Effective health',
  threatShort: 'Threat/s',
} as const;

/** The gap column's text for a row: "Top"/"Least" on the leader, "−4.8%"/"+15.2%" below. */
export function gapText(row: Pick<TierRow, 'rank' | 'role' | 'gapPercent'>): string {
  if (row.rank === 1) return row.role === 'tank' ? tiersCopy.leastGap : tiersCopy.topGap;
  const figure = row.gapPercent.toFixed(1);
  return row.role === 'tank' ? `+${figure}%` : `−${figure}%`;
}
