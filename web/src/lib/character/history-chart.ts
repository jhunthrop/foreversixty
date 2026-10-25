// web/src/lib/character/history-chart.ts
// The History section's view model as pure functions (experience spec
// 2026-10-24-character-history 3.3 to 3.6, 7 and 8): the API payload and a measured size in,
// pixel geometry, readouts and notes out. The component draws; nothing here touches the DOM.
import { utcDayMonthYear } from '../dates';
import { difficultyLabel } from '../difficulty';
import { percentileToken } from '../report/format';
import { SCORE_MAX, TIER_FROM, shortDate, tierColor, tierLabel, tierTextColor } from '../score/score';
import type { Tier } from '../score/score';
import {
  type CharacterHistory,
  type HistoryEncounter,
  type HistoryGap,
  type HistoryScorePoint,
  type HistoryVisibility,
} from './history-api';
import { fightKey } from './fight-key';
import { historyCopy } from './history-copy';
import { wrapReason } from './history-layout';
import {
  type AxisTick,
  type ChartSize,
  type Domain,
  type PlotRect,
  axisTicks,
  parseDomain,
  parseTicks,
  plotRect,
  scoreDomain,
  scoreTicks,
  yScale,
} from './history-scales';

const MIN_BAND_LABEL_HEIGHT = 22;
const MIN_GAP_LABEL_WIDTH = 72;
const DOTS_ALL_UP_TO = 40;
const COMMON_BAND_OPACITY = 0.04;
const BAND_OPACITY = 0.08;
const PATH_DECIMALS = 2;
/** Distance between the baselines of a gap reason's two lines. */
const GAP_LABEL_LINE_PITCH = 14;
const TIER_BOUNDARY_VALUES: ReadonlySet<number> = new Set(Object.values(TIER_FROM));

/** No points, exactly one (no chart box, a sentence), or a line of two or more. */
export type ChartMode = 'empty' | 'single' | 'line';

function modeFor(count: number): ChartMode {
  return count === 0 ? 'empty' : count === 1 ? 'single' : 'line';
}

// ---------------------------------------------------------------- score, one point a day

/** One UTC date of the score chart: the day's last rated fight, and every fight of the day. */
export interface ScoreDay {
  /** YYYY-MM-DD, UTC. */
  date: string;
  /** Midnight UTC of `date`, in ms. */
  time: number;
  score: number;
  tier: string;
  visibility: HistoryVisibility;
  /** How many rated fights the day held. */
  fightCount: number;
  /** The last fight's key. */
  key: string;
  /** Every fight of the day, for the list marks. */
  keys: readonly string[];
}

function utcDate(at: string): string {
  return new Date(at).toISOString().slice(0, 10);
}

function dayStart(date: string): number {
  return Date.parse(`${date}T00:00:00Z`);
}

/** Collapses the API's every-rated-fight list to one point per UTC date, keeping the last fight. */
export function collapseScorePoints(points: readonly HistoryScorePoint[]): ScoreDay[] {
  const byDate = new Map<string, HistoryScorePoint[]>();
  for (const point of [...points].sort((a, b) => Date.parse(a.at) - Date.parse(b.at))) {
    const date = utcDate(point.at);
    byDate.set(date, [...(byDate.get(date) ?? []), point]);
  }
  return [...byDate.entries()].map(([date, fights]) => {
    const last = fights[fights.length - 1];
    return {
      date,
      time: dayStart(date),
      score: last.score,
      tier: last.tier,
      visibility: last.visibility,
      fightCount: fights.length,
      key: fightKey(last.report_id, last.fight_index),
      keys: fights.map((fight) => fightKey(fight.report_id, fight.fight_index)),
    };
  });
}

// ---------------------------------------------------------------- geometry

export interface ChartDot {
  x: number;
  y: number;
}

export interface YTick {
  value: number;
  y: number;
  /** A tier's lower bound on a chart with tier bands: drawn by the band's own edge, not as grid. */
  boundary: boolean;
}

export interface XTick extends Omit<AxisTick, 'time' | 'x'> {
  /** Pixels from the chart's left edge. */
  x: number;
  label: string;
}

export interface TierBand {
  tier: Tier;
  top: number;
  bottom: number;
  color: string | null;
  opacity: number;
  /** Tall enough (22 px) to carry its name. */
  labelled: boolean;
}

export interface GapLabelLine {
  text: string;
  y: number;
}

export interface GapBand {
  left: number;
  right: number;
  reason: string;
  /** Wide enough (72 px) for the reason inside; otherwise it goes under the chart. */
  labelInside: boolean;
  /** The reason's lines with their baseline y, centred on the band; none when it goes under the chart. */
  labelLines: GapLabelLine[];
  /** "20 Jul" and "9 Aug". */
  from: string;
  to: string;
}

