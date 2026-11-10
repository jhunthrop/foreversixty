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
  import { nightTops, outputFigure, topOutputLabel } from '../../lib/guild/output-figure';
  import { TRASH_LABEL, bossIdsOf, nightFights, raidTallyParts } from '../../lib/guild/raid-view';
  import Page from '../ui/Page.svelte';
  import Section from '../ui/Section.svelte';
  import Tag from '../ui/Tag.svelte';
  import { LINK, ROW_ITEM } from '../../lib/ui/primitives';
  import FactLine from './FactLine.svelte';
  import { attendanceCopy } from '../../lib/guild/attendance-copy';
  import { SEP, raidCopy } from '../../lib/guild/raid-copy';
  import { nightFromHash } from '../../lib/guild/raids-hash';
  import { guildRaidHref } from '../../lib/characters';
  import { RAID_CAP, lastPlannedRaid, openRaidCount, type RaidSubmitResult } from '../../lib/guild/raid-form';
  import { createRaid, type RaidInput } from '../../lib/guild/raids-api';
  import { planRaid } from '../../lib/guild/raids-tab-controller';
  import { button } from '../../lib/ui/button';
  import type { GateFacts } from '../../lib/guild/plan-gate';
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
    viewerTimezone = new Intl.DateTimeFormat().resolvedOptions().timeZone,
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
  <Section title={raidCopy.subheads.raidNights} panel testid="guild-raid-nights">
    {#if rows.length === 0}
      <p class="text-muted text-[14px] leading-[22px]" data-testid="guild-raids-empty">
        {publicView ? NO_PUBLIC_RAID_NIGHTS : 'No raid nights logged yet.'}
      </p>
    {:else}
      <div class="flex flex-col" data-testid="guild-raids-rows">
        {#each rows as row (row.id)}
          {@const isOpen = expanded === row.id}
          <div class={ROW_ITEM} id={`guild-raid-night-${row.id}`}>
            <button
              class="flex min-h-[var(--row-h)] w-full flex-wrap items-center gap-x-3 gap-y-1 py-[10px] text-left"
              onclick={() => (expanded = isOpen ? null : row.id)}
              aria-expanded={isOpen}
              data-testid={`guild-raid-row-${row.id}`}
            >
              <span class="min-w-0">
                <span class="text-strong block text-[15px] leading-5 font-bold">{row.zone}</span>
                <span class="text-muted tabular block text-[13px] leading-[18px]"
                  >{raidNightDay(row.created_at)}{SEP}{minutes(row.duration_ms)}</span
                >
              </span>
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
                    class="{LINK} inline-flex min-h-11 items-center self-start text-[14px]"
                    data-testid={`guild-raid-link-${row.id}`}>{attendanceCopy.match.openRaid}</a
                  >
                {/if}
                {#if row.present.length > 0}
                  <div class="flex flex-wrap items-center gap-1">
                    {#each row.present.slice(0, 20) as p (p.character_key)}
                      <ClassCrestRing characterClass={p.class} size={22} />
                    {/each}
                    <span class="text-muted ml-2 text-[12px]"
                      >{row.present.length} of {row.raiders} present</span
                    >
                  </div>
                {/if}
                {#if nightTops(row).length > 0}
                  <!-- Healing then damage, worded and laid out as Progression's per-role line. -->
                  <p class="text-[13px]" data-testid={`guild-raid-top-${row.id}`}>
                    {#each nightTops(row) as top, index (top.metric)}
                      {#if index > 0}<span class="hidden md:inline" aria-hidden="true">&nbsp;·&nbsp;</span
                        >{/if}
                      <span class="block md:inline">
                        <span class="text-muted">{topOutputLabel(top.metric)}</span>
                        <span style={`color:${classColorVar(top.class)};font-weight:600`}>{top.name}</span>
                        <span class="numeric">{outputFigure(top.value)}</span>
                      </span>
                    {/each}
                  </p>
                {/if}
                <ul class="flex flex-col" data-testid={`guild-raid-fights-${row.id}`}>
                  {#each nightFights(row, bossIds) as fight (fight.index)}
                    <li
                      class="{ROW_ITEM} flex min-h-11 items-center gap-3 py-1.5"
                      data-testid={fight.trash ? 'guild-raid-fight-trash' : 'guild-raid-fight'}
                    >
                      {#if fight.trash}
                        <!-- Trash: named, quiet, no pull number and no badge (guild round 2, WP-12). -->
                        <span class="text-muted min-w-[160px] text-[13px]">{fight.label}</span>
                        <span class="text-muted numeric text-[13px]">{clock(fight.duration_ms)}</span>
                        <span class="text-muted ml-auto text-[13px]">{TRASH_LABEL}</span>
                      {:else if !fight.unnamed}
                        <span class="text-strong min-w-[160px] text-[14px] font-semibold">{fight.label}</span>
                        <span class="text-muted numeric text-[13px]">{clock(fight.duration_ms)}</span>
                        <span class="ml-auto">
                          <Tag tone={fight.kill ? 'success' : 'danger'}>{fight.kill ? 'Kill' : 'Wipe'}</Tag>
                        </span>
                      {:else}
                        <span class="min-w-[160px] text-[14px]">{fight.label}</span>
                        <span class="text-muted text-[13px]">no named encounter published</span>
                      {/if}
                    </li>
                  {/each}
                </ul>
              </div>
            {/if}
          </div>
        {/each}
      </div>
    {/if}
  </Section>
</Page>
