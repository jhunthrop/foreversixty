import { ctaCopy } from './cta-copy';
import { copiedThen } from '../ui/copy';
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
  /** A portrait-style capture (settings, character panel, guild tab, build dialogs): shown at
   *  native width under a 900 px cap instead of the 420 px landscape cap. */
  tall?: boolean;
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
  copiedForGame: copiedThen('Paste it in game with /fs'),
  /** The clipboard refused (permission, an insecure page, no clipboard at all): the code is
   *  shown in a read-only field under this line instead. */
  copyFailed: 'Could not copy; select the code below instead',
  /** The accessible name of that read-only field. */
  buildCodeLabel: 'Build code',
  addonCodeHint:
    'Paste this in game with /fs; the Forever Sixty line and the talent window follow it from then on.',

  // --- planner import box ---
  importTitle: 'Import from addon',
  importPlaceholder: 'FS1:…',
  importAction: 'Import',
  importNoExport: "Don't have an export?",
  importGetAddon: ctaCopy.getAddonAndCompanion,
  importOrderApproximated:
    'Nothing in the game records the order a build was spent in, so this order is our best reconstruction: the main tree first, each talent as soon as its row opens. The finished tree is exact.',
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
  installCurseForge: 'Install from CurseForge',
  installWago: 'Install from Wago Addons',
  installZip: 'Download the zip',
  /** The one quiet line under the install buttons (UX review r5, finding 2). */
  installDataHint:
    'Forever Sixty Data adds the item scores the addon uses, updated nightly. CurseForge and Wago install it for you.',
  installDataRelease: 'Latest data release',
  currentDataBuild: (build: string): string => `Current data build: ${build}`,
  flowOutTitle: 'Out of the game',
  /** The addon's copy mode (L8.5 ruling 5): `/fs copy` opens it with the box selected, and
   *  Ctrl+C is the action; there is no button, and the old export command is gone. */
  flowOutBody:
    'In game, type /fs copy, press Ctrl+C, then paste it into the planner’s Import from addon box.',
  flowInTitle: 'Into the game',
  flowInBody: 'Save a build here, press Copy for the game, then type /fs in game and paste it.',

  // --- /setup page's three steps (spec 2026-09-25 §3.4; the companion second, owner 2026-11-10) ---
  setupSignInTitle: 'Sign in',
  setupCompanionTitle: 'The companion',
  setupAddonTitle: 'The addon',
  setupSignInBody: 'Brings your Forever characters in from Battle.net, and saves your builds, sims and logs.',
  /** Same wording as `home-panel-copy.ts`'s `signInButton`; kept as its own key here so
   *  every visible string on this page comes from `addonCopy`, matching this page's own
   *  copy-module convention rather than reaching into the home page's module. */
  setupSignInAction: 'Sign in with Battle.net',
  /** Under the companion step's title: which of the three steps can be skipped. */
  setupCompanionNote: 'Sends bags, bank and live logs after every session.',
  /** Advanced Combat Logging is a setting that applies at once (no restart); /combatlog turns off
   *  when you log out, so it is typed each time you log in (setup rulings 2026-11-07). The path is
   *  the Forever client's own: Esc > Options, the System group, Network (Blizzard UI 1.60.1.70888,
   *  Blizzard_SettingsDefinitions_Shared/Network.lua registers Network under SETTING_GROUP_SYSTEM). */
  setupCompanionBody:
    'The companion logs live from your desktop. Turn on Options > System > Network > Advanced Combat Logging once. Type /combatlog each time you log in, and every fight appears under Your reports within seconds of the pull ending.',
  /** The page's title, and its one H1 (setup spec 2026-11-07 section 4). */
  setupPageTitle: 'Get set up',
  /** A signed-out visitor who pasted an export: the one done line a step can show without a session. */
  setupAddonDoneLine: 'An export is loaded; the site is pointed at that character.',
  /** The done and not-done lines of each step for a signed-in account (W.9). */
  setupSignedInAs: (name: string): string => `Signed in as ${name}`,
  setupAddonArrived: (character: string, when: string): string =>
    `${character} arrived from the addon ${when}`,
  /** The addon step before anything shows the addon: the parallel of `setupCompanionNote`. */
  setupAddonNote: 'Optional. Gear tips on every tooltip, talent marks, bags and bank.',
  /** The addon step once a companion found it in the game folder ('' when its TOC names no version). */
  setupAddonInstalled: (device: string, version: string): string =>
    version === '' ? `Installed on ${device}` : `Installed on ${device} · ${version}`,
  /** Under the note while no companion has looked in the game folder: an installed addon is not
   *  failed, only not seen yet (wow-player 2026-11-10). */
  setupAddonUnseen: 'Already installed? The companion sees it.',
  /** A companion's scan found no Forever Sixty addon in that PC's game folder: a to-do, never a fault. */
  setupAddonNotOn: (device: string): string => `Not in _forever_ on ${device} yet`,
  /** Under the installed line when the feed has a newer addon release; the row's Update buttons follow. */
  setupAddonBehind: (latest: string): string => `${latest} is out`,
  /** The addon step's install buttons, as one group (its accessible name). */
  setupInstallAddon: 'Install the addon',
  /** The same group on an installed addon behind the feed. */
  setupUpdateAddon: 'Update the addon',
  setupUpdateCurseForge: 'Update from CurseForge',
  setupUpdateWago: 'Update from Wago Addons',
  /** Step 1 while the Battle.net import runs or finished under two minutes ago (`/v1/me` `bnet_import`). */
  setupImportWait: 'Your characters can take up to two minutes to arrive.',
  /** Step 1 after `?signin=failed` (the API's callback redirect). */
  setupSignInFailed: 'Battle.net sign-in did not finish. Nothing was saved.',
  /** Step 1 when the last import could not key some characters (`bnet_import.state` partial, or
   *  failed with `unresolved`): the visit refresh brings them in once Battle.net names their realm. */
  setupImportMissing: (count: number): string =>
    count === 1
      ? 'One character is not in yet. It comes in on its own when you next visit.'
      : `${count} characters are not in yet. They come in on their own when you next visit.`,
  /** Step 1 when the last import keyed nothing (`bnet_import.state` failed, no count): the row's
   *  Refresh runs it again. */
  setupImportFailed: 'The Battle.net import did not finish. Refresh to run it again.',
  /** The hand install under The addon: the folders each zip holds (addon-release.yml ships
   *  ForeverSixty and ForeverSixtyInbox; addon-data-release.yml ForeverSixtyData). */
  setupHandInstall: 'Installing by hand',
  /** The path with the visitor's own separator: a Mac reads slashes, every other system the
   *  Windows backslashes (wow-player 2026-11-10). */
  setupHandInstallBody: (mac: boolean): string => {
    const path = ['World of Warcraft', '_forever_', 'Interface', 'AddOns'].join(mac ? '/' : '\\');
    return `Unzip both into ${path}: the addon zip holds ForeverSixty and ForeverSixtyInbox, the data zip holds ForeverSixtyData.`;
  },
  /** The paste field's label (the Section title is `setupTryExport`). */
  pasteFieldLabel: 'Addon export',
  setupCompanionLine: (device: string, seen: string): string =>
    `Companion signed in on ${device} · last seen ${seen}`,
  setupCompanionWaiting: 'No companion signed in yet',
  setupHowItWorks: 'How the addon works',
  setupCompanionPair: 'Sign in with Battle.net inside the companion; it connects itself.',
  /** The companion step on a phone or tablet: the companion is a desktop app, so no download is offered. */
  setupCompanionMobile: 'The companion runs on your PC or Mac.',
  /** The button that copies this page's address, to open it on the computer. */
  setupCopyLink: 'Copy link',
  setupLinkCopied: 'Link copied',
  /** Under a download button, from the feed's manifest: "Version 0.2.37", then " · 16.8 MB" for one file. */
  setupCompanionVersion: (version: string): string => `Version ${version}`,
  setupCompanionFileMeta: (version: string, size: string): string => `Version ${version} · ${size}`,
  /** The quiet disclosure under the visitor's own download that lists the other builds. */
  setupOtherPlatforms: 'Other platforms',
  /** The primary download button: "Download for Windows", "Download for Mac". */
  setupCompanionDownload: (platform: string): string => `Download for ${platform}`,
  /** Both builds side by side, when the page cannot tell the visitor's system: the group's name. */
  setupDownloadCompanion: 'Download the companion',
  /** A 2x crop of the 0.2.44 companion's Status page (round-5 window, seal v4), the Logging state
   *  credited to the fixtures' fictional two-word name; shot by design/mocks/companion/shoot_ui.cjs
   *  (state logging-credited, 800 px window). */
  setupCompanionCapture: {
    src: '/addon/companion-status.jpg',
    width: 1600,
    height: 460,
    /** The phone source (W.10, P.1): the headline's title and sub line only, a 2x crop of the same
     *  Status capture, shown under 640px so the sub line stays near 11px at 314px wide. */
    narrow: { src: '/addon/companion-status-phone.jpg', width: 912, height: 140 },
    /** The label on the picture while the visitor is signed out: the name in it is not theirs. */
    exampleLabel: 'Example',
    alt: 'The companion’s Status page: Logging, uploading your fights as they end, logging as Bow Jackzon.',
  },
  /** The disclosure that holds the paste box. */
  setupTryExport: 'Try an export',
  /** What an export is, first thing inside its disclosure. */
  setupTryExportBody:
    'Paste what the addon saved (type /fs copy in game) to see your character without the companion.',

  // --- /setup: what the addon does in the game (Task L8.3) ---
  /** One row per surface of the addon-v2 spec's §4, in its order; `title` is the spec's own
   *  heading and `body` one plain sentence from that section. `images` are crops of real
   *  client captures (docs/superpowers/reviews/week1/), each alt text saying only what is on
   *  screen in it; a surface with no capture yet (messaging) has none, and says nothing
   *  about it. `width`/`height` are the files' own pixel sizes, for the fixed aspect box. */
  surfacesTitle: 'What it does in the game',
  /** On a phone the surfaces fold behind this one row (setup rulings U5); open from 768. */
  surfacesShow: 'Show the screenshots',
  surfaces: [
    {
      section: '4.1',
      title: 'The Forever Sixty line',
      body: 'A small icon and line beside the minimap that names your next step: the talent to take, a bag item to equip, loot here for you, or what a raid check found missing.',
      images: [
        {
          src: '/addon/chip-tooltip-legend.jpg',
          width: 1019,
          height: 360,
          alt: 'The Forever Sixty line beside the minimap with its hover open: “Forever Sixty · Beast Mastery / Marksmanship build”, “Next talent at 25: Lethal Attacks”, the legend lines “Green arrow: a bag item beats what you wear.” and “Gold corner: next upgrade to chase.”, then the click, right-click and shift-drag hints.',
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
          alt: 'A hunter’s talent window, Beast Mastery, Marksmanship and Survival, with small order marks on the planned talents, one ringed in gold in Marksmanship, and the strip along the bottom naming Lethal Attacks (1 of 5) as the next talent, 12 of 51 points spent and 3 off plan, the last part in red.',
        },
      ],
    },
    {
      section: '4.3',
      title: 'Item tooltip',
      body: 'Every item tooltip gains one line saying whether the item is an upgrade and by how much, or on what you wear, the next upgrade and where it drops, with the details under Shift.',
      images: [
        {
          src: '/addon/item-tooltip-ranged.jpg',
          width: 910,
          height: 615,
          alt: 'A level 24 hunter’s worn Short Ash Bow tooltip on the character panel. Straight after the game’s own lines, one white line, “Next upgrade: Ranger Bow +3.5 DPS”, and under it “World drop”, with the Ranger Bow’s own tooltip beside it. Held Shift adds “+1 Agi”, “Also good: Lil Timmy’s Peashooter +2.2 (gun)” and “Cracked Blacksmith Hammer +0.8 (thrown)”.',
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
          alt: 'The top of the Combined Backpack with a small green arrow on the top-right corner of an item in the second row of the backpack: a bag item that beats what the character wears.',
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
          tall: true,
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
          tall: true,
          width: 728,
          height: 612,
          alt: 'The Forever Sixty guild tab: “Online only” unticked, “97 of 97 shown · 2 of 3 online run Forever Sixty · 1-20 of 97”, and columns Name, Level, Spec, Gear, Score (sorted), Attendance and Gap, with scored rows for Belee Dat (Frost, 2310), Onearm Lopen (Arms, 1720, gap 14 DPS) and Bow Jackzon (Beast Mastery, 1240, gap 0 DPS), then two dimmed rows and “no logs” for the rest.',
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
          tall: true,
          width: 770,
          height: 553,
          alt: 'The build dialog for a hunter on Paste a build: “Following: Beast Mastery / Marksmanship build · Hunter · 51 po…”, “Next talent: Take Lethal Attacks (Marksmanship, 1 of 5)”, and the empty box “Paste a build code from foreversixty.gg (Ctrl+V)”, with “Tab to switch” at the bottom.',
        },
        {
          /** Cropped to the window's head and tabs (setup rulings 2026-11-07): the full capture's
           *  selected export read as a broken pink image at grid size. */
          src: '/addon/build-dialog-copy-tab.jpg',
          width: 770,
          height: 100,
          alt: 'The build dialog’s head, “Forever Sixty”, with its two tabs, Paste a build and Copy my character, the second one selected.',
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
          tall: true,
          width: 1166,
          height: 790,
          alt: 'The Forever Sixty settings page with its eight switches ticked, each with its hint: Show the Forever Sixty line, Hide the Forever Sixty line in combat, Marks on the talent window, Show the minimap button, Gear tips on item tooltips, Score badges in chat, Share my build with my guild and Save my character for the companion app, then the Reset positions button.',
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
    { command: '/fs build', description: 'switch between saved builds' },
    { command: '/fs chip', description: 'show or hide the Forever Sixty line' },
    { command: '/fs tour', description: 'see the tour again' },
    { command: '/fs inspect', description: 'gear and talents of your target' },
    { command: '/fs link', description: "copy a player's site page" },
    { command: '/fs group', description: "your group's specs, gear and buffs" },
    { command: '/fs ready', description: 'check flasks, food and buffs now' },
    { command: '/fs job', description: 'your jobs for the next raid' },
    { command: '/fs loot', description: 'what an item is worth to you' },
    { command: '/fs floor N', description: 'tag members scored under N' },
    { command: '/fs options', description: 'open the settings page' },
    { command: '/fs diag', description: 'check what is wrong' },
    { command: '/fs help', description: 'every command' },
  ],
  /** The companion handshake (addon-v2 spec §4.10 and §6.3–6.4), three sentences. */
  handshakeTitle: 'The companion handshake',
  handshake: [
    'When you log out, the addon saves your character export (level, spec, talents, equipped gear, and bags when the bag switch is on), any sim request and the guild tab’s live roster.',
    'The roster is saved only while “Share my build with my guild” is on, and members who opted out appear as “seen HH:MM” only; the companion uploads all of it and brings back what the site has for you, builds and sim results included.',
    'The build dialog’s Copy my character tab then says “Site has your save from” and how long ago once the site has your latest copy, and “Not sent to the site yet · log out to send” until then.',
  ],

  // --- gear panel weights ---
  scoreColumn: 'Score',
  sortByScore: 'Sort by score',
  weightsTitle: (specName: string): string => `${specName} stat weights`,
  weightsAreOpinions: 'Weights are opinions. These are ours, with the sources we took them from.',
  weightsMissing: 'No stat weights for this spec yet, so nothing is scored.',
} as const;
