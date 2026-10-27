// web/src/components/checkout/CheckoutGuildChooser.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { MeGuild } from '../../lib/account/api';
import CheckoutGuildChooser from './CheckoutGuildChooser.svelte';

const base = { region: 'us', ruleset: 'x', verified: true };
const GUILDS: MeGuild[] = [
  { ...base, id: 1, name: 'Olympus', verified_count: 12, can: ['billing'], plan: null },
  { ...base, id: 2, name: 'Titans', verified_count: 28, can: ['billing', 'settings'] },
  {
    ...base,
    id: 3,
    name: 'Planned',
    verified_count: 5,
    can: ['billing'],
    plan: { tier: 'guild' } as MeGuild['plan'],
  },
  { ...base, id: 4, name: 'Member', verified_count: 5, can: [] },
];

function html(props: {
  guilds: MeGuild[] | null;
  failed?: boolean;
  crests?: Record<number, string | null> | null;
  summary?: string;
  plan?: 'guild' | 'raid_guild' | null;
}): string {
  return render(CheckoutGuildChooser, {
    props: { onChoose: () => {}, onRetry: () => {}, ...props },
  }).body.replace(/<!--.*?-->/g, '');
}

describe('CheckoutGuildChooser', () => {
  it('lists only billable guilds, as two 56px rows', () => {
    const body = html({ guilds: GUILDS });
    expect(body.match(/data-testid="checkout-guild-row"/g)).toHaveLength(2);
    expect(body.match(/\bh-14\b/g)).toHaveLength(2);
    expect(body).toContain('Which guild?');
    expect(body).toContain('12 verified members');
    expect(body).not.toContain('Planned');
  });

  it('names the plan that fits: Guild up to 25, Raid guild above', () => {
    const body = html({ guilds: GUILDS });
    expect(body).toContain('Raid guild fits 28 members');
    expect(body).toContain('>Guild<');
  });

  it('with a picked plan, a row only flags a guild that plan does not fit', () => {
    const raid = html({ guilds: GUILDS, plan: 'raid_guild' });
    expect(raid).not.toContain('>Guild<');
    expect(raid).not.toContain('Raid guild fits');
    const guild = html({ guilds: GUILDS, plan: 'guild' });
    expect(guild).toContain('Raid guild fits 28 members');
    expect(guild).not.toContain('>Guild<');
  });

  it('says so, with the Find your guild link, when nothing is billable', () => {
    const body = html({ guilds: [GUILDS[3]!] });
    expect(body).toContain('You can subscribe a guild you lead or have claimed.');
    expect(body).toContain('href="/guild"');
    expect(body).toContain('Find your guild');
    expect(body).not.toContain('checkout-guild-row');
  });

  it('shows the load error with Try again', () => {
    const body = html({ guilds: null, failed: true });
    expect(body).toContain('Your guilds did not load.');
    expect(body).toContain('Try again');
  });

  it('shows three skeleton rows while loading', () => {
    const body = html({ guilds: null });
    expect(body).toContain('checkout-chooser-loading');
    expect(body.match(/skeleton-block/g)).toHaveLength(3);
  });

  it('keeps the rows in the Skeleton until the crests settle, so each crest arrives with its row', () => {
    const body = html({ guilds: GUILDS, crests: null });
    expect(body).toContain('checkout-chooser-loading');
    expect(body).not.toContain('checkout-guild-row');
  });

  it('names the plan and price picked on /premium under the heading', () => {
    const body = html({ guilds: GUILDS, summary: 'Raid guild · $200 a year' });
    expect(body).toContain('data-testid="checkout-chooser-plan">Raid guild · $200 a year<');
  });

  it('draws a known crest inside the ring, and the heading is focusable', () => {
    const body = html({ guilds: GUILDS, crests: { 1: '/crest/1.png' } });
    expect(body.match(/<img /g)).toHaveLength(1);
    expect(body).toContain('src="/crest/1.png"');
    expect(body).toContain('tabindex="-1"');
  });
});
