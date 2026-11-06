<!-- web/src/components/ui/Paperdoll.svelte -->
<!-- The one paperdoll (web layout spec 2026-11-05 §4, owner decisions D3 and D10): two left-aligned
     columns in the character panel's order, icon first, each cell two lines. Line 1: the slot label
     and its level flag, the icon in its quality frame, the item name in its quality colour and its
     short stats, the gain at the right. Line 2: the full source and the next best (or the worn item).
     One cell is selected; below 1760 px the selected cell expands in place under itself, from 1760
     the detail sits in the side inspector. Four modes (pick, compare, worn, report) differ only in
     the cells the caller maps (lib/bis/doll-view.ts for BiS, lib/gear/doll-cells.ts for the worn
     and report dolls); the frame, cell, states and keyboard are one.
     It draws no box of its own: the page's Plate holds it. Every icon and name carries
     `data-bis-item`, so the page's one tooltip host opens the real item. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { onMount } from 'svelte';
  import { rarityClassFor } from '../../lib/planner/items';
  import { basisText } from '../../lib/ui/primitives';
  import {
    defaultSelection,
    dollColumns,
    dollKeyTarget,
    selectable,
    type DollCell,
    type DollMode,
  } from '../../lib/ui/paperdoll';
  import Figure from './Figure.svelte';
  import ItemIcon from './ItemIcon.svelte';
  import LoadError from './LoadError.svelte';

  let {
    cells,
    mode,
    build,
    selected = $bindable(null),
    onselect,
    detail,
    groups,
    loading = false,
    failed = false,
    failedMessage = 'We could not load the gear list.',
    onretry,
    label = 'Gear',
    testid = 'paperdoll',
  }: {
    cells: readonly DollCell[];
    mode: DollMode;
    build: string;
    /** The selected slot; null for none (below 1760 nothing is expanded until a tap). */
    selected?: string | null;
    onselect?: (slot: string | null) => void;
    /** The selected cell's detail: the side inspector (`side`) or the in-place expansion (`inline`). */
    detail?: Snippet<[DollCell, 'side' | 'inline']>;
    /** The "Where to go" view: groups of slots in place of the two columns. */
    groups?: readonly { key: string; title: string; count: string; slots: readonly string[] }[];
    loading?: boolean;
    failed?: boolean;
    failedMessage?: string;
    onretry?: () => void;
    label?: string;
    testid?: string;
  } = $props();

  /** Server-rendered as wide, so the side inspector's space is reserved; corrected on mount. */
  let wide = $state(true);
  let root = $state<HTMLElement | undefined>(undefined);

  onMount(() => {
    const query = window.matchMedia('(min-width: 1760px)');
    const apply = () => (wide = query.matches);
    apply();
    query.addEventListener('change', apply);
    return () => query.removeEventListener('change', apply);
  });

  const columns = $derived(dollColumns(cells));
  const bySlot = $derived(new Map(cells.map((cell) => [cell.slot, cell])));
  /** The shown slot: the chosen one, or from 1760 the default (§4.5: the side panel is never empty). */
  const shown = $derived(selected ?? (wide && detail !== undefined ? defaultSelection(cells) : null));
  const selectedCell = $derived(shown === null ? undefined : bySlot.get(shown));
  const focusSlot = $derived(selectedCell?.slot ?? cells.find(selectable)?.slot ?? null);

  function choose(cell: DollCell): void {
    if (!selectable(cell)) return;
    // A second tap on the open cell collapses it below 1760; the side inspector always shows one.
    const next = shown === cell.slot && !wide ? null : cell.slot;
    selected = next;
    onselect?.(next);
  }

  function onkeydown(event: KeyboardEvent, cell: DollCell): void {
    if (event.key === 'Enter' || event.key === ' ') {
      event.preventDefault();
      choose(cell);
      return;
    }
    const target = dollKeyTarget(cells, cell.slot, event.key);
    if (target === null) return;
    event.preventDefault();
    root?.querySelector<HTMLElement>(`[data-doll-slot="${CSS.escape(target)}"]`)?.focus();
    if (wide) {
      selected = target;
      onselect?.(target);
    }
  }

  const GAIN_ROWS = 17;
