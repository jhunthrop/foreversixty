import { JSDOM } from 'jsdom';
import { describe, expect, it } from 'vitest';
import type { PremiumSource } from './billing/contract';
import { applySessionPremium, nextPremiumNavState, premiumNavShown } from './premium-nav';

// Spec 2026-10-29-premium-entry P2 and section 4.
const SOURCES: [PremiumSource | null, boolean][] = [
  [null, true], // signed out
  ['none', true], // free, and an officer whose guild pays but who holds no seat
  ['own', false],
  ['own_then_seat', false],
  ['grant', false],
  ['seat', false],
  ['fallback', false],
];

const rootWith = (state?: string): HTMLElement => {
  const root = new JSDOM('<!doctype html><html></html>').window.document.documentElement;
  if (state !== undefined) root.dataset.premiumNav = state;
  return root;
};

describe('premiumNavShown', () => {
  it.each(SOURCES)('%s shows the tag: %s', (source, shown) => {
    expect(premiumNavShown(source)).toBe(shown);
  });
});

describe('nextPremiumNavState', () => {
  it.each(SOURCES)('a showing tag, %s: held only for Premium', (source, shown) => {
    expect(nextPremiumNavState(null, source)).toBe(shown ? null : 'hold');
  });

  it.each(SOURCES)('a tag hidden before paint stays hidden, %s', (source) => {
    expect(nextPremiumNavState('hide', source)).toBe('hide');
  });

  it('a held tag stays held for the page view', () => {
    expect(nextPremiumNavState('hold', 'none')).toBe('hold');
  });
});

describe('applySessionPremium', () => {
  it('holds the box for a Premium account the hint missed', () => {
    const root = rootWith();
    applySessionPremium('grant', root);
    expect(root.dataset.premiumNav).toBe('hold');
  });

  it('leaves a free or signed-out tag alone', () => {
    for (const source of ['none', null] as const) {
      const root = rootWith();
      applySessionPremium(source, root);
      expect(root.hasAttribute('data-premium-nav')).toBe(false);
    }
  });

  it('never brings back a tag the pre-paint hint hid', () => {
    const root = rootWith('hide');
    applySessionPremium('none', root);
    expect(root.dataset.premiumNav).toBe('hide');
  });
});
