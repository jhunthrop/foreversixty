// web/src/lib/guild/raid-copy.ts
// Every string of design/specs/2026-10-26-guild-raids.md 6.1, as amended by the raid comp spec
// (2026-11-05-guild-raid-comp.md 7), in one module: constants where the words are fixed, functions
// where the data fills them. No number is typed into prose; raid-copy.test.ts reads the spec tables
// themselves, so a drift in either file fails.
import type {
  GuildRaidGroup,
  GuildRaidRole,
  GuildRaidSignup,
  GuildRaidView,
  GuildSignupState,
  GuildViewerRole,
  GuildZoneKind,
  NeedsRaid,
  NextRaidView,
} from './api';
import { calendarDayIn, clock, viewerZone, weekdayDayMonth, weekdayName, weekdayShort } from '../dates';
import { MINUTE_MS, spokenDuration } from '../format/duration';
import { SITE_LOCALE } from '../format/number';
import { countdownText } from '../ui/countdown';
import { instantToWallClock, offsetsDiffer, wallClockToInstant, zoneAbbreviation } from './raid-time';

/** A day further out than this is written as a date, not a weekday (spec 6.1 `{when}`). */
const WEEKDAY_WINDOW_DAYS = 6;
/** A start on the guild's own day reads `Today` before this hour of the guild's clock and
 *  `Tonight` from it (guild loop round 3: a 6:49 am start read `Tonight`). */
export const TONIGHT_FROM_HOUR = 17;
/** The middle dot between a line's parts. Svelte trims a literal's edge spaces, so pages write it
 *  as an expression (`{SEP}`). */
export const SEP = ' · ';

/** The raid page's Post now / Repost (owner ruling 2026-11-09). */
export const raidPostCopy = {
  postNow: 'Post now',
  repost: 'Repost',
  notPosted: 'Not on Discord yet.',
  postedTo: (channel: string | null, when: string): string => `Posted to ${channel ?? 'Discord'} ${when}.`,
  posted: 'Posted.',
  reposted: 'Reposted as a new message.',
  noChannel: 'Pick a Discord channel in Settings first.',
  failed: 'The post did not go out. Try again.',
} as const;

