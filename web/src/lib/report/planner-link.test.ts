// web/src/lib/report/planner-link.test.ts
import { describe, expect, it } from 'vitest';
import fixtureSummary from '../../fixtures/report/fights/3/summary.json';
import type { Summary } from './types';
import {
  GEAR_SLOT_ORDER,
  classSlugOf,
  gearFromCombatant,
  plannerLinkFor,
  treeRanksFromTalents,
} from './planner-link';

const summary = fixtureSummary as Summary;
const tank = summary.combatants.find((row) => row.name === 'Baelgrim-Nightslayer')!;

describe('gear from a combatant', () => {
  it('maps COMBATANT_INFO’s slot order onto the planner’s slots', () => {
    expect(GEAR_SLOT_ORDER[0]).toBe('head');
    expect(GEAR_SLOT_ORDER[3]).toBeNull(); // shirt
    expect(GEAR_SLOT_ORDER[14]).toBe('back');
    expect(GEAR_SLOT_ORDER[17]).toBeNull(); // tabard
  });

  it('keeps the items that exist and drops the empty and unslotted ones', () => {
    expect(gearFromCombatant(tank.gear)).toEqual({ head: 175850, neck: 175885 });
  });
});

describe('talent ranks', () => {
  it('splits a rank-per-talent array across the trees', () => {
    expect(treeRanksFromTalents([5, 0, 3, 1, 2, 0], [3, 2, 1])).toEqual([[5, 0, 3], [1, 2], [0]]);
  });

  it('refuses an array whose length is not the class’s talent count', () => {
    // The fixture's combatant carries seven retail spell ids, not ranks.
    expect(treeRanksFromTalents(tank.talents, [3, 2, 1])).toBeNull();
    expect(treeRanksFromTalents([1, 2], [3, 2, 1])).toBeNull();
  });
});

describe('class slugs', () => {
  it('slugs the engine’s class names, and refuses one the planner has no data for', () => {
    expect(classSlugOf('Warrior')).toBe('warrior');
    expect(classSlugOf('Death Knight')).toBeNull();
    expect(classSlugOf(undefined)).toBeNull();
  });
});

describe('plannerLinkFor', () => {
  const base = { dataBuild: '1.15.9.69722', className: 'Warrior', treeSizes: [3, 2, 1] };

  it('builds a gear-only link when the log did not record ranks', () => {
    const link = plannerLinkFor({ ...base, combatant: tank })!;
    expect(link.talents).toBe(false);
    expect(link.label).toBe('Gear in the planner');
    expect(link.href).toBe(
      `/planner?code=${encodeURIComponent('FS1:1.15.9.69722:warrior::0/0/0:head=175850,neck=175885')}`,
    );
  });

  it('carries the talents when the log did record ranks', () => {
    const link = plannerLinkFor({
      ...base,
      combatant: { ...tank, talents: [3, 0, 2, 1, 0, 0] },
    })!;
    expect(link.talents).toBe(true);
    expect(link.label).toBe('Build in the planner');
    // Tree 2's ranks are [1, 0]; encodeTree (fs1.ts, settled by Task 13) trims trailing
    // zeros per tree, so the trailing 0 is dropped rather than spelled out.
    expect(decodeURIComponent(link.href)).toContain(':302/1/0:');
  });

  it('has no link at all for a class the planner has no data for', () => {
    expect(plannerLinkFor({ ...base, className: 'Death Knight', combatant: tank })).toBeNull();
  });

  it('leaves the race field empty, because no combat-log event records it', () => {
    const link = plannerLinkFor({ ...base, combatant: tank })!;
    expect(decodeURIComponent(link.href)).toContain(':warrior::');
  });
});
