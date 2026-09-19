<!-- web/src/components/sim/SimHistory.svelte -->
<!-- A signed-in player's saved sims. Rendered only when there is a session: a signed-out
     placeholder inviting an account would be an advertisement on a page that works fine
     without one, and the free browser lane is the product. -->
<script lang="ts">
  import { formatAmount } from '../../lib/report/format';
  import { simCopy } from '../../lib/sim/copy';
  import { specLabel } from '../../lib/sim/spec-label';
  import type { SimListRow } from '../../lib/sim/types';

  let { rows, error }: { rows: SimListRow[] | null; error: string | null } = $props();
</script>

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-history">
  <h2 class="section-title text-[15px]">{simCopy.yourSims}</h2>
  {#if error !== null}
    <p role="alert" class="text-strong text-[13px]">{error}</p>
  {:else if rows === null}
    <p class="text-muted text-[13px]">{simCopy.historyLoading}</p>
  {:else if rows.length === 0}
    <p class="text-muted text-[13px]">{simCopy.historyEmpty}</p>
  {:else}
    <ul class="border-line bg-raised rounded-panel flex flex-col border">
      {#each rows as row (row.sim_id)}
        <li class="border-line-soft border-b last:border-b-0">
          <a
            class="grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 px-3 py-2 text-[14px] md:grid-cols-[minmax(0,1.6fr)_minmax(0,1fr)_88px_96px]"
            href={`/sim/${row.sim_id}`}
            data-testid={`sim-history-${row.sim_id}`}
          >
            <span class="text-strong truncate font-semibold"
              >{row.title === '' ? specLabel(row.spec) : row.title}</span
            >
            <span class="text-muted hidden truncate text-[13px] md:inline">{specLabel(row.spec)}</span>
            <span class="tabular text-right font-mono">{formatAmount(Math.round(row.dps))}</span>
            <span class="tabular text-muted hidden text-right font-mono text-[13px] md:inline">
              {row.created_at.slice(0, 10)}
            </span>
          </a>
        </li>
      {/each}
    </ul>
  {/if}
</section>
