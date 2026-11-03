import { describe, expect, it } from 'vitest';
import {
  castText,
  cooldownText,
  costText,
  displayDescription,
  rangeText,
  spellTipModel,
  talentTipModel,
  tipFromBundle,
  type GuideTipBundle,
  type SpellTipRow,
} from './game-tip';

const MORTAL_STRIKE: SpellTipRow = {
  name: 'Mortal Strike',
  rank: 4,
  cost: 300,
  cost_type: 1,
  cost_pct: 0,
  cast_time_ms: 0,
  channeled: false,
  passive: false,
  cooldown_ms: 6000,
  range_yd: 5,
  description: 'A vicious strike that deals weapon damage plus 160.  Lasts 10 sec.',
};

describe('the client tooltip lines', () => {
  it('prints cost in the power the client names, rage in whole points', () => {
    expect(costText({ cost: 300, cost_type: 1, cost_pct: 0 })).toBe('30 Rage');
    expect(costText({ cost: 250, cost_type: 0, cost_pct: 0 })).toBe('250 Mana');
    expect(costText({ cost: 35, cost_type: 3, cost_pct: 0 })).toBe('35 Energy');
    expect(costText({ cost: 0, cost_type: 0, cost_pct: 12 })).toBe('12% of base mana');
    expect(costText({ cost: 0, cost_type: 0, cost_pct: 0 })).toBeUndefined();
  });

  it('prints range, cast and cooldown as the client does', () => {
    expect(rangeText(5)).toBe('Melee Range');
    expect(rangeText(30)).toBe('30 yd range');
    expect(rangeText(0)).toBeUndefined();
    expect(castText({ cast_time_ms: 0, channeled: false })).toBe('Instant');
    expect(castText({ cast_time_ms: 2500, channeled: false })).toBe('2.5 sec cast');
    expect(castText({ cast_time_ms: 0, channeled: true })).toBe('Channeled');
    expect(cooldownText(6000)).toBe('6 sec cooldown');
    expect(cooldownText(120_000)).toBe('2 min cooldown');
    expect(cooldownText(90_000)).toBe('1.5 min cooldown');
    expect(cooldownText(1500)).toBeUndefined();
  });

  it('builds the spell tooltip: rank, cost and range, cast and cooldown, the description', () => {
    expect(spellTipModel(MORTAL_STRIKE, 'ability_warrior_savageblow')).toEqual({
      name: 'Mortal Strike',
      icon: 'ability_warrior_savageblow',
      subtitle: 'Rank 4',
      lines: [
        { left: '30 Rage', right: 'Melee Range' },
        { left: 'Instant', right: '6 sec cooldown' },
      ],
      description: 'A vicious strike that deals weapon damage plus 160. Lasts 10 sec.',
    });
  });

  it('says Passive for a passive and nothing about casting it', () => {
    const model = spellTipModel({ ...MORTAL_STRIKE, passive: true }, undefined, 'Racial');
    expect(model.lines).toEqual([{ left: 'Passive' }]);
    expect(model.subtitle).toBe('Racial');
  });
});

describe('displayDescription', () => {
  it('reads a conditional for a character without the talent', () => {
    expect(displayDescription('Absorbs 162 damage$?s11094[ and reflects][].  Lasts 30 sec.')).toBe(
      'Absorbs 162 damage. Lasts 30 sec.',
    );
  });

  it('names both words of a gendered pair and agrees a plural with its number', () => {
    expect(displayDescription('the mage and $ghis:her; allies')).toBe('the mage and his or her allies');
    expect(displayDescription('The next 1 melee attack that $Lstrikes:strike; the caster')).toBe(
      'The next 1 melee attack that strikes the caster',
    );
  });
});

describe('talentTipModel', () => {
  const entry = {
    name: 'Cruelty',
    maxRank: 5,
    ranks: ['1% crit', '2% crit', '3% crit', '4% crit', '5% crit'],
  };

  it('shows the build rank and the next one', () => {
    expect(talentTipModel(entry, 2)).toEqual({
      name: 'Cruelty',
      subtitle: 'Rank 2/5',
      lines: [],
      description: '2% crit',
      nextRank: '3% crit',
    });
  });

  it('shows the first rank of an untaken talent and no next rank at the last', () => {
    expect(talentTipModel(entry, 0)).toMatchObject({ subtitle: 'Rank 0/5', description: '1% crit' });
    expect(talentTipModel(entry, 5).nextRank).toBeUndefined();
  });
});

describe('tipFromBundle', () => {
  const bundle: GuideTipBundle = {
    tips: { 'spell:1': spellTipModel(MORTAL_STRIKE, undefined) },
    talents: { '9': { name: 'Cruelty', maxRank: 5, ranks: ['a', 'b', 'c', 'd', 'e'] } },
    items: {},
  };

  it('reads a spell by key and a talent at the key rank', () => {
    expect(tipFromBundle(bundle, 'spell:1')?.name).toBe('Mortal Strike');
    expect(tipFromBundle(bundle, 'talent:9:3')?.description).toBe('c');
    expect(tipFromBundle(bundle, 'talent:8:1')).toBeNull();
    expect(tipFromBundle(bundle, 'spell:2')).toBeNull();
  });
});
