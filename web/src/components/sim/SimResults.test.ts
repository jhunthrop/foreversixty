// web/src/components/sim/SimResults.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import SimResults from './SimResults.svelte';
import { afterSimCopy } from '../../lib/sim/after-sim-copy';
import summaryFixture from '../../fixtures/sim/result.json';
import type { Summary } from '../../lib/report/types';
import type { Estimate } from '../../lib/sim/types';

// The fixture's JSON-inferred literal types (e.g. a `misses` object missing one key on some
// rows) don't structurally satisfy `Summary`/`Estimate` -- the same `as unknown as <Type>`
// cast every other sim test using this fixture already applies (details.test.ts,
// margin-agreement.test.ts), rather than trusting the raw JSON shape.
const BASE_PROPS = {
  summary: summaryFixture.summary as unknown as Summary,
  estimate: summaryFixture.dps as Estimate,
  iterationsRun: summaryFixture.iterations_run,
  actionNames: null,
  sample: undefined,
};

describe('SimResults', () => {
  it('the first line points at Top Gear when there is no recorded upgrade', () => {
    const { body } = render(SimResults, { props: { ...BASE_PROPS, topUpgrade: null } });
    expect(body).toContain(afterSimCopy.noUpgrade);
    expect(body).toContain('href="/sim/gear"');
  });

  it('the first line names the item, source and gain when an upgrade is on record', () => {
    const { body } = render(SimResults, {
      props: {
        ...BASE_PROPS,
        topUpgrade: { itemName: 'Bracers of X', sourceName: 'Blackfathom Deeps', gain: '+14', savedAt: '' },
      },
    });
    expect(body).toContain('Upgrade: Bracers of X from Blackfathom Deeps, +14 DPS.');
  });
});
