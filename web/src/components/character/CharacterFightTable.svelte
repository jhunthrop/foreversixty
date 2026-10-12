<!-- web/src/components/character/CharacterFightTable.svelte -->
<!-- One of the character page's two fight tables, "Best per encounter" and "Every ranked fight"
     (fix round 2 item 6): both share one column order under a header row -- Boss, Spec, Parse %,
     the amount (HPS or DPS), Date, Execution -- and dates read "9 Dec 2026". On a phone the Spec
     and Date columns go; Execution stays in both tables (L8.4 fix round 4 items 3 and 7), at
     least 44px wide and 16px after the amount. From md Date sits 8px further from the amount
     (fix round 4 item 8).
     L8.4 fix round 3: both Execution states (a scored link, an unscored dash) take one
     row-aligned cell style (lib/character-profile.ts's executionCellClass, items 4 and 6); the
     column is explained by a muted 12px caption under the title, at every width (item 7, fix
     round 4 items 6 and 7). The dash is named "Not scored yet" on hover and aloud (fix round 5
     item 7). -->
<script lang="ts">
  import { amountHeader, characterTableCopy, executionCellClass } from '../../lib/character-profile';
  import { formatDayMonthYear } from '../../lib/dates';
  import type { CharacterFight } from '../../lib/rankings/api';
  import { RANKING_METRICS } from '../../lib/rankings/url';
  import { formatAmount, percentileToken, rowLink } from '../../lib/report/format';
  import { executionHref, executionLabel, executionTitle } from '../../lib/sim/execution';

  let {
    title,
    rows,
    testid,
    bossHref,
    rowKey,
    executionTestid,
    rowId = null,
    rated = null,
  }: {
    title: string;
    rows: readonly CharacterFight[];
    testid: string;
    /** Where the boss name links: its rankings board, or the fight's report. */
    bossHref: (row: CharacterFight) => string;
    rowKey: (row: CharacterFight, index: number) => string;
    executionTestid: string;
    /** Each row's id, for a link to it ("Every ranked fight", L8.5 ruling 1); none without. */
    rowId?: ((row: CharacterFight) => string) | null;
    /** The latest rated fight and its rating (L8.5 fix round 2 item 3): its row says "rated 65"
     *  under its parse. */
    rated?: { report_id: string; fight_index: number; overall: number } | null;
  } = $props();

  /** The address's target (L8.5 fix round 2 items 3 and 5): a 2px gold left edge and a gold tint
   *  over the page's ground, eased over 120ms, landing 96px under the window's top. */
  const TARGET =
    ' scroll-mt-24 transition-[background-color,box-shadow] duration-[120ms] target:bg-[rgba(229,185,85,.08)] target:shadow-[inset_2px_0_0_var(--color-gold)]';
  /** "rated 65" on a line of its own, right-aligned under the parse: it spans Boss to Parse %, so
   *  it never widens a column (Spec is not drawn on a phone, so two columns there, three from
   *  md). */
  const RATED_NOTE = 'text-muted col-span-2 col-start-1 -mt-2 text-right text-[12px] md:col-span-3';

  /**
   * A percentile beside a formatted amount has no column heading at any breakpoint --
   * unlike the report island's tables, this list never had one to lose, so the two bare
   * numbers need their own labels rather than relying on position. `RANKING_METRICS`
   * already holds the one label table for `dps`/`hps`/`damage_taken`; echoed as-is if the
   * API ever returns an id this list has not heard of.
   */
  function metricLabel(id: string): string {
    return RANKING_METRICS.find((metric) => metric.id === id)?.label ?? id;
  }

  function percentileAriaLabel(percentile: number | undefined): string | undefined {
    return percentile === undefined ? undefined : `${Math.round(percentile)} percentile`;
  }

  const GRID =
    'grid-cols-[minmax(0,1fr)_auto_auto_auto] md:grid-cols-[minmax(0,1fr)_112px_64px_88px_112px_72px]';
</script>

<section class="flex flex-col gap-2">
  <h2 class="section-title text-[18px]">{title}</h2>
  <p class="text-muted text-[12px]" data-testid="character-execution-definition">
    {characterTableCopy.executionDefinition}
  </p>
  <ul class="flex flex-col" data-testid={testid}>
    <li
      class={`label text-muted border-line-soft grid min-h-9 items-end gap-3 border-b px-2 pb-2 ${GRID}`}
      aria-hidden="true"
      data-testid="character-table-header"
    >
      <span>{characterTableCopy.boss}</span>
      <span class="hidden md:inline">{characterTableCopy.spec}</span>
      <span class="text-right">{characterTableCopy.parse}</span>
      <span class="text-right">{amountHeader(rows)}</span>
      <span class="hidden md:ml-2 md:inline">{characterTableCopy.date}</span>
      <span class="text-right max-md:ml-1">{characterTableCopy.execution}</span>
    </li>
    {#each rows as row, index (rowKey(row, index))}
      <!-- L8.5 ruling 1: the rating card's weakest part links here (fix round 2: the target's
           gold edge and tint, 96px under the window's top). -->
      <li
        class={`border-line-soft grid min-h-11 items-center gap-3 border-b px-2 py-2 text-[14px] ${GRID}${rowId === null ? '' : TARGET}`}
        id={rowId === null ? undefined : rowId(row)}
      >
        <a class="{rowLink} truncate" href={bossHref(row)}>{row.encounter}</a>
        <span class="text-muted hidden truncate text-[13px] md:inline">{row.spec ?? ''}</span>
        <span
          class="tabular text-right font-mono text-[13px]"
          style={row.percentile === undefined ? undefined : `color: ${percentileToken(row.percentile)}`}
          aria-label={percentileAriaLabel(row.percentile)}
        >
          {row.percentile === undefined ? '' : Math.round(row.percentile)}
        </span>
        <a
          class="{rowLink} tabular justify-end text-right font-mono"
          href={`/reports/${row.report_id}?fight=${row.fight_index}`}
          aria-label={`${formatAmount(Math.round(row.value))} ${metricLabel(row.metric)}`}
        >
          {formatAmount(Math.round(row.value))}
        </a>
        <span
          class="text-muted tabular hidden font-mono text-[13px] md:ml-2 md:inline"
          data-testid="character-fight-date">{formatDayMonthYear(new Date(row.fought_at))}</span
        >
        {#if row.execution_score === null}
          <span
            class={executionCellClass(false)}
            title={characterTableCopy.executionUnscored}
            aria-label={characterTableCopy.executionUnscored}
            data-testid={executionTestid}>{executionLabel(null)}</span
          >
        {:else}
          <a
            class={executionCellClass(true)}
            href={executionHref(row.report_id, row.fight_index)}
            title={executionTitle(row.execution_score)}
            data-testid={executionTestid}>{executionLabel(row.execution_score)}</a
          >
        {/if}
        {#if rated !== null && rated.report_id === row.report_id && rated.fight_index === row.fight_index}
          <span class={RATED_NOTE} data-testid="character-fight-rated"
            >{characterTableCopy.ratedNote(rated.overall)}</span
          >
        {/if}
      </li>
    {/each}
  </ul>
</section>
