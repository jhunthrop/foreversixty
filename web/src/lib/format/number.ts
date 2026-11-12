// web/src/lib/format/number.ts
// Every number the site prints is grouped here (owner ruling 2026-11-10: one locale, never a bare
// toLocaleString). The words and the grouping are the site's: a page rendered at build time and
// the same page hydrated in a German browser must read the same, so the locale is fixed, not the
// viewer's. apps/shared/formats.json holds the canonical cases every surface replays.

/** The locale every figure is grouped in (formats.json `site_locale`). */
export const SITE_LOCALE = 'en-US';

/** An empty figure cell: an en dash, never an em dash or a hyphen. */
export const EMPTY_FIGURE = '–';

// Built once, on first use: constructing a NumberFormat loads locale data, tens of milliseconds on
// a throttled phone, so a module that only imports a helper pays nothing until it formats.
const formats = new Map<string, Intl.NumberFormat>();
function numberFormat(minFraction: number, maxFraction: number): Intl.NumberFormat {
  const key = `${minFraction}:${maxFraction}`;
  let format = formats.get(key);
  if (format === undefined) {
    format = new Intl.NumberFormat(SITE_LOCALE, {
      minimumFractionDigits: minFraction,
      maximumFractionDigits: maxFraction,
    });
    formats.set(key, format);
  }
  return format;
}

function finite(n: number): number {
  return Number.isFinite(n) ? n : 0;
}

/** A whole figure, grouped: `1,234`. Rounds to the nearest whole. */
export function grouped(n: number): string {
  return numberFormat(0, 0).format(Math.round(finite(n)));
}

/** A figure to exactly one decimal, grouped: `1,234.6`, `3.0`. */
export function oneDecimal(n: number): string {
  return numberFormat(1, 1).format(finite(n));
}

/** A figure with up to `max` decimals and no trailing zeros, grouped: `1.5`, `12`. */
export function upTo(n: number, max: number): string {
  return numberFormat(0, max).format(finite(n));
}

/** 1.25 stays 1.25, 12.50 becomes 12.5, 3.00 becomes 3. */
function trimZero(value: number): string {
  return value.toFixed(2).replace(/\.?0+$/, '');
}

/** `81k`, `239.2k`: a figure in thousands to one trimmed decimal. */
export function thousands(n: number): string {
  return `${(finite(n) / 1000).toFixed(1).replace(/\.0$/, '')}k`;
}

/**
 * A large amount as a column reads it: grouped below 100k (`99,999`), then `239.2k`, `1.09M`,
 * `2.5B`. From 100k the grouped form is seven characters against a neighbour's `1.09M`, and a
 * column mixing the two reads the longer string as the bigger number.
 */
export function compact(n: number): string {
  const safe = finite(n);
  const abs = Math.abs(safe);
  if (abs >= 1_000_000_000) return `${trimZero(safe / 1_000_000_000)}B`;
  if (abs >= 1_000_000) return `${trimZero(safe / 1_000_000)}M`;
  if (abs >= 100_000) return thousands(safe);
  return grouped(safe);
}

/** A percentage to one decimal: `12.5%`. */
export function percent(value: number): string {
  return `${finite(value).toFixed(1)}%`;
}

/** 1st, 2nd, 3rd, 4th … 11th, 12th, 13th … 21st. */
export function ordinal(n: number): string {
  const rest = n % 100;
  if (rest >= 11 && rest <= 13) return `${n}th`;
  const last = n % 10;
  return `${n}${last === 1 ? 'st' : last === 2 ? 'nd' : last === 3 ? 'rd' : 'th'}`;
}

/** A grouped count with its noun: `1 kill`, `1,204 views`. */
export function counted(n: number, one: string, many = `${one}s`): string {
  return `${grouped(n)} ${n === 1 ? one : many}`;
}

/** An amount in cents as money, always to the cent and grouped: `$2.00`, `$1,234.56`. A column
 *  of amounts lines up; a whole amount reads `$10.00`, not `$10`, beside `$9.50`. */
export function usdCents(cents: number): string {
  const amount = Math.abs(finite(cents)) / 100;
  return `${cents < 0 ? '-' : ''}$${numberFormat(2, 2).format(amount)}`;
}
