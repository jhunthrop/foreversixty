<!-- web/src/components/report/RatingTab.svelte -->
<!-- The Rating tab (logs spec 2026-11-07 section 6.2): Section `Performance rating` holding the
     tab's one Plate. A trash pull reads the shipped `noRatingForTrash` line. A boss pull draws,
     per rated player in the source scope, a first row (crest, `{name}'s rating`, the basis
     sentence, the overall as a big `70 of 100` in its parse colour), then a StatTable row per part: its name and
     one-line definition, a score Meter (green from 75), `Of 100`, `Weight`, `Basis`; an excluded
     part reads `ratingCopy.excludedReason` with dashes. A part's moments link into the tab that
     shows them. The rail's rated players and The cap live in RatingRail.svelte; `onRatings`
     hands this tab's fetch to it so the page fetches once. -->
<script lang="ts">
  import CharacterName from '../ui/CharacterName.svelte';
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar } from '../../lib/report/format';
  import { classSlugOf } from '../../lib/report/planner-link';
  import { SOURCE_ENEMIES, SOURCE_FRIENDLIES } from '../../lib/report/url';
  import { createReportRatingsFetch, type ReportRatingsStatus } from '../../lib/rating/report-ratings.svelte';
  import { ratingCopy, componentLabel } from '../../lib/rating/copy';
  import { momentHref } from '../../lib/rating/moments';
  import { classForPlayer, guidForPlayer } from '../../lib/rating/roster-join';
  import {
    RATING_LOADING,
    basisColumn,
    basisSentence,
    momentText,
    orderedComponents,
    partDefinition,
    playerRating,
    ratingColour,
    scoreTone,
  } from '../../lib/report/rating-card';
  import type { RatingCardPlayer, RatingComponent, ReportRatings } from '../../lib/rating/types';
  import type { StatColumn, StatFigure } from '../../lib/ui/stat-table';
  import Crest from '../ui/Crest.svelte';
  import Divided from '../ui/Divided.svelte';
  import Figure from '../ui/Figure.svelte';
  import Meter from '../ui/Meter.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';

  let {
    reportId,
    fightIndex,
    roster,
    source,
    onSelectPlayer,
    apiBase = undefined,
    onRatings = undefined,
  }: {
    reportId: string;
    fightIndex: number;
    roster: { guid: string; name: string; class?: string }[];
    source: string;
    onSelectPlayer: (guid: string) => void;
    apiBase?: string;
    /** The fetch's state on every change, for the rail (RatingRail's `ratings` and `status`). */
    onRatings?: (state: { data: ReportRatings | null; status: ReportRatingsStatus }) => void;
  } = $props();

  const fetcher = createReportRatingsFetch(apiBase);
  $effect(() => {
    fetcher.load(reportId, fightIndex);
  });
  $effect(() => {
    onRatings?.({ data: fetcher.data, status: fetcher.status });
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

  const COLUMNS: readonly StatColumn[] = [
    { id: 'score', label: 'Of 100' },
    { id: 'weight', label: 'Weight' },
    { id: 'basis', label: 'Basis' },
  ];
  const DASH: StatFigure = { value: '–', muted: true };

  function figuresOf(part: RatingComponent): StatFigure[] {
    if (part.excluded) return [DASH, DASH, DASH];
    return [
      { value: String(Math.round(part.score ?? 0)) },
      { value: `${Math.round(part.weight)}%` },
      { value: basisColumn(part), muted: true },
    ];
  }

  const classOf = (player: RatingCardPlayer, guid: string): string | undefined =>
    classForPlayer(guid, player.class, roster);
</script>

{#snippet line(text: string, testid: string, alert = false)}
  <p class="text-muted px-4 py-4 text-[14px]" role={alert ? 'alert' : undefined} data-testid={testid}>
    {text}
  </p>
{/snippet}

{#snippet partRow(guid: string, part: RatingComponent)}
  <StatRow
    figures={figuresOf(part)}
    phone={part.excluded ? undefined : 'of 100'}
    testid={`rating-component-${part.name}`}
  >
    <!-- What a part measures is its name's hover (logs lane review 2026-11-09): the row carries the
         score, and a line only when the part was not scored or has moments to open. -->
    {#snippet name()}<span
        title={part.name === 'utility' && !part.excluded
          ? `${partDefinition(part.name)} ${ratingCopy.threatNotModeled}`
          : partDefinition(part.name)}
        data-testid={part.name === 'utility' && !part.excluded ? 'rating-threat-note' : undefined}
        >{componentLabel(part.name)}</span
      >{/snippet}
    {#snippet sub()}
      {#if part.excluded}
        <div>{ratingCopy.excludedReason(componentLabel(part.name), part.reason)}</div>
      {/if}
      {#if part.moments.length > 0}
        <ul class="flex flex-wrap gap-x-3" data-testid={`rating-moments-${part.name}`}>
          {#each part.moments as moment, index (index)}
            {@const href = momentHref(fightIndex, guid, moment)}
            <li>
              {#if href !== null}
                <a class="text-gold underline-offset-2 hover:underline" {href}>{momentText(moment)}</a>
              {:else}
                {momentText(moment)}
              {/if}
            </li>
          {/each}
        </ul>
      {/if}
    {/snippet}
    {#snippet meter()}
      {#if !part.excluded}<Meter value={(part.score ?? 0) / 100} tone={scoreTone(part.score ?? 0)} />{/if}
    {/snippet}
  </StatRow>
{/snippet}

{#snippet card(player: RatingCardPlayer, guid: string)}
  {@const name = splitUnitName(player.player_name).name}
  <div class="w-full" data-testid={`rating-card-${guid || player.player_key}`}>
    <Row title={playerRating(name)} sub={basisSentence(player)}>
      {#snippet lead()}<Crest cls={classSlugOf(classOf(player, guid)) ?? ''} size={32} />{/snippet}
      {#snippet titleContent()}
        {#if guid !== ''}
          <button
            type="button"
            class="focus-visible:outline-gold text-left underline-offset-2 hover:underline focus-visible:outline-2"
            style:color={classColorVar(classOf(player, guid))}
            title="Show only this player"
            onclick={() => onSelectPlayer(guid)}>{playerRating(name)}</button
          >
        {:else}
          <CharacterName name={playerRating(name)} cls={classOf(player, guid)} weight="inherit" />
        {/if}
      {/snippet}
      {#snippet trail()}
        {#if player.insufficient}
          <span class="text-muted text-[15px] font-semibold" data-testid="rating-insufficient"
            >{ratingCopy.insufficient}</span
          >
        {:else}
          <Figure
            value={String(Math.round(player.overall))}
            unit="of 100"
            color={ratingColour(player.overall)}
            size="big"
            testid="rating-overall"
          />
        {/if}
      {/snippet}
    </Row>
    {#if player.insufficient}
      <p class="text-muted px-4 pb-3 text-[13px]" data-testid="rating-insufficient-note">
        {ratingCopy.insufficientNote(player.insufficient_reason)}
      </p>
    {:else if player.overall_capped}
      <p class="text-muted px-4 pb-3 text-[13px]" data-testid="rating-capped-note">
        <span class="font-semibold tabular-nums">{Math.round(player.overall)}</span>, {ratingCopy.cappedNote}
      </p>
    {/if}
    <StatTable
      columns={COLUMNS}
      nameHead="Part"
      barHead="Score"
      label={playerRating(name)}
      testid="rating-components"
    >
      {#each orderedComponents(player) as part (part.name)}
        {@render partRow(guid, part)}
      {/each}
    </StatTable>
  </div>
{/snippet}

<Section
  title={ratingCopy.panelHeading}
  action={{ label: ratingCopy.explainLink, href: '/ratings' }}
  testid="rating-tab"
  panel
  flush
>
  <div>
    {#if fetcher.status === 'loading' || fetcher.status === 'idle'}
      {@render line(RATING_LOADING, 'rating-tab-loading')}
    {:else if fetcher.status === 'flagged'}
      {@render line(ratingCopy.flagged, 'rating-tab-flagged')}
    {:else if fetcher.status === 'none'}
      {@render line(ratingCopy.noRatingForTrash, 'rating-tab-none')}
    {:else if fetcher.status === 'failed'}
      {@render line(ratingCopy.fetchFailed, 'rating-tab-error', true)}
    {:else if source === SOURCE_ENEMIES}
      {@render line(ratingCopy.enemiesHaveNone, 'rating-tab-enemies')}
    {:else if visible.length === 0 && withGuids.length === 0}
      {@render line(ratingCopy.tabEmpty, 'rating-tab-empty')}
    {:else if visible.length === 0}
      {@render line(ratingCopy.noMatchingRow, 'rating-tab-no-match')}
    {:else}
      {#each visible as { player, guid }, index (player.player_key)}
        {#if index === 0}{@render card(player, guid)}
        {:else}<Divided gap={0}>{@render card(player, guid)}</Divided>{/if}
      {/each}
    {/if}
  </div>
</Section>
