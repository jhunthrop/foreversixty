// web/src/lib/home/home-b-view.ts
// Home B's facts and words (design/specs/2026-11-07-home-b.md §1): the head's three figures, the
// countdown panels, Your spec this week, Top movers and Guild sims. Pure: the island
// (components/home/HomeCharacter.svelte) and the unit tests read them.
import classes from '../../data/classes.json';
import { characterHref } from '../characters';
import { SPECS } from '../sim/specs';
import type { GuildSims, ImportSim, Movers, SpecWeek } from './home-answer';
import { bandWords, shortDate } from './home-view';

const MAX_LEVEL = 60;
const BAND_STEP = 10;
/** Slots on a best-in-slot list. */
export const BIS_SLOTS = 17;
/** The guild ladder's rows above the viewer's pinned row (ruling R2). */
export const GUILD_SIMS_SHOWN = 5;


export const homeBCopy = {
  simmedOnImport: 'Simmed on import',
  simmed: (date: string) => `Simmed ${date}`,
  simPending: 'Sim pending',
  simFailed: 'Sim failed',
  runItYourself: 'Run it yourself',
  noGear: 'No gear imported',
  importCharacter: 'Import your character',
  upgradesLabel: (n: number) => `${n} ${n === 1 ? 'upgrade' : 'upgrades'}`,
  toBand: (band: number) => `To band ${band}`,
  levels: 'levels',
  onBis: 'On best in slot',
  ofSlots: (n: number) => `${n} of ${BIS_SLOTS}`,
  raidsEyebrow: (day: string) => `First raids open · ${day}`,
  raidsOpen: 'Raids open now',
  guildEyebrow: (guild: string) => `${guild} · next guild event`,
  guildNone: (guild: string) => `${guild}: no raid planned`,
  findGuild: 'Find a guild',
  findGuildSub: 'Raid with a guild that plans its nights here.',
  guildsLink: 'Guilds',
  raidPageLink: 'Raid page',
  signUpLink: 'Sign up on the raid page',
  you: (state: string) => `you: ${state}`,
  notSigned: 'not signed',
  inCount: (n: number) => `${n} in`,
  upgrades: 'Your upgrades',
  bisLink: 'Best in slot',
  guildSims: 'Guild sims',
  guildPage: 'Guild page',
  youNoSim: 'You: no sim yet',
  runSim: 'Run a sim',
  specWeek: 'Your spec this week',
  tierList: 'Tier list',
  movers: 'Top movers',
  topThree: (spec: string, band: string) => `Top 3 for ${spec}, ${band}`,
  rankings: 'Rankings',
} as const;

export const homeBHrefs = {
  sim: '/sim',
  setup: '/setup',
  guilds: '/guild',
  tiers: '/tiers',
  rankings: '/rankings',
} as const;

export interface HeadFigure {
  key: 'sim' | 'gain' | 'band';
  value: string;
  unit?: string;
  label: string;
  caption?: string;
  tone: 'strong' | 'kill' | 'muted';
  link?: { label: string; href: string };
}

const oneDecimal = (n: number) =>
  n.toLocaleString('en-US', { minimumFractionDigits: 1, maximumFractionDigits: 1 });

/** The first figure: the sim on import (ruling P11 head 1). */
export function simFigure(sim: ImportSim | null): HeadFigure {
  if (sim === null) {
    return {
      key: 'sim',
      value: homeBCopy.noGear,
      label: homeBCopy.simmedOnImport,
      tone: 'muted',
      link: { label: homeBCopy.importCharacter, href: homeBHrefs.setup },
    };
  }
  if (sim.state === 'pending')
    return { key: 'sim', value: homeBCopy.simPending, label: homeBCopy.simmedOnImport, tone: 'muted' };
  if (sim.state === 'failed' || sim.dps === null) {
    return {
      key: 'sim',
      value: homeBCopy.simFailed,
      label: homeBCopy.simmedOnImport,
      tone: 'muted',
      link: { label: homeBCopy.runItYourself, href: homeBHrefs.sim },
    };
  }
  return {
    key: 'sim',
    value: oneDecimal(sim.dps),
    unit: 'DPS',
    label: homeBCopy.simmedOnImport,
    caption: sim.simmed_at === null ? undefined : homeBCopy.simmed(shortDate(sim.simmed_at)),
    tone: 'strong',
  };
}

