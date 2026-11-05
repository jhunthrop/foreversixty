// web/src/lib/report/ability-table.ts
// Damage by ability as data (spec 2026-11-04-sim-results 3.6): one row per damaging ability with
// the figures a player reads. Pure, so every number is pinned against the fixture by hand.
//
// The adapter's counts (sim/adapter/adapter.go ability()): `hits` already includes crits and
// glances, and `misses` holds MISS, DODGE, PARRY, BLOCK and GLANCING, so a glance and a block are
// in both. Hence:
//   attempts = hits + MISS + DODGE + PARRY
//   landed   = hits
// Crit is crits over landed with glances counted (build ruling R7, the Warcraft Logs convention).
import { attackHand, parseActionKey, resolveActionName, type ActionNames, type AttackHand } from '../sim/action-names';
import { damageCopy } from '../sim/result-copy';
import type { Ability, Actor, AuraTrack, CastRow } from './types';

/** The outcomes the sub-line prints, in the order it prints them. */
export const OUTCOME_ORDER = ['MISS', 'DODGE', 'PARRY', 'GLANCING', 'BLOCK'] as const;
export type OutcomeKind = (typeof OUTCOME_ORDER)[number];

export interface OutcomeShare {
  kind: OutcomeKind;
  /** Share of attempts, 0..100. */
  pct: number;
}

export interface AbilityRow {
  /** The action key: the row's identity and the source of its name. */
  key: string;
  /** The client spell id, 0 for a white swing. */
  spellId: number;
  name: string;
  /** The hand of a white swing row, null otherwise. */
  hand: AttackHand | null;
  periodic: boolean;
  /** The pet's name when the row belongs to one, null for the player's own. */
  pet: string | null;
  school: number | undefined;
  total: number;
  dps: number;
  /** Share of the table's total, 0..100. */
  share: number;
  /** Null on a periodic row. A white row counts swings. */
  casts: number | null;
  attempts: number;
  landed: number;
  avgHit: number;
  /** Crits over landed (ticks for a periodic row), 0..100; null with nothing landed. */
  crit: number | null;
  biggestHit: number;
  /** The matching target debuff's uptime, 0..100; null when none is on record. */
  uptime: number | null;
  ticks: number;
  outcomes: OutcomeShare[];
}

export type SortColumn = 'dps' | 'share' | 'casts' | 'avgHit' | 'crit' | 'biggestHit';
export type SortDirection = 'desc' | 'asc';

export interface AbilitySort {
  column: SortColumn;
  direction: SortDirection;
}

export const DEFAULT_SORT: AbilitySort = { column: 'dps', direction: 'desc' };

export interface AbilityTableInput {
  actors: readonly Actor[];
  casts: readonly CastRow[];
  auras: readonly AuraTrack[];
  durationMs: number;
  names: ActionNames | null;
}

export interface AbilityTable {
  rows: AbilityRow[];
  total: number;
  dps: number;
  hasBiggestHit: boolean;
  hasUptime: boolean;
}

const missCount = (ability: Ability, kind: string): number => ability.misses?.[kind] ?? 0;

/** Every swing or press the engine rolled: `hits` holds glances and blocks already. */
export function attemptsOf(ability: Ability): number {
  return ability.hits + missCount(ability, 'MISS') + missCount(ability, 'DODGE') + missCount(ability, 'PARRY');
}

function outcomesOf(ability: Ability, attempts: number): OutcomeShare[] {
  if (attempts === 0) return [];
  return OUTCOME_ORDER.flatMap((kind) => {
    const count = missCount(ability, kind);
    return count > 0 ? [{ kind, pct: (count / attempts) * 100 }] : [];
  });
}

function buildRow(
  ability: Ability,
  pet: string | null,
  { casts, auras, durationMs, names }: AbilityTableInput,
  grandTotal: number,
): AbilityRow {
  const parsed = parseActionKey(ability.name);
  const hand = parsed?.kind === 'other' && parsed.label === 'attack' ? attackHand(parsed.tag) : null;
  const periodic = ability.ticks > 0 && ability.hits === 0;
  const landed = periodic ? ability.ticks : ability.hits + ability.ticks;
  const attempts = attemptsOf(ability);
  const spellId = parsed?.kind === 'spell' ? parsed.id : 0;
  const castRow = casts.find((row) => row.spell_name === ability.name);
  const track = periodic
    ? auras.find((row) => row.type === 'DEBUFF' && row.spell_id === ability.spell_id)
    : undefined;
  return {
    key: ability.name,
    spellId,
    name: hand === null ? resolveActionName(ability.name, names) : damageCopy.hands[hand],
    hand,
    periodic,
    pet,
    school: ability.school,
    total: ability.total,
    dps: durationMs > 0 ? ability.total / (durationMs / 1000) : 0,
    share: grandTotal > 0 ? (ability.total / grandTotal) * 100 : 0,
    casts: periodic ? null : hand !== null ? attempts : (castRow?.succeeded ?? attempts),
    attempts,
    landed,
    avgHit: landed > 0 ? ability.total / landed : 0,
    crit: landed > 0 ? (ability.crits / landed) * 100 : null,
    biggestHit: ability.max,
    uptime: track !== undefined && durationMs > 0 ? Math.min((track.uptime_ms / durationMs) * 100, 100) : null,
    ticks: ability.ticks,
    outcomes: outcomesOf(ability, attempts),
  };
}

/** The table for a result's damage rows. Only abilities that dealt damage get a row. */
export function abilityTable(input: AbilityTableInput): AbilityTable {
  const [player, ...pets] = input.actors;
  if (player === undefined) return { rows: [], total: 0, dps: 0, hasBiggestHit: false, hasUptime: false };
  const dealt = (actor: Actor): Ability[] => actor.abilities.filter((ability) => ability.total > 0);
  const total = input.actors.flatMap(dealt).reduce((sum, ability) => sum + ability.total, 0);
  const rows = [
    ...dealt(player).map((ability) => buildRow(ability, null, input, total)),
    ...pets.flatMap((pet) => dealt(pet).map((ability) => buildRow(ability, pet.name, input, total))),
  ];
  return {
    rows,
    total,
    dps: input.durationMs > 0 ? total / (input.durationMs / 1000) : 0,
    hasBiggestHit: rows.some((row) => row.biggestHit > 0),
    hasUptime: rows.some((row) => row.uptime !== null),
  };
}

const SORT_VALUE: Readonly<Record<SortColumn, (row: AbilityRow) => number>> = {
  dps: (row) => row.dps,
  share: (row) => row.share,
  casts: (row) => row.casts ?? -1,
  avgHit: (row) => row.avgHit,
  crit: (row) => row.crit ?? -1,
  biggestHit: (row) => row.biggestHit,
};

/** Rows in the chosen order; a tie keeps the table's own order, so a sort never reshuffles it. */
export function sortAbilityRows(rows: readonly AbilityRow[], { column, direction }: AbilitySort): AbilityRow[] {
  const value = SORT_VALUE[column];
  const sign = direction === 'desc' ? -1 : 1;
  return rows
    .map((row, index) => ({ row, index }))
    .sort((a, b) => sign * (value(a.row) - value(b.row)) || a.index - b.index)
    .map(({ row }) => row);
}

/** A first press sorts descending, a second press ascending. */
export function nextSort(current: AbilitySort, column: SortColumn): AbilitySort {
  if (current.column !== column) return { column, direction: 'desc' };
  return { column, direction: current.direction === 'desc' ? 'asc' : 'desc' };
}
