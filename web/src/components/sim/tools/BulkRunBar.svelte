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
  import { bulkCopy } from '../../../lib/sim/copy';
  import RequestDrawer from '../RequestDrawer.svelte';
  import SettingsBar from '../SettingsBar.svelte';

  let { store }: { store: BulkStore } = $props();

  const running = $derived(store.phase === 'running' || store.serverRunning);
  const countLabel = $derived(
    store.combinations === null ? bulkCopy.combinationsCounting : bulkCopy.combinations(store.combinations),
  );
  const PRECISION_LABELS: Record<Precision, string> = {
    fast: bulkCopy.precisionFast,
    normal: bulkCopy.precisionNormal,
    high: bulkCopy.precisionHigh,
  };

  function shareUnavailable(): null {
    // No tool page reads `?req=` yet (only /sim does) -- sharing a bulk/weights request by
    // URL is not part of this task's scope. Returning null renders the drawer's own
    // "too long" message, which is not quite the right words for "not supported here" but
    // is the only failure state the component already has words for; a dedicated message
    // is a request-drawer change, out of this task's file list.
    return null;
  }
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
    <span class="tabular text-strong font-mono text-[14px]" data-testid="sim-combo-count">
      {countLabel}
    </span>

    <label class="text-muted flex items-center gap-2 text-[12px]">
      {bulkCopy.precisionLabel}
      <select
        data-testid="sim-precision"
        class="border-line-warm rounded-control bg-bg text-text h-11 border px-2 text-[14px]"
        disabled={running}
        value={store.precision}
        onchange={(event) => store.setPrecision(event.currentTarget.value as Precision)}
      >
        {#each BULK_PRECISIONS as precision (precision)}
          <option value={precision}>{PRECISION_LABELS[precision]}</option>
        {/each}
      </select>
    </label>

    <button
      type="button"
      class="{SECONDARY_BUTTON} border-line-warm text-nav px-4"
      data-testid="sim-run-bulk"
      disabled={store.capNotice !== null || store.character === null}
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

  <p class="text-muted text-[12px]">{bulkCopy.precisionNote[store.precision]}</p>

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
     both go through the store's own request methods (bulk-store-request.ts); Share is not
     wired to a URL here (see `shareUnavailable` above). -->
<RequestDrawer
  request={store.requestPreview}
  disabled={running}
  onvalidate={(json) => store.validateRequest(json)}
  onapply={(next) => store.applyRequest(next)}
  onrun={(next) => store.runRequest(next)}
  onshare={shareUnavailable}
/>
