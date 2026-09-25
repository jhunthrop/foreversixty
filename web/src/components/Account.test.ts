// web/src/components/Account.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import Account from './Account.svelte';

describe('Account', () => {
  it('renders the nav mode placeholder before the session resolves', () => {
    const { body } = render(Account, { props: { mode: 'nav' } });
    expect(body).toContain('data-testid="session-nav"');
  });
});

describe('Account "My guilds" block', () => {
  it('renders nothing before the session resolves (loading state has no guild rows yet)', () => {
    const { body } = render(Account, { props: { mode: 'account' } });
    expect(body).not.toContain('data-testid="account-guilds"');
  });
});

describe('Account refreshed toast', () => {
  it('renders no toast on a plain server render (no window/query string to read yet)', () => {
    const { body } = render(Account, { props: { mode: 'account' } });
    expect(body).not.toContain('data-testid="account-toast"');
  });
});

describe('Account "reports" mode (logs page)', () => {
  it('renders nothing before the session resolves, not a "Loading your reports" message', () => {
    const { body } = render(Account, { props: { mode: 'reports' } });
    expect(body).not.toContain('Loading your reports');
    expect(body).not.toContain('data-testid="my-reports"');
  });
});