export interface ChartGeometry {
  plot: PlotRect;
  domain: Domain;
  yTicks: YTick[];
  xTicks: XTick[];
  /** One per data point, oldest first. */
  dots: ChartDot[];
  /** One SVG path per unbroken run of two or more points. */
  runs: string[];
  bands: TierBand[];
  gaps: GapBand[];
}

interface Series {
  times: readonly number[];
  values: readonly number[];
  domain: Domain;
  yTicks: readonly number[];
  gaps: readonly HistoryGap[];
  bands: boolean;
}

function round(value: number): number {
  return Number(value.toFixed(PATH_DECIMALS));
}

function pathOf(dots: readonly ChartDot[]): string {
  return dots.map((dot, index) => `${index === 0 ? 'M' : 'L'}${round(dot.x)} ${round(dot.y)}`).join(' ');
}

/** The points split where an API gap lies between two neighbours: no segment crosses a gap. */
function unbrokenRuns(
  dots: readonly ChartDot[],
  times: readonly number[],
  gaps: readonly HistoryGap[],
): ChartDot[][] {
  const runs: ChartDot[][] = [];
  let current: ChartDot[] = [];
  dots.forEach((dot, index) => {
    current = [...current, dot];
    const next = times[index + 1];
    const broken =
      next !== undefined &&
      gaps.some((gap) => times[index] >= dayStart(gap.from) && next <= dayStart(gap.to));
    if (next === undefined || broken) {
      runs.push(current);
      current = [];
    }
  });
  return runs;
}

function axisLabel(time: number, withYear: boolean): string {
  const date = new Date(time).toISOString().slice(0, 10);
  const label = shortDate(date) ?? date;
  return withYear ? `${label} ${date.slice(0, 4)}` : label;
}

function xTicksFor(times: readonly number[], plot: PlotRect, phone: boolean): XTick[] {
  const first = times[0];
  const last = times[times.length - 1];
  const withYear = new Date(first).getUTCFullYear() !== new Date(last).getUTCFullYear();
  return axisTicks(first, last, plot.width, phone).map((tick) => ({
    x: plot.left + tick.x,
    anchor: tick.anchor,
    label: axisLabel(tick.time, withYear),
  }));
}

function tierBands(domain: Domain, plot: PlotRect): TierBand[] {
  const y = yScale(domain, plot);
  const tiers = Object.keys(TIER_FROM) as Tier[];
  return tiers.flatMap((tier, index) => {
    const from = TIER_FROM[tier];
    const to = index + 1 < tiers.length ? TIER_FROM[tiers[index + 1]] : SCORE_MAX;
    if (to <= domain.min || from >= domain.max) return [];
    const top = y(Math.min(to, domain.max));
    const bottom = y(Math.max(from, domain.min));
    return [
      {
        tier,
        top,
        bottom,
        color: tierColor(tier),
        opacity: tier === 'common' ? COMMON_BAND_OPACITY : BAND_OPACITY,
        labelled: bottom - top >= MIN_BAND_LABEL_HEIGHT,
      },
    ];
  });
}

function gapLabelLines(reason: string, width: number, plot: PlotRect): GapLabelLine[] {
  const lines = wrapReason(reason, width);
  const middle = plot.top + plot.height / 2;
  return lines.map((text, row) => ({
    text,
    y: middle + (row - (lines.length - 1) / 2) * GAP_LABEL_LINE_PITCH,
  }));
}

function gapBands(gaps: readonly HistoryGap[], times: readonly number[], plot: PlotRect): GapBand[] {
  const first = times[0];
  const last = times[times.length - 1];
  const x = (time: number): number =>
    last === first ? plot.left : plot.left + ((time - first) / (last - first)) * plot.width;
  return gaps
    .filter((gap) => dayStart(gap.to) > first && dayStart(gap.from) < last)
    .map((gap) => {
      const left = x(Math.max(dayStart(gap.from), first));
      const right = x(Math.min(dayStart(gap.to), last));
      const labelInside = right - left >= MIN_GAP_LABEL_WIDTH;
      return {
        left,
        right,
        reason: gap.reason,
        labelInside,
        labelLines: labelInside ? gapLabelLines(gap.reason, right - left, plot) : [],
        from: shortDate(gap.from) ?? gap.from,
        to: shortDate(gap.to) ?? gap.to,
      };
    });
}