</script>

{#snippet cellView(cell: DollCell)}
  {@const isSelected = cell.slot === shown}
  {@const canSelect = selectable(cell)}
  <div
    class="ui-doll-cell"
    role="option"
    aria-selected={canSelect ? isSelected : undefined}
    aria-disabled={canSelect ? undefined : 'true'}
    tabindex={canSelect ? (cell.slot === focusSlot ? 0 : -1) : undefined}
    data-doll-slot={cell.slot}
    data-bis-item={cell.item && !cell.unknown ? cell.item.id : undefined}
    data-state={cell.item === null ? 'empty' : cell.unknown ? 'unknown' : 'item'}
    data-testid={`${testid}-slot-${cell.slot}`}
    onclick={() => choose(cell)}
    onkeydown={(event) => onkeydown(event, cell)}
  >
    <span class="ui-doll-label">
      <span class="text-muted text-[13px] leading-[18px] font-semibold">{cell.label}</span>
      {#if cell.flag}<span
          class="ui-doll-flag"
          data-tone={cell.flag.tone}
          data-testid={`${testid}-flag-${cell.slot}`}>{cell.flag.text}</span
        >{/if}
    </span>
    <span
      class="ui-doll-icon"
      data-bis-item={cell.item && !cell.unknown ? cell.item.id : undefined}
      data-covered={cell.item === null && cell.coveredBy ? 'true' : undefined}
    >
      <ItemIcon
        icon={cell.item?.icon ?? cell.coveredBy?.icon}
        quality={cell.item?.quality ?? cell.coveredBy?.quality}
        {build}
        size={40}
        badge={cell.badge}
        unknown={cell.unknown}
        deferred
      />
    </span>
    <span class="ui-doll-body">
      <span class="ui-doll-line1">
        {#if cell.item}
          <span
            class="text-[15px] leading-5 font-bold {cell.unknown
              ? 'text-muted'
              : rarityClassFor(cell.item.quality)}"
            data-bis-item={cell.unknown ? undefined : cell.item.id}
            data-testid={`item-hover-${cell.item.id}`}>{cell.item.name}</span
          >
          {#if cell.stats}<span class="ui-doll-stats text-muted text-[13px] leading-5" title={cell.stats}
              >{cell.stats}</span
            >{/if}
        {:else}
          <span class="text-muted text-[15px] leading-5">{cell.emptyTitle ?? 'Nothing better yet'}</span>
        {/if}
      </span>
      <span class="ui-doll-line2 text-muted text-[13px] leading-[18px]">
        {#if cell.item === null}
          {cell.empty ?? ''}
        {:else}
          {#if cell.source}<span
              class="ui-doll-src"
              class:text-rarity-uncommon={cell.sourceTone === 'enchant'}
              data-source-kind={cell.sourceKind}
              data-testid={`${testid}-line2-${cell.slot}`}>{cell.source}</span
            >{/if}
        {/if}
      </span>
    </span>
    <span class="ui-doll-gain">
      {#if cell.worn}
        <span data-bis-item={cell.worn.id} title={`You wear ${cell.worn.name}`}>
          <ItemIcon icon={cell.worn.icon} quality={cell.worn.quality} {build} size={28} />
        </span>
      {/if}
      {#if cell.gain}
        <span class="ui-doll-figure">
          <Figure
            value={cell.gain.value}
            unit={cell.gain.unit}
            caption={cell.gain.basis ? undefined : cell.gain.caption}
            tone={cell.gain.tone ?? 'strong'}
            size="sm"
            align="end"
          />
          {#if cell.gain.basis}
            <span
              class="ui-doll-basis"
              title={basisText(cell.gain.basis)}
              data-testid={`${testid}-basis-${cell.slot}`}
              >{cell.gain.basis.prefix}{#if cell.gain.basis.item}{' '}<span
                  class={rarityClassFor(cell.gain.basis.item.quality)}
                  data-bis-item={cell.gain.basis.item.id}>{cell.gain.basis.item.name}</span
                >{/if}</span
            >
          {/if}
        </span>
      {/if}
    </span>
  </div>
  {#if isSelected && !wide && detail}
    <div class="ui-doll-inline" data-testid={`${testid}-detail-${cell.slot}`}>
      {@render detail(cell, 'inline')}
    </div>
  {/if}
{/snippet}

{#snippet skeleton()}
  {#each Array.from({ length: GAIN_ROWS }, (_, index) => index) as index (index)}
    <div class="ui-doll-cell" aria-hidden="true">
      <span class="ui-doll-label"><span class="skeleton-block h-3 w-12"></span></span>
      <span class="skeleton-block h-10 w-10 rounded-[5px]"></span>
      <span class="ui-doll-body gap-2"
        ><span class="skeleton-block h-4 w-40"></span><span class="skeleton-block h-3 w-56"></span></span
      >
      <span class="skeleton-block h-5 w-10"></span>
    </div>
  {/each}
{/snippet}

<div
  bind:this={root}
  class="ui-doll"
  data-mode={mode}
  data-view={groups ? 'where' : 'slot'}
  data-wide={wide && detail !== undefined ? 'true' : undefined}
  data-testid={testid}
  aria-busy={loading ? 'true' : undefined}
>
  {#if failed}
    <div class="p-4"><LoadError message={failedMessage} onRetry={onretry} /></div>
  {:else if loading}
    <div class="ui-doll-columns" role="presentation">{@render skeleton()}</div>
  {:else}
    <div class="ui-doll-columns" role="listbox" aria-label={label}>
      {#if groups}
        {#each groups as group (group.key)}
          <div
            class="ui-doll-group"
            role="group"
            aria-labelledby={`${testid}-group-${group.key}-head`}
            data-testid={`${testid}-group-${group.key}`}
          >
            <div class="ui-doll-group-head" id={`${testid}-group-${group.key}-head`}>
              <span class="text-strong text-[18px] leading-6 font-bold">{group.title}</span>
              <span class="text-muted ml-auto text-[13px]">{group.count}</span>
            </div>
            {#each group.slots as slot (slot)}
              {@const cell = bySlot.get(slot)}
              {#if cell}{@render cellView(cell)}{/if}
            {/each}
          </div>
        {/each}
      {:else}
        <div class="ui-doll-column" data-column="left">
          {#each columns.left as cell (cell.slot)}{@render cellView(cell)}{/each}
        </div>
        <div class="ui-doll-column" data-column="right">
          {#each columns.right as cell (cell.slot)}{@render cellView(cell)}{/each}
        </div>
      {/if}
    </div>
    {#if wide && detail && selectedCell}
      <div class="ui-doll-side" data-testid={`${testid}-inspector`}>
        {@render detail(selectedCell, 'side')}
      </div>
    {/if}
  {/if}
</div>

<style>
  .ui-doll {
    display: grid;
    grid-template-columns: minmax(0, 1fr);
    min-width: 0;
  }
  .ui-doll-columns {
    display: grid;
    grid-template-columns: minmax(0, 1fr);
    min-width: 0;
  }
  .ui-doll-column,
  .ui-doll-group {
    min-width: 0;
  }
  .ui-doll-cell {
    display: grid;
    grid-template-columns: 64px 40px minmax(0, 1fr) auto;
    column-gap: 12px;
    align-items: start;
    min-height: var(--row-h);
    padding: 9px 16px;
    border-bottom: 1px solid var(--color-line-soft);
    cursor: pointer;
    position: relative;
  }
  .ui-doll-cell[aria-disabled='true'] {
    cursor: default;
  }
  .ui-doll-cell:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: -2px;
  }
  .ui-doll-cell:active:not([aria-disabled='true']) {
    background: var(--color-pressed);
  }
  .ui-doll-cell[aria-selected='true'] {
    background: var(--color-selected);
    box-shadow: inset 3px 0 0 var(--color-gold);
  }
  @media (hover: hover) and (pointer: fine) {
    .ui-doll-cell:not([aria-selected='true']):not([aria-disabled='true']):hover {
      background: var(--color-hover);
    }
  }
  .ui-doll-label {
    display: flex;
    flex-direction: column;
    padding-top: 1px;
    min-width: 0;
  }
  .ui-doll-flag {
    font-size: 12px;
    line-height: 16px;
    font-weight: 700;
    color: var(--color-muted);
  }
  .ui-doll-flag[data-tone='kill'] {
    color: var(--color-kill);
  }
  .ui-doll-flag[data-tone='gold'] {
    color: var(--color-gold);
  }
  /* A two-hander covering the off hand: greyed in the socket, as the game's pane draws it. */
  .ui-doll-icon[data-covered='true'] {
    filter: grayscale(1);
    opacity: 0.45;
  }
  .ui-doll-body {
    display: flex;
    flex-direction: column;
    gap: 1px;
    min-width: 0;
  }
  /* Line 1 holds one line from 768: the name never gives way; the short stats after it do (an
     ellipsis, every stat in the tooltip and the detail). A phone lets them wrap. */
  .ui-doll-line1 {
    display: flex;
    flex-wrap: wrap;
    align-items: baseline;
    column-gap: 10px;
    min-width: 0;
  }
  .ui-doll-stats {
    white-space: nowrap;
  }
  @media (min-width: 768px) {
    .ui-doll-line1 {
      flex-wrap: nowrap;
    }
    .ui-doll-line1 > :first-child {
      flex: none;
      max-width: 100%;
    }
    .ui-doll-stats {
      min-width: 0;
      overflow: hidden;
      text-overflow: ellipsis;
    }
  }
  .ui-doll-line2 {
    overflow-wrap: anywhere;
  }
  .ui-doll-gain {
    display: flex;
    align-items: center;
    gap: 10px;
    justify-self: end;
    min-width: 0;
  }
  /* The gain column (option C): the figure, then what it is against on one line that may
     ellipsize, its full words in the title. The cheapest width on the row. */
  .ui-doll-figure {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    width: var(--doll-gain-w, 144px);
    min-width: 0;
  }
  .ui-doll-basis {
    max-width: 100%;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    font-size: 12px;
    line-height: 16px;
    color: var(--color-muted);
    text-align: right;
  }
  .ui-doll-inline {
    border-bottom: 1px solid var(--color-line-soft);
  }
  .ui-doll-group-head {
    display: flex;
    align-items: baseline;
    gap: 12px;
    padding: 14px 16px 6px;
    margin: 0;
  }
  /* Below 400 px the slot label sits above the name (§4.3), 12 px caps. */
  @media (max-width: 479px) {
    .ui-doll-cell {
      grid-template-columns: 40px minmax(0, 1fr) auto;
      grid-template-rows: auto auto;
      padding-inline: 12px;
    }
    .ui-doll-label {
      grid-column: 2 / -1;
      flex-direction: row;
      gap: 8px;
      align-items: baseline;
    }
    .ui-doll-label > span:first-child {
      font-size: 12px;
      line-height: 16px;
      font-weight: 700;
      letter-spacing: var(--tracking-label);
      text-transform: uppercase;
    }
    .ui-doll-icon {
      grid-column: 1;
      grid-row: 1 / span 2;
    }
    .ui-doll-body {
      grid-column: 2;
    }
    .ui-doll-gain {
      grid-column: 3;
      grid-row: 2;
    }
    .ui-doll-figure {
      width: 112px;
    }
  }
  @media (min-width: 1024px) {
    .ui-doll-columns {
      grid-template-columns: minmax(0, 1fr) minmax(0, 1fr);
    }
    .ui-doll-column + .ui-doll-column {
      border-left: 1px solid var(--color-line-soft);
    }
    .ui-doll[data-view='where'] .ui-doll-columns {
      display: block;
      columns: 2;
      column-gap: 0;
    }
    .ui-doll-group {
      break-inside: avoid;
    }
  }
  @media (min-width: 1760px) {
    .ui-doll[data-wide='true'] {
      grid-template-columns: minmax(0, 1fr) var(--inspector-w);
    }
    .ui-doll-side {
      border-left: 1px solid var(--color-line);
      min-width: 0;
    }
  }
  @media (max-width: 1759px) {
    .ui-doll-side {
      display: none;
    }
  }
</style>
