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
  localDown: 'The companion stopped responding. Right-click its icon in the system tray (bottom right, under the ^ arrow), choose Quit, then open it again.',
  trayNotice: 'Still running in the tray. Use Quit from the tray icon to stop logging.',
  dismiss: 'Dismiss',

  footer: {
    version: (v) => 'Companion ' + v,
    site: 'foreversixty.gg',
    updateReady: 'Update ready. It installs the next time you open the companion.',
    // The first half minute after the first launch of a new version.
    updatedTo: (v) => 'Updated to ' + v,
  },

  status: {
    reportLogging: 'Open this report on the site',
    // The headline while the companion's own service does not answer; the
    // cards under it keep the last numbers it sent, dimmed.
    unavailableTitle: 'Status unavailable',
    unavailableSub: 'The numbers below are the last ones the companion sent.',
    inProgress: 'In progress',
    rows: {
      report: 'Current report',
      fights: 'Fights uploaded',
      queued: 'Queued',
      addon: 'Addon',
      lastSync: 'Last sync',
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
      // The second line of the Current report row while it is not credited.
      notYetHelp: 'Log out once in game with the addon loaded. The companion credits this report then.',
      separator: ' · ',
    },
    // The Addon row's Check for builds now button and the line it leaves.
    checkBuilds: 'Check for builds now',
    checking: 'Checking…',
    checked: (when, found) => 'Checked ' + when + ' · ' + found,
    checkedJustNow: 'just now',
    checkedBuilds: (n, who) => n + (n === 1 ? ' build' : ' builds') + (who.length ? ' for ' + who.join(', ') : ''),
    checkedNothing: 'nothing new',
    noReport: 'None yet',
    starting: 'Starting...',
    folders: 'Game folders',
    noFolders: 'None found yet.',
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
    emptyBody: 'Type /combatlog in game, fight something, and the report shows up here within seconds.',
    unpairedTitle: 'Sign in first.',
    unpairedButton: 'Sign in with Battle.net',
    unnamed: 'Unnamed report',
    inProgress: 'in progress',
    notUploaded: 'not uploaded yet',
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
    fallbackBusy: 'Connecting...',
    tryAgain: 'Try again',
    starting: 'Opening your browser...',
    // Signing in
    waiting: 'Waiting for you to sign in...',
    checkCode: 'Check that your browser shows this code:',
    lasts: 'This link works for 10 minutes.',
    countdown: (m, s) => 'Expires in ' + m + ':' + String(s).padStart(2, '0'),
    reopen: 'Open the browser again',
    cancel: 'Cancel',
    trouble: 'Having trouble reaching the site. Still trying...',
    // Browser blocked
    blocked: 'Could not open your browser. Copy this link into any browser, on this PC or another.',
    copyLink: 'Copy link',
    copied: 'Copied.',
    linkLabel: 'Sign-in link',
    // Connected
    connectedAs: (who) => (who ? 'Connected as ' + who : 'Connected'),
    flash: 'Connected. The companion is ready.',
    savedIn: 'Saved in',
    details: 'Details',
    deviceName: 'Device name',
    api: 'API',
    openReports: 'Open my reports',
    forget: 'Forget this device',
    forgetAsk: 'Forget this device? Uploads stop until you sign in again. Your reports on the site stay.',
    forgetYes: 'Forget',
    forgetNo: 'Cancel',
    forgottenRemoved: 'Forgotten. This computer was removed from your account.',
    forgottenLocal: 'Forgotten on this PC. To remove it from your account too, open ',
    forgottenLink: 'Account > Devices',
    forgottenLocalEnd: ' on foreversixty.gg.',
    forgetting: 'Forgetting...',
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
    removing: 'Removing...',
    addFolder: 'Add a folder...',
    adding: 'Choosing...',
    folderAdded: 'Added.',
    folderAddedPending: 'Added. Switching to the new folder when the current report closes.',
    folderRemoved: 'Removed.',
    noFolders: 'No game folder found yet.',
    character: 'Reports are credited to',
    characterAuto: 'Automatic (read from the log)',
    characterAutoNow: (name) => 'Automatic (currently ' + name + ')',
    characterHelp: 'Automatic reads who is logging from the combat log itself. Pick a character only if that is wrong.',
    save: 'Save changes',
    saving: 'Saving...',
    saved: 'Saved.',
    savedVisibility: (words) => 'Saved. New reports will be ' + words + '. The report in progress keeps its setting.',
    rulesets: { normal: 'Normal', pvp: 'PvP', rp: 'Roleplaying', hardcore: 'Hardcore' },
  },
};

if (typeof module !== 'undefined') module.exports = COPY;
