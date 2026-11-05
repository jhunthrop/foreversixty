// web/src/lib/sim/timeline.test.ts
// The timeline's rules on small hand-made results, and the fixture only where a rule is about the
// shape of a real sample, never about its figures: lane E regenerates that fixture from a real run.
import { describe, expect, it } from 'vitest';
import fullFixture from '../../fixtures/sim/result-full.json';
import olderFixture from '../../fixtures/sim/result.json';
import type { ActionNames } from './action-names';
import {
  baseActionKey,
  buildTimeline,
  caretAt,
  clipSegments,
  closeUpWindow,
  layoutRows,
  MAX_EVENTS,
  mergeTicks,
  MIN_WINDOW_MS,
  moveCaret,
  panWindow,
  preciseClock,
  pressesInView,
  readoutAt,
  resizeWindow,
  resourceAt,
  stacksAt,
  tickFollow,
  tickNear,
  windowIncluding,
  type TickLane,
  type TimelineEvent,
  type TimelineModel,
} from './timeline';
import type { SampleAura, SampleCast, SimResult } from './types';

const NAMES: ActionNames = {
  spell: {
    '20662': 'Execute',
    '11567': 'Heroic Strike',
    '12328': 'Death Wish',
    '12867': 'Deep Wounds',
    '12319': 'Flurry',
    '11585': 'Overpower',
    '11597': 'Sunder Armor',
    '2687': 'Bloodrage',
  },
  item: {},
};

function result(
  sample: SampleCast[],
  extra: Partial<SimResult> = {},
  abilities: Partial<SimResult['summary']['damage_done'][number]['abilities'][number]>[] = [],
  casts: { spell_name: string; succeeded: number }[] = [],
): SimResult {
  return {
    sample,
    summary: {
      damage_done: [{ abilities }],
      casts,
    },
    request: { character: { gear: [] } },
    ...extra,
  } as unknown as SimResult;
}

const cast = (at_ms: number, action: string, rage?: number, more: Partial<SampleCast> = {}): SampleCast => ({
  at_ms,
  action,
  ...(rage === undefined ? {} : { resources: { rage } }),
  ...more,
});

function built(input: SimResult): TimelineModel {
  const model = buildTimeline(input, NAMES);
  if (model === null) throw new Error('expected a timeline');
  return model;
}

describe('baseActionKey', () => {
  it('folds a tag and a rank into the base spell', () => {
    expect(baseActionKey('spell:11567/1')).toBe('spell:11567');
    expect(baseActionKey('spell:25286/1+r3')).toBe('spell:25286');
    expect(baseActionKey('item:13442')).toBe('item:13442');
  });

  it('keeps an other key whole: a swing tag names a hand', () => {
    expect(baseActionKey('other:attack/2')).toBe('other:attack/2');
  });
});

describe('buildTimeline: axis', () => {
  it('runs from the first event to the sample length', () => {
    const model = built(
      result([cast(-2500, 'spell:2687'), cast(1000, 'spell:20662')], { sample_duration_ms: 90_000 }),
    );
    expect(model.startMs).toBe(-2500);
    expect(model.endMs).toBe(90_000);
  });

  it('ends a second after the last event on a result without the length', () => {
    const model = built(result([cast(0, 'spell:2687'), cast(5000, 'spell:20662')]));
    expect(model.endMs).toBe(6000);
  });

  it('has no timeline at all without a sample', () => {
    expect(buildTimeline(result([]), NAMES)).toBeNull();
    expect(buildTimeline(olderFixture as unknown as SimResult, NAMES)).toBeNull();
  });

  it('caps the events it reads', () => {
    const sample = Array.from({ length: MAX_EVENTS + 20 }, (_, index) => cast(index, 'spell:20662'));
    expect(built(result(sample)).casts[0]?.events).toHaveLength(MAX_EVENTS);
  });
});

