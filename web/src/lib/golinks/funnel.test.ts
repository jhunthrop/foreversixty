import { describe, expect, it } from 'vitest';
import type { GoLinkFunnel } from './api';
import { countLabel, funnelTotals, shareOf, shortLink, sortedByClicks } from './funnel';

const row = (slug: string, clicks: number, sign_ins = 0, linked = 0, premium = 0): GoLinkFunnel => ({
  slug,
  target: '/',
  label: slug,
  created_at: '2026-11-02T00:00:00Z',
  clicks,
  bot_clicks: 1,
  sign_ins,
  linked,
  premium,
});

describe('shortLink', () => {
  it('is the link people paste', () => {
    expect(shortLink('reddit')).toBe('https://foreversixty.gg/go/reddit');
  });
});

describe('shareOf', () => {
  it('is a whole percent of the step before, or a dash with nothing to divide', () => {
    expect(shareOf(3, 12)).toBe('25%');
    expect(shareOf(1, 3)).toBe('33%');
    expect(shareOf(0, 0)).toBe('–');
    expect(shareOf(0, 40)).toBe('0%');
  });
});

describe('countLabel', () => {
  it('groups thousands', () => {
    expect(countLabel(6132)).toBe('6,132');
    expect(countLabel(0)).toBe('0');
  });
});

describe('funnelTotals', () => {
  it('adds every link up', () => {
    expect(funnelTotals([row('a', 10, 2, 1, 1), row('b', 5, 1, 1, 0)])).toEqual({
      clicks: 15,
      bot_clicks: 2,
      sign_ins: 3,
      linked: 2,
      premium: 1,
    });
  });
});

describe('sortedByClicks', () => {
  it('puts the busiest link first, ties by slug, without touching the input', () => {
    const input = [row('b', 1), row('a', 1), row('c', 9)];
    expect(sortedByClicks(input).map((r) => r.slug)).toEqual(['c', 'a', 'b']);
    expect(input.map((r) => r.slug)).toEqual(['b', 'a', 'c']);
  });
});
