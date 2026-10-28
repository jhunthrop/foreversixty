// SSR render only (this project's vitest config hands components Svelte's server entry): the
// first paint is the reserved loading slot and makes no request. The settled states are covered
// in lib/coach/last-raid.test.ts and the clicks and the refresh in Playwright.
import { render } from 'svelte/server';
import { describe, expect, it, vi } from 'vitest';
import CoachLastRaidLine from './CoachLastRaidLine.svelte';

describe('CoachLastRaidLine', () => {
  it('paints the reserved slot first, busy, and does not fetch on the server', () => {
    const fetchRecent = vi.fn();
    const { body } = render(CoachLastRaidLine, { props: { premium: true, fetchRecent } });
    expect(body).toContain('data-testid="coach-last-raid"');
    expect(body).toContain('data-state="loading"');
    expect(body).toContain('aria-busy="true"');
    expect(body).toContain('min-h-[77px]');
    expect(body).toContain('md:min-h-[57px]');
    expect(body).toContain('role="status"');
    expect(body).toContain('aria-label="Last raid"');
    expect(body).toContain('coach-last-raid-skeleton');
    expect(fetchRecent).not.toHaveBeenCalled();
  });
});
