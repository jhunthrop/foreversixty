// The Premium entry's strings against the spec's own copy table (2026-10-29-premium-entry
// section 10) and its anti-nag word list (section 2): a drift in either file fails here.
import { describe, expect, it } from 'vitest';
import { specCell } from '../test-support/spec-table';
import { PREMIUM_LINKS } from './billing/premium-links';
import { guildBillingCopy } from './billing/guild-billing-copy';
import { selectorCopy } from './character-selector/copy';
import { PREMIUM_NAV_ITEM } from './nav';
import { simCopy } from './sim/copy';

const SPEC = new URL('../../../design/specs/2026-10-29-premium-entry.md', import.meta.url);
const HEADING = '## 10. Copy';
const row = (key: string): string => specCell(SPEC, HEADING, key)[0] ?? '';

describe('Premium entry copy (spec section 10)', () => {
  it('the nav tag', () => expect(PREMIUM_NAV_ITEM.label).toBe(row('`PREMIUM_NAV_ITEM.label`')));
  it('the Plan row', () => {
    expect(selectorCopy.account.seePlans).toBe(row('`selectorCopy.account.seePlans`'));
    expect(selectorCopy.account.freePlanHref).toBe(row('`selectorCopy.account.freePlanHref`'));
  });
  it('the sim link', () => expect(simCopy.runOnServersFree).toBe(row('`simCopy.runOnServersFree`')));
  it('Get premium is gone', () => expect('getPremium' in simCopy).toBe(false));
  it('the guild chooser', () =>
    expect(guildBillingCopy.chooser.compare).toBe(row('`guildBillingCopy.chooser.compare`')));
});

describe('Premium entry hrefs (spec sections 7 and 8)', () => {
  it('every contextual link lands on its anchor, never the page top', () => {
    expect(PREMIUM_LINKS).toEqual({
      card: '/premium#premium',
      coach: '/premium#coach',
      serverSims: '/premium#server-sims',
      history: '/premium#history',
      notifications: '/premium#notifications',
      guildLeaders: '/premium#guild-leaders',
    });
  });
});

describe('the anti-nag words (spec section 2)', () => {
  it('none of the new strings uses a forbidden word or an exclamation mark', () => {
    const strings = [
      PREMIUM_NAV_ITEM.label,
      selectorCopy.account.seePlans,
      simCopy.runOnServersFree,
      guildBillingCopy.chooser.compare,
    ];
    for (const text of strings) {
      expect(text).not.toMatch(/Go Premium|Unlock|\bPro\b|Limited time|Upgrade|!/);
    }
  });
});
