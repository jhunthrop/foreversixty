// web/src/components/SessionNav.test.ts
// Static-render check, same pattern as SubstitutionChips.test.ts: `svelte/server`'s
// `render` compiles real markup into a string without a browser. Svelte 5's `$effect`
// never runs during this render, so the output is always the component's initial state --
// exactly the "session not resolved yet" placeholder every page paints before hydration,
// which is the state this test needs to pin (it is what reserves the header's width).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import SessionNav from './SessionNav.svelte';

describe('SessionNav', () => {
  it('renders the session-nav testid every page-level test and e2e spec queries', () => {
    const { body } = render(SessionNav);
    expect(body).toContain('data-testid="session-nav"');
  });

  it('reserves its width before the session check resolves: an invisible "Sign in", no href', () => {
    const { body } = render(SessionNav);
    expect(body).toContain('Sign in');
    expect(body).not.toContain('href="/login"');
    expect(body).not.toContain('href="/account"');
  });
});
