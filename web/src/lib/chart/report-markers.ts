// web/src/lib/chart/report-markers.ts
// A fight's markers from its summary (controls and charts spec §B5.6): every death with its
// killing blow and its Deaths-tab link; the raid cooldowns `isRaidCooldown` names and the
// life-saving casts (by spell id); on the threat chart the taunts and the boss target changes.
import { isRaidCooldown } from '../report/raid-cooldowns';
import type { CastRow, Death, Taunt, TargetChange } from '../report/types';
import { castEvent, deathEvent, targetChangeEvent, tauntEvent } from './events';
import { markedCast } from './marked-casts';
import type { ChartEvent } from './types';

/** The Deaths tab's key for a death: the one DeathsTab opens and scrolls to. */
export function deathKey(death: Pick<Death, 'guid' | 'at_ms'>): string {
  return `${death.guid}-${death.at_ms}`;
}

export interface MarkerInput {
  deaths: readonly Death[];
  casts?: readonly CastRow[];
  /** One player as the Source: only their casts, and only their deaths. */
  sourcePlayer?: string;
  deathHref: (death: Death) => string;
  /** A spell's icon url, when the build has one. */
  iconOf?: (spellId: number, stem?: string) => string | null;
}

export function rateMarkers({
  deaths,
  casts = [],
  sourcePlayer,
  deathHref,
  iconOf,
}: MarkerInput): ChartEvent[] {
  const shownDeaths =
    sourcePlayer === undefined ? deaths : deaths.filter((death) => death.guid === sourcePlayer);
  const castEvents = casts
    .filter(
      (row) => sourcePlayer === undefined || row.guid === sourcePlayer || row.owner_guid === sourcePlayer,
    )
    .flatMap((row) => {
      const marked = markedCast(row.spell_id);
      if (marked === undefined && !isRaidCooldown(row.spell_name)) return [];
      const label = marked?.label ?? row.spell_name;
      return row.sequence.map((at) =>
        castEvent(at, label, row.name, row.spell_id, iconOf?.(row.spell_id, marked?.icon) ?? null),
      );
    });
  return [...shownDeaths.map((death) => deathEvent(death, deathHref(death))), ...castEvents].sort(
    (a, b) => a.atMs - b.atMs,
  );
}

export interface ThreatMarkerInput {
  deaths: readonly Death[];
  taunts: readonly Taunt[];
  targetChanges?: readonly TargetChange[];
  enemyGuids: readonly string[];
  deathHref: (death: Death) => string;
  iconOf?: (spellId: number, stem?: string) => string | null;
}

/** The threat chart's markers: deaths, the taunts on this enemy, and its target changes when the summary has them. */
export function threatMarkers(input: ThreatMarkerInput): ChartEvent[] {
  const enemies = new Set(input.enemyGuids);
  return [
    ...input.deaths.map((death) => deathEvent(death, input.deathHref(death))),
    ...input.taunts
      .filter((taunt) => enemies.has(taunt.target_guid))
      .map((taunt) => tauntEvent(taunt, input.iconOf?.(taunt.spell_id) ?? null)),
    ...(input.targetChanges ?? []).filter((change) => enemies.has(change.enemy_guid)).map(targetChangeEvent),
  ].sort((a, b) => a.atMs - b.atMs);
}
