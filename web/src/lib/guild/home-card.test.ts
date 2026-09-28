// web/src/lib/guild/home-card.test.ts
import { describe, expect, it } from 'vitest';
import type { Me } from '../account/api';
import type { GuildHome, GuildRosterRow } from './api';
import { homeGuildCardView } from './home-card';

function me(guilds: Me['guilds']): Me {
  return {
    user: { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false },
    characters: [],
    guilds,
  };
}

function row(overrides: Partial<GuildRosterRow> = {}): GuildRosterRow {
  return {
    character_key: 'us/normal/alt',
    region: 'us',
    ruleset: 'normal',
    name: 'Alt',
    rank: 'member',
    verified: true,
    logged_recently: false,
    consent: 'roster',
    may_remove: false,
    ...overrides,
  };
}

function home(overrides: Partial<GuildHome> = {}): GuildHome {
  return {
    guild: { id: 501, region: 'us', ruleset: 'hardcore', name: 'The Last Watch' },
    claim: { state: 'claimed', frozen: false },
    reports: [],
    roster: [],
    ...overrides,
  };
}

const MEMBER_GUILD = {
  id: 501,
  region: 'us',
  ruleset: 'hardcore',
  name: 'The Last Watch',
  rank: 'member',
  verified: true,
};

const OFFICER_GUILD = { ...MEMBER_GUILD, rank: 'officer' };
const LEADER_GUILD = { ...MEMBER_GUILD, rank: 'leader' };

describe('homeGuildCardView', () => {
  it('is the no-guild state when the account has no guilds, pointing at Get set up', () => {
    const view = homeGuildCardView(me([]), null, null);
    expect(view.state).toBe('no-guild');
    expect(view.line).toBe(
      'Ask an officer of your guild for its invite link, or save a character that is in one on Get set up.',
    );
    expect(view.action).toEqual({ label: 'Get set up', href: '/setup' });
  });

  it('is the no-guild state when there is no account at all', () => {
    expect(homeGuildCardView(null, null, null).state).toBe('no-guild');
  });

  it('is home-failed when the guild home fetch failed or was refused, still pointing at the guild', () => {
    const view = homeGuildCardView(me([MEMBER_GUILD]), null, null);
    expect(view.state).toBe('home-failed');
    expect(view.line).toBe('');
    expect(view.action).toEqual({ label: 'View guild', href: '/guild/us/hardcore/the-last-watch' });
  });

  it('is claim-pending for an officer of a guild with a pending claim', () => {
    const view = homeGuildCardView(
      me([OFFICER_GUILD]),
      home({ claim: { state: 'pending', frozen: false } }),
      null,
    );
    expect(view.state).toBe('claim-pending');
    expect(view.line).toBe('A claim is waiting for a second officer.');
    expect(view.action).toEqual({
      label: 'Confirm the claim',
      href: '/guild/us/hardcore/the-last-watch/claim',
    });
  });

  it('is claim-pending for a leader too', () => {
    const view = homeGuildCardView(
      me([LEADER_GUILD]),
      home({ claim: { state: 'pending', frozen: false } }),
      null,
    );
    expect(view.state).toBe('claim-pending');
  });

  it('is claim-unclaimed for an officer of an unclaimed guild, naming the guild', () => {
    const view = homeGuildCardView(
      me([OFFICER_GUILD]),
      home({ claim: { state: 'unclaimed', frozen: false } }),
      null,
    );
    expect(view.state).toBe('claim-unclaimed');
    expect(view.line).toBe(
      'Nobody has claimed The Last Watch yet. Claiming unlocks settings, the invite link and roster approval.',
    );
    expect(view.action).toEqual({
      label: 'Claim this guild',
      href: '/guild/us/hardcore/the-last-watch/claim',
    });
  });

  it('is waiting-approval for an officer with unverified roster rows once the claim is settled', () => {
    const view = homeGuildCardView(
      me([OFFICER_GUILD]),
      home({ roster: [row({ verified: false }), row({ verified: false }), row({ verified: true })] }),
      null,
    );
    expect(view.state).toBe('waiting-approval');
    expect(view.line).toBe('2 members waiting for approval.');
    expect(view.action).toEqual({ label: 'Review roster', href: '/guild/us/hardcore/the-last-watch' });
  });

  it('checks claim-pending before waiting-approval', () => {
    const view = homeGuildCardView(
      me([OFFICER_GUILD]),
      home({ claim: { state: 'pending', frozen: false }, roster: [row({ verified: false })] }),
      null,
    );
    expect(view.state).toBe('claim-pending');
  });

  it('checks claim-unclaimed before waiting-approval', () => {
    const view = homeGuildCardView(
      me([OFFICER_GUILD]),
      home({ claim: { state: 'unclaimed', frozen: false }, roster: [row({ verified: false })] }),
      null,
    );
    expect(view.state).toBe('claim-unclaimed');
  });

  it('is steady for an officer once claimed with nobody waiting, showing the stats line and Guild settings', () => {
    const view = homeGuildCardView(
      me([OFFICER_GUILD]),
      home({
        reports: [
          {
            id: 'r1',
            title: 'Night one',
            zone: 'Onyxia',
            created_at: '2026-09-20',
            fight_count: 5,
            kill_count: 3,
          },
        ],
        roster: [
          row({ logged_recently: true }),
          row({ logged_recently: true }),
          row({ logged_recently: false }),
        ],
      }),
      { killed: 3, total: 8 },
    );
    expect(view.state).toBe('steady');
    expect(view.line).toBe('2 logged in the last day · 1 reports this week · 3/8 bosses');
    expect(view.action).toEqual({
      label: 'Guild settings',
      href: '/guild/us/hardcore/the-last-watch/settings',
    });
  });

  it('is steady for a plain member, showing View guild instead of Guild settings', () => {
    const view = homeGuildCardView(me([MEMBER_GUILD]), home({ roster: [row({ logged_recently: true })] }), {
      killed: 1,
      total: 8,
    });
    expect(view.state).toBe('steady');
    expect(view.action).toEqual({ label: 'View guild', href: '/guild/us/hardcore/the-last-watch' });
  });

  it('never shows claim state to a plain member -- an unclaimed or pending claim still lands on steady', () => {
    const pending = homeGuildCardView(
      me([MEMBER_GUILD]),
      home({ claim: { state: 'pending', frozen: false } }),
      null,
    );
    expect(pending.state).toBe('steady');
    const unclaimed = homeGuildCardView(
      me([MEMBER_GUILD]),
      home({ claim: { state: 'unclaimed', frozen: false } }),
      null,
    );
    expect(unclaimed.state).toBe('steady');
  });

  it('never shows a waiting-approval count to a plain member', () => {
    const view = homeGuildCardView(me([MEMBER_GUILD]), home({ roster: [row({ verified: false })] }), null);
    expect(view.state).toBe('steady');
  });

  it('omits the bosses figure from the steady line when progression failed to load', () => {
    const view = homeGuildCardView(me([MEMBER_GUILD]), home({ reports: [] }), null);
    expect(view.line).toBe('0 logged in the last day · 0 reports this week');
  });
});
