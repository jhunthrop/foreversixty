// web/src/lib/report/frame.test.ts
// The report frame's words (logs spec 2026-11-07 6.1): the tab strip, the Fight list, the head's
// answer and This pull, on the bundled real log's own pulls.
import { describe, expect, it } from 'vitest';
import sample from '../../../public/sample-report/sample/meta.json';
import {
  NO_BOSS_PULLS,
  TABS_MD,
  fightLabel,
  fightOptions,
  formatRate,
  idleToggleLabel,
  listedTabs,
  pullLength,
  reportFigure,
  reportMeta,
  thisPullFacts,
} from './frame';
import type { FightEntry } from './types';

const FIGHTS = sample.fights as FightEntry[];
const PULL_6 = FIGHTS.find((fight) => fight.index === 6) as FightEntry;
const BOSS: FightEntry = {
  ...PULL_6,
  index: 90,
  kind: 'encounter',
  name: 'Warden Kelthas',
  kill: false,
  boss_health_pct: 34,
  duration_ms: 72_000,
};

describe('the report tabs (decision R3)', () => {
  it('lists the seven first tabs in the spec’s order, Coach only when listed', () => {
    expect(
      listedTabs(false)
        .slice(0, TABS_MD)
        .map((tab) => tab.label),
    ).toEqual(['Summary', 'Damage done', 'Healing', 'Damage taken', 'Deaths', 'Mechanics', 'Timeline']);
    expect(listedTabs(false).some((tab) => tab.id === 'coach')).toBe(false);
    expect(listedTabs(true).some((tab) => tab.id === 'coach')).toBe(true);
    expect(listedTabs(true).at(-1)?.label).toBe('Compare');
  });
});

describe('the Fight list (decision R2)', () => {
  it('says a report with no boss pull has none, and numbers trash by length', () => {
    const options = fightOptions(FIGHTS, 6, false);
    expect(options[0]).toMatchObject({ label: NO_BOSS_PULLS, disabled: true, group: 'Boss pulls' });
    expect(options.find((option) => option.value === '6')).toMatchObject({
      label: 'Trash · 12:43',
      description: '21:07:51 · 177 mobs died',
    });
  });

  it('hides the pulls that killed nothing until asked, and says how many', () => {
    const idle = FIGHTS.filter((fight) => fight.kind !== 'encounter' && (fight.mob_deaths ?? 0) === 0);
    expect(idle.length).toBeGreaterThan(0);
    expect(fightOptions(FIGHTS, 6, false).some((option) => option.value === String(idle[0].index))).toBe(
      false,
    );
    expect(fightOptions(FIGHTS, 6, true).some((option) => option.value === String(idle[0].index))).toBe(true);
    expect(idleToggleLabel(FIGHTS)).toBe(`Show ${idle.length} pulls that killed nothing`);
  });

  it('reads one mob died, never one mobs', () => {
    const one = { ...PULL_6, index: 99, mob_deaths: 1 };
    expect(fightOptions([one], 99, false)[1].description).toMatch(/· 1 mob died$/);
  });

  it('numbers a boss pull and its wipe', () => {
    expect(fightLabel(BOSS, 2)).toBe('Warden Kelthas #2 · Wipe 34% · 1:12');
    const options = fightOptions([BOSS], 90, false);
    expect(options.map((option) => option.label)).toEqual([
      'All boss pulls',
      'Warden Kelthas #1 · Wipe 34% · 1:12',
    ]);
  });
});

describe('the band’s meta line (directions b1)', () => {
  it('reads a trash pull as zone, pull, length and mobs died', () => {
    expect(reportMeta('Zephras Isle', PULL_6, undefined, 5)).toBe(
      'Zephras Isle · Pull 6, Trash · 12:43 · 177 mobs died',
    );
  });

  it('reads a boss pull with its attempt, outcome, length and players', () => {
    expect(reportMeta('', BOSS, 2, 40)).toBe('Pull 90, Warden Kelthas #2 · Wipe 34% · 1:12 · 40 players');
  });
});

describe('the band’s figure follows the tab (directions b1)', () => {
  const totals = { damageDone: 2_400, healing: 142, damageTaken: 737, deaths: 2, durationMs: 763_000 };

  it('is the group’s rate on the rate tabs', () => {
    expect(reportFigure({ ...totals, tab: 'summary', player: false })).toEqual({
      value: '3.1',
      unit: 'group DPS',
    });
    expect(reportFigure({ ...totals, tab: 'healing', player: false })).toEqual({
      value: '0.2',
      unit: 'group HPS',
    });
    expect(reportFigure({ ...totals, tab: 'damage-taken', player: false })).toEqual({
      value: '1.0',
      unit: 'group DTPS',
    });
  });

  it('is the player’s own rate on their view, and deaths on Deaths and Mechanics', () => {
    expect(reportFigure({ ...totals, tab: 'damage-done', player: true })).toEqual({
      value: '3.1',
      unit: 'DPS',
    });
    expect(reportFigure({ ...totals, tab: 'deaths', player: false })).toEqual({ value: '2', unit: 'deaths' });
    expect(reportFigure({ ...totals, tab: 'mechanics', player: false, deaths: 1 })).toEqual({
      value: '1',
      unit: 'death',
    });
  });

  it('is absent where no one figure answers the tab', () => {
    for (const tab of ['timeline', 'threat', 'resources', 'queries', 'coach'] as const)
      expect(reportFigure({ ...totals, tab, player: false })).toBeNull();
  });

  it('reads a rate with one decimal under 100 and whole from 100', () => {
    expect(formatRate(1.84)).toBe('1.8');
    expect(formatRate(1374.4)).toBe('1,374');
    expect(pullLength(4_000)).toBe('0:04');
  });
});

describe('This pull (directions b1)', () => {
  it('says only what the band does not: the start, the scope’s count and its deaths', () => {
    const facts = thisPullFacts({ fight: PULL_6, scope: 'Your group', players: 5, deaths: 0, phases: [] });
    expect(facts.map((fact) => [fact.label, fact.value])).toEqual([
      ['Started', '21:07:51'],
      ['In your group', '5'],
      ['Deaths', '0'],
    ]);
  });

  it('adds a boss pull’s phases', () => {
    const facts = thisPullFacts({
      fight: BOSS,
      scope: 'Group',
      players: 5,
      deaths: 1,
      phases: [{ name: 'Phase 1', start_ms: 0, end_ms: 14_000 }],
    });
    expect(facts.map((fact) => fact.label)).toEqual(['Started', 'In group', 'Deaths', 'Phases']);
  });

  it('reads everyone the logger saw as In range', () => {
    const facts = thisPullFacts({
      fight: PULL_6,
      scope: 'Everyone in range',
      players: 265,
      deaths: 2,
      phases: [],
    });
    expect(facts[1]).toMatchObject({ label: 'In range', value: '265' });
  });
});
