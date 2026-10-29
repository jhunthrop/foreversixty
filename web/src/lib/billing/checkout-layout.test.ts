// web/src/lib/billing/checkout-layout.test.ts
// Pins the checkout frame's reserves to the spec heights (§5.7, §7) and the panel they hold.
import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import {
  CHECKOUT_BAND_ATTRIBUTE,
  CHECKOUT_BAND_RESERVE,
  CHECKOUT_BODY_RESERVE,
  CHECKOUT_FRAME_PX,
  GUILD_CARDS_PROPERTY,
  GUILD_LOADING_ATTRIBUTE,
  GUILD_SKELETON_CARD,
  PREMIUM_BAND_PX,
  GUILD_LEAD_RESERVE,
  guildFrameReserve,
  releaseGuildFrameReserve,
  showFooterKeepReserve,
} from './checkout-layout';

describe('checkout layout reserves', () => {
  it('the band holds one box at phone, tablet and desktop', () => {
    expect(CHECKOUT_BAND_RESERVE).toBe('min-h-[163px] md:min-h-[206px] lg:min-h-[150px]');
  });

  it('the body holds the PremiumOwnedPanel, not just the 21 px message line', () => {
    expect(CHECKOUT_BODY_RESERVE).toBe('min-h-[204px] md:min-h-[184px]');
  });

  it('the guild lead holds the longest lead: 2 lines at 1440, 3 at 390', () => {
    expect(GUILD_LEAD_RESERVE).toBe('min-h-[63px] md:min-h-[42px]');
  });

  it('the skeleton card is the ready height: 260 px phone, 156 px from 768', () => {
    expect(GUILD_SKELETON_CARD).toBe('h-[260px] md:h-[156px]');
  });
});

describe('the pixel table matches the classes', () => {
  const px = CHECKOUT_FRAME_PX;

  it('band, lead and card figures are the ones the classes carry', () => {
    expect(CHECKOUT_BAND_RESERVE).toBe(
      `min-h-[${px.band.phone}px] md:min-h-[${px.band.md}px] lg:min-h-[${px.band.lg}px]`,
    );
    expect(GUILD_LEAD_RESERVE).toBe(`min-h-[${px.lead.phone}px] md:min-h-[${px.lead.md}px]`);
    expect(GUILD_SKELETON_CARD).toBe(`h-[${px.card.phone}px] md:h-[${px.card.md}px]`);
  });
});

describe('guildFrameReserve', () => {
  it('phone: band 163 + 20 + lead 63 + 16 + foot 63 + 64 - 12, then 260 + 12 a card', () => {
    expect(guildFrameReserve('phone')).toEqual({ fixedPx: 377, perCardPx: 272 });
  });

  it('md: band 206 + 32 + lead 42 + 16 + 63 + 64 - 12, then 156 + 12 a card', () => {
    expect(guildFrameReserve('md')).toEqual({ fixedPx: 411, perCardPx: 168 });
  });

  it('lg: the band is 150, the rest as md', () => {
    expect(guildFrameReserve('lg')).toEqual({ fixedPx: 355, perCardPx: 168 });
  });
});

describe('the premium band', () => {
  it('CheckoutBand carries PREMIUM_BAND_PX under the head script mark', () => {
    const band = readFileSync(
      new URL('../../components/checkout/CheckoutBand.svelte', import.meta.url),
      'utf8',
    );
    expect(band).toMatch(
      new RegExp(
        `data-checkout-band='premium'\\]\\) \\.checkout-band-row \\{\\s*min-height: ${PREMIUM_BAND_PX.phone}px;`,
      ),
    );
    expect(band).toContain(`min-height: ${PREMIUM_BAND_PX.md}px;`);
    expect(CHECKOUT_BAND_ATTRIBUTE).toBe('data-checkout-band');
  });
});

describe('showFooterKeepReserve', () => {
  it('drops the loading mark and keeps the card count', () => {
    const properties = new Map<string, string>([[GUILD_CARDS_PROPERTY, '2']]);
    const attributes = new Set<string>([GUILD_LOADING_ATTRIBUTE]);
    const root = {
      style: { removeProperty: (name: string) => properties.delete(name) },
      removeAttribute: (name: string) => attributes.delete(name),
    } as unknown as HTMLElement;
    showFooterKeepReserve(root);
    expect([...properties.keys()]).toEqual([GUILD_CARDS_PROPERTY]);
    expect([...attributes]).toEqual([]);
  });
});

describe('releaseGuildFrameReserve', () => {
  it('removes the card count the min-height reads and the loading mark, and only those', () => {
    const properties = new Map<string, string>([
      [GUILD_CARDS_PROPERTY, '2'],
      ['color', 'red'],
    ]);
    const attributes = new Set<string>([GUILD_LOADING_ATTRIBUTE, 'lang']);
    const root = {
      style: { removeProperty: (name: string) => properties.delete(name) },
      removeAttribute: (name: string) => attributes.delete(name),
    } as unknown as HTMLElement;
    releaseGuildFrameReserve(root);
    expect([...properties.keys()]).toEqual(['color']);
    expect([...attributes]).toEqual(['lang']);
  });
});
