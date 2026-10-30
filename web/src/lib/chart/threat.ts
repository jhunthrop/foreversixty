// web/src/lib/chart/threat.ts
// The estimated-threat chart's model (controls and charts spec §B5.3, §B12): one cumulative line
// per player on the charted enemy, the top 8 by standing at the window's end, and always the
// tanks, everyone who taunted it in the window and everyone it attacked in the window.
import { splitUnitName } from '../characters';
import { classColorVar } from '../report/format';
import type { Actor, Taunt, TargetChange, ThreatPair } from '../report/types';
import { BUCKET_MS, sliceSeries, type TimeWindow } from '../report/window';
import { cumulative, selectDrawn, standingAt } from './drawn';
import type { ChartSeries, PlayerEntry } from './types';

export interface ThreatPlayer {
  name: string;
  className?: string;
  role: PlayerEntry['role'];
}

export interface ThreatInput {
  /** Whole-fight pairs: each player's threat on each enemy, one value per second. */
  pairs: readonly ThreatPair[];
  /** Every GUID folded under the charted enemy's name. */
  enemyGuids: readonly string[];
  players: ReadonlyMap<string, ThreatPlayer>;
  window: TimeWindow;
  taunts: readonly Taunt[];
  /** Absent until the engine infers them (spec §B12). */
  targetChanges?: readonly TargetChange[];
  /** Damage taken, whole fight: who the enemy hit, when the summary carries no target changes. */
  damageTaken?: readonly Actor[];
  /** One player as the Source: that player and the always-drawn ones only. */
  sourcePlayer?: string;
  youGuid?: string;
  chosen: ReadonlySet<string> | null;
}

export interface ThreatModel {
  series: ChartSeries[];
  entries: PlayerEntry[];
  /** Every player's estimated threat on the enemy, per second: a tooltip share's denominator. */
  totals: number[];
  total: number;
  isDefault: boolean;
}

const inWindow = (atMs: number, window: TimeWindow): boolean =>
  atMs >= window.startMs && atMs <= window.endMs;

/**
 * The players the enemy attacked in the window: from its inferred target changes when the summary
 * has them (including the one it was on as the window opened), else every player whose damage
 * taken in the window includes a hit from it.
 */
export function attackedPlayers(
  enemyGuids: readonly string[],
  window: TimeWindow,
  targetChanges: readonly TargetChange[] | undefined,
  damageTaken: readonly Actor[] = [],
): Set<string> {
  const enemies = new Set(enemyGuids);
  if (targetChanges !== undefined) {
    const mine = targetChanges
      .filter((change) => enemies.has(change.enemy_guid))
      .sort((a, b) => a.at_ms - b.at_ms);
    const before = mine.filter((change) => change.at_ms < window.startMs).at(-1);
    return new Set([
      ...(before === undefined ? [] : [before.to_guid]),
      ...mine.filter((change) => inWindow(change.at_ms, window)).map((change) => change.to_guid),
    ]);
  }
  return new Set(
    damageTaken
      .filter((actor) => actor.targets.some((pair) => enemies.has(pair.guid)))
      .filter((actor) => sliceSeries(actor.series, window).some((value) => value > 0))
      .map((actor) => actor.guid),
  );
}

/** Each player's summed series on the enemy (six adds of one name are one enemy). */
function seriesByPlayer(
  pairs: readonly ThreatPair[],
  enemyGuids: readonly string[],
  players: ReadonlyMap<string, ThreatPlayer>,
): Map<string, number[]> {
  const enemies = new Set(enemyGuids);
  const out = new Map<string, number[]>();
  for (const pair of pairs) {
    if (pair.series === undefined || !enemies.has(pair.target_guid) || !players.has(pair.guid)) continue;
    const found = out.get(pair.guid) ?? [];
    out.set(
      pair.guid,
      Array.from(
        { length: Math.max(found.length, pair.series.length) },
        (_, index) => (found[index] ?? 0) + (pair.series?.[index] ?? 0),
      ),
    );
  }
  return out;
}

export function buildThreatModel(input: ThreatInput): ThreatModel {
  const { players, window } = input;
  const running = new Map(
    [...seriesByPlayer(input.pairs, input.enemyGuids, players)].map(([guid, series]) => [
      guid,
      cumulative(series),
    ]),
  );
  const figures = new Map([...running].map(([guid, line]) => [guid, standingAt(line, window)]));
  const enemies = new Set(input.enemyGuids);
  const taunters = input.taunts
    .filter((taunt) => enemies.has(taunt.target_guid) && inWindow(taunt.at_ms, window))
    .map((taunt) => taunt.source_guid);
  const tanks = [...players].filter(([, player]) => player.role === 'tank').map(([guid]) => guid);
  const attacked = attackedPlayers(input.enemyGuids, window, input.targetChanges, input.damageTaken);
  const pinned = new Set(
    [...tanks, ...taunters, ...attacked, input.sourcePlayer, input.youGuid].filter(
      (guid): guid is string => guid !== undefined && running.has(guid),
    ),
  );
  const chosen = input.sourcePlayer === undefined ? input.chosen : (input.chosen ?? new Set<string>());
  const drawnIds = selectDrawn({
    candidates: [...figures].map(([id, figure]) => ({ id, figure })),
    pinned,
    chosen,
  });
  const drawn = new Set(drawnIds);
  const length = [...running.values()].reduce((longest, line) => Math.max(longest, line.length), 0);
  const totals = Array.from({ length }, (_, index) =>
    [...running.values()].reduce((sum, line) => sum + (line[index] ?? line.at(-1) ?? 0), 0),
  );
  const nameOf = (guid: string): string => splitUnitName(players.get(guid)?.name ?? guid).name;
  const series: ChartSeries[] = drawnIds.map((guid) => ({
    id: guid,
    label: nameOf(guid),
    colour: classColorVar(players.get(guid)?.className),
    kind: 'player',
    values: running.get(guid) ?? [],
    figure: figures.get(guid) ?? 0,
    pinned: pinned.has(guid),
    you: guid === input.youGuid,
  }));
  const entries: PlayerEntry[] = [...figures]
    .sort((a, b) => b[1] - a[1])
    .map(([guid, figure]) => ({
      id: guid,
      label: nameOf(guid),
      colour: classColorVar(players.get(guid)?.className),
      className: players.get(guid)?.className,
      role: players.get(guid)?.role ?? 'other',
      figure,
      drawn: drawn.has(guid),
      pinned: pinned.has(guid),
      you: guid === input.youGuid,
    }));
  return { series, entries, totals, total: running.size, isDefault: input.chosen === null };
}

/** A player's share of every player's estimated threat on the enemy at a second, 0 to 1. */
export function shareAt(model: Pick<ThreatModel, 'totals'>, value: number, ms: number): number {
  const index = Math.min(model.totals.length - 1, Math.floor(ms / BUCKET_MS));
  const total = index < 0 ? 0 : model.totals[index];
  return total <= 0 ? 0 : value / total;
}
