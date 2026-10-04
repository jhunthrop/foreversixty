<!-- web/src/components/guild/GuildRosterTable.svelte -->
<!-- Guild control-centre spec §4.B: the Roster tab -- filter bar, five-column sort header,
     rating/attendance/parses/professions/main-alt columns, the viewer's own row pinned
     first. Member/officer only (unchanged visibility from v1). -->
<script lang="ts">
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import GuildRosterHandoff from '../GuildRosterHandoff.svelte';
  import { classColorVar } from '../../lib/report/format';
  import type { CharacterPath } from '../../lib/characters';
  import { characterSlug } from '../../lib/characters';
  import type { GuildRosterRow } from '../../lib/guild/api';
  import {
    DEFAULT_ROSTER_FILTERS,
    orderRosterForTab,
    ratingFloor,
    type RosterFilters,
    type RosterRoleFilter,
    type RosterSortKey,
  } from '../../lib/guild/roster-sort';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { BUSY_CLASS } from '../../lib/ui/busy';
  import EmptyState from '../ui/EmptyState.svelte';

  let {
    roster,
    pending,
    officer,
    myCharacterKey,
    rosterBusy,
    rosterActionError,
    frozen,
    onApprove,
    onRemove,
    onApproveAll,
    rowPath,
  }: {
    roster: GuildRosterRow[];
    pending: GuildRosterRow[];
    officer: boolean;
    myCharacterKey: string | null;
    rosterBusy: string | null;
    rosterActionError: string;
    frozen: boolean;
    onApprove: (row: GuildRosterRow) => void;
    onRemove: (row: GuildRosterRow) => void;
    onApproveAll: () => void;
    rowPath: (row: GuildRosterRow) => CharacterPath;
  } = $props();

  let filters = $state<RosterFilters>({ ...DEFAULT_ROSTER_FILTERS });
  let sortKey = $state<RosterSortKey>('ilvl');

  const floor = $derived(ratingFloor(roster));
  const ordered = $derived(
    orderRosterForTab(roster, { filters, floor, sortKey, myCharacterKey }),
  );
  const classOptions = $derived([...new Set(roster.map((r) => r.class).filter((c): c is string => c !== undefined))].sort());

  function toggle(key: 'verifiedOnly' | 'belowFloorOnly'): void {
    filters = { ...filters, [key]: !filters[key] };
  }
  function cycleRole(): void {
    const order: RosterRoleFilter[] = ['all', 'tank', 'healer', 'dps'];
    const next = order[(order.indexOf(filters.role) + 1) % order.length];
    filters = { ...filters, role: next };
  }
  function cycleClass(): void {
    const order = ['all', ...classOptions];
    const next = order[(order.indexOf(filters.classFilter) + 1) % order.length];
    filters = { ...filters, classFilter: next };
  }

  const SORT_COLUMNS: { key: RosterSortKey; label: string }[] = [
    { key: 'rank', label: 'Rank' },
    { key: 'ilvl', label: 'Item level' },
    { key: 'rating', label: 'Rating' },
    { key: 'attendance', label: 'Attendance' },
    { key: 'seen', label: 'Last seen' },
  ];

  function ratingTooltip(row: GuildRosterRow): string {
    const r = row.rating;
    if (r === undefined || r === null) return '';
    return `Output ${r.output} · Survival ${r.survival} · Mechanics ${r.mechanics} · Utility ${r.utility} · Preparation ${r.preparation} · Activity ${r.activity}`;
  }
</script>

