import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { createRawSnippet } from 'svelte';
import AccountPanel from './AccountPanel.svelte';

describe('AccountPanel', () => {
  it('wraps its children in the bordered, raised box and carries a testid', () => {
    const children = createRawSnippet(() => ({
      render: () => '<p data-testid="inner">Inside</p>',
    }));
    const { body } = render(AccountPanel, { props: { testid: 'account-more', children } });
    expect(body).toContain('data-testid="account-more"');
    expect(body).toContain('bg-raised');
    expect(body).toContain('data-testid="inner"');
  });
});
