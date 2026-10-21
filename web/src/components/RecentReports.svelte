<!-- web/src/components/RecentReports.svelte -->
<!-- The public "Recent public reports" feed: GET /v1/reports/recent, needing no session.
     Mounted on /logs below "Your reports"; `compact` trims it to a handful of rows with no
     pagination, for a panel that is not the page's main content. Under the feed sits the
     site's one sample report as a labelled row of its own. See ReportRow.svelte for the shared
     row markup this shares with MyReports.svelte ("Your reports"). -->
<script lang="ts">
  import { recentReportsCopy } from '../lib/reports/copy';
  import { RECENT_REPORTS_LOADING_MIN_H } from '../lib/reports/layout';
  import { fetchRecentReports, isPublicListable, type RecentReport } from '../lib/reports/recent';
  import { publicReportRow } from '../lib/logs/report-summary';
  import { SAMPLE_FALLBACK_ROW, SAMPLE_LABEL, SAMPLE_REPORT_ID, loadSampleRow } from '../lib/logs/sample-row';
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
      // Logs spec (2026-10-11) section 9: the canonical sample report is shown once, as its own
      // labelled row under the feed -- it must never double as a feed entry, or this feed's
      // genuinely-empty state (and its already-correct empty copy) is never reachable.
      const filtered = result.rows.filter((row) => row.id !== SAMPLE_REPORT_ID && isPublicListable(row));
      rows = compact ? filtered.slice(0, COMPACT_ROWS) : filtered;
      nextCursor = result.next_cursor;
      status = 'ready';
    } catch {
      status = 'failed';
    }
  }

  let sample = $state(SAMPLE_FALLBACK_ROW);

  $effect(() => {
    void load();
    void loadSampleRow().then((row) => {
      sample = row;
    });
  });

  const hasMore = $derived(!compact && nextCursor !== undefined);
</script>

<section class="flex flex-col gap-3" data-testid="recent-reports">
  {#if heading}<h2 class="section-title text-[18px]">{recentReportsCopy.heading}</h2>{/if}
  {#if status === 'loading'}
    <!-- One line over a reserved sentence's height, not a screenful of rows: the feed's ready
         state is its empty sentence (RECENT_REPORTS_LOADING_MIN_H), and the sample row under it
         must not move when the feed answers. -->
    <Skeleton
      lines={1}
      minHeight={RECENT_REPORTS_LOADING_MIN_H}
      label={recentReportsCopy.placeholderLabel}
      testid="recent-reports-skeleton"
    />
  {:else if status === 'failed'}
    <LoadError
      message={recentReportsCopy.failed}
      onRetry={() => void load(attemptedCursor)}
      testid="recent-reports-error"
    />
  {:else if rows.length === 0 && compact}
    <!-- The landing's panel with nothing in the feed shows only the sample row under it (P.1). -->
  {:else if rows.length === 0}
    <EmptyState message={recentReportsCopy.empty} testid="recent-reports-empty" />
  {:else}
    <!-- No `reveal` fade on the rows: the sample row under them is drawn at once, and a feed
         row caught mid-fade read as a dimmed, disabled item beside it (W.9). -->
    <ul class="flex flex-col" data-testid="recent-reports-list">
      {#each rows as report (report.id)}
        {@const row = publicReportRow({
          title: report.title,
          zone: report.zone,
          killCount: report.kill_count,
          bossCount: report.boss_count,
          fightCount: report.fight_count,
          guildName: report.guild_name,
          loggingCharacter: report.logging_character,
        })}
        <ReportRow
          href={`/reports/${report.id}`}
          title={row.title}
          createdAt={report.created_at}
          summary={row.summary}
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
  <div class="flex flex-col" data-testid="recent-reports-sample">
    <span class="label text-muted">{SAMPLE_LABEL}</span>
    <ul>
      <ReportRow
        href={`/reports/${sample.id}`}
        title={sample.title}
        createdAt={sample.createdAt}
        dateLabel={sample.dateLabel}
        summary={sample.summary}
      />
    </ul>
  </div>
</section>
