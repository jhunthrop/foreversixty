import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import { createPlannerStore } from '../../lib/planner/store.svelte';
import { indexTalents } from '../../lib/planner/rules';
import type { TalentFile } from '../../lib/planner/types';
import TalentCell from './TalentCell.svelte';

const talents = indexTalents(fixtureTalents as TalentFile);
const arms = talents.trees[0];
const talent = arms.talents[0];

function storeWithRank(rank: number) {
  const store = createPlannerStore({ treeVersion: 'test', classSlug: 'warrior', raceSlug: 'human' });
  store.setTalents(fixtureTalents as TalentFile);
  for (let i = 0; i < rank; i++) store.addPoint(talent.id);
  return store;
}

describe('TalentCell read-only mode', () => {
  it('renders a div, not a button, when readOnly', () => {
    const store = storeWithRank(1);
    const { body } = render(TalentCell, {
      props: { store, talent, focused: false, onfocuscell: () => {}, readOnly: true },
    });
    expect(body).not.toContain('<button');
    expect(body).toContain(`data-testid="talent-${talent.id}"`);
  });

  it('still shows the real rank and max rank when readOnly', () => {
    const store = storeWithRank(2);
    const { body } = render(TalentCell, {
      props: { store, talent, focused: false, onfocuscell: () => {}, readOnly: true },
    });
    expect(body).toContain(`${2}/${talent.max_rank}`);
  });

  it('renders the normal interactive button when readOnly is left out', () => {
    const store = storeWithRank(0);
    const { body } = render(TalentCell, { props: { store, talent, focused: false, onfocuscell: () => {} } });
    expect(body).toContain('<button');
  });
});
