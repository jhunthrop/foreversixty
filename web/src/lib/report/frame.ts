// web/src/lib/report/frame.ts
// The report page's frame in words (design/specs/2026-11-07-logs-reports.md section 6.1): the one
// tab strip, the Fight select's rows, the band's meta and figure and the rail's This pull facts. Pure, so
// the island and its tests read the same rules.
import type { SelectOption } from '../ui/listbox-model';
import type { Fact } from '../ui/facts';
import { joinFacts, outcomeLabel } from './format';
import { logClock } from '../dates';
import { pullNumbers, trashGroups } from './fights';
import type { FightEntry } from './types';
import { ALL_FIGHTS, trashScope, type Tab } from './url';
import { counted, grouped, oneDecimal } from '../format/number';
import { clockSeconds } from '../format/duration';

/** Every tab in the strip's order (decision R3); the modes and views of old are tabs here. */
export const REPORT_TABS: readonly { id: Tab; label: string }[] = [
  { id: 'summary', label: 'Summary' },
  { id: 'damage-done', label: 'Damage done' },
  { id: 'healing', label: 'Healing' },
  { id: 'damage-taken', label: 'Damage taken' },
  { id: 'deaths', label: 'Deaths' },
  { id: 'mechanics', label: 'Mechanics' },
  { id: 'timeline', label: 'Timeline' },
  { id: 'threat', label: 'Threat' },
  { id: 'casts', label: 'Casts' },
  { id: 'buffs', label: 'Buffs' },
  { id: 'debuffs', label: 'Debuffs' },
  { id: 'rating', label: 'Rating' },
  { id: 'coach', label: 'Coach' },
  { id: 'interrupts', label: 'Interrupts' },
  { id: 'dispels', label: 'Dispels' },
  { id: 'resources', label: 'Resources' },
  { id: 'events', label: 'Events' },
  { id: 'queries', label: 'Queries' },
  { id: 'compare', label: 'Compare' },
];

/** Tabs in the strip below 1760 px; the rest sit in More. */
export const TABS_MD = 7;
/** Tabs in the strip from 1760 px when Coach is listed (12 without it). */
export const TABS_XL = 13;

/** The tabs this viewer sees: Coach only for an owner of a character in the pull, or an officer. */
export function listedTabs(coachListed: boolean): { id: Tab; label: string }[] {
  return REPORT_TABS.filter((tab) => tab.id !== 'coach' || coachListed);
}

/** The tabs that mean something over a whole night (section 6.3); the rest say pick a pull. */
export const NIGHT_TABS: ReadonlySet<Tab> = new Set([
  'summary',
  'damage-done',
  'healing',
  'damage-taken',
  'mechanics',
]);

/** The tabs whose Source scope narrows a table with target and ability filters. */
export const ACTOR_TABS: ReadonlySet<Tab> = new Set(['damage-done', 'healing', 'damage-taken']);

const plural = counted;

/** A pull's length as m:ss, short pulls included ("0:04"), the way the fight menu reads. */
export function pullLength(ms: number): string {
  return clockSeconds(Math.round(ms / 1000));
}

/** The Fight select's face: "Trash · 12:43", "Warden Kelthas #2 · Wipe 34% · 1:12". */
export function fightLabel(fight: FightEntry, attempt?: number): string {
  if (fight.kind !== 'encounter') return `${fight.name} · ${pullLength(fight.duration_ms)}`;
  const outcome = fight.in_progress ? 'Live' : fight.kill ? 'Kill' : outcomeLabel(fight);
  return joinFacts(`${fight.name} #${attempt ?? 1}`, outcome, pullLength(fight.duration_ms));
}

/** The Fight select's option for "every boss pull" (the night view). */
export const ALL_BOSS_PULLS = 'All boss pulls';
export const NO_BOSS_PULLS = 'No boss pulls in this report.';
const BOSS_GROUP = 'Boss pulls';

/** The trash pulls that killed nothing: hidden behind the menu's foot toggle until asked for. */
export function idlePulls(fights: readonly FightEntry[]): FightEntry[] {
  return fights.filter((fight) => fight.kind !== 'encounter' && fight.mob_deaths === 0);
}

/** The foot toggle's words ("Show 7 pulls that killed nothing"); '' when every pull killed something. */
export function idleToggleLabel(fights: readonly FightEntry[]): string {
  const idle = idlePulls(fights).length;
  return idle === 0 ? '' : `Show ${plural(idle, 'pull')} that killed nothing`;
}

/**
 * The Fight select's rows (decision R2 as ruled): Boss pulls first (the night, then each pull
 * numbered per boss) or one disabled line when there are none; then Trash, each pull that killed
 * something, and the idle ones only when `showIdle` (or when one is the selected pull).
 */
