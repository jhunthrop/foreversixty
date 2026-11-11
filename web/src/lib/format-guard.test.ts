// web/src/lib/format-guard.test.ts
// One house style for every date, clock, number, duration, parse figure, countdown and class name
// (owner rulings 2026-11-10, component audit clusters 3, 4, 7, 8 and 26). The formatting lives in
// lib/dates.ts, lib/format/, lib/report/format.ts (parse figures), lib/ui/countdown.ts and
// lib/text.ts; a page that formats for itself drifts ("25 Sept" beside "25 Sep", a 98.6 parse
// reading "99" on one page and "98" on the next). This scans every .ts, .svelte and .astro file
// under src/ (tests, data, fixtures and content excluded) and fails on a local copy. Each rule
// carries a sample it must catch, so a rule that stops matching fails too.
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';

const SRC = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
const SKIPPED_DIRS = new Set(['data', 'fixtures', 'content']);

interface Rule {
  name: string;
  pattern: RegExp;
  /** Files (relative to src/) that own the thing the rule guards. */
  owners: readonly string[];
  /** A line the rule must catch. */
  sample: string;
}

const RULES: readonly Rule[] = [
  {
    name: 'a month-name list outside lib/dates.ts',
    pattern: /['"](?:Jan|January)['"]\s*,\s*['"]Feb/,
    owners: ['lib/dates.ts'],
    sample: "const M = ['Jan', 'Feb', 'Mar'];",
  },
  {
    name: '"Sept": months are three letters',
    pattern: /\bSept\b(?!ember)/,
    owners: [],
    sample: 'return `${day} Sept`;',
  },
  {
    name: 'date or number formatting through Intl or toLocale*String outside the format modules',
    pattern: /\bIntl\.(?:DateTimeFormat|NumberFormat|RelativeTimeFormat)\b|\.toLocale(?:Date|Time)?String\(/,
    owners: ['lib/dates.ts', 'lib/format/number.ts', 'lib/guild/raid-time.ts'],
    sample: "n.toLocaleString('en-US')",
  },
  {
    name: 'rounding or flooring a parse percentile outside lib/report/format.ts',
    pattern: /Math\.(?:round|floor|ceil)\([^)]*(?:percentile|parse_pct|bestParse)/i,
    owners: ['lib/report/format.ts'],
    sample: 'String(Math.floor(row.percentile))',
  },
  {
    name: 'a local duration, countdown or relative-time formatter',
    pattern:
      /\b(?:function|const)\s+(?:durationLabel|killTimeLabel|formatReportLength|formatCooldown|plaqueClock|agoText|relativeTime|formatDuration|formatDurationPrecise|formatAmount|formatPercent|ordinal)\b|\bfunction\s+span\b/,
    owners: ['lib/dates.ts', 'lib/format/number.ts', 'lib/format/duration.ts'],
    sample: 'export function durationLabel(ms: number): string {',
  },
  {
    name: 'a second list of the nine classes',
    pattern: /['"]warrior['"]\s*,\s*['"]paladin['"]/,
    owners: ['lib/ui/primitives.ts'],
    sample: "const CLASSES = ['warrior', 'paladin', 'hunter'];",
  },
  {
    name: 'a local capitalise helper',
    pattern: /\bfunction\s+(?:capitali[sz]e|upperFirst|titleCase)\b/,
    owners: ['lib/text.ts'],
    sample: 'function capitalise(word: string): string {',
  },
  {
    name: 'a hand-written "Updated" stamp outside UpdatedStamp',
    pattern: /(?:>|^)\s*Updated \{/m,
    owners: ['components/UpdatedStamp.astro', 'components/UpdatedStamp.svelte'],
    sample: '<span slot="meta">Updated {dayMonth(updated)}</span>',
  },
  {
    name: 'a time-zone lookup outside lib/dates.ts viewerZone()',
    pattern: /resolvedOptions\(\)\.timeZone/,
    owners: ['lib/dates.ts'],
    sample: 'const tz = Intl.DateTimeFormat().resolvedOptions().timeZone;',
  },
];

function scannedFiles(dir: string): string[] {
  return readdirSync(dir, { withFileTypes: true }).flatMap((entry) => {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) return SKIPPED_DIRS.has(entry.name) ? [] : scannedFiles(full);
    if (entry.name.includes('.test.') || entry.name.endsWith('.generated.ts')) return [];
    return /\.(ts|svelte|astro)$/.test(entry.name) ? [full] : [];
  });
}

function relative(file: string): string {
  return path.relative(SRC, file).split(path.sep).join('/');
}

/** Comments name the old formats on purpose (why a helper exists); only code counts. */
function withoutComments(text: string): string {
  return text
    .replace(/\/\*[\s\S]*?\*\//g, '')
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/(^|[^:'"`])\/\/.*$/gm, '$1');
}

function offenders(rule: Rule): string[] {
  // Scanned over the whole file, not line by line: a class list spreads one name per line.
  const global = new RegExp(rule.pattern.source, `${rule.pattern.flags.replace('g', '')}g`);
  return scannedFiles(SRC)
    .filter((file) => !rule.owners.includes(relative(file)))
    .flatMap((file) => {
      const text = withoutComments(readFileSync(file, 'utf8'));
      return [...text.matchAll(global)].map((match) => {
        const line = text.slice(0, match.index).split('\n').length;
        return `${relative(file)}:${line}: ${match[0].trim()}`;
      });
    });
}

/** Hand-written prose (the changelog, the content pages) carries dates the formatters never see. */
const PROSE_DATE = /\bSept\b(?!ember)|\b\d{1,2}:\d{2} [ap]\.m\.|\b[ap]\.m\. P[DS]T\b/g;

function markdownFiles(dir: string): string[] {
  return readdirSync(dir, { withFileTypes: true }).flatMap((entry) => {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) return markdownFiles(full);
    return entry.name.endsWith('.md') ? [full] : [];
  });
}

describe('format guard, prose', () => {
  it('catches "Sept 17" and "10:00 a.m."', () => {
    expect('on Sept 17 at 10:00 a.m. PDT'.match(PROSE_DATE)?.length).toBe(2);
  });
  it('finds no "Sept" or "a.m./p.m." in the content collections', () => {
    const found = markdownFiles(path.join(SRC, 'content')).flatMap((file) =>
      [...readFileSync(file, 'utf8').matchAll(PROSE_DATE)].map((match) => `${relative(file)}: ${match[0]}`),
    );
    expect(found).toEqual([]);
  });
});

describe('format guard', () => {
  for (const rule of RULES) {
    it(`catches ${rule.name}`, () => {
      expect(rule.pattern.test(rule.sample)).toBe(true);
    });
    it(`finds no ${rule.name}`, () => {
      expect(offenders(rule)).toEqual([]);
    });
  }
});
