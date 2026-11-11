// web/src/lib/guild/copy.ts
// Every visible string the guild home, claim, settings, join and account-consent views
// use, in the site's own honest-copy voice: no exclamation marks, never a nag inside a
// tool. Wording quoted directly from the spec is marked as such below.
import { GUILD_TRIAL_DAYS } from '../billing/plans';
import { counted } from '../plural';
import { uiCopy, withRetry } from '../ui/copy';
import { possessive } from '../possessive';

export const guildHomeCopy = {
  eyebrow: 'Guild',
  // The line every not-yet-ready state paints (loading, failed, a bad address), so the first
  // thing on screen is server-rendered text that hydration does not change: the page's
  // largest-contentful-paint candidate lands with the HTML instead of waiting for, or
  // being replaced by, the island's scripts (W.16).
  // A guild address that names no guild the site knows (guild round 3): where it looked, in words.
  notFound: (where: string): string =>
    `No guild by that name on ${where}. Check the spelling and the ruleset.`,
  lookUpAgain: 'Look it up again',
  /** Day one: a guild exists on the site only once a member's character brings it (an addon
   *  sync or a Battle.net import names the guild), so a right name can still miss. */
  notHereYet:
    'Not on Forever Sixty yet. A guild shows up once a member signs in and imports a character, or syncs the addon.',
  notHereYetAction: 'Get set up',
  notHereYetHref: '/setup',
  /** `Did you mean ` + up to three links `{Guild} ({where})`, joined `, ` and ` or `, + `?`
   *  (guild loop round 3). */
  didYouMean: 'Did you mean ',
  suggestionLink: (name: string, where: string): string => `${name} (${where})`,
  /** What goes before the suggestion at `index` of `count`: nothing, `, `, or ` or ` before the last. */
  suggestionJoin: (index: number, count: number): string => {
    if (index === 0) return '';
    return index === count - 1 ? ' or ' : ', ';
  },
  lede: 'Progression, roster, raid nights and loot.',
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
  // The band's claim row (spec 2026-11-05 §3.1).
  claimRow: {
    guildMaster: (guildName: string): string =>
      `You lead ${guildName} in game. Claim it here to approve members and share the invite link.`,
    claimButton: (guildName: string): string => `Claim ${guildName}`,
    officer: (guildName: string): string =>
      `Nobody has claimed ${guildName} here yet. Your guild master claims it.`,
    copyLink: 'Copy the claim link for your guild master',
    copied: uiCopy.copied,
    pending: (guildName: string): string => `A claim on ${guildName} is waiting for a second officer.`,
    open: 'Open the claim',
  },
  claimedBy: 'Claimed by',
  claimedByMaster: 'Claimed by the guild master',
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
  contesting: 'Contesting…',
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
  // What an officer who may claim reads.
  unclaimedForOfficer: 'Nobody has claimed this guild.',
  // Claiming asks first; the rules sit in their own card right below, so the question is bare.
  claimConfirmLine: (guildName: string): string => `Claim ${guildName === '' ? 'this guild' : guildName}?`,
  claimConfirmButton: 'Yes, claim it',
  claiming: 'Claiming…',
  // After the claim: the page says what the viewer holds and where to go next.
  youHold: (guildName: string): string => `You hold ${guildName === '' ? 'this guild' : guildName}.`,
  nextHeading: 'Next',
  nextRoster: 'Approve your officers and members',
  releaseConfirmLine: (guildName: string): string =>
    `Release your claim on ${guildName === '' ? 'this guild' : guildName}? Officers will have to claim it again.`,
  releaseConfirmButton: 'Release',
  releasing: 'Releasing…',
  cancel: 'Cancel',
  claimedBySomeoneElse: (name: string, since?: string): string =>
    since === undefined ? `Claimed by ${name}.` : `Claimed by ${name} since ${since}.`,
  theGuildMaster: 'the guild master',
  // `expiresAt` is only known in the same session the viewer just triggered the pending
  // claim (the POST .../claim response) -- a pending claim loaded later from settings
  // carries no expiry, hence the optional form rather than a second, separate string.
  pending: (expiresAt?: string): string =>
    expiresAt === undefined
      ? 'A claim is pending, confirmed by a second officer or the guild master.'
      : `A claim is pending, confirmed by a second officer or the guild master. Expires ${expiresAt.slice(0, 10)}.`,
  claimButton: 'Claim this guild',
  confirmButton: 'Confirm this claim',
  confirming: 'Confirming…',
  releaseButton: 'Release claim',
  notEligible: 'Only an officer or the guild master of this guild can claim it.',
  // A claim link made for a plan (spec 2026-10-27 section 9.3): what the link is for, then the order.
  planLine: (guildName: string, planName: string, price: string): string =>
    `This link is for subscribing ${guildName} to the ${planName} plan (${price}). Claim the guild first, then subscribe.`,
  // A plain member, the common visitor here: what is true, then what to do with the link below.
  memberHandoff: (guildName: string): string =>
    `You are a member of ${guildName}, so only its guild master can claim it. Send them this page's link.`,
  // After a save on /setup that came from a claim link (spec section 9.3).
  backToClaim: (guildName: string): string => `Back to ${possessive(guildName)} claim`,
  // /setup's Notice over that link, once the account knows a character in the guild (setup spec 4.6).
  setupClaimReady: (guildName: string): string => `You have a character in ${guildName}.`,
  // Signed in, but no character of theirs in this guild as far as the site knows.
  notAMember: 'None of your characters is in this guild. Add one on Get set up.',
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
  // Spec 2026-10-27 section 9.3, under the read-only billing sentence.
  sendThemThisLink: 'Send them this link.',
  defaultVisibility: 'Default report visibility',
  officerThreshold: 'Officer rank threshold',
  membersHeading: 'Members and access',
  recordsHeading: 'Records and crest',
  recordsOnlyHeading: 'Records',
  // The invite link (spec 2026-11-05 §3.3, §4.1). The API keeps the link sealed so it shows again
  // on a later visit; where it could not keep it (no key on the server, or a link made before it
  // could), the link shows only right after it is made and the row says so.
  newLinkButton: 'Make a new link',
  firstLinkButton: 'Make the invite link',
  making: 'Making…',
  noLinkYet: 'Make a link to share in guild chat or Discord.',
  inviteFailed: withRetry('That did not work.'),
  rotateConfirmLine:
    'Make a new link? The old one stops working, so anyone you sent it to needs the new one.',
  rotateCancel: 'Cancel',
  tokenShownOnce: 'Copy it now. It cannot be shown again.',
  linkLabel: 'Invite link',
  copyButton: uiCopy.copyLink,
  copied: uiCopy.copiedLink,
  copyFailed: uiCopy.copyLinkFailed,
  shareButton: uiCopy.shareLink,
  inviteShareTitle: 'Join our guild on Forever Sixty',
  currentLinkMade: (date: string): string =>
    `Made ${date}. It still works, but it cannot be shown again. Make a new one to copy it.`,
  /** A kept link's line: `Made 12 Nov.`, or `Made 12 Nov by` + the maker's name + `.`. */
  linkMade: (date: string): string => `Made ${date}.`,
  linkMadeBy: (date: string): string => `Made ${date} by`,
  saved: 'Saved.',
  // Shown when a save or rotate action itself fails (as opposed to the initial load).
  actionFailed: 'That did not save; try again.',
  frozenNotice: 'This guild’s claim is contested. Officer actions are frozen until a moderator resolves it.',
} as const;

