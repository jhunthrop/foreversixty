<!-- web/src/components/sim/tools/StatWeights.svelte -->
<!-- Design 7. The caution opens the page, in our words, and links to Top Gear -- it is not
     a footnote, because the whole reason this page exists is that addons ask for a number
     the number itself cannot justify.

     The reference stat always leads `store.stats`: the engine normalises its weights to it,
     so the request must weigh it. It is seeded from the spec's own
     `reference_stat` (weights.ts's `referenceFor`/`defaultStatsFor`, run by the store) and
     is never itself untickable here -- the checkbox for it is disabled, not merely
     defaulted, so `WeightsSpec.Reference` (required, contract 10.8) can never go empty
     through this control. `store.stats` cannot reach [] this way; only `setStats([])`
     directly (the store's own pre-send refusal, `bulkCopy.weightsNeedStats`) can.

     The results read as the published weights do (owner, 2026-09-30, round 11): the
     SimulationCraft scale-factor table (WeightsScaleTable.svelte over weights.ts's
     `toolScaleView`), top stat = 1.00 with DPS per point and the error, haste per 1% in the
     caption, no Pawn string. A row the engine flagged `insignificant` greys there. The
     picker above them offers only
     `store.weightStats` when the spec sent one (sub-item 4): a stat this 1.60 sim does not
     model for this spec is never on screen to tick in the first place. -->
<script lang="ts">
  import type { Me } from '../../../lib/account/api';
  import { SECONDARY_BUTTON } from '../../../lib/planner/styles';
  import type { BulkStore } from '../../../lib/sim/bulk-store.svelte';
  import { bulkCopy, WEIGHTS_STATS_FROM_ENGINE } from '../../../lib/sim/copy';
  import { specLabel } from '../../../lib/sim/spec-label';
  import { hasWeightStats, pickableStatsFor, statLabel, toolScaleView } from '../../../lib/sim/weights';
  import WeightsScaleTable from '../WeightsScaleTable.svelte';
  import BulkRunBar from './BulkRunBar.svelte';
  import SaveSimForm from './SaveSimForm.svelte';

  let { store }: { store: BulkStore; me: Me | null } = $props();

  const view = $derived(toolScaleView(store.weights, store.character?.spec ?? '', store.referenceStat));
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
    <h3 class="section-title">{bulkCopy.weightsPick}</h3>
    {#if store.referenceStat}
      <p class="text-muted text-[12px]" data-testid="sim-weights-reference">
        {bulkCopy.weightsReference}: {statLabel(store.referenceStat)}
      </p>
    {/if}
    {#if hasWeightStats(store.weightStats)}
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
           (the table's own "Not significant", which the caveat's own first sentence
           restates in prose), not a second warning box stacked above this one. -->
      <p class="text-muted text-[12px]" data-testid="sim-weights-error-caveat">
        {bulkCopy.weightsErrorCaveat(store.weights.map((row) => row.stat))}
      </p>
      <WeightsScaleTable {view} />

      <!-- Contract 10.6: saved sims of every kind are public at /sim/<id>. A finished
           weights run saves through the identical component and flow Top Gear and talent
           compare already use (SaveSimForm.svelte, ComboResults.svelte's own extraction),
           rather than a second, drifting copy of the same save form. -->
      <SaveSimForm onsave={(title) => store.save(title)} titleFor={saveTitleFor} {canSave} />
    </section>
  {/if}
</div>
