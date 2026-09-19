// web/src/fixtures/sim/bulk-fixture.test.ts
// The fixtures are data, so their invariants are asserted rather than assumed: every later
// task's tests read these files, and a fixture that quietly stops matching the contract
// turns a real failure into a green run.
//
// Two assertions below deliberately diverge from the task brief's draft, because checking
// the real files settled the question the draft got wrong:
//   - phases[0].start: api/internal/phase.go's zero Boundary marshals through Go's
//     encoding/json as "0001-01-01T00:00:00Z", not "" -- and the phases.json already
//     checked in on this branch (part A) already carries that value. "" would have been
//     wrong the moment a real GET /v1/phases response was compared against it.
//   - weights: the reference stat's weight is pinned to exactly 1 by definition, with zero
//     uncertainty -- it is not a ratio of two noisy estimates the way every other row is.
//     Both fixture files agree on this (weights-result.json's attack_power row already
//     carries error: 0); only the brief's own test snippet asked for every row's error to
//     be positive, which its own fixture could never pass.
import { describe, expect, it } from 'vitest';
import bulkJson from './bulk-result.json';
import weightsJson from './weights-result.json';
import specsJson from './specs.json';
import lootJson from '../planner/loot.json';
import enchantsJson from '../planner/enchants.json';
import suffixesJson from '../planner/suffixes.json';
import simbuffsJson from '../planner/simbuffs.json';
import itemsJson from '../planner/items/warrior.json';
import phasesJson from '../../data/phases.json';
import type { BulkResult, WeightsResult } from '../../lib/sim/bulk-types';
import type { SpecFidelity } from '../../lib/sim/types';

const bulk = bulkJson as unknown as BulkResult;
const weights = weightsJson as unknown as WeightsResult;

describe('the bulk result fixture', () => {
  it('is a gear request with an equipped baseline and three stages', () => {
    expect(bulk.request.bulk?.mode).toBe('gear');
    expect(bulk.request.bulk?.precision).toBe('fast');
    expect(bulk.equipped.mean).toBeGreaterThan(0);
    expect(bulk.stages.map((stage) => stage.iterations)).toEqual([100, 1000, 3000]);
  });

  it('is ranked best first, with every delta measured against the equipped set', () => {
    const means = bulk.combos.map((combo) => combo.dps.mean);
    expect([...means].sort((a, b) => b - a)).toEqual(means);
    for (const combo of bulk.combos) {
      expect(combo.delta.mean).toBeCloseTo(combo.dps.mean - bulk.equipped.mean, 6);
      expect(combo.delta.error).toBeGreaterThan(0);
    }
  });

  it('carries a within-error group at the top and a separated group below it', () => {
    expect(bulk.combos.filter((combo) => combo.group === 0).length).toBeGreaterThan(1);
    expect(new Set(bulk.combos.map((combo) => combo.group))).toEqual(new Set([0, 1, 2]));
  });

  it('names a slot on every ring and trinket substitution, the way sim/bulk resolves them', () => {
    for (const combo of bulk.combos) {
      for (const sub of combo.substitutions) {
        if (sub.kind !== 'item') continue;
        expect(sub.slot).toBeTruthy();
        expect(sub.item_id).toBeGreaterThan(0);
        expect(sub.origin).toBeTruthy();
      }
    }
  });
});

describe('the weights result fixture', () => {
  it('normalises the reference stat to exactly 1, with zero error', () => {
    const reference = weights.request.weights?.reference ?? '';
    const row = weights.weights.find((entry) => entry.stat === reference);
    expect(row?.weight).toBe(1);
    expect(row?.error).toBe(0);
  });

  it('gives every non-reference stat a measured, positive error', () => {
    const reference = weights.request.weights?.reference ?? '';
    for (const row of weights.weights) {
      if (row.stat === reference) continue;
      expect(row.error).toBeGreaterThan(0);
    }
  });
});

describe('the spec fixture', () => {
  it('gives every spec a reference stat', () => {
    for (const row of specsJson as unknown as SpecFidelity[]) {
      expect(row.reference_stat).toBeTruthy();
    }
  });
});