describe('buildTimeline: lanes', () => {
  it('orders Casts by first use and folds a queue press into its base as hollow', () => {
    const model = built(
      result([
        cast(0, 'spell:2687'),
        cast(100, 'spell:11567/1'),
        cast(900, 'spell:11567', 20, { outcome: 'hit', amount: 500 }),
        cast(1000, 'spell:20662'),
      ]),
    );
    expect(model.casts.map((lane) => lane.name)).toEqual(['Bloodrage', 'Heroic Strike', 'Execute']);
    const strike = model.casts[1]!;
    expect(strike.events.map((event) => event.atMs)).toEqual([900]);
    expect(strike.queued.map((event) => event.atMs)).toEqual([100]);
  });

  it('counts a queued spell by its presses, never presses and landings', () => {
    const model = built(
      result([cast(100, 'spell:11567/1'), cast(900, 'spell:11567'), cast(1000, 'spell:20662')]),
    );
    expect(model.presses.map((press) => press.action)).toEqual(['spell:11567/1', 'spell:20662']);
  });

  it('treats a tagged key with no landing as the landing itself', () => {
    const model = built(result([cast(100, 'spell:20662/1')]));
    expect(model.casts[0]?.events).toHaveLength(1);
    expect(model.casts[0]?.queued).toHaveLength(0);
  });

  it('puts swings and a periodic row in Swings and ticks, named for the player', () => {
    const model = built(
      result(
        [
          cast(0, 'other:attack/1'),
          cast(1, 'other:attack/2'),
          cast(2, 'other:attack/3'),
          cast(3, 'spell:12867', undefined, { outcome: 'tick', amount: 141 }),
          cast(4, 'other:rage_gain'),
        ],
        {},
        [{ name: 'spell:12867', ticks: 59, total: 100, school: 1 }],
      ),
    );
    expect(model.swings.map((lane) => lane.name)).toEqual([
      'Main hand',
      'Off hand',
      'Extra attacks',
      'Deep Wounds ticks',
    ]);
    expect(model.casts).toHaveLength(0);
    expect(model.periodicName).toBe('Deep Wounds');
  });

  it('orders the hands by hand, whichever swung first, then the periodic lanes', () => {
    const model = built(
      result(
        [
          cast(0, 'spell:12867', undefined, { outcome: 'tick' }),
          cast(1, 'other:attack/2'),
          cast(2, 'other:attack/1'),
        ],
        {},
        [{ name: 'spell:12867', ticks: 1, total: 1, school: 1 }],
      ),
    );
    expect(model.swings.map((lane) => lane.name)).toEqual(['Main hand', 'Off hand', 'Deep Wounds ticks']);
  });

  it('never counts a periodic application as a tick or a press in a run with outcomes', () => {
    const model = built(
      result(
        [cast(0, 'spell:12867'), cast(1000, 'spell:12867', undefined, { outcome: 'tick_crit', amount: 9 })],
        {},
        [{ name: 'spell:12867', ticks: 1, total: 9, school: 1 }],
      ),
    );
    expect(model.swings[0]?.events.map((event) => event.atMs)).toEqual([1000]);
    expect(model.casts).toHaveLength(0);
    expect(model.presses).toHaveLength(0);
  });

  it('keeps an older run periodic rows without calling them ticks', () => {
    const model = built(
      result([cast(0, 'spell:12867')], {}, [{ name: 'spell:12867', ticks: 1, total: 9, school: 1 }]),
    );
    expect(model.swings[0]?.name).toBe('Deep Wounds');
    expect(model.periodicName).toBeNull();
  });

  it('puts a hit nobody pressed on Swings and ticks, the off-hand strike named for its hand', () => {
    const model = built(
      result(
        [
          cast(0, 'spell:1680', 5, { outcome: 'hit' }),
          cast(0, 'spell:462891', 5, { outcome: 'crit', amount: 7 }),
        ],
        {},
        [
          { name: 'spell:1680', total: 9, school: 1 },
          { name: 'spell:462891', total: 7, school: 1 },
        ],
        [{ spell_name: 'spell:1680', succeeded: 1 }],
      ),
    );
    expect(model.casts.map((lane) => lane.action)).toEqual(['spell:1680']);
    expect(model.swings[0]?.action).toBe('spell:462891');
    expect(model.swings[0]?.name).toMatch(/\(off hand\)$/);
    expect(model.presses).toHaveLength(1);
  });

  it('never draws Thrash, a lane or a bar', () => {
    const model = built(
      result([cast(0, 'spell:21919'), cast(1, 'spell:2687')], {
        sample_auras: [{ action: 'spell:21919', on: 'player', segments: [[0, 10]] }],
        sample_duration_ms: 1000,
      }),
    );
    expect(model.casts.map((lane) => lane.action)).toEqual(['spell:2687']);
    expect(model.onYou).toHaveLength(0);
  });

  it('splits a castable DoT: the cast on Casts, its ticks on the ticks lane', () => {
    const model = built(
      result(
        [
          cast(0, 'spell:772', 10, { outcome: 'hit' }),
          cast(3000, 'spell:772', 10, { outcome: 'tick', amount: 9 }),
        ],
        {},
        [{ name: 'spell:772', ticks: 5, total: 45, school: 1 }],
        [{ spell_name: 'spell:772', succeeded: 1 }],
      ),
    );
    expect(model.casts[0]?.events).toHaveLength(1);
    expect(model.swings[0]?.events).toHaveLength(1);
  });

  it('colours a damaging cast by its school and a cooldown by none', () => {
    const model = built(
      result([cast(0, 'spell:20662'), cast(1, 'spell:2687')], {}, [
        { name: 'spell:20662', total: 10, school: 1 },
        { name: 'spell:2687', total: 0 },
      ]),
    );
    expect(model.casts[0]?.schoolToken).toBe('--color-school-physical');
    expect(model.casts[1]?.schoolToken).toBeNull();
  });
});

