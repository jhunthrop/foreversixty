// web/src/lib/ui/field.ts
// The text field's classes (web layout spec 2026-11-05 §3A.2): ui/Field.svelte and the static
// lookup form (CharacterSearch.astro, whose input carries the typeahead's combobox attributes)
// draw the one field.

/** The input itself: the field radius, the page ground, 14 px. */
export const FIELD_CONTROL =
  'rounded-field bg-bg text-strong placeholder:text-muted scheme-dark w-full border px-3 text-[14px]';
/** The focus ring. */
export const FIELD_FOCUS =
  'focus-visible:outline-gold focus-visible:outline-2 focus-visible:outline-offset-2';
/** The one-line height: 44 px on a phone, 40 from 768. */
export const FIELD_HEIGHT = 'h-11 md:h-10';
/** The edge: the line, or the wipe colour when invalid. */
export function fieldEdge(invalid: boolean): string {
  return invalid ? 'border-wipe' : 'border-line';
}
/** The label over it: 12 px caps, muted. */
export const FIELD_LABEL = 'text-label text-muted font-semibold';
