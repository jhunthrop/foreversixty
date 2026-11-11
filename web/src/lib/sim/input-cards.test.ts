// web/src/lib/sim/input-cards.test.ts
// The saved run's lower cards against the real-engine fixture and the older shape: every figure a
// card prints is pinned to the fixture's own numbers, read from the fixture rather than typed.
import { describe, expect, it } from 'vitest';
import fullFixture from '../../fixtures/sim/result-full.json';
import olderFixture from '../../fixtures/sim/result.json';
import furyApl from '../../../../data/curated/apl/warrior-fury.json';
import { indexTalents } from '../planner/rules';
import type { TalentFile } from '../planner/types';
import { rotationLinesOf } from '../../../scripts/rotation-apl.mjs';
import type { ActionNames } from './action-names';
import { auraCard } from './aura-card';
import { damageSchools } from './buff-effect';
import { enchantNames, gearSlots, setLines, statsBlock } from './gear-card';
import { resourceCard } from './resource-card';
import { rotationRows } from './rotation-card';
import type { SimBuffRow } from './sim-buffs';
import { talentsFit } from './talents-fit';
import type { EnchantRow } from './enchants';
import type { SimResult } from './types';

const full = fullFixture as unknown as SimResult;
const older = olderFixture as unknown as SimResult;

const names: ActionNames = {
  spell: {
    '2687': 'Bloodrage',
    '11597': 'Sunder Armor',
    '11722': 'Curse of the Elements',
    '9907': 'Faerie Fire',
    '12319': 'Flurry',
    '11551': 'Battle Shout',
    '23894': 'Bloodthirst',
    '20662': 'Execute',
    '1680': 'Whirlwind',
  },
  item: { '13442': 'Mighty Rage Potion' },
};

describe('resourceCard', () => {
  const card = resourceCard(full, names)!;
  const source = full.resource_sources![0]!;
  const sum = (rows: readonly { amount: number }[]): number =>
    rows.reduce((total, row) => total + row.amount, 0);

  it('titles by power and fills the three tiles from the summary', () => {
    const summary = full.summary.resources[0]!;
    expect(card.title).toBe('Rage');
    expect(card.tiles.map((tile) => tile.label)).toEqual(['Gained', 'Spent', 'Lost at the cap']);
    expect(card.tiles[0]!.value).toBe(Math.round(summary.gained).toLocaleString('en-US'));
    const share = ((summary.wasted! / summary.gained) * 100).toFixed(1);
    expect(card.tiles[2]!.value).toBe(`${summary.wasted} · ${share}%`);
  });

  it('lists that sum to the tiles within 0.5', () => {
    const summary = full.summary.resources[0]!;
    expect(Math.abs(sum(source.gained) - summary.gained)).toBeLessThanOrEqual(0.5);
    expect(Math.abs(sum(source.spent) - summary.spent)).toBeLessThanOrEqual(0.5);
  });

  it('sorts by amount, scales bars to the top row and writes the overflow', () => {
    const rows = card.lists!.gained;
    expect(rows[0]!.action).toBe('other:attack/1');
    expect(rows[0]!.name).toBe('Main hand');
    expect(rows[0]!.barPct).toBe(100);
    expect(rows[0]!.lost).toBe('−17.2 lost');
    expect(rows.find((row) => row.action === 'other:refund')).toMatchObject({
      name: 'Refunded on a miss, dodge or parry',
      lost: null,
    });
    expect(rows.find((row) => row.action === 'spell:2687')?.name).toBe('Bloodrage');
    expect(card.lists!.spent.some((row) => row.action === 'spell:2687')).toBe(false);
  });

  it('keeps the tiles and drops the lists on an older run', () => {
    const card = resourceCard(older, names);
    expect(card?.tiles.length).toBeGreaterThanOrEqual(2);
    expect(card?.lists).toBeNull();
  });
});

