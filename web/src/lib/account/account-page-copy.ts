import { ctaCopy } from '../addon/cta-copy';
import { fillShared } from '../copy/shared';
import {
  sharedAccountDeleteBusy,
  sharedAccountDeleteButton,
  sharedAccountDeleteConfirmLabel,
  sharedAccountDeleteConfirmWord,
  sharedAccountDeleteDeletedDiscord,
  sharedAccountDeleteDeletedGuild,
  sharedAccountDeleteDeletedIdentity,
  sharedAccountDeleteDeletedPrivate,
  sharedAccountDeleteDeletedRefundPolicy,
  sharedAccountDeleteDeletedSubscription,
  sharedAccountDeleteDeletedTitle,
  sharedAccountDeleteDone,
  sharedAccountDeleteIntro,
  sharedAccountDeleteKeptBilling,
  sharedAccountDeleteKeptBuilds,
  sharedAccountDeleteKeptCharacters,
  sharedAccountDeleteKeptReports,
  sharedAccountDeleteKeptTitle,
  sharedAccountDeleteLine,
  sharedAccountDeleteMismatch,
  sharedAccountDeleteQuestion,
  sharedAccountDeleteTitle,
} from '../copy/shared/account.generated';
// web/src/lib/account/account-page-copy.ts
// Every visible string /account's head and sections use (premium spec 2026-11-08 sections 7, 10)
// that character-list-copy.ts and signin-copy.ts do not already carry.

export const accountPageCopy = {
  /** The Characters panel when /v1/me failed; LoadError adds Try again. */
  loadFailed: 'Could not reach Forever Sixty. Check your connection.',
  title: 'Your account',
  /** The You panel's facts: who is signed in, and how. */
  battletagLabel: 'BattleTag',
  emailLabel: 'Email',
  signInLabel: 'Sign-in',
  signInBattlenet: 'Battle.net',
  signInEmail: 'Email',
  signOut: 'Sign out',
  devicesLabel: 'Devices',
  youLabel: 'You',
  guildLabel: 'Guild',
  planLabel: 'Plan',
  noDevices:
    'No devices yet. The companion signs in with Battle.net from its Device page; a pairing code is only for the fallback.',
  /** A player with no companion yet: the next step is the download, and the code is the fallback (W.9). */
  downloadCompanion: 'Download the companion',
  downloadCompanionHref: '/setup#companion',
  pairADevice: 'Use a pairing code instead',
  pseudonymLabel: 'Show a pseudonym instead of my character names',
  pseudonymNote:
    'Applies everywhere your characters appear, on reports and rankings alike. Private reports go when you delete them or your account; public and guild reports stay, with your name removed.',
  /** The hub-arrival banner, shown once per `?signed_in=1` visit. */
  signedInBanner: (name: string): string =>
    `Signed in. ${name} is your current character; change it from any row below.`,
  /** The same banner for an account the site has no characters for yet. A Battle.net
   *  account imports its Forever characters on sign-in, so with none there are none to bring;
   *  an email account has only the addon. */
  signedInBannerEmpty: (who: string, viaBattlenet: boolean): string =>
    viaBattlenet
      ? `Signed in as ${who}. Battle.net lists no Forever character on this account yet; make one in game, then refresh from Battle.net.`
      : `Signed in as ${who}. Sign in with Battle.net to bring your Forever characters here, or ${ctaCopy.installAddonAndCompanionInline}.`,
  /** Under the Characters Plate: the one link to the reports page (ruling A3). */
  reportsLink: 'Your reports',
  /** The rail's last Section (gap G6, closed by the account-delete lane 2026-11-07). */
  deleteSection: {
    title: sharedAccountDeleteTitle,
    line: sharedAccountDeleteLine,
    button: sharedAccountDeleteButton,
  },
  /** The confirm dialog the Section opens (DELETE /v1/me, operation deleteAccount). */
  deleteDialog: {
    title: sharedAccountDeleteQuestion,
    intro: sharedAccountDeleteIntro,
    deletedLabel: sharedAccountDeleteDeletedTitle,
    /** Each line, and the words in it that link (the refund policy). */
    deleted: [
      {
        text: sharedAccountDeleteDeletedSubscription,
        link: { label: sharedAccountDeleteDeletedRefundPolicy, href: '/refunds' },
      },
      { text: sharedAccountDeleteDeletedIdentity },
      { text: sharedAccountDeleteDeletedDiscord },
      { text: sharedAccountDeleteDeletedPrivate },
      { text: sharedAccountDeleteDeletedGuild },
    ],
    keptLabel: sharedAccountDeleteKeptTitle,
    kept: [
      { text: sharedAccountDeleteKeptBilling },
      { text: sharedAccountDeleteKeptReports },
      { text: sharedAccountDeleteKeptCharacters },
      { text: sharedAccountDeleteKeptBuilds },
    ],
    /** The field's label: the account's BattleTag, or the word for an account with none. */
    confirmLabel: (word: string): string => fillShared(sharedAccountDeleteConfirmLabel, { word }),
    /** What an account without a BattleTag types (the server takes it from every account). */
    confirmWord: sharedAccountDeleteConfirmWord,
    /** Under the field after the server refused what was typed (400 confirm_mismatch, invalid). */
    mismatch: (word: string): string => fillShared(sharedAccountDeleteMismatch, { battletag: word }),
    confirm: sharedAccountDeleteButton,
    busy: sharedAccountDeleteBusy,
    cancel: 'Cancel',
  },
  /** The home page's one-time notice after a deletion (`/?deleted=1`, read once). */
  deletedNotice: sharedAccountDeleteDone,
  signOutAsk: 'Sign out of Forever Sixty on this browser?',
} as const;
