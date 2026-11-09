// web/src/lib/guild/field-styles.ts
// The guild forms' native field chrome (raids spec 3.5): a 44 px field for a date, time, text or
// textarea. Selects are the shared ui/Select.
import { FIELD_CONTROL, FIELD_FOCUS, fieldEdge } from '../ui/field';

// The skin is the shared Field's (lib/ui/field.ts): the page ground, the line edge, the gold focus
// ring; the warm-bordered raised field it replaced was the old card look.
const FIELD_SKIN = `${FIELD_CONTROL} ${FIELD_FOCUS} ${fieldEdge(false)} aria-invalid:border-wipe min-w-0 disabled:cursor-not-allowed`;

/** A one-line field, 44 px. */
export const FIELD_CHROME = `${FIELD_SKIN} h-11 px-3`;

/** A textarea: the same chrome, its height from `rows`. */
export const TEXTAREA_CHROME = `${FIELD_SKIN} px-3 py-2.5 leading-5`;

/** A checkbox row's label. */
export const CHECK_ROW = 'flex min-h-11 items-center gap-3 text-[14px]';

/** An inline switch: a 36 x 20 track with a 14 px thumb on a native checkbox, so the row's label
 *  stays the 44 px hit area. `SWITCH_ON` adds the gold track and the thumb at the right; a switch
 *  that cannot be used (disabled by another switch) leaves it off and reads neutral. */
export const SWITCH_CHROME =
  'relative h-5 w-9 shrink-0 cursor-pointer appearance-none rounded-full border border-line bg-bg transition-colors motion-reduce:transition-none ' +
  "before:absolute before:top-1/2 before:left-0.5 before:h-3.5 before:w-3.5 before:-translate-y-1/2 before:rounded-full before:bg-muted before:transition-transform motion-reduce:before:transition-none before:content-[''] " +
  'focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gold disabled:cursor-not-allowed';
export const SWITCH_ON =
  'checked:border-teal-edge checked:bg-teal checked:before:translate-x-4 checked:before:bg-strong';

/** The shared disabled look, for a row whose control cannot be used at all (not merely busy). */
export const DISABLED_LOOK = 'cursor-not-allowed opacity-50';

/** Below md a 44 px ::after, centred on an inline link, is its phone hit area. */
export const PHONE_HIT_AREA =
  "relative after:absolute after:inset-x-0 after:top-1/2 after:h-11 after:-translate-y-1/2 after:content-[''] md:after:hidden";

/** An inline link inside a line of prose (the GuildPlanGate recipe): the phone hit area; the focus
 *  ring stays on the text. */
export const INLINE_LINK = `text-gold hover:text-gold-hover active:text-gold-hover focus-visible:outline-gold whitespace-nowrap underline focus-visible:outline-2 focus-visible:outline-offset-2 ${PHONE_HIT_AREA}`;