export const raidCopy = {
  /** A raider's unmet raid requirements, the one word on every board ("2 missing"). */
  missing: (count: number): string => `${count} missing`,
  planButton: 'Plan an event',
  cap: 'A guild plans at most 12 raids ahead.',
  subheads: {
    planned: 'Planned',
    raidNights: 'Raid nights',
    yourSignup: 'Your sign-up',
    signups: 'Sign-ups',
    notes: 'Notes',
    tanks: 'Tanks',
    healers: 'Healers',
    melee: 'Melee',
    ranged: 'Ranged',
    late: 'Late',
    bench: 'Bench',
    out: 'Out',
    notSigned: 'Not signed up',
  },
  segments: { in: 'In', late: 'Late', bench: 'Bench', out: 'Out' },
  roleLabels: { tank: 'Tank', healer: 'Healer', dps: 'DPS' } satisfies Record<GuildRaidRole, string>,
  /** Beside a sign-up control: the role the answer carries ("as DPS"). */
  signingAs: (role: string): string => `as ${role}`,
  emptyPlanned: {
    member: 'Nothing planned yet. Officers post dungeon runs, events and raids here.',
    officer: 'Nothing planned yet. Post a dungeon run or a guild event and the guild can sign up.',
  },
  bandNone: {
    member: 'Nothing planned yet.',
    officer: 'Nothing planned yet.',
    planLink: 'Plan an event',
  },
  placeholders: {
    note: 'Note for the officers (optional)',
    chooseWhat: 'Choose a dungeon, event or raid',
    eventTitle: 'What is it? (a PvP premade, a quest group)',
  },
  stripStatus: {
    saved: 'Saved.',
    failed: 'That did not save. Try again.',
    movedToBench: (name: string): string => `${name} moved to Bench.`,
    left: (n: number): string => `${n} left`,
    /** On an answered character that is not the account's roster row: `Officers see Quillfeather (Late).` */
    officersSee: (name: string, state: string): string => `Officers see ${name} (${state}).`,
  },
  /** The roster row's line naming its account's other answered characters (raids spec 3.7):
   *  `Also: Rootbound (Restoration Druid, Bench)`, several joined by `; ` (guild round 3). */
  also: (parts: readonly string[]): string => `Also: ${parts.join('; ')}`,
  /** One character of the Also line: `Rootbound (Restoration Druid, Bench)`. */
  alsoPart: (name: string, specClass: string, state: string): string => `${name} (${specClass}, ${state})`,
  /** The strip's one spec picker (raid comp spec 3.2, 7). */
  spec: {
    /** The select's visually hidden label. */
    label: (name: string): string => `Spec for ${name}`,
    placeholder: 'Choose a spec',
    /** In, Late or Bench pressed with no spec chosen. */
    pickFirst: 'Pick a spec first.',
  },
  waitingCharacter: 'Waiting for approval. You can sign it up once an officer approves it.',
  raidStates: {
    locked: 'Sign-ups are locked. An officer can still change yours.',
  },
  races: {
    locked: 'This raid was locked while you were signing up.',
    cancelled: 'This raid was cancelled while you were signing up.',
    officerToolsEnded: "Not saved. This guild's officer tools have ended.",
  },
  poll: 'Live updates paused. Reload to see the latest.',
  /** Any officer write the API refused with 429: the wait, not "did not save". */
  rateLimited: 'Too many requests right now. Wait a minute and try again.',
  officerRow: {
    edit: 'Edit',
    lock: 'Lock sign-ups',
    locking: 'Locking…',
    unlock: 'Unlock sign-ups',
    copy: 'Copy the sign-up list',
    copied: 'Copied',
    failed: 'That did not save; try again.',
  },
  cancelConfirm: {
    postToDiscord: 'Post the cancellation to Discord',
    confirming: 'Cancelling…',
    keep: 'Keep it',
  },
  copyFailed: 'Could not copy. Select the list and copy it.',
  /** The copied list's class line head (raid comp spec 4.2). */
  copyClasses: 'Classes',
  form: {
    what: 'What',
    groupSize: 'Group size',
    noLimit: 'No limit',
    groups: { event: 'Guild event', dungeons: 'Dungeons', raids: 'Raids' },
    needsKey: 'Needs a key: ',
    /** A key with a keyless way in: `Key to the City opens the service entrance.` */
    keyOpens: (door: string): string => ` opens ${door}.`,
    opens: (raid: string, day: string): string => `${raid} opens ${day}.`,
    title: 'Title',
    date: 'Date',
    time: 'Time',
    duration: 'Duration',
    guildZone: 'Guild time zone',
    notes: 'Notes',
    repeat: 'Also plan the next 3 weeks at this time',
    postToDiscord: 'Post to Discord',
    postToDiscordIfMoved: 'Post to Discord if the time changes',
    remindDayBefore: 'Remind the day before',
    postNow: (channel: string | null): string =>
      `Posts to ${channel ?? 'your Discord channel'} as soon as you save.`,
    postRepeat: 'All four raids post as soon as you save.',
    postOff: 'This raid will not post to Discord.',
    remind: 'Posts a reminder 24 hours before the start.',
    remindNeedsPost: 'Turn on Post to Discord to send a reminder.',
    posted: (channel: string): string => `Posted to ${channel}. The post updates itself.`,
    reminded: 'Reminder sent.',
    submitCreate: 'Post the event',
    submitCreating: 'Posting…',
    submitEdit: 'Save changes',
    submitEditing: 'Saving…',
    cancel: 'Cancel',
    forYou: (when: string): string => `For you: ${when}.`,
    forYouSame: 'For you: same time.',
    markdownHelper: 'Markdown: bold, lists and links. Members see this on the raid page.',
    onlyFit: (n: number): string => `Only ${n} more fit: a guild plans at most 12 raids ahead.`,
    connectDiscord: 'Connect a Discord channel in Settings to post raids there.',
    connectDiscordLink: 'Settings',
    errors: {
      what: 'Choose a dungeon, event or raid.',
      eventTitle: 'Give the event a title.',
      opens: (raid: string, day: string): string => `${raid} opens ${day}. Pick that day or later.`,
      date: 'Pick a date.',
      time: 'Pick a time after now.',
      cap: 'A guild plans at most 12 raids ahead.',
      failed: 'That did not save. Try again.',
    },
  },
  emptyGroup: {
    tank: 'No tanks yet.',
    healer: 'No healers yet.',
    melee: 'No melee yet.',
    ranged: 'No ranged yet.',
  } satisfies Record<GuildRaidGroup, string>,
  notVisible: {
    members: (guild: string): string => `Planned raids are for members of ${guild}.`,
    unverified: 'You are waiting for approval. Planned raids show once an officer approves you.',
  },
  loadErrors: { raid: 'This raid did not load.', planned: 'Planned raids did not load.' },
  backLink: (guild: string): string => `‹ ${guild} · Raids`,
  /** The raid page's own words outside the 6.1 table (spec 3.2, 3.7, 4 "Page states"). */
  page: {
    eyebrow: 'Raid',
    loading: 'Loading the raid',
    /** The officer state select's accessible name (3.7). */
    signupFor: (name: string): string => `Sign-up for ${name}`,
    /** The not-signed row's resting option; its accessible name is `Not signed up`. */
    notSignedMark: '—',
    /** Between the start and end clocks: `8:00 pm to 11:00 pm`. */
    timeTo: 'to',
  },
} as const;

