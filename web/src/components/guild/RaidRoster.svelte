<!-- web/src/components/guild/RaidRoster.svelte -->
<!-- Sign-ups (raids spec 3.7, 4 "Officer state select"): In rows under Tanks, Healers and DPS (four
     columns from 1280 with DPS across two, three from 1024, two from 768), then Late, Bench, Out and
     Not signed up. Officers get a label-less state select at the end of every row; a refusal reverts
     it, with the line under the roster, or under the row for the plan gate (focused). A change that
     moves the row to another group hands focus to its select there; a saving select stays enabled
     (aria-busy), so the focused control is never disabled under the officer's hand. -->
<script lang="ts">
  import { tick } from 'svelte';
  import type { GuildRaidRole, GuildRaidView, GuildSignupState } from '../../lib/guild/api';
  import { raidCopy } from '../../lib/guild/raid-copy';
  import { focusOnMount } from '../../lib/guild/focus-now';
  import { raidCount } from '../../lib/guild/raid-page';
  import type { OfficerLine } from '../../lib/guild/raid-page-types';
  import { OWN_ROW_WASH } from '../../lib/guild/raid-report-styles';
  import { rosterGroups, rosterSelectId, type RosterRow } from '../../lib/guild/raid-roster';
  import { SIGNUP_STATES } from '../../lib/guild/segments-keys';
  import type { FilterOption } from '../../lib/rankings/options';
  import { classColorVar } from '../../lib/report/format';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import FilterSelect from '../rankings/FilterSelect.svelte';
  import RaidDisclosure from './RaidDisclosure.svelte';

  let {
    raid,
    officer,
    officerTools,
    gateId,
    cancelled = false,
    closed = false,
    collapsed = false,
    busyKey = null,
    line = null,
    onSetState,
  }: {
    raid: GuildRaidView;
    officer: boolean;
    officerTools: boolean;
    /** The gate sentence a read-only select names in aria-describedby. */
    gateId: string;
    cancelled?: boolean;
    /** A cancelled or done raid takes no sign-up changes: the selects stay, disabled. */
    closed?: boolean;
    /** A done raid (R-A4): the roster sits closed below Attendance, its heading in the summary. */
    collapsed?: boolean;
    busyKey?: string | null;
    line?: OfficerLine | null;
    /** Resolves false on a refusal, so the select reverts. */
    onSetState: (row: RosterRow, state: GuildSignupState) => Promise<boolean>;
  } = $props();

  const ROLE_COLUMNS: readonly { role: GuildRaidRole; label: string; place: string }[] = [
    { role: 'tank', label: raidCopy.subheads.tanks, place: 'md:col-start-1 md:row-start-1' },
    {
      role: 'healer',
      label: raidCopy.subheads.healers,
      place: 'md:col-start-1 md:row-start-2 lg:col-start-2 lg:row-start-1',
    },
    {
      role: 'dps',
      label: raidCopy.subheads.dps,
      place: 'md:col-start-2 md:row-span-2 md:row-start-1 lg:col-start-3 lg:row-span-1 xl:col-span-2',
    },
  ];
  const STATE_OPTIONS: readonly FilterOption[] = SIGNUP_STATES.map((state) => ({
    value: state,
    label: raidCopy.segments[state],
  }));
  const NOT_SIGNED_OPTION: FilterOption = {
    value: '',
    label: raidCopy.page.notSignedMark,
    disabled: true,
    ariaLabel: raidCopy.subheads.notSigned,
  };

  const groups = $derived(rosterGroups(raid));
  const count = $derived(raidCount(raid));
  const others = $derived([
    { id: 'late', label: raidCopy.subheads.late, rows: groups.late, withRole: true },
    { id: 'bench', label: raidCopy.subheads.bench, rows: groups.bench, withRole: false },
    { id: 'out', label: raidCopy.subheads.out, rows: groups.out, withRole: false },
  ]);
  const shownOthers = $derived(others.filter((group) => group.rows.length > 0));
  /** The value each select shows while its save is in flight; dropped when it settles, so the
   *  select reads the raid again (a refusal's revert, a success's saved state). */
  let chosen = $state<Readonly<Record<string, string>>>({});

  function isState(value: string): value is GuildSignupState {
    return (SIGNUP_STATES as readonly string[]).includes(value);
  }

  function shownValue(entry: RosterRow): string {
    return chosen[entry.character_key] ?? entry.state ?? '';
  }

  async function change(row: RosterRow, value: string): Promise<void> {
    if (!isState(value)) return;
    const key = row.character_key;
    chosen = { ...chosen, [key]: value };
    const saved = await onSetState(row, value);
    chosen = Object.fromEntries(Object.entries(chosen).filter(([k]) => k !== key));
    if (!saved) return;
    await tick();
    document.getElementById(rosterSelectId(key))?.focus();
  }
</script>

