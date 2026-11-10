// The one countdown (lib/ui/countdown.ts, owner ruling 2026-11-10); apps/shared/formats.json holds
// the cross-surface cases (lib/formats.test.ts).
import { describe, expect, it } from 'vitest';
import { clockLabel, clockParts, countdownText } from './countdown';

const AT = new Date('2026-12-09T00:00:00Z');

describe('countdown clock', () => {
  it('counts days, hours and minutes', () => {
    const parts = clockParts(AT, new Date('2026-11-06T18:00:00Z'));
    expect(parts).toEqual([
      { value: '32', unit: 'days' },
      { value: '6', unit: 'hours' },
      { value: '0', unit: 'min' },
    ]);
    expect(clockLabel(parts)).toBe('32 days, 6 hours, 0 min');
  });

  it('drops the units that are zero ahead of the first, and never shows seconds', () => {
    expect(clockParts(AT, new Date('2026-12-08T23:58:30Z'))).toEqual([{ value: '2', unit: 'min' }]);
    expect(countdownText(AT, new Date('2026-12-08T21:00:00Z'))).toBe('3 hours, 0 min');
  });

  it('is gone once the instant has passed', () => {
    expect(clockParts(AT, AT)).toEqual([]);
    expect(countdownText(AT, AT)).toBe('');
  });
});