/** What each kind of night is called in a sentence (guild events spec 3.5). */
const KIND_WORDS: Record<GuildZoneKind, string> = { raid: 'raid', dungeon: 'run', event: 'event' };
const KIND_EYEBROWS: Record<GuildZoneKind, string> = { raid: 'Raid', dungeon: 'Dungeon', event: 'Event' };

/** The zone kind a view carries; a view from before kinds is a raid. */
export function kindOf(view: { zone_kind?: GuildZoneKind }): GuildZoneKind {
  return view.zone_kind ?? 'raid';
}

/** Whether a night counts roles. A guild event is a meeting: no role grid, no class row, no role
 *  words, no spec picker; raids and dungeon runs (which need a tank and a healer) keep them
 *  (guild events spec 3.5, guild round 3). */
export function kindCountsRoles(kind: GuildZoneKind): boolean {
  return kind !== 'event';
}

/** The zone beside a title that is not the zone's name, with the size an officer picked for a guild
 *  event: `Guild event · 20 players` (a bare `20` read as nothing, guild round 3). The planned row's
 *  second line and the event page's h1 read it. */
export function zoneLineOf(view: {
  zone_kind?: GuildZoneKind;
  zone_name: string;
  size: number | null;
}): string {
  return kindOf(view) === 'event' && view.size !== null
    ? `${view.zone_name}${SEP}${view.size} players`
    : view.zone_name;
}

/** The raid's kind while the page loads, in the Label role: `Raid`, `Dungeon`, `Event`. */
export const eyebrowOf = (kind: GuildZoneKind): string => KIND_EYEBROWS[kind];
/** `Cancel raid`, `Cancel run`, `Cancel event`. */
export const cancelActionOf = (kind: GuildZoneKind): string => `Cancel ${KIND_WORDS[kind]}`;
/** `This raid has ended.`, `This run has ended.`, `This event has ended.` */
export const endedLineOf = (kind: GuildZoneKind): string => `This ${KIND_WORDS[kind]} has ended.`;
/** `This raid was cancelled.`, `This run was cancelled.`, `This event was cancelled.` */
export const cancelledLineOf = (kind: GuildZoneKind): string => `This ${KIND_WORDS[kind]} was cancelled.`;
/** The cancel confirm's question. */
export const cancelQuestionOf = (kind: GuildZoneKind): string =>
  `Cancel this ${KIND_WORDS[kind]}? Members see it as cancelled.`;
