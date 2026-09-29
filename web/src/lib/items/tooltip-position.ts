// web/src/lib/items/tooltip-position.ts
// The tooltip panel's own "keep it inside the viewport" shift (ItemTooltip.svelte's
// `--item-tooltip-shift`, consumed as its `style:left`): shared by every place that mounts
// the panel next to an anchor element -- ItemHover.svelte's own per-item mount and
// BisTooltipHost.svelte's single delegated mount for the /bis pages -- so the edge-aware
// math lives in exactly one place instead of two copies that could drift apart.
const TOOLTIP_PANEL_WIDTH_PX = 288; // w-72, ItemTooltip.svelte's own panel width
const VIEWPORT_MARGIN_PX = 8;

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

/**
 * DOM wrapper: measures `host` (the element the tooltip panel is mounted into) and sets the
 * CSS variable ItemTooltip.svelte reads. Called once right after the panel mounts, and
 * again on demand if the host moves -- this only ever needs the DOM, never a re-mount,
 * since the panel's `mount()` props are otherwise a point-in-time snapshot.
 */
export function positionTooltipPanel(host: HTMLElement): void {
  const anchor = host.getBoundingClientRect();
  const viewport = document.documentElement.clientWidth;
  host.style.setProperty('--item-tooltip-shift', `${tooltipShiftPx(anchor.left, viewport)}px`);
}
