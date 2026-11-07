// web/src/lib/ui/countdown.ts
// The countdown panel's clock (ui/Countdown.svelte, home B spec 2026-11-07 §1): days, hours and
// minutes to a known instant, seconds only in the final hour (they add nothing before it and
// repaint every second). Pure; the panel ticks it.

export interface ClockPart {
  value: string;
  unit: 'days' | 'hours' | 'min' | 'sec';
}

const SECOND = 1000;
const MINUTE = 60 * SECOND;
const HOUR = 60 * MINUTE;
const DAY = 24 * HOUR;

const two = (n: number) => String(n).padStart(2, '0');

/** The clock's parts, or null once the instant has passed. */
export function clockParts(target: Date, now: Date): ClockPart[] | null {
  const left = target.getTime() - now.getTime();
  if (left <= 0) return null;
  if (left < HOUR) {
    return [
      { value: two(Math.floor(left / MINUTE)), unit: 'min' },
      { value: two(Math.floor((left % MINUTE) / SECOND)), unit: 'sec' },
    ];
  }
  return [
    { value: two(Math.floor(left / DAY)), unit: 'days' },
    { value: two(Math.floor((left % DAY) / HOUR)), unit: 'hours' },
    { value: two(Math.floor((left % HOUR) / MINUTE)), unit: 'min' },
  ];
}

/** How long until the clock next changes: a second in the final hour, else to the next minute. */
export function nextTickMs(target: Date, now: Date): number {
  const left = target.getTime() - now.getTime();
  if (left <= 0) return MINUTE;
  if (left < HOUR) return SECOND;
  return left % MINUTE || MINUTE;
}

/** The clock read aloud: "33 days, 4 hours, 12 min". */
export function clockLabel(parts: readonly ClockPart[]): string {
  return parts.map((part) => `${Number(part.value)} ${part.unit}`).join(', ');
}
