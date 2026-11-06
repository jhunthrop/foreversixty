import { describe, expect, it } from 'vitest';
import { raidReport } from '../../test-support/report-fixtures';
import { awardTooltipModel, lootIconOf, lootInitials } from './loot-icon';

const award = raidReport().loot[0]!.awards[0]!;

describe('loot icons', () => {
  it('uses the award icon, and nothing (the neutral tile) when there is none', () => {
    expect(lootIconOf({ icon: 'inv_misc_head_dragon_01' })).toBe('inv_misc_head_dragon_01');
    expect(lootIconOf({ icon: '' })).toBeNull();
    expect(lootInitials('Head of Onyxia')).toBe('He');
  });

  it('builds a tooltip from the award alone, with no invented facts', () => {
    const model = awardTooltipModel({ ...award, icon: 'inv_misc_head_dragon_01' });
    expect(model).toMatchObject({
      id: award.item_id,
      name: award.item_name,
      quality: award.quality,
      icon: 'inv_misc_head_dragon_01',
      itemLevel: 0,
      stats: [],
      sourceLines: [],
    });
  });
});
