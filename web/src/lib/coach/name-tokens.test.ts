import { describe, expect, it } from 'vitest';
import { buildDictionary, tokenizeNames } from './name-tokens';

const dictionary = buildDictionary(
  { '23881': 'Bloodthirst', '1': 'Heroic Strike', '2': 'Heroic' },
  { '19019': 'Thunderfury, Blessed Blade of the Windseeker' },
);

describe('tokenizeNames', () => {
  it('turns an exact spell name into a token and keeps punctuation outside', () => {
    expect(tokenizeNames('Bloodthirst: 9 casts lost', dictionary)).toEqual([
      { text: 'Bloodthirst', entry: { kind: 'spell', id: 23881 } },
      { text: ': 9 casts lost' },
    ]);
  });
  it('prefers the longest match', () => {
    const parts = tokenizeNames('Use Heroic Strike now', dictionary);
    expect(parts[1]).toEqual({ text: 'Heroic Strike', entry: { kind: 'spell', id: 1 } });
  });
  it('matches an item name that contains a comma', () => {
    const parts = tokenizeNames('Wield Thunderfury, Blessed Blade of the Windseeker.', dictionary);
    expect(parts.find((p) => 'entry' in p)).toEqual({
      text: 'Thunderfury, Blessed Blade of the Windseeker',
      entry: { kind: 'item', id: 19019 },
    });
  });
  it('leaves an unknown name and a partial word plain', () => {
    expect(tokenizeNames('Bloodthirsty Execute', dictionary)).toEqual([{ text: 'Bloodthirsty Execute' }]);
  });
  it('an empty dictionary is plain text', () => {
    expect(tokenizeNames('Bloodthirst', new Map())).toEqual([{ text: 'Bloodthirst' }]);
  });
});