describe('buildTimeline: bars', () => {
  const auras: SampleAura[] = [
    { action: 'spell:12319', on: 'player', segments: [[5000, 90_000]] },
    { action: 'spell:12328', on: 'player', segments: [[1000, 4000]] },
    { action: 'spell:11597', on: 'target', segments: [[2000, 100_000]], stacks: [[2000, 5]] },
  ];

  it('orders each group by first application', () => {
    const model = built(
      result([cast(0, 'spell:20662')], { sample_auras: auras, sample_duration_ms: 100_000 }),
    );
    expect(model.onYou.map((bar) => bar.name)).toEqual(['Death Wish', 'Flurry']);
    expect(model.onTarget.map((bar) => bar.name)).toEqual(['Sunder Armor']);
    expect(model.hasAuras).toBe(true);
  });

  it('leaves raid buffs and debuffs up all fight to the Buffs card, and a queue window off the bars', () => {
    const model = built(
      result([cast(0, 'spell:11567/1'), cast(5, 'spell:11567')], {
        sample_duration_ms: 1000,
        sample_auras: [
          { action: 'spell:10442', on: 'player', segments: [[0, 1000]] },
          { action: 'spell:12328', on: 'player', segments: [[0, 400]] },
          { action: 'spell:11567/1', on: 'player', segments: [[0, 5]] },
          { action: 'spell:11597', on: 'target', segments: [[0, 1000]] },
        ],
      }),
    );
    expect(model.onYou.map((bar) => bar.action)).toEqual(['spell:12328']);
    expect(model.onTarget).toHaveLength(0);
  });

  it('keeps a debuff applied after the pull, even when it then stays up to the end', () => {
    const model = built(
      result([cast(0, 'spell:20662')], {
        sample_duration_ms: 100_000,
        sample_auras: [{ action: 'spell:12867', on: 'target', segments: [[10, 100_000]] }],
      }),
    );
    expect(model.onTarget.map((bar) => bar.name)).toEqual(['Deep Wounds']);
  });

  it('says so when the result has no aura data', () => {
    expect(built(result([cast(0, 'spell:20662')])).hasAuras).toBe(false);
  });

  it('draws the Overpower window only when the rotation casts Overpower, and never as a buff', () => {
    const window: SampleAura = { action: 'spell:11585', on: 'player', segments: [[0, 5000]] };
    const silent = built(result([cast(0, 'spell:20662')], { sample_auras: [window] }));
    expect(silent.overpower).toBeNull();
    expect(silent.onYou).toHaveLength(0);

    const casting = built(
      result(
        [cast(0, 'spell:11585')],
        { sample_auras: [window] },
        [],
        [{ spell_name: 'spell:11585', succeeded: 3 }],
      ),
    );
    expect(casting.overpower?.name).toBe('Overpower window');
    expect(casting.onYou).toHaveLength(0);
  });
});

