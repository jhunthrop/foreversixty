// web/src/lib/bis/page-state.test.ts
// A folded band (lib/bis/band-path.ts) has no path of its own: its hash opens it on the spec page.
import { describe, expect, it } from 'vitest';
import { parseState, stateSearch, type BisStateContext } from './page-state';

const context: BisStateContext = {
  bands: [20, 30, 40, 50, 60],
  pathBand: 20,
  presetsFor: () => ['bare'],
  racesFor: (faction) => (faction === 'horde' ? ['orc'] : ['human']),
  simRaceFor: (faction) => (faction === 'horde' ? 'orc' : 'human'),
};

describe('a folded band hash', () => {
  it('opens its band on the spec page, with the query for the rest', () => {
    const state = parseState('?faction=horde', '#levels-40-49', context);
    expect(state.band).toBe(40);
    expect(state.faction).toBe('horde');
  });

  it('keeps the band out of the query: the hash carries it', () => {
    const state = parseState('', '#levels-30-39', context);
    expect(state.band).toBe(30);
    expect(stateSearch(state, context)).toBe('');
  });

  it('leaves the path band for any other hash', () => {
    expect(parseState('', '#leveling', context).band).toBe(20);
  });
});
