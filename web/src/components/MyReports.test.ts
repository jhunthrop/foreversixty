// web/src/components/MyReports.test.ts
// Server render only, the same constraint RecentReports.test.ts notes: this project's vitest
// config gives `render()` Svelte's server entry, where $effect never runs, so the fetch-driven
// states (rows, failed) are Playwright's job. What a server render can prove is the initial,
// pre-hydration DOM: a Skeleton reserves the ready height, and the retired "Reload the page"
// copy is gone. The signed-out sentence belongs to the page that mounts this (LogsLanding), and
// its e2e proves it.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import MyReports from './MyReports.svelte';

describe('MyReports', () => {
  it('reserves height with a Skeleton while loading', () => {
    const { body } = render(MyReports, { props: {} });
    expect(body).toContain('data-testid="my-reports-skeleton"');
    expect(body).not.toContain('Reload the page');
  });

  it('draws no sign-in control of its own', () => {
    const { body } = render(MyReports, { props: {} });
    expect(body).not.toContain('Sign in');
    expect(body).not.toContain('<button');
  });
});
