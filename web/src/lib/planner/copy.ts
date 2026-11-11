// web/src/lib/planner/copy.ts
// Planner-specific copy. The planner's other components read addonCopy/simCopy for their
// strings; this is the first module that is genuinely the planner's own -- the confirm step
// Share opens into before it writes a build to a public link (one-product spec section 1).
import { itemNotInDataLine } from '../items/unknown-item-copy';
import { ctaCopy } from '../addon/cta-copy';

import { grouped } from '../format/number';
import { uiCopy, withRetry } from '../ui/copy';
/** "Level 60" or "Levels 20 to 29": the band's label (`60`, `20 to 29`) as a phrase that reads in a sentence. */
export const bandPhrase = (bandLabel: string): string =>
  /^\d+$/.test(bandLabel) ? `Level ${bandLabel}` : `Levels ${bandLabel}`;

const lowerBandPhrase = (bandLabel: string): string => bandPhrase(bandLabel).toLowerCase();

export const plannerCopy = {
  /** The standalone planner's lead: what it does for the player. */
  lead: 'Plan every point up to level 60 in the order you take them, check it against the guide builds, then copy it into the game or share the link.',
  /** A slot an import names with an item this build's data does not carry; see
   *  simCopy.unknownItem, the same line on the simulator's strip. */
  itemNotSimmed: (name: string): string => `${name} · not simmed`,
  unknownItem: itemNotInDataLine,
  shareTitle: 'Share this build',
  pointOrderTitle: 'Point order',
  /** The point-order card's meta: "0 spent". */
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
  copiedUnsavedLink: uiCopy.copied,
  /** Rebuild spec §4.H, review finding 8: a plain-text summary for someone with no account
   *  and no addon, beside "Copy for the game". Always computed from real data, never a
   *  placeholder-filled template. */
  /** Share's own copy: the build as one line for chat, beside In the game's "Copy for the game". */
  copyAsText: 'Copy for chat',
  copiedText: uiCopy.copied,
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
    `Loaded the ${bandLabel} build. The order shown is one way to level into it: the main tree first, each talent as soon as its row opens.`,
  /** A link that carries gear but no talent ranks (a build link from a log). */
  noTalentRanksNote: 'This link carries no talent ranks.',
  /** The record carries more points than any legal order reaches: name what the planner kept. */
  keptNote: (kept: number, total: number, name: string): string =>
    `Loaded ${kept} of ${name}'s ${total} points; the rest cannot be reached in any legal order.`,
  /** A character's export or an addon import that lost points: the game records ranks, never the order. */
  reconstructedDroppedNote: (dropped: number): string =>
    `Talents loaded from a character, minus ${dropped} that no legal order reaches.`,
  olderBuildNote: 'This code is from an older client build; check the talents.',
  /** A code that names no character (a guide's Load this build, a shared ?code=), bug B1: the
   *  long sentence from 768, the short one on a phone, where the note is one line. */
  linkedNote: 'Opened from a link. Point order: the main tree first.',
  linkedNoteShort: 'Opened from a link.',
  /** A code for a class the planner has no talent data for. */
  noClassNote: (slug: string): string => `That code names a class this planner does not have: ${slug}.`,
  inGameTitle: 'In the game',
  /** The ControlBar's accessible name. */
  barLabel: 'Planner choices',
  /** The head answer's unit after the points spent: "51 of 51 points". */
  pointsOfUnit: (available: number): string => `of ${available} points`,
  /** The trees' panel heading. */
  talentsTitle: 'Talents',
  /** The rail's guide builds for the class on screen. */
  guideBuildsTitle: 'Guide builds',
  /** The guide build row that is the build on screen. */
  guideBuildLoaded: 'Loaded',
  /** A guide build row's distance from the build on screen: "3 points off". */
  guideBuildOff: (points: number): string => `${points} ${points === 1 ? 'point' : 'points'} off`,
  /** A guide build row's first talent the build on screen lacks: "Missing Lacerate". */
  guideBuildMissing: (talent: string): string => `Missing ${talent}`,
  /** The guide's own build names (guidesCopy.levelingBuild, raidBuild; guide-builds.test.ts holds them equal):
   *  repeated here so the planner island does not carry the guides' whole copy. */
  guideLevelingBuild: 'Level 60 build',
  guideRaidBuild: 'Raid build',
  /** In the game, under Copy for the game: the one line on using it, linked to the addon's setup step. */
  pasteHint: 'In game: /fs, then paste.',
  pasteHintHref: '/setup#addon',
  /** Point order's column heads: the level past the row's first (ruling 2026-11-09: digits read as ranks). */
  orderMarkers: Array.from({ length: 10 }, (_, step) => `+${step}`) as readonly string[],
  /** Points spent past a character's level (decision P4). */
  plannedAhead: (n: number): string => `${n} planned ahead`,
  /** Point order's line after the character's level row (decision P4). */
  levelLine: (name: string, level: number): string => `${name} is level ${level}.`,
  /** Above an empty build's Point order, the spec guide's order it previews, dimmed. */
  previewLine: (specName: string): string => `${specName} guide order`,
  /** The same line when the spec guide's order fills the sockets past the character's level. */
  aheadLine: (name: string, level: number, specName: string): string =>
    `${name} is ${level}. Then ${specName} guide order.`,
  /** A talent cell whose last points are planned past the character's level: its accessible name's tail. */
  plannedCell: (level: number): string => `planned past level ${level}`,
  /** Point order's row labels: "Levels 20 to 29", "Level 60". */
  decade: (lo: number): string => (lo >= 60 ? 'Level 60' : `Levels ${lo} to ${lo + 9}`),
  /** The next socket's tooltip in Point order. */
  nextPoint: (level: number): string => `Next point: level ${level}`,
  /** A spent socket's tooltip in Point order. */
  pointAt: (level: number, name: string, rank: number): string => `Level ${level}: ${name} ${rank}`,
  /** An empty socket in Point order. */
  levelSocket: (level: number): string => `Level ${level}`,
  noPointsYet: 'No points spent yet.',
  /** A locked cell's tooltip at a character's level cap (decision P3). */
  noPointsLeft: (level: number): string => `No points left at level ${level}.`,
  reset: 'Reset…',
  resetAsk: 'Clear every point in this build?',
  resetYes: 'Clear all points',
  resetNo: 'Keep the build',
  /** The small reset link in each tree head (Q9). */
  resetTree: 'Reset',
  resetTreeLabel: (treeName: string): string => `Reset ${treeName}`,
  sharedLinkNote: 'Read only.',
  fork: 'Edit a copy',
  /** The shared page's meta lead and its view count (spec 2026-11-08 section 8). */
  sharedMeta: 'Shared build',
  sharedViews: (views: number): string => `${grouped(views)} ${views === 1 ? 'view' : 'views'}`,
  shareTitleLabel: 'Title',
  shareTitlePlaceholder: 'Optional',
  copyLink: 'Copy link',
  copied: uiCopy.copied,
  retry: uiCopy.retry,
  loadFailedDetail: (build: string): string => `Build ${build} did not return the files the planner needs.`,
  reloadTalents: 'Reload talents',
} as const;

