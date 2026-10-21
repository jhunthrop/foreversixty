// web/src/lib/guild/copy.ts
// Every visible string the guild home, claim, settings, join and account-consent views
// use, in the site's own honest-copy voice: no exclamation marks, never a nag inside a
// tool. Wording quoted directly from the spec is marked as such below.
export const guildHomeCopy = {
  reportsHeading: "This week's reports",
  // The public block's own reports heading (UX review defect 2, 2026-09-28): distinguishes
  // it from the member section's "This week's reports" above -- both can render on the same
  // page and, until now, both said the bare "Reports".
  allReportsHeading: 'All reports',
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
  /** The officer-only action on the empty-roster state: where the invite link lives. */
  manageInvite: 'Guild settings',
  loggedRecently: 'Logged in the last day',
  unverified: 'Unverified',
  // UX review defect 3 (2026-09-28): the roster summary an officer sees above the (now
  // unverified-first) roster list, so waiting rows are announced, not just sorted forward.
  waitingForApproval: (count: number): string => `${count} waiting for approval`,
  itemLevelLabel: 'ilvl',
  approve: 'Approve',
  remove: 'Remove',
  settingsLink: 'Guild settings',
  claimLink: 'Claim this guild',
  // Shown in the same header slot as claimLink, but only once an officer/leader claim is
  // already pending confirmation (UX review defect 1, 2026-09-28).
  confirmClaimLink: 'Confirm the claim',
  // A plain, unverified member has no claim standing at all -- this replaces the claim
  // link for them in that same slot, rather than leaving it silently absent.
  notVerifiedNote: 'You are not verified yet. An officer can approve you from the roster.',
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
  // What an officer who may claim reads: what claiming gets them.
  unclaimedForOfficer:
    'Nobody has claimed this guild. Claiming lets you approve members, set report visibility and manage billing.',
  // Claiming asks first, with the rules in one line.
  claimConfirmLine: (guildName: string): string =>
    `Claim ${guildName === '' ? 'this guild' : guildName}? You may attempt one claim every 30 days, hold one claimed guild at a time, and a claim can be contested.`,
  claimConfirmButton: 'Yes, claim it',
  // After the claim: the page says what the viewer holds and where to go next.
  youHold: (guildName: string): string => `You hold ${guildName === '' ? 'this guild' : guildName}.`,
  nextHeading: 'Next',
  nextRoster: 'Approve your officers and members on the roster',
  nextInvite: 'Make an invite link in settings',
  releaseConfirmLine: (guildName: string): string =>
    `Release your claim on ${guildName === '' ? 'this guild' : guildName}? Officers will have to claim it again.`,
  releaseConfirmButton: 'Release',
  cancel: 'Cancel',
  claimedBySomeoneElse: (battletag: string, since?: string): string =>
    since === undefined ? `Claimed by ${battletag}.` : `Claimed by ${battletag} since ${since}.`,
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
  // A plain member, the common visitor here (the guild page's own "Claim this guild" link
  // brought them): what is true, then what to do, since they cannot act themselves.
  memberCannotClaim:
    'You are a member of this guild, so you cannot claim it. Ask the guild master or an officer to save their character on Get set up, then open this page: the guild master claims at once, an officer starts a claim a second officer confirms.',
  // Signed in, but no character of theirs in this guild as far as the site knows.
  notAMember:
    'The site does not know a character of yours in this guild. Save one on Get set up and this page will know your rank.',
  // The state as a member may read it: the home says claimed, but never by whom.
  claimedByAnOfficer: 'This guild has been claimed by one of its officers.',
  backToGuild: (guildName: string): string => (guildName === '' ? 'Guild page' : guildName),
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
  // The invite link: rotate asks first, then the full URL with a Copy button. The API keeps
  // only a hash of the token, so the link can be shown once and never again; the page
  // says so instead of offering a "show current link" it cannot honour.
  rotateConfirmLine: 'Rotate the invite link? The old link stops working.',
  rotateConfirmButton: 'Rotate',
  rotateCancel: 'Cancel',
  tokenShownOnce: 'Shown once. Copy it now.',
  linkLabel: 'Invite link',
  copyButton: 'Copy link',
  copied: 'Copied',
  copyFailed: 'Could not copy. Select the link and copy it.',
  shareButton: 'Share',
  currentLinkMade: (date: string): string =>
    `Current link made ${date}. It cannot be shown again; rotate it for a new one.`,
  saved: 'Saved.',
  // Shown when a save or rotate action itself fails (as opposed to the initial load).
  actionFailed: 'That did not save; try again.',
  frozenNotice: 'This guild’s claim is contested. Officer actions are frozen until a moderator resolves it.',
} as const;

