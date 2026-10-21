// web/src/lib/home/outcome-rows.test.ts
import { describe, expect, it } from 'vitest';
import type { CurrentCharacter } from '../current-character';
import { outcomeHrefs, outcomeRows, outcomeRowsHeading } from './outcome-rows';

const CURRENT: CurrentCharacter = {
  source: 'addon',
  ref: 'FS1:abc',
  label: 'Bow Jackzon · Marksmanship Hunter',
  classSlug: 'hunter',
  savedAt: '2026-10-20T10:00:00Z',
};

describe('outcome rows (P.1)', () => {
  it('are three, in the player’s words, with no product name or count as a title', () => {
    expect(outcomeRows.map((row) => row.title)).toEqual([
      'Plan your talents',
      'Sim your gear',
      'Upload tonight’s raid',
    ]);
    for (const row of outcomeRows) {
      expect(row.line.length).toBeLessThan(60);
      expect(row.title + row.line).not.toMatch(/\d|Planner|Simulator|Rankings|Guides/);
    }
    expect(outcomeRowsHeading).not.toMatch(/what the site does/i);
  });

  it('open the tools with no current character', () => {
    expect(outcomeHrefs(null)).toEqual({ planner: '/planner', sim: '/sim', logs: '/logs' });
  });

  it('open the current character’s planner and simulator when there is one', () => {
    const hrefs = outcomeHrefs(CURRENT);
    expect(hrefs.planner).toBe('/planner?code=FS1%3Aabc');
    expect(hrefs.sim).toContain('code=');
    expect(hrefs.sim.startsWith('/sim')).toBe(true);
    expect(hrefs.logs).toBe('/logs');
  });
});
