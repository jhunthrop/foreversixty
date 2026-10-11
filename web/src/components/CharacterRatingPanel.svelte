<!-- web/src/components/CharacterRatingPanel.svelte -->
<!-- The character page's rating panel. At the top, the Forever Score (spec 2026-10-11 §4.6):
     the badge, its tier and "as of" date, and the five parts as plain bars against their
     ceilings -- or the addon's no-logs line where the number would be, never a zero. It is
     plain data from the caller (Character.svelte already holds the GET /v1/characters/...
     answer that carries it). Beneath it, unchanged, the execution rating: overall score,
     six-segment bar, a trend sparkline once there are enough fights, and the best/worst
     component -- spec §6.2. That part fetches independently, the same self-contained
     pattern Character.svelte's own $effect already uses for fetchCharacter. -->
<script lang="ts">
  import { fetchCharacterRating, RankingsError } from '../lib/rankings/api';
  import { ratingCopy, componentLabel, MIN_TREND_SAMPLES } from '../lib/rating/copy';
  import { RATING_COMPONENT_ORDER } from '../lib/rating/types';
  import { scoreCopy } from '../lib/score/copy';
  import { COMPONENT_MAX, asOfText, hasScore, tierLabel, type ForeverScore } from '../lib/score/score';
  import type { CharacterPath } from '../lib/characters';
  import type { CharacterRating } from '../lib/rating/types';
  import RatingTrend from './RatingTrend.svelte';
  import ScoreBadge from './ScoreBadge.svelte';
  import EmptyState from './ui/EmptyState.svelte';

  let {
    path,
    apiBase = undefined,
    heading = true,
    forever = undefined,
  }: {
    path: CharacterPath;
    apiBase?: string;
    /** False when the mounting panel already titles it (the account page's "Your ratings"). */
    heading?: boolean;
    /** The character's Forever Score; null (or a null score) for a character the site has no
     *  logs of. Left out, the panel draws no score block at all (the account page). */
    forever?: ForeverScore | null;
  } = $props();

  /** A bar's fill, to one decimal place: 412 of 500 is 82.4%. */
  function share(points: number, max: number): number {
    return Math.round(Math.min(1, Math.max(0, points / max)) * 1000) / 10;
  }

  let data = $state<CharacterRating | null>(null);
  let status = $state<'loading' | 'ready' | 'hidden'>('loading');

  $effect(() => {
    const requested = path;
    status = 'loading';
    void fetchCharacterRating(requested, apiBase)
      .then((result) => {
        if (path !== requested) return;
        data = result;
        status = 'ready';
      })
      .catch((thrown: unknown) => {
        if (path !== requested) return;
        // Ruling 2: an anonymized character 404s. Every other failure also hides the
        // panel rather than showing an alarming error for what is, from a visitor's
        // seat, the same as "nothing here yet" -- the rest of the character page already
        // loaded and still has value.
        if (!(thrown instanceof RankingsError) || thrown.status !== 404) {
          // Non-anonymize failures still hide rather than alarm, but are worth a console
          // trace for whoever is watching the browser's own devtools during development.
        }
        data = null;
        status = 'hidden';
      });
  });
</script>

