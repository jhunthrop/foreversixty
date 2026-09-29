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
const VIEWPORT_MARGIN_PX = 8;
// The visual gap between the anchor and the panel (ItemTooltip.svelte's own `mt-2`/`mb-2`,
// 0.5rem) -- named separately from VIEWPORT_MARGIN_PX even though both are 8px today, since
// one is "how far from the anchor" and the other "how far from the viewport edge": two
// different reasons that happen to share a value, not one value with two names.
const TOOLTIP_ANCHOR_GAP_PX = 8;

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
  margin: number = VIEWPORT_MARGIN_PX,
): number {
  const maxLeft = viewportWidth - margin - panelWidth - anchorLeft;
  return Math.min(0, maxLeft);
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

/**
 * DOM wrapper: measures `host` (the element the tooltip panel is mounted into) and the panel
 * itself (already mounted as `host`'s child, at its natural, unclamped size -- the caller
 * always mounts before positioning), then sets the CSS variables ItemTooltip.svelte's own
 * inline styles read for both axes. Called once right after the panel mounts, and again on
 * demand if the host moves -- this only ever needs the DOM, never a re-mount, since the
 * panel's `mount()` props are otherwise a point-in-time snapshot.
 */
export function positionTooltipPanel(host: HTMLElement): void {
  const anchor = host.getBoundingClientRect();
  const viewportWidth = document.documentElement.clientWidth;
  const viewportHeight = document.documentElement.clientHeight;
  const panel = host.querySelector<HTMLElement>('[data-testid="item-tooltip"]');
  const panelRect = panel?.getBoundingClientRect();

  host.style.setProperty(
    '--item-tooltip-shift',
    `${tooltipShiftPx(anchor.left, viewportWidth, panelRect?.width ?? TOOLTIP_PANEL_WIDTH_PX)}px`,
  );

  const { side, maxHeightPx } = tooltipVerticalPlacementFor(
    anchor.top,
    anchor.bottom,
    panelRect?.height ?? 0,
    viewportHeight,
  );
  const below = side === 'below';
  host.style.setProperty('--item-tooltip-top', below ? '100%' : 'auto');
  host.style.setProperty('--item-tooltip-bottom', below ? 'auto' : '100%');
  host.style.setProperty('--item-tooltip-mt', below ? `${TOOLTIP_ANCHOR_GAP_PX}px` : '0px');
  host.style.setProperty('--item-tooltip-mb', below ? '0px' : `${TOOLTIP_ANCHOR_GAP_PX}px`);
  host.style.setProperty('--item-tooltip-max-h', maxHeightPx === null ? 'none' : `${maxHeightPx}px`);
  host.style.setProperty('--item-tooltip-overflow-y', maxHeightPx === null ? 'visible' : 'auto');
}
