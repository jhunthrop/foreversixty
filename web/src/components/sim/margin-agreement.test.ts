// web/src/components/sim/margin-agreement.test.ts
// Task 3: RunControl's headline band and DetailsCard's "Margin of error" row must never
// disagree -- both read from estimate.ts's one formatMargin, so a reader who checks the
// details card against the number beside the DPS figure always finds the same figure. A
// static-render check (svelte/server, no jsdom), the pattern tools/DropResults.test.ts
// already uses for a component whose whole surface is a few props.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import fixture from '../../fixtures/sim/result.json';
import type { Lane, PrecisionId } from '../../lib/sim/precision';
import type { SimPhase } from '../../lib/sim/store.svelte';
import type { Estimate, SimResult } from '../../lib/sim/types';
import DetailsCard from './DetailsCard.svelte';
import RunControl from './RunControl.svelte';

// 1.96 * 4 = 7.84, which formatMargin renders as one decimal (below 10): "7.8".
const dps: Estimate = { mean: 1000, stddev: 100, error: 4, min: 0, max: 0 };

function renderHeadline(estimate: Estimate): string {
  const { body } = render(RunControl, {
    props: {
      phase: 'done' as SimPhase,
      estimate,
      iterationsDone: 3000,
      iterationsTotal: 3000,
      precisionId: 'normal' as PrecisionId,
      relativeError: 0.004,
      lane: 'browser' as Lane,
      premium: false,
      message: null,
      detail: '',
      racePending: false,
      staleVersion: null,
      serverRunning: false,
      onrun: () => {},
      onstop: () => {},
      onprecision: () => {},
      onserver: () => {},
      onrerun: () => {},
    },
  });
  return body;
}

function renderDetails(result: SimResult): string {
  const { body } = render(DetailsCard, { props: { result } });
  return body;
}

describe('the margin reads the same on the headline and the details card', () => {
  it('shows the identical "± N" figure in both, from the same formatter', () => {
    const result = { ...(fixture as unknown as SimResult), dps };
    expect(renderHeadline(dps)).toContain('± 7.8');
    expect(renderDetails(result)).toContain('± 7.8');
  });
});
