// web/src/lib/sim/compare.ts
// A sim beside the logged fight it was built from, per ability and per buff.
//
// The one hard part is the join, and it is hard because the two sides identify a row
// differently. sim/adapter gives a summary row a unique integer: the client spell id for a
// plain untagged spell, and a derived id at or above SYNTHETIC_ID_BASE for a tagged or
// ranked spell, an item or an "other" action. A combat log carries the client's own ids and
// the client's own names, and gives every melee swing spell id 0.
//
// So the key is two-tiered, in this order:
//
//   1. `id:<n>` when the row's id is a real client id (below SYNTHETIC_ID_BASE and non-zero).
//      That is the right key for every ordinary ability and the only one that works with no
//      name table loaded.
//   2. `name:<resolved, de-variant, aliased, case-folded>` otherwise -- the tagged, item and
//      "other" rows, whose ids are the adapter's own arithmetic, plus the log's id-zero
//      rows. The variant suffix is dropped so the sim's tagged melee rows fold onto
//      one, and simCopy.actionAliases turns the engine's "Attack" into the log's "Melee".
//
// Anything matched by neither is its own row with a zero on one side. That is the honest
// outcome: the table then says the sim did something the fight did not, which is visible,
// rather than folding two different actions together, which is not. In particular two
// id-zero rows on the log side stay apart, because tier 2 keys them by name.
import { SYNTHETIC_ID_BASE, resolveActionName, type ActionNames } from './action-names';
import { simCopy } from './copy';
import type { Summary } from '../report/types';
import { playerActor } from './sentence';

export interface AbilityDiff {
  /**
   * The row's own identity, and what the `{#each}` key and the test id are built from --
   * never the name, which is display text and can contain a space, a colon or a slash.
   * It is the sim side's `spell_id` for the key when there is one, else the fight side's:
   * both are stable, and join keys are unique, so no two rows can collide.
   */
  spellId: number;
  name: string;
  simCasts: number;
  actualCasts: number;
  simDamage: number;
  actualDamage: number;
}

export interface AuraDiff {
  name: string;
  simUptimeMs: number;
  actualUptimeMs: number;
}

export interface Comparison {
  headline: string;
  simDps: number;
  actualDps: number;
  /** actual / simulated, clamped to [0, 2] as the contract clamps the stored column. */
  executionScore: number | null;
  abilities: AbilityDiff[];
  auras: AuraDiff[];
  /** The differences in words, largest gap first: at most five abilities and three auras. */
  lines: string[];
}

export const MAX_ABILITY_LINES = 5;
export const MAX_AURA_LINES = 3;

/** "Heroic Strike (2, Rank 3)" is the same action as "Heroic Strike" for a join. */
const VARIANT = /\s*\((?:\d+(?:, Rank \d+)?|Rank \d+)\)$/;

/** One row's join key, and the name it should be shown under. */
interface Keyed {
  key: string;
  label: string;
}

function keyed(spellId: number, rawName: string, names: ActionNames | null): Keyed {
  const label = resolveActionName(rawName, names);
  if (spellId > 0 && spellId < SYNTHETIC_ID_BASE) return { key: `id:${spellId}`, label };
  const base = label.replace(VARIANT, '');
  const aliased = simCopy.actionAliases[base] ?? base;
  return { key: `name:${aliased.trim().toLowerCase()}`, label: aliased };
}

function dpsOf(summary: Summary, total: number): number {
  return summary.duration_ms <= 0 ? 0 : (total / summary.duration_ms) * 1000;
}

/** Successful casts per join key, for one actor. */
function castsByKey(summary: Summary, guid: string, names: ActionNames | null): Map<string, number> {
  const casts = new Map<string, number>();
  for (const row of summary.casts) {
    if (row.guid !== guid) continue;
    const { key } = keyed(row.spell_id, row.spell_name, names);
    casts.set(key, (casts.get(key) ?? 0) + row.succeeded);
  }
  return casts;
}

/** Damage per join key, and the label each key should be shown under. */
function damageByKey(
  abilities: readonly { spell_id: number; name: string; total: number }[],
  names: ActionNames | null,
): { damage: Map<string, number>; labels: Map<string, string>; ids: Map<string, number> } {
  const damage = new Map<string, number>();
  const labels = new Map<string, string>();
  const ids = new Map<string, number>();
  for (const ability of abilities) {
    const { key, label } = keyed(ability.spell_id, ability.name, names);
    damage.set(key, (damage.get(key) ?? 0) + ability.total);
    // First one wins: the sim side is read first, so a resolved engine name beats a raw key.
    if (!labels.has(key)) labels.set(key, label);
    if (!ids.has(key)) ids.set(key, ability.spell_id);
  }
  return { damage, labels, ids };
}

/** Buff uptime per join key, for one actor. */
function upByKey(summary: Summary, guid: string, names: ActionNames | null): Map<string, number> {
  const up = new Map<string, number>();
  for (const track of summary.auras) {
    if (track.target_guid !== guid || track.type !== 'BUFF') continue;
    const { key } = keyed(track.spell_id, track.name, names);
    up.set(key, Math.max(up.get(key) ?? 0, track.uptime_ms));
  }
  return up;
}

