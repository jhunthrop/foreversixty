<!-- web/src/components/ui/TabStripMore.svelte -->
<!-- TabStrip with an overflow (logs spec 2026-11-07 section 3A; the mock library's TabStripMore):
     link tabs, the first `md` below 1760 px and the first `xl` from 1760, the rest in a More menu
     on the overlay surface. More reads "More: {tab}" and carries the gold rule while the current
     tab is inside it. Below 768 the strip is one line that scrolls sideways, edge to edge. A click
     on a tab calls `onSelect` (the page routes in place); its href keeps a copied link and a
     middle click working. Escape or a click outside closes the menu and focus returns to More. -->
<script lang="ts">
  import { TAB_STRIP, tabClass, type TabItem } from '../../lib/ui/tabs';
  import Overlay from './Overlay.svelte';

  let {
    tabs,
    active,
    md,
    xl,
    label,
    onSelect,
    moreLabel = 'More',
    testid,
  }: {
    tabs: readonly TabItem[];
    active: string;
    /** Tabs shown below 1760 px. */
    md: number;
    /** Tabs shown from 1760 px. */
    xl: number;
    label: string;
    onSelect?: (id: string) => void;
    moreLabel?: string;
    testid?: string;
  } = $props();

  /** Which More is open: the one below 1760 ('md') or from it ('xl'). */
  let open = $state<'md' | 'xl' | null>(null);
  let root = $state<HTMLElement | null>(null);

  const shown = $derived(tabs.slice(0, xl));
  const rest = $derived({ md: tabs.slice(md), xl: tabs.slice(xl) });

  function moreText(list: readonly TabItem[]): string {
    const current = list.find((tab) => tab.id === active);
    return current === undefined ? moreLabel : `${moreLabel}: ${current.label}`;
  }

  function choose(tab: TabItem, event: MouseEvent): void {
    open = null;
    if (tab.disabled === true) {
      event.preventDefault();
      return;
    }
    if (onSelect === undefined) return;
    if (event.metaKey || event.ctrlKey || event.shiftKey || event.button !== 0) return;
    event.preventDefault();
    onSelect(tab.id);
  }

  function close(refocus: boolean): void {
    const which = open;
    open = null;
    if (refocus && which !== null) root?.querySelector<HTMLElement>(`[data-more="${which}"]`)?.focus();
  }

  function onKey(event: KeyboardEvent): void {
    if (open === null) return;
    if (event.key === 'Escape') {
      event.preventDefault();
      close(true);
      return;
    }
    if (event.key !== 'ArrowDown' && event.key !== 'ArrowUp') return;
    const items = [...(root?.querySelectorAll<HTMLElement>('[role="menuitemradio"]') ?? [])];
    if (items.length === 0) return;
    event.preventDefault();
    const at = items.indexOf(document.activeElement as HTMLElement);
    const step = event.key === 'ArrowDown' ? 1 : -1;
    items[(at + step + items.length) % items.length]?.focus();
  }

  $effect(() => {
    if (open === null) return;
    const outside = (event: PointerEvent): void => {
      if (root !== null && !root.contains(event.target as Node)) close(false);
    };
    document.addEventListener('pointerdown', outside);
    root?.querySelector<HTMLElement>('[role="menuitemradio"]')?.focus();
    return () => document.removeEventListener('pointerdown', outside);
  });
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="ui-tsm" bind:this={root} onkeydown={onKey}>
  <nav class={`${TAB_STRIP.strip} ui-tsm-strip`} aria-label={label} data-testid={testid}>
    {#each shown as tab, index (tab.id)}
      {@const current = tab.id === active}
      <a
        href={tab.href}
        aria-current={current ? 'page' : undefined}
        aria-disabled={tab.disabled === true ? 'true' : undefined}
        class={[tabClass(current, tab.disabled === true), index >= md && 'ui-tsm-xl']}
        data-tab-id={tab.id}
        data-testid={tab.testid}
        onclick={(event) => choose(tab, event)}>{tab.label}</a
      >
    {/each}
    {#each ['md', 'xl'] as const as which (which)}
      {@const list = rest[which]}
      {@const inside = list.some((tab) => tab.id === active)}
      <button
        type="button"
        class={[tabClass(inside, false), 'ui-tsm-more', `ui-tsm-more-${which}`]}
        aria-haspopup="menu"
        aria-expanded={open === which}
        data-more={which}
        data-testid={testid === undefined ? undefined : `${testid}-more`}
        onclick={() => (open = open === which ? null : which)}
        >{moreText(list)}<span aria-hidden="true" class="ml-[6px] text-[11px]">▾</span></button
      >
    {/each}
  </nav>
  {#if open !== null}
    <div class="ui-tsm-menu">
      <Overlay kind="menu" label={moreLabel}>
        <div class="flex flex-col py-1">
          {#each rest[open] as tab (tab.id)}
            {@const current = tab.id === active}
            <a
              href={tab.href}
              role="menuitemradio"
              aria-checked={current}
              class="ui-tsm-item"
              data-current={current ? 'true' : undefined}
              data-tab-id={tab.id}
              data-testid={tab.testid}
              onclick={(event) => choose(tab, event)}>{tab.label}</a
            >
          {/each}
        </div>
      </Overlay>
    </div>
  {/if}
</div>

<style>
  .ui-tsm {
    position: relative;
  }
  .ui-tsm-more-xl,
  .ui-tsm-xl {
    display: none;
  }
  @media (min-width: 1760px) {
    .ui-tsm-xl,
    .ui-tsm-more-xl {
      display: inline-flex;
    }
    .ui-tsm-more-md {
      display: none;
    }
  }
  @media (max-width: 767px) {
    .ui-tsm-strip {
      margin-inline: calc(-1 * var(--page-inset));
      padding-inline: calc(var(--page-inset) - 14px);
    }
  }
  .ui-tsm-menu {
    position: absolute;
    right: 0;
    top: calc(100% + 4px);
    z-index: 40;
    min-width: 220px;
  }
  .ui-tsm-item {
    display: flex;
    align-items: center;
    min-height: 44px;
    padding: 0 16px;
    font-size: 14px;
    font-weight: 600;
    color: var(--color-nav);
    box-shadow: inset 3px 0 0 transparent;
  }
  .ui-tsm-item:hover,
  .ui-tsm-item:focus-visible {
    color: var(--color-strong);
    outline: none;
  }
  .ui-tsm-item[data-current='true'] {
    color: var(--color-strong);
    box-shadow: inset 3px 0 0 var(--color-gold);
  }
</style>
