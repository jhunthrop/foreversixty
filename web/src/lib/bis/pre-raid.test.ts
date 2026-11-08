// web/src/lib/bis/pre-raid.test.ts
import { describe, expect, it } from 'vitest';
import { filteredBand, type FilterableSlot } from './source-filter';
import { PRE_RAID_KINDS, isPreRaidKinds, isPreRaidView, listIsPreRaid, preRaidOffered } from './pre-raid';
import { bisContext } from './page-context';
import { isDefaultState, parseState, stateSearch, withChange } from './page-state';
import { setFigureFor } from './source-choice';
import type { BisBandModel } from './page-model';

/** A level 60 band once raids open: the head pick is a raid drop, its best dungeon piece in by_kind;
 *  the hands pick is a PvP reward with a crafted runner-up. */
const RAID_BAND: FilterableSlot[] = [
  {
    slot: 'head',
    item_id: 16866,
    source_kind: 'raid',
    by_kind: { dungeon: { item_id: 12587, rank: 2, source: 'Blackrock Spire', dps_delta: -11.1 } },
  },
  {
    slot: 'hands',
    item_id: 16548,
    source_kind: 'pvp',
    by_kind: { crafted: { item_id: 15063, rank: 3, source: 'Leatherworking', dps_delta: -4.2 } },
  },
  { slot: 'feet', item_id: 13967, source_kind: 'dungeon' },
];

/** Today's level 60 lists: no pick from a raid or PvP (by_kind may still list PvP runners-up). */
const PRE_RAID_BAND: FilterableSlot[] = [
  {
    slot: 'head',
    item_id: 12640,
    source_kind: 'crafted',
    by_kind: { pvp: { item_id: 231538, rank: 2, source: 'PvP reward', dps_delta: -5.6 } },
  },
  { slot: 'feet', item_id: 13967, source_kind: 'dungeon' },
];

describe('the Pre-raid toggle', () => {
  it('turns off raid and PvP and nothing else', () => {
    expect(PRE_RAID_KINDS).not.toContain('raid');
    expect(PRE_RAID_KINDS).not.toContain('pvp');
    expect(PRE_RAID_KINDS).toHaveLength(6);
    expect(isPreRaidKinds([...PRE_RAID_KINDS].reverse())).toBe(true);
    expect(isPreRaidView(50, PRE_RAID_KINDS)).toBe(false);
  });

  it('shows on a level 60 band with a raid or PvP pick, and hides on a list that is already pre-raid', () => {
    expect(preRaidOffered(60, RAID_BAND)).toBe(true);
    expect(preRaidOffered(60, PRE_RAID_BAND)).toBe(false);
    expect(preRaidOffered(50, RAID_BAND)).toBe(false);
  });

  it('resolves the pre-raid picks from by_kind', () => {
    const picks = filteredBand(RAID_BAND, new Set(PRE_RAID_KINDS));
    expect(picks.get('head')).toMatchObject({ item_id: 12587, source_kind: 'dungeon', dps_delta: -11.1 });
    expect(picks.get('hands')).toMatchObject({ item_id: 15063, source_kind: 'crafted' });
    expect(picks.get('feet')).toMatchObject({ item_id: 13967, dps_delta: 0 });
  });

  it('calls a list pre-raid only while no pick is had only from a raid or PvP', () => {
    expect(listIsPreRaid(PRE_RAID_BAND)).toBe(true);
    expect(listIsPreRaid(RAID_BAND)).toBe(false);
    // An item a dungeon also drops stays in the pre-raid list: the toggle does not move it.
    expect(listIsPreRaid([{ item_id: 1, source_kind: 'raid', kinds: ['raid', 'dungeon'] }])).toBe(true);
  });
});

describe('the pre-raid page state', () => {
  const data = { bands: [20, 30, 40, 50, 60], models: {} };
  const races = { alliance: [{ value: 'human' }], horde: [{ value: 'orc' }] };

  it('opens /60/pre-raid on the pre-raid kinds with nothing in the query', () => {
    const context = bisContext(data, races, 60, true);
    const state = parseState('', '', context);
    expect(isPreRaidKinds(state.sources)).toBe(true);
    expect(stateSearch(state, context)).toBe('');
    expect(isDefaultState(state, context)).toBe(true);
  });

  it('keeps the pre-raid kinds out of the query on /60 (the path carries them) and moves off them on /60/pre-raid', () => {
    const plain = bisContext(data, races, 60, false);
    const preRaid = withChange(parseState('', '', plain), { sources: [...PRE_RAID_KINDS] }, plain);
    expect(stateSearch(preRaid, plain)).toBe('');
    expect(isDefaultState(preRaid, plain)).toBe(false);
    const onPage = bisContext(data, races, 60, true);
    const all = withChange(
      parseState('', '', onPage),
      { sources: [...PRE_RAID_KINDS, 'raid', 'pvp'] },
      onPage,
    );
    expect(isDefaultState(all, onPage)).toBe(false);
  });
});

describe('the set figure under the toggle', () => {
  const model = (filtered?: BisBandModel['answer']['filtered']): BisBandModel =>
    ({
      band: 60,
      unit: 'DPS',
      answer: { figure: '1004.2', simmed: 'Simmed 7 Nov', error: 1.1, ...(filtered ? { filtered } : {}) },
    }) as unknown as BisBandModel;

  it('shows the unfiltered set, labelled as such, until the ranker publishes pre_raid', () => {
    const figure = setFigureFor(model(), PRE_RAID_KINDS, 2);
    expect(figure?.value).toBe('1004.2');
    expect(figure?.label).toBe('Best set, all sources');
  });

  it('shows the pre_raid figure once present', () => {
    const figure = setFigureFor(
      model({ pre_raid: { dps: 990.4, error: 1.3, changed_slots: 2 } }),
      PRE_RAID_KINDS,
      2,
    );
    expect(figure?.value).toBe('990.4');
    expect(figure?.label).toBe('Set: pre-raid');
  });
});
