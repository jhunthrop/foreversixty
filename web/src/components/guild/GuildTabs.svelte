<!-- web/src/components/guild/GuildTabs.svelte -->
<!-- Guild control-centre spec §4.0/§7: the seven-tab strip, hash-anchored on the one guild
     URL (#roster, #raids, ...), never a separate route. Drawn by the shared TabStrip. -->
<script lang="ts" module>
  import { tabIdOfHash } from '../../lib/guild/raids-hash';

  export type GuildTabId =
    'overview' | 'roster' | 'raids' | 'progression' | 'readiness' | 'loot' | 'settings';

  export const GUILD_TABS: { id: GuildTabId; label: string }[] = [
    { id: 'overview', label: 'Overview' },
    { id: 'roster', label: 'Roster' },
    { id: 'raids', label: 'Raids' },
    { id: 'progression', label: 'Progression' },
    { id: 'readiness', label: 'Readiness' },
    { id: 'loot', label: 'Loot' },
    { id: 'settings', label: 'Settings' },
  ];

  const PUBLIC_TABS: readonly GuildTabId[] = ['overview', 'raids', 'progression'];
  const MEMBER_TABS: readonly GuildTabId[] = [
    'overview',
    'roster',
    'raids',
    'progression',
    'readiness',
    'loot',
  ];
  const OFFICER_TABS: readonly GuildTabId[] = [...MEMBER_TABS, 'settings'];

  /** Spec §4.0's visibility matrix, stated once: which tab ids a role may ever reach,
   *  including by a stale/bookmarked hash (never a 403 page, never a tab rendered with its
   *  content withheld row by row). */
  export function tabsForRole(role: 'public' | 'member' | 'officer' | 'moderator'): readonly GuildTabId[] {
    if (role === 'public') return PUBLIC_TABS;
    if (role === 'member') return MEMBER_TABS;
    return OFFICER_TABS;
  }

  export function tabFromHash(hash: string, allowed: readonly GuildTabId[]): GuildTabId {
    const id = tabIdOfHash(hash);
    return (allowed as readonly string[]).includes(id) ? (id as GuildTabId) : allowed[0];
  }

  /** A hash change: one that names a tab (or none) routes; an in-page anchor that names no tab
   *  (Settings' `#billing`, `#discord`) keeps the open tab, so following it never unmounts it. */
  export function tabAfterHashChange(
    hash: string,
    allowed: readonly GuildTabId[],
    current: GuildTabId,
  ): GuildTabId {
    const id = tabIdOfHash(hash);
    const namesTab = id === '' || GUILD_TABS.some((tab) => tab.id === id);
    return namesTab ? tabFromHash(hash, allowed) : current;
  }
</script>

<script lang="ts">
  import TabStrip from '../ui/TabStrip.svelte';
  import type { TabItem } from '../../lib/ui/tabs';

  let {
    active,
    role,
    onSelect,
    hrefBase,
  }: {
    active: GuildTabId;
    role: 'public' | 'member' | 'officer' | 'moderator';
    onSelect?: (tab: GuildTabId) => void;
    /** Off the guild page (a raid page): each tab is a plain link to the guild page's hash. */
    hrefBase?: string;
  } = $props();

  const tabs = $derived<TabItem[]>(
    GUILD_TABS.filter((tab) => tabsForRole(role).includes(tab.id)).map((tab) => ({
      id: tab.id,
      label: tab.label,
      href: `${hrefBase ?? ''}#${tab.id}`,
      suffix: tab.id === 'loot' && role === 'member' ? ' · read-only' : undefined,
      testid: `guild-tab-${tab.id}`,
    })),
  );
</script>

<!-- The shared TabStrip (site alignment spec 2026-10-29 D3): in-page tabs on the guild page, where
     a click selects in place and the page routes the hash; link tabs to the guild page's hash
     off it (a raid page). It scrolls with the edge fades and keeps the current tab in view. -->
<TabStrip
  {tabs}
  {active}
  label="Guild sections"
  onSelect={onSelect === undefined ? undefined : (id) => onSelect(id as GuildTabId)}
  testid="guild-tabs"
/>
