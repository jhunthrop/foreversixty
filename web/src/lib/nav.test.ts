import { describe, expect, it } from 'vitest';
import {
  GUILD_GROUP_INDEX,
  NAV_GROUPS,
  SETUP_NAV_ITEM,
  isNavGroupCurrent,
  isNavItemCurrent,
} from './nav';

const [tools, logs, guild, guides] = NAV_GROUPS;
const rowPairs = (group: (typeof NAV_GROUPS)[number]) =>
  group.rows.map((row) => [row.label, row.href]);

describe('nav groups', () => {
  it('orders the four doors Tools, Logs, Guild, Guides', () => {
    expect(NAV_GROUPS.map((group) => [group.label, group.href])).toEqual([
      ['Tools', '/planner'],
      ['Logs', '/logs'],
      ['Guild', '/guild'],
      ['Guides', '/guides'],
    ]);
    expect(NAV_GROUPS[GUILD_GROUP_INDEX]).toBe(guild);
  });

  it('lists the rows of each panel in order', () => {
    expect(rowPairs(tools)).toEqual([
      ['Planner', '/planner'],
      ['BiS', '/bis'],
      ['Simulator', '/sim'],
      ['Get set up', '/setup'],
    ]);
    expect(rowPairs(logs)).toEqual([
      ['Upload a log', '/logs'],
      ['Your reports', '/reports'],
      ['Rankings', '/rankings'],
      ['Tier list', '/tiers'],
    ]);
    expect(rowPairs(guild)).toEqual([
      ['Your guild', '/guild'],
      ['Roster', '/guild#roster'],
      ['Loot council', '/guild#loot'],
      ['Guild logs', '/guild#raids'],
    ]);
    expect(guides.rows).toEqual([]);
  });

  it('gives every row a description and the right foots', () => {
    for (const group of NAV_GROUPS) {
      for (const row of group.rows) expect(row.description.length).toBeGreaterThan(0);
    }
    expect(guild.foot).toEqual({ label: 'Guild settings', href: '/guild#settings' });
    expect(guides.foot).toEqual({ label: 'All guides', href: '/guides' });
    expect(tools.foot).toBeUndefined();
    expect(logs.foot).toBeUndefined();
  });

  it('names the setup nav item "Get set up" at /setup', () => {
    expect(SETUP_NAV_ITEM).toEqual({ label: 'Get set up', href: '/setup' });
  });
});

describe('isNavGroupCurrent', () => {
  it('marks a group when the page is one of its rows or a sub-path', () => {
    expect(isNavGroupCurrent(logs, '/tiers/healer')).toBe(true);
    expect(isNavGroupCurrent(tools, '/sim/gear')).toBe(true);
    expect(isNavGroupCurrent(guild, '/guild/us/pvp/olympus-xxvii')).toBe(true);
    expect(isNavGroupCurrent(guides, '/guides/warrior')).toBe(true);
  });

  it('does not mark an unrelated group', () => {
    expect(isNavGroupCurrent(tools, '/logs')).toBe(false);
    expect(isNavGroupCurrent(logs, '/plannerx')).toBe(false);
  });
});

describe('isNavItemCurrent', () => {
  const planner = tools.rows[0];

  it('matches the exact path and sub-paths only', () => {
    expect(isNavItemCurrent(planner, '/planner')).toBe(true);
    expect(isNavItemCurrent(planner, '/planner/x')).toBe(true);
    expect(isNavItemCurrent(planner, '/sim')).toBe(false);
    expect(isNavItemCurrent(planner, '/plannerx')).toBe(false);
  });
});
