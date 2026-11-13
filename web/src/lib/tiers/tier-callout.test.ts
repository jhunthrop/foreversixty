// web/src/lib/tiers/tier-callout.test.ts
import { describe, expect, it } from 'vitest';
import { dpsInput, healerInput, tankInput } from './tier-test-support';
import { rankRole } from './tier-list';
import { specViewKey, specViewKeyForCharacter, tierSpecViews } from './tier-callout';

const dps = rankRole(
  [
    dpsInput('warrior', 'fury', 'Fury', 100),
    dpsInput('warrior', 'arms', 'Arms', 95.2),
    dpsInput('mage', 'fire', 'Fire', 80),
  ],
  'dps',
);
const tank = rankRole(
  [
    tankInput('paladin', 'protection', 'Protection', { dtps: 300, effective_health: 1, tps: 1 }),
    tankInput('warrior', 'protection', 'Protection', { dtps: 345, effective_health: 1, tps: 1 }),
  ],
  'tank',
);
const views = tierSpecViews({ dps, tank, healer: [] });

describe('tierSpecViews', () => {
  it('keys a view by class and spec slug', () => {
    expect(specViewKey('warrior', 'fury')).toBe('warrior/fury');
    expect(Object.keys(views).sort()).toEqual([
      'mage/fire',
      'paladin/protection',
      'warrior/arms',
      'warrior/fury',
      'warrior/protection',
    ]);
  });

  it('says the top DPS spec is level with itself', () => {
    expect(views['warrior/fury']).toMatchObject({
      title: 'Fury Warrior is 1st of 3 DPS specs.',
      standing: 'Tier S · #1 · 100.0 DPS',
      detail: 'Level with the top spec.',
      bisHref: '/bis/warrior/fury/60',
      guideHref: '/guides/warrior/fury',
      rolePath: '/tiers',
    });
  });

  it('names the tier, the gap in percent and the top spec for a DPS spec behind it', () => {
    expect(views['warrior/arms']!.detail).toBe('Tier S · 4.8% behind Fury Warrior');
    expect(views['mage/fire']!.detail).toBe('Tier C · 20.0% behind Fury Warrior');
    expect(views['mage/fire']!.title).toBe('Fire Mage is 3rd of 3 DPS specs.');
  });

  it('words a tank in damage taken', () => {
    expect(views['paladin/protection']!.detail).toBe('Takes the least damage.');
    expect(views['warrior/protection']).toMatchObject({
      title: 'Protection Warrior is 2nd of 2 Tank specs.',
      standing: '#2 · 345 taken / s',
      detail: 'Takes 15% more damage a second than Protection Paladin',
      rolePath: '/tiers/tank',
    });
  });

  it('gives a healer whose list is not ranked yet its figure alone, and a tie its shared place', () => {
    const unranked = tierSpecViews({
      dps: [],
      tank: [],
      healer: rankRole(
        [healerInput('druid', 'restoration', 'Restoration', 700), healerInput('priest', 'holy', 'Holy', 600)],
        'healer',
      ),
    });
    expect(unranked['priest/holy']!.title).toBe(
      'Holy Priest is on the Healer list; ranks appear once the next ranking publishes the measured error.',
    );
    const bounds = (hps: number) => [
      { scale: 1.5, hps, hps_error: 1 },
      { scale: 1.65, hps, hps_error: 1 },
    ];
    const tied = tierSpecViews({
      dps: [],
      tank: [],
      healer: rankRole(
        [
          healerInput('druid', 'restoration', 'Restoration', 700, 1, bounds(700)),
          healerInput('priest', 'holy', 'Holy', 600, 1, bounds(600)),
          healerInput('shaman', 'restoration', 'Restoration', 599.5, 1, bounds(599.5)),
        ],
        'healer',
      ),
    });
    expect(tied['priest/holy']!.title).toBe('Holy Priest is tied 2nd of 3 Healer specs.');
    expect(tied['priest/holy']).toMatchObject({ standing: '=2 · 600.0 HPS', detail: 'Tied' });
    expect(unranked['priest/holy']!.standing).toBe('600.0 HPS');
  });

  it('skips an empty role', () => {
    expect(Object.values(views).some((v) => v.role === 'healer')).toBe(false);
  });
});

describe('specViewKeyForCharacter', () => {
  it('slugs the character class and spec names', () => {
    expect(specViewKeyForCharacter({ class: 'Warrior', spec: 'Fury' })).toBe('warrior/fury');
    expect(specViewKeyForCharacter({ class: 'Druid', spec: 'Feral Bear' })).toBe('druid/feral-bear');
  });

  it('is undefined until both are known', () => {
    expect(specViewKeyForCharacter({ class: 'Warrior' })).toBeUndefined();
    expect(specViewKeyForCharacter({ spec: 'Fury' })).toBeUndefined();
  });
});
