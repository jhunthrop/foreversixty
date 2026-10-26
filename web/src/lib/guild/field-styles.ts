// web/src/lib/guild/field-styles.ts
// The guild forms' native field chrome (raids spec 3.5): FilterSelect's 44 px warm border on a
// date, time, text, textarea or grouped select that FilterSelect itself cannot draw.

const FIELD_SKIN =
  'border-line-warm bg-raised text-text rounded-control placeholder:text-muted hover:border-line-warm-strong focus-visible:outline-gold aria-invalid:border-wipe w-full min-w-0 border text-[13px] transition-colors duration-[120ms] focus-visible:outline-2 disabled:cursor-not-allowed';

/** A one-line field, 44 px. */
export const FIELD_CHROME = `${FIELD_SKIN} h-11 px-3`;

/** A select: FilterSelect's own box, no native arrow, room on the right for `SelectChevron`. */
export const SELECT_CHROME = `${FIELD_SKIN} h-11 cursor-pointer appearance-none truncate pr-8 pl-3`;

/** A textarea: the same chrome, its height from `rows`. */
export const TEXTAREA_CHROME = `${FIELD_SKIN} px-3 py-2.5 leading-5`;

/** A checkbox row's label. */
export const CHECK_ROW = 'flex min-h-11 items-center gap-3 text-[14px]';

/** The shared disabled look, for a row whose control cannot be used at all (not merely busy). */
export const DISABLED_LOOK = 'cursor-not-allowed opacity-50';

/** An inline link inside a line of prose (the GuildPlanGate recipe): below md a 44 px ::after,
 *  centred on the link, is its phone hit area; the focus ring stays on the text. */
export const INLINE_LINK =
  "text-gold hover:text-gold-hover active:text-gold-hover focus-visible:outline-gold relative whitespace-nowrap underline after:absolute after:inset-x-0 after:top-1/2 after:h-11 after:-translate-y-1/2 after:content-[''] focus-visible:outline-2 focus-visible:outline-offset-2 md:after:hidden";
