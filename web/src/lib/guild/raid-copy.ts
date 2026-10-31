// web/src/lib/guild/raid-copy.ts
// Every string of design/specs/2026-10-26-guild-raids.md 6.1, in one module: constants where the
// words are fixed, functions where the data fills them. No number is typed into prose; raid-copy.test.ts
// reads the spec table itself, so a drift in either file fails.
import type { GuildRaidRole, GuildRaidSignup, GuildRaidView, GuildSignupState, GuildViewerRole } from './api';
import {
  DEFAULT_LOCALE,
  clockText,
  dateLabel,
  offsetsDiffer,
  wallClockToInstant,
  zoneAbbreviation,
} from './raid-time';

const MINUTE_MS = 60_000;
const HOUR_MS = 60 * MINUTE_MS;
const DAY_MS = 24 * HOUR_MS;
/** A day further out than this is written as a date, not a weekday (spec 6.1 `{when}`). */
const WEEKDAY_WINDOW_DAYS = 6;
/** The middle dot between a line's parts. Svelte trims a literal's edge spaces, so pages write it
 *  as an expression (`{SEP}`). */
export const SEP = ' · ';

/** A sentence whose middle word is bold. */
export interface SpecLead {
  before: string;
  name: string;
  after: string;
}

export const raidCopy = {
  planButton: 'Plan a raid',
  gate: 'Planning and running raids are part of the guild plan. Subscribe the guild.',
  gateLink: 'Subscribe the guild',
  cap: 'A guild plans at most 12 raids ahead.',
  subheads: {
    planned: 'Planned',
    raidNights: 'Raid nights',
    yourSignup: 'Your sign-up',
    signups: 'Sign-ups',
    notes: 'Notes',
    tanks: 'Tanks',
    healers: 'Healers',
    dps: 'DPS',
    late: 'Late',
    bench: 'Bench',
    out: 'Out',
    notSigned: 'Not signed up',
  },
  segments: { in: 'In', late: 'Late', bench: 'Bench', out: 'Out' },
  roleLabels: { tank: 'Tank', healer: 'Healer', dps: 'DPS' } satisfies Record<GuildRaidRole, string>,
  emptyPlanned: {
    member: 'No raid planned yet. Officers plan raids here.',
    officer: 'No raid planned yet. Plan one and the guild can sign up.',
  },
  bandNone: {
    member: 'No raid planned yet.',
    officer: 'No raid planned yet.',
    planLink: 'Plan a raid',
  },
  placeholders: { note: 'Note for the officers (optional)', chooseRaid: 'Choose a raid' },
  stripStatus: {
    saved: 'Saved.',
    failed: 'That did not save. Try again.',
    movedToBench: (name: string): string => `${name} moved to Bench.`,
    left: (n: number): string => `${n} left`,
  },
  /** The raid page's spec ask: the Discord spec card's keys (guided-cards spec P.7). */
  spec: {
    title: 'Which spec tonight?',
    placeholder: 'Choose a spec',
    change: 'Change spec',
    /** `spec.lead`: the name is the bold part. */
    lead: (name: string, state: string, title: string, date: string): SpecLead => ({
      before: 'The pick signs ',
      name,
      after: ` up as ${state} for ${title}, ${date}, and is kept for the next raid.`,
    }),
    /** `spec.lead.change`. */
    leadChange: (name: string, title: string, date: string): SpecLead => ({
      before: 'The pick changes the spec ',
      name,
      after: ` brings to ${title}, ${date}, and is kept for the next raid.`,
    }),
    /** `spec.kept`, alone (without its leading space). */
    kept: (spec: string, name: string): string => `${spec} is kept for ${name}.`,
  },
  waitingCharacter: 'Waiting for approval. You can sign it up once an officer approves it.',
  raidStates: {
    locked: 'Sign-ups are locked. An officer can still change yours.',
    cancelled: 'This raid was cancelled.',
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
    cancel: 'Cancel raid',
    failed: 'That did not save; try again.',
  },
  cancelConfirm: {
    question: 'Cancel this raid? Members see it as cancelled.',
    postToDiscord: 'Post the cancellation to Discord',
    confirm: 'Cancel the raid',
    confirming: 'Cancelling…',
    keep: 'Keep it',
  },
  copyFailed: 'Could not copy. Select the list and copy it.',
  form: {
    raid: 'Raid',
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
    postLater: (channel: string | null, day: string): string =>
      `Posts to ${channel ?? 'your Discord channel'} on ${day}, a week before the raid.`,
    postRepeat: (firstWithinWeek: boolean): string =>
      firstWithinWeek
        ? 'Each raid posts a week before it starts; the first posts as soon as you save.'
        : 'Each raid posts a week before it starts.',
    postOff: 'This raid will not post to Discord.',
    remind: 'Posts a reminder 24 hours before the start.',
    remindNeedsPost: 'Turn on Post to Discord to send a reminder.',
    posted: (channel: string): string => `Posted to ${channel}. The post updates itself.`,
    reminded: 'Reminder sent.',
    submitCreate: 'Post the raid',
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
      raid: 'Choose a raid.',
      date: 'Pick a date.',
      time: 'Pick a time after now.',
      cap: 'A guild plans at most 12 raids ahead.',
      failed: 'That did not save. Try again.',
    },
  },
  notesEmpty: {
    member: 'No notes for this raid.',
    officer: 'No notes yet. Edit the raid to add some.',
  },
  emptyRole: { tank: 'No tanks yet.', healer: 'No healers yet.', dps: 'No DPS yet.' } satisfies Record<
    GuildRaidRole,
    string
  >,
  notVisible: {
    members: (guild: string): string => `Planned raids are for members of ${guild}.`,
    unverified: 'You are waiting for approval. Planned raids show once an officer approves you.',
  },
  loadErrors: { raid: 'This raid did not load.', planned: 'Planned raids did not load.' },
  backLink: (guild: string): string => `‹ ${guild} · Raids`,
  /** The raid page's own words outside the 6.1 table (spec 3.2, 3.7, 4 "Page states"). */
  page: {
    eyebrow: 'Raid',
    /** The guild band's own eyebrow (guild page spec 4.A), on the not-visible and failed pages. */
    guildEyebrow: 'Guild',
    loading: 'Loading the raid',
    /** The officer state select's accessible name (3.7). */
    signupFor: (name: string): string => `Sign-up for ${name}`,
    /** The not-signed row's resting option; its accessible name is `Not signed up`. */
    notSignedMark: '—',
    /** Between the start and end clocks: `8:00 pm to 11:00 pm`. */
    timeTo: 'to',
  },
  changelog:
    "Added raid sign-ups. Officers plan raids on the guild's Raids tab and members sign up in one tap, from the raid page or the guild page.",
} as const;

