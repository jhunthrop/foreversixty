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
    'Applies everywhere your characters appear, on reports and rankings alike. Your reports are kept until you delete them or your account.',
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
  deleteTitle: 'Delete your account',
  /** Ruling MUST 4, matched to the privacy page's requests section. */
  deleteBody:
    'Email help@foreversixty.gg with your BattleTag and we delete your account and everything on it, within thirty days.',
  signOutAsk: 'Sign out of Forever Sixty on this browser?',
} as const;
