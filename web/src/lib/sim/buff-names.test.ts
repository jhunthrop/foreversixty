import { describe, expect, it } from 'vitest';
import { EMPTY_BUFF_NAMES, buffIcon, buffLabel } from './buff-names';

const names = {
  entries: {
    battle_shout: { name: 'Battle Shout', icon: 'ability_warrior_battleshout' },
    'main_hand_imbue:shadow_oil': { name: 'Shadow Oil', icon: 'inv_potion_21' },
  },
};

describe('buffLabel', () => {
  it('is the build’s own name when it has one', () => {
    expect(buffLabel('battle_shout', names)).toBe('Battle Shout');
    expect(buffLabel('main_hand_imbue:shadow_oil', names)).toBe('Shadow Oil');
  });

  it('humanises the id when the build has no table, rather than showing “Unknown”', () => {
    expect(buffLabel('flask_of_supreme_power', EMPTY_BUFF_NAMES)).toBe('Flask of supreme power');
    expect(buffLabel('flask_of_supreme_power', null)).toBe('Flask of supreme power');
  });

  it('humanises a qualified id by its own half, keeping the qualifier', () => {
    expect(buffLabel('off_hand_imbue:frost_oil', EMPTY_BUFF_NAMES)).toBe('Off hand imbue: Frost oil');
  });

  it('leaves a client item id alone: there is nothing to humanise', () => {
    expect(buffLabel('item:13452', EMPTY_BUFF_NAMES)).toBe('item:13452');
  });
});

describe('buffIcon', () => {
  it('is the build’s icon path when the table names one', () => {
    expect(buffIcon('1.15.9', 'battle_shout', names)).toBe(
      '/data/1.15.9/icons/ability_warrior_battleshout.webp',
    );
  });

  it('is null when there is no icon, so the row renders without one', () => {
    expect(buffIcon('1.15.9', 'thorns', names)).toBeNull();
    expect(buffIcon('1.15.9', 'thorns', null)).toBeNull();
  });
});
