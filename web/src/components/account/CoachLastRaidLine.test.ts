// SSR render only (this project's vitest config hands components Svelte's server entry): the
// first paint draws nothing and makes no request; the line appears only once a last raid is read
// (premium spec 2026-11-08 G5). The settled states are covered in lib/coach/last-raid.test.ts and
// the line itself in Playwright (coach-account.spec.ts).
import { render } from 'svelte/server';
import { describe, expect, it, vi } from 'vitest';
import CoachLastRaidLine from './CoachLastRaidLine.svelte';

describe('CoachLastRaidLine', () => {
  it('draws nothing while it loads, and does not fetch on the server', () => {
    const fetchRecent = vi.fn();
    const { body } = render(CoachLastRaidLine, { props: { premium: true, fetchRecent } });
    expect(body).not.toContain('coach-last-raid');
    expect(body.replace(/<!--.*?-->/g, '').trim()).toBe('');
    expect(fetchRecent).not.toHaveBeenCalled();
  });
});
