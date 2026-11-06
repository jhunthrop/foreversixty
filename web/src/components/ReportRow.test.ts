// web/src/components/ReportRow.test.ts
// Static-render check: ReportRow takes no effects and no client state, so the server render is
// the real render (logs spec 2026-11-07 sections 4 and 5).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import ReportRow from './ReportRow.svelte';

const THREE_DAYS_MS = 3 * 24 * 60 * 60 * 1000;

const BASE = {
  href: '/reports/abc123def456',
  title: 'Progress night',
  createdAt: new Date(Date.now() - THREE_DAYS_MS - 60_000).toISOString(),
  summary: "Onyxia's Lair · 1 boss · 2 fights · Not credited",
};

const html = (props: Record<string, unknown>) =>
  render(ReportRow, { props: { ...BASE, ...props } }).body.replace(/<!--[^>]*-->/g, '');

describe('ReportRow', () => {
  it('is one link to the report, with a relative age inside two weeks', () => {
    const body = html({});
    expect(body).toMatch(/^<a[^>]*href="\/reports\/abc123def456"/);
    expect(body).toContain('Progress night');
    expect(body).toContain('3 days ago');
  });

  it('shows a date from two weeks on', () => {
    expect(html({ createdAt: '2026-09-15T12:55:13Z' })).toContain('Sep 15');
  });

  it('shows the line the caller built under the title', () => {
    expect(html({})).toContain("Onyxia's Lair · 1 boss · 2 fights · Not credited");
  });

  it('shows a status Tag only while the report is not finished', () => {
    const live = html({ status: 'live' });
    expect(live).toContain('data-testid="report-status"');
    expect(live).toContain('Live');
    for (const status of ['complete', 'ready', '']) {
      expect(html({ status })).not.toContain('data-testid="report-status"');
    }
  });

  it("carries the caller's own Tag (the sample's) and its own date", () => {
    const body = html({ tag: 'Sample report', dateLabel: '17 Sep 2026' });
    expect(body).toContain('Sample report');
    expect(body).toContain('17 Sep 2026');
    expect(body).not.toContain('days ago');
  });

  it('adds the credit hint only when given', () => {
    expect(html({})).not.toContain('data-testid="report-row-hint"');
    expect(html({ status: 'live', creditHint: 'Log out once in game.' })).toContain('Log out once in game.');
  });

  it('draws no surface of its own: the row is the ui Row', () => {
    expect(html({})).toContain('ui-row');
    expect(html({})).not.toMatch(/rounded-panel|bg-raised|card-surface/);
  });
});
