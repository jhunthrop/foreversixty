<!-- web/src/components/ReportRow.svelte -->
<!-- One report's row (logs spec 2026-11-07 sections 4 and 5), the ui Row as a link to the report:
     the title in the link gold (directions b1-logs), one muted line under it (the caller builds it: lib/logs/report-summary.ts), and at
     the right a muted status Tag only while the report is not finished (Live, Processing,
     Uploading; a complete or ready report carries none, MUST 4) or the caller's own Tag (the
     sample's), then the age (relative inside two weeks, a date beyond, `relativeDay`; a caller
     may pass its own `dateLabel`). Every list that shows a report draws this one row. -->
<script lang="ts">
  import { relativeDay } from '../lib/dates';
  import { reportStatusLabel } from '../lib/logs/report-summary';
  import Row from './ui/Row.svelte';
  import Tag from './ui/Tag.svelte';

  let {
    href,
    title,
    createdAt,
    summary,
    status = '',
    tag = '',
    creditHint = '',
    dateLabel = '',
    open = false,
  }: {
    href: string;
    title: string;
    createdAt: string;
    summary: string;
    /** A report's processing status. A Tag shows only while it is not yet finished. */
    status?: string;
    /** A Tag of the caller's own (the sample's "Sample report"). */
    tag?: string;
    /** A line under the summary, for a live row nobody is credited on yet. */
    creditHint?: string;
    /** The date as the caller words it, in place of the relative age. */
    dateLabel?: string;
    /** An open Section's row (no inline padding), not a Plate's. */
    open?: boolean;
  } = $props();

  const when = $derived(dateLabel !== '' ? dateLabel : relativeDay(new Date(createdAt)));
  const badge = $derived(reportStatusLabel(status));
</script>

<Row {title} {href} {open} testid="report-row">
  {#snippet titleContent()}<span class="text-gold">{title}</span>{/snippet}
  {#snippet subContent()}
    <span data-testid="report-row-summary">{summary}</span>
    {#if creditHint !== ''}<span class="block text-[12px]" data-testid="report-row-hint">{creditHint}</span
      >{/if}
  {/snippet}
  {#snippet trail()}
    {#if badge !== ''}<span data-testid="report-status"><Tag tone="muted">{badge}</Tag></span>
    {:else if tag !== ''}<Tag tone="muted">{tag}</Tag>{/if}
    <span class="text-small text-muted tabular-nums" data-testid="report-row-when">{when}</span>
  {/snippet}
</Row>
