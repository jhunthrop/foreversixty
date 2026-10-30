// web/src/lib/chart/tooltip.ts
// The tooltip's rows at one second (controls and charts spec §B5.7): highest value first, the
// legend-hovered series always in, players who died before this second last as "Died m:ss", and
// at most 12 rows then a count.
import { chartCopy } from './copy';
import type { ChartSeries } from './types';

export const TOOLTIP_ROWS = 12;

export interface TooltipLine {
  id: string;
  label: string;
  colour: string;
  /** The value at the second; null for a dead player. */
  value: number | null;
  valueText: string;
  dead: boolean;
  highlighted: boolean;
}

export function tooltipLines(
  series: readonly ChartSeries[],
  ms: number,
  options: {
    hidden: ReadonlySet<string>;
    highlightId: string | null;
    format: (value: number) => string;
  },
): { lines: TooltipLine[]; more: number } {
  const index = Math.floor(ms / 1000);
  const all: TooltipLine[] = series
    .filter((line) => !options.hidden.has(line.id))
    .map((line) => {
      const raw = line.values[index] ?? null;
      const dead = raw === null && line.diedAtMs !== undefined && line.diedAtMs <= ms;
      return {
        id: line.id,
        label: line.label,
        colour: line.colour,
        value: raw,
        valueText: dead ? chartCopy.diedAt(line.diedAtMs!) : options.format(raw ?? 0),
        dead,
        highlighted: line.id === options.highlightId,
      };
    });
  const living = all.filter((line) => !line.dead).sort((a, b) => (b.value ?? 0) - (a.value ?? 0));
  const ordered = [...living, ...all.filter((line) => line.dead)];
  const shown = ordered.slice(0, TOOLTIP_ROWS);
  const pinned = ordered.find((line) => line.highlighted);
  if (pinned !== undefined && !shown.includes(pinned)) shown[shown.length - 1] = pinned;
  return { lines: shown, more: Math.max(0, ordered.length - TOOLTIP_ROWS) };
}
