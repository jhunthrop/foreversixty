// web/src/lib/sim/timeline-lanes.test.ts
// The light lane pass the page's main chunk sizes the timeline with must agree with the drawn
// model on real samples, at both widths; and the fixture's Deep Wounds bar starts where its
// sample says the debuff went on (ux-designer build review, deviation 3).
import { describe, expect, it } from 'vitest';
import fullFixture from '../../fixtures/sim/result-full.json';
import olderFixture from '../../fixtures/sim/result.json';
import { buildTimeline, layoutRows } from './timeline';
import { plotHeightOf, timelineHeadMeta, timelineShape, upAllFight } from './timeline-lanes';
import type { SimResult } from './types';

const full = fullFixture as unknown as SimResult;
const older = olderFixture as unknown as SimResult;

describe('timelineShape agrees with the drawn model', () => {
  it.each([
    ['the real-engine fixture', full],
    ['the older fixture', older],
  ])('on %s, at phone and wide', (_, result) => {
    const model = buildTimeline(result, null);
    const shape = timelineShape(result);
    if (model === null) {
      expect(shape).toBeNull();
      return;
    }
    expect(shape).not.toBeNull();
    expect(shape!.presses).toBe(model.presses.length);
    for (const phone of [false, true]) {
      expect(plotHeightOf(shape!, phone)).toBe(layoutRows(model, false, phone).height);
    }
  });

  it('counts the fixture’s lanes: twelve casts, ten on you, Deep Wounds alone on the target', () => {
    expect(timelineShape(full)).toMatchObject({
      castLanes: 12,
      windowLane: false,
      playerLanes: 10,
      targetLanes: 1,
      hasAuras: true,
      hasResource: true,
    });
  });

  it('has no shape and no meta without a fight', () => {
    const empty = { ...full, sample: [] } as SimResult;
    expect(timelineShape(empty)).toBeNull();
    expect(timelineHeadMeta(empty)).toBeUndefined();
    expect(timelineHeadMeta(full)).toBe('The median run · 2:53 long');
  });
});

describe('upAllFight', () => {
  it('is up from the pull to the end, never applied after the pull', () => {
    expect(upAllFight({ segments: [[0, 1000]] }, 1000)).toBe(true);
    expect(upAllFight({ segments: [[10, 1000]] }, 1000)).toBe(false);
    expect(upAllFight({ segments: [[0, 500]] }, 1000)).toBe(false);
    expect(upAllFight({ segments: [] }, 1000)).toBe(false);
  });
});

describe('the fixture’s Deep Wounds bar (spec 10, check 8)', () => {
  const TICK_MS = 3000;
  const rows = full.sample!.filter((row) => row.action === 'spell:12867');
  const ticks = rows.filter((row) => row.outcome === 'tick' || row.outcome === 'tick_crit');
  const applications = rows.filter((row) => row.outcome === undefined);
  const bar = full.sample_auras!.find((aura) => aura.on === 'target' && aura.action === 'spell:12867')!;

  it('starts at the first application, one tick period before the first tick', () => {
    // The off-hand swing crits at the pull and applies it at 0:00.01; it first ticks at 0:03.
    expect(bar.segments[0]![0]).toBe(applications[0]!.at_ms);
    expect(ticks[0]!.at_ms - bar.segments[0]![0]).toBe(TICK_MS);
  });

  it('is up at every tick', () => {
    for (const tick of ticks) {
      expect(bar.segments.some(([from, to]) => from <= tick.at_ms && tick.at_ms <= to)).toBe(true);
    }
  });
});
