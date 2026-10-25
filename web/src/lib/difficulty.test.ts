import { describe, expect, it } from 'vitest';
import { difficultyLabel } from './difficulty';

describe('difficultyLabel', () => {
  it.each([
    [3, '10 Player'],
    [4, '25 Player'],
    [5, '10 Heroic'],
    [6, '25 Heroic'],
    [9, '40 Player'],
    [148, '20 Player'],
    [186, '40 Player'],
  ])('names %i as %s', (id, label) => {
    expect(difficultyLabel(id)).toBe(label);
  });

  it('falls back to the id for one it cannot name', () => {
    expect(difficultyLabel(77)).toBe('Difficulty 77');
    expect(difficultyLabel(0)).toBe('Difficulty 0');
  });
});
