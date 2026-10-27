<!-- web/src/components/guild/RaidReport.svelte -->
<!-- The raid-night report (attendance spec 3.4, `<section id="report">`): the header with the officer
     buttons (`Copy for Discord`; `Post to Discord` only with a webhook, once, then its posted line),
     and the blocks in two columns from 1024 (kill order and deaths left, loot and changes right,
     consumables across) or one column in the phone order. Members read the kill order and loot only
     (ruling R-A1); officers read everything, with or without the plan. -->
<script lang="ts">
  import { onDestroy } from 'svelte';
  import type { GuildRaidView } from '../../lib/guild/api';
  import type { RaidNight, RaidReport as Report } from '../../lib/guild/attendance-contract';
  import { attendanceCopy, postedLine, postedTime } from '../../lib/guild/attendance-copy';
  import { selectOnMount } from '../../lib/guild/focus-now';
  import { discordReportText } from '../../lib/guild/night';
  import { RAID_BUTTON } from '../../lib/guild/raid-buttons';
  import { raidCopy } from '../../lib/guild/raid-copy';
  import { splitAtLink } from '../../lib/guild/raid-form';
  import { ReportCopy } from '../../lib/guild/report-copy.svelte';
  import { reportUrlOf } from '../../lib/guild/report-view';
  import GuildPlanGate from './GuildPlanGate.svelte';
  import RaidConsumables from './RaidConsumables.svelte';
  import RaidDeaths from './RaidDeaths.svelte';
  import RaidKillOrder from './RaidKillOrder.svelte';
  import RaidLoot from './RaidLoot.svelte';
  import RaidMoves from './RaidMoves.svelte';

  const COPY_FALLBACK_ROWS = 8;
  const gateLead = splitAtLink(raidCopy.gate, raidCopy.gateLink).before;
  const connectParts = splitAtLink(raidCopy.form.connectDiscord, raidCopy.form.connectDiscordLink);

  let {
    raid,
    report,
    night,
    reportId,
    guildName,
    guildPath,
    settingsHref,
    officer,
    officerTools,
    gateId,
    posting,
    postLine,
    onPost,
    viewerTimezone,
    locale,
    copy = new ReportCopy((text) => navigator.clipboard.writeText(text)),
    deathsExpanded = $bindable(false),
    consumablesExpanded = $bindable(false),
  }: {
    raid: GuildRaidView;
    report: Report;
    night: RaidNight;
    /** The matched log's id, for the kill order's fight links. */
    reportId: string;
    guildName: string;
    guildPath: string;
    settingsHref: string;
    officer: boolean;
    /** False in officer-read-only: Post is disabled, the gate sentence shows, Copy stays live. */
    officerTools: boolean;
    gateId: string;
    posting: boolean;
    postLine: string | null;
    onPost: () => void;
    viewerTimezone: string;
    locale: string;
    /** The clipboard state; a test seeds it. */
    copy?: ReportCopy;
    deathsExpanded?: boolean;
    consumablesExpanded?: boolean;
  } = $props();

  onDestroy(() => copy.stop());

  const ownKeys = $derived((raid.attendance ?? []).filter((row) => row.own).map((row) => row.character_key));
  const posted = $derived(report.posted_at);
  const copyLabel = $derived(
    copy.state === 'copied' ? attendanceCopy.report.copied : attendanceCopy.report.copy,
  );

  function copyForDiscord(): void {
    void copy.copy(
      discordReportText({
        title: raid.title,
        startsAt: raid.starts_at,
        guildName,
        tz: raid.guild_timezone,
        night,
        report,
        reportUrl: reportUrlOf(location),
        locale,
      }),
    );
  }
</script>

<section
  id="report"
  class="flex min-w-0 scroll-mt-4 flex-col gap-4"
  aria-labelledby="raid-report-heading"
  data-testid="raid-report"
>
  <div class="flex flex-wrap items-center justify-between gap-3">
    <h2 id="raid-report-heading" class="section-title text-[18px]">{attendanceCopy.report.head}</h2>
    {#if officer}
      <div class="flex flex-wrap items-center gap-3">
        <button type="button" class={RAID_BUTTON} onclick={copyForDiscord} data-testid="raid-report-copy"
          >{copyLabel}</button
        >
        <span class="sr-only" role="status" aria-live="polite" data-testid="raid-report-copied"
          >{copy.state === 'copied' ? attendanceCopy.report.copied : ''}</span
        >
        {#if posted !== null}
          <p class="text-muted text-[13px]" data-testid="raid-report-posted">
            {postedLine(postedTime(posted, viewerTimezone, locale))}
          </p>
        {:else if raid.discord.connected}
          <button
            type="button"
            class={RAID_BUTTON}
            disabled={!officerTools || posting}
            aria-describedby={officerTools ? undefined : gateId}
            onclick={onPost}
            data-testid="raid-report-post"
            >{posting ? attendanceCopy.report.posting : attendanceCopy.report.post}</button
          >
        {:else}
          <p class="text-muted text-[13px]" data-testid="raid-report-connect">
            {connectParts.before}<a href={settingsHref} class="text-text underline"
              >{raidCopy.form.connectDiscordLink}</a
            >{connectParts.after}
          </p>
        {/if}
      </div>
    {/if}
  </div>
  {#if officer && !officerTools}
    <GuildPlanGate
      tab="raids"
      {settingsHref}
      sentence={gateLead}
      id={gateId}
      testid="guild-raids-plan-gate-report"
    />
  {/if}
  {#if postLine !== null}
    <p class="text-wipe text-[13px]" role="alert" data-testid="raid-report-post-error">{postLine}</p>
  {/if}
  {#if copy.state === 'failed'}
    <div class="flex flex-col gap-2" data-testid="raid-report-copy-failed">
      <p class="text-wipe text-[13px]" role="alert">{attendanceCopy.errors.copyFailed}</p>
      <textarea
        class="border-line-warm bg-raised rounded-control text-text w-full p-3 font-mono text-[12px]"
        rows={COPY_FALLBACK_ROWS}
        readonly
        aria-label={attendanceCopy.report.copy}
        use:selectOnMount>{copy.text}</textarea
      >
    </div>
  {/if}

  <div class="flex flex-col gap-8 lg:grid lg:grid-cols-2 lg:gap-x-8">
    <div class="contents lg:flex lg:min-w-0 lg:flex-col lg:gap-8">
      <div class="order-1 lg:order-none">
        <RaidKillOrder {report} {reportId} />
      </div>
      {#if officer}
        <div class="order-3 lg:order-none">
          <RaidDeaths rows={report.deaths} {ownKeys} bind:expanded={deathsExpanded} />
        </div>
      {/if}
    </div>
    <div class="contents lg:flex lg:min-w-0 lg:flex-col lg:gap-8">
      <div class="order-2 lg:order-none">
        <RaidLoot groups={report.loot} {officer} lootHref={`${guildPath}#loot`} {ownKeys} />
      </div>
      {#if officer}
        <div class="order-4 lg:order-none">
          <RaidMoves
            moves={report.moves}
            zoneName={raid.zone_name}
            guildTimezone={raid.guild_timezone}
            {locale}
            {ownKeys}
          />
        </div>
      {/if}
    </div>
    {#if officer}
      <div class="order-5 lg:order-none lg:col-span-2">
        <RaidConsumables consumables={report.consumables} {ownKeys} bind:expanded={consumablesExpanded} />
      </div>
    {/if}
  </div>
</section>
