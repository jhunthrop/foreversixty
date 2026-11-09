// web/src/lib/report/copy.test.ts
import { describe, expect, it } from 'vitest';
import { reportCopy } from './copy';

describe('the sample report’s muted line (P.1)', () => {
  it('says what the sample is, in the fewest words', () => {
    expect(reportCopy.sampleNote).toBe('A real dungeon night, names changed.');
  });
});
