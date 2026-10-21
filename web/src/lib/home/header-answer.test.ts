// web/src/lib/home/header-answer.test.ts
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../account/api';
import { homeAction, upgradesAnswer, upgradesNote } from './header-answer';
import type { CachedUpgrades } from './upgrades-loader';

function cached(overrides: Partial<CachedUpgrades['result']> = {}, band = 60): CachedUpgrades {
  return {
    specKey: 'warrior-fury',
    band,
    items: [],
    result: {
      upgrades: Array.from({ length: 9 }, () => ({}) as CachedUpgrades['result']['upgrades'][number]),
      alreadyBis: [],
      totalGainDps: 11.2,
      notSimCheckedCount: 0,
      scoreUnit: 'dps',
      ...overrides,
    } as CachedUpgrades['result'],
  };
}

const CHARACTER: MeCharacter = {
  key: 'us/normal/tester',
  region: 'us',
  ruleset: 'normal',
  name: 'Tester',
  class: 'Warrior',
  spec: 'Fury',
  level: 12,
  faction: 'horde',
};

describe('upgradesAnswer', () => {
  it('is the count and the gain as two figures at the same size', () => {
    expect(upgradesAnswer(cached())).toEqual({
      kind: 'figures',
      label: 'Upgrades for your level range',
      figures: [
        { value: '9', unit: 'upgrades' },
        { value: '+11.2', unit: 'DPS' },
      ],
      note: { text: 'Best in slot for 60, together' },
    });
  });

  it('says "upgrade" for one', () => {
    const answer = upgradesAnswer(cached({ upgrades: [{} as never] }));
    expect(answer.kind === 'figures' && answer.figures[0].unit).toBe('upgrade');
  });

  it('leaves the gain out when no slot has a measured gain, and says so in the note', () => {
    const answer = upgradesAnswer(cached({ totalGainDps: 0, notSimCheckedCount: 2 }, 20));
    expect(answer).toMatchObject({
      kind: 'figures',
      figures: [{ value: '9', unit: 'upgrades' }],
      note: { text: 'Best in slot for 20 to 29, together · 2 slots not sim-checked' },
    });
    expect(answer.kind === 'figures' && answer.figures).toHaveLength(1);
  });

  it('is one sentence, not a zero, when every slot already matches', () => {
    expect(upgradesAnswer(cached({ upgrades: [] }))).toEqual({
      kind: 'sentence',
      text: 'Every slot matches the 60 list.',
    });
  });

  it('carries the unit the band is scored in', () => {
    const answer = upgradesAnswer(cached({ scoreUnit: 'hps' }));
    expect(answer.kind === 'figures' && answer.figures[1]?.unit).toBe('HPS');
  });
});

describe('upgradesNote', () => {
  it('names the single slot left out in the singular', () => {
    expect(upgradesNote(60, 1)).toBe('Best in slot for 60, together · one slot not sim-checked');
  });
});

describe('homeAction', () => {
  it('jumps to the upgrades table', () => {
    expect(homeAction({ character: { ...CHARACTER, level: 60 }, belowLevelFloor: false })).toEqual({
      label: 'See your upgrades →',
      href: '#upgrades',
    });
  });

  it('under the level floor opens the first band’s list for the character’s own spec and faction', () => {
    expect(homeAction({ character: CHARACTER, belowLevelFloor: true })).toEqual({
      label: 'Fury gear for 20 to 29 →',
      href: '/bis/warrior/fury?faction=horde#band-horde-20',
    });
  });

  it('under the level floor with no spec points at the BiS index', () => {
    const bare = { ...CHARACTER, spec: undefined };
    expect(homeAction({ character: bare, belowLevelFloor: true })).toEqual({
      label: 'Pick a spec →',
      href: '/bis',
    });
  });
});
