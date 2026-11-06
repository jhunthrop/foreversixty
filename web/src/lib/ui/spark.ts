// web/src/lib/ui/spark.ts
// The Spark primitive's geometry (logs spec 2026-11-07 section 3A): a series drawn as one line
// across a 1000 x 24 box, or a strip of ticks at positions 0 to 1 (a cast sequence).
const WIDTH = 1000;
const BASELINE = 22;
const RISE = 20;
const TICK_W = 4;

/** The polyline's points, each value over `top` (the series' own largest by default). */
export function sparkPoints(values: readonly number[], top?: number): string {
  const ceiling = top ?? Math.max(1, ...values);
  const step = WIDTH / Math.max(values.length - 1, 1);
  return values
    .map((value, index) => `${(index * step).toFixed(1)},${(BASELINE - (value / ceiling) * RISE).toFixed(1)}`)
    .join(' ');
}

/** A tick's x for a position 0 to 1. */
export function sparkTickX(position: number): number {
  return Math.round(Math.min(1, Math.max(0, position)) * (WIDTH - TICK_W) * 10) / 10;
}
