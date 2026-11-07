// The countdown panel's clock (lib/ui/countdown.ts, home B spec 2026-11-07 §1).
import { describe, expect, it } from 'vitest';
import { clockLabel, clockParts, nextTickMs } from './countdown';

const AT = new Date('2026-12-09T00:00:00Z');

describe('countdown clock', () => {
  it('counts days, hours and minutes, two digits each', () => {
    const parts = clockParts(AT, new Date('2026-11-06T18:00:00Z'));
    expect(parts).toEqual([
      { value: '32', unit: 'days' },
      { value: '06', unit: 'hours' },
      { value: '00', unit: 'min' },
    ]);
    expect(clockLabel(parts ?? [])).toBe('32 days, 6 hours, 0 min');
  });

  it('shows seconds only in the final hour', () => {
    expect(clockParts(AT, new Date('2026-12-08T23:58:30Z'))).toEqual([
      { value: '01', unit: 'min' },
      { value: '30', unit: 'sec' },
    ]);
    expect(nextTickMs(AT, new Date('2026-12-08T23:58:30Z'))).toBe(1000);
  });

  it('is gone once the instant has passed', () => {
    expect(clockParts(AT, AT)).toBeNull();
  });

  it('ticks on the next minute boundary of the time left', () => {
    expect(nextTickMs(AT, new Date('2026-11-06T18:00:20Z'))).toBe(40_000);
  });
});
