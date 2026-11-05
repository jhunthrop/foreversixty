// web/src/lib/guild/landing-mine.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { Me } from '../account/api';
import GuildLanding from '../../components/GuildLanding.svelte';
import { lookupDefaults, myGuildRows } from './landing-mine';

/** An account whose current character is an alt with no guild: its guilds come from its other characters. */
const ME = {
  characters: [],
  guilds: [
    { id: 7, region: 'us', ruleset: 'pvp', name: 'Lifecycle Test', verified: true, verified_count: 3 },
    { id: 9, region: 'eu', ruleset: 'hardcore', name: 'The Last Watch', verified: false, verified_count: 0 },
  ],
} as unknown as Me;

describe('myGuildRows', () => {
  it('lists every guild of the account with where it lives', () => {
    expect(myGuildRows(ME)).toEqual([
      { key: 'us-pvp-7', name: 'Lifecycle Test', href: '/guild/us/pvp/lifecycle-test', where: 'US PvP' },
      {
        key: 'eu-hardcore-9',
        name: 'The Last Watch',
        href: '/guild/eu/hardcore/the-last-watch',
        where: 'EU Hardcore',
      },
    ]);
  });

  it('lists nothing signed out', () => {
    expect(myGuildRows(null)).toEqual([]);
  });
});

describe('lookupDefaults', () => {
  it("starts the lookup on the first guild's region and ruleset", () => {
    expect(lookupDefaults(ME)).toEqual({ region: 'us', ruleset: 'pvp' });
  });

  it('keeps the page defaults with no guild', () => {
    expect(lookupDefaults(null)).toBeNull();
  });
});

describe('the landing for a signed-in player', () => {
  it('puts Your guilds above the lookup', () => {
    const { body } = render(GuildLanding, { props: { me: ME } });
    const mine = body.indexOf('guild-landing-mine');
    expect(mine).toBeGreaterThan(-1);
    expect(mine).toBeLessThan(body.indexOf('guild-lookup-card'));
    expect(body).toContain('Your guilds');
    expect(body).toContain('href="/guild/us/pvp/lifecycle-test"');
    expect(body).toContain('EU Hardcore');
  });

  it('draws no Your guilds card signed out', () => {
    const { body } = render(GuildLanding, { props: {} });
    expect(body).not.toContain('guild-landing-mine');
  });
});