{#if forever !== undefined}
  <section class="flex flex-col gap-3" data-testid="forever-score">
    <h2 class="section-title text-[18px]">{scoreCopy.heading}</h2>
    {#if hasScore(forever)}
      {@const tier = tierLabel(forever.tier)}
      {@const asOf = asOfText(forever)}
      <div class="flex flex-wrap items-baseline gap-x-3 gap-y-1">
        <ScoreBadge score={forever.score} tier={forever.tier} size="lg" />
        {#if tier !== null}<span class="text-[14px] font-semibold">{tier}</span>{/if}
        {#if asOf !== null}<span class="text-muted text-[13px]">{asOf}</span>{/if}
      </div>
      {#if forever.components !== null}
        <ul class="flex max-w-[520px] flex-col gap-2">
          {#each scoreCopy.components as component (component.key)}
            {@const points = forever.components[component.key]}
            {@const max = COMPONENT_MAX[component.key]}
            {#if typeof points === 'number'}
              <li
                class="grid grid-cols-[88px_minmax(0,1fr)_88px] items-center gap-3 text-[13px]"
                data-testid={`forever-score-${component.key}`}
              >
                <span>{component.label}</span>
                <span class="bg-line-soft h-2" aria-hidden="true"
                  ><span class="bg-gold block h-full" style={`width: ${share(points, max)}%`}></span></span
                >
                <span class="text-muted tabular text-right font-mono whitespace-nowrap"
                  >{scoreCopy.ofMax(points, max)}</span
                >
              </li>
            {/if}
          {/each}
        </ul>
      {/if}
    {:else}
      <p class="text-muted text-[14px]" data-testid="forever-score-none">{scoreCopy.noLogs}</p>
    {/if}
    <a
      class="text-gold inline-flex min-h-11 w-fit items-center text-[12px] underline-offset-2 hover:underline md:min-h-0"
      href={scoreCopy.explainHref}>{scoreCopy.explainLink}</a
    >
  </section>
{/if}

{#if status === 'loading'}
  <!-- Reserves height for the common loading -> ready path so the panel's own container
       does not pop in and push "Best per encounter" down once the fetch resolves -- the
       plan's binding "no layout shift" constraint. Mirrors RatingPanel.svelte's own
       loading-skeleton idiom. The rare loading -> hidden path (an anonymized character)
       still fully collapses to nothing once resolved -- Ruling 2's own explicit
       requirement, which takes precedence over reserving permanent dead space for a
       panel that is supposed not to exist for that character. -->
  <section class="flex flex-col gap-2" data-testid="character-rating-loading" aria-hidden="true">
    {#if heading}<h2 class="section-title text-[18px]">{ratingCopy.panelHeading}</h2>{/if}
    <div class="bg-line-soft h-24 w-full animate-pulse rounded"></div>
  </section>
{:else if status === 'ready' && data !== null}
  <section class="flex flex-col gap-2" data-testid="character-rating">
    {#if heading}<h2 class="section-title text-[18px]">{ratingCopy.panelHeading}</h2>{/if}
    {#if data.sample_size === 0}
      <EmptyState
        message={ratingCopy.characterEmpty}
        action={{ label: ratingCopy.characterEmptyAction, href: '/logs' }}
        testid="character-rating-empty"
      />
    {:else}
      {#if data.latest !== null}
        <div class="flex items-center gap-4">
          <span class="text-gold tabular text-[32px] font-bold" data-testid="character-rating-overall">
            {Math.round(data.latest.overall)}
          </span>
          <span class="flex w-32 shrink-0 gap-[2px]" aria-hidden="true">
            {#each RATING_COMPONENT_ORDER as name (name)}
              {@const part = data.latest.components.find((c) => c.name === name)}
              <span class="bg-line-soft h-2 flex-1" title={componentLabel(name)}
                ><span
                  class="bg-gold block h-full"
                  style={`width: ${part && !part.excluded && part.score !== null ? part.score : 0}%`}
                ></span></span
              >
            {/each}
          </span>
        </div>
      {/if}
      {#if data.sample_size < MIN_TREND_SAMPLES}
        <p class="text-muted text-[13px]" data-testid="character-rating-too-few">
          {ratingCopy.trendTooFew(data.sample_size)}
        </p>
      {:else}
        <RatingTrend points={data.trend} />
      {/if}
      {#if data.best_component !== '' || data.worst_component !== ''}
        <p class="text-muted text-[13px]">
          {#if data.best_component !== ''}Best: {componentLabel(data.best_component)}.{/if}
          {#if data.worst_component !== ''}Needs work: {componentLabel(data.worst_component)}.{/if}
        </p>
      {/if}
    {/if}
    <a
      class="text-gold inline-flex min-h-11 items-center text-[12px] underline-offset-2 hover:underline md:min-h-0"
      href="/ratings">{ratingCopy.explainLink}</a
    >
  </section>
{/if}
