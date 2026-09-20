// web/src/components/sim/tools/StatWeights.test.ts
// D45/D46's own component-level check on the *live* tool page -- fix round 1 found coverage
// resting only on SavedWeights.svelte's static render plus the shared predicate, but D46 was
// specifically about the live page (the table above the COPY FOR PAWN button). A weight the
// engine flagged `insignificant` renders greyed with the label here, and the Pawn string
// below the table never contains it -- both off `createBulkStore`'s real `weights`/`pawn`
// derivations, not a hand-mocked store, so this exercises the actual component wiring
// (`StatWeights.svelte`'s `$derived(pawnString(...))`, its `{@const significant =
// isSignificant(row)}`), not a re-statement of weights.ts's own unit tests.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import weightsFixture from '../../../fixtures/sim/weights-result.json';
import { createBulkStore } from '../../../lib/sim/bulk-store.svelte';
import type { WeightsResult } from '../../../lib/sim/bulk-types';
import { WEIGHT_INSIGNIFICANT_LABEL } from '../../../lib/sim/copy';
import type { StatWeight } from '../../../lib/sim/types';
import StatWeights from './StatWeights.svelte';

const fixture = weightsFixture as unknown as WeightsResult;

/**
 * A store with `weights` set through the real `adoptResult` (the same method a finished run
 * calls), not a hand-built fake object -- so this renders through `createBulkStore`'s own
 * `get weights()`/`get referenceStat()`/`get stats()` getters, the same ones the live page
 * reads. No character is loaded: none of the assertions below need one, and `BulkRunBar`
 * (which `StatWeights` always mounts) renders fine with `store.character === null`, the
 * same as it does for a moment on every real page load before a character finishes loading.
 */
function storeWith(weights: StatWeight[]) {
  const store = createBulkStore({ tool: 'weights', treeVersion: 'test' });
  store.setStats(weights.map((row) => row.stat));
  store.adoptResult({ ...fixture, weights });
  return store;
}

function renderWeights(weights: StatWeight[]): string {
  const { body } = render(StatWeights, { props: { store: storeWith(weights), me: null } });
  return body;
}

/** The whole `<li data-testid="sim-weight-<stat>">…</li>` element for one row -- a class
 *  that renders before the test id (it does: `class` comes before `data-testid`) is not
 *  lost the way slicing the body on the test id string alone would lose it. */
function rowFor(body: string, stat: string): string {
  const match = new RegExp(`<li[^>]*data-testid="sim-weight-${stat}"[^>]*>.*?</li>`, 's').exec(body);
  if (match === null) throw new Error(`no <li> rendered for stat "${stat}"`);
  return match[0];
}

describe('StatWeights: the greyed row and the Pawn string agree (D45/D46, live page)', () => {
  it('greys exactly the row the engine flagged insignificant, and no other', () => {
    const weights: StatWeight[] = [
      { stat: 'attack_power', weight: 1, error: 0 },
      { stat: 'strength', weight: 2.14, error: 0.06 },
      { stat: 'crit', weight: 14.37, error: 22.55, insignificant: true },
    ];
    const body = renderWeights(weights);
    const critRow = rowFor(body, 'crit');
    const strengthRow = rowFor(body, 'strength');
    expect(critRow).toContain('opacity-50');
    expect(critRow).toContain(WEIGHT_INSIGNIFICANT_LABEL);
    expect(strengthRow).not.toContain('opacity-50');
    expect(strengthRow).not.toContain(WEIGHT_INSIGNIFICANT_LABEL);
  });

  it('D46: the Pawn string omits exactly the greyed row, never disagreeing with the table', () => {
    const weights: StatWeight[] = [
      { stat: 'attack_power', weight: 1, error: 0 },
      { stat: 'strength', weight: 2.14, error: 0.06 },
      { stat: 'crit', weight: 14.37, error: 22.55, insignificant: true },
    ];
    const body = renderWeights(weights);
    expect(body).not.toContain('CritRating=');
    expect(body).toContain('AttackPower=1.00');
    expect(body).toContain('Strength=2.14');
  });

  it('a result with no insignificant rows greys nothing', () => {
    const weights: StatWeight[] = [
      { stat: 'attack_power', weight: 1, error: 0 },
      { stat: 'strength', weight: 2.14, error: 0.06 },
    ];
    const body = renderWeights(weights);
    // Scoped to the two row <li>s, not the whole body: unrelated elements elsewhere on the
    // page (the save button's own `disabled:opacity-50` utility class) legitimately carry
    // the substring "opacity-50" without meaning a row is greyed.
    expect(rowFor(body, 'attack_power')).not.toContain('opacity-50');
    expect(rowFor(body, 'strength')).not.toContain('opacity-50');
    expect(body).not.toContain(WEIGHT_INSIGNIFICANT_LABEL);
    expect(body).toContain('Strength=2.14');
  });
});
