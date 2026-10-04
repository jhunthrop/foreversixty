<!-- web/src/components/ReportRow.svelte -->
<!-- One report's line: a title, its day, its fight and kill counts, and one line of
     trailing detail. Shared by MyReports.svelte ("Your reports", where `meta` is the
     report's own visibility and `status` shows a pill while it is still processing) and
     RecentReports.svelte (the public feed, where `meta` is the report's guild name and
     `status` is always complete, so it never shows). Extracted so the row markup that used
     to live only in MyReports.svelte exists in exactly one place.

     Player review ruling, logs mock (2026-10-04) finding 3: under 640px the facts line
     stacks under the title rather than sharing its line with it -- a title and a pinned
     date must never be left to collide on one line. `flex-col` is the default (phone); the
     facts group only joins the title's own line from `sm:` (640px) up. -->

<script lang="ts">
  import { reportDay } from '../lib/reports/hero';

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
  class="border-line-soft flex min-h-11 flex-col gap-1 border-b py-2 text-[14px] sm:flex-row sm:flex-wrap sm:items-center sm:gap-x-3 sm:gap-y-1"
  data-testid="report-row"
>
  <a {href} class="font-semibold">{title}</a>
  <span class="flex flex-wrap items-center gap-x-3 gap-y-1">
    <span class="text-muted tabular font-mono text-[13px]">{reportDay(createdAt)}</span>
    <span class="text-muted text-[13px]">
      <span class="tabular font-mono">{fightCount} fights · {killCount} kills</span>{meta === ''
        ? ''
        : ` · ${meta}`}
    </span>
    {#if status !== '' && status !== 'complete'}
      <!-- Outline, not the filled gold: that fill is the account page's Main pill, and one
           solid gold shape should mean one thing on a page. -->
      <span class="pill pill-sample" data-testid="report-status">{status}</span>
    {/if}
  </span>
</li>
