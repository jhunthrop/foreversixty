// web/src/lib/billing/copy.test.ts
import { describe, expect, it } from 'vitest';
import { specCell } from '../../test-support/spec-table';
import { premiumCopy, checkoutCopy, billingBlockCopy, premiumBlockCopy } from './copy';

describe('premiumCopy', () => {
  it('builds a priced plan name from the one price data file, never a hardcoded literal', () => {
    expect(premiumCopy.pricedName('premium', 'monthly')).toBe('Premium — $6 a month');
    expect(premiumCopy.pricedName('guild', 'yearly')).toBe('Guild — $150 a year');
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

describe('checkout chooser', () => {
  it('matches spec 6.1', () => {
    const spec = new URL('../../../../design/specs/2026-10-27-launch-pricing.md', import.meta.url);
    expect([
      checkoutCopy.whichGuild,
      checkoutCopy.verifiedMembers(0).replace('0', '{n}'),
      'Guild',
      checkoutCopy.fitsRaidGuild(0).replace('0', '{n}'),
      checkoutCopy.noBillableGuild,
      checkoutCopy.findYourGuild,
      checkoutCopy.guildsDidNotLoad,
      checkoutCopy.tryAgain,
    ]).toEqual(specCell(spec, '### 6.1', 'Checkout chooser'));
  });
});

describe('billingBlockCopy', () => {
  it('names the guild billing contact without inventing a fact', () => {
    expect(billingBlockCopy.guildBilledBy('Fixture#1234')).toContain('Fixture#1234');
  });
});

describe('what the page promises', () => {
  it('promises only what exists, and dates what is coming', () => {
    expect(premiumCopy.premiumFeatures).toEqual([
      'Run sims on our servers (5,000 combinations, any precision, no tab left open).',
      'Character history charts over time.',
      'Discord notifications when a server sim finishes or a log finishes parsing.',
      'A supporter mark on your characters.',
    ]);
    expect(premiumCopy.guildFeatures).toEqual([
      'Everything in Premium, for everyone in the guild.',
      'Officer tools: the loot council helper and the raid readiness board, with more as they ship.',
    ]);
    expect(premiumCopy.faq.map((f) => f.question)).not.toContain(
      'What happens to my long-retention logs if I stop?',
    );
    for (const text of [
      ...premiumCopy.premiumFeatures,
      ...premiumCopy.guildFeatures,
      ...premiumCopy.faq.map((f) => f.answer),
    ]) {
      expect(text).not.toMatch(/retention|ninety|two-year|two years|Compare more than two/);
    }
  });
});

describe('premiumBlockCopy history row', () => {
  it('carries the three states in the spec words', () => {
    expect(premiumBlockCopy.rows.history).toMatchObject({
      on: 'On',
      onWithMain: 'Open on your main character',
      off: 'Sparkline only',
    });
  });

  it('dates nothing: every benefit on the page has shipped', () => {
    expect(premiumCopy.premiumFeatures.join(' ')).not.toMatch(/November|December|ships/i);
  });
});

describe('premiumBlockCopy server sims row', () => {
  const { on } = premiumBlockCopy.rows.serverSims;

  it('counts the month in plain words', () => {
    expect(on(4)).toBe('On · 4 runs this month');
    expect(on(1)).toBe('On · 1 run this month');
    expect(on(0)).toBe('On · no runs yet this month');
    expect(on(1000)).toBe('On · 1,000 runs this month');
  });

  it('is plain On when an older /v1/me sent no count', () => {
    expect(on(undefined)).toBe('On');
  });
});
