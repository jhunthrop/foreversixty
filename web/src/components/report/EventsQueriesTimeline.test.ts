// web/src/components/report/EventsQueriesTimeline.test.ts
// The Events, Queries and Timeline tab bodies in round 5 (b1-report-events, b1-report-queries,
// b1-report-timeline): each is a panel with its own filter chips under the head, and the Queries
// rail is its own Section. A static render (svelte/server), as ResourceGraphs.test.ts does.
import { render } from 'svelte/server';
import type { Component } from 'svelte';
import { describe, expect, it } from 'vitest';
import fixtureSummary from '../../fixtures/report/fights/3/summary.json';
import type { Summary } from '../../lib/report/types';
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
    expect(body).toMatch(/[\d,]+ in this window/);
    expect(body).toContain('data-dense="true"');
  });

  it('colours a player’s name in their class colour', () => {
    const death = summary.deaths[0];
    expect(body).toContain(death.name.split('-')[0]);
    expect(body).toMatch(/font-semibold" style="color: var\(--color-class-/);
  });

  it('draws the Show chips under the head, each a toggle the url keeps', () => {
    expect(body).toContain('role="group" aria-label="Show" data-testid="events-chips"');
    const off = html(EventsView, { summary, classOf, off: ['aura-applied', 'aura-removed'] });
    expect(off).toMatch(/aria-pressed="false"[^>]*data-testid="events-chips-auras"/);
  });

  it('says No rows matched. when nothing is left', () => {
    expect(html(EventsView, { summary, classOf, search: 'zzzz-nothing' })).toContain('No rows matched.');
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
    expect(body).toContain('role="group" aria-label="Starting points"');
    expect(body).toMatch(/aria-pressed="true"[^>]*data-testid="template-damage-by-ability"/);
    expect(body).toContain('data-testid="template-event-counts"');
    expect(body).toContain('data-testid="query-sql"');
    expect(body).toContain('data-testid="query-run"');
  });

  it('rails the common columns, every one behind All', () => {
    const rail = html(QueriesRail, {});
    expect(rail).toContain('>Columns</h2>');
    expect(rail).toContain('<li>time_unix_nano</li>');
    expect(rail).not.toContain('<li>adv_level</li>');
    expect(rail).toMatch(/data-testid="queries-rail-all"[^>]*>All 66</);
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
    expect(body).toContain('role="group" aria-label="Show" data-testid="timeline-show"');
    expect(body).toContain('>Buffs</button>');
    expect(body).toContain('grid-cols-[124px_minmax(0,1fr)] md:grid-cols-[180px_minmax(0,1fr)]');
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