function layout(series: Series, size: ChartSize): ChartGeometry {
  const plot = plotRect(size);
  const y = yScale(series.domain, plot);
  const first = series.times[0];
  const span = series.times[series.times.length - 1] - first;
  const dots = series.times.map((time, index) => ({
    x: plot.left + (span === 0 ? 0 : ((time - first) / span) * plot.width),
    y: y(series.values[index]),
  }));
  return {
    plot,
    domain: series.domain,
    yTicks: series.yTicks.map((value) => ({
      value,
      y: y(value),
      boundary: series.bands && TIER_BOUNDARY_VALUES.has(value),
    })),
    xTicks: xTicksFor(series.times, plot, size.phone),
    dots,
    runs: unbrokenRuns(dots, series.times, series.gaps)
      .filter((run) => run.length >= 2)
      .map(pathOf),
    bands: series.bands ? tierBands(series.domain, plot) : [],
    gaps: gapBands(series.gaps, series.times, plot),
  };
}

/** Above 40 points only the selected and the newest are drawn as dots. */
export function drawnDotIndexes(count: number, selected: number): number[] {
  if (count <= DOTS_ALL_UP_TO) return Array.from({ length: count }, (_, index) => index);
  return [...new Set([selected, count - 1])].sort((a, b) => a - b);
}

// ---------------------------------------------------------------- the score chart

export interface ScoreChart {
  mode: ChartMode;
  days: readonly ScoreDay[];
  /** A private or guild point is drawn, so the control row names the rings. */
  hasRings: boolean;
  geometry: ChartGeometry | null;
  /** One muted line per gap too narrow for an inside label. */
  gapNotes: string[];
  /** "Up 1090 since 12 Oct", or null with fewer than two points. */
  change: string | null;
}

function netChange(first: number, last: number, sinceDate: string): string {
  return historyCopy.netChange(last - first, shortDate(sinceDate) ?? sinceDate);
}

export function buildScoreChart(
  history: Pick<CharacterHistory, 'score' | 'gaps'>,
  size: ChartSize,
): ScoreChart {
  const days = collapseScorePoints(history.score);
  const mode = modeFor(days.length);
  const hasRings = days.some((day) => day.visibility === 'private');
  if (mode !== 'line') return { mode, days, hasRings, geometry: null, gapNotes: [], change: null };
  const scores = days.map((day) => day.score);
  const domain = scoreDomain(scores);
  const geometry = layout(
    {
      times: days.map((day) => day.time),
      values: scores,
      domain,
      yTicks: scoreTicks(domain),
      gaps: history.gaps,
      bands: true,
    },
    size,
  );
  return {
    mode,
    days,
    hasRings,
    geometry,
    gapNotes: geometry.gaps
      .filter((gap) => !gap.labelInside)
      .map((gap) => historyCopy.gapNote(gap.from, gap.to, gap.reason)),
    change: netChange(scores[0], scores[scores.length - 1], days[0].date),
  };
}

// ---------------------------------------------------------------- the parse chart

export interface ParseKill {
  at: string;
  time: number;
  parse: number;
  /** The whole number the page writes; `parse` keeps the percentile for geometry and colour. */
  rounded: number;
  key: string;
  /** The percentile colour ladder's token. */
  color: string;
}

export interface ParseChart {
  mode: ChartMode;
  kills: readonly ParseKill[];
  geometry: ChartGeometry | null;
  change: string | null;
}

/** Parse draws no gaps: a boss is not killed every week, and that is not a gap in the record. */
export function buildParseChart(encounter: HistoryEncounter | undefined, size: ChartSize): ParseChart {
  const kills: ParseKill[] = (encounter?.points ?? []).map((kill) => ({
    at: kill.at,
    time: Date.parse(kill.at),
    parse: kill.parse_percentile,
    rounded: Math.round(kill.parse_percentile),
    key: fightKey(kill.report_id, kill.fight_index),
    color: percentileToken(kill.parse_percentile),
  }));
  const mode = modeFor(kills.length);
  // One kill is a point, not a line; the component writes the one-kill sentence instead.
  if (mode !== 'line') return { mode, kills, geometry: null, change: null };
  const values = kills.map((kill) => kill.parse);
  const domain = parseDomain(values);
  return {
    mode,
    kills,
    geometry: layout(
      {
        times: kills.map((kill) => kill.time),
        values,
        domain,
        yTicks: parseTicks(domain),
        gaps: [],
        bands: false,
      },
      size,
    ),
    change: netChange(kills[0].rounded, kills[kills.length - 1].rounded, utcDate(kills[0].at)),
  };
}

// ---------------------------------------------------------------- readouts

