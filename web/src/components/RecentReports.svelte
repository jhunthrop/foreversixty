<!-- web/src/components/RecentReports.svelte -->
<!-- The public "Recent public reports" feed: GET /v1/reports/recent, needing no session.
     Mounted on /logs above "Your reports" in full, and -- per spec section 4 -- meant for
     a compact mount on the homepage too; `compact` trims it to a handful of rows with no
     pagination, for a panel that is not the page's main content. See ReportRow.svelte for
     the shared row markup this shares with MyReports.svelte ("Your reports"). -->
<script lang="ts">
  import { recentReportsCopy } from '../lib/reports/copy';
  import { REPORTS_LOADING_MIN_H } from '../lib/reports/layout';
  import { fetchRecentReports, type RecentReport } from '../lib/reports/recent';
  import ReportRow from './ReportRow.svelte';
  import EmptyState from './ui/EmptyState.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  /** `heading` is off where the page already titles the block, as /logs' panel does --
   *  the same convention MyReports.svelte uses for "Your reports". */
  let { compact = false, heading = true }: { compact?: boolean; heading?: boolean } = $props();

  /** How many rows a compact mount shows. There is no "older reports" at this size. */
  const COMPACT_ROWS = 5;

  let rows = $state<RecentReport[]>([]);
  let nextCursor = $state<string | undefined>(undefined);
  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  // The cursor the in-flight (or last-failed) request used, so Try again re-fires that
  // page rather than silently bouncing the reader back to the first one -- the same fix
  // MyReports.svelte's `attemptedPage` makes for its page number.
  let attemptedCursor = $state<string | undefined>(undefined);

  async function load(cursor?: string): Promise<void> {
    attemptedCursor = cursor;
    status = 'loading';
    try {
      const result = await fetchRecentReports(cursor);
      rows = compact ? result.rows.slice(0, COMPACT_ROWS) : result.rows;
      nextCursor = result.next_cursor;
      status = 'ready';
    } catch {
      status = 'failed';
    }
  }

  $effect(() => {
    void load();
  });

  const hasMore = $derived(!compact && nextCursor !== undefined);
</script>

<section class="flex flex-col gap-3" data-testid="recent-reports">
  {#if heading}<h2 class="section-title text-[18px]">{recentReportsCopy.heading}</h2>{/if}
  {#if status === 'loading'}
    <!-- Five rows, not the page's real row count: the feed has no fixed cap, and a
         hundred shimmering rows would be its own kind of noise. `minHeight` carries the
         reservation instead, sized to a phone screenful of rows the way Rankings.svelte's
         own 8-line skeleton reserves 440px -- /logs is one of the CLS-0.05 URLs in
         lighthouserc.json, so the reserve, not the row count, is what has to hold. -->
    <Skeleton lines={5} rowHeight="h-11" minHeight={REPORTS_LOADING_MIN_H} testid="recent-reports-skeleton" />
  {:else if status === 'failed'}
    <LoadError
      message={recentReportsCopy.failed}
      onRetry={() => void load(attemptedCursor)}
      testid="recent-reports-error"
    />
  {:else if rows.length === 0}
    <EmptyState message={recentReportsCopy.empty} testid="recent-reports-empty" />
  {:else}
    <ul class="reveal flex flex-col">
      {#each rows as report (report.id)}
        <ReportRow
          href={`/reports/${report.id}`}
          title={report.title}
          createdAt={report.created_at}
          fightCount={report.fight_count}
          killCount={report.kill_count}
          meta={report.guild_name ?? ''}
        />
      {/each}
    </ul>
    {#if hasMore}
      <button
        class="border-line-warm rounded-control text-text inline-flex h-11 w-fit items-center border px-4 text-[12px] font-bold tracking-[0.06em] uppercase"
        data-testid="recent-reports-older"
        onclick={() => void load(nextCursor)}
      >
        {recentReportsCopy.older}
      </button>
    {/if}
  {/if}
</section>
