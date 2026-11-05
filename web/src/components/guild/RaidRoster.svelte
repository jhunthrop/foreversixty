<!-- web/src/components/guild/RaidRoster.svelte -->
<!-- Sign-ups (raids spec 3.7, raid comp spec 3.3, 4 "Officer state select"): In rows under Tanks,
     Healers, Melee and Ranged by the API's group (four columns from 1280; two from 768, Tanks over
     Healers beside Melee over Ranged; one below), then Late, Bench, Out and Not signed up. Every row
     reads its spec and class. Officers get a label-less state select at the end of every row; a refusal reverts
     it, with the line under the roster, or under the row for the plan gate (focused). A change that
     moves the row to another group hands focus to its select there; a saving select stays enabled
     (aria-busy), so the focused control is never disabled under the officer's hand. A guild event
     counts no roles: its sign-ups are one list by state (In, Late, Bench, Out), each row crest,
     name and class (guild events spec 3.5, guild round 3). -->
<script lang="ts">
  import { tick } from 'svelte';
  import type { GuildRaidGroup, GuildRaidView, GuildSignupState } from '../../lib/guild/api';
  import { kindCountsRoles, kindOf, raidCopy } from '../../lib/guild/raid-copy';
  import { focusOnMount } from '../../lib/guild/focus-now';
  import { raidCount } from '../../lib/guild/raid-page';
  import type { OfficerLine } from '../../lib/guild/raid-page-types';
  import { ownRowBandInside } from '../../lib/ui/row';
  import { rosterGroups, rosterSelectId, type RosterRow } from '../../lib/guild/raid-roster';
  import { SIGNUP_STATES } from '../../lib/guild/segments-keys';
  import type { FilterOption } from '../../lib/rankings/options';
  import { classColorVar } from '../../lib/report/format';
  import { classWord, specClassLine } from '../../lib/guild/raid-comp';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import Select from '../ui/Select.svelte';
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

  const GROUP_COLUMNS: readonly { group: GuildRaidGroup; label: string; place: string }[] = [
    { group: 'tank', label: raidCopy.subheads.tanks, place: 'md:col-start-1 md:row-start-1' },
    {
      group: 'healer',
      label: raidCopy.subheads.healers,
      place: 'md:col-start-1 md:row-start-2 xl:col-start-2 xl:row-start-1',
    },
    {
      group: 'melee',
      label: raidCopy.subheads.melee,
      place: 'md:col-start-2 md:row-start-1 xl:col-start-3',
    },
    {
      group: 'ranged',
      label: raidCopy.subheads.ranged,
      place: 'md:col-start-2 md:row-start-2 xl:col-start-4 xl:row-start-1',
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
  /** A raid or a run splits In by role; a guild event lists In beside Late, Bench and Out. */
  const roles = $derived(kindCountsRoles(kindOf(raid)));
  const others = $derived([
    ...(roles ? [] : [{ id: 'in', label: raidCopy.segments.in, rows: groups.in, withRole: false }]),
    { id: 'late', label: raidCopy.subheads.late, rows: groups.late, withRole: roles },
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
  {@const specLine = roles
    ? `${specClassLine(entry.class, entry.spec_key)}${withRole ? ` · ${raidCopy.roleLabels[entry.role]}` : ''}`
    : classWord(entry.class)}
  <!-- §7 caps: 36 px desktop, 44 phone before a note. A bare row centres crest, name and the
       officer select (32 desktop, 44 phone) with no padding past them; a noted row runs from the top. -->
  <li
    class="rounded-control flex min-h-11 break-inside-avoid flex-wrap gap-x-3 {ownRowBandInside(
      entry.own,
    )} md:min-h-9 {entry.note === '' && entry.also === null ? 'items-center' : 'items-start py-1.5'}"
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
      <!-- One row per account: its other answered characters, so a benched alt is never hidden. -->
      {#if entry.also !== null}
        <!-- Phone: wrapped under the spec line (guild round 3). -->
        <p
          class="text-muted text-[12px] leading-4 break-words md:hidden"
          data-testid="raid-roster-also-phone"
        >
          {entry.also}
        </p>
      {/if}
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
      <!-- 104 px, under the select's bare-field floor (select.css .sel-root, 120 px, which
           `min-w-0` lifts): `Bench` and the chevron fit, and the name block gains the 16 px that
           lets `Beast Mastery Hunter` read in full in a four-column roster at 1440 and 2000 (guild
           loop round 3). The field and its chevron stay inside the row and its band. -->
      <div class="w-26 shrink-0">
        <Select
          class="min-w-0"
          id={rosterSelectId(entry.character_key)}
          size="dense"
          align="end"
          label={raidCopy.page.signupFor(entry.name)}
          value={shownValue(entry)}
          options={entry.state === null ? [NOT_SIGNED_OPTION, ...STATE_OPTIONS] : STATE_OPTIONS}
          disabled={!officerTools || closed}
          describedBy={officerTools ? undefined : gateId}
          {busy}
          testid={`raid-roster-select-${entry.name}`}
          onchange={(event) => void change(entry, (event.currentTarget as HTMLSelectElement).value)}
        />
      </div>
    {/if}
    {#if entry.also !== null}
      <!-- From md the Also line spans the row under the name block and the select, indented past
           the crest (24 + 12 px), so it reads in full; cut with the whole line as its title only
           when the row still cannot hold it (guild round 3: the name column ellipsised it). -->
      <p
        class="text-muted hidden basis-full truncate pl-9 text-[12px] leading-4 md:block"
        title={entry.also}
        data-testid="raid-roster-also"
      >
        {entry.also}
      </p>
    {/if}
    {#if entry.note !== ''}
      <!-- The note spans the row under the Also line at every width, indented past the crest, so
           Also always reads before the note (guild loop round 3: phone had Also first, desktop
           the note first). -->
      <p class="text-text basis-full pl-9 text-[12px] break-words" data-testid="raid-roster-note">
        {entry.note}
      </p>
    {/if}
  </li>
{/snippet}

{#snippet body()}
  {#if roles}
    <div class="grid grid-cols-1 gap-x-4 gap-y-6 md:grid-cols-2 xl:grid-cols-4">
      {#each GROUP_COLUMNS as column (column.group)}
        {@const rows = groups.groups[column.group]}
        <div class="flex min-w-0 flex-col gap-2 {column.place}" data-testid={`raid-roster-${column.group}`}>
          {@render heading(column.label, rows.length, `raid-roster-${column.group}-heading`)}
          {#if rows.length === 0}
            <p class="text-muted text-[13px]">{raidCopy.emptyGroup[column.group]}</p>
          {:else}
            <ul>
              {#each rows as entry (entry.character_key)}{@render row(entry, false)}{/each}
            </ul>
          {/if}
        </div>
      {/each}
    </div>
  {/if}
  <div class="flex flex-col gap-6 {roles ? 'mt-2' : ''}">
    <!-- Late, Bench and Out side by side from 1024, two columns from 768, stacked below; an empty
           group is left out, so the rest flow left (spec 3.7). -->
    {#if shownOthers.length > 0}
      <div
        class="grid grid-cols-1 gap-x-4 gap-y-6 md:grid-cols-2 {roles ? 'lg:grid-cols-3' : 'xl:grid-cols-4'}"
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
          <!-- A strict grid: every cell one person, rows 8 px apart, so two names in a column never
               sit with their selects touching (guild round 2). -->
          <ul
            class="grid grid-cols-1 gap-x-4 gap-y-2 md:grid-cols-2 xl:grid-cols-4"
            data-testid="raid-roster-not-signed-grid"
          >
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
