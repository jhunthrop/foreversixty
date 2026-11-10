// companion/ui/copy.js
// Every sentence the window itself owns, in one object so the strings
// are greppable. What the Go side owns (the Status headline, the addon
// row, the failure messages, the tray) arrives in the snapshot and is
// shown as it is: see internal/app/copy.go. Nothing here is ever built
// from an error's text.
'use strict';

const COPY = {
  title: 'Forever Sixty',

  nav: { status: 'Status', reports: 'Reports', device: 'Device', settings: 'Settings' },
  pages: { status: 'Status', reports: 'Your reports', device: 'This device', settings: 'Settings' },

  // Shown when the companion's own service does not answer. The cause
  // is never printed.
  localDown: 'The companion stopped answering, so fights are not uploading. Quit it from the tray icon and open it again.',
  trayNotice: 'Still running in the tray. Use Quit from the tray icon to stop logging.',
  dismiss: 'Dismiss',

  // Shown in the body until the companion has answered once.
  checking: 'Checking…',

  footer: {
    // The footer is "Companion" at first paint and "Companion 0.2.13" from the first answer.
    product: 'Companion',
    version: (v) => 'Companion ' + v,
    site: 'foreversixty.gg',
    updateReady: 'Update ready. Quit the companion from the tray icon and open it again to install it.',
    // After "Companion 0.2.13", for the first half minute after a new version's first launch.
    justUpdated: 'just updated',
    joint: ' · ',
  },

  status: {
    reportLogging: 'Open this report on the site',
    // The figure at the right of the headline: the fights the site holds.
    figFights: (n) => (n === 1 ? 'fight uploaded' : 'fights uploaded'),
    // Under the Queued row once the fight at the head has failed: it holds the rest, in order.
    held: (h) => (h.fight ? 'Fight ' + h.fight : 'The raw log') + ' failed ' + h.attempts + (h.attempts === 1 ? ' time' : ' times') +
      ' and retries on its own; the rest wait behind it.',
    // The Current report row's second line: when the combat log last grew.
    logWritten: (ago) => 'Log written ' + ago,
    upToDate: 'up to date',
    // The headline while the companion's own service does not answer; the
    // cards under it keep the last numbers it sent, dimmed.
    unavailableTitle: 'Status unavailable',
    unavailableSub: 'The numbers below are the last ones the companion sent.',
    // The caption over the greyed numbers while Status is unavailable.
    lastKnown: 'Last known',
    // The Fights uploaded row while the open report has none yet.
    noFightsYet: '0 so far',
    noFightsHelp: 'The first fight uploads when it ends.',
    // The one row while every fight of the report is still waiting to go up.
    uploadingFights: (n) => 'Uploading ' + n + (n === 1 ? ' fight' : ' fights') + '…',
    inProgress: 'In progress',
    rows: {
      report: 'Current report',
      fights: 'Fights uploaded',
      queued: 'Queued',
      addon: 'Addon',
      lastSync: 'Characters sent',
    },
    // The Last sync row's sent line; the other lines come from the snapshot.
    lastSyncSent: (n, ago) => n + (n === 1 ? ' character, ' : ' characters, ') + ago,
    // How long ago, in the Addon row's words (copy.go since).
    ago: {
      lessThanMinute: 'less than a minute ago',
      oneMinute: '1 minute ago',
      minutes: (n) => n + ' minutes ago',
      today: (clock) => 'today at ' + clock,
      yesterday: 'yesterday',
      days: (n) => n + ' days ago',
    },
    // Who a report is credited to. Two states: the name-versus-realm split of
    // why it is not stays in companion.log, because the player's step is the same.
    credit: {
      credited: (name) => 'credited to ' + name,
      notYet: 'not credited yet',
      // The log has named nobody yet on a report still open: not a fault, so muted.
      noName: 'name not read yet',
      // The second line of the Current report row while it is not credited.
      notYetHelp: 'Log out once in game with the addon loaded. The companion credits this report then.',
      separator: ' · ',
    },
    // The Addon row's Check for builds now button and the line it leaves.
    checkBuilds: 'Fetch my builds from the site',
    checking: 'Fetching…',
    checked: (when, found) => 'Checked ' + when + ' · ' + found,
    checkedJustNow: 'just now',
    checkedBuilds: (n, who) => n + (n === 1 ? ' build' : ' builds') + (who.length ? ' for ' + who.join(', ') : ''),
    checkedNothing: 'nothing new',
    // The Install the addon into <folder> button while it works. The button's label and the
    // version lines come from the snapshot (copy.go).
    installingAddon: 'Installing…',
    noReport: 'None yet',
    starting: 'Starting…',
    folders: 'Game folders',
    noFolders: 'None found yet.',
    noInstall: 'No WoW Forever install found. Pick its _forever_ folder in',
    noInstallLink: 'Settings',
    watching: 'Watching',
    lastPlayed: 'Last played',
    notPlayed: 'Not played yet',
    watchingTitle: 'The companion reads this folder’s combat log',
    // The watched install is a fault and shouts; any other is just a fact.
    advancedOffWatched: 'Advanced logging: OFF',
    advancedOff: 'Advanced logging: off',
    advancedOn: 'Advanced logging: on',
    advancedUnknown: 'Not checked yet. Log in once.',
    // The line a folder row keeps comes from the snapshot (copy.go).
  },

  reports: {
    all: 'Open all my reports on the site',
    emptyTitle: 'No reports yet.',
    emptyBody: 'Turn /combatlog on in game and your raids land here.',
    unpairedTitle: 'Sign in first.',
    unpairedButton: 'Sign in with Battle.net',
    unnamed: 'Unnamed report',
    inProgress: 'in progress',
    // A report the site has not created yet. The companion creates it on its next upload try.
    notOnSite: 'not on the site yet',
    notOnSiteWhen: 'uploads when the site answers',
    fights: (n) => n + (n === 1 ? ' fight' : ' fights'),
    copyLink: 'Copy link',
    copied: 'Copied.',
    today: 'Today',
    yesterday: 'Yesterday',
  },

  device: {
    // Not connected
    notConnected: 'Not connected',
    signedOutBar: (who) => (who ? 'Your sign-in as ' + who : 'Your sign-in') + ' is no longer valid. Sign in again to keep uploading.',
    intro: 'Sign in on foreversixty.gg so the companion can upload your fights and keep the addon up to date.',
    signIn: 'Sign in with Battle.net',
    signInNote: 'Opens your browser. Your Battle.net password never reaches this app.',
    fallbackToggle: 'Having trouble? Paste a code from the site',
    fallbackLabel: 'Code from the site',
    fallbackPlaceholder: 'Paste the code',
    fallbackHelp: 'On foreversixty.gg open Logs and press Pair a device.',
    fallbackButton: 'Connect with code',
    fallbackBusy: 'Connecting…',
    tryAgain: 'Try again',
    starting: 'Opening your browser…',
    // Signing in
    waiting: 'Waiting for you to sign in…',
    checkCode: 'Check that your browser shows this code:',
    // Under the heading while the sign-in page is open in the player's browser.
    opened: 'Your browser has opened. Sign in with Battle.net there, then come back.',
    countdown: (m, s) => 'Expires in ' + m + ':' + String(s).padStart(2, '0'),
    reopen: 'Open the browser again',
    cancel: 'Cancel',
    trouble: 'Having trouble reaching the site. Still trying…',
    // Browser blocked
    blocked: 'Could not open your browser. Copy this link into any browser, on this PC or another.',
    copyLink: 'Copy link',
    copied: 'Copied.',
    linkLabel: 'Sign-in link',
    // Connected
    connectedAs: (who) => (who ? 'Connected as ' + who : 'Connected'),
    flash: 'Connected. The companion is ready.',
    savedIn: 'Saved in',
    // The connected page's facts, from the snapshot.
    engine: 'Log parser',
    uploaded: 'Uploaded from this PC',
    uploadedValue: (reports, fights) => reports + (reports === 1 ? ' report, ' : ' reports, ') + fights + (fights === 1 ? ' fight' : ' fights'),
    addonSeen: 'Addon last seen in game',
    joint: ' · ',
    openReports: 'Open my reports',
    forget: 'Forget this device',
    forgetAsk: 'Forget this device? Uploads stop until you sign in again. Your reports on the site stay.',
    forgetYes: 'Forget',
    forgetNo: 'Cancel',
    forgottenRemoved: 'Forgotten. This computer was removed from your account.',
    forgottenLocal: 'Forgotten on this PC. To remove it from your account too, open ',
    forgottenLink: 'Account > Devices',
    forgottenLocalEnd: ' on foreversixty.gg.',
    forgetting: 'Forgetting…',
    emptyCode: 'Paste the code from the site first.',
  },

  settings: {
    visibility: 'New reports are',
    visibilityOptions: { public: 'Public', unlisted: 'Unlisted', private: 'Private', guild: 'Guild only' },
    // What each choice means, under the select.
    visibilityHelp: {
      public: 'Anyone can find it in rankings and search.',
      unlisted: 'Anyone with the link can open it. It is not listed.',
      private: 'Only you.',
      guild: 'Members of your guild.',
    },
    // What the Saved note says about the choice.
    visibilitySaved: {
      public: 'public',
      unlisted: 'unlisted',
      private: 'private',
      guild: 'guild only',
    },
    folders: 'Game folders',
    foundAutomatically: 'Found automatically',
    remove: 'Remove',
    removing: 'Removing…',
    addFolder: 'Add a folder…',
    adding: 'Choosing…',
    folderAdded: 'Added.',
    folderAddedPending: 'Added. Switching to the new folder when the current report closes.',
    folderRemoved: 'Removed.',
    noFolders: 'No WoW Forever install found yet. Add its _forever_ folder.',
    character: 'Reports are credited to',
    characterAuto: 'Automatic (read from the log)',
    characterAutoNow: (name) => 'Automatic (currently ' + name + ')',
    characterHelp: 'Pick a character only if Automatic is wrong.',
    keepAddon: 'Keep the addon up to date',
    keepAddonHelp: 'New versions go in when the game is closed. CurseForge and Wago users can leave this on; they get the same files.',
    savedKeep: { on: 'Saved. The addon will be kept up to date.', off: 'Saved. The companion will not update the addon.' },
    save: 'Save changes',
    noChanges: 'No changes',
    unsaved: 'Unsaved changes',
    saving: 'Saving…',
    saved: 'Saved.',
    savedVisibility: (words) => 'Saved. New reports will be ' + words + '. The report in progress keeps its setting.',
    rulesets: { normal: 'Normal', pvp: 'PvP', rp: 'Roleplaying', hardcore: 'Hardcore' },
  },

  // The Mac's first launch from a download (install.html, internal/install).
  install: {
    title: 'Move Forever Sixty to Applications?',
    lede: 'It runs from there and keeps itself up to date. The download goes in the Trash.',
    openAtLogin: 'Open at login',
    move: 'Move to Applications',
    moving: 'Moving…',
    failed: 'That did not work. Close this window and open Forever Sixty again to try once more.',
  },
};

if (typeof module !== 'undefined') module.exports = COPY;
