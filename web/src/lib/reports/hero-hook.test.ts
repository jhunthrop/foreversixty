// web/src/lib/reports/hero-hook.test.ts
import { describe, expect, it } from 'vitest';
import type { FightEntry, RosterRow } from '../report/types';
import { buildHookLinks, formatHookDps, pickLastKillFight, selectHookRow } from './hero-hook';

function fight(overrides: Partial<FightEntry>): FightEntry {
  return {
    index: 1,
    kind: 'encounter',
    name: 'Trash',
    kill: false,
    in_progress: false,
    start: '',
    end: '',
    duration_ms: 0,
    players: [],
    deaths: 0,
    npc_kills: 0,
    ...overrides,
  };
}

function row(overrides: Partial<RosterRow>): RosterRow {
  return {
    guid: 'Player-1',
    name: 'Zulmara',
    role: 'dps',
    active_ms: 0,
    activity_pct: 100,
    deaths: 0,
    damage_done: 0,
    healing_done: 0,
    damage_taken: 0,
    dps: 0,
    hps: 0,
    dtps: 0,
    ...overrides,
  };
}

describe('pickLastKillFight', () => {
  it('is null for a report with no fights', () => {
    expect(pickLastKillFight([])).toBeNull();
  });

  it('picks the last kill when one exists', () => {
    const fights = [
      fight({ index: 1, kind: 'trash', name: 'Trash' }),
      fight({ index: 2, kind: 'encounter', name: 'Boss A', kill: true }),
      fight({ index: 3, kind: 'encounter', name: 'Boss B', kill: false }),
      fight({ index: 4, kind: 'encounter', name: 'Boss C', kill: true }),
    ];
    expect(pickLastKillFight(fights)?.name).toBe('Boss C');
  });

  it('falls back to the last fight overall when the report has no kill', () => {
    const fights = [
      fight({ index: 1, kind: 'trash', name: 'Trash' }),
      fight({ index: 2, name: 'Wipe', kill: false }),
    ];
    expect(pickLastKillFight(fights)?.name).toBe('Wipe');
  });
});

describe('selectHookRow', () => {
  it('is null for an empty roster', () => {
    expect(selectHookRow([], ['Zulmara'])).toBeNull();
  });

  it("matches the visitor's own character by name, case-insensitive", () => {
    const roster = [row({ name: 'Mishvamp', dps: 2231 }), row({ name: 'ZULMARA', dps: 500 })];
    expect(selectHookRow(roster, ['zulmara'])?.name).toBe('ZULMARA');
  });

  it('falls back to the top-dps row with no match', () => {
    const roster = [row({ name: 'Mishvamp', dps: 2231 }), row({ name: 'Reglitch', dps: 2364 })];
    expect(selectHookRow(roster, ['someone-else'])?.name).toBe('Reglitch');
  });
});

describe('formatHookDps', () => {
  it('shows one decimal under 1,000', () => {
    expect(formatHookDps(345.678)).toBe('345.7');
  });

  it('shows a whole number with a thousands separator at or above 1,000', () => {
    expect(formatHookDps(2364.4)).toBe('2,364');
    expect(formatHookDps(1000)).toBe('1,000');
  });
});

describe('buildHookLinks', () => {
  it("omits the planner link when no combatant record matches the row's guid", () => {
    const result = buildHookLinks({
      reportId: 'rep1',
      fightIndex: 3,
      dataBuild: '1.60.1.70009',
      row: row({ guid: 'Player-1', name: 'Zulmara', class: 'Hunter', dps: 500 }),
      combatants: [],
      treeSizesFor: () => [],
    });
    expect(result.plannerLink).toBeNull();
    expect(result.simLink.href).toContain('rep1');
    expect(result.name).toBe('Zulmara');
    expect(result.dpsText).toBe('500.0');
  });

  it('builds a planner link when a matching combatant exists for a known class', () => {
    const result = buildHookLinks({
      reportId: 'rep1',
      fightIndex: 3,
      dataBuild: '1.60.1.70009',
      row: row({ guid: 'Player-1', name: 'Zulmara', class: 'Hunter', dps: 500 }),
      combatants: [
        {
          guid: 'Player-1',
          name: 'Zulmara',
          gear: [],
          talents: [],
          consumables: [],
          raid_buffs: [],
          missing_buffs: [],
        },
      ],
      treeSizesFor: () => [],
    });
    expect(result.plannerLink).not.toBeNull();
    expect(result.plannerLink?.href).toContain('/planner?code=');
  });
});
