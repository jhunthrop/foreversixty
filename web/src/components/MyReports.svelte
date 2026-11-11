<!-- web/src/components/MyReports.svelte -->
<!-- The reports you own, for a signed-in visitor (logs spec 2026-11-07 sections 4 and 5): the page
     that mounts it has read the session and says what a signed-out visitor needs. Rows are
     ReportRow (the ui Row), newest first, `Older reports` at the foot when the API has more.
     `plate` puts every state in the page's one Plate (/reports); without it the rows are the
     Plate on /logs and the empty line is open text in its Section. A guild member gets the guild
     TabStrip (`Yours`, each guild) over the list; each guild's tab lists that guild's reports. -->
<script lang="ts">
  import {
    REPORTS_PER_PAGE,
    listGuildReports,
    listMyReports,
    reportNightAt,
    type MyReport,
  } from '../lib/account/api';
  import { nightDate } from '../lib/dates';
  import { LOG_OUT_HINT, creditedName, ownReportRow } from '../lib/logs/report-summary';
  import { logsCopy } from '../lib/reports/copy';
  import { myReportsCopy } from '../lib/reports/my-reports-copy';
  import ReportRow from './ReportRow.svelte';
  import EmptyState from './ui/EmptyState.svelte';
  import LoadError from './ui/LoadError.svelte';
  import LoadOlder from './ui/LoadOlder.svelte';
  import Plate from './ui/Plate.svelte';
  import Skeleton from './ui/Skeleton.svelte';
  import TabStrip from './ui/TabStrip.svelte';

  type Link = { label: string; href: string };

  let {
    guilds = [],
    plate,
    open = false,
    nightDates = false,
    emptyMessage: ownEmptyMessage = myReportsCopy.empty,
    emptyAction,
    emptySecondary,
  }: {
    /** Every guild this account belongs to (`Me.guilds`): one tab each beside `Yours`. */
    guilds?: readonly { id: number; name: string }[];
    /** The Plate's accessible name: every state sits in the page's one Plate (/reports). */
    plate?: string;
    /** The rows open, in the caller's own panel (/logs' Your reports Section panel). */
    open?: boolean;
    /** The account page's rows print the night's date ("Tue 20 Oct") instead of the age. */
    nightDates?: boolean;
    /** The `Yours` tab's empty sentence, where the page says it in its own words. */
    emptyMessage?: string;
    /** The empty state's primary action (/reports: Upload a log). */
    emptyAction?: Link;
    /** A quieter second link beside it (/reports: the sample report). */
    emptySecondary?: Link;
  } = $props();

  const YOURS = 'yours';

  /** `null` is the `Yours` tab (`listMyReports`); a guild id selects that guild's own list. */
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
      rows = next === 1 ? result.rows : [...rows, ...result.rows];
      total = result.total;
      page = result.page;
      status = 'ready';
    } catch {
      status = 'failed';
    }
  }

  function selectTab(id: string): void {
    const guildId = id === YOURS ? null : Number(id);
    if (guildId === activeGuild) return;
    activeGuild = guildId;
    rows = [];
    void load(1);
  }

  $effect(() => {
    void load(1);
  });

  /** The `Yours` tab lists reports the viewer owns; a guild's tab lists other people's. */
  const ownsRows = $derived(activeGuild === null);
  const hasMore = $derived(rows.length > 0 && page * REPORTS_PER_PAGE < total);
  const activeGuildName = $derived(guilds.find((guild) => guild.id === activeGuild)?.name ?? '');
  const emptyMessage = $derived(
    activeGuild === null ? ownEmptyMessage : logsCopy.guildEmpty(activeGuildName),
  );
  /** The designed empty: its first sentence leads, the rest says what fills it. */
  const emptyLead = $derived(emptyMessage.split(/(?<=\.)\s/, 1)[0]);
  const emptyFill = $derived(emptyMessage.slice(emptyLead.length).trim());
  const tabs = $derived([
    { id: YOURS, label: 'Yours' },
    ...guilds.map((guild) => ({ id: String(guild.id), label: guild.name, testid: 'my-reports-guild-tab' })),
  ]);
</script>

{#snippet list()}
  {#each rows as report (report.id)}
    {@const row = ownReportRow({
      title: report.title,
      zone: report.zone,
      startedAt: reportNightAt(report),
      fightCount: report.fight_count,
      killCount: report.kill_count,
      bossCount: report.boss_count,
      loggingCharacter: report.logging_character,
    })}
    <ReportRow
      href={`/reports/${report.id}`}
      title={row.title}
      createdAt={report.created_at}
      dateLabel={nightDates ? nightDate(new Date(reportNightAt(report))) : ''}
      summary={row.sub}
      {open}
      status={report.status}
      creditHint={ownsRows && report.status === 'live' && creditedName(report.logging_character) === null
        ? LOG_OUT_HINT
        : ''}
    />
  {/each}
{/snippet}

{#snippet state()}
  {#if status === 'loading' && rows.length === 0}
    <div class="p-4"><Skeleton lines={3} rowHeight="h-11" testid="my-reports-skeleton" /></div>
  {:else if status === 'failed' && rows.length === 0}
    <div class="p-4">
      <LoadError
        message={myReportsCopy.failed}
        onRetry={() => void load(attemptedPage)}
        testid="my-reports-error"
      />
    </div>
  {:else if rows.length === 0}
    <div class="p-4">
      {#if activeGuild === null && emptyAction}
        <EmptyState
          message={emptyLead}
          fill={emptyFill === '' ? undefined : emptyFill}
          action={emptyAction}
          secondary={emptySecondary}
          testid="my-reports-empty"
        />
      {:else}
        <EmptyState message={emptyMessage} testid="my-reports-empty" />
      {/if}
    </div>
  {:else}
    {@render list()}
  {/if}
{/snippet}

<div class="flex min-w-0 flex-col gap-3" data-testid="my-reports">
  {#if guilds.length > 0}
    <TabStrip
      {tabs}
      active={activeGuild === null ? YOURS : String(activeGuild)}
      label="Reports"
      onSelect={selectTab}
      testid="my-reports-guild-tabs"
    />
  {/if}
  {#if plate !== undefined}
    <Plate label={plate}>{@render state()}</Plate>
  {:else if rows.length > 0 && open}
    <div>{@render list()}</div>
  {:else if rows.length > 0}
    <Plate label="Your reports">{@render list()}</Plate>
  {:else}
    {@render state()}
  {/if}
  {#if hasMore}
    <LoadOlder
      label={myReportsCopy.older}
      loading={status === 'loading'}
      failed={status === 'failed'}
      onLoad={() => void load(status === 'failed' ? attemptedPage : page + 1)}
      testid="my-reports-older"
    />
  {/if}
</div>
