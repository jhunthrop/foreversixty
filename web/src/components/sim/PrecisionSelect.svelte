<!-- web/src/components/sim/PrecisionSelect.svelte -->
<!-- The one precision field (data-testid="sim-precision") every precision control on the sim lane
     renders -- /sim's own RunControl (the full four options, including "until ±0.5%") and
     every bulk tool page's BulkRunBar (the three fixed counts; /sim/weights now among
     them, contract 10.8's page having shipped with no working precision control at all --
     dps-minmaxer review round 1, D45). Extracted rather than left as two hand-copied
     select blocks (which is what BulkRunBar's and RunControl's markup used to be,
     byte-for-byte the same shape with a different option list): a caller owns its own
     option array and its own label text, because the two lanes' precisions genuinely do
     not mean the same thing -- a bulk run's `fast`/`normal`/`high` pick a stage ladder, a
     plain or a weights run's pick a flat iteration count, and only /sim's own run has a
     stopping rule at all (a weights run is one wasm call with no adaptive loop to stop --
     see bulk-run.ts's runWeightsRun -- so offering "until ±0.5%" there would describe a
     behaviour that does not exist). What both lanes share, and now share literally, is the
     field itself: its wiring, its disabled state, its test id. A thin wrapper over the shared
     Select (design/specs/2026-10-30-controls-charts.md §A10 row 18). -->
<script lang="ts">
  import Select from '../ui/Select.svelte';

  let {
    id,
    value,
    options,
    labelFor,
    disabled = false,
    onchange,
    testid = 'sim-precision',
    class: className = '',
  }: {
    /** The field's id, which the caller's own `<label for>` names. */
    id: string;
    value: string;
    options: readonly string[];
    labelFor: (id: string) => string;
    disabled?: boolean;
    onchange: (value: string) => void;
    testid?: string;
    class?: string;
  } = $props();
</script>

<Select
  {id}
  class={className}
  size="compact"
  {value}
  options={options.map((option) => ({ value: option, label: labelFor(option) }))}
  {disabled}
  {testid}
  onchange={(event) => onchange((event.currentTarget as HTMLSelectElement).value)}
/>
