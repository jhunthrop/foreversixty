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
     /v1/me), never from the rank.

     GUILD-ALL 2026-11-07, the flat system (web layout spec 2026-11-05 §3A): two open Sections,
     Waiting for approval and Roster; the roster's rows are the page's one Plate. No card, no icon
     chip, no hand-drawn rule: ui/Ruled draws the hairlines. -->
<script lang="ts">
  import CharacterName from '../ui/CharacterName.svelte';
  import Crest from '../ui/Crest.svelte';
  import Tag from '../ui/Tag.svelte';
  import SupporterChip from '../SupporterChip.svelte';
  import GuildRosterHandoff from '../GuildRosterHandoff.svelte';
  import type { CharacterPath } from '../../lib/characters';
  import type { GuildRosterCounts, GuildRosterEntry, GuildRosterRow } from '../../lib/guild/api';
  import { guildHomeCopy, guildRosterAddonCopy, guildRosterModelCopy } from '../../lib/guild/copy';
  import { focusOnMountIf } from '../../lib/guild/focus-now';
  import type { AttendanceSource } from '../../lib/guild/attendance-contract';
  import { professionName } from '../../lib/guild/professions';
  import { rosterHoldsHandoff } from '../../lib/guild/roster-links';
  import {
    RANKS_PENDING_LINE,
    entrySpecLine,
    entryFacts,
    inGameFacts,
    ROSTER_COLUMNS,
    rosterColumns,
    rosterSpecLine,
    rosterStatChips,
    type RosterColumnKey,
    type InGameFact,
  } from '../../lib/guild/roster-facts';
  import {
    approveAllQuestion,
    can as hasCapability,
    countLine,
    derivedCounts,
    entriesWithoutRow,
    entryRankWord,
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
  import { button } from '../../lib/ui/button';
  import { ROW_RULE, ownRowBand } from '../../lib/ui/row';
  import Page from '../ui/Page.svelte';
  import Ruled from '../ui/Ruled.svelte';
  import Section from '../ui/Section.svelte';
  import Button from '../ui/Button.svelte';
  import Chips from '../ui/Chips.svelte';
  import Select from '../ui/Select.svelte';
  import { upperFirst } from '../../lib/text';
  import FactLine from './FactLine.svelte';
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
    onRemoveMember = () => {},
    onDecline,
    onApproveAll,
    rowPath,
    profileHref = () => null,
    mySiteSpec = null,
    now = new Date(),
    attendanceSource = 'reports',
    focusApproveAll = false,
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
    /** Remove a member who joined by invite link: their entry has no character row, so the API
     *  removes them by account (`account_key`). */
    onRemoveMember?: (entry: GuildRosterEntry) => void;
    /** Turn a waiting character away: the API's remove, so the row leaves the list. */
    onDecline: (row: GuildRosterRow) => void;
    onApproveAll: () => void;
    rowPath: (row: GuildRosterRow) => CharacterPath;
    /** The row's member profile, when the viewer may open it (an officer, or their own row). */
    profileHref?: (row: GuildRosterRow) => string | null;
    /** The spec the site's own record (the account's character, `/v1/me`) holds for the viewer's
     *  character; the other rows carry only the roster's spec. Null when unknown. */
    mySiteSpec?: string | null;
    /** The clock the snapshot's "received … ago" and each "seen … ago" read against; a test pins it. */
    now?: Date;
    /** The home's `attendance_source`: in `raids` mode the chip counts raids over two windows. */
    attendanceSource?: AttendanceSource;
    /** Approve all takes focus as it mounts: the Before raids open card's Approve button came here. */
    focusApproveAll?: boolean;
  } = $props();

  const canApprove = $derived(hasCapability(can, 'approve'));
  const canRemove = $derived(hasCapability(can, 'remove'));

  let filters = $state<RosterFilters>({ ...DEFAULT_ROSTER_FILTERS });
  let noAccountOnly = $state(false);
  let sortKey = $state<RosterSortKey>(DEFAULT_ROSTER_SORT);
  /** The waiting row whose Decline is asking first, and whether Approve all is. */
  let decliningKey = $state<string | null>(null);
  /** The verified row (character key) or invited entry (account key) whose Remove is asking first. */
  let removingKey = $state<string | null>(null);
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
  /** The figure columns some row fills: an empty column is not drawn. */
  const shownColumns = $derived(
    ROSTER_COLUMNS.filter((key) =>
      ordered.some((row) =>
        rosterColumns(row, attendanceSource).some((cell) => cell.key === key && cell.text !== ''),
      ),
    ),
  );
  /** A row's cells for the shown columns. */
  function rowColumns(row: GuildRosterRow): { key: RosterColumnKey; text: string }[] {
    return rosterColumns(row, attendanceSource).filter((cell) => shownColumns.includes(cell.key));
  }
  /** Any row that can be removed: only then does the action column take its room (a handoff sizes to itself). */
  const anyActions = $derived(canRemove && ordered.some((row) => row.may_remove));
  /** Any row the addon carried: only then does the In game slot take room on every row. */
  const anyInGame = $derived(roster.some((row) => inGameFacts(row, now).length > 0));
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
  const ROLE_OPTIONS: readonly { value: RosterRoleFilter; label: string }[] = [
    { value: 'all', label: 'All roles' },
    { value: 'tank', label: 'Tank' },
    { value: 'healer', label: 'Healer' },
    { value: 'dps', label: 'Damage' },
  ];
  const classSelectOptions = $derived([
    { value: 'all', label: 'All classes' },
    ...classOptions.map((cls) => ({ value: cls, label: upperFirst(cls) })),
  ]);
  /** The on/off filters, as the shared chips. */
  const toggleChips = $derived([
    { id: 'verified', label: 'Verified only', on: filters.verifiedOnly },
    { id: 'noaccount', label: guildRosterModelCopy.noAccountFilter, on: noAccountOnly },
    { id: 'floor', label: 'Below rating floor', on: filters.belowFloorOnly },
  ]);
  function onToggleChip(id: string): void {
    if (id === 'verified') toggle('verifiedOnly');
    else if (id === 'floor') toggle('belowFloorOnly');
    else noAccountOnly = !noAccountOnly;
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
  function doRemove(row: GuildRosterRow): void {
    removingKey = null;
    onRemove(row);
  }
  function doRemoveMember(entry: GuildRosterEntry): void {
    removingKey = null;
    onRemoveMember(entry);
  }
  function doDecline(row: GuildRosterRow): void {
    decliningKey = null;
    onDecline(row);
  }

  /** The sort chips, in the order a guild master reads them. The addon's gear gap is a fact on the
   *  row, not a sort the page offers. */
  const SORT_CHIPS: { key: RosterSortKey; label: string }[] = [
    { key: 'rank', label: 'Rank' },
    { key: 'ilvl', label: 'Item level' },
    { key: 'rating', label: 'Rating' },
    { key: 'attendance', label: 'Attendance' },
    { key: 'seen', label: 'Last seen' },
  ];
  /** Rating sorts nothing while no row has one. */
  const sortChips = $derived(
    SORT_CHIPS.filter(
      (chip) =>
        chip.key !== 'rating' || roster.some((row) => row.rating !== undefined && row.rating !== null),
    ),
  );

  /** What a row's spec line reads: the addon's spec in game, the site's where it differs. The viewer's
   *  own row compares against the site's character record, which is what their header chip shows. */
  function specLineOf(row: GuildRosterRow): string {
    const siteSpec = row.character_key === myCharacterKey && mySiteSpec !== null ? mySiteSpec : row.spec;
    return rosterSpecLine({ siteSpec, addonSpec: row.addon?.spec ?? null, className: row.class });
  }

  /** Focus lands on Cancel when a confirm opens, so a stray Enter backs out rather than removes. */
  function focusOnMount(node: HTMLElement): void {
    node.focus();
  }

  /** The hover on a stat chip: the rating's six parts, or which boss kill the best figure is from. */
  function chipTitle(chip: string, row: GuildRosterRow): string | undefined {
    if (chip.startsWith('Rating')) return ratingTooltip(row);
    if (chip.startsWith('Best') && row.best_parse !== undefined && row.best_parse !== null)
      return `Best boss kill: ${row.best_parse.encounter}`;
    return undefined;
  }

  function ratingTooltip(row: GuildRosterRow): string {
    const r = row.rating;
    if (r === undefined || r === null) return '';
    return `Output ${r.output} · Survival ${r.survival} · Mechanics ${r.mechanics} · Utility ${r.utility} · Preparation ${r.preparation} · Activity ${r.activity}`;
  }

  // Tailwind utilities only (no scoped <style>): see GuildTabs.svelte's own header comment
  // -- this tab's rows only ever mount after the guild's async roster fetch resolves, so
  // a scoped <style> block's CSS is silently dropped from this build's SSR-critical-CSS
  // inlining pass.
  const ROW = `guild-roster-row flex flex-wrap items-center gap-3 py-2.5 xl:flex-nowrap ${ROW_RULE}`;
  // A name with no account row: the same shape, a different class so the account rows stay
  // the only `guild-roster-row`s.
  const ENTRY_ROW = 'guild-roster-entry flex flex-wrap items-center gap-x-3 gap-y-2 py-2.5 xl:flex-nowrap';
  /** Each figure column's width from lg, the header's and every row's alike. */
  const COLUMN_WIDTH: Readonly<Record<RosterColumnKey, string>> = {
    level: 'w-[88px]',
    raids: 'w-[132px]',
    best: 'w-[104px]',
    rating: 'w-[80px]',
  };
</script>

{#snippet factList(facts: InGameFact[])}
  <!-- Each fact is one unbreakable piece (it wraps only as a whole) and carries its own trailing
       dot, so a wrap never leaves a dot at the start of a line. The spaces are mustaches because
       Svelte drops a bare space at the end of an {#each} body. -->
  <!-- eslint-disable svelte/no-useless-mustaches -->
  {#each facts as fact, i (fact.key)}
    <span class="whitespace-nowrap"
      ><span
        class={`${fact.key === 'seen' ? 'text-muted' : ''} ${fact.key === 'gap' ? 'tabular' : ''}`}
        style={fact.hover === undefined ? '' : 'cursor:default'}
        title={fact.hover}
        data-testid={fact.key === 'gap' ? 'guild-roster-ingame-gap' : undefined}>{fact.text}</span
      >{#if i < facts.length - 1}{' '}<span class="text-muted" aria-hidden="true"
          >{guildRosterAddonCopy.separator}</span
        >{/if}</span
    >{' '}
  {/each}
  <!-- eslint-enable svelte/no-useless-mustaches -->
{/snippet}

{#snippet inGame(row: GuildRosterRow)}
  {@const facts = inGameFacts(row, now)}
  <!-- The slot is the same width on every row, so the row's height never depends on how much
       the addon knew: its text wraps inside it. Below xl it has a line of its own (at 1024 a 260 px
       slot beside the actions would leave the stats too little room), and a row
       the snapshot did not carry has none there at all (no lone "—"). -->
  {#if anyInGame}
    <div
      class={`min-w-0 basis-full pl-12 text-[12px] break-words max-xl:order-2 xl:shrink-0 xl:basis-[260px] xl:pl-0 ${facts.length === 0 ? 'hidden xl:block' : ''}`}
      data-testid="guild-roster-ingame-slot"
    >
      {#if facts.length > 0}
        <span class="sr-only">{guildRosterAddonCopy.srLabel}</span>
        <span data-testid="guild-roster-ingame">{@render factList(facts)}</span>
      {/if}
    </div>
  {/if}
{/snippet}

{#snippet entryInGame(entry: GuildRosterEntry)}
  {@const facts = entryFacts(entry, now)}
  <div
    class={`text-nav min-w-0 basis-full pl-12 text-[12px] break-words max-xl:order-2 xl:shrink-0 xl:basis-[260px] xl:pl-0 ${facts.length === 0 ? 'hidden xl:block' : ''}`}
    data-testid="guild-roster-entry-ingame"
  >
    {#if facts.length > 0}
      <span class="sr-only">{guildRosterAddonCopy.srLabel}</span>
      <span>{@render factList(facts)}</span>
    {/if}
  </div>
{/snippet}

{#snippet removeConfirm(name: string, rejoin: string, onYes: () => void)}
  <!-- Indented to the name column at every width, as the award confirm is. Remove is the danger
   token (as Leave guild is), Cancel a ghost with the focus ring on open. -->
  <div
    class="flex flex-wrap items-center gap-x-4 gap-y-2 py-2.5 pl-12"
    role="group"
    aria-label={guildRosterModelCopy.removeQuestion(name)}
    data-testid="guild-roster-remove-confirm"
  >
    <div class="flex basis-full flex-col gap-0.5">
      <span class="text-[13px]">{guildRosterModelCopy.removeQuestion(name)}</span>
      <span class="text-muted text-[12px]" data-testid="guild-roster-remove-rejoin">{rejoin}</span>
    </div>
    <button class={button('danger', 'sm')} onclick={onYes} data-testid="guild-roster-remove-yes">
      {guildRosterModelCopy.removeButton}
    </button>
    <button
      class={button('ghost', 'sm')}
      onclick={() => (removingKey = null)}
      use:focusOnMount
      data-testid="guild-roster-remove-cancel"
    >
      {guildRosterModelCopy.cancel}
    </button>
  </div>
{/snippet}

{#snippet pendingRows()}
  <Ruled id="guild-roster-unverified" testid="guild-roster-unverified-list">
    {#each shownPending as row (row.character_key)}
      <div class={ROW}>
        <Crest cls={row.class ?? ''} size={36} />
        <div class="flex min-w-[150px] flex-col">
          <span class="flex min-w-0 items-center gap-2">
            <CharacterName
              name={row.name}
              cls={row.class}
              weight="inherit"
              class="heading-name text-[14px]"
            />
            {#if row.supporter === true}<SupporterChip />{/if}
          </span>
          <span class="text-muted text-[12px]">{specLineOf(row)}</span>
        </div>
        {#if row.rank !== 'member'}
          <Tag testid="guild-roster-pending-rank">{rankWord(row.rank)}</Tag>
        {/if}
        <div class="ml-auto flex shrink-0 flex-nowrap items-center justify-end gap-3">
          {#if canApprove}
            <Button
              size="sm"
              onclick={() => onApprove(row)}
              busy={rosterBusy === row.character_key}
              busyLabel={guildRosterModelCopy.approving}
              disabled={frozen}
              label={`Approve ${row.name}`}
              testid="guild-roster-approve"
            >
              Approve
            </Button>
          {/if}
          {#if canRemove && row.may_remove}
            <Button
              tone="ghost"
              size="sm"
              onclick={() => (decliningKey = row.character_key)}
              busy={rosterBusy === row.character_key}
              busyLabel={guildRosterModelCopy.declining}
              disabled={frozen && !removeOverridesFrozen(row)}
              label={`Decline ${row.name}`}
              testid="guild-roster-decline"
            >
              {guildRosterModelCopy.decline}
            </Button>
          {/if}
        </div>
      </div>
      {#if decliningKey === row.character_key}
        <div class="flex flex-wrap items-center gap-3 py-2.5" data-testid="guild-roster-decline-confirm">
          <span class="text-[13px]">{guildRosterModelCopy.declineQuestion(row.name)}</span>
          <button
            class={button('secondary', 'sm')}
            onclick={() => doDecline(row)}
            data-testid="guild-roster-decline-yes"
          >
            {guildRosterModelCopy.decline}
          </button>
          <button class={button('ghost', 'sm')} onclick={() => (decliningKey = null)}>
            {guildRosterModelCopy.cancel}
          </button>
        </div>
      {/if}
    {/each}
  </Ruled>
{/snippet}

<div data-testid="guild-roster-tab">
  <Page layout="ledger" panels>
    {#if shownPending.length > 0}
      <Section
        title="Waiting for approval"
        panel
        frame="framed"
        testid={canApprove ? 'guild-roster-waiting' : undefined}
      >
        {#snippet head()}
          {#if canApprove}<button
              class={button('primary', 'sm')}
              onclick={askApproveAll}
              disabled={frozen}
              data-testid="guild-roster-approve-all"
              use:focusOnMountIf={focusApproveAll}
            >
              Approve all
            </button>{/if}
        {/snippet}
        <div class="flex flex-col gap-3">
          {#if confirmingApproveAll}
            <div class="flex flex-wrap items-center gap-3" data-testid="guild-roster-approve-all-confirm">
              <span class="text-[14px]">{approveAllQuestion(pending)}</span>
              <button
                class={button('secondary', 'sm')}
                onclick={doApproveAll}
                disabled={frozen}
                data-testid="guild-roster-approve-all-yes"
              >
                {guildRosterModelCopy.confirm}
              </button>
              <button class={button('ghost', 'sm')} onclick={() => (confirmingApproveAll = false)}>
                {guildRosterModelCopy.cancel}
              </button>
            </div>
          {/if}
          {@render pendingRows()}
        </div>
      </Section>
    {/if}

    <Section
      title="Roster"
      panel
      meta={countLine(countsView)}
      metaTestid="guild-roster-counts"
      testid="guild-roster-card"
    >
      <div class="flex flex-col gap-4">
        <!-- Filter and Sort share one wrapping line (batch 3 review: the raiders start higher). -->
        <div class="flex flex-wrap items-center gap-x-6 gap-y-2">
          <div class="flex flex-wrap items-center gap-2" data-testid="guild-roster-filters">
            <span class="quiet-label">Filter</span>
            <Select
              id="guild-roster-filter-role"
              size="dense"
              label="Role"
              value={filters.role}
              options={ROLE_OPTIONS}
              testid="guild-roster-filter-role"
              onchange={(event) =>
                (filters = {
                  ...filters,
                  role: (event.currentTarget as HTMLSelectElement).value as RosterRoleFilter,
                })}
            />
            <Select
              id="guild-roster-filter-class"
              size="dense"
              label="Class"
              value={filters.classFilter}
              options={classSelectOptions}
              testid="guild-roster-filter-class"
              onchange={(event) =>
                (filters = { ...filters, classFilter: (event.currentTarget as HTMLSelectElement).value })}
            />
            <Chips
              chips={toggleChips}
              label="Filter"
              onToggle={onToggleChip}
              chipTestid="guild-roster-filter"
            />
          </div>

          <div class="flex flex-wrap items-center gap-2" data-testid="guild-roster-sort">
            <span class="quiet-label">Sort</span>
            <Chips
              chips={sortChips.map((chip) => ({
                id: chip.key,
                label: sortKey === chip.key ? `${chip.label} ▾` : chip.label,
                on: sortKey === chip.key,
              }))}
              label="Sort"
              onToggle={(id) => (sortKey = id as RosterSortKey)}
              chipTestid="guild-roster-sort"
            />
          </div>
        </div>
        <!-- Only while no rank is known: a roster row's rank is the addon's, so the line would read
           over ranks it already sent (guild round 3). -->
        {#if allRows.length === 0 && extras.some((entry) => entry.state === 'no_account')}
          <p class="text-muted text-[12px]" data-testid="guild-roster-ranks-pending">{RANKS_PENDING_LINE}</p>
        {/if}

        {#if ordered.length === 0 && extras.length === 0}
          <EmptyState message="No raiders match these filters." testid="guild-roster-empty" />
        {/if}
        {#if ordered.length > 0 || extras.length > 0}
          <div role="region" aria-label="Roster" data-testid="guild-roster-plate">
            <div>
              {#if ordered.length > 0}
                <div
                  class="text-label text-muted hidden items-center gap-3 pb-2 font-bold lg:flex"
                  data-testid="guild-roster-heads"
                  aria-hidden="true"
                >
                  <span class="w-9 shrink-0"></span>
                  <span class="w-[240px] shrink-0">{guildRosterModelCopy.heads.raider}</span>
                  <span class="flex flex-1 justify-end">
                    {#each shownColumns as key (key)}<span class={`text-right ${COLUMN_WIDTH[key]}`}
                        >{guildRosterModelCopy.heads[key]}</span
                      >{/each}
                  </span>
                  {#if anyInGame}<span class="hidden shrink-0 basis-[260px] xl:block"></span>{/if}
                  {#if anyActions}<span class="w-[216px] shrink-0"></span>{/if}
                </div>
                <Ruled testid="guild-roster-rows">
                  {#each ordered as row (row.character_key)}
                    {@const pinned = row.character_key === myCharacterKey}
                    {@const chips = rosterStatChips(row, attendanceSource)}
                    {@const profile = profileHref(row)}
                    {@const specLine = specLineOf(row)}
                    {@const subtitle = altOf.has(row.character_key)
                      ? `${specLine}${specLine === '' ? '' : ' · '}alt of ${altOf.get(row.character_key)}`
                      : specLine}
                    <div class={`${ROW} ${ownRowBand(pinned, 0)}`}>
                      <Crest cls={row.class ?? ''} size={36} />
                      <!-- At lg the name block is a fixed 240 px and its subtitle one ellipsised line (the whole
               text in the title), so every row's rank and stats start at the same left edge. -->
                      <div class="flex min-w-0 flex-1 flex-col lg:w-[240px] lg:flex-none lg:shrink-0">
                        <span class="flex min-w-0 items-center gap-2">
                          <CharacterName
                            name={`${row.name}${pinned ? ' (you)' : ''}`}
                            cls={row.class}
                            href={profile ?? undefined}
                            title={`${row.name}${pinned ? ' (you)' : ''}`}
                            weight="inherit"
                            class="heading-name min-w-0 text-[14px] focus-visible:underline lg:truncate"
                            testid="guild-roster-name"
                          />
                          {#if row.supporter === true}<SupporterChip />{/if}
                        </span>
                        <!-- One line per part at lg ("Beast Mastery Hunter in game", then "Survival on the site"), each
                 ellipsised to the 240 px block, so the disagreeing record is never the part cut off; one
                 wrapping run, parts joined by a dot, below lg. -->
                        <span
                          class="text-muted text-[12px] max-lg:break-words"
                          title={subtitle}
                          data-testid="guild-roster-spec"
                          >{#each subtitle.split(' · ') as part, i (i)}<span class="lg:block lg:truncate"
                              >{#if i > 0}<span class="lg:hidden" aria-hidden="true">&#32;·&#32;</span
                                >{/if}{part}{#if i === subtitle.split(' · ').length - 1 && row.rank !== 'member'}&#32;·&#32;<span
                                  data-testid="guild-roster-rank">{rankWord(row.rank)}</span
                                >{/if}</span
                            >{/each}</span
                        >
                      </div>
                      <!-- Below lg the row is three blocks under the name (indented past the crest): the stat
               chips, the professions on one line, and the in-game facts; Remove rides the name
               row. At lg the chips and professions wrap inside their own group, so the actions keep
               their place at the row's end; the In game slot joins them on the row from xl. -->
                      <div
                        class="flex min-w-0 basis-full flex-col gap-1 pl-12 max-lg:order-1 lg:flex-1 lg:basis-0 lg:pl-0"
                        data-testid="guild-roster-facts"
                      >
                        <div class="flex min-w-0 flex-wrap items-center gap-x-3 gap-y-1 lg:flex-nowrap">
                          <!-- At lg the stats wrap inside their own box beside the rank slot, never under it,
                   so a long stat line starts at the same left edge as a short one. Below lg the box
                   is `contents` and the stats wrap with the rank as before. -->
                          <div
                            class="contents lg:flex lg:min-w-0 lg:flex-1 lg:flex-wrap lg:items-center lg:gap-x-3 lg:gap-y-1"
                          >
                            {#if chips.length > 0}
                              <!-- Below lg the stats as one sentence in the body face, " · " between them (P.1);
                                   from lg the columns under the header carry the units. -->
                              <FactLine
                                class="tabular text-[12px] lg:hidden"
                                facts={chips.map((chip, i) => ({
                                  text: chip,
                                  title: chipTitle(chip, row),
                                  class: i === 0 ? '' : 'text-muted',
                                  testid: 'guild-roster-chip',
                                }))}
                              />
                            {/if}
                            {#if row.logged_recently && !everyoneOnlineToday}
                              <span class="text-muted text-[12px]" data-testid="guild-roster-online">
                                {guildHomeCopy.loggedRecently}
                              </span>
                            {/if}
                            <span
                              class="hidden text-[14px] tabular-nums lg:ml-auto lg:flex lg:items-center"
                              data-testid="guild-roster-columns"
                            >
                              {#each rowColumns(row) as cell (cell.key)}
                                <span
                                  class={`text-right ${COLUMN_WIDTH[cell.key]} ${cell.key === 'level' ? 'text-strong' : 'text-text'}`}
                                  >{cell.text}</span
                                >
                              {/each}
                            </span>
                          </div>
                        </div>
                        {#if row.professions !== undefined && row.professions.length > 0}
                          <span
                            class="text-muted block truncate text-[12px] lg:pl-[84px]"
                            data-testid="guild-roster-professions"
                            >{row.professions.map(professionName).join(' / ')}</span
                          >
                        {/if}
                      </div>
                      {@render inGame(row)}
                      <div
                        class="contents lg:ml-auto lg:flex lg:shrink-0 lg:flex-nowrap lg:items-center lg:justify-end lg:gap-3 {anyActions
                          ? 'lg:w-[216px]'
                          : 'lg:w-auto'}"
                        data-testid="guild-roster-actions"
                      >
                        {#if rosterHoldsHandoff(row)}
                          <div
                            class="basis-full pl-12 has-[[data-state=empty]]:hidden max-lg:order-3 lg:basis-auto lg:pl-0"
                          >
                            <GuildRosterHandoff path={rowPath(row)} stacked />
                          </div>
                        {/if}
                        {#if canRemove && row.may_remove}
                          <Button
                            size="sm"
                            onclick={() => (removingKey = row.character_key)}
                            busy={rosterBusy === row.character_key}
                            busyLabel={guildRosterModelCopy.removing}
                            disabled={frozen && !removeOverridesFrozen(row)}
                            label={`Remove ${row.name}`}
                            testid="guild-roster-remove"
                          >
                            {guildRosterModelCopy.removeButton}
                          </Button>
                        {/if}
                      </div>
                    </div>
                    {#if removingKey === row.character_key}
                      {@render removeConfirm(row.name, guildRosterModelCopy.removeRejoin, () =>
                        doRemove(row),
                      )}
                    {/if}
                  {/each}
                </Ruled>
              {/if}
              {#if extras.length > 0}
                <Ruled testid="guild-roster-entries">
                  {#each extras as entry (`${entry.state}:${entry.name}`)}
                    <div class={ENTRY_ROW} data-testid="guild-roster-entry" data-state={entry.state}>
                      <!-- No class known: the component's neutral ringed disc, never an empty outline. -->
                      <Crest cls={entry.class ?? ''} size={36} testid="guild-roster-entry-crest" />
                      <div class="flex min-w-0 flex-1 flex-col lg:w-[240px] lg:flex-none lg:shrink-0">
                        <CharacterName
                          name={entry.name}
                          cls={entry.class}
                          weight="inherit"
                          class="heading-name text-[14px]"
                          testid="guild-roster-entry-name"
                        />
                        {#if entrySpecLine(entry) !== ''}
                          <span
                            class="text-nav text-[12px] max-lg:break-words lg:truncate"
                            title={entrySpecLine(entry)}
                            data-testid="guild-roster-entry-spec">{entrySpecLine(entry)}</span
                          >
                        {/if}
                      </div>
                      <div
                        class="flex min-w-0 basis-full flex-col gap-2 pl-12 max-lg:order-1 lg:flex-1 lg:basis-0 lg:flex-row lg:items-center lg:gap-3 lg:pl-0"
                      >
                        <!-- A name only the addon saw carries the state word alone: the snapshot has no rank for
                 it, and saying "Member" would be a rank nobody gave it. -->
                        {#if entry.state !== 'no_account'}
                          <span class="text-nav text-[12px]" data-testid="guild-roster-entry-rank"
                            >{entryRankWord(entry)}</span
                          >
                        {/if}
                        {#if entry.state === 'no_account'}
                          <Tag
                            tone="muted"
                            title={`The addon saw this character in the guild; no account on the site has claimed it. ${entryRankWord(entry)} until the addon sends ranks.`}
                            testid="guild-roster-entry-state">{stateWord(entry.state)}</Tag
                          >
                        {:else}
                          <span class="text-nav text-[12px]" data-testid="guild-roster-entry-state"
                            >{stateWord(entry.state)}</span
                          >
                        {/if}
                      </div>
                      {@render entryInGame(entry)}
                      {#if canRemove && entry.may_remove === true && entry.account_key !== undefined}
                        {@const accountKey = entry.account_key}
                        <div
                          class="contents lg:ml-auto lg:flex lg:w-[216px] lg:shrink-0 lg:items-center lg:justify-end"
                          data-testid="guild-roster-entry-actions"
                        >
                          <Button
                            size="sm"
                            onclick={() => (removingKey = accountKey)}
                            busy={rosterBusy === accountKey}
                            busyLabel={guildRosterModelCopy.removing}
                            disabled={frozen}
                            label={`Remove ${entry.name}`}
                            testid="guild-roster-entry-remove"
                          >
                            {guildRosterModelCopy.removeButton}
                          </Button>
                        </div>
                      {:else}
                        <div class="ml-auto shrink-0 max-lg:hidden lg:w-[216px]" aria-hidden="true"></div>
                      {/if}
                    </div>
                    {#if entry.account_key !== undefined && removingKey === entry.account_key}
                      {@render removeConfirm(entry.name, guildRosterModelCopy.removeInvitedRejoin, () =>
                        doRemoveMember(entry),
                      )}
                    {/if}
                  {/each}
                </Ruled>
              {/if}
            </div>
          </div>
        {/if}
        {#if rosterActionError !== ''}
          <p class="text-[14px]" role="alert" data-testid="guild-roster-action-error">{rosterActionError}</p>
        {/if}
      </div>
    </Section>
  </Page>
</div>
