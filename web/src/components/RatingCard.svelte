<!-- web/src/components/RatingCard.svelte -->
<!-- The Performance Rating in the same card shell as the Forever Score (fix round 2 item 6):
     "65 of 100", the six-segment bar when any segment has a score, the trend sparkline with a
     visible "Last 6 fights" label (the count is the data's), and one line naming the weakest
     part of a fight. The profile's CharacterRatingPanel mounts it with the fetched rating;
     /ratings mounts it as a sample. Plain data in. -->
<script lang="ts">
  import { componentLabel, MIN_TREND_SAMPLES, ratingCopy } from '../lib/rating/copy';
  import { RATING_COMPONENT_ORDER, type RatingComponent } from '../lib/rating/types';
  import { scoreCopy } from '../lib/score/copy';
  import RatingTrend from './RatingTrend.svelte';

  let {
    overall,
    components = [],
    trend,
    sampleSize,
    worst = '',
    sample = false,
  }: {
    overall: number | null;
    components?: RatingComponent[];
    trend: { fought_at: string; overall: number }[];
    /** Rated fights on record; under MIN_TREND_SAMPLES the sparkline gives way to a line. */
    sampleSize: number;
    /** The weakest component's name, or '' for none. */
    worst?: string;
    sample?: boolean;
  } = $props();

  const CARD =
    'bg-raised border-line rounded-panel flex flex-col gap-3 border p-4 shadow-[inset_0_1px_0_rgba(229,185,85,.35)] md:p-5';
  /** Underlined; the gold is the base `a` rule's, so hover and pressed take --gold-hover
   *  (L8.4 fix round 5 item 2). */
  const LINK = 'inline-flex min-h-11 items-center underline underline-offset-2 md:min-h-0';

  /** A segment's fill, 0 when the part was not scored. */
  function segmentWidth(part: RatingComponent | undefined): number {
    return part && !part.excluded && part.score !== null ? part.score : 0;
  }

  const segments = $derived(
    RATING_COMPONENT_ORDER.map((name) => ({
      name,
      width: segmentWidth(components.find((c) => c.name === name)),
    })),
  );
</script>

<div class={CARD} data-testid="rating-card">
  {#if overall !== null}
    <div class="flex flex-wrap items-center gap-x-4 gap-y-2">
      <span class="flex items-baseline gap-2">
        <span
          class="text-gold tabular text-[36px] leading-none font-bold md:text-[48px]"
          data-testid="character-rating-overall">{Math.round(overall)}</span
        >
        <span class="text-muted tabular font-mono text-[14px]">{ratingCopy.outOf100}</span>
      </span>
      {#if segments.some((segment) => segment.width > 0)}
        <span class="flex w-32 shrink-0 gap-[2px]" aria-hidden="true" data-testid="character-rating-segments">
          {#each segments as segment (segment.name)}
            <span class="bg-line-soft h-2 flex-1" title={componentLabel(segment.name)}
              ><span class="bg-gold block h-full" style={`width: ${segment.width}%`}></span></span
            >
          {/each}
        </span>
      {/if}
      {#if sample}<span class="pill pill-sample ml-auto">{scoreCopy.samplePill}</span>{/if}
    </div>
  {/if}
  {#if sampleSize < MIN_TREND_SAMPLES}
    <p class="text-muted text-[13px]" data-testid="character-rating-too-few">
      {ratingCopy.trendTooFew(sampleSize)}
    </p>
  {:else}
    <figure class="flex flex-col gap-1" data-testid="rating-card-trend">
      <figcaption class="label text-muted" data-testid="rating-card-trend-label">
        {ratingCopy.lastFights(trend.length)}
      </figcaption>
      <RatingTrend points={trend} />
    </figure>
  {/if}
  <div class="flex flex-wrap items-center justify-between gap-x-4 gap-y-1">
    {#if worst !== ''}
      {@const definition = ratingCopy.componentDefinition(worst)}
      <!-- L8.4 fix round 4 item 3: the part named is defined under the line, a muted 12px caption
           at every width, as the fight tables define Execution. -->
      <div class="flex flex-col gap-1">
        <p class="text-[14px] leading-[22px]" data-testid="rating-card-weakest">
          {ratingCopy.weakestPart(componentLabel(worst))}
        </p>
        {#if definition !== null}
          <p class="text-muted text-[12px]" data-testid="rating-card-weakest-definition">{definition}</p>
        {/if}
      </div>
    {/if}
    {#if !sample}
      <a class={`${LINK} ml-auto text-[14px]`} href="/ratings">{ratingCopy.explainLink}</a>
    {/if}
  </div>
</div>
