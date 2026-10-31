// web/src/lib/planner/copy.ts
// Planner-specific copy. The planner's other components read addonCopy/simCopy for their
// strings; this is the first module that is genuinely the planner's own -- the confirm step
// Share opens into before it writes a build to a public link (one-product spec section 1).
import { itemNotInDataLine } from '../items/unknown-item-copy';

export const plannerCopy = {
  /** A slot an import names with an item this build's data does not carry; see
   *  simCopy.unknownItem, the same line on the simulator's strip. */
  itemNotSimmed: (name: string): string => `${name} · not simmed`,
  unknownItem: itemNotInDataLine,
  shareTitle: 'Share this build',
  pointOrderTitle: 'Point order',
  /** The point-order card's meta: "0 spent". */
  pointOrderMeta: (spent: number): string => `${spent} spent`,
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
  shareConfirmCopyUnsaved: 'Copy an unsaved link instead',
  copiedUnsavedLink: 'Copied',
  /** Rebuild spec §4.H, review finding 8: a plain-text summary for someone with no account
   *  and no addon, beside "Copy for the game". Always computed from real data, never a
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
  /** Fix round 2 (ux-designer review finding 2): a band load (the BiS page's "Talents in
   *  planner" link, a guide's "Open this build in the planner", or `BandCompare`'s own Load
   *  button) is never a character or an addon import -- the reconstruction note's existing
   *  "Talents loaded from a character" wording is false for it, so it gets its own, honest
   *  copy naming the band instead. */
  bandLoadedNote: (bandLabel: string): string =>
    `Loaded the ${bandLabel} build. The order shown is one legal way to reach it, lowest tier first; it is not a recommended leveling order.`,
  /** A link that carries gear but no talent ranks (a build link from a log). */
  noTalentRanksNote: 'This link carries no talent ranks.',
  /** The record carries more points than any legal order reaches: name what the planner kept. */
  keptNote: (kept: number, total: number, name: string): string =>
    `Loaded ${kept} of ${name}'s ${total} points; the rest cannot be reached in any legal order. The order is a reconstruction: the game does not record the order points were spent in.`,
} as const;

/** Rebuild spec §4.A: the header's own copy -- crest, h1, suffix. */
export const plannerHeaderCopy = {
  eyebrow: 'Build planner',
  h1: (className: string): string => `${className} talents`,
  /** Fix round 1, item 2.a (the approved mock, `day3/shots/boards/Planner.png`): one line
   *  under the h1, always the same sentence -- never computed, this page has one job. */
  description:
    "Spend your points, compare them with the level range's build, and send the result to the addon or the simulator.",
  /** Sits beside the race select: the planner has no separate FactionToggle (§4.A's own
   *  reasoning -- a second faction control would ask the visitor to agree with their race
   *  twice), and the mock names that choice explicitly rather than leaving it unsaid. */
  factionFollowsRace: 'Faction follows your race',
} as const;

/** The planner's action on the shared character header (spec 2026-10-11 section 4). */
export const plannerCharacterCardCopy = {
  /** The one send button names who it sends to: the build goes to that character's addon. */
  sendBuildTo: (name: string): string => `Send this build to ${name}`,
  sending: 'Sending',
  sent: 'Sent',
  /** A visitor with no current character cannot be sent anything: the button is the way in. */
  getTheAddon: 'Get the addon to use this build in game',
  getTheAddonHref: '/setup#addon',
  /** A /reload never picks a delivery up: the build arrives at a login from the character screen. */
  sentLine: (characterName: string): string =>
    `Sent to ${characterName}. Log out to the character screen and back in to see it. /reload will not show it.`,
  /** The second, muted line: the companion fetches the inbox every ten minutes and after a logout. */
  sentDetail: 'The companion checks every 10 minutes, and again when you log out.',
  signIn: 'Sign in to send builds to your addon.',
  signInLink: 'Sign in',
  signInHref: '/login',
  noDevice: 'Run the companion once so builds can reach your addon.',
  noDeviceLink: 'Set up the companion',
  noDeviceHref: '/setup#companion',
  failed: 'Could not send just now. Try again.',
  /** The API refused the save and said why: the first field's message, never a generic retry. */
  refused: (reason: string): string => `The site refused this build: ${reason.replace(/[.s]+$/, '')}.`,
  /** Worn pieces the planner has no item for: named by slot (the character's export carries only ids). */
  leftOut: (slots: readonly string[], names: Readonly<Record<string, string>> = {}): string => {
    const parts = slots.map((slot) => names[slot] ?? leftOutSlotPhrase[slot] ?? 'a worn piece');
    const list =
      parts.length < 3
        ? parts.join(' and ')
        : `${parts.slice(0, -1).join(', ')} and ${parts[parts.length - 1]}`;
    return `Left out: ${list} (not ${parts.length === 1 ? 'a planner item' : 'planner items'}).`;
  },
  /** An empty draft is never sent; a muted hint, not an error (nothing went wrong). */
  spendFirst: "Spend some points first, or load the level range's build.",
  /** A build with fewer points than the character has spent: say both counts, nothing more. */
  shorterThanSpend: (built: number, characterName: string, spent: number): string =>
    `This build has ${built} point${built === 1 ? '' : 's'}; ${characterName} has ${spent}.`,
  /** The build on screen is another class's: the button says so and stays off. */
  otherClass: (buildClass: string, characterName: string, characterClass: string): string =>
    `This is ${article(buildClass)} ${buildClass} build; ${characterName} is ${article(characterClass)} ${characterClass}`,
  /** Another player's build of the character's own class: it replaces what the character has. */
  sendOthersBuildTo: (name: string, replaces: number | null): string =>
    replaces === null || replaces <= 0
      ? `Send this build to ${name}`
      : `Send this build to ${name} (replaces their ${replaces} points)`,
  /** The one question before another player's build goes out. */
  confirmQuestion: (owner: string, name: string): string =>
    owner === '' ? `Send this build to ${name}?` : `Send ${owner}'s build to ${name}?`,
  confirmYes: 'Yes, send it',
  confirmNo: 'Cancel',
  /** The planner header's line for a build that is somebody else's: "Ashenveil's build · Holy Paladin". */
  originLine: (owner: string, specClass: string): string =>
    owner === '' ? `A linked build · ${specClass}` : `${owner}'s build · ${specClass}`,
  /** The simulator link under the send button. */
  simThisBuild: 'Sim this build',
} as const;

