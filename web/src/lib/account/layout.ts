// web/src/lib/account/layout.ts
// Reserved-height constants for /account's loading state (experience spec
// 2026-10-25-discord-notifications section 9.1): the loading page draws the ready page's
// layout, so these tables are the single source for the floors that global.css's
// "account loading reserve" block repeats as CSS (layout-reserve.test.ts asserts the two agree) and for
// SIGNED_OUT_MIN_H. `npm run lhci`'s own account.html run cannot confirm any of it: its sandbox
// has no CORS allowance for localhost to call the real API, so it always lands on the
// `failed`/LoadError branch instead.

import type { Me } from './api';
import { storage } from '../layout-reserve';

/** The gap between the page's stacked blocks (Tailwind `gap-8`). */
export const SECTION_GAP_PX = 32;

/** The gap between the rail's Plan block and the Discord block under it (`gap-4`). */
export const PLAN_GAP_PX = 16;

/** The signed-out prompt's floor. A visitor with the session hint sees the loading reserve until
 *  `/v1/me` answers, so a hinted visitor who turns out signed out only avoids a layout shift if
 *  the prompt's block holds the same height as the reserved grid under the header row (no
 *  Discord hint then: `data-discord` is `none`, so the rail is Devices + You + Plan). Both
 *  figures are `SECTION_GAP_PX` arithmetic over the tables below -- `layout-reserve.test.ts` checks
 *  them -- and Tailwind needs the literal class text. Unverifiable against `npm run lhci`'s own
 *  account.html run (its sandbox has no CORS allowance, so it always lands on LoadError). */
export const SIGNED_OUT_MIN_H = 'min-h-[1364px] lg:min-h-[829px]';

// ---- The loading reserve (experience spec 2026-10-25-discord-notifications §9.1) ----------
// The loading state draws the ready page's skeleton at its real positions. Two buckets:
// `lg` is the 12-column grid at 1024 and up (measured at 1440), `stack` is the one-column
// page below it (measured at 390). The header, left and rail floors were measured on a
// FOREVER_DATA=fixture build with a stubbed, empty account (no characters, devices, reports or
// guilds) and rounded down, so a floor never out-grows the real panel. The Discord floors were
// measured on the shipped panels (Task 3a) and rounded up to 5.
// All apply only when no rendered height is stored.

export type ReserveBucket = 'lg' | 'stack';

/** One number per bucket: the 1440 desktop layout and the 390 phone layout. */
export interface BucketPx {
  lg: number;
  stack: number;
}

/** What `data-discord` on `<html>` says the Discord block will be. */
export type HintState = 'none' | 'free' | 'connect' | 'form' | 'saved' | 'off' | 'lapsed';

/** The ready header row: 44 px beside the title at 1024 and up, the stacked row below it. */
export const HEADER_ROW_H: BucketPx = { lg: 44, stack: 83 };

/** The Characters panel's last-raid coach line (experience spec 2026-10-28-log-coach 3.3): 44 px
 *  on one line, 64 px on the phone's two, held in every state, plus the panel's 16 px (`gap-4`)
 *  flex gap above the list. Counted once, in `LEFT_FLOOR`. */
export const COACH_LAST_RAID_RESERVE: BucketPx = { lg: 44, stack: 64 };

/** The gap between a StatePanel's label row, the coach line and its body (Tailwind `gap-4`). */
const PANEL_GAP_PX = 16;

/** An empty Characters panel with the coach line, an empty Your reports panel and the 32 px gap between them
 *  (298 / 403 measured before the coach line existed). */
export const LEFT_FLOOR: BucketPx = {
  lg: 298 + COACH_LAST_RAID_RESERVE.lg + PANEL_GAP_PX,
  stack: 403 + COACH_LAST_RAID_RESERVE.stack + PANEL_GAP_PX,
};

/** The Plan panel with the free plan line only (measured on the empty fixture account). */
const PLAN_EMPTY_FLOOR: BucketPx = { lg: 219, stack: 211 };

/**
 * What the tallest Plan headline adds to the free panel (a seat holder with a personal portal):
 * the `Billing history` button row (12 px flex gap + the 44 px fixed button), line 2 under the plan
 * line (4 px gap + 20 px per 13 px line; two lines at 1440, three at 390 where the seat note
 * wraps). The past-due alert and a grant's line 2 are never taller. The panel never draws shorter
 * than the stored reserve says, so only a first load with nothing stored sees this floor.
 */
export const PLAN_HEADLINE_RESERVE: BucketPx = { lg: 12 + 44 + 4 + 2 * 20, stack: 12 + 44 + 4 + 3 * 20 };

/** One 13 px `{Guild} has no plan.` line under a guild row in the Plan panel (20 px line). The
 *  reserve holds one row; an account with more no-plan guilds grows the panel, which is not a
 *  first-load shift (the stored height takes over). */
export const PLAN_NO_PLAN_ROW_RESERVE = 20;

/** The rail's panels with nothing extra in them (Devices, You, Plan at its tallest headline). */
export const RAIL_FLOORS: Readonly<Record<'devices' | 'you' | 'plan', BucketPx>> = {
  devices: { lg: 233, stack: 249 },
  you: { lg: 193, stack: 185 },
  plan: {
    lg: PLAN_EMPTY_FLOOR.lg + PLAN_HEADLINE_RESERVE.lg + PLAN_NO_PLAN_ROW_RESERVE,
    stack: PLAN_EMPTY_FLOOR.stack + PLAN_HEADLINE_RESERVE.stack + PLAN_NO_PLAN_ROW_RESERVE,
  },
};

/** The Discord block's height by the state `/v1/me` puts it in; `none` and `free` draw no panel. */
export const DISCORD_FLOORS: Readonly<Record<HintState, BucketPx>> = {
  none: { lg: 0, stack: 0 },
  free: { lg: 44, stack: 44 },
  connect: { lg: 320, stack: 340 },
  form: { lg: 650, stack: 640 },
  saved: { lg: 500, stack: 490 },
  off: { lg: 565, stack: 555 },
  lapsed: { lg: 480, stack: 470 },
};

/** What `data-discord-link` on `<html>` says the You panel's Discord link row will be. */
export type LinkHintState = 'none' | 'unlinked' | 'suggested' | 'linked';

/** The link row's height in the You panel, its gap included (experience spec 2026-10-27-discord-bot 3.2). */
export const ACCOUNT_DISCORD_LINK_RESERVE: Readonly<Record<LinkHintState, BucketPx>> = {
  none: { lg: 0, stack: 0 },
  unlinked: { lg: 115, stack: 135 },
  suggested: { lg: 120, stack: 140 },
  linked: { lg: 115, stack: 120 },
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

/** The rail's floor with nothing stored: its three panels, the gaps, and the Discord block under Plan. */
export function railFloor(state: HintState, bucket: ReserveBucket): number {
  const panels = RAIL_FLOORS.devices[bucket] + RAIL_FLOORS.you[bucket] + RAIL_FLOORS.plan[bucket];
  const discord = DISCORD_FLOORS[state][bucket];
  return panels + 2 * SECTION_GAP_PX + (discord > 0 ? PLAN_GAP_PX + discord : 0);
}

/** The reserved grid's height under the header row when nothing is stored (the taller column). */
export function gridFloor(state: HintState, bucket: ReserveBucket): number {
  const left = LEFT_FLOOR[bucket];
  const rail = railFloor(state, bucket);
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
