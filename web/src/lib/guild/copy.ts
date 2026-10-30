// web/src/lib/guild/copy.ts
// Every visible string the guild home, claim, settings, join and account-consent views
// use, in the site's own honest-copy voice: no exclamation marks, never a nag inside a
// tool. Wording quoted directly from the spec is marked as such below.
import { uiCopy } from '../ui/copy';
import { possessive } from '../possessive';

export const guildHomeCopy = {
  eyebrow: 'Guild',
  // The line every not-yet-ready state paints (loading, failed, a bad address), so the first
  // thing on screen is server-rendered text that hydration does not change: the page's
  // largest-contentful-paint candidate lands with the HTML instead of waiting for, or
  // being replaced by, the island's scripts (W.16).
  lede: "A guild's progression per boss, pull counts, kill dates and roster bests. Its roster, raid nights and loot sit in the tabs once it has loaded.",
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
  loggedRecently: 'Online today',
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
  // Spec 2026-10-27 section 9.3: when nobody in the guild can approve members yet.
  notVerifiedNoApprover: (guildName: string): string =>
    `You are not verified yet. Nobody in ${guildName} can approve members until its guild master claims it.`,
  sendGuildMasterLink: 'Send your guild master a link',
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
  stateHeading: 'Claim',
  failed: 'That did not load.',
  // Guild name isn't known until the public guild page resolves, so the heading falls
  // back to a nameless form rather than the page ever hardcoding either half itself.
  heading: (guildName: string): string => (guildName === '' ? 'Claim this guild' : `Claim ${guildName}`),
  // Once the viewer holds the claim the page is about the claim, not about making one.
  headingHeld: (guildName: string): string => (guildName === '' ? 'Your claim' : `${guildName} · your claim`),
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
  nextRoster: 'Approve your officers and members',
  nextInvite: 'Make an invite link',
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
  // A claim link made for a plan (spec 2026-10-27 section 9.3): what the link is for, then the order.
  planLine: (guildName: string, planName: string, price: string): string =>
    `This link is for subscribing ${guildName} to the ${planName} plan (${price}). Claim the guild first; then the plan is one step.`,
  // A plain member, the common visitor here: what is true, then what to do with the link below.
  memberHandoff: (guildName: string): string =>
    `You are a member of ${guildName}, so only its guild master can claim it. Send them this page's link.`,
  // After a save on /setup that came from a claim link (spec section 9.3).
  backToClaim: (guildName: string): string => `Back to ${possessive(guildName)} claim`,
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
  eyebrow: 'Guild',
  // Spec 2026-10-27 section 9.3, under the read-only billing sentence.
  sendThemThisLink: 'Send them this link.',
  defaultVisibility: 'Default report visibility',
  officerThreshold: 'Officer rank threshold',
  inviteHeading: 'Invite link',
  inviteHelp: 'Anyone with this link can join as a member.',
  membersHeading: 'Members and access',
  recordsHeading: 'Records and crest',
  recordsOnlyHeading: 'Records',
  rotateButton: 'Rotate invite link',
  // The invite link: rotate asks first, then the full URL with a Copy button. The API keeps
  // only a hash of the token, so the link can be shown once and never again; the page
  // says so instead of offering a "show current link" it cannot honour.
  rotateConfirmLine: 'Rotate the invite link? The old link stops working.',
  rotateConfirmButton: 'Rotate',
  rotateCancel: 'Cancel',
  tokenShownOnce: 'Shown once. Copy it now.',
  linkLabel: 'Invite link',
  copyButton: uiCopy.copyLink,
  copied: uiCopy.copiedLink,
  copyFailed: uiCopy.copyLinkFailed,
  shareButton: uiCopy.shareLink,
  inviteShareTitle: 'Join our guild on Forever Sixty',
  inviteShareText: 'Use this link to join our guild on Forever Sixty.',
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
  /** The styled picker button (P.1): a ghost button, never the browser's own file control. */
  chooseButton: 'Choose an image',
  /** Beside the button until a file is picked. */
  noneChosen: 'No image chosen',
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
  // Spec 2026-10-27 section 9.3, invite joined: the guild's plan, from `/v1/me` `covered`.
  onPlan: (guildName: string, tier: string): string => `${guildName} is on the ${tier} plan.`,
  noPlanYet: (guildName: string): string => `${guildName} has no plan yet.`,
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

const LEAVE_GUILD = 'Leave guild';

export const guildConsentCopy = {
  heading: 'My guilds',
  roster: 'Roster only',
  gear: 'Gear',
  gearBags: 'Gear and bags',
  leave: LEAVE_GUILD,
  left: 'Left.',
  leaveConfirmLine: (guildName: string, characters: number): string =>
    `Leave ${guildName}? This removes ${characterCount(characters)} from its roster.`,
  /** The second sentence of the confirm for the account that holds the guild's claim: leaving
   *  releases it, and the officers who work from it lose their tools. Shown when `/v1/me`'s
   *  entry for the guild carries the release capability. */
  leaveReleasesClaim: (guildName: string): string =>
    `Leaving ${guildName} means you stop running it here. Your officers lose the roster and loot tools and can no longer see your gear.`,
  /** The muted line under it: no self-serve hand-over exists, so the way out is to ask first. */
  leaveNoHandOver:
    'There is no hand-over yet; if you want an officer to run the guild after you, ask us on Discord before you leave.',
  /** The same line cut around its one link, so the page can make "Discord" the site's invite. */
  leaveNoHandOverParts: [
    'There is no hand-over yet; if you want an officer to run the guild after you, ask us on ',
    'Discord',
    ' before you leave.',
  ] as const,
  /** The line a plain officer (tools, but no claim) reads in the confirm. */
  leaveLosesOfficer: (guildName: string): string => `You lose your officer tools in ${guildName}.`,
  /** The label on the consent select, which sets what the guild sees of this account. */
  shareLabel: 'Share with this guild',
  leaveConfirmButton: LEAVE_GUILD,
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
  /** "Gear gap 12 DPS": the figure with the word, so it does not read as a bare number. */
  gap: (gap: number): string => `Gear gap ${gap} DPS`,
  /** "seen 2 hours ago": `when` is the relative phrase for the member's `seen_at`. */
  seen: (when: string): string => `seen ${when}`,
  seenHover: (clock: string): string => `Seen at ${clock} on the member's clock`,
  /** The hover on a row whose member shares nothing but the time the addon saw them. */
  withheldHover: 'Details withheld by the member',
  separator: '·',
  /** Who sent the newest snapshot and when; the counts live on their own line. */
  snapshotLine: (by: string, when: string): string => `Roster from the addon · sent by ${by} ${when}`,
  noSnapshot: 'No roster from the addon yet. Log in with the addon on your guild master and it sends one.',
} as const;

/** The bare `/guild` landing: the lookup, the top guilds and the officer line. */
export const guildLandingCopy = {
  eyebrow: 'Guild',
  heading: 'Guilds',
  lookupHeading: 'Open a guild',
  namePlaceholder: 'Guild name',
  lead: "Your guild's progress, roster, raid nights and loot, from the logs and the addon.",
  nameLabel: 'Guild name',
  regionLabel: 'Region',
  rulesetLabel: 'Ruleset',
  submit: 'Open guild',
  emptyName: "Type the guild's name as the game writes it.",
  topHeading: 'Top guilds',
  officerLine: 'Run a guild? Open its page and choose Claim this guild.',
} as const;

/** The roster as everyone the addon saw, with a word beside each (lifecycle review round 4,
 *  N4/N5/N6/N13): the state words, the rank words, the count line, what approving does and
 *  what only the guild master can do. */
export const guildRosterModelCopy = {
  states: {
    approved: 'Approved',
    waiting: 'Waiting for approval',
    no_account: 'No account yet',
    not_seen: 'Not seen in game yet',
  },
  ranks: { leader: 'Leader', officer: 'Officer', member: 'Member' },
  /** A name only the addon saw: its snapshot carries no guild rank, so none is said. */
  rankNotKnown: 'Rank not known',
  /** "7 on the roster · 2 in game without an account · 1 waiting for approval · 1 not seen in
   *  game yet": the parts that are zero are left out. */
  countLine: (c: {
    on_roster: number;
    in_game_no_account: number;
    waiting: number;
    not_seen: number;
  }): string =>
    [
      `${c.on_roster} on the roster`,
      c.in_game_no_account > 0 ? `${c.in_game_no_account} in game without an account` : '',
      c.waiting > 0 ? `${c.waiting} waiting for approval` : '',
      c.not_seen > 0 ? `${c.not_seen} not seen in game yet` : '',
    ]
      .filter((part) => part !== '')
      .join(' · '),
  gearNote: 'Loot and readiness list members with gear on the site.',
  noAccountFilter: 'No account',
  /** Above the waiting list: what Approve gives the character. */
  approveExplains:
    'Approving lets this character see the roster, loot and readiness; officers can also remove members and change settings.',
  decline: 'Decline',
  /** Asked before a verified character is taken off the roster. */
  removeQuestion: (name: string): string => `Remove ${name} from the roster?`,
  removeButton: 'Remove',
  /** The one muted line under the Remove question: nobody is locked out by it. */
  removeRejoin: 'They can rejoin with the invite link.',
  declineQuestion: (name: string): string => `Remove ${name} from the waiting list? They can ask again.`,
  /** "Approve 4, including 1 officer?" */
  approveAllQuestion: (total: number, officers: number): string =>
    `Approve ${total}, including ${officers} officer${officers === 1 ? '' : 's'}?`,
  confirm: 'Yes',
  cancel: 'Cancel',
} as const;

/** The officer-tools gate (design/specs/2026-10-25-officer-tools-gate.md section 7): the sentence
 *  and its link are separate so the link can be an anchor; together they read
 *  "Award and Nudge are part of the guild plan. Subscribe the guild." */
export const guildGateCopy = {
  sentence: 'Award and Nudge are part of the guild plan. ',
  link: 'Subscribe the guild',
  /** A Guild-tier guild (officer tools on, no Raid guild): launch pricing spec §3.2.8. */
  raidGuildSentence: 'Loot council and the readiness board are part of Raid guild. ',
  raidGuildLink: 'Move to Raid guild',
  awardRefused: "Not awarded. This guild's officer tools have ended.",
  undoRefused: "Not undone. This guild's officer tools have ended.",
  readOnly: 'Read-only',
} as const;
