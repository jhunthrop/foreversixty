<!-- web/src/components/CharacterRatingPanel.svelte -->
<!-- The character page's rating panel. At the top, the Forever Score card
     (ForeverScoreCard.svelte, spec 2026-10-11 §4.6): plain data from the caller
     (Character.svelte already holds the GET /v1/characters/... answer that carries it).
     Beneath it the execution rating in the same card shell (RatingCard.svelte, fix round 2
     item 6): "65 of 100", the six-segment bar when any segment has a score, the line with the
     latest fight's rating and the count, the trend sparkline once there are enough fights, and
     the weakest part of a fight, linked to its row in "Every ranked fight" (L8.5 ruling 1) --
     spec §6.2. That part fetches independently, the same self-contained pattern
     Character.svelte's own $effect already uses for fetchCharacter. -->
<script lang="ts">
  import { latestRatedFight, weakestFightHref } from '../lib/character-profile';
  import { fetchCharacterRating, RankingsError, type CharacterFight } from '../lib/rankings/api';
  import { ratingCopy } from '../lib/rating/copy';
  import type { ForeverScore, ScorePoint } from '../lib/score/score';
  import type { CharacterPath } from '../lib/characters';
  import type { CharacterRating } from '../lib/rating/types';
  import ForeverScoreCard from './ForeverScoreCard.svelte';
  import RatingCard from './RatingCard.svelte';
  import EmptyState from './ui/EmptyState.svelte';

  let {
    path,
    apiBase = undefined,
    heading = true,
    forever = undefined,
    scoreHistory = null,
    bisHref = null,
    gearHref = null,
    owned = false,
    compactEmpty = false,
    name = '',
    fights = [],
    onrated = undefined,
    onshown = undefined,
  }: {
    path: CharacterPath;
    apiBase?: string;
    /** False when the mounting panel already titles it (the account page's "Your ratings"). */
    heading?: boolean;
    /** The character's Forever Score; null (or a null score) for a character the site has no
     *  logs of. Left out, the panel draws no score card at all (the account page). */
    forever?: ForeverScore | null;
    /** The character's score_history, for the score card's sparkline. */
    scoreHistory?: readonly ScorePoint[] | null;
    /** The spec's best-in-slot page, when known (the score card's "Most to gain: Gear"). */
    bisHref?: string | null;
    /** The profile's Gear section's anchor, when the page has one (L8.6 ruling 2): the score
     *  card's "Most to gain: Gear" lands there instead of the BiS page. */
    gearHref?: string | null;
    /** The viewer's own character: the captions say "Your" only then (fix round 2 item 1). */
    owned?: boolean;
    /** The page leads with its gear: a character with no score shows one "No logs yet" line, not the
     *  card of empty bars (W.10). */
    compactEmpty?: boolean;
    /** The character's name, for the no-logs card's "Look up another character". */
    name?: string;
    /** "Every ranked fight"'s rows, which the rating card's weakest part links into (L8.5
     *  ruling 1); none, no link. */
    fights?: readonly CharacterFight[];
    /** Told the latest rated fight and its rating once the rating loads (null when there is
     *  none), so "Every ranked fight" can mark that row "rated 65" (L8.5 fix round 2 item 3). */
    onrated?: (rated: { report_id: string; fight_index: number; overall: number } | null) => void;
    /** Told whether the panel has any rating to show (a fight rated), so the account page can
     *  leave its "Your ratings" section out when there is nothing in it (W.3). */
    onshown?: (shown: boolean) => void;
  } = $props();

  let data = $state<CharacterRating | null>(null);
  let status = $state<'loading' | 'ready' | 'hidden'>('loading');
  /** The latest rated fight, named in the rating card's line (L8.5 fix round 2 item 1). */
  const latest = $derived(data === null ? null : latestRatedFight(data.trend, fights));

  $effect(() => {
    const requested = path;
    status = 'loading';
    void fetchCharacterRating(requested, apiBase)
      .then((result) => {
        if (path !== requested) return;
        data = result;
        status = 'ready';
        onshown?.(result.sample_size > 0);
        const newest = latestRatedFight(result.trend, []);
        onrated?.(
          newest === null || result.latest === null
            ? null
            : {
                report_id: newest.report_id,
                fight_index: newest.fight_index,
                overall: result.latest.overall,
              },
        );
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
        onshown?.(false);
        onrated?.(null);
      });
  });
</script>

{#if forever !== undefined}
  <ForeverScoreCard
    {forever}
    history={scoreHistory}
    {bisHref}
    {gearHref}
    {owned}
    compact={compactEmpty}
    {name}
    region={path.region}
    ruleset={path.ruleset}
  />
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
    {#if heading}<h2 class="section-title">{ratingCopy.panelHeading}</h2>{/if}
    <div class="bg-line-soft h-24 w-full animate-pulse rounded"></div>
  </section>
{:else if status === 'ready' && data !== null}
  <section class="flex flex-col gap-2" data-testid="character-rating">
    {#if heading}
      <h2 class="section-title">{ratingCopy.panelHeading}</h2>
      <p class="text-muted text-[13px]" data-testid="character-rating-caption">
        {ratingCopy.panelCaption(owned)}
      </p>
    {/if}
    {#if data.sample_size === 0}
      <EmptyState
        message={ratingCopy.characterEmpty}
        action={{ label: ratingCopy.characterEmptyAction, href: '/logs' }}
        testid="character-rating-empty"
      />
    {:else}
      <RatingCard
        overall={data.latest === null ? null : data.latest.overall}
        components={data.latest?.components ?? []}
        trend={data.trend}
        sampleSize={data.sample_size}
        worst={data.worst_component}
        weakestHref={weakestFightHref(data.trend, fights)}
        latestFight={latest === null ? null : { encounter: latest.encounter, date: latest.date }}
      />
    {/if}
  </section>
{/if}