/** The cancel confirm's button. */
export const cancelConfirmOf = (kind: GuildZoneKind): string => `Cancel the ${KIND_WORDS[kind]}`;
/** The notes panel with nothing in it: `No notes for this run.` / `No notes yet. Edit the event to add some.` */
export const notesEmptyOf = (kind: GuildZoneKind, officer: boolean): string =>
  officer
    ? `No notes yet. Edit the ${KIND_WORDS[kind]} to add some.`
    : `No notes for this ${KIND_WORDS[kind]}.`;

/** The needs line reads raids only (guild events spec 3.4): a dungeon run or an event at the head of
 *  the band leaves it at `before the next raid`, unless the API names the raid (`needsDeadline`). */
export function raidDeadlineStart(next: NextRaidView | null | undefined): string | undefined {
  return next != null && kindOf(next) === 'raid' ? next.starts_at : undefined;
}

const SIGNED_UP_STATE_WORDS: Record<GuildSignupState, string> = {
  in: raidCopy.segments.in,
  late: raidCopy.segments.late,
  bench: raidCopy.segments.bench,
  out: raidCopy.segments.out,
};

function plural(n: number, singular: string, pluralForm: string): string {
  return n === 1 ? singular : pluralForm;
}

/** A count line at zero: never `0 signed up` (guild round 2). */
export const NOBODY_SIGNED_UP = 'Nobody signed up yet';

/** `23 of 40 signed up`; `23 signed up` while the size is unknown (ruling R6); `Nobody signed up yet` at 0. */
export function countLine(n: number, size: number | null): string {
  if (n === 0) return NOBODY_SIGNED_UP;
  return size === null ? `${n} signed up` : `${n} of ${size} signed up`;
}

/** `23/40`; the bare count while the size is unknown. */
export function countRatio(n: number, size: number | null): string {
  return size === null ? `${n}` : `${n}/${size}`;
}

/** The raid comp's words (raid comp spec 7): the count clauses, the class row and the spec
 *  picker's options. The groups count In rows, so the late clause (only when someone is late)
 *  makes the parts add up to the signed-up count. */
export const compCopy = {
  tanks: (n: number): string => `${n} ${plural(n, 'tank', 'tanks')}`,
  healers: (n: number): string => `${n} ${plural(n, 'healer', 'healers')}`,
  melee: (n: number): string => `${n} melee`,
  ranged: (n: number): string => `${n} ranged`,
  /** `2 late`; `2 late (1 tank, 1 healer)`; a single part that is all of them drops its number:
   *  `1 late (healer)`. */
  late: (n: number, tanks: number, healers: number): string => {
    const parts = [
      { n: tanks, word: plural(tanks, 'tank', 'tanks') },
      { n: healers, word: plural(healers, 'healer', 'healers') },
    ].filter((part) => part.n > 0);
    if (parts.length === 0) return `${n} late`;
    const only = parts.length === 1 && parts[0]?.n === n;
    const inner = only ? (parts[0]?.word ?? '') : parts.map((part) => `${part.n} ${part.word}`).join(', ');
    return `${n} late (${inner})`;
  },
  /** The class row's list name. */
  classRowLabel: 'Classes signed up',
  /** `Holy 3`. */
  specCount: (spec: string, n: number): string => `${spec} ${n}`,
  /** `Fury (tank)`: an off-tank's part. */
  offTank: (spec: string): string => `${spec} (tank)`,
  unknownSpec: 'Unknown',
  /** `Priest 5: Holy 3, Discipline 1, Shadow 1`; `Mage 0` with nobody in. */
  classItem: (className: string, n: number, specs: string): string =>
    specs === '' ? `${className} ${n}` : `${className} ${n}: ${specs}`,
  /** `Holy · Healer`. */
  specOption: (spec: string, role: string): string => `${spec}${SEP}${role}`,
  /** `9 pm EST your time`. */
  yourTime: (clock: string): string => `${clock} your time`,
} as const;

