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
  import { classForPlayer, guidForPlayer } from '../../lib/rating/roster-join';
  import type { ReportRatings } from '../../lib/rating/types';
  import { classTone } from '../../lib/ui/primitives';
  import Crest from '../ui/Crest.svelte';
  import Figure from '../ui/Figure.svelte';
  import { ratingColour } from '../../lib/report/rating-card';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let {
    ratings,
    roster,
  }: {
    ratings: ReportRatings | null;
    roster: { guid: string; name: string; class?: string }[];
  } = $props();

  const rows = $derived(
    (ratings?.players ?? []).map((player) => {
      const guid = guidForPlayer(player, roster);
      return { player, className: classForPlayer(guid, player.class, roster) ?? '' };
    }),
  );
</script>

<!-- Only with ratings to list: without them the Rating panel beside it already says why (logs lane
     review 2026-11-09), and the cap has nothing to cap. -->
{#if rows.length > 0}
  <div class="flex flex-col gap-[var(--section-gap)]" data-testid="rating-rail">
    <Section title={ratingCopy.panelHeading} panel>
      <div data-testid="rating-rail-rows">
        {#each rows as { player, className } (player.player_key)}
          <Row
            title={splitUnitName(player.player_name).name}
            titleTone={classTone(classSlugOf(className) ?? '')}
            open
          >
            {#snippet lead()}<Crest
                cls={classSlugOf(className) ?? ''}
                size={28}
              />{/snippet}
            {#snippet trail()}
              {#if player.insufficient}
                <span class="text-muted text-[13px]">{ratingCopy.insufficient}</span>
              {:else}
                <Figure
                  value={String(Math.round(player.overall))}
                  unit="of 100"
                  size="sm"
                  color={ratingColour(player.overall)}
                />
              {/if}
            {/snippet}
          </Row>
        {/each}
      </div>
    </Section>
    <Section title="The cap" panel>
      <p class="text-muted text-[14px] leading-5">{ratingCopy.cappedRule}</p>
    </Section>
  </div>
{/if}