const entry = (name: string, extra: Partial<SimBuffRow> = {}): SimBuffRow => ({
  name,
  icon: 'x',
  target: false,
  ...extra,
});
const buffs: Record<string, SimBuffRow> = {
  faerie_fire: entry('Faerie Fire', { target: true, stats: { target_armor: -505 } }),
  curse_of_elements: entry('Curse of the Elements', { target: true, school_damage: ['fire', 'frost'] }),
  sunder_armor: entry('Sunder Armor', {
    target: true,
    stats: { target_armor: -2250 },
    exclusive_with: { group: 'major_armor', rank: 1 },
  }),
  arcane_brilliance: entry('Arcane Brilliance', { stats: { intellect: 31 } }),
  battle_shout: entry('Battle Shout', { stats: { attack_power: 112 } }),
};
const fury = { attack_power: 1, strength: 1.9, agility: 0.7, hit: 32, crit: 18 };

describe('auraCard', () => {
  const build = (result: SimResult, weights: typeof fury | null = fury) =>
    auraCard({
      result,
      buffs,
      weights,
      schools: damageSchools(result.summary.damage_done),
      names,
      fallbackTargetIds: new Set(['faerie_fire']),
    });
  const namesIn = (rows: readonly { name: string }[]): string[] => rows.map((row) => row.name);

  it('files uptimes from the fixture: Flurry a bar, Battle Shout up all fight, sorted by uptime', () => {
    const { player } = build(full);
    const flurry = full.summary.auras.find((row) => row.name === 'spell:12319')!;
    const bar = player.bars.find((row) => row.name === 'Flurry')!;
    expect(bar.uptimePct).toBeCloseTo((flurry.uptime_ms / full.summary.duration_ms) * 100, 5);
    expect(bar.applications).toBe(flurry.applications);
    expect(namesIn(player.upAllFight)).toContain('Battle Shout');
    const uptimes = player.bars.map((row) => row.uptimePct);
    expect(uptimes).toEqual([...uptimes].sort((a, b) => b - a));
  });

  it('gives the same answer as Settings: R2 Faerie Fire shows, Curse of the Elements is quiet', () => {
    const { target, player } = build(full);
    expect(namesIn(target.upAllFight)).toContain('Faerie Fire');
    expect(namesIn(target.noEffect)).toContain('Curse of the Elements');
    expect(namesIn(player.noEffect)).toContain('Arcane Brilliance');
  });

  it('shows Sunder Armor with its stacks', () => {
    expect(build(full).target.upAllFight.find((row) => row.name === 'Sunder Armor')?.stacks).toBe(5);
  });

  it('never shows the queue track, the Overpower window, Thrash, zero uptime or rage bookkeeping', () => {
    const track = (name: string, uptime = 1000) => ({
      ...full.summary.auras[0]!,
      name,
      uptime_ms: uptime,
      target_guid: `g-${name}`,
    });
    const withNoise: SimResult = {
      ...full,
      summary: {
        ...full.summary,
        auras: [
          ...full.summary.auras,
          track('spell:21919'),
          track('other:rage_gain'),
          track('spell:12292', 0),
          track('spell:20007/1'),
          track('spell:20007/2'),
        ],
      },
    };
    const { player } = build(withNoise);
    const keys = [...player.bars, ...player.upAllFight, ...player.noEffect].map((row) => row.actionKey);
    for (const hidden of ['spell:21919', 'other:rage_gain', 'spell:12292', 'spell:11567/1', 'spell:11585'])
      expect(keys).not.toContain(hidden);
    const hands = player.bars
      .filter((row) => row.actionKey?.startsWith('spell:20007'))
      .map((row) => row.name);
    expect(hands.some((name) => name.endsWith('(main hand)'))).toBe(true);
    expect(hands.some((name) => name.endsWith('(off hand)'))).toBe(true);
  });

  it('holds off a debuff a lower rank in the same request keeps out', () => {
    const both: SimResult = {
      ...full,
      request: {
        ...full.request,
        character: { ...full.request.character, buffs: ['sunder_armor', 'expose_armor'] },
      },
    };
    const card = auraCard({
      result: both,
      buffs: {
        ...buffs,
        expose_armor: entry('Expose Armor', {
          target: true,
          exclusive_with: { group: 'major_armor', rank: 2 },
        }),
      },
      weights: fury,
      schools: ['physical'],
      names,
      fallbackTargetIds: new Set(),
    });
    expect(namesIn(card.target.upAllFight)).not.toContain('Expose Armor');
  });

  it('files nothing apart while the weights are unread', () => {
    expect(build(full, null).player.noEffect).toEqual([]);
  });

  it('an older run with no target tracks draws the request list and says so', () => {
    const card = build(older);
    expect(card.targetFromRequest).toBe(true);
    expect(card.target.bars).toEqual([]);
  });
});

