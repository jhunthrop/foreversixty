// The one source of the primary nav's structure: labels, hrefs, and "is this page current"
// logic. Header.astro renders from this; Header.test.ts asserts against it directly, so a
// reordering shows up as a data-level test failure rather than a markup diff.

export interface NavItem {
  readonly label: string;
  readonly href: string;
}

/** The four tools, first in the nav on every breakpoint (spec 2026-09-21, section 3). */
export const PRIMARY_NAV_ITEMS: readonly NavItem[] = [
  { label: 'Planner', href: '/planner' },
  { label: 'Simulator', href: '/sim' },
  { label: 'Logs', href: '/logs' },
  { label: 'Rankings', href: '/rankings' },
];

/** Folded under the "Reference" disclosure. */
export const REFERENCE_NAV_ITEMS: readonly NavItem[] = [
  { label: 'Classes', href: '/classes' },
  { label: 'Guides', href: '/guides' },
  { label: 'Zones', href: '/zones' },
  { label: 'Dungeons', href: '/dungeons' },
];

/** Promoted from the footer into the primary nav. */
export const ADDON_NAV_ITEM: NavItem = { label: 'The addon', href: '/addon' };

/** Last in the row: what the site sells is one click from every page, not a footer link. */
export const PREMIUM_NAV_ITEM: NavItem = { label: 'Premium', href: '/premium' };

/** The items after the Reference disclosure, in order. */
export const TRAILING_NAV_ITEMS: readonly NavItem[] = [ADDON_NAV_ITEM, PREMIUM_NAV_ITEM];

/**
 * True for the item's own page and any sub-path of it (`/sim` matches `/sim/gear`), so a
 * tool's own sub-pages (the sim tab strip, a saved permalink) still mark the top-level nav
 * entry current. Never matches a different page that merely shares a prefix (`/planner`
 * does not match `/plannerx`).
 */
export function isNavItemCurrent(item: NavItem, path: string): boolean {
  return path === item.href || path.startsWith(`${item.href}/`);
}

/** True when the current page is one the Reference disclosure holds, so its summary can
 * carry `aria-current` even though none of its own children render as the trigger. */
export function isReferenceCurrent(path: string): boolean {
  return REFERENCE_NAV_ITEMS.some((item) => isNavItemCurrent(item, path));
}
