<!-- web/src/components/report/RatingRail.svelte -->
<!-- The Rating tab's rail under This pull (logs spec 2026-11-07 section 6.2): Section
     `Performance rating`, a Row per rated player (crest, name in the class colour, `{n} of 100`;
     `Not rated` when under half the weight was measurable), then Section `The cap` with
     `ratingCopy.cappedRule`. It draws the RatingTab's own fetch, handed over through RatingTab's
     `onRatings`, so the page fetches the ratings once. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { classSlugOf } from '../../lib/report/planner-link';
  import { ratingCopy } from '../../lib/rating/copy';
  import type { ReportRatingsStatus } from '../../lib/rating/report-ratings.svelte';
  import { classForPlayer, guidForPlayer } from '../../lib/rating/roster-join';
  import type { ReportRatings } from '../../lib/rating/types';
  import { classTone } from '../../lib/ui/primitives';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import Figure from '../ui/Figure.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let {
    ratings,
    status,
    roster,
  }: {
    ratings: ReportRatings | null;
    status: ReportRatingsStatus;
    roster: { guid: string; name: string; class?: string }[];
  } = $props();

  const rows = $derived(
    (ratings?.players ?? []).map((player) => {
      const guid = guidForPlayer(player, roster);
      return { player, className: classForPlayer(guid, player.class, roster) ?? '' };
    }),
  );

  const emptyLine = $derived(
    status === 'none'
      ? ratingCopy.noRatingForTrash
      : status === 'flagged'
        ? ratingCopy.flagged
        : status === 'failed'
          ? ratingCopy.fetchFailed
          : status === 'ready'
            ? ratingCopy.panelEmpty
            : '',
  );
</script>

<div class="flex flex-col gap-[var(--section-gap)]" data-testid="rating-rail">
  <Section title={ratingCopy.panelHeading}>
    {#if rows.length > 0}
      <div data-testid="rating-rail-rows">
        {#each rows as { player, className } (player.player_key)}
          <Row
            title={splitUnitName(player.player_name).name}
            titleTone={classTone(classSlugOf(className) ?? '')}
            open
          >
            {#snippet lead()}<ClassCrestRing
                characterClass={classSlugOf(className) ?? ''}
                size={28}
              />{/snippet}
            {#snippet trail()}
              {#if player.insufficient}
                <span class="text-muted text-[13px]">{ratingCopy.insufficient}</span>
              {:else}
                <Figure value={String(Math.round(player.overall))} unit="of 100" size="sm" tone="gold" />
              {/if}
            {/snippet}
          </Row>
        {/each}
      </div>
    {:else if emptyLine !== ''}
      <p class="text-muted text-[14px]" data-testid="rating-rail-empty">{emptyLine}</p>
    {/if}
  </Section>
  <Section title="The cap">
    <p class="text-muted text-[14px] leading-5">{ratingCopy.cappedRule}</p>
  </Section>
</div>
