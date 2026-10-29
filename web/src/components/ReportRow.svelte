<!-- web/src/components/ReportRow.svelte -->
<!-- One report's row (design/specs/2026-10-11-logs.md section 2.5): the title on the first line, one
     muted line under it that ends in the age (relative inside two weeks, a date beyond,
     `relativeDay`; a caller may pass its own `dateLabel`). Every list that shows a report uses this
     one row: the public feed, the one sample under it and "Your reports" look the same, the title
     always 16 px in the strong text colour (W.9). The caller builds the summary (`lib/logs/report-summary.ts`): "Your reports" says
     where and how many fights, the public feed says kills and guild. Shared by MyReports.svelte
     and RecentReports.svelte so the row exists in exactly one place.

     A caller that passes `character` (the reports you own) gets the player's order on the summary
     line: "today · Bow Jackzon · Darnassus · 0 fights", then the badge. `character` is the credited
     name, or null for "Not credited" in muted text; `creditHint` adds a line under it ("Log out
     once in game."). Without `character` the age stays beside the title, as the public feed has it.

     The whole row is the target: the title is the link and its ::after covers the row, so a hover
     underlines the title and moves nothing, and every row is 56 px or taller. A report still being
     processed or live shows its status as an outline pill; a complete one shows nothing. -->
<script lang="ts">
  import { relativeDay } from '../lib/dates';
  import { NOT_CREDITED_LABEL, reportStatusLabel } from '../lib/logs/report-summary';

  let {
    href,
    title,
    createdAt,
    summary,
    status = '',
    character = undefined,
    creditHint = '',
    dateLabel = '',
  }: {
    href: string;
    title: string;
    createdAt: string;
    summary: string;
    /** A report's processing status. A pill shows only while it is not yet complete. */
    status?: string;
    /** The credited character's name; null when none is credited; undefined to leave the part out. */
    character?: string | null;
    /** A line under the summary, for a row nobody is credited on yet. */
    creditHint?: string;
    /** The date as the caller words it, in place of the relative age (the sample's, which reads
     *  as the report page's own date). */
    dateLabel?: string;
  } = $props();

  const when = $derived(dateLabel !== '' ? dateLabel : relativeDay(new Date(createdAt)));
  const badge = $derived(reportStatusLabel(status));
</script>

<li
  class="report-row border-line-soft border-b"
  class:report-row-owned={character !== undefined}
  data-testid="report-row"
>
  <a {href} class="report-title text-strong text-[16px] font-semibold">{title}</a>
  <span class="text-muted flex flex-wrap items-center gap-x-2 text-[13px]">
    {#if character !== undefined}
      <span class="numeric" data-testid="report-row-when">{when}</span>
      {#if character === null}
        <span class="opacity-75" data-testid="report-row-character">{NOT_CREDITED_LABEL}</span>
      {:else}
        <span class="text-text" data-testid="report-row-character">{character}</span>
      {/if}
    {/if}
    {#if summary !== ''}<span data-testid="report-row-summary">{summary}</span>{/if}
    {#if character === undefined}
      {#if summary !== ''}<span aria-hidden="true">·</span>{/if}
      <span class="tabular" data-testid="report-row-when">{when}</span>
    {/if}
    {#if badge !== ''}
      <span class="pill pill-sample" data-testid="report-status">{badge}</span>
    {/if}
  </span>
  {#if creditHint !== ''}
    <span class="text-muted text-[12px]" data-testid="report-row-hint">{creditHint}</span>
  {/if}
</li>

<style>
  .report-row {
    position: relative;
    display: grid;
    grid-template-columns: minmax(0, 1fr) auto;
    column-gap: 16px;
    row-gap: 2px;
    align-items: baseline;
    min-height: 56px;
    padding-block: 10px;
  }
  .report-row > :not(.report-title) {
    grid-column: 1 / -1;
  }
  .report-title::after {
    content: '';
    position: absolute;
    inset: 0;
  }
  .report-row:hover .report-title {
    text-decoration: underline;
    text-underline-offset: 4px;
  }
</style>
