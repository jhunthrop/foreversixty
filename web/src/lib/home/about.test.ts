import { describe, expect, it } from 'vitest';
import { SPECS } from '../sim/specs';
import { aboutView, listParts, type AboutFacts, type Part } from './about';

const FACTS: AboutFacts = {
  build: '1.60.1.70790',
  readAt: '2026-11-11T07:32:30Z',
  bisRankedAt: '2026-11-10T06:00:00Z',
  bisSpecs: new Set(SPECS.map((spec) => spec.spec)),
  bands: { first: 20, last: 60 },
  raidsOpen: '2026-12-09T00:00:00.000Z',
  ranked: [
    { name: 'Saltspine', slug: 'saltspine' },
    { name: 'Relic Guardian', slug: 'relic-guardian' },
  ],
};

const prose = (parts: readonly Part[]) => parts.map((part) => part.text).join('');

describe('aboutView', () => {
  it('reads every fact from the data: bands, build, read day, ranking day, raid day', () => {
    const text = aboutView(FACTS).paragraphs.map(prose).join('\n');
    expect(text).toContain('each ten-level band from 20 to 60');
    expect(text).toContain('client build 1.60.1.70790, read from the game files on 11 Nov.');
    expect(text).toContain('BiS lists ranked 10 Nov.');
    expect(text).toContain("opens on 9 Dec with Barrow Deeps, Hyjal Summit and Onyxia's Lair.");
  });

  it('links the most logged bosses to their boards, and says nothing of them with none ranked', () => {
    const raid = aboutView(FACTS).paragraphs[2] ?? [];
    expect(raid.filter((part) => part.href?.startsWith('/rankings/')).map((part) => part.href)).toEqual([
      '/rankings/saltspine',
      '/rankings/relic-guardian',
    ]);
    expect(prose(aboutView({ ...FACTS, ranked: [] }).paragraphs[2] ?? [])).not.toContain('rankings');
  });

  it('says the tier is open once its day has passed', () => {
    expect(prose(aboutView({ ...FACTS, raidsOpen: null }).paragraphs[2] ?? [])).toMatch(
      /^The first raid tier is open: Barrow Deeps/,
    );
  });

  it('lists every spec with a BiS file once, by role, in its class colour, linking its BiS page', () => {
    const roles = aboutView(FACTS).roles;
    expect(roles.map((group) => group.label)).toEqual(['Tanks', 'Healers', 'Damage']);
    const specs = roles.flatMap((group) => group.specs);
    expect(specs).toHaveLength(SPECS.length);
    expect(specs).toContainEqual({ text: 'Fury Warrior', href: '/bis/warrior/fury', classSlug: 'warrior' });
    expect(aboutView({ ...FACTS, bisSpecs: new Set(['warrior-fury']) }).roles).toEqual([
      {
        role: 'dps',
        label: 'Damage',
        specs: [{ text: 'Fury Warrior', href: '/bis/warrior/fury', classSlug: 'warrior' }],
      },
    ]);
  });

  it('writes no em dash and no definitional blurb', () => {
    const text = aboutView(FACTS).paragraphs.map(prose).join(' ');
    expect(text).not.toMatch(/[—–]/);
    expect(text).not.toMatch(/\b(is|are) the [^.]*:/);
  });
});

describe('listParts', () => {
  it('joins with commas and a final "and"', () => {
    expect(prose(listParts([{ text: 'a' }, { text: 'b' }, { text: 'c' }]))).toBe('a, b and c');
    expect(prose(listParts([{ text: 'a' }]))).toBe('a');
  });
});
