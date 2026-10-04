<!-- web/src/components/MyReports.svelte -->
<!-- The reports you own. Rendered inside Account.svelte's `reports` and `account` modes so
     there is one "who is signed in" fetch per page rather than two.

     Logs landing spec (2026-10-04) §4.C.1 adds the guild tab strip: when `guilds` is
     non-empty, a small pill row (`Mine` + one pill per guild, the same active-pill
     convention `ClassHeader.astro`'s own band strip already uses) switches the list
     between `listMyReports` and `listGuildReports(guild.id)`. `heading` is off where the
     page already titles the block, as /logs' panel does. -->
<script lang="ts">
  import { REPORTS_PER_PAGE, listGuildReports, listMyReports, type MyReport } from '../lib/account/api';
  import { logsCopy } from '../lib/reports/copy';
  import { REPORTS_LOADING_MIN_H } from '../lib/reports/layout';
  import { myReportsCopy } from '../lib/reports/my-reports-copy';
  import ReportRow from './ReportRow.svelte';
  import SignInPrompt from './SignInPrompt.svelte';
  import EmptyState from './ui/EmptyState.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  let {
    signedIn,
    heading = true,
    guilds = [],
  }: {
    signedIn: boolean;
    heading?: boolean;
    /** §4.C.1's own tab strip source: `Me.guilds`, every guild this account belongs to. */
    guilds?: readonly { id: number; name: string }[];
  } = $props();

  /** `null` is the "Mine" tab (`listMyReports`); a guild id selects that guild's own tab
   *  (`listGuildReports`). Kept as one piece of state rather than an index into `guilds` so
   *  switching tabs never has to guess which list a stale index pointed at. */
  let activeGuild = $state<number | null>(null);

  let rows = $state<MyReport[]>([]);
  let total = $state(0);
  let page = $state(1);
  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let attemptedPage = $state(1);

  async function load(next: number): Promise<void> {
    attemptedPage = next;
    status = 'loading';
    try {
      const result =
        activeGuild === null ? await listMyReports(next) : await listGuildReports(activeGuild, next);
      rows = result.rows;
      total = result.total;
      page = result.page;
      status = 'ready';
    } catch {
      status = 'failed';
    }
  }

  function selectTab(guildId: number | null): void {
    if (guildId === activeGuild) return;
    activeGuild = guildId;
    void load(1);
  }

  $effect(() => {
    if (signedIn) void load(1);
    else status = 'ready';
  });

  const hasMore = $derived(rows.length > 0 && page * REPORTS_PER_PAGE < total);
  const activeGuildName = $derived(guilds.find((guild) => guild.id === activeGuild)?.name ?? '');
  const emptyMessage = $derived(
    activeGuild === null ? myReportsCopy.empty : logsCopy.guildEmpty(activeGuildName),
  );
</script>

<section class="flex flex-col gap-3" data-testid="my-reports">
  {#if heading}<h2 class="section-title text-[18px]">Your reports</h2>{/if}
  {#if !signedIn}
    <SignInPrompt line="Sign in to see the reports you own." testid="reports-signin" />
  {:else}
    {#if guilds.length > 0}
      <!-- Horizontally scrolling on phone, same rule every other pill row on the site
           follows (spec §7); never wraps, so every pill keeps its 44px hit height. -->
      <div
        class="my-reports-tabs -mx-1 flex gap-2 overflow-x-auto px-1"
        role="tablist"
        aria-label="Reports"
        data-testid="my-reports-guild-tabs"
      >
        <button
          type="button"
          role="tab"
          aria-selected={activeGuild === null}
          class="rounded-control inline-flex h-11 shrink-0 items-center border px-4 text-[13px] font-semibold"
          class:border-gold={activeGuild === null}
          class:text-strong={activeGuild === null}
          class:border-line={activeGuild !== null}
          style={activeGuild === null ? 'background: #e5b95514' : undefined}
          onclick={() => selectTab(null)}
        >
          Mine
        </button>
        {#each guilds as guild (guild.id)}
          <button
            type="button"
            role="tab"
            aria-selected={activeGuild === guild.id}
            class="rounded-control inline-flex h-11 shrink-0 items-center border px-4 text-[13px] font-semibold"
            class:border-gold={activeGuild === guild.id}
            class:text-strong={activeGuild === guild.id}
            class:border-line={activeGuild !== guild.id}
            style={activeGuild === guild.id ? 'background: #e5b95514' : undefined}
            data-testid="my-reports-guild-tab"
            onclick={() => selectTab(guild.id)}
          >
            {guild.name}
          </button>
        {/each}
      </div>
    {/if}
    {#if status === 'loading'}
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
        message={emptyMessage}
        action={activeGuild === null ? { label: myReportsCopy.uploadAction, href: '/logs' } : undefined}
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
  {/if}
</section>

<style>
  /* Scrollbar hidden, same convention ClassHeader.astro's own band-tab strip uses, so a
     phone-width tab row scrolls (spec §7) without a visible scrollbar track. */
  .my-reports-tabs {
    scrollbar-width: none;
  }
  .my-reports-tabs::-webkit-scrollbar {
    display: none;
  }
</style>