describe('buildTimeline: execute and resource', () => {
  it('starts the wash at the first Execute of any rank', () => {
    expect(built(result([cast(0, 'spell:2687'), cast(500, 'spell:20661')])).executeAtMs).toBe(500);
    expect(built(result([cast(0, 'spell:2687')])).executeAtMs).toBeNull();
  });

  it('reads the class pool, holds it between events, and reports its peak and cap', () => {
    const model = built(
      result([cast(0, 'spell:2687', 15), cast(1000, 'spell:20662', 60), cast(2000, 'spell:2687', 0)]),
    );
    expect(model.resource?.key).toBe('rage');
    expect(model.resource?.peak).toBe(60);
    expect(model.resource?.cap).toBe(100);
    expect(resourceAt(model.resource, 1500)).toBe(60);
    expect(resourceAt(model.resource, 2000)).toBe(0);
    expect(resourceAt(model.resource, -1)).toBeNull();
  });

  it('prefers energy over mana and caps mana at its peak', () => {
    const cat = built(result([{ at_ms: 0, action: 'spell:1', resources: { mana: 3000, energy: 80 } }]));
    expect(cat.resource?.key).toBe('energy');
    const mage = built(result([{ at_ms: 0, action: 'spell:1', resources: { mana: 3000 } }]));
    expect(mage.resource?.cap).toBe(3000);
  });

  it('has no line without resources', () => {
    expect(built(result([cast(0, 'spell:2687')])).resource).toBeNull();
  });
});

describe('clocks', () => {
  it('prints a tenth, and a sign before the pull', () => {
    expect(preciseClock(145_000)).toBe('2:25.0');
    expect(preciseClock(-2500)).toBe('-0:02.5');
    expect(preciseClock(59_960)).toBe('1:00.0');
  });
});

describe('zoom window', () => {
  const fight = { startMs: -2500, endMs: 173_950 };

  it('opens at the pull with no crosshair, centred on one when there is', () => {
    expect(closeUpWindow(fight, null)).toEqual({ startMs: 0, endMs: 30_000 });
    expect(closeUpWindow(fight, 100_000)).toEqual({ startMs: 85_000, endMs: 115_000 });
  });

  it('stays inside the fight', () => {
    expect(closeUpWindow(fight, 173_000)).toEqual({ startMs: 143_950, endMs: 173_950 });
    expect(closeUpWindow(fight, -2000)).toEqual({ startMs: -2500, endMs: 27_500 });
    expect(closeUpWindow({ startMs: 0, endMs: 12_000 }, null)).toEqual({ startMs: 0, endMs: 12_000 });
  });

  it('pans with its length kept', () => {
    expect(panWindow({ startMs: 0, endMs: 30_000 }, 5000, fight)).toEqual({ startMs: 5000, endMs: 35_000 });
    expect(panWindow({ startMs: 0, endMs: 30_000 }, -9000, fight)).toEqual({ startMs: -2500, endMs: 27_500 });
  });

  it('resizes an edge, never under the minimum and never out of the fight', () => {
    const start = { startMs: 0, endMs: 30_000 };
    expect(resizeWindow(start, 'end', 5000, fight).endMs).toBe(35_000);
    expect(resizeWindow(start, 'end', -60_000, fight).endMs).toBe(MIN_WINDOW_MS);
    expect(resizeWindow(start, 'start', 60_000, fight).startMs).toBe(30_000 - MIN_WINDOW_MS);
    expect(resizeWindow(start, 'start', -60_000, fight).startMs).toBe(-2500);
  });

  it('pulls itself over a time outside it, and leaves a time inside alone', () => {
    const window = { startMs: 0, endMs: 30_000 };
    expect(windowIncluding(window, 10_000, fight)).toBe(window);
    expect(windowIncluding(window, 100_000, fight)).toEqual({ startMs: 85_000, endMs: 115_000 });
  });
});

