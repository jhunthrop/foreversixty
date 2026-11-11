<!-- web/src/components/guild/GuildRaids.svelte -->
<!-- Guild control-centre spec §4.C: every raid night, one expandable in place (no
     navigation). A night in Barrow Deeps or Hyjal Summit shows "Pull N: no named
     encounter published" rather than inventing a boss name (§12.1). Shown to every role
     (public sees public reports only, per the existing report-visibility rule server
     side) -- this component renders whatever `rows` it is given.
     Raids spec 3.1: above the nights, verified members and officers get `Planned`; officers get
     `Plan an event` (disabled with the gate sentence in officer-read-only; the cap line at 12 raids
     limits raids only, so the button stays live) and the form in place under the header row. One minute tick feeds every row's countdown. -->
<script lang="ts">
  import EmptyState from '../ui/EmptyState.svelte';
  import { minuteClock } from '../../lib/ui/minute-clock.svelte';
  import { tick } from 'svelte';
  import {
    fetchGuildProgression,
    type GuildRaidDiscord,
    type GuildRaidRow,
    type GuildRaidView,
    type GuildZoneView,
  } from '../../lib/guild/api';
  import { NO_PUBLIC_RAID_NIGHTS } from '../../lib/guild/overview-view';
  import { bossIdsOf } from '../../lib/guild/raid-view';
  import { fetchLootNights } from '../../lib/guild/loot-council-api';
  import {
    ATTENDANCE_GRID_NIGHTS,
    absentByRaid,
    awardsByReport,
    raidNightsMeta,
    type NightMisses,
  } from '../../lib/guild/raid-nights-table';
  import { fetchRecentAttendance } from '../../lib/guild/overview-api';
  import Page from '../ui/Page.svelte';
  import Section from '../ui/Section.svelte';
  import { raidCopy } from '../../lib/guild/raid-copy';
  import { nightFromHash } from '../../lib/guild/raids-hash';
  import GuildRaidNights from './GuildRaidNights.svelte';
  import { RAID_CAP, lastPlannedRaid, openRaidCount, type RaidSubmitResult } from '../../lib/guild/raid-form';
  import { createRaid, type RaidInput } from '../../lib/guild/raids-api';
  import { planRaid } from '../../lib/guild/raids-tab-controller';
  import { button } from '../../lib/ui/button';
  import type { GateFacts } from '../../lib/guild/plan-gate';
  import { viewerZone } from '../../lib/dates';
  import GuildPlanGate from './GuildPlanGate.svelte';
  import GuildPlannedRaids from './GuildPlannedRaids.svelte';
  import RaidForm from './RaidForm.svelte';

  interface Props {
    rows: GuildRaidRow[];
    guildId?: number | null;
    publicView?: boolean;
    /** Verified members and officers only (R1); absent for everyone else. */
    planned?: readonly GuildRaidView[];
    zones?: readonly GuildZoneView[];
    guildTimezone?: string | null;
    /** The guild channel's state from the raids page; falls back to a planned raid's. */
    discord?: GuildRaidDiscord | null;
    officer?: boolean;
    officerTools?: boolean;
    /** A verified member of this guild. */
    memberView?: boolean;
    settingsHref?: string;
    /** The gate sentence's facts (spec 2026-11-05 §3.2). */
    gateFacts: GateFacts;
    /** The guild page's own path; each planned row links under it. */
    guildPath?: string;
    /** Pins the clock (tests); otherwise the tab ticks once a minute. */
    now?: Date;
    viewerTimezone?: string;
    /** Re-reads the raids page after a raid is planned. */
    onPlanned?: () => Promise<void>;
    /** A create was refused with `guild_plan_required`: the tab turns officer-read-only. */
    onPlanRefused?: () => void;
    /** The tab opens with the plan form up: the Before raids open card's Plan an event. */
    openForm?: boolean;
  }

  let {
    rows,
    guildId = null,
    publicView = false,
    planned,
    zones = [],
    guildTimezone = null,
    discord = null,
    officer = false,
    officerTools = true,
    memberView = false,
    settingsHref = '',
    gateFacts,
    guildPath = '',
    now,
    viewerTimezone = viewerZone(),
    onPlanned = () => Promise.resolve(),
    onPlanRefused = () => {},
    openForm = false,
  }: Props = $props();

  const GATE_ID = 'officer-tools-gate-raids';
  const CAP_ID = 'guild-raids-cap';

  const showPlanned = $derived(planned !== undefined && (memberView || officer));
  const readOnly = $derived(officer && !officerTools);
  /** A gated officer with nothing planned gets no Planned panel: its empty line invites a plan the
   *  disabled button above cannot make (raids spec 3.1, guild round 3). Rows still planned (a lapsed
   *  plan) keep the panel, so their sign-ups stay reachable. */
  const showPlannedPanel = $derived(showPlanned && !(readOnly && (planned?.length ?? 0) === 0));
  const plannedCount = $derived(openRaidCount(planned ?? []));
  const atCap = $derived(plannedCount >= RAID_CAP);
  const planDescribedBy = $derived(readOnly ? GATE_ID : atCap ? CAP_ID : undefined);

  const rowClock = minuteClock(
    () => now,
    () => showPlanned,
  );

  let formOpen = $state(false);
  // The Before raids open card's Plan an event: the tab opens with the form up.
  $effect(() => {
    if (openForm) formOpen = true;
  });
  let planButton = $state<HTMLButtonElement | null>(null);

  async function closeForm(): Promise<void> {
    formOpen = false;
    await tick();
    planButton?.focus();
  }

  async function submitPlan(payload: RaidInput): Promise<RaidSubmitResult> {
    if (guildId === null) return { ok: false, code: 'no_guild', message: raidCopy.form.errors.failed };
    const result = await planRaid(guildId, payload, { create: createRaid, reload: onPlanned, onPlanRefused });
    if (result.ok) await closeForm();
    return result;
  }

  /** The tier's boss ids, so a night's kills split into bosses and trash packs; unknown (the
   *  progression read has not answered, or failed) until then, when the API's own count shows. */
  let bossIds = $state<ReadonlySet<number> | null>(null);
  $effect(() => {
    if (guildId === null) return;
    void fetchGuildProgression(guildId)
      .then((page) => {
        bossIds = bossIdsOf(page);
      })
      .catch(() => {
        bossIds = null;
      });
  });

  /** Items the council awarded per night, for a verified member or officer (the loot nights read);
   *  null for anyone else and until it answers, which leaves the Loot column out. */
  let lootByReport = $state<ReadonlyMap<string, number> | null>(null);
  $effect(() => {
    if (guildId === null || !(memberView || officer)) return;
    void fetchLootNights(guildId)
      .then((page) => {
        lootByReport = awardsByReport(page.nights);
      })
      .catch(() => {
        lootByReport = null;
      });
  });

  /** Officers: raiders absent per guild raid (the attendance grid over every account). */
  let absentByRaidId = $state<ReadonlyMap<number, NightMisses> | null>(null);
  $effect(() => {
    if (guildId === null || !officer) return;
    void fetchRecentAttendance(guildId, ATTENDANCE_GRID_NIGHTS, true)
      .then((grid) => {
        absentByRaidId = absentByRaid(grid);
      })
      .catch(() => {
        absentByRaidId = null;
      });
  });

  /** The open night: the table is the summary, a night's pulls open in place on a click or its
   *  deep link. */
  let expanded = $state<string | null>(null);

  /** `#raids-{report_id}` opens that night once per hash (now and on a later hash change); rows
   *  arriving after the hash are waited for, and a night the viewer collapsed stays collapsed. */
  let hash = $state(typeof window === 'undefined' ? '' : window.location.hash);
  let openedHash: string | null = null;
  $effect(() => {
    const onChange = (): void => {
      openedHash = null;
      hash = window.location.hash;
    };
    window.addEventListener('hashchange', onChange);
    return () => window.removeEventListener('hashchange', onChange);
  });
  $effect(() => {
    const id = nightFromHash(hash);
    if (id === null || id === openedHash || !rows.some((row) => row.id === id)) return;
    openedHash = id;
    expanded = id;
    void tick().then(() => document.getElementById(`guild-raid-night-${id}`)?.scrollIntoView?.());
  });
