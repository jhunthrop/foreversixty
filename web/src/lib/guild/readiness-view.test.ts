// web/src/lib/guild/readiness-view.test.ts
import { describe, expect, it } from 'vitest';
import type { GuildReadinessRow } from './api';
import {
  consumablesKind,
  consumablesLabel,
  enchantKind,
  enchantLabel,
  gearGapKind,
  gearGapLabel,
  itemLevelKind,
  itemLevelLabel,
  nudgeText,
  pinReadinessOwnRow,
  readinessFails,
  readinessScore,
  sortReadinessWorstFirst,
  talentPointsKind,
  talentPointsLabel,
} from './readiness-view';

function row(overrides: Partial<GuildReadinessRow> = {}): GuildReadinessRow {
  return {
    character_key: 'us/pvp/thornhide',
    name: 'Thornhide',
    class: 'druid',
    spec: 'Feral',
    consent: 'gear_bags',
    gear_gap: { upgrades: 3, gain_dps: 23.8, not_sim_checked: 0 },
    enchants: { missing_slots: ['chest', 'boots'], checked: true },
    consumables: { state: 'short' },
    talent_points_unspent: 1,
    item_level: 58,
    item_level_delta: -5,
    logged_at: '2026-12-20T20:00:00Z',
    failing: 4,
    ...overrides,
  };
}

