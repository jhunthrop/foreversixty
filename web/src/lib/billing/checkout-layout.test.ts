// web/src/lib/billing/checkout-layout.test.ts
// Pins the checkout frame's reserves to the spec heights (§5.7, §7) and the panel they hold.
import { describe, expect, it } from 'vitest';
import {
  CHECKOUT_BAND_RESERVE,
  CHECKOUT_BODY_RESERVE,
  GUILD_LEAD_RESERVE,
  GUILD_SKELETON_CARD,
} from './checkout-layout';

describe('checkout layout reserves', () => {
  it('the band holds one box at phone, tablet and desktop', () => {
    expect(CHECKOUT_BAND_RESERVE).toBe('min-h-[156px] md:min-h-[184px] lg:min-h-[148px]');
  });

  it('the body holds the PremiumOwnedPanel, not just the 21 px message line', () => {
    expect(CHECKOUT_BODY_RESERVE).toBe('min-h-[204px] md:min-h-[184px]');
  });

  it('the guild lead reserves 21 px, 42 px below 768', () => {
    expect(GUILD_LEAD_RESERVE).toBe('min-h-[42px] md:min-h-[21px]');
  });

  it('a skeleton card is the ready card: 156 px desktop, 260 px phone', () => {
    expect(GUILD_SKELETON_CARD).toBe('h-[260px] md:h-[156px]');
  });
});
