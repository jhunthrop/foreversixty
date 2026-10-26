<!-- web/src/components/guild/RaidDeaths.svelte -->
<!-- Deaths (attendance spec 3.4, officers): one 40 px row per raider who died, most first; an 8 px
     muted bar (never `--wipe`, R-A3), the count, and the wipe/kill split. The first ten show, then
     `Show all n raiders`. None: `Nobody died.` -->
<script lang="ts">
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

  const ordered = $derived(deathOrder(rows));
  const most = $derived(ordered.length === 0 ? 1 : deathCount(ordered[0]));
  const shown = $derived(visibleRows(ordered, expanded));
</script>

<RaidBlock heading={attendanceCopy.report.deaths} count={ordered.length} testid="raid-deaths">
  {#if ordered.length === 0}
    <p class="text-muted text-[14px]" data-testid="raid-deaths-none">{attendanceCopy.deaths.none}</p>
  {:else}
    <ul>
      {#each shown as row (row.character_key)}
        <li
          class="{ROW_DIVIDER} flex min-h-10 items-center gap-3 px-1"
          style={ownKeys.includes(row.character_key) ? `background:${OWN_ROW_WASH}` : undefined}
          data-testid="raid-death-row"
        >
          <ClassCrestRing characterClass={row.class} size={20} />
          <span
            class="w-28 truncate text-[13px] font-semibold"
            style={`color:${classColorVar(row.class)}`}
            title={row.name}>{row.name}</span
          >
          <span class="{BAR_TRACK} hidden w-full max-w-[160px] sm:inline-block" aria-hidden="true"
            ><span class={BAR_FILL} style:width={`${(deathCount(row) / most) * 100}%`}></span></span
          >
          <span class="text-strong font-mono text-[13px]">{deathCount(row)}</span>
          <span class="text-muted text-[12px]">{deathSplit(row.on_wipes, row.on_kills)}</span>
        </li>
      {/each}
    </ul>
    {#if !expanded && ordered.length > shown.length}
      <RaidShowAll total={ordered.length} onShow={() => (expanded = true)} testid="raid-deaths-show-all" />
    {/if}
  {/if}
</RaidBlock>
