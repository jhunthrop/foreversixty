// web/src/lib/guild/billing-styles.ts
// The control recipes the guild settings #billing blocks share, spelled once.
import { SECONDARY_BUTTON_FIXED } from '../planner/styles';

export const BILLING_OUTLINE_BUTTON = `${SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-4`;
/** `Move to Raid guild`: the gold outline, as the plans page's card button. */
export const BILLING_MOVE_BUTTON = `${SECONDARY_BUTTON_FIXED} border-gold text-gold w-fit px-4`;
/** A 44 px underlined text button (`Keep Guild`). */
export const BILLING_TEXT_BUTTON =
  'text-text hover:text-gold-hover focus-visible:outline-gold inline-flex min-h-11 items-center text-[13px] underline focus-visible:outline-2';
