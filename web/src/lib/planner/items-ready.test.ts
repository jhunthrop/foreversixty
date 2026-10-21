import { describe, expect, it, vi } from 'vitest';
import fixtureClasses from '../../fixtures/planner/classes.json';
import fixtureItems from '../../fixtures/planner/items/warrior.json';
import { ensureItems } from './items-ready';
import { createPlannerStore } from './store.svelte';
import type { ClassRow, ItemFile } from './types';

function store() {
  const s = createPlannerStore({
    treeVersion: '1.15.9.69722',
    classSlug: 'warrior',
    raceSlug: 'human',
    gear: { head: 12640, finger1: 264908 },
  });
  s.setReference({ classes: fixtureClasses as ClassRow[], races: [], combos: [] });
  return s;
}

describe('ensureItems', () => {
  it('loads the class file once, and the draft then drops what it does not list', async () => {
    const s = store();
    const load = vi.fn().mockResolvedValue(fixtureItems as ItemFile);
    await Promise.all([ensureItems(s, load), ensureItems(s, load)]);
    await ensureItems(s, load);
    expect(load).toHaveBeenCalledTimes(1);
    expect(s.leftOutSlots).toEqual(['finger1']);
  });

  it('leaves the gear as worn when the file cannot be read', async () => {
    const s = store();
    await ensureItems(s, vi.fn().mockRejectedValue(new Error('offline')));
    expect(s.leftOutSlots).toEqual([]);
  });
});