describe('gear card', () => {
  const rows = [
    { id: 1900, display: 'Crusader', name: 'Enchant Weapon - Crusader' },
    { id: 1, name: 'Enchant Gloves - Greater Strength' },
    { id: 856, display: '+7 Strength', name: 'Enchant Bracer - Superior Strength' },
  ] as unknown as EnchantRow[];
  const enchanted = full.request.character.gear.map((slot) =>
    slot.slot === 'main_hand'
      ? { ...slot, enchant: 1900 }
      : slot.slot === 'wrist'
        ? { ...slot, enchant: 856 }
        : slot.slot === 'hands'
          ? { ...slot, enchant: 1 }
          : slot,
  );
  const slotsOf = (gear: typeof enchanted) =>
    gearSlots({ gear, items: new Map(), failed: true, enchants: enchantNames(rows) });

  it('prints the client enchant text and never the recipe', () => {
    const slots = slotsOf(enchanted);
    expect(slots.find((slot) => slot.slot === 'main_hand')?.enchant).toBe('Crusader');
    expect(slots.find((slot) => slot.slot === 'wrist')?.enchant).toBe('+7 Strength');
    expect(slots.find((slot) => slot.slot === 'hands')?.enchant).toBeUndefined();
    expect(slots.find((slot) => slot.slot === 'head')?.enchant).toBeUndefined();
  });

  it('a set lights the bonuses its pieces reach and a lone piece names no set', () => {
    const set = {
      id: 1,
      name: 'Battlegear',
      item_ids: [12640, 13089, 99, 98, 97],
      bonuses: [
        { pieces: 2, description: 'A' },
        { pieces: 5, description: 'B' },
      ],
    };
    expect(setLines(full.request.character.gear, [set])).toEqual([
      {
        id: 1,
        head: 'Set: Battlegear 2 of 5',
        bonuses: [
          { text: '(2) A', active: true },
          { text: '(5) B', active: false },
        ],
      },
    ]);
    expect(setLines(full.request.character.gear, [{ ...set, item_ids: [12640, 99] }])).toEqual([]);
  });

  it('reads every stat from final_stats: 0 haste reads 0%', () => {
    const block = statsBlock(full)!;
    const stats = full.final_stats!;
    expect(block.head).toBe('In the fight, against a level 63 boss.');
    expect(block.hit?.value).toBe(`${stats.melee_hit_pct!.toFixed(1)}%`);
    const value = (id: string) => block.cells.find((cell) => cell.id === id)?.value;
    expect(value('attack-power')).toBe(stats.attack_power!.toLocaleString('en-US'));
    expect(value('crit')).toBe(`${stats.melee_crit_pct_vs_target!.toFixed(1)}%`);
    expect(value('skill')).toBe(`${stats.weapon_skill!.main_hand} · ${stats.weapon_skill!.off_hand}`);
    expect(value('haste')).toBe('0%');
  });

  it('puts the run’s own measured miss rate under hit, never a fixed cap', () => {
    // Specials: 0 MISS in the fixture's yellow rows. White: 12 + 7 + 4 misses over 63 + 82 + 21 swings.
    expect(statsBlock(full)?.hit?.reference?.note).toBe('Miss 0% specials · 13.9% white');
  });

  it('one weapon keeps the hit line; the skill shows one figure', () => {
    const oneHand: SimResult = {
      ...full,
      request: {
        ...full.request,
        character: {
          ...full.request.character,
          gear: full.request.character.gear.filter((slot) => slot.slot !== 'off_hand'),
        },
      },
    };
    expect(statsBlock(oneHand)?.cells.find((cell) => cell.id === 'skill')?.value).toBe(
      String(full.final_stats!.weapon_skill!.main_hand),
    );
  });

  it('hides the block without final_stats, never dashes', () => {
    expect(statsBlock(older)).toBeNull();
  });
});

