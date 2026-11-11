// companion/ui/copy.js
// Every sentence the window itself owns, in one object, so app.js and install.js read
// COPY.status.rows.report and never a literal. The words are apps/shared/copy/companion.json,
// generated into copy_gen.js (COPY_TEXT) by web/scripts/gen-mobile-copy.mjs: edit the JSON and
// run the generator, never this file's words. What the Go side owns (the Status headline, the
// addon row, the failure messages, the tray) arrives in the snapshot from the same JSON through
// internal/copygen. Nothing here is ever built from an error's text.
'use strict';

const COPY = (() => {
  const T = typeof COPY_TEXT !== 'undefined' ? COPY_TEXT : require('./copy_gen.js');

  // fill puts each value in its {name} marker: fill('Log written {ago}', { ago: '5 min ago' }).
  const fill = (text, values) => text.replace(/\{([a-zA-Z][a-zA-Z0-9]*)\}/g, (m, name) => (name in values ? String(values[name]) : m));
  // count is a plural's form for n, the count filled in: "1 fight", "3 fights".
  const count = (plural, n) => fill(n === 1 ? plural.one : plural.other, { [plural.count]: n });

  return {
    title: T.title,

    nav: { status: T.navStatus, reports: T.navReports, device: T.navDevice, settings: T.navSettings },
    pages: { status: T.pagesStatus, reports: T.pagesReports, device: T.pagesDevice, settings: T.pagesSettings },

    // Shown when the companion's own service does not answer. The cause is never printed.
    localDown: T.localDown,
    trayNotice: T.trayNotice,
    dismiss: T.dismiss,

    // Shown in the body until the companion has answered once.
    checking: T.checking,

    footer: {
      // The footer is "Companion" at first paint and "Companion 0.2.13" from the first answer.
      product: T.footerProduct,
      version: (v) => fill(T.footerVersion, { version: v }),
      site: T.footerSite,
      updateReady: T.footerUpdateReady,
      // After "Companion 0.2.13", for the first half minute after a new version's first launch.
      justUpdated: T.footerJustUpdated,
      joint: T.joint,
    },

    status: {
      reportLogging: T.statusReportLogging,
      // The figure at the right of the headline: the fights the site holds.
      figFights: (n) => (n === 1 ? T.statusFigFightsOne : T.statusFigFightsOther),
      // Under the Queued row once the fight at the head has failed: it holds the rest, in order.
      held: (h) => fill(T.statusHeld, {
        what: h.fight ? fill(T.statusHeldFight, { fight: h.fight }) : T.statusHeldRawLog,
        times: count(T.statusHeldTimes, h.attempts),
      }),
      // The Current report row's second line: when the combat log last grew.
      logWritten: (ago) => fill(T.statusLogWritten, { ago }),
      upToDate: T.versionUpToDate,
      // The headline while the companion's own service does not answer; the cards under it keep
      // the last numbers it sent, dimmed.
      unavailableTitle: T.statusUnavailableTitle,
      unavailableSub: T.statusUnavailableSub,
      // The caption over the greyed numbers while Status is unavailable.
      lastKnown: T.statusLastKnown,
      // The Fights uploaded row while the open report has none yet.
      noFightsYet: T.statusNoFightsYet,
      noFightsHelp: T.statusNoFightsHelp,
      // The one row while every fight of the report is still waiting to go up.
      uploadingFights: (n) => fill(T.statusUploadingFights, { fights: count(T.countFights, n) }),
      inProgress: T.statusInProgress,
      rows: {
        report: T.statusRowsReport,
        fights: T.statusRowsFights,
        queued: T.statusRowsQueued,
        addon: T.statusRowsAddon,
        lastSync: T.statusRowsLastSync,
      },
      // The Last sync row's sent line; the other lines come from the snapshot.
      lastSyncSent: (n, ago) => fill(T.statusLastSyncSent, { characters: count(T.countCharacters, n), ago }),
      // Who a report is credited to. Two states: the name-versus-realm split of why it is not
      // stays in companion.log, because the player's step is the same.
      credit: {
        credited: (name) => fill(T.statusCreditCredited, { name }),
        notYet: T.statusCreditNotYet,
        // The log has named nobody yet on a report still open: not a fault, so muted.
        noName: T.statusCreditNoName,
        // The second line of the Current report row while it is not credited.
        notYetHelp: T.statusCreditNotYetHelp,
        separator: T.joint,
      },
      // The Addon row's Check for builds now button and the line it leaves.
      checkBuilds: T.statusCheckBuilds,
      checking: T.statusChecking,
      checked: (when, found) => fill(T.statusChecked, { when, found }),
      checkedBuilds: (n, who) => {
        const builds = count(T.countBuilds, n);
        return who.length ? fill(T.statusCheckedBuildsFor, { builds, who: who.join(', ') }) : builds;
      },
      checkedNothing: T.statusCheckedNothing,
      // The Install the addon into <folder> button while it works. The button's label and the
      // version lines come from the snapshot (copy.go).
      installingAddon: T.statusInstallingAddon,
      noReport: T.statusNoReport,
      starting: T.statusStarting,
      folders: T.statusFolders,
      noFolders: T.statusNoFolders,
      noInstall: T.statusNoInstall,
      noInstallLink: T.statusNoInstallLink,
      watching: T.statusWatching,
      lastPlayed: T.statusLastPlayed,
      notPlayed: T.statusNotPlayed,
      watchingTitle: T.statusWatchingTitle,
      // The watched install is a fault and shouts; any other is just a fact.
      advancedOffWatched: T.statusAdvancedOffWatched,
      advancedOff: T.statusAdvancedOff,
      advancedOn: T.statusAdvancedOn,
      advancedUnknown: T.statusAdvancedUnknown,
      // The line a folder row keeps comes from the snapshot (copy.go).
    },

    reports: {
      // The one name for the reports job, as the site's: the tray and the Device page say it too.
      all: T.reportsOnSite,
      emptyTitle: T.reportsEmptyTitle,
      emptyBody: T.reportsEmptyBody,
      unpairedTitle: T.reportsUnpairedTitle,
      unpairedButton: T.signIn,
      unnamed: T.reportsUnnamed,
      inProgress: T.reportsInProgress,
      // A report the site has not created yet. The companion creates it on its next upload try.
      notOnSite: T.reportsNotOnSite,
      notOnSiteWhen: T.reportsNotOnSiteWhen,
      fights: (n) => count(T.countFights, n),
      copyLink: T.copyLink,
      copied: T.copied,
      today: T.reportsToday,
      yesterday: T.reportsYesterday,
    },

    device: {
      // Not connected
      notConnected: T.deviceNotConnected,
      signedOutBar: (who) => (who ? fill(T.deviceSignedOutBarAs, { who }) : T.deviceSignedOutBar),
      intro: T.deviceIntro,
      signIn: T.signIn,
      signInNote: T.deviceSignInNote,
      fallbackToggle: T.deviceFallbackToggle,
      fallbackLabel: T.deviceFallbackLabel,
      fallbackPlaceholder: T.deviceFallbackPlaceholder,
      fallbackHelp: T.deviceFallbackHelp,
      fallbackButton: T.deviceFallbackButton,
      fallbackBusy: T.deviceFallbackBusy,
      tryAgain: T.deviceTryAgain,
      starting: T.deviceStarting,
      // Signing in
      waiting: T.deviceWaiting,
      checkCode: T.deviceCheckCode,
      // Under the heading while the sign-in page is open in the player's browser.
      opened: T.deviceOpened,
      countdown: (m, s) => fill(T.deviceCountdown, { m, s: String(s).padStart(2, '0') }),
      reopen: T.deviceReopen,
      cancel: T.deviceCancel,
      trouble: T.deviceTrouble,
      // Browser blocked
      blocked: T.deviceBlocked,
      copyLink: T.copyLink,
      copied: T.copied,
      linkLabel: T.deviceLinkLabel,
      // Connected
      connectedAs: (who) => (who ? fill(T.deviceConnectedAs, { who }) : T.deviceConnected),
      flash: T.deviceFlash,
      savedIn: T.deviceSavedIn,
      // The connected page's facts, from the snapshot.
      engine: T.deviceEngine,
      uploaded: T.deviceUploaded,
      uploadedValue: (reports, fights) => fill(T.deviceUploadedValue, { reports: count(T.countReports, reports), fights: count(T.countFights, fights) }),
      addonSeen: T.deviceAddonSeen,
      joint: T.joint,
      openReports: T.reportsOpen,
      forget: T.deviceForget,
      forgetAsk: T.deviceForgetAsk,
      forgetYes: T.deviceForgetYes,
      forgetNo: T.deviceForgetNo,
      forgottenRemoved: T.deviceForgottenRemoved,
      forgottenLocal: T.deviceForgottenLocal,
      forgottenLink: T.deviceForgottenLink,
      forgottenLocalEnd: T.deviceForgottenLocalEnd,
      forgetting: T.deviceForgetting,
      emptyCode: T.failureEmptyCode,
    },

    settings: {
      visibility: T.settingsVisibility,
      visibilityOptions: {
        public: T.settingsVisibilityOptionsPublic,
        unlisted: T.settingsVisibilityOptionsUnlisted,
        private: T.settingsVisibilityOptionsPrivate,
        guild: T.settingsVisibilityOptionsGuild,
      },
      // What each choice means, under the select.
      visibilityHelp: {
        public: T.settingsVisibilityHelpPublic,
        unlisted: T.settingsVisibilityHelpUnlisted,
        private: T.settingsVisibilityHelpPrivate,
        guild: T.settingsVisibilityHelpGuild,
      },
      // What the Saved note says about the choice.
      visibilitySaved: {
        public: T.settingsVisibilitySavedPublic,
        unlisted: T.settingsVisibilitySavedUnlisted,
        private: T.settingsVisibilitySavedPrivate,
        guild: T.settingsVisibilitySavedGuild,
      },
      folders: T.settingsFolders,
      foundAutomatically: T.settingsFoundAutomatically,
      remove: T.settingsRemove,
      removing: T.settingsRemoving,
      addFolder: T.settingsAddFolder,
      adding: T.settingsAdding,
      folderAdded: T.settingsFolderAdded,
      folderAddedPending: T.settingsFolderAddedPending,
      folderRemoved: T.settingsFolderRemoved,
      noFolders: T.settingsNoFolders,
      character: T.settingsCharacter,
      characterAuto: T.settingsCharacterAuto,
      characterAutoNow: (name) => fill(T.settingsCharacterAutoNow, { name }),
      characterHelp: T.settingsCharacterHelp,
      keepAddon: T.settingsKeepAddon,
      keepAddonHelp: T.settingsKeepAddonHelp,
      savedKeep: { on: T.settingsSavedKeepOn, off: T.settingsSavedKeepOff },
      save: T.settingsSave,
      noChanges: T.settingsNoChanges,
      unsaved: T.settingsUnsaved,
      saving: T.settingsSaving,
      saved: T.settingsSaved,
      savedVisibility: (words) => fill(T.settingsSavedVisibility, { words }),
      rulesets: {
        normal: T.settingsRulesetsNormal,
        pvp: T.settingsRulesetsPvp,
        rp: T.settingsRulesetsRp,
        hardcore: T.settingsRulesetsHardcore,
      },
    },

    // The install windows (install.html, internal/install).
    install: {
      // The Mac's first launch from the disk image or a download (internal/install, offer.go).
      move: {
        title: T.installMoveTitle,
        lede: T.installMoveLede,
        toggle: T.installMoveToggle,
        action: T.installMoveAction,
        busy: T.installMoveBusy,
        failed: T.installMoveFailed,
      },
      // The Windows setup download (setup.go); an upgrade when a copy is installed already.
      setup: {
        title: T.installSetupTitle,
        lede: T.installSetupLede,
        titleUpgrade: T.installSetupTitleUpgrade,
        ledeUpgrade: T.installSetupLedeUpgrade,
        toggle: T.installSetupToggle,
        action: T.installSetupAction,
        actionUpgrade: T.installSetupActionUpgrade,
        busy: T.installSetupBusy,
        busyUpgrade: T.installSetupBusyUpgrade,
        done: T.installSetupDone,
        doneUpgrade: T.installSetupDoneUpgrade,
        failed: T.installSetupFailed,
      },
      // Apps & Features' Uninstall (setup.go).
      uninstall: {
        title: T.installUninstallTitle,
        lede: T.installUninstallLede,
        toggle: T.installUninstallToggle,
        action: T.installUninstallAction,
        busy: T.installUninstallBusy,
        failed: T.installUninstallFailed,
        done: T.installUninstallDone,
      },
      // The step under the bar, as each one starts.
      steps: {
        stop: T.installStepsStop,
        copy: T.installStepsCopy,
        shortcut: T.installStepsShortcut,
        register: T.installStepsRegister,
        start: T.installStepsStart,
        remove: T.installStepsRemove,
      },
    },
  };
})();

if (typeof module !== 'undefined') module.exports = COPY;
