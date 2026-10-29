<!-- web/src/components/guild/RaidAttendance.svelte -->
<!-- Attendance (attendance spec 3.5, officers, `<section id="attendance">`): the summary sentence,
     then five groups in order, each with a G1 group head. The first three groups take full rows
     (who, the signed chip, the outcome, pulls with a bar, and `Excused` on Not in the log and Late
     rows only); Present and Bench and out take compact rows. An Excused write is optimistic: the
     controller moves the row and the summary, and a refusal puts them back with the row's line. -->
<script lang="ts">
  import type { RaidAttendanceRow, RaidNight } from '../../lib/guild/attendance-contract';
  import { attendanceCopy, attendanceSummary, pullsRatio } from '../../lib/guild/attendance-copy';
  import { bossPullsHead, excusedLabel, pullsOf } from '../../lib/guild/night-copy';
  import { raidCopy } from '../../lib/guild/raid-copy';
  import { canExcuse, groupAttendance, outcomeText } from '../../lib/guild/night';
  import { BAR_FILL, BAR_TRACK, OWN_ROW_WASH, ROW_DIVIDER } from '../../lib/guild/raid-report-styles';
  import { classColorVar } from '../../lib/report/format';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import Figures from './Figures.svelte';
  import RaidChip from './RaidChip.svelte';

  let {
    night,
    rows,
    busy = {},
    lines = {},
    readOnly,
    gateId,
    onExcuse,
  }: {
    night: RaidNight;
    rows: readonly RaidAttendanceRow[];
    /** Characters whose Excused write is in flight. */
    busy?: Readonly<Record<string, boolean>>;
    /** A refused write's line, by character. */
    lines?: Readonly<Record<string, string>>;
    /** Officer-read-only: the checkboxes stay, disabled, and name the gate sentence. */
    readOnly: boolean;
    gateId: string;
    onExcuse: (characterKey: string, excused: boolean) => void;
  } = $props();

  const groups = $derived(groupAttendance(rows));
  const washOf = (row: RaidAttendanceRow): string | undefined =>
    row.own ? `background:${OWN_ROW_WASH}` : undefined;
  const missed = (row: RaidAttendanceRow): boolean => row.pulls_present < row.pulls_total;
</script>