export function fightOptions(
  fights: readonly FightEntry[],
  selected: number,
  showIdle: boolean,
): SelectOption[] {
  const pulls = pullNumbers(fights);
  const bosses = fights.filter((fight) => fight.kind === 'encounter');
  const bossRows: SelectOption[] =
    bosses.length === 0
      ? [{ value: 'none', label: NO_BOSS_PULLS, group: BOSS_GROUP, disabled: true }]
      : [
          {
            value: String(ALL_FIGHTS),
            label: ALL_BOSS_PULLS,
            description: plural(bosses.length, 'pull'),
            group: BOSS_GROUP,
          },
          ...bosses.map((fight) => ({
            value: String(fight.index),
            label: fightLabel(fight, pulls.get(fight.index)?.pull),
            description: `${logClock(fight.start)} · ${plural(fight.deaths ?? 0, 'death')}`,
            group: BOSS_GROUP,
          })),
        ];
  const trash = fights.filter(
    (fight) => fight.kind !== 'encounter' && (showIdle || fight.mob_deaths !== 0 || fight.index === selected),
  );
  const trashGroup = `Trash · ${plural(trash.length, 'pull')}`;
  // Each zone's trash folds into one entry first (the night's trash, the way Warcraft Logs
  // collapses it): its pulls read together, the night view's Summary and tables over them.
  const zoneRows = trashGroups(fights)
    .filter((group) => group.fights.length > 1)
    .map((group) => ({
      value: String(trashScope(group.key)),
      label: `All trash in ${group.zone}`,
      description: plural(group.fights.length, 'pull'),
      group: trashGroup,
    }));
  // A pull indexed before engine 0.6.0 has no true kill count: its row says when, never a count.
  const trashRows = trash.map((fight) => ({
    value: String(fight.index),
    label: fightLabel(fight),
    description:
      fight.mob_deaths === undefined
        ? logClock(fight.start)
        : `${logClock(fight.start)} · ${plural(fight.mob_deaths, 'mob')} died`,
    group: trashGroup,
  }));
  return [...bossRows, ...zoneRows, ...trashRows];
}

/** A per-second figure: one decimal under 100 (1.8), whole from 100 (1,374). */
export function formatRate(perSecond: number): string {
  const safe = Number.isFinite(perSecond) ? perSecond : 0;
  return safe < 100 ? oneDecimal(safe) : grouped(safe);
}

/** The band's meta line (directions b1): "Zephras Isle · Pull 6, Trash · 12:43 · 177 mobs died",
 *  "Pull 2, Onyxia #1 · Wipe at 34% · 4:10 · 40 players". Without `players` the line stops at the
 *  pull's length: the band then draws `reportMetaTail` itself, from 768 only, so a phone's meta
 *  stays one line. */
export function reportMeta(
  zone: string,
  fight: FightEntry,
  attempt: number | undefined,
  players?: number,
): string {
  const tail = players === undefined ? '' : reportMetaTail(fight, players).replace(/^ · /, '');
  if (fight.kind !== 'encounter') {
    return joinFacts(zone, `Pull ${fight.index}, ${fight.name}`, pullLength(fight.duration_ms), tail);
  }
  const outcome = fight.in_progress ? 'Live' : fight.kill ? 'Kill' : outcomeLabel(fight);
  return joinFacts(
    zone,
    `Pull ${fight.index}, ${fight.name} #${attempt ?? 1}`,
    outcome,
    pullLength(fight.duration_ms),
    tail,
  );
}

/** The meta line's last fact: " · 40 players" on a boss pull, " · 177 mobs died" on trash. */
export function reportMetaTail(fight: FightEntry, players: number): string {
  return fight.kind === 'encounter'
    ? ` · ${plural(players, 'player')}`
    : ` · ${plural(fight.mob_deaths ?? 0, 'mob')} died`;
}

/** The band's one figure and its unit (directions b1: the figure follows the tab). */
export interface ReportFigure {
  value: string;
  unit: string;
}

export interface FigureFacts {
  tab: Tab;
  /** One player's own view: their DPS, HPS or DTPS, not the group's. */
  player: boolean;
  /** A boss pull: Mechanics counts its deaths; on trash Mechanics has nothing to count. */
  bossPull: boolean;
  /** The scope's totals over the pull (or the window). */
  damageDone: number;
  healing: number;
  damageTaken: number;
  deaths: number;
  durationMs: number;
}

const RATE_UNIT: Partial<Record<Tab, { key: 'damageDone' | 'healing' | 'damageTaken'; unit: string }>> = {
  summary: { key: 'damageDone', unit: 'DPS' },
  'damage-done': { key: 'damageDone', unit: 'DPS' },
  healing: { key: 'healing', unit: 'HPS' },
  'damage-taken': { key: 'damageTaken', unit: 'DTPS' },
};

/** Group DPS, HPS or DTPS on the rate tabs, the player's own on their view, deaths on Deaths and
 *  Mechanics (a boss pull's); null where no single figure answers the tab (Timeline, Threat, Resources, Queries,
 *  Coach and the rest). */
export function reportFigure(facts: FigureFacts): ReportFigure | null {
  if (facts.tab === 'deaths' || (facts.tab === 'mechanics' && facts.bossPull))
    return { value: grouped(facts.deaths), unit: facts.deaths === 1 ? 'death' : 'deaths' };
  const rate = RATE_UNIT[facts.tab];
  if (rate === undefined) return null;
  const seconds = Math.max(facts.durationMs, 1) / 1000;
  return {
    value: formatRate(facts[rate.key] / seconds),
    unit: facts.player ? rate.unit : `group ${rate.unit}`,
  };
}

export interface PullFacts {
  fight: FightEntry;
  phases: readonly { name: string; start_ms: number; end_ms: number }[];
}

/** This pull's facts (directions b1): only what the band does not say, the start and a boss pull's
 *  phases. The scope's count and deaths are the band's and the tables' (h-audit 2026-11-09). */
export function thisPullFacts({ fight, phases }: PullFacts): Fact[] {
  return [
    { id: 'started', label: 'Started', value: logClock(fight.start) },
    ...phases.map((phase, index) => ({
      id: `phase-${index}`,
      label: phase.name,
      value: `${pullLength(phase.start_ms)}–${pullLength(phase.end_ms)}`,
    })),
  ];
}
