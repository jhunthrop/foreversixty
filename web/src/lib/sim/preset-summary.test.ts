import { beforeAll, describe, expect, it } from 'vitest';
import { EMPTY_BUFF_NAMES } from './buff-names';
import { presetSummary } from './preset-summary';
import { presetConsumables, RAID_BUFFS } from './settings';
import { loadRaidConsumes } from './raid-consumes';

// The Raid-buffed consumables are a chunk loaded on demand (raid-consumes.ts).
beforeAll(() => loadRaidConsumes());

describe('presetSummary', () => {
  it('is empty for solo, which applies nothing', () => {
    expect(presetSummary('solo', 'warrior-fury', 60, null)).toEqual([]);
  });

  it('carries every raid-buffed id, grouped, none lost and none invented', () => {
    const groups = presetSummary('raid-buffed', 'warrior-fury', 60, null);
    const ids = groups.flatMap((group) => group.rows.map((row) => row.id));
    expect(new Set(ids)).toEqual(
      new Set([...RAID_BUFFS, ...presetConsumables('raid-buffed', 'warrior-fury', 60)]),
    );
    // Every group this task-4-brief.md's motivating finding named as under-populated
    // (raid buffs, on the target, weapon oils and stones) actually has rows now.
    expect(groups.some((group) => group.group === 'raid-buffs')).toBe(true);
    expect(groups.some((group) => group.group === 'debuffs')).toBe(true);
    expect(groups.some((group) => group.group === 'weapon-imbue')).toBe(true);
  });

  it('carries each spec’s own published consumables: a rogue’s poisons, never a mage’s oil', () => {
    const rogue = presetSummary('raid-buffed', 'rogue-combat', 60, null);
    const mage = presetSummary('raid-buffed', 'mage-fire', 60, null);
    const rogueIds = new Set(rogue.flatMap((group) => group.rows.map((row) => row.id)));
    const mageIds = new Set(mage.flatMap((group) => group.rows.map((row) => row.id)));
    for (const id of presetConsumables('raid-buffed', 'rogue-combat', 60))
      expect(rogueIds.has(id), id).toBe(true);
    for (const id of presetConsumables('raid-buffed', 'mage-fire', 60))
      expect(mageIds.has(id), id).toBe(true);
    expect(rogueIds.has('main_hand_imbue:deadly_poison')).toBe(true);
    expect(rogueIds.has('main_hand_imbue:brilliant_wizard_oil')).toBe(false);
    // The buff half never splits: both answers carry the identical raid list.
    for (const id of RAID_BUFFS) {
      expect(rogueIds.has(id)).toBe(true);
      expect(mageIds.has(id)).toBe(true);
    }
    // A poisoned hand's stone is filed with the weapon oils and stones, not the potions.
    const stones = rogue.find((group) => group.group === 'weapon-imbue')?.rows.map((row) => row.id) ?? [];
    expect(stones).toContain('main_hand_stone:elemental_sharpening_stone');
  });

  /**
   * The raid-leader review's own finding: a bare engine id in the UI. buffLabel already
   * humanises an id the build's table has no row for (buff-names.test.ts covers that rule
   * generically); this names the gap for the two ids the brief's own acceptance test reads
   * off the page, over a names table exactly as sparse as this repo's fixture data really
   * is (src/fixtures/planner/simbuffs.json has three rows, and sunder_armor and
   * trueshot_aura are not among them).
   */
  it('never prints a bare id, even for a build whose table has no row for it', () => {
    const groups = presetSummary('raid-buffed', 'warrior-fury', 60, EMPTY_BUFF_NAMES);
    for (const group of groups) {
      for (const row of group.rows) {
        expect(row.label, row.id).not.toBe(row.id);
        expect(row.label, row.id).not.toMatch(/_/);
      }
    }
    const labels = groups.flatMap((group) => group.rows.map((row) => row.label));
    expect(labels).toContain('Sunder armor');
    expect(labels).toContain('Trueshot aura');
  });
});
