// web/src/lib/guild/output-figure.ts
// One output figure, written one way on every guild surface (guild round 2, WP-13). To a Classic
// player a "parse" is a percentile, so a raw HPS or DPS figure is never called one: the Raids tab
// reads `Top HPS`, Progression `Best HPS` per role, the roster chip `Best HPS 465`. The number is
// rounded and grouped by the one formatter, so the same kill reads the same digits on all three
// (round 1 printed 464.6 on two tabs and 465 on the third). The API picks the figure by one rule
// (api/internal/guilds/parse_source.go: boss kills only, the role's own metric).
import { formatAmount } from '../report/format';

/** `HPS` or `DPS`, from the API's `metric`. */
export function outputUnit(metric: string): string {
  return metric.toUpperCase();
}

/** `465`, `1,204`: whole, grouped. */
export function outputFigure(value: number): string {
  return formatAmount(Math.round(value));
}

/** The Raids tab's night line label: `Top HPS`, `Top DPS`. */
export function topOutputLabel(metric: string): string {
  return `Top ${outputUnit(metric)}`;
}

/** Progression's per-role label: `Best HPS`, `Best DPS`, and a tank's `Best tank DPS`. */
export function bestOutputLabel(role: string, metric: string): string {
  return role === 'tank' ? `Best tank ${outputUnit(metric)}` : `Best ${outputUnit(metric)}`;
}
