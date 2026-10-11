<!-- web/src/components/ReportRow.svelte -->
<!-- One report's row (design/specs/2026-10-11-logs.md section 2.5): the title and its age on the
     first line (relative inside two weeks, a date beyond, `relativeDay`), one summary line under
     the title. The caller builds the summary (`lib/logs/report-summary.ts`): "Your reports" says
     zone, bosses and length, the public feed says kills and guild. Shared by MyReports.svelte and
     RecentReports.svelte so the row exists in exactly one place.

     The whole row is the target: the title is the link and its ::after covers the row, so a hover
     underlines the title and moves nothing, and every row is 56 px or taller. A report still being
     processed shows its status as an outline pill; a complete one shows nothing. -->
<script lang="ts">
  import { relativeDay } from '../lib/dates';

  let {
    href,
    title,
    createdAt,
    summary,
    status = '',
  }: {
    href: string;
    title: string;
    createdAt: string;
    summary: string;
    /** A report's processing status. A pill shows only while it is not yet complete. */
    status?: string;
  } = $props();
</script>

<li class="report-row border-line-soft border-b" data-testid="report-row">
  <a {href} class="report-title text-text font-semibold">{title}</a>
  <span class="text-muted tabular font-mono text-[13px]" data-testid="report-row-when"
    >{relativeDay(new Date(createdAt))}</span
  >
  <span class="text-muted flex flex-wrap items-center gap-x-2 text-[12px]">
    {#if summary !== ''}<span data-testid="report-row-summary">{summary}</span>{/if}
    {#if status !== '' && status !== 'complete'}
      <span class="pill pill-sample" data-testid="report-status">{status}</span>
    {/if}
  </span>
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
  .report-row > :last-child {
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
