<!-- web/src/components/CharacterRatingPanel.svelte -->
<!-- The character page's rating panel: overall score, six-segment bar, a trend
     sparkline once there are enough fights, and the best/worst component -- spec §6.2.
     Fetches independently, the same self-contained pattern Character.svelte's own
     $effect already uses for fetchCharacter. -->
<script lang="ts">
  import { fetchCharacterRating, RankingsError } from '../lib/rankings/api';
  import { ratingCopy, componentLabel, MIN_TREND_SAMPLES } from '../lib/rating/copy';
  import { RATING_COMPONENT_ORDER } from '../lib/rating/types';
  import type { CharacterPath } from '../lib/characters';
  import type { CharacterRating } from '../lib/rating/types';
  import RatingTrend from './RatingTrend.svelte';
  import EmptyState from './ui/EmptyState.svelte';

  let { path, apiBase = undefined }: { path: CharacterPath; apiBase?: string } = $props();

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

{#if status === 'loading'}
  <!-- Reserves height for the common loading -> ready path so the panel's own container
       does not pop in and push "Best per encounter" down once the fetch resolves -- the
       plan's binding "no layout shift" constraint. Mirrors RatingPanel.svelte's own
       loading-skeleton idiom. The rare loading -> hidden path (an anonymized character)
       still fully collapses to nothing once resolved -- Ruling 2's own explicit
       requirement, which takes precedence over reserving permanent dead space for a
       panel that is supposed not to exist for that character. -->
  <section class="flex flex-col gap-2" data-testid="character-rating-loading" aria-hidden="true">
    <h2 class="section-title text-[18px]">{ratingCopy.panelHeading}</h2>
    <div class="bg-line-soft h-24 w-full animate-pulse rounded"></div>
  </section>
{:else if status === 'ready' && data !== null}
  <section class="flex flex-col gap-2" data-testid="character-rating">
    <h2 class="section-title text-[18px]">{ratingCopy.panelHeading}</h2>
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
