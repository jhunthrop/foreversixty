import { describe, expect, it } from 'vitest';
import { guidesCopy } from './copy';

describe('the /guides index line', () => {
  it('is true on launch day: no beta cap, no dashes, no talk about the site', () => {
    expect(guidesCopy.indexLine).not.toMatch(/beta|level 30/i);
    expect(guidesCopy.indexLine).not.toMatch(/[—–]|\s-\s/);
    expect(guidesCopy.indexLine).not.toMatch(/simulat|this site|\bsim\b/i);
  });
});

describe('guidesCopy', () => {
  it('has no blank strings', () => {
    for (const [key, value] of Object.entries(guidesCopy)) {
      if (typeof value !== 'string') continue;
      expect(value.trim(), `${key} is blank`).not.toBe('');
    }
  });

  it('has no blank-returning copy functions', () => {
    for (const [key, value] of Object.entries(guidesCopy)) {
      if (typeof value !== 'function') continue;
      const sample = (value as (...args: (string | number)[]) => string)(1, 'Example');
      expect(sample.trim(), `${key} returns blank`).not.toBe('');
    }
  });
});

describe('the stat table per-point column', () => {
  it('says "or rating" when the table prices a rating stat, and not otherwise', () => {
    expect(guidesCopy.perPointColumn('dps', true)).toBe('DPS per point or rating');
    expect(guidesCopy.perPointColumn('dps', false)).toBe('DPS per point');
    expect(guidesCopy.perPointColumn('healer', true)).toBe('HPS per point or rating');
    expect(guidesCopy.perPointColumn('tank', false)).toBe('Tank score per point');
  });
});

describe('the stat table unit line', () => {
  it('reads a plain stat per point', () => {
    expect(guidesCopy.unitLine('spell power', '0.39', 'dps')).toBe(
      '1 point of spell power = 0.39 DPS at level 60',
    );
  });

  it('names a rating stat as rating', () => {
    expect(guidesCopy.unitLine('hit', '0.46', 'dps', true)).toBe('1 hit rating = 0.46 DPS at level 60.');
  });

  it('gives every rating stat in the table its rating per 1%', () => {
    const hit = guidesCopy.ratingConversion('hit', 10);
    const crit = guidesCopy.ratingConversion('critical strike', 14);
    expect(guidesCopy.ratingNote('Hit', false, hit)).toBe(
      'Hit is per rating point: 10 hit rating is 1% hit.',
    );
    expect(guidesCopy.ratingNote('Hit and Critical strike', true, `${hit} and ${crit}`)).toBe(
      'Hit and Critical strike are per rating point: 10 hit rating is 1% hit and 14 critical strike rating is 1% critical strike.',
    );
  });
});

describe('the BiS caption', () => {
  it('takes "an" before a vowel and "a" otherwise', () => {
    expect(guidesCopy.bisCaption('Undead', 'Horde')).toBe('Picks for an Undead (Horde).');
    expect(guidesCopy.bisCaption('Orc', 'Horde')).toBe('Picks for an Orc (Horde).');
    expect(guidesCopy.bisCaption('Human', 'Alliance')).toBe('Picks for a Human (Alliance).');
  });
});
