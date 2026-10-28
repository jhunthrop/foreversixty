// web/src/lib/coach/last-raid.ts
// The account's "last raid" line (experience spec 2026-10-28 section 3.3) as data: the copy
// module's sentence split into styled segments plus the one link it carries. Pure, so the
// component and its tests agree. The line is never hidden: no raid in 14 days is the muted
// `lastRaidNone` line (controller ruling), not `null`.
import { coachCopy, fill } from './copy';
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
}

export const PREMIUM_HREF = '/premium';
const SEGMENT_SEPARATOR = ' · ';
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
  if (raid === null) return { segments: [{ role: 'muted', text: coachCopy.lastRaidNone }], href: null };
  const fields = {
    zone: raid.zone,
    date: formatRaidDate(raid.fought_at, timeZone),
    stem: raid.first?.stem ?? '',
    n: raid.first?.on ?? 0,
    m: raid.first?.of ?? 0,
    count: raid.finding_count,
  };
  if (recent.pending && raid.first === null) {
    return { segments: toSegments(fill(coachCopy.lastRaidPending, fields), false), href: null };
  }
  const template = premium ? premiumTemplate(raid) : freeTemplate(raid);
  const href = !premium ? (template.endsWith('Premium') ? PREMIUM_HREF : null) : raid.url;
  const linked = href !== null && template !== coachCopy.lastRaidCleanFree;
  return { segments: toSegments(fill(template, fields), linked), href: linked ? href : null };
}

export type LastRaidState =
  { kind: 'loading' } | { kind: 'error' } | { kind: 'ready'; view: LastRaidView; pending: boolean };

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
      pending: recent.pending && recent.last_raid !== null,
    };
  } catch {
    return { kind: 'error' };
  }
}