const leftOutSlotPhrase: Record<string, string> = {
  head: 'the helm in your head slot',
  neck: 'the necklace in your neck slot',
  shoulder: 'the piece in your shoulder slot',
  back: 'the cloak in your back slot',
  chest: 'the piece in your chest slot',
  wrist: 'the piece in your wrist slot',
  hands: 'the piece in your hands slot',
  waist: 'the belt in your waist slot',
  legs: 'the piece in your legs slot',
  feet: 'the piece in your feet slot',
  finger1: 'the ring in your first ring slot',
  finger2: 'the ring in your second ring slot',
  trinket1: 'the trinket in your first trinket slot',
  trinket2: 'the trinket in your second trinket slot',
  main_hand: 'the weapon in your main hand',
  off_hand: 'the item in your off hand',
  ranged: 'the piece in your ranged slot',
};

function article(word: string): string {
  return /^[aeiou]/i.test(word) ? 'an' : 'a';
}

/** Rebuild spec §4.D: `BandCompare`, the new rail panel comparing a build against its own
 *  level band's published build. */
/** "Level 60" or "Levels 20 to 29": the band's label (`60`, `20 to 29`) as a phrase that reads in a sentence. */
export const bandPhrase = (bandLabel: string): string =>
  /^\d+$/.test(bandLabel) ? `Level ${bandLabel}` : `Levels ${bandLabel}`;

const lowerBandPhrase = (bandLabel: string): string => bandPhrase(bandLabel).toLowerCase();

export const bandCompareCopy = {
  heading: (bandLabel: string): string => `${bandPhrase(bandLabel)} build`,
  /** "No level 9 build for Arms. Pick Fury or Protection." -- the level and the class's other specs come from data. */
  noLevelBuild: (level: number, specName: string, otherSpecs: readonly string[]): string => {
    const head = `No level ${level} build for ${specName}.`;
    if (otherSpecs.length === 0) return head;
    const last = otherSpecs.at(-1);
    const names = otherSpecs.length === 1 ? last : `${otherSpecs.slice(0, -1).join(', ')} or ${last}`;
    return `${head} Pick ${names}.`;
  },
  setDpsCaption: (bandLabel: string): string =>
    `the ${lowerBandPhrase(bandLabel)} range's own gear set, fully equipped — not your build`,
  differs: (n: number): string => `${n} of your points differ from this build`,
  /** The same sentence as `differs`, minus its own leading number -- `BandCompare.svelte`
   *  renders the count in its own mono/gold span and this text beside it, rather than
   *  string-surgery on `differs`' own output. */
  differsSuffix: 'of your points differ from this build',
  /** A build inside the band's, only short of it: the count leads, this text follows it. */
  toTakeSuffix: "of this build's points still to take",
  matches: 'Matches',
  /** The quiet note under the line when the build holds points past the band's own total. */
  beyondGuide: (n: number): string => `+${n} beyond the guide`,
  loadButton: (bandLabel: string): string => `Load the ${lowerBandPhrase(bandLabel)} build`,
  /** The line under the panel after a load: cut at a character's level, or the band's whole build. */
  loadedLine: (level: number, points: number, of: number): string =>
    points < of
      ? `Loaded for level ${level} · ${points} of ${of} points`
      : `Loaded at level ${level} · ${points} points`,
  seeWhatDiffers: 'See what differs',
  loadConfirmMessage:
    "Loading replaces every point you've spent with this level range's own build. Load anyway?",
  loadConfirmProceed: 'Load it',
  loadConfirmCancel: 'Keep my build',
  gearRowLabel: (bandLabel: string): string => `Gear for ${bandLabel}`,
  bestInSlotList: 'Best in slot list',
  /** Rebuild spec §4.E.4: one more line in the talent cell's own existing tooltip, once a
   *  band has loaded and this cell's rank differs from it. */
  bandRankTooltipLine: (bandLabel: string, bandRank: number, maxRank: number): string =>
    `${bandPhrase(bandLabel)} build: ${bandRank} of ${maxRank}`,
} as const;
