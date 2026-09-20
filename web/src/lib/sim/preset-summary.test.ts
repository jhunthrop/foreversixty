import { describe, expect, it } from 'vitest';
import { EMPTY_BUFF_NAMES } from './buff-names';
import { presetSummary } from './preset-summary';
import { CASTER_CONSUMABLES, PHYSICAL_CONSUMABLES, RAID_BUFFS } from './settings';

describe('presetSummary', () => {
  it('is empty for solo, which applies nothing', () => {
    expect(presetSummary('solo', 'attack_power', null)).toEqual([]);
  });

  it('carries every raid-buffed id, grouped, none lost and none invented', () => {
    const groups = presetSummary('raid-buffed', 'attack_power', null);
    const ids = groups.flatMap((group) => group.rows.map((row) => row.id));
    expect(new Set(ids)).toEqual(new Set([...RAID_BUFFS, ...PHYSICAL_CONSUMABLES]));
    // Every group this task-4-brief.md's motivating finding named as under-populated
    // (raid buffs, on the target, weapon oils and stones) actually has rows now.
    expect(groups.some((group) => group.group === 'raid-buffs')).toBe(true);
    expect(groups.some((group) => group.group === 'debuffs')).toBe(true);
    expect(groups.some((group) => group.group === 'weapon-imbue')).toBe(true);
  });

  it('splits the consumable half by reference_stat, physical from caster', () => {
    const physical = presetSummary('raid-buffed', 'attack_power', null);
    const caster = presetSummary('raid-buffed', 'spell_power', null);
    const physicalIds = new Set(physical.flatMap((group) => group.rows.map((row) => row.id)));
    const casterIds = new Set(caster.flatMap((group) => group.rows.map((row) => row.id)));
    for (const id of PHYSICAL_CONSUMABLES) expect(physicalIds.has(id)).toBe(true);
    for (const id of CASTER_CONSUMABLES) expect(casterIds.has(id)).toBe(true);
    for (const id of CASTER_CONSUMABLES) expect(physicalIds.has(id)).toBe(false);
    // The buff half never splits (Ruling 2): both answers carry the identical 54 ids.
    for (const id of RAID_BUFFS) {
      expect(physicalIds.has(id)).toBe(true);
      expect(casterIds.has(id)).toBe(true);
    }
  });

  /**
   * The raid-leader review's own finding: a bare engine id in the UI. buffLabel already
   * humanises an id the build's table has no row for (buff-names.test.ts covers that rule
   * generically); this names the gap for the two ids the brief's own acceptance test reads
   * off the page, over a names table exactly as sparse as this repo's fixture data really
   * is (src/fixtures/planner/simbuffs.json has three rows, and sunder_armor and
   * songflower_serenade are not among them).
   */
  it('never prints a bare id, even for a build whose table has no row for it', () => {
    const groups = presetSummary('raid-buffed', 'attack_power', EMPTY_BUFF_NAMES);
    for (const group of groups) {
      for (const row of group.rows) {
        expect(row.label, row.id).not.toBe(row.id);
        expect(row.label, row.id).not.toMatch(/_/);
      }
    }
    const labels = groups.flatMap((group) => group.rows.map((row) => row.label));
    expect(labels).toContain('Sunder armor');
    expect(labels).toContain('Songflower serenade');
  });
});
