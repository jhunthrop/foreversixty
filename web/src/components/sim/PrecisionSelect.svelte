<!-- web/src/components/sim/PrecisionSelect.svelte -->
<!-- The one <select data-testid="sim-precision"> every precision control on the sim lane
     renders -- /sim's own RunControl (the full four options, including "until ±0.5%") and
     every bulk tool page's BulkRunBar (the three fixed counts; /sim/weights now among
     them, contract 10.8's page having shipped with no working precision control at all --
     dps-minmaxer review round 1, D45). Extracted rather than left as two hand-copied
     `<select>` blocks (which is what BulkRunBar's and RunControl's markup used to be,
     byte-for-byte the same shape with a different option list): a caller owns its own
     option array and its own label text, because the two lanes' precisions genuinely do
     not mean the same thing -- a bulk run's `fast`/`normal`/`high` pick a stage ladder, a
     plain or a weights run's pick a flat iteration count, and only /sim's own run has a
     stopping rule at all (a weights run is one wasm call with no adaptive loop to stop --
     see bulk-run.ts's runWeightsRun -- so offering "until ±0.5%" there would describe a
     behaviour that does not exist). What both lanes share, and now share literally, is the
     element itself: its wiring, its disabled state, its test id. -->
<script lang="ts">
  let {
    value,
    options,
    labelFor,
    disabled = false,
    onchange,
    testid = 'sim-precision',
    class: className = '',
  }: {
    value: string;
    options: readonly string[];
    labelFor: (id: string) => string;
    disabled?: boolean;
    onchange: (value: string) => void;
    testid?: string;
    class?: string;
  } = $props();
</script>

<select
  class={className}
  data-testid={testid}
  {disabled}
  {value}
  onchange={(event) => onchange(event.currentTarget.value)}
>
  {#each options as id (id)}
    <option value={id}>{labelFor(id)}</option>
  {/each}
</select>
