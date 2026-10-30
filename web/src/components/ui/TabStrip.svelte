<!-- web/src/components/ui/TabStrip.svelte -->
<!-- The one tab strip (site alignment spec 2026-10-29 D3, S7): the report's underline tabs.
     Two kinds, one look:
     - In-page tabs (`onSelect` given): a `tablist` of `tab`s with `aria-selected`, a roving
       tabindex (only the current tab is a tab stop), and the arrow keys, Home and End moving
       and selecting (automatic activation). A tab with an `href` stays a real link (a hash the
       page routes on) whose click selects in place. `idPrefix` wires each tab to its panel:
       tab `${idPrefix}-tab-${id}` controls `${idPrefix}-panel-${id}`.
     - Link tabs (no `onSelect`): a `nav` of links, the current one `aria-current="page"`.
     The strip scrolls sideways on a phone with a 24 px fade on a clipped edge, and the current
     tab scrolls into view in the strip, never moving the page. TabStrip.astro is the static
     link-tab twin. Tailwind utilities only: the guild tabs mount client-side after a fetch. -->
<script lang="ts">
  import { scrollFade } from '../../lib/scroll-fade';
  import { TAB_STRIP, scrollTabIntoView, tabClass, tabIdForKey, type TabItem } from '../../lib/ui/tabs';

  let {
    tabs,
    active,
    label,
    onSelect,
    idPrefix,
    testid,
  }: {
    tabs: readonly TabItem[];
    /** The current tab's id. */
    active: string;
    /** The strip's accessible name ("Guild sections"). */
    label: string;
    /** In-page tabs: called with the chosen tab's id. Omitted, the tabs are links. */
    onSelect?: (id: string) => void;
    idPrefix?: string;
    testid?: string;
  } = $props();

  let strip = $state<HTMLElement | undefined>(undefined);

  const tablist = $derived(onSelect !== undefined);

  function tabElement(id: string): HTMLElement | null {
    return strip?.querySelector<HTMLElement>(`[data-tab-id="${CSS.escape(id)}"]`) ?? null;
  }

  function choose(tab: TabItem, event: MouseEvent): void {
    if (onSelect === undefined) return;
    event.preventDefault();
    if (tab.disabled === true) return;
    onSelect(tab.id);
  }

  function onKeydown(event: KeyboardEvent): void {
    if (onSelect === undefined) return;
    const id = tabIdForKey(event.key, tabs, active);
    if (id === null) return;
    event.preventDefault();
    onSelect(id);
    tabElement(id)?.focus();
  }

  $effect(() => {
    const current = active;
    if (strip === undefined) return;
    const tab = tabElement(current);
    if (tab !== null) scrollTabIntoView(strip, tab);
  });
</script>

{#snippet tabLabel(tab: TabItem)}
  {tab.label}{#if tab.suffix !== undefined}<span class={TAB_STRIP.suffix}>{tab.suffix}</span>{/if}
{/snippet}

<div class={TAB_STRIP.wrapper} data-testid={testid}>
  {#if tablist}
    <!-- The tabs are the focus stops (roving tabindex); the list only hears their keys. -->
    <!-- svelte-ignore a11y_interactive_supports_focus -->
    <div
      bind:this={strip}
      use:scrollFade
      class={TAB_STRIP.strip}
      role="tablist"
      aria-label={label}
      aria-orientation="horizontal"
      onkeydown={onKeydown}
    >
      {#each tabs as tab (tab.id)}
        {@const current = tab.id === active}
        {@const disabled = tab.disabled === true}
        {#if tab.href !== undefined && !disabled}
          <a
            href={tab.href}
            role="tab"
            id={idPrefix === undefined ? undefined : `${idPrefix}-tab-${tab.id}`}
            aria-controls={idPrefix === undefined ? undefined : `${idPrefix}-panel-${tab.id}`}
            aria-selected={current}
            tabindex={current ? 0 : -1}
            class={tabClass(current, false)}
            data-tab-id={tab.id}
            data-testid={tab.testid}
            onclick={(event) => choose(tab, event)}>{@render tabLabel(tab)}</a
          >
        {:else}
          <button
            type="button"
            role="tab"
            id={idPrefix === undefined ? undefined : `${idPrefix}-tab-${tab.id}`}
            aria-controls={idPrefix === undefined ? undefined : `${idPrefix}-panel-${tab.id}`}
            aria-selected={current}
            aria-disabled={disabled || undefined}
            tabindex={current && !disabled ? 0 : -1}
            class={tabClass(current, disabled)}
            data-tab-id={tab.id}
            data-testid={tab.testid}
            onclick={(event) => choose(tab, event)}>{@render tabLabel(tab)}</button
          >
        {/if}
      {/each}
    </div>
  {:else}
    <nav bind:this={strip} use:scrollFade class={TAB_STRIP.strip} aria-label={label}>
      {#each tabs as tab (tab.id)}
        {@const current = tab.id === active}
        {#if tab.href !== undefined && tab.disabled !== true}
          <a
            href={tab.href}
            aria-current={current ? 'page' : undefined}
            class={tabClass(current, false)}
            data-tab-id={tab.id}
            data-testid={tab.testid}>{@render tabLabel(tab)}</a
          >
        {:else}
          <span
            aria-disabled="true"
            class={tabClass(false, true)}
            data-tab-id={tab.id}
            data-testid={tab.testid}>{@render tabLabel(tab)}</span
          >
        {/if}
      {/each}
    </nav>
  {/if}
  <div class={TAB_STRIP.fadeStart} aria-hidden="true" data-testid="tab-strip-fade-start"></div>
  <div class={TAB_STRIP.fadeEnd} aria-hidden="true" data-testid="tab-strip-fade-end"></div>
</div>
