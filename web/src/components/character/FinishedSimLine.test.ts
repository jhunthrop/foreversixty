// The line's markup, rendered statically (svelte/server, the project's component-test setup).
// Its dismiss flow needs a browser: the seen fan-out is unit-tested in
// lib/character-header/finished-sim-line.test.ts, and the busy, failure-then-restore and
// focus behaviour in tests/e2e/header-finished-sim.spec.ts.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { FinishedSim } from '../../lib/account/api';
import { serverRunsCopy } from '../../lib/sim/server-runs-copy';
import FinishedSimLine from './FinishedSimLine.svelte';

function run(id: string, overrides: Partial<FinishedSim> = {}): FinishedSim {
  return {
    sim_id: id,
    spec: 'mage-frost',
    title: '',
    dps: 1204.4,
    finished_at: '2026-10-24T11:00:00Z',
    ...overrides,
  };
}

function body(runs: [FinishedSim, ...FinishedSim[]]): string {
  return render(FinishedSimLine, { props: { runs, onDismissed: () => {} } }).body;
}

describe('FinishedSimLine', () => {
  it('one unseen run: the sentence, the figure, the link, no more-link, a status region', () => {
    const html = body([run('sim-1', { title: 'Frost night' })]);
    expect(html).toContain('role="status"');
    expect(html).toContain('href="/sim/sim-1"');
    expect(html).toContain('title="Frost night"');
    expect(html).toContain('Your Frost Mage sim finished: ');
    expect(html).toMatch(/class="fsl-figure[^"]*">1,204 DPS</);
    expect(html).toContain(`>${serverRunsCopy.bandOpen}<`);
    expect(html).not.toContain('finished-sim-more');
    expect(html).toContain('aria-label="Dismiss: your Frost Mage sim finished"');
    expect(html).toContain('aria-busy="false"');
  });

  it('phone: the arrow stands in for Open result inside the same link', () => {
    const html = body([run('sim-1')]);
    expect(html).toMatch(/data-testid="finished-sim-open-phone">→</);
    expect(html.match(/<a /g)).toHaveLength(1);
  });

  it('three unseen runs: names the newest and links the other two to the list', () => {
    const html = body([run('sim-3'), run('sim-2'), run('sim-1')]);
    expect(html).toContain('href="/sim/sim-3"');
    expect(html).toMatch(/href="\/sim#server-runs"[^>]*>and 2 more</);
  });
});