/** Rebuild spec §4.A: the header's own copy -- crest, h1, suffix. */
export const plannerHeaderCopy = {
  h1: (className: string): string => `${className} talents`,
} as const;

/** The planner's action on the shared character header (spec 2026-10-11 section 4). */
export const plannerCharacterCardCopy = {
  /** The one send button names who it sends to: the build goes to that character's addon. */
  sendBuildTo: (name: string): string => `Send this build to ${name}`,
  sending: 'Sending',
  sent: 'Sent',
  /** A visitor with no current character cannot be sent anything: the button is the way in. */
  getTheAddon: `${ctaCopy.getAddonAndCompanion} to use this build in game`,
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
  failed: withRetry('Could not send just now.'),
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

/** "A", "A and B", "A, B and C". */
function listOf(parts: readonly string[]): string {
  return parts.length < 3 ? parts.join(' and ') : `${parts.slice(0, -1).join(', ')} and ${parts.at(-1)}`;
}

function article(word: string): string {
  return /^[aeiou]/i.test(word) ? 'an' : 'a';
}

/** Rebuild spec §4.D: `BandCompare`, the new rail panel comparing a build against its own
 *  level band's published build. */

export const bandCompareCopy = {
  heading: (bandLabel: string): string => `${bandPhrase(bandLabel)} build`,
  /** The same panel on a shared build (ruling SHOULD, 2026-11-08): the guide build it is set against. */
  sharedHeading: 'Guide build',
  /** "No level 9 build for Arms. Pick Fury or Protection." -- the level and the class's other specs come from data. */
  noLevelBuild: (level: number, specName: string, otherSpecs: readonly string[]): string => {
    const head = `No level ${level} build for ${specName}.`;
    if (otherSpecs.length === 0) return head;
    const last = otherSpecs.at(-1);
    const names = otherSpecs.length === 1 ? last : `${otherSpecs.slice(0, -1).join(', ')} or ${last}`;
    return `${head} Pick ${names}.`;
  },
  /** The count leads in gold, this text follows it. */
  differsSuffix: 'of your points differ from this build',
  /** A build inside the band's, only short of it: the count leads, this text follows it. */
  toTakeSuffix: 'points left',
  /** The level card on a build that is its guide's: "Matches the Survival guide build". */
  matches: (specName: string): string =>
    specName === '' ? 'Matches' : `Matches the ${specName} guide build`,
  /** The quiet note under the line when the build holds points past the band's own total. */
  beyondGuide: (n: number): string => `+${n} beyond the build`,
  /** The difference named, not counted (ruling Q1): "Differs by 2 points: this build takes
   *  Rapid Killing 2/2; the level 60 build takes Deterrence and Counterattack." */
  differsNamed: (
    count: number,
    mine: readonly string[],
    theirs: readonly string[],
    bandLabel: string,
  ): string => {
    const sides = [
      mine.length > 0 ? `this build takes ${listOf(mine)}` : '',
      theirs.length > 0 ? `the ${lowerBandPhrase(bandLabel)} build takes ${listOf(theirs)}` : '',
    ].filter(Boolean);
    return `Differs by ${count} point${count === 1 ? '' : 's'}: ${sides.join('; ')}.`;
  },
  /** A band build that covers fewer points than the character has (ruling MUST 3): never "Matches". */
  coversFewer: (covered: number, total: number): string => `${covered} of ${total} points`,
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
  bestInSlotList: 'Best in slot list',
  /** Rebuild spec §4.E.4: one more line in the talent cell's own existing tooltip, once a
   *  band has loaded and this cell's rank differs from it. */
  bandRankTooltipLine: (bandLabel: string, bandRank: number, maxRank: number): string =>
    `${bandPhrase(bandLabel)} build: ${bandRank} of ${maxRank}`,
} as const;
