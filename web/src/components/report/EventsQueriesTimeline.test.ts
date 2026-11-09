// web/src/components/report/EventsQueriesTimeline.test.ts
// The Events, Queries and Timeline tab bodies in the flat system (logs spec 2026-11-07 section
// 6.2): each is one Section holding the one Plate, no card, and their rails are their own
// Sections. A static render (svelte/server), as ResourceGraphs.test.ts does.
import { render } from 'svelte/server';
import type { Component } from 'svelte';
import { describe, expect, it } from 'vitest';
import fixtureSummary from '../../fixtures/report/fights/3/summary.json';
import type { Summary } from '../../lib/report/types';
import EventsRail from './EventsRail.svelte';
import EventsView from './EventsView.svelte';
import QueriesRail from './QueriesRail.svelte';
import QueriesView from './QueriesView.svelte';
import TimelinesView from './TimelinesView.svelte';

const summary = fixtureSummary as Summary;
const classOf = new Map(summary.roster.map((row) => [row.guid, row.class ?? '']));

function html<P extends Record<string, unknown>>(component: Component<P>, props: P): string {
  return render(component, { props }).body.replace(/<!--[^>]*-->/g, '');
}

const plates = (body: string): number => body.match(/data-region="plate"/g)?.length ?? 0;

describe('EventsView', () => {
  const body = html(EventsView, { summary, classOf });

  it('is a Section holding one Plate with the events table', () => {
    expect(body).toContain('data-testid="events-view"');
    expect(plates(body)).toBe(1);
    expect(body).toContain('data-testid="event-list"');
    expect(body).toMatch(/events in this window, in order/);
  });

  it('colours a player’s name in their class colour', () => {
    const death = summary.deaths[0];
    expect(body).toContain(death.name.split('-')[0]);
    expect(body).toMatch(/font-semibold" style="color: var\(--color-class-/);
  });

  it('draws the Show chips until the rail takes them', () => {
    expect(body).toContain('data-testid="events-chips"');
    expect(html(EventsView, { summary, classOf, chips: false })).not.toContain('data-testid="events-chips"');
  });

  it('says No rows matched. when nothing is left', () => {
    expect(html(EventsView, { summary, classOf, search: 'zzzz-nothing' })).toContain('No rows matched.');
  });
});

describe('EventsRail', () => {
  it('is the Show panel with the chips checked by the url', () => {
    const body = html(EventsRail, { off: ['aura-removed'] });
    expect(body).toContain('>Show</h2>');
    expect(body).toMatch(/events-chip-auras"/);
    expect(plates(body)).toBe(1);
  });
});

describe('QueriesView', () => {
  const body = html(QueriesView, {
    dataBaseUrl: '/sample-report/sample',
    fightIndex: 3,
    window: { startMs: 0, endMs: 60_000 },
  });

  it('is one Plate with the starting points, the SQL and Run', () => {
    expect(plates(body)).toBe(1);
    expect(body).toContain('role="radiogroup" aria-label="Starting points"');
    expect(body).toContain('data-testid="template-event-counts"');
    expect(body).toContain('data-testid="query-sql"');
    expect(body).toContain('data-testid="query-run"');
  });

  it('rails the parquet’s columns', () => {
    const rail = html(QueriesRail, {});
    expect(rail).toContain('>Columns</h2>');
    expect(rail).toContain('<li>time_unix_nano</li>');
  });
});

describe('TimelinesView', () => {
  const body = html(TimelinesView, {
    summary,
    classOf,
    window: { startMs: 0, endMs: summary.duration_ms },
  });

  it('is one Plate with a lane per player and a legend', () => {
    expect(plates(body)).toBe(1);
    expect(body).toContain(`data-testid="lane-${summary.roster[0].guid}"`);
    expect(body).toContain('Buff up');
    expect(body).toContain('grid-cols-[96px_minmax(0,1fr)] md:grid-cols-[180px_minmax(0,1fr)]');
  });

  it('reads Nothing in this window. with nobody in it', () => {
    const empty = html(TimelinesView, {
      summary: { ...summary, roster: [] },
      classOf,
      window: { startMs: 0, endMs: 1000 },
    });
    expect(empty).toContain('Nothing in this window.');
  });

  it('marks each phase change inside the window down the lanes, never the pull start', () => {
    const half = Math.round(summary.duration_ms / 2);
    const phased = html(TimelinesView, {
      summary: {
        ...summary,
        phases: [
          { name: 'Phase 1', start_ms: 0, end_ms: half },
          { name: 'Phase 2', start_ms: half, end_ms: summary.duration_ms },
        ],
      },
      classOf,
      window: { startMs: 0, endMs: summary.duration_ms },
    });
    expect(phased.match(/data-testid="timeline-phase"/g)).toHaveLength(1);
    expect(phased).toContain('title="Phase 2"');
  });
});
