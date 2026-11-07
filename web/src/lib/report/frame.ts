// web/src/lib/report/frame.ts
// The report page's frame in words (design/specs/2026-11-07-logs-reports.md section 6.1): the one
// tab strip, the Fight select's rows, the head's answer and the rail's This pull facts. Pure, so
// the island and its tests read the same rules.
import type { SelectOption } from '../ui/listbox-model';
import type { Fact } from '../ui/facts';
import { formatClock, formatDuration, joinFacts, outcomeLabel } from './format';
import { pullNumbers } from './fights';
import type { FightEntry } from './types';
import { ALL_FIGHTS, type Tab } from './url';

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

const plural = (n: number, one: string, many = `${one}s`): string =>
  `${n.toLocaleString('en-US')} ${n === 1 ? one : many}`;

/** A pull's length as m:ss, short pulls included ("0:04"), the way the fight menu reads. */
export function pullLength(ms: number): string {
  const total = Math.max(0, Math.round(ms / 1000));
  return `${Math.floor(total / 60)}:${String(total % 60).padStart(2, '0')}`;
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
            description: `${formatClock(fight.start)} · ${plural(fight.deaths ?? 0, 'death')}`,
            group: BOSS_GROUP,
          })),
        ];
  const trash = fights.filter(
    (fight) => fight.kind !== 'encounter' && (showIdle || fight.mob_deaths !== 0 || fight.index === selected),
  );
  const trashGroup = `Trash · ${plural(trash.length, 'pull')}`;
  // A pull indexed before engine 0.6.0 has no true kill count: its row says when, never a count.
  const trashRows = trash.map((fight) => ({
    value: String(fight.index),
    label: fightLabel(fight),
    description:
      fight.mob_deaths === undefined
        ? formatClock(fight.start)
        : `${formatClock(fight.start)} · ${plural(fight.mob_deaths, 'mob')} died`,
    group: trashGroup,
  }));
  return [...bossRows, ...trashRows];
}

/** The head's answer (decision R1 as ruled). A trash pull leads with its mobs and length, never a DPS. */
export interface ReportAnswer {
  label: string;
  value: string;
  unit: string;
  caption: string;
  /** The value's tone: strong on trash and the group's DPS, gold for the viewer's own. */
  tone: 'strong' | 'gold';
}

/** A per-second figure: one decimal under 100 (1.8), whole from 100 (1,374). */
export function formatRate(perSecond: number): string {
  const safe = Number.isFinite(perSecond) ? perSecond : 0;
  return safe < 100
    ? safe.toLocaleString('en-US', { minimumFractionDigits: 1, maximumFractionDigits: 1 })
    : Math.round(safe).toLocaleString('en-US');
}

export interface AnswerFacts {
  fight: FightEntry;
  attempt?: number;
  /** "Your group", "Group" or "Everyone in range". */
  scope: string;
  players: number;
  deaths: number;
  /** The scope's damage per second over the pull (boss pulls). */
  groupDps: number;
}

export function reportAnswer(facts: AnswerFacts): ReportAnswer {
  const { fight } = facts;
  if (fight.kind !== 'encounter') {
    const mobs = fight.mob_deaths ?? 0;
    return {
      label: `Trash pull · started ${formatClock(fight.start)}`,
      value: mobs.toLocaleString('en-US'),
      unit: `${mobs === 1 ? 'mob' : 'mobs'} · ${pullLength(fight.duration_ms)}`,
      caption: joinFacts(
        facts.scope,
        plural(facts.players, 'player'),
        facts.deaths === 0 ? 'no deaths' : plural(facts.deaths, 'death'),
      ),
      tone: 'strong',
    };
  }
  // The viewer's own DPS and parse lead once a boss pull with their character is reachable (G14).
  const outcome = fight.in_progress
    ? 'Live'
    : fight.kill
      ? 'Kill'
      : `Wipe${fight.boss_health_pct === undefined ? '' : ` at ${Math.round(fight.boss_health_pct)}%`}`;
  return {
    label: joinFacts(`${fight.name} #${facts.attempt ?? 1}`, outcome, pullLength(fight.duration_ms)),
    value: formatRate(facts.groupDps),
    unit: 'DPS',
    caption: `${facts.scope} damage · ${plural(facts.deaths, 'death')}`,
    tone: 'strong',
  };
}

export interface PullFacts {
  fight: FightEntry;
  players: number;
  deaths: number;
  phases: readonly { name: string; start_ms: number; end_ms: number }[];
}

/** The rail's This pull (decision R7 as ruled): trash reads Mobs died, never "Result: Trash". */
export function thisPullFacts({ fight, players, deaths, phases }: PullFacts): Fact[] {
  const trash = fight.kind !== 'encounter';
  const facts: Fact[] = [{ id: 'pull', label: 'Pull', value: fight.name }];
  if (!trash)
    facts.push({
      id: 'result',
      label: 'Result',
      value: fight.in_progress ? 'Live' : fight.kill ? 'Kill' : outcomeLabel(fight),
    });
  facts.push({ id: 'length', label: 'Length', value: pullLength(fight.duration_ms) });
  if (trash)
    facts.push({ id: 'mobs', label: 'Mobs died', value: (fight.mob_deaths ?? 0).toLocaleString('en-US') });
  facts.push(
    { id: 'players', label: 'Players', value: players.toLocaleString('en-US') },
    { id: 'deaths', label: 'Deaths', value: String(deaths) },
  );
  if (phases.length > 0)
    facts.push({
      id: 'phases',
      label: 'Phases',
      value: phases
        .map((phase) => `${phase.name} ${formatDuration(phase.start_ms)}–${formatDuration(phase.end_ms)}`)
        .join(' · '),
    });
  return facts;
}
