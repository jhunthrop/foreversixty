// web/src/components/account/AccountMenu.test.ts
// Static-render check, same pattern as SessionNav.test.ts (the component this replaces):
// `svelte/server`'s `render` compiles real markup into a string without a browser. Svelte
// 5's `$effect` never runs during this render, so the output is always the component's
// initial state -- exactly the "session not resolved yet" placeholder every page paints
// before hydration, which is the state this test needs to pin (it is what reserves the
// header's width).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import AccountMenu from './AccountMenu.svelte';

describe('AccountMenu', () => {
  it('renders the invisible loading placeholder before the session resolves', () => {
    const { body } = render(AccountMenu);
    expect(body).toContain('data-testid="session-nav"');
    expect(body).toContain('invisible');
  });

  it('server-renders no menu content before the session answers (avoids a signed-out flash)', () => {
    const { body } = render(AccountMenu);
    expect(body).not.toContain('data-testid="account-menu"');
    expect(body).not.toContain('href="/login"');
  });
});
