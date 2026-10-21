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
    expect(live).toContain('>Live<');

    const complete = render(ReportRow, { props: { ...BASE, status: 'complete' } }).body;
    expect(complete).not.toContain('data-testid="report-status"');

    const unset = render(ReportRow, { props: BASE }).body;
    expect(unset).not.toContain('data-testid="report-status"');
  });

  it('leads the summary line with the age and the credited character, then the badge', () => {
    const { body } = render(ReportRow, {
      props: { ...BASE, summary: 'Darnassus · 0 fights', status: 'live', character: 'Bow Jackzon' },
    });
    const when = body.indexOf('data-testid="report-row-when"');
    const who = body.indexOf('Bow Jackzon');
    const where = body.indexOf('Darnassus · 0 fights');
    const badge = body.indexOf('data-testid="report-status"');
    expect(when).toBeGreaterThan(-1);
    expect(when).toBeLessThan(who);
    expect(who).toBeLessThan(where);
    expect(where).toBeLessThan(badge);
  });

  it('says Not credited when the character is null, with the hint only when given', () => {
    const plain = render(ReportRow, { props: { ...BASE, character: null } }).body;
    expect(plain).toContain('Not credited');
    expect(plain).not.toContain('data-testid="report-row-hint"');
    const hinted = render(ReportRow, {
      props: { ...BASE, status: 'live', character: null, creditHint: 'Log out once in game.' },
    }).body;
    expect(hinted).toContain('Log out once in game.');
  });

  it('leaves the credit part out for the public feed, which passes no character', () => {
    const { body } = render(ReportRow, { props: BASE });
    expect(body).not.toContain('report-row-character');
  });

  // W.9: one row style for every list, the age at the end of the muted line.
  it('draws the title 16 px in the strong colour, never dimmed, the same for every row', () => {
    const feed = render(ReportRow, { props: BASE }).body;
    const sample = render(ReportRow, { props: { ...BASE, dateLabel: '26 Sept 2026' } }).body;
    const owned = render(ReportRow, { props: { ...BASE, character: 'Bow Jackzon' } }).body;
    for (const body of [feed, sample, owned]) {
      const title = /<a [^>]*class="([^"]*report-title[^"]*)"/.exec(body)?.[1] ?? '';
      expect(title).toContain('text-strong');
      expect(title).toContain('text-[16px]');
      expect(title).not.toMatch(/text-muted|opacity/);
      expect(body).not.toContain('reveal');
    }
  });

  it('ends the muted line of a public row with the age, after the summary', () => {
    const { body } = render(ReportRow, { props: BASE });
    const summary = body.indexOf("Onyxia's Lair");
    const when = body.indexOf('data-testid="report-row-when"');
    expect(summary).toBeGreaterThan(-1);
    expect(when).toBeGreaterThan(summary);
    expect(body).toContain('3 days ago');
  });

  it('says the date label a caller gives in place of the relative age', () => {
    const { body } = render(ReportRow, { props: { ...BASE, dateLabel: '26 Sept 2026' } });
    expect(body).toContain('26 Sept 2026');
    expect(body).not.toContain('3 days ago');
  });
});
