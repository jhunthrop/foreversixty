<!-- web/src/components/planner/PlannerDps.svelte -->
<!-- The fourth figure in the summary bar. It is the planner's own shape -- label caption,
     mono value at the same size -- so it reads as one more fact about the build rather
     than a widget bolted on.
     It dims rather than blanks while a new estimate runs: a number that disappears on every
     click is worse than one that is briefly a click behind. -->
<script lang="ts">
  import { confidenceBand, formatMargin } from '../../lib/sim/estimate';
  import { plannerSimCopy as simCopy } from '../../lib/sim/planner-copy';
  import type { LiveDps } from '../../lib/planner/live-dps.svelte';
  import type { LiveGate } from '../../lib/planner/live-gate';
  import { FACT_FIGURE_CLASS, PRIMARY_BUTTON_FIXED, SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { plannerScoreCopy } from '../../lib/planner/score-unit';
  import type { SlotScoreUnit } from '../../lib/bis/tank-view';
  import { poolQualityCopy } from '../../lib/sim/pool-quality-copy';
  import StatTile from '../ui/StatTile.svelte';

  let {
    live,
    href = '',
    gate,
    pointsLeft,
    onshow,
    ownGearCaveat = false,
    showSimLink = true,
    unfinishedNote = 'dynamic',
    unit = 'dps',
    tile = false,
  }: {
    live: LiveDps;
    /** Unused when `showSimLink` is false (the planner rebuild's header facts rail -- the
     *  link itself moved to the rail, above the Share panel, spec fix round 1 item 2.a). */
    href?: string;
    gate: LiveGate;
    pointsLeft: number;
    onshow: () => void;
    /**
     * newcomer round 4 (review.md:83-110): true only when this preview is TalentCandidates'
     * own inline ADD A BUILD editor (SummaryBar.svelte's `!standalone`), where the gear on
     * screen can be a pasted build's own -- imported wholesale by ImportBox's `onimport`
     * (`store.svelte.ts`'s `loadImported` replaces the whole gear map) -- rather than the
     * player's current gear the ranked comparison beside it locks every row to. Root cause
     * of the round-4 finding (Fire Mage: 681 ± 6.9 here, 400 in the ranked table): neither
     * number is wrong, this one is just never the player's own gear once a build is pasted.
     */
    ownGearCaveat?: boolean;
    /** False inside `PlannerHeadControls`'s own facts rail (fix round 1): "Sim this build" is now
     *  its own outlined gold button in the rail, above the Share panel, not bundled with the
     *  figure. Still true (default) for the non-standalone inline embed, unchanged. */
    showSimLink?: boolean;
    /** 'static' inside the header's facts rail: the dynamic "{n} points to go" count now
     *  lives under "Points left" instead (so the two never repeat the same number twice),
     *  and this slot reads a plain explainer of what the figure is waiting for. 'dynamic'
     *  (default) keeps the non-standalone embed's original behaviour. */
    unfinishedNote?: 'dynamic' | 'static';
    /** What the figure is called (the build's role): a healer's is HPS, a tank's a tank
     *  score. The live run asks the engine for the role's own figure (`role_metrics`), so a
     *  DPS number is never relabelled as HPS. */
    unit?: SlotScoreUnit;
    /** True inside the header's facts rail: the figure is a `StatTile` (heading-face
     *  figure, the note wrapping under it on a phone) instead of the embed's mono column. */
    tile?: boolean;
  } = $props();

  // `off` with a figure still held is a build that stopped being simmed -- a point came out,
  // or the device has not been asked yet -- so it dims exactly as a run in flight does.
  const stale = $derived(live.state === 'pending' || live.state === 'running' || live.state === 'off');
  // An error blanks to em dash rather than dimming: the module does not clear `estimate` on
  // a failed run (a cancel keeps the figure that is still in flight), but a number left over
  // from a build the engine could not run this time -- most often a spec switch -- is not a
  // stale version of the current answer, it is an answer to a different question.
  const figureText = $derived(
    live.state !== 'error' && live.estimate.mean > 0
      ? plannerScoreCopy.figure(live.estimate.mean, unit)
      : '—',
  );
  const band = $derived(
    live.state === 'ready' && live.estimate.error > 0
      ? `± ${formatMargin(confidenceBand(live.estimate))}`
      : '',
  );
  // The line under the figure says why there is no fresh number, when there is not one.
  const dpsNote = $derived(
    gate === 'unfinished'
      ? unfinishedNote === 'static'
        ? plannerScoreCopy.waiting
        : simCopy.plannerDpsPointsToGo(pointsLeft)
      : gate === 'ask'
        ? simCopy.plannerDpsShowNote
        : band,
  );
</script>

{#snippet askButton()}
  <!-- The same 44px the figure occupies, so asking moves nothing. -->
  <button
    type="button"
    class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-3"
    onclick={onshow}
    data-testid="planner-dps-show"
  >
    {plannerScoreCopy.show(unit)}
  </button>
{/snippet}

{#snippet figureValue(figureClass: string)}
  <!-- min-w-[7ch]: the row wrap reads this column's own width, which would otherwise shrink for
       the em dash and grow again for a real figure -- the same swap "keeps its height through
       a run" (planner-dps.spec.ts) guards against. -->
  <span
    class={`tabular inline-block min-w-[7ch] ${figureClass} ${
      stale || live.state === 'error' || live.estimate.mean === 0 ? 'text-muted' : 'text-gold'
    }`}
    data-testid="planner-dps"
    title={`${plannerScoreCopy.label(unit)}: ${plannerScoreCopy.ratePhrase(unit)}`}
  >
    {figureText}
  </span>
{/snippet}

{#if tile}
  <!-- Polish spec 5.1: a tile like its neighbours. The note wraps under the figure instead of
       running off the right edge (it is nowrap in the embed below); two lines are reserved
       below 640px, where the half-width tile wraps "Appears once your points are spent.",
       and one from there up, so the answer arriving moves nothing. -->
  <StatTile label={plannerScoreCopy.label(unit)} figureClass={FACT_FIGURE_CLASS}>
    {#snippet figure()}
      {#if gate === 'ask'}
        {@render askButton()}
      {:else}
        {@render figureValue('')}
      {/if}
    {/snippet}
    {#snippet note()}
      <span
        class="tabular text-muted block min-h-[36px] basis-full text-[13px] leading-[18px] sm:min-h-[18px] sm:basis-auto"
        data-testid="planner-dps-error">{dpsNote}</span
      >
    {/snippet}
  </StatTile>
{:else}
  <!-- One column, like Level and Points beside it: caption, then the value row, then the line
       that qualifies it. "Sim this build" sits IN the value row rather than beside the column,
       and the bar aligns its columns to the top (SummaryBar.svelte), so this column being one
       line taller than its neighbours leaves every caption and every value on one baseline. -->
  <div class="flex flex-col gap-1">
    <span class="quiet-label">{plannerScoreCopy.label(unit)}</span>
    <!-- A column below md, one row from md: with the figure and the button side by side on a
         phone, the runner's wider fallback font wrapped the pair in ways a Mac never showed
         (planner-dps.spec.ts read the button a row above the figure on CI). -->
    <div class="flex flex-col items-start gap-2 md:flex-row md:items-center md:gap-4">
      {#if gate === 'ask'}
        {@render askButton()}
      {:else}
        {@render figureValue('numeric text-[20px] leading-11')}
      {/if}
      <!-- The bar's one primary button: simming a build is one of the planner's own jobs
           (build review round 1, finding 5), a real action rather than an underlined link.
           PRIMARY_BUTTON_FIXED, not the shrinking PRIMARY_BUTTON: no other control in this
           lane shrinks its target on desktop. -->
      {#if showSimLink}
        <a class="{PRIMARY_BUTTON_FIXED} px-4" {href} data-testid="planner-sim-link">
          {simCopy.simThisBuild}
        </a>
      {/if}
    </div>
    <!-- Always on the page, at a fixed height, even with nothing to say: the band only exists
         once a run is ready, and removing the line while the next run was pending made the
         whole summary bar one line shorter on every talent click, which moved the trees. -->
    <span class="text-muted numeric block h-[18px] text-[12px] leading-[18px]" data-testid="planner-dps-error"
      >{dpsNote}</span
    >
    {#if ownGearCaveat}
      <span class="text-muted block text-[12px]" data-testid="planner-dps-own-gear">
        {poolQualityCopy.plannerDpsOwnGearNote}
      </span>
    {/if}
  </div>
{/if}
