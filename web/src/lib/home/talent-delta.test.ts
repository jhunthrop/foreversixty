// web/src/lib/home/talent-delta.test.ts
import { describe, expect, it } from 'vitest';
import { talentDeltaFor } from './talent-delta';

describe('talentDeltaFor', () => {
  it('is zero when the character matches the band exactly', () => {
    const trees = ['0000000000000000', '35300000000000000', '000000000000000000'];
    expect(talentDeltaFor(trees, trees.join('-'))).toBe(0);
  });

  it('counts the band points the character has not matched in a cell', () => {
    // Band tree 1 spends rank 3 in its first cell; the character spends rank 1 there.
    const band = '0000000000000000-35300000000000000-000000000000000000';
    const trees = ['0000000000000000', '15300000000000000', '000000000000000000'];
    expect(talentDeltaFor(trees, band)).toBe(2);
  });

  it('does not count points the character spent where the band spends none', () => {
    const band = '0000000000000000-00000000000000000-000000000000000000';
    const trees = ['5000000000000000', '00000000000000000', '000000000000000000'];
    expect(talentDeltaFor(trees, band)).toBe(0);
  });

  it('counts points the band spends that the character has not matched', () => {
    const band = '0000000000000000-50000000000000000-000000000000000000';
    const trees = ['0000000000000000', '00000000000000000', '000000000000000000'];
    expect(talentDeltaFor(trees, band)).toBe(5);
  });

  it('never exceeds the band build points: a build with every point elsewhere reads 11 of 11, not 22', () => {
    const band = '3300000000000000-50000000000000000-000000000000000000';
    const trees = ['0000000000000000', '00000000000000000', '330000000000000000'];
    const bandPoints = 11;
    expect(talentDeltaFor(trees, band)).toBe(bandPoints);
  });

  it('an over-spent cell does not cancel a short one elsewhere', () => {
    const band = '0000000000000000-20000000000000000-200000000000000000';
    const trees = ['0000000000000000', '50000000000000000', '000000000000000000'];
    expect(talentDeltaFor(trees, band)).toBe(2);
  });

  it('treats a shorter string (missing trailing cells) as trailing zeros on either side', () => {
    expect(talentDeltaFor(['3'], '0')).toBe(0);
    expect(talentDeltaFor(['0'], '3')).toBe(3);
  });
});
