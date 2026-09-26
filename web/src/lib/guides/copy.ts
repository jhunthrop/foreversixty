// Every user-visible string the guides lane's own components (GuideBuildTree,
// BuildActionButtons, RacePillRow, StatPriorityPills) render, reference voice throughout
// (design/DESIGN-SYSTEM.md: state the thing and stop).
export const guidesCopy = {
  loadThisBuild: 'Load this build',
  simThisBuild: 'Sim this build',
  treeLoadFailed: 'This build did not load',
  recommendedRaceLabel: 'Recommended',
} as const;
