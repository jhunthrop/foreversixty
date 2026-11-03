// web/src/content/pages/privacy-analytics.test.ts
// The privacy page says what the visit beacon and the tracked /go/ links collect
// (lib/golinks): counts without cookies or personal data, and an account's first link.
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { describe, expect, it } from 'vitest';

const text = readFileSync(path.join(import.meta.dirname, 'privacy.md'), 'utf8');

describe('privacy page', () => {
  it('says visits and link clicks are counted without cookies or personal data', () => {
    expect(text).toMatch(/count visits and clicks on our links without cookies or any personal data/);
  });

  it('says an account keeps the first of our links it arrived from', () => {
    expect(text).toMatch(/which of our links, if any, an account first arrived from/);
  });
});
