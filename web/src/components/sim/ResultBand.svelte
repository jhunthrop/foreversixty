<!-- web/src/components/sim/ResultBand.svelte -->
<!-- The saved run's band (spec 2026-11-04-sim-results 3.3): who ran it on the left, the page's
     one answer on the right from 1024 (under it below that), the one primary action, and the
     quiet facts row under both. Full bleed on the site's band ground; the content sits in the
     page column. The band's figure is the page's one primary element. -->
<script lang="ts">
  import { onDestroy } from 'svelte';
  import { classColorVar } from '../../lib/report/format';
  import { afterSimSentence } from '../../lib/sim/after-sim-copy';
  import type { LastUpgrade } from '../../lib/sim/last-upgrade';
  import { bandDescriptor, bandError, bandFacts, bandHeading } from '../../lib/sim/result-band';
  import { bandCopy, oneDecimal } from '../../lib/sim/result-copy';
  import type { SimResult } from '../../lib/sim/types';
  import { button } from '../../lib/ui/button';
  import { createCopyFeedback, type CopyFeedbackState } from '../../lib/ui/copy-feedback';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import ActionLink from '../ui/ActionLink.svelte';
  import Glyph from '../ui/Glyph.svelte';

  let {
    result,
    onrerun,
    raidBuffs = null,
    topUpgrade = null,
    live = false,
    figureElement = $bindable(),
  }: {
    result: SimResult;
    /** Opens /sim with this run's whole setup; the page never runs the sim for the visitor.
     *  Absent on a live run, where the band carries no actions. */
    onrerun?: () => void;
    /** A finished run on /sim: the page's h1 is its own, and its run control and save panel are
     *  the actions, so the band drops "Open in the sim" and "Copy link". */
    live?: boolean;
    /** The spec's raid preset buffs, null until the preset table's chunk has loaded. */
    raidBuffs?: readonly string[] | null;
    /** A Top Gear upgrade on record, when one is. */
    topUpgrade?: LastUpgrade | null;
    /** The figure's element, for the section nav's compact figure. */
    figureElement?: HTMLElement;
  } = $props();

  /** How long "Copy failed" shows before the label returns. */
  const COPY_FAILED_MS = 3000;

  let copyState = $state<CopyFeedbackState>('idle');
  let failedTimer: ReturnType<typeof setTimeout> | undefined;
  const copier = createCopyFeedback(
    (text) => navigator.clipboard.writeText(text),
    (next) => {
      copyState = next;
      clearTimeout(failedTimer);
      if (next === 'failed') failedTimer = setTimeout(() => (copyState = 'idle'), COPY_FAILED_MS);
    },
  );
  onDestroy(() => {
    copier.dispose();
    clearTimeout(failedTimer);
  });

  const copyLabel = $derived.by(() => {
    if (copyState === 'copied') return bandCopy.copied;
    return copyState === 'failed' ? bandCopy.copyFailed : bandCopy.copyLink;
  });

  const heading = $derived(bandHeading(result));
  const [who, ...rest] = $derived(bandDescriptor(result));
  const error = $derived(bandError(result));
  const facts = $derived(bandFacts(result, raidBuffs));
  const hasFigure = $derived(result.dps.mean > 0);
  const classColor = $derived(classColorVar(result.request.character.class));
  // /sim's run control already owns `sim-dps` and `sim-error` above a live band.
  const figureTestid = $derived(live ? 'sim-report-dps' : 'sim-dps');
  const errorTestid = $derived(live ? 'sim-report-error' : 'sim-error');
</script>

<section class="full-bleed surface-band" data-testid="sim-result-band">
  <div class="page-column px-[18px] pt-5 md:pt-7">
    <div class="flex flex-col gap-6 lg:flex-row lg:items-end lg:justify-between lg:gap-10">
      <div class="flex min-w-0 items-center gap-4 md:gap-5">
        <ClassCrestRing
          characterClass={result.request.character.class}
          size={64}
          class="h-14 w-14 md:h-16 md:w-16"
          testid="sim-result-crest"
        />
        <div class="flex min-w-0 flex-col gap-0.5">
          <p class="label text-gold" data-testid="sim-result-eyebrow">
            {live ? bandCopy.liveEyebrow(result.lane) : bandCopy.eyebrow}
          </p>
          <svelte:element
            this={live ? 'h2' : 'h1'}
            class="heading-page"
            style:color={classColor}
            data-testid="sim-saved-title">{heading}</svelte:element
          >
          <p class="text-muted text-[13px] leading-[22px]" data-testid="sim-result-descriptor">
            {#if who !== undefined && who !== ''}<span class="text-text font-semibold">{who}</span>{/if}
            {#each rest as piece (piece)}
              <span class="before:mx-1 before:content-['·']">{piece}</span>
            {/each}
          </p>
          <p
            class="text-text mt-1 flex flex-wrap items-center gap-x-1 text-[14px]"
            data-testid="sim-next-action"
          >
            {#if topUpgrade === null}
              {bandCopy.nextStep}
              <ActionLink href={bandCopy.nextStepHref}>{bandCopy.nextStepLink}</ActionLink>
            {:else}
              {afterSimSentence(topUpgrade)}
            {/if}
          </p>
        </div>
      </div>

      <div class="flex flex-col lg:items-end lg:text-right">
        <p class="flex items-baseline gap-2 lg:justify-end" data-testid="sim-result-figure">
          <span
            bind:this={figureElement}
            class={`font-heading tabular text-[56px] leading-[56px] font-extrabold md:text-[64px] md:leading-[60px] ${hasFigure ? 'gold-text' : 'text-muted'}`}
            data-testid={figureTestid}>{hasFigure ? oneDecimal(result.dps.mean) : '–'}</span
          >
          <span class="font-heading text-muted text-[18px] font-bold tracking-[.04em]">{bandCopy.dps}</span>
        </p>
        {#if error !== null}
          <p class="text-muted mt-0.5 text-[15px]" data-testid={errorTestid}>
            <span class="text-text font-semibold">{error.margin}</span> · {error.percent}
            {bandCopy.error}
          </p>
        {/if}
        {#if !live && onrerun !== undefined}
          <div class="mt-3.5 flex flex-col gap-3 md:flex-row lg:justify-end">
            <button type="button" class={button('primary')} onclick={onrerun} data-testid="sim-run-yourself">
              <Glyph name="play" />
              {bandCopy.openInSim}
            </button>
            <button
              type="button"
              class={button('secondary')}
              onclick={() => void copier.copy(window.location.href)}
              data-testid="sim-saved-copy"
            >
              <Glyph name={copyState === 'copied' ? 'check' : 'link'} />
              <span aria-live="polite">{copyLabel}</span>
            </button>
          </div>
        {/if}
      </div>
    </div>

    <dl
      class="mt-6 grid grid-cols-2 border-t border-[rgba(229,185,85,.14)] md:flex md:flex-wrap"
      data-testid="sim-result-facts"
    >
      {#each facts as fact (fact.id)}
        <div
          class="border-line-soft flex flex-col border-b py-3 md:border-r md:border-b-0 md:pr-6 md:not-first:pl-6 md:last:border-r-0"
          data-testid={`sim-fact-${fact.id}`}
        >
          <dt class="text-muted text-[13px] leading-[18px] font-semibold">{fact.label}</dt>
          <dd class="text-strong numeric m-0 text-[15px] leading-[22px] font-semibold">{fact.value}</dd>
        </div>
      {/each}
    </dl>
  </div>
</section>
