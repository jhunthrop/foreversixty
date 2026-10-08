import { describe, expect, it } from 'vitest';
import { guideBuildBlocks } from './guide-builds';

const LEVELING = 'FS1:1.60.1.70009:rogue:night-elf:32502110551501001/302303/512:';
const RAID = 'FS1:1.60.1.70009:rogue:night-elf:01532310421501/315303000015/002:';

describe('guideBuildBlocks', () => {
  it('shows the leveling and the raid build as two labelled blocks when they differ', () => {
    expect(guideBuildBlocks(LEVELING, RAID)).toEqual([
      {
        kind: 'leveling',
        id: 'leveling-build',
        label: 'Leveling build',
        code: LEVELING,
        loadAriaLabel: 'Load the leveling build in the planner',
        simAriaLabel: 'Sim the leveling build',
      },
      {
        kind: 'raid',
        id: 'raid-build',
        label: 'Raid build',
        code: RAID,
        loadAriaLabel: 'Load the raid build in the planner',
        simAriaLabel: 'Sim the raid build',
      },
    ]);
  });

  it('shows one block for both when the guide has no raid build', () => {
    expect(guideBuildBlocks(LEVELING)).toEqual([
      { kind: 'both', id: 'leveling-and-raid-build', label: 'Leveling and raid build', code: LEVELING },
    ]);
  });

  it('shows one block for both when the raid build repeats the leveling build', () => {
    expect(guideBuildBlocks(LEVELING, LEVELING)).toHaveLength(1);
  });
});
