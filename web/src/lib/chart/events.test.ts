// web/src/lib/chart/events.test.ts
import { describe, expect, it } from 'vitest';
import type { Death, Taunt } from '../report/types';
import {
  castEvent,
  clusterRows,
  deathEvent,
  eventsNear,
  laneItems,
  phaseEvent,
  tauntEvent,
  targetChangeEvent,
} from './events';
import type { ChartEvent } from './types';

const death: Death = {
  guid: 'Player-1',
  name: 'Brokk-Nightslayer',
  at_ms: 126_000,
  killing_blow: {
    at_ms: 126_000,
    source_guid: 'Creature-1',
    source_name: 'Ragnaros',
    spell_id: 0,
    spell_name: 'Melee',
    amount: 3862,
    overkill: 1104,
  },
  last: [],
  auras_held: [],
  auras_lost: [],
};

describe('event lines (spec §B8)', () => {
  it('names the killing blow on a death, with the overkill', () => {
    const event = deathEvent(death, '?tab=deaths');
    expect(event.text).toBe('Brokk died · Melee from Ragnaros · 3,862 (1,104 overkill)');
    expect(event.shortText).toBe('Brokk died · Melee · 3,862');
    expect(event.href).toBe('?tab=deaths');
    expect(event.clusterText).toBe('Brokk · Melee');
  });

  it('leaves the overkill out when there was none, and the blow out when the summary has none', () => {
    const blow = { ...death.killing_blow!, overkill: 0 };
    expect(deathEvent({ ...death, killing_blow: blow }, '').text).toBe(
      'Brokk died · Melee from Ragnaros · 3,862',
    );
    expect(deathEvent({ ...death, killing_blow: undefined }, '').text).toBe('Brokk died');
  });

  it('reads a taunt, a target change, a cast and a phase', () => {
    const taunt: Taunt = {
      at_ms: 127_000,
      source_guid: 'Player-2',
      source_name: 'Thandril-Nightslayer',
      target_guid: 'Creature-1',
      target_name: 'Ragnaros',
      spell_id: 355,
      spell_name: 'Taunt',
    };
    expect(tauntEvent(taunt).text).toBe('Taunt · Thandril on Ragnaros');
    expect(
      targetChangeEvent({
        at_ms: 127_000,
        enemy_guid: 'Creature-1',
        enemy_name: 'Ragnaros',
        to_guid: 'Player-2',
        to_name: 'Thandril-Nightslayer',
        basis: 'melee',
      }).text,
    ).toBe('Ragnaros turned to Thandril');
    expect(castEvent(5000, 'Rebirth', 'Sunwick-Nightslayer', 20484).text).toBe('Rebirth · Sunwick');
    expect(phaseEvent({ name: 'Submerge', startMs: 180_000, endMs: 270_000 }).text).toBe('Phase · Submerge');
  });
});

const at = (atMs: number, kind: ChartEvent['kind'] = 'cast'): ChartEvent => ({
  key: `${kind}-${atMs}`,
  kind,
  atMs,
  text: `${kind} at ${atMs}`,
});

describe('eventsNear', () => {
  it('lists every event within ±2 s of the second, deaths included, in time order', () => {
    const events = [
      at(130_000),
      at(126_000, 'death'),
      at(127_000, 'target-change'),
      at(127_400, 'taunt'),
      at(124_000),
    ];
    expect(eventsNear(events, 127_000).map((event) => event.atMs)).toEqual([126_000, 127_000, 127_400]);
  });

  it('orders one second’s events death, then target change, then taunt', () => {
    const same = [at(127_400, 'taunt'), at(127_000, 'target-change'), at(127_200, 'death')];
    expect(eventsNear(same, 127_000).map((event) => event.kind)).toEqual(['death', 'target-change', 'taunt']);
  });

  it('includes the edges of the window', () => {
    expect(eventsNear([at(125_000), at(129_000)], 127_000)).toHaveLength(2);
    expect(eventsNear([at(124_999)], 127_000)).toHaveLength(0);
  });
});

describe('laneItems', () => {
  const xOf = (ms: number): number => ms / 100; // 10 px per second

  it('steps a marker 22 px right of a neighbour, with a leader back to its time', () => {
    const items = laneItems([at(10_000), at(10_500)], xOf);
    expect(items.map((item) => item.x)).toEqual([100, 122]);
    expect(items[1].leader).toBe(true);
    expect(items[0].leader).toBe(false);
  });

  it('collapses three or more deaths inside 44 px into one counted marker', () => {
    const wipe = [0, 500, 1000, 1500, 2000].map((offset) => at(302_000 + offset, 'death'));
    const items = laneItems(wipe, xOf);
    expect(items).toHaveLength(1);
    expect(items[0].kind).toBe('cluster');
    expect(items[0].tone).toBe('death');
    expect(items[0].events).toHaveLength(5);
    expect(items[0].startMs).toBe(302_000);
    expect(items[0].endMs).toBe(304_000);
  });

  it('keeps two deaths apart, and collapses three other markers with a gold ring', () => {
    expect(laneItems([at(1000, 'death'), at(2000, 'death')], xOf)).toHaveLength(2);
    const casts = laneItems([at(1000), at(1500), at(2000)], xOf);
    expect(casts).toHaveLength(1);
    expect(casts[0].tone).toBe('gold');
  });
});

describe('clusterRows', () => {
  it('lists six deaths, then counts the rest', () => {
    const deaths = Array.from({ length: 14 }, (_, index) => at(302_000 + index * 100, 'death'));
    const rows = clusterRows(deaths);
    expect(rows.shown).toHaveLength(6);
    expect(rows.more).toBe(8);
  });
});