/** `23 of 40 in`; `23 in` while the size is unknown: the unsigned band line's count. No verb, so
 *  `1 of 5 in` never reads `1 of 5 are in` (guild loop round 3). */
function inLine(n: number, size: number | null): string {
  return size === null ? `${n} in` : `${n} of ${size} in`;
}

export function showAll(n: number): string {
  return `Show all ${n} planned raids`;
}

/** The needs line's deadline: `before Thursday`, or `before the next raid` with none planned. */
export function needsBefore(weekday: string | null): string {
  return weekday === null ? 'before the next raid' : `before ${weekday}`;
}

/** The needs line's deadline from `home.next_raid.starts_at`: the raid's weekday in the viewer's
 *  zone, or `before the next raid` with none planned (spec 3.4, 6.1 "Needs line"). */
export function nextRaidDeadline(startsAt: string | undefined, viewerTz: string = viewerZone()): string {
  if (startsAt === undefined) return needsBefore(null);
  return needsBefore(weekdayName(new Date(startsAt), viewerTz));
}

/** The needs line's deadline (guild loop round 3): with the raid named by the API, `before
 *  Onyxia's Lair, Thu 10 Dec` on the guild's calendar, so a dungeon run at the head of the band no
 *  longer leaves it at `before the next raid`; without one, `nextRaidDeadline` as before. */
export function needsDeadline(
  needsRaid: NeedsRaid | null | undefined,
  next: NextRaidView | null | undefined,
  viewerTz: string = viewerZone(),
): string {
  if (needsRaid == null) return nextRaidDeadline(raidDeadlineStart(next), viewerTz);
  const tz = next?.guild_timezone ?? viewerTz;
  return `before ${needsRaid.title}, ${weekdayDayMonth(new Date(needsRaid.starts_at), tz)}`;
}

/** `8 pm CST`, `8:30 pm CST`: a clock with its zone abbreviation. */
export function zonedClock(at: Date, tz: string, locale: string = SITE_LOCALE): string {
  return `${clock(at, tz, { locale })} ${zoneAbbreviation(tz, at)}`;
}

function guildClock(startsAt: Date, tz: string, locale: string): string {
  return zonedClock(startsAt, tz, locale);
}

/** The viewer's own clock under the guild's (raid comp spec 3.1): `9 pm EST your time`, with the
 *  short weekday first when the viewer's date is another (`Fri 2 am GMT your time`); `null`
 *  when the two zones read the same offset. */
export function yourTimeLine(
  startsAt: Date,
  guildTz: string,
  viewerTz: string,
  locale: string = SITE_LOCALE,
): string | null {
  if (!offsetsDiffer(startsAt, viewerTz, guildTz)) return null;
  const time = zonedClock(startsAt, viewerTz, locale);
  const otherDay = calendarDayIn(startsAt, viewerTz) !== calendarDayIn(startsAt, guildTz);
  return compCopy.yourTime(otherDay ? `${weekdayShort(startsAt, viewerTz)} ${time}` : time);
}

/** The form's line under the zone select: the entered guild wall time as the viewer reads it. */
export function forYouLine(
  date: string,
  time: string,
  guildTz: string,
  viewerTz: string,
  locale: string = SITE_LOCALE,
): string {
  const instant = wallClockToInstant(date, time, guildTz);
  if (!offsetsDiffer(instant, guildTz, viewerTz)) return raidCopy.form.forYouSame;
  return raidCopy.form.forYou(
    `${weekdayDayMonth(instant, viewerTz)}, ${clock(instant, viewerTz, { locale })}`,
  );
}

/** `Today` before 5 pm on the guild's clock, `Tonight` from it. */
function sameDayWord(startsAt: Date, guildTz: string): string {
  const hour = Number(instantToWallClock(startsAt, guildTz).time.slice(0, 2));
  return hour < TONIGHT_FROM_HOUR ? 'Today' : 'Tonight';
}