/** The second figure: what the upgrades add, or nothing to say. */
export function gainFigure(total: number, count: number): HeadFigure | null {
  if (count === 0) return null;
  return {
    key: 'gain',
    value: `+${oneDecimal(total)}`,
    unit: 'DPS',
    label: homeBCopy.upgradesLabel(count),
    tone: 'kill',
  };
}

/** The third figure: levels to the next band, or at 60 the slots already on the list. */
export function bandFigure(level: number | undefined, onBis: number | null): HeadFigure | null {
  if (level === undefined) return null;
  if (level >= MAX_LEVEL) {
    return onBis === null
      ? null
      : { key: 'band', value: homeBCopy.ofSlots(onBis), label: homeBCopy.onBis, tone: 'strong' };
  }
  const next = (Math.floor(level / BAND_STEP) + 1) * BAND_STEP;
  return {
    key: 'band',
    value: String(next - level),
    unit: homeBCopy.levels,
    label: homeBCopy.toBand(next),
    tone: 'strong',
  };
}

/** `Fury Warrior` from `warrior-fury`. */
export function specName(specKey: string): string {
  const spec = SPECS.find((row) => row.spec === specKey);
  if (spec === undefined) return specKey;
  const cls = classes.find((row) => row.slug === spec.class_slug)?.name ?? spec.class_slug;
  return `${spec.name} ${cls}`;
}

const ROLE_WORDS = { dps: 'damage specs', tank: 'tank specs', healer: 'healer specs' } as const;

/** Your spec this week: the title and the line under it (ruling R3). No delta without last week. */
export function specWeekView(week: SpecWeek): { title: string; sub: string } {
  const title = `${specName(week.spec)} at ${bandWords(week.band)}`;
  const place = `#${week.place} of ${week.of} ${ROLE_WORDS[week.role]}`;
  if (week.grade === null || week.last_week_grade === null) return { title, sub: place };
  const order = 'SABCD';
  const now = order.indexOf(week.grade);
  const before = order.indexOf(week.last_week_grade);
  const move =
    now === before
      ? 'Same as last week'
      : now < before
        ? `Up from ${week.last_week_grade} last week`
        : `Down from ${week.last_week_grade} last week`;
  return { title, sub: `${move}, ${place}` };
}

/** The movers block's title: movers with a baseline, else the honest top 3 (ruling R4). */
export function moversTitle(movers: Movers): string {
  return movers.baseline
    ? homeBCopy.movers
    : homeBCopy.topThree(specName(movers.spec), bandWords(movers.band));
}

export function moversMeta(movers: Movers): string | undefined {
  if (!movers.baseline) return undefined;
  const spec = SPECS.find((row) => row.spec === movers.spec)?.name ?? movers.spec;
  return `${spec}, ${bandWords(movers.band)}`;
}

/** `▲ 6`, `▼ 2`, or nothing (no baseline, or no move). */
export function deltaText(delta: number | null): string | null {
  if (delta === null || delta === 0) return null;
  return delta > 0 ? `▲ ${delta}` : `▼ ${-delta}`;
}

/** `DPS, 60 raid` */
export function guildSimsMeta(sims: GuildSims): string {
  const role = sims.role === 'dps' ? 'DPS' : sims.role === 'tank' ? 'Tanks' : 'Healers';
  return `${role}, ${bandWords(sims.band)} ${sims.preset}`;
}

export function figureText(value: number): string {
  return oneDecimal(value);
}

/** A character's page from its key (`us/normal/thalrin`). */
export function characterLink(key: string, name: string): string {
  const [region = '', ruleset = ''] = key.split('/');
  return characterHref(region, ruleset, name);
}
