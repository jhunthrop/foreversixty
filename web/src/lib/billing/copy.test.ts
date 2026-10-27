// web/src/lib/billing/copy.test.ts
import { describe, expect, it } from 'vitest';
import { specCell } from '../../test-support/spec-table';
import { checkoutCopy, billingBlockCopy, premiumBlockCopy } from './copy';

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
      '{Plan} · {price}',
    ]).toEqual(specCell(spec, '### 6.1', 'Checkout chooser'));
  });

  it('names the plan and price picked on /premium', () => {
    expect(checkoutCopy.chosenPlan('raid_guild', 'yearly', true)).toBe('Raid guild · $200 a year');
    expect(checkoutCopy.chosenPlan('guild', 'monthly', false)).toBe('Guild · $15 a month');
  });
});

describe('billingBlockCopy', () => {
  it('names the guild billing contact without inventing a fact', () => {
    expect(billingBlockCopy.guildBilledBy('Fixture#1234')).toContain('Fixture#1234');
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
