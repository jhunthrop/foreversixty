<!-- web/src/components/ui/FilterPopover.svelte -->
<!-- The one filter popover (component audit 2026-11-10 cluster 14; the BiS Sources pattern, web
     layout spec 2026-11-05 §5.11): from 768 px a compact button with a chevron ("Filters (2)",
     "Sources (6)") opens the caller's `panel` on the shared overlay surface; Escape, a click outside
     or focus leaving closes it, and Escape returns focus to the button. Below 768 px the `stack`
     snippet takes its place in the expanded control stack, under the group label. No Popover API
     (an old browser): the panel hangs under the button in the page instead. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import Icon from './Icon.svelte';
  import Overlay from './Overlay.svelte';
  import { ICON_CHEVRON_DOWN } from '../../lib/ui/ui-icons';
  import { button as buttonClass } from '../../lib/ui/button';
  import { popoverSupported } from '../../lib/ui/listbox-model';
  import type { PopoverHandle } from '../../lib/ui/popover';

  let {
    label,
    buttonLabel,
    buttonAriaLabel,
    id,
    testid,
    buttonTestid,
    popoverTestid,
    width = 300,
    focusSelector = 'input, select, button',
    panel,
    foot,
    stack,
  }: {
    /** The group's name: the dialog's accessible name and the phone stack's label. */
    label: string;
    /** The button's words, which name the state ("All sources", "Filters (2)"). */
    buttonLabel: string;
    /** The button's accessible name when the words say less than the state ("Sources: all"). */
    buttonAriaLabel?: string;
    /** The panel's element id (aria-controls), and the label's id prefix. */
    id: string;
    testid: string;
    buttonTestid: string;
    popoverTestid: string;
    /** The panel's width in px. */
    width?: number;
    /** What takes focus when the panel opens: its first match. */
    focusSelector?: string;
    panel: Snippet;
    /** A muted note under a hairline at the panel's foot. */
    foot?: Snippet;
    /** The same choice below 768 px, in the expanded control stack. */
    stack: Snippet;
  } = $props();

  let wrap = $state<HTMLElement | null>(null);
  let trigger = $state<HTMLButtonElement | null>(null);
  let panelEl = $state<HTMLElement | null>(null);
  let open = $state(false);
  let fallback = $state(false);
  let handle: PopoverHandle | null = null;

  async function show(): Promise<void> {
    if (open || wrap === null || trigger === null || panelEl === null) return;
    open = true;
    fallback = !popoverSupported();
    if (!fallback) {
      const { openPopover } = await import('../../lib/ui/popover');
      if (!open) return;
      handle = openPopover({
        anchor: trigger,
        panel: panelEl,
        container: wrap,
        align: 'start',
        matchWidth: false,
        onDismiss: () => {
          handle = null;
          open = false;
        },
      });
    }
    panelEl.querySelector<HTMLElement>(focusSelector)?.focus();
  }

  function hide(refocus: boolean): void {
    if (!open) return;
    handle?.close();
    handle = null;
    open = false;
    if (refocus) trigger?.focus();
  }

  function onPanelKey(event: KeyboardEvent): void {
    if (event.key !== 'Escape') return;
    event.preventDefault();
    event.stopPropagation();
    hide(true);
  }

  function onPanelFocusOut(event: FocusEvent): void {
    const next = event.relatedTarget;
    if (next instanceof Node && wrap?.contains(next)) return;
    if (next !== null) hide(false);
  }
</script>

<div class="ui-cbar-group ui-filter-pop" data-kind="select" data-testid={testid}>
  <span class="ui-cbar-label text-label text-muted font-bold" id="{id}-label">{label}</span>
  <div class="ui-filter-pop-anchor" bind:this={wrap}>
    <button
      bind:this={trigger}
      type="button"
      class="{buttonClass('secondary', 'sm')} ui-filter-pop-button"
      aria-haspopup="dialog"
      aria-expanded={open}
      aria-controls={id}
      aria-label={buttonAriaLabel}
      data-testid={buttonTestid}
      onclick={() => (open ? hide(true) : void show())}
    >
      <span>{buttonLabel}</span>
      <Icon icon={ICON_CHEVRON_DOWN} size={14} class="ui-filter-pop-chev" />
    </button>
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    <div
      bind:this={panelEl}
      {id}
      class="ui-filter-pop-panel"
      style:width="{width}px"
      popover="manual"
      data-fallback={fallback && open ? 'true' : undefined}
      onkeydown={onPanelKey}
      onfocusout={onPanelFocusOut}
    >
      <Overlay kind="menu" role="dialog" {label} testid={popoverTestid} {foot}>{@render panel()}</Overlay>
    </div>
  </div>
  <div class="ui-filter-pop-stack" role="group" aria-labelledby="{id}-label">{@render stack()}</div>
</div>

<style>
  .ui-filter-pop {
    display: flex;
    flex-direction: column;
    gap: 4px;
    min-width: 0;
  }
  .ui-filter-pop-anchor {
    display: none;
    position: relative;
  }
  .ui-filter-pop-stack {
    display: flex;
    flex-direction: column;
    gap: 12px;
  }
  .ui-filter-pop-button {
    display: inline-flex;
    align-items: center;
    gap: 8px;
    height: 40px;
    padding: 0 12px;
    color: var(--color-strong);
    font-size: 14px;
    font-weight: 600;
    white-space: nowrap;
    transition: border-color 120ms;
  }
  .ui-filter-pop-button:hover,
  .ui-filter-pop-button[aria-expanded='true'] {
    border-color: var(--color-gold);
  }
  .ui-filter-pop-button :global(.ui-filter-pop-chev) {
    color: var(--color-muted);
    transition: rotate 120ms;
  }
  .ui-filter-pop-button[aria-expanded='true'] :global(.ui-filter-pop-chev) {
    rotate: 180deg;
  }
  /* The popover element only places the panel; Overlay draws the surface. */
  .ui-filter-pop-panel {
    position: fixed;
    inset: auto;
    margin: 0;
    padding: 0;
    border: 0;
    background: transparent;
    overflow: visible;
    color: var(--color-strong);
    display: none;
  }
  .ui-filter-pop-panel:popover-open {
    display: block;
  }
  .ui-filter-pop-panel[data-fallback='true'] {
    display: block;
    position: absolute;
    top: calc(100% + 4px);
    left: 0;
    z-index: 40;
  }
  @media (min-width: 768px) {
    .ui-filter-pop {
      flex-direction: row;
      align-items: center;
      gap: 8px;
    }
    /* The button names itself, so the group label stays on the phone. */
    .ui-filter-pop > .ui-cbar-label,
    .ui-filter-pop-stack {
      display: none;
    }
    .ui-filter-pop-anchor {
      display: block;
    }
  }
</style>