/** The band's `{when}` in the guild's clock first (raid comp spec 3.4): `Tonight 8 pm CST`
 *  (`Today 10 am CST` before 5 pm),
 *  `Thursday 8 pm CST`, `Thu 17 Dec, 8 pm CST`, then ` (9 pm EST your time)` when the viewer's
 *  offset differs. */
export function whenPhrase(
  startsAt: Date,
  now: Date,
  locale: string,
  guildTz: string,
  viewerTz: string = viewerZone(),
): string {
  const time = zonedClock(startsAt, guildTz, locale);
  const days = calendarDayIn(startsAt, guildTz) - calendarDayIn(now, guildTz);
  let phrase: string;
  if (days === 0) phrase = `${sameDayWord(startsAt, guildTz)} ${time}`;
  else if (days === 1) phrase = `Tomorrow ${time}`;
  else if (days > 1 && days <= WEEKDAY_WINDOW_DAYS) {
    phrase = `${weekdayName(startsAt, guildTz)} ${time}`;
  } else phrase = `${weekdayDayMonth(startsAt, guildTz)}, ${time}`;
  const yours = yourTimeLine(startsAt, guildTz, viewerTz, locale);
  return yours === null ? phrase : `${phrase} (${yours})`;
}

/** `in 1 day, 4 hours, 12 min` (the one countdown), then `started 25 min ago`, `ended 2 h ago`. */
/** A past span to the whole minute: the band steps on the minute tick, so seconds would sit still. */
function minutesAgo(ms: number): string {
  return spokenDuration(Math.round(ms / MINUTE_MS) * MINUTE_MS);
}

export function countdown(startsAt: Date, endsAt: Date, now: Date): string {
  const untilStart = startsAt.getTime() - now.getTime();
  if (Math.abs(untilStart) < MINUTE_MS) return 'starting now';
  if (untilStart > 0) return `in ${countdownText(startsAt, now)}`;
  if (now.getTime() < endsAt.getTime()) return `started ${minutesAgo(-untilStart)} ago`;
  return `ended ${minutesAgo(now.getTime() - endsAt.getTime())} ago`;
}

export interface BandLineArgs {
  /** The `whenPhrase` output. */
  when: string;
  /** The viewer's state; `null` while unsigned. */
  state: GuildSignupState | null;
  role?: GuildRaidRole;
  /** Set when the account has two or more characters. */
  characterName?: string;
  n: number;
  size: number | null;
}

/** The clause around its state word: `you are ` / `In` / ` as DPS`; no word while unsigned. The
 *  Overview's own row reads it alone, after the character's spec and level. */
export function bandClause({ state, role, characterName }: BandLineArgs): BandLineParts {
  if (state === null) return { head: 'you have not signed up yet', word: '', tail: '' };
  const word = SIGNED_UP_STATE_WORDS[state];
  if (state === 'bench') return { head: 'you are on the ', word, tail: '' };
  if (state === 'out') return { head: 'you are ', word, tail: '' };
  const on = characterName === undefined ? '' : ` on ${characterName}`;
  const as = role === undefined ? '' : ` as ${raidCopy.roleLabels[role]}`;
  return { head: 'you are ', word, tail: `${on}${as}` };
}

/** The band sentence in three parts, so the state word can read 600 `--strong` (spec 3.4). */
export interface BandLineParts {
  head: string;
  word: string;
  tail: string;
}

export function bandLineParts(args: BandLineArgs): BandLineParts {
  const clause = bandClause(args);
  const count =
    args.state === null
      ? `. ${inLine(args.n, args.size)}.`
      : `, ${args.n === 0 ? NOBODY_SIGNED_UP.toLowerCase() : countLine(args.n, args.size)}.`;
  return { head: `${args.when}, ${clause.head}`, word: clause.word, tail: `${clause.tail}${count}` };
}

