<!-- web/src/components/sim/tools/StatWeights.svelte -->
<!-- Design 7. The caution opens the page, in our words, and links to Top Gear -- it is not
     a footnote, because the whole reason this page exists is that addons ask for a number
     the number itself cannot justify.

     The reference stat always leads `store.stats`: pawnString and every row on this page
     read `stats[0]` as the one normalised to 1.00, and a picker that let it move would
     silently rescale every other number on screen. It is seeded from the spec's own
     `reference_stat` (weights.ts's `referenceFor`/`defaultStatsFor`, run by the store) and
     is never itself untickable here -- the checkbox for it is disabled, not merely
     defaulted, so `WeightsSpec.Reference` (required, contract 10.8) can never go empty
     through this control. `store.stats` cannot reach [] this way; only `setStats([])`
     directly (the store's own pre-send refusal, `bulkCopy.weightsNeedStats`) can.

     No statistic is computed here: a weight and its error come straight off
     `store.weights`' own `StatWeight` rows, and `weightScale`/`pawnString` are Task 9's. -->
<script lang="ts">
  import type { Me } from '../../../lib/account/api';
  import { SECONDARY_BUTTON } from '../../../lib/planner/styles';
  import type { BulkStore } from '../../../lib/sim/bulk-store.svelte';
  import { bulkCopy } from '../../../lib/sim/copy';
  import { pawnString, statLabel, weightScale, WEIGHT_STATS } from '../../../lib/sim/weights';
  import BulkRunBar from './BulkRunBar.svelte';

  let { store }: { store: BulkStore; me: Me | null } = $props();

  let copied = $state(false);

  const scale = $derived(weightScale(store.weights));
  const pawn = $derived(
    store.weights.length === 0 ? '' : pawnString(store.character?.spec ?? '', store.weights),
  );

  function togglePick(id: string): void {
    if (id === store.referenceStat) return;
    const next = store.stats.includes(id) ? store.stats.filter((stat) => stat !== id) : [...store.stats, id];
    // The reference stat always leads (see the file comment above).
    const reference = store.referenceStat;
    store.setStats(reference === '' ? next : [reference, ...next.filter((stat) => stat !== reference)]);
  }

  async function copyPawn(): Promise<void> {
    try {
      await navigator.clipboard.writeText(pawn);
      copied = true;
    } catch {
      copied = false;
    }
  }

  /** Bar geometry as percentages of the widest weight plus its error. */
  function bar(weight: number, error: number): { width: string; error: string } {
    return {
      width: `${Math.max(0, (weight / scale) * 100)}%`,
      error: `${Math.max(0, ((2 * error) / scale) * 100)}%`,
    };
  }
</script>

<div class="flex flex-col gap-[22px] md:gap-8" data-testid="sim-stat-weights">
  <section
    class="border-line-warm rounded-panel mx-[18px] border p-4 md:mx-0"
    data-testid="sim-weights-warning"
  >
    <p class="text-strong text-[14px]">{bulkCopy.weightsWarning}</p>
    <a class="{SECONDARY_BUTTON} border-line-warm text-nav mt-2 px-4" href="/sim/gear">
      {bulkCopy.weightsWarningLink}
    </a>
  </section>

  <section class="border-line rounded-panel mx-[18px] flex flex-col gap-2 border p-3 md:mx-0">
    <h3 class="section-title text-[14px]">{bulkCopy.weightsPick}</h3>
    <p class="text-muted text-[12px]" data-testid="sim-weights-reference">
      {bulkCopy.weightsReference}: {statLabel(store.referenceStat)}
    </p>
    <ul class="flex flex-wrap gap-3">
      {#each WEIGHT_STATS as stat (stat.id)}
        <li>
          <label class="text-text flex min-h-11 items-center gap-2 text-[13px]">
            <input
              type="checkbox"
              class="h-5 w-5"
              data-testid={`sim-weight-pick-${stat.id}`}
              checked={store.stats.includes(stat.id)}
              disabled={stat.id === store.referenceStat}
              onchange={() => togglePick(stat.id)}
            />
            {stat.label}
          </label>
        </li>
      {/each}
    </ul>
  </section>

  <BulkRunBar {store} />

  {#if store.weights.length > 0}
    <section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-weights">
      <ul class="flex flex-col">
        {#each store.weights as row (row.stat)}
          {@const geometry = bar(row.weight, row.error)}
          <li
            class="border-line-soft grid min-h-11 grid-cols-[minmax(120px,1fr)_minmax(0,3fr)_88px] items-center gap-x-3 border-b px-2 py-2"
            data-testid={`sim-weight-${row.stat}`}
          >
            <span class="text-text text-[13px]">{statLabel(row.stat)}</span>
            <span class="bg-card-top relative block h-[8px] w-full">
              <span class="bg-gold absolute top-0 left-0 block h-[8px]" style={`width:${geometry.width}`}
              ></span>
              <span
                class="border-line-warm absolute top-0 block h-[8px] border-x"
                style={`left:calc(${geometry.width} - ${geometry.error}/2);width:${geometry.error}`}
                aria-hidden="true"
              ></span>
            </span>
            <span class="tabular text-strong ml-auto font-mono text-[13px]">
              {row.weight.toFixed(2)}
              <span class="text-muted">± {row.error.toFixed(2)}</span>
            </span>
          </li>
        {/each}
      </ul>

      <div class="border-line rounded-panel flex flex-wrap items-center gap-3 border p-3">
        <code class="text-muted flex-1 font-mono text-[12px] break-all" data-testid="sim-pawn">{pawn}</code>
        <button
          type="button"
          class="{SECONDARY_BUTTON} border-line-warm text-nav px-4"
          data-testid="sim-pawn-copy"
          onclick={() => void copyPawn()}
        >
          {copied ? bulkCopy.weightsCopied : bulkCopy.weightsCopyPawn}
        </button>
      </div>
    </section>
  {/if}
</div>
