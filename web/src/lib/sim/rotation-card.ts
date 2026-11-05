// web/src/lib/sim/rotation-card.ts
// The Rotation card's rows (spec 2026-11-04-sim-results 3.11): the curated APL's lines against
// this run's casts. A cast counts by the spell's name, so a line written at one rank finds a run
// that cast another. Tagged rows and `other:` rows are not button presses: a white swing, a
// periodic tick and a queue track are skipped.
import { parseActionKey, resolveActionName, type ActionNames } from './action-names';
import { rotationCardCopy as copy } from './input-cards-copy';
import type { RotationLines } from './rotations';
import type { SimResult } from './types';

export interface PrepullRow {
  key: string;
  action: string;
  name: string;
  /** "-2.5 s". */
  at: string;
  /** Null when the run kept no sample, so nothing can be said either way. */
  cast: boolean | null;
}

export interface PriorityRow {
  key: string;
  number: number;
  /** The key the icon, rank and tooltip read: this run's cast of the line where it has one. */
  action: string | null;
  name: string;
  note: string;
  /** Casts per fight; null on a later line of a spell an earlier line already counted. */
  casts: number | null;
  dimmed: boolean;
}

export interface RotationRows {
  prepull: PrepullRow[];
  priority: PriorityRow[];
}

interface Pressed {
  key: string;
  name: string;
  casts: number;
}

function pressedInRun(result: SimResult, names: ActionNames | null): Pressed[] {
  const byName = new Map<string, Pressed>();
  for (const row of result.summary.casts) {
    const parsed = parseActionKey(row.spell_name);
    if (parsed === null || parsed.kind === 'other' || parsed.kind === 'unknown' || parsed.tag !== 0) continue;
    const key = `${parsed.kind}:${parsed.label}`;
    const name = resolveActionName(key, names);
    const found = byName.get(name);
    byName.set(name, { key: found?.key ?? key, name, casts: (found?.casts ?? 0) + row.succeeded });
  }
  return [...byName.values()];
}

function castBeforePull(result: SimResult, names: ActionNames | null): ReadonlySet<string> | null {
  if (result.sample === undefined) return null;
  return new Set(
    result.sample
      .filter((row) => row.at_ms < 0 && row.outcome !== 'tick' && row.outcome !== 'tick_crit')
      .map((row) => {
        const parsed = parseActionKey(row.action);
        return parsed === null ? row.action : resolveActionName(`${parsed.kind}:${parsed.label}`, names);
      }),
  );
}

export function rotationRows(
  result: SimResult,
  lines: RotationLines,
  names: ActionNames | null,
): RotationRows {
  const pressed = pressedInRun(result, names);
  const pressedByName = new Map(pressed.map((row) => [row.name, row]));
  const named = new Set(
    [...lines.prepull, ...lines.priority].flatMap((line) =>
      line.action === null ? [] : [resolveActionName(line.action, names)],
    ),
  );
  const others = pressed.filter((row) => !named.has(row.name) && row.casts > 0);
  const before = castBeforePull(result, names);
  const counted = new Set<string>();

  const priority = lines.priority.map((line, at): PriorityRow => {
    const key = `line-${at}`;
    if (line.action === null) {
      const total = others.reduce((sum, row) => sum + row.casts, 0);
      return {
        key,
        number: at + 1,
        action: others[0]?.key ?? null,
        name: copy.otherCooldowns,
        note: others.length === 0 ? '' : copy.thisRun(others.map((row) => row.name).join(', ')),
        casts: Math.round(total),
        dimmed: total === 0,
      };
    }
    const name = resolveActionName(line.action, names);
    const run = pressedByName.get(name);
    const repeat = counted.has(name);
    counted.add(name);
    return {
      key,
      number: at + 1,
      action: run?.key ?? line.action,
      name,
      note: line.note,
      casts: repeat ? null : Math.round(run?.casts ?? 0),
      dimmed: Math.round(run?.casts ?? 0) === 0,
    };
  });

  return {
    prepull: lines.prepull.flatMap((line, at): PrepullRow[] =>
      line.action === null
        ? []
        : [
            {
              key: `prepull-${at}`,
              action: line.action,
              name: resolveActionName(line.action, names),
              at: copy.prepullAt(line.at_seconds),
              cast: before === null ? null : before.has(resolveActionName(line.action, names)),
            },
          ],
    ),
    priority,
  };
}

/** The action keys the rows name, so the card can read the shard names they need. */
export function rotationActionKeys(lines: RotationLines): string[] {
  return [...lines.prepull, ...lines.priority].flatMap((line) => (line.action === null ? [] : [line.action]));
}
