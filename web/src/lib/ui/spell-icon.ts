// web/src/lib/ui/spell-icon.ts
// What ui/SpellIcon.svelte and its static twin ui/SpellIcon.astro share (component audit
// 2026-11-10 cluster 1): every spell, ability, buff, enchant, consumable and spec icon on the site
// is one frame, a 1 px warm line at the icon radius on the dark socket, at one of these sizes.
// Items keep ui/ItemIcon, whose frame is the item's quality.

/** The sizes a spell icon is drawn at, in CSS px. */
export const SPELL_ICON_SIZES = [14, 16, 18, 20, 22, 24, 28, 32, 36, 40] as const;
export type SpellIconSize = (typeof SPELL_ICON_SIZES)[number];

/** The frame: the warm line, the icon radius, the dark socket behind the art. */
export const SPELL_ICON_FRAME =
  'ui-spell-icon border-rule-warm rounded-icon inline-flex shrink-0 overflow-hidden border bg-[var(--color-bg)]';

/** The art inside the frame. */
export const SPELL_ICON_IMAGE = 'block h-full w-full object-cover';

/** An icon's accessible face. */
export interface SpellIconA11y {
  role?: 'img';
  'aria-label'?: string;
  'aria-hidden'?: 'true';
  title?: string;
}

/** A named icon (one that stands alone) is an image with that name; an icon beside its printed
 *  name is decoration, hidden. */
export function spellIconA11y(label: string): SpellIconA11y {
  return label === '' ? { 'aria-hidden': 'true' } : { role: 'img', 'aria-label': label, title: label };
}
