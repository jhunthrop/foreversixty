// web/src/components/sim/SavedWeights.test.ts
// D45/D46's own component-level check: a weight the engine flagged `insignificant` renders
// greyed with the "not distinguishable from zero" label, and the Pawn string beside the
// table never contains it -- the table and the Pawn string can never disagree, because both
// come off the identical rendered output here. A static-render check (svelte/server, no
// jsdom), the pattern margin-agreement.test.ts and tools/DropResults.test.ts already use for
// a component whose whole surface is a few props -- `SavedWeights` takes exactly one,
// `result`, unlike `StatWeights.svelte`'s stateful `BulkStore`.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import fixture from '../../fixtures/sim/weights-result.json';
import { bulkCopy, WEIGHT_INSIGNIFICANT_LABEL } from '../../lib/sim/copy';
import type { WeightsResult } from '../../lib/sim/bulk-types';
import SavedWeights from './SavedWeights.svelte';

const result = fixture as unknown as WeightsResult;

function renderWeights(next: WeightsResult): string {
  const { body } = render(SavedWeights, { props: { result: next } });
  return body;
}

/** The whole `<li data-testid="sim-weight-<stat>">…</li>` element for one row, so a class
 *  that renders before the test id (it does: `class` comes before `data-testid`) is not
 *  lost the way slicing the body on the test id string alone would lose it. */
function rowFor(body: string, stat: string): string {
  const match = new RegExp(`<li[^>]*data-testid="sim-weight-${stat}"[^>]*>.*?</li>`, 's').exec(body);
  if (match === null) throw new Error(`no <li> rendered for stat "${stat}"`);
  return match[0];
}

describe('SavedWeights: the greyed row and its label (D45)', () => {
  it('greys exactly the row the engine flagged insignificant -- the fixture’s own agility', () => {
    const body = renderWeights(result);
    const agilityRow = rowFor(body, 'agility');
    const strengthRow = rowFor(body, 'strength');
    expect(agilityRow).toContain('opacity-50');
    expect(agilityRow).toContain(WEIGHT_INSIGNIFICANT_LABEL);
    expect(strengthRow).not.toContain('opacity-50');
    expect(strengthRow).not.toContain(WEIGHT_INSIGNIFICANT_LABEL);
  });

  it('D46: the Pawn string omits exactly the greyed row, never disagreeing with the table', () => {
    const body = renderWeights(result);
    // Agility's own Pawn key ("Agility=") must not appear -- it is both greyed in the
    // table above and the row the D46 repro named as disagreeing with the Pawn string.
    expect(body).not.toContain('Agility=');
    // Every other row's own Pawn key IS in the string: the table's included rows and the
    // Pawn string's are the same set.
    expect(body).toContain('Strength=1.50');
    expect(body).toContain('AttackPower=1.00');
    expect(body).toContain('CritRating=7.03');
    expect(body).toContain('HitRating=6.87');
    expect(body).toContain('HasteRating=5.85');
  });

  it('a result with no insignificant rows greys nothing', () => {
    const allSignificant: WeightsResult = {
      ...result,
      weights: result.weights.map((row) => ({ ...row, insignificant: false })),
    };
    const body = renderWeights(allSignificant);
    expect(body).not.toContain('opacity-50');
    expect(body).not.toContain(WEIGHT_INSIGNIFICANT_LABEL);
    expect(body).toContain('Agility=0.16');
  });
});

// Task 8, sub-item 1: the same caveat travels with a saved result -- a shared link is
// exactly where someone meets these weights without having read the live page (this file's
// own header comment), so the caution has to be here too, in the same words.
describe('SavedWeights: the error-is-a-lower-bound caveat (Task 8, sub-item 1)', () => {
  it('renders once, in the same words as the live page', () => {
    const body = renderWeights(result);
    expect(body).toContain(bulkCopy.weightsErrorCaveat);
  });
});
