// web/src/lib/ui/card.ts
// The design system's card surface as class strings, for boxes that are not yet a `Card`
// component (site alignment spec 2026-10-29 S1b, polish spec 4.2). The utilities live in
// global.css: the warm gradient, the gold-tinted border, the 12 px panel radius and the deep shadow.
/** A primary card: the surface with the inset gold top edge (the /account card). */
export const CARD = 'card-surface';
/** A secondary card (prose asides, tables inside prose, tool-canvas frames): no top edge. */
export const CARD_QUIET = 'card-surface-quiet';

/** What a card's body holds, which sets its padding (site alignment spec 2026-10-29 §3.4). */
export type CardBody = 'text' | 'rows' | 'table' | 'none';

/** A card's own box: the surface, clipped so rows and dividers run edge to edge. */
export function cardClass(quiet: boolean): string {
  return `${quiet ? CARD_QUIET : CARD} overflow-hidden`;
}

/**
 * The card head (§3.3): a 36 px tile, the title and a right slot in a 68 px head; the compact head
 * (§3.3.1) is 48 px with a 28 px tile and the 18 px card title, for dense data cards.
 */
export const CARD_HEAD = {
  root: 'flex min-h-[68px] flex-wrap items-center gap-x-3 gap-y-2 px-5 py-4',
  compact: 'flex min-h-12 flex-wrap items-center gap-x-[10px] gap-y-1 px-4 py-[10px]',
  titleBlock: 'flex min-w-0 flex-1 flex-col gap-1',
  label: 'text-muted text-[11px] leading-none font-bold tracking-[.14em] uppercase',
  titleRow: 'flex flex-wrap items-center gap-x-2 gap-y-1',
  aside: 'ml-auto min-w-0 text-right text-[13px]',
  meta: 'text-muted',
} as const;

/** The tile size in each head. */
export function cardHeadTileSize(compact: boolean): 36 | 28 {
  return compact ? 28 : 36;
}

/** The title's recipe: the card title (18) when compact; 24 uppercase alone; its copy's case under a micro-label. */
export function cardHeadTitleClass(compact: boolean, labelled: boolean): string {
  if (compact) return 'section-title leading-6';
  return labelled ? 'heading-plan' : 'heading-plan uppercase';
}

/** The body's padding per kind (§3.4): text and table bodies draw the head's divider; rows bring their own top rules. */
export const CARD_BODY: Readonly<Record<CardBody, string>> = {
  text: 'border-line-soft flex flex-col gap-3 border-t px-5 py-4 text-[14px] leading-[22px]',
  rows: 'flex flex-col',
  table: 'border-line-soft border-t px-5 pb-2',
  none: '',
};

/** The foot (§3.4.4): left-aligned buttons 12 px apart over a 1 px rule. */
export const CARD_FOOT = 'border-line-soft flex flex-wrap items-center gap-3 border-t px-5 py-4';