{#snippet heading(label: string, n: number, testid: string)}
  <h3 class="label text-muted" data-testid={testid}>
    {label} <span class="numeric">{n}</span>
  </h3>
{/snippet}

{#snippet row(entry: RosterRow, withRole: boolean)}
  {@const busy = busyKey === entry.character_key}
  {@const specLine = `${entry.spec}${withRole ? ` · ${raidCopy.roleLabels[entry.role]}` : ''}`}
  <!-- §7 caps: 36 px desktop, 44 phone before a note. A bare row centres crest, name and the
       officer select (32 desktop, 44 phone) with no padding past them; a noted row runs from the top. -->
  <li
    class="rounded-control flex min-h-11 break-inside-avoid gap-3 px-2 md:min-h-9 {entry.note === ''
      ? 'items-center'
      : 'items-start py-1.5'}"
    style={entry.own ? `background: ${OWN_ROW_WASH}` : undefined}
    aria-busy={busy}
    data-testid="raid-roster-row"
  >
    <span class="flex shrink-0"><ClassCrestRing characterClass={entry.class} size={24} /></span>
    <div class="min-w-0 flex-1">
      <!-- The name over the spec at every width, one line each (34 px), so a column's rows line up. -->
      <p
        class="block truncate text-[14px] leading-[18px] font-semibold"
        style={`color:${classColorVar(entry.class)}`}
        title={entry.name}
      >
        {entry.name}
      </p>
      <p class="text-muted block truncate text-[12px] leading-4" title={specLine}>{specLine}</p>
      {#if entry.note !== ''}<p class="text-text text-[12px] break-words">{entry.note}</p>{/if}
      {#if line?.control === `roster:${entry.character_key}`}
        <p
          class="text-wipe text-[13px]"
          role="alert"
          tabindex="-1"
          use:focusOnMount
          data-testid="raid-roster-refusal"
        >
          {line.message}
        </p>
      {/if}
    </div>
    {#if officer}
      <div class="w-24 shrink-0">
        <FilterSelect
          id={rosterSelectId(entry.character_key)}
          label={raidCopy.page.signupFor(entry.name)}
          labelHidden
          compact
          value={shownValue(entry)}
          options={entry.state === null ? [NOT_SIGNED_OPTION, ...STATE_OPTIONS] : STATE_OPTIONS}
          disabled={!officerTools || closed}
          describedBy={officerTools ? undefined : gateId}
          {busy}
          onchange={(value) => void change(entry, value)}
          testid={`raid-roster-select-${entry.name}`}
        />
      </div>
    {/if}
  </li>
{/snippet}

{#snippet body()}
  <div class="grid grid-cols-1 gap-x-4 gap-y-6 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
    {#each ROLE_COLUMNS as column (column.role)}
      {@const rows = groups.roles[column.role]}
      <div class="flex min-w-0 flex-col gap-2 {column.place}" data-testid={`raid-roster-${column.role}`}>
        {@render heading(column.label, rows.length, `raid-roster-${column.role}-heading`)}
        {#if rows.length === 0}
          <p class="text-muted text-[13px]">{raidCopy.emptyRole[column.role]}</p>
        {:else}
          <ul class={column.role === 'dps' ? 'xl:columns-2 xl:gap-4' : ''}>
            {#each rows as entry (entry.character_key)}{@render row(entry, false)}{/each}
          </ul>
        {/if}
      </div>
    {/each}
  </div>
  <div class="mt-2 flex flex-col gap-6">
    <!-- Late, Bench and Out side by side from 1024, two columns from 768, stacked below; an empty
           group is left out, so the rest flow left (spec 3.7). -->
    {#if shownOthers.length > 0}
      <div
        class="grid grid-cols-1 gap-x-4 gap-y-6 md:grid-cols-2 lg:grid-cols-3"
        data-testid="raid-roster-others"
      >
        {#each shownOthers as group (group.id)}
          <div class="flex min-w-0 flex-col gap-2" data-testid={`raid-roster-${group.id}`}>
            {@render heading(group.label, group.rows.length, `raid-roster-${group.id}-heading`)}
            <ul>
              {#each group.rows as entry (entry.character_key)}{@render row(entry, group.withRole)}{/each}
            </ul>
          </div>
        {/each}
      </div>
    {/if}
    {#if groups.notSigned.length > 0}
      <div class="flex flex-col gap-2" data-testid="raid-roster-not-signed">
        {@render heading(
          raidCopy.subheads.notSigned,
          groups.notSigned.length,
          'raid-roster-not-signed-heading',
        )}
        {#if officer}
          <ul class="grid grid-cols-1 gap-x-4 md:grid-cols-2 xl:grid-cols-4">
            {#each groups.notSigned as entry (entry.character_key)}{@render row(entry, false)}{/each}
          </ul>
        {:else}
          <p class="flex flex-wrap gap-x-3 gap-y-1 text-[13px]">
            {#each groups.notSigned as entry (entry.character_key)}
              <span style={`color:${classColorVar(entry.class)}`}>{entry.name}</span>
            {/each}
          </p>
        {/if}
      </div>
    {/if}
  </div>
  {#if line?.control === 'roster'}
    <p class="text-wipe text-[13px]" role="alert" data-testid="raid-roster-error">{line.message}</p>
  {/if}
{/snippet}

{#if collapsed}
  <RaidDisclosure
    heading={raidCopy.subheads.signups}
    headingId="raid-roster-heading"
    count={count.ratio}
    testid="raid-roster-disclosure"
  >
    <section class="flex flex-col gap-4" aria-labelledby="raid-roster-heading" data-testid="raid-roster">
      {@render body()}
    </section>
  </RaidDisclosure>
{:else}
  <section
    class="flex flex-col gap-4 {cancelled ? 'opacity-60' : ''}"
    aria-labelledby="raid-roster-heading"
    data-testid="raid-roster"
  >
    <div class="flex items-baseline gap-3">
      <h2 id="raid-roster-heading" class="section-title">{raidCopy.subheads.signups}</h2>
      <span class="text-muted numeric text-[13px]" data-testid="raid-roster-ratio">{count.ratio}</span>
    </div>
    {@render body()}
  </section>
{/if}
