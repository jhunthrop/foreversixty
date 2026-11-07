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
  reportAnswer,
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

describe('the head’s answer (decision R1 as ruled)', () => {
  it('leads a trash pull with its mobs and length, never a DPS', () => {
    const answer = reportAnswer({ fight: PULL_6, scope: 'Your group', players: 5, deaths: 0, groupDps: 3 });
    expect(answer).toMatchObject({
      label: 'Trash pull · started 21:07:51',
      value: '177',
      unit: 'mobs · 12:43',
      caption: 'Your group · 5 players · no deaths',
    });
  });

  it('leads a boss pull with the DPS', () => {
    const answer = reportAnswer({
      fight: BOSS,
      attempt: 2,
      scope: 'Group',
      players: 5,
      deaths: 1,
      groupDps: 1374.4,
    });
    expect(answer).toMatchObject({
      label: 'Warden Kelthas #2 · Wipe at 34% · 1:12',
      value: '1,374',
      unit: 'DPS',
    });
    expect(answer.caption).toBe('Group damage · 1 death');
  });

  it('reads a rate with one decimal under 100 and whole from 100', () => {
    expect(formatRate(1.84)).toBe('1.8');
    expect(formatRate(1374.4)).toBe('1,374');
    expect(pullLength(4_000)).toBe('0:04');
  });
});

describe('This pull (decision R7 as ruled)', () => {
  it('reads Mobs died on trash, never a Result', () => {
    const facts = thisPullFacts({ fight: PULL_6, players: 5, deaths: 0, phases: [] });
    expect(facts.map((fact) => [fact.label, fact.value])).toEqual([
      ['Pull', 'Trash'],
      ['Length', '12:43'],
      ['Mobs died', '177'],
      ['Players', '5'],
      ['Deaths', '0'],
    ]);
  });

  it('reads Result and Phases on a boss pull', () => {
    const facts = thisPullFacts({
      fight: BOSS,
      players: 5,
      deaths: 1,
      phases: [{ name: 'Phase 1', start_ms: 0, end_ms: 14_000 }],
    });
    expect(facts.map((fact) => fact.label)).toEqual([
      'Pull',
      'Result',
      'Length',
      'Players',
      'Deaths',
      'Phases',
    ]);
  });
});
