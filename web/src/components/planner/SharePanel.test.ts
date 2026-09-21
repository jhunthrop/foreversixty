// web/src/components/planner/SharePanel.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import fixtureClasses from '../../fixtures/planner/classes.json';
import fixtureCombos from '../../fixtures/planner/combos.json';
import fixtureRaces from '../../fixtures/planner/races.json';
import fixtureTalents from '../../fixtures/planner/talents/warrior.json';
import type { LiveDps } from '../../lib/planner/live-dps.svelte';
import { createPlannerStore } from '../../lib/planner/store.svelte';
import type { ClassRow, Combo, RaceRow, TalentFile } from '../../lib/planner/types';
import SharePanel from './SharePanel.svelte';

// Same loading sequence store.test.ts's own `loaded()` uses.
function fixtureStore() {
  const store = createPlannerStore({ treeVersion: '1.15.9.69722', classSlug: 'warrior', raceSlug: 'human' });
  store.setReference({
    classes: fixtureClasses as ClassRow[],
    races: fixtureRaces as RaceRow[],
    combos: fixtureCombos as Combo[],
  });
  store.setTalents(fixtureTalents as TalentFile);
  store.addPoint(1001);
  return store;
}

// SharePanel only reads `live.state`; a minimal stub is enough for a server render.
const fixtureLive = { state: 'off' } as unknown as LiveDps;

describe('SharePanel', () => {
  it('shows the Share button, not a saved link, on first render', () => {
    const { body } = render(SharePanel, { props: { store: fixtureStore(), live: fixtureLive } });
    expect(body).not.toContain('data-testid="share-link"');
    expect(body).toContain('>Share<');
  });

  it('does not show the confirm step before Share is clicked', () => {
    const { body } = render(SharePanel, { props: { store: fixtureStore(), live: fixtureLive } });
    expect(body).not.toContain('data-testid="share-confirm"');
  });

  it('gives the Share button a stable test id for the e2e confirm flow', () => {
    const { body } = render(SharePanel, { props: { store: fixtureStore(), live: fixtureLive } });
    expect(body).toContain('data-testid="share-open"');
  });
});
