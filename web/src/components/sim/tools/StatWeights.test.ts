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

/** One stat's StatRow, from its own opening tag up to the next row (or the table's end). */
function rowFor(body: string, stat: string): string {
  const match = new RegExp(
    `<div[^>]*data-testid="sim-weight-${stat}"[^>]*>.*?(?=<div[^>]*class="ui-st-r|$)`,
    's',
  ).exec(body);
  if (match === null) throw new Error(`no row rendered for stat "${stat}"`);
  return match[0];
}

/** A greyed row: its primary figure is drawn muted (the StatRow's not-significant reading). */
const GREYED = 'data-primary="true" data-muted="true"';

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
    expect(critRow).toContain(GREYED);
    expect(critRow).toContain(bisCopy.weightsNotSignificant);
    expect(strengthRow).not.toContain(GREYED);
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
    // The Scale column says the normalisation; the caption keeps only haste (h-audit 2026-11-09).
    expect(body).not.toContain('normalized to');
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
    expect(rowFor(body, 'attack_power')).not.toContain(GREYED);
    expect(rowFor(body, 'strength')).not.toContain(GREYED);
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