function talentFile(): TalentFile {
  const talent = (id: number, tier: number, column: number, max: number, prereq: number | null = null) => ({
    id,
    name: `T${id}`,
    icon: 'i',
    max_rank: max,
    tier,
    column,
    prereq_talent_id: prereq,
    prereq_rank: prereq === null ? null : 1,
    ranks: [],
  });
  return {
    build: 'b',
    class_id: 1,
    class_slug: 'warrior',
    trees: [
      {
        id: 1,
        name: 'A',
        position: 0,
        talents: [talent(1, 0, 0, 5), talent(2, 1, 0, 1, 1), talent(3, 1, 1, 3)],
      },
      { id: 2, name: 'B', position: 1, talents: [talent(4, 0, 0, 3)] },
    ],
  } as unknown as TalentFile;
}

describe('talentsFit', () => {
  const index = indexTalents(talentFile());

  it('takes a legal string, and one the engine trimmed', () => {
    expect(talentsFit('510-3', index)).toBe(true);
    expect(talentsFit('5-', index)).toBe(true);
  });
  it.each([
    ['too few parts', '510'],
    ['a part longer than its tree', '5100-3'],
    ['a digit past the top rank', '610-3'],
    ['a tier gate not met', '101-3'],
    ['a prerequisite not met', '410-3'.replace('410', '011')],
    ['not digits', '5x0-3'],
  ])('refuses %s', (_why, talents) => {
    expect(talentsFit(talents, index)).toBe(false);
  });
});