const event = (atMs: number, more: Partial<TimelineEvent> = {}): TimelineEvent => ({
  atMs,
  action: 'spell:20662',
  resource: null,
  queued: false,
  ...more,
});

describe('mergeTicks', () => {
  const window = { startMs: 0, endMs: 1000 };

  it('merges ticks under 1 px apart into the first', () => {
    const ticks = mergeTicks([event(0), event(5), event(9), event(100)], window, 0.1);
    expect(ticks.map((tick) => tick.count)).toEqual([3, 1]);
    expect(ticks.map((tick) => tick.index)).toEqual([0, 3]);
  });

  it('keeps ticks a pixel apart, and measures from the merged tick, not the last merged', () => {
    expect(mergeTicks([event(0), event(10)], window, 0.1)).toHaveLength(2);
    expect(mergeTicks([event(0), event(6), event(12)], window, 0.1).map((tick) => tick.count)).toEqual([
      2, 1,
    ]);
  });

  it('draws only what is in the window, positioned from its start', () => {
    const ticks = mergeTicks([event(-5), event(500), event(1001)], { startMs: 100, endMs: 1000 }, 1);
    expect(ticks).toHaveLength(1);
    expect(ticks[0]?.x).toBe(400);
    expect(ticks[0]?.index).toBe(1);
  });
});

describe('clipSegments', () => {
  it('cuts segments to the window and drops those outside', () => {
    expect(
      clipSegments(
        [
          [0, 50],
          [90, 110],
          [120, 150],
          [190, 300],
          [300, 400],
        ],
        { startMs: 100, endMs: 200 },
      ),
    ).toEqual([
      [100, 110],
      [120, 150],
      [190, 200],
    ]);
  });
});

describe('stacksAt', () => {
  it('reads the last change point at or before the time', () => {
    const stacks: [number, number][] = [
      [0, 1],
      [100, 3],
    ];
    expect(stacksAt(stacks, 50)).toBe(1);
    expect(stacksAt(stacks, 100)).toBe(3);
    expect(stacksAt(stacks, -1)).toBe(0);
  });
});

describe('tickNear', () => {
  const ticks = [10, 20, 40].map((x, index) => ({ x, event: event(x), count: 1, index }));

  it('finds the nearest within the tolerance and nothing outside it', () => {
    expect(tickNear(ticks, 19, 5)?.x).toBe(20);
    expect(tickNear(ticks, 30, 5)).toBeNull();
    expect(tickNear(ticks, 14, 5)?.x).toBe(10);
  });
});

