<!-- web/src/components/MyReports.svelte -->
<!-- The reports you own. Rendered inside Account.svelte's `reports` and `account` modes so
     there is one "who is signed in" fetch per page rather than two. -->
<script lang="ts">
  import { REPORTS_PER_PAGE, listMyReports, type MyReport } from '../lib/account/api';
  import { REPORTS_LOADING_MIN_H } from '../lib/reports/layout';
  import { myReportsCopy } from '../lib/reports/my-reports-copy';
  import ReportRow from './ReportRow.svelte';
  import SignInPrompt from './SignInPrompt.svelte';
  import EmptyState from './ui/EmptyState.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  /** `heading` is off where the page already titles the block, as /logs' panel does. */
  let { signedIn, heading = true }: { signedIn: boolean; heading?: boolean } = $props();

  let rows = $state<MyReport[]>([]);
  let total = $state(0);
  let page = $state(1);
  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let attemptedPage = $state(1);

  async function load(next: number): Promise<void> {
    attemptedPage = next;
    status = 'loading';
    try {
      const result = await listMyReports(next);
      rows = result.rows;
      total = result.total;
      page = result.page;
      status = 'ready';
    } catch {
      status = 'failed';
    }
  }

  $effect(() => {
    if (signedIn) void load(1);
    else status = 'ready';
  });

  const hasMore = $derived(rows.length > 0 && page * REPORTS_PER_PAGE < total);
</script>

<section class="flex flex-col gap-3" data-testid="my-reports">
  {#if heading}<h2 class="section-title text-[18px]">Your reports</h2>{/if}
  {#if !signedIn}
    <SignInPrompt line="Sign in to see the reports you own." testid="reports-signin" />
  {:else if status === 'loading'}
    <!-- Five rows, not REPORTS_PER_PAGE's hundred: a hundred shimmering rows would be
         its own kind of noise, so the row count is a legible stand-in and
         REPORTS_LOADING_MIN_H carries the reservation /logs' CLS 0.05 budget needs. -->
    <Skeleton lines={5} rowHeight="h-11" minHeight={REPORTS_LOADING_MIN_H} testid="my-reports-skeleton" />
  {:else if status === 'failed'}
    <LoadError
      message={myReportsCopy.failed}
      onRetry={() => void load(attemptedPage)}
      testid="my-reports-error"
    />
  {:else if rows.length === 0}
    <EmptyState
      message={myReportsCopy.empty}
      action={{ label: myReportsCopy.uploadAction, href: '/logs' }}
      testid="my-reports-empty"
    />
  {:else}
    <ul class="reveal flex flex-col">
      {#each rows as report (report.id)}
        <ReportRow
          href={`/reports/${report.id}`}
          title={report.title === '' ? report.zone : report.title}
          createdAt={report.created_at}
          fightCount={report.fight_count}
          killCount={report.kill_count}
          meta={report.visibility}
          status={report.status}
        />
      {/each}
    </ul>
    {#if hasMore}
      <button
        class="border-line-warm rounded-control text-text inline-flex h-11 w-fit items-center border px-4 text-[12px] font-bold tracking-[0.06em] uppercase"
        onclick={() => void load(page + 1)}
      >
        Older reports
      </button>
    {/if}
  {/if}
</section>
