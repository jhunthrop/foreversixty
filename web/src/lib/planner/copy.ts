// web/src/lib/planner/copy.ts
// Planner-specific copy. The planner's other components read addonCopy/simCopy for their
// strings; this is the first module that is genuinely the planner's own -- the confirm step
// Share opens into before it writes a build to a public link (one-product spec section 1).
export const plannerCopy = {
  share: 'Share',
  shareConfirmTitle: 'Share this build',
  shareConfirmIntro:
    'Sharing posts the following to a link anyone can open. Once created, that link cannot be edited:',
  shareFieldClassRace: 'Class and race',
  shareFieldTalents: 'Talent order',
  shareFieldGear: 'Gear',
  shareFieldTitle: (title: string): string => `Title: "${title}"`,
  shareFieldSim: 'A simmed DPS result for the card',
  shareConfirmProceed: 'Share anyway',
  shareConfirmCancel: 'Cancel',
  shareConfirmCopyCode: 'Copy addon code instead',
  shareConfirmCopyUnsaved: 'Copy an unsaved link instead',
  copiedUnsavedLink: 'Copied',
} as const;
