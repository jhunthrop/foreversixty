// web/src/components/MyReports.test.ts
// Server render only, the same constraint RecentReports.test.ts notes: this project's vitest
// config gives `render()` Svelte's server entry, where $effect never runs, so the fetch-driven
// states (rows, failed) are Playwright's job. What a server render can prove is the initial,
// pre-hydration DOM: a Skeleton reserves the ready height while signed in, the signed-out
// prompt renders in its place otherwise, and the retired "Reload the page" copy is gone.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import MyReports from './MyReports.svelte';

describe('MyReports', () => {
  it('reserves height with a Skeleton while signed in and loading', () => {
    const { body } = render(MyReports, { props: { signedIn: true } });
    expect(body).toContain('data-testid="my-reports-skeleton"');
    expect(body).not.toContain('Reload the page');
  });

  it('renders nothing report-related when signed out', () => {
    const { body } = render(MyReports, { props: { signedIn: false } });
    expect(body).toContain('data-testid="reports-signin"');
    expect(body).not.toContain('my-reports-skeleton');
  });
});
