// web/src/components/guild/GuildBillingSection.test.ts
// The orchestrator's server render is its loading state: the read happens on the client, so the
// states themselves are pinned in billing/*.test.ts and lib/guild/billing-state.test.ts.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { GUILD_BILLING_LOADING, guildBillingReserveKey } from '../../lib/guild/layout';
import GuildBillingSection from './GuildBillingSection.svelte';

describe('GuildBillingSection', () => {
  it('reserves its height with Skeleton lines until the billing read answers', () => {
    const body = render(GuildBillingSection, {
      props: {
        guildId: 7,
        path: { region: 'us', ruleset: 'pvp', slug: 'olympus-xxvii' },
        guildName: 'OLYMPUS XXVII',
      },
    }).body;
    expect(body).toContain('id="billing"');
    expect(body).toContain('guild-billing-skeleton');
    expect(body).toContain(GUILD_BILLING_LOADING.minHeight);
    expect(GUILD_BILLING_LOADING.minHeight).toBe('min-h-[760px] md:min-h-[520px]');
    expect(body).not.toContain('guild-records');
  });

  it('keys the stored ready height per guild and per layout', () => {
    expect(guildBillingReserveKey(7, true)).toBe('fs.guildBilling.reserve.7.phone');
    expect(guildBillingReserveKey(7, false)).toBe('fs.guildBilling.reserve.7.wide');
  });
});
