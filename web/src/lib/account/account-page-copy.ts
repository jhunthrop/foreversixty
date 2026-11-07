import { ctaCopy } from '../addon/cta-copy';
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
    title: 'Delete account',
    line: 'Delete your account and the data that is only yours. The next step lists exactly what goes and what stays.',
    button: 'Delete account',
  },
  /** The confirm dialog the Section opens (DELETE /v1/me, operation deleteAccount). */
  deleteDialog: {
    title: 'Delete your account?',
    intro: "This can't be undone.",
    deletedLabel: 'Deleted',
    /** Each line, and the words in it that link (the refund policy). */
    deleted: [
      {
        text: 'Any Premium subscription is cancelled at once. There is no automatic refund; the refund policy applies.',
        link: { label: 'refund policy', href: '/refunds' },
      },
      {
        text: 'Your BattleTag, email and Battle.net sign-in. Every browser, companion and app is signed out.',
      },
      { text: 'Your Discord link and Premium role, notification settings and watch list.' },
      { text: 'Your private reports, saved sims, addon exports, and characters known only from your addon.' },
      { text: 'Your guild memberships, applications and raid sign-ups.' },
    ],
    keptLabel: 'Kept',
    kept: [
      {
        text: 'Stripe billing records, which we must keep for tax and legal reasons. Your link to them is removed.',
      },
      { text: 'Public and guild reports, with your name removed.' },
      {
        text: 'Character data from public sources, such as the Battle.net profile and public logs, no longer linked to you.',
      },
      { text: 'Builds you shared, as anonymous links.' },
    ],
    /** The field's label: the account's BattleTag, or the word for an account with none. */
    confirmLabel: (word: string): string => `Type ${word} to confirm`,
    /** What an account without a BattleTag types (the server takes it from every account). */
    confirmWord: 'delete',
    /** Under the field after the server refused what was typed (400 confirm_mismatch, invalid). */
    mismatch: (word: string): string => `Type ${word} to confirm.`,
    confirm: 'Delete account',
    busy: 'Deleting',
    cancel: 'Cancel',
  },
  /** The home page's one-time notice after a deletion (`/?deleted=1`, read once). */
  deletedNotice: 'Your account is deleted. Thanks for playing with us.',
  signOutAsk: 'Sign out of Forever Sixty on this browser?',
} as const;
