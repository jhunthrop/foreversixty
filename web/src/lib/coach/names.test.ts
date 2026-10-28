// web/src/lib/coach/names.test.ts
import { describe, expect, it } from 'vitest';
import { EMPTY_NAMES, withAnchorNames } from './names';
import { tokenizeNames } from './name-tokens';

describe('withAnchorNames', () => {
  it("adds the answer's anchor spell names to the dictionary and the chip names", () => {
    const names = withAnchorNames(EMPTY_NAMES, {
      a1: { kind: 'cast', spell_id: 23894, spell_name: 'Bloodthirst' },
    });
    expect(names.spellName(23894)).toBe('Bloodthirst');
    expect(tokenizeNames('Press Bloodthirst.', names.dictionary)).toEqual([
      { text: 'Press ' },
      { text: 'Bloodthirst', entry: { kind: 'spell', id: 23894 } },
      { text: '.' },
    ]);
  });
  it('leaves the tables alone when no anchor names a spell', () => {
    expect(withAnchorNames(EMPTY_NAMES, { a1: { kind: 'buffs' } })).toBe(EMPTY_NAMES);
  });
});
