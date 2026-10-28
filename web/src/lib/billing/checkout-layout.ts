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
 * A guild card's floor, and the skeleton card's height (ready: 156 px desktop, 260 px phone): a shorter card (an on_plan one) would shrink the
 * list when the rows replace the skeleton. The shell's skeleton must use this same constant.
 */
export const GUILD_CARD_MIN = 'min-h-[260px] md:min-h-[156px]';
