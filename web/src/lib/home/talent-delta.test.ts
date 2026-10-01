// web/src/lib/home/talent-delta.test.ts
import { describe, expect, it } from 'vitest';
import { talentDeltaFor } from './talent-delta';

describe('talentDeltaFor', () => {
  it('is zero when the character matches the band exactly', () => {
    const trees = ['0000000000000000', '35300000000000000', '000000000000000000'];
    expect(talentDeltaFor(trees, trees.join('-'))).toBe(0);
  });

  it('sums the absolute per-cell rank difference across every tree', () => {
    // Band tree 1 spends rank 3 in its first cell; the character spends rank 1 there --
    // |1 - 3| = 2, every other cell equal.
    const band = '0000000000000000-35300000000000000-000000000000000000';
    const trees = ['0000000000000000', '15300000000000000', '000000000000000000'];
    expect(talentDeltaFor(trees, band)).toBe(2);
  });

  it('counts points the character spent that the band does not (band spends none there)', () => {
    const band = '0000000000000000-00000000000000000-000000000000000000';
    const trees = ['5000000000000000', '00000000000000000', '000000000000000000'];
    expect(talentDeltaFor(trees, band)).toBe(5);
  });

  it('counts points the band spends that the character has not matched', () => {
    const band = '0000000000000000-50000000000000000-000000000000000000';
    const trees = ['0000000000000000', '00000000000000000', '000000000000000000'];
    expect(talentDeltaFor(trees, band)).toBe(5);
  });

  it('is symmetric: a mismatch in both directions on the same cell adds, never cancels', () => {
    // Character has 5 points in a cell the band spends 2 in -- |5 - 2| = 3, not 5 - 2
    // double-counted or net to zero.
    const band = '0000000000000000-20000000000000000-000000000000000000';
    const trees = ['0000000000000000', '50000000000000000', '000000000000000000'];
    expect(talentDeltaFor(trees, band)).toBe(3);
  });

  it('treats a shorter string (missing trailing cells) as trailing zeros on either side', () => {
    expect(talentDeltaFor(['3'], '0')).toBe(3);
    expect(talentDeltaFor(['0'], '3')).toBe(3);
  });
});
