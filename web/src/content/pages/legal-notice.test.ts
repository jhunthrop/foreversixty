// web/src/content/pages/legal-notice.test.ts
// The full trademark sentence and the operator line left the footer in round 5; /about and /terms
// carry them verbatim (lib/legal-notice.ts), so the site still says both on a page of its own.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { describe, expect, it } from 'vitest';
import { LEGAL_NOTICE_PAGES, LEGAL_OPERATOR, LEGAL_TRADEMARK } from '../../lib/legal-notice';

const PAGES_DIR = path.resolve(import.meta.dirname);

describe('the legal notice', () => {
  for (const file of LEGAL_NOTICE_PAGES) {
    const text = readFileSync(path.join(PAGES_DIR, file), 'utf8');
    it(`${file} carries the trademark sentence and the operator line verbatim`, () => {
      expect(text).toContain(LEGAL_TRADEMARK);
      expect(text).toContain(LEGAL_OPERATOR);
    });
  }
});
