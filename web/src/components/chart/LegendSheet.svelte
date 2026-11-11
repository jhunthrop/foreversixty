<!-- web/src/components/chart/LegendSheet.svelte -->
<!-- The phone sheet behind the legend's `+{n}` and the coarse-pointer Players control (controls
     and charts spec §B6.8): a modal bottom sheet on the account selector's sheet recipe. Players
     in two groups, Drawn and Not drawn, each a 52 px row with a Switch; an always-drawn row's
     Switch is on and disabled. Done, the scrim or Escape closes it and focus goes back. -->
<script lang="ts">
  import { matchesFilter } from '../../lib/chart/choice';
  import { chartCopy } from '../../lib/chart/copy';
  import type { DrawPreset, PlayersChoice } from '../../lib/chart/drawn';
  import type { MedianInfo } from '../../lib/chart/rate';
  import type { PlayerEntry } from '../../lib/chart/types';
  import { claimOverlay } from '../../lib/ui/overlay';
  import Crest from '../ui/Crest.svelte';
  import PresetChips from './PresetChips.svelte';
  import Button from '../ui/Button.svelte';
  import Field from '../ui/Field.svelte';
  import Icon from '../ui/Icon.svelte';
  import { ICON_PIN } from '../../lib/ui/ui-icons';
  import Overlay from '../ui/Overlay.svelte';
  import Row from '../ui/Row.svelte';
  import Switch from '../ui/Switch.svelte';

  let {
    open,
    entries,
    medians = [],
    choice,
    presets,
    formatValue,
    onToggle,
    onMedian,
    onPreset,
    onClose,
  }: {
    open: boolean;
    entries: readonly PlayerEntry[];
    medians?: readonly MedianInfo[];
    /** What is drawn now: a preset, or the reader's own set. */
    choice: PlayersChoice;
    /** The presets this chart offers: all three on the damage chart, Top 8 on threat. */
    presets: readonly DrawPreset[];
    formatValue: (value: number) => string;
    onToggle: (id: string) => void;
    onMedian?: (id: string) => void;
    onPreset: (preset: DrawPreset) => void;
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
    <Overlay kind="sheet" role="none">
      <div class="panel">
        <div class="grid h-5 place-items-center" aria-hidden="true">
          <span class="h-1 w-9 rounded-[2px] bg-white/[.18]"></span>
        </div>
        <div class="flex h-12 items-center gap-3 px-[18px]">
          <h2 id="legend-sheet-title" class="text-strong text-[18px] font-bold">{chartCopy.players}</h2>
          <span class="text-muted numeric mr-auto text-[13px]"
            >{chartCopy.sheetMeta(drawnCount, entries.length)}</span
          >
          <Button tone="ghost" onclick={onClose} testid="legend-sheet-done">{chartCopy.done}</Button>
        </div>
        <div class="px-[18px] pb-2"><PresetChips {presets} {choice} {onPreset} testid="sheet-presets" /></div>
        <div class="px-[18px] pt-1 pb-2.5">
          <Field
            id="legend-sheet-filter"
            type="search"
            label={chartCopy.filterPlayers}
            labelHidden
            placeholder={chartCopy.filterPlayers}
            bind:value={query}
          />
        </div>
        <div class="list px-[18px]">
          <p class="text-muted text-small pt-2.5 pb-1 font-semibold">{chartCopy.sheetDrawn}</p>
          {#each onMedians as median (median.id)}
            <Row title={median.label} density="slim" open>
              {#snippet lead()}<span class="swatch" aria-hidden="true"></span>{/snippet}
              {#snippet trail()}<Switch
                  checked
                  label={median.label}
                  onchange={() => onMedian?.(median.id)}
                />{/snippet}
            </Row>
          {/each}
          {#each drawn as entry (entry.id)}
            {@render row(entry)}
          {/each}
          {#if notDrawn.length > 0}
            <p class="text-muted text-small pt-4 pb-1 font-semibold">{chartCopy.sheetNotDrawn}</p>
            {#each notDrawn as entry (entry.id)}
              {@render row(entry)}
            {/each}
          {/if}
        </div>
      </div>
    </Overlay>
  {/if}
</dialog>

{#snippet row(entry: PlayerEntry)}
  <Row title={entry.label} density="slim" open testid="legend-sheet-row">
    {#snippet lead()}<Crest cls={entry.className} size={24} />{/snippet}
    {#snippet titleContent()}<span class="block truncate" style:color={entry.colour}>{entry.label}</span
      >{/snippet}
    {#snippet trail()}
      {#if entry.pinned}<Icon icon={ICON_PIN} size={12} class="text-muted" />{/if}
      <span class="text-muted numeric text-[13px]">{formatValue(entry.figure)}</span>
      <Switch
        checked={entry.drawn}
        disabled={entry.pinned}
        label={entry.label}
        onchange={() => onToggle(entry.id)}
      />
    {/snippet}
  </Row>
{/snippet}

<style>
  .sheet {
    position: fixed;
    inset: auto 0 0 0;
    width: 100%;
    max-width: 100%;
    max-height: 80vh;
    margin: 0;
    padding: 0;
    border: 0;
    background: transparent;
    color: var(--color-text);
    overflow: hidden;
  }
  .sheet[open] {
    display: flex;
    flex-direction: column;
  }
  .sheet::backdrop {
    background: color-mix(in srgb, black 65%, transparent);
  }
  .sheet > :global(.ui-overlay) {
    display: flex;
    flex-direction: column;
    min-height: 0;
    max-height: 80vh;
  }
  .panel {
    display: flex;
    flex-direction: column;
    min-height: 0;
    padding-bottom: calc(14px + env(safe-area-inset-bottom));
  }
  .list {
    flex: 1;
    overflow-y: auto;
    overscroll-behavior: contain;
  }
  .swatch {
    display: block;
    width: 24px;
    height: 3px;
    background: repeating-linear-gradient(90deg, var(--color-text) 0 4px, transparent 4px 7px);
  }
</style>
