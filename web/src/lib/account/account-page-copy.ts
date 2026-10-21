// web/src/lib/account/account-page-copy.ts
// Every visible string /account's page header and rail panels use (brief 2026-09-22 §B1,
// §B2) that character-list-copy.ts and signin-copy.ts do not already carry.

export const accountPageCopy = {
  title: 'Your account',
  signedInWithBattlenet: 'Signed in with Battle.net',
  signedInByEmail: 'Signed in by email',
  signOut: 'Sign out',
  devicesLabel: 'Devices',
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
  /** The same banner for an account the site has no characters for yet. */
  signedInBannerEmpty: (who: string): string =>
    `Signed in as ${who}. Install the addon and the companion to bring your characters here.`,
  /** Shown in the hero band instead of the handoff links when no character on the account
   *  has a build at all (spec 2026-09-22 §3.1). */
  noBattlenetDataForRealm: 'Blizzard serves no data for this realm type yet.',
  /** "Your ratings" panel heading (spec §3.1). */
  yourRatingsLabel: 'Your ratings',
} as const;