const SIGNED_UP_STATE_WORDS: Record<GuildSignupState, string> = {
  in: raidCopy.segments.in,
  late: raidCopy.segments.late,
  bench: raidCopy.segments.bench,
  out: raidCopy.segments.out,
};

function plural(n: number, singular: string, pluralForm: string): string {
  return n === 1 ? singular : pluralForm;
}

/** `23 of 40 signed up`; `23 signed up` while the size is unknown (ruling R6). */
export function countLine(n: number, size: number | null): string {
  return size === null ? `${n} signed up` : `${n} of ${size} signed up`;
}

/** `23/40`; the bare count while the size is unknown. */
export function countRatio(n: number, size: number | null): string {
  return size === null ? `${n}` : `${n}/${size}`;
}

/** `3 tanks · 7 healers · 13 DPS · 1 late`: the roles count In rows, so the late clause (only when
 *  someone is late) makes the parts add up to the signed-up count. */
export function roleBreakdown(tanks: number, healers: number, dps: number, late: number): string {
  return roleBreakdownParts(tanks, healers, dps, late).join(SEP);
}

/** The breakdown's clauses, for a line that joins them its own way (the copied list's commas). */
export function roleBreakdownParts(tanks: number, healers: number, dps: number, late: number): string[] {
  const parts = [
    `${tanks} ${plural(tanks, 'tank', 'tanks')}`,
    `${healers} ${plural(healers, 'healer', 'healers')}`,
    `${dps} DPS`,
  ];
  return late > 0 ? [...parts, `${late} late`] : parts;
}

