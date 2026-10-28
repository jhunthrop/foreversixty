// web/src/lib/billing/checkout-plan-switch.ts
// Spec §5.3: `Switch to Guild` / `Switch to Raid guild` rewrite the checkout URL's `plan` in
// place (history.replaceState), so the band and the buttons follow with no reload.
import type { PlanKey } from './plans';

/** The parts of `window.location` the rewrite reads. */
export interface PlanSwitchLocation {
  pathname: string;
  search: string;
  hash: string;
}

/** The parts of `window.history` the rewrite uses. */
export interface PlanSwitchHistory {
  state: unknown;
  replaceState: (data: unknown, unused: string, url: string) => void;
}

/** `search` with its `plan` set to `plan`; every other param kept in order. */
export function withPlan(search: string, plan: PlanKey): string {
  const params = new URLSearchParams(search);
  params.set('plan', plan);
  return `?${params.toString()}`;
}

/** Rewrites the current URL's plan in place: no navigation, no new history entry. */
export function replacePlanInUrl(
  plan: PlanKey,
  location: PlanSwitchLocation = window.location,
  history: PlanSwitchHistory = window.history,
): void {
  const url = `${location.pathname}${withPlan(location.search, plan)}${location.hash}`;
  history.replaceState(history.state, '', url);
}