// Guild crest round (docs/contracts/2026-10-05-guild-crest-api.md), Settings tab block.
export const guildCrestCopy = {
  heading: 'Guild crest',
  defaultCaption: "Default: your faction's logo",
  pickerLabel: 'Choose a crest image',
  save: 'Save',
  saving: 'Saving…',
  saved: 'Saved.',
  remove: 'Remove',
  removing: 'Removing…',
  removeConfirmLine: 'Remove the guild crest? It falls back to your faction logo.',
  removeConfirmButton: 'Yes, remove the crest',
  actionFailed: 'That did not work; try again',
} as const;

export const guildJoinCopy = {
  failed: 'That invite link did not work.',
  invitedYou: (guildName: string, realm: string): string => `${guildName} · ${realm} invited you.`,
  consequence:
    'Joining puts your characters on its roster and shares your gear with its officers; you can change that on your account.',
  joinButton: (guildName: string): string => `Join ${guildName}`,
  joining: 'Joining…',
  joined: 'You joined the guild.',
  signInButton: 'Sign in with Battle.net to join',
  signInLine: 'Sign in to join this guild.',
  notNow: 'Not now',
  alreadyIn: (guildName: string): string => `You are already in ${guildName}.`,
  openGuild: 'Open the guild',
  invalid: 'This invite link no longer works. Ask a guild officer for a new one.',
} as const;

/** "all 3 of your characters" (consent and leave cover every character of the account in
 *  the guild): counted when the account page knows the number, plain otherwise. */
function characterCount(characters: number): string {
  if (characters <= 0) return 'your characters';
  if (characters === 1) return 'your character';
  return `all ${characters} of your characters`;
}

export const guildConsentCopy = {
  heading: 'My guilds',
  roster: 'Roster only',
  gear: 'Gear',
  gearBags: 'Gear and bags',
  leave: 'Leave',
  left: 'Left.',
  leaveConfirmLine: (guildName: string, characters: number): string =>
    `Leave ${guildName}? This removes ${characterCount(characters)} from its roster.`,
  leaveConfirmButton: 'Leave',
  leaveCancel: 'Cancel',
  consentSaved: (characters: number): string => `Saved. Applies to ${characterCount(characters)}.`,
} as const;

/** Task L8.3: the Roster tab's "In game" column and the line above the table, from the
 *  guild's newest addon roster snapshot (Task L6.14). The gap is read as DPS: the addon's
 *  hello carries no unit, so the column's hover says it is the addon's own figure. */
export const guildRosterAddonCopy = {
  column: 'In game',
  columnHover: 'Gear gap in DPS, as the addon reports it',
  /** Read before each row's cell, for a screen reader: the rows have no table headers. */
  srLabel: 'In game:',
  notSeen: '—',
  notSeenHover: 'Not seen by the addon in the last snapshot',
  level: (level: number): string => `Level ${level}`,
  /** Shown after the addon's spec, only when it differs from the site's or the site has none. */
  addonTag: 'per addon',
  gap: (gap: number): string => `gap ${gap} DPS`,
  seen: (clock: string): string => `seen ${clock}`,
  /** "seen Mon 21:40": the day is named when `seen_at` is not today in the viewer's date. */
  seenOnDay: (day: string, clock: string): string => `seen ${day} ${clock}`,
  seenHover: (clock: string): string => `Seen at ${clock} on the member's clock`,
  /** The hover on a row whose member shares nothing but the time the addon saw them. */
  withheldHover: 'Details withheld by the member',
  separator: '·',
  /** The build name in the In game slot, so it does not read as a stray word. */
  build: (name: string): string => `build: ${name}`,
  /** `seen` counts the roster rows the snapshot matched; `total` is every roster row. */
  snapshotLine: (by: string, when: string, seen: number, total: number): string =>
    `Roster from the addon · sent by ${by} ${when} · ${seen} of ${total} members seen`,
  noSnapshot: 'No addon snapshot yet — a guild member with the addon and the companion sends one on logout.',
} as const;

/** The bare `/guild` landing: the lookup, the top guilds and the officer line. */
export const guildLandingCopy = {
  heading: 'Guilds',
  lead: "Your guild's progress, roster, raid nights and loot, from the logs and the addon.",
  nameLabel: 'Guild name',
  regionLabel: 'Region',
  rulesetLabel: 'Ruleset',
  submit: 'Open guild',
  emptyName: "Type the guild's name as the game writes it.",
  topHeading: 'Top guilds',
  officerLine: 'Run a guild? Open its page and choose Claim this guild.',
} as const;
