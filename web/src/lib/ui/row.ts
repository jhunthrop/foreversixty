// web/src/lib/ui/row.ts
// The settings row's class strings (polish spec 4.2), shared by ui/SettingsRow.svelte and by the
// one row that must be a different element (the Discord reconnect link, which keeps its opening
// watchdog): 56 px minimum, 20 px side padding, 14 px gap, a 1 px top rule, a hover ground only
// when the row is interactive.
export const ROW =
  'border-line-soft flex min-h-14 w-full flex-wrap items-center gap-x-[14px] gap-y-1 border-t px-5 py-[10px] text-left';

/** The same row open, in a Section rather than a card (web layout spec 2026-11-05 §3A.2): no side
 *  padding, a hairline under each row but the last. */
export const ROW_OPEN = ROW.replace('border-t px-5', 'border-b px-0 last:border-b-0');

/** The hover ground alone, for a row that holds its own links and buttons (a character row). */
export const ROW_HOVER = 'transition-colors duration-[120ms] hover:bg-hover';

export const ROW_INTERACTIVE =
  'cursor-pointer transition-colors duration-[120ms] hover:bg-hover active:bg-gold/10 focus-visible:outline-2 focus-visible:-outline-offset-2 focus-visible:outline-gold';

/**
 * Row bands (ui/row.css `.row-band`): a row that lights up inside a list is a rounded band whose
 * content sits 12 px in. A row whose list pads it by less reaches out by the difference, so its
 * content keeps the list's left edge and the band grows into the gutter, keyed by the list's own
 * side padding in Tailwind steps (0, `px-1`, `px-2`). Apply it to the lit row only.
 */
export const ROW_BAND_REACH = {
  0: '-mx-3 px-3',
  1: '-mx-2 px-3',
  2: '-mx-1 px-3',
} as const;

/** Every row of a ruled list that holds a band: its rule steps back when the row beside it lights
 *  (ui/row.css keys the neighbour selectors to it, never to a universal subject). */
export const ROW_RULE = 'row-rule';

/** The viewer's own row: the selected state (accent tint and bar) on a band. */
export const ROW_BAND_OWN = 'row-band row-band-own';

/** A marked row (a selected fight): the selected state (accent tint and bar) on a band. */
export const ROW_BAND_MARK = 'row-band row-band-mark';

/** A row the address may name (`#id`): the marked look while it is the target. */
export const ROW_BAND_TARGET = 'row-band row-band-target';

/** A link row's hover and press ground, on a band. */
export const ROW_BAND_HOVER = 'row-band row-band-hover';

/** A list's own side padding in the same steps: the class an unlit row keeps. */
const ROW_PAD = { 0: '', 1: 'px-1', 2: 'px-2' } as const;

/** A row's side padding and band in a list padded `pad` steps: lit (the viewer's own row, a marked
 *  fight), the band and its reach; unlit, the list's padding. The row's class string carries no
 *  `px-*` of its own, so the two never compete. */
export function rowBand(band: string | null, pad: keyof typeof ROW_BAND_REACH): string {
  return band === null ? ROW_PAD[pad] : `${band} ${ROW_BAND_REACH[pad]}`;
}

/** The viewer's own row, or the list's padding: `rowBand` with the own band. */
export function ownRowBand(own: boolean, pad: keyof typeof ROW_BAND_REACH): string {
  return rowBand(own ? ROW_BAND_OWN : null, pad);
}

/** The band's own content inset (12 px), as a row class. */
const BAND_INSET = 'px-3';

/** The viewer's own row in a list that must not grow past its edges (a grid column at the page's
 *  content edge, guild round 3: the raid roster's band ran 4 px past 1344 at 1440). Every row,
 *  lit or not, holds its content 12 px in, so the band fills the column and never reaches out. */
export function ownRowBandInside(own: boolean): string {
  return own ? `${ROW_BAND_OWN} ${BAND_INSET}` : BAND_INSET;
}

/** The label, sub and value classes of a row's text. */
export const ROW_LABEL = 'text-[15px] font-semibold [overflow-wrap:anywhere]';
export const ROW_SUB = 'text-muted text-[13px]';