/** `23 of 40 are in`; `23 are in` while the size is unknown: the unsigned band line's count. */
function inLine(n: number, size: number | null): string {
  return size === null ? `${n} are in` : `${n} of ${size} are in`;
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
export function nextRaidDeadline(
  startsAt: string | undefined,
  locale: string = DEFAULT_LOCALE,
  viewerTz: string = browserZone(),
): string {
  if (startsAt === undefined) return needsBefore(null);
  return needsBefore(
    new Intl.DateTimeFormat(locale, { timeZone: viewerTz, weekday: 'long' }).format(new Date(startsAt)),
  );
}

/** `8:00 pm CST guild time`. */
export function guildTimeLabel(startsAt: Date, tz: string, locale: string = DEFAULT_LOCALE): string {
  return `${guildClock(startsAt, tz, locale)} guild time`;
}

/** `8:00 pm CST`: the guild clock with its zone abbreviation. */
function guildClock(startsAt: Date, tz: string, locale: string): string {
  return `${clockText(startsAt, tz, locale, true)} ${zoneAbbreviation(tz, startsAt)}`;
}

/** The form's line under the zone select: the entered guild wall time as the viewer reads it. */
export function forYouLine(
  date: string,
  time: string,
  guildTz: string,
  viewerTz: string,
  locale: string = DEFAULT_LOCALE,
): string {
  const instant = wallClockToInstant(date, time, guildTz);
  if (!offsetsDiffer(instant, guildTz, viewerTz)) return raidCopy.form.forYouSame;
  return raidCopy.form.forYou(
    `${dateLabel(instant, viewerTz, locale)}, ${clockText(instant, viewerTz, locale, true)}`,
  );
}

function calendarDayDiff(from: Date, to: Date, tz: string): number {
  const asDay = (d: Date): number =>
    Date.parse(`${new Intl.DateTimeFormat('en-CA', { timeZone: tz }).format(d)}T00:00:00Z`);
  return Math.round((asDay(to) - asDay(from)) / DAY_MS);
}

function browserZone(): string {
  return new Intl.DateTimeFormat().resolvedOptions().timeZone;
}

/** The spec's `{when}`: `Tonight 8 pm`, `Thursday 8 pm`, `Thu 17 Dec, 8 pm`, then the guild's own
 *  clock when its offset differs from the viewer's. */
export function whenPhrase(
  startsAt: Date,
  now: Date,
  locale: string,
  guildTz: string,
  viewerTz: string = browserZone(),
): string {
  const time = clockText(startsAt, viewerTz, locale);
  const days = calendarDayDiff(now, startsAt, viewerTz);
  let phrase: string;
  if (days === 0) phrase = `Tonight ${time}`;
  else if (days === 1) phrase = `Tomorrow ${time}`;
  else if (days > 1 && days <= WEEKDAY_WINDOW_DAYS) {
    const weekday = new Intl.DateTimeFormat(locale, { timeZone: viewerTz, weekday: 'long' }).format(startsAt);
    phrase = `${weekday} ${time}`;
  } else phrase = `${dateLabel(startsAt, viewerTz, locale)}, ${time}`;
  if (!offsetsDiffer(startsAt, viewerTz, guildTz)) return phrase;
  return `${phrase} (${clockText(startsAt, guildTz, locale)} guild time)`;
}

/** `2 days`, `1 day 4 h`, `5 h 20 min`, `12 min`. */
function span(ms: number): string {
  const days = Math.floor(ms / DAY_MS);
  const hours = Math.floor((ms % DAY_MS) / HOUR_MS);
  const minutes = Math.floor((ms % HOUR_MS) / MINUTE_MS);
  if (days > 0) return `${days} ${plural(days, 'day', 'days')}${hours > 0 ? ` ${hours} h` : ''}`;
  if (hours > 0) return `${hours} h${minutes > 0 ? ` ${minutes} min` : ''}`;
  return `${minutes} min`;
}

export function countdown(startsAt: Date, endsAt: Date, now: Date): string {
  const untilStart = startsAt.getTime() - now.getTime();
  if (Math.abs(untilStart) < MINUTE_MS) return 'starting now';
  if (untilStart > 0) return `in ${span(untilStart)}`;
  if (now.getTime() < endsAt.getTime()) return `started ${span(-untilStart)} ago`;
  return `ended ${span(now.getTime() - endsAt.getTime())} ago`;
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

/** The clause around its state word: `you are ` / `In` / ` as DPS`; no word while unsigned. */
function bandClause({ state, role, characterName }: BandLineArgs): BandLineParts {
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
    args.state === null ? `. ${inLine(args.n, args.size)}.` : `, ${countLine(args.n, args.size)}.`;
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

function signupsIn(signups: readonly GuildRaidSignup[], state: GuildSignupState, role?: GuildRaidRole) {
  return signups.filter((s) => s.state === state && (role === undefined || s.role === role));
}

function namesLine(label: string, rows: readonly GuildRaidSignup[]): string | null {
  return rows.length === 0 ? null : `${label}: ${rows.map((row) => row.name).join(', ')}`;
}

function breakdownList(counts: GuildRaidView['counts']): string {
  return roleBreakdownParts(counts.tanks, counts.healers, counts.dps, counts.late).join(', ');
}

/** Discord's `<t:{unix}:F>` stamp, which each reader's client draws in their own zone. */
function discordStamp(at: Date): string {
  return `<t:${Math.floor(at.getTime() / 1000)}:F>`;
}

/** Spec 4 "Copy the sign-up list": plain text in guild time, groups in roster order, empty groups left out. */
export function copyList(view: GuildRaidView, tz: string, locale: string = DEFAULT_LOCALE): string {
  const startsAt = new Date(view.starts_at);
  const heads = raidCopy.subheads;
  const { counts } = view;
  const header = [
    view.title,
    `${dateLabel(startsAt, tz, locale)}, ${guildClock(startsAt, tz, locale)}`,
    discordStamp(startsAt),
  ].join(SEP);
  const size =
    view.size === null
      ? `${counts.in + counts.late} signed up`
      : `${counts.in + counts.late} of ${view.size}`;
  const groups = [
    namesLine(heads.tanks, signupsIn(view.signups, 'in', 'tank')),
    namesLine(heads.healers, signupsIn(view.signups, 'in', 'healer')),
    namesLine(heads.dps, signupsIn(view.signups, 'in', 'dps')),
    namesLine(heads.late, signupsIn(view.signups, 'late')),
    namesLine(heads.bench, signupsIn(view.signups, 'bench')),
    namesLine(heads.out, signupsIn(view.signups, 'out')),
  ];
  const lines = [header, `${size}: ${breakdownList(counts)}`];
  return [...lines, ...groups.filter((line): line is string => line !== null)].join('\n');
}
