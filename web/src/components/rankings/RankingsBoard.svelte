<!-- web/src/components/rankings/RankingsBoard.svelte -->
<!-- The character board: the column line, the viewer's best kill pinned on top only when it is
     off the visible page (design/specs/2026-10-11-rankings.md section 3, decision 5), then the
     rows from the top of the page. When the viewer's row is on the page it is marked where it
     is, never repeated. Crests appear on the rows only when the board spans more than one
     class (decision 7). -->
<script lang="ts">
  import type { RankingRow as RankingRowData, RankingsPage } from '../../lib/rankings/api';
  import { rankingsCopy } from '../../lib/rankings/copy';
  import { RANKING_ROW_GRID, TABLE_UP, WIDE_ONLY } from '../../lib/rankings/layout';
  import { ownRowPlacement, showRowCrests } from '../../lib/rankings/row-view';
  import RankingRow from './RankingRow.svelte';

  let {
    page,
    metricLabel,
    ownKey,
    now,
  }: {
    page: RankingsPage;
    metricLabel: string;
    /** The viewer's player key while the board is the viewer's own; null otherwise. */
    ownKey: string | null;
    now: Date;
  } = $props();

  const placement = $derived(ownKey === null ? 'none' : ownRowPlacement(page.rows, page.character_rank));
  const pinned = $derived<RankingRowData | null>(
    placement === 'pinned' ? (page.character_rank?.row ?? null) : null,
  );
  const crest = $derived(showRowCrests(page.rows));
  const columnClass = 'quiet-label';
</script>

<div class="bg-raised border-line rounded-panel overflow-hidden border" data-testid="ranking-board">
  <!-- The column line is drawn from the table row up, where every cell has its own column; on a
       phone the row folds its figures into a line of its own. -->
  <div
    class="{RANKING_ROW_GRID} border-line-soft hidden min-h-9 items-center gap-x-3 border-b border-l-2 border-l-transparent pr-3.5 pl-3 min-[600px]:grid"
    data-testid="ranking-columns"
  >
    <span class={columnClass}>#</span>
    <span class={columnClass}>Character</span>
    <span class="{columnClass} {WIDE_ONLY}">Guild</span>
    <span class="{columnClass} {WIDE_ONLY} text-right">Size</span>
    <span class="{columnClass} text-right">{metricLabel}</span>
    <span class="{columnClass} {TABLE_UP} text-right">Executed</span>
    <span class="{columnClass} {WIDE_ONLY} text-right">Length</span>
    <span class="{columnClass} {TABLE_UP} text-right">Date</span>
    <span class="{columnClass} {TABLE_UP} text-right">Build</span>
    <span class="{columnClass} text-right">Report</span>
  </div>

  {#if pinned !== null}
    <div data-testid="ranking-pinned">
      <p class="label text-muted px-3.5 pt-2.5">{rankingsCopy.ownBest}</p>
      <ul>
        <RankingRow row={pinned} total={page.total} own {crest} {now} testid="ranking-pinned-row" />
      </ul>
      <div class="bg-line-warm-strong h-px"></div>
    </div>
  {/if}

  <ul class="reveal flex flex-col" data-testid="ranking-rows">
    {#each page.rows as row (`${row.report_id}-${row.fight_index}-${row.player.key}`)}
      <RankingRow
        {row}
        total={page.total}
        own={ownKey !== null && row.player.key === ownKey}
        {crest}
        {now}
        testid={`ranking-${row.rank}`}
      />
    {/each}
  </ul>
</div>
