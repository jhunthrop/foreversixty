// web/src/lib/reports/copy.test.ts
import { describe, expect, it } from 'vitest';
import { logsCopy, pairingCopy } from './copy';

describe('logsCopy.guildEmpty', () => {
  it('names the guild, verbatim per spec §6', () => {
    expect(logsCopy.guildEmpty('The Last Watch')).toBe('No The Last Watch reports yet.');
  });
});

describe('pairingCopy.success', () => {
  it('names the device, verbatim per spec §6', () => {
    expect(pairingCopy.success('MacBook Pro')).toBe(
      'MacBook Pro paired. It starts uploading as soon as you are logging.',
    );
  });
});
