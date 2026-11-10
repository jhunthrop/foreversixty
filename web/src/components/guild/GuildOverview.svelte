<!-- web/src/components/guild/GuildOverview.svelte -->
<!-- The role-aware Overview (web layout spec 2026-11-05 §7.1), built only from the ui/ primitives:
     one Plate (the next raid, with the viewer's own sign-up row), open Sections around it.
     Member: the verdict Notice, the private checklist (D12), their numbers, the raid's comp.
     Officer: the dense officer's week (dense/OfficerWeek, design/specs/2026-11-09-guild-dense.md 5.1). Everyone: the last raid, with
     progression as its one line. A module whose data has no source yet is not drawn (7.4). -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type {
    GuildHomeSummary,
    GuildRaidView,
    GuildRaidsPage,
    GuildRosterRow,
    NextRaidView,
  } from '../../lib/guild/api';
  import { fetchGuildProgression, fetchGuildRaids } from '../../lib/guild/api';
  import { guildRaidHref } from '../../lib/characters';
  import { minuteClock } from '../../lib/ui/minute-clock.svelte';
  import {
    fetchMyChecklist,
    fetchOfficerActions,
    type GuildChecklist,
    type GuildOfficerActions,
  } from '../../lib/guild/overview-api';
  import { checklistRows } from '../../lib/guild/overview-role-view';
  import {
    downEncounterNames,
    lastRaidProgressLine,
    raidNightDay,
    type TierCounts,
  } from '../../lib/guild/overview-view';
  import { fetchGuildRaid } from '../../lib/guild/raids-api';
  import { bossIdsOf, raidTallyParts } from '../../lib/guild/raid-view';
  import { counted } from '../../lib/plural';
  import Notice from '../ui/Notice.svelte';
  import Page from '../ui/Page.svelte';
  import LastRaid from './overview/LastRaid.svelte';
  import MemberChecklist from './overview/MemberChecklist.svelte';
  import NextRaidPlate from './overview/NextRaidPlate.svelte';
  import RecruitLine from './overview/RecruitLine.svelte';
  import MemberWeek from './dense/MemberWeek.svelte';
  import CompGaps from './dense/CompGaps.svelte';
  import OfficerWeek from './dense/OfficerWeek.svelte';
  import type { GuildTabId } from './GuildTabs.svelte';

  /** A report row the Last raid line can summarize (home.reports or the public page's). */
  interface OverviewReport {
    id: string;
    title: string;
    zone: string;
    created_at: string;
    fight_count?: number;
    kill_count?: number;
  }

  let {
    guildId,
    role,
    officer = false,
    reports,
    summary,
    publicCounts = { killed: 0, pulls: 0 },
    showCounts = true,
    nextRaid = null,
    guildPath,
    myRow = null,
    roster = [],
    onSelectTab,
    lead,
    you,
    now,
  }: {
    guildId: number | null;
    role: 'public' | 'member' | 'officer' | 'moderator';
    /** The viewer may run the guild (officer tools), whatever their rank word. */
    officer?: boolean;
    reports: readonly OverviewReport[];
    summary: GuildHomeSummary | undefined;
    /** The public read's tier counts: the Last raid line for a page with no home or progression read. */
    publicCounts?: TierCounts;
    /** False for a visitor when the only report is guild-only. */
    showCounts?: boolean;
    nextRaid?: NextRaidView | null;
    guildPath: string;
    /** The viewer's own roster row (their raid character), for Your numbers. */
    myRow?: GuildRosterRow | null;
    /** The guild's roster (the home's), for the officer's week. */
    roster?: readonly GuildRosterRow[];
    onSelectTab: (tab: GuildTabId) => void;
    /** The first block on the tab, when the page has one to lead with (Before raids open, which then
     *  carries the next step: no "No raid planned" line and no numbers before the first raid). */
    lead?: Snippet;
    /** The viewer's own sign-up row for the next raid (the band line). */
    you?: Snippet;
    /** A fixed clock (tests, captures); live otherwise. */
    now?: Date;
  } = $props();

  const viewerTz = Intl.DateTimeFormat().resolvedOptions().timeZone;
  const clock = minuteClock(() => now);
  const member = $derived(role !== 'public');

  let progression = $state<{ named: number; names: string[] } | null>(null);
  let raidsPage = $state<GuildRaidsPage | null>(null);
  let bossIds = $state<ReadonlySet<number> | null>(null);
  let raid = $state<GuildRaidView | null>(null);
  let checklist = $state<GuildChecklist | null>(null);
  let actions = $state<GuildOfficerActions | null>(null);

  $effect(() => {
    if (guildId === null) return;
    void fetchGuildProgression(guildId)
      .then((page) => {
        progression = { named: page.tier.named_encounters, names: downEncounterNames(page.encounters) };
        bossIds = bossIdsOf(page);
      })
      .catch(() => (progression = null));
    void fetchGuildRaids(guildId)
      .then((page) => (raidsPage = page))
      .catch(() => (raidsPage = null));
  });

  function reloadRaid(): void {
    if (guildId === null || nextRaid === null) return;
    void fetchGuildRaid(guildId, nextRaid.id)
      .then((view) => (raid = view))
      .catch(() => (raid = null));
  }

  $effect(() => {
    if (!member) return;
    reloadRaid();
  });

  $effect(() => {
    if (guildId === null || !member || officer) return;
    void fetchMyChecklist(guildId, nextRaid?.id ?? null)
      .then((list) => (checklist = list))
      .catch(() => (checklist = null));
  });

  $effect(() => {
    if (guildId === null || !officer) return;
    void fetchOfficerActions(guildId)
      .then((page) => (actions = page))
      .catch(() => (actions = null));
  });

  const rows = $derived(checklist === null ? [] : checklistRows(checklist));
  const progressionLine = $derived(lastRaidProgressLine(progression, summary?.tier, publicCounts));
  const last = $derived(reports[0] ?? null);
  const lastFacts = $derived.by((): string[] => {
    if (last === null) return [];
    const night = raidsPage?.rows.find((row) => row.id === last.id);
    if (night !== undefined) return raidTallyParts(night, bossIds);
    if (last.kill_count === undefined || last.fight_count === undefined) return [];
    return [
      counted(last.kill_count, 'kill'),
      counted(Math.max(0, last.fight_count - last.kill_count), 'wipe'),
    ];
  });
  const planHref = $derived(nextRaid === null ? '#raids' : guildRaidHref(guildPath, nextRaid.id));
</script>

{#snippet lastRaid(withProgression = true)}
  <LastRaid
    report={last === null ? null : { ...last, day: raidNightDay(last.created_at) }}
    facts={lastFacts}
    progression={showCounts && withProgression ? progressionLine : ''}
  />
{/snippet}

{#snippet plate()}
  {#if nextRaid !== null}
    <NextRaidPlate
      next={nextRaid}
      {raid}
      {officer}
      {planHref}
      {viewerTz}
      now={clock.now}
      you={officer ? undefined : you}
    />
  {:else if member && lead === undefined}
    <Notice
      text="No raid planned yet."
      action={officer ? { label: 'Plan one', href: '#raids' } : undefined}
      tone="neutral"
      testid="guild-overview-no-raid"
    />
  {/if}
{/snippet}

<div data-testid="guild-overview-tab">
  {#if role === 'public'}
    <Page layout="split" panels testid="guild-overview-page">
      {#if guildId !== null && guildPath !== ''}<RecruitLine {guildId} {guildPath} />{/if}
      {@render lastRaid()}
    </Page>
  {:else if officer}
    {#if guildId !== null}
      <OfficerWeek
        {lead}
        {guildId}
        next={nextRaid}
        {raid}
        {roster}
        {raidsPage}
        {actions}
        {guildPath}
        {planHref}
        {viewerTz}
        now={clock.now}
        {onSelectTab}
        onRaidChanged={reloadRaid}
      />
    {/if}
  {:else}
    <Page layout="split" panels testid="guild-overview-page">
      {@render lead?.()}
      {#if guildId !== null && nextRaid !== null && lead === undefined}
        <MemberWeek
          part="week"
          {guildId}
          next={nextRaid}
          {raid}
          {myRow}
          {roster}
          {raidsPage}
          {viewerTz}
          now={clock.now}
          {you}
        />
      {:else}
        {@render plate()}
      {/if}
      {#if checklist !== null && rows.length > 0}
        <div id="guild-overview-checklist">
          <MemberChecklist {rows} />
        </div>
      {/if}
      {#if guildId !== null && lead === undefined}
        <MemberWeek
          part="nights"
          {guildId}
          next={nextRaid}
          {raid}
          {myRow}
          {roster}
          {raidsPage}
          {viewerTz}
          now={clock.now}
        />
      {/if}
      {#snippet rail()}
        {#if raid?.comp_target}<CompGaps {raid} />{/if}
        {#if guildId !== null}
          <MemberWeek
            part="loot"
            {guildId}
            next={nextRaid}
            {raid}
            {myRow}
            {roster}
            {raidsPage}
            {viewerTz}
            now={clock.now}
          />
        {/if}
        {@render lastRaid(false)}
      {/snippet}
    </Page>
  {/if}
</div>
