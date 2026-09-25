import { describe, expect, it } from 'vitest';
import { PRIMARY_NAV_ITEMS, SETUP_NAV_ITEM, TRAILING_NAV_ITEMS, isNavItemCurrent } from './nav';

describe('nav structure', () => {
  it('orders the five doors Planner, Simulator, Logs, Rankings, Guides', () => {
    expect(PRIMARY_NAV_ITEMS.map((item) => item.label)).toEqual([
      'Planner',
      'Simulator',
      'Logs',
      'Rankings',
      'Guides',
    ]);
    expect(PRIMARY_NAV_ITEMS.map((item) => item.href)).toEqual([
      '/planner',
      '/sim',
      '/logs',
      '/rankings',
      '/guides',
    ]);
  });

  it('names the setup nav item "Get set up" at /setup, alone in the trailing row', () => {
    expect(SETUP_NAV_ITEM).toEqual({ label: 'Get set up', href: '/setup' });
    expect(TRAILING_NAV_ITEMS).toEqual([SETUP_NAV_ITEM]);
  });
});

describe('isNavItemCurrent', () => {
  const planner = PRIMARY_NAV_ITEMS[0];

  it('matches the exact path', () => {
    expect(isNavItemCurrent(planner, '/planner')).toBe(true);
  });

  it('matches a sub-path', () => {
    const simulator = PRIMARY_NAV_ITEMS[1];
    expect(isNavItemCurrent(simulator, '/sim/gear')).toBe(true);
  });

  it('does not match an unrelated path', () => {
    expect(isNavItemCurrent(planner, '/sim')).toBe(false);
  });

  it('does not match a path that merely starts with the same letters', () => {
    expect(isNavItemCurrent(planner, '/plannerx')).toBe(false);
  });
});
