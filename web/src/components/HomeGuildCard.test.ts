// web/src/components/HomeGuildCard.test.ts
// svelte/server's render only ever sees pre-$effect state, so these pin what the card
// shows the instant HomeAccountPanel mounts it with the account: the height-reserving
// skeleton while a guild's data is still to come, and the finished no-guild card at once
// when there is nothing to fetch.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { Me } from '../lib/account/api';
import HomeGuildCard from './HomeGuildCard.svelte';

const base = { user: { battletag: 'Ash#1' }, characters: [] } as unknown as Me;

describe('HomeGuildCard first paint', () => {
  it('reserves the card height with a Skeleton while a guild is still loading', () => {
    const me = {
      ...base,
      guilds: [
        { id: 5, region: 'us', ruleset: 'pvp', name: 'The Last Watch', rank: 'member', verified: true },
      ],
    } as Me;
    const { body } = render(HomeGuildCard, { props: { me } });
    expect(body).toContain('data-testid="home-guild-card-skeleton"');
    expect(body).toContain('min-h-[168px]');
    expect(body).not.toContain('data-testid="home-guild-card"');
  });

  it('draws the no-guild card at once when the account has no guild to fetch', () => {
    const { body } = render(HomeGuildCard, { props: { me: { ...base, guilds: [] } as Me } });
    expect(body).toContain('data-testid="home-guild-card"');
    expect(body).toContain('No guild yet.');
    expect(body).not.toContain('data-testid="home-guild-card-skeleton"');
  });
});
