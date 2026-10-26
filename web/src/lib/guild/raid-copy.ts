// web/src/lib/guild/raid-copy.ts
// Every string of design/specs/2026-10-26-guild-raids.md 6.1, in one module: constants where the
// words are fixed, functions where the data fills them. No number is typed into prose; raid-copy.test.ts
// reads the spec table itself, so a drift in either file fails.
import type { GuildRaidRole, GuildRaidSignup, GuildRaidView, GuildSignupState } from './api';
import { clockText, dateLabel, offsetsDiffer } from './raid-time';

const DEFAULT_LOCALE = 'en-US';
const MINUTE_MS = 60_000;
const HOUR_MS = 60 * MINUTE_MS;
const DAY_MS = 24 * HOUR_MS;
/** A day further out than this is written as a date, not a weekday (spec 6.1 `{when}`). */
const WEEKDAY_WINDOW_DAYS = 6;

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
  waitingCharacter: 'Waiting for approval. You can sign it up once an officer approves it.',
  raidStates: {
    locked: 'Sign-ups are locked. An officer can still change yours.',
    ended: 'This raid has ended.',
    cancelled: 'This raid was cancelled.',
    done: 'Logged as a raid night.',
    openNight: 'Open the night',
  },
  races: {
    locked: 'This raid was locked while you were signing up.',
    cancelled: 'This raid was cancelled while you were signing up.',
    officerToolsEnded: "Not saved. This guild's officer tools have ended.",
  },
  poll: 'Live updates paused. Reload to see the latest.',
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
    submitCreate: 'Post the raid',
    submitCreating: 'Posting…',
    submitEdit: 'Save changes',
    submitEditing: 'Saving…',
    cancel: 'Cancel',
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

export function roleBreakdown(tanks: number, healers: number, dps: number): string {
  return [
    `${tanks} ${plural(tanks, 'tank', 'tanks')}`,
    `${healers} ${plural(healers, 'healer', 'healers')}`,
    `${dps} DPS`,
  ].join(' · ');
}

export function showAll(n: number): string {
  return `Show all ${n} planned raids`;
}

/** The needs line's deadline: `before Thursday`, or `before the next raid` with none planned. */
export function needsBefore(weekday: string | null): string {
  return weekday === null ? 'before the next raid' : `before ${weekday}`;
}

/** `8:00 pm CST guild time`. */
export function guildTimeLabel(startsAt: Date, tz: string, locale: string = DEFAULT_LOCALE): string {
  const zone = new Intl.DateTimeFormat(DEFAULT_LOCALE, { timeZone: tz, timeZoneName: 'short' })
    .formatToParts(startsAt)
    .find((part) => part.type === 'timeZoneName')?.value;
  return `${clockText(startsAt, tz, locale, true)} ${zone ?? tz} guild time`;
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

function bandClause({ state, role, characterName }: BandLineArgs): string {
  if (state === null) return 'you have not signed up';
  if (state === 'bench') return 'you are on the Bench';
  if (state === 'out') return 'you are Out';
  const on = characterName === undefined ? '' : ` on ${characterName}`;
  const as = role === undefined ? '' : ` as ${raidCopy.roleLabels[role]}`;
  return `you are ${SIGNED_UP_STATE_WORDS[state]}${on}${as}`;
}

/** The guild band's raid sentence: `Tonight 8 pm, you are In as DPS, 23 of 40 signed up.` */
export function bandLine(args: BandLineArgs): string {
  return `${args.when}, ${bandClause(args)}, ${countLine(args.n, args.size)}.`;
}

function signupsIn(signups: readonly GuildRaidSignup[], state: GuildSignupState, role?: GuildRaidRole) {
  return signups.filter((s) => s.state === state && (role === undefined || s.role === role));
}

/** `8:00 pm CST`: the guild clock with its zone abbreviation. */
function guildClock(startsAt: Date, tz: string, locale: string): string {
  const zone = new Intl.DateTimeFormat(DEFAULT_LOCALE, { timeZone: tz, timeZoneName: 'short' })
    .formatToParts(startsAt)
    .find((part) => part.type === 'timeZoneName')?.value;
  return `${clockText(startsAt, tz, locale, true)} ${zone ?? tz}`;
}

function namesLine(label: string, rows: readonly GuildRaidSignup[]): string | null {
  return rows.length === 0 ? null : `${label}: ${rows.map((row) => row.name).join(', ')}`;
}

function breakdownList(tanks: number, healers: number, dps: number): string {
  return roleBreakdown(tanks, healers, dps).replaceAll(' · ', ', ');
}

/** Spec 4 "Copy the sign-up list": plain text in guild time, groups in roster order, empty groups left out. */
export function copyList(view: GuildRaidView, tz: string, locale: string = DEFAULT_LOCALE): string {
  const startsAt = new Date(view.starts_at);
  const heads = raidCopy.subheads;
  const { counts } = view;
  const header = `${view.title} · ${dateLabel(startsAt, tz, locale)}, ${guildClock(startsAt, tz, locale)}`;
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
  const lines = [header, `${size}: ${breakdownList(counts.tanks, counts.healers, counts.dps)}`];
  return [...lines, ...groups.filter((line): line is string => line !== null)].join('\n');
}