// Guild crest round (docs/contracts/2026-10-05-guild-crest-api.md), Settings tab block.
export const guildCrestCopy = {
  heading: 'Guild crest',
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
  /** The one Section's heading under the guild's head (web layout spec 2026-11-05 §3A). */
  sectionHeading: 'Your invite',
  guildLine: (realm: string, members: number | undefined): string =>
    members === undefined ? realm : `${realm} · ${counted(members, 'member')} on Forever Sixty`,
  invited: (guildName: string): string => `You are invited to join ${guildName}.`,
  /** After the inviter's crest and name (spec 2026-11-05 §3.4 row 2). */
  invitedBy: 'invited you to join.',
  joinAs: (name: string): string => `You join as ${name}`,
  level: (level: number): string => `Level ${level}`,
  noneOnRoster: (guildName: string): string => `None of your characters is in ${guildName} in game yet.`,
  noneOnRosterSub: (name: string): string => `Officers see ${name} until one of yours shows on the roster.`,
  joinWithAccount: 'You join with your account',
  joinWithAccountSub: 'Your characters show once you sync the addon or import from Battle.net.',
  verifiedAtOnce: 'Joining by this link verifies you at once. No officer has to approve you.',
  consequence:
    'Joining puts your characters on its roster and shares your gear with its officers; you can change that on your account.',
  joinButton: (guildName: string): string => `Join ${guildName}`,
  joining: 'Joining…',
  joined: 'You joined the guild.',
  // Spec 2026-10-27 section 9.3, invite joined: the guild's plan, from `/v1/me` `covered`.
  onPlan: (guildName: string, tier: string): string => `${guildName} is on the ${tier} plan.`,
  noPlanYet: (guildName: string): string => `${guildName} is on the free plan.`,
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
 *  guild's newest addon roster snapshot (Task L6.14). Its gear gap is the readiness board's
 *  (guild round 3), read as DPS by the chip's own words (`gap`). */
export const guildRosterAddonCopy = {
  column: 'In game',
  /** Read before each row's cell, for a screen reader: the rows have no table headers. */
  srLabel: 'In game:',
  level: (level: number): string => `Level ${level}`,
  /** "Gear gap 12 DPS": the figure with the word, so it does not read as a bare number. */
  gap: (gap: number): string => `Gear gap ${gap} DPS`,
  /** "seen 2 hours ago": `when` is the relative phrase for the member's `seen_at`. */
  seen: (when: string): string => `seen ${when}`,
  seenHover: (clock: string): string => `Seen at ${clock} on the member's clock`,
  /** The hover on a row whose member shares nothing but the time the addon saw them. */
  withheldHover: 'Details withheld by the member',
  separator: '·',
} as const;

/** The bare `/guild` landing: the lookup, the top guilds and the officer line. */
export const guildLandingCopy = {
  /** The bare /guild page's <title> and meta description, written for search. */
  docTitle: 'Find a WoW Forever Guild: Progress, Roster and Raids',
  docDescription:
    'Look up any World of Warcraft: Forever guild by name and realm: progression per boss, kill dates, the roster and its raid nights, from logs and the Forever Sixty addon.',
  eyebrow: 'Guild',
  heading: 'Guilds',
  lookupHeading: 'Open a guild',
  mineHeading: 'Your guilds',
  /** A Your guilds row's line: where the guild lives and, when the API counts it, its roster. */
  mineWhere: (where: string, members: number | null): string =>
    members === null ? where : `${where} · ${members === 1 ? '1 member' : `${members} members`}`,
  namePlaceholder: 'Guild name',
  nameLabel: 'Guild name',
  regionLabel: 'Region',
  rulesetLabel: 'Ruleset',
  submit: 'Open guild',
  emptyName: "Type the guild's name as the game writes it.",
  topHeading: 'Top guilds',
  /** The rail's officer section (web layout spec 2026-11-05 §6.1): its heading and its one line. */
  runHeading: 'Run a guild?',
  officerLine: 'Sign in and import your character, then open the guild page and choose Claim this guild.',
  /** In the officer section's place for a player who already holds a claim (guild loop round 3). */
  heldHeading: 'Your claim',
  holdsClaim: (guildName: string): string => `You hold ${possessive(guildName)} claim.`,
  signIn: 'Sign in with Battle.net',
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
  ranks: { leader: 'Guild master', officer: 'Officer', member: 'Member' },
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
  noAccountFilter: 'No account',
  /** The roster's column heads from lg: the units the cells leave out. */
  heads: { raider: 'Raider', level: 'Item level', raids: 'Raids', best: 'Top output', rating: 'Rating' },
  decline: 'Decline',
  approving: 'Approving…',
  /** Asked before a verified character is taken off the roster. */
  removeQuestion: (name: string): string => `Remove ${name} from the roster?`,
  removeButton: 'Remove',
  /** The one muted line under the Remove question: nobody is locked out by it. */
  removeRejoin: 'They can rejoin with the invite link.',
  /** The same line for a member who came in by the invite link: it works until an officer replaces it. */
  removeInvitedRejoin: 'The invite link still lets them back in until you rotate it in Settings.',
  declineQuestion: (name: string): string => `Remove ${name} from the waiting list? They can ask again.`,
  /** "Approve 4, including 1 officer?" */
  approveAllQuestion: (total: number, officers: number): string =>
    `Approve ${total}, including ${officers} officer${officers === 1 ? '' : 's'}?`,
  confirm: 'Yes',
  cancel: 'Cancel',
} as const;

/** The officer-tools gate (design/specs/2026-10-25-officer-tools-gate.md section 7): the sentence
 *  and its link are separate so the link can be an anchor; together they read
 *  "Award and Nudge are part of the guild plan. Subscribe your guild." */
export const guildGateCopy = {
  // Spec 2026-11-05 §3.2, part 1 by feature; part 2 by viewer (lib/guild/plan-gate.ts).
  lead: {
    raids: 'Comp targets, pings, assignments and log matching are part of the Guild and Raid guild plans.',
    loot: 'Awarding loot is part of Raid guild.',
    readiness: 'Nudge is part of Raid guild.',
  },
  startTrial: `Start the ${GUILD_TRIAL_DAYS}-day trial`,
  subscribe: 'Subscribe your guild',
  moveToRaidGuild: 'Move to Raid guild',
  /** A plan Forever Sixty granted (guild round 3): no Move control exists, so the gate asks us. */
  grantedMove: {
    text: 'Your plan was granted by Forever Sixty. Ask us on Discord to move it to Raid guild.',
    link: 'Ask us on Discord',
  },
  ask: 'Ask ',
  askTail: ' to turn it on.',
  askMaster: 'Ask your guild master to turn it on.',
  claimFirst: 'Claim the guild first',
  masterClaimsFirst: 'Your guild master claims the guild first.',
  awardRefused: "Not awarded. This guild's officer tools have ended.",
  undoRefused: "Not undone. This guild's officer tools have ended.",
  readOnly: 'Read-only',
} as const;
