import { describe, expect, it } from 'vitest';
import { PHASES, phaseAt, phaseLabel } from './phases';

describe('phaseAt', () => {
  it('files a fight under the phase that was open when it happened', () => {
    expect(phaseAt('2026-09-01T12:00:00Z')).toBe('pre-beta');
    expect(phaseAt('2026-09-26T02:01:07.173Z')).toBe('beta');
    expect(phaseAt('2026-11-04T22:59:59Z')).toBe('beta');
    expect(phaseAt('2026-11-04T23:00:00Z')).toBe('launch');
    expect(phaseAt('2027-01-01T00:00:00Z')).toBe('raids-1');
  });

  it('falls back to the first phase for a date it cannot read', () => {
    expect(phaseAt('not a date')).toBe(PHASES[0].id);
  });
});

describe('phaseLabel', () => {
  it('echoes an id it does not know rather than hiding it', () => {
    expect(phaseLabel('launch')).toBe('Launch');
    expect(phaseLabel('raids-9')).toBe('raids-9');
  });
});
