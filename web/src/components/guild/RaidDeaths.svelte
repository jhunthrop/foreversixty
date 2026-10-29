<!-- web/src/components/guild/RaidDeaths.svelte -->
<!-- Deaths (attendance spec 3.4, officers): one 40 px row per raider who died, most first; an 8 px
     muted bar (never `--wipe`, R-A3), the count, and the wipe/kill split. The first ten show, then
     `Show all n raiders`. None: `Nobody died.` -->
<script lang="ts">
  import { tick } from 'svelte';
  import { focusRowAt } from '../../lib/guild/focus-now';
  import type { RaidDeathRow } from '../../lib/guild/attendance-contract';
  import { attendanceCopy, deathSplit } from '../../lib/guild/attendance-copy';
  import { BAR_FILL, BAR_TRACK, OWN_ROW_WASH, ROW_DIVIDER } from '../../lib/guild/raid-report-styles';
  import { deathCount, deathOrder } from '../../lib/guild/report-view';
  import { visibleRows } from '../../lib/guild/night';
  import { classColorVar } from '../../lib/report/format';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import RaidBlock from './RaidBlock.svelte';
  import RaidShowAll from './RaidShowAll.svelte';

  let {
    rows,
    ownKeys = [],
    expanded = $bindable(false),
  }: {
    rows: readonly RaidDeathRow[];
    /** The viewer's own characters (the own-row wash). */
    ownKeys?: readonly string[];
    /** Every raider shows; a test seeds it. */
    expanded?: boolean;
  } = $props();

  let list = $state<HTMLElement | null>(null);

  async function reveal(): Promise<void> {
    const from = shown.length;
    expanded = true;
    await tick();
    focusRowAt(list, from);
  }

  const ordered = $derived(deathOrder(rows));
  const most = $derived(ordered.length === 0 ? 1 : deathCount(ordered[0]));
  const shown = $derived(visibleRows(ordered, expanded));
  /** The sub-head counts deaths (the night line's figure), `Show all` counts raiders. */
  const total = $derived(ordered.reduce((sum, row) => sum + deathCount(row), 0));
</script>

<RaidBlock heading={attendanceCopy.report.deaths} count={total} testid="raid-deaths">
  {#if ordered.length === 0}
    <p class="text-muted text-[14px]" data-testid="raid-deaths-none">{attendanceCopy.deaths.none}</p>
  {:else}
    <ul bind:this={list}>
      {#each shown as row (row.character_key)}
        <!-- Phone: line 1 crest, the whole name and the count; line 2 the bar and the split.
             From 768 one 40 px row: crest, name, bar, count, split. -->
        <li
          class="{ROW_DIVIDER} grid min-h-[52px] grid-cols-[20px_minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 px-1 py-1.5 md:flex md:min-h-10 md:py-0"
          style={ownKeys.includes(row.character_key) ? `background:${OWN_ROW_WASH}` : undefined}
          tabindex="-1"
          data-testid="raid-death-row"
        >
          <ClassCrestRing characterClass={row.class} size={20} />
          <span
            class="text-[13px] font-semibold break-words md:w-28 md:truncate"
            style={`color:${classColorVar(row.class)}`}
            title={row.name}>{row.name}</span
          >
          <span class="text-strong numeric text-[13px] md:order-2">{deathCount(row)}</span>
          <span class="col-span-3 flex items-center gap-3 pl-8 md:contents">
            <span
              class="{BAR_TRACK} inline-block w-full max-w-[160px] flex-1 md:order-1 lg:max-w-[96px] xl:max-w-[160px]"
              aria-hidden="true"
              ><span class={BAR_FILL} style:width={`${(deathCount(row) / most) * 100}%`}></span></span
            >
            <span class="text-muted text-[12px] whitespace-nowrap md:order-3"
              >{deathSplit(row.on_wipes, row.on_kills)}</span
            >
          </span>
        </li>
      {/each}
    </ul>
    {#if !expanded && ordered.length > shown.length}
      <RaidShowAll total={ordered.length} onShow={() => void reveal()} testid="raid-deaths-show-all" />
    {/if}
  {/if}
</RaidBlock>
