import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { serverRunsCopy } from '../../lib/sim/server-runs-copy';
import type { SimListRow } from '../../lib/sim/types';
import ServerRuns from './ServerRuns.svelte';

const NOW = new Date('2026-10-24T12:00:00Z');

function row(index: number): SimListRow {
  return {
    sim_id: `sim-${index}`,
    spec: 'mage-frost',
    dps: 1200 + index,
    engine_version: 'e',
    created_at: `2026-10-24T11:0${index}:00Z`,
    title: '',
    state: 'done',
    seen: true,
  };
}

function renderList(rows: SimListRow[] | null, error: string | null = null): string {
  return render(ServerRuns, { props: { rows, error, now: NOW, onretry: () => {}, onstarted: () => {} } })
    .body;
}

describe('ServerRuns', () => {
  it('has the heading, the section id and the test id', () => {
    const body = renderList([row(1)]);
    expect(body).toContain('id="server-runs"');
    expect(body).toContain('data-testid="server-runs"');
    expect(body).toContain(serverRunsCopy.heading);
  });

  it('loading: a busy skeleton box with only screen reader text', () => {
    const body = renderList(null);
    expect(body).toContain('aria-busy="true"');
    expect(body).toContain(serverRunsCopy.loadingSr);
    expect(body).toContain('server-runs-rows-reserve');
  });

  it('empty: one EmptyState with the instruction and no action', () => {
    const body = renderList([]);
    expect(body).toContain('data-testid="server-runs-empty"');
    expect(body).toContain('Run on our servers');
    expect(body).not.toContain('<a');
  });

  it('error: the failure line and Try again', () => {
    const body = renderList(null, 'boom');
    expect(body).toContain(serverRunsCopy.loadFailed);
    expect(body).toContain('Try again');
    expect(body).not.toContain('boom');
  });

  it('shows the newest five and the overflow footer when there are more', () => {
    const body = renderList([1, 2, 3, 4, 5, 6, 7].map(row));
    expect(body.match(/data-testid="server-run-sim-/g)).toHaveLength(5);
    expect(body).toContain('server-run-sim-7');
    expect(body).not.toContain('server-run-sim-1"');
    expect(body).toContain(serverRunsCopy.overflow);
  });

  it('shows no overflow footer at five rows or fewer', () => {
    expect(renderList([1, 2, 3, 4, 5].map(row))).not.toContain(serverRunsCopy.overflow);
  });
});
