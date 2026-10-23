// web/src/lib/guides/gear-extras.test.ts
import activeBuild from '../../data/active-build.json';
import { readFileSync, readdirSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import {
  CONSUMABLE_GIVES,
  consumableKind,
  consumableRows,
  enchantGives,
  enchantScore,
  loadGearExtras,
  pickEnchants,
  type EnchantRecord,
} from './gear-extras';
import type { BisStatWeight } from '../bis/types';

const BUILD = activeBuild.build;
const REPO = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');

function enchant(overrides: Partial<EnchantRecord>): EnchantRecord {
  return {
    id: 1,
    name: 'Enchant X',
    icon: 'icon_x',
    slots: ['wrist'],
    item_types: ['normal'],
    classes: [],
    stats: {},
    phase: 0,
    ...overrides,
  };
}

const WEIGHTS: BisStatWeight[] = [
  { stat: 'spell_power', weight: 1 },
  { stat: 'intellect', weight: 0.2 },
  { stat: 'stamina', weight: 0.5, insignificant: true },
];

const CONTEXT = { classSlug: 'mage', twoHanded: false, shield: false, hasOffHand: false };

describe('enchantGives', () => {
  it('words flat stats in the weights’ own terms', () => {
    expect(enchantGives({ spell_damage: 30 })).toBe('+30 spell power');
    expect(enchantGives({ intellect: 7 })).toBe('+7 intellect');
    expect(enchantGives({ healing: 55 })).toBe('+55 healing power');
  });

  it('collapses an all-stats enchant', () => {
    expect(enchantGives({ strength: 4, agility: 4, stamina: 4, intellect: 4, spirit: 4 })).toBe(
      '+4 to all stats',
    );
  });
});

describe('enchantScore', () => {
  const byStat = new Map([
    ['spell_power', 1],
    ['intellect', 0.2],
  ]);

  it('adds up value times weight, through the aliases', () => {
    expect(enchantScore({ spell_damage: 30 }, byStat)).toBe(30);
    expect(enchantScore({ intellect: 10 }, byStat)).toBeCloseTo(2);
  });

  it('does not score crit, hit or haste, whose weights are per rating point', () => {
    expect(
      enchantScore(
        { crit: 2, hit: 2, spell_haste: 3 },
        new Map([
          ['crit', 5],
          ['hit', 5],
        ]),
      ),
    ).toBe(0);
  });
});

describe('pickEnchants', () => {
  it('picks the best-scoring enchant for each slot, in slot order, skipping what scores nothing', () => {
    const rows = pickEnchants(
      [
        enchant({ id: 1, name: 'Weak Bracer', slots: ['wrist'], stats: { intellect: 3 } }),
        enchant({ id: 2, name: 'Strong Bracer', slots: ['wrist'], stats: { intellect: 7 } }),
        enchant({ id: 3, name: 'Weapon SP', slots: ['main_hand', 'off_hand'], stats: { spell_damage: 30 } }),
        enchant({ id: 4, name: 'Boots Nothing', slots: ['feet'], stats: { strength: 5 } }),
      ],
      WEIGHTS,
      CONTEXT,
    );
    expect(rows.map((row) => [row.slotLabel, row.name, row.gives])).toEqual([
      ['Wrist', 'Strong Bracer', '+7 intellect'],
      ['Weapon', 'Weapon SP', '+30 spell power'],
    ]);
  });

  it('leaves out a later phase, an armor kit and another class’s enchant', () => {
    const rows = pickEnchants(
      [
        enchant({ id: 1, name: 'Phase Four', stats: { intellect: 9 }, phase: 4 }),
        enchant({ id: 2, name: 'Kit', stats: { intellect: 9 }, item_types: ['kit'] }),
        enchant({ id: 3, name: 'Warrior Only', stats: { intellect: 9 }, classes: ['warrior'] }),
        enchant({ id: 4, name: 'Mage Only', stats: { intellect: 1 }, classes: ['mage'] }),
      ],
      WEIGHTS,
      CONTEXT,
    );
    expect(rows.map((row) => row.name)).toEqual(['Mage Only']);
  });

  it('offers a two-hand weapon enchant only to a two-hander and an off-hand row only with an off hand', () => {
    const records = [
      enchant({
        id: 1,
        name: 'Two Hand',
        slots: ['main_hand', 'off_hand'],
        item_types: ['two_hand'],
        stats: { intellect: 20 },
      }),
      enchant({ id: 2, name: 'One Hand', slots: ['main_hand', 'off_hand'], stats: { intellect: 5 } }),
    ];
    expect(pickEnchants(records, WEIGHTS, CONTEXT).map((row) => row.name)).toEqual(['One Hand']);
    expect(pickEnchants(records, WEIGHTS, { ...CONTEXT, twoHanded: true }).map((row) => row.name)).toEqual([
      'Two Hand',
    ]);
    const withOffHand = pickEnchants(records, WEIGHTS, { ...CONTEXT, hasOffHand: true });
    expect(withOffHand.map((row) => row.slot)).toEqual(['main_hand', 'off_hand']);
  });

  it('ignores a weight the sweep marked insignificant', () => {
    const rows = pickEnchants(
      [enchant({ id: 1, name: 'Stamina Bracer', stats: { stamina: 9 } })],
      WEIGHTS,
      CONTEXT,
    );
    expect(rows).toEqual([]);
  });
});

describe('consumableRows', () => {
  const simBuffs = {
    flask_of_supreme_power: { name: 'Flask', icon: 'inv_potion_41' },
    juju_power: { name: 'Juju', icon: 'temp' },
  };

  it('joins the preset’s consumes to their line and their icon', () => {
    const rows = consumableRows(
      [
        { id: 'flask_of_supreme_power', label: 'Flask of Supreme Power' },
        { id: 'juju_power', label: 'Juju Power' },
        { id: 'not_in_the_table', label: 'Mystery' },
      ],
      simBuffs,
      (icon) => icon !== 'temp',
    );
    expect(rows).toEqual([
      {
        id: 'flask_of_supreme_power',
        kind: 'flask',
        name: 'Flask of Supreme Power',
        icon: 'inv_potion_41',
        gives: 'Spell power',
      },
      { id: 'juju_power', kind: 'other', name: 'Juju Power', icon: undefined, gives: 'Strength' },
    ]);
  });
});

describe('consumableKind', () => {
  it('tells a flask from an elixir from everything else', () => {
    expect(consumableKind('flask_of_supreme_power')).toBe('flask');
    expect(consumableKind('greater_arcane_elixir')).toBe('elixir');
    expect(consumableKind('elixir_of_frost_power')).toBe('elixir');
    expect(consumableKind('major_mana_potion')).toBe('other');
  });
});

describe('enchant icons', () => {
  it('drops an icon the data shares among many enchants, keeps one an enchant owns', () => {
    const shared = Array.from({ length: 4 }, (_, i) =>
      enchant({ id: i + 10, name: `Stand-in ${i}`, icon: 'generic', stats: { intellect: 1 } }),
    );
    const own = enchant({ id: 1, name: 'Own', icon: 'own_icon', slots: ['feet'], stats: { intellect: 5 } });
    const rows = pickEnchants([...shared, own], WEIGHTS, CONTEXT);
    expect(rows.find((row) => row.name === 'Own')?.icon).toBe('own_icon');
    expect(rows.find((row) => row.slot === 'wrist')?.icon).toBeUndefined();
  });
});

describe('consumableRows with a repeated consumable', () => {
  it('lists it once', () => {
    const rows = consumableRows(
      [
        { id: 'major_mana_potion', label: 'Major Mana Potion' },
        { id: 'major_mana_potion', label: 'Major Mana Potion' },
      ],
      {},
      () => false,
    );
    expect(rows.map((row) => row.name)).toEqual(['Major Mana Potion']);
  });
});

describe('the consumable lines against the published presets', () => {
  it('has a line for every consumable any spec’s raid-ready preset uses', () => {
    const dir = path.join(REPO, 'data/builds', BUILD, 'bis');
    const missing = new Set<string>();
    for (const name of readdirSync(dir).filter((file) => file.endsWith('.json'))) {
      const file = JSON.parse(readFileSync(path.join(dir, name), 'utf8')) as {
        presets?: { raid?: { consumes: { id: string }[] } };
      };
      for (const consume of file.presets?.raid?.consumes ?? []) {
        if (CONSUMABLE_GIVES[consume.id] === undefined) missing.add(consume.id);
      }
    }
    expect([...missing]).toEqual([]);
  });
});

describe('loadGearExtras for Frost Mage, from the real data', () => {
  const extras = loadGearExtras(BUILD, 'mage', 'mage-frost', ['gnome', 'troll'], {});

  it('picks a weapon enchant for spell power and lists the preset’s consumables', () => {
    if (extras === undefined) throw new Error('mage-frost has no ranked file');
    expect(extras.enchants.find((row) => row.slot === 'main_hand')?.gives).toBe('+30 spell power');
    const names = extras.consumables.map((row) => row.name);
    expect(names).toContain('Flask of Supreme Power');
    expect(names).toContain('Elixir of Frost Power');
  });

  it('is undefined for a spec with no ranked file', () => {
    expect(loadGearExtras(BUILD, 'mage', 'mage-nonexistent', ['gnome'], {})).toBeUndefined();
  });
});