function labelsOf(summary: Summary, guid: string, names: ActionNames | null): Map<string, string> {
  const labels = new Map<string, string>();
  for (const track of summary.auras) {
    if (track.target_guid !== guid || track.type !== 'BUFF') continue;
    const { key, label } = keyed(track.spell_id, track.name, names);
    if (!labels.has(key)) labels.set(key, label);
  }
  return labels;
}

function pct(part: number, whole: number): number {
  return whole <= 0 ? 0 : Math.round((part / whole) * 100);
}

function round(value: number): string {
  return Math.round(value).toLocaleString('en-US');
}

/**
 * `actualPlayerName` names the character in the logged fight; the sim has exactly one
 * player, so its actor is whichever did the most damage. `names` is the build's name table
 * (Task 23) and may be null, in which case a sim row keeps its action key and only the
 * id-keyed rows pair up -- which is still most of the table.
 *
 * `sim` is the raw summary, not a `namedSummary()` copy: this function resolves names
 * itself, because it needs the raw `spell_id` for tier 1 and the resolved name for tier 2.
 */
export function compareSummaries(
  sim: Summary,
  actual: Summary,
  actualPlayerName: string,
  names: ActionNames | null,
): Comparison {
  const simActor = playerActor(sim);
  const actualActor = actual.damage_done.find((row) => row.name === actualPlayerName) ?? null;

  if (simActor === null || actualActor === null) {
    return {
      headline: simCopy.compareNoPlayer,
      simDps: 0,
      actualDps: 0,
      executionScore: null,
      abilities: [],
      auras: [],
      lines: [],
    };
  }

  const simDps = dpsOf(sim, simActor.total);
  const actualDps = dpsOf(actual, actualActor.total);
  const executionScore = simDps <= 0 ? null : Math.min(2, Math.max(0, actualDps / simDps));

  const simCasts = castsByKey(sim, simActor.guid, names);
  const actualCasts = castsByKey(actual, actualActor.guid, names);
  const simSide = damageByKey(simActor.abilities, names);
  const actualSide = damageByKey(actualActor.abilities, names);

  const abilityKeys = [...new Set([...simSide.damage.keys(), ...actualSide.damage.keys()])];
  const abilities: AbilityDiff[] = abilityKeys
    .map((key) => ({
      spellId: simSide.ids.get(key) ?? actualSide.ids.get(key) ?? 0,
      name: simSide.labels.get(key) ?? actualSide.labels.get(key) ?? key,
      simCasts: simCasts.get(key) ?? 0,
      actualCasts: actualCasts.get(key) ?? 0,
      simDamage: simSide.damage.get(key) ?? 0,
      actualDamage: actualSide.damage.get(key) ?? 0,
    }))
    .sort((a, b) => Math.max(b.simDamage, b.actualDamage) - Math.max(a.simDamage, a.actualDamage));

  const simUp = upByKey(sim, simActor.guid, names);
  const actualUp = upByKey(actual, actualActor.guid, names);
  const simAuraLabels = labelsOf(sim, simActor.guid, names);
  const actualAuraLabels = labelsOf(actual, actualActor.guid, names);
  const auras: AuraDiff[] = [...new Set([...simUp.keys(), ...actualUp.keys()])]
    .map((key) => ({
      name: simAuraLabels.get(key) ?? actualAuraLabels.get(key) ?? key,
      simUptimeMs: simUp.get(key) ?? 0,
      actualUptimeMs: actualUp.get(key) ?? 0,
    }))
    .sort((a, b) => Math.max(b.simUptimeMs, b.actualUptimeMs) - Math.max(a.simUptimeMs, a.actualUptimeMs));

  const castLines = [...abilities]
    .filter((row) => row.simCasts !== row.actualCasts)
    .sort((a, b) => Math.abs(b.simCasts - b.actualCasts) - Math.abs(a.simCasts - a.actualCasts))
    .slice(0, MAX_ABILITY_LINES)
    .map((row) => `${row.name} cast ${row.actualCasts} times, the sim expects ${row.simCasts}.`);

  const share = (ms: number, summary: Summary): number => ms / Math.max(summary.duration_ms, 1);
  const auraLines = [...auras]
    .filter((row) => pct(row.simUptimeMs, sim.duration_ms) !== pct(row.actualUptimeMs, actual.duration_ms))
    .sort(
      (a, b) =>
        Math.abs(share(b.simUptimeMs, sim) - share(b.actualUptimeMs, actual)) -
        Math.abs(share(a.simUptimeMs, sim) - share(a.actualUptimeMs, actual)),
    )
    .slice(0, MAX_AURA_LINES)
    .map(
      (row) =>
        `${row.name} uptime ${pct(row.actualUptimeMs, actual.duration_ms)}% against ${pct(row.simUptimeMs, sim.duration_ms)}%.`,
    );

  const headline =
    executionScore === null
      ? `This fight did ${round(actualDps)} DPS; the sim produced no damage to compare it with.`
      : `This fight did ${round(actualDps)} DPS; the sim expects ${round(simDps)}. That is ${Math.round(executionScore * 100)}% of what this gear can do.`;

  return {
    headline,
    simDps,
    actualDps,
    executionScore,
    abilities,
    auras,
    lines: [...castLines, ...auraLines],
  };
}
