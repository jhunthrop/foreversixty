// web/src/components/sim/SimHistory.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import SimHistory from './SimHistory.svelte';
import { landingCopy } from '../../lib/sim/landing-copy';
import { simCopy } from '../../lib/sim/copy';
import type { SimListRow } from '../../lib/sim/types';

const ONE_SPEC: SimListRow[] = [
  {
    sim_id: 'a',
    spec: 'warrior-fury',
    dps: 1204,
    engine_version: 'e',
    created_at: '2026-09-14T10:02:00Z',
    title: '',
  },
  {
    sim_id: 'b',
    spec: 'warrior-fury',
    dps: 1245,
    engine_version: 'e',
    created_at: '2026-09-15T10:02:00Z',
    title: '',
  },
];

const TWO_SPECS: SimListRow[] = [
  ...ONE_SPEC,
  {
    sim_id: 'c',
    spec: 'mage-fire',
    dps: 900,
    engine_version: 'e',
    created_at: '2026-09-16T10:02:00Z',
    title: '',
  },
];

describe('SimHistory', () => {
  it('renders the empty state through EmptyState, not an ad-hoc paragraph', () => {
    const { body } = render(SimHistory, {
      props: { rows: [], error: null, kind: 'all', onkind: () => {} },
    });
    expect(body).toContain('data-testid="sim-history-empty"');
    expect(body).toContain(simCopy.historyEmpty);
  });

  it('offers no action -- running a sim is already the page this panel sits on', () => {
    const { body } = render(SimHistory, {
      props: { rows: [], error: null, kind: 'all', onkind: () => {} },
    });
    // EmptyState renders an <a> only when given an `action` prop.
    expect(body).not.toContain('<a');
  });

  // Finding 4, 2026-09-26 layout pass: the best-per-spec table needs two or more specs to
  // say anything a single-spec list does not already say.
  it('renders no best-per-spec table for a character with sims in only one spec', () => {
    const { body } = render(SimHistory, {
      props: { rows: ONE_SPEC, error: null, kind: 'all', onkind: () => {} },
    });
    expect(body).not.toContain('data-testid="sim-best-per-spec"');
  });

  it('renders the best-per-spec table once the character has sims in two or more specs', () => {
    const { body } = render(SimHistory, {
      props: { rows: TWO_SPECS, error: null, kind: 'all', onkind: () => {} },
    });
    expect(body).toContain('data-testid="sim-best-per-spec"');
    expect(body).toContain(landingCopy.bestPerSpecTitle);
    expect(body).toContain('data-testid="sim-best-per-spec-warrior-fury"');
    expect(body).toContain('data-testid="sim-best-per-spec-mage-fire"');
  });

  // Finding 6: the date used to be `hidden md:inline`, dropped from the phone layout
  // entirely. It now always renders (CSS alone repositions it at each breakpoint).
  it('always renders the row date, never hidden outright', () => {
    const { body } = render(SimHistory, {
      props: { rows: ONE_SPEC, error: null, kind: 'all', onkind: () => {} },
    });
    expect(body).not.toContain('hidden md:inline');
    expect(body).toContain('Sept 14');
    expect(body).toContain('Sept 15');
  });
});
