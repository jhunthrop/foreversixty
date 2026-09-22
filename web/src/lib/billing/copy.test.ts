// web/src/lib/billing/copy.test.ts
import { describe, expect, it } from 'vitest';
import { premiumCopy, checkoutCopy, billingBlockCopy } from './copy';

describe('premiumCopy', () => {
  it('builds a priced plan name from the one price data file, never a hardcoded literal', () => {
    expect(premiumCopy.pricedName('premium', 'monthly')).toBe('Premium — $4/month');
    expect(premiumCopy.pricedName('guild', 'yearly')).toBe('Guild — $150/year');
  });

  it('has no exclamation marks anywhere (honest-copy rule)', () => {
    const strings = [
      premiumCopy.freeForeverHeading,
      ...premiumCopy.freeForeverBody,
      premiumCopy.freeForeverNote,
      premiumCopy.noAds,
      ...premiumCopy.premiumFeatures,
      ...premiumCopy.guildFeatures,
      ...premiumCopy.faq.flatMap((entry) => [entry.question, entry.answer]),
    ];
    for (const text of strings) expect(text).not.toContain('!');
  });

  it('states plainly that nothing from Battle.net is ever paid', () => {
    const hit = premiumCopy.faq.find((entry) => entry.question.includes('Battle.net'));
    expect(hit?.answer).toMatch(/never paid|stays free|no\./i);
  });
});

describe('checkoutCopy', () => {
  it('has an honest, non-failing sentence for a not-yet-open storefront', () => {
    expect(checkoutCopy.notOpenYet).not.toContain('!');
    expect(checkoutCopy.notOpenYet.length).toBeGreaterThan(0);
  });
});

describe('billingBlockCopy', () => {
  it('names the guild billing contact without inventing a fact', () => {
    expect(billingBlockCopy.guildBilledBy('Fixture#1234')).toContain('Fixture#1234');
  });
});
