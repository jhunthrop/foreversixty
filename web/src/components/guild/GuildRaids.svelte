<!-- web/src/components/guild/GuildRaids.svelte -->
<!-- Guild control-centre spec §4.C: every raid night, one expandable in place (no
     navigation). A night in Barrow Deeps or Hyjal Summit shows "Pull N — no named
     encounter published" rather than inventing a boss name (§12.1). Shown to every role
     (public sees public reports only, per the existing report-visibility rule server
     side) -- this component renders whatever `rows` it is given.
     Raids spec 3.1: above the nights, verified members and officers get `Planned`; officers get
     `Plan a raid` (disabled with the gate sentence in officer-read-only, and with the cap line at
     12) and the form in place under the header row. One minute tick feeds every row's countdown. -->
<script lang="ts">
  import { minuteClock } from '../../lib/guild/minute-clock.svelte';
  import { tick } from 'svelte';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import { classColorVar } from '../../lib/report/format';
  import {
    fetchGuildProgression,
    type GuildRaidDiscord,
    type GuildRaidRow,
    type GuildRaidView,
    type GuildZoneView,
  } from '../../lib/guild/api';
  import { NO_PUBLIC_RAID_NIGHTS, raidNightDay } from '../../lib/guild/overview-view';
  import { bossIdsOf, raidTallyParts } from '../../lib/guild/raid-view';
  import EmptyState from '../ui/EmptyState.svelte';
  import FactLine from './FactLine.svelte';
  import { attendanceCopy } from '../../lib/guild/attendance-copy';
  import { raidCopy } from '../../lib/guild/raid-copy';
  import { nightFromHash } from '../../lib/guild/raids-hash';
  import { guildRaidHref } from '../../lib/characters';
  import {
    RAID_CAP,
    lastPlannedRaid,
    openRaidCount,
    splitAtLink,
    type RaidSubmitResult,
  } from '../../lib/guild/raid-form';
  import { createRaid, type RaidInput } from '../../lib/guild/raids-api';
  import { planRaid } from '../../lib/guild/raids-tab-controller';
  import { PRIMARY_BUTTON_FIXED } from '../../lib/planner/styles';
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
    /** The guild page's own path; each planned row links under it. */
    guildPath?: string;
    /** Pins the clock (tests); otherwise the tab ticks once a minute. */
    now?: Date;
    viewerTimezone?: string;
    /** Re-reads the raids page after a raid is planned. */
    onPlanned?: () => Promise<void>;
    /** A create was refused with `guild_plan_required`: the tab turns officer-read-only. */
    onPlanRefused?: () => void;
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
    guildPath = '',
    now,
    viewerTimezone = new Intl.DateTimeFormat().resolvedOptions().timeZone,
    onPlanned = () => Promise.resolve(),
    onPlanRefused = () => {},
  }: Props = $props();

  const GATE_ID = 'officer-tools-gate-raids';
  const CAP_ID = 'guild-raids-cap';
  const gateLead = splitAtLink(raidCopy.gate, raidCopy.gateLink).before;

  const showPlanned = $derived(planned !== undefined && (memberView || officer));
  const readOnly = $derived(officer && !officerTools);
  const plannedCount = $derived(openRaidCount(planned ?? []));
  const atCap = $derived(plannedCount >= RAID_CAP);
  const planDescribedBy = $derived(readOnly ? GATE_ID : atCap ? CAP_ID : undefined);

  const rowClock = minuteClock(
    () => now,
    () => showPlanned,
  );

  let formOpen = $state(false);
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

  // The first night opens from the first paint; rows that arrive later open theirs below.
  // svelte-ignore state_referenced_locally
  let expanded = $state<string | null>(rows[0]?.id ?? null);
  // svelte-ignore state_referenced_locally
  let defaulted = $state(rows.length > 0);

  /** `rows` arrives after this component's own mount (Guild.svelte only renders it once
   *  the raids fetch resolves, but a future caller could pass an empty array first) -- a
   *  plain `$state` initializer only captures its value once, so the first night's own
   *  default-open state is set here instead, exactly once. */
  $effect(() => {
    if (defaulted || rows.length === 0) return;
    expanded = rows[0].id;
    defaulted = true;
  });

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
    defaulted = true;
    void tick().then(() => document.getElementById(`guild-raid-night-${id}`)?.scrollIntoView?.());
  });

  function minutes(ms: number): string {
    const total = Math.round(ms / 60000);
    return `${total} min`;
  }
  function clock(ms: number): string {
    const totalSeconds = Math.round(ms / 1000);
    const m = Math.floor(totalSeconds / 60);
    const s = totalSeconds % 60;
    return `${m}:${String(s).padStart(2, '0')}`;
  }
</script>

