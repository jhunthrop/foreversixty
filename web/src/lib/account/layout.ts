// web/src/lib/account/layout.ts
// Reserved-height constants for /account's loading state (experience spec
// 2026-10-25-discord-notifications section 9.1): the loading page draws the ready page's
// layout, so these tables are the single source for the floors that global.css's
// "account loading reserve" block repeats as CSS (layout-reserve.test.ts asserts the two agree).
// The page on panels (redesign 2026-11-09, directions mock b4-account) keeps the mechanism: `left`
// is the main column (Characters, Guild, Devices, Discord) and `rail` the rail (You, Plan, Delete
// account); a stored height is what reserves them, and the floors below are the shipped page's,
// met only on a first visit. `npm run lhci`'s own account.html run cannot confirm any of it: its
// sandbox has no CORS allowance for localhost to call the real API, so it always lands on the
// `failed`/LoadError branch instead.

import type { Me } from './api';
import { storage } from '../layout-reserve';

/** The gap between the page's panels (ui/Page `panels`: `--section-gap` 22 px). */
export const SECTION_GAP_PX = 22;

// ---- The loading reserve (experience spec 2026-10-25-discord-notifications §9.1) ----------
// The loading state draws the ready page's skeleton at its real positions. Two buckets:
// `lg` is the split page at 1024 and up (measured at 1440), `stack` is the one-column page
// below it (measured at 390). Each floor is a panel of the ready page measured on a
// FOREVER_DATA=fixture build (2026-11-09) and rounded down to 5, so a floor never out-grows the
// real panel. All apply only when no rendered height is stored.

export type ReserveBucket = 'lg' | 'stack';

/** One number per bucket: the 1440 desktop layout and the 390 phone layout. */
export interface BucketPx {
  lg: number;
  stack: number;
}

/** What `data-discord` on `<html>` says the Discord block will be. */
export type HintState = 'none' | 'free' | 'connect' | 'form' | 'saved' | 'off' | 'lapsed';

/** The main column's panels with nothing extra in them: an empty Characters panel and Devices with no device. */
export const LEFT_FLOORS: Readonly<Record<'characters' | 'devices', BucketPx>> = {
  characters: { lg: 225, stack: 270 },
  devices: { lg: 145, stack: 215 },
};

/** The rail's panels: You (no Discord rows: they are the Discord panel's), the free Plan, Delete account. */
export const RAIL_FLOORS: Readonly<Record<'you' | 'plan' | 'delete', BucketPx>> = {
  you: { lg: 350, stack: 350 },
  plan: { lg: 295, stack: 340 },
  delete: { lg: 180, stack: 160 },
};

/** The Discord panel's notifications block by the state `/v1/me` puts it in; `none` draws no panel. */
export const DISCORD_FLOORS: Readonly<Record<HintState, BucketPx>> = {
  none: { lg: 0, stack: 0 },
  free: { lg: 140, stack: 155 },
  connect: { lg: 295, stack: 355 },
  form: { lg: 535, stack: 660 },
  saved: { lg: 345, stack: 440 },
  off: { lg: 420, stack: 535 },
  lapsed: { lg: 385, stack: 410 },
};

/** What `data-discord-link` on `<html>` says the Discord panel's link rows will be. */
export type LinkHintState = 'none' | 'unlinked' | 'suggested' | 'linked';

/** The link row and the Premium role row under it (both show whenever the link row does), by the link row's state. */
export const ACCOUNT_DISCORD_LINK_RESERVE: Readonly<Record<LinkHintState, BucketPx>> = {
  none: { lg: 0, stack: 0 },
  unlinked: { lg: 120, stack: 200 },
  suggested: { lg: 120, stack: 200 },
  linked: { lg: 110, stack: 180 },
};

/** Read back by the pre-paint script in Base.astro. */
export const ACCOUNT_RESERVE_KEY = 'fs.accountReserve';

/** A stored height above this is nonsense (tampered or from another layout) and is ignored. */
export const ACCOUNT_RESERVE_MAX_PX = 4000;

const RESERVE_VERSION = 1;

export interface BucketHeights {
  left: number;
  rail: number;
}

export interface AccountReserve {
  lg: BucketHeights;
  stack: BucketHeights;
}

function validHeight(value: unknown): value is number {
  return typeof value === 'number' && Number.isFinite(value) && value > 0 && value <= ACCOUNT_RESERVE_MAX_PX;
}

function validBucket(value: unknown): BucketHeights | null {
  if (typeof value !== 'object' || value === null) return null;
  const { left, rail } = value as Partial<BucketHeights>;
  return validHeight(left) && validHeight(rail) ? { left, rail } : null;
}

/** Whatever valid buckets storage holds (a first visit has only the one it was on). */
function readStoredBuckets(): Partial<AccountReserve> {
  try {
    const raw = storage()?.getItem(ACCOUNT_RESERVE_KEY);
    if (raw === null || raw === undefined) return {};
    const parsed = JSON.parse(raw) as { v?: unknown; lg?: unknown; stack?: unknown };
    if (parsed.v !== RESERVE_VERSION) return {};
    const lg = validBucket(parsed.lg);
    const stack = validBucket(parsed.stack);
    return { ...(lg === null ? {} : { lg }), ...(stack === null ? {} : { stack }) };
  } catch {
    return {};
  }
}

