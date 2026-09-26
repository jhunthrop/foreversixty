// web/src/lib/sim/after-sim-copy.test.ts
import { describe, expect, it } from 'vitest';
import { afterSimCopy, afterSimSentence } from './after-sim-copy';

describe('afterSimSentence', () => {
  it('names the item, the source and the gain when an upgrade exists', () => {
    expect(
      afterSimSentence({
        itemName: 'Bracers of X',
        sourceName: 'Blackfathom Deeps',
        gain: '+14',
        savedAt: '',
      }),
    ).toBe('Upgrade: Bracers of X from Blackfathom Deeps, +14 DPS.');
  });

  it('points at Top Gear when there is no upgrade on record', () => {
    expect(afterSimSentence(null)).toBe(afterSimCopy.noUpgrade);
  });
});
