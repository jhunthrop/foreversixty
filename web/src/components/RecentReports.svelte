<!-- web/src/components/RecentReports.svelte -->
<!-- The public "Recent public reports" feed: GET /v1/reports/recent, needing no session.
     Mounted on /logs above "Your reports" in full, and -- per spec section 4 -- meant for
     a compact mount on the homepage too; `compact` trims it to a handful of rows with no
     pagination, for a panel that is not the page's main content. See ReportRow.svelte for
     the shared row markup this shares with MyReports.svelte ("Your reports"). -->
<script lang="ts">
  import { recentReportsCopy } from '../lib/reports/copy';
  import { fetchRecentReports, type RecentReport } from '../lib/reports/recent';
  import ReportRow from './ReportRow.svelte';

  /** `heading` is off where the page already titles the block, as /logs' panel does --
   *  the same convention MyReports.svelte uses for "Your reports". */
  let { compact = false, heading = true }: { compact?: boolean; heading?: boolean } = $props();

  /** How many rows a compact mount shows. There is no "older reports" at this size. */
  const COMPACT_ROWS = 5;

  let rows = $state<RecentReport[]>([]);
  let nextCursor = $state<string | undefined>(undefined);
  let status = $state<'loading' | 'ready' | 'failed'>('loading');

  async function load(cursor?: string): Promise<void> {
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
    <p class="text-muted text-[14px]">{recentReportsCopy.loading}</p>
  {:else if status === 'failed'}
    <p class="text-[14px]" role="alert">{recentReportsCopy.failed}</p>
  {:else if rows.length === 0}
    <p class="text-muted text-[14px]" data-testid="recent-reports-empty">{recentReportsCopy.empty}</p>
  {:else}
    <ul class="flex flex-col">
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