/** The guild band's raid sentence: `Tonight 8 pm, you are In as DPS, 23 of 40 signed up.` */
export function bandLine(args: BandLineArgs): string {
  const { head, word, tail } = bandLineParts(args);
  return `${head}${word}${tail}`;
}

/** Who sees the band line (spec 3.4, R1): verified members and officers, never a visitor or a
 *  member still waiting for approval. */
export function showsBandLine(verified: boolean, role: GuildViewerRole): boolean {
  return verified && role !== 'public';
}

/** What the copied list needs from the comp rules (raid-comp.ts), passed in so this module stays
 *  free of the spec catalogue. */
export interface CopyListComp {
  /** The breakdown clauses: `3 tanks`, ..., `1 late (healer)`. */
  breakdown: readonly string[];
  /** The class row's items in the game's order. */
  classes: readonly { name: string; count: number }[];
  /** A row's roster group. */
  groupOf: (row: GuildRaidSignup) => GuildRaidGroup;
  /** `Protection Warrior`, or the class alone. */
  specClass: (row: GuildRaidSignup) => string;
  /** The roster's order within a group. */
  order: (a: GuildRaidSignup, b: GuildRaidSignup) => number;
  /** False for a guild event: the In rows are one list, not four role groups. */
  roles: boolean;
}

function namesLine(label: string, rows: readonly GuildRaidSignup[], comp: CopyListComp): string | null {
  if (rows.length === 0) return null;
  return `${label}: ${[...rows]
    .sort(comp.order)
    .map((row) => `${row.name} (${comp.specClass(row)})`)
    .join(', ')}`;
}

/** Discord's `<t:{unix}:F>` stamp, which each reader's client draws in their own zone. */
function discordStamp(at: Date): string {
  return `<t:${Math.floor(at.getTime() / 1000)}:F>`;
}

/** Spec 4 "Copy the sign-up list" (raid comp spec 4.2): plain text in guild time, the class line,
 *  then the groups in roster order, empty groups and classes at 0 left out. */
export function copyList(
  view: GuildRaidView,
  tz: string,
  comp: CopyListComp,
  locale: string = SITE_LOCALE,
): string {
  const startsAt = new Date(view.starts_at);
  const heads = raidCopy.subheads;
  const { counts } = view;
  const header = [
    view.title,
    `${weekdayDayMonth(startsAt, tz)}, ${guildClock(startsAt, tz, locale)}`,
    discordStamp(startsAt),
  ].join(SEP);
  const size =
    view.size === null
      ? `${counts.in + counts.late} signed up`
      : `${counts.in + counts.late} of ${view.size}`;
  const inState = (state: GuildSignupState): GuildRaidSignup[] =>
    view.signups.filter((s) => s.state === state);
  const inGroup = (group: GuildRaidGroup): GuildRaidSignup[] =>
    inState('in').filter((s) => comp.groupOf(s) === group);
  const inLines = comp.roles
    ? [
        namesLine(heads.tanks, inGroup('tank'), comp),
        namesLine(heads.healers, inGroup('healer'), comp),
        namesLine(heads.melee, inGroup('melee'), comp),
        namesLine(heads.ranged, inGroup('ranged'), comp),
      ]
    : [namesLine(raidCopy.segments.in, inState('in'), comp)];
  const groups = [
    ...inLines,
    namesLine(heads.late, inState('late'), comp),
    namesLine(heads.bench, inState('bench'), comp),
    namesLine(heads.out, inState('out'), comp),
  ];
  const classes = comp.classes.filter((c) => c.count > 0).map((c) => `${c.name} ${c.count}`);
  const lines = [
    header,
    comp.breakdown.length === 0 ? size : `${size}: ${comp.breakdown.join(', ')}`,
    ...(classes.length === 0 ? [] : [`${raidCopy.copyClasses}: ${classes.join(', ')}`]),
  ];
  return [...lines, ...groups.filter((line): line is string => line !== null)].join('\n');
}
