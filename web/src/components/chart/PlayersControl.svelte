<!-- web/src/components/chart/PlayersControl.svelte -->
<!-- The chart head's Players control (controls and charts spec §A4.5, §B5.3): who the chart draws.
     A searchable multi-select with a Compare group (the raid median, a class median), then Tanks,
     Healers and Damage; an always-drawn player is on and disabled. The presets row under the
     filter (PresetChips) picks Top 4 + bottom 4, Top 8 or Bottom 8; a toggle makes it Custom. It opens bottom-end in the top layer and is the page's one open overlay.
     On a coarse pointer it opens the legend sheet instead (onCoarse).

     TODO(dropdown lane): this is a temporary checkbox list. When the shared listbox lands on
     main, the trigger and list become `Select multiple searchable size="dense"` and this file
     keeps only the option mapping. -->
<script lang="ts">
  import { matchesFilter } from '../../lib/chart/choice';
  import { chartCopy } from '../../lib/chart/copy';
  import type { DrawPreset, PlayersChoice } from '../../lib/chart/drawn';
  import type { MedianInfo } from '../../lib/chart/rate';
  import type { PlayerEntry } from '../../lib/chart/types';
  import { claimOverlay } from '../../lib/ui/overlay';
  import ClassIcon from '../report/ClassIcon.svelte';
  import PresetChips from './PresetChips.svelte';
  import SelectChevron from '../ui/SelectChevron.svelte';

  let {
    id,
    entries,
    medians = [],
    choice,
    presets,
    formatValue,
    coarse = false,
    onToggle,
    onMedian,
    onPreset,
    onCoarse,
  }: {
    id: string;
    entries: readonly PlayerEntry[];
    medians?: readonly MedianInfo[];
    /** What is drawn now: a preset, or the reader's own set. */
    choice: PlayersChoice;
    /** The presets this chart offers: all three on the damage chart, Top 8 on threat. */
    presets: readonly DrawPreset[];
    formatValue: (value: number) => string;
    coarse?: boolean;
    onToggle: (id: string) => void;
    onMedian?: (id: string) => void;
    onPreset: (preset: DrawPreset) => void;
    onCoarse: () => void;
  } = $props();

  let open = $state(false);
  let query = $state('');
  let trigger = $state<HTMLButtonElement | null>(null);
  let panel = $state<HTMLDivElement | null>(null);
  let filter = $state<HTMLInputElement | null>(null);
  let place = $state({ top: 0, right: 0 });
  let release: (() => void) | null = null;

  const drawnCount = $derived(
    entries.filter((entry) => entry.drawn).length + medians.filter((m) => m.on).length,
  );
  const fieldText = $derived(
    chartCopy.playersField(entries.filter((entry) => entry.drawn).length, entries.length, choice),
  );
  const ROLE_ORDER = ['tank', 'healer', 'dps', 'other'] as const;
  const groups = $derived(
    ROLE_ORDER.map((role) => ({
      role,
      rows: entries.filter((entry) => entry.role === role && matchesFilter(entry.label, query)),
    })).filter((group) => group.rows.length > 0),
  );
  const shownMedians = $derived(medians.filter((median) => matchesFilter(median.label, query)));

  function close(): void {
    if (!open) return;
    open = false;
    query = '';
    panel?.hidePopover?.();
    release?.();
    release = null;
    trigger?.focus();
  }

  function show(): void {
    if (coarse) {
      onCoarse();
      return;
    }
    if (trigger === null) return;
    const rect = trigger.getBoundingClientRect();
    place = { top: rect.bottom + 4, right: Math.max(8, window.innerWidth - rect.right) };
    open = true;
    release = claimOverlay(() => {
      release = null;
      open = false;
      panel?.hidePopover?.();
    });
    queueMicrotask(() => {
      if (panel !== null && typeof panel.showPopover === 'function' && !panel.matches(':popover-open'))
        panel.showPopover();
      filter?.focus();
    });
  }

  $effect(() => {
    if (!open) return;
    const outside = (event: PointerEvent): void => {
      const target = event.target as Node;
      if (panel?.contains(target) || trigger?.contains(target)) return;
      close();
    };
    document.addEventListener('pointerdown', outside);
    return () => document.removeEventListener('pointerdown', outside);
  });

  function onKey(event: KeyboardEvent): void {
    if (event.key !== 'Escape') return;
    event.preventDefault();
    event.stopPropagation();
    if (query !== '') query = '';
    else close();
  }
</script>

