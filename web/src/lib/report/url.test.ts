// web/src/lib/report/url.test.ts
import { describe, expect, it } from 'vitest';
import {
  ALL_FIGHTS,
  MODES,
  TABS,
  VIEWS,
  defaultState,
  parseReportState,
  reportSearch,
  withState,
} from './url';

describe('the report URL state', () => {
  it('offers the five modes, with Mechanics and Replay disabled', () => {
    expect(MODES.map((m) => m.id)).toEqual(['analyze', 'compare', 'rankings', 'mechanics', 'replay']);
    expect(MODES.filter((m) => !m.enabled).map((m) => m.label)).toEqual(['Mechanics', 'Replay']);
    expect(MODES.filter((m) => !m.enabled).every((m) => m.note === 'later')).toBe(true);
  });

  it('offers the four views and the twelve tabs in the spec’s order', () => {
    expect(VIEWS.map((v) => v.id)).toEqual(['tables', 'timelines', 'events', 'queries']);
    expect(TABS.map((t) => t.id)).toEqual([
      'summary',
      'damage-done',
      'damage-taken',
      'healing',
      'threat',
      'buffs',
      'debuffs',
      'deaths',
      'interrupts',
      'dispels',
      'resources',
      'casts',
    ]);
    expect(TABS.map((t) => t.label)).toEqual([
      'Summary',
      'Damage Done',
      'Damage Taken',
      'Healing',
      'Threat',
      'Buffs',
      'Debuffs',
      'Deaths',
      'Interrupts',
      'Dispels',
      'Resources',
      'Casts',
    ]);
  });

  it('defaults to the report’s first fight, analyze, tables, summary, all friendlies', () => {
    expect(defaultState(1)).toEqual({
      fight: 1,
      mode: 'analyze',
      view: 'tables',
      tab: 'summary',
      source: 'friendlies',
      start: null,
      end: null,
      target: '',
      ability: null,
      flags: [],
      compareWith: null,
      compareMetric: '',
      eventsOff: [],
      find: '',
      openDeaths: [],
      rankingsSpec: '',
      rankingsMetric: '',
    });
  });

  it('parses every field out of a full query string', () => {
    const state = parseReportState(
      '?fight=3&mode=compare&view=events&tab=deaths&source=Player-4184-000000A1&start=4000&end=12000',
      1,
    );
    expect(state).toEqual({
      fight: 3,
      mode: 'compare',
      view: 'events',
      tab: 'deaths',
      source: 'Player-4184-000000A1',
      start: 4000,
      end: 12000,
      target: '',
      ability: null,
      flags: [],
      compareWith: null,
      compareMetric: '',
      eventsOff: [],
      find: '',
      openDeaths: [],
      rankingsSpec: '',
      rankingsMetric: '',
    });
  });

  it('falls back to the default for anything it does not recognise', () => {
    const state = parseReportState('?fight=nope&mode=replay&view=sideways&tab=gear&start=-5&end=abc', 2);
    expect(state).toEqual({
      fight: 2,
      mode: 'analyze',
      view: 'tables',
      tab: 'summary',
      source: 'friendlies',
      start: null,
      end: null,
      target: '',
      ability: null,
      flags: [],
      compareWith: null,
      compareMetric: '',
      eventsOff: [],
      find: '',
      openDeaths: [],
      rankingsSpec: '',
      rankingsMetric: '',
    });
  });

  it('drops a window whose end is not after its start', () => {
    expect(parseReportState('?start=9000&end=9000', 1).start).toBeNull();
    expect(parseReportState('?start=9000&end=1000', 1).end).toBeNull();
  });

  it('serialises only what differs from the default', () => {
    expect(reportSearch(defaultState(1), 1)).toBe('');
    expect(reportSearch(withState(defaultState(1), { tab: 'healing' }), 1)).toBe('?tab=healing');
    expect(reportSearch(withState(defaultState(1), { fight: 3, start: 1000, end: 5000 }), 1)).toBe(
      '?fight=3&start=1000&end=5000',
    );
  });

  it('serialises the fields in the contract’s order', () => {
    const state = withState(defaultState(1), {
      fight: 3,
      mode: 'rankings',
      view: 'queries',
      tab: 'casts',
      source: 'enemies',
      start: 10,
      end: 20,
      target: '',
      ability: null,
      flags: [],
      compareWith: null,
      compareMetric: '',
      eventsOff: [],
      find: '',
      openDeaths: [],
      rankingsSpec: '',
      rankingsMetric: '',
    });
    expect(reportSearch(state, 1)).toBe(
      '?fight=3&mode=rankings&view=queries&tab=casts&source=enemies&start=10&end=20',
    );
  });

  it('round-trips through parse', () => {
    const state = withState(defaultState(1), { fight: 2, tab: 'buffs', source: 'enemies', start: 3, end: 9 });
    expect(parseReportState(reportSearch(state, 1), 1)).toEqual(state);
  });

  it('withState copies rather than mutating', () => {
    const state = defaultState(1);
    const next = withState(state, { tab: 'threat' });
    expect(state.tab).toBe('summary');
    expect(next.tab).toBe('threat');
    expect(next).not.toBe(state);
  });

  it('spells the whole night as fight=all and reads it back as ALL_FIGHTS', () => {
    expect(parseReportState('?fight=all', 1).fight).toBe(ALL_FIGHTS);
    expect(reportSearch(withState(defaultState(1), { fight: ALL_FIGHTS }), 1)).toBe('?fight=all');
    expect(
      parseReportState(reportSearch(withState(defaultState(1), { fight: ALL_FIGHTS }), 1), 1).fight,
    ).toBe(ALL_FIGHTS);
  });

  it('carries the filter bar and the compare pairing, so a copied link is the view', () => {
    const state = withState(defaultState(1), {
      target: 'Creature-1-2',
      ability: 116,
      flags: ['bossOnly', 'ignoreAfterDeath'],
      compareWith: 4,
      compareMetric: 'dps',
      eventsOff: [],
      find: '',
      openDeaths: [],
      rankingsSpec: '',
      rankingsMetric: '',
    });
    const search = reportSearch(state, 1);
    expect(search).toBe('?target=Creature-1-2&ability=116&flags=bd&with=4&cmetric=dps');
    expect(parseReportState(search, 1)).toEqual(state);
  });
});
