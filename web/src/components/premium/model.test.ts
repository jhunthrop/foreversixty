// web/src/components/premium/model.test.ts
import { describe, expect, it } from 'vitest';
import { GUILD_TRIAL_DAYS, PLANS, formatPrice, perMonthLine } from '../../lib/billing/plans';
import { perMemberLine, perRaiderLine, plansPageCopy } from '../../lib/billing/plans-page-copy';
import {
  buttonLabel,
  checkoutHref,
  costLine,
  featureLines,
  founderKey,
  planBlock,
  priceCell,
  trialParts,
} from './model';

describe('checkoutHref', () => {
  it('carries the plan and interval and no guild id', () => {
    expect(checkoutHref('premium', 'yearly')).toBe('/premium/checkout?plan=premium&interval=yearly');
    expect(checkoutHref('raid_guild', 'monthly')).toBe('/premium/checkout?plan=raid_guild&interval=monthly');
  });
});

describe('buttonLabel', () => {
  it('says Subscribe for Premium and Subscribe a guild for the guild plans', () => {
    expect(buttonLabel('premium', 'yearly', true)).toBe(
      `Subscribe · ${formatPrice(PLANS.premium.founderYearly, 'yearly')}`,
    );
    expect(buttonLabel('guild', 'monthly', false)).toBe(
      `Subscribe a guild · ${formatPrice(PLANS.guild.monthly, 'monthly')}`,
    );
  });
});

describe('priceCell', () => {
  it('puts the list price and the founder date on one founder line while founder is open', () => {
    const cell = priceCell('premium', 'yearly', true);
    expect(cell.amount).toBe(`$${PLANS.premium.founderYearly}`);
    expect(cell.founderLine).toBe('Founder price until 6 Jan 2027, then $50.');
    expect(priceCell('premium', 'monthly', true).founderLine).toBe(
      'Founder price until 6 Jan 2027, then $6 a month.',
    );
    expect(priceCell('guild', 'yearly', true).founderLine).toBe('Founder price until 6 Jan 2027, then $150.');
    expect(priceCell('raid_guild', 'monthly', true).founderLine).toBe(
      'Founder price until 6 Jan 2027, then $25 a month.',
    );
    expect(cell.suffix).toBe(plansPageCopy.price.suffix.yearly);
    expect(cell.secondLine).toBe(perMonthLine(PLANS.premium.founderYearly));
  });

  it('shows list prices with no founder line once it is closed', () => {
    const cell = priceCell('guild', 'monthly', false);
    expect(cell.amount).toBe(`$${PLANS.guild.monthly}`);
    expect(cell.founderLine).toBeNull();
    expect(cell.secondLine).toBe(plansPageCopy.price.orYearly(formatPrice(PLANS.guild.yearly, 'yearly')));
  });
});

describe('costLine', () => {
  it('is the per-member line for Guild, the per-raider line for Raid guild, none for Premium', () => {
    expect(costLine('premium', 'yearly', true)).toBeNull();
    expect(costLine('guild', 'yearly', true)).toBe(perMemberLine('yearly', true));
    expect(costLine('raid_guild', 'monthly', false)).toBe(perRaiderLine('monthly', false));
  });
});

describe('trialParts', () => {
  it('slices the strong lead off the trial sentence without losing a word', () => {
    const { lead, rest } = trialParts();
    expect(lead).toBe(`First ${GUILD_TRIAL_DAYS} days free`);
    expect(lead + rest).toBe(plansPageCopy.trial);
    expect(rest.startsWith(', once per guild')).toBe(true);
  });
});

describe('planBlock and founderKey', () => {
  it('puts the coach line first in the Premium features', () => {
    expect(planBlock('premium').features).toBe(plansPageCopy.premium.features);
    expect(featureLines(planBlock('premium').features)[0]?.text).toMatch(/^Coach: /);
  });

  it('maps each plan to its copy block', () => {
    expect(planBlock('raid_guild')).toBe(plansPageCopy.raidGuild);
    expect(founderKey(true)).toBe('open');
    expect(founderKey(false)).toBe('closed');
  });
});
