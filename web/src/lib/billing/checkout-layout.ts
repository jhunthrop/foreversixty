// web/src/lib/billing/checkout-layout.ts
// The /premium/checkout frame's height reserves (spec 2026-10-27-subscribe-and-join §5.7, §7,
// §11): every box that fills in after the static render or after /v1/me answers holds its
// ready height from the first paint, so the page's CLS stays within 0.02. Tailwind classes,
// written whole so the class scanner sees them.

/**
 * The guild step's lead slot while loading: the LONGEST lead in checkout-guild-copy.ts, the
 * 136-character `allUnclaimed(2)` sentence, at 21 px a line. At 1440 the 720 px column at 15 px
 * (about 100 characters a line) wraps it to 2 lines, 42 px. At 390 the 354 px column at 14 px
 * (about 53 characters a line) wraps it to 3 lines, 63 px.
 */
export const GUILD_LEAD_RESERVE = 'min-h-[63px] md:min-h-[42px]';

/**
 * The skeleton card's height while /v1/me is in flight: a `ready` card, 156 px from 768 and 260 px
 * on a phone (spec §5.7, §7). Real cards carry no floor and size to their content; a taller card
 * (the owner's `ask_leader`) grows below the fold, and the CLS e2e holds the result at 0.02.
 */
export const GUILD_SKELETON_CARD = 'h-[260px] md:h-[156px]';

/** The pixel figures behind the classes above, for the pre-paint frame reserve below. */
export const CHECKOUT_FRAME_PX = {
  /** The PageHeader with the plan summary (CheckoutHead), measured on the build. */
  head: { phone: 163, md: 206, lg: 150 },
  /** The guild step's top padding (pt-5 / md:pt-8). */
  stepTop: { phone: 20, md: 32 },
  lead: { phone: 63, md: 42 },
  /** The lead's margin below (mb-4). */
  leadGap: 16,
  card: { phone: 260, md: 156 },
  /** The card list's gap (gap-3). */
  cardGap: 12,
  /** The foot: 24 px above, two 13 px lines. */
  foot: 63,
  /** The guild step's bottom padding (pb-16). */
  bottom: 64,
} as const;

/** The custom property checkout.astro's head script sets to the skeleton card count; the main
 *  column's min-height reads it, so removing it releases the reserve. */
export const GUILD_CARDS_PROPERTY = '--checkout-guild-cards';

/**
 * Set on <html> by the head script while the guild step loads. The site footer is hidden meanwhile:
 * a real card can be taller or shorter than its 156 / 260 px skeleton (the owner's `ask_leader`
 * card is about 330 px), and the footer would otherwise be seen moving when the rows arrive (a
 * layout shift, which Chromium counts even for a move into view); first drawn at its final place,
 * it never moves on screen.
 */
export const GUILD_LOADING_ATTRIBUTE = 'data-checkout-loading';

/** Spec §7 (G2): the pre-paint reserve holds only while the step loads; once rows or an empty state
 *  render, the page is as tall as its content and the footer shows. */
export function releaseGuildFrameReserve(root: HTMLElement): void {
  root.style.removeProperty(GUILD_CARDS_PROPERTY);
  root.removeAttribute(GUILD_LOADING_ATTRIBUTE);
}

/**
 * The load error keeps the reserve (at least the viewport, so the footer stays below the fold) and
 * only shows the footer. Released there, the footer would be first drawn in view seconds after
 * load and become the page's largest paint (Lighthouse's signed-out run reaches this state).
 */
export function showFooterKeepReserve(root: HTMLElement): void {
  root.removeAttribute(GUILD_LOADING_ATTRIBUTE);
}

export type FrameBreakpoint = 'phone' | 'md' | 'lg';

export interface GuildFrameReserve {
  /** Everything but the cards: band, step padding, lead, foot, less one card gap. */
  fixedPx: number;
  /** One card and its gap. */
  perCardPx: number;
}

/**
 * Spec §7 / §11: the main column's height for the guild step while it loads, as fixed + cards x
 * perCard, held from the first paint by checkout.astro's head script so hydration moves no footer.
 */
export function guildFrameReserve(breakpoint: FrameBreakpoint): GuildFrameReserve {
  const px = CHECKOUT_FRAME_PX;
  const wide = breakpoint !== 'phone';
  const fixedPx =
    px.head[breakpoint] +
    (wide ? px.stepTop.md : px.stepTop.phone) +
    (wide ? px.lead.md : px.lead.phone) +
    px.leadGap +
    px.foot +
    px.bottom -
    px.cardGap;
  return { fixedPx, perCardPx: (wide ? px.card.md : px.card.phone) + px.cardGap };
}
