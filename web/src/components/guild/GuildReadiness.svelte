<!-- web/src/components/guild/GuildReadiness.svelte -->
<!-- Readiness by exception (batch 3, b3-guild-readiness; guild control-centre spec §4.E): the raiders
     short of something first, on the framed panel, each with what they are short of in gold and
     (officers) Nudge, which copies a message to paste to them: there is no network path from the
     site to the addon. A raid that asks for fire resistance lists each signed tank's on its own
     panel. Then who is ready, by name in class colour, and who is still leveling. The next-raid
     toggle cuts the board to the raid's In and Late sign-ups. Officer tools gate §3: with the tools
     off the plan sentence sits under the heading and every Nudge stays on screen, disabled. -->
<script lang="ts">
  import { classColorVar } from '../../lib/report/format';
  import type { GuildRaidSignup, GuildReadinessRow } from '../../lib/guild/api';
  import type { ReadinessRaid } from '../../lib/guild/attendance-contract';
  import { attendanceCopy } from '../../lib/guild/attendance-copy';
  import { CopyState } from '../../lib/ui/copy-state.svelte';

  import { fetchGuildRaid } from '../../lib/guild/raids-api';
  import Chips from '../ui/Chips.svelte';
  import EmptyState from '../ui/EmptyState.svelte';
  import Page from '../ui/Page.svelte';
  import Section from '../ui/Section.svelte';
  import { guildGateCopy } from '../../lib/guild/copy';
  import type { GateFacts } from '../../lib/guild/plan-gate';
  import {
    READINESS_EMPTY,
    NEXT_RAID_ONLY,
    levelingLabel,
    nudgeText,
    pinReadinessOwnRow,
    readinessLists,
  } from '../../lib/guild/readiness-view';
  import { viewerZone, weekdayDayMonth } from '../../lib/dates';
  import GuildPlanGate from './GuildPlanGate.svelte';
  import ReadinessBoard from './dense/ReadinessBoard.svelte';

  let {
    rows,
    officer,
    officerTools,
    raidGuildTools,
    myCharacterKey,
    gateFacts,
    guildId = null,
    raid = null,
    nextRaidOnly = false,
    onToggleNextRaid = () => {},
    guildTimezone = null,
    viewerTimezone = viewerZone(),
    median = null,
    now = new Date(),
  }: {
    rows: GuildReadinessRow[];
    officer: boolean;
    /** False when the guild's officer tools are off (gate spec §3.1, officer-read-only). */
    officerTools: boolean;
    /** False for an officer of a Guild-tier guild (pricing spec §3.2.8). */
    raidGuildTools: boolean;
    myCharacterKey: string | null;
    /** The gate sentence's facts (spec 2026-11-05 §3.2). */
    gateFacts: GateFacts;
    /** The guild, for the raid's sign-ups (who tanks) when the raid asks for fire resistance. */
    guildId?: number | null;
    /** The next planned or locked raid's filtered board header; null shows no toggle. */
    raid?: ReadinessRaid | null;
    /** The toggle's state, held by the page so it survives a re-read. */
    nextRaidOnly?: boolean;
    onToggleNextRaid?: () => void;
    guildTimezone?: string | null;
    viewerTimezone?: string;
    /** The board's median item level, the item level column's baseline. */
    median?: number | null;
    now?: Date;
  } = $props();

  const filtered = $derived(raid !== null && nextRaidOnly);
  const raidDay = $derived(
    raid === null ? '' : weekdayDayMonth(new Date(raid.starts_at), guildTimezone ?? viewerTimezone),
  );
  const lists = $derived(readinessLists(rows, filtered ? raid : null, false));
  /** A member's own row leads the list it is in; an officer's board stays worst first. */
  const short = $derived(
    officer
      ? lists.short
      : pinReadinessOwnRow(
          lists.short.map((entry) => entry.row),
          myCharacterKey,
        ).map((row) => lists.short.find((entry) => entry.row === row)!),
  );

  const officerReadOnly = $derived(officer && !(officerTools && raidGuildTools));
  /** Native disabled plus the sentence as description (§3.4); undefined leaves no attribute. */
  const gateDescribedBy = $derived(officerReadOnly ? 'officer-tools-gate-readiness' : undefined);

  /** The raid's signed tanks (In or Late), read only when the raid asks for fire resistance. */
  let tanks = $state<GuildRaidSignup[] | null>(null);
  $effect(() => {
    const raidId = raid?.id;
    if (!filtered || raid?.fire_resistance !== true || guildId === null || raidId === undefined) {
      tanks = null;
      return;
    }
    void fetchGuildRaid(guildId, raidId)
      .then((view) => {
        tanks = view.signups.filter((s) => s.role === 'tank' && (s.state === 'in' || s.state === 'late'));
      })
      .catch(() => (tanks = null));
  });

  // A refused clipboard leaves no confirmation: the label never promised more than Copied.
  const nudge = new CopyState();
  $effect(() => () => nudge.dispose());

  function onNudge(row: GuildReadinessRow): Promise<void> {
    return nudge.copy(nudgeText(row), row.character_key);
  }
</script>

<div data-testid="guild-readiness-tab">
  <Page layout="ledger" panels>
    <Section
      title={`${lists.ready.length} of ${lists.ready.length + short.length} ready${filtered ? ` for ${raidDay}` : ''}`}
      panel
      flush
      meta={officer ? undefined : guildGateCopy.readOnly}
      metaTestid="guild-readiness-read-only"
      testid="guild-readiness-short"
    >
      {#snippet head()}
        {#if raid !== null}
          <Chips
            chips={[{ id: 'next-raid', label: NEXT_RAID_ONLY, on: nextRaidOnly }]}
            label={NEXT_RAID_ONLY}
            onToggle={onToggleNextRaid}
            chipTestid="guild-readiness"
          />
        {/if}
      {/snippet}
      {#if officerReadOnly}<GuildPlanGate tab="readiness" facts={gateFacts} />{/if}
      {#if filtered && rows.length === 0}
        <EmptyState message={attendanceCopy.readiness.nobodySigned} testid="guild-readiness-nobody" />
      {:else if rows.length === 0}
        <EmptyState
          message={READINESS_EMPTY}
          action={officer ? { label: 'Open the roster', href: '#roster' } : undefined}
          testid="guild-readiness-empty"
        />
      {:else}
        <ReadinessBoard
          {short}
          ready={lists.ready}
          {officer}
          readOnly={officerReadOnly}
          describedBy={gateDescribedBy}
          copiedKey={nudge.state === 'copied' ? nudge.key : null}
          {median}
          fireKeys={tanks === null ? null : new Set(tanks.map((t) => t.character_key))}
          {now}
          onNudge={(row) => void onNudge(row)}
        />
      {/if}
    </Section>

    {#if lists.leveling.length > 0}
      <Section title="Leveling" panel testid="guild-readiness-leveling-group">
        <p class="text-[14px] leading-[23px]">
          {#each lists.leveling as row, index (row.character_key)}<span
              style={`color:${classColorVar(row.class)}`}
              data-testid={`guild-readiness-level-${row.character_key}`}>{row.name}</span
            ><span class="text-muted">&#32;{levelingLabel(row.level ?? 0)}</span
            >{#if index < lists.leveling.length - 1}<span class="text-muted">,&#32;</span>{/if}{/each}
        </p>
      </Section>
    {/if}
  </Page>
</div>
