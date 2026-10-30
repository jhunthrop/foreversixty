// web/src/components/ui/select-sites.test.ts
// One dropdown site-wide (design/specs/2026-10-30-controls-charts.md §A10, §E): no component
// draws a raw <select> except the shared Select itself. TimeChart and ThreatTable belong to the
// chart lane, which moves them onto Select when it lands; they leave this list then.
import { readFileSync, readdirSync, statSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';

const SRC = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../..');
const ALLOWED = new Set([
  'components/ui/Select.svelte',
  'components/report/TimeChart.svelte',
  'components/report/ThreatTable.svelte',
]);

function* componentFiles(dir: string): Generator<string> {
  for (const name of readdirSync(dir)) {
    const full = path.join(dir, name);
    if (statSync(full).isDirectory()) yield* componentFiles(full);
    else if (/\.(svelte|astro)$/.test(name)) yield full;
  }
}

describe('raw selects', () => {
  it('only the shared Select renders a <select>', () => {
    const offenders = [...componentFiles(SRC)]
      .map((file) => path.relative(SRC, file))
      .filter((file) => !ALLOWED.has(file))
      .filter((file) => readFileSync(path.join(SRC, file), 'utf8').includes('<select'));
    expect(offenders).toEqual([]);
  });
});
