// web/src/lib/report/copy.test.ts
import { describe, expect, it } from 'vitest';
import { reportCopy } from './copy';

describe('the sample report’s muted line (P.1)', () => {
  it('says why the numbers are small, in the fewest words', () => {
    expect(reportCopy.sampleNote).toBe('Level 10 to 30 players, names changed.');
  });
});
