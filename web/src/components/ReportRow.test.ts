// web/src/components/ReportRow.test.ts
// Static-render check, the pattern sim/HelpNote.test.ts and sim/SavedWeights.test.ts use for
// a component whose whole surface is its props: ReportRow takes no effects and no client
// state, so the server render is the real render.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import ReportRow from './ReportRow.svelte';

const BASE = {
  href: '/reports/abc123def456',
  title: 'Progress night',
  createdAt: '2026-12-09T22:10:00Z',
  fightCount: 8,
  killCount: 3,
};

describe('ReportRow', () => {
  it('links the title to the report and shows the day, not the full timestamp', () => {
    const { body } = render(ReportRow, { props: BASE });
    expect(body).toContain('href="/reports/abc123def456"');
    expect(body).toContain('Progress night');
    expect(body).toContain('2026-12-09');
    expect(body).not.toContain('22:10:00');
  });

  it('shows the fight and kill counts', () => {
    const { body } = render(ReportRow, { props: BASE });
    expect(body).toContain('8 fights');
    expect(body).toContain('3 kills');
  });

  it('appends meta after the counts when given, and omits the separator when not', () => {
    const withMeta = render(ReportRow, { props: { ...BASE, meta: 'The Last Watch' } }).body;
    expect(withMeta).toContain('8 fights · 3 kills</span> · The Last Watch');

    const withoutMeta = render(ReportRow, { props: BASE }).body;
    expect(withoutMeta).not.toContain('· ·');
    expect(withoutMeta).toMatch(/8 fights · 3 kills<\/span>\s*<\/span>/);
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
