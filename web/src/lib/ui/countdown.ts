// web/src/lib/ui/countdown.ts
// The one countdown (owner ruling 2026-11-10; apps/shared/formats.json `countdown`): days, hours
// and minutes to a known instant, the largest non-zero unit first and at most three, minutes
// rounded up so the clock never reads 0 before the instant. No seconds: every countdown on the
// site steps on the one minute tick (minute-clock.svelte.ts). Pure; the surfaces tick it.

import { DAY_MS, HOUR_MS, MINUTE_MS } from '../format/duration';

export type CountdownUnit = 'day' | 'days' | 'hour' | 'hours' | 'min';

export interface ClockPart {
  value: string;
  unit: CountdownUnit;
}

const MINUTES_PER_HOUR = HOUR_MS / MINUTE_MS;
const MINUTES_PER_DAY = DAY_MS / MINUTE_MS;

function part(n: number, one: 'day' | 'hour' | 'min', many: 'days' | 'hours' | 'min'): ClockPart {
  return { value: String(n), unit: n === 1 ? one : many };
}

/** The clock's parts; empty once the instant has passed. */
export function clockParts(target: Date, now: Date): ClockPart[] {
  const left = target.getTime() - now.getTime();
  if (!(left > 0)) return [];
  const totalMinutes = Math.ceil(left / MINUTE_MS);
  const days = Math.floor(totalMinutes / MINUTES_PER_DAY);
  const hours = Math.floor((totalMinutes % MINUTES_PER_DAY) / MINUTES_PER_HOUR);
  const minutes = totalMinutes % MINUTES_PER_HOUR;
  const min = part(minutes, 'min', 'min');
  if (days > 0) return [part(days, 'day', 'days'), part(hours, 'hour', 'hours'), min];
  if (hours > 0) return [part(hours, 'hour', 'hours'), min];
  return [min];
}

/** The clock read aloud, and its inline text form: "32 days, 5 hours, 12 min". */
export function clockLabel(parts: readonly ClockPart[]): string {
  return parts.map((entry) => `${entry.value} ${entry.unit}`).join(', ');
}

/** The countdown to `target` as text: "32 days, 5 hours, 12 min"; empty once it has passed. */
export function countdownText(target: Date, now: Date): string {
  return clockLabel(clockParts(target, now));
}
