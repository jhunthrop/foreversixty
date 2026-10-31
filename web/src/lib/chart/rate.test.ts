// web/src/lib/chart/rate.test.ts
import { describe, expect, it } from 'vitest';
import { buildRateModel, classMedianInfo, RAID_MEDIAN_ID, type RatePlayer } from './rate';

const CLASSES = ['Warrior', 'Rogue', 'Mage', 'Warlock', 'Hunter', 'Priest', 'Druid', 'Shaman'];

/** A 40-player Ragnaros: 340 s, submerged (no damage at all) from 180 s to 270 s. */
function raid(): RatePlayer[] {
  return Array.from({ length: 40 }, (_, index) => {
    const role = index < 2 ? 'tank' : index < 10 ? 'healer' : 'dps';
    const base = 600 - index * 7.3;
    return {
      guid: `Player-${index}`,
      name: `Raider${index}`,
      className: CLASSES[index % CLASSES.length],
      role,
      series: Array.from({ length: 340 }, (_, second) =>
        second >= 180 && second < 270 ? 0 : Math.round(base + ((second * 37 + index * 11) % 50)),
      ),
    } satisfies RatePlayer;
  });
}

const WHOLE = { startMs: 0, endMs: 340_000 };

describe('buildRateModel', () => {
  it('draws the top 4 and the bottom 4 damage dealers by default, plus you and the median', () => {
    const model = buildRateModel({
      metric: 'damage-done',
      players: raid(),
      window: WHOLE,
      durationMs: 340_000,
      pinned: new Set(['Player-20']),
      youGuid: 'Player-20',
      chosen: null,
      compare: { raid: true, className: null },
    });
    // The bottom is taken from the damage dealers: the healers' small damage is not a low performer.
    expect(model.series.filter((line) => line.kind === 'player').map((line) => line.id)).toEqual([
      'Player-0',
      'Player-1',
      'Player-2',
      'Player-3',
      'Player-20',
      'Player-36',
      'Player-37',
      'Player-38',
      'Player-39',
    ]);
  });

  it('draws the top 8 of N with the Top 8 preset, plus you, plus the raid median under them', () => {
    const model = buildRateModel({
      metric: 'damage-done',
      players: raid(),
      window: WHOLE,
      durationMs: 340_000,
      pinned: new Set(['Player-35']),
      youGuid: 'Player-35',
      chosen: null,
      preset: 'top',
      compare: { raid: true, className: null },
    });
    expect(model.total).toBe(40);
    expect(model.series[0].id).toBe(RAID_MEDIAN_ID);
    expect(model.series[0].kind).toBe('median');
    const players = model.series.filter((line) => line.kind === 'player');
    expect(players.map((line) => line.id)).toEqual([
      'Player-0',
      'Player-1',
      'Player-2',
      'Player-3',
      'Player-4',
      'Player-5',
      'Player-6',
      'Player-7',
      'Player-35',
    ]);
    const you = players.at(-1)!;
    expect(you.pinned).toBe(true);
    expect(you.you).toBe(true);
    expect(model.entries.filter((entry) => entry.drawn)).toHaveLength(9);
    expect(model.isDefault).toBe(true);
  });

  it('puts the sum of every player’s rounded figure in the head', () => {
    const model = buildRateModel({
      metric: 'damage-done',
      players: raid(),
      window: WHOLE,
      durationMs: 340_000,
      pinned: new Set(),
      chosen: null,
      compare: { raid: true, className: null },
    });
    const legendSum = model.entries.reduce((sum, entry) => sum + Math.round(entry.figure), 0);
    expect(model.raidFigure).toBe(legendSum);
    expect(model.raidFigure).toBeGreaterThan(0);
  });

  it('draws the submerge flat at zero: no gaps and no slope across it', () => {
    const model = buildRateModel({
      metric: 'damage-done',
      players: raid(),
      window: WHOLE,
      durationMs: 340_000,
      pinned: new Set(),
      chosen: null,
      compare: { raid: true, className: null },
    });
    for (const line of model.series) {
      // 5 s points: 180 s to 270 s is points 36 to 53.
      const submerged = line.values.slice(36, 54);
      expect(
        submerged.every((value) => value === 0),
        line.id,
      ).toBe(true);
    }
  });

  it('takes the raid median over the damage dealers, and the class median over that class', () => {
    const model = buildRateModel({
      metric: 'damage-done',
      players: raid(),
      window: WHOLE,
      durationMs: 340_000,
      pinned: new Set(),
      chosen: null,
      compare: { raid: true, className: 'Shaman' },
    });
    const [raidMedian, classMedian] = model.medians;
    expect(raidMedian.count).toBe(30);
    expect(raidMedian.description).toBe('Median DPS of the 30 damage dealers');
    expect(classMedian.label).toBe('Shaman median');
    expect(classMedian.description).toMatch(/^Median DPS of the \d+ damage-dealing shamans$/);
    expect(model.series[1].id).toBe('median:shaman');
    expect(model.series[1].colour).toBe('var(--color-class-shaman)');
  });

  it('healing takes the healers’ median', () => {
    const model = buildRateModel({
      metric: 'healing',
      players: raid(),
      window: WHOLE,
      durationMs: 340_000,
      pinned: new Set(),
      chosen: null,
      compare: { raid: true, className: null },
    });
    expect(model.medians[0].count).toBe(8);
    expect(model.medians[0].description).toBe('Median HPS of the 8 healers');
  });

  it('draws the chosen players once the Players control changes', () => {
    const model = buildRateModel({
      metric: 'damage-done',
      players: raid(),
      window: WHOLE,
      durationMs: 340_000,
      pinned: new Set(),
      chosen: new Set(['Player-20']),
      compare: { raid: false, className: null },
    });
    expect(model.series.map((line) => line.id)).toEqual(['Player-20']);
    expect(model.isDefault).toBe(false);
  });

  it('ranks by the window’s figure, so a zoom re-orders the legend', () => {
    const players: RatePlayer[] = [
      { guid: 'a', name: 'A', role: 'dps', series: [100, 100, 0, 0] },
      { guid: 'b', name: 'B', role: 'dps', series: [0, 0, 100, 100] },
    ];
    const late = buildRateModel({
      metric: 'damage-done',
      players,
      window: { startMs: 2000, endMs: 4000 },
      durationMs: 340_000,
      pinned: new Set(),
      chosen: null,
      compare: { raid: false, className: null },
    });
    expect(late.series.map((line) => line.id)).toEqual(['b', 'a']);
    expect(late.series[0].figure).toBe(100);
  });

  it('draws one Source player as the gold area', () => {
    const model = buildRateModel({
      metric: 'damage-done',
      players: raid(),
      window: WHOLE,
      durationMs: 340_000,
      pinned: new Set(['Player-12']),
      chosen: new Set(),
      compare: { raid: true, className: null },
      single: 'Player-12',
    });
    const single = model.series.find((line) => line.id === 'Player-12')!;
    expect(single.kind).toBe('single');
    expect(single.colour).toBe('var(--color-gold)');
  });

  it('ends a dead player’s line at their death', () => {
    const model = buildRateModel({
      metric: 'damage-done',
      players: raid(),
      window: WHOLE,
      durationMs: 340_000,
      pinned: new Set(),
      chosen: null,
      compare: { raid: false, className: null },
      deadSpans: new Map([['Player-0', [{ startMs: 126_000, endMs: 340_000 }]]]),
    });
    const line = model.series.find((entry) => entry.id === 'Player-0')!;
    // 5 s points on a 340 s fight: the point at 125 s is the last before the death at 126 s.
    expect(line.values[25]).not.toBeNull();
    expect(line.values[26]).toBeNull();
    expect(line.diedAtMs).toBe(126_000);
  });

  it('draws a long fight as 5 s averages and a 10 s pull every second', () => {
    const long = buildRateModel({
      metric: 'damage-done',
      players: raid(),
      window: WHOLE,
      durationMs: 340_000,
      pinned: new Set(),
      chosen: null,
      compare: { raid: true, className: null },
    });
    expect(long.bucket.seconds).toBe(5);
    expect(long.series.every((line) => line.stepMs === 5000 && line.values.length === 68)).toBe(true);
    const first = raid()[0].series.slice(0, 5);
    expect(long.series.find((line) => line.id === 'Player-0')!.values[0]).toBe(
      first.reduce((sum, value) => sum + value, 0) / 5,
    );

    const pull = raid().map((player) => ({ ...player, series: player.series.slice(0, 10) }));
    const short = buildRateModel({
      metric: 'damage-done',
      players: pull,
      window: { startMs: 0, endMs: 10_000 },
      durationMs: 10_000,
      pinned: new Set(),
      chosen: null,
      compare: { raid: true, className: null },
    });
    expect(short.bucket).toEqual({ seconds: 1, labelFloorS: 2 });
    const player0 = short.series.find((line) => line.id === 'Player-0')!;
    expect(player0.stepMs).toBe(1000);
    expect(player0.values).toEqual(pull[0].series);
  });
});

describe('classMedianInfo', () => {
  it('names and describes the class median the same whether it is on or off', () => {
    const off = classMedianInfo(raid(), 'Shaman', false);
    expect(off.on).toBe(false);
    expect(off.label).toBe('Shaman median');
    expect(off.description).toBe(`Median DPS of the ${off.count} damage-dealing shamans`);
  });
});

describe('the median’s name', () => {
  it('follows the set the chart draws', () => {
    const base = {
      metric: 'damage-done' as const,
      players: raid(),
      window: WHOLE,
      durationMs: 340_000,
      pinned: new Set<string>(),
      chosen: null,
      compare: { raid: true, className: null },
    };
    expect(buildRateModel(base).medians[0].label).toBe('Raid median');
    expect(buildRateModel({ ...base, noun: 'everyone' }).medians[0].label).toBe('Median player');
    const line = buildRateModel({ ...base, noun: 'group' }).series.find(
      (entry) => entry.id === RAID_MEDIAN_ID,
    );
    expect(line?.label).toBe('Group median');
  });
});