describe('readoutAt', () => {
  const sample = [
    cast(0, 'spell:2687', 10),
    cast(1000, 'spell:20662', 40),
    cast(2000, 'other:attack/1', 55),
    cast(3000, 'spell:20662', 5),
  ];
  const model = built(
    result(sample, {
      sample_duration_ms: 10_000,
      sample_auras: [
        { action: 'spell:12328', on: 'player', segments: [[500, 2500]] },
        { action: 'spell:12319', on: 'player', segments: [[0, 9000]] },
      ],
    }),
  );

  it('reads the resource and every aura up now, even one up most of the fight', () => {
    const readout = readoutAt(model, 2200);
    expect(readout.time).toBe('0:02.2');
    expect(readout.resource).toEqual({ label: 'Rage', value: '55' });
    expect(readout.onYou).toEqual(['Flurry', 'Death Wish']);
    expect(readout.onTarget).toEqual([]);
  });

  it('lists recent casts newest first and never a swing', () => {
    const readout = readoutAt(model, 3500);
    expect(readout.casts.map((row) => [row.name, row.agoMs])).toEqual([
      ['Execute', 500],
      ['Execute', 2500],
      ['Bloodrage', 3500],
    ]);
    expect(readout.onYou).toEqual(['Flurry']);
  });

  it('is empty before the first cast', () => {
    expect(readoutAt(model, -100).casts).toEqual([]);
    expect(readoutAt(model, -100).resource).toBeNull();
  });
});

describe('tickFollow', () => {
  const lane = { name: 'Execute', events: [event(0), event(1500)] } as unknown as TickLane;

  it('names the next on the lane and the last of the fight', () => {
    expect(tickFollow(lane, 0)).toEqual({ name: 'Execute', nextInMs: 1500 });
    expect(tickFollow(lane, 1).nextInMs).toBeNull();
  });
});

describe('the keyboard model', () => {
  const lane = (...times: number[]): TickLane =>
    ({ events: times.map((at) => event(at)) }) as unknown as TickLane;
  const lanes = [lane(0, 1000), lane(500, 4000), lane(), lane(2000)];

  it('steps left and right through time across lanes', () => {
    expect(moveCaret(lanes, { lane: 0, index: 0 }, 'ArrowRight')).toEqual({ lane: 1, index: 0 });
    expect(moveCaret(lanes, { lane: 1, index: 0 }, 'ArrowRight')).toEqual({ lane: 0, index: 1 });
    expect(moveCaret(lanes, { lane: 3, index: 0 }, 'ArrowLeft')).toEqual({ lane: 0, index: 1 });
  });

  it('stops at either end', () => {
    expect(moveCaret(lanes, { lane: 0, index: 0 }, 'ArrowLeft')).toBeNull();
    expect(moveCaret(lanes, { lane: 1, index: 1 }, 'ArrowRight')).toBeNull();
  });

  it('moves up and down to the nearest in time, skipping an empty lane', () => {
    expect(moveCaret(lanes, { lane: 0, index: 1 }, 'ArrowDown')).toEqual({ lane: 1, index: 0 });
    expect(moveCaret(lanes, { lane: 1, index: 1 }, 'ArrowDown')).toEqual({ lane: 3, index: 0 });
    expect(moveCaret(lanes, { lane: 3, index: 0 }, 'ArrowUp')).toEqual({ lane: 1, index: 0 });
    expect(moveCaret(lanes, { lane: 0, index: 0 }, 'ArrowUp')).toBeNull();
    expect(moveCaret(lanes, { lane: 3, index: 0 }, 'ArrowDown')).toBeNull();
  });

  it('goes home and to the end, and starts from home with no caret', () => {
    expect(moveCaret(lanes, { lane: 1, index: 0 }, 'Home')).toEqual({ lane: 0, index: 0 });
    expect(moveCaret(lanes, { lane: 1, index: 0 }, 'End')).toEqual({ lane: 1, index: 1 });
    expect(moveCaret(lanes, null, 'ArrowRight')).toEqual({ lane: 0, index: 0 });
  });

  it('finds the landing nearest a time', () => {
    expect(caretAt(lanes, 1900)).toEqual({ lane: 3, index: 0 });
    expect(caretAt([], 0)).toBeNull();
  });

  it('has nowhere to go with no landings', () => {
    expect(moveCaret([lane()], null, 'Home')).toBeNull();
  });
});

