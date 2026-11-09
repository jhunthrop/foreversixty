import { describe, expect, it } from 'vitest';
import communityLinks from '../data/links.json';
import {
  GUILD_GROUP_INDEX,
  NAV_GROUPS,
  SETUP_NAV_ITEM,
  isNavGroupCurrent,
  isNavItemCurrent,
  navKey,
} from './nav';

const [bis, guides, tiers, sim, planner, logs, guild] = NAV_GROUPS;
const rowPairs = (group: (typeof NAV_GROUPS)[number]) => group.rows.map((row) => [row.label, row.href]);

describe('nav groups', () => {
  it("orders the mock's doors: Best in slot, Guides, Tier list, Sim, Planner, Logs, Guild", () => {
    expect(NAV_GROUPS.map((group) => [group.label, group.href])).toEqual([
      ['Best in slot', '/bis'],
      ['Guides', '/guides'],
      ['Tier list', '/tiers'],
      ['Sim', '/sim'],
      ['Planner', '/planner'],
      ['Logs', '/logs'],
      ['Guild', '/guild'],
    ]);
    expect(NAV_GROUPS[GUILD_GROUP_INDEX]).toBe(guild);
  });

  it('gives a panel only to the doors with sub-pages', () => {
    expect(NAV_GROUPS.filter((group) => group.panel).map((group) => group.label)).toEqual([
      'Guides',
      'Logs',
      'Guild',
    ]);
    for (const plain of [bis, tiers, sim, planner]) {
      expect(plain.rows).toEqual([]);
      expect(plain.foot).toBeUndefined();
    }
  });

  it('lists the rows of each panel in order', () => {
    expect(rowPairs(logs)).toEqual([
      ['Upload a log', '/logs'],
      ['Your reports', '/reports'],
      ['Rankings', '/rankings'],
    ]);
    expect(rowPairs(guild)).toEqual([
      ['Your guild', '/guild'],
      ['Roster', '/guild#roster'],
      ['Loot council', '/guild#loot'],
      ['Guild logs', '/guild#raids'],
      ['Forever Sixty Discord', communityLinks.discordInvite],
    ]);
    expect(guides.rows).toEqual([]);
  });

  it('gives every row a description and the right foots', () => {
    for (const group of NAV_GROUPS) {
      for (const row of group.rows) expect(row.description.length).toBeGreaterThan(0);
    }
    expect(guild.foot).toEqual({ label: 'Guild settings', href: '/guild#settings' });
    expect(guides.foot).toEqual({ label: 'All guides', href: '/guides' });
    expect(logs.foot).toBeUndefined();
  });

  it('ends the Guild panel with the Discord row, the one community link in the header', () => {
    const last = guild.rows.at(-1)!;
    expect(last.description).toBe('Find a guild, ask about the site, report a bug.');
    expect(last.mark).toBe('discord');
    expect(NAV_GROUPS.flatMap((group) => group.rows).filter((row) => row.mark === 'discord')).toHaveLength(1);
  });

  it('names the setup nav item "Get set up" at /setup', () => {
    expect(SETUP_NAV_ITEM).toEqual({ label: 'Get set up', href: '/setup' });
  });
});

describe('isNavGroupCurrent', () => {
  it('marks a group when the page is one of its rows or a sub-path', () => {
    expect(isNavGroupCurrent(tiers, '/tiers/healer')).toBe(true);
    expect(isNavGroupCurrent(logs, '/tiers/healer')).toBe(false);
    expect(isNavGroupCurrent(sim, '/sim/gear')).toBe(true);
    expect(isNavGroupCurrent(logs, '/rankings')).toBe(true);
    expect(isNavGroupCurrent(guild, '/guild/us/pvp/olympus-xxvii')).toBe(true);
    expect(isNavGroupCurrent(guides, '/guides/warrior')).toBe(true);
  });

  it('does not mark an unrelated group', () => {
    expect(isNavGroupCurrent(bis, '/logs')).toBe(false);
    expect(isNavGroupCurrent(planner, '/plannerx')).toBe(false);
  });
});

describe('navKey', () => {
  it('turns a label into a lowercase hyphenated key', () => {
    expect(navKey('Upload a log')).toBe('upload-a-log');
    expect(navKey('BiS')).toBe('bis');
  });
});

describe('isNavItemCurrent', () => {
  it('matches the exact path and sub-paths only', () => {
    expect(isNavItemCurrent(planner, '/planner')).toBe(true);
    expect(isNavItemCurrent(planner, '/planner/x')).toBe(true);
    expect(isNavItemCurrent(planner, '/sim')).toBe(false);
    expect(isNavItemCurrent(planner, '/plannerx')).toBe(false);
  });
});
