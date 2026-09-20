<!-- web/src/components/sim/tools/BulkRunBar.svelte -->
<!-- Design 3.1.8: the combination count, the precision, the cap notice, the lane switch and
     one button. Every tool page mounts this, so the run affordance is identical on all four
     and the settings bar and the request drawer arrive with it.

     The cap notice never trims the list. It says what would exceed the cap and by how much,
     and offers the lane that would allow it -- design 2.3's own rule.

     "SettingsPanel" in this task's brief is part A's SettingsBar.svelte: it takes exactly
     the `{ settings, spec, disabled, onchange }` shape the brief describes (it mounts
     SettingsSheet's own disclosure itself, so nothing else here needs to). -->
<script lang="ts">
  import { SECONDARY_BUTTON } from '../../../lib/planner/styles';
  import type { BulkStore } from '../../../lib/sim/bulk-store.svelte';
  import { BULK_PRECISIONS, LOW_CORE_CAP, type Precision } from '../../../lib/sim/bulk-types';
  import { bulkCopy, simCopy } from '../../../lib/sim/copy';
  import PrecisionSelect from '../PrecisionSelect.svelte';
  import RequestDrawer from '../RequestDrawer.svelte';
  import SettingsBar from '../SettingsBar.svelte';

  let { store }: { store: BulkStore } = $props();

  /**
   * The weights tool sends no `bulk` block at all (bulk-store.svelte.ts's `mode` is null
   * for it), so `recount` never runs and `combinations` never leaves null -- the count
   * would read "Counting combinations…" for as long as the page is open, so the count and
   * its cap machinery stay behind this flag. Precision is NOT dead there any more
   * (dps-minmaxer review round 1, D45: the weights page shipped with no working precision
   * control at all): `buildRequest` now sends `deps.getPrecision()`'s own iteration count
   * for a weights run, the same `store.precision` this bar already reads and writes for
   * the other three tools -- see the select below, which renders for every tool and only
   * swaps its label text and drops the staged note when `combinationTool` is false, since
   * a weights run is one flat iteration count with no stage ladder to describe.
   */
  const combinationTool = $derived(store.tool !== 'weights');
  const running = $derived(store.phase === 'running' || store.serverRunning);
  /**
   * `store.combinations` stays `null` both while a count is in flight and after one has
   * failed (bulk-store-request.ts's `recount`, generic branch) -- the two used to look
   * identical here, so a failed count read "Counting combinations…" forever with no
   * indication anything had gone wrong. `store.message` is what tells them apart: it is
   * cleared at the top of every `recount` call and only set on a real failure, so a
   * non-null message with a null count means the answer already arrived (badly) rather
   * than still being awaited. The real explanation is the `sim-message` block below this
   * bar; this placeholder just gets out of its way instead of contradicting it.
   */
  const countLabel = $derived(
    store.combinations !== null
      ? bulkCopy.combinations(store.combinations)
      : store.message !== null
        ? ''
        : store.phase === 'counting'
          ? bulkCopy.combinationsCounting
          : bulkCopy.combinationsNone,
  );
  /**
   * A combination tool with nothing to run: no count yet (nothing ticked, or a source
   * whose drops none of this character can wear, which `recount` exits on before it ever
   * counts) or a count of zero. The weights page has no count and is never gated on one.
   */
  const nothingToRun = $derived(
    combinationTool && (store.combinations === null || store.combinations === 0) && store.phase !== 'running',
  );
  const PRECISION_LABELS: Record<Precision, string> = {
    fast: bulkCopy.precisionFast,
    normal: bulkCopy.precisionNormal,
    high: bulkCopy.precisionHigh,
  };
  /**
   * A combination tool's word is a bare "Fast"/"Normal"/"High" -- `precisionNote` below
   * says what it means. The weights tool has no stage note to say it for, so it takes
   * `/sim`'s own three fixed-count labels instead ("Fast, 500 iterations", ...) -- the
   * identical `simCopy.precisionLabel` map `RunControl.svelte` reads from, unmodified,
   * rather than a second sentence saying the same thing a different way.
   */
  const precisionLabelFor = (id: Precision): string =>
    combinationTool ? PRECISION_LABELS[id] : (simCopy.precisionLabel[id] ?? id);
</script>

<SettingsBar
  settings={store.settings}
  spec={store.character?.spec ?? ''}
  disabled={running}
  onchange={(next) => store.setSettings(next)}
/>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-3 border p-4 md:mx-0"
  data-testid="sim-run-bulk-bar"
>
  <div class="flex flex-wrap items-center gap-4">
    {#if combinationTool}
      <span class="tabular text-strong font-mono text-[14px]" data-testid="sim-combo-count">
        {countLabel}
      </span>
    {/if}

    <label class="text-muted flex items-center gap-2 text-[12px]">
      {bulkCopy.precisionLabel}
      <PrecisionSelect
        class="border-line-warm rounded-control bg-bg text-text h-11 border px-2 text-[14px]"
        value={store.precision}
        options={BULK_PRECISIONS}
        labelFor={precisionLabelFor}
        disabled={running}
        onchange={(value) => store.setPrecision(value as Precision)}
      />
    </label>

    <button
      type="button"
      class="{SECONDARY_BUTTON} border-line-warm text-nav px-4"
      data-testid="sim-run-bulk"
      disabled={store.capNotice !== null ||
        store.character === null ||
        store.phase === 'loading-character' ||
        (nothingToRun && !running)}
      onclick={() => (running ? store.stop() : void store.run())}
    >
      {#if running}{bulkCopy.stopBulk}{:else if store.result !== null}{bulkCopy.runBulkAgain}{:else}{bulkCopy.runBulk}{/if}
    </button>

    {#if store.premium}
      <button
        type="button"
        class="{SECONDARY_BUTTON} border-line-warm text-nav px-4"
        data-testid="sim-server-run"
        disabled={running || store.serverCapNotice !== null}
        onclick={() => void store.runOnServer()}>{bulkCopy.capPremium}</button
      >
    {/if}
  </div>

  {#if combinationTool}
    <p class="text-muted text-[12px]">{bulkCopy.precisionNote[store.precision]}</p>
  {/if}

  {#if store.cap === LOW_CORE_CAP}
    <p class="text-muted text-[12px]" data-testid="sim-low-core">{bulkCopy.lowCoreNote(store.cap)}</p>
  {/if}

  {#if store.capNotice !== null}
    <p class="text-strong text-[13px]" role="alert" data-testid="sim-cap-notice">
      {bulkCopy.capNotice(store.capNotice.cap, store.capNotice.combinations)}
      {#if store.serverCapNotice === null}
        <!-- Contract 10.1 A2: the premium lane's cap is 5,000, and bulkCopy.capPremiumNote
             says that number. Offering it when the list is past 5,000 too would send the
             player to a lane that refuses the same request. -->
        <span class="text-muted">{bulkCopy.capPremiumNote}</span>
      {:else}
        <span class="text-muted"
          >{bulkCopy.serverCapNotice(store.serverCapNotice.cap, store.serverCapNotice.combinations)}</span
        >
      {/if}
    </p>
  {/if}

  {#if store.progressLine !== ''}
    <p class="tabular text-muted font-mono text-[13px]" data-testid="sim-stage-progress">
      {store.progressLine}
    </p>
  {/if}

  {#if store.message !== null}
    <p class="text-strong text-[13px]" role="alert" data-testid="sim-message">
      {store.message}
      {#if store.detail !== ''}
        <span class="text-muted block font-mono text-[12px]">{store.detail}</span>
      {/if}
    </p>
  {/if}
</section>

<!-- Design 8: the exact JSON the run will send, editable and validated by the engine's own
     Validate through the wasm's simValidate (contract 10.2). Part A owns the component;
     mounting it here is what gives every tool page the same escape hatch. Apply and Run
     both go through the store's own request methods (bulk-store-request.ts). `canShare` is
     false: no tool page reads `?req=` back (only /sim does), so the Share affordance is
     hidden entirely rather than the drawer inventing a message for a feature this page does
     not have. -->
<RequestDrawer
  request={store.requestPreview}
  disabled={running}
  onvalidate={(json) => store.validateRequest(json)}
  onapply={(next) => store.applyRequest(next)}
  onrun={(next) => void store.runRequest(next)}
  canShare={false}
/>
