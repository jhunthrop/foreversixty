// web/src/lib/guild/landing-mine.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { Me } from '../account/api';
import GuildLanding from '../../components/GuildLanding.svelte';
import { factionLogoSrc } from '../faction-mark';
import { heldClaimRow, lookupDefaults, myGuildRows } from './landing-mine';

/** An account whose current character is an alt with no guild: its guilds come from its other characters. */
const ME = {
  characters: [],
  guilds: [
    {
      id: 7,
      region: 'us',
      ruleset: 'pvp',
      name: 'Lifecycle Test',
      rank: 'officer',
      verified: true,
      verified_count: 3,
      faction: 'alliance',
      roster_count: 24,
    },
    { id: 9, region: 'eu', ruleset: 'hardcore', name: 'The Last Watch', verified: false, verified_count: 0 },
  ],
} as unknown as Me;

describe('myGuildRows', () => {
  it('lists every guild of the account with its mark, where it lives and its roster', () => {
    expect(myGuildRows(ME)).toEqual([
      {
        key: 'us-pvp-7',
        name: 'Lifecycle Test',
        href: '/guild/us/pvp/lifecycle-test',
        where: 'US PvP · 24 members',
        mark: factionLogoSrc('alliance'),
        rank: 'Officer',
      },
      {
        key: 'eu-hardcore-9',
        name: 'The Last Watch',
        href: '/guild/eu/hardcore/the-last-watch',
        where: 'EU Hardcore',
        mark: null,
        rank: null,
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
  it('leads the rail with Your guilds, above Run a guild?', () => {
    const { body } = render(GuildLanding, { props: { me: ME } });
    const mine = body.indexOf('guild-landing-mine');
    expect(mine).toBeGreaterThan(body.indexOf('ui-page-rail'));
    expect(mine).toBeLessThan(body.indexOf('guild-landing-claim'));
    expect(body).toContain('>Officer<');
    expect(body).toContain('Your guilds');
    expect(body).toContain('href="/guild/us/pvp/lifecycle-test"');
    expect(body).toContain('EU Hardcore');
  });

  it('draws no Your guilds section signed out', () => {
    const { body } = render(GuildLanding, { props: {} });
    expect(body).not.toContain('guild-landing-mine');
  });
});

describe('the foot line for a claim holder', () => {
  const HOLDER = {
    ...ME,
    guilds: [
      ME.guilds[0],
      { ...ME.guilds[1], claim: { state: 'claimed', yours: true, pending_yours: false, holder_name: null } },
    ],
  } as unknown as Me;

  it('finds the guild whose claim the account holds', () => {
    expect(heldClaimRow(HOLDER)?.name).toBe('The Last Watch');
    expect(heldClaimRow(ME)).toBeNull();
    expect(heldClaimRow(null)).toBeNull();
  });

  it('says the claim is held instead of asking Run a guild?', () => {
    const { body } = render(GuildLanding, { props: { me: HOLDER } });
    expect(body).toContain("You hold The Last Watch's claim.");
    expect(body).not.toContain('Run a guild?');
    expect(render(GuildLanding, { props: { me: ME } }).body).toContain('Run a guild?');
  });
});