{#snippet who(row: RaidAttendanceRow, size: number)}
  <span class="flex min-w-0 items-center gap-3">
    <ClassCrestRing characterClass={row.class} {size} />
    <span class="min-w-0">
      <span
        class="block truncate text-[14px] font-semibold"
        style={`color:${classColorVar(row.class)}`}
        title={row.name}>{row.name}</span
      >
      <span class="text-muted block truncate text-[12px]">{row.spec}</span>
    </span>
  </span>
{/snippet}

{#snippet fullRow(row: RaidAttendanceRow)}
  <li
    class="{ROW_DIVIDER} grid grid-cols-[minmax(0,1fr)_auto] items-center gap-x-4 gap-y-1 px-2 py-2 md:min-h-[52px] md:grid-cols-[minmax(200px,1.4fr)_132px_minmax(180px,1fr)_120px_112px]"
    style={washOf(row)}
    aria-busy={busy[row.character_key] === true}
    data-testid="raid-attendance-row"
  >
    {@render who(row, 24)}
    <span class="justify-self-end md:justify-self-start"
      ><RaidChip kind="viewer" state={row.signed} role={row.role ?? 'dps'} /></span
    >
    <!-- Phone: outcome and pulls share line 2; from md each takes its own grid cell. -->
    <div class="col-span-2 flex items-center justify-between gap-3 md:contents">
      <span
        class="text-[13px] {row.outcome === 'no_show' ? 'text-muted' : 'text-text'}"
        data-testid="raid-attendance-outcome">{outcomeText(row)}</span
      >
      <span class="flex items-center gap-2">
        <span class="quiet-label whitespace-nowrap md:hidden">{bossPullsHead}</span>
        <span class="text-text text-[13px]"
          ><Figures text={pullsOf(row.pulls_present, row.pulls_total)} /></span
        >
        <span class="{BAR_TRACK} w-12" aria-hidden="true"
          ><span
            class={BAR_FILL}
            style:width={`${row.pulls_total === 0 ? 0 : (row.pulls_present / row.pulls_total) * 100}%`}
          ></span></span
        >
      </span>
    </div>
    <span class="col-span-2 md:col-span-1">
      {#if canExcuse(row)}
        <label class="text-text flex min-h-11 items-center gap-2 text-[13px]">
          <input
            type="checkbox"
            class="accent-gold size-4"
            checked={row.excused}
            disabled={readOnly}
            aria-label={excusedLabel(row.name)}
            aria-describedby={readOnly ? gateId : undefined}
            onchange={(event) => onExcuse(row.character_key, event.currentTarget.checked)}
            data-testid="raid-attendance-excused"
          />
          {attendanceCopy.attendance.excused}
        </label>
      {/if}
    </span>
    {#if lines[row.character_key] !== undefined}
      <p
        class="text-wipe col-span-2 text-[13px] md:col-span-5"
        role="alert"
        data-testid="raid-attendance-line"
      >
        {lines[row.character_key]}
      </p>
    {/if}
  </li>
{/snippet}

{#snippet compactRow(row: RaidAttendanceRow)}
  <li
    class="rounded-control flex min-h-11 items-center gap-3 px-2 md:min-h-9"
    style={washOf(row)}
    data-testid="raid-attendance-compact-row"
  >
    <ClassCrestRing characterClass={row.class} size={24} />
    <span class="min-w-0 flex-1 truncate">
      <span class="text-[14px] font-semibold" style={`color:${classColorVar(row.class)}`}>{row.name}</span>
      {#if row.came_as !== null}<span class="text-muted text-[12px]"> as {row.came_as}</span>{/if}
    </span>
    {#if row.outcome === 'absent' && row.signed !== null}
      <!-- Bench and out: which one, muted, in place of a pull count (UX 20, WP 17). -->
      <span class="text-muted text-[13px]">{raidCopy.segments[row.signed]}</span>
    {:else}
      <span class="{missed(row) ? 'text-text' : 'text-muted'} numeric text-[13px]"
        >{pullsRatio(row.pulls_present, row.pulls_total)}</span
      >
    {/if}
  </li>
{/snippet}

<section
  id="attendance"
  class="flex min-w-0 scroll-mt-4 flex-col gap-4"
  aria-labelledby="raid-attendance-heading"
  data-testid="raid-attendance"
>
  <h2 id="raid-attendance-heading" class="section-title">{attendanceCopy.attendance.head}</h2>
  <p class="text-text text-[14px]" data-testid="raid-attendance-summary">{attendanceSummary(night)}</p>
  {#if groups.some((group) => !group.compact)}
    <div
      class="hidden gap-x-4 px-2 md:grid md:grid-cols-[minmax(200px,1.4fr)_132px_minmax(180px,1fr)_120px_112px]"
      data-testid="raid-attendance-heads"
    >
      <span class="quiet-label col-start-4">{bossPullsHead}</span>
    </div>
  {/if}
  {#each groups as group (group.key)}
    <div class="flex min-w-0 flex-col gap-2" data-testid={`raid-attendance-${group.key}`}>
      <h3 class="label text-muted flex items-baseline justify-between gap-3">
        <span>{group.label} <span class="numeric">{group.rows.length}</span></span>
        {#if group.key === 'present'}<span class="text-[11px] whitespace-nowrap md:hidden"
            >{bossPullsHead}</span
          >{/if}
      </h3>
      {#if group.compact}
        <ul class="grid grid-cols-1 gap-x-4 md:grid-cols-2 lg:grid-cols-3">
          {#each group.rows as row (row.character_key)}{@render compactRow(row)}{/each}
        </ul>
      {:else}
        <ul>
          {#each group.rows as row (row.character_key)}{@render fullRow(row)}{/each}
        </ul>
      {/if}
    </div>
  {/each}
</section>
