// The one source of the primary nav's structure: round 5's doors (design/specs/2026-11-09-directions.md
// section 5, coordinator ruling 2026-11-09: the mock's flat doors), their panel rows where a door
// has sub-pages, and "is this page current" logic. Header.astro renders from this; Header.test.ts
// asserts against it directly, so a reordering shows up as a data-level test failure rather than
// a markup diff. Get set up left the bar for the account menu and the footer.

import communityLinks from '../data/links.json' with { type: 'json' };
import type { UiIconName } from './ui/ui-icons';

/** The Guild panel's last row: the site's one header link to the community Discord. */
export const DISCORD_ROW_LABEL = 'Forever Sixty Discord';

export interface NavItem {
  readonly label: string;
  readonly href: string;
}

/** A door on the bar: its label leads with the owner's custom icon (lib/ui/ui-icons.ts). */
export interface NavDoor extends NavItem {
  readonly icon: UiIconName;
}

/** A panel row: its icon, or the Discord brand tile in its place (the Guild panel's last row). */
export type NavRow = NavItem & { readonly description: string } & (
    { readonly icon: UiIconName; readonly mark?: never } | { readonly mark: 'discord'; readonly icon?: never }
  );

export interface NavGroup extends NavDoor {
  /** The panel's rows. Empty for Guides (the header renders its class grid from classes.json)
   *  and for a plain door. */
  readonly rows: readonly NavRow[];
  readonly foot?: NavItem;
  /** True for a door with sub-pages: a chevron and a panel. A plain door is its link alone. */
  readonly panel: boolean;
}

/** A plain door: a link with no panel. */
const door = (label: string, href: string, icon: UiIconName): NavGroup => ({
  label,
  href,
  icon,
  rows: [],
  panel: false,
});

/**
 * The doors, in the mock's order on every breakpoint. A door is a real link to its own href; a
 * door with sub-pages (Guides, Logs, Guild) has its rows behind a chevron panel.
 */
export const NAV_GROUPS: readonly NavGroup[] = [
  door('Best in slot', '/bis', 'bis'),
  {
    label: 'Guides',
    href: '/guides',
    icon: 'guides',
    rows: [],
    foot: { label: 'All guides', href: '/guides' },
    panel: true,
  },
  door('Tier list', '/tiers', 'tier-list'),
  door('Sim', '/sim', 'sim'),
  door('Planner', '/planner', 'planner'),
  {
    label: 'Logs',
    href: '/logs',
    icon: 'logs',
    rows: [
      {
        label: 'Upload a log',
        href: '/logs',
        icon: 'upload',
        description: 'Drop a combat log, or let the companion stream it live',
      },
      {
        label: 'Your reports',
        href: '/reports',
        icon: 'logs',
        description: 'Every fight you logged, newest first',
      },
      {
        label: 'Rankings',
        href: '/rankings',
        icon: 'rankings',
        description: 'Parses by boss, spec and level range, from every log',
      },
    ],
    panel: true,
  },
  {
    label: 'Guild',
    href: '/guild',
    icon: 'guild',
    rows: [
      {
        label: 'Your guild',
        href: '/guild',
        icon: 'guild',
        description: "Progress, the week's raids, who logged what",
      },
      {
        label: 'Roster',
        href: '/guild#roster',
        icon: 'manage',
        description: 'Everyone the addon has seen, with gear score and last sync',
      },
      {
        label: 'Loot council',
        href: '/guild#loot',
        icon: 'bis',
        description: 'Who gains most from each drop, by the sim, before the roll',
      },
      {
        label: 'Guild logs',
        href: '/guild#raids',
        icon: 'logs',
        description: 'Reports from guild members, raid night by raid night',
      },
      {
        label: DISCORD_ROW_LABEL,
        href: communityLinks.discordInvite,
        description: 'Find a guild, ask about the site, report a bug.',
        mark: 'discord',
      },
    ],
    foot: { label: 'Guild settings', href: '/guild#settings' },
    panel: true,
  },
];

/** Where the Guild group sits in NAV_GROUPS: the session fills its rows (spec section 2). */
export const GUILD_GROUP_INDEX = NAV_GROUPS.findIndex((group) => group.label === 'Guild');

/**
 * What the Guild panel's two member-only doors say for a visitor with no guild of their own
 * (signed out or guildless), keyed by the tab they open for a member. The rows above are the
 * member's: the header script swaps these in when the session lists no guild, so neither door is
 * a link to a section the landing page does not have.
 */
export const GUILDLESS_SWAPS: Readonly<Record<string, NavItem & { readonly description?: string }>> = {
  roster: {
    label: 'Find a guild',
    href: '/guild',
    description: 'Look up a guild by name, or open one of the top guilds',
  },
  settings: { label: 'Claim your guild', href: '/guild#claim' },
};

/** The Guild row without a hash: the guild's overview (`data-guild-tab` in the header). */
export const GUILD_OVERVIEW_TAB = 'overview';

/** Get set up: the account menu's row and the footer's link (round 5: off the bar's doors); also
 *  the guild claim flow's link. */
export const SETUP_NAV_ITEM: NavItem = { label: 'Get set up', href: '/setup' };

/** Premium, their teal Subscribe at the bar's end: a link with no panel and no chevron, hidden for
 *  an account with Premium from any source (spec 2026-10-29-premium-entry sections 3 and 4). */
export const PREMIUM_NAV_ITEM: NavDoor = { label: 'Premium', href: '/premium', icon: 'premium' };

/**
 * True for the item's own page and any sub-path of it (`/sim` matches `/sim/gear`), so a
 * tool's own sub-pages (the sim tab strip, a saved permalink) still mark the top-level nav
 * entry current. Never matches a different page that merely shares a prefix (`/planner`
 * does not match `/plannerx`).
 */
export function isNavItemCurrent(item: NavItem, path: string): boolean {
  return path === item.href || path.startsWith(`${item.href}/`);
}

/** The stable key of a door or row in ids and test ids: "Upload a log" -> "upload-a-log". */
export function navKey(label: string): string {
  return label
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, '-')
    .replace(/^-|-$/g, '');
}

const withoutHash = (href: string): string => href.split('#')[0];

/** True when the page is the group's own page or any of its rows' pages (hashes ignored). */
export function isNavGroupCurrent(group: NavGroup, path: string): boolean {
  return [group, ...group.rows].some((item) =>
    isNavItemCurrent({ label: item.label, href: withoutHash(item.href) }, path),
  );
}
