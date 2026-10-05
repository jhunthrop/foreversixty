<!-- web/src/components/guild/GuildTabs.svelte -->
<!-- Guild control-centre spec §4.0/§7: the seven-tab strip, hash-anchored on the one guild
     URL (#roster, #raids, ...), never a separate route. 44px pills (ClassHeader.astro's own
     `SpecTabs` family, rebuilt fresh here since this page has no ClassHeader instance to
     extend). Scrolls horizontally with a peek on phone; the active tab is scrolled into
     view programmatically on load (spec §7: a 2px gold sliver off-screen is not enough on
     its own to say which tab is active). -->
<script lang="ts" module>
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
    const id = hash.replace(/^#/, '').toLowerCase();
    return (allowed as readonly string[]).includes(id) ? (id as GuildTabId) : allowed[0];
  }
</script>

<script lang="ts">
  let {
    active,
    role,
    onSelect,
  }: {
    active: GuildTabId;
    role: 'public' | 'member' | 'officer' | 'moderator';
    onSelect: (tab: GuildTabId) => void;
  } = $props();

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

<!-- Tailwind utilities only, never a scoped <style> block: this component (and every
     other tab panel reached through it) only ever mounts client-side after the guild's
     own async fetch resolves, so it is absent from the page's server-rendered markup --
     this build's own Astro/Vite pipeline only inlines a Svelte component's scoped CSS for
     classes present in that initial SSR output, silently dropping styles for anything
     that renders later (found when this round's own capture showed unstyled, borderless
     tab pills). Tailwind's utility stylesheet has no such gap: it is generated from every
     class literal anywhere in the source tree, independent of what actually renders for a
     given request, so it is the only reliable way to style a lazily-mounted region on this
     page. Every guild/* component in this round follows the same rule. -->
<div
  bind:this={stripEl}
  class="flex [scrollbar-width:none] gap-2 overflow-x-auto [&::-webkit-scrollbar]:hidden"
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
      class={`rounded-control box-border inline-flex h-11 shrink-0 items-center border px-3.5 text-[12px] font-bold tracking-[0.06em] whitespace-nowrap uppercase ${isActive ? 'border-gold text-gold' : 'border-line text-text'}`}
      style={isActive ? 'background: rgba(229,185,85,.08)' : ''}
      onclick={(event) => select(tab.id, event)}
    >
      {tab.label}
      {#if readOnlySuffix}<span class="text-muted ml-1 text-[10px] tracking-normal normal-case">
          · read-only</span
        >{/if}
    </a>
  {/each}
</div>
