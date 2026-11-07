// web/src/components/sim/SavedWeights.test.ts
// A saved weights run reads as the published weights do (owner, 2026-09-30; round 11): top
// stat 1.00, sorted, no Pawn string; a weight the engine flagged `insignificant` renders
// greyed and "Not significant". A static-render check (svelte/server, no
// jsdom), the pattern margin-agreement.test.ts and tools/DropResults.test.ts already use for
// a component whose whole surface is a few props -- `SavedWeights` takes exactly one,
// `result`, unlike `StatWeights.svelte`'s stateful `BulkStore`.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import fixture from '../../fixtures/sim/weights-result.json';
import { bisCopy } from '../../lib/bis/copy';
import { bulkCopy } from '../../lib/sim/copy';
import type { WeightsResult } from '../../lib/sim/bulk-types';
import SavedWeights from './SavedWeights.svelte';

const result = fixture as unknown as WeightsResult;
/** A greyed weight's name (WeightsScaleTable's StatRow): muted. */
const GREYED = '<span class="text-muted">';

function renderWeights(next: WeightsResult): string {
  const { body } = render(SavedWeights, { props: { result: next } });
  return body;
}

/** One stat's StatRow markup: from its own row up to the next row (or the end of the page). */
function rowFor(body: string, stat: string): string {
  const match = new RegExp(
    `<div[^>]*data-testid="sim-weight-${stat}"[^>]*>.*?(?=data-testid="sim-weight-|$)`,
    's',
  ).exec(body);
  if (match === null) throw new Error(`no row rendered for stat "${stat}"`);
  return match[0];
}

describe('SavedWeights: the greyed row and its label (D45)', () => {
  it('greys exactly the row the engine flagged insignificant -- the fixture’s own agility', () => {
    const body = renderWeights(result);
    const agilityRow = rowFor(body, 'agility');
    const strengthRow = rowFor(body, 'strength');
    expect(agilityRow).toContain(GREYED);
    expect(agilityRow).toContain(bisCopy.weightsNotSignificant);
    expect(strengthRow).not.toContain(GREYED);
    expect(strengthRow).not.toContain(bisCopy.weightsNotSignificant);
  });

  it('normalizes to the top per-point stat, puts haste in the caption, and has no Pawn string', () => {
    const body = renderWeights(result);
    expect(body).not.toContain('Pawn');
    expect(body).toContain('= 1.00');
    expect(body).toMatch(/Haste: [0-9.]+ per 1%/);
    expect(() => rowFor(body, 'melee_haste')).toThrow();
  });

  // The shared weight reading (lib/bis/weight-reading.ts) also greys a weight inside its own
  // error, so a fully significant result is one with no flag and every weight beyond its error.
  it('a result with no insignificant rows greys nothing', () => {
    const allSignificant: WeightsResult = {
      ...result,
      weights: result.weights.map((row) => ({
        ...row,
        insignificant: false,
        error: Math.abs(row.weight) / 10,
      })),
    };
    const body = renderWeights(allSignificant);
    expect(body).not.toContain(GREYED);
    expect(body).not.toContain(bisCopy.weightsNotSignificant);
  });
});

// Task 8, sub-item 1: the same caveat travels with a saved result -- a shared link is
// exactly where someone meets these weights without having read the live page (this file's
// own header comment), so the caution has to be here too, in the same words.
describe('SavedWeights: the error-is-a-lower-bound caveat (Task 8, sub-item 1)', () => {
  it('renders once, in the same words as the live page, naming this result’s own weighed stats', () => {
    const body = renderWeights(result);
    expect(body).toContain(bulkCopy.weightsErrorCaveat(result.weights.map((row) => row.stat)));
  });
});
