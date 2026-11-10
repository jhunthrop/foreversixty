<!-- web/src/components/guild/report/RaidNightReport.svelte -->
<!-- The done raid (web layout spec 2026-11-05 §7.2 "Raid report"): "How did the night go and who
     needs a word?", mock web-guild-report, built only from the ui/ primitives. Page split: the back link
     and the head (the raid, its Done Tag and date, the answer: kills, wipes, attempts, deaths, the
     guild tabs, as the planner's), then for a
     member Your night, Pulls (framed, each pull's deaths in the Inspector), for officers the deaths
     by raider, DPS and HPS by boss kill and the consumables by pull. Rail: Attendance (the match line, the summary, and for
     officers the fix rows behind Fix attendance), Loot awarded, the Discord summary (officers), the
     notes. A member's report carries only their own deaths,
     output and consumables (the API scopes it), and nothing here lists other raiders worst first.
     A matched log that turned private keeps the head and the match line and draws nothing else. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import type { GuildRaidView } from '../../../lib/guild/api';
  import { attendanceCopy } from '../../../lib/guild/attendance-copy';
  import { discordReportText } from '../../../lib/guild/night';
  import type { GateFacts } from '../../../lib/guild/plan-gate';
  import { eyebrowOf, kindOf, raidCopy, SEP, zoneLineOf } from '../../../lib/guild/raid-copy';
  import { raidWhen, type RaidPageFlags } from '../../../lib/guild/raid-page';
  import type { RaidPageController } from '../../../lib/guild/raid-page-controller.svelte';
  import { plannerStatusTag } from '../../../lib/guild/raid-planner-view';
  import { headlineFigures } from '../../../lib/guild/report-night-view';
  import { reportUrlOf } from '../../../lib/guild/report-view';
  import FactionCrest from '../../character/FactionCrest.svelte';
  import Figure from '../../ui/Figure.svelte';
  import Page from '../../ui/Page.svelte';
  import PageHeader from '../../ui/PageHeader.svelte';
  import Section from '../../ui/Section.svelte';
  import Tag from '../../ui/Tag.svelte';
  import GuildPlanGate from '../GuildPlanGate.svelte';
  import GuildTabs from '../GuildTabs.svelte';
  import RaidDoneBand from '../RaidDoneBand.svelte';
  import ReportAttendance from './ReportAttendance.svelte';
  import ReportConsumables from './ReportConsumables.svelte';
  import ReportDeaths from './ReportDeaths.svelte';
  import ReportDiscord from './ReportDiscord.svelte';
  import ReportLoot from './ReportLoot.svelte';
  import ReportOutput from './ReportOutput.svelte';
  import ReportPulls from './ReportPulls.svelte';
  import ReportYourNight from './ReportYourNight.svelte';

  let {
    c,
    raid,
    flags,
    now,
    viewerTimezone,
    locale,
    guildName,
    guildPath,
    settingsHref,
    gateFacts,
    gateId,
  }: {
    c: RaidPageController;
    raid: GuildRaidView;
    flags: RaidPageFlags;
    now: Date;
    viewerTimezone: string;
    locale: string;
    guildName: string;
    guildPath: string;
    settingsHref: string;
    gateFacts: GateFacts;
    gateId: string;
  } = $props();

  const OPEN_NIGHT_ID = 'raid-open-night';
  const officer = $derived(c.role === 'officer');
  const kind = $derived(kindOf(raid));
  const when = $derived(raidWhen(raid, now, viewerTimezone, locale));
  const status = $derived(plannerStatusTag(flags.status));
  const guildTimezone = $derived(raid.guild_timezone ?? viewerTimezone);
  const report = $derived(raid.report);
  const night = $derived(raid.night);
  /** Your night is a member's: officers read their rows in the full report. */
  const showYourNight = $derived(!officer && (raid.attendance ?? []).some((row) => row.own));

  // A poll that matched the log while the officer was in the match picker swaps the planner for
  // this page; focus would fall to the body, so it goes to the match line's link instead.
  onMount(() => {
    if (c.logMatched && document.activeElement === document.body) {
      document.getElementById(OPEN_NIGHT_ID)?.focus();
    }
  });

  function copyText(): string {
    if (report === null || night === null) return '';
    return discordReportText({
      title: raid.title,
      startsAt: raid.starts_at,
      guildName,
      tz: raid.guild_timezone,
      night,
      report,
      reportUrl: reportUrlOf(location),
      locale,
    });
  }
</script>

{#snippet meta()}
  <Tag tone={status.tone} testid="raid-status-chip">{status.text}</Tag>
  {#if kind !== 'raid'}<span data-testid="raid-kind">{eyebrowOf(kind)}</span>{/if}
  {#if raid.title !== raid.zone_name}<span data-testid="raid-title-zone">{zoneLineOf(raid)}</span>{/if}
  <span data-testid="raid-when">{when.date}{SEP}{when.start} to {when.end}</span>
{/snippet}

{#snippet answer()}
  {#if report !== null}
    <div class="grid grid-cols-4 gap-x-6 gap-y-5 max-[600px]:grid-cols-2" data-testid="raid-report-figures">
      {#each headlineFigures(report) as figure (figure.key)}
        <Figure
          value={figure.value}
          label={figure.label}
          tone={figure.tone}
          testid={`raid-report-figure-${figure.key}`}
        />
      {/each}
    </div>
  {/if}
{/snippet}

{#snippet header()}
  <PageHeader title={raid.title} tone="strong" {meta} {answer} testid="raid-head">
    {#snippet crest()}<FactionCrest
        faction={c.guild?.faction ?? null}
        src={c.guild?.mark ?? null}
      />{/snippet}
    {#snippet crumb()}<a href={`${guildPath}#raids`} data-testid="raid-back-link"
        >{raidCopy.backLink(guildName)}</a
      >{/snippet}
    {#snippet tabs()}
      <GuildTabs active="raids" role={officer ? 'officer' : 'member'} hrefBase={guildPath} />
    {/snippet}
  </PageHeader>
{/snippet}

<!-- The log the night was read from: in the Attendance section once the report draws, alone under
     the head while it cannot (a matched log that turned private). -->
{#snippet matchLine()}
  <RaidDoneBand
    {raid}
    {now}
    {guildPath}
    {viewerTimezone}
    {guildTimezone}
    {locale}
    {officer}
    nightLine={false}
    readOnly={!c.officerTools}
    {gateId}
    loadCandidates={c.loadCandidates}
    onMatch={c.matchLog}
    onClear={c.clearMatch}
  />
{/snippet}

<!-- One block: the head and the Page sit 20 px apart (the Page's own top), never the guild view's
     gap between them. The live line is outside the main column, where it would take a gap. -->
<div>
  <p class="sr-only" role="status" data-testid="raid-live-region">
    {c.logMatched ? attendanceCopy.liveRegion : ''}
  </p>
  <Page layout="split" panels {header} testid="raid-page">
    {#if report === null || night === null}{@render matchLine()}{/if}
    {#if officer && !c.officerTools}
      <GuildPlanGate tab="raids" facts={gateFacts} id={gateId} testid="guild-raids-plan-gate-report" />
    {/if}
    {#if c.pollPaused}<p class="text-muted text-[13px]" data-testid="raid-poll-paused">
        {raidCopy.poll}
      </p>{/if}
    {#if report !== null && night !== null}
      {#if showYourNight}<ReportYourNight {raid} currentKey={c.currentKey} {locale} />{/if}
      <ReportPulls {report} />
      {#if officer}<ReportDeaths {report} />{/if}
      <ReportOutput
        parses={report.parses}
        {officer}
        ownKeys={(raid.attendance ?? []).filter((row) => row.own).map((row) => row.character_key)}
      />
      <ReportConsumables byPull={report.consumables.by_pull} />
    {/if}

    {#snippet rail()}
      {#if report !== null && night !== null}
        <ReportAttendance
          log={matchLine}
          {night}
          rows={raid.attendance ?? []}
          {officer}
          readOnly={!c.officerTools}
          {gateId}
          fixing={c.night.fixing}
          lines={c.night.fixLines}
          onFix={(key, fix) => void c.night.fix(key, fix)}
        />
        <ReportLoot groups={report.loot} {officer} lootHref={`${guildPath}#loot`} />
        {#if officer}
          <ReportDiscord
            title={raid.title}
            {report}
            {night}
            connected={raid.discord.connected}
            {settingsHref}
            officerTools={c.officerTools}
            posting={c.night.posting}
            postLine={c.night.postLine}
            {viewerTimezone}
            {locale}
            {copyText}
            onPost={(repost) => void c.night.post(repost)}
          />
        {/if}
      {/if}
      {#if raid.notes_html.trim() !== ''}
        <Section title={raidCopy.subheads.notes} anchor="raid-notes" panel testid="raid-notes">
          <div
            class="prose-forever [&_a]:text-gold [&_li]:marker:text-muted text-[14px] leading-[1.6] [&_ol]:list-decimal [&_ul]:list-disc"
            data-testid="raid-notes-prose"
          >
            <!-- The API renders and sanitises the notes (paragraphs, bold, italic, lists, links). -->
            <!-- eslint-disable-next-line svelte/no-at-html-tags -->
            {@html raid.notes_html}
          </div>
        </Section>
      {/if}
    {/snippet}
  </Page>
</div>