<section class="flex flex-col gap-2" data-testid="guild-raids-tab">
  <div class="flex items-center justify-between gap-3" data-testid="guild-raids-header">
    <h2 class="section-title">Raids</h2>
    <!-- Hidden while the form is open, so the form's own submit is the screen's one fill (3.1). -->
    {#if officer && showPlanned && !formOpen}
      <button
        type="button"
        bind:this={planButton}
        class="{PRIMARY_BUTTON_FIXED} focus-visible:outline-gold px-4 focus-visible:outline-2 focus-visible:outline-offset-2 active:bg-[#cfa64a] disabled:cursor-not-allowed disabled:opacity-50"
        disabled={readOnly || atCap}
        aria-describedby={planDescribedBy}
        onclick={() => (formOpen = true)}
        data-testid="guild-raids-plan">{raidCopy.planButton}</button
      >
    {/if}
  </div>
  {#if officer && showPlanned && readOnly}
    <GuildPlanGate tab="raids" {settingsHref} sentence={gateLead} />
  {:else if officer && showPlanned && atCap}
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
  {#if showPlanned && planned !== undefined}
    <GuildPlannedRaids {planned} {officer} {guildPath} now={rowClock.now} {viewerTimezone} />
    <h3 class="label text-muted mt-4" data-testid="guild-raid-nights-heading">
      {raidCopy.subheads.raidNights}
    </h3>
  {/if}
  {#if rows.length === 0}
    <EmptyState
      message={publicView
        ? NO_PUBLIC_RAID_NIGHTS
        : 'No raid nights logged yet. Upload a raid log to get started.'}
      testid="guild-raids-empty"
    />
  {:else}
    <div class="flex flex-col" data-testid="guild-raids-rows">
      {#each rows as row (row.id)}
        {@const isOpen = expanded === row.id}
        <div class="border-line-soft border-b" id={`guild-raid-night-${row.id}`}>
          <button
            class="flex w-full flex-wrap items-center gap-3 py-2 text-left"
            onclick={() => (expanded = isOpen ? null : row.id)}
            aria-expanded={isOpen}
            data-testid={`guild-raid-row-${row.id}`}
          >
            <span class="min-w-[140px] text-[14px] font-semibold">{row.zone}</span>
            <span class="text-muted tabular text-[13px]">{raidNightDay(row.created_at)}</span>
            <span class="text-muted numeric text-[13px]">{minutes(row.duration_ms)}</span>
            <FactLine
              facts={raidTallyParts(row, bossIds)}
              class="text-muted tabular ml-auto text-right text-[13px]"
            />
          </button>
          {#if isOpen}
            <div class="flex flex-col gap-3 pb-4" data-testid={`guild-raid-expanded-${row.id}`}>
              {#if row.raid_id !== undefined && row.raid_id !== null}
                <a
                  href={guildRaidHref(guildPath, row.raid_id)}
                  class="text-gold inline-flex min-h-11 items-center self-start text-[13px] font-bold"
                  data-testid={`guild-raid-link-${row.id}`}>{attendanceCopy.match.openRaid}</a
                >
              {/if}
              <div class="flex flex-wrap items-center gap-1">
                {#each row.present.slice(0, 20) as p (p.character_key)}
                  <ClassCrestRing characterClass={p.class} size={22} />
                {/each}
                <span class="text-muted ml-2 text-[12px]">{row.present.length} of {row.raiders} present</span>
              </div>
              {#if row.top_parse !== null}
                <p class="text-[13px]">
                  Top parse <span style={`color:${classColorVar(row.top_parse.class)};font-weight:600`}
                    >{row.top_parse.name}</span
                  >
                  {row.top_parse.value.toFixed(1)}
                  {row.top_parse.metric.toUpperCase()}
                </p>
              {/if}
              <div class="flex flex-col" data-testid={`guild-raid-fights-${row.id}`}>
                {#each row.fights as fight (fight.index)}
                  <div class="border-line-soft flex items-center gap-3 border-b py-1.5">
                    {#if fight.name !== null}
                      <span class="min-w-[120px] text-[13px]">{fight.name} · pull {fight.index + 1}</span>
                      <span class="text-muted numeric text-[12px]">{clock(fight.duration_ms)}</span>
                      <span
                        class="pill"
                        class:pill-kill={fight.kill}
                        style={fight.kill
                          ? 'color:#7bff5c;background:rgba(30,255,0,.10);border-color:rgba(30,255,0,.25)'
                          : ''}
                      >
                        {fight.kill ? 'Kill' : 'Wipe'}
                      </span>
                    {:else}
                      <span class="min-w-[120px] text-[13px]">Pull {fight.index + 1}</span>
                      <span class="text-muted text-[12px]">no named encounter published</span>
                    {/if}
                  </div>
                {/each}
              </div>
            </div>
          {/if}
        </div>
      {/each}
    </div>
  {/if}
</section>
