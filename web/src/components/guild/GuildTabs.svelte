<!-- web/src/components/guild/GuildTabs.svelte -->
<!-- Guild control-centre spec §4.0/§7: the seven-tab strip, hash-anchored on the one guild
     URL (#roster, #raids, ...), never a separate route. 44px pills (ClassHeader.astro's own
     `SpecTabs` family, rebuilt fresh here since this page has no ClassHeader instance to
     extend). Scrolls horizontally with a peek on phone; the active tab is scrolled into
     view programmatically on load (spec §7: a 2px gold sliver off-screen is not enough on
     its own to say which tab is active). -->
<script lang="ts" module>
  export type GuildTabId = 'overview' | 'roster' | 'raids' | 'progression' | 'readiness' | 'loot' | 'settings';

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
  const MEMBER_TABS: readonly GuildTabId[] = ['overview', 'roster', 'raids', 'progression', 'readiness', 'loot'];
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
    const id = hash.replace(/^#/, '').toLowerCase();
    return (allowed as readonly string[]).includes(id) ? (id as GuildTabId) : allowed[0];
  }
</script>

<script lang="ts">
  let {
    active,
    role,
    onSelect,
  }: { active: GuildTabId; role: 'public' | 'member' | 'officer' | 'moderator'; onSelect: (tab: GuildTabId) => void } =
    $props();

  const allowed = $derived(tabsForRole(role));
  let stripEl = $state<HTMLDivElement | undefined>(undefined);

  function select(tab: GuildTabId, event: MouseEvent): void {
    event.preventDefault();
    onSelect(tab);
  }

  /** Spec §7: scroll the active pill into view on load/activation rather than relying on
   *  the horizontal-scroll peek alone. */
  $effect(() => {
    const currentActive = active;
    if (stripEl === undefined) return;
    const pill = stripEl.querySelector<HTMLElement>(`[data-guild-tab="${currentActive}"]`);
    pill?.scrollIntoView({ block: 'nearest', inline: 'nearest' });
  });
</script>

<div
  bind:this={stripEl}
  class="guild-tab-strip flex gap-2 overflow-x-auto"
  role="tablist"
  aria-label="Guild sections"
  data-testid="guild-tabs"
>
  {#each GUILD_TABS.filter((tab) => allowed.includes(tab.id)) as tab (tab.id)}
    {@const isActive = tab.id === active}
    {@const readOnlySuffix = tab.id === 'loot' && role === 'member'}
    <a
      href={`#${tab.id}`}
      role="tab"
      aria-selected={isActive}
      data-guild-tab={tab.id}
      data-testid={`guild-tab-${tab.id}`}
      class="guild-tab-pill"
      class:is-active={isActive}
      onclick={(event) => select(tab.id, event)}
    >
      {tab.label}
      {#if readOnlySuffix}<span class="guild-tab-readonly"> · read-only</span>{/if}
    </a>
  {/each}
</div>

<style>
  .guild-tab-strip {
    scrollbar-width: none;
  }
  .guild-tab-strip::-webkit-scrollbar {
    display: none;
  }
  .guild-tab-pill {
    display: inline-flex;
    flex-shrink: 0;
    align-items: center;
    height: 44px;
    box-sizing: border-box;
    padding: 0 14px;
    border-radius: var(--radius-control);
    border: 1px solid var(--color-line);
    font-size: 12px;
    font-weight: 700;
    letter-spacing: 0.06em;
    text-transform: uppercase;
    color: var(--color-text);
    white-space: nowrap;
  }
  .guild-tab-pill.is-active {
    border-color: var(--color-gold);
    background: rgba(229, 185, 85, 0.08);
    color: var(--color-gold);
  }
  .guild-tab-pill:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .guild-tab-readonly {
    color: var(--color-muted);
    font-size: 10px;
    text-transform: none;
    letter-spacing: normal;
  }
</style>
