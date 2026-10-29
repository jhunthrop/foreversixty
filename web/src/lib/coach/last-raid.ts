// web/src/lib/coach/last-raid.ts
// The account's "last raid" line (experience spec 2026-10-28 section 3.3) as data: the copy
// module's sentence split into styled segments plus the one link it carries. Pure, so the
// component and its tests agree. The line is never hidden: no raid in 14 days is the muted
// `lastRaidNone` line (controller ruling), not `null`.
import { PREMIUM_LINKS } from '../billing/premium-links';
import { COACH_SEPARATOR, coachCopy, fill } from './copy';
import type { CoachRecent, LastRaid } from './types';

export type LastRaidRole = 'lead' | 'place' | 'plain' | 'link' | 'muted';

export interface LastRaidSegment {
  role: LastRaidRole;
  text: string;
}

export interface LastRaidView {
  segments: LastRaidSegment[];
  /** Where the link segment goes; null when the sentence carries no link. */
  href: string | null;
  /** The fix-first stem inside the sentence, so the view can clamp it to one line; null when absent. */
  stem: string | null;
}

const SEGMENT_SEPARATOR = COACH_SEPARATOR;

/** The line's accessible name: the copy's own lead, so it is never typed twice. */
export const LAST_RAID_LABEL = coachCopy.lastRaidNone.split(SEGMENT_SEPARATOR)[0]!;

/** The inline Try again control: a 44 px target whose negative block margin (44 - 2 x 12 = 20) keeps it one 20 px line tall, so the reserve holds. */
export const LAST_RAID_RETRY_CLASS =
  '-my-3 inline-flex min-h-11 items-center px-1 text-gold underline hover:text-gold-hover focus-visible:outline-2 focus-visible:outline-gold disabled:opacity-60';

export const PREMIUM_HREF = PREMIUM_LINKS.coach;
const DATE_LOCALE = 'en-GB';

/** `ddd d MMM` in the viewer's zone (the browser's own unless `timeZone` is given). */
export function formatRaidDate(iso: string, timeZone?: string): string {
  const parts = new Intl.DateTimeFormat(DATE_LOCALE, {
    weekday: 'short',
    day: 'numeric',
    month: 'short',
    timeZone,
  }).formatToParts(new Date(iso));
  const part = (type: string): string => parts.find((p) => p.type === type)?.value ?? '';
  return `${part('weekday')} ${part('day')} ${part('month')}`;
}

/** Splits a filled sentence on its separators: the lead, the place and date, then plain text, with the last one a link when asked. */
function toSegments(sentence: string, linked: boolean): LastRaidSegment[] {
  const pieces = sentence.split(SEGMENT_SEPARATOR);
  return pieces.map((text, index) => {
    if (index === 0) return { role: 'lead', text };
    if (index === 1) return { role: 'place', text };
    return { role: linked && index === pieces.length - 1 ? 'link' : 'plain', text };
  });
}

function premiumTemplate(raid: LastRaid): string {
  const first = raid.first;
  if (raid.finding_count === 0 || first === null) return coachCopy.lastRaidClean;
  if (first.of === 1) return coachCopy.lastRaidPremiumOne;
  return first.on === first.of ? coachCopy.lastRaidPremiumAll : coachCopy.lastRaidPremium;
}

function freeTemplate(raid: LastRaid): string {
  if (raid.finding_count === 0) return coachCopy.lastRaidCleanFree;
  if (raid.first === null) return coachCopy.lastRaidFreeNoFirst;
  return raid.finding_count === 1 ? coachCopy.lastRaidFreeOne : coachCopy.lastRaidFree;
}

/** The line for a settled `/v1/me/coach/recent` answer; `premium` is `me.entitlements.log_coach`. */
export function lastRaidLine(recent: CoachRecent, premium: boolean, timeZone?: string): LastRaidView {
  const raid = recent.last_raid;
  if (raid === null)
    return { segments: [{ role: 'muted', text: coachCopy.lastRaidNone }], href: null, stem: null };
  const fields = {
    zone: raid.zone,
    date: formatRaidDate(raid.fought_at, timeZone),
    stem: raid.first?.stem ?? '',
    n: raid.first?.on ?? 0,
    m: raid.first?.of ?? 0,
    count: raid.finding_count,
  };
  if (recent.pending && raid.first === null) {
    return { segments: toSegments(fill(coachCopy.lastRaidPending, fields), false), href: null, stem: null };
  }
  const template = premium ? premiumTemplate(raid) : freeTemplate(raid);
  const href = linkTarget(raid, premium, template);
  const stem = template.includes('{stem}') ? fields.stem : null;
  return { segments: toSegments(fill(template, fields), href !== null), href, stem };
}

