// web/src/fixtures/planner/fixture.test.ts
import { readdirSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import classes from './classes.json';
import combos from './combos.json';
import items from './items/warrior.json';
import races from './races.json';
import sets from './sets.json';
import talents from './talents/warrior.json';
import { MAX_POINTS, SLOTS } from '../../lib/planner/types';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const iconFiles = new Set(readdirSync(path.join(HERE, 'icons')));
const ITEM_FILE_SLOTS = new Set<string>([...SLOTS, 'finger', 'trinket']);

describe('planner fixture talents', () => {
  it('is the warrior file for the active build with two ordered trees', () => {
    expect(talents.build).toBe('1.15.9.69722');
    expect(talents.class_slug).toBe('warrior');
    expect(talents.class_id).toBe(1);
    expect(talents.trees.map((t) => t.position)).toEqual([0, 1]);
    expect(talents.trees.map((t) => t.name)).toEqual(['Arms', 'Fury']);
  });

  it('gives every talent exactly max_rank ranks and an icon that exists', () => {
    for (const tree of talents.trees) {
      for (const talent of tree.talents) {
        expect(talent.ranks).toHaveLength(talent.max_rank);
        expect(iconFiles.has(`${talent.icon}.webp`)).toBe(true);
        for (const rank of talent.ranks) {
          expect(rank.description.length).toBeGreaterThan(0);
          expect(rank.spell_id).toBeGreaterThan(0);
        }
      }
    }
  });

  it('uses 0-based tiers starting at 0 and 0-based columns', () => {
    for (const tree of talents.trees) {
      const tiers = [...new Set(tree.talents.map((t) => t.tier))].sort((a, b) => a - b);
      expect(tiers[0]).toBe(0);
      expect(tiers).toEqual(tiers.map((_, i) => i));
      expect(tree.talents.every((t) => t.column >= 0)).toBe(true);
    }
  });

  it('points every prerequisite at an earlier tier in the same tree', () => {
    for (const tree of talents.trees) {
      const byId = new Map(tree.talents.map((t) => [t.id, t]));
      for (const talent of tree.talents) {
        if (talent.prereq_talent_id === null) {
          expect(talent.prereq_rank).toBeNull();
          continue;
        }
        const prereq = byId.get(talent.prereq_talent_id);
        expect(prereq).toBeDefined();
        expect(prereq!.tier).toBeLessThan(talent.tier);
        expect(talent.prereq_rank).toBeGreaterThan(0);
        expect(talent.prereq_rank!).toBeLessThanOrEqual(prereq!.max_rank);
      }
    }
  });

  it('has more rank capacity than the 51-point cap so the cap is reachable', () => {
    const capacity = talents.trees.flatMap((t) => t.talents).reduce((n, t) => n + t.max_rank, 0);
    expect(capacity).toBeGreaterThan(MAX_POINTS);
  });
});

describe('planner fixture reference data', () => {
  it('lists all nine classes and nine races, each change carrying a source', () => {
    expect(classes).toHaveLength(9);
    expect(races).toHaveLength(9);
    for (const row of [...classes, ...races]) {
      expect(row.forever_changes.length).toBeGreaterThan(0);
      for (const change of row.forever_changes) {
        expect(change.sources.length).toBeGreaterThan(0);
        expect(change.sources.every((s) => s.url.startsWith('https://'))).toBe(true);
      }
    }
    expect(races.find((r) => r.slug === 'skyborne')?.placeholder).toBe(true);
  });

  it('only pairs known race ids with known class ids and marks the Forever additions', () => {
    const classIds = new Set(classes.map((c) => c.id));
    const raceIds = new Set(races.map((r) => r.id));
    for (const combo of combos) {
      expect(classIds.has(combo.class_id)).toBe(true);
      expect(raceIds.has(combo.race_id)).toBe(true);
    }
    const undead = races.find((r) => r.slug === 'undead')!.id;
    const paladin = classes.find((c) => c.slug === 'paladin')!.id;
    expect(combos.find((c) => c.race_id === undead && c.class_id === paladin)?.new_in_forever).toBe(true);
  });
});

describe('planner fixture items', () => {
  it('uses item-file slot names and resolvable set ids and icons', () => {
    const setIds = new Set(sets.map((s) => s.id));
    for (const item of items.items) {
      expect(ITEM_FILE_SLOTS.has(item.slot)).toBe(true);
      expect(iconFiles.has(`${item.icon}.webp`)).toBe(true);
      if (item.set_id !== null) expect(setIds.has(item.set_id)).toBe(true);
    }
    for (const set of sets) {
      expect(set.bonuses.length).toBeGreaterThan(0);
      expect(set.item_ids.length).toBeGreaterThan(0);
    }
  });
});
