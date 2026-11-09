// web/src/lib/guild/readiness-view.test.ts
import { describe, expect, it } from 'vitest';
import type { GuildReadinessRow } from './api';
import {
  isLeveling,
  levelingLabel,
  nudgeText,
  passesEveryCheck,
  pinReadinessOwnRow,
  readinessFails,
  readinessGroups,
  readinessScore,
  sortReadinessWorstFirst,
} from './readiness-view';

function row(overrides: Partial<GuildReadinessRow> = {}): GuildReadinessRow {
  return {
    character_key: 'us/pvp/thornhide',
    name: 'Thornhide',
    class: 'druid',
    spec: 'Feral',
    consent: 'gear_bags',
    level: 60,
    gear_gap: { upgrades: 3, gain_dps: 23.8, not_sim_checked: 0, unit: 'dps' },
    enchants: { missing_slots: ['chest', 'boots'], checked: true },
    consumables: { state: 'short' },
    talent_points_unspent: 1,
    item_level: 58,
    item_level_delta: -5,
    logged_at: '2026-12-20T20:00:00Z',
    failing: 4,
    no_export: false,
    synced_before_cutoff: null,
    ...overrides,
  };
}

describe('readinessScore / sortReadinessWorstFirst', () => {
  it('sorts failCount*100 + gearGainDps worst-first', () => {
    const low = row({
      character_key: 'a',
      failing: 1,
      gear_gap: { upgrades: 1, gain_dps: 10, not_sim_checked: 0, unit: 'dps' },
    });
    const high = row({
      character_key: 'b',
      failing: 3,
      gear_gap: { upgrades: 3, gain_dps: 5, not_sim_checked: 0, unit: 'dps' },
    });
    expect(readinessScore(high)).toBeGreaterThan(readinessScore(low));
    expect(sortReadinessWorstFirst([low, high]).map((r) => r.character_key)).toEqual(['b', 'a']);
  });

  it('breaks ties on the bigger gear gain', () => {
    const smaller = row({
      character_key: 'a',
      failing: 2,
      gear_gap: { upgrades: 2, gain_dps: 10, not_sim_checked: 0, unit: 'dps' },
    });
    const bigger = row({
      character_key: 'b',
      failing: 2,
      gear_gap: { upgrades: 2, gain_dps: 40, not_sim_checked: 0, unit: 'dps' },
    });
    expect(sortReadinessWorstFirst([smaller, bigger]).map((r) => r.character_key)).toEqual(['b', 'a']);
  });

  it('never mutates the input array', () => {
    const rows = [row({ character_key: 'a' }), row({ character_key: 'b', failing: 9 })];
    const original = [...rows];
    sortReadinessWorstFirst(rows);
    expect(rows).toEqual(original);
  });
});

describe('pinReadinessOwnRow', () => {
  it('pins the viewer’s own row first without disturbing worst-first order otherwise', () => {
    const worst = row({ character_key: 'worst', failing: 5 });
    const mine = row({ character_key: 'mine', failing: 1 });
    const middle = row({ character_key: 'middle', failing: 3 });
    const sorted = sortReadinessWorstFirst([mine, worst, middle]);
    const pinned = pinReadinessOwnRow(sorted, 'mine');
    expect(pinned.map((r) => r.character_key)).toEqual(['mine', 'worst', 'middle']);
  });

  it('is a no-op for a null viewer key', () => {
    const rows = [row({ character_key: 'a' }), row({ character_key: 'b' })];
    expect(pinReadinessOwnRow(rows, null)).toEqual(rows);
  });
});

describe('leveling characters', () => {
  const leveling = row({
    level: 52,
    gear_gap: null,
    enchants: { missing_slots: [], checked: false },
    consumables: { state: 'unknown' },
    talent_points_unspent: 0,
    item_level_delta: null,
    failing: 0,
    name: 'Bravo',
    character_key: 'b',
  });

  it('isLeveling is a known level under 60 only', () => {
    expect(isLeveling(leveling)).toBe(true);
    expect(isLeveling(row({ level: 60 }))).toBe(false);
    expect(isLeveling(row({ level: null }))).toBe(false);
  });

  it('reads the level and the distance to 60, one level singular', () => {
    expect(levelingLabel(52)).toBe('Level 52 · 8 levels to 60');
    expect(levelingLabel(59)).toBe('Level 59 · 1 level to 60');
  });

  it('sorts after every raider, highest level first, then by name', () => {
    const alpha = { ...leveling, name: 'Alpha', character_key: 'a' };
    const low = { ...leveling, level: 20, name: 'Low', character_key: 'l' };
    const clean = row({ failing: 0, name: 'Clean', character_key: 'c' });
    const noExport = row({ failing: 0, no_export: true, name: 'None', character_key: 'n' });
    const names = sortReadinessWorstFirst([low, leveling, noExport, clean, alpha, row()]).map((r) => r.name);
    expect(names).toEqual(['Thornhide', 'Clean', 'None', 'Alpha', 'Bravo', 'Low']);
    const groups = readinessGroups([clean, alpha]);
    expect(groups.graded.map((r) => r.name)).toEqual(['Clean']);
    expect(groups.leveling.map((r) => r.name)).toEqual(['Alpha']);
  });

  it('passes nothing: a leveling row, a no-export row and a row with nothing read never read Ready', () => {
    expect(passesEveryCheck(leveling)).toBe(false);
    expect(passesEveryCheck(row({ failing: 0, no_export: true }))).toBe(false);
    expect(
      passesEveryCheck(
        row({
          failing: 0,
          gear_gap: null,
          enchants: { missing_slots: [], checked: false },
          consumables: { state: 'unknown' },
          talent_points_read: false,
        }),
      ),
    ).toBe(false);
    expect(passesEveryCheck(row({ failing: 0 }))).toBe(true);
  });
});

describe('readinessFails / nudgeText', () => {
  it('lists fails worst-first phrasing, matching the spec template clauses', () => {
    expect(readinessFails(row())).toEqual([
      '3 gear upgrades waiting (+24 DPS)',
      'no enchant: Chest, Boots',
      'bags short on consumables',
      '1 unspent talent point',
    ]);
  });

  it('counts any upgrade, as the board does, and prints a gain only when it was sim-checked', () => {
    expect(
      readinessFails(row({ gear_gap: { upgrades: 2, gain_dps: 10, not_sim_checked: 0, unit: 'dps' } }))[0],
    ).toBe('2 gear upgrades waiting (+10 DPS)');
    expect(
      readinessFails(row({ gear_gap: { upgrades: 1, gain_dps: 0, not_sim_checked: 1, unit: 'dps' } }))[0],
    ).toBe('1 gear upgrade waiting');
  });

  it('nudgeText prefers the API’s own officer nudge_text', () => {
    expect(nudgeText(row({ nudge_text: 'Thornhide: custom message' }))).toBe('Thornhide: custom message');
  });

  it('nudgeText falls back to a built message from readinessFails', () => {
    expect(nudgeText(row({ nudge_text: undefined }))).toBe(
      'Thornhide: 3 gear upgrades waiting (+24 DPS); no enchant: Chest, Boots; bags short on consumables; ' +
        '1 unspent talent point.',
    );
  });

  it('nudgeText reads an all-clear line when nothing fails', () => {
    const clean = row({
      gear_gap: { upgrades: 0, gain_dps: 0, not_sim_checked: 0, unit: 'dps' },
      enchants: { missing_slots: [], checked: true },
      consumables: { state: 'stocked' },
      talent_points_unspent: 0,
      nudge_text: undefined,
    });
    expect(nudgeText(clean)).toBe('Thornhide: every readiness check passes.');
  });
});
