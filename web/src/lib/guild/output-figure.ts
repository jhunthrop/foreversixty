// web/src/lib/guild/output-figure.ts
// One output figure, written one way on every guild surface (guild round 2, WP-13 and fix round).
// To a Classic player a "parse" is a percentile, so a raw HPS or DPS figure is never called one,
// and it carries one word everywhere: a raid night reads `Top HPS` and `Top DPS`, Progression
// `Top HPS` per role, the roster chip `Top HPS 465`. The number is rounded and grouped by the one
// formatter, so the same kill reads the same digits on all three. The API picks the figure by one
// rule (api/internal/guilds/parse_source.go: boss kills only, the role's own metric).
import { formatAmount } from '../report/format';
import type { GuildRaidRow, GuildRaidTopParse } from './api';

/** `HPS` or `DPS`, from the API's `metric`. */
export function outputUnit(metric: string): string {
  return metric.toUpperCase();
}

/** `465`, `1,204`: whole, grouped. */
export function outputFigure(value: number): string {
  return formatAmount(Math.round(value));
}

/** `Top HPS`, `Top DPS`; a tank's DPS on Progression reads `Top tank DPS`. */
export function topOutputLabel(metric: string, role?: string): string {
  return role === 'tank' ? `Top tank ${outputUnit(metric)}` : `Top ${outputUnit(metric)}`;
}

/** A raid night's figures in reading order, healing then damage; the API leaves out one it lacks. */
export function nightTops(row: Pick<GuildRaidRow, 'top_hps' | 'top_dps'>): GuildRaidTopParse[] {
  return [row.top_hps, row.top_dps].filter((top): top is GuildRaidTopParse => top != null);
}
