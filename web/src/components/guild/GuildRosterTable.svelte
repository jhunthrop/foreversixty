<!-- web/src/components/guild/GuildRosterTable.svelte -->
<!-- Guild control-centre spec §4.B: the Roster tab -- filter bar, five-column sort header,
     rating/attendance/parses/professions/main-alt columns, the viewer's own row pinned
     first. Member/officer only (unchanged visibility from v1). Task L8.3 adds the "In game"
     column (the addon's level/spec where they differ, gear gap in DPS, build, seen time) from
     the guild's newest addon roster snapshot, and the line above the table naming it.

     Lifecycle round 4 (the roster model): the roster is everyone the addon saw, not only the
     verified accounts. The count line names how many are on it, in game without an account,
     waiting and not seen yet; the account rows keep their table, and the names with no
     account row (and an invited member the addon has not seen) list after them with a word
     beside each. Every row carries its rank word, Member included. The waiting list says what
     Approve does, shows the rank the addon sent and offers Decline; Approve all asks first when
     an officer is among them. Controls draw from `can` (the account's capability list from
     /v1/me), never from the rank. -->
<script lang="ts">
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import GuildRosterHandoff from '../GuildRosterHandoff.svelte';
  import { classColorVar } from '../../lib/report/format';
  import type { CharacterPath } from '../../lib/characters';
  import type {
    GuildAddonSnapshot,
    GuildRosterCounts,
    GuildRosterEntry,
    GuildRosterRow,
  } from '../../lib/guild/api';
  import { guildHomeCopy, guildRosterAddonCopy, guildRosterModelCopy } from '../../lib/guild/copy';
  import {
    addonGapLabel,
    addonIdentity,
    addonLevelSpec,
    addonSeen,
    addonSnapshotLine,
    addonSpecName,
  } from '../../lib/guild/roster-addon';
  import {
    approveAllQuestion,
    can as hasCapability,
    countLine,
    derivedCounts,
    entriesWithoutRow,
    filterEntries,
    rankWord,
    stateWord,
  } from '../../lib/guild/roster-model';
  import {
    DEFAULT_ROSTER_FILTERS,
    DEFAULT_ROSTER_SORT,
    altOfMainName,
    applyRosterFilters,
    pinOwnRowFirst,
    ratingFloor,
    sortRoster,
    type RosterFilters,
    type RosterRoleFilter,
    type RosterSortKey,
  } from '../../lib/guild/roster-sort';
  import { specName } from '../../lib/guild/spec-name';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { BUSY_CLASS } from '../../lib/ui/busy';
  import EmptyState from '../ui/EmptyState.svelte';

  let {
    roster,
    pending,
    members,
    counts,
    can,
    myCharacterKey,
    rosterBusy,
    rosterActionError,
    frozen,
    removeOverridesFrozen = () => false,
    onApprove,
    onRemove,
    onDecline,
    onApproveAll,
    rowPath,
    addonSnapshot = null,
    now = new Date(),
  }: {
    roster: GuildRosterRow[];
    pending: GuildRosterRow[];
    /** The home's `members`: everyone the addon saw plus every verified character. Absent on a
     *  home from before the roster model, read as the account rows alone. */
    members?: GuildRosterEntry[];
    /** The home's `counts`; absent, they are derived from the account rows. */
    counts?: GuildRosterCounts;
    /** The account's capability words for this guild (`/v1/me`'s `can`). */
    can: readonly string[];
    myCharacterKey: string | null;
    rosterBusy: string | null;
    rosterActionError: string;
    frozen: boolean;
    /** Unchanged v1 rule: a contested claim freezes Remove for everyone except the row's
     *  own account or a site moderator. */
    removeOverridesFrozen?: (row: GuildRosterRow) => boolean;
    onApprove: (row: GuildRosterRow) => void;
    onRemove: (row: GuildRosterRow) => void;
    /** Turn a waiting character away: the API's remove, so the row leaves the list. */
    onDecline: (row: GuildRosterRow) => void;
    onApproveAll: () => void;
    rowPath: (row: GuildRosterRow) => CharacterPath;
    /** The home's `addon_snapshot` (Task L6.14); absent on a home from before it, read as null. */
    addonSnapshot?: GuildAddonSnapshot | null;
    /** The clock the snapshot's "received … ago" and each "seen … ago" read against; a test pins it. */
    now?: Date;
  } = $props();

  const canApprove = $derived(hasCapability(can, 'approve'));
  const canRemove = $derived(hasCapability(can, 'remove'));

  let filters = $state<RosterFilters>({ ...DEFAULT_ROSTER_FILTERS });
  let noAccountOnly = $state(false);
  let sortKey = $state<RosterSortKey>(DEFAULT_ROSTER_SORT);
  /** The waiting row whose Decline is asking first, and whether Approve all is. */
  let decliningKey = $state<string | null>(null);
  let confirmingApproveAll = $state(false);

  const floor = $derived(ratingFloor(roster));
  /** "Online today" on every row says nothing, so it is shown only when it tells rows apart. */
  const everyoneOnlineToday = $derived(
    roster.filter((row) => row.verified).every((row) => row.logged_recently),
  );
  // Verified rows only -- `pending`'s own explicit unverified block above already covers
  // every unverified row (spec §4.B), so this must never also lead with them (that doubled
  // every unverified raider in the roster-rows region and pushed the pinned viewer row
  // down, found on this round's own e2e run).
  const ordered = $derived(
    noAccountOnly
      ? []
      : pinOwnRowFirst(
          sortRoster(
            applyRosterFilters(
              roster.filter((row) => row.verified),
              filters,
              floor,
            ),
            sortKey,
          ),
          myCharacterKey,
        ),
  );
  const classOptions = $derived(
    [
      ...new Set(
        [...roster.map((r) => r.class), ...(members ?? []).map((m) => m.class)].filter(
          (c): c is string => c !== undefined,
        ),
      ),
    ].sort(),
  );
  const altOf = $derived(altOfMainName(roster));
  // Every roster row once (a pending row may also be in `roster`): the sender's name lookup
  // and the derived counts both read this.
  const allRows = $derived([...new Map([...roster, ...pending].map((r) => [r.character_key, r])).values()]);
  const snapshotLine = $derived(addonSnapshotLine(addonSnapshot, allRows, now));
  const countsView = $derived(counts ?? derivedCounts(allRows));
  const extras = $derived(
    filterEntries(entriesWithoutRow(members), {
      role: filters.role,
      classFilter: filters.classFilter,
      verifiedOnly: filters.verifiedOnly,
      noAccountOnly,
    }),
  );
  const shownPending = $derived(noAccountOnly ? [] : pending);

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

  /** Approve all goes straight through unless an officer is waiting, then it asks first. */
  function askApproveAll(): void {
    if (approveAllQuestion(pending) === null) onApproveAll();
    else confirmingApproveAll = true;
  }
  function doApproveAll(): void {
    confirmingApproveAll = false;
    onApproveAll();
  }
  function doDecline(row: GuildRosterRow): void {
    decliningKey = null;
    onDecline(row);
  }

  const SORT_COLUMNS: { key: RosterSortKey; label: string; hover?: string }[] = [
    { key: 'rank', label: 'Rank' },
    { key: 'ilvl', label: 'Item level' },
    { key: 'rating', label: 'Rating' },
    { key: 'attendance', label: 'Attendance' },
    { key: 'seen', label: 'Last seen' },
    { key: 'addon_gap', label: guildRosterAddonCopy.column, hover: guildRosterAddonCopy.columnHover },
  ];

  function ratingTooltip(row: GuildRosterRow): string {
    const r = row.rating;
    if (r === undefined || r === null) return '';
    return `Output ${r.output} · Survival ${r.survival} · Mechanics ${r.mechanics} · Utility ${r.utility} · Preparation ${r.preparation} · Activity ${r.activity}`;
  }

  // Tailwind utilities only (no scoped <style>): see GuildTabs.svelte's own header comment
  // -- this tab's rows only ever mount after the guild's async roster fetch resolves, so
  // a scoped <style> block's CSS is silently dropped from this build's SSR-critical-CSS
  // inlining pass.
  const FILTER_BTN = 'border-line h-8 rounded-control border px-3 text-[12px] font-semibold text-text';
  const FILTER_BTN_ON = 'border-gold text-gold bg-[rgba(229,185,85,.08)]';
  const SORT_HEADER =
    'h-9 bg-none border-none p-0 text-[11px] font-bold tracking-[0.1em] uppercase text-muted';
  const SORT_HEADER_ACTIVE = 'text-gold';
  const ROW =
    'guild-roster-row flex flex-wrap items-center gap-3 border-line-soft border-b py-2.5 lg:flex-nowrap';
  // A name with no account row: the same shape, a different class so the account rows stay
  // the only `guild-roster-row`s.
  const ENTRY_ROW =
    'guild-roster-entry flex flex-wrap items-center gap-3 border-line-soft border-b py-2.5 lg:flex-nowrap';
  const MUTED_WORD = 'text-muted text-[12px]';
</script>

{#snippet inGame(row: GuildRosterRow)}
  {@const identity = addonIdentity(row)}
  {@const none = identity === null || row.addon === undefined || row.addon === null}
  <!-- The slot is the same width on every row, so the row's height never depends on how much
       the addon knew: its text wraps inside it. Below lg it has a line of its own, and a row
       the snapshot did not carry has none there at all (no lone "—"). -->
  <div
    class={`min-w-0 basis-full text-[12px] break-words lg:shrink-0 lg:basis-[260px] ${none ? 'hidden lg:block' : ''}`}
    data-testid="guild-roster-ingame-slot"
  >
    <span class="sr-only">{guildRosterAddonCopy.srLabel}</span>
    {#if identity === null || row.addon === undefined || row.addon === null}
      <span
        class="text-muted"
        style="cursor:help"
        title={guildRosterAddonCopy.notSeenHover}
        data-testid="guild-roster-ingame-none">{guildRosterAddonCopy.notSeen}</span
      >
    {:else}
      {@const withheld =
        row.addon.level === null &&
        row.addon.spec === null &&
        row.addon.gear_gap === null &&
        row.addon.build === null}
      {@const seen = addonSeen(row.addon.seen_at, now, withheld)}
      {@const gapLabel = addonGapLabel(row.addon.gear_gap)}
      {@const build = row.addon.build !== null && row.addon.build !== '' ? row.addon.build : null}
      <!-- The spaces are mustaches because Svelte drops a bare space at the start of an {#if}. -->
      <!-- eslint-disable svelte/no-useless-mustaches -->
      <span data-testid="guild-roster-ingame"
        >{#if identity.level !== null}{addonLevelSpec(
            identity.level,
            identity.spec,
          )}{#if identity.spec !== null}
            {' '}<span class="text-muted">{guildRosterAddonCopy.addonTag}</span
            >{/if}{/if}{#if gapLabel !== null}{#if identity.level !== null}
            {' '}<span class="text-muted">{guildRosterAddonCopy.separator}</span>
          {/if}
          <span class="tabular font-mono" data-testid="guild-roster-ingame-gap">{gapLabel}</span
          >{/if}{#if build !== null}{#if identity.level !== null || gapLabel !== null}
            {' '}<span class="text-muted">{guildRosterAddonCopy.separator}</span>
          {/if}
          {guildRosterAddonCopy.build(
            build,
          )}{/if}{#if seen !== null}{#if identity.level !== null || gapLabel !== null || build !== null}
            {' '}<span class="text-muted">{guildRosterAddonCopy.separator}</span>
          {/if}
          <span class="text-muted" style="cursor:help" title={seen.hover}>{seen.label}</span>{/if}</span
      >
      <!-- eslint-enable svelte/no-useless-mustaches -->
    {/if}
  </div>
{/snippet}

{#snippet entryInGame(entry: GuildRosterEntry)}
  {@const seen = entry.seen_at === undefined ? null : addonSeen(entry.seen_at, now)}
  {@const level = entry.level === undefined ? null : guildRosterAddonCopy.level(entry.level)}
  {@const none = level === null && seen === null}
  <div
    class={`min-w-0 basis-full text-[12px] break-words lg:shrink-0 lg:basis-[260px] ${none ? 'hidden lg:block' : ''}`}
    data-testid="guild-roster-entry-ingame"
  >
    <span class="sr-only">{guildRosterAddonCopy.srLabel}</span>
    {#if none}
      <span class="text-muted" title={guildRosterModelCopy.states.not_seen}
        >{guildRosterAddonCopy.notSeen}</span
      >
    {:else}
      <span
        >{level ?? ''}{#if level !== null && seen !== null}
          &nbsp;<span class="text-muted">{guildRosterAddonCopy.separator}</span
          >&nbsp;{/if}{#if seen !== null}<span class="text-muted" style="cursor:help" title={seen.hover}
            >{seen.label}</span
          >{/if}</span
      >
    {/if}
  </div>
{/snippet}

<section class="flex flex-col gap-4" data-testid="guild-roster-tab">
  <div class="flex flex-col gap-1">
    <p class="text-[14px] font-semibold" data-testid="guild-roster-counts">{countLine(countsView)}</p>
    <p class="text-muted text-[12px]" data-testid="guild-roster-gear-note">
      {guildRosterModelCopy.gearNote}
    </p>
  </div>

  <div class="flex flex-wrap items-center gap-2" data-testid="guild-roster-filters">
    <span class="label text-muted">Filter</span>
    <button class={FILTER_BTN} onclick={cycleRole} data-testid="guild-roster-filter-role">
      Role: {filters.role === 'all' ? 'All' : filters.role}
    </button>
    <button class={FILTER_BTN} onclick={cycleClass} data-testid="guild-roster-filter-class">
      Class: {filters.classFilter === 'all' ? 'All' : filters.classFilter}
    </button>
    <button
      class={`${FILTER_BTN} ${filters.verifiedOnly ? FILTER_BTN_ON : ''}`}
      onclick={() => toggle('verifiedOnly')}
      data-testid="guild-roster-filter-verified"
    >
      Verified only
    </button>
    <button
      class={`${FILTER_BTN} ${noAccountOnly ? FILTER_BTN_ON : ''}`}
      onclick={() => (noAccountOnly = !noAccountOnly)}
      data-testid="guild-roster-filter-noaccount"
    >
      {guildRosterModelCopy.noAccountFilter}
    </button>
    <button
      class={`${FILTER_BTN} ${filters.belowFloorOnly ? FILTER_BTN_ON : ''}`}
      onclick={() => toggle('belowFloorOnly')}
      data-testid="guild-roster-filter-floor"
    >
      Below rating floor
    </button>
  </div>

  {#if canApprove && shownPending.length > 0}
    <div class="flex flex-col gap-2" data-testid="guild-roster-waiting">
      <div class="flex items-center justify-between gap-3">
        <span class="text-muted text-[13px]">{guildHomeCopy.waitingForApproval(shownPending.length)}</span>
        <button
          class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3`}
          onclick={askApproveAll}
          disabled={frozen}
          data-testid="guild-roster-approve-all"
        >
          Approve all
        </button>
      </div>
      {#if confirmingApproveAll}
        <div
          class="border-line-soft flex flex-wrap items-center gap-3 border p-3"
          data-testid="guild-roster-approve-all-confirm"
        >
          <span class="text-[13px]">{approveAllQuestion(pending)}</span>
          <button
            class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-3`}
            onclick={doApproveAll}
            disabled={frozen}
            data-testid="guild-roster-approve-all-yes"
          >
            {guildRosterModelCopy.confirm}
          </button>
          <button
            class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3`}
            onclick={() => (confirmingApproveAll = false)}
          >
            {guildRosterModelCopy.cancel}
          </button>
        </div>
      {/if}
      <p class="text-muted text-[12px]" data-testid="guild-roster-approve-explains">
        {guildRosterModelCopy.approveExplains}
      </p>
    </div>
  {/if}

  {#if shownPending.length > 0}
    <div class="flex flex-col" id="guild-roster-unverified" data-testid="guild-roster-unverified-list">
      {#each shownPending as row (row.character_key)}
        <div class={ROW}>
          <ClassCrestRing characterClass={row.class ?? ''} size={36} />
          <div class="flex min-w-[150px] flex-col">
            <span
              class="font-display text-[14px] font-bold"
              style={`color:${classColorVar(row.class ?? '')}`}
            >
              {row.name}
            </span>
            <span class="text-muted text-[12px]">{specName(row.spec ?? '')}</span>
          </div>
          <span
            class={row.rank === 'member' ? MUTED_WORD : 'pill pill-site'}
            data-testid="guild-roster-pending-rank">{rankWord(row.rank)}</span
          >
          <span class="pill pill-site" data-testid="guild-roster-state"
            >{guildRosterModelCopy.states.waiting}</span
          >
          <div class="ml-auto flex shrink-0 flex-nowrap items-center justify-end gap-3">
            {#if canApprove}
              <button
                class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3 ${rosterBusy === row.character_key ? BUSY_CLASS : ''}`}
                onclick={() => onApprove(row)}
                disabled={rosterBusy === row.character_key || frozen}
                aria-label={`Approve ${row.name}`}
                data-testid="guild-roster-approve"
              >
                Approve
              </button>
            {/if}
            {#if canRemove && row.may_remove}
              <button
                class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3 ${rosterBusy === row.character_key ? BUSY_CLASS : ''}`}
                onclick={() => (decliningKey = row.character_key)}
                disabled={rosterBusy === row.character_key || (frozen && !removeOverridesFrozen(row))}
                aria-label={`Decline ${row.name}`}
                data-testid="guild-roster-decline"
              >
                {guildRosterModelCopy.decline}
              </button>
            {/if}
          </div>
        </div>
        {#if decliningKey === row.character_key}
          <div
            class="border-line-soft flex flex-wrap items-center gap-3 border-b py-2.5"
            data-testid="guild-roster-decline-confirm"
          >
            <span class="text-[13px]">{guildRosterModelCopy.declineQuestion(row.name)}</span>
            <button
              class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-3`}
              onclick={() => doDecline(row)}
              data-testid="guild-roster-decline-yes"
            >
              {guildRosterModelCopy.decline}
            </button>
            <button
              class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3`}
              onclick={() => (decliningKey = null)}
            >
              {guildRosterModelCopy.cancel}
            </button>
          </div>
        {/if}
      {/each}
    </div>
  {/if}

  <p class="text-muted text-[13px]" data-testid="guild-roster-addon-snapshot">{snapshotLine}</p>

  <div class="flex flex-wrap gap-6 pl-1 lg:gap-3 lg:pr-0" data-testid="guild-roster-sort-header">
    {#each SORT_COLUMNS as col (col.key)}
      <button
        class={`${SORT_HEADER} ${col.key === 'addon_gap' ? 'min-w-0 text-left lg:ml-auto lg:w-[260px] lg:shrink-0' : ''} ${sortKey === col.key ? SORT_HEADER_ACTIVE : ''}`}
        title={col.hover}
        onclick={() => (sortKey = col.key)}
        data-testid={`guild-roster-sort-${col.key}`}
      >
        {col.label}{sortKey === col.key ? ' ▾' : ''}
      </button>
    {/each}
    <!-- Where each row's actions group sits, so the In game heading stays over its slot. -->
    <span class="hidden lg:block lg:w-[190px] lg:shrink-0" aria-hidden="true"></span>
  </div>

  {#if ordered.length === 0 && extras.length === 0}
    <EmptyState message="No raiders match these filters." testid="guild-roster-empty" />
  {/if}
  {#if ordered.length > 0}
    <div class="flex flex-col" data-testid="guild-roster-rows">
      {#each ordered as row (row.character_key)}
        {@const pinned = row.character_key === myCharacterKey}
        <div class={`${ROW} ${pinned ? 'bg-[rgba(229,185,85,.06)]' : ''}`}>
          <ClassCrestRing characterClass={row.class ?? ''} size={36} />
          <div class="flex min-w-[150px] flex-col">
            <span
              class="font-display text-[14px] font-bold"
              style={`color:${classColorVar(row.class ?? '')}`}
              data-testid="guild-roster-name"
            >
              {row.name}{pinned ? ' (you)' : ''}
            </span>
            <span class="text-muted text-[12px]" data-testid="guild-roster-spec"
              >{specName(row.spec ?? '')}{#if altOf.has(row.character_key)}{row.spec ? ' · ' : ''}<span
                  class="text-muted text-[11px]">alt of {altOf.get(row.character_key)}</span
                >{/if}</span
            >
          </div>
          <!-- Below lg the facts flow with the rest of the row; at lg they wrap inside their own
               group, so the In game slot and the actions keep their place at the row's end. -->
          <div
            class="contents lg:flex lg:min-w-0 lg:flex-1 lg:flex-wrap lg:items-center lg:gap-x-3 lg:gap-y-1"
          >
            <span
              class={row.rank === 'member' ? MUTED_WORD : 'pill pill-site'}
              data-testid="guild-roster-rank">{rankWord(row.rank)}</span
            >
            {#if row.logged_recently && !everyoneOnlineToday}
              <span class="text-muted text-[12px]" data-testid="guild-roster-online">
                {guildHomeCopy.loggedRecently}
              </span>
            {/if}
            {#if row.item_level !== undefined}
              <span class="tabular font-mono text-[13px]">ilvl {row.item_level}</span>
            {/if}
            {#if row.attendance !== undefined}
              <span class="text-muted tabular font-mono text-[12px]"
                >{row.attendance.present}/{row.attendance.nights} nights</span
              >
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
          </div>
          {@render inGame(row)}
          <div
            class="ml-auto flex shrink-0 flex-nowrap items-center justify-end gap-3 lg:w-[190px]"
            data-testid="guild-roster-actions"
          >
            {#if row.consent !== 'roster'}
              <GuildRosterHandoff path={rowPath(row)} stacked />
            {/if}
            {#if canRemove && row.may_remove}
              <button
                class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3 ${rosterBusy === row.character_key ? BUSY_CLASS : ''}`}
                onclick={() => onRemove(row)}
                disabled={rosterBusy === row.character_key || (frozen && !removeOverridesFrozen(row))}
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
  {#if extras.length > 0}
    <div class="flex flex-col" data-testid="guild-roster-entries">
      {#each extras as entry (`${entry.state}:${entry.name}`)}
        <div class={ENTRY_ROW} data-testid="guild-roster-entry" data-state={entry.state}>
          <ClassCrestRing characterClass={entry.class ?? ''} size={36} />
          <div class="flex min-w-[150px] flex-col">
            <span
              class="font-display text-[14px] font-bold"
              style={entry.class === undefined ? '' : `color:${classColorVar(entry.class)}`}
              data-testid="guild-roster-entry-name">{entry.name}</span
            >
            <span class="text-muted text-[12px]"
              >{entry.spec === undefined ? '' : addonSpecName(entry.spec)}</span
            >
          </div>
          <div
            class="contents lg:flex lg:min-w-0 lg:flex-1 lg:flex-wrap lg:items-center lg:gap-x-3 lg:gap-y-1"
          >
            <span class={MUTED_WORD} data-testid="guild-roster-entry-rank">{rankWord(entry.rank)}</span>
            <span
              class="text-muted text-[12px]"
              data-testid="guild-roster-entry-state"
              title={entry.state === 'no_account'
                ? 'The addon saw this character in the guild; no account on the site has claimed it.'
                : undefined}>{stateWord(entry.state)}</span
            >
          </div>
          {@render entryInGame(entry)}
          <div class="ml-auto shrink-0 lg:w-[190px]" aria-hidden="true"></div>
        </div>
      {/each}
    </div>
  {/if}
  {#if rosterActionError !== ''}
    <p class="text-[13px]" role="alert" data-testid="guild-roster-action-error">{rosterActionError}</p>
  {/if}
</section>
