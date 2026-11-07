// web/src/lib/billing/checkout-layout.test.ts
// Pins the checkout frame's reserves to the spec heights (§5.7, §7) and the panel they hold.
import { describe, expect, it } from 'vitest';
import {
  CHECKOUT_FRAME_PX,
  GUILD_CARDS_PROPERTY,
  GUILD_LOADING_ATTRIBUTE,
  GUILD_SKELETON_CARD,
  GUILD_LEAD_RESERVE,
  guildFrameReserve,
  releaseGuildFrameReserve,
  showFooterKeepReserve,
} from './checkout-layout';

describe('checkout layout reserves', () => {
  it('the guild lead holds the longest lead: 2 lines at 1440, 3 at 390', () => {
    expect(GUILD_LEAD_RESERVE).toBe('min-h-[63px] md:min-h-[42px]');
  });

  it('the skeleton card is the ready height: 260 px phone, 156 px from 768', () => {
    expect(GUILD_SKELETON_CARD).toBe('h-[260px] md:h-[156px]');
  });
});

describe('the pixel table matches the classes', () => {
  const px = CHECKOUT_FRAME_PX;

  it('lead and card figures are the ones the classes carry', () => {
    expect(GUILD_LEAD_RESERVE).toBe(`min-h-[${px.lead.phone}px] md:min-h-[${px.lead.md}px]`);
    expect(GUILD_SKELETON_CARD).toBe(`h-[${px.card.phone}px] md:h-[${px.card.md}px]`);
  });
});

describe('guildFrameReserve', () => {
  const px = CHECKOUT_FRAME_PX;
  it.each(['phone', 'md', 'lg'] as const)(
    '%s: head, step top, lead, gap, foot, bottom less one card gap',
    (bp) => {
      const wide = bp !== 'phone';
      expect(guildFrameReserve(bp)).toEqual({
        fixedPx:
          px.head[bp] +
          (wide ? px.stepTop.md : px.stepTop.phone) +
          (wide ? px.lead.md : px.lead.phone) +
          px.leadGap +
          px.foot +
          px.bottom -
          px.cardGap,
        perCardPx: (wide ? px.card.md : px.card.phone) + px.cardGap,
      });
    },
  );
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
