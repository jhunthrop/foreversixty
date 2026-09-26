// web/src/components/sim/SimHistory.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import SimHistory from './SimHistory.svelte';
import { simCopy } from '../../lib/sim/copy';

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
});
