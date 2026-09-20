<!-- web/src/components/sim/RunControl.svelte -->
<!-- One button and one number. Everything here exists to keep that number honest while it
     moves: the band beside it is the 95% confidence interval, the line under it says how
     many iterations bought that band, and the figure is tabular so the digits do not
     jitter as it refines ten times a second.
     aria-live is on the message and never on the figure: a live region on a number that
     changes that often makes a screen reader unusable. -->
<script lang="ts">
  import { confidenceBand, formatMargin } from '../../lib/sim/estimate';
  import { simCopy } from '../../lib/sim/copy';
  import { percentLabel } from '../../lib/sim/details';
  import { LANE_ITERATION_CEILING, PRECISIONS, type Lane, type PrecisionId } from '../../lib/sim/precision';
  import { isSimulatedSpec, specDisplayName } from '../../lib/sim/spec-label';
  import type { SimPhase } from '../../lib/sim/store.svelte';
  import type { Estimate } from '../../lib/sim/types';
  import { engineLabel } from '../../lib/sim/version';
  import HelpNote from './HelpNote.svelte';
  import PrecisionSelect from './PrecisionSelect.svelte';

  let {
    spec,
    phase,
    estimate,
    iterationsDone,
    iterationsTotal,
    precisionId,
    relativeError,
    lane,
    premium,
    message,
    detail,
    racePending,
    staleVersion,
    serverRunning,
    onrun,
    onstop,
    onprecision,
    onserver,
    onrerun,
  }: {
    /** The loaded character's spec. Task 3 (healer review): a healer or tank spec disables
     *  the button and shows the honest line beside it instead of the run this control
     *  would otherwise offer. */
    spec: string;
    phase: SimPhase;
    estimate: Estimate;
    iterationsDone: number;
    iterationsTotal: number;
    precisionId: PrecisionId;
    /** `error / mean` of the figure on screen, for the progress line's per cent. */
    relativeError: number;
    /**
     * Which lane the target-error note's ceiling names. `store.lane` -- reflects the lane
     * the figure on screen most recently ran on, browser by default -- so a premium player
     * who just ran on the server sees that lane's own ceiling (100,000), not the browser's,
     * pinned regardless of which button they are looking at.
     */
    lane: Lane;
    premium: boolean;
    message: string | null;
    /** The engine's own words for the failure, shown verbatim under the message. Empty when there are none. */
    detail: string;
    /**
     * True while the loaded character still owes the strip a race (`store.needsRace`). It
     * is a prop rather than a `SimPhase` because it is orthogonal to the run's own state:
     * the page is `idle` either way, and a sixth phase would have to be re-entered after
     * every run. The run control is disabled while it is true.
     */
    racePending: boolean;
    staleVersion: string | null;
    /**
     * True while `store.runOnServer()` is in flight (fix round 1). Kept apart from
     * `SimPhase` the way `racePending` is: a server run and a browser run are two
     * independent lanes, so this cannot just be another value of `phase`. The primary
     * button is disabled and relabelled while it is true -- the server lane has no cancel
     * path yet, and a button that reads "Stop" over a click that does nothing is worse
     * than one that is honestly unavailable.
     */
    serverRunning: boolean;
    onrun: () => void;
    onstop: () => void;
    onprecision: (value: PrecisionId) => void;
    onserver: () => void;
    onrerun: () => void;
  } = $props();

  const running = $derived(phase === 'running');
  const loadingEngine = $derived(phase === 'loading-engine');
  const hasFigure = $derived(estimate.mean > 0);
  const simulated = $derived(isSimulatedSpec(spec));

  const figure = $derived(hasFigure ? Math.round(estimate.mean).toLocaleString('en-US') : '—');
  const band = $derived(hasFigure && estimate.error > 0 ? `± ${formatMargin(confidenceBand(estimate))}` : '');
  const percent = $derived(
    iterationsTotal > 0 ? Math.min(100, Math.round((iterationsDone / iterationsTotal) * 100)) : 0,
  );
  const showBar = $derived(running || loadingEngine || serverRunning || phase === 'done');
  const label = $derived(
    running
      ? simCopy.stop
      : loadingEngine
        ? simCopy.engineLoadingButton
        : serverRunning
          ? simCopy.serverRunButton
          : phase === 'done'
            ? simCopy.runAgain
            : simCopy.run,
  );
  const errorText = $derived(relativeError > 0 ? ` · ${percentLabel(relativeError)}` : '');
  const progressLine = $derived(
    phase === 'done'
      ? `${iterationsDone.toLocaleString('en-US')} ${simCopy.iterations}${errorText}`
      : `${iterationsDone.toLocaleString('en-US')} of ${iterationsTotal.toLocaleString('en-US')} ${simCopy.iterations}${errorText}`,
  );
