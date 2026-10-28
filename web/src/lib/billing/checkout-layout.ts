// web/src/lib/billing/checkout-layout.ts
// The /premium/checkout frame's height reserves (spec 2026-10-27-subscribe-and-join §5.7, §7,
// §11): every box that fills in after the static render or after /v1/me answers holds its
// ready height from the first paint, so the page's CLS stays within 0.02. Tailwind classes,
// written whole so the class scanner sees them.

/**
 * The band: the same box for a guild plan, Premium, an invalid link and the moment before the
 * URL is read. From the classes (line-height 1.5, label 11 px): phone 24 + h1 26.4 + 8 + summary
 * 97.5 (size line on two lines) = 155.9; md 52 + label/h1 50.9 + 8 + summary 78 = 188.9; lg 52 +
 * summary 97.5 (Guild's size line wraps in the 440 px column) = 149.5. Estimates: Task 25's CLS
 * e2e asserts the band height before and after the summary appears.
 */
export const CHECKOUT_BAND_RESERVE = 'min-h-[156px] md:min-h-[190px] lg:min-h-[150px]';

/**
 * The narrow body under the band (redirect message, refusals, invalid link): it holds Task 12's
 * PremiumOwnedPanel (§4.2) — padding 20, title, plan line, line 2, the 44 px button 12 px below —
 * so the panel replacing the one-line message moves nothing. Line 2 wraps once on a phone.
 */
export const CHECKOUT_BODY_RESERVE = 'min-h-[204px] md:min-h-[184px]';

/**
 * The guild step's lead slot while loading: the LONGEST lead in checkout-guild-copy.ts, the
 * 136-character `allUnclaimed(2)` sentence, at 21 px a line. At 1440 the 720 px column at 15 px
 * (about 100 characters a line) wraps it to 2 lines, 42 px. At 390 the 354 px column at 14 px
 * (about 53 characters a line) wraps it to 3 lines, 63 px.
 */
export const GUILD_LEAD_RESERVE = 'min-h-[63px] md:min-h-[42px]';

/**
 * A guild card's floor and the skeleton card's height: the tallest common card, `ask_leader`
 * (the owner's case), so a skeleton never grows when its row arrives. From the classes (line
 * height 1.5; the step block is 320 px wide at md and about 322 px on a 390 phone, gap 10):
 * sentence 2 lines 42 + CopyLinkField (label 16.5 + 4 + field 44 + 8 + buttons 44 = 116.5) +
 * short verify line 2 lines 39 + other-character line 3 lines 58.5 + Discord fallback 2 lines 39
 * + 4 gaps 40 = 335.
 * - md (two columns): padding 40 + max(step block 335, identity/you/facts ~110) = 375 -> 376.
 * - phone (one column): padding 32 + identity 40 + you row 8 + 42 (two lines) + facts 6 + 39
 *   + gap 14 + rule 1 + 14 + step block 335 = 531 -> 532.
 * Shorter cards (ready, on_plan) keep this floor. Estimates: Task 28 reads the CLS e2e annotation
 * and adjusts this once from the measurement if it is still above 0.02.
 */
export const GUILD_CARD_MIN = 'min-h-[532px] md:min-h-[376px]';

/** The pixel figures behind the classes above, for the pre-paint frame reserve below. */
export const CHECKOUT_FRAME_PX = {
  band: { phone: 156, md: 190, lg: 150 },
  /** The guild step's top padding (pt-5 / md:pt-8). */
  stepTop: { phone: 20, md: 32 },
  lead: { phone: 63, md: 42 },
  /** The lead's margin below (mb-4). */
  leadGap: 16,
  card: { phone: 532, md: 376 },
  /** The card list's gap (gap-3). */
  cardGap: 12,
  /** The foot: 24 px above, two 13 px lines. */
  foot: 63,
  /** The guild step's bottom padding (pb-16). */
  bottom: 64,
} as const;

export type FrameBreakpoint = 'phone' | 'md' | 'lg';

export interface GuildFrameReserve {
  /** Everything but the cards: band, step padding, lead, foot, less one card gap. */
  fixedPx: number;
  /** One card and its gap. */
  perCardPx: number;
}

/**
 * Spec §7 / §11: the main column's height for the guild step, as fixed + cards x perCard, held
 * from the first paint by checkout.astro's head script so hydration and /v1/me move no footer.
 */
export function guildFrameReserve(breakpoint: FrameBreakpoint): GuildFrameReserve {
  const px = CHECKOUT_FRAME_PX;
  const wide = breakpoint !== 'phone';
  const fixedPx =
    px.band[breakpoint] +
    (wide ? px.stepTop.md : px.stepTop.phone) +
    (wide ? px.lead.md : px.lead.phone) +
    px.leadGap +
    px.foot +
    px.bottom -
    px.cardGap;
  return { fixedPx, perCardPx: (wide ? px.card.md : px.card.phone) + px.cardGap };
}
