// web/src/lib/chart/types.ts
// The shapes the shared chart (components/chart) draws. The report's adapters (lib/chart/rate.ts,
// lib/chart/threat.ts, lib/chart/events.ts) build them from a summary; the chart never reads a
// summary itself.

/** What a line is: a player, a median reference, the one gold area, or an extra line beside it. */
export type SeriesKind = 'player' | 'median' | 'single' | 'extra';

/** One drawn line, one value per whole second of the fight. `null` is a gap (a dead player). */
export interface ChartSeries {
  /** A unit GUID, `median:raid`, `median:<class>` or another stable key. */
  id: string;
  label: string;
  /** A CSS colour expression, a token (`var(--color-class-mage)`) and never a hex. */
  colour: string;
  kind: SeriesKind;
  values: readonly (number | null)[];
  /** The window's figure the legend shows: a rate, or the standing at the window's end. */
  figure: number;
  /** Always drawn: the legend chip carries the pin and cannot be turned off. */
  pinned?: boolean;
  /** The viewer's own character. */
  you?: boolean;
  /** When the player died, so the legend shows the value at death and the line ends there. */
  diedAtMs?: number;
  /** A canvas dash pattern; medians are 5/4. */
  dash?: readonly number[];
}

/** The kinds of moment a chart marks (spec §B5.6). */
export type ChartEventKind = 'death' | 'target-change' | 'taunt' | 'cast' | 'phase';

/** One marked moment: a lane marker, a plot line, and a line in the tooltip's event list. */
export interface ChartEvent {
  key: string;
  kind: ChartEventKind;
  atMs: number;
  /** The tooltip's line (spec §B8). */
  text: string;
  /** The phone event line's shorter reading; the tooltip's when absent. */
  shortText?: string;
  /** A death's row in a collapsed cluster: the name and the spell, without the source. */
  clusterText?: string;
  /** The link a death opens on the Deaths tab. */
  href?: string;
  /** A spell icon's url, when the build has one. */
  iconUrl?: string | null;
}

/** A named stretch of the fight. */
export interface ChartPhase {
  name: string;
  startMs: number;
  endMs: number;
}

/** One entry of the Players control and the legend sheet: everyone who could be drawn. */
export interface PlayerEntry {
  id: string;
  label: string;
  colour: string;
  className?: string;
  role: 'tank' | 'healer' | 'dps' | 'other';
  figure: number;
  drawn: boolean;
  pinned: boolean;
  you: boolean;
}
