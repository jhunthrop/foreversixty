// web/src/content/pages/legal-markers.test.ts
// The legal pages shipped with visible OWNER: markers and a draft banner until the owner
// supplied every fact (docs/superpowers/plans/2026-09-21-pay-web.md, Task 12). They were all
// filled in on 2026-10-24 (Commish LLC, Arkansas; help@foreversixty.gg; the registered agent's
// address; no partial-period refunds). This test keeps a marker or the banner from ever coming
// back.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { describe, expect, it } from 'vitest';

const PAGES_DIR = path.resolve(import.meta.dirname);
const LEGAL_FILES = ['terms.md', 'privacy.md', 'refunds.md'];

describe('legal pages', () => {
  for (const file of LEGAL_FILES) {
    const text = readFileSync(path.join(PAGES_DIR, file), 'utf8');
    it(`${file} carries no OWNER: marker and no draft banner`, () => {
      expect(text).not.toMatch(/OWNER:/);
      expect(text).not.toContain('legal-draft-banner');
    });
    it(`${file} names Commish LLC and the support address`, () => {
      expect(text).toContain('Commish LLC');
      expect(text).toContain('help@foreversixty.gg');
    });
    it(`${file} promises no retention window`, () => {
      expect(text).not.toMatch(/ninety|two-year|two years|retention falls back|long-retention/i);
    });
  }
});
