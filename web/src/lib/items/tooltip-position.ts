// web/src/lib/items/tooltip-position.ts
// The tooltip panel's own "keep it inside the viewport" placement (ItemTooltip.svelte's
// `--item-tooltip-shift` and its vertical CSS variables, consumed as its own `style:` props):
// shared by every place that mounts the panel next to an anchor element -- ItemHover.svelte's
// own per-item mount and BisTooltipHost.svelte's single delegated mount for the /bis pages --
// so the edge-aware math lives in exactly one place instead of two copies that could drift
// apart.
//
// Tooltip-polish brief item 1 (owner screenshot 2026-09-29): the panel's own width is no
// longer the fixed 288px `w-72` -- it now clamps between 280px and 360px (full width minus
// 32px on a phone, ItemTooltip.svelte's own `w-[clamp(...)]`), so `TOOLTIP_PANEL_WIDTH_PX`
// below is `tooltipShiftPx`'s pure-function default/test fixture only; `positionTooltipPanel`
// measures the panel's own real rendered width instead of assuming a constant.
const TOOLTIP_PANEL_WIDTH_PX = 288;

/** The panel's own placement, read from the variables positionTooltipPanel sets on the host: the
 *  `style` both tooltip panels (ItemTooltip, GameTooltip) hand their Overlay. */
export const TOOLTIP_PANEL_PLACEMENT = [
  'position: var(--item-tooltip-position, absolute)',
  'left: var(--item-tooltip-shift, 0px)',
  'top: var(--item-tooltip-top, 100%)',
  'bottom: var(--item-tooltip-bottom, auto)',
  'margin-top: var(--item-tooltip-mt, 8px)',
  'margin-bottom: var(--item-tooltip-mb, 0px)',
  'max-height: var(--item-tooltip-max-h, none)',
  'overflow-y: var(--item-tooltip-overflow-y, visible)',
].join('; ');
// The site's own 16 px side gutter: a panel never sits closer to a screen edge than the page text does.
const SIDE_MARGIN_PX = 16;
// Above and below the anchor, the panel may come this close to the top or bottom edge.
const VIEWPORT_MARGIN_PX = 8;
// The visual gap between the anchor and the panel (ItemTooltip.svelte's own `mt-2`/`mb-2`,
// 0.5rem) -- named separately from VIEWPORT_MARGIN_PX even though both are 8px today, since
// one is "how far from the anchor" and the other "how far from the viewport edge": two
// different reasons that happen to share a value, not one value with two names.
const TOOLTIP_ANCHOR_GAP_PX = 8;
/** ItemTooltip.svelte's phone width is `100vw - 32px`: a 16 px gutter each side. */
const PHONE_GUTTER_PX = 16;

/**
 * Pure: the CSS shift (always <= 0) that keeps a `panelWidth`-wide panel hung from
 * `anchorLeft` inside a `viewportWidth`-wide viewport -- near the right edge of a phone,
 * hanging the panel from the anchor's left edge would otherwise push it off-screen and
 * scroll the page sideways. `0` (the panel's own unshifted default) whenever there is
 * already room.
 */
export function tooltipShiftPx(
  anchorLeft: number,
  viewportWidth: number,
  panelWidth: number = TOOLTIP_PANEL_WIDTH_PX,
  margin: number = SIDE_MARGIN_PX,
): number {
  const maxLeft = viewportWidth - margin - panelWidth - anchorLeft;
  return Math.min(0, maxLeft);
}

/** Which anchor edge the panel's matching edge meets (paperdoll spec §4.2): `start`, left to
 *  left (the default, every list and the paperdoll's left column); `end`, right to right (the
 *  paperdoll's right column, so the panel opens toward the centre, never over the far column). */
export type TooltipAlign = 'start' | 'end';

/**
 * Pure: the panel's left offset from its anchor's left edge. `start` is `tooltipShiftPx`;
 * `end` meets the panel's right edge with the anchor's, then keeps it `margin` inside both
 * viewport edges.
 */
export function tooltipLeftPx(
  align: TooltipAlign,
  anchorLeft: number,
  anchorWidth: number,
  viewportWidth: number,
  panelWidth: number = TOOLTIP_PANEL_WIDTH_PX,
  margin: number = SIDE_MARGIN_PX,
): number {
  // On a phone the panel is the viewport minus two gutters: it spans exactly those gutters,
  // whichever cell or column was tapped (paperdoll spec §4.2 rule 5).
  if (viewportWidth - panelWidth <= 2 * PHONE_GUTTER_PX) return PHONE_GUTTER_PX - anchorLeft;
  if (align === 'start') return tooltipShiftPx(anchorLeft, viewportWidth, panelWidth, margin);
  const maxLeft = viewportWidth - margin - panelWidth - anchorLeft;
  return Math.min(maxLeft, Math.max(margin - anchorLeft, anchorWidth - panelWidth));
}

