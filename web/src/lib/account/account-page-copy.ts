import { ctaCopy } from '../addon/cta-copy';
import { sharedText } from '../copy/shared';
// web/src/lib/account/account-page-copy.ts
// Every visible string /account's head and sections use (premium spec 2026-11-08 sections 7, 10)
// that character-list-copy.ts and signin-copy.ts do not already carry.

export const accountPageCopy = {
  title: 'Your account',
  signedInWithBattlenet: 'Signed in with Battle.net',
  signedInByEmail: 'Signed in by email',
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
    title: sharedText('shared.account.delete.title'),
    line: sharedText('shared.account.delete.line'),
    button: sharedText('shared.account.delete.button'),
  },
  /** The confirm dialog the Section opens (DELETE /v1/me, operation deleteAccount). */
  deleteDialog: {
    title: sharedText('shared.account.delete.question'),
    intro: sharedText('shared.account.delete.intro'),
    deletedLabel: sharedText('shared.account.delete.deleted.title'),
    /** Each line, and the words in it that link (the refund policy). */
    deleted: [
      {
        text: sharedText('shared.account.delete.deleted.subscription'),
        link: { label: sharedText('shared.account.delete.deleted.refundPolicy'), href: '/refunds' },
      },
      { text: sharedText('shared.account.delete.deleted.identity') },
      { text: sharedText('shared.account.delete.deleted.discord') },
      { text: sharedText('shared.account.delete.deleted.private') },
      { text: sharedText('shared.account.delete.deleted.guild') },
    ],
    keptLabel: sharedText('shared.account.delete.kept.title'),
    kept: [
      { text: sharedText('shared.account.delete.kept.billing') },
      { text: sharedText('shared.account.delete.kept.reports') },
      { text: sharedText('shared.account.delete.kept.characters') },
      { text: sharedText('shared.account.delete.kept.builds') },
    ],
    /** The field's label: the account's BattleTag, or the word for an account with none. */
    confirmLabel: (word: string): string => sharedText('shared.account.delete.confirmLabel', { word }),
    /** What an account without a BattleTag types (the server takes it from every account). */
    confirmWord: sharedText('shared.account.delete.confirmWord'),
    /** Under the field after the server refused what was typed (400 confirm_mismatch, invalid). */
    mismatch: (word: string): string => sharedText('shared.account.delete.mismatch', { battletag: word }),
    confirm: sharedText('shared.account.delete.button'),
    busy: sharedText('shared.account.delete.busy'),
    cancel: 'Cancel',
  },
  /** The home page's one-time notice after a deletion (`/?deleted=1`, read once). */
  deletedNotice: sharedText('shared.account.delete.done'),
  signOutAsk: 'Sign out of Forever Sixty on this browser?',
} as const;
