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

function html(props: { guilds: MeGuild[] | null; failed?: boolean }): string {
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
});
