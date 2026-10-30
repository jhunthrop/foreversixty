// web/src/lib/chart/value-axis.ts
// The y axis of every report chart (controls and charts spec §B5.4 and §B6.5): a "nice" top tick,
// 4 or 5 gridlines from zero, short labels, and the log scale's powers of ten.

const NICE_STEPS = [1, 1.5, 2, 2.5, 3, 4, 5, 10] as const;

/** The log scale's floor: a zero is drawn there. */
export const LOG_FLOOR = 1;

/** The smallest of 1, 1.5, 2, 2.5, 3, 4, 5 or 10 times a power of ten that holds `peak`. */
export function niceTop(peak: number): number {
  if (!Number.isFinite(peak) || peak <= 0) return 1;
  const power = 10 ** Math.floor(Math.log10(peak));
  const factor = NICE_STEPS.find((step) => step * power >= peak - 1e-9) ?? 10;
  return factor * power;
}

/** True when `step` is 1, 2, 2.5, 3, 4 or 5 times a power of ten: a label that reads round. */
function isRoundStep(step: number): boolean {
  if (step <= 0) return false;
  const mantissa = step / 10 ** Math.floor(Math.log10(step));
  return [1, 2, 2.5, 3, 4, 5].some((nice) => Math.abs(mantissa - nice) < 1e-9);
}

/** Zero to `top` in 4, 3 or 5 equal steps, the first count whose step reads round. */
export function valueTicks(top: number): number[] {
  const steps = [4, 3, 5].find((count) => isRoundStep(top / count)) ?? 4;
  return Array.from({ length: steps + 1 }, (_, index) => (top / steps) * index);
}

/** `0`, `400`, `1.2k`, `60k`, `1.2M`. */
export function valueLabel(value: number): string {
  const abs = Math.abs(value);
  const trim = (n: number): string => n.toFixed(1).replace(/\.0$/, '');
  if (abs >= 1_000_000) return `${trim(value / 1_000_000)}M`;
  if (abs >= 1000) return `${trim(value / 1000)}k`;
  return String(Math.round(value));
}

/** Powers of ten from the floor to the largest at or under `peak`. */
export function logTicks(peak: number): number[] {
  const ticks: number[] = [];
  for (let at = LOG_FLOOR; at <= Math.max(peak, LOG_FLOOR); at *= 10) ticks.push(at);
  return ticks;
}
