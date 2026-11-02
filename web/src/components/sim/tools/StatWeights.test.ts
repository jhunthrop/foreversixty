// web/src/components/sim/tools/StatWeights.test.ts
// The live tool page's own component-level check, off `createBulkStore`'s real `weights`
// derivation: a weight the engine flagged `insignificant` renders greyed, and the table reads
// as the published weights do (owner, 2026-09-30; round 11) -- top stat 1.00, no Pawn string.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import weightsFixture from '../../../fixtures/sim/weights-result.json';
import { createBulkStore } from '../../../lib/sim/bulk-store.svelte';
import type { WeightsResult } from '../../../lib/sim/bulk-types';
import { bisCopy } from '../../../lib/bis/copy';
import { bulkCopy } from '../../../lib/sim/copy';
import type { StatWeight } from '../../../lib/sim/types';
import { weightsEngineIterations, WEIGHTS_BROWSER_DEFAULT_ITERATIONS } from '../../../lib/sim/weights';
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

/** The whole `<div data-testid="sim-weight-<stat>">…</div>` row element for one stat. */
function rowFor(body: string, stat: string): string {
  const match = new RegExp(`<div[^>]*data-testid="sim-weight-${stat}"[^>]*>.*?</div>`, 's').exec(body);
  if (match === null) throw new Error(`no row rendered for stat "${stat}"`);
  return match[0];
}

describe('StatWeights: the published SimulationCraft form (live page)', () => {
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
    expect(critRow).toContain(bisCopy.weightsNotSignificant);
    expect(strengthRow).not.toContain('opacity-50');
    expect(strengthRow).not.toContain(bisCopy.weightsNotSignificant);
  });

  it('normalizes to the top stat, shows DPS per point, and carries no Pawn string', () => {
    const weights: StatWeight[] = [
      { stat: 'attack_power', weight: 1, error: 0.01, dps_per_point: 0.06 },
      { stat: 'strength', weight: 2, error: 0.06, dps_per_point: 0.12 },
    ];
    const body = renderWeights(weights);
    expect(rowFor(body, 'strength')).toContain('1.00');
    expect(rowFor(body, 'strength')).toContain('0.120 DPS');
    expect(rowFor(body, 'attack_power')).toContain('0.50');
    expect(body).toContain('normalized to Strength = 1.00');
    expect(body).not.toContain('Pawn');
    expect(body.indexOf('sim-weight-strength')).toBeLessThan(body.indexOf('sim-weight-attack_power'));
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
    expect(body).not.toContain(bisCopy.weightsNotSignificant);
  });
});

// Task 8, sub-item 1: contract 10.9's error-is-a-lower-bound caveat, printed once beside the
// table -- not per row, and not only when a row happens to be greyed (the fixture below
// carries none), since the caveat is about every row's own ± figure, not only the ones D45
// already greys.
describe('StatWeights: the error-is-a-lower-bound caveat (Task 8, sub-item 1)', () => {
  it('renders once, in plain words, beside the table', () => {
    const weights: StatWeight[] = [
      { stat: 'attack_power', weight: 1, error: 0 },
      { stat: 'strength', weight: 2.14, error: 0.06 },
    ];
    const body = renderWeights(weights);
    const caveat = bulkCopy.weightsErrorCaveat(weights.map((row) => row.stat));
    expect(body).toContain(caveat);
    // Plain words, not the contract citation a player would have to look up.
    expect(caveat).not.toMatch(/10\.9|lower bound/i);
  });

  it('says nothing about the error bar when there are no weights to caution about yet', () => {
    const body = renderWeights([]);
    expect(body).not.toContain(bulkCopy.weightsErrorCaveat([]));
  });
});

// Task 8, sub-item 2: the precision control discloses the real engine cost -- not the wire's
// own nominal "N iterations" -- for whatever precision and however many stats are currently
// ticked, so "Normal" costs something the reader chooses knowingly rather than discovers a
// minute later.
describe('StatWeights: the precision control discloses the real cost (Task 8, sub-item 2)', () => {
  it('shows the true total for the default precision and the ticked stat count, not the wire’s bare iteration count', () => {
    const weights: StatWeight[] = [
      { stat: 'attack_power', weight: 1, error: 0 },
      { stat: 'strength', weight: 2.14, error: 0.06 },
      { stat: 'crit', weight: 14.37, error: 22.55 },
    ];
    const store = storeWith(weights);
    // createBulkStore's own default precision is 'fast'; the browser lane's own guarded
    // count for it is WEIGHTS_BROWSER_DEFAULT_ITERATIONS, not PRECISION_ITERATIONS.fast.
    const total = weightsEngineIterations(WEIGHTS_BROWSER_DEFAULT_ITERATIONS, store.stats.length);
    const { body } = render(StatWeights, { props: { store, me: null } });
    expect(body).toContain(bulkCopy.weightsCostNote(total));
    // Not the old, misleadingly bare per-tier number a player would otherwise read as the
    // real cost.
    expect(body).not.toContain('Fast, 500 iterations');
  });
});
