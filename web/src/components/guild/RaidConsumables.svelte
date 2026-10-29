<!-- web/src/components/guild/RaidConsumables.svelte -->
<!-- Consumables (attendance spec 3.4, officers): who, `Before` (the readiness pill as the board read it
     at the raid's cutoff) and `In the log` (chips for what was used, muted words for what was not),
     worst first, ten rows then `Show all n raiders`. A log with no buff events drops the In the log
     column and says so in one line under the sub-head. Rows stack below 768. -->
<script lang="ts">
  import { tick } from 'svelte';
  import { focusRowAt } from '../../lib/guild/focus-now';
  import type { ConsumableBefore, RaidConsumables } from '../../lib/guild/attendance-contract';
  import { attendanceCopy, usedChips } from '../../lib/guild/attendance-copy';
  import { PILL_WARN } from '../../lib/guild/pill-styles';
  import { OWN_ROW_WASH, ROW_DIVIDER } from '../../lib/guild/raid-report-styles';
  import { consumableOrder } from '../../lib/guild/report-view';
  import { visibleRows } from '../../lib/guild/night';
  import { classColorVar } from '../../lib/report/format';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import RaidBlock from './RaidBlock.svelte';
  import RaidShowAll from './RaidShowAll.svelte';

  let {
    consumables,
    ownKeys = [],
    expanded = $bindable(false),
  }: {
    consumables: RaidConsumables;
    ownKeys?: readonly string[];
    expanded?: boolean;
  } = $props();

  // The readiness board's own pill colours (GuildReadiness.svelte), by the state it read.
  const BEFORE_PILL: Record<ConsumableBefore, string> = {
    stocked: 'pill-community',
    short: PILL_WARN,
    unknown: 'pill-sample',
  };

  let list = $state<HTMLElement | null>(null);

  async function reveal(): Promise<void> {
    const from = shown.length;
    expanded = true;
    await tick();
    focusRowAt(list, from);
  }

  const ordered = $derived(consumableOrder(consumables.rows, consumables.log_has_buffs));
  const shown = $derived(visibleRows(ordered, expanded));
  const hasBuffs = $derived(consumables.log_has_buffs);
  /** Three columns with or without buff data, so `Before` holds its x (spec 3.4). */
  const COLUMNS = 'md:grid-cols-[minmax(180px,1fr)_120px_minmax(0,1.5fr)]';
  const LEAD = 'quiet-label md:hidden';
</script>

<RaidBlock heading={attendanceCopy.report.consumables} count={ordered.length} testid="raid-consumables">
  {#if !hasBuffs}
    <p class="text-muted text-[13px]" data-testid="raid-consumables-no-buffs">
      {attendanceCopy.consumables.noBuffData}
    </p>
  {/if}
  <div
    class="border-line-soft hidden min-h-6 items-end gap-x-4 border-b px-1 pb-1 md:grid {COLUMNS}"
    data-testid="raid-consumables-heads"
  >
    <span></span>
    <span class="quiet-label">{attendanceCopy.consumables.before}</span>
    {#if hasBuffs}<span class="quiet-label">{attendanceCopy.consumables.inTheLog}</span>{/if}
  </div>
  <ul bind:this={list}>
    {#each shown as row (row.character_key)}
      <li
        class="{ROW_DIVIDER} grid min-h-11 grid-cols-1 items-center gap-x-4 gap-y-1 px-1 py-2 {COLUMNS}"
        style={ownKeys.includes(row.character_key) ? `background:${OWN_ROW_WASH}` : undefined}
        tabindex="-1"
        data-testid="raid-consumable-row"
      >
        <span class="flex min-w-0 items-center gap-2">
          <ClassCrestRing characterClass={row.class} size={20} />
          <span
            class="truncate text-[13px] font-semibold"
            style={`color:${classColorVar(row.class)}`}
            title={row.name}>{row.name}</span
          >
        </span>
        <span class="flex items-center gap-2"
          ><span class={LEAD}>{attendanceCopy.consumables.before}</span><span
            class="pill {BEFORE_PILL[row.before]} justify-center md:w-full"
            >{attendanceCopy.consumables.state[row.before]}</span
          ></span
        >
        {#if hasBuffs}
          <span class="flex flex-wrap items-center gap-1.5">
            <span class={LEAD}>{attendanceCopy.consumables.inTheLog}</span>
            {#each usedChips(row) as chip (chip.text)}
              {#if chip.used}<span class="pill pill-sample">{chip.text}</span>{:else}<span
                  class="text-muted text-[12px]">{chip.text}</span
                >{/if}
            {/each}
          </span>
        {/if}
      </li>
    {/each}
  </ul>
  {#if !expanded && ordered.length > shown.length}
    <RaidShowAll total={ordered.length} onShow={() => void reveal()} testid="raid-consumables-show-all" />
  {/if}
</RaidBlock>