describe('rotationRows', () => {
  const lines = {
    sha: 'x',
    prepull: [
      { action: 'spell:2687', at_seconds: 2.5 },
      { action: 'spell:11551', at_seconds: 1.5 },
    ],
    priority: [
      { action: null, note: '' },
      { action: 'spell:20662', note: 'First.' },
      { action: 'spell:23894', note: 'Second.' },
      { action: 'spell:23894', note: 'Again.' },
      { action: 'spell:11585', note: 'Never.' },
    ],
  };
  const rows = rotationRows(full, lines, { ...names, spell: { ...names.spell, '11585': 'Overpower' } });
  const casts = (key: string) => full.summary.casts.find((row) => row.spell_name === key)!.succeeded;

  it('counts casts per fight from the summary, once per spell', () => {
    expect(rows.priority[1]).toMatchObject({
      name: 'Execute',
      casts: casts('spell:20662'),
      dimmed: false,
      note: 'First.',
    });
    expect(rows.priority[2]!.casts).toBe(casts('spell:23894'));
    expect(rows.priority[3]!.casts).toBeNull();
  });
  it('dims a line never cast', () => {
    expect(rows.priority[4]).toMatchObject({ dimmed: true, casts: 0 });
  });
  it('the cooldown pass lists only what no other line names, swings left out', () => {
    expect(rows.priority[0]!.note).toContain('Mighty Rage Potion');
    expect(rows.priority[0]!.note).not.toContain('Bloodthirst');
    expect(rows.priority[0]!.note).not.toContain('Main hand');
  });
  it('a pre-pull row is Cast only when the sample has it before the pull', () => {
    expect(rows.prepull[0]).toMatchObject({ name: 'Bloodrage', at: '-2.5 s', cast: true });
    expect(rows.prepull[1]).toMatchObject({ name: 'Battle Shout', cast: false });
  });
  it('an older run with no sample says nothing about the pull', () => {
    const { sample: _sample, ...noSample } = full;
    expect(rotationRows(noSample as SimResult, lines, names).prepull[0]!.cast).toBeNull();
  });
  it('reads a curated APL: pass, sequence and leak guard', () => {
    const parsed = rotationLinesOf({
      rotation: {
        prepullActions: [
          { action: { castSpell: { spellId: { spellId: 1 } } }, doAtValue: { const: { val: '-1.5s' } } },
          { action: { castSpell: { spellId: { spellId: 2 } } }, doAtValue: { const: { val: '-3s' } } },
        ],
        priorityList: [
          { action: { autocastOtherCooldowns: {} } },
          {
            notes: 'Hit it.',
            action: { strictSequence: { actions: [{ castSpell: { spellId: { spellId: 3 } } }] } },
          },
          { action: { castSpell: { spellId: { itemId: 4 } } } },
          { action: { wait: {} } },
        ],
      },
    } as Parameters<typeof rotationLinesOf>[0]);
    expect(parsed.prepull.map((row) => row.action)).toEqual(['spell:2', 'spell:1']);
    expect(parsed.priority).toEqual([
      { action: null, note: '' },
      { action: 'spell:3', note: 'Hit it.' },
      { action: 'item:4', note: '' },
    ]);
  });

  const leveled = {
    rotation: {
      priorityList: [
        { action: { autocastOtherCooldowns: {} } },
        { notes: 'From 40 it goes first.', action: { castSpell: { spellId: { spellId: 20662 } } } },
      ],
    },
    level_notes: [
      { from_level: 60, step: 1, notes: 'First.' },
      { from_level: 50, step: 1, notes: 'Fifty.' },
    ],
  } as Parameters<typeof rotationLinesOf>[0];

  it('carries level notes per step, lowest level first, and rejects a bad one', () => {
    expect(rotationLinesOf(leveled).priority[1]).toEqual({
      action: 'spell:20662',
      note: 'From 40 it goes first.',
      level_notes: [
        { from_level: 50, note: 'Fifty.' },
        { from_level: 60, note: 'First.' },
      ],
    });
    const onPass = { ...leveled, level_notes: [{ from_level: 60, step: 0, notes: 'x' }] };
    expect(() => rotationLinesOf(onPass as typeof leveled)).toThrow(/casts nothing/);
    const outOfRange = { ...leveled, level_notes: [{ from_level: 60, step: 9, notes: 'x' }] };
    expect(() => rotationLinesOf(outOfRange as typeof leveled)).toThrow(/level_notes\[0\]/);
  });

  it('prints the note for the run’s level, and the general note below every level note', () => {
    const parsed = { ...rotationLinesOf(leveled), sha: null };
    const at = (level: number) => {
      const run: SimResult = {
        ...full,
        request: { ...full.request, character: { ...full.request.character, level } },
      };
      return rotationRows(run, parsed, names).priority[1]!.note;
    };
    expect(at(60)).toBe('First.');
    expect(at(55)).toBe('Fifty.');
    expect(at(40)).toBe('From 40 it goes first.');
  });

  it('the curated Fury notes at 60 carry no leveling clause and one Execute wording', () => {
    const fury = rotationLinesOf(furyApl, 'warrior-fury');
    const notes = rotationRows(full, { ...fury, sha: null }, names).priority.map((row) => row.note);
    expect(full.request.character.level).toBe(60);
    for (const note of notes) {
      expect(note).not.toMatch(/\b[Ff]rom (36|40)\b|before level|\bat 59\b|below 20%|from 25% down/);
      expect(note).not.toMatch(/the execute phase/);
    }
  });

  it('an older run counts a spell with no cast row from its damage row, as the table does', () => {
    // result.json: Execute dealt 6 hits and kept no cast row; Bloodthirst has neither.
    const fury = { ...rotationLinesOf(furyApl, 'warrior-fury'), sha: null };
    const rows = rotationRows(older, fury, names).priority;
    expect(rows.find((row) => row.name === 'Execute')).toMatchObject({ casts: 6, dimmed: false });
    expect(rows.find((row) => row.name === 'Bloodthirst')).toMatchObject({ casts: 0, dimmed: true });
    // A newer run's second-hand hit (Whirlwind's 462891) is not a second press of Whirlwind.
    const current = rotationRows(full, fury, { ...names, spell: { ...names.spell, '462891': 'Whirlwind' } });
    const whirlwind = full.summary.casts.find((row) => row.spell_name === 'spell:1680')!.succeeded;
    expect(current.priority.find((row) => row.name === 'Whirlwind')?.casts).toBe(Math.round(whirlwind));
  });
});
