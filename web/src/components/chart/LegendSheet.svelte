<!-- web/src/components/chart/LegendSheet.svelte -->
<!-- The phone sheet behind the legend's `+{n}` and the coarse-pointer Players control (controls
     and charts spec §B6.8): a modal bottom sheet on the account selector's sheet recipe. Players
     in two groups, Drawn and Not drawn, each a 52 px row with a Switch; an always-drawn row's
     Switch is on and disabled. Done, the scrim or Escape closes it and focus goes back. -->
<script lang="ts">
  import { matchesFilter } from '../../lib/chart/choice';
  import { chartCopy } from '../../lib/chart/copy';
  import type { MedianInfo } from '../../lib/chart/rate';
  import type { PlayerEntry } from '../../lib/chart/types';
  import { claimOverlay } from '../../lib/ui/overlay';
  import ClassIcon from '../report/ClassIcon.svelte';
  import Glyph from '../ui/Glyph.svelte';
  import Switch from '../ui/Switch.svelte';

  let {
    open,
    entries,
    medians = [],
    isDefault,
    formatValue,
    onToggle,
    onMedian,
    onReset,
    onClose,
  }: {
    open: boolean;
    entries: readonly PlayerEntry[];
    medians?: readonly MedianInfo[];
    isDefault: boolean;
    formatValue: (value: number) => string;
    onToggle: (id: string) => void;
    onMedian?: (id: string) => void;
    onReset: () => void;
    onClose: () => void;
  } = $props();

  let dialog = $state<HTMLDialogElement | null>(null);
  let query = $state('');
  let returnTo: HTMLElement | null = null;

  $effect(() => {
    const element = dialog;
    if (element === null || !open) return;
    returnTo = document.activeElement as HTMLElement | null;
    if (typeof element.showModal === 'function' && !element.open) element.showModal();
    const release = claimOverlay(onClose);
    return () => {
      release();
      if (element.open) element.close();
      query = '';
      returnTo?.focus();
    };
  });

  const visible = $derived(entries.filter((entry) => matchesFilter(entry.label, query)));
  const drawn = $derived(visible.filter((entry) => entry.drawn));
  const notDrawn = $derived(visible.filter((entry) => !entry.drawn));
  const onMedians = $derived(medians.filter((median) => median.on));
  const drawnCount = $derived(entries.filter((entry) => entry.drawn).length);
</script>

<dialog
  bind:this={dialog}
  class="sheet"
  aria-labelledby="legend-sheet-title"
  data-testid="legend-sheet"
  oncancel={(event) => {
    event.preventDefault();
    onClose();
  }}
  onclick={(event) => {
    if (event.target === dialog) onClose();
  }}
>
  {#if open}
    <div class="grab" aria-hidden="true"><span></span></div>
    <div class="head flex items-center gap-3">
      <h2 id="legend-sheet-title" class="title">{chartCopy.players}</h2>
      <span class="text-muted numeric mr-auto text-[13px]"
        >{chartCopy.sheetMeta(drawnCount, entries.length)}</span
      >
      <button type="button" class="action" disabled={isDefault} onclick={onReset}>{chartCopy.topN}</button>
      <button type="button" class="action" onclick={onClose} data-testid="legend-sheet-done"
        >{chartCopy.done}</button
      >
    </div>
    <div class="filter">
      <Glyph name="search" size={16} class="text-muted" />
      <input
        bind:value={query}
        type="search"
        placeholder={chartCopy.filterPlayers}
        aria-label={chartCopy.filterPlayers}
      />
    </div>
    <div class="list">
      <p class="group-head">{chartCopy.sheetDrawn}</p>
      {#each onMedians as median (median.id)}
        <div class="row">
          <span class="swatch" aria-hidden="true"></span>
          <span class="name flex-1">{median.label}</span>
          <Switch checked label={median.label} onchange={() => onMedian?.(median.id)} />
        </div>
      {/each}
      {#each drawn as entry (entry.id)}
        {@render row(entry)}
      {/each}
      {#if notDrawn.length > 0}
        <p class="group-head">{chartCopy.sheetNotDrawn}</p>
        {#each notDrawn as entry (entry.id)}
          {@render row(entry)}
        {/each}
      {/if}
    </div>
  {/if}
</dialog>

{#snippet row(entry: PlayerEntry)}
  <div class="row" data-testid="legend-sheet-row">
    <ClassIcon className={entry.className} size={24} />
    <span class="name min-w-0 flex-1 truncate" style:color={entry.colour}>{entry.label}</span>
    {#if entry.pinned}<Glyph name="pin" size={12} class="text-muted" />{/if}
    <span class="text-muted numeric text-[13px]">{formatValue(entry.figure)}</span>
    <Switch
      checked={entry.drawn}
      disabled={entry.pinned}
      label={entry.label}
      onchange={() => onToggle(entry.id)}
    />
  </div>
{/snippet}

<style>
  .sheet {
    position: fixed;
    inset: auto 0 0 0;
    width: 100%;
    max-width: 100%;
    max-height: 80vh;
    margin: 0;
    padding: 0 0 calc(14px + env(safe-area-inset-bottom));
    border: 1px solid var(--color-line-card);
    border-bottom: 0;
    border-radius: 12px 12px 0 0;
    background: var(--surface-card);
    color: var(--color-text);
    overflow: hidden;
    flex-direction: column;
  }
  .sheet[open] {
    display: flex;
  }
  .sheet::backdrop {
    background: rgb(0 0 0 / 0.65);
  }
  .grab {
    display: grid;
    place-items: center;
    height: 20px;
  }
  .grab span {
    width: 36px;
    height: 4px;
    border-radius: 2px;
    background: var(--color-line-warm);
  }
  .head {
    padding: 0 18px;
    height: 48px;
  }
  .title {
    font-family: var(--font-display, inherit);
    font-size: 18px;
    font-weight: 700;
    text-transform: uppercase;
    color: var(--color-strong);
  }
  .action {
    min-height: 44px;
    color: var(--color-gold);
    font-size: 13px;
    font-weight: 700;
    letter-spacing: 0.06em;
    text-transform: uppercase;
  }
  .action:disabled {
    opacity: 0.45;
  }
  .filter {
    display: flex;
    align-items: center;
    gap: 8px;
    margin: 4px 18px 10px;
    padding: 0 12px;
    height: 44px;
    border: 1px solid var(--color-line);
    border-radius: 6px;
    background: var(--color-bg);
  }
  .filter input {
    flex: 1;
    background: transparent;
    color: var(--color-strong);
    font-size: 14px;
    outline: none;
  }
  .list {
    flex: 1;
    overflow-y: auto;
    overscroll-behavior: contain;
  }
  .group-head {
    padding: 10px 18px 4px;
    font-size: 11px;
    font-weight: 700;
    letter-spacing: 0.14em;
    text-transform: uppercase;
    color: var(--color-muted);
    border-bottom: 1px solid var(--color-line-soft);
  }
  .row {
    display: flex;
    align-items: center;
    gap: 10px;
    height: 52px;
    padding: 0 18px;
    border-bottom: 1px solid var(--color-line-soft);
  }
  .name {
    font-size: 15px;
    font-weight: 600;
  }
  .swatch {
    width: 24px;
    height: 3px;
    background: repeating-linear-gradient(90deg, var(--color-text) 0 4px, transparent 4px 7px);
  }
</style>
