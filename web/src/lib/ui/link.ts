// web/src/lib/ui/link.ts
// The action link (polish spec 4.2): a standalone action that is not prose. 14 px 600 gold, a
// trailing chevron (ui/ActionLink.svelte draws it), 44 px tall, no underline; hover lightens the
// gold and the chevron moves 2 px right. Inline links in prose keep INLINE_LINK (guild/field-styles).
export const ACTION_LINK =
  'group text-gold inline-flex min-h-11 w-fit items-center gap-1 text-[14px] font-semibold transition-colors duration-[120ms] ' +
  'hover:text-gold-hover active:text-gold-hover focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gold rounded-control';

/** The chevron's own motion, inside an ACTION_LINK (`group`). */
export const ACTION_LINK_CHEVRON = 'transition-transform duration-[120ms] group-hover:translate-x-0.5';
