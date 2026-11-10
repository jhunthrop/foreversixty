// web/src/lib/formats.test.ts
// The web replays every case of apps/shared/formats.json, the house formats the server pages, the
// companion and both apps replay too (owner rulings 2026-11-10). A case that fails here is a
// format that drifted from the one every other surface prints.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import {
  MONTHS,
  clock,
  compactAgo,
  dayMonth,
  formatSyncStamp,
  longDay,
  monthHeading,
  relativeDay,
  updatedLabel,
  weekdayDayMonth,
  zonedDayMonthYear,
} from './dates';
import { clockDuration, clockDurationPrecise, spokenDuration } from './format/duration';
import { EMPTY_FIGURE, SITE_LOCALE, compact, grouped, oneDecimal, ordinal, percent } from './format/number';
import { parseLabel, parseText, parseTier } from './report/format';
import { clockLabel, clockParts } from './ui/countdown';

interface DateCase {
  in: string;
  zone: string;
  now?: string;
  locale?: string;
  always_minutes?: boolean;
  out: string;
}
interface Pair<I> {
  in: I;
  out: string;
}
interface Formats {
  site_locale: string;
  months: string[];
  dates: Record<
    'day_month' | 'day_month_year' | 'weekday_day_month' | 'long_day' | 'month_heading' | 'updated' | 'stamp',
    DateCase[]
  >;
  clock: DateCase[];
  relative: { now: string; compact_ago: Pair<string>[]; relative_day: Pair<string>[] };
  durations: Record<'clock' | 'clock_precise' | 'spoken', { ms: number; out: string }[]>;
  numbers: { empty: string } & Record<
    'grouped' | 'one_decimal' | 'compact' | 'percent' | 'ordinal',
    Pair<number>[]
  >;
  parse: { cases: { in: number; out: string; tier: string; label: string }[] };
  countdown: {
    now: string;
    cases: { target: string; parts: { value: string; unit: string }[]; label: string }[];
  };
}

const FIXTURE = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  '../../../apps/shared/formats.json',
);
const formats = JSON.parse(readFileSync(FIXTURE, 'utf8')) as Formats;

function each<T extends { out: string }>(cases: T[], run: (c: T) => string): void {
  for (const c of cases) expect(run(c), JSON.stringify(c)).toBe(c.out);
}

describe('apps/shared/formats.json', () => {
  it('shares the site locale and the month table', () => {
    expect(formats.site_locale).toBe(SITE_LOCALE);
    expect(formats.months).toEqual([...MONTHS]);
  });

  it('dates', () => {
    const d = formats.dates;
    each(d.day_month, (c) => dayMonth(new Date(c.in), c.zone));
    each(d.day_month_year, (c) => zonedDayMonthYear(c.in, c.zone));
    each(d.weekday_day_month, (c) => weekdayDayMonth(new Date(c.in), c.zone, new Date(c.now ?? c.in)));
    each(d.long_day, (c) => longDay(new Date(c.in), c.zone));
    each(d.month_heading, (c) => monthHeading(new Date(c.in)));
    each(d.updated, (c) => updatedLabel(new Date(c.in)));
    each(d.stamp, (c) => formatSyncStamp(new Date(c.in), c.zone, c.locale));
  });

  it('clock', () => {
    each(formats.clock, (c) =>
      clock(new Date(c.in), c.zone, { locale: c.locale, alwaysMinutes: c.always_minutes === true }),
    );
  });

  it('relative', () => {
    const now = new Date(formats.relative.now);
    each(formats.relative.compact_ago, (c) => compactAgo(new Date(c.in), now));
    each(formats.relative.relative_day, (c) => relativeDay(new Date(c.in), now));
  });

  it('durations', () => {
    each(formats.durations.clock, (c) => clockDuration(c.ms));
    each(formats.durations.clock_precise, (c) => clockDurationPrecise(c.ms));
    each(formats.durations.spoken, (c) => spokenDuration(c.ms));
  });

  it('numbers', () => {
    const n = formats.numbers;
    expect(EMPTY_FIGURE).toBe(n.empty);
    each(n.grouped, (c) => grouped(c.in));
    each(n.one_decimal, (c) => oneDecimal(c.in));
    each(n.compact, (c) => compact(c.in));
    each(n.percent, (c) => percent(c.in));
    each(n.ordinal, (c) => ordinal(c.in));
  });

  it('parse figures', () => {
    for (const c of formats.parse.cases) {
      expect(parseText(c.in)).toBe(c.out);
      expect(parseTier(c.in)).toBe(c.tier);
      expect(parseLabel(c.in)).toBe(c.label);
    }
  });

  it('countdowns', () => {
    const now = new Date(formats.countdown.now);
    for (const c of formats.countdown.cases) {
      const parts = clockParts(new Date(c.target), now);
      expect(parts).toEqual(c.parts);
      expect(clockLabel(parts)).toBe(c.label);
    }
  });
});
