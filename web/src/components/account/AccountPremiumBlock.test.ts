// web/src/components/account/AccountPremiumBlock.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { mockEntitlements } from '../../test-support/billing-fixtures';
import AccountPremiumBlock from './AccountPremiumBlock.svelte';

function text(html: string): string {
  return html
    .replace(/<!--.*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/\s+/g, ' ');
}

function blockText(entitlements?: ReturnType<typeof mockEntitlements>): string {
  return text(render(AccountPremiumBlock, { props: { entitlements } }).body);
}

describe('AccountPremiumBlock', () => {
  it('shows the live states when the entitlements are on', () => {
    const body = blockText(mockEntitlements({ server_sims: true, supporter_mark: true }));
    expect(body).toContain('Premium');
    expect(body).toContain('Server sims · On');
    expect(body).toContain('Supporter mark · Shown on your characters');
  });

  it('shows the free states by default', () => {
    const body = blockText(mockEntitlements());
    expect(body).toContain('Server sims · Browser lane only');
    expect(body).toContain('Supporter mark · Not shown');
  });

  it('carries the month on the two benefits that have not shipped', () => {
    const body = blockText(mockEntitlements({ server_sims: true }));
    expect(body).toContain('History charts · Ships November');
    expect(body).toContain('Discord notifications · Ships December');
  });

  it('renders the free state for an older /v1/me with no entitlements at all', () => {
    const html = render(AccountPremiumBlock, { props: { entitlements: undefined } }).body;
    expect(html).toContain('data-testid="account-premium-block"');
    expect(text(html)).toContain('Server sims · Browser lane only');
    expect(text(html)).toContain('Supporter mark · Not shown');
  });
});
