<!-- web/src/components/guild/GuildReadiness.svelte -->
<!-- Readiness by exception (batch 3, b3-guild-readiness; guild control-centre spec §4.E): the raiders
     short of something first, on the framed panel, each with what they are short of in gold and
     (officers) Nudge, which copies a message to paste to them: there is no network path from the
     site to the addon. A raid that asks for fire resistance lists each signed tank's on its own
     panel. Then who is ready, by name in class colour, and who is still leveling. The next-raid
     toggle cuts the board to the raid's In and Late sign-ups. Officer tools gate §3: with the tools
     off the plan sentence sits under the heading and every Nudge stays on screen, disabled. -->
<script lang="ts">
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import { classColorVar } from '../../lib/report/format';
  import { classTone } from '../../lib/ui/primitives';
  import type { GuildRaidSignup, GuildReadinessRow } from '../../lib/guild/api';
  import type { ReadinessRaid } from '../../lib/guild/attendance-contract';
  import { attendanceCopy, nextRaidToggle } from '../../lib/guild/attendance-copy';
  import { FILTER_BTN, FILTER_BTN_ON } from '../../lib/guild/filter-styles';
  import { dateLabel } from '../../lib/guild/raid-time';
  import { fetchGuildRaid } from '../../lib/guild/raids-api';
  import EmptyState from '../ui/EmptyState.svelte';
  import Page from '../ui/Page.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import { button } from '../../lib/ui/button';
  import { guildGateCopy } from '../../lib/guild/copy';
  import type { GateFacts } from '../../lib/guild/plan-gate';
  import { counted } from '../../lib/plural';
  import {
    READINESS_EMPTY,
    READY_MEANS,
    levelingLabel,
    nudgeText,
    pinReadinessOwnRow,
    readinessLists,
    tankFireLine,
  } from '../../lib/guild/readiness-view';
  import GuildPlanGate from './GuildPlanGate.svelte';

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
    viewerTimezone = new Intl.DateTimeFormat().resolvedOptions().timeZone,
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
  } = $props();

  const filtered = $derived(raid !== null && nextRaidOnly);
  const raidDay = $derived(
    raid === null ? '' : dateLabel(new Date(raid.starts_at), guildTimezone ?? viewerTimezone),
  );
  const lists = $derived(readinessLists(rows, filtered ? raid : null));
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
  const fireOf = (key: string): number | null =>
    rows.find((row) => row.character_key === key)?.resistances?.fire ?? null;

  let copiedKey = $state<string | null>(null);

  async function onNudge(row: GuildReadinessRow): Promise<void> {
    try {
      await navigator.clipboard.writeText(nudgeText(row));
      copiedKey = row.character_key;
      setTimeout(() => {
        if (copiedKey === row.character_key) copiedKey = null;
      }, 2000);
    } catch {
      // Clipboard access can be refused (no permission, insecure context): the label never
      // promised more than "Copied", so a refusal leaves no confirmation rather than an error.
    }
  }

  const specLine = (row: GuildReadinessRow): string =>
    [
      `${row.spec} ${row.class.charAt(0).toUpperCase()}${row.class.slice(1)}`,
      row.item_level === null ? '' : `item level ${row.item_level}`,
    ]
      .filter((part) => part !== '')
      .join(' · ');
</script>

<div data-testid="guild-readiness-tab">
  <Page layout="ledger" panels>
    <Section
      title="Short of something"
      panel
      frame="framed"
      meta={officer ? (filtered ? `In or Late for ${raidDay}` : undefined) : guildGateCopy.readOnly}
      metaTestid="guild-readiness-read-only"
      testid="guild-readiness-short"
    >
      {#snippet head()}
        {#if raid !== null}
          <button
            class={`${FILTER_BTN} ${nextRaidOnly ? FILTER_BTN_ON : ''}`}
            aria-pressed={nextRaidOnly}
            onclick={onToggleNextRaid}
            data-testid="guild-readiness-next-raid"
          >
            {nextRaidToggle(raidDay)}
          </button>
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
      {:else if short.length === 0}
        <p class="text-muted text-[14px]" data-testid="guild-readiness-nobody-short">Nobody.</p>
      {:else}
        <ul data-testid="guild-readiness-rows">
          {#each short as { row, shortfalls } (row.character_key)}
            {@const pinned = row.character_key === myCharacterKey}
            <li data-testid={`guild-readiness-row-${row.character_key}`}>
              <Row
                title={`${row.name}${pinned ? ' (you)' : ''}`}
                titleTone={classTone(row.class)}
                sub={specLine(row)}
                selected={pinned}
                open
              >
                {#snippet lead()}<ClassCrestRing characterClass={row.class} size={30} />{/snippet}
                {#snippet foot()}<span class="text-gold text-[13px]" data-testid="guild-readiness-shortfalls"
                    >{shortfalls.join(' · ')}</span
                  >{/snippet}
                {#snippet trail()}
                  {#if officer}
                    <button
                      class={button('secondary', 'sm')}
                      onclick={() => void onNudge(row)}
                      disabled={officerReadOnly}
                      aria-describedby={gateDescribedBy}
                      aria-label={`Nudge ${row.name}`}
                      data-testid={`guild-readiness-nudge-${row.character_key}`}
                    >
                      {copiedKey === row.character_key ? 'Copied' : 'Nudge'}
                    </button>
                  {/if}
                {/snippet}
              </Row>
            </li>
          {/each}
        </ul>
      {/if}
    </Section>

    {#if tanks !== null}
      <Section
        title="Fire resistance on tanks"
        panel
        meta={counted(tanks.length, 'tank')}
        testid="guild-readiness-fire"
      >
        {#if tanks.length === 0}
          <p class="text-muted text-[14px]">No tank signed yet.</p>
        {:else}
          <ul>
            {#each tanks as tank (tank.character_key)}
              <li>
                <Row title={tank.name} titleTone={classTone(tank.class)} density="slim" open>
                  {#snippet trail()}<span class="text-[14px] font-semibold tabular-nums"
                      >{tankFireLine(fireOf(tank.character_key))}</span
                    >{/snippet}
                </Row>
              </li>
            {/each}
          </ul>
        {/if}
      </Section>
    {/if}

    {#if lists.ready.length > 0}
      <Section
        title={`${lists.ready.length} ready`}
        panel
        meta={READY_MEANS}
        testid="guild-readiness-ready-list"
      >
        <p class="text-[14px] leading-[23px]">
          {#each lists.ready as row, index (row.character_key)}<span
              style={`color:${classColorVar(row.class)}`}>{row.name}</span
            >{#if index < lists.ready.length - 1}<span class="text-muted">,&#32;</span>{/if}{/each}
        </p>
      </Section>
    {/if}

    {#if lists.leveling.length > 0}
      <Section title="Leveling" panel testid="guild-readiness-leveling-group">
        <p class="text-[14px] leading-[23px]">
          {#each lists.leveling as row, index (row.character_key)}<span
              style={`color:${classColorVar(row.class)}`}>{row.name}</span
            ><span class="text-muted">&#32;{levelingLabel(row.level ?? 0)}</span
            >{#if index < lists.leveling.length - 1}<span class="text-muted">,&#32;</span>{/if}{/each}
        </p>
      </Section>
    {/if}
  </Page>
</div>
