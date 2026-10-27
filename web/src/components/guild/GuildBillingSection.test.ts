// web/src/components/guild/GuildBillingSection.test.ts
// The orchestrator's server render is its loading state: the read happens on the client, so the
// states themselves are pinned in billing/*.test.ts and lib/guild/billing-state.test.ts.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import GuildBillingSection from './GuildBillingSection.svelte';

describe('GuildBillingSection', () => {
  it('reserves its height with Skeleton lines until the billing read answers', () => {
    const body = render(GuildBillingSection, { props: { guildId: 7 } }).body;
    expect(body).toContain('id="billing"');
    expect(body).toContain('guild-billing-skeleton');
    expect(body).toContain('min-h-[640px] md:min-h-[420px]');
    expect(body).not.toContain('guild-records');
  });
});
