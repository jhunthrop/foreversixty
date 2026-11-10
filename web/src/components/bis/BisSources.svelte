<!-- web/src/components/bis/BisSources.svelte -->
<!-- The control bar's Sources group (web layout spec 2026-11-05 §5.11). From 768 px a compact
     button ("All sources", "Sources (6)": the kinds that are on) opens a popover on the shared
     overlay surface with the checklist, each kind with its count (slots with a ranked option from
     that source), and one footer line; Escape or a click outside closes it and focus returns to
     the button. The button names itself, so the group label shows only in the phone stack. Below 768 px the same choice is a row of toggle chips in the expanded control
     stack. Turning a kind off re-picks; nothing hides. -->
<script lang="ts">
  import Overlay from '../ui/Overlay.svelte';
  import { button as buttonClass } from '../../lib/ui/button';
  import { popoverSupported } from '../../lib/ui/listbox-model';
  import type { PopoverHandle } from '../../lib/ui/popover';
  import { SOURCE_FILTER_KINDS } from '../../lib/bis/source-filter';
  import { sourceKindLabel, sourcesButtonLabel, sourcesFieldLabel } from '../../lib/bis/source-choice';
  import type { SourceFilterKind } from '../../lib/bis/types';

  let {
    on,
    counts,
    band,
    onchange,
  }: {
    on: readonly SourceFilterKind[];
    counts: Readonly<Record<SourceFilterKind, number>>;
    band: number;
    onchange: (next: SourceFilterKind[]) => void;
  } = $props();

  const FOOT =
    'Counts are slots with a ranked option from that source. Turning one off re-picks each slot; nothing is hidden.';
  const PANEL_ID = 'bis-sources-panel';

  let wrap = $state<HTMLElement | null>(null);
  let button = $state<HTMLButtonElement | null>(null);
  let panel = $state<HTMLElement | null>(null);
  let open = $state(false);
  let handle: PopoverHandle | null = null;
  /** No Popover API (an old browser): the panel hangs under the button in the page instead. */
  let fallback = $state(false);

  function ordered(kinds: Iterable<string>): SourceFilterKind[] {
    const set = new Set(kinds);
    return SOURCE_FILTER_KINDS.filter((kind) => set.has(kind));
  }

  function toggle(kind: SourceFilterKind): void {
    onchange(on.includes(kind) ? on.filter((other) => other !== kind) : ordered([...on, kind]));
  }

  async function show(): Promise<void> {
    if (open || wrap === null || button === null || panel === null) return;
    open = true;
    fallback = !popoverSupported();
    if (!fallback) {
      const { openPopover } = await import('../../lib/ui/popover');
      if (!open) return;
      handle = openPopover({
        anchor: button,
        panel,
        container: wrap,
        align: 'start',
        matchWidth: false,
        onDismiss: () => {
          handle = null;
          open = false;
        },
      });
    }
    panel.querySelector<HTMLElement>('input')?.focus();
  }

  function hide(refocus: boolean): void {
    if (!open) return;
    handle?.close();
    handle = null;
    open = false;
    if (refocus) button?.focus();
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

<div class="ui-cbar-group bis-sources" data-kind="select" data-testid="bis-sources">
  <span class="ui-cbar-label text-label text-muted font-bold" id="bis-sources-label">Sources</span>
  <div class="bis-sources-pop" bind:this={wrap}>
    <button
      bind:this={button}
      type="button"
      class="{buttonClass('secondary', 'sm', true)} bis-sources-button"
      aria-haspopup="dialog"
      aria-expanded={open}
      aria-controls={PANEL_ID}
      aria-label={`Sources: ${sourcesFieldLabel(on, band)}`}
      data-testid="bis-sources-button"
      onclick={() => (open ? hide(true) : void show())}
    >
      <span>{sourcesButtonLabel(on)}</span>
      <svg
        class="bis-sources-chev"
        width="14"
        height="14"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2.5"
        stroke-linecap="round"
        stroke-linejoin="round"
        aria-hidden="true"><path d="M6 9l6 6 6-6" /></svg
      >
    </button>
    <div
      bind:this={panel}
      id={PANEL_ID}
      class="bis-sources-panel"
      popover="manual"
      data-fallback={fallback && open ? 'true' : undefined}
      onkeydown={onPanelKey}
      onfocusout={onPanelFocusOut}
    >
      <Overlay kind="menu" role="dialog" label="Sources" testid="bis-sources-popover">
        <fieldset class="bis-sources-list">
          <legend class="sr-only">Sources</legend>
          {#each SOURCE_FILTER_KINDS as kind (kind)}
            <label class="bis-sources-row" data-testid={`bis-source-check-${kind}`}>
              <input type="checkbox" checked={on.includes(kind)} onchange={() => toggle(kind)} />
              <span class="bis-sources-name">{sourceKindLabel(kind)}</span>
              <span class="bis-sources-count">{counts[kind]}</span>
            </label>
          {/each}
        </fieldset>
        {#snippet foot()}{FOOT}{/snippet}
      </Overlay>
    </div>
  </div>
  <div class="bis-sources-chips" role="group" aria-labelledby="bis-sources-label">
    {#each SOURCE_FILTER_KINDS as kind (kind)}
      <button
        type="button"
        class="bis-chip"
        aria-pressed={on.includes(kind)}
        data-testid={`bis-source-chip-${kind}`}
        onclick={() => toggle(kind)}
      >
        {sourceKindLabel(kind)} <span class="bis-chip-count">{counts[kind]}</span>
      </button>
    {/each}
    <p class="bis-sources-foot text-muted">{FOOT}</p>
  </div>
</div>

<style>
  .bis-sources {
    display: flex;
    flex-direction: column;
    gap: 4px;
    min-width: 0;
  }
  .bis-sources-chips {
    display: flex;
    flex-wrap: wrap;
    gap: 6px;
  }
  .bis-sources-pop {
    display: none;
    position: relative;
  }
  .bis-sources-button {
    display: inline-flex;
    align-items: center;
    gap: 8px;
    height: 36px;
    padding: 0 10px;
    color: var(--color-strong);
    font-size: 13px;
    font-weight: 600;
    white-space: nowrap;
    transition: border-color 120ms;
  }
  .bis-sources-button:hover,
  .bis-sources-button[aria-expanded='true'] {
    border-color: var(--color-gold);
  }
  .bis-sources-button:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .bis-sources-chev {
    color: var(--color-muted);
    transition: rotate 120ms;
  }
  .bis-sources-button[aria-expanded='true'] .bis-sources-chev {
    rotate: 180deg;
  }
  /* The popover element only places the panel; Overlay draws the surface. */
  .bis-sources-panel {
    position: fixed;
    inset: auto;
    margin: 0;
    padding: 0;
    border: 0;
    background: transparent;
    overflow: visible;
    width: 280px;
    color: var(--color-strong);
    display: none;
  }
  .bis-sources-panel:popover-open {
    display: block;
  }
  .bis-sources-panel[data-fallback='true'] {
    display: block;
    position: absolute;
    top: calc(100% + 4px);
    left: 0;
    z-index: 40;
  }
  .bis-sources-list {
    margin: 0;
    padding: 6px 0;
    border: 0;
  }
  .bis-sources-row {
    display: flex;
    align-items: center;
    gap: 10px;
    min-height: 36px;
    padding: 0 14px;
    font-size: 14px;
    font-weight: 600;
    cursor: pointer;
  }
  .bis-sources-row:hover {
    background: var(--color-hover);
  }
  .bis-sources-row input {
    width: 16px;
    height: 16px;
    accent-color: var(--color-gold);
  }
  .bis-sources-name {
    flex: 1;
  }
  .bis-sources-count {
    color: var(--color-muted);
    font-variant-numeric: tabular-nums;
    font-size: 13px;
  }
  .bis-chip {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    min-height: 32px;
    padding: 0 10px;
    border: 1px solid color-mix(in srgb, currentColor 40%, transparent);
    border-radius: var(--radius-control, 6px);
    color: var(--color-muted);
    font-size: 13px;
    font-weight: 600;
  }
  .bis-chip[aria-pressed='true'] {
    border-color: var(--color-selected-mark);
    background: var(--color-selected);
    color: var(--color-strong);
  }
  .bis-chip:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .bis-chip-count {
    font-variant-numeric: tabular-nums;
    opacity: 0.8;
  }
  .bis-sources-foot {
    flex-basis: 100%;
    font-size: 12px;
    line-height: 16px;
  }
  @media (min-width: 768px) {
    .bis-sources {
      flex-direction: row;
      align-items: center;
      gap: 8px;
    }
    /* The button names itself ("All sources", "Sources (6)"), so the group label stays on the phone. */
    .bis-sources > .ui-cbar-label,
    .bis-sources-chips {
      display: none;
    }
    .bis-sources-pop {
      display: block;
    }
  }
</style>
