<!-- web/src/components/rankings/RankingsLegend.svelte -->
<!-- The one-line colour legend (design/specs/2026-10-11-rankings.md section 3), the board card's
     foot: the percentile steps as six tags. It repeats no count: the header already says it. -->
<script lang="ts">
  import { rankingsCopy } from '../../lib/rankings/copy';
  import { RANK_LEGEND } from '../../lib/rankings/row-view';
</script>

<p class="text-muted flex flex-wrap items-center gap-x-2 gap-y-1.5 text-[12px]" data-testid="rankings-legend">
  {rankingsCopy.legendLead}
  {#each RANK_LEGEND as step (step.label)}
    <!-- Tag's shape (22 px, 11 px 700 uppercase, 1 px border) in the parse colour; Tag
         itself has four state tones only. -->
    <span
      class="inline-flex h-[22px] items-center rounded-full border border-current/40 bg-current/10 px-2 text-[11px] font-bold tracking-[.08em] whitespace-nowrap uppercase"
      style:color={step.color}>{step.label}</span
    >
  {/each}
</p>
