// web/src/lib/addon/copy.ts
// Every string this lane adds to the site: the FSB1 decoder's refusals, the share
// panel's button, the planner's import box and the /addon page. Components import from
// here and tests assert against the same constants, the way web/src/lib/sim/copy.ts and
// web/src/lib/planner/rules.ts's `messages` already do.
//
// Voice, per design/DESIGN-SYSTEM.md: reference, not pitch. State the thing and stop.

export interface AddonSurfaceImage {
  src: string;
  width: number;
  height: number;
  alt: string;
}
export interface AddonSurface {
  section: string;
  title: string;
  body: string;
  images: readonly AddonSurfaceImage[];
}

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

  // --- "Copy for the game" (CopyForGame.svelte): the planner's share panel, Top Gear and a
  //     saved sim all copy the same FSB1 build code through the same button. ---
  copyForGame: 'Copy for the game',
  /** Stands for 3 s after a copy, then the button reads `copyForGame` again. */
  copiedForGame: 'Copied, paste it in game with /fs',
  /** The clipboard refused (permission, an insecure page, no clipboard at all): the code is
   *  shown in a read-only field under this line instead. */
  copyFailed: 'Could not copy; select the code below instead',
  /** The accessible name of that read-only field. */
  buildCodeLabel: 'Build code',
  addonCodeHint: 'Paste this in game with /fs; the chip and the talent window follow it from then on.',

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
  /** The standalone planner switches class itself and imports once those talents load. */
  importSwitchingClass: (exported: string): string =>
    `That export is for ${exported}; switching the planner to it.`,

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
  /** Said under the save button when it is pressed with that field still empty, instead of
   *  a disabled button (which looks the same as a live one and read as broken). */
  pasteNeedsName: 'Give the character a name first.',
  pasteNeedsRegion: 'Pick the region first.',
  pasteNeedsRuleset: 'Pick the ruleset first.',
  pasteSaved: 'Saved to your account',

  // --- /addon page ---
  pageDescription:
    'The addon brings your character here: gear and talents, bags and bank, your professions, an in-game build guide, and a fresh export the moment you log out.',
  pageNoNetwork:
    'The addon never talks to the network. Everything it knows is generated here and carried in the strings you copy.',
  installCurseForge: 'Install from CurseForge',
  installWago: 'Install from Wago Addons',
  installGitHub: 'Download the latest release',
  currentDataBuild: (build: string): string => `Current data build: ${build}`,
  flowOutTitle: 'Out of the game',
  /** The addon's copy mode (L8.5 ruling 5): `/fs copy` opens it with the box selected, and
   *  Ctrl+C is the action; there is no button, and the old export command is gone. */
  flowOutBody:
    'In game, type /fs copy, press Ctrl+C, then paste it into the planner’s Import from addon box.',
  flowInTitle: 'Into the game',
  flowInBody: 'Save a build here, press Copy for the game, then type /fs in game and paste it.',
  inGameTitle: 'The in-game window',
  /** L8.3: the surfaces now have their own section with captures from the client, so this
   *  card points at it rather than saying none exists. */
  inGameBody:
    'The addon is in beta testing in game. Every surface it draws is listed above under What it does in the game, with a capture from the client where we have one.',

  // --- /setup page's three steps (spec 2026-09-25 §3.4) ---
  setupStep1Title: '1. Sign in with Battle.net',
  setupStep2Title: '2. The addon',
  setupStep3Title: '3. The companion',
  setupSignInBody:
    'Saves your builds, sims and logs to an account. Character import waits on Blizzard serving Forever data.',
  /** Second line under step 1: same family of copy as `accountPageCopy.noBattlenetDataForRealm`
   *  (spec 2026-09-25 §3.4) -- that key names the gap per character, on the account page,
   *  after sign-in; this one names it up front, before signing in, as a standalone fact. */
  setupSignInRealmNote: 'Until then, the addon below adds your characters.',
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
  setupSignedInLine: 'Signed in. Paste an addon export to add your characters.',
  setupSignedInLink: 'Your account',
  setupAddonDoneLine: 'An export is loaded; the site is pointed at that character.',
  setupHowItWorks: 'How the addon works',
  setupCompanionDownloadLead: 'Download the companion',
  setupCompanionPair: 'Pair it on your account',
  disclosureOpen: 'Show',
  disclosureClose: 'Hide',

  // --- /setup: what the addon does in the game (Task L8.3) ---
  /** One row per surface of the addon-v2 spec's §4, in its order; `title` is the spec's own
   *  heading and `body` one plain sentence from that section. `images` are crops of real
   *  client captures (docs/superpowers/reviews/week1/), each alt text saying only what is on
   *  screen in it; a surface with no capture yet (messaging) has none, and says nothing
   *  about it. `width`/`height` are the files' own pixel sizes, for the fixed aspect box. */
  surfacesTitle: 'What it does in the game',
  surfaces: [
    {
      section: '4.1',
      title: 'The chip',
      body: 'A small icon and line beside the minimap that names your next step: the talent to take, a bag item to equip, loot here for you, or what a raid check found missing.',
      images: [
        {
          src: '/addon/chip-tooltip-legend.jpg',
          width: 1019,
          height: 360,
          alt: 'The chip beside the minimap with its hover open: “Forever Sixty · Beast Mastery / Marksmanship build”, “Next talent at 25: Lethal Attacks”, the legend lines “Green arrow: a bag item beats what you wear.” and “Gold corner: next upgrade to chase.”, then the click, right-click and shift-drag hints.',
        },
      ],
    },
    {
      section: '4.2',
      title: 'Talent marks',
      body: 'In the talent window, a gold ring marks the next talent, every talent your build takes carries its order number, and a strip along the bottom counts your progress and any points spent off plan.',
      images: [
        {
          src: '/addon/talent-marks-trees.jpg',
          width: 504,
          height: 298,
          alt:
            'A hunter’s talent window, Beast Mastery, Marksmanship and Survival, with small order marks on the planned talents, one ringed in gold in Marksmanship, and the strip along the bottom: “Next: Lethal Attacks (1 of 5) · 12 of 51 points · 3 off plan”, the last part in red.',
        },
      ],
    },
    {
      section: '4.3',
      title: 'Item tooltip',
      body: 'Every item tooltip gains a Forever Sixty section that says whether the item is an upgrade and by how much, the next upgrade for that slot and where it comes from, and two alternatives.',
      images: [
        {
          src: '/addon/item-tooltip-ranged.jpg',
          width: 910,
          height: 615,
          alt: 'A level 24 hunter’s worn Short Ash Bow tooltip on the character panel, ending in the Forever Sixty section: “Forever Sixty · Beast Mastery · Estimate”, “Equipped”, “Next upgrade: Ranger Bow +3.5 DPS”, “(+1 Agi, +3.4 weapon DPS)”, “Planned for your build”, “World drop · check the AH”, and “Also good:” Lil Timmy’s Peashooter +2.2 (needs bullets) and Cracked Blacksmith Hammer +0.8, with the Ranger Bow’s own tooltip beside it.',
        },
      ],
    },
    {
      section: '4.4',
      title: 'Paper doll, bags, quest rewards, rolls',
      body: 'A green arrow marks a paper-doll slot and the bag item that beats it, a gold corner marks a slot with an upgrade to chase, the quest reward your spec prefers is marked, and the roll frame shows the tooltip’s verdict beside the item.',
      images: [
        {
          src: '/addon/bag-arrow.jpg',
          width: 560,
          height: 300,
          alt:
            'The top of the Combined Backpack with a small green arrow on the top-right corner of an item in the second row of the backpack: a bag item that beats what the character wears.',
        },
      ],
    },
    {
      section: '4.5',
      title: 'The "Forever Sixty" character-panel tab',
      body: 'A tab on the character panel with your next gear targets, your rotation for your level and your build; at 60 it adds your sim and your last raid.',
      images: [
        {
          src: '/addon/character-panel-tab.jpg',
          width: 580,
          height: 796,
          alt: 'The Forever Sixty tab for a level 24 hunter: Next targets lists Ranger Bow +3.5 DPS (Ranged, World drop) and Blackened Defias Belt +1.1 DPS (Waist, The Deadmines, Captain Greenskin, 30% drop); Rotation lists Aimed Shot, Multi-Shot, Serpent Sting and Arcane Shot with their conditions, and Rapid Fire dimmed at level 26; Build reads “Plan: Beast Mastery 31 / Marksmanship 20”, “12 of 15 points on plan” and “3 off plan” in red, over the buttons Use a different build and Copy my character.',
        },
      ],
    },
    {
      section: '4.6',
      title: 'The player tooltip and the score',
      body: 'Every player tooltip carries that player’s Forever Score, and in LFG, Trade, guild chat and whispers a sender with a score gets an FS badge before their message.',
      images: [
        {
          src: '/addon/player-tooltip-score.jpg',
          width: 524,
          height: 307,
          alt: 'A player tooltip for Bow Jackzon, level 24 hunter, with the line “Forever Score 1240 · Rare · as of 13 Oct” in blue and the muted hint “Shift: score details · target them, then /fs link” under the client’s own lines.',
        },
        {
          src: '/addon/chat-badge-guild.jpg',
          width: 860,
          height: 90,
          alt: 'A guild chat line, “[Guild] [Bow Jackzon]: FS 1240 · (addon test 3)”, with “FS 1240” in blue and a muted dot before the message.',
        },
      ],
    },
    {
      section: '4.7',
      title: 'Guild tab',
      body: 'A tab on the guild frame lists the roster’s level, spec, gear, Forever Score, attendance and gear gap, sortable, live for members who run the addon.',
      images: [
        {
          src: '/addon/guild-tab.jpg',
          width: 728,
          height: 612,
          alt:
            'The Forever Sixty guild tab: “Online only” unticked, “97 of 97 shown · 2 of 3 online run Forever Sixty · 1-20 of 97”, and columns Name, Level, Spec, Gear, Score (sorted), Attendance and Gap, with scored rows for Belee Dat (Frost, 2310), Onearm Lopen (Arms, 1720, gap 14 DPS) and Bow Jackzon (Beast Mastery, 1240, gap 0 DPS), then two dimmed rows and “no logs” for the rest.',
        },
      ],
    },
    {
      section: '4.8',
      title: 'Group panel',
      body: 'A tab on the raid frame shows each member’s spec, level, gear gap and score before the first pull, and on a ready check lists who lacks flask, food, weapon imbue or a class buff.',
      images: [
        {
          src: '/addon/group-panel-solo.jpg',
          width: 883,
          height: 157,
          alt: 'The Forever Sixty group panel while solo: “Not in a group.” then, muted, “A party or raid shows each member’s spec, gear gap, score and buffs. /fs ready tries the check on you.”',
        },
      ],
    },
    {
      section: '4.9',
      title: 'Guild and raid messaging',
      body: 'Addons in your guild, raid or party share level, spec, gear gap and build name, pass builds to each other, and tell the master looter what a drop is worth to each member.',
      images: [],
    },
    {
      section: '4.10',
      title: 'Build dialog',
      body: 'The addon’s one window of its own, where you paste a build from this site to follow it or copy your character to paste here.',
      images: [
        {
          src: '/addon/build-dialog-paste.jpg',
          width: 770,
          height: 567,
          alt:
            'The build dialog for a hunter on Paste a build: “Following: Beast Mastery / Marksmanship build · Hunter · 51 po…”, “Next talent: Take Lethal Attacks (Marksmanship, 1 of 5)”, and the empty box “Paste a build code from foreversixty.gg (Ctrl+V)”, with “Tab to switch” at the bottom.',
        },
        {
          src: '/addon/build-dialog-copy.jpg',
          width: 633,
          height: 464,
          alt: 'The build dialog on Copy my character: “Selected. Press Ctrl+C, then paste it on foreversixty.gg.” over the selected export, then “The site has your character from 10 Oct 20:00. Yours is newer.”, “Sim done from an earlier copy · a new sim is queued” and “Your next step: Take Deflection (1 of 2)”.',
        },
      ],
    },
    {
      section: '4.11',
      title: 'Settings',
      body: 'Eight switches on the client’s Settings page, each with one line saying what it does, then Reset positions and the command list.',
      images: [
        {
          src: '/addon/settings-switches.jpg',
          width: 1166,
          height: 1036,
          alt:
            'The Forever Sixty settings page with its eight switches ticked, each with its hint: Show the chip, Hide the chip in combat, Marks on the talent window, Show the minimap button, Gear tips on item tooltips, Score badges in chat, Share my build with my guild and Save my character for the companion app, then the Reset positions button and the start of the Commands list.',
        },
      ],
    },
  ] as readonly AddonSurface[],
  /** The addon's own list (addon/ForeverSixty/Locale.lua `settingsCommands`, which /fs help,
   *  an unknown command and the Settings page all print), split at its first ": ". */
  commandsTitle: 'Commands',
  commands: [
    { command: '/fs', description: 'paste or copy a build' },
    { command: '/fs copy', description: 'copy my character for the site' },
    { command: '/fs chip', description: 'show or hide the chip' },
    { command: '/fs inspect', description: 'gear and talents of your target' },
    { command: '/fs link', description: "copy a player's site page" },
    { command: '/fs group', description: "your group's specs, gear and buffs" },
    { command: '/fs ready', description: 'check flasks, food and buffs now' },
    { command: '/fs loot', description: 'what an item is worth to you' },
    { command: '/fs floor N', description: 'tag members scored under N' },
    { command: '/fs options', description: 'open the settings page' },
    { command: '/fs diag', description: 'what is not working' },
    { command: '/fs help', description: 'every command' },
  ],
  /** The testers' commands, kept on the page because /fs help lists them: a small label, then
   *  one muted line. */
  commandsTestersLabel: 'For testers',
  commandsTestersLine: '/fs diag verbose, /fs inbox, /fs hello, /fs seen, /fs record',
  /** The companion handshake (addon-v2 spec §4.10 and §6.3–6.4), three sentences. */
  handshakeTitle: 'The companion handshake',
  handshake: [
    'When you log out, the addon saves your character export (level, spec, talents, equipped gear, and bags when the bag switch is on), any sim request and the guild tab’s live roster.',
    'The roster is saved only while “Share my build with my guild” is on, and members who opted out appear as “seen HH:MM” only; the companion uploads all of it and brings back what the site has for you, builds and sim results included.',
    'The build dialog’s Copy my character tab then says “Seen by the site” with the time, once the site has your latest copy.',
  ],

  // --- gear panel weights ---
  scoreColumn: 'Score',
  sortByScore: 'Sort by score',
  weightsTitle: (specName: string): string => `${specName} stat weights`,
  weightsAreOpinions: 'Weights are opinions. These are ours, with the sources we took them from.',
  weightsMissing: 'No stat weights for this spec yet, so nothing is scored.',
} as const;
