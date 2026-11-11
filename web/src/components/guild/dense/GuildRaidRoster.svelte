<!-- web/src/components/guild/dense/GuildRaidRoster.svelte -->
<!-- The Roster tab, dense (design/specs/2026-11-09-guild-dense.md section 5.3): one table of every
     raider grouped by raid group with the comp target's seats, each row the figures an officer keeps
     a sheet for: rating, item level against the guild median, the gain to BiS, best parse, the last
     nights, attendance, items this tier, the readiness count and the days since the last log. Sorts
     on every figure, filters by group and name. Waiting approvals sit above it on one line. Reads the
     home it is given plus the attendance grid (officers), the loot history and the readiness board. -->
<script lang="ts">
  import { onMount, type ComponentProps } from 'svelte';
  import GuildRosterTable from '../GuildRosterTable.svelte';
  import Disclosure from '../../ui/Disclosure.svelte';
  import type { GuildRaidGroup, GuildReadinessPage, GuildRosterRow } from '../../../lib/guild/api';
  import { fetchGuildReadiness } from '../../../lib/guild/api';
  import { fetchLootNights, type LootNights } from '../../../lib/guild/loot-council-api';
  import {
    fetchRecentAttendance,
    type GuildCompTarget,
    type GuildRecentAttendance,
  } from '../../../lib/guild/overview-api';
  import { fetchGuildRaid } from '../../../lib/guild/raids-api';
  import { classColorVar, parseLabel, parseText, percentileToken } from '../../../lib/report/format';
  import type { StatColumn, StatFigure, StatSort } from '../../../lib/ui/stat-table';
  import {
    GROUP_WORD,
    RAID_GROUP_ORDER,
    filterBoard,
    groupLine,
    rosterBoard,
    seenText,
    sortBoard,
    STALE_DAYS,
    type BoardRow,
    type BoardSortKey,
  } from '../../../lib/guild/dense/roster-board';
  import ClassCrestRing from '../../character/ClassCrestRing.svelte';
  import { toCsv } from '../../../lib/report/csv';
  import Button from '../../ui/Button.svelte';
  import CopyButton from '../../ui/CopyButton.svelte';
  import Field from '../../ui/Field.svelte';
  import Page from '../../ui/Page.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';
  import Segmented from '../../ui/Segmented.svelte';
  import StatRow from '../../ui/StatRow.svelte';
  import StatTable from '../../ui/StatTable.svelte';
  import AttendanceStrip from './AttendanceStrip.svelte';

  let {
    guildId,
    roster,
    pending = [],
    nextRaidId = null,
    officer = false,
    canApprove = false,
    busy = false,
    profileHref,
    onApprove,
    onDecline,
    onApproveAll,
    table,
    now,
  }: {
    guildId: number;
    roster: readonly GuildRosterRow[];
    pending?: readonly GuildRosterRow[];
    nextRaidId?: number | null;
    officer?: boolean;
    canApprove?: boolean;
    busy?: boolean;
    profileHref: (row: GuildRosterRow) => string | null;
    onApprove: (row: GuildRosterRow) => void;
    onDecline: (row: GuildRosterRow) => void;
    onApproveAll: () => void;
    /** The account view (ranks, Remove, members with no account, mains and alts), under the board. */
    table?: ComponentProps<typeof GuildRosterTable>;
    now?: Date;
  } = $props();

  const NIGHTS = 5;
  let attendance = $state<GuildRecentAttendance | null>(null);
  let loot = $state<LootNights | null>(null);
  let readiness = $state<GuildReadinessPage | null>(null);
  let target = $state<GuildCompTarget | null>(null);
  let only = $state<GuildRaidGroup | 'all'>('all');
  let find = $state('');
  let sort = $state<StatSort>({ column: 'rating', direction: 'desc' });
  let reviewing = $state(false);

  onMount(() => {
    if (officer)
      void fetchRecentAttendance(guildId, NIGHTS, true)
        .then((page) => (attendance = page))
        .catch(() => (attendance = null));
    void fetchLootNights(guildId)
      .then((page) => (loot = page))
      .catch(() => (loot = null));
    void fetchGuildReadiness(guildId)
      .then((page) => (readiness = page))
      .catch(() => (readiness = null));
    if (nextRaidId !== null)
      void fetchGuildRaid(guildId, nextRaidId)
        .then((view) => (target = view.comp_target ?? null))
        .catch(() => (target = null));
  });

  const clock = $derived(now ?? new Date());
  const groups = $derived(
    rosterBoard({
      roster,
      attendance,
      loot,
      readiness: readiness?.rows ?? null,
      medianItemLevel: readiness?.median_item_level ?? null,
      target,
      now: clock,
    }),
  );
  const shown = $derived(
    filterBoard(groups, only, find).map((g) => ({
      ...g,
      rows: sortBoard(g.rows, sort.column as BoardSortKey),
    })),
  );
  const counts = $derived(new Map(groups.map((g) => [g.group, g.rows.length] as const)));
  const filterOptions = $derived([
    { value: 'all' as const, label: `All ${roster.length}` },
    ...RAID_GROUP_ORDER.filter((g) => counts.has(g)).map((g) => ({
      value: g,
      label: `${GROUP_WORD[g]} ${counts.get(g)}`,
    })),
  ]);
  const hasCells = $derived(attendance !== null);
  const columns: readonly StatColumn[] = [
    { id: 'rating', label: 'Rating', sortable: true },
    { id: 'ilvl', label: 'Item level', sortable: true },
    { id: 'gap', label: 'To BiS', sortable: true },
    { id: 'best', label: 'Parse %', sortable: true },
    { id: 'nights', label: 'Nights', sortable: true },
    { id: 'loot', label: 'Loot', sortable: true },
    { id: 'ready', label: 'Ready', sortable: true },
    { id: 'seen', label: 'Last log', sortable: true },
  ];
  const meta = $derived(
    [`${roster.length} raiders`, readiness === null ? '' : `median item level ${readiness.median_item_level}`]
      .filter((p) => p !== '')
      .join(' · '),
  );

  function figures(row: BoardRow): StatFigure[] {
    const dash = { value: '–', muted: true };
    return [
      row.rating === null ? dash : { value: String(row.rating) },
      row.itemLevel === null
        ? dash
        : {
            value:
              row.itemLevelDelta === null
                ? String(row.itemLevel)
                : `${row.itemLevel} ${signed(row.itemLevelDelta)}`,
          },
      row.gap === null || Math.round(row.gap) === 0
        ? dash
        : { value: `+${Math.round(row.gap)}${row.gapUnit === null ? '' : ` ${row.gapUnit.toUpperCase()}`}` },
      row.bestPercentile === null
        ? dash
        : {
            value: parseText(row.bestPercentile),
            color: percentileToken(row.bestPercentile),
            label: parseLabel(row.bestPercentile),
          },
      row.attendancePct === null ? dash : { value: `${row.attendancePct}%` },
      { value: String(row.items) },
      row.short === null
        ? dash
        : row.short === 0
          ? { value: '✓', tone: 'kill' }
          : { value: `${row.short} missing`, tone: 'wipe' },
      row.seenDays === null || row.seenDays < 2
        ? { value: seenText(row.seenDays), muted: true }
        : {
            value: seenText(row.seenDays),
            tone: (row.seenDays ?? 0) >= STALE_DAYS ? 'wipe' : undefined,
            muted: (row.seenDays ?? 0) < STALE_DAYS,
          },
    ];
  }

  function signed(n: number): string {
    return n > 0 ? `+${n}` : n === 0 ? '±0' : `−${-n}`;
  }

  function rankWord(rank: BoardRow['rank']): string {
    return rank === 'leader' ? ' · Guild master' : rank === 'officer' ? ' · Officer' : '';
  }

  function onsort(column: string): void {
    sort = { column, direction: 'desc' };
  }

  function csv(): string {
    const head = ['Raider', 'Class', 'Spec', ...columns.map((c) => c.label)];
    const lines = groups.flatMap((g) =>
      g.rows.map((r) => [r.name, r.cls, r.spec, ...figures(r).map((f) => f.value)]),
    );
    return toCsv([head, ...lines]);
  }