</script>

<!-- Two open Sections (web layout spec 2026-11-05 §3A.2): Planned, whose rows are the tab's one
     Plate, with the officer's Plan an event at the head's right and the form open under it; then
     the logged raid nights as open rows, each opening its pulls in place. -->
<Page layout="ledger" panels testid="guild-raids-tab">
  {#if showPlanned}
    <Section
      title={raidCopy.subheads.planned}
      panel
      meta={(planned?.length ?? 0) > 0 ? String(planned?.length) : undefined}
      metaTestid="guild-planned-count"
      testid="guild-planned"
    >
      {#snippet head()}
        <!-- Hidden while the form is open, so the form's own submit is the screen's one fill (3.1). -->
        {#if officer && !formOpen}
          <button
            type="button"
            bind:this={planButton}
            class={button(readOnly ? 'secondary' : 'primary', 'sm')}
            disabled={readOnly}
            aria-describedby={planDescribedBy}
            onclick={() => (formOpen = true)}
            data-testid="guild-raids-plan">{raidCopy.planButton}</button
          >
        {/if}
      {/snippet}
      <div class="flex flex-col gap-4">
        {#if officer && readOnly}
          <GuildPlanGate tab="raids" facts={gateFacts} />
        {:else if officer && atCap}
          <p id={CAP_ID} class="text-muted text-[13px]" data-testid="guild-raids-cap">{raidCopy.cap}</p>
        {/if}
        {#if formOpen && officer && planned !== undefined}
          <RaidForm
            mode="create"
            {zones}
            {guildTimezone}
            {viewerTimezone}
            lastRaid={lastPlannedRaid(planned)}
            {plannedCount}
            discord={discord ?? planned[0]?.discord ?? null}
            {settingsHref}
            onSubmit={submitPlan}
            onCancel={() => void closeForm()}
          />
        {/if}
        {#if showPlannedPanel && planned !== undefined}
          <GuildPlannedRaids {planned} {officer} {guildPath} now={rowClock.now} {viewerTimezone} />
        {/if}
      </div>
    </Section>
  {/if}
  <Section
    title={raidCopy.subheads.raidNights}
    panel
    flush={rows.length > 0}
    meta={rows.length > 0 ? raidNightsMeta(rows, bossIds) : undefined}
    testid="guild-raid-nights"
  >
    {#if rows.length === 0}
      <EmptyState
        message={publicView ? NO_PUBLIC_RAID_NIGHTS : 'No raid nights logged yet.'}
        testid="guild-raids-empty"
      />
    {:else}
      <GuildRaidNights
        {rows}
        {bossIds}
        {guildPath}
        {lootByReport}
        absentByRaid={absentByRaidId}
        {expanded}
        onToggle={(id) => (expanded = expanded === id ? null : id)}
      />
    {/if}
  </Section>
</Page>
