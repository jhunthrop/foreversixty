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

  let { live, href }: { live: LiveDps; href: string } = $props();

  const stale = $derived(live.state === 'pending' || live.state === 'running');
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
</script>

<div class="flex flex-col gap-1">
  <span class="label text-muted">{simCopy.plannerDpsLabel}</span>
  <span
    class={`tabular font-mono text-[20px] leading-11 ${
      stale || live.state === 'error' || live.estimate.mean === 0 ? 'text-muted' : 'text-gold'
    }`}
    data-testid="planner-dps"
  >
    {figure}
  </span>
  {#if band !== ''}
    <span class="tabular text-muted font-mono text-[12px]" data-testid="planner-dps-error">{band}</span>
  {/if}
</div>

<!-- min-h-11 at every width, the same as GearPanel.svelte's slot buttons: no other control
     in this lane shrinks its target on desktop, and a summary-bar row is exactly where a
     mouse-only "it's fine above 44px on desktop" argument would first break the pattern. -->
<a class="label text-nav flex min-h-11 items-center underline" {href} data-testid="planner-sim-link">
  {simCopy.simThisBuild}
</a>
