<!-- web/src/components/sim/SimHistory.svelte -->
<!-- A signed-in player's saved sims. Rendered only when there is a session: a signed-out
     placeholder inviting an account would be an advertisement on a page that works fine
     without one, and the free browser lane is the product. -->
<script lang="ts">
  import { formatDate } from '../../lib/dates';
  import { bestPerSpec } from '../../lib/sim/best-per-spec';
  import { simCopy } from '../../lib/sim/copy';
  import { KIND_FILTERS, headlineOf, kindOf, titleOf, type KindFilter } from '../../lib/sim/history';
  import type { SimListRow } from '../../lib/sim/types';
  import BestPerSpecTable from './BestPerSpecTable.svelte';
  import EmptyState from '../ui/EmptyState.svelte';

  let {
    rows,
    error,
    kind,
    onkind,
  }: {
    rows: SimListRow[] | null;
    error: string | null;
    kind: KindFilter;
    onkind: (next: KindFilter) => void;
  } = $props();

  // 2026-09-26 layout pass, Finding 4: the current character's own rows, by spec, best DPS
  // first -- derived from the exact rows the list below renders, so the two never disagree.
  // Hidden with fewer than two specs: one spec has nothing to compare yet.
  const perSpec = $derived(rows === null ? [] : bestPerSpec(rows));
</script>

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-history">
  <div class="flex flex-wrap items-baseline justify-between gap-3">
    <h2 class="section-title text-[15px]">{simCopy.yourSims}</h2>
    <label class="flex items-center gap-2 text-[13px]">
      <span class="label text-muted">{simCopy.historyFilter}</span>
      <select
        class="border-line-warm rounded-control bg-raised text-text min-h-11 border px-2 text-[13px] md:min-h-9"
        value={kind}
        onchange={(event) => onkind(event.currentTarget.value as KindFilter)}
        data-testid="sim-history-filter"
      >
        {#each KIND_FILTERS as filter (filter)}
          <option value={filter}>{simCopy.kindLabel[filter]}</option>
        {/each}
      </select>
    </label>
  </div>
  {#if perSpec.length >= 2}
    <BestPerSpecTable rows={perSpec} />
  {/if}
  {#if error !== null}
    <p role="alert" class="text-strong text-[13px]">{error}</p>
  {:else if rows === null}
    <p class="text-muted text-[13px]">{simCopy.historyLoading}</p>
  {:else if rows.length === 0}
    <EmptyState message={simCopy.historyEmpty} testid="sim-history-empty" />
  {:else}
    <ul class="border-line bg-raised rounded-panel flex flex-col border">
      {#each rows as row (row.sim_id)}
        <li class="border-line-soft border-b last:border-b-0">
          <!-- Finding 6, 2026-09-26 layout pass: the date used to be `hidden md:inline`,
               dropped on phone with no sign anything was cut. The two-column phone grid
               below (`minmax(0,1fr)_auto`) auto-flows: pill/title on the first line, the
               headline spanning both columns (`col-span-2`) on its own second line, and the
               date -- now always rendered -- on its own third line, under the DPS the
               headline names, rather than vanishing. Desktop drops the span and the four
               items map straight onto the four explicit columns as before. -->
          <a
            class="grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 px-3 py-2 text-[14px] md:grid-cols-[88px_minmax(0,1.4fr)_minmax(0,1.2fr)_96px]"
            href={`/sim/${row.sim_id}`}
            data-testid={`sim-history-${row.sim_id}`}
          >
            <span class="pill pill-sample shrink-0" data-testid={`sim-history-kind-${row.sim_id}`}>
              {simCopy.kindLabel[kindOf(row)]}
            </span>
            <span class="text-strong truncate font-semibold">{titleOf(row)}</span>
            <!-- The API's own sentence replaces the bare DPS column: "3 upgrades on
                 Ragnaros" says what a Droptimizer row is and a number does not. -->
            <span
              class="text-muted col-span-2 truncate text-[13px] md:col-span-1"
              data-testid={`sim-history-headline-${row.sim_id}`}
            >
              {headlineOf(row)}
            </span>
            <span class="tabular text-muted text-left font-mono text-[12px] md:text-right md:text-[13px]">
              {formatDate(new Date(row.created_at))}
            </span>
          </a>
        </li>
      {/each}
    </ul>
  {/if}
</section>