describe('the loot fixture', () => {
  const itemIds = new Set((itemsJson.items as { id: number }[]).map((item) => item.id));

  it('covers every source kind the picker groups by', () => {
    expect(new Set(lootJson.sources.map((source) => source.kind))).toEqual(
      new Set(['raid', 'dungeon', 'world', 'crafted', 'rep', 'pvp', 'quest']),
    );
  });

  it('only drops items the fixture item file actually has', () => {
    for (const source of lootJson.sources) {
      for (const id of [...(source.items ?? []), ...(source.trash ?? [])]) {
        expect(itemIds.has(id), `source ${source.id} drops ${id}`).toBe(true);
      }
      for (const boss of source.bosses ?? []) {
        for (const id of boss.items) expect(itemIds.has(id), `boss ${boss.id} drops ${id}`).toBe(true);
      }
    }
  });

  it('names a phase only from the phase table, or the literal "later"', () => {
    const names = new Set([...phasesJson.map((phase) => phase.name), 'later']);
    for (const source of lootJson.sources) {
      if (source.opens === undefined) continue;
      expect(names.has(source.opens), `source ${source.id} opens in ${source.opens}`).toBe(true);
    }
  });

  it('spells every boss id as <source>:<npc-id>, per contract 10.4', () => {
    for (const source of lootJson.sources) {
      for (const boss of source.bosses ?? []) {
        expect(boss.id).toBe(`${source.id}:${boss.npc_id}`);
      }
    }
  });
});

describe('the enchant and suffix fixtures', () => {
  // Real shape (data/builds/1.60.1.69893/enchants.json), not the brief's draft: the fork
  // database's UIEnchant rows key on "id" (the effect id), every row carries a nonzero
  // spell_id, item_id is 0 unless a physical item (a kit or a scroll) applies the effect,
  // and item_types is the EnchantType string enum ("normal", "kit", "shield", "two_hand"),
  // not a numeric list.
  it('keys every enchant by id (the effect id) with a spell that applies it, per the fork database UIEnchant shape', () => {
    for (const enchant of enchantsJson) {
      expect(enchant.id).toBeGreaterThan(0);
      expect(enchant.spell_id).toBeGreaterThan(0);
      expect(enchant.item_id).toBeGreaterThanOrEqual(0);
      expect(enchant.slots.length).toBeGreaterThan(0);
      expect(enchant.item_types.length).toBeGreaterThan(0);
      expect(enchant.icon.startsWith('fixture_')).toBe(true);
    }
  });

  it('points every suffix on an item at a suffix the file defines', () => {
    const suffixIds = new Set(suffixesJson.map((suffix) => suffix.id));
    for (const item of itemsJson.items as { suffixes?: number[] }[]) {
      for (const id of item.suffixes ?? []) expect(suffixIds.has(id)).toBe(true);
    }
  });
});

describe('the phase table', () => {
  it('is the four phases api/internal/phase names, in order, with the launch instant', () => {
    expect(phasesJson.map((phase) => phase.name)).toEqual(['pre-beta', 'beta', 'launch', 'raids-1']);
    expect(phasesJson[2].start).toBe('2026-11-04T23:00:00Z');
    expect(phasesJson[3].start).toBe('2026-12-09T00:00:00Z');
    // api/internal/phase.go: pre-beta's zero time.Time marshals as this exact RFC3339
    // instant, which is what a real GET /v1/phases response sends -- not an empty string.
    expect(phasesJson[0].start).toBe('0001-01-01T00:00:00Z');
  });

  it('carries no display label: contract 10.4 pins the shape at name and start', () => {
    for (const phase of phasesJson) expect(Object.keys(phase).sort()).toEqual(['name', 'start']);
  });
});

describe('the simbuffs fixture', () => {
  it('names and illustrates every consumable the preset applies', () => {
    const entries: Record<string, { name: string; icon: string }> = simbuffsJson.entries;
    for (const id of ['flask_of_supreme_power', 'elixir_of_the_mongoose']) {
      expect(entries[id]?.name.length).toBeGreaterThan(0);
      expect(entries[id]?.icon.startsWith('fixture_')).toBe(true);
    }
  });
});
