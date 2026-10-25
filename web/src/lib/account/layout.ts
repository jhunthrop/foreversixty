// web/src/lib/account/layout.ts
// Reserved-height constants for /account's loading skeletons and its signed-out/signed-in
// first-screen match (brief 2026-09-22 §B2/§B6/§B8). Re-measured against a real
// `FOREVER_DATA=fixture` build at 360px after the page-header/two-column/rail-panel
// redesign (account-visual lane, 2026-09-22), with route-stubbed /v1/me responses so the
// measurement isn't confounded by network conditions. The title itself (`<h1>`) now
// renders identically in every status -- loading, failed, signed-out, ready -- so it is
// common to all four and excluded from the figures below; what each constant reserves is
// everything from just under the title to the bottom of its section. As with the previous
// lane's measurement, `npm run lhci`'s own account.html run cannot confirm this directly:
// its sandbox has no CORS allowance for localhost to call the real API, so it always lands
// on the `failed`/LoadError branch instead -- see SIGNED_OUT_MIN_H's own comment.

import type { Me } from './api';
import { storage } from '../layout-reserve';

/** The identity line (BattleTag/email, sign-in method, Sign out), the compact
 *  current-character chip, and the Characters panel -- everything the ready view shows
 *  before "Your reports". measured against meBnetFixture at 360px, 2026-09-22 */
export const CHARACTERS_SKELETON_MIN_H = 'min-h-[692px]';

/** "Your reports" and the rail's Devices panel -- the two sections that follow Characters
 *  in reading order at 360px (the two-column grid only applies at `lg`; below that every
 *  panel stacks in one column). measured against meBnetFixture at 360px, 2026-09-22 */
export const IDENTITY_SKELETON_MIN_H = 'min-h-[332px]';

/** The rail's You (pseudonym toggle) and Guilds-and-plan panels.
 *  measured against meBnetFixture at 360px, 2026-09-22 */
export const MORE_SKELETON_MIN_H = 'min-h-[308px]';

/** The signed-out prompt's floor. `/account`'s loading branch always renders all three
 *  skeletons above (Account.svelte does not know `signedIn` until `/v1/me` answers), so a
 *  genuinely signed-out visitor only avoids a layout shift if this block reserves the SAME
 *  total height the loading skeletons already reserved. Same fixed-height-regardless-of-
 *  branch pattern CurrentCharacterChip.svelte's CHIP_HEIGHT and the pairing-mode block
 *  above already use in this file, for the same reason.
 *  1396px = CHARACTERS_SKELETON_MIN_H (692) + IDENTITY_SKELETON_MIN_H (332) +
 *  MORE_SKELETON_MIN_H (308) + two gap-8 gaps (2 x 32) between them, matching the loading
 *  branch's `flex flex-col gap-8` wrapper exactly. Unverifiable against `npm run lhci`'s
 *  own account.html run for the CORS reason above; the previous lane confirmed via direct
 *  DOM/CLS instrumentation that the exact figure does not change lhci's measured CLS
 *  (the sandbox never reaches this branch at all), so 1396px is kept because it is the
 *  mathematically correct value for a real signed-out visitor. */
export const SIGNED_OUT_MIN_H = 'min-h-[1396px]';

// ---- The loading reserve (experience spec 2026-10-25-discord-notifications §9.1) ----------
// The loading state draws the ready page's skeleton at its real positions. Two buckets:
// `lg` is the 12-column grid at 1024 and up (measured at 1440), `stack` is the one-column
// page below it (measured at 390). The floors are the spec's targets until the lane's
// measurement pass replaces them; they apply only when no rendered height is stored.

export type ReserveBucket = 'lg' | 'stack';

/** One number per bucket: the 1440 desktop layout and the 390 phone layout. */
export interface BucketPx {
  lg: number;
  stack: number;
}

/** What `data-discord` on `<html>` says the Discord block will be. */
export type HintState = 'none' | 'free' | 'form' | 'saved' | 'off' | 'lapsed';

/** The ready header row: 44 px beside the title at 1024 and up, the stacked row below it. */
export const HEADER_ROW_H: BucketPx = { lg: 44, stack: 84 };

/** An empty Characters panel, an empty Your reports panel and the 32 px gap between them. */
export const LEFT_FLOOR: BucketPx = { lg: 300, stack: 300 };

/** The rail's panels with nothing extra in them (Devices, You, Plan). */
export const RAIL_FLOORS: Readonly<Record<'devices' | 'you' | 'plan', BucketPx>> = {
  devices: { lg: 232, stack: 232 },
  you: { lg: 192, stack: 192 },
  plan: { lg: 219, stack: 219 },
};

/** The Discord block's height by the state `/v1/me` puts it in; `none` and `free` draw no panel. */
export const DISCORD_FLOORS: Readonly<Record<HintState, BucketPx>> = {
  none: { lg: 0, stack: 0 },
  free: { lg: 44, stack: 44 },
  form: { lg: 610, stack: 630 },
  saved: { lg: 460, stack: 475 },
  off: { lg: 540, stack: 565 },
  lapsed: { lg: 475, stack: 475 },
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

/** The stored heights of the last ready pages; null unless both buckets are present and valid. */
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
 * Which Discord block the page will draw, from the cached `/v1/me` (spec 4.1 precedence). The
 * pre-paint script in Base.astro copies this logic verbatim: keep it small and import-free.
 */
export function discordHintState(me: Me | null): HintState {
  if (me === null || me.notifications === undefined) return 'none';
  const discord = me.notifications.discord;
  const entitled = me.entitlements?.notifications === true;
  if (discord === null || discord === undefined) return entitled ? 'form' : 'free';
  if (!entitled) return 'lapsed';
  return discord.disabled_reason === null ? 'saved' : 'off';
}
