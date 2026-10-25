import { describe, expect, it } from 'vitest';
import { characterHistoryFixture } from '../../fixtures/character-history';
import { tierColor } from '../score/score';
import type { HistoryEncounter, HistoryScorePoint } from './history-api';
import {
  bossPills,
  buildParseChart,
  buildScoreChart,
  collapseScorePoints,
  defaultBossPill,
  drawnDotIndexes,
  parseReadout,
  parseSrRows,
  scoreReadout,
  scoreSrRows,
  shortBossName,
} from './history-chart';

const DESKTOP = { width: 588, height: 294, phone: false };
const PHONE = { width: 320, height: 213, phone: true };

function fight(at: string, score: number, overrides: Partial<HistoryScorePoint> = {}): HistoryScorePoint {
  return {
    at,
    score,
    tier: 'rare',
    report_id: `r${at}`,
    fight_index: 1,
    encounter: 'Ragnaros',
    difficulty: 3,
    kill: true,
    visibility: 'public',
    ...overrides,
  };
}

function encounter(name: string, id: number, difficulty: number, ats: string[]): HistoryEncounter {
  return {
    encounter: name,
    encounter_id: id,
    difficulty,
    points: ats.map((at, i) => ({ at, parse_percentile: 50 + i, report_id: `k${id}${i}`, fight_index: i })),
  };
}

describe('collapseScorePoints', () => {
  it('keeps the last fight of each UTC date and counts the fights', () => {
    const days = collapseScorePoints([
      fight('2026-12-11T19:00:00Z', 1790),
      fight('2026-12-11T21:14:03Z', 1830, { tier: 'epic' }),
      fight('2026-12-10T08:00:00Z', 1700),
    ]);
    expect(days.map((d) => [d.date, d.score, d.fightCount])).toEqual([
      ['2026-12-10', 1700, 1],
      ['2026-12-11', 1830, 2],
    ]);
    expect(days[1].tier).toBe('epic');
    expect(days[1].keys).toHaveLength(2);
    expect(days[1].key).toBe('r2026-12-11T21:14:03Z:1');
  });

  it('reads the date in UTC, not local time', () => {
    expect(collapseScorePoints([fight('2026-12-11T23:59:59Z', 1)])[0].date).toBe('2026-12-11');
    expect(collapseScorePoints([fight('2026-12-12T00:00:01Z', 1)])[0].date).toBe('2026-12-12');
  });
});

describe('buildScoreChart', () => {
  const fixture = characterHistoryFixture();

  it('is empty for no points and single for one date', () => {
    expect(buildScoreChart({ score: [], gaps: [] }, DESKTOP).mode).toBe('empty');
    const two = [fight('2026-12-11T10:00:00Z', 1), fight('2026-12-11T11:00:00Z', 2)];
    const single = buildScoreChart({ score: two, gaps: [] }, DESKTOP);
    expect(single.mode).toBe('single');
    expect(single.geometry).toBeNull();
    expect(single.change).toBeNull();
  });

  it('puts x proportional to time, not evenly', () => {
    const chart = buildScoreChart(
      {
        score: [
          fight('2026-10-01T10:00:00Z', 1000),
          fight('2026-10-02T10:00:00Z', 1100),
          fight('2026-10-11T10:00:00Z', 1200),
        ],
        gaps: [],
      },
      DESKTOP,
    );
    const [a, b, c] = chart.geometry!.dots.map((dot) => dot.x);
    expect(b - a).toBeCloseTo((c - a) / 10, 5);
    expect(a).toBe(chart.geometry!.plot.left);
    expect(c).toBeCloseTo(chart.geometry!.plot.left + chart.geometry!.plot.width, 5);
  });

  it('breaks the line at an API gap and draws the band', () => {
    const chart = buildScoreChart(fixture, DESKTOP);
    const geometry = chart.geometry!;
    expect(chart.days).toHaveLength(8);
    expect(geometry.runs).toHaveLength(2);
    const gap = geometry.gaps[0];
    expect(gap).toMatchObject({ reason: 'No rated fights', from: '26 Oct', to: '17 Nov' });
    expect(geometry.dots[2].x).toBeCloseTo(gap.left, 5);
    expect(geometry.dots[3].x).toBeCloseTo(gap.right, 5);
  });

  it('draws one unbroken line without gaps', () => {
    const chart = buildScoreChart({ score: fixture.score, gaps: [] }, DESKTOP);
    expect(chart.geometry!.runs).toHaveLength(1);
    expect(chart.geometry!.runs[0].startsWith('M')).toBe(true);
  });

  it('labels a gap inside its band from 72 px, else writes the line under the chart', () => {
    const wide = buildScoreChart(fixture, { ...DESKTOP, width: 1000 });
    expect(wide.geometry!.gaps[0].labelInside).toBe(true);
    expect(wide.gapNotes).toEqual([]);

    const narrow = buildScoreChart(fixture, { width: 200, height: 133, phone: true });
    expect(narrow.geometry!.gaps[0].labelInside).toBe(false);
    expect(narrow.gapNotes).toEqual(['Gap 26 Oct to 17 Nov: No rated fights']);
  });

  it('draws the tiers inside the domain, labelling only bands 22 px tall', () => {
    const geometry = buildScoreChart(fixture, DESKTOP).geometry!;
    expect(geometry.domain).toEqual({ min: 600, max: 1900 });
    expect(geometry.bands.map((band) => band.tier)).toEqual(['uncommon', 'rare', 'epic']);
    expect(geometry.bands.every((band) => band.labelled)).toBe(true);
    expect(geometry.yTicks.map((tick) => tick.value)).toEqual([1000, 1500]);
    const short = buildScoreChart(fixture, { width: 320, height: 60, phone: true }).geometry!;
    expect(short.bands.some((band) => !band.labelled)).toBe(true);
  });

  it('marks rings when a private or guild point is drawn', () => {
    expect(buildScoreChart(fixture, DESKTOP).hasRings).toBe(true);
    expect(buildScoreChart({ score: [fight('2026-12-01T00:00:00Z', 1)], gaps: [] }, DESKTOP).hasRings).toBe(
      false,
    );
  });

  it('writes the net change from the first date', () => {
    expect(buildScoreChart(fixture, DESKTOP).change).toBe('Up 1090 since 12 Oct');
    const down = buildScoreChart(
      { score: [fight('2026-10-12T00:00:00Z', 900), fight('2026-10-20T00:00:00Z', 860)], gaps: [] },
      DESKTOP,
    );
    expect(down.change).toBe('Down 40 since 12 Oct');
    const flat = buildScoreChart(
      { score: [fight('2026-10-12T00:00:00Z', 900), fight('2026-10-20T00:00:00Z', 900)], gaps: [] },
      DESKTOP,
    );
    expect(flat.change).toBe('No change since 12 Oct');
  });

  it('thins x labels to three on a phone and adds the year across a year boundary', () => {
    expect(buildScoreChart(fixture, PHONE).geometry!.xTicks).toHaveLength(3);
    const labels = buildScoreChart(fixture, DESKTOP).geometry!.xTicks.map((tick) => tick.label);
    expect(labels[0]).toBe('12 Oct');
    expect(labels[labels.length - 1]).toBe('11 Dec');
    const crossing = buildScoreChart(
      { score: [fight('2026-12-01T00:00:00Z', 900), fight('2027-01-20T00:00:00Z', 950)], gaps: [] },
      DESKTOP,
    ).geometry!.xTicks.map((tick) => tick.label);
    expect(crossing[crossing.length - 1]).toBe('20 Jan 2027');
  });
});