export interface Readout {
  /** The big number: "1830" or "88". */
  value: string;
  /** The tier word, or "parse". */
  word: string;
  /** "11 Dec 2026 · 2 fights" or "11 Dec 2026 · kill 5 of 5". */
  detail: string;
  /** The colour of `value` and `word`. */
  color: string | null;
  /** The whole readout on one line: "1830 Epic · 11 Dec 2026 · 2 fights". */
  text: string;
}

export function scoreReadout(day: ScoreDay): Readout {
  const word = tierLabel(day.tier) ?? '';
  const detail = historyCopy.scoreDetail(utcDayMonthYear(day.date), day.fightCount);
  return {
    value: String(day.score),
    word,
    detail,
    color: tierTextColor(day.tier),
    text: historyCopy.scoreReadoutText(day.score, word, detail),
  };
}

export function parseReadout(kills: readonly ParseKill[], index: number): Readout {
  const kill = kills[index];
  const detail = historyCopy.parseDetail(utcDayMonthYear(kill.at), index + 1, kills.length);
  return {
    value: String(kill.rounded),
    word: historyCopy.parseWord,
    detail,
    color: kill.color,
    text: `${kill.rounded} ${historyCopy.parseWord} · ${detail}`,
  };
}

// ---------------------------------------------------------------- the boss picker

export interface BossPill {
  /** `${encounter_id}:${difficulty}`. */
  id: string;
  /** "Golemagg", or "Golemagg · Hard" for a boss killed on two difficulties. */
  label: string;
  count: number;
  /** "Golemagg, 5 kills". */
  accessibleName: string;
  encounter: HistoryEncounter;
}

/** The words before "the": Golemagg for Golemagg the Incinerator. */
export function shortBossName(name: string): string {
  return name.split(' the ')[0];
}

function newestKill(encounter: HistoryEncounter): number {
  return Math.max(...encounter.points.map((point) => Date.parse(point.at)));
}

function newestKillOfBoss(encounters: readonly HistoryEncounter[], encounterId: number): number {
  return Math.max(...encounters.filter((e) => e.encounter_id === encounterId).map(newestKill));
}

/**
 * One pill per boss and difficulty with a kill. Bosses run newest kill first; one boss's
 * difficulties sit together by id ascending, and each writes its word ("Golemagg · 40 Player")
 * only when the boss has kills at two or more ids.
 */
export function bossPills(encounters: readonly HistoryEncounter[]): BossPill[] {
  const killed = encounters.filter((encounter) => encounter.points.length > 0);
  return [...killed]
    .sort(
      (a, b) =>
        newestKillOfBoss(killed, b.encounter_id) - newestKillOfBoss(killed, a.encounter_id) ||
        a.encounter_id - b.encounter_id ||
        a.difficulty - b.difficulty,
    )
    .map((encounter) => {
      const several = killed.filter((other) => other.encounter_id === encounter.encounter_id).length > 1;
      const label = historyCopy.pillBoss(
        shortBossName(encounter.encounter),
        several ? difficultyLabel(encounter.difficulty) : null,
      );
      return {
        id: `${encounter.encounter_id}:${encounter.difficulty}`,
        label,
        count: encounter.points.length,
        accessibleName: historyCopy.pillName(label, encounter.points.length),
        encounter,
      };
    });
}

/** The pill of the newest kill. */
export function defaultBossPill(pills: readonly BossPill[]): BossPill | undefined {
  return pills.reduce<BossPill | undefined>(
    (best, pill) =>
      best === undefined || newestKill(pill.encounter) > newestKill(best.encounter) ? pill : best,
    undefined,
  );
}

// ---------------------------------------------------------------- screen-reader rows

export interface SrRow {
  text: string;
  value?: string;
}

/** The screen-reader table's rows, oldest first, with each gap between its two days. */
export function scoreSrRows(chart: ScoreChart, gaps: readonly HistoryGap[]): SrRow[] {
  return chart.days.flatMap((day, index) => {
    const row: SrRow = { text: utcDayMonthYear(day.date), value: String(day.score) };
    const next = chart.days[index + 1];
    const gap = next === undefined ? undefined : gaps.find((g) => g.from === day.date && g.to === next.date);
    return gap === undefined
      ? [row]
      : [
          row,
          {
            text: historyCopy.srGapRow(
              shortDate(gap.from) ?? gap.from,
              shortDate(gap.to) ?? gap.to,
              gap.reason,
            ),
          },
        ];
  });
}

export function parseSrRows(chart: ParseChart): SrRow[] {
  return chart.kills.map((kill) => ({ text: utcDayMonthYear(kill.at), value: String(kill.rounded) }));
}
