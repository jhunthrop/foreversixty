<!-- web/src/components/RecentReports.svelte -->
<!-- /logs' "Recent public reports" Section (logs spec 2026-11-07 section 4): the public feed as open
     report Rows (none yet: the sample row alone, no empty sentence over it), then the one sample report as its own Row with the
     `Sample report` Tag (bundled with the site, read from its static meta; never a feed entry,
     so the feed's own empty state stays reachable). `Older reports` pages the feed in place. The
     viewer's own reports (`exclude`) sit in Your reports above and are left out here. -->
<script lang="ts">
  import { recentReportsCopy } from '../lib/reports/copy';
  import { reportCopy } from '../lib/report/copy';
  import { RECENT_REPORTS_LOADING_MIN_H } from '../lib/reports/layout';
  import { fetchRecentReports, isPublicListable, type RecentReport } from '../lib/reports/recent';
  import { publicReportRow } from '../lib/logs/report-summary';
  import { SAMPLE_FALLBACK_ROW, SAMPLE_REPORT_ID, loadSampleRow } from '../lib/logs/sample-row';
  import ReportRow from './ReportRow.svelte';
  import Button from './ui/Button.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Section from './ui/Section.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  /** The viewer's own report ids: listed just above on /logs, so the feed leaves them out. */
  let { exclude = new Set<string>() }: { exclude?: ReadonlySet<string> } = $props();

  let rows = $state<RecentReport[]>([]);
  let nextCursor = $state<string | undefined>(undefined);
  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  // The cursor the in-flight (or last-failed) request used, so Try again re-fires that page.
  let attemptedCursor = $state<string | undefined>(undefined);
  let sample = $state(SAMPLE_FALLBACK_ROW);

  async function load(cursor?: string): Promise<void> {
    attemptedCursor = cursor;
    status = 'loading';
    try {
      const result = await fetchRecentReports(cursor);
      // The sample is shown once, as its own row under the feed, never as a feed entry.
      const listed = result.rows.filter((row) => row.id !== SAMPLE_REPORT_ID && isPublicListable(row));
      rows = cursor === undefined ? listed : [...rows, ...listed];
      nextCursor = result.next_cursor;
      status = 'ready';
    } catch {
      status = 'failed';
    }
  }

  const shown = $derived(rows.filter((row) => !exclude.has(row.id)));

  $effect(() => {
    void load();
    void loadSampleRow().then((row) => {
      sample = row;
    });
  });
</script>

<Section title={recentReportsCopy.heading} panel testid="recent-reports">
  {#if status === 'loading' && rows.length === 0}
    <Skeleton
      lines={1}
      minHeight={RECENT_REPORTS_LOADING_MIN_H}
      label={recentReportsCopy.placeholderLabel}
      testid="recent-reports-skeleton"
    />
  {:else if status === 'failed' && rows.length === 0}
    <div class="p-4">
      <LoadError
        message={recentReportsCopy.failed}
        onRetry={() => void load(attemptedCursor)}
        testid="recent-reports-error"
      />
    </div>
  {/if}
  <div class="flex flex-col" data-testid="recent-reports-list">
    {#each shown as report (report.id)}
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
        open
      />
    {/each}
    <div data-testid="recent-reports-sample">
      <ReportRow
        href={`/reports/${sample.id}`}
        title={sample.title}
        createdAt={sample.createdAt}
        dateLabel={sample.dateLabel}
        summary={sample.summary}
        tag={sample.createdAt === '' ? '' : reportCopy.sampleBadge}
        open
      />
    </div>
  </div>
  {#if nextCursor !== undefined && status !== 'failed'}
    <div class="mt-3">
      <Button size="sm" onclick={() => void load(nextCursor)} testid="recent-reports-older">
        {recentReportsCopy.older}
      </Button>
    </div>
  {/if}
</Section>
