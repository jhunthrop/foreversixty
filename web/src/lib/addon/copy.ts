// web/src/lib/addon/copy.ts
// Every string this lane adds to the site: the FSB1 decoder's refusals, the share
// panel's button, the planner's import box and the /addon page. Components import from
// here and tests assert against the same constants, the way web/src/lib/sim/copy.ts and
// web/src/lib/planner/rules.ts's `messages` already do.
//
// Voice, per design/DESIGN-SYSTEM.md: reference, not pitch. State the thing and stop.

export const addonCopy = {
  // --- codec refusals. Each names what is wrong; never a generic failure. ---
  wrongPrefix: (found: string, wanted: string): string => `That code is ${found}; this site reads ${wanted}.`,
  /** Stands in for `found` in `wrongPrefix` when the code has no prefix segment to name. */
  unlabelledCode: 'unlabelled',
  tooLong: 'That code is too long to read.',
  shortCode: 'That code is missing its talent and gear fields.',
  orderLength: 'That code’s talent order is not a whole number of points.',
  orderCell: (cell: string): string => `That code names talent cell ${cell}, which is not on any tree.`,
  unknownSlot: (slot: string): string => `That code names a slot this planner does not have: ${slot}.`,
  gearEntry: (entry: string): string => `That code has an unreadable gear entry: ${entry}.`,
  statPair: (pair: string): string => `That code has an unreadable stat: ${pair}.`,
  emptyField: (field: string): string => `That code’s ${field} field is empty.`,
  /** Names the wire field in `emptyField`; the grammar's own two mandatory fields. */
  dataBuildField: 'data build',
  classField: 'class',

  // --- bags and bank (Battle.net-sourced characters) ---
  /** Spec 2026-09-22 §3.3: shown wherever a Battle.net-sourced character's bags/bank are
   *  unavailable (no current render site yet — the bulk gear tool has no per-slot "why is
   *  this empty" note today; this is the copy ready for when it does). */
  bagsNeedAddonForBlizzard: 'Bags and bank come from the addon. Install it to include them.',

  // --- share panel ---
  copyAddonCode: 'Copy addon code',
  copiedAddonCode: 'Copied',
  addonCodeHint: 'Paste this into the game with /fs follow to see the next talent point and score your bags.',

  // --- planner import box ---
  importTitle: 'Import from addon',
  importPlaceholder: 'FS1:…',
  importAction: 'Import',
  importOrderApproximated:
    'Nothing in the game records the order a build was spent in, so this order is our best reconstruction: lowest tier first, left to right. The finished tree is exact.',
  importDropped: (names: string): string =>
    `Imported, but these could not be placed under the planner’s own rules: ${names}.`,
  importOlderBuild: (exported: string, active: string): string =>
    `That export is from data build ${exported}; this site is on ${active}. Update the addon.`,
  importWrongClass: (exported: string, current: string): string =>
    `That export is for ${exported}; this planner is on ${current}. Switch to ${exported} and import again.`,

  // --- /addon page's own paste box ---
  pasteTitle: 'Try an export',
  pastePlaceholder: 'FS1:…',
  pasteAction: 'Load',
  pasteOpenPlanner: 'Open in planner',
  pasteOpenSim: 'Open in simulator',

  // --- /addon page's signed-in save (spec 2026-09-22 §7.4) ---
  pasteSignInHint: 'Sign in to keep this character on your account.',
  pasteNameLabel: 'Character name',
  pasteRegionLabel: 'Region',
  pasteRulesetLabel: 'Ruleset',
  pasteSaveAction: 'Save to your account',
  pasteSaved: 'Saved to your account',

  // --- /addon page ---
  pageDescription:
    'Battle.net already gives your gear and talents. The addon adds your bags and bank, your professions, an in-game build guide, and refreshes the moment you log out.',
  pageNoNetwork:
    'The addon never talks to the network. Everything it knows is generated here and carried in the strings you copy.',
  installCurseForge: 'Install from CurseForge',
  installWago: 'Install from Wago Addons',
  installGitHub: 'Download the latest release',
  currentDataBuild: (build: string): string => `Current data build: ${build}`,
  flowOutTitle: 'Out of the game',
  flowOutBody: 'Type /fs export in game, copy the string, paste it into the planner’s Import from addon box.',
  flowInTitle: 'Into the game',
  flowInBody: 'Save a build here, press Copy addon code, then type /fs follow and paste it.',
  inGameTitle: 'The in-game window',
  inGameBody:
    'The addon window, tracker, talent glow, gear tab and minimap button are in beta testing in game. There is no screenshot here because we have not captured one yet.',

  // --- /setup page's three steps (spec 2026-09-25 §3.4) ---
  setupStep1Title: '1. Sign in with Battle.net',
  setupStep2Title: '2. The addon',
  setupStep3Title: '3. The companion',
  setupSignInBody: 'Imports your characters -- gear, talents and guild, refreshed nightly.',
  /** Second line under step 1: same family of copy as `accountPageCopy.noBattlenetDataForRealm`
   *  (spec 2026-09-25 §3.4) -- that key names the gap per character, on the account page,
   *  after sign-in; this one names it up front, before signing in, as a standalone fact. */
  setupSignInRealmNote: 'Blizzard does not yet serve character data for every realm type.',
  /** Same wording as `home-panel-copy.ts`'s `signInButton`; kept as its own key here so
   *  every visible string on this page comes from `addonCopy`, matching this page's own
   *  copy-module convention rather than reaching into the home page's module. */
  setupSignInAction: 'Sign in with Battle.net',
  setupCompanionBody:
    'The companion logs live from your desktop: turn on advanced combat logging, type /combatlog in game, and every fight appears under Your reports within seconds of the pull ending.',
  /** The page as a checklist (design loop, setup round): its own title, a status row
   *  under it, a done line for each step a visitor has already finished, the addon's
   *  how-it-works cards behind one disclosure, and the companion's downloads. */
  setupPageTitle: 'Get set up',
  setupStatusSignIn: 'Not signed in',
  setupStatusSignedIn: 'Signed in',
  setupStatusAddon: 'Addon export',
  setupStatusAddonDone: 'Addon export loaded',
  setupStatusCompanion: 'Companion app',
  setupSignedInLine: 'Signed in. Your characters import from Battle.net nightly.',
  setupSignedInLink: 'Your account',
  setupAddonDoneLine: 'An export is loaded; the site is pointed at that character.',
  setupHowItWorks: 'How the addon works',
  setupCompanionDownloadLead: 'Download the companion',
  setupCompanionPair: 'Pair it on your account',
  disclosureOpen: 'Show',
  disclosureClose: 'Hide',

  // --- gear panel weights ---
  scoreColumn: 'Score',
  sortByScore: 'Sort by score',
  weightsTitle: (specName: string): string => `${specName} stat weights`,
  weightsAreOpinions: 'Weights are opinions. These are ours, with the sources we took them from.',
  weightsMissing: 'No stat weights for this spec yet, so nothing is scored.',
} as const;
