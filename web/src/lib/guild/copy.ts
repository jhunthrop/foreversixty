// web/src/lib/guild/copy.ts
// Every visible string the guild home, claim, settings, join and account-consent views
// use, in the site's own honest-copy voice: no exclamation marks, never a nag inside a
// tool. Wording quoted directly from the spec is marked as such below.
export const guildHomeCopy = {
  reportsHeading: "This week's reports",
  // Spec section 4.1, exact wording.
  noReports: 'No reports this week yet.',
  // One report row's kill/wipe tally -- combined here rather than as two separate labels
  // so a translation can reorder or repunctuate the pair as a unit.
  reportSummary: (killCount: number, wipeCount: number): string => `${killCount} kills · ${wipeCount} wipes`,
  progressionHeading: 'Progression',
  // The progression list's empty line, moved out of Guild.svelte's markup when it became
  // an EmptyState: section 1.2 of the 2026-09-22 design wants every visible string here.
  noProgression: 'No pulls recorded yet.',
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
  untitledReport: 'Untitled report',
  contestButton: 'Contest this claim',
  // The true contest rules. One plain fact per line, in this order: account eligibility,
  // the per-account attempt limit, the per-account permanent block after an uphold, the
  // per-guild cooldown after an uphold, then the freeze consequence. Simplified (no more
  // young-claim/corroboration nuance) and extended (the guild-wide upheld cooldown) by a
  // later security-review response -- read directly against contest.go's ContestClaim,
  // recentlyUpheld, and previouslyUpheld.
  contestRules: [
    'Contesting needs a Battle.net-linked account.',
    'You may attempt one contest every 30 days.',
    'If a moderator upholds this guild’s claim, you cannot contest it again.',
    'A guild whose claim was upheld in the last 30 days cannot be contested again, unless a moderator reopens it.',
    'Officer tools freeze until a moderator reviews the contest. Members can still read and upload.',
  ] as readonly string[],
  contestConfirmButton: 'Yes, contest this claim',
  cancel: 'Cancel',
  frozenNotice: 'This guild’s claim is contested. Officer actions are frozen until a moderator resolves it.',
  // The contest call itself succeeded but the page's own follow-up refresh failed -- a
  // distinct case from the contest failing outright, so it gets its own honest sentence
  // rather than the generic action-failed message that would wrongly imply the contest
  // itself did not go through.
  contestRecordedRefreshFailed: 'This claim is now contested. Reload the page to see the latest state.',
  reportsUnverifiedNote: 'You are not verified yet, so only public reports show here.',
} as const;

export const guildClaimCopy = {
  failed: 'That did not load.',
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
  rulesHeading: 'How claiming works',
  rules: [
    'Claiming needs a Battle.net-linked account.',
    'You may attempt one claim every 30 days.',
    'You may hold only one claimed guild at a time.',
    'A claim can be contested.',
  ] as readonly string[],
  contested: 'This guild’s claim is contested and under review by a moderator.',
} as const;

export const guildSettingsCopy = {
  failed: 'Settings did not load.',
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
  frozenNotice: 'This guild’s claim is contested. Officer actions are frozen until a moderator resolves it.',
} as const;

export const guildJoinCopy = {
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
