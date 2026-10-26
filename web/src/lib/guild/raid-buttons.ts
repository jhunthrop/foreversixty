// web/src/lib/guild/raid-buttons.ts
// The raid page's button recipes, shared by the officer row, the cancel confirm and the
// attendance surfaces (raids spec 4 "Officer buttons"): one outline, and the failure-coloured
// outline that only a destructive answer wears.
import { SECONDARY_BUTTON_FIXED } from '../planner/styles';

export const RAID_BUTTON = `${SECONDARY_BUTTON_FIXED} border-line-warm text-text px-4 active:bg-card-top disabled:cursor-not-allowed disabled:opacity-50`;

/** The confirm's destructive answer: the same outline in the failure colour. */
export const RAID_WIPE_BUTTON =
  'rounded-control border-wipe/60 text-wipe hover:border-wipe hover:bg-wipe/10 active:bg-wipe/15 inline-flex h-11 items-center border px-4 text-[12px] font-bold tracking-[0.06em] uppercase transition-colors duration-[120ms] disabled:cursor-not-allowed disabled:opacity-50';

/** A 44 px underlined text button (`Cancel`, `Not this log`). */
export const RAID_TEXT_BUTTON =
  'text-text focus-visible:outline-gold inline-flex min-h-11 items-center text-[13px] underline focus-visible:outline-2 focus-visible:outline-offset-2 disabled:cursor-not-allowed disabled:opacity-50';
