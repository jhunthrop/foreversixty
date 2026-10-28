import { describe, expect, it } from 'vitest';
import { possessive } from './possessive';

describe('possessive', () => {
  it.each([
    ['OLYMPUS XXVII', "OLYMPUS XXVII's"],
    ['Employed Millennials', "Employed Millennials'"],
    ['THE BOYS', "THE BOYS'"],
    ['Thornwick', "Thornwick's"],
  ])('%s', (name, expected) => {
    expect(possessive(name)).toBe(expected);
  });
});
