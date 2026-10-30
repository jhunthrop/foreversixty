// web/src/lib/chart/time-axis.ts
// Every time axis on the site reads its labels and its tick step from here (controls and charts
// spec §B5.4): `m:ss` under an hour, `h:mm:ss` from one, and the smallest of a short list of
// steps that keeps two labels 64 px apart.

/** The steps a time axis may take, in seconds. */
export const TIME_STEPS_S = [5, 10, 15, 30, 60, 120, 300] as const;

/** The least room between two time labels, in px. */
export const MIN_LABEL_GAP_PX = 64;

/** A fight time on the chart's clock: `2:07`, or `1:02:05` past an hour. */
export function clockLabel(ms: number): string {
  const safe = Number.isFinite(ms) && ms > 0 ? ms : 0;
  const total = Math.floor(safe / 1000);
  const hours = Math.floor(total / 3600);
  const minutes = Math.floor((total % 3600) / 60);
  const seconds = String(total % 60).padStart(2, '0');
  return hours > 0 ? `${hours}:${String(minutes).padStart(2, '0')}:${seconds}` : `${minutes}:${seconds}`;
}

/** The smallest listed step whose labels sit at least 64 px apart across `widthPx`. */
export function timeStepSeconds(spanS: number, widthPx: number): number {
  const perSecond = spanS <= 0 ? 0 : widthPx / spanS;
  return TIME_STEPS_S.find((step) => step * perSecond >= MIN_LABEL_GAP_PX) ?? TIME_STEPS_S.at(-1)!;
}

/** The tick positions, in seconds, on whole steps between `startS` and `endS` inclusive. */
export function timeTicks(startS: number, endS: number, widthPx: number): number[] {
  if (endS <= startS) return [];
  const step = timeStepSeconds(endS - startS, widthPx);
  const ticks: number[] = [];
  for (let at = Math.ceil(startS / step) * step; at <= endS; at += step) ticks.push(at);
  return ticks;
}
