<!-- web/src/components/report/RatingTab.svelte -->
<!-- The full report card: one number first, then the six components, then each
     component's moments -- spec section 6.1's exact three-level structure. Native
     <details>/<summary> disclosures throughout (the report/ directory's own ActorRow.svelte
     and DeathsTab.svelte roll a button + aria-expanded + local state instead, but this
     card's nesting -- a details per player, a details per component inside it -- reads
     more plainly with the browser's own disclosure element and needs no click-driven state
     to open the single-player case correctly on first render). -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar } from '../../lib/report/format';
  import { SOURCE_ENEMIES, SOURCE_FRIENDLIES } from '../../lib/report/url';
  import { createReportRatingsFetch } from '../../lib/rating/report-ratings.svelte';
  import { ratingCopy, componentLabel } from '../../lib/rating/copy';
  import { momentHref } from '../../lib/rating/moments';
  import { classForPlayer, guidForPlayer } from '../../lib/rating/roster-join';
  import { RATING_COMPONENT_ORDER } from '../../lib/rating/types';
  import type { RatingCardPlayer, RatingComponent, RatingMoment } from '../../lib/rating/types';

  let {
    reportId,
    fightIndex,
    roster,
    source,
    onSelectPlayer,
    apiBase = undefined,
  }: {
    reportId: string;
    fightIndex: number;
    roster: { guid: string; name: string; class?: string }[];
    source: string;
    onSelectPlayer: (guid: string) => void;
    apiBase?: string;
  } = $props();

  const fetcher = createReportRatingsFetch(apiBase);
  $effect(() => {
    fetcher.load(reportId, fightIndex);
  });

  const withGuids = $derived(
    (fetcher.data?.players ?? []).map((player) => ({ player, guid: guidForPlayer(player, roster) })),
  );
  const visible = $derived(
    source === SOURCE_ENEMIES
      ? []
      : source === SOURCE_FRIENDLIES
        ? withGuids
        : withGuids.filter((row) => row.guid === source),
  );

  function orderedComponents(player: RatingCardPlayer): RatingComponent[] {
    return RATING_COMPONENT_ORDER.map(
      (name) =>
        player.components.find((c) => c.name === name) ?? {
          name,
          score: null,
          weight: 0,
          basis: '',
          percentile: null,
          bracket_n: 0,
          excluded: true,
          reason: '',
          moments: [],
        },
    );
  }

  function basisLine(player: RatingCardPlayer, part: RatingComponent): string {
    if (part.excluded) return ratingCopy.excludedReason(componentLabel(part.name), part.reason);
    if (part.basis === 'percentile' && part.percentile !== null)
      return ratingCopy.percentileBasis(
        player.spec,
        player.class,
        player.role,
        part.percentile,
        part.bracket_n,
      );
    if (part.basis === 'absolute') return ratingCopy.absoluteBasis;
    return '';
  }

  function momentText(moment: RatingMoment): string {
    const time = moment.at_ms !== undefined ? `${Math.round(moment.at_ms / 1000)}s — ` : '';
    return `${time}${moment.spell_name ?? 'Unnamed'}`;
  }
</script>

<div class="flex flex-col gap-4" data-testid="rating-tab">
  {#if fetcher.status === 'loading' || fetcher.status === 'idle'}
    <p class="text-muted text-[14px]" data-testid="rating-tab-loading">Loading ratings.</p>
  {:else if fetcher.status === 'failed'}
    <p class="text-[14px]" role="alert" data-testid="rating-tab-error">{ratingCopy.fetchFailed}</p>
  {:else if source === SOURCE_ENEMIES}
    <p class="text-muted text-[14px]" data-testid="rating-tab-enemies">{ratingCopy.enemiesHaveNone}</p>
  {:else if visible.length === 0 && withGuids.length === 0}
    <p class="text-muted text-[14px]" data-testid="rating-tab-empty">{ratingCopy.tabEmpty}</p>
  {:else if visible.length === 0}
    <p class="text-muted text-[14px]" data-testid="rating-tab-no-match">{ratingCopy.noMatchingRow}</p>
  {:else}
    {#each visible as { player, guid } (player.player_key)}
      <details
        class="border-line rounded-panel bg-raised border p-3"
        data-testid={`rating-card-${guid || player.player_key}`}
        open={visible.length === 1}
      >
        <summary class="flex cursor-pointer items-center gap-3">
          {#if guid !== ''}
            <button
              type="button"
              class="flex-1 truncate text-left font-semibold underline-offset-2 hover:underline"
              style={`color: ${classColorVar(classForPlayer(guid, player.class, roster))}`}
              title="Show only this player"
              onclick={(event) => {
                // A button nested in <summary> still bubbles its click up to toggle the
                // details unless the click's default action is stopped here -- selecting a
                // player should narrow the source, not just fold this one card.
                event.preventDefault();
                onSelectPlayer(guid);
              }}>{splitUnitName(player.player_name).name}</button
            >
          {:else}
            <span
              class="flex-1 truncate font-semibold"
              style={`color: ${classColorVar(classForPlayer(guid, player.class, roster))}`}
            >
              {splitUnitName(player.player_name).name}
            </span>
          {/if}
          <span class="text-gold tabular text-[28px] font-bold" data-testid="rating-overall"
            >{Math.round(player.overall)}</span
          >
        </summary>
        {#if player.overall_capped}
          <p class="text-muted mt-2 text-[13px]" data-testid="rating-capped-note">
            <span class="tabular font-semibold">{Math.round(player.overall)}</span>, {ratingCopy.cappedNote}
          </p>
        {/if}
        <ul class="mt-3 flex flex-col gap-2" data-testid="rating-components">
          {#each orderedComponents(player) as part (part.name)}
            <li>
              <details
                class="border-line-soft rounded-control border p-2"
                data-testid={`rating-component-${part.name}`}
              >
                <summary class="flex cursor-pointer items-center gap-3 text-[13px]">
                  <span class="flex-1 font-semibold">{componentLabel(part.name)}</span>
                  {#if !part.excluded}
                    <span class="tabular text-muted text-[12px]">weight {Math.round(part.weight)}%</span>
                    <span class="text-gold tabular w-10 text-right font-mono font-bold"
                      >{Math.round(part.score ?? 0)}</span
                    >
                  {:else}
                    <span class="text-muted text-[12px]">not scored</span>
                  {/if}
                </summary>
                <p class="text-muted mt-2 text-[12px]">{basisLine(player, part)}</p>
                {#if part.name === 'utility' && !part.excluded}
                  <p class="text-muted mt-1 text-[12px]" data-testid="rating-threat-note">
                    {ratingCopy.threatNotModeled}
                  </p>
                {/if}
                {#if part.moments.length > 0}
                  <ul class="mt-2 flex flex-col gap-1" data-testid={`rating-moments-${part.name}`}>
                    {#each part.moments as moment, index (index)}
                      {@const href = momentHref(fightIndex, guid, moment)}
                      <li class="text-[12px]">
                        {#if href !== null}
                          <a class="text-gold underline-offset-2 hover:underline" {href}
                            >{momentText(moment)}</a
                          >
                        {:else}
                          <span class="text-muted">{momentText(moment)}</span>
                        {/if}
                      </li>
                    {/each}
                  </ul>
                {/if}
              </details>
            </li>
          {/each}
        </ul>
      </details>
    {/each}
  {/if}
  <a
    class="text-gold inline-flex min-h-11 items-center text-[12px] underline-offset-2 hover:underline md:min-h-0"
    href="/ratings">{ratingCopy.explainLink}</a
  >
</div>