<div class="players relative max-md:w-full">
  <button
    bind:this={trigger}
    type="button"
    {id}
    class="field"
    class:open
    aria-haspopup="listbox"
    aria-expanded={open}
    aria-controls="{id}-list"
    data-testid="players-control"
    onclick={() => (open ? close() : show())}
  >
    <span class="truncate">{fieldText}</span>
    <SelectChevron />
  </button>
  <div
    bind:this={panel}
    popover="manual"
    class="panel"
    class:shown={open}
    style:top="{place.top}px"
    style:right="{place.right}px"
    data-testid="players-list"
    role="presentation"
    onkeydown={onKey}
  >
    <div class="flex items-center gap-2 filter">
      <input
        bind:this={filter}
        bind:value={query}
        type="search"
        placeholder={chartCopy.filterPlayers}
        aria-label={chartCopy.filterPlayers}
        aria-controls="{id}-list"
        data-testid="players-filter"
      />
      <span class="text-muted shrink-0 text-[12px]" aria-live="polite">{entries.length} players</span>
    </div>
    <div class="preset-row"><PresetChips {presets} {choice} {onPreset} /></div>
    <div
      class="scroll"
      id="{id}-list"
      role="listbox"
      aria-multiselectable="true"
      aria-label={chartCopy.players}
      tabindex="-1"
    >
      {#if shownMedians.length > 0}
        <p class="group-head">{chartCopy.groups.compare}</p>
        {#each shownMedians as median (median.id)}
          <label class="row" role="option" aria-selected={median.on}>
            <input
              type="checkbox"
              checked={median.on}
              onchange={() => onMedian?.(median.id)}
              data-testid="players-median"
            />
            <span class="flex min-w-0 flex-col">
              <span class="pc-label">{median.label}</span>
              <span class="pc-description">{median.description}</span>
            </span>
          </label>
        {/each}
      {/if}
      {#each groups as group (group.role)}
        <p class="group-head">{chartCopy.groups[group.role]}</p>
        {#each group.rows as entry (entry.id)}
          <label class="row" role="option" aria-selected={entry.drawn} aria-disabled={entry.pinned}>
            <input
              type="checkbox"
              checked={entry.drawn}
              disabled={entry.pinned}
              onchange={() => onToggle(entry.id)}
              data-testid="players-option"
              data-id={entry.id}
            />
            <ClassIcon className={entry.className} size={20} />
            <span class="flex min-w-0 flex-1 flex-col">
              <span class="pc-label truncate" style:color={entry.colour}
                >{entry.label}{#if entry.you}<span class="you">{chartCopy.youList}</span>{/if}</span
              >
              {#if entry.pinned}<span class="pc-description">{chartCopy.alwaysDrawn}</span>{/if}
            </span>
            <span class="text-muted numeric text-[13px]" data-testid="players-option-figure"
              >{formatValue(entry.figure)}</span
            >
          </label>
        {/each}
      {/each}
    </div>
    <div class="foot flex items-center justify-end">
      <span class="text-muted numeric text-[12px]">{chartCopy.drawnCount(drawnCount)}</span>
    </div>
  </div>
</div>

<style>
  .field {
    display: inline-flex;
    align-items: center;
    justify-content: space-between;
    gap: 8px;
    width: 196px;
    height: 36px;
    padding: 0 10px;
    border: 1px solid var(--color-line-card);
    border-radius: var(--radius-control);
    background: rgba(255, 255, 255, 0.045);
    color: var(--color-strong);
    font-size: 13px;
    font-weight: 600;
  }
  .field:hover {
    border-color: rgba(229, 185, 85, 0.45);
    background: rgba(255, 255, 255, 0.06);
  }
  .field.open {
    border-color: var(--color-gold);
    box-shadow: var(--glow-gold);
  }
  .field:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  @media (max-width: 767px) {
    .field {
      width: 100%;
      height: 44px;
    }
  }
  .panel {
    text-align: left;
    text-transform: none;
    letter-spacing: normal;
    position: fixed;
    inset: auto;
    left: auto;
    margin: 0;
    width: 320px;
    max-height: min(420px, calc(100vh - 120px));
    padding: 0;
    flex-direction: column;
    border: 1px solid var(--color-line-card);
    border-radius: var(--radius-panel);
    background: var(--surface-card);
    color: var(--color-text);
    box-shadow:
      inset 0 1px 0 var(--card-highlight),
      var(--shadow-card);
    display: none;
  }
  .panel.shown {
    display: flex;
    z-index: 60;
  }
  .filter {
    padding: 6px 8px 8px;
    border-bottom: 1px solid var(--color-line-soft);
  }
  .filter input {
    flex: 1;
    height: 36px;
    padding: 0 10px;
    border: 1px solid var(--color-line);
    border-radius: 6px;
    background: var(--color-bg);
    color: var(--color-strong);
    font-size: 14px;
  }
  .filter input:focus {
    outline: none;
    border-color: var(--color-gold);
    box-shadow: 0 0 0 2px rgba(229, 185, 85, 0.35);
  }
  .scroll {
    flex: 1;
    overflow-y: auto;
    overscroll-behavior: contain;
    padding: 4px 0;
  }
  .group-head {
    padding: 10px 12px 4px;
    font-size: 11px;
    font-weight: 700;
    letter-spacing: 0.14em;
    text-transform: uppercase;
    color: var(--color-muted);
  }
  .row {
    display: flex;
    align-items: center;
    gap: 10px;
    min-height: 36px;
    padding: 4px 12px;
    cursor: pointer;
  }
  .row:hover {
    background: rgba(229, 185, 85, 0.1);
  }
  .row[aria-disabled='true'] {
    cursor: default;
  }
  .row input {
    width: 16px;
    height: 16px;
    accent-color: var(--color-gold);
  }
  .pc-label {
    font-size: 13px;
    font-weight: 600;
    color: var(--color-strong);
  }
  .pc-description {
    font-size: 12px;
    line-height: 16px;
    color: var(--color-muted);
  }
  .you {
    margin-left: 6px;
    padding: 0 4px;
    border: 1px solid rgba(229, 185, 85, 0.38);
    border-radius: var(--radius-pill);
    font-size: 11px;
    font-weight: 700;
    letter-spacing: 0.08em;
    text-transform: uppercase;
    color: var(--color-gold);
  }
  .foot {
    padding: 6px 8px;
    border-top: 1px solid var(--color-line-soft);
  }
  .preset-row {
    padding: 6px 8px;
    border-bottom: 1px solid var(--color-line-soft);
  }
</style>
