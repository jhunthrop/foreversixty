// web/src/components/state-guard.test.ts
// One set of controls and states (component audit 2026-11-10 clusters 10 to 16, 19 and 28): the
// copy button and its "Copied" (owner ruling 2026-11-10: clears after about two seconds), the error
// line with Try again, the empty state, the filter popover, the focus trap, the tabs and the
// disclosure each live once in components/ui/ and lib/ui/. A page that builds its own drifts (a
// "Copied" that never clears, a retry that is a link on one page and a gold button on the next, a
// tab strip whose arrows stop at the ends). This scans every .ts, .svelte and .astro file under
// src/ (tests, data, fixtures, content and generated files excluded) and fails on a local copy.
// Each rule carries a sample it must catch, so a rule that stops matching fails too.
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';

const SRC = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
const SKIPPED_DIRS = new Set(['data', 'fixtures', 'content']);
const UI = 'components/ui/';

interface Rule {
  name: string;
  pattern: RegExp;
  /** Files (relative to src/) that own the thing the rule guards; a trailing `/` owns a directory. */
  owners: readonly string[];
  /** A line the rule must catch. */
  sample: string;
  /** Only files this returns true for are scanned (the empty-state heuristic). */
  applies?: (text: string) => boolean;
}

const RULES: readonly Rule[] = [
  {
    name: 'a clipboard call outside lib/ui/copy-feedback.ts',
    pattern: /navigator\.clipboard\b/,
    owners: ['lib/ui/copy-feedback.ts'],
    sample: 'await navigator.clipboard.writeText(url);',
  },
  {
    name: 'a tab list, dialog or popover outside components/ui/',
    pattern: /role=["'{](?:tablist|dialog|alertdialog)\b|aria-modal=|\bpopover=["'](?:manual|auto)["']/,
    owners: [
      UI,
      // The header's character selector: a dialog on phones, trapped through lib/ui/focus-trap.ts.
      'components/nav/SelectorPanel.svelte',
      // The typed delete confirmation, trapped through lib/ui/focus-trap.ts.
      'components/account/AccountDeleteDialog.svelte',
      // The chart's tooltip, raised into the top layer and placed by lib/items/tooltip-position.ts.
      'components/chart/ChartTooltip.svelte',
    ],
    sample: '<div role="tablist" aria-label="Bosses">',
  },
  {
    name: 'a <details> outside ui/Disclosure',
    pattern: /<details\b/,
    owners: ['components/ui/Disclosure.svelte', 'components/ui/Disclosure.astro'],
    sample: '<details class="mt-2"><summary>More</summary>',
  },
  {
    name: '"Try again", "Retry" or "Copied" written outside lib/ui/copy.ts',
    pattern: /[\s'"`>(](?:Try again|Retry|Copied)\b(?=[^\n]*['"`<]|\s*\n\s*<)/,
    owners: ['lib/ui/copy.ts'],
    sample: "  failed: 'That did not save. Try again.',",
  },
  {
    name: 'a muted "No …" paragraph in a file that does not use EmptyState',
    pattern: /<(?:p|div|span)\b[^>]*\btext-muted\b[^>]*>\s*No\s/,
    owners: [UI],
    sample: '<p class="text-muted text-[14px]">No raids planned yet.</p>',
    applies: (text) => !/\bEmptyState\b/.test(text),
  },
  {
    name: 'a focusable-element selector outside lib/ui/focus-trap.ts',
    pattern: /button:not\(\[disabled\]\)|a\[href\]\s*,\s*button/,
    owners: ['lib/ui/focus-trap.ts'],
    sample: "const FOCUSABLE = 'a[href],button:not([disabled]),input';",
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

function owned(rule: Rule, rel: string): boolean {
  return rule.owners.some((owner) => (owner.endsWith('/') ? rel.startsWith(owner) : rel === owner));
}

/** Comments name the old patterns on purpose (why a helper exists); only code counts. */
function withoutComments(text: string): string {
  return text
    .replace(/\/\*[\s\S]*?\*\//g, '')
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/(^|[^:'"`])\/\/.*$/gm, '$1');
}

function offenders(rule: Rule): string[] {
  const global = new RegExp(rule.pattern.source, `${rule.pattern.flags.replace('g', '')}g`);
  return scannedFiles(SRC)
    .filter((file) => !owned(rule, relative(file)))
    .flatMap((file) => {
      const text = withoutComments(readFileSync(file, 'utf8'));
      if (rule.applies !== undefined && !rule.applies(text)) return [];
      return [...text.matchAll(global)].map((match) => {
        const line = text.slice(0, match.index).split('\n').length;
        return `${relative(file)}:${line}: ${match[0].trim()}`;
      });
    });
}

describe('state guard', () => {
  for (const rule of RULES) {
    it(`catches ${rule.name}`, () => {
      expect(rule.pattern.test(rule.sample)).toBe(true);
      expect(rule.applies?.(rule.sample) ?? true).toBe(true);
    });
    it(`finds no ${rule.name}`, () => {
      expect(offenders(rule)).toEqual([]);
    });
  }

  it('the copy rule leaves a lower-case "copied" and an identifier alone', () => {
    const rule = RULES[3];
    expect(rule.pattern.test("setupLinkCopied: 'Link copied',")).toBe(false);
    expect(rule.pattern.test('onRetry={() => load()}')).toBe(false);
  });
});
