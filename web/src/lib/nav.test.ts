import { describe, expect, it } from 'vitest';
import {
  ADDON_NAV_ITEM,
  PRIMARY_NAV_ITEMS,
  REFERENCE_NAV_ITEMS,
  isNavItemCurrent,
  isReferenceCurrent,
} from './nav';

describe('nav structure', () => {
  it('orders the four tools Planner, Simulator, Logs, Rankings', () => {
    expect(PRIMARY_NAV_ITEMS.map((item) => item.label)).toEqual(['Planner', 'Simulator', 'Logs', 'Rankings']);
    expect(PRIMARY_NAV_ITEMS.map((item) => item.href)).toEqual(['/planner', '/sim', '/logs', '/rankings']);
  });

  it('orders the reference group Classes, Guides, Zones, Dungeons', () => {
    expect(REFERENCE_NAV_ITEMS.map((item) => item.label)).toEqual(['Classes', 'Guides', 'Zones', 'Dungeons']);
    expect(REFERENCE_NAV_ITEMS.map((item) => item.href)).toEqual([
      '/classes',
      '/guides',
      '/zones',
      '/dungeons',
    ]);
  });

  it('names the addon nav item "The addon" at /addon', () => {
    expect(ADDON_NAV_ITEM).toEqual({ label: 'The addon', href: '/addon' });
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
    // /plannerx is not a sub-path of /planner; only an exact "/planner" or "/planner/..." is.
    expect(isNavItemCurrent(planner, '/plannerx')).toBe(false);
  });
});

describe('isReferenceCurrent', () => {
  it('is true when the path is one of the reference pages', () => {
    expect(isReferenceCurrent('/classes')).toBe(true);
    expect(isReferenceCurrent('/dungeons/hall-of-thanes')).toBe(true);
  });

  it('is false elsewhere', () => {
    expect(isReferenceCurrent('/planner')).toBe(false);
    expect(isReferenceCurrent('/')).toBe(false);
  });
});
