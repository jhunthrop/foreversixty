// web/src/lib/guild/raid-buttons.ts
// The raid page's button recipes, shared by the officer row, the cancel confirm and the
// attendance surfaces (raids spec 4 "Officer buttons"): the system's secondary button (lib/ui/button),
// and the danger button that only a destructive answer wears.
import { button } from '../ui/button';

export const RAID_BUTTON = button('secondary', 'md');

/** The confirm's destructive answer: the danger outline. */
export const RAID_WIPE_BUTTON = button('danger', 'md');

/** A 44 px underlined text button (`Cancel`, `Not this log`). */
export const RAID_TEXT_BUTTON =
  'text-text focus-visible:outline-gold inline-flex min-h-11 items-center text-[13px] underline focus-visible:outline-2 focus-visible:outline-offset-2 disabled:cursor-not-allowed disabled:opacity-50';
