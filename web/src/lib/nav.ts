// The one source of the primary nav's structure: labels, hrefs, and "is this page current"
// logic. Header.astro renders from this; Header.test.ts asserts against it directly, so a
// reordering shows up as a data-level test failure rather than a markup diff.

export interface NavItem {
  readonly label: string;
  readonly href: string;
}

/** The five doors, in this order on every breakpoint (spec 2026-09-25, section 2). */
export const PRIMARY_NAV_ITEMS: readonly NavItem[] = [
  { label: 'Planner', href: '/planner' },
  { label: 'Simulator', href: '/sim' },
  { label: 'Logs', href: '/logs' },
  { label: 'Rankings', href: '/rankings' },
  { label: 'Guides', href: '/guides' },
];

/** The sixth, quieter item at the end of the row (spec section 2). */
export const SETUP_NAV_ITEM: NavItem = { label: 'Get set up', href: '/setup' };

/** The items after the five doors, in order. One entry today; kept as an array (rather than
 *  inlining SETUP_NAV_ITEM at the call site) so Header.astro's render loop needs no special
 *  case if a second trailing item is ever added. */
export const TRAILING_NAV_ITEMS: readonly NavItem[] = [SETUP_NAV_ITEM];

/**
 * True for the item's own page and any sub-path of it (`/sim` matches `/sim/gear`), so a
 * tool's own sub-pages (the sim tab strip, a saved permalink) still mark the top-level nav
 * entry current. Never matches a different page that merely shares a prefix (`/planner`
 * does not match `/plannerx`).
 */
export function isNavItemCurrent(item: NavItem, path: string): boolean {
  return path === item.href || path.startsWith(`${item.href}/`);
}
