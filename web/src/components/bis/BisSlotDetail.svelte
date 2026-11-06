<!-- web/src/components/bis/BisSlotDetail.svelte -->
<!-- The selected slot's detail (web layout spec 2026-11-05 §4.5, §5.5): the side inspector from
     1760 px, the in-place expansion below. Everything the two-line row leaves out lives here:
     stats, enchant and effect, the worn item, where to get it, every runner-up with its gap to the pick,
     the trinket ranking and the simulator link. All of it is in the page's data (R2). -->
<script lang="ts">
  import Inspector, { type InspectorSection } from '../ui/Inspector.svelte';
  import ItemIcon from '../ui/ItemIcon.svelte';
  import { rarityClassFor } from '../../lib/planner/items';
  import { isTie, shortStats } from '../../lib/bis/doll-view';
  import { signedFigure } from '../../lib/ui/primitives';
  import type { BisCandidate, BisItem, BisRowModel, GainUnit } from '../../lib/bis/page-model';
  import type { DollCell } from '../../lib/ui/paperdoll';

  let {
    cell,
    row,
    items,
    unit,
    build,
    simHref,
    placement,
  }: {
    cell: DollCell;
    row: BisRowModel | undefined;
    items: Readonly<Record<number, BisItem>>;
    unit: GainUnit;
    build: string;
    simHref?: string;
    placement: 'side' | 'inline';
  } = $props();

  const shownId = $derived(cell.item?.id);
  const shown = $derived(shownId === undefined ? undefined : items[shownId]);
  const source = $derived(shownId === undefined ? undefined : row?.sources[shownId]?.line);
  const pickRows = $derived<BisCandidate[]>(
    row?.id === undefined ? [] : [{ id: row.id, delta: 0 }, ...row.alts],
  );

  function gapText(delta: number): string {
    return isTie(delta) ? `same ${unit}` : `${signedFigure(delta)} ${unit}`;
  }

  const sections = $derived.by(() => {
    const list: InspectorSection[] = [];
    if (cell.note && cell.note.label !== 'You wear')
      list.push({ label: cell.note.label, body: noteLine, testid: 'bis-detail-repick' });
    if (cell.worn) list.push({ label: 'You wear', body: worn, testid: 'bis-detail-worn' });
    if (source) list.push({ label: 'Where to get it', body: where, testid: 'bis-detail-source' });
    if (pickRows.length > 1)
      list.push({
        label: `Runners-up · ${unit} against the pick`,
        body: runners,
        testid: 'bis-detail-runners',
      });
    if (row?.ranking && row.ranking.length > 0)
      list.push({
        label: `Trinkets ranked · ${unit} against the pick`,
        body: ranking,
        testid: 'bis-detail-ranking',
      });
    return list;
  });
</script>

{#snippet candidate(entry: BisCandidate, tag: string)}
  {@const item = items[entry.id]}
  <li class="bis-detail-cand">
    <span data-bis-item={entry.id}>
      <ItemIcon icon={item?.icon} quality={item?.quality ?? 1} {build} size={28} />
    </span>
    <span class={['bis-detail-name', rarityClassFor(item?.quality ?? 1)]} data-bis-item={entry.id}
      >{item?.name ?? `Item ${entry.id}`}</span
    >
    <span class="bis-detail-gap">{tag}</span>
    {#if row?.sources[entry.id]}<span class="bis-detail-src text-muted">{row.sources[entry.id]!.line}</span
      >{/if}
  </li>
{/snippet}

{#snippet worn()}
  {#if cell.worn}
    <ul class="bis-detail-list">
      {@render candidate({ id: cell.worn.id, delta: 0 }, cell.gain?.value ?? '')}
    </ul>
  {/if}
  {#if cell.note?.label === 'You wear'}{@render noteLine()}{/if}
{/snippet}

{#snippet noteLine()}
  {#if cell.note}
    <p class="bis-detail-note text-muted" data-testid="bis-detail-note">
      {cell.note.lead}
      {#if cell.note.item}<span
          class={rarityClassFor(cell.note.item.quality)}
          data-bis-item={cell.note.item.id}>{cell.note.item.name}</span
        >{/if}
      {#if cell.note.tail}{cell.note.tail}{/if}
    </p>
  {/if}
{/snippet}

{#snippet where()}
  <p class="bis-detail-text">{source}</p>
{/snippet}

{#snippet runners()}
  <ul class="bis-detail-list">
    {#each pickRows as entry, index (entry.id)}
      {@render candidate(entry, index === 0 ? 'Pick' : gapText(entry.delta))}
    {/each}
  </ul>
{/snippet}

{#snippet ranking()}
  <ul class="bis-detail-list">
    {#each row?.ranking ?? [] as entry (entry.id)}{@render candidate(entry, gapText(entry.delta))}{/each}
  </ul>
{/snippet}

{#snippet head()}
  {#if cell.item}
    <p class={['bis-detail-title', rarityClassFor(cell.item.quality)]} data-bis-item={cell.item.id}>
      {cell.item.name}
    </p>
    <p class="bis-detail-text text-muted">
      {[shown?.ilvl ? `Item level ${shown.ilvl}` : undefined, shortStats(shown?.stats, 6)]
        .filter(Boolean)
        .join(' · ')}
    </p>
    {#if row?.enchant}<p class="bis-detail-text" data-testid="bis-detail-enchant">
        <span class="text-muted">Enchant</span> <span class="text-rarity-uncommon">{row.enchant}</span>
      </p>{/if}
    {#if shown?.effect}<p class="bis-detail-text">{shown.effect}</p>{/if}
    {#if row?.setBonus}<p class="bis-detail-text text-muted" data-testid="bis-row-set-bonus">
        {row.setBonus}
      </p>{/if}
    {#if simHref}<p class="bis-detail-text">
        <a href={simHref} class="bis-detail-link">Compare in the simulator ›</a>
      </p>{/if}
  {:else if cell.empty}
    <p class="bis-detail-text">{cell.empty}</p>
  {/if}
{/snippet}

<Inspector selection={cell.label} {head} {sections} {placement} testid={`bis-detail-${placement}`} />

<style>
  .bis-detail-title {
    font-size: 16px;
    font-weight: 700;
  }
  .bis-detail-text {
    font-size: 13px;
    line-height: 20px;
  }
  .bis-detail-link {
    color: var(--color-gold);
    font-weight: 600;
  }
  .bis-detail-list {
    display: flex;
    flex-direction: column;
    gap: 6px;
  }
  .bis-detail-cand {
    display: grid;
    grid-template-columns: 28px minmax(0, 1fr) auto;
    column-gap: 8px;
    align-items: center;
    font-size: 14px;
  }
  .bis-detail-gap {
    font-variant-numeric: tabular-nums;
    color: var(--color-muted);
    font-size: 13px;
  }
  .bis-detail-src {
    grid-column: 2 / 4;
    font-size: 12px;
    line-height: 16px;
  }
</style>
