<!-- web/src/components/report/RatingPanel.svelte -->
<!-- The Summary dashboard's rating row: one line per roster player, the overall score plus
     a compact six-segment bar, linking into the full RatingTab (spec §6.1). Mirrors
     SummaryPanels.svelte's own dashboard pattern (a short list, a "go deeper" link) but
     fetches its own data -- unlike SummaryPanels' siblings, a rating is not part of the
     report's already-loaded summary.json (see docs/superpowers/plans/2026-09-21-rating-web
     .md Ruling 8 on why this is its own component rather than a SummaryPanels addition). -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar } from '../../lib/report/format';
  import { createReportRatingsFetch } from '../../lib/rating/report-ratings.svelte';
  import { ratingCopy, componentLabel } from '../../lib/rating/copy';
  import { RATING_COMPONENT_ORDER } from '../../lib/rating/types';
  import { classForPlayer, guidForPlayer } from '../../lib/rating/roster-join';

  let {
    reportId,
    fightIndex,
    roster,
    onTab,
    onSelectPlayer = undefined,
    apiBase = undefined,
  }: {
    reportId: string;
    fightIndex: number;
    roster: { guid: string; name: string; class?: string }[];
    onTab: () => void;
    onSelectPlayer?: (guid: string) => void;
    apiBase?: string;
  } = $props();

  const fetcher = createReportRatingsFetch(apiBase);
  $effect(() => {
    fetcher.load(reportId, fightIndex);
  });

  const rows = $derived(
    (fetcher.data?.players ?? []).map((player) => ({ player, guid: guidForPlayer(player, roster) })),
  );

  const panel = 'border-line rounded-panel bg-raised flex flex-col gap-2 border p-3';
  const heading = 'label text-muted flex items-center justify-between';
  const more =
    'text-gold inline-flex min-h-11 items-center text-[11px] normal-case tracking-normal md:min-h-0';
</script>

<section class={panel} data-testid="rating-panel">
  <h2 class={heading}>
    {ratingCopy.panelHeading}
    <button type="button" class={more} onclick={onTab}>{ratingCopy.panelMore}</button>
  </h2>
  {#if fetcher.status === 'loading' || fetcher.status === 'idle'}
    <div class="flex h-8 flex-col gap-2" data-testid="rating-panel-loading" aria-hidden="true">
      <div class="bg-line-soft h-8 w-full animate-pulse rounded"></div>
    </div>
  {:else if fetcher.status === 'failed'}
    <p class="text-muted text-[13px]" role="alert" data-testid="rating-panel-error">
      {ratingCopy.fetchFailed}
    </p>
  {:else if rows.length === 0}
    <p class="text-muted text-[13px]" data-testid="rating-panel-empty">{ratingCopy.panelEmpty}</p>
  {:else}
    <ul class="flex flex-col" data-testid="rating-panel-rows">
      {#each rows as { player, guid } (player.player_key)}
        <li class="border-line-soft flex min-h-8 items-center gap-3 border-b py-1 text-[13px] last:border-0">
          <span
            class="flex min-w-0 flex-1 items-center truncate font-semibold"
            style={`color: ${classColorVar(classForPlayer(guid, player.class, roster))}`}
          >
            {#if onSelectPlayer && guid !== ''}
              <button
                type="button"
                class="inline-flex min-h-11 items-center truncate underline-offset-2 hover:underline md:min-h-0"
                style={`color: ${classColorVar(classForPlayer(guid, player.class, roster))}`}
                title="Show only this player"
                onclick={() => onSelectPlayer(guid)}>{splitUnitName(player.player_name).name}</button
              >
            {:else}
              {splitUnitName(player.player_name).name}
            {/if}
          </span>
          <span
            class={`tabular w-8 text-right font-mono font-bold ${player.insufficient ? 'text-muted' : 'text-gold'}`}
            data-testid="rating-panel-overall"
            title={player.insufficient ? ratingCopy.insufficientNote(player.insufficient_reason) : undefined}
            >{player.insufficient ? '—' : Math.round(player.overall)}</span
          >
          <span class="flex w-24 shrink-0 gap-[2px]" aria-hidden="true">
            {#each RATING_COMPONENT_ORDER as name (name)}
              {@const part = player.components.find((c) => c.name === name)}
              <span class="bg-line-soft h-2 flex-1" title={part ? componentLabel(name) : undefined}
                ><span
                  class="bg-gold block h-full"
                  style={`width: ${part && !part.excluded && part.score !== null ? part.score : 0}%`}
                ></span></span
              >
            {/each}
          </span>
        </li>
      {/each}
    </ul>
  {/if}
</section>
