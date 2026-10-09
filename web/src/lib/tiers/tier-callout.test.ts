// web/src/lib/tiers/tier-callout.test.ts
import { describe, expect, it } from 'vitest';
import { dpsInput, tankInput } from './tier-test-support';
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
      detail: 'Level with the top spec.',
      bisHref: '/bis/warrior/fury',
      rolePath: '/tiers',
    });
  });

  it('names the top spec and the gap for a DPS spec behind it', () => {
    expect(views['warrior/arms']!.detail).toBe('4.8% behind Fury Warrior.');
    expect(views['mage/fire']!.title).toBe('Fire Mage is 3rd of 3 DPS specs.');
  });

  it('words a tank in damage taken', () => {
    expect(views['paladin/protection']!.detail).toBe('Takes the least damage.');
    expect(views['warrior/protection']).toMatchObject({
      title: 'Protection Warrior is 2nd of 2 Tank specs.',
      detail: '15.0% more damage taken than Protection Paladin.',
      pointer: 'Protection Warrior is on the Tank list. See where it stands →',
      rolePath: '/tiers/tank',
    });
  });

  it('colours with the class token and the crest path', () => {
    expect(views['mage/fire']).toMatchObject({
      colorVar: 'var(--color-class-mage)',
      crestSrc: '/icons/hd/crests/mage.webp',
    });
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
