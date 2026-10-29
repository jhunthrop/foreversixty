// web/src/lib/premium-nav.ts
// Whether the header's Premium tag shows (spec 2026-10-29-premium-entry sections 3 and 4). The
// tag is static markup in every page's header; Base.astro's head script hides it before first
// paint from the stored hint (`html[data-premium-nav="hide"]`). Once `/v1/me` answers, the
// selector island calls `applySessionPremium`: an account with Premium whose tag is still
// showing gets `hold` (visibility hidden, the box kept), so the doors never move after paint.
import type { PremiumSource } from './billing/contract';

/** The root attribute's name in `dataset` form; `data-premium-nav` in markup and CSS. */
export const PREMIUM_NAV_DATASET_KEY = 'premiumNav';

/** `hide`: decided before paint (display none). `hold`: decided after paint (visibility hidden). */
export type PremiumNavState = 'hide' | 'hold';

/**
 * Who sees the tag: signed-out visitors (`null`) and free accounts. Premium from any source
 * (own plan, seat, grant, the fallback entitlement) hides it. An officer whose guild pays but
 * who holds no seat has no personal Premium, so `premiumSource` is `none` and the tag shows.
 */
export function premiumNavShown(source: PremiumSource | null): boolean {
  return source === null || source === 'none';
}

/**
 * The state after `/v1/me`: a hidden tag never comes back during this page view (the next
 * page's pre-paint reads the refreshed hint), and a showing tag is only ever held, never
 * removed from the flow.
 */
export function nextPremiumNavState(
  current: PremiumNavState | null,
  source: PremiumSource | null,
): PremiumNavState | null {
  if (current !== null) return current;
  return premiumNavShown(source) ? null : 'hold';
}

function readState(root: HTMLElement): PremiumNavState | null {
  const value = root.dataset[PREMIUM_NAV_DATASET_KEY];
  return value === 'hide' || value === 'hold' ? value : null;
}

/** Applies the session's answer to the root element; a no-op when nothing changes. */
export function applySessionPremium(
  source: PremiumSource | null,
  root: HTMLElement = document.documentElement,
): void {
  const current = readState(root);
  const next = nextPremiumNavState(current, source);
  if (next !== null && next !== current) root.dataset[PREMIUM_NAV_DATASET_KEY] = next;
}
