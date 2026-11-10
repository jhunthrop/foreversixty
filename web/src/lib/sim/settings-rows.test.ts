// web/src/lib/sim/settings-rows.test.ts
// Spec 2026-11-04-sim-results 3.12 and section 10, check 2.
import { describe, expect, it } from 'vitest';
import full from '../../fixtures/sim/result-full.json';
import {
  actionsPerMinute,
  chipName,
  settingsChipGroups,
  settingsRows,
  targetFact,
  type SettingsBuffs,
} from './settings-rows';
import type { SimResult } from './types';

const result = full as unknown as SimResult;
const BUFFS: SettingsBuffs = {
  arcane_brilliance: { name: 'Arcane Brilliance', icon: 'a', target: false, stats: { intellect: 31 } },
  blessing_of_might: { name: 'Blessing of Might', icon: 'b', target: false, stats: { attack_power: 112 } },
  sunder_armor: {
    name: 'Sunder Armor',
    icon: 's',
    target: true,
    stats: { target_armor: -2250 },
    exclusive_with: { group: 'major_armor', rank: 1 },
  },
  expose_armor: {
    name: 'Expose Armor',
    icon: 'e',
    target: true,
    stats: { target_armor: -2250 },
    exclusive_with: { group: 'major_armor', rank: 2 },
  },
  'main_hand_imbue:elemental_sharpening_stone': { name: 'Elemental Sharpening Stone', icon: 'm' },
  mana_spring_totem: { name: 'Mana Spring Totem', icon: 'n' },
};
const WEIGHTS = { attack_power: 1, strength: 1.9 };
const withBuffs = (buffs: string[]): SimResult => ({
  ...result,
  request: { ...result.request, character: { ...result.request.character, buffs } },
});

describe('settingsRows', () => {
  it('reads 40.4 actions per minute: 121 presses over 179.92 s', () => {
    expect(actionsPerMinute(result).toFixed(1)).toBe('40.4');
  });

  it('prints the fight, the target and the Execute phase in the spec words', () => {
    const all = settingsRows(result);
    const rows = Object.fromEntries(all.map((row) => [row.id, row.value]));
    const subs = Object.fromEntries(all.map((row) => [row.id, row.sub]));
    expect(rows.fight).toBe('3:00, ± 20%');
    expect(subs.fight).toBe('2:24 to 3:36');
    expect(rows.target).toBe('Level 63 boss');
    expect(subs.target).toBe('3,731 armor, 471 after debuffs');
    expect(rows.execute).toBe('Last 25% of the fight');
    // The share is the value ("Last 25% of the fight"); no sub restates it (h-audit 2026-11-09).
    expect(subs.execute).toBeUndefined();
    expect(rows.race).toBe('Troll');
    expect(rows.runs).toBe('10,000');
  });

  it('adds only the armor parts the run knows (build ruling R4)', () => {
    expect(
      targetFact({ ...result, target_stats: { level: 63, armor: 3731, armor_after_debuffs: 1471 } }),
    ).toEqual({ value: 'Level 63 boss', sub: '3,731 armor, 1,471 after debuffs' });
    expect(targetFact({ ...result, target_stats: { level: 62, armor: 3400 } })).toEqual({
      value: 'Level 62 boss',
      sub: '3,400 armor',
    });
  });
});

describe('the chips', () => {
  it('names the hand and the grade, never a raw key', () => {
    expect(chipName('main_hand_imbue:elemental_sharpening_stone', BUFFS)).toBe(
      'Elemental Sharpening Stone (main hand)',
    );
    expect(chipName('off_hand_imbue:dense_sharpening_stone', BUFFS)).toBe(
      'Dense sharpening stone (off hand)',
    );
    expect(chipName('mana_spring_totem:improved', BUFFS)).toBe('Mana Spring Totem, improved');
  });

  it('splits the sides, counts every chip and carries a suppressed debuff its reason', () => {
    const groups = settingsChipGroups(
      withBuffs(['arcane_brilliance', 'blessing_of_might', 'sunder_armor', 'expose_armor']),
      BUFFS,
      WEIGHTS,
      ['physical'],
      (stem) => `/${stem}.webp`,
    );
    const [player, target, consumables] = groups;
    expect(player?.count).toBe(2);
    expect(player?.chips.map((chip) => chip.id)).toEqual(['blessing_of_might']);
    expect(player?.noEffect.map((chip) => chip.id)).toEqual(['arcane_brilliance']);
    expect(target?.chips.map((chip) => chip.id)).toEqual(['sunder_armor', 'expose_armor']);
    expect(target?.chips[0]?.reason).toBeNull();
    expect(target?.chips[1]?.reason).toBe('replaced by Sunder Armor');
    expect(target?.chips[0]?.icon).toBe('/s.webp');
    expect(consumables?.count).toBe(8);
  });
});
