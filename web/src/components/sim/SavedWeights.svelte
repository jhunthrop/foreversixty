<!-- web/src/components/sim/SavedWeights.svelte -->
<!-- A saved stat-weights result at /sim/<id>: read-only. The caution travels with the
     numbers -- a shared link is exactly where someone meets these weights without having
     read the page that made them (design 7) -- so the same warning that opens /sim/weights
     opens this page too, in the same words.

     No statistic is computed here: a weight and its error come straight off
     `result.weights`' own `StatWeight` rows; `weightScale`/`pawnString`/`statLabel` are
     Task 9's, unmodified. The bar geometry is presentation math over those two numbers, the
     same local helper `StatWeights.svelte` (Task 19) already uses -- not extracted into a
     shared function, per that task's own report, so it is re-declared here rather than
     imported from a component this page does not otherwise depend on. Task 20.

     D45/D46: a saved run's own greyed rows and Pawn string follow the live tool page's
     rules exactly, off the same `isSignificant`/`pawnString` -- a shared link is exactly
     where someone meets these weights without having read the caution above (this file's
     own header), so a saved page that still bolded a row the engine called noise, or
     printed it into the Pawn string, would be D45/D46 again with a permalink. -->
<script lang="ts">
  import { SECONDARY_BUTTON } from '../../lib/planner/styles';
  import type { WeightsResult } from '../../lib/sim/bulk-types';
  import { bulkCopy, WEIGHT_INSIGNIFICANT_LABEL } from '../../lib/sim/copy';
  import {
    formatWeightError,
    isSignificant,
    pawnString,
    statLabel,
    weightScale,
  } from '../../lib/sim/weights';

  let { result }: { result: WeightsResult } = $props();

  let copied = $state(false);

  const scale = $derived(weightScale(result.weights));
  const pawn = $derived(pawnString(result.request.spec, result.weights));

  /** Bar geometry as percentages of the widest weight plus its own error. */
  function bar(weight: number, error: number): { width: string; error: string } {
    return {
      width: `${Math.max(0, (weight / scale) * 100)}%`,
      error: `${Math.max(0, ((2 * error) / scale) * 100)}%`,
    };
  }

  async function copyPawn(): Promise<void> {
    try {
      await navigator.clipboard.writeText(pawn);
      copied = true;
    } catch {
      copied = false;
    }
  }
</script>

<section
  class="border-line-warm rounded-panel mx-[18px] border p-4 md:mx-0"
  data-testid="sim-weights-warning"
>
  <p class="text-strong text-[14px]">{bulkCopy.weightsWarning}</p>
  <a class="text-nav underline" href="/sim/gear">{bulkCopy.weightsWarningLink}</a>
</section>

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-weights">
  <!-- Task 8, sub-item 1: the same caveat as the live page (StatWeights.svelte), in the
       same words, right before the numbers it caveats -- see that file's own comment. -->
  <p class="text-muted text-[12px]" data-testid="sim-weights-error-caveat">
    {bulkCopy.weightsErrorCaveat(result.weights.map((row) => row.stat))}
  </p>
  <ul class="flex flex-col">
    {#each result.weights as row (row.stat)}
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
          <span class="bg-gold absolute top-0 left-0 block h-[8px]" style={`width:${geometry.width}`}></span>
          <span
            class="border-line-warm absolute top-0 block h-[8px] border-x"
            style={`left:calc(${geometry.width} - ${geometry.error}/2);width:${geometry.error}`}
            aria-hidden="true"
          ></span>
        </span>
        <span class="tabular text-strong ml-auto font-mono text-[13px]">
          {row.weight.toFixed(2)}
          <span class="text-muted">± {formatWeightError(row.error)}</span>
          {#if !significant}
            <span class="text-muted block font-sans text-[11px]" data-testid={`sim-weight-note-${row.stat}`}>
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
</section>