describe('readinessScore / sortReadinessWorstFirst', () => {
  it('sorts failCount*100 + gearGainDps worst-first', () => {
    const low = row({
      character_key: 'a',
      failing: 1,
      gear_gap: { upgrades: 1, gain_dps: 10, not_sim_checked: 0 },
    });
    const high = row({
      character_key: 'b',
      failing: 3,
      gear_gap: { upgrades: 3, gain_dps: 5, not_sim_checked: 0 },
    });
    expect(readinessScore(high)).toBeGreaterThan(readinessScore(low));
    expect(sortReadinessWorstFirst([low, high]).map((r) => r.character_key)).toEqual(['b', 'a']);
  });

  it('breaks ties on the bigger gear gain', () => {
    const smaller = row({
      character_key: 'a',
      failing: 2,
      gear_gap: { upgrades: 2, gain_dps: 10, not_sim_checked: 0 },
    });
    const bigger = row({
      character_key: 'b',
      failing: 2,
      gear_gap: { upgrades: 2, gain_dps: 40, not_sim_checked: 0 },
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

describe('cell labels', () => {
  it('gearGapLabel reads the upgrade count and gain', () => {
    expect(gearGapLabel(row())).toBe('3 upgrades · +24 DPS');
  });

  it('gearGapLabel reads "no gear consent" when the API withholds gear_gap', () => {
    expect(gearGapLabel(row({ gear_gap: null }))).toBe('no gear consent');
    expect(gearGapLabel(row({ gear_gap: { upgrades: null, gain_dps: null, not_sim_checked: 0 } }))).toBe(
      'no gear consent',
    );
  });

  it('enchantLabel lists missing slots, capitalized', () => {
    expect(enchantLabel(row())).toBe('Chest, Boots');
  });

  it('enchantLabel reads "All enchanted" with nothing missing', () => {
    expect(enchantLabel(row({ enchants: { missing_slots: [], checked: true } }))).toBe('All enchanted');
  });

  it('enchantLabel reads "consent needed" when unchecked', () => {
    expect(enchantLabel(row({ enchants: { missing_slots: [], checked: false } }))).toBe('consent needed');
  });

  it('consumablesLabel reads Stocked/Short/consent needed', () => {
    expect(consumablesLabel(row({ consumables: { state: 'stocked' } }))).toBe('Stocked');
    expect(consumablesLabel(row({ consumables: { state: 'short' } }))).toBe('Short');
    expect(consumablesLabel(row({ consumables: { state: 'unknown' } }))).toBe('consent needed');
  });

  it('talentPointsLabel reads the unspent count or a bare 0', () => {
    expect(talentPointsLabel(row({ talent_points_unspent: 2 }))).toBe('2 unspent');
    expect(talentPointsLabel(row({ talent_points_unspent: 0 }))).toBe('0');
  });

  it('itemLevelLabel signs the delta and dashes an unknown item level', () => {
    expect(itemLevelLabel(row({ item_level: 58, item_level_delta: -5 }))).toBe('58 (-5)');
    expect(itemLevelLabel(row({ item_level: 70, item_level_delta: 5 }))).toBe('70 (+5)');
    expect(itemLevelLabel(row({ item_level: null, item_level_delta: null }))).toBe('—');
  });
});

describe('readinessFails / nudgeText', () => {
  it('lists fails worst-first phrasing, matching the spec template clauses', () => {
    expect(readinessFails(row())).toEqual([
      '3 gear upgrades waiting (24 DPS)',
      'no enchant: Chest, Boots',
      'bags short on consumables',
      '1 unspent talent point',
    ]);
  });

  it('omits an upgrade clause under 3 upgrades (spec §4.A.2 threshold)', () => {
    const fails = readinessFails(row({ gear_gap: { upgrades: 2, gain_dps: 10, not_sim_checked: 0 } }));
    expect(fails.some((fail) => fail.includes('gear upgrades waiting'))).toBe(false);
  });

  it('nudgeText prefers the API’s own officer nudge_text', () => {
    expect(nudgeText(row({ nudge_text: 'Thornhide: custom message' }))).toBe('Thornhide: custom message');
  });

  it('nudgeText falls back to a built message from readinessFails', () => {
    expect(nudgeText(row({ nudge_text: undefined }))).toBe(
      'Thornhide: 3 gear upgrades waiting (24 DPS), no enchant: Chest, Boots, bags short on consumables, ' +
        '1 unspent talent point -- check before Thursday.',
    );
  });

  it('nudgeText reads an all-clear line when nothing fails', () => {
    const clean = row({
      gear_gap: { upgrades: 0, gain_dps: 0, not_sim_checked: 0 },
      enchants: { missing_slots: [], checked: true },
      consumables: { state: 'stocked' },
      talent_points_unspent: 0,
      nudge_text: undefined,
    });
    expect(nudgeText(clean)).toBe('Thornhide: every readiness check passes.');
  });
});

describe('cell kinds (styling)', () => {
  it('gearGapKind reads consent without gear_gap, neutral otherwise', () => {
    expect(gearGapKind(row({ gear_gap: null }))).toBe('consent');
    expect(gearGapKind(row())).toBe('neutral');
  });

  it('enchantKind reads consent/ok/warn', () => {
    expect(enchantKind(row({ enchants: { missing_slots: [], checked: false } }))).toBe('consent');
    expect(enchantKind(row({ enchants: { missing_slots: [], checked: true } }))).toBe('ok');
    expect(enchantKind(row({ enchants: { missing_slots: ['chest'], checked: true } }))).toBe('warn');
  });

  it('consumablesKind reads consent/ok/warn', () => {
    expect(consumablesKind(row({ consumables: { state: 'unknown' } }))).toBe('consent');
    expect(consumablesKind(row({ consumables: { state: 'stocked' } }))).toBe('ok');
    expect(consumablesKind(row({ consumables: { state: 'short' } }))).toBe('warn');
  });

  it('talentPointsKind warns on any unspent point, neutral at 0', () => {
    expect(talentPointsKind(row({ talent_points_unspent: 1 }))).toBe('warn');
    expect(talentPointsKind(row({ talent_points_unspent: 0 }))).toBe('neutral');
  });

  it('itemLevelKind reads consent/ok/warn', () => {
    expect(itemLevelKind(row({ item_level: null, item_level_delta: null }))).toBe('consent');
    expect(itemLevelKind(row({ item_level: 70, item_level_delta: 5 }))).toBe('ok');
    expect(itemLevelKind(row({ item_level: 55, item_level_delta: -5 }))).toBe('warn');
  });
});
