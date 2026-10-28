import { describe, expect, it } from 'vitest';
import { FIXTURE_GUID } from '../../test-support/coach-fixtures';
import { coachAnchorHref } from './anchors';
import { chipsFor, clock, firstSpellId } from './chips';
import type { CoachAnchor } from './types';

const anchors: Record<string, CoachAnchor> = {
  a1: { kind: 'cast', spell_id: 23881 },
  a2: { kind: 'death', at_ms: 200000 },
  a3: { kind: 'time', at_ms: 65_000 },
  a4: { kind: 'aura', spell_id: 5, debuff: true },
  a5: { kind: 'resource', at_ms: 10_000, to_ms: 25_000 },
};
const context = {
  anchors,
  fightIndex: 3,
  firstFight: 1,
  playerGuid: FIXTURE_GUID,
  spellName: (id: number) => (id === 23881 ? 'Bloodthirst' : 'Sunder'),
};

describe('chips', () => {
  it('labels by the copy grammar and links through coachAnchorHref', () => {
    const chips = chipsFor(['a1', 'a2', 'a3', 'a4', 'a5'], context);
    expect(chips.map((chip) => chip.label)).toEqual([
      'Bloodthirst · Casts',
      'Death · 3:20',
      'Timelines · 1:05',
      'Sunder · Debuffs',
      'Resources · 0:10–0:25',
    ]);
    expect(chips[0]?.href).toBe(coachAnchorHref(3, 1, FIXTURE_GUID, anchors['a1']!));
  });
  it('a chip whose spell name does not resolve reads the kind word alone', () => {
    const unnamed = { ...context, spellName: () => '' };
    const healing = { h: { kind: 'healing', spell_id: 2060 } } as const;
    expect(chipsFor(['a1', 'a4'], unnamed).map((chip) => chip.label)).toEqual(['Casts', 'Debuffs']);
    expect(chipsFor(['h'], { ...unnamed, anchors: healing }).map((chip) => chip.label)).toEqual(['Healing']);
  });
  it('a death chip renders from at_ms, the shape the API emits, and links the rebuilt anchor', () => {
    const [chip] = chipsFor(['a2'], context);
    expect(chip?.label).toBe('Death · 3:20');
    expect(chip?.href).toBe(`?fight=3&tab=deaths#death-${FIXTURE_GUID}-200000`);
  });
  it('a legacy death string still labels', () => {
    const legacy = { d: { kind: 'death', death: `death-${FIXTURE_GUID}-200000` } } as const;
    expect(chipsFor(['d'], { ...context, anchors: legacy })[0]?.label).toBe('Death · 3:20');
  });
  it('an unknown anchor id renders no chip', () => {
    expect(chipsFor(['zz'], context)).toEqual([]);
  });
  it('formats clocks', () => {
    expect(clock(0)).toBe('0:00');
    expect(clock(232_000)).toBe('3:52');
  });
  it('picks the first spell anchor', () => {
    expect(firstSpellId(['a2', 'a1'], anchors)).toBe(23881);
    expect(firstSpellId(['a2'], anchors)).toBeNull();
  });
});
