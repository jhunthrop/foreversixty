// web/src/lib/guild/billing-styles.ts
// The control recipes the guild settings #billing blocks share, spelled once, over the button
// recipes (lib/ui/button.ts): sentence case, the 44 px hit target.
import { button } from '../ui/button';

export const BILLING_OUTLINE_BUTTON = `${button('secondary')} w-fit`;
/** `Move to Raid guild`: the gold outline, the one gold control on the plan line. */
export const BILLING_MOVE_BUTTON = `${button('secondary')} border-gold/60 text-gold w-fit`;
/** A 44 px underlined text button (`Keep Guild`). */
export const BILLING_TEXT_BUTTON =
  'text-text hover:text-gold-hover focus-visible:outline-gold inline-flex min-h-11 items-center text-[13px] underline focus-visible:outline-2';
/** The safe choice a confirm block focuses in code (spec 4): the 2 px gold ring shows on any focus,
 *  not only keyboard focus-visible, so the reader sees where focus went after a click. */
export const BILLING_SAFE_TEXT_BUTTON = `${BILLING_TEXT_BUTTON} focus:outline-gold focus:outline-2 focus:outline-offset-2`;
/** An opener (`Give seat`, `Move to Raid guild`) while its confirm block is open: drawn pressed. */
export const OPENER_PRESSED = 'bg-white/[.08]';
/** A confirm block's buttons: full width below 768, commit first (spec 7); inline from 768. */
export const CONFIRM_ACTIONS =
  'flex flex-col items-stretch gap-2 md:flex-row md:flex-wrap md:items-center md:gap-3';
/** A button inside CONFIRM_ACTIONS: full width and centred below 768. */
export const CONFIRM_ACTION_WIDTH = 'w-full justify-center md:w-fit';
/** A confirm block under the row that opened it (a seat, a transfer, the upgrade): open, set in
 *  from the row's text by a gold rule, never a box of its own. */
export const CONFIRM_BLOCK = 'flex flex-col gap-3 border-l-2 border-gold/40 py-1 pl-4';
