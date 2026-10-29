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
<script lang="ts" generics="Row extends CharacterFight">
  import { amountHeader, characterTableCopy, executionCellClass } from '../../lib/character-profile';
  import { fightKey } from '../../lib/character/fight-key';
  import { formatDayMonthYear } from '../../lib/dates';
  import { specName } from '../../lib/guild/spec-name';
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
    heading = true,
    columns = 'full',
    lead = 'boss',
    selected = null,
    note = null,
  }: {
    title: string;
    rows: readonly Row[];
    testid: string;
    /** Where the boss name links: its rankings board, or the fight's report. */
    bossHref: (row: Row) => string;
    rowKey: (row: Row, index: number) => string;
    executionTestid: string;
    /** Each row's id, for a link to it ("Every ranked fight", L8.5 ruling 1); none without. */
    rowId?: ((row: Row) => string) | null;
    /** The latest rated fight and its rating (L8.5 fix round 2 item 3): its row says "rated 65"
     *  under its parse. */
    rated?: { report_id: string; fight_index: number; overall: number } | null;
    /** False: no `h2` and no execution caption; `title` becomes the list's own small label (the
     *  History section's lists, experience spec 2026-10-24-character-history 3.7). */
    heading?: boolean;
    /** `no-spec` drops the Spec column (the History lists are one spec's, in a 588px column). */
    columns?: 'full' | 'no-spec';
    /** `date`: the first column is the fight's date, linked, and the Date column is blank (a
     *  list of one boss's kills). */
    lead?: 'boss' | 'date';
    /** `report_id:fight_index` keys to mark with the target treatment, with no id or scroll. */
    selected?: ReadonlySet<string> | null;
    /** A muted word after the boss name (a private or guild report), or null. */
    note?: ((row: Row) => string | null) | null;
  } = $props();

  /** The address's target (L8.5 fix round 2 items 3 and 5): a 2px gold left edge and a gold tint
   *  over the page's ground, eased over 120ms, landing 96px under the window's top. */
  const TARGET =
    ' scroll-mt-24 transition-[background-color,box-shadow] duration-[120ms] target:bg-[rgba(229,185,85,.08)] target:shadow-[inset_2px_0_0_var(--color-gold)]';
  /** "rated 65" on a line of its own, right-aligned under the parse: it spans Boss to Parse %, so
   *  it never widens a column (Spec is not drawn on a phone, so two columns there, three from
   *  md). */
  const RATED_NOTE = 'text-muted col-span-2 col-start-1 -mt-2 text-right text-[12px] md:col-span-3';
  /** The same without the Spec column: two columns from md too. */
  const RATED_NOTE_NO_SPEC = 'text-muted col-span-2 col-start-1 -mt-2 text-right text-[12px] md:col-span-2';

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

  const GRIDS = {
    full: 'grid-cols-[minmax(0,1fr)_auto_auto_auto] md:grid-cols-[minmax(0,1fr)_112px_64px_88px_112px_72px]',
    'no-spec': 'grid-cols-[minmax(0,1fr)_auto_auto_auto] md:grid-cols-[minmax(0,1fr)_64px_88px_112px_72px]',
  } as const;
  const GRID = $derived(GRIDS[columns]);
  const ratedNoteClass = $derived(columns === 'full' ? RATED_NOTE : RATED_NOTE_NO_SPEC);
  /** The marked row's gold edge and tint: the address target's look, without its id or scroll. */
  const MARKED = ' scroll-mt-24 bg-[rgba(229,185,85,.08)] shadow-[inset_2px_0_0_var(--color-gold)]';
  const isMarked = (row: Row): boolean =>
    selected !== null && selected.has(fightKey(row.report_id, row.fight_index));
</script>