/** How far a pointer-following tooltip (the chart's crosshair) sits beside the pointer. */
const POINTER_GAP_PX = 16;

/**
 * Pure: a pointer-following tooltip's viewport position (the chart crosshair's, controls and
 * charts spec §B5.7): `POINTER_GAP_PX` right of `x`, or left of it when `flip`, its top at `y`,
 * clamped `VIEWPORT_MARGIN_PX` inside the viewport on every side.
 */
export function pointerTipPosition(
  pointer: { x: number; y: number; flip: boolean },
  panel: { width: number; height: number },
  viewport: { width: number; height: number },
): { left: number; top: number } {
  const wanted = pointer.flip ? pointer.x - POINTER_GAP_PX - panel.width : pointer.x + POINTER_GAP_PX;
  return {
    left: Math.max(VIEWPORT_MARGIN_PX, Math.min(viewport.width - panel.width - VIEWPORT_MARGIN_PX, wanted)),
    top: Math.max(
      VIEWPORT_MARGIN_PX,
      Math.min(viewport.height - panel.height - VIEWPORT_MARGIN_PX, pointer.y),
    ),
  };
}

export type TooltipVerticalSide = 'below' | 'above';

export interface TooltipVerticalPlacement {
  side: TooltipVerticalSide;
  /** `null` when the panel fits at full height on its chosen side -- no clamp needed. A
   *  number when neither side has room for the panel at full height: the larger of the two,
   *  so the panel scrolls internally rather than spilling past the viewport either way. */
  maxHeightPx: number | null;
}

/**
 * Pure: where a `panelHeight`-tall panel should open relative to its anchor inside a
 * `viewportHeight`-tall viewport (tooltip-polish brief item 5 -- "never extends below the
 * viewport; flips above the anchor or clamps and, if still taller than the viewport, scrolls
 * internally"). Prefers `below` (the panel's own unshifted default, matching every other
 * popover on the site); flips to `above` only when `below` cannot fit the panel at full
 * height but `above` can; when NEITHER side fits, picks whichever has more room and clamps
 * to it, so the caller can turn on internal scrolling instead of ever letting the panel spill
 * past the viewport.
 */
export function tooltipVerticalPlacementFor(
  anchorTop: number,
  anchorBottom: number,
  panelHeight: number,
  viewportHeight: number,
  gap: number = TOOLTIP_ANCHOR_GAP_PX,
  margin: number = VIEWPORT_MARGIN_PX,
): TooltipVerticalPlacement {
  const spaceBelow = viewportHeight - anchorBottom - gap - margin;
  const spaceAbove = anchorTop - gap - margin;
  if (panelHeight <= spaceBelow) return { side: 'below', maxHeightPx: null };
  if (panelHeight <= spaceAbove) return { side: 'above', maxHeightPx: null };
  const side: TooltipVerticalSide = spaceAbove > spaceBelow ? 'above' : 'below';
  return { side, maxHeightPx: Math.max(0, Math.max(spaceAbove, spaceBelow)) };
}

/** Whether a box between `host` and the page body cuts off what overflows it (a table that
 *  scrolls sideways on a phone). The body's own sideways clip is the page's, not a box's. */
export function hasClippingAncestor(host: HTMLElement): boolean {
  for (let node = host.parentElement; node !== null && node !== document.body; node = node.parentElement) {
    const style = getComputedStyle(node);
    if (style.overflowX !== 'visible' || style.overflowY !== 'visible') return true;
  }
  return false;
}

/**
 * DOM wrapper: measures `host` (the element the tooltip panel is mounted into) and the panel
 * itself (already mounted as `host`'s child, at its natural, unclamped size -- the caller
 * always mounts before positioning), then sets the CSS variables ItemTooltip.svelte's own
 * inline styles read for both axes. Called once right after the panel mounts, and again on
 * demand if the host moves -- this only ever needs the DOM, never a re-mount, since the
 * panel's `mount()` props are otherwise a point-in-time snapshot.
 */