describe('drawnDotIndexes', () => {
  it('draws every dot up to 40 points, then only the selected and the newest', () => {
    expect(drawnDotIndexes(40, 3)).toHaveLength(40);
    expect(drawnDotIndexes(41, 3)).toEqual([3, 40]);
    expect(drawnDotIndexes(41, 40)).toEqual([40]);
  });
});

describe('scoreReadout', () => {
  it('reads "1830 Epic · 11 Dec 2026 · 2 fights"', () => {
    const day = buildScoreChart(characterHistoryFixture(), DESKTOP).days.at(-1)!;
    const readout = scoreReadout(day);
    expect(readout.text).toBe('1830 Epic · 11 Dec 2026 · 2 fights');
    expect(readout).toMatchObject({ value: '1830', word: 'Epic', detail: '11 Dec 2026 · 2 fights' });
  });

  it('says "1 fight" for one', () => {
    const day = collapseScorePoints([fight('2026-12-11T10:00:00Z', 1830, { tier: 'epic' })])[0];
    expect(scoreReadout(day).detail).toBe('11 Dec 2026 · 1 fight');
  });
});

describe('buildParseChart', () => {
  const golemagg = characterHistoryFixture().encounters[0];

  it('draws a line from two kills and a point for one', () => {
    const line = buildParseChart(golemagg, DESKTOP);
    expect(line.mode).toBe('line');
    expect(line.geometry!.runs).toHaveLength(1);
    const one = buildParseChart({ ...golemagg, points: golemagg.points.slice(0, 1) }, DESKTOP);
    expect(one.mode).toBe('single');
    expect(one.geometry).toBeNull();
    expect(buildParseChart(undefined, DESKTOP).mode).toBe('empty');
  });

  it('follows the colour ladder for y bounds and ticks, and draws no gaps', () => {
    const geometry = buildParseChart(golemagg, DESKTOP).geometry!;
    expect(geometry.domain).toEqual({ min: 25, max: 100 });
    expect(geometry.yTicks.map((tick) => tick.value)).toEqual([25, 50, 75, 100]);
    expect(geometry.gaps).toEqual([]);
    expect(geometry.bands).toEqual([]);
  });

  it('writes the net change and the kill readout', () => {
    const chart = buildParseChart(golemagg, DESKTOP);
    expect(chart.change).toBe('Up 34 since 17 Nov');
    const readout = parseReadout(chart.kills, 4);
    expect(readout).toMatchObject({ value: '88', word: 'parse', detail: '11 Dec 2026 · kill 5 of 5' });
    expect(readout.color).toBe('var(--color-parse-purple)');
    expect(parseReadout(chart.kills, 1).color).toBe('var(--color-parse-blue)');
  });
});

