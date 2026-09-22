// web/src/content/pages/legal-markers.test.ts
// Coordinator instruction (docs/superpowers/plans/2026-09-21-pay-web.md, Task 12): the
// legal pages render with visible OWNER: markers and a draft banner until the owner fills
// every marker in. This test is the backstop that keeps a forgotten marker from shipping in
// a real production build once the owner believes the pages are final: it only enforces
// anything when LEGAL_PAGES_FINAL=1 is set, and never on a fixture build (FOREVER_DATA=
// fixture), which intentionally never sets that flag.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { describe, expect, it } from 'vitest';

const PAGES_DIR = path.resolve(import.meta.dirname);
const LEGAL_FILES = ['terms.md', 'privacy.md', 'refunds.md'];

describe('legal page OWNER markers', () => {
  const isFinal = process.env.LEGAL_PAGES_FINAL === '1';
  const isFixtureBuild = process.env.FOREVER_DATA === 'fixture';

  it.runIf(isFinal && !isFixtureBuild)('no OWNER: marker remains once LEGAL_PAGES_FINAL=1 is set', () => {
    for (const file of LEGAL_FILES) {
      const text = readFileSync(path.join(PAGES_DIR, file), 'utf8');
      expect(text, `${file} still has an OWNER: marker`).not.toMatch(/OWNER:/);
    }
  });

  it('every legal page currently carries at least one OWNER: marker (sanity check this test can fail)', () => {
    for (const file of LEGAL_FILES) {
      const text = readFileSync(path.join(PAGES_DIR, file), 'utf8');
      expect(text).toMatch(/OWNER:/);
    }
  });
});