export function positionTooltipPanel(host: HTMLElement, align: TooltipAlign = 'start'): void {
  const anchor = host.getBoundingClientRect();
  const viewportWidth = document.documentElement.clientWidth;
  const viewportHeight = document.documentElement.clientHeight;
  const panel = host.querySelector<HTMLElement>('[data-testid="item-tooltip"], [data-testid="game-tooltip"]');
  const panelRect = panel?.getBoundingClientRect();
  if (host.dataset.tipBeside !== undefined && placeBeside(host, anchor, panelRect)) return;

  const shift = tooltipLeftPx(
    align,
    anchor.left,
    anchor.width,
    viewportWidth,
    panelRect?.width ?? TOOLTIP_PANEL_WIDTH_PX,
  );
  const { side, maxHeightPx } = tooltipVerticalPlacementFor(
    anchor.top,
    anchor.bottom,
    panelRect?.height ?? 0,
    viewportHeight,
  );
  const below = side === 'below';
  // Inside a scrolling box (a guide's table) an absolute panel is cut off at the box's edge,
  // so it is placed on the viewport instead, at the same spot beside its anchor.
  const fixed = hasClippingAncestor(host);
  host.style.setProperty('--item-tooltip-position', fixed ? 'fixed' : 'absolute');
  host.style.setProperty('--item-tooltip-shift', `${fixed ? anchor.left + shift : shift}px`);
  const top = fixed ? `${anchor.bottom}px` : '100%';
  const bottom = fixed ? `${viewportHeight - anchor.top}px` : '100%';
  host.style.setProperty('--item-tooltip-top', below ? top : 'auto');
  host.style.setProperty('--item-tooltip-bottom', below ? 'auto' : bottom);
  host.style.setProperty('--item-tooltip-mt', below ? `${TOOLTIP_ANCHOR_GAP_PX}px` : '0px');
  host.style.setProperty('--item-tooltip-mb', below ? '0px' : `${TOOLTIP_ANCHOR_GAP_PX}px`);
  host.style.setProperty('--item-tooltip-max-h', maxHeightPx === null ? 'none' : `${maxHeightPx}px`);
  host.style.setProperty('--item-tooltip-overflow-y', maxHeightPx === null ? 'visible' : 'auto');
}

/** The side a talent's tooltip opens on, as the client draws it: beside the cell (right, else
 *  left), so it never covers the tiles around it. Null when neither side has room (a phone),
 *  where the tooltip falls back to below or above the cell. */
export type TalentTipSide = 'right' | 'left';

export function talentTipSide(
  anchorLeft: number,
  anchorRight: number,
  tipWidth: number,
  viewportWidth: number,
): TalentTipSide | null {
  if (anchorRight + TOOLTIP_ANCHOR_GAP_PX + tipWidth <= viewportWidth - SIDE_MARGIN_PX) return 'right';
  if (anchorLeft - TOOLTIP_ANCHOR_GAP_PX - tipWidth >= SIDE_MARGIN_PX) return 'left';
  return null;
}

/** The tooltip's left offset from its cell when it opens beside it. */
export function talentTipSideLeft(side: TalentTipSide, anchorWidth: number, tipWidth: number): number {
  return side === 'right' ? anchorWidth + TOOLTIP_ANCHOR_GAP_PX : -(tipWidth + TOOLTIP_ANCHOR_GAP_PX);
}

/**
 * A host marked `data-tip-beside` (every talent cell and Point order socket): the panel opens
 * beside the cell, its top on the cell's top, the way the client's talent frame draws it. False,
 * placing nothing, when neither side has room or the panel would run past the viewport's foot;
 * the caller then places it below or above.
 */
function placeBeside(host: HTMLElement, anchor: DOMRect, panelRect: DOMRect | undefined): boolean {
  const width = panelRect?.width ?? TOOLTIP_PANEL_WIDTH_PX;
  const side = talentTipSide(anchor.left, anchor.right, width, document.documentElement.clientWidth);
  const height = panelRect?.height ?? 0;
  if (side === null || anchor.top + height > document.documentElement.clientHeight - VIEWPORT_MARGIN_PX) {
    return false;
  }
  const left = talentTipSideLeft(side, anchor.width, width);
  const fixed = hasClippingAncestor(host);
  host.style.setProperty('--item-tooltip-position', fixed ? 'fixed' : 'absolute');
  host.style.setProperty('--item-tooltip-shift', `${fixed ? anchor.left + left : left}px`);
  host.style.setProperty('--item-tooltip-top', fixed ? `${anchor.top}px` : '0px');
  host.style.setProperty('--item-tooltip-bottom', 'auto');
  host.style.setProperty('--item-tooltip-mt', '0px');
  host.style.setProperty('--item-tooltip-mb', '0px');
  host.style.setProperty('--item-tooltip-max-h', 'none');
  host.style.setProperty('--item-tooltip-overflow-y', 'visible');
  return true;
}
