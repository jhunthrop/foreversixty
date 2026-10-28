// web/src/lib/coach/copy.ts
// Every string the coach shows on the web, word for word from the experience spec's section 8.2
// (design/specs/2026-10-28-log-coach.md). copy.test.ts reads the spec table and fails on any drift.
// A `{field}` is filled by `fill`.

/**
 * The muted ` · ` between a line's segments, rendered as an expression: Svelte trims the edges of
 * a tag's literal text, which would drop the space before the dot.
 */
export const COACH_SEPARATOR = ' · ';

export const coachCopy = {
  rowPremium: 'Fix first · {title}',
  rowFree: '{n} things to fix · First: {title} · Read the full review with Premium',
  rowFreeOne: '1 thing to fix · First: {title} · Read the full review with Premium',
  rowFreeNoFirst: '{n} things to fix · Read the full review with Premium',
  rowFreeNoFirstOne: '1 thing to fix · Read the full review with Premium',
  rowPremiumNoFirst: "{n} things to fix · Open your coach's review",
  rowPremiumNoFirstOne: "1 thing to fix · Open your coach's review",
  rowClean: 'Nothing to fix on this pull that the log can show.',
  rowOfficer: 'Coach review',
  rowPreparing: 'Top {Spec} logs for this boss are still loading.',
  rowError: "The coach couldn't read this pull just now.",
  rowRetry: 'Try again',
  loading: 'Reading the pull…',
  writing: 'Your coach is writing…',
  writingLate: 'Still writing. Check back in a few minutes.',
  ready: "Your coach's review is ready.",
  busy: "Your coach is busy right now. This is the review from your log alone; the coach's version will be here when you next open it.",
  error: "The coach couldn't read this pull just now.",
  titleModel: 'Your {Spec} coach',
  titleRules: 'Your {Spec} review',
  titleOfficer: "{Name}'s {Spec} review",
  titleStarter: 'Your {Spec} coach · early version',
  titleGeneric: 'Your coach',
  starterHint: "This coach reads your log against our guide's {Spec} rotation.",
  officerNote: 'Written to {Name}. You see it as an officer of the guild.',
  officerPill: 'Officer view',
  officerPick: "Open a player's row and choose Coach review.",
  labelFixFirst: 'Fix first',
  labelPractice: 'Practice',
  labelThen: 'Then',
  labelWell: 'Doing well',
  labelCannot: "Couldn't check",
  footer: 'Reviewed {date} · Based on: {labels}.',
  history: 'History',
  historyOlder: 'Show older',
  historyNone: 'Nothing to fix',
  historyRules: 'Log only',
  rulesNothing: 'Nothing to fix on this pull that the log can show.',
  notReviewable:
    'This pull is too short to review. The coach reads boss pulls long enough to show a rotation.',
  preparing: 'Top {Spec} logs for this boss are still loading. Check again in a few minutes.',
  checkAgain: 'Check again',
  notAvailable: "This review isn't available on this account.",
  signedOut: "Sign in to read your coach's review.",
  signIn: 'Sign in',
  night: 'Coach reviews cover one pull at a time. Pick a boss pull from the list.',
  teaserCount: '{n} things to fix',
  teaserCountOne: '1 thing to fix',
  teaserFirst: '{title}',
  upgrade: 'Read the full review with Premium',
  anchorCast: '{Spell} · Casts',
  anchorBuffs: '{Spell} · Buffs',
  anchorDebuffs: '{Spell} · Debuffs',
  anchorBuffsPull: 'Buffs · at the pull',
  anchorHealing: '{Spell} · Healing',
  anchorTime: 'Timelines · {t}',
  anchorResource: 'Resources · {t}',
  anchorDeath: 'Death · {t}',
  lastRaidPremium: 'Last raid · {zone}, {date} · Fix first: {stem}, on {n} of {m} bosses · Open review',
  lastRaidPremiumOne: 'Last raid · {zone}, {date} · Fix first: {stem} · Open review',
  lastRaidPremiumAll: 'Last raid · {zone}, {date} · Fix first: {stem}, on all {m} bosses · Open review',
  lastRaidFree:
    'Last raid · {zone}, {date} · {count} things to fix · First: {stem} · Read the full review with Premium',
  lastRaidFreeOne:
    'Last raid · {zone}, {date} · 1 thing to fix · First: {stem} · Read the full review with Premium',
  lastRaidFreeNoFirst:
    'Last raid · {zone}, {date} · {count} things to fix · Read the full review with Premium',
  lastRaidClean: 'Last raid · {zone}, {date} · Nothing to fix that the log can show · Open review',
  lastRaidCleanFree: 'Last raid · {zone}, {date} · Nothing to fix that the log can show',
  lastRaidPending: 'Last raid · {zone}, {date} · Reading your pulls…',
  lastRaidNone: 'Last raid · no log in the last 14 days',
  lastRaidError: "Your last raid isn't available right now.",
} as const;

export type CoachCopyKey = keyof typeof coachCopy;

/** Replaces every `{field}` in a template; a field the template lacks is ignored, one it needs and lacks is left visible. */
export function fill(template: string, fields: Readonly<Record<string, string | number>>): string {
  return template.replace(/\{(\w+)\}/g, (slot, name: string) =>
    name in fields ? String(fields[name]) : slot,
  );
}
