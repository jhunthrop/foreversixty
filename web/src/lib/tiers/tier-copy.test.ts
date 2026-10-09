// web/src/lib/tiers/tier-copy.test.ts
import { describe, expect, it } from 'vitest';
import { gapText, howToReadBullets, ordinal, tierNotes, tiersCopy, type NoteSegment } from './tier-copy';

const flat = (note: NoteSegment[]): string => note.map((segment) => segment.text).join('');
const CTX = { faction: 'horde', presetLabel: 'Raid-ready, Phase 1' } as const;

describe('ordinal', () => {
  it.each([
    [1, '1st'],
    [2, '2nd'],
    [3, '3rd'],
    [4, '4th'],
    [11, '11th'],
    [12, '12th'],
    [21, '21st'],
    [112, '112th'],
  ])('%i is %s', (n, expected) => {
    expect(ordinal(n)).toBe(expected);
  });
});

describe('tierNotes', () => {
  it('gives each role three notes, naming the faction race and linking the leveling bands', () => {
    for (const role of ['dps', 'tank', 'healer'] as const) {
      const notes = tierNotes({ ...CTX, role, healSeconds: 300, bossSwingSeconds: 2 });
      expect(notes).toHaveLength(3);
      const text = notes.map(flat).join(' ');
      expect(text).toContain('best Horde race');
      expect(text).toContain('Raid-ready, Phase 1');
      expect(notes.flat().some((s) => s.href === '/bis')).toBe(true);
    }
  });

  it('renders the healer duration and the tank swing from the data it is given', () => {
    expect(flat(tierNotes({ ...CTX, role: 'healer', healSeconds: 240 })[0]!)).toContain('over 240 seconds');
    expect(flat(tierNotes({ ...CTX, role: 'tank', bossSwingSeconds: 2.5 })[1]!)).toContain(
      'every 2.5 seconds',
    );
  });

  it('fails fast when a profile figure is missing', () => {
    expect(() => tierNotes({ ...CTX, role: 'healer' })).toThrow(/heal profile duration/);
    expect(() => tierNotes({ ...CTX, role: 'tank' })).toThrow(/boss swing speed/);
  });

  it('states the tie rule once, on the DPS list', () => {
    const dps = tierNotes({ ...CTX, role: 'dps' })
      .map(flat)
      .join(' ');
    expect(dps).toContain('≈ tie marks two specs within 1% of each other.');
  });
});

describe('how to read this', () => {
  it('keeps only lines the notes above do not say', () => {
    const above = tierNotes({ ...CTX, role: 'dps' })
      .map(flat)
      .join(' ');
    for (const bullet of howToReadBullets('dps')) expect(above).not.toContain(bullet);
    expect(howToReadBullets('dps')).toHaveLength(1);
    expect(howToReadBullets('tank')[0]).toContain('Same boss profile');
  });
});

describe('gapText', () => {
  it('reads Top or Least on the leader, a signed gap below', () => {
    expect(gapText({ rank: 1, role: 'dps', gapPercent: 0 })).toBe(tiersCopy.topGap);
    expect(gapText({ rank: 1, role: 'tank', gapPercent: 0 })).toBe(tiersCopy.leastGap);
    expect(gapText({ rank: 2, role: 'dps', gapPercent: 4.84 })).toBe('−4.8%');
    expect(gapText({ rank: 2, role: 'tank', gapPercent: 15.23 })).toBe('+15.2%');
  });
});

describe('tie titles', () => {
  it('names the neighbour a row ties with', () => {
    expect(tiersCopy.tieTitle('above')).toBe('Within 1% of the spec above: a tie');
    expect(tiersCopy.tieTitle('below')).toContain('below');
    expect(tiersCopy.tieTitle('both')).toContain('above and below');
  });
});
