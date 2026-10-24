// web/src/components/account/AccountPremiumBlock.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { EntitlementsView, MeCharacter } from '../../lib/account/api';
import { mockEntitlements } from '../../test-support/billing-fixtures';
import AccountPremiumBlock from './AccountPremiumBlock.svelte';

const CHARACTER: MeCharacter = {
  key: 'us/hardcore/elyra-duskvale',
  region: 'us',
  ruleset: 'hardcore',
  name: 'Elyra Duskvale',
  class: 'Priest',
  level: 60,
};

function text(html: string): string {
  return html
    .replace(/<!--.*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/\s+/g, ' ');
}

function renderBlock(entitlements?: EntitlementsView, characters: MeCharacter[] = []): string {
  return render(AccountPremiumBlock, { props: { entitlements, characters } }).body;
}

describe('AccountPremiumBlock', () => {
  it('has no subheading of its own: the plan line above names the plan', () => {
    expect(text(renderBlock(mockEntitlements()))).not.toContain('Premium');
  });

  it('shows the live values when the entitlements are on', () => {
    const body = text(
      renderBlock(mockEntitlements({ server_sims: true, supporter_mark: true }), [CHARACTER]),
    );
    expect(body).toContain('Server sims On');
    expect(body).toContain('Supporter mark Shown on your characters');
  });

  it('shows the free values by default', () => {
    const body = text(renderBlock(mockEntitlements()));
    expect(body).toContain('Server sims Browser lane only');
    expect(body).toContain('Supporter mark Not shown');
  });

  it('carries the month on the two benefits that have not shipped, muted and never linked', () => {
    const html = renderBlock(mockEntitlements({ server_sims: true, supporter_mark: true }), [CHARACTER]);
    expect(text(html)).toContain('History charts Ships November');
    expect(text(html)).toContain('Discord notifications Ships December');
    expect(html).toContain('text-muted">Ships November<');
    expect(html).toContain('text-muted">Ships December<');
    expect(html).not.toContain('account-premium-history-link');
    expect(html).not.toContain('account-premium-notifications-link');
  });

  it('links a live value to where the benefit lives', () => {
    const html = renderBlock(mockEntitlements({ server_sims: true, supporter_mark: true }), [CHARACTER]);
    expect(html).toMatch(/href="\/sim"[^>]*data-testid="account-premium-serverSims-link"/);
    expect(html).toMatch(
      /href="\/character\/us\/hardcore\/elyra-duskvale"[^>]*account-premium-supporter-link/,
    );
  });

  it('keeps a live supporter value plain text when the account has no character', () => {
    const html = renderBlock(mockEntitlements({ supporter_mark: true }), []);
    expect(text(html)).toContain('Shown on your characters');
    expect(html).not.toContain('account-premium-supporter-link');
  });

  it('never links a free value', () => {
    const html = renderBlock(mockEntitlements(), [CHARACTER]);
    expect(html).not.toContain('-link');
  });

  it('renders the free state for an older /v1/me with no entitlements at all', () => {
    const html = renderBlock(undefined);
    expect(html).toContain('data-testid="account-premium-block"');
    expect(text(html)).toContain('Server sims Browser lane only');
    expect(text(html)).toContain('Supporter mark Not shown');
  });
});