/** Premium always opens the review; free sends everything but the clean night to the upgrade page. */
function linkTarget(raid: LastRaid, premium: boolean, template: string): string | null {
  if (premium) return raid.url;
  return template === coachCopy.lastRaidCleanFree ? null : PREMIUM_HREF;
}

export type LastRaidState =
  | { kind: 'loading' }
  | { kind: 'retrying' }
  | { kind: 'error' }
  | { kind: 'ready'; view: LastRaidView; recent: CoachRecent; pending: boolean };

/** Fetches the recent answer and settles it into a state; a failed call is `error`, never a throw. */
export async function settleLastRaid(
  fetchRecent: () => Promise<CoachRecent>,
  premium: boolean,
  timeZone?: string,
): Promise<LastRaidState> {
  try {
    const recent = await fetchRecent();
    return {
      kind: 'ready',
      view: lastRaidLine(recent, premium, timeZone),
      recent,
      pending: recent.pending && recent.last_raid !== null,
    };
  } catch {
    return { kind: 'error' };
  }
}

export interface LastRaidLoaderOptions {
  fetchRecent: () => Promise<CoachRecent>;
  /** Read at each settle, so a caller may pass a getter that follows live state. */
  premium: boolean;
  /** One refresh this long after a first answer that is still pending. */
  refreshDelayMs: number;
  onState: (state: LastRaidState) => void;
  schedule?: (run: () => void, delayMs: number) => unknown;
  cancel?: (handle: unknown) => void;
}

export interface LastRaidLoader {
  start: () => void;
  /** Try again after an error: shows `retrying` while the call is in flight. */
  retry: () => void;
  /** Cancels the pending refresh and ignores any answer still in flight. */
  stop: () => void;
}

// The last ready answer on this page. The account page re-creates the line whenever /v1/me
// refreshes (a Discord save, a plan change), and a remount must not flash the skeleton over a
// line the player has already read: it starts from this answer and refreshes it silently.
let lastReady: Extract<LastRaidState, { kind: 'ready' }> | null = null;

/** The ready answer a remounted line starts from, if this page has one. */
export function cachedLastRaid(): LastRaidState | null {
  return lastReady;
}

/** Forgets the cached answer (tests, and a sign-out). */
export function resetLastRaidCache(): void {
  lastReady = null;
}

/** Loads the line once, refreshes once if the first answer is pending, and offers a retry. */
export function createLastRaidLoader(options: LastRaidLoaderOptions): LastRaidLoader {
  const { fetchRecent } = options;
  const schedule = options.schedule ?? ((run, ms) => setTimeout(run, ms));
  const cancel = options.cancel ?? ((handle) => clearTimeout(handle as ReturnType<typeof setTimeout>));
  let stopped = false;
  let timer: unknown;

  const onState = (state: LastRaidState): void => {
    if (state.kind === 'ready') lastReady = state;
    options.onState(state);
  };

  const refresh = async (): Promise<void> => {
    const next = await settleLastRaid(fetchRecent, options.premium);
    // A failed refresh keeps the line already on screen.
    if (!stopped && next.kind === 'ready') onState(next);
  };

  const load = async (before: LastRaidState, armRefresh: boolean): Promise<void> => {
    onState(before);
    const settled = await settleLastRaid(fetchRecent, options.premium);
    if (stopped) return;
    onState(settled);
    if (armRefresh && settled.kind === 'ready' && settled.pending) {
      timer = schedule(() => void refresh(), options.refreshDelayMs);
    }
  };

  return {
    // A remount after an answer shows that answer and refreshes it in place; only the first
    // mount on a page shows the skeleton.
    start: () => void (lastReady === null ? load({ kind: 'loading' }, true) : refresh()),
    retry: () => void load({ kind: 'retrying' }, false),
    stop: () => {
      stopped = true;
      if (timer !== undefined) cancel(timer);
    },
  };
}
