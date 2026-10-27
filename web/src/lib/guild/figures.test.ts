import { describe, expect, it } from 'vitest';
import { figureParts } from './figures';

describe('figureParts', () => {
  it('splits numerals from the words around them', () => {
    expect(figureParts('2 h 41 min')).toEqual([
      { text: '2', figure: true },
      { text: ' h ', figure: false },
      { text: '41', figure: true },
      { text: ' min', figure: false },
    ]);
  });

  it('keeps a kill time and a signed change whole', () => {
    expect(figureParts('Onyxia 4:12')).toEqual([
      { text: 'Onyxia ', figure: false },
      { text: '4:12', figure: true },
    ]);
    expect(figureParts('up −12')).toEqual([
      { text: 'up ', figure: false },
      { text: '−12', figure: true },
    ]);
  });

  it('reads a line with no numerals as words', () => {
    expect(figureParts('Excused')).toEqual([{ text: 'Excused', figure: false }]);
  });
});
