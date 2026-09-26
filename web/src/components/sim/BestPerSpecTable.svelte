<!-- web/src/components/sim/BestPerSpecTable.svelte -->
<!-- 2026-09-26 layout pass, review round 1 Finding 4: a compact "spec, best DPS, when"
     summary above the sim history list -- SimHistory.svelte hides this entirely unless the
     current character has sims in two or more specs (one spec has nothing to compare). Rows
     come from `bestPerSpec` (best-per-spec.ts), a pure function over the same rows the list
     below renders, so the two can never disagree on which run was best. -->
<script lang="ts">
  import { formatDate } from '../../lib/dates';
  import { formatAmount } from '../../lib/report/format';
  import type { BestPerSpecRow } from '../../lib/sim/best-per-spec';
  import { landingCopy } from '../../lib/sim/landing-copy';
  import { specDisplayName } from '../../lib/sim/spec-label';

  let { rows }: { rows: BestPerSpecRow[] } = $props();
</script>

<div class="border-line bg-raised rounded-panel flex flex-col border" data-testid="sim-best-per-spec">
  <h3 class="label text-muted px-3 pt-3 pb-1">{landingCopy.bestPerSpecTitle}</h3>
  <ul class="flex flex-col">
    {#each rows as row (row.spec)}
      <li
        class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto_auto] items-center gap-x-3 border-b px-3 py-2 text-[14px] last:border-b-0"
        data-testid={`sim-best-per-spec-${row.spec}`}
      >
        <span class="text-strong truncate font-semibold">{specDisplayName(row.spec)}</span>
        <span class="tabular text-text text-right font-mono text-[13px]">
          {formatAmount(Math.round(row.dps))} DPS
        </span>
        <span class="tabular text-muted text-right font-mono text-[12px]">
          {formatDate(new Date(row.createdAt))}
        </span>
      </li>
    {/each}
  </ul>
</div>
