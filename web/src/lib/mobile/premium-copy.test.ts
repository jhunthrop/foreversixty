// The apps' Premium screen (mobile spec 5.8.1) lists what the plan includes in the site's own
// words (decision G: the apps carry the web's exact Premium rules). Every `plan.feature.*` line in
// apps/shared/copy/en.json must be a sentence of a Premium line on /premium, so the two cannot
// drift; a new app line means a new line on /premium first.
import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { plansPageCopy } from '../billing/plans-page-copy';
import { accountPlanCopy } from '../billing/account-plan-copy';
import { premiumSourceCopy } from '../billing/premium-source-copy';

const CATALOG = new URL('../../../../apps/shared/copy/en.json', import.meta.url);
const strings = (JSON.parse(readFileSync(CATALOG, 'utf8')) as { strings: Record<string, unknown> }).strings;
const text = (key: string): string => {
  const value = strings[key];
  if (typeof value !== 'string') throw new Error(`en.json has no plain string ${key}`);
  return value;
};

const siteLines = plansPageCopy.premium.features.map((feature) => feature.text);

describe('the apps say what Premium includes in the site words', () => {
  const featureKeys = Object.keys(strings).filter((key) => key.startsWith('plan.feature.'));

  it('lists at least one feature line', () => {
    expect(featureKeys.length).toBeGreaterThan(0);
  });

  it.each(featureKeys)('%s is a sentence of a /premium line', (key) => {
    expect(siteLines.some((line) => line.includes(text(key)))).toBe(true);
  });

  it('reads the plan lines the account page reads', () => {
    expect(text('plan.line.grantNoEnd')).toBe(premiumSourceCopy.grant.noEnd);
    expect(text('plan.given.grantNoEnd')).toBe(premiumSourceCopy.grant.givenNoEnd);
    expect(text('plan.given.grantUntil')).toBe(premiumSourceCopy.grant.givenUntil);
    expect(text('plan.seatGivenBy').replace('{name}', 'X')).toBe(accountPlanCopy.planLine.givenBy('X'));
    expect(text('plan.line.seat').replace('{guild}', 'G')).toBe(accountPlanCopy.planLine.seat('G'));
    expect(text('plan.founderNote')).toBe(accountPlanCopy.founder.title);
  });
});
