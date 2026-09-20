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
     `store.weights`' own `StatWeight` rows, and `weightScale`/`pawnString` are Task 9's.

     D45/D46: a row the engine flagged `insignificant` greys here and drops out of the Pawn
     string below -- both read `isSignificant`, weights.ts's one predicate for "is this row
     meaningful", so the table and the Pawn string cannot disagree about which rows those
     are (weights.test.ts's own regression guard). The picker above them offers only
     `store.weightStats` when the spec sent one (sub-item 4): a stat this 1.60 sim does not
     model for this spec is never on screen to tick in the first place. -->
<script lang="ts">
  import type { Me } from '../../../lib/account/api';
  import { SECONDARY_BUTTON } from '../../../lib/planner/styles';
  import type { BulkStore } from '../../../lib/sim/bulk-store.svelte';
  import { bulkCopy, WEIGHT_INSIGNIFICANT_LABEL, WEIGHTS_STATS_FROM_ENGINE } from '../../../lib/sim/copy';
  import { specLabel } from '../../../lib/sim/spec-label';
  import {
    isSignificant,
    pawnString,
    pickableStatsFor,
    statLabel,
    weightScale,
  } from '../../../lib/sim/weights';
  import BulkRunBar from './BulkRunBar.svelte';
  import SaveSimForm from './SaveSimForm.svelte';

  let { store }: { store: BulkStore; me: Me | null } = $props();

  let copied = $state(false);

  const scale = $derived(weightScale(store.weights));
  const pawn = $derived(
    store.weights.length === 0 ? '' : pawnString(store.character?.spec ?? '', store.weights),
  );
  /** The picker's own list: `store.weightStats` when the spec sent one, the full pinned
   *  vocabulary otherwise (sub-item 4; `weights.ts`'s own fallback rule). */
  const pickable = $derived(pickableStatsFor(store.weightStats));

  /**
   * `store.save` (bulk-store.svelte.ts) has no fallback title of its own, the same as every
   * other tool page's save form -- see ComboResults.svelte's identical comment. A weights
   * run has no winner to headline, so this names the tool and the spec instead, the same
   * two-part title SavedSim.svelte gives a saved weights page's own <h1>.
   */
  const saveTitleFor = $derived(`${bulkCopy.weightsTitle} · ${specLabel(store.character?.spec ?? '')}`);

  /** A stopped run has nothing finished worth naming and saving. */
  const canSave = $derived(store.result?.aborted !== true);

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
    {#if store.weightStats !== undefined}
      <p class="text-muted text-[12px]" data-testid="sim-weights-stats-note">{WEIGHTS_STATS_FROM_ENGINE}</p>
    {/if}
    <ul class="flex flex-wrap gap-3">
      {#each pickable as stat (stat.id)}
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
      <!-- Task 8, sub-item 1: contract 10.9's error-is-a-lower-bound caveat, read right
           before the numbers it caveats -- one thought together with the greying below
           (D45's own WEIGHT_INSIGNIFICANT_LABEL, which the caveat's own first sentence
           restates in prose), not a second warning box stacked above this one. -->
      <p class="text-muted text-[12px]" data-testid="sim-weights-error-caveat">
        {bulkCopy.weightsErrorCaveat}
      </p>
      <ul class="flex flex-col">
        {#each store.weights as row (row.stat)}
          {@const geometry = bar(row.weight, row.error)}
          {@const significant = isSignificant(row)}
          <li
            class="border-line-soft grid min-h-11 grid-cols-[minmax(120px,1fr)_minmax(0,3fr)_88px] items-center gap-x-3 border-b px-2 py-2 {significant
              ? ''
              : 'opacity-50'}"
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
              {#if !significant}
                <span
                  class="text-muted block font-sans text-[11px]"
                  data-testid={`sim-weight-note-${row.stat}`}
                >
                  {WEIGHT_INSIGNIFICANT_LABEL}
                </span>
              {/if}
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

      <!-- Contract 10.6: saved sims of every kind are public at /sim/<id>. A finished
           weights run saves through the identical component and flow Top Gear and talent
           compare already use (SaveSimForm.svelte, ComboResults.svelte's own extraction),
           rather than a second, drifting copy of the same save form. -->
      <SaveSimForm onsave={(title) => store.save(title)} titleFor={saveTitleFor} {canSave} />
    </section>
  {/if}
</div>
