import { describe, expect, it } from 'vitest';
import generated from '../../data/generated/sim-ids.json';
import { simCopy } from './copy';
import { BUFF_GROUPS, CATALOGUE, IMPROVED_SUFFIX, buildCatalogue, gradeOf, rowsIn, setGrade } from './buffs';

describe('the catalogue', () => {
  it('groups every buff by the message IDS.md says it lands in', () => {
    const groups = new Map(CATALOGUE.map((row) => [row.id, row.group]));
    expect(groups.get('battle_shout')).toBe('raid-buffs');
    expect(groups.get('sunder_armor')).toBe('debuffs');
    expect(groups.get('atiesh_mage')).toBe('party-buffs');
    expect(groups.get('blessing_of_kings')).toBe('player-buffs');
  });

  it('puts the world buffs in their own section rather than among the blessings', () => {
    const groups = new Map(CATALOGUE.map((row) => [row.id, row.group]));
    for (const id of [
      'rallying_cry_of_the_dragonslayer',
      'songflower_serenade',
      'spirit_of_zandalar',
      'warchiefs_blessing',
      'fengus_ferocity',
      'moldars_moxie',
      'slipkiks_savvy',
      'sayges_fortune',
    ]) {
      expect(groups.get(id), id).toBe('world-buffs');
    }
  });

  it('groups consumables by the Consumes field they set', () => {
    const groups = new Map(CATALOGUE.map((row) => [row.id, row.group]));
    expect(groups.get('flask_of_supreme_power')).toBe('flask');
    expect(groups.get('elixir_of_the_mongoose')).toBe('battle-elixir');
    expect(groups.get('elixir_of_superior_defense')).toBe('guardian-elixir');
    expect(groups.get('food_grilled_squid')).toBe('food');
    expect(groups.get('main_hand_imbue:shadow_oil')).toBe('weapon-imbue');
    expect(groups.get('major_mana_potion')).toBe('potion');
    expect(groups.get('explosive_thorium_grenade')).toBe('explosive');
  });

  it('holds every id the engine publishes and invents none', () => {
    const published = new Set([
      ...generated.buffs.map((row) => row.id),
      ...generated.consumables.map((row) => row.id),
    ]);
    for (const row of CATALOGUE) {
      // A graded row is folded into its plain id and never appears on its own.
      expect(published.has(row.id), row.id).toBe(true);
    }
    for (const id of published) {
      if (id.endsWith(IMPROVED_SUFFIX)) continue;
      expect(
        CATALOGUE.some((row) => row.id === id),
        id,
      ).toBe(true);
    }
  });

  it('folds `<id>:improved` into the plain id as a grade, never as a second row', () => {
    const catalogue = buildCatalogue({
      buffs: [
        { id: 'battle_shout', message: 'RaidBuffs' },
        { id: 'battle_shout:improved', message: 'RaidBuffs' },
        { id: 'thorns', message: 'RaidBuffs' },
      ],
      consumables: [],
      professions: [],
      worldBuffs: [],
      stats: [],
    });
    expect(catalogue.map((row) => row.id)).toEqual(['battle_shout', 'thorns']);
    expect(catalogue.find((row) => row.id === 'battle_shout')?.graded).toBe(true);
    expect(catalogue.find((row) => row.id === 'thorns')?.graded).toBe(false);
  });

  it('names every group', () => {
    for (const group of BUFF_GROUPS) {
      expect(simCopy.buffGroupLabel[group], group).toBeTruthy();
    }
  });

  it('rowsIn is the group’s rows, alphabetical by id', () => {
    const flasks = rowsIn('flask').map((row) => row.id);
    expect(flasks.length).toBeGreaterThan(1);
    expect([...flasks].sort()).toEqual(flasks);
  });
});

describe('the three-way grade', () => {
  it('reads off, on and improved out of a selection', () => {
    expect(gradeOf([], 'battle_shout')).toBe('off');
    expect(gradeOf(['battle_shout'], 'battle_shout')).toBe('on');
    expect(gradeOf(['battle_shout:improved'], 'battle_shout')).toBe('improved');
  });

  it('never leaves both forms in the list', () => {
    expect(setGrade(['battle_shout'], 'battle_shout', 'improved')).toEqual(['battle_shout:improved']);
    expect(setGrade(['battle_shout:improved'], 'battle_shout', 'on')).toEqual(['battle_shout']);
    expect(setGrade(['battle_shout:improved'], 'battle_shout', 'off')).toEqual([]);
  });

  it('leaves every other id alone and its order intact', () => {
    expect(setGrade(['thorns', 'battle_shout', 'blood_pact'], 'battle_shout', 'off')).toEqual([
      'thorns',
      'blood_pact',
    ]);
  });

  it('does not mutate the list it was given', () => {
    const selected = ['battle_shout'];
    setGrade(selected, 'battle_shout', 'off');
    expect(selected).toEqual(['battle_shout']);
  });

  it('names the three grades', () => {
    expect(Object.keys(simCopy.gradeLabel).sort()).toEqual(['improved', 'off', 'on']);
  });
});
