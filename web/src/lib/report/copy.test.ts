// web/src/lib/report/copy.test.ts
import { describe, expect, it } from 'vitest';
import { reportCopy } from './copy';

describe('the sample report’s muted line (P.1)', () => {
  it('says why the numbers are small, and that yours look like it', () => {
    expect(reportCopy.sampleNote).toBe(
      'A real beta night with the names changed. Level 10 to 30 players, so the numbers are small. Your own reports look like this.',
    );
  });
});
