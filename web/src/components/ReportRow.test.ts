// web/src/components/ReportRow.test.ts
// Static-render check, the pattern sim/HelpNote.test.ts and sim/SavedWeights.test.ts use for
// a component whose whole surface is its props: ReportRow takes no effects and no client
// state, so the server render is the real render.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import ReportRow from './ReportRow.svelte';

const THREE_DAYS_MS = 3 * 24 * 60 * 60 * 1000;

const BASE = {
  href: '/reports/abc123def456',
  title: 'Progress night',
  createdAt: new Date(Date.now() - THREE_DAYS_MS - 60_000).toISOString(),
  summary: "Onyxia's Lair · 1 boss · 48 m",
};

describe('ReportRow', () => {
  it('links the title to the report and shows a relative age inside two weeks', () => {
    const { body } = render(ReportRow, { props: BASE });
    expect(body).toContain('href="/reports/abc123def456"');
    expect(body).toContain('Progress night');
    expect(body).toContain('3 days ago');
  });

  it('shows a date from two weeks on', () => {
    const { body } = render(ReportRow, { props: { ...BASE, createdAt: '2026-09-15T12:55:13Z' } });
    expect(body).toContain('Sep 15');
  });

  it('shows the summary line the caller built', () => {
    const { body } = render(ReportRow, { props: BASE });
    expect(body).toContain("Onyxia's Lair · 1 boss · 48 m");
  });

  it('shows a status pill only when the status is not complete', () => {
    const live = render(ReportRow, { props: { ...BASE, status: 'live' } }).body;
    expect(live).toContain('data-testid="report-status"');
    expect(live).toContain('live');

    const complete = render(ReportRow, { props: { ...BASE, status: 'complete' } }).body;
    expect(complete).not.toContain('data-testid="report-status"');

    const unset = render(ReportRow, { props: BASE }).body;
    expect(unset).not.toContain('data-testid="report-status"');
  });
});
