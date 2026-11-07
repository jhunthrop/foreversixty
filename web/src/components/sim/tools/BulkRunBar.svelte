<!-- web/src/components/sim/tools/BulkRunBar.svelte -->
<!-- The run line every tool page shares (mock design/mocks/ws_simtools.py): the quick sim's own
     ControlBar (SimRunBar: fight, length, targets, buffs, precision, More settings and the one
     primary; a payer's "Run on our servers" is in More settings, as on /sim), with the tool's precision choices and words, then the standing lines under it: the
     combination count, what the precision costs, the cap and anything that blocks a run. The cap
     notice never trims the list: it says what would exceed the cap and by how much, and offers the
     lane that would allow it. More settings is the quick sim's sheet (target, execute, buffs, the
     raw request). -->
<script lang="ts">
  import { PREMIUM_LINKS } from '../../../lib/billing/premium-links';
  import type { BulkStore } from '../../../lib/sim/bulk-store.svelte';
  import { BULK_PRECISIONS, LOW_CORE_CAP, type Precision } from '../../../lib/sim/bulk-types';
  import { bulkCopy, simCopy } from '../../../lib/sim/copy';
  import { poolQualityCopy } from '../../../lib/sim/pool-quality-copy';
  import { isSimulatedSpec, specDisplayName } from '../../../lib/sim/spec-label';
  import { weightsEngineIterations, weightsIterationsFor } from '../../../lib/sim/weights';
  import type { PrimaryAction } from '../../../lib/sim/answer';
  import type { PrecisionId } from '../../../lib/sim/precision';
  import { LINK } from '../../../lib/ui/primitives';
  import MoreSettingsSheet from '../MoreSettingsSheet.svelte';
  import SimRunBar from '../SimRunBar.svelte';

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
   * Task 3 (healer review BLOCKER, `/sim/weights` "fails silently" for 64s; MAJOR on the
   * other three tools, a raw Go error after "setting up a run"): every tool page mounts
   * this bar, so gating the one run button here covers `/sim/gear`, `/sim/drops`,
   * `/sim/talents` and `/sim/weights` at once, the same way `isSimulatedSpec` gates
   * RunControl.svelte on `/sim` itself. `true` while no character is loaded -- the existing
   * `store.character === null` guard on the button already covers that case, and this must
   * not additionally disable the empty-character state with a spec name it does not have.
   */
  const notSimulated = $derived(store.character !== null && !isSimulatedSpec(store.character.spec));
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
   * Bare "Fast"/"Normal"/"High" for every tool now, combination and weights alike --
   * `precisionNote`/`weightsCostNote` below say what each actually costs. Task 8, sub-item
   * 2: the weights tool used to borrow `/sim`'s own fixed-count labels ("Fast, 500
   * iterations"), but that number was never the real cost of a weights sweep (contract
   * 10.9's per-stat multiplier -- `weights.ts`'s own `weightsEngineIterations`), so the
   * label was quietly wrong by up to 68x. `weightsCostNote` states the true total instead;
   * the option text itself no longer claims a number it cannot back up.
   */
  const precisionOptions = BULK_PRECISIONS.map((id) => ({ value: id, label: PRECISION_LABELS[id] }));
  /**
   * Task 8, sub-item 2: the real engine cost of the currently selected precision, for
   * however many stats are ticked right now -- `store.stats` is generic on `BulkStore`
   * (every tool has one), so this reads fine even before `combinationTool` gates the
   * note below to weights only. Always the browser lane: `weightsIterationsFor`'s server
   * number is `runOnServer()`'s own concern, not this free, default button's.
   */
  const weightsCost = $derived(
    weightsEngineIterations(weightsIterationsFor(store.precision, 'browser'), store.stats.length),
  );

  let moreOpen = $state(false);
  const runDisabled = $derived(
    store.capNotice !== null ||
      store.character === null ||
      store.phase === 'loading-character' ||
      (nothingToRun && !running) ||
      (notSimulated && !running) ||
      (store.pendingCustomBuild && !running),
  );
  const primary = $derived<PrimaryAction>(running ? 'stop' : runDisabled ? 'disabled' : 'run');
  const primaryLabel = $derived(
    running ? bulkCopy.stopBulk : store.result !== null ? bulkCopy.runBulkAgain : bulkCopy.runBulk,
  );
  const NOTE = 'text-small text-muted m-0';
  const BLOCK = 'm-0 text-[14px] leading-5 text-strong';