describe('rows and the ready height', () => {
  const sample = [cast(0, 'spell:20662', 5), cast(1, 'spell:2687', 5), cast(2, 'other:attack/1', 5)];

  it('draws no group without lanes, and adds Swings and ticks only when on', () => {
    const model = built(result(sample, { sample_auras: [] }));
    const groups = (swings: boolean): string[] =>
      layoutRows(model, swings).rows.flatMap((row) => (row.kind === 'group' ? [row.group] : []));
    expect(groups(false)).toEqual(['casts']);
    expect(groups(true)).toEqual(['casts', 'swings']);
  });

  it('replaces the bar groups with one line on an older result', () => {
    const { rows } = layoutRows(built(result(sample)), false);
    expect(rows.map((row) => row.kind)).toEqual(['axis', 'group', 'lane', 'lane', 'notice', 'resource']);
  });

  it('gives the older line room to wrap on a phone', () => {
    const notice = (phone: boolean): number | undefined =>
      layoutRows(built(result(sample)), false, phone).rows.find((row) => row.kind === 'notice')?.height;
    expect(notice(true)).toBeGreaterThan(notice(false)!);
  });

  it('prints every lane name on its own row above the lane on a phone', () => {
    const { rows } = layoutRows(built(result(sample)), false, true);
    expect(rows.map((row) => row.kind)).toEqual([
      'axis',
      'group',
      'name',
      'lane',
      'name',
      'lane',
      'notice',
      'resource',
    ]);
    const lane = rows.find((row) => row.kind === 'lane');
    const name = rows.find((row) => row.kind === 'name');
    expect(name!.y + name!.height).toBe(lane!.y);
  });
});

describe('the fixture', () => {
  const model = built(fullFixture as unknown as SimResult);

  it('builds its lanes in the groups the spec names', () => {
    expect(model.casts.length).toBeGreaterThan(0);
    expect(model.swings.map((lane) => lane.name)).toContain('Main hand');
    expect(model.onYou.length).toBeGreaterThan(0);
    expect(model.onTarget.length).toBeGreaterThan(0);
    expect(model.hasOutcomes).toBe(true);
  });

  it('never lists a queued spell twice, and every press is inside the whole fight', () => {
    const pressKeys = new Set(model.presses.map((press) => press.action));
    expect(pressKeys.has('spell:11567/1')).toBe(true);
    expect(pressKeys.has('spell:11567')).toBe(false);
    expect(pressesInView(model, { startMs: model.startMs, endMs: model.endMs })).toHaveLength(
      model.presses.length,
    );
  });
});

describe('the fixture readout (wow-player M8)', () => {
  const names: ActionNames = {
    spell: {
      ...NAMES.spell,
      '1719': 'Recklessness',
      '18499': 'Berserker Rage',
      '20554': 'Berserking',
      '1249110': 'Frozen Heart of the Mountain',
      '10610': 'Windfury Totem',
      '1329555': 'Reckless Swings',
    },
    item: { '13442': 'Mighty Rage Potion' },
  };
  const model = buildTimeline(fullFixture as unknown as SimResult, names)!;

  it('lists every aura up at 0:16.5, Flurry included', () => {
    const readout = readoutAt(model, 16_500);
    expect(readout.onYou).toEqual(['Mighty Rage Potion', 'Flurry']);
    expect(readout.onTarget).toEqual(['Deep Wounds']);
  });

  it('lists the cooldowns first and then Flurry at 2:21.5', () => {
    const readout = readoutAt(model, 141_500);
    expect(readout.onYou).toEqual(['Death Wish', 'Recklessness', 'Flurry']);
    expect(readout.onTarget).toEqual(['Deep Wounds']);
  });

  it('leaves the debuffs up from pull to end to the Buffs card, and keeps Deep Wounds', () => {
    expect(model.onTarget.map((bar) => bar.name)).toEqual(['Deep Wounds']);
  });
});