</script>

<Page layout="ledger" testid="guild-roster-page">
  {#if officer && pending.length > 0}
    <div data-testid="guild-raid-roster-waiting">
      <p class="text-gold flex flex-wrap items-center gap-x-3 gap-y-2 text-[15px]">
        {pending.length} waiting for approval: {pending.map((p) => p.name).join(', ')}
        <Button tone="secondary" size="sm" onclick={() => (reviewing = !reviewing)} pressed={reviewing}
          >Review</Button
        >
        {#if canApprove}
          <Button
            tone="primary"
            size="sm"
            disabled={busy}
            onclick={onApproveAll}
            testid="guild-raid-roster-approve-all">Approve all</Button
          >
        {/if}
      </p>
      {#if reviewing}
        <div class="pt-2">
          {#each pending as row (row.character_key)}
            <Row
              title={row.name}
              titleTone={{ classSlug: (row.class ?? 'warrior') as never }}
              sub={`${row.spec_name ?? row.spec ?? ''} · item level ${row.item_level ?? '–'}`}
              open
            >
              {#snippet lead()}<ClassCrestRing characterClass={row.class ?? ''} size={32} />{/snippet}
              {#snippet trail()}
                <Button tone="secondary" size="sm" disabled={busy} onclick={() => onApprove(row)}
                  >Approve</Button
                >
                <Button tone="ghost" size="sm" disabled={busy} onclick={() => onDecline(row)}>Decline</Button>
              {/snippet}
            </Row>
          {/each}
        </div>
      {/if}
    </div>
  {/if}

  <Section title="Roster" {meta} panel flush testid="guild-raid-roster">
    {#snippet head()}
      <CopyButton tone="ghost" text={csv} label="Copy CSV" testid="guild-roster-copy-csv" />
    {/snippet}
    <div class="flex flex-wrap items-center gap-3 px-5 pb-3">
      <div class="filter-scroll max-w-full min-w-0 overflow-x-auto">
        <Segmented
          options={filterOptions}
          value={only}
          label="Raid group"
          size="sm"
          onChange={(value) => (only = value)}
        />
      </div>
      <div class="w-[220px]">
        <Field
          id="guild-roster-find"
          type="search"
          label="Find a raider"
          labelHidden
          bind:value={find}
          testid="guild-roster-find"
        />
      </div>
      {#if hasCells}<span class="ml-auto hidden lg:inline"><AttendanceStrip legend /></span>{/if}
    </div>
    <StatTable
      {columns}
      nameHead="Raider"
      barHead={hasCells ? `Last ${attendance?.raids.length ?? NIGHTS}` : ''}
      bar={hasCells}
      dense
      {sort}
      {onsort}
      label="Roster"
      testid="guild-roster-board"
    >
      {#each shown as group (group.group)}
        <StatRow
          total
          bar={hasCells}
          figures={columns.map(() => ({ value: '' }))}
          testid={`guild-roster-group-${group.group}`}
        >
          {#snippet name()}{GROUP_WORD[group.group]} {group.rows.length}{/snippet}
          {#snippet sub()}{groupLine(group)}{/snippet}
        </StatRow>
        {#each group.rows as row (row.key)}
          {@const href = profileHref(roster.find((r) => r.character_key === row.key) as GuildRosterRow)}
          <StatRow
            bar={hasCells}
            figures={figures(row)}
            phone={row.itemLevel === null ? undefined : `${row.itemLevel} ilvl`}
            testid={`guild-roster-row-${row.name.toLowerCase()}`}
          >
            {#snippet icon()}<ClassCrestRing characterClass={row.cls} size={28} />{/snippet}
            {#snippet name()}
              {#if href}<a {href} style:color={classColorVar(row.cls)}>{row.name}</a>{:else}<span
                  style:color={classColorVar(row.cls)}>{row.name}</span
                >{/if}
            {/snippet}
            {#snippet sub()}{row.spec}{rankWord(row.rank)}{#if row.cells.length > 0}<span class="ui-st-phone"
                  ><span class="block pt-0.5"
                    ><AttendanceStrip cells={row.cells} />{#if row.short}<span class="text-wipe ml-2"
                        >{row.short} missing</span
                      >{/if}{#if (row.seenDays ?? 0) >= 2}<span
                        class={['ml-2', (row.seenDays ?? 0) >= STALE_DAYS ? 'text-wipe' : 'text-muted']}
                        >{seenText(row.seenDays)}</span
                      >{/if}</span
                  ></span
                >{/if}{/snippet}
            {#snippet meter()}<AttendanceStrip cells={row.cells} />{/snippet}
          </StatRow>
        {/each}
      {/each}
    </StatTable>
  </Section>
  {#if table !== undefined}
    <Disclosure
      summary="Ranks, accounts and removals"
      meta={`${table.roster.length + (table.members?.length ?? 0)} on the roster`}
      testid="guild-roster-accounts"
    >
      <GuildRosterTable {...table} />
    </Disclosure>
  {/if}
</Page>

<style>
  @media (max-width: 1023.98px) {
    .filter-scroll {
      mask-image: linear-gradient(to right, black calc(100% - 24px), transparent);
      scroll-snap-type: x proximity;
    }
  }
</style>