<section class="flex flex-col gap-4" data-testid="guild-roster-tab">
  <div class="flex flex-wrap items-center gap-2" data-testid="guild-roster-filters">
    <span class="label text-muted">Filter</span>
    <button class="guild-filter-btn" onclick={cycleRole} data-testid="guild-roster-filter-role">
      Role: {filters.role === 'all' ? 'All' : filters.role}
    </button>
    <button class="guild-filter-btn" onclick={cycleClass} data-testid="guild-roster-filter-class">
      Class: {filters.classFilter === 'all' ? 'All' : filters.classFilter}
    </button>
    <button
      class="guild-filter-btn"
      class:is-on={filters.verifiedOnly}
      onclick={() => toggle('verifiedOnly')}
      data-testid="guild-roster-filter-verified"
    >
      Verified only
    </button>
    <button
      class="guild-filter-btn"
      class:is-on={filters.belowFloorOnly}
      onclick={() => toggle('belowFloorOnly')}
      data-testid="guild-roster-filter-floor"
    >
      Below rating floor
    </button>
  </div>

  {#if officer && pending.length > 0}
    <div class="flex items-center justify-between gap-3" data-testid="guild-roster-waiting">
      <span class="text-muted text-[13px]">{pending.length} waiting for approval</span>
      <button
        class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3`}
        onclick={onApproveAll}
        disabled={frozen}
        data-testid="guild-roster-approve-all"
      >
        Approve all
      </button>
    </div>
  {/if}

  {#if pending.length > 0}
    <div class="flex flex-col" id="guild-roster-unverified" data-testid="guild-roster-unverified-list">
      {#each pending as row (row.character_key)}
        <div class="guild-roster-row">
          <ClassCrestRing characterClass={row.class ?? ''} size={36} />
          <div class="flex min-w-[150px] flex-col">
            <span class="font-display text-[14px] font-bold" style={`color:${classColorVar(row.class ?? '')}`}>
              {row.name}
            </span>
            <span class="text-muted text-[12px]">{row.spec ?? ''}</span>
          </div>
          <span class="pill pill-site">Unverified</span>
          {#if officer}
            <button
              class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text ml-auto px-3 ${rosterBusy === row.character_key ? BUSY_CLASS : ''}`}
              onclick={() => onApprove(row)}
              disabled={rosterBusy === row.character_key || frozen}
              aria-label={`Approve ${row.name}`}
              data-testid="guild-roster-approve"
            >
              Approve
            </button>
          {/if}
        </div>
      {/each}
    </div>
  {/if}

  <div class="flex flex-wrap gap-6 px-1" data-testid="guild-roster-sort-header">
    {#each SORT_COLUMNS as col (col.key)}
      <button
        class="guild-sort-header"
        class:is-active={sortKey === col.key}
        onclick={() => (sortKey = col.key)}
        data-testid={`guild-roster-sort-${col.key}`}
      >
        {col.label}{sortKey === col.key ? ' ▾' : ''}
      </button>
    {/each}
  </div>

  {#if ordered.length === 0}
    <EmptyState message="No raiders match these filters." testid="guild-roster-empty" />
  {:else}
    <div class="flex flex-col" data-testid="guild-roster-rows">
      {#each ordered as row (row.character_key)}
        {@const pinned = row.character_key === myCharacterKey}
        <div class="guild-roster-row" class:is-pinned={pinned}>
          <ClassCrestRing characterClass={row.class ?? ''} size={36} />
          <div class="flex min-w-[150px] flex-col">
            <span class="font-display text-[14px] font-bold" style={`color:${classColorVar(row.class ?? '')}`}>
              {row.name}{pinned ? ' (you)' : ''}
            </span>
            <span class="text-muted text-[12px]">{row.spec ?? ''}</span>
          </div>
          {#if row.rank === 'officer' || row.rank === 'leader'}
            <span class="pill pill-site">{row.rank === 'leader' ? 'Leader' : 'Officer'}</span>
          {/if}
          {#if row.logged_recently}
            <span class="pill" style="color:#7bff5c;background:rgba(30,255,0,.10);border-color:rgba(30,255,0,.25)">
              Logged in the last day
            </span>
          {/if}
          {#if row.item_level !== undefined}
            <span class="tabular font-mono text-[13px]">ilvl {row.item_level}</span>
          {/if}
          {#if row.attendance !== undefined}
            <span class="text-muted tabular font-mono text-[12px]">{row.attendance.present}/{row.attendance.nights} nights</span>
          {/if}
          {#if row.best_parse !== undefined && row.best_parse !== null}
            <span class="text-muted tabular font-mono text-[12px]">
              {Math.round(row.best_parse.value)} best {row.best_parse.metric.toUpperCase()}
            </span>
          {/if}
          {#if row.professions !== undefined && row.professions.length > 0}
            <span class="text-muted text-[11px]">{row.professions.join(' / ')}</span>
          {/if}
          {#if row.rating !== undefined && row.rating !== null}
            <span class="tabular font-mono text-[13px]" style="cursor:help" title={ratingTooltip(row)}>
              {row.rating.overall} <span class="text-muted text-[11px]">rating</span>
            </span>
          {/if}
          <div class="ml-auto flex items-center gap-3">
            {#if row.consent !== 'roster'}
              <GuildRosterHandoff path={rowPath(row)} />
            {/if}
            {#if officer && row.rank !== 'leader'}
              <button
                class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3 ${rosterBusy === row.character_key ? BUSY_CLASS : ''}`}
                onclick={() => onRemove(row)}
                disabled={rosterBusy === row.character_key || frozen}
                aria-label={`Remove ${row.name}`}
                data-testid="guild-roster-remove"
              >
                Remove
              </button>
            {/if}
          </div>
        </div>
      {/each}
    </div>
  {/if}
  {#if rosterActionError !== ''}
    <p class="text-[13px]" role="alert" data-testid="guild-roster-action-error">{rosterActionError}</p>
  {/if}
</section>

<style>
  .guild-roster-row {
    display: flex;
    align-items: center;
    gap: 12px;
    flex-wrap: wrap;
    padding: 9px 4px;
    border-bottom: 1px solid var(--color-border-soft);
  }
  .guild-roster-row.is-pinned {
    background: rgba(229, 185, 85, 0.06);
  }
  .guild-filter-btn {
    height: 32px;
    padding: 0 12px;
    border: 1px solid var(--color-line);
    border-radius: var(--radius-control);
    background: none;
    color: var(--color-text);
    font-size: 12px;
    font-weight: 600;
  }
  .guild-filter-btn.is-on {
    border-color: var(--color-gold);
    color: var(--color-gold);
    background: rgba(229, 185, 85, 0.08);
  }
  .guild-sort-header {
    background: none;
    border: none;
    height: 36px;
    padding: 0;
    font-size: 11px;
    font-weight: 700;
    letter-spacing: 0.1em;
    text-transform: uppercase;
    color: var(--color-muted);
  }
  .guild-sort-header.is-active {
    color: var(--color-gold);
  }
</style>
