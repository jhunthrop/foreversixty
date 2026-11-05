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
  isLeveling,
  itemLevelChip,
  itemLevelKind,
  itemLevelLabel,
  levelingLabel,
  medianCaption,
  nudgeText,
  passesEveryCheck,
  pinReadinessOwnRow,
  readinessFails,
  readinessGroups,
  readinessScore,
  sortReadinessWorstFirst,
  talentPointsKind,
  talentPointsLabel,
  worstCheck,
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

describe('cell labels', () => {
  it('gearGapLabel reads the upgrade count and gain', () => {
    expect(gearGapLabel(row())).toBe('3 upgrades · +24 DPS');
  });

  it('gearGapLabel names the real state when the API withholds gear_gap', () => {
    // Roster-only consent: the member does not share gear.
    expect(gearGapLabel(row({ consent: 'roster', gear_gap: null }))).toBe('Not shared');
    // A member whose consent IS gear is never told "no gear consent": the site simply has not checked.
    expect(gearGapLabel(row({ consent: 'gear', gear_gap: null }))).toBe('Not checked yet');
    expect(
      gearGapLabel(
        row({
          consent: 'gear',
          gear_gap: { upgrades: null, gain_dps: null, not_sim_checked: 0, unit: 'dps' },
        }),
      ),
    ).toBe('Not checked yet');
  });

  it('gearGapLabel never prints +0 DPS for upgrades no simulation produced', () => {
    expect(
      gearGapLabel(row({ gear_gap: { upgrades: 15, gain_dps: 0, not_sim_checked: 15, unit: 'dps' } })),
    ).toBe('15 upgrades · not simmed yet');
    expect(
      gearGapLabel(row({ gear_gap: { upgrades: 1, gain_dps: 0, not_sim_checked: 1, unit: 'dps' } })),
    ).toBe('1 upgrade · not simmed yet');
    expect(
      gearGapLabel(row({ gear_gap: { upgrades: 5, gain_dps: 48, not_sim_checked: 1, unit: 'dps' } })),
    ).toBe('5 upgrades · +48 DPS · 1 not simmed yet');
    expect(
      gearGapLabel(row({ gear_gap: { upgrades: 0, gain_dps: 0, not_sim_checked: 0, unit: 'dps' } })),
    ).toBe('No upgrades');
  });

  it('enchantLabel lists missing slots, capitalized', () => {
    expect(enchantLabel(row())).toBe('Chest, Boots');
  });

  it('enchantLabel reads "All enchanted" with nothing missing', () => {
    expect(enchantLabel(row({ enchants: { missing_slots: [], checked: true } }))).toBe('All enchanted');
  });

  it('enchantLabel reads "Not shared" for roster-only consent when unchecked', () => {
    expect(enchantLabel(row({ consent: 'roster', enchants: { missing_slots: [], checked: false } }))).toBe(
      'Not shared',
    );
  });

  // Live-fix round, defect 4: the four real contract slot keys (readiness_core.go's
  // enchantableSlots) humanise via the shared SLOT_LABELS table, not a bare capitalised
  // slug (main_hand -> "Main Hand", the client's casing, never "Main_hand").
  it('enchantLabel humanises the real contract slot keys via SLOT_LABELS', () => {
    const missing = row({
      enchants: { missing_slots: ['main_hand', 'chest', 'back', 'feet'], checked: true },
    });
    expect(enchantLabel(missing)).toBe('Main Hand, Chest, Back, Feet');
  });

  it('consumablesLabel reads Stocked/Short, else why not', () => {
    expect(consumablesLabel(row({ consumables: { state: 'stocked' } }))).toBe('Stocked');
    expect(consumablesLabel(row({ consumables: { state: 'short' } }))).toBe('Short');
    expect(consumablesLabel(row({ consent: 'gear', consumables: { state: 'unknown' } }))).toBe('Not shared');
    expect(consumablesLabel(row({ consent: 'roster', consumables: { state: 'unknown' } }))).toBe(
      'Not shared',
    );
    expect(consumablesLabel(row({ consent: 'gear_bags', consumables: { state: 'unknown' } }))).toBe(
      'Not checked yet',
    );
  });

  it('talentPointsLabel reads the unspent count, a real zero, or that the talents were not read', () => {
    expect(talentPointsLabel(row({ talent_points_unspent: 2 }))).toBe('2 unspent');
    expect(talentPointsLabel(row({ talent_points_unspent: 0, talent_points_read: true }))).toBe(
      'None unspent',
    );
    expect(talentPointsLabel(row({ talent_points_unspent: 0, talent_points_read: false }))).toBe(
      'Not read yet',
    );
    // A response from before the field claims nothing.
    expect(talentPointsLabel(row({ talent_points_unspent: 0 }))).toBe('—');
  });

  it('itemLevelLabel is the figure alone and dashes an unknown item level', () => {
    expect(itemLevelLabel(row({ item_level: 58, item_level_delta: -5 }))).toBe('58');
    expect(itemLevelLabel(row({ item_level: null, item_level_delta: null }))).toBe('—');
  });

  it('medianCaption says how far from the raid median, and nothing with no delta', () => {
    expect(medianCaption(row({ item_level_delta: 3 }))).toBe('3 above the raid median');
    expect(medianCaption(row({ item_level_delta: -1 }))).toBe('1 below the raid median');
    expect(medianCaption(row({ item_level_delta: 0 }))).toBe('at the raid median');
    expect(medianCaption(row({ item_level_delta: null }))).toBe('');
  });

  it('itemLevelChip reads the phone line', () => {
    expect(itemLevelChip(row({ item_level: 60, item_level_delta: -1 }))).toBe(
      'Item level 60 · 1 below median',
    );
    expect(itemLevelChip(row({ item_level: 60, item_level_delta: 0 }))).toBe('Item level 60 · at median');
    expect(itemLevelChip(row({ item_level: 60, item_level_delta: null }))).toBe('Item level 60');
    expect(itemLevelChip(row({ item_level: null, item_level_delta: null }))).toBe('—');
  });
});

