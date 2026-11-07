// web/src/components/Account.test.ts
// The server render of both modes: /account's first paint is its loading state (the ready page's
// shape, nothing from the account yet), /login's the Battle.net button.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import Account from './Account.svelte';

describe('Account', () => {
  it('server-renders /account as the loading state, its columns carrying the reserve', () => {
    const { body } = render(Account, { props: { mode: 'account' } });
    expect(body).toContain('data-testid="account"');
    expect(body).toContain('data-testid="account-left-skeleton"');
    expect(body).toContain('data-testid="account-rail-skeleton"');
    expect(body).toContain('Your account');
    expect(body).not.toContain('data-testid="account-guilds"');
    expect(body).not.toContain('data-testid="account-signed-in-banner"');
  });

  it('server-renders /login with the Battle.net button', () => {
    const { body } = render(Account, { props: { mode: 'login' } });
    expect(body).toContain('data-testid="battlenet"');
  });
});
