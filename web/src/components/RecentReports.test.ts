// web/src/components/RecentReports.test.ts
// Server render only, the same constraint HelpNote.test.ts and SavedWeights.test.ts note:
// this project's vitest config gives `render()` Svelte's server entry, where $effect never
// runs, so the fetch-driven states (rows, empty, failed) are Playwright's job
// (tests/e2e/logs-recent-reports.spec.ts). What a server render can prove is the initial,
// pre-hydration DOM: the loading copy renders in place rather than a blank panel, so the
// panel does not jump when the client takes over -- the no-layout-shift rule spec section 0
// sets for anything that appears after hydration.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { recentReportsCopy } from '../lib/reports/copy';
import RecentReports from './RecentReports.svelte';

describe('RecentReports (initial render)', () => {
  it('shows its own heading by default', () => {
    const { body } = render(RecentReports, { props: {} });
    expect(body).toContain(recentReportsCopy.heading);
  });

  it('omits the heading when told the page already titles the panel', () => {
    const { body } = render(RecentReports, { props: { heading: false } });
    expect(body).not.toContain(recentReportsCopy.heading);
  });

  it('renders the loading copy before the client fetch runs, not a blank panel', () => {
    const { body } = render(RecentReports, { props: {} });
    expect(body).toContain(recentReportsCopy.loading);
  });
});
