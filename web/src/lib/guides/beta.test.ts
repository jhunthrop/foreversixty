import { describe, expect, it } from 'vitest';
import { betaOver } from './beta';
import { guidesCopy } from './copy';

const BEFORE = new Date('2026-11-04T22:59:59Z');
const AT_LAUNCH = new Date('2026-11-04T23:00:00Z');

describe('the beta cap copy', () => {
  it('flips at launch', () => {
    expect(betaOver(BEFORE)).toBe(false);
    expect(betaOver(AT_LAUNCH)).toBe(true);
  });

  it('names the level 30 cap only before launch', () => {
    expect(guidesCopy.statusNote(BEFORE)).toContain('The beta stops at level 30.');
    expect(guidesCopy.indexHonestyLine(BEFORE)).toContain('The beta stops at level 30.');
    expect(guidesCopy.statusNote(AT_LAUNCH)).not.toMatch(/beta|level 30/i);
    expect(guidesCopy.indexHonestyLine(AT_LAUNCH)).not.toMatch(/beta|level 30/i);
  });
});