describe('a healer or tank gap', () => {
  it('prints no figure and no not-simmed suffix', () => {
    const healer = row({ gear_gap: { upgrades: 4, gain_dps: 31, not_sim_checked: 2, unit: null } });
    expect(gearGapLabel(healer)).toBe('4 upgrades');
    expect(readinessFails(healer)[0]).toBe('4 gear upgrades waiting');
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

describe('worstCheck', () => {
  it('names the worst failing check in words, enchants first', () => {
    expect(worstCheck(row())).toEqual({ text: 'No enchant: Chest, Boots', more: 3 });
    expect(
      worstCheck(
        row({
          enchants: { missing_slots: ['main_hand', 'chest', 'back', 'feet'], checked: true },
          failing: 1,
        }),
      ),
    ).toEqual({ text: 'No enchant: Main Hand, Chest and 2 more slots', more: 0 });
    expect(
      worstCheck(
        row({ enchants: { missing_slots: ['main_hand', 'chest', 'back'], checked: true }, failing: 1 }),
      ),
    ).toEqual({ text: 'No enchant: Main Hand, Chest and 1 more slot', more: 0 });
  });

  it('falls through talents, consumables and upgrades', () => {
    const noEnchant = { enchants: { missing_slots: [], checked: true } };
    expect(worstCheck(row({ ...noEnchant, talent_points_unspent: 3, failing: 1 }))?.text).toBe(
      '3 talent points unspent',
    );
    expect(worstCheck(row({ ...noEnchant, talent_points_unspent: 0, failing: 2 }))).toEqual({
      text: 'Bags short on consumables',
      more: 1,
    });
    expect(
      worstCheck(
        row({ ...noEnchant, talent_points_unspent: 0, consumables: { state: 'stocked' }, failing: 1 }),
      )?.text,
    ).toBe('3 upgrades waiting');
  });

  it('is null for a clean row', () => {
    expect(worstCheck(row({ failing: 0 }))).toBeNull();
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
