// web/src/lib/guild/field-styles.ts
// The guild forms' native field chrome (raids spec 3.5): FilterSelect's 44 px warm border on a
// date, time, text, textarea or grouped select that FilterSelect itself cannot draw.

const FIELD_BASE =
  'border-line-warm bg-raised text-text rounded-control placeholder:text-muted hover:border-line-warm-strong focus-visible:outline-gold aria-invalid:border-wipe w-full min-w-0 border px-3 text-[13px] transition-colors duration-[120ms] focus-visible:outline-2 disabled:cursor-not-allowed';

/** A one-line field: input or select, 44 px. */
export const FIELD_CHROME = `${FIELD_BASE} h-11`;

/** A textarea: the same chrome, its height from `rows`. */
export const TEXTAREA_CHROME = `${FIELD_BASE} py-2.5 leading-5`;

/** An inline link inside a line of prose (the GuildPlanGate recipe): below md a 44 px ::after,
 *  centred on the link, is its phone hit area; the focus ring stays on the text. */
export const INLINE_LINK =
  "text-gold hover:text-gold-hover active:text-gold-hover focus-visible:outline-gold relative whitespace-nowrap underline after:absolute after:inset-x-0 after:top-1/2 after:h-11 after:-translate-y-1/2 after:content-[''] focus-visible:outline-2 focus-visible:outline-offset-2 md:after:hidden";
