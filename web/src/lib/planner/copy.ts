// web/src/lib/planner/copy.ts
// Planner-specific copy. The planner's other components read addonCopy/simCopy for their
// strings; this is the first module that is genuinely the planner's own -- the confirm step
// Share opens into before it writes a build to a public link (one-product spec section 1).
export const plannerCopy = {
  /** A slot an import names with an item this build's data does not carry; see
   *  simCopy.unknownItem, the same line on the simulator's strip. */
  itemNotSimmed: (name: string): string => `${name} · not simmed`,
  unknownItem: (id: number): string => `Unknown item ${id} · not in our data yet`,
  shareTitle: 'Share this build',
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
  /** Rebuild spec §4.H, review finding 8: a plain-text summary for someone with no account
   *  and no addon, beside "Copy addon code". Always computed from real data, never a
   *  placeholder-filled template. */
  copyAsText: 'Copy as text',
  copiedText: 'Copied',
  shareTextSummary: (
    raceName: string,
    className: string,
    trees: readonly { name: string; points: number }[],
    level: number,
  ): string =>
    `${raceName} ${className} · ${trees.map((tree) => `${tree.name} ${tree.points}`).join('/')} · Level ${level}`,
} as const;

/** Rebuild spec §4.A: the header's own copy -- crest, h1, suffix. */
export const plannerHeaderCopy = {
  eyebrow: 'Build planner',
  h1: (className: string): string => `${className} talents`,
  /** Fix round 1, item 2.a (the approved mock, `day3/shots/boards/Planner.png`): one line
   *  under the h1, always the same sentence -- never computed, this page has one job. */
  description:
    "Spend your points, compare them with the band's build, and send the result to the addon or the simulator.",
  /** Sits beside the race select: the planner has no separate FactionToggle (§4.A's own
   *  reasoning -- a second faction control would ask the visitor to agree with their race
   *  twice), and the mock names that choice explicitly rather than leaving it unsaid. */
  factionFollowsRace: 'Faction follows your race',
} as const;

/** Rebuild spec §4.B: the Character card paired with the header, signed in only. */
export const plannerCharacterCardCopy = {
  label: 'Your character',
  sendBuildToAddon: 'Send this build to the addon',
  loadError: 'Your character did not load.',
} as const;

/** Rebuild spec §4.D: `BandCompare`, the new rail panel comparing a build against its own
 *  level band's published build. */
export const bandCompareCopy = {
  heading: (bandLabel: string): string => `${bandLabel} build`,
  headingAside: "your level's band",
  noLevelingList: (specName: string): string => `No leveling list for ${specName} yet.`,
  setDpsCaption: (bandLabel: string): string =>
    `the ${bandLabel} band's own gear set, fully equipped — not your build`,
  differs: (n: number): string => `${n} of your points differ from this build`,
  /** The same sentence as `differs`, minus its own leading number -- `BandCompare.svelte`
   *  renders the count in its own mono/gold span and this text beside it, rather than
   *  string-surgery on `differs`' own output. */
  differsSuffix: 'of your points differ from this build',
  matches: 'Matches',
  loadButton: (bandLabel: string): string => `Load the ${bandLabel} build`,
  seeWhatDiffers: 'See what differs',
  loadConfirmMessage: "Loading replaces every point you've spent with this band's own build. Load anyway?",
  loadConfirmProceed: 'Load it',
  loadConfirmCancel: 'Keep my build',
  gearRowLabel: (bandLabel: string): string => `Gear for ${bandLabel}`,
  bestInSlotList: 'Best in slot list',
  /** Rebuild spec §4.E.4: one more line in the talent cell's own existing tooltip, once a
   *  band has loaded and this cell's rank differs from it. */
  bandRankTooltipLine: (bandLabel: string, bandRank: number, maxRank: number): string =>
    `${bandLabel} build: ${bandRank} of ${maxRank}`,
} as const;
