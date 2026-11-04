// web/src/lib/account/account-page-copy.ts
// Every visible string /account's page header and rail panels use (brief 2026-09-22 §B1,
// §B2) that character-list-copy.ts and signin-copy.ts do not already carry.

export const accountPageCopy = {
  title: 'Your account',
  /** The micro label over the page title. */
  eyebrow: 'Account',
  signedInWithBattlenet: 'Signed in with Battle.net',
  signedInByEmail: 'Signed in by email',
  signOut: 'Sign out',
  devicesLabel: 'Devices',
  reportsLabel: 'Your reports',
  /** The current-character card's prompt when it has no export (polish spec 4.3). */
  noExportPrompt:
    'No gear or talents yet. They appear once Battle.net publishes this character; paste an export from the addon to see them now.',
  youLabel: 'You',
  guildLabel: 'Guild',
  planLabel: 'Plan',
  currentCharacterLabel: 'Current character',
  noDevices:
    'No devices yet. The companion signs in with Battle.net from its Device page; a pairing code is only for the fallback.',
  /** A player with no companion yet: the next step is the download, and the code is the fallback (W.9). */
  downloadCompanion: 'Download the companion',
  downloadCompanionHref: '/setup#companion',
  pairADevice: 'Use a pairing code instead',
  pseudonymLabel: 'Show a pseudonym instead of my character names',
  pseudonymNote:
    'Applies everywhere your characters appear, on reports and rankings alike. Reports themselves are never deleted or rewritten.',
  planKey: 'Plan',
  /** spec 2026-09-22 §3.1: the hero band's Logs action, always shown beside the
   *  handoff links (Open in simulator/planner or the paste fallback). */
  heroLogs: 'Logs',
  /** The hub-arrival banner, shown once per `?signed_in=1` visit. */
  signedInBanner: (name: string): string =>
    `Signed in. ${name} is your current character; change it from any row below.`,
  /** The same banner for an account the site has no characters for yet. A Battle.net
   *  account imports its Forever characters on sign-in, so with none there are none to bring;
   *  an email account has only the addon. */
  signedInBannerEmpty: (who: string, viaBattlenet: boolean): string =>
    viaBattlenet
      ? `Signed in as ${who}. Battle.net lists no Forever character on this account yet; make one in game, then refresh from Battle.net.`
      : `Signed in as ${who}. Sign in with Battle.net to bring your Forever characters here, or install the addon and the companion.`,
  /** Shown in the hero band instead of the handoff links when no character on the account
   *  has a build at all (spec 2026-09-22 §3.1). */
  noBattlenetDataForRealm: 'Battle.net sends gear and talents after you next log out of the game.',
  /** "Your ratings" panel heading (spec §3.1). */
  yourRatingsLabel: 'Your ratings',
} as const;
