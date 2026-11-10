// web/src/lib/chart/theme.ts
// The one chart theme (round 5, design/specs/2026-11-09-directions.md): every chart on the site
// (the report's time series, the sim timeline and distribution, the character score and parse
// charts) takes its ink from here, so the charts read as one system on the dark card. Faint white
// gridlines and frame, half-strength white tick labels in the body face with tabular numerals, the
// accent blue for the primary series and the selection, gold only for a highlight a player acts on.
// Class, school, parse and outcome colours stay their own tokens.
//
// Every ink names a token and an alpha, never a hex: `inkCss` gives the CSS/SVG string (a
// color-mix over the token), `inkCanvas` resolves it against an element for a canvas, which
// cannot read custom properties.
import { resolveColour, withAlpha } from './palette';

/** A chart colour: a design token and the strength it is drawn at. */
export interface ChartInk {
  readonly token: `--${string}`;
  readonly alpha: number;
}

const ink = (token: `--${string}`, alpha = 1): ChartInk => ({ token, alpha });

/** The chart roles. */
export const CHART_INK = {
  /** Horizontal gridlines inside the plot. */
  grid: ink('--color-strong', 0.08),
  /** The axis line and its ticks. */
  frame: ink('--color-strong', 0.18),
  /** Tick labels and axis captions. */
  label: ink('--color-strong', 0.55),
  /** The primary series (one player's line, a distribution's bars), the selection and crosshair. */
  accent: ink('--color-accent'),
  /** The primary series' filled area under its line. */
  accentArea: ink('--color-accent', 0.18),
  /** A selected window or brushed range. */
  selection: ink('--color-accent', 0.14),
  /** Every second phase: a faint wash under the lines. */
  band: ink('--color-accent', 0.06),
  /** A phase start's dashed line. */
  phase: ink('--color-accent', 0.7),
  /** A highlight a player acts on (a marked cast, the pull line): gold, sparingly. */
  highlight: ink('--color-gold'),
  /** A highlight's faint plot line. */
  highlightLine: ink('--color-gold', 0.35),
  /** A quieter series beside the primary one (a distribution's tails, a muted reference). */
  quiet: ink('--color-strong', 0.22),
  /** A hollow marker's fill and the card a readout sits on. */
  surface: ink('--color-raised'),
  /** The readout's edge: the panel's warm ring. */
  edge: ink('--color-panel-edge'),
  /** Plain ink on the chart. */
  text: ink('--color-text'),
  /** A death's line and marker. */
  death: ink('--color-death'),
  /** A warning wash (a capped resource, an ember note). */
  warn: ink('--color-warn', 0.07),
} as const satisfies Record<string, ChartInk>;

export type ChartRole = keyof typeof CHART_INK;

/** The token as a CSS value: `var(--x)`, or a color-mix at its alpha (CSS and inline SVG). */
export function inkCss(value: ChartInk): string {
  if (value.alpha >= 1) return `var(${value.token})`;
  return `color-mix(in srgb, var(${value.token}) ${Math.round(value.alpha * 100)}%, transparent)`;
}

/** The ink resolved against `element`'s computed tokens, for a canvas. */
export function inkCanvas(element: Element, value: ChartInk): string {
  const resolved = resolveColour(element, `var(${value.token})`);
  return value.alpha >= 1 ? resolved : withAlpha(resolved, value.alpha);
}

/** The primary series' colour as a CSS value (one player's line, a distribution's bars). */
export const CHART_ACCENT = inkCss(CHART_INK.accent);

/** The tick and label size on every chart. */
export const CHART_FONT_PX = 12;

/** The canvas font: the body face (Open Sans) from --font-body, so a chart never names a family. */
export function chartFont(element: Element, px: number = CHART_FONT_PX, weight = 400): string {
  const family = getComputedStyle(element).getPropertyValue('--font-body').trim() || 'sans-serif';
  return `${weight} ${px}px ${family}`;
}
