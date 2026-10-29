import { readFileSync } from 'node:fs';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { serverRunsCopy } from '../../lib/sim/server-runs-copy';
import type { RerunStatus } from '../../lib/sim/server-runs-poll';
import type { SimListRow } from '../../lib/sim/types';
import ServerRunRow from './ServerRunRow.svelte';

const NOW = new Date('2026-10-24T12:00:00Z');

function row(overrides: Partial<SimListRow>): SimListRow {
  return {
    sim_id: 'sim-1',
    spec: 'mage-frost',
    dps: 1204,
    engine_version: 'e',
    created_at: '2026-10-24T11:58:00Z',
    title: '',
    kind: 'run',
    lane: 'server',
    ...overrides,
  };
}

function text(html: string): string {
  return html
    .replace(/<!--.*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/\s+/g, ' ')
    .trim();
}

function renderRow(r: SimListRow, rerun: RerunStatus = 'idle'): string {
  return render(ServerRunRow, { props: { row: r, now: NOW, rerun, onrerun: () => {} } }).body;
}

describe('ServerRunRow states', () => {
  it('queued: Starting, started just now, not a link, no button', () => {
    const body = renderRow(row({ state: 'queued', created_at: '2026-10-24T11:59:50Z' }));
    expect(text(body)).toContain('Starting');
    expect(text(body)).toContain('Started just now');
    expect(body).not.toContain('<a');
    expect(body).not.toContain('<button');
    expect(body).not.toContain('server-run-dot-pulse');
  });

  it('running with an estimate: Running, takes about 3 minutes, pulsing dot', () => {
    const body = renderRow(row({ state: 'running', estimate_seconds: 180 }));
    expect(text(body)).toContain('Running · takes about 3 minutes');
    expect(text(body)).toContain('Started 2 minutes ago');
    expect(body).toContain('server-run-dot-pulse');
    expect(body).not.toContain('<a');
    expect(body).not.toContain('<button');
  });

  it('running overdue: taking longer than usual', () => {
    const body = renderRow(
      row({ state: 'running', estimate_seconds: 60, created_at: '2026-10-24T11:51:00Z' }),
    );
    expect(text(body)).toContain('Running · taking longer than usual');
    expect(text(body)).toContain('Started 9 minutes ago');
  });

  it('running without an estimate: Running alone', () => {
    const body = renderRow(row({ state: 'running' }));
    expect(text(body)).toContain('Running');
    expect(text(body)).not.toContain('takes about');
    expect(text(body)).not.toContain('taking longer');
  });

  it('done and opened: a whole-row link to the result with the arrow, no dot', () => {
    const body = renderRow(
      row({ state: 'done', seen: true, headline: '1,204 DPS', finished_at: '2026-10-24T11:46:00Z' }),
    );
    expect(body).toContain('href="/sim/sim-1"');
    expect(text(body)).toContain('1,204 DPS');
    expect(text(body)).toContain('Finished 14 minutes ago');
    expect(body).toContain('server-run-arrow');
    expect(body).not.toContain(serverRunsCopy.unreadSr);
    expect(body).not.toContain('<button');
  });

  it('done and unopened: the gold dot with its screen reader text and the tint class', () => {
    const body = renderRow(row({ state: 'done', seen: false, finished_at: '2026-10-24T11:46:00Z' }));
    expect(body).toContain(serverRunsCopy.unreadSr);
    expect(body).toContain('server-run-unseen');
    expect(text(body)).toContain('1,204 DPS');
  });

  it('error: the reason in plain words, Ended, a Run again button and no link', () => {
    const body = renderRow(
      row({
        state: 'error',
        failure_reason: 'stopped before completion',
        finished_at: '2026-10-24T10:00:00Z',
      }),
    );
    expect(text(body)).toContain('Failed: stopped before completion');
    expect(text(body)).toContain('Ended 2 hours ago');
    expect(text(body)).toContain('Run again');
    expect(body).not.toContain('href="/sim/');
  });

  it('error without a reason reads Failed alone', () => {
    const body = renderRow(row({ state: 'error', finished_at: '2026-10-24T10:00:00Z' }));
    expect(text(body)).toContain('Failed Ended');
  });

  it('titled rows show the title and the spec in the meta line', () => {
    const body = renderRow(row({ state: 'done', seen: true, title: 'Raid night' }));
    expect(text(body)).toContain('Raid night');
    expect(text(body)).toContain('Frost Mage · Sim');
  });

  it('untitled rows show the spec label and the kind alone', () => {
    const body = renderRow(row({ state: 'done', seen: true, kind: 'gear' }));
    expect(text(body)).toContain('Frost Mage');
    expect(text(body)).toContain('Gear');
  });

  it('shows no queue position, cancel or percentage', () => {
    for (const state of ['queued', 'running', 'done', 'error'] as const) {
      const body = text(renderRow(row({ state, estimate_seconds: 180 })));
      expect(body).not.toMatch(/cancel|position|queue|%/i);
    }
  });
});

describe('Run again', () => {
  const failed = row({ state: 'error', finished_at: '2026-10-24T10:00:00Z' });

  it('pending: Starting…, disabled and busy', () => {
    const body = renderRow(failed, 'pending');
    expect(text(body)).toContain('Starting…');
    expect(body).toContain('disabled');
    expect(body).toContain('aria-busy="true"');
  });

  it('started: Started, disabled', () => {
    const body = renderRow(failed, 'started');
    expect(text(body)).toContain('Started');
    expect(body).toContain('disabled');
  });

  it('failed: the could-not-start line, the button back at rest', () => {
    const body = renderRow(failed, 'failed');
    expect(text(body)).toContain(serverRunsCopy.runAgainFailed);
    expect(body).not.toContain('disabled');
    expect(body).not.toContain('/premium');
  });

  it('lapsed (402): the Premium line and See plans to /premium', () => {
    const body = renderRow(failed, 'lapsed');
    expect(text(body)).toContain(serverRunsCopy.planLapsed);
    expect(body).toContain('href="/premium#server-sims"');
    expect(text(body)).toContain(serverRunsCopy.seePlans);
    expect(text(body)).not.toContain(serverRunsCopy.runAgainFailed);
  });
});

describe('the pulse', () => {
  it('is switched off under prefers-reduced-motion in the component stylesheet', () => {
    const source = readFileSync(new URL('./ServerRunRow.svelte', import.meta.url), 'utf8');
    expect(source).toMatch(
      /@media \(prefers-reduced-motion: reduce\)[\s\S]*server-run-dot-pulse[\s\S]*animation: none/,
    );
  });
});
