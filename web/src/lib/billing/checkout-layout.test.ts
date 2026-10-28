// web/src/lib/billing/checkout-layout.test.ts
// Pins the checkout frame's reserves to the spec heights (§5.7, §7) and the panel they hold.
import { describe, expect, it } from 'vitest';
import {
  CHECKOUT_BAND_RESERVE,
  CHECKOUT_BODY_RESERVE,
  GUILD_CARD_MIN,
  GUILD_LEAD_RESERVE,
} from './checkout-layout';

describe('checkout layout reserves', () => {
  it('the band holds one box at phone, tablet and desktop', () => {
    expect(CHECKOUT_BAND_RESERVE).toBe('min-h-[156px] md:min-h-[190px] lg:min-h-[150px]');
  });

  it('the body holds the PremiumOwnedPanel, not just the 21 px message line', () => {
    expect(CHECKOUT_BODY_RESERVE).toBe('min-h-[204px] md:min-h-[184px]');
  });

  it('the guild lead holds the longest lead: 2 lines at 1440, 3 at 390', () => {
    expect(GUILD_LEAD_RESERVE).toBe('min-h-[63px] md:min-h-[42px]');
  });

  it('a card never falls below the skeleton it replaces', () => {
    expect(GUILD_CARD_MIN).toBe('min-h-[260px] md:min-h-[156px]');
  });
});
