// web/src/lib/tiers/tier-copy.test.ts
import { describe, expect, it } from 'vitest';
import {
  formatMetric,
  gapText,
  howToReadBullets,
  ordinal,
  tierNotes,
  tiersCopy,
  type NoteSegment,
} from './tier-copy';

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
    expect(dps).toContain('≈ tie marks a spec within 1% of the one above.');
    expect(dps).toContain(
      'S is within 5% of the top, A within 10%, B within 20%, C within 30%, D beyond; a ≈ tie across a line is a tie.',
    );
  });
});

describe('how to read this', () => {
  it('carries the spec bullets, with the preset and the fight length from the data and constants', () => {
    const dps = howToReadBullets('dps', 'Raid-ready, Phase 1');
    expect(dps).toEqual([
      'Every spec is simmed in its best Raid-ready, Phase 1 gear with raid buffs and consumables.',
      'One target for 180 seconds. Cleave, add fights, movement and what a spec brings the raid are not counted, so a spec can sit lower here than it plays in your raid.',
    ]);
    for (const role of ['tank', 'healer'] as const) {
      expect(howToReadBullets(role, 'x')[1]).toContain('Same boss profile as the BiS pages');
    }
  });

  it('is not repeated word for word in the notes above the list', () => {
    for (const role of ['dps', 'tank', 'healer'] as const) {
      const above = tierNotes({ ...CTX, role, healSeconds: 300, bossSwingSeconds: 2 })
        .map(flat)
        .join(' ');
      for (const bullet of howToReadBullets(role, CTX.presetLabel)) expect(above).not.toContain(bullet);
    }
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

describe('formatMetric', () => {
  it('writes tank figures whole and damage and healing rates to one decimal', () => {
    expect(formatMetric('tank', 322.14)).toBe('322');
    expect(formatMetric('tank', 35711.4)).toBe('35,711');
    expect(formatMetric('dps', 921.44)).toBe('921.4');
    expect(formatMetric('healer', 1204)).toBe('1,204.0');
  });
});

describe('tier bands', () => {
  it('words each band from its bounds', () => {
    expect(tiersCopy.tierBand({ from: 0, to: 5 })).toBe('Within 5% of the top');
    expect(tiersCopy.tierBand({ from: 5, to: 10 })).toBe('5 to 10% behind');
    expect(tiersCopy.tierBand({ from: 30, to: null })).toBe('30% or more behind');
  });

  it('counts specs in the singular and the plural', () => {
    expect(tiersCopy.tierCount(1)).toBe('1 spec');
    expect(tiersCopy.tierCount(4)).toBe('4 specs');
  });
});
