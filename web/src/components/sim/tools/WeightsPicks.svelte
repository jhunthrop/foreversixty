<!-- web/src/components/sim/tools/WeightsPicks.svelte -->
<!-- /sim/weights' stats to weigh, as one row of ticks over the table they fill (h-audit 2026-11-09:
     alone in the rail it sat 400 px short of the table). The reference stat is its locked tick
     (StatWeights.svelte's file comment says why). Without `ontoggle` (WeightsShell's server
     render) the boxes draw disabled until the island mounts. -->
<script lang="ts">
  import { bulkCopy } from '../../../lib/sim/copy';
  import type { WeightStat } from '../../../lib/sim/weights';
  import Section from '../../ui/Section.svelte';

  let {
    pickable,
    picked,
    reference,
    ontoggle,
  }: {
    pickable: readonly WeightStat[];
    picked: readonly string[];
    /** The stat the weights normalise to; '' before a character is loaded. */
    reference: string;
    ontoggle?: (id: string) => void;
  } = $props();
</script>

<Section title={bulkCopy.weightsPick} panel testid="sim-weights-pick">
  <ul class="m-0 flex list-none flex-wrap gap-x-6 p-0">
    {#each pickable as stat (stat.id)}
      <li>
        <label class="text-text flex min-h-11 items-center gap-2 text-[14px]">
          <input
            type="checkbox"
            class="accent-gold h-[18px] w-[18px]"
            data-testid={`sim-weight-pick-${stat.id}`}
            checked={picked.includes(stat.id)}
            disabled={stat.id === reference || ontoggle === undefined}
            onchange={() => ontoggle?.(stat.id)}
          />
          {stat.label}
        </label>
      </li>
    {/each}
  </ul>
</Section>
