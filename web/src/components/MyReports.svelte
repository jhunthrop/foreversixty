<!-- web/src/components/MyReports.svelte -->
<!-- The reports you own. Rendered inside Account.svelte's `reports` and `account` modes so
     there is one "who is signed in" fetch per page rather than two. -->
<script lang="ts">
  import { REPORTS_PER_PAGE, listMyReports, type MyReport } from '../lib/account/api';
  import SignInPrompt from './SignInPrompt.svelte';

  /** `heading` is off where the page already titles the block, as /logs' panel does. */
  let { signedIn, heading = true }: { signedIn: boolean; heading?: boolean } = $props();

  let rows = $state<MyReport[]>([]);
  let total = $state(0);
  let page = $state(1);
  let status = $state<'loading' | 'ready' | 'failed'>('loading');

  async function load(next: number): Promise<void> {
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

  const day = (iso: string): string => iso.slice(0, 10);
  const hasMore = $derived(rows.length > 0 && page * REPORTS_PER_PAGE < total);
</script>

<section class="flex flex-col gap-3" data-testid="my-reports">
  {#if heading}<h2 class="section-title text-[18px]">Your reports</h2>{/if}
  {#if !signedIn}
    <SignInPrompt line="Sign in to see the reports you own." testid="reports-signin" />
  {:else if status === 'loading'}
    <p class="text-muted text-[14px]">Loading your reports.</p>
  {:else if status === 'failed'}
    <p class="text-[14px]" role="alert">Your reports did not load. Reload the page to try again.</p>
  {:else if rows.length === 0}
    <p class="text-muted text-[14px]">No reports yet. Upload a log or start the companion.</p>
  {:else}
    <ul class="flex flex-col">
      {#each rows as report (report.id)}
        <li
          class="border-line-soft flex min-h-11 flex-wrap items-center gap-x-3 gap-y-1 border-b py-2 text-[14px]"
        >
          <a href={`/reports/${report.id}`} class="font-semibold"
            >{report.title === '' ? report.zone : report.title}</a
          >
          <span class="text-muted tabular font-mono text-[13px]">{day(report.created_at)}</span>
          <span class="text-muted text-[13px]">
            <span class="tabular font-mono">{report.fight_count} fights · {report.kill_count} kills</span> ·
            {report.visibility}
          </span>
          {#if report.status !== 'complete'}
            <span class="pill pill-site" data-testid="report-status">{report.status}</span>
          {/if}
        </li>
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
