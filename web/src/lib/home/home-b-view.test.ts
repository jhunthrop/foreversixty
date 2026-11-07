// Home B's head figures and rail words (lib/home/home-b-view.ts, spec 2026-11-07 §1).
import { describe, expect, it } from 'vitest';
import {
  bandFigure,
  deltaText,
  gainFigure,
  guildSimsMeta,
  moversMeta,
  moversTitle,
  simFigure,
  specName,
  specWeekView,
} from './home-b-view';
import type { SpecWeek } from './home-answer';

const WEEK: SpecWeek = {
  spec: 'warrior-fury',
  band: 20,
  faction: 'horde',
  role: 'dps',
  grade: 'A',
  last_week_grade: 'B',
  place: 4,
  of: 23,
};

describe('home B view', () => {
  it('reads the sim on import in each of its states', () => {
    expect(
      simFigure({ state: 'done', dps: 26.24, simmed_at: '2026-11-06T12:00:00Z', sim_id: 's' }),
    ).toMatchObject({
      value: '26.2',
      unit: 'DPS',
      label: 'Simmed on import',
      caption: 'Simmed 6 Nov',
    });
    expect(simFigure({ state: 'pending', dps: null, simmed_at: null, sim_id: null }).value).toBe(
      'Sim pending',
    );
    expect(simFigure({ state: 'failed', dps: null, simmed_at: null, sim_id: null }).link?.href).toBe('/sim');
    expect(simFigure(null).value).toBe('No gear imported');
  });

  it('says what the upgrades add and how far the next band is, or the slots at 60', () => {
    expect(gainFigure(10.34, 6)).toMatchObject({ value: '+10.3', label: '6 upgrades' });
    expect(gainFigure(0, 0)).toBeNull();
    expect(bandFigure(20, null)).toMatchObject({ value: '10', unit: 'levels', label: 'To band 30' });
    expect(bandFigure(34, null)).toMatchObject({ value: '6', label: 'To band 40' });
    expect(bandFigure(60, 12)).toMatchObject({ value: '12 of 17', label: 'On best in slot' });
    expect(bandFigure(60, null)).toBeNull();
  });

  it('moves the grade only against a real last week', () => {
    expect(specName('warrior-fury')).toBe('Fury Warrior');
    expect(specWeekView(WEEK)).toEqual({
      title: 'Fury Warrior at 20 to 29',
      sub: 'Up from B last week, #4 of 23 damage specs',
    });
    expect(specWeekView({ ...WEEK, grade: 'B', last_week_grade: 'A' }).sub).toMatch(/^Down from A last week/);
    expect(specWeekView({ ...WEEK, last_week_grade: 'A' }).sub).toMatch(/^Same as last week/);
    expect(specWeekView({ ...WEEK, last_week_grade: null }).sub).toBe('#4 of 23 damage specs');
  });

  it('calls the movers movers only with a baseline', () => {
    const movers = { spec: 'warrior-fury', band: 20, baseline: true, rows: [] };
    expect(moversTitle(movers)).toBe('Top movers');
    expect(moversMeta(movers)).toBe('Fury, 20 to 29');
    expect(moversTitle({ ...movers, baseline: false })).toBe('Top 3 for Fury Warrior, 20 to 29');
    expect(deltaText(6)).toBe('▲ 6');
    expect(deltaText(-2)).toBe('▼ 2');
    expect(deltaText(null)).toBeNull();
  });

  it('names the guild ladder by role, band and preset', () => {
    expect(
      guildSimsMeta({
        guild_id: 2,
        guild_name: 'O',
        role: 'dps',
        band: 60,
        preset: 'raid',
        unit: 'DPS',
        of: 14,
        rows: [],
        you: null,
      }),
    ).toBe('DPS, 60 raid');
  });
});
