// The one source of the primary nav's structure (spec design/specs/2026-10-15-grouped-header.md,
// which supersedes the eight doors of 2026-10-09-nav-character-selector.md section 2.2): the
// four doors, their panel rows, the trailing Get set up door, and "is this page current" logic. Header.astro renders from
// this; Header.test.ts asserts against it directly, so a reordering shows up as a data-level
// test failure rather than a markup diff.

import communityLinks from '../data/links.json' with { type: 'json' };

/** The Guild panel's last row: the site's one header link to the community Discord. */
export const DISCORD_ROW_LABEL = 'Forever Sixty Discord';

export interface NavItem {
  readonly label: string;
  readonly href: string;
}

export interface NavRow extends NavItem {
  readonly description: string;
  /** A tinted brand tile in place of the line icon (the Guild panel's Discord row). */
  readonly mark?: 'discord';
}

export interface NavGroup extends NavItem {
  /** The panel's rows. Empty for Guides: the header renders its class grid from classes.json. */
  readonly rows: readonly NavRow[];
  readonly foot?: NavItem;
}

/**
 * The four doors, in this order on every breakpoint. A door is a real link to its own href;
 * the rows are the panel behind it.
 */
export const NAV_GROUPS: readonly NavGroup[] = [
  {
    label: 'Tools',
    href: '/planner',
    rows: [
      {
        label: 'Planner',
        href: '/planner',
        description: 'Talents, the build you follow, and what to train next',
      },
      { label: 'BiS', href: '/bis', description: 'The gear to chase at your level, slot by slot' },
      {
        label: 'Simulator',
        href: '/sim',
        description: 'What your gear does, simulated fight by fight',
      },
    ],
  },
  {
    label: 'Logs',
    href: '/logs',
    rows: [
      {
        label: 'Upload a log',
        href: '/logs',
        description: 'Drop a combat log, or let the companion stream it live',
      },
      {
        label: 'Your reports',
        href: '/reports',
        description: 'Every fight you logged, newest first',
      },
      {
        label: 'Rankings',
        href: '/rankings',
        description: 'Parses by boss, spec and level range, from every log',
      },
      {
        label: 'Tier list',
        href: '/tiers',
        description: 'DPS, healer and tank, ranked from the same logs',
      },
    ],
  },
  {
    label: 'Guild',
    href: '/guild',
    rows: [
      {
        label: 'Your guild',
        href: '/guild',
        description: "Progress, the week's raids, who logged what",
      },
      {
        label: 'Roster',
        href: '/guild#roster',
        description: 'Everyone the addon has seen, with gear score and last sync',
      },
      {
        label: 'Loot council',
        href: '/guild#loot',
        description: 'Who gains most from each drop, by the sim, before the roll',
      },
      {
        label: 'Guild logs',
        href: '/guild#raids',
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
  },
  {
    label: 'Guides',
    href: '/guides',
    rows: [],
    foot: { label: 'All guides', href: '/guides' },
  },
];

/** Where the Guild group sits in NAV_GROUPS: the session fills its rows (spec section 2). */
export const GUILD_GROUP_INDEX = 2;

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

/** The bar's last door, after the four groups, with no panel of its own (owner, 2026-10-15:
 *  setup stays on the primary bar); also the guild claim flow's link. */
export const SETUP_NAV_ITEM: NavItem = { label: 'Get set up', href: '/setup' };

/** The bar's gold Premium tag after Get set up: a link with no panel and no chevron, hidden for
 *  an account with Premium from any source (spec 2026-10-29-premium-entry sections 3 and 4). */
export const PREMIUM_NAV_ITEM: NavItem = { label: 'Premium', href: '/premium' };

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
