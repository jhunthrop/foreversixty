// Home's sim figure and rail words (lib/home/home-b-view.ts, spec 2026-11-07 §1).
import { describe, expect, it } from 'vitest';
import { simFigure, specName, specWeekView } from './home-b-view';
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
    expect(
      simFigure(
        { state: 'done', dps: 122.1, simmed_at: '2026-11-06T04:00:00Z', sim_id: 'imp34' },
        new Date('2026-11-06T18:00:00Z'),
      ).link,
    ).toEqual({ label: 'Sim 14 h ago', href: '/sim/imp34' });
    expect(simFigure({ state: 'done', dps: 1, simmed_at: null, sim_id: null }).link).toBeUndefined();
    expect(simFigure({ state: 'pending', dps: null, simmed_at: null, sim_id: null }).value).toBe(
      'Sim pending',
    );
    expect(simFigure({ state: 'failed', dps: null, simmed_at: null, sim_id: null }).link?.href).toBe('/sim');
    expect(simFigure(null).value).toBe('No gear imported');
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
});
