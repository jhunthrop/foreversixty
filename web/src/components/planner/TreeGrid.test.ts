import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import { createPlannerStore } from '../../lib/planner/store.svelte';
import type { TalentFile } from '../../lib/planner/types';
import TreeGrid from './TreeGrid.svelte';

const talentFile = fixtureTalents as TalentFile;
const tree = talentFile.trees[0];

function readyStore() {
  const store = createPlannerStore({ treeVersion: 'test', classSlug: 'warrior', raceSlug: 'human' });
  store.setTalents(talentFile);
  return store;
}

describe('TreeGrid read-only mode', () => {
  it('renders no <button> elements when readOnly', () => {
    const { body } = render(TreeGrid, { props: { store: readyStore(), tree, readOnly: true } });
    expect(body).not.toContain('<button');
  });

  it('still renders one cell per talent', () => {
    const { body } = render(TreeGrid, { props: { store: readyStore(), tree, readOnly: true } });
    for (const talent of tree.talents) {
      expect(body).toContain(`data-testid="talent-${talent.id}"`);
    }
  });
});