</script>

<section
  class="border-line bg-raised rounded-panel mx-[18px] flex flex-col gap-4 border p-4 md:mx-0 md:flex-row md:items-center md:gap-6"
  data-testid="sim-run"
>
  <button
    type="button"
    class="border-line-warm-strong rounded-control bg-card-top text-strong label min-h-11 w-full border px-5 disabled:opacity-50 md:w-auto md:min-w-[9rem]"
    disabled={loadingEngine ||
      phase === 'loading-character' ||
      (racePending && !running) ||
      (serverRunning && !running) ||
      (!simulated && !running)}
    onclick={() => (running ? onstop() : onrun())}
    data-testid="sim-run-button"
  >
    {label}
  </button>

  <div class="flex min-w-0 flex-1 flex-col gap-1">
    <div class="flex flex-wrap items-baseline gap-2">
      <span
        class={`tabular font-mono text-[32px] leading-none ${hasFigure ? 'text-gold' : 'text-muted'}`}
        data-testid="sim-dps"
      >
        {figure}
      </span>
      <span class="tabular text-muted font-mono text-[14px]" data-testid="sim-error">{band}</span>
      <span class="label text-muted">{simCopy.dps}</span>
      {#if staleVersion}
        <span class="pill pill-sample" data-testid="sim-stale-pill">{engineLabel(staleVersion)}</span>
      {/if}
    </div>

    {#if showBar}
      <div
        class="bg-line-soft h-[6px] w-full overflow-hidden rounded-full"
        role="progressbar"
        aria-valuemin="0"
        aria-valuemax="100"
        aria-valuenow={percent}
        aria-label={simCopy.progressLabel}
        data-testid="sim-progress-bar"
      >
        <div class="bg-gold h-full transition-[width] duration-120" style={`width:${percent}%`}></div>
      </div>
      <p class="text-muted tabular font-mono text-[12px]" data-testid="sim-progress">{progressLine}</p>
    {/if}

    {#if loadingEngine}
      <p class="text-muted text-[12px]" data-testid="sim-engine-loading">{simCopy.engineLoading}</p>
    {/if}
  </div>

  <div class="flex flex-wrap items-end gap-3">
    <div class="flex flex-col gap-1">
      <label class="flex flex-col gap-1">
        <span class="label text-muted">{simCopy.precision}</span>
        <PrecisionSelect
          class="border-line-warm rounded-control bg-raised text-text min-h-11 min-w-0 border px-3 text-[14px] font-semibold md:min-h-9"
          value={precisionId}
          options={PRECISIONS}
          labelFor={(id) => simCopy.precisionLabel[id] ?? id}
          disabled={running || serverRunning}
          onchange={(value) => onprecision(value as PrecisionId)}
        />
      </label>
      <HelpNote label={simCopy.precision} id="sim-precision" disabled={running || serverRunning}>
        <p>{simCopy.precisionHelp}</p>
      </HelpNote>
    </div>
    {#if premium}
      <button
        type="button"
        class="border-line-warm rounded-control text-nav label min-h-11 border px-4 md:min-h-9"
        disabled={running || serverRunning || !simulated}
        onclick={onserver}
        data-testid="sim-server-run">{simCopy.runOnServers}</button
      >
    {/if}
  </div>

  {#if !simulated}
    <p class="text-strong order-last w-full text-[13px]" data-testid="sim-run-not-simulated">
      {simCopy.runNotSimulated(specDisplayName(spec))}
    </p>
  {/if}

  {#if precisionId === 'target-error'}
    <p class="text-muted order-last w-full text-[12px]" data-testid="sim-target-error">
      {simCopy.targetErrorNote(LANE_ITERATION_CEILING[lane].toLocaleString('en-US'))}
    </p>
  {/if}

  {#if staleVersion}
    <p class="order-last w-full text-[13px]" data-testid="sim-stale">
      <span class="text-strong">{simCopy.staleEngine}</span>
      <button type="button" class="ml-2 underline" onclick={onrerun} data-testid="sim-rerun">
        {simCopy.runAgain}
      </button>
    </p>
  {/if}

  {#if message}
    <p
      role="alert"
      aria-live="polite"
      class="text-strong order-last w-full text-[13px]"
      data-testid="sim-message"
    >
      {message}
    </p>
  {/if}

  <!-- The engine's own words, verbatim. sim/request refuses an unknown or ambiguous buff
       or consumable id rather than dropping it, and its message names the id -- which is
       the only thing here that says what to change. Never paraphrased, never hidden, and
       mono because it is a machine's sentence, not ours. -->
  {#if detail}
    <p class="text-muted order-last w-full font-mono text-[12px]" data-testid="sim-detail">
      {detail}
    </p>
  {/if}
</section>
