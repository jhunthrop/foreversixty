import { describe, expect, it } from 'vitest';
import { raidReport } from '../../test-support/report-fixtures';
import { QUESTION_MARK_ICON, awardTooltipModel, lootIconOf } from './loot-icon';

const award = raidReport().loot[0]!.awards[0]!;

describe('loot icons', () => {
  it('uses the award icon, else the client question mark', () => {
    expect(lootIconOf({ icon: 'inv_helmet_70' })).toBe('inv_helmet_70');
    expect(lootIconOf({ icon: '' })).toBe(QUESTION_MARK_ICON);
  });

  it('builds a tooltip from the award alone, with no invented facts', () => {
    const model = awardTooltipModel({ ...award, icon: '' });
    expect(model).toMatchObject({
      id: award.item_id,
      name: award.item_name,
      quality: award.quality,
      icon: QUESTION_MARK_ICON,
      itemLevel: 0,
      stats: [],
      sourceLines: [],
    });
  });
});
