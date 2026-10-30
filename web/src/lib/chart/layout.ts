// web/src/lib/chart/layout.ts
// The chart frame's fixed heights (controls and charts spec §B5.4, §B7, §B11.2): the shell
// reserves exactly what the loaded chart draws, so loading, empty, error and ready are one height
// and nothing shifts when the chunk arrives.

export interface PlotHeights {
  /** Below 768 px. */
  base: number;
  md: number;
  lg: number;
  xl: number;
}

export const CHART_AXIS_PX = 22;
export const CHART_LANE_PX = 26;
export const CHART_EVENT_LINE_PX = 28;
/** The brush: 6 px above, the strip, 16 px of labels, 8 px below. */
export const BRUSH_STRIP_PX = { fine: 36, coarse: 32 } as const;

export function plotHeightFor(viewport: number, heights: PlotHeights): number {
  if (viewport >= 1440) return heights.xl;
  if (viewport >= 1024) return heights.lg;
  if (viewport >= 768) return heights.md;
  return heights.base;
}

/** The y gutter: 52 px, 40 px below 768. */
export function gutterFor(viewport: number): number {
  return viewport >= 768 ? 52 : 40;
}

export function brushHeight(coarse: boolean): number {
  return 6 + (coarse ? BRUSH_STRIP_PX.coarse : BRUSH_STRIP_PX.fine) + 16 + 8;
}
