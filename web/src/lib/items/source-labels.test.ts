// web/src/lib/items/source-labels.test.ts
import { describe, expect, it } from 'vitest';
import { craftedSourceLabel, pvpRankTitle, pvpSourceLabel } from './source-labels';

describe('craftedSourceLabel', () => {
  it('never doubles the profession when the source name already is it', () => {
    expect(craftedSourceLabel('Blacksmithing', 'blacksmithing')).toBe('Crafted: Blacksmithing');
  });

  it('names the profession when it genuinely differs from the source name', () => {
    expect(craftedSourceLabel('Recipe: Thing', 'tailoring')).toBe('Crafted: Tailoring · Recipe: Thing');
  });

  it('omits the profession entirely when none is given', () => {
    expect(craftedSourceLabel('Blacksmithing')).toBe('Crafted: Blacksmithing');
  });
});

describe('pvpRankTitle', () => {
  it('reads Vanilla’s own rank ladder, offset by the client’s +4 RequiredPVPRank', () => {
    expect(pvpRankTitle('alliance', 5)).toBe('Private');
    expect(pvpRankTitle('alliance', 18)).toBe('Grand Marshal');
    expect(pvpRankTitle('horde', 5)).toBe('Scout');
    expect(pvpRankTitle('horde', 18)).toBe('High Warlord');
  });

  it('is undefined for a rank outside the 14-rank ladder', () => {
    expect(pvpRankTitle('alliance', 4)).toBeUndefined();
  });
});

describe('pvpSourceLabel', () => {
  it('names the rank, its title and the faction', () => {
    expect(pvpSourceLabel(11, 'alliance')).toBe('PvP rank 11 · Knight-Lieutenant · Alliance');
  });

  it('falls back to rank and faction alone when the rank has no known title', () => {
    expect(pvpSourceLabel(4, 'alliance')).toBe('PvP rank 4 · Alliance');
  });
});
