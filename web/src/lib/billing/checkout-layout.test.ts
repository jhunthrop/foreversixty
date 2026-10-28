// web/src/lib/billing/checkout-layout.test.ts
// Pins the checkout frame's reserves to the spec heights (§5.7, §7) and the panel they hold.
import { describe, expect, it } from 'vitest';
import {
  CHECKOUT_BAND_RESERVE,
  CHECKOUT_BODY_RESERVE,
  CHECKOUT_FRAME_PX,
  GUILD_CARD_MIN,
  GUILD_LEAD_RESERVE,
  guildFrameReserve,
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

  it('a card and its skeleton hold the ask_leader card: 532 px phone, 376 px from 768', () => {
    expect(GUILD_CARD_MIN).toBe('min-h-[532px] md:min-h-[376px]');
  });
});

describe('the pixel table matches the classes', () => {
  const px = CHECKOUT_FRAME_PX;

  it('band, lead and card figures are the ones the classes carry', () => {
    expect(CHECKOUT_BAND_RESERVE).toBe(
      `min-h-[${px.band.phone}px] md:min-h-[${px.band.md}px] lg:min-h-[${px.band.lg}px]`,
    );
    expect(GUILD_LEAD_RESERVE).toBe(`min-h-[${px.lead.phone}px] md:min-h-[${px.lead.md}px]`);
    expect(GUILD_CARD_MIN).toBe(`min-h-[${px.card.phone}px] md:min-h-[${px.card.md}px]`);
  });
});

describe('guildFrameReserve', () => {
  it('phone: band 156 + 20 + lead 63 + 16 + foot 63 + 64 - 12, then 532 + 12 a card', () => {
    expect(guildFrameReserve('phone')).toEqual({ fixedPx: 370, perCardPx: 544 });
  });

  it('md: band 190 + 32 + lead 42 + 16 + 63 + 64 - 12, then 376 + 12 a card', () => {
    expect(guildFrameReserve('md')).toEqual({ fixedPx: 395, perCardPx: 388 });
  });

  it('lg: the band is 150, the rest as md', () => {
    expect(guildFrameReserve('lg')).toEqual({ fixedPx: 355, perCardPx: 388 });
  });
});
