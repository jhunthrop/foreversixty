<!-- web/src/components/ReportRow.svelte -->
<!-- One report's line: a title, its day, its fight and kill counts, and one line of
     trailing detail. Shared by MyReports.svelte ("Your reports", where `meta` is the
     report's own visibility and `status` shows a pill while it is still processing) and
     RecentReports.svelte (the public feed, where `meta` is the report's guild name and
     `status` is always complete, so it never shows). Extracted so the row markup that used
     to live only in MyReports.svelte exists in exactly one place. -->
<script lang="ts">
  const day = (iso: string): string => iso.slice(0, 10);

  let {
    href,
    title,
    createdAt,
    fightCount,
    killCount,
    meta = '',
    status = '',
  }: {
    href: string;
    title: string;
    createdAt: string;
    fightCount: number;
    killCount: number;
    /** Trailing detail after the fight/kill count: a visibility, a guild name, or nothing. */
    meta?: string;
    /** A report's processing status. A pill shows only while it is not yet complete. */
    status?: string;
  } = $props();
</script>

<li
  class="border-line-soft flex min-h-11 flex-wrap items-center gap-x-3 gap-y-1 border-b py-2 text-[14px]"
  data-testid="report-row"
>
  <a {href} class="font-semibold">{title}</a>
  <span class="text-muted tabular font-mono text-[13px]">{day(createdAt)}</span>
  <span class="text-muted text-[13px]">
    <span class="tabular font-mono">{fightCount} fights · {killCount} kills</span>{meta === ''
      ? ''
      : ` · ${meta}`}
  </span>
  {#if status !== '' && status !== 'complete'}
    <span class="pill pill-site" data-testid="report-status">{status}</span>
  {/if}
</li>
