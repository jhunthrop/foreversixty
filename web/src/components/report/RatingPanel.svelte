<!-- web/src/components/report/RatingPanel.svelte -->
<!-- The Summary rail's Performance rating (logs spec 2026-11-07 6.2): a Row per rated player,
     the crest, the name in its class colour and the score as "70 of 100", with the Rating tab ›
     link. A boss pull's only: the report shows no rating panel on trash (directions b1). A rating is not part of summary.json, so the panel fetches its
     own (docs/superpowers/plans/2026-09-21-rating-web.md Ruling 8). -->
<script lang="ts">
  import EmptyState from '../ui/EmptyState.svelte';
  import { unitLabel } from '../../lib/report/names.svelte';
  import { classColorVar } from '../../lib/report/format';
  import { ratingColour } from '../../lib/report/rating-card';
  import { createReportRatingsFetch } from '../../lib/rating/report-ratings.svelte';
  import { ratingCopy } from '../../lib/rating/copy';
  import { classForPlayer, guidForPlayer } from '../../lib/rating/roster-join';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import Crest from '../ui/Crest.svelte';

  let {
    reportId,
    fightIndex,
    roster,
    tabHref,
    apiBase = undefined,
  }: {
    reportId: string;
    fightIndex: number;
    roster: { guid: string; name: string; class?: string }[];
    /** The Rating tab's url, kept with the fight. */
    tabHref: string;
    apiBase?: string;
  } = $props();

  const fetcher = createReportRatingsFetch(apiBase);
  $effect(() => {
    fetcher.load(reportId, fightIndex);
  });

  const rows = $derived(
    (fetcher.data?.players ?? []).map((player) => ({ player, guid: guidForPlayer(player, roster) })),
  );
</script>

<Section
  title={ratingCopy.panelHeading}
  action={{ label: ratingCopy.panelMore, href: tabHref }}
  testid="rating-panel"
  panel
>
  {#if fetcher.status === 'loading' || fetcher.status === 'idle'}
    <div data-testid="rating-panel-loading" aria-hidden="true"><Skeleton minHeight="min-h-[58px]" /></div>
  {:else if fetcher.status === 'flagged'}
    <EmptyState inset message={ratingCopy.flagged} testid="rating-panel-flagged" />
  {:else if fetcher.status === 'none'}
    <EmptyState inset message={ratingCopy.noRatingForTrash} testid="rating-panel-none" />
  {:else if fetcher.status === 'failed'}
    <EmptyState inset message={ratingCopy.fetchFailed} role="alert" testid="rating-panel-error" />
  {:else if rows.length === 0}
    <EmptyState inset message={ratingCopy.panelEmpty} testid="rating-panel-empty" />
  {:else}
    <div data-testid="rating-panel-rows">
      {#each rows as { player, guid } (player.player_key)}
        {@const colour = classColorVar(classForPlayer(guid, player.class, roster))}
        <Row title={unitLabel(player.player_name).text} open>
          {#snippet lead()}<Crest cls={classForPlayer(guid, player.class, roster)} size={28} />{/snippet}
          {#snippet titleContent()}<span style:color={colour}>{unitLabel(player.player_name).text}</span
            >{/snippet}
          {#snippet trail()}<span
              class={['text-[15px] font-bold tabular-nums', player.insufficient && 'text-muted']}
              style:color={player.insufficient ? undefined : ratingColour(player.overall)}
              data-testid="rating-panel-overall"
              title={player.insufficient
                ? ratingCopy.insufficientNote(player.insufficient_reason)
                : undefined}>{player.insufficient ? '–' : `${Math.round(player.overall)} of 100`}</span
            >{/snippet}
        </Row>
      {/each}
    </div>
  {/if}
</Section>