</script>

<SimRunBar
  settings={store.settings}
  precisionId={store.precision as PrecisionId}
  spec={store.character?.spec ?? ''}
  level={store.character?.level ?? 60}
  {primary}
  {primaryLabel}
  primaryTestid="sim-run-bulk"
  testid="sim-run-bulk-bar"
  busy={running}
  {precisionOptions}
  onsettings={(next) => store.setSettings(next)}
  onprecision={(value) => store.setPrecision(value as Precision)}
  onprimary={() => (running ? store.stop() : void store.run())}
  onmore={() => (moreOpen = true)}
/>

<div class="flex flex-col gap-1 px-[var(--page-inset)] pt-3" data-testid="sim-run-notes">
  <p class="m-0 flex flex-wrap items-baseline gap-x-3 text-[14px] leading-5">
    {#if combinationTool && countLabel !== ''}<span
        class="text-strong font-semibold tabular-nums"
        data-testid="sim-combo-count">{countLabel}</span
      >{/if}
    {#if combinationTool}
      <span class="text-muted">{bulkCopy.precisionNote[store.precision]}</span>
    {:else}
      <span class="text-muted" data-testid="sim-weights-cost-note"
        >{bulkCopy.weightsCostNote(weightsCost)}</span
      >
    {/if}
  </p>
  {#if store.cap === LOW_CORE_CAP}
    <p class={NOTE} data-testid="sim-low-core">{bulkCopy.lowCoreNote(store.cap)}</p>
  {/if}
  {#if notSimulated}
    <p class={BLOCK} data-testid="sim-run-not-simulated">
      {simCopy.runNotSimulated(specDisplayName(store.character?.spec ?? ''))}
    </p>
  {/if}
  {#if store.pendingCustomBuild}
    <p class={BLOCK} data-testid="sim-run-pending-custom">{poolQualityCopy.talentsPendingCustomBuild}</p>
  {/if}
  {#if store.capNotice !== null}
    <p class={BLOCK} role="alert" data-testid="sim-cap-notice">
      {bulkCopy.capNotice(store.capNotice.cap, store.capNotice.combinations)}
      {#if store.serverCapNotice === null}
        <span class="text-muted"
          >{bulkCopy.capPremiumNote}
          <a href={PREMIUM_LINKS.serverSims} class={LINK}>{bulkCopy.seePlans}</a></span
        >
      {:else}
        <span class="text-muted"
          >{bulkCopy.serverCapNotice(store.serverCapNotice.cap, store.serverCapNotice.combinations)}</span
        >
      {/if}
    </p>
  {/if}
  {#if store.progressLine !== ''}
    <p class="{NOTE} tabular-nums" data-testid="sim-stage-progress">{store.progressLine}</p>
  {/if}
  {#if store.message !== null}
    <p class="text-wipe m-0 text-[14px] leading-5" role="alert" data-testid="sim-message">
      {store.message}
      {#if store.detail !== ''}<span class="text-muted block text-[13px]">{store.detail}</span>{/if}
    </p>
  {/if}
</div>

{#if store.character !== null}
  <MoreSettingsSheet
    open={moreOpen}
    settings={store.settings}
    spec={store.character.spec}
    level={store.character.level}
    build={store.character.tree_version}
    classSlug={store.character.class_slug}
    names={store.simBuffs}
    request={store.requestPreview}
    disabled={running || store.serverCapNotice !== null}
    premium={store.premium}
    onserver={() => {
      moreOpen = false;
      void store.runOnServer();
    }}
    onclose={() => (moreOpen = false)}
    onchange={(next) => store.setSettings(next)}
    onvalidate={(json) => store.validateRequest(json)}
    onapply={(next) => store.applyRequest(next)}
    onrun={(next) => {
      moreOpen = false;
      void store.runRequest(next);
    }}
    onshare={() => null}
  />
{/if}
