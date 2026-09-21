<!-- web/src/components/planner/PlannerDps.svelte -->
<!-- The fourth figure in the summary bar. It is the planner's own shape -- label caption,
     mono value at the same size -- so it reads as one more fact about the build rather
     than a widget bolted on.
     It dims rather than blanks while a new estimate runs: a number that disappears on every
     click is worse than one that is briefly a click behind. -->
<script lang="ts">
  import { confidenceBand, formatMargin } from '../../lib/sim/estimate';
  import { simCopy } from '../../lib/sim/copy';
  import type { LiveDps } from '../../lib/planner/live-dps.svelte';
  import type { LiveGate } from '../../lib/planner/live-gate';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';

  let {
    live,
    href,
    gate,
    pointsLeft,
    onshow,
  }: { live: LiveDps; href: string; gate: LiveGate; pointsLeft: number; onshow: () => void } = $props();

  // `off` with a figure still held is a build that stopped being simmed -- a point came out,
  // or the device has not been asked yet -- so it dims exactly as a run in flight does.
  const stale = $derived(live.state === 'pending' || live.state === 'running' || live.state === 'off');
  // An error blanks to em dash rather than dimming: the module does not clear `estimate` on
  // a failed run (a cancel keeps the figure that is still in flight), but a number left over
  // from a build the engine could not run this time -- most often a spec switch -- is not a
  // stale version of the current answer, it is an answer to a different question.
  const figure = $derived(
    live.state !== 'error' && live.estimate.mean > 0
      ? Math.round(live.estimate.mean).toLocaleString('en-US')
      : '—',
  );
  const band = $derived(
    live.state === 'ready' && live.estimate.error > 0
      ? `± ${formatMargin(confidenceBand(live.estimate))}`
      : '',
  );
  // The line under the figure says why there is no fresh number, when there is not one.
  const note = $derived(
    gate === 'unfinished'
      ? simCopy.plannerDpsPointsToGo(pointsLeft)
      : gate === 'ask'
        ? simCopy.plannerDpsShowNote
        : band,
  );
</script>

<!-- One column, like Level and Points beside it: caption, then the value row, then the line
     that qualifies it. "Sim this build" sits IN the value row rather than beside the column,
     and the bar aligns its columns to the top (SummaryBar.svelte), so this column being one
     line taller than its neighbours leaves every caption and every value on one baseline. -->
<div class="flex flex-col gap-1">
  <span class="label text-muted">{simCopy.plannerDpsLabel}</span>
  <div class="flex items-center gap-4">
    {#if gate === 'ask'}
      <!-- The same 44px the figure occupies, so asking moves nothing. -->
      <button
        type="button"
        class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-3"
        onclick={onshow}
        data-testid="planner-dps-show"
      >
        {simCopy.plannerDpsShow}
      </button>
    {:else}
      <span
        class={`tabular font-mono text-[20px] leading-11 ${
          stale || live.state === 'error' || live.estimate.mean === 0 ? 'text-muted' : 'text-gold'
        }`}
        data-testid="planner-dps"
      >
        {figure}
      </span>
    {/if}
    <!-- min-h-11 at every width, the same as GearPanel.svelte's slot buttons: no other
         control in this lane shrinks its target on desktop, and a summary-bar row is exactly
         where a mouse-only "it's fine above 44px on desktop" argument would first break the
         pattern. -->
    <a class="label text-nav flex min-h-11 items-center underline" {href} data-testid="planner-sim-link">
      {simCopy.simThisBuild}
    </a>
  </div>
  <!-- Always on the page, at a fixed height, even with nothing to say: the band only exists
       once a run is ready, and removing the line while the next run was pending made the
       whole summary bar one line shorter on every talent click, which moved the trees. -->
  <span
    class="tabular text-muted block h-[18px] font-mono text-[12px] leading-[18px]"
    data-testid="planner-dps-error">{note}</span
  >
</div>
