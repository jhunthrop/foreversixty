// web/src/lib/ui/icon.ts
// The one set of <svg> attributes an icon from lib/ui/ui-icons.ts is drawn with, shared by
// ui/Icon.svelte, ui/Icon.astro and the listbox's DOM builder, so a line glyph's stroke and a
// custom icon's fill never differ between the three. Decorative (aria-hidden): the control that
// holds an icon carries its name.
import type { UiIconShape } from './ui-icons';

/** The line family's stroke on its 24 px grid (design/assets/icons/ui/line). */
export const LINE_STROKE_WIDTH = 2;

export function iconAttributes(icon: UiIconShape, size: number): Record<string, string> {
  const paint: Record<string, string> =
    icon.kind === 'fill'
      ? { fill: 'currentColor' }
      : {
          fill: 'none',
          stroke: 'currentColor',
          'stroke-width': String(LINE_STROKE_WIDTH),
          'stroke-linecap': 'round',
          'stroke-linejoin': 'round',
        };
  return {
    width: String(size),
    height: String(size),
    viewBox: `0 0 ${icon.box} ${icon.box}`,
    ...paint,
    'aria-hidden': 'true',
    'data-icon': icon.name,
  };
}

/**
 * The icon as a CSS mask image, for a glyph only a stylesheet can draw (a browser's own control,
 * such as the search field's clear button): the same shape, painted by the element's background.
 */
export function iconMaskUrl(icon: UiIconShape): string {
  const paint =
    icon.kind === 'fill'
      ? "fill='black'"
      : `fill='none' stroke='black' stroke-width='${LINE_STROKE_WIDTH}' stroke-linecap='round' stroke-linejoin='round'`;
  const svg = `<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 ${icon.box} ${icon.box}' ${paint}>${icon.body.replaceAll('"', "'")}</svg>`;
  return `url("data:image/svg+xml,${encodeURIComponent(svg)}")`;
}