{#snippet firstLink(row: Row, noted: boolean)}
  {#if lead === 'date'}
    <a class="{rowLink} text-muted numeric truncate text-[13px]" href={bossHref(row)}
      >{formatDayMonthYear(new Date(row.fought_at))}</a
    >
  {:else if noted}
    <!-- Beside a note the name takes the width left and ends in an ellipsis (text-overflow does
         not apply to the link's own flex box, so the text sits in a span). -->
    <a class="{rowLink} min-w-0" href={bossHref(row)}><span class="truncate">{row.encounter}</span></a>
  {:else}
    <a class="{rowLink} truncate" href={bossHref(row)}>{row.encounter}</a>
  {/if}
{/snippet}

<section class="flex flex-col gap-2">
  {#if heading}
    <h2 class="section-title">{title}</h2>
    <p class="text-muted text-[12px]" data-testid="character-execution-definition">
      {characterTableCopy.executionDefinition}
    </p>
  {:else}
    <p class="label text-muted pt-1.5" data-testid="character-table-label">{title}</p>
  {/if}
  <ul class="flex flex-col" data-testid={testid}>
    <li
      class={`quiet-label border-line-soft grid min-h-9 items-end gap-3 border-b px-2 pb-2 ${GRID}`}
      aria-hidden="true"
      data-testid="character-table-header"
    >
      <span>{lead === 'date' ? characterTableCopy.date : characterTableCopy.boss}</span>
      {#if columns === 'full'}<span class="hidden md:inline">{characterTableCopy.spec}</span>{/if}
      <span class="text-right">{characterTableCopy.parse}</span>
      <span class="text-right">{amountHeader(rows)}</span>
      <span class="hidden md:ml-2 md:inline">{lead === 'date' ? '' : characterTableCopy.date}</span>
      <span class="text-right max-md:ml-1">{characterTableCopy.execution}</span>
    </li>
    {#each rows as row, index (rowKey(row, index))}
      {@const word = note === null ? null : note(row)}
      <!-- L8.5 ruling 1: the rating card's weakest part links here (fix round 2: the target's
           gold edge and tint, 96px under the window's top). -->
      <li
        class={`border-line-soft grid min-h-11 items-center gap-3 border-b px-2 py-2 text-[14px] ${GRID}${rowId === null ? '' : TARGET}${isMarked(row) ? MARKED : ''}`}
        id={rowId === null ? undefined : rowId(row)}
        data-marked={isMarked(row) ? 'true' : undefined}
      >
        {#if word === null}
          {@render firstLink(row, false)}
        {:else}
          <!-- On a phone the note takes the line under the name; from md it sits beside it. -->
          <span
            class="flex min-w-0 md:items-center md:gap-2 {lead === 'boss'
              ? 'max-md:flex-col'
              : 'items-center gap-2'}"
          >
            {@render firstLink(row, true)}
            <span class="text-muted flex-none text-[12px]" data-testid="character-fight-note">{word}</span>
          </span>
        {/if}
        {#if columns === 'full'}<span class="text-muted hidden truncate text-[13px] md:inline"
            >{specName(row.spec ?? '')}</span
          >{/if}
        <span
          class="numeric text-right text-[13px]"
          style={row.percentile === undefined ? undefined : `color: ${percentileToken(row.percentile)}`}
          aria-label={percentileAriaLabel(row.percentile)}
        >
          {row.percentile === undefined ? '' : Math.round(row.percentile)}
        </span>
        <a
          class="{rowLink} numeric justify-end text-right"
          href={`/reports/${row.report_id}?fight=${row.fight_index}`}
          aria-label={`${formatAmount(Math.round(row.value))} ${metricLabel(row.metric)}`}
        >
          {formatAmount(Math.round(row.value))}
        </a>
        <span
          class="text-muted numeric hidden text-[13px] md:ml-2 md:inline"
          data-testid="character-fight-date"
          >{lead === 'date' ? '' : formatDayMonthYear(new Date(row.fought_at))}</span
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
          <span class={ratedNoteClass} data-testid="character-fight-rated"
            >{characterTableCopy.ratedNote(rated.overall)}</span
          >
        {/if}
      </li>
    {/each}
  </ul>
</section>
