// web/src/lib/chart/report-markers.test.ts
import { describe, expect, it } from 'vitest';
import type { CastRow, Death, Taunt, TargetChange } from '../report/types';
import { deathKey, rateMarkers, threatMarkers } from './report-markers';

const BOSS = 'Creature-1';
const death: Death = {
  guid: 'Player-1',
  name: 'Brokk',
  at_ms: 126_000,
  last: [],
  auras_held: [],
  auras_lost: [],
};
const cast = (spell_id: number, spell_name: string, guid = 'Player-2'): CastRow => ({
  guid,
  name: 'Sunwick',
  spell_id,
  spell_name,
  started: 1,
  succeeded: 1,
  failed: 0,
  cast_time_ms: 0,
  sequence: [5000],
});
const href = (entry: Death): string => `?tab=deaths&death=${deathKey(entry)}`;

describe('rateMarkers', () => {
  it('marks deaths with their Deaths-tab link, raid cooldowns by name and life-saving casts by id', () => {
    const events = rateMarkers({
      deaths: [death],
      casts: [
        cast(20484, 'Rebirth'),
        cast(10060, 'Power Infusion'),
        cast(133, 'Fireball'),
        cast(17534, 'Healing Potion'),
      ],
      deathHref: href,
    });
    expect(events.map((event) => event.text)).toEqual([
      'Rebirth · Sunwick',
      'Power Infusion · Sunwick',
      'Major Healing Potion · Sunwick',
      'Brokk died',
    ]);
    expect(events.at(-1)!.href).toBe('?tab=deaths&death=Player-1-126000');
  });

  it('does not mark a cast that is only named like a life-saving one', () => {
    expect(
      rateMarkers({ deaths: [], casts: [cast(999_999, 'Rebirth of the Ancients')], deathHref: href }),
    ).toEqual([]);
  });

  it('with one player as the Source, keeps their casts and deaths only', () => {
    const events = rateMarkers({
      deaths: [death],
      casts: [cast(10060, 'Power Infusion', 'Player-2'), cast(10060, 'Power Infusion', 'Player-3')],
      sourcePlayer: 'Player-2',
      deathHref: href,
    });
    expect(events).toHaveLength(1);
  });
});

describe('threatMarkers', () => {
  const taunt: Taunt = {
    at_ms: 127_000,
    source_guid: 'Player-9',
    source_name: 'Thandril',
    target_guid: BOSS,
    target_name: 'Ragnaros',
    spell_id: 355,
    spell_name: 'Taunt',
  };
  const change: TargetChange = {
    at_ms: 127_000,
    enemy_guid: BOSS,
    enemy_name: 'Ragnaros',
    to_guid: 'Player-9',
    to_name: 'Thandril',
    basis: 'melee',
  };

  it('draws target-change markers when the summary carries target_changes', () => {
    const events = threatMarkers({
      deaths: [death],
      taunts: [taunt],
      targetChanges: [change],
      enemyGuids: [BOSS],
      deathHref: href,
    });
    expect(events.map((event) => event.kind)).toEqual(['death', 'taunt', 'target-change']);
  });

  it('draws none, and nothing else changes, when it does not', () => {
    const events = threatMarkers({ deaths: [death], taunts: [taunt], enemyGuids: [BOSS], deathHref: href });
    expect(events.map((event) => event.kind)).toEqual(['death', 'taunt']);
  });

  it('keeps only this enemy’s taunts and changes', () => {
    const events = threatMarkers({
      deaths: [],
      taunts: [{ ...taunt, target_guid: 'other' }],
      targetChanges: [{ ...change, enemy_guid: 'other' }],
      enemyGuids: [BOSS],
      deathHref: href,
    });
    expect(events).toEqual([]);
  });
});
