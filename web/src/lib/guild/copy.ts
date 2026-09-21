// web/src/lib/guild/copy.ts
// Every visible string the guild home, claim, settings, join and account-consent views
// use, in the site's own honest-copy voice: no exclamation marks, never a nag inside a
// tool. Wording quoted directly from the spec is marked as such below.
export const guildHomeCopy = {
  loading: 'Loading your guild.',
  reportsHeading: "This week's reports",
  // Spec section 4.1, exact wording.
  noReports: 'No reports this week yet.',
  // One report row's kill/wipe tally -- combined here rather than as two separate labels
  // so a translation can reorder or repunctuate the pair as a unit.
  reportSummary: (killCount: number, wipeCount: number): string => `${killCount} kills · ${wipeCount} wipes`,
  rosterHeading: 'Roster',
  // Spec section 4.1, exact wording, officer viewer.
  emptyRosterOfficer:
    "You're the only member the site knows about. Share the invite link to bring the rest of the guild in.",
  emptyRosterMember: "You're the only member the site knows about.",
  loggedRecently: 'Logged in the last day',
  unverified: 'Unverified',
  itemLevelLabel: 'ilvl',
  approve: 'Approve',
  remove: 'Remove',
  settingsLink: 'Guild settings',
  claimLink: 'Claim this guild',
  openSim: 'Open in simulator',
  openPlanner: 'Open in planner',
} as const;

export const guildClaimCopy = {
  loading: "Checking this guild's claim.",
  failed: 'That did not load. Reload the page to try again.',
  // Guild name isn't known until the public guild page resolves, so the heading falls
  // back to a nameless form rather than the page ever hardcoding either half itself.
  heading: (guildName: string): string => (guildName === '' ? 'Claim this guild' : `Claim ${guildName}`),
  unclaimed: 'Nobody has claimed this guild yet.',
  claimedByYou: 'You claimed this guild.',
  claimedBySomeoneElse: (battletag: string): string => `Claimed by ${battletag}.`,
  // `expiresAt` is only known in the same session the viewer just triggered the pending
  // claim (the POST .../claim response) -- a pending claim loaded later from settings
  // carries no expiry, hence the optional form rather than a second, separate string.
  pending: (expiresAt?: string): string =>
    expiresAt === undefined
      ? 'A claim is pending, confirmed by a second officer or the guild master.'
      : `A claim is pending, confirmed by a second officer or the guild master. Expires ${expiresAt.slice(0, 10)}.`,
  claimButton: 'Claim this guild',
  confirmButton: 'Confirm this claim',
  releaseButton: 'Release claim',
  notEligible: 'Only an officer or the guild master of this guild can claim it.',
  signInLine: 'Sign in to claim this guild.',
} as const;

export const guildSettingsCopy = {
  loading: 'Loading settings.',
  failed: 'Settings did not load. Reload the page to try again.',
  // Spec section 4.4, exact wording.
  forbidden: 'You need to be a verified officer of this guild to see its settings.',
  heading: 'Guild settings',
  defaultVisibility: 'Default report visibility',
  officerThreshold: 'Officer rank threshold',
  inviteHeading: 'Invite link',
  // Spec section 3.3, exact wording.
  inviteWarning: 'Anyone with this link can join as a member. Rotate it if it leaks.',
  rotateButton: 'Rotate invite link',
  tokenShownOnce: 'This link is shown once. Copy it now.',
  saved: 'Saved.',
  // Shown when a save or rotate action itself fails (as opposed to the initial load).
  actionFailed: 'That did not save; try again.',
} as const;

export const guildJoinCopy = {
  loading: 'Loading this invite.',
  failed: 'That invite link did not work.',
  joinButton: 'Join as a member',
  joined: 'You joined the guild.',
  signInLine: 'Sign in to join this guild.',
} as const;

export const guildConsentCopy = {
  heading: 'My guilds',
  roster: 'Roster only',
  gear: 'Gear',
  gearBags: 'Gear and bags',
  leave: 'Leave',
  left: 'Left.',
} as const;