/** The stored heights of the last ready pages (the test reader: Base.astro holds its own copy); null unless both buckets are present and valid. */
export function readAccountReserve(): AccountReserve | null {
  const { lg, stack } = readStoredBuckets();
  return lg === undefined || stack === undefined ? null : { lg, stack };
}

/** Stores the rendered column heights for one bucket, keeping the other; a 0 or oversized value skips the write. */
export function writeAccountReserve(bucket: ReserveBucket, left: number, rail: number): void {
  if (!validHeight(left) || !validHeight(rail)) return;
  const kept = readStoredBuckets();
  const next = { ...kept, v: RESERVE_VERSION, [bucket]: { left: Math.round(left), rail: Math.round(rail) } };
  try {
    storage()?.setItem(ACCOUNT_RESERVE_KEY, JSON.stringify(next));
  } catch {
    // Private mode or quota: the hint is a convenience.
  }
}

/** Sign-out forgets the heights: the next account may have a different page. */
export function clearAccountReserve(): void {
  try {
    storage()?.removeItem(ACCOUNT_RESERVE_KEY);
  } catch {
    // Best-effort.
  }
}

/**
 * Which Discord block the page will draw, from the cached `/v1/me` (spec 4.1 precedence; `form` with connect on is `connect`). The
 * pre-paint script in Base.astro copies this logic verbatim: keep it small and import-free.
 */
export function discordHintState(me: Me | null): HintState {
  if (me === null || me.notifications === undefined) return 'none';
  const discord = me.notifications.discord;
  const entitled = me.entitlements && me.entitlements.notifications === true;
  if (discord === null || discord === undefined) {
    if (!entitled) return 'free';
    return me.notifications.connect === true ? 'connect' : 'form';
  }
  if (!entitled) return 'lapsed';
  return discord.disabled_reason === null ? 'saved' : 'off';
}

/**
 * Which link row the You panel will draw, from the cached `/v1/me`. The pre-paint script in
 * Base.astro copies this logic verbatim: keep it small and import-free.
 */
export function discordLinkHint(me: Me | null): LinkHintState {
  if (me === null || me.discord_link_available !== true) return 'none';
  if (me.discord_link && me.discord_link.display_name) return 'linked';
  return me.discord_link_suggested ? 'suggested' : 'unlinked';
}

/** The main column's floor with nothing stored: Characters and Devices, and the Discord panel when there is one. */
export function leftFloor(state: HintState, bucket: ReserveBucket): number {
  const panels = LEFT_FLOORS.characters[bucket] + SECTION_GAP_PX + LEFT_FLOORS.devices[bucket];
  const discord = DISCORD_FLOORS[state][bucket];
  return panels + (discord > 0 ? SECTION_GAP_PX + discord : 0);
}

/** The rail's floor with nothing stored: its three panels and the two gaps. */
export function railFloor(bucket: ReserveBucket): number {
  return RAIL_FLOORS.you[bucket] + RAIL_FLOORS.plan[bucket] + RAIL_FLOORS.delete[bucket] + 2 * SECTION_GAP_PX;
}

/** The reserved split's height under the head when nothing is stored (the taller column; stacked below 1024). */
export function gridFloor(state: HintState, bucket: ReserveBucket): number {
  const left = leftFloor(state, bucket);
  const rail = railFloor(bucket);
  return bucket === 'lg' ? Math.max(left, rail) : left + SECTION_GAP_PX + rail;
}

const LG_MEDIA = '(min-width: 1024px)';

/** The bucket this viewport's page lays out in: the 12-column grid at 1024 and up, else one column. */
export function currentBucket(): ReserveBucket {
  return typeof window !== 'undefined' && window.matchMedia(LG_MEDIA).matches ? 'lg' : 'stack';
}

const TRANSIENT_SELECTOR = '[data-testid="discord-cancel"], [data-testid="discord-remove-confirm"]';

/** True while the Discord block shows a passing view (replacing a webhook, or the remove question). */
export function isTransientAccountState(doc: Document): boolean {
  return doc.querySelector(TRANSIENT_SELECTOR) !== null;
}

/** Stores the two columns' rendered heights for the current bucket (a missing column or a passing view writes nothing). */
export function recordAccountReserve(left: Element | null, rail: Element | null): void {
  if (left === null || rail === null || isTransientAccountState(document)) return;
  writeAccountReserve(
    currentBucket(),
    left.getBoundingClientRect().height,
    rail.getBoundingClientRect().height,
  );
}

/**
 * Runs `record` two animation frames from now (the ready view has laid out) and again on
 * `pagehide` (the Your reports list that arrives later is counted). Returns the cleanup, which
 * also stops a later `pagehide` write -- sign-out needs that, or it would store the heights it
 * just cleared.
 */
export function scheduleAccountReserveWrites(record: () => void): () => void {
  let outer = 0;
  let inner = 0;
  outer = requestAnimationFrame(() => {
    inner = requestAnimationFrame(record);
  });
  window.addEventListener('pagehide', record);
  return () => {
    cancelAnimationFrame(outer);
    cancelAnimationFrame(inner);
    window.removeEventListener('pagehide', record);
  };
}
