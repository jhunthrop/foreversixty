// web/src/lib/format/duration.ts
// The two ways the site writes a length of time (owner ruling 2026-11-10: one clock style, one
// spoken style, everywhere). The clock is the game's and the log's: `1:42`, `1:02:09`, tenths
// under a minute (`0:04.2`, where a second is a tenth of the pull). Spoken is for nights,
// cooldowns and settings: `2 h 41 min`, `1 min 30 s`. apps/shared/formats.json holds the cases.

export const SECOND_MS = 1000;
export const MINUTE_MS = 60 * SECOND_MS;
export const HOUR_MS = 60 * MINUTE_MS;
export const DAY_MS = 24 * HOUR_MS;
/** From this length the spoken form drops the seconds: a 41-minute night is `41 min`. */
const SPOKEN_SECONDS_UNDER_MS = 10 * MINUTE_MS;
const TENTHS_PER_MINUTE = 600;
const TENTHS_PER_HOUR = 36_000;

const two = (n: number): string => String(n).padStart(2, '0');

function safeMs(ms: number): number {
  return Number.isFinite(ms) && ms > 0 ? ms : 0;
}

/** Whole seconds as `m:ss`, or `h:mm:ss` from an hour. */
function wholeClock(totalSeconds: number): string {
  const hours = Math.floor(totalSeconds / 3600);
  const minutes = Math.floor((totalSeconds % 3600) / 60);
  const seconds = two(totalSeconds % 60);
  return hours > 0 ? `${hours}:${two(minutes)}:${seconds}` : `${minutes}:${seconds}`;
}

/** `0:04.2`, `1:01.4`: minutes and seconds to the tenth, inside an hour. */
function tenthsClock(tenths: number): string {
  const minutes = Math.floor(tenths / TENTHS_PER_MINUTE);
  const seconds = ((tenths % TENTHS_PER_MINUTE) / 10).toFixed(1).padStart(4, '0');
  return `${minutes}:${seconds}`;
}

/**
 * A length on the game's clock: `1:42`, `12:05`, `1:02:09`; under a minute to the tenth
 * (`0:04.2`) unless it is a whole second (`0:45`).
 */
export function clockDuration(ms: number): string {
  const tenths = Math.round(safeMs(ms) / 100);
  if (tenths < TENTHS_PER_MINUTE && tenths % 10 !== 0) return tenthsClock(tenths);
  return wholeClock(Math.round(safeMs(ms) / SECOND_MS));
}

/** The same clock always to the tenth, for event lists: `0:04.2`, `1:01.4`, `1:00:01.4`. */
export function clockDurationPrecise(ms: number): string {
  const tenths = Math.round(safeMs(ms) / 100);
  if (tenths < TENTHS_PER_HOUR) return tenthsClock(tenths);
  const rest = tenthsClock(tenths % TENTHS_PER_HOUR).padStart(7, '0');
  return `${Math.floor(tenths / TENTHS_PER_HOUR)}:${rest}`;
}

/** A length in whole seconds on the clock (sim settings carry seconds): 180 is `3:00`. */
export function clockSeconds(seconds: number): string {
  return wholeClock(Math.max(0, Math.round(Number.isFinite(seconds) ? seconds : 0)));
}

/**
 * A length in words: `30 s`, `1 min 30 s`, `41 min`, `2 h`, `2 h 41 min`. Seconds show under ten
 * minutes; from there the length rounds to the minute.
 */
export function spokenDuration(ms: number): string {
  const safe = safeMs(ms);
  if (safe < SPOKEN_SECONDS_UNDER_MS) {
    const total = Math.round(safe / SECOND_MS);
    const minutes = Math.floor(total / 60);
    const seconds = total % 60;
    if (minutes === 0) return `${seconds} s`;
    return seconds === 0 ? `${minutes} min` : `${minutes} min ${seconds} s`;
  }
  const total = Math.round(safe / MINUTE_MS);
  const hours = Math.floor(total / 60);
  const minutes = total % 60;
  if (hours === 0) return `${minutes} min`;
  return minutes === 0 ? `${hours} h` : `${hours} h ${minutes} min`;
}

/** The spoken form for a length in minutes (the raid form's options): 90 is `1 h 30 min`. */
export function spokenMinutes(minutes: number): string {
  return spokenDuration(minutes * MINUTE_MS);
}
