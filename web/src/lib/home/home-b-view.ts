// web/src/lib/home/home-b-view.ts
// Home's facts and words (design/specs/2026-11-07-home-b.md §1, round 5 2026-11-09): the head's
// sim figure, the plaque's guild lines, Your spec this week and Guild sims. Pure: the island
// (components/home/HomeCharacter.svelte) and the unit tests read them.
import { characterHref } from '../characters';
import { SPECS } from '../sim/specs';
import type { ImportSim, SpecWeek } from './home-answer';
import { bandWords, shortDate } from './home-view';
import { fillShared } from '../copy/shared';
import { compactAgo } from '../dates';
import {
  sharedHomeBisLink,
  sharedHomeFindGuildSub,
  sharedHomeGuildNone,
  sharedHomeGuildSims,
  sharedHomeGuildsLink,
  sharedHomeImportCharacter,
  sharedHomeNoGear,
  sharedHomeRaidPageLink,
  sharedHomeRaidsEyebrow,
  sharedHomeRaidsOpen,
  sharedHomeRunItYourself,
  sharedHomeRunSim,
  sharedHomeSignUpLink,
  sharedHomeSimFailed,
  sharedHomeSimPending,
  sharedHomeSimmed,
  sharedHomeSimmedOnImport,
  sharedHomeSpecWeek,
  sharedHomeTierList,
  sharedHomeUpgrades,
  sharedHomeYouNoSim,
} from '../copy/shared/home.generated';

/** The guild ladder's rows above the viewer's pinned row (ruling R2). */
export const GUILD_SIMS_SHOWN = 5;

export const homeBCopy = {
  simmedOnImport: sharedHomeSimmedOnImport,
  /** The import sim's link in the head: "Sim 8 h ago". */
  simLink: 'Sim',
  simmed: (date: string) => fillShared(sharedHomeSimmed, { date }),
  simPending: sharedHomeSimPending,
  simFailed: sharedHomeSimFailed,
  runItYourself: sharedHomeRunItYourself,
  noGear: sharedHomeNoGear,
  importCharacter: sharedHomeImportCharacter,
  raidsEyebrow: (day: string) => fillShared(sharedHomeRaidsEyebrow, { day }),
  raidsOpen: sharedHomeRaidsOpen,
  guildNone: (guild: string) => fillShared(sharedHomeGuildNone, { guild }),
  findGuildSub: sharedHomeFindGuildSub,
  guildsLink: sharedHomeGuildsLink,
  raidPageLink: sharedHomeRaidPageLink,
  signUpLink: sharedHomeSignUpLink,
  upgrades: sharedHomeUpgrades,
  bisLink: sharedHomeBisLink,
  guildSims: sharedHomeGuildSims,
  youNoSim: sharedHomeYouNoSim,
  runSim: sharedHomeRunSim,
  specWeek: sharedHomeSpecWeek,
  tierList: sharedHomeTierList,
} as const;

export const homeBHrefs = {
  sim: '/sim',
  setup: '/setup',
  guilds: '/guild',
  tiers: '/tiers',
} as const;

export interface HeadFigure {
  key: 'sim';
  value: string;
  unit?: string;
  label: string;
  caption?: string;
  tone: 'strong' | 'kill' | 'muted';
  link?: { label: string; href: string };
}

const oneDecimal = (n: number) =>
  n.toLocaleString('en-US', { minimumFractionDigits: 1, maximumFractionDigits: 1 });

/** The finished import sim's one line in the head: which sim and how old, linking that result. */
function simLine(sim: ImportSim, now: Date): HeadFigure['link'] {
  if (sim.sim_id === null) return undefined;
  const age = sim.simmed_at === null ? '' : compactAgo(new Date(sim.simmed_at), now);
  return {
    label: age === '' ? homeBCopy.simLink : `${homeBCopy.simLink} ${age}`,
    href: `${homeBHrefs.sim}/${encodeURIComponent(sim.sim_id)}`,
  };
}

/** The first figure: the sim on import (ruling P11 head 1). */
export function simFigure(sim: ImportSim | null, now: Date = new Date()): HeadFigure {
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
    link: simLine(sim, now),
  };
}

/** `Fury Warrior` from `warrior-fury`. */
export function specName(specKey: string): string {
  return SPECS.find((row) => row.spec === specKey)?.full_name ?? specKey;
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

export function figureText(value: number): string {
  return oneDecimal(value);
}

/** A character's page from its key (`us/normal/thalrin`). */
export function characterLink(key: string, name: string): string {
  const [region = '', ruleset = ''] = key.split('/');
  return characterHref(region, ruleset, name);
}
