// web/src/lib/guides/gear-extras.test.ts
import activeBuild from '../../data/active-build.json';
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import {
  CONSUMABLE_GIVES,
  KIT_IMBUES,
  consumableKind,
  consumableLines,
  consumableRows,
  raidConsumeIds,
  simRaidConsumes,
  enchantGives,
  enchantScore,
  loadGearExtras,
  pickEnchants,
  type ConsumableRow,
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

const RAID_CONSUMES = JSON.parse(readFileSync(path.join(REPO, 'sim/request/raid-consumes.json'), 'utf8')) as {
  specs: Record<string, { from_level: number; consumes: string[] }[]>;
};
const SPECS = Object.keys(RAID_CONSUMES.specs).sort();

describe('the consumable lines against the sim’s raid preset', () => {
  it('has a line for every consumable any spec’s level 60 raid preset carries', () => {
    const missing = new Set(
      SPECS.flatMap((spec) => raidConsumeIds(RAID_CONSUMES, spec, 60)).filter(
        (id) => !KIT_IMBUES.has(id) && CONSUMABLE_GIVES[id] === undefined,
      ),
    );
    expect([...missing]).toEqual([]);
  });

  it('has a line for every consumable the curated raid preset names, before a refresh publishes it', () => {
    const curated = JSON.parse(readFileSync(path.join(REPO, 'data/curated/presets.json'), 'utf8')) as {
      presets: {
        raid: {
          consumes: {
            groups: Record<string, { ids: { id: string }[]; dual_wield_ids?: { id: string }[] }>;
            by_class: Record<string, { id: string }[]>;
            by_spec: Record<string, { id: string }[]>;
          };
        };
      };
    };
    const consumes = curated.presets.raid.consumes;
    const entries = [
      ...Object.values(consumes.groups).flatMap((group) => [...group.ids, ...(group.dual_wield_ids ?? [])]),
      ...Object.values(consumes.by_class).flat(),
      ...Object.values(consumes.by_spec).flat(),
    ];
    const missing = new Set(
      entries.map((entry) => entry.id).filter((id) => CONSUMABLE_GIVES[id] === undefined),
    );
    expect([...missing]).toEqual([]);
  });
});

describe('the guide’s consumables are the sim’s', () => {
  it.each(SPECS)('%s lists exactly what the sim’s level 60 raid preset carries', (spec) => {
    const want = raidConsumeIds(RAID_CONSUMES, spec, 60).filter((id) => !KIT_IMBUES.has(id));
    const shown = consumableRows(simRaidConsumes(spec, {}), {}, () => false).map((row) => row.id);
    expect(shown).toEqual([...new Set(want)]);
  });

  it('reads the current preset, not the held BiS file: Destruction carries every caster elixir', () => {
    const names = simRaidConsumes('warlock-destruction', {}).map((row) => row.label);
    expect(names).toEqual(
      expect.arrayContaining([
        'Flask of Supreme Power',
        'Greater Arcane Elixir',
        'Elixir of Shadow Power',
        'Elixir of Fire Power',
        'Baked Salmon',
      ]),
    );
    expect(names).not.toContain('Runn Tum Tuber Surprise');
  });

  it('gives each hunter what its play uses: bow hunters no Strength items, Survival its frenzy potion', () => {
    for (const spec of ['hunter-beast-mastery', 'hunter-marksmanship']) {
      const ids = simRaidConsumes(spec, {}).map((row) => row.id);
      for (const id of ['juju_power', 'r_o_i_d_s', 'food_grilled_squid']) expect(ids).not.toContain(id);
      // Forever's Juju Might (16329) also gives 40 ranged attack power (aura 124).
      expect(ids).toContain('juju_might');
      expect(ids).toContain('major_mana_potion');
    }
    const survival = simRaidConsumes('hunter-survival', {}).map((row) => row.id);
    expect(survival).toContain('major_frenzy_potion');
    expect(survival).toContain('juju_might');
    expect(survival).not.toContain('major_mana_potion');
  });

  it('takes the level 60 row, not an earlier one', () => {
    const file = {
      specs: {
        x: [
          { from_level: 1, consumes: ['a'] },
          { from_level: 50, consumes: ['b'] },
        ],
      },
    };
    expect(raidConsumeIds(file, 'x', 60)).toEqual(['b']);
    expect(raidConsumeIds(file, 'x', 20)).toEqual(['a']);
    expect(raidConsumeIds(file, 'missing', 60)).toEqual([]);
  });
});

function row(id: string): ConsumableRow {
  return { id, kind: consumableKind(id), name: id, icon: undefined, gives: 'x' };
}

describe('consumableLines', () => {
  it('lists the flask beside the elixirs, never instead of them, then the rest under a heading', () => {
    const lines = consumableLines(
      ['food_baked_salmon', 'elixir_of_shadow_power', 'flask_of_supreme_power', 'greater_arcane_elixir'].map(
        row,
      ),
    );
    expect(lines.map((line) => ('header' in line ? `# ${line.header}` : line.row.id))).toEqual([
      'flask_of_supreme_power',
      'elixir_of_shadow_power',
      'greater_arcane_elixir',
      '# Food, oil and potions',
      'food_baked_salmon',
    ]);
  });

  it('adds no heading when every row is of one group', () => {
    expect(consumableLines([row('food_baked_salmon')]).every((line) => !('header' in line))).toBe(true);
  });

  it.each(SPECS)('%s: no flask is ever shown as replacing an elixir', (spec) => {
    const rows = consumableRows(simRaidConsumes(spec, {}), {}, () => false);
    const lines = consumableLines(rows);
    const text = JSON.stringify(lines);
    expect(text).not.toMatch(/without a flask|replaces both elixirs/i);
    expect(lines.filter((line) => 'row' in line)).toHaveLength(rows.length);
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
