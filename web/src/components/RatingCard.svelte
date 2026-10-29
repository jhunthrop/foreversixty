<!-- web/src/components/RatingCard.svelte -->
<!-- The Performance Rating in the same card shell as the Forever Score (fix round 2 item 6),
     with that card's treatment (L8.5 ruling 1): "65 of 100" and the six-segment bar when any
     segment has a score; one line saying the number is the latest fight's and how many fights
     are rated ("Latest fight rated 65 of 100 · 6 fights rated", fix round 1 items 1 and 2: never
     an average, and no claim about the Forever Score's Parses); the trend as the score card's sparkline (the shared
     Sparkline figure and lib/score/sparkline.ts's model), "Up 9 since 4 Dec" under it; and one
     line naming the weakest part of a fight, whose name links to that fight's row in "Every
     ranked fight" when the page has one. From 1280 two columns: the number and that line
     left, the sparkline filling the right; a caller whose container is narrow at every window
     width (/ratings' 560px sample) turns the split off with `columns={false}` (fix round 1
     item 5). The profile's CharacterRatingPanel mounts it with the
     fetched rating; /ratings mounts it as a sample. Plain data in. -->
<script lang="ts">
  import { componentLabel, MIN_TREND_SAMPLES, ratingCopy } from '../lib/rating/copy';
  import { RATING_COMPONENT_ORDER, type RatingComponent } from '../lib/rating/types';
  import { scoreCopy } from '../lib/score/copy';
  import { ratingSparkline } from '../lib/score/sparkline';
  import { CARD } from '../lib/ui/card';
  import Sparkline from './Sparkline.svelte';

  let {
    overall,
    components = [],
    trend,
    sampleSize,
    worst = '',
    weakestHref = null,
    latestFight = null,
    sample = false,
    columns = true,
  }: {
    overall: number | null;
    components?: RatingComponent[];
    trend: { fought_at: string; overall: number }[];
    /** Rated fights on record; under MIN_TREND_SAMPLES the sparkline gives way to a line. */
    sampleSize: number;
    /** The weakest component's name, or '' for none. */
    worst?: string;
    /** The weakest part's fight's row ("#fight-<report>-<n>"), or null for no link. */
    weakestHref?: string | null;
    /** The latest rated fight's encounter and "9 Dec" date, named in the line (fix round 2
     *  item 1); null names no fight. */
    latestFight?: { encounter: string | null; date: string | null } | null;
    sample?: boolean;
    /** Two columns from 1280; false keeps one column for a narrow container. */
    columns?: boolean;
  } = $props();

  const PANEL = `${CARD} flex flex-col gap-3 p-4 md:p-5`;
  /** Underlined; the gold is the base `a` rule's, so hover and pressed take --gold-hover
   *  (L8.4 fix round 5 item 2). */
  const LINK = 'inline-flex min-h-11 items-center underline underline-offset-2 md:min-h-0';
  /** The link inside a sentence, the score card's "Most to gain: Gear" one: inline, its 44px
   *  phone target vertical padding round the text, stacked above the definition under it. */
  const SENTENCE_LINK = 'relative z-10 underline underline-offset-2 max-md:py-[13.5px]';

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
  const line = $derived(ratingSparkline(trend));
  const rated = $derived(overall === null ? null : Math.round(overall));
</script>

<div class={PANEL} data-testid="rating-card">
  <div
    class={columns ? 'flex flex-col gap-3 xl:grid xl:grid-cols-2 xl:gap-x-8' : 'flex flex-col gap-3'}
    data-testid="rating-card-columns"
  >
    <div class="flex flex-col gap-3" data-testid="rating-card-main">
      {#if overall !== null}
        <div class="flex flex-wrap items-center gap-x-4 gap-y-2">
          <span class="flex items-baseline gap-2">
            <span
              class="text-gold tabular text-[36px] leading-none font-bold md:text-[48px]"
              data-testid="character-rating-overall">{Math.round(overall)}</span
            >
            <span class="text-muted numeric text-[14px]">{ratingCopy.outOf100}</span>
          </span>
          {#if segments.some((segment) => segment.width > 0)}
            <span
              class="flex w-32 shrink-0 gap-[2px]"
              aria-hidden="true"
              data-testid="character-rating-segments"
            >
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
      {#if rated !== null}
        <!-- Fix round 2 item 7: three spans; "rated 65 of 100" never breaks inside, so a long
             fight name breaks before "rated"; under 360px the dot goes and "6 fights rated" takes
             a line of its own. -->
        <p class="text-[14px] leading-[22px]" data-testid="rating-card-verdict">
          <span class="text-text">{ratingCopy.verdictLead(latestFight)}</span>
          <span class="whitespace-nowrap">{ratingCopy.verdictRated(rated)}</span><span
            class="max-[360px]:hidden">{ratingCopy.verdictSeparator}</span
          ><span class="whitespace-nowrap max-[360px]:block">{ratingCopy.verdictCount(sampleSize)}</span>
        </p>
        <p class="text-muted text-[12px]" data-testid="rating-card-parse-vs-rating">
          {ratingCopy.parseVsRating}
        </p>
      {/if}
    </div>
    <div class="flex flex-col gap-1" data-testid="rating-card-side">
      {#if sampleSize < MIN_TREND_SAMPLES}
        <p class="text-muted text-[13px]" data-testid="character-rating-too-few">
          {ratingCopy.trendTooFew(sampleSize)}
        </p>
      {:else if line !== null}
        <Sparkline
          {line}
          class="w-full"
          testids={{
            figure: 'rating-card-history',
            plot: 'rating-card-plot',
            svg: 'rating-trend',
            last: 'rating-card-sparkline-last',
            dates: 'rating-card-history-dates',
            caption: 'rating-card-history-caption',
            lastLabel: 'rating-card-history-last',
            table: 'rating-card-history-table',
          }}
          tableCaption={ratingCopy.trendTableCaption}
          dateHeader={ratingCopy.trendTableDate}
          valueHeader={ratingCopy.trendTableRating}
        />
      {/if}
    </div>
  </div>
  <div class="flex flex-wrap items-start justify-between gap-x-4 gap-y-1">
    {#if worst !== ''}
      {@const definition = ratingCopy.componentDefinition(worst)}
      <!-- L8.4 fix round 4 item 3: the part named is defined under the line, a muted 12px caption
           at every width, as the fight tables define Execution. -->
      <div class="flex flex-col gap-1">
        <p class="text-[14px] leading-[22px]" data-testid="rating-card-weakest">
          {#if weakestHref === null}
            {ratingCopy.weakestPart(componentLabel(worst))}
          {:else}
            {ratingCopy.weakestPartLead}
            <a class={SENTENCE_LINK} href={weakestHref}>{componentLabel(worst)}</a>{ratingCopy.weakestPartEnd}
          {/if}
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
