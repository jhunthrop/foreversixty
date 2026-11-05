<!-- web/src/components/report/AuraBars.svelte -->
<!-- AuraTable's `bars` variant: one column of the saved run's Buffs and debuffs card (spec
     2026-11-04-sim-results 3.10). Uptime bars for what comes and goes, then an icon-and-name row for
     what stayed up, then what did nothing for this spec, quiet. -->
<script lang="ts">
  import type { AuraColumn, AuraEntry } from '../../lib/sim/aura-card';
  import type { CardTables } from '../../lib/sim/card-tables.svelte';
  import { aurasCardCopy as copy } from '../../lib/sim/input-cards-copy';
  import { oneDecimal } from '../../lib/sim/result-copy';
  import ActionMark from '../sim/ActionMark.svelte';

  let {
    column,
    heading,
    side,
    tables,
    note,
  }: {
    column: AuraColumn;
    heading: string;
    side: 'player' | 'target';
    tables: CardTables;
    /** A line under the heading (an older run's target column). */
    note?: string;
  } = $props();
</script>

{#snippet mark(entry: AuraEntry, size: number)}
  <ActionMark action={entry.actionKey} {tables} {size} fallbackIcon={entry.buffIcon}>{entry.name}</ActionMark>
{/snippet}

{#snippet inline(entry: AuraEntry)}
  <li class="flex items-center gap-1.5 text-[14px]" data-testid="aura-inline">
    {@render mark(entry, 20)}
    {#if entry.stacks > 1}<span class="text-muted numeric text-[13px]">{copy.stacks(entry.stacks)}</span>{/if}
  </li>
{/snippet}

<div class="flex min-w-0 flex-col gap-3 px-5 py-4" data-testid={`aura-column-${side}`}>
  <div class="quiet-label border-line-soft flex justify-between border-b pb-2">
    <span>{heading}</span>
    {#if column.bars.length > 0}<span>{copy.comeAndGo(column.bars.length)}</span>{/if}
  </div>
  {#if note !== undefined}<p class="text-muted text-[13px]" data-testid="aura-column-note">{note}</p>{/if}
  {#if column.bars.length > 0}
    <ul class="-mt-3 flex flex-col">
      {#each column.bars as entry (entry.key)}
        <li
          class="border-line-soft grid min-h-10 grid-cols-[minmax(0,1fr)_56px_52px_36px] items-center gap-2 border-b text-[14px] md:grid-cols-[minmax(0,1fr)_96px_60px_44px] md:gap-3"
          data-testid="aura-bar-row"
        >
          {@render mark(entry, 24)}
          <span class="bg-line-soft block h-1.5 rounded-full">
            <span
              class={`block h-full rounded-full ${side === 'player' ? 'bg-text/55' : 'bg-gold-deep'}`}
              style:width={`${entry.uptimePct}%`}
            ></span>
          </span>
          <span class="text-strong numeric text-right font-semibold">{oneDecimal(entry.uptimePct)}%</span>
          <span class="text-muted numeric text-right text-[13px]">{copy.applied(entry.applications)}</span>
        </li>
      {/each}
    </ul>
  {/if}
  {#if column.upAllFight.length > 0}
    <div class="flex flex-col gap-2" data-testid={`aura-up-all-fight-${side}`}>
      <span class="text-muted text-[13px] font-semibold">{copy.upAllFight}</span>
      <ul class="flex flex-wrap gap-x-4 gap-y-2">
        {#each column.upAllFight as entry (entry.key)}{@render inline(entry)}{/each}
      </ul>
    </div>
  {/if}
  {#if column.noEffect.length > 0}
    <div class="flex flex-col gap-2" data-testid={`aura-no-effect-${side}`}>
      <span class="text-muted text-[13px] font-semibold">{copy.noEffect}</span>
      <ul class="flex flex-wrap gap-x-4 gap-y-2 opacity-[.62]">
        {#each column.noEffect as entry (entry.key)}{@render inline(entry)}{/each}
      </ul>
    </div>
  {/if}
</div>
