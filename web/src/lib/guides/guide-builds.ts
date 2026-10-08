// web/src/lib/guides/guide-builds.ts
// The talent builds a spec guide shows under "Talents and builds": the leveling build
// (frontmatter `build:`) and, when the guide carries one, the raid build (`raidBuild:`), the
// build the raid-ready preset's level-60 entry is simmed on. A spec whose raid search found
// nothing better has no `raidBuild:` (or repeats `build:`), so the one build is both and is
// shown once under a single label rather than as two identical trees.
import { guidesCopy } from './copy';

export type GuideBuildKind = 'leveling' | 'raid' | 'both';

export interface GuideBuildBlock {
  kind: GuideBuildKind;
  /** The heading's anchor id. */
  id: string;
  /** The visible heading above this build's tree. */
  label: string;
  /** The FS1 code this block's tree, planner link and sim link are built from. */
  code: string;
  /** Accessible names for the block's planner and sim links; absent when the guide has one
   *  build, whose links keep their visible text as their name. */
  loadAriaLabel?: string;
  simAriaLabel?: string;
}

export function guideBuildBlocks(levelingCode: string, raidCode?: string): GuideBuildBlock[] {
  if (raidCode === undefined || raidCode === levelingCode) {
    return [
      {
        kind: 'both',
        id: 'leveling-and-raid-build',
        label: guidesCopy.levelingAndRaidBuild,
        code: levelingCode,
      },
    ];
  }
  return [
    {
      kind: 'leveling',
      id: 'leveling-build',
      label: guidesCopy.levelingBuild,
      code: levelingCode,
      loadAriaLabel: guidesCopy.loadLevelingBuildAriaLabel,
      simAriaLabel: guidesCopy.simLevelingBuildAriaLabel,
    },
    {
      kind: 'raid',
      id: 'raid-build',
      label: guidesCopy.raidBuild,
      code: raidCode,
      loadAriaLabel: guidesCopy.loadRaidBuildAriaLabel,
      simAriaLabel: guidesCopy.simRaidBuildAriaLabel,
    },
  ];
}