describe('bossPills', () => {
  it('orders by newest kill and defaults to the newest boss', () => {
    const pills = bossPills([
      encounter('Gehennas', 1, 3, ['2026-10-01T00:00:00Z']),
      encounter('Golemagg the Incinerator', 2, 3, ['2026-11-01T00:00:00Z', '2026-12-01T00:00:00Z']),
    ]);
    expect(pills.map((pill) => pill.label)).toEqual(['Golemagg', 'Gehennas']);
    expect(defaultBossPill(pills)?.label).toBe('Golemagg');
    expect(pills[0]).toMatchObject({ count: 2, accessibleName: 'Golemagg, 2 kills' });
    expect(pills[1].accessibleName).toBe('Gehennas, 1 kill');
  });

  it('writes the difficulty on every pill of a boss killed at two ids, ordered by id', () => {
    const pills = bossPills([
      encounter('Golemagg the Incinerator', 2, 9, ['2026-12-01T00:00:00Z']),
      encounter('Golemagg the Incinerator', 2, 3, ['2026-12-02T00:00:00Z']),
      encounter('Gehennas', 1, 3, ['2026-12-03T00:00:00Z']),
    ]);
    expect(pills.map((pill) => pill.label)).toEqual([
      'Gehennas',
      'Golemagg · 10 Player',
      'Golemagg · 40 Player',
    ]);
    expect(pills[2].accessibleName).toBe('Golemagg · 40 Player, 1 kill');
  });

  it('defaults to the pill of the newest kill', () => {
    const pills = bossPills([
      encounter('Golemagg the Incinerator', 2, 3, ['2026-12-01T00:00:00Z']),
      encounter('Golemagg the Incinerator', 2, 9, ['2026-12-05T00:00:00Z']),
    ]);
    expect(defaultBossPill(pills)?.label).toBe('Golemagg · 40 Player');
  });

  it('shortens a name to the words before "the" and skips bosses with no kills', () => {
    expect(shortBossName('Golemagg the Incinerator')).toBe('Golemagg');
    expect(shortBossName('The Prophet Skeram')).toBe('The Prophet Skeram');
    expect(bossPills([encounter('Nobody', 3, 3, [])])).toEqual([]);
  });
});

describe('screen-reader rows', () => {
  it('lists the days oldest first with each gap as a row', () => {
    const fixture = characterHistoryFixture();
    const rows = scoreSrRows(buildScoreChart(fixture, DESKTOP), fixture.gaps);
    expect(rows[2]).toEqual({ text: '26 Oct 2026', value: '1090' });
    expect(rows[3]).toEqual({ text: 'Gap from 26 Oct to 17 Nov: No rated fights' });
    expect(rows).toHaveLength(9);
  });

  it('lists the kills', () => {
    const chart = buildParseChart(characterHistoryFixture().encounters[0], DESKTOP);
    expect(parseSrRows(chart)[0]).toEqual({ text: '17 Nov 2026', value: '54' });
  });
});

describe('score chart thresholds and bands', () => {
  const flat = (width: number, height: number, phone = false) =>
    buildScoreChart(
      {
        score: [
          fight('2026-10-01T00:00:00Z', 1000),
          fight('2026-10-11T00:00:00Z', 1100),
          fight('2026-10-21T00:00:00Z', 1100),
        ],
        gaps: [{ from: '2026-10-01', to: '2026-10-11', reason: 'No rated fights' }],
      },
      { width, height, phone },
    ).geometry!;

  it('puts the gap reason inside its band at exactly 72 px, not at 71.5', () => {
    expect(flat(198, 200).gaps[0]).toMatchObject({ labelInside: true });
    expect(flat(198, 200).gaps[0].right - flat(198, 200).gaps[0].left).toBe(72);
    expect(flat(197, 200).gaps[0].labelInside).toBe(false);
  });

  it('labels a tier band at exactly 22 px tall, not at 21.5', () => {
    const at22 = flat(300, 82).bands;
    expect(at22.map((b) => b.bottom - b.top)).toEqual([22, 22]);
    expect(at22.every((b) => b.labelled)).toBe(true);
    expect(flat(300, 81).bands.every((b) => !b.labelled)).toBe(true);
  });

  it('draws common at 4 percent and the other tiers at 8 percent in the quality colour', () => {
    const wide = buildScoreChart(
      { score: [fight('2026-10-01T00:00:00Z', 100), fight('2026-10-02T00:00:00Z', 2000)], gaps: [] },
      DESKTOP,
    ).geometry!;
    expect(wide.bands.map((b) => [b.tier, b.opacity])).toEqual([
      ['common', 0.04],
      ['uncommon', 0.08],
      ['rare', 0.08],
      ['epic', 0.08],
      ['legendary', 0.08],
    ]);
    expect(wide.bands.map((b) => b.color)).toEqual(
      ['common', 'uncommon', 'rare', 'epic', 'legendary'].map((tier) => tierColor(tier)),
    );
  });
});
