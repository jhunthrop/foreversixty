<!-- web/src/components/report/CastTable.svelte -->
<!-- Casts (logs spec 2026-11-07 section 6.2): the tab's one Plate, a StatTable of the spell's
     icon, the spell over its caster (class colour; a pet reads `{pet} ({owner}'s pet)`), a Spark
     tick strip of the sequence, Casts, Per min and Failed. The sequence is the only way to see a
     rotation's shape without opening the events file.

     Three different truths sit in one row. The sequence is exact under any window --
     window.ts's scopeCastRow filters the millisecond offsets themselves, rather than
     scaling anything. Cast and Failed are not: the engine keeps only a whole-fight count,
     so scopeCastRow scales both by the window's share of the sequence, and they carry the
     `~` mark ActorRow uses for its per-ability and per-target splits. Cast time is a third
     kind of lie: logs/engine/summary/deaths.go accumulates it across every successful
     cast in the whole fight and scopeCastRow never touches it, so under any window it is
     silently the whole fight's total -- the `†` mark SummaryTab's Active column uses. -->
<script lang="ts">
  import EmptyState from '../ui/EmptyState.svelte';
  import { splitUnitName } from '../../lib/characters';
  import { approximateMark, classColorVar, wholeFightMark } from '../../lib/report/format';
  import type { CastRow } from '../../lib/report/types';
  import { castKey } from '../../lib/report/cast-key';
  import type { CastCounts } from '../../lib/report/exact';
  import { SpellIcons } from '../../lib/report/spell-icons.svelte';
  import type { StatColumn, StatFigure, StatSort } from '../../lib/ui/stat-table';
  import { clockDuration } from '../../lib/format/duration';
  import Section from '../ui/Section.svelte';
  import Spark from '../ui/Spark.svelte';
  import SpellIcon from '../ui/SpellIcon.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import CopyCsv from './CopyCsv.svelte';

  let {
    rows,
    everyone = rows,
    durationMs,
    startMs = 0,
    classOf = new Map<string, string>(),
    approximate = false,
    measured = undefined,
    measureError = '',
    whole = undefined,
    names = new Map<string, string>(),
  }: {
    rows: CastRow[];
    /** Every caster's whole-fight rows: who recorded failures is a fact about the log, not the scope or the window. */
    everyone?: CastRow[];
    durationMs: number;
    startMs?: number;
    classOf?: Map<string, string>;
    /** GUID to unit name, so a pet's row can be filed under its owner's name. */
    names?: ReadonlyMap<string, string>;
    /** True when the window is brushed, so Cast and Failed are a scaled share. */
    approximate?: boolean;
    /** The window's own counts per caster and spell, once the fight's events were read. */
    measured?: ReadonlyMap<string, CastCounts>;
    /**
     * The whole fight's rows in this source scope: a spell only started, never landed,
     * inside the window has no windowed row (the window keeps rows by their successes),
     * and its measured counts need a row to sit on.
     */
    whole?: CastRow[];
    /** Why the measure did not run, when it did not: the scaled figures stay, marked. */
    measureError?: string;
  } = $props();
  /** The player a row belongs to: a pet's owner, or the caster themselves. */
  const ownerOf = (row: CastRow): string => row.owner_guid ?? row.guid;
  /** The pet's own name when the row is a pet's; '' for a caster's own row. */
  const viaOf = (row: CastRow): string => (ownerOf(row) === row.guid ? '' : splitUnitName(row.name).name);
  /** The name the Caster column shows: the owner's, since that is whose page this is. */
  const casterName = (row: CastRow): string =>
    splitUnitName(names.get(ownerOf(row)) ?? (ownerOf(row) === row.guid ? row.name : ownerOf(row))).name;

  /**
   * Under a brush the measured counts stand in for the scaled ones, row by row, and a
   * spell the window only saw started or refused gets a row of its own: the cancelled
   * cast a player most wants to see is the one that never went off.
   */
  const shown = $derived.by(() => {
    if (measured === undefined) return rows;
    const present = new Set(rows.map(castKey));
    const clippedOnly = (whole ?? rows)
      .filter((row) => !present.has(castKey(row)))
      .flatMap((row) => {
        const counts = measured.get(castKey(row));
        return counts === undefined || (counts.started === 0 && counts.failed === 0)
          ? []
          : [{ ...row, ...counts, sequence: [] }];
      });
    return [
      ...rows.map((row) => {
        const counts = measured.get(castKey(row));
        return counts === undefined ? row : { ...row, ...counts };
      }),
      ...clippedOnly,
    ];
  });
  /** True while the figures are the summary's scaled ones: a brush with no measure yet. */
  const scaled = $derived(approximate && measured === undefined);

  /** Failure reasons that mean a cast bar was cut short, not a press the game refused. */
  const CUT_SHORT = /interrupt|moving|cancel/i;
  const cutShort = (row: CastRow): number =>
    Object.entries(row.fail_reasons ?? {})
      .filter(([reason]) => CUT_SHORT.test(reason))
      .reduce((sum, [, count]) => sum + count, 0);
  /**
   * Hardcasts begun that never went off: the sequence's gaps a player can close. Two
   * signals say so -- a cast start with no success, and a failure the game filed as
   * interrupted or moved -- and a client logs one, the other or both, so the larger wins.
   */
  const cancelled = (row: CastRow): number => Math.max(row.started - row.succeeded, cutShort(row));
  /** Presses the game refused outright: range, target, mana. The cut-short ones are Cancelled. */
  const refused = (row: CastRow): number => Math.max(row.failed - cutShort(row), 0);
  /**
   * Only the client that wrote the log records failed casts, so a zero on everyone else
   * is silence, not a clean sheet: their cell reads a dash.
   */
  const recorders = $derived(new Set(everyone.filter((row) => row.failed > 0).map((row) => row.guid)));
  const failedKnown = (row: CastRow): boolean => recorders.size === 0 || recorders.has(row.guid);
  type CastSortColumn = 'casts' | 'perMin' | 'failed';
  let sort = $state<StatSort>({ column: 'casts', direction: 'desc' });
  function onsort(column: string): void {
    sort =
      sort.column === column
        ? { column, direction: sort.direction === 'desc' ? 'asc' : 'desc' }
        : { column, direction: 'desc' };
  }
  /** Casts and Per min order alike (one window, one length); Failed orders an unknown count last. */
  const sortKey = (row: CastRow, column: CastSortColumn): number =>
    column === 'failed' ? (failedKnown(row) ? refused(row) : -1) : row.succeeded;
  const ordered = $derived(
    [...shown].sort((a, b) => {
      const order = sortKey(b, sort.column as CastSortColumn) - sortKey(a, sort.column as CastSortColumn);
      return (sort.direction === 'desc' ? order : -order) || a.spell_name.localeCompare(b.spell_name);
    }),
  );
  /** Cancelled casts over the rows on screen: the total the rhythm line adds up. */
  const cancelledTotal = $derived(ordered.reduce((sum, row) => sum + cancelled(row), 0));
  /** The table as lines: one per caster and spell, with the casting time in seconds. */
  function csvLines(): string[][] {
    return [
      ['Player', 'Via', 'Spell', 'Spell id', 'Casts', 'Started', 'Cancelled', 'Failed', 'Casting s'],
      ...ordered.map((row) => [
        casterName(row),
        viaOf(row),
        row.spell_name,
        String(row.spell_id),
        String(row.succeeded),
        String(row.started),
        scaled ? '' : String(cancelled(row)),
        failedKnown(row) ? String(refused(row)) : '',
        (row.cast_time_ms / 1000).toFixed(1),
      ]),
    ];
  }
  /** Spell names two different spell ids share for one caster, shown with the id to tell them apart. */
  const sameName = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const ids = new Map<string, Set<number>>();
    for (const row of rows) {
      const key = `${row.guid}|${row.spell_name}`;
      // eslint-disable-next-line svelte/prefer-svelte-reactivity
      const found = ids.get(key) ?? new Set<number>();
      found.add(row.spell_id);
      ids.set(key, found);
    }
    return new Set([...ids.entries()].filter(([, set]) => set.size > 1).map(([key]) => key));
  });
  const minutes = $derived(durationMs / 60_000);
  const perMinute = (count: number): string => (minutes <= 0 ? '0' : (count / minutes).toFixed(1));

  /**
   * One caster's rhythm across every spell: how many casts, how many a minute, and the
   * longest stretch with none. Only for a single caster, which is what the Source picker
   * gives; over a whole raid the longest gap is meaningless.
   */
  const rhythm = $derived.by(() => {
    const casters = new Set(rows.map(ownerOf));
    if (casters.size !== 1) return null;
    const ticks = rows.flatMap((row) => row.sequence).sort((a, b) => a - b);
    if (ticks.length < 2) return null;
    let gap = { from: startMs, to: ticks[0] };
    for (let i = 1; i < ticks.length; i += 1) {
      if (ticks[i] - ticks[i - 1] > gap.to - gap.from) gap = { from: ticks[i - 1], to: ticks[i] };
    }
    // Not the stretch after the last cast: on a kill that is the boss dying, and on a
    // wipe it is the player dead, neither a gap a rotation can close.
    return { casts: ticks.length, gap };
  });
  /** Where a cast sits along the window, 0 to 1: the Spark's tick. */
  const position = (ms: number): number => (durationMs === 0 ? 0 : (ms - startMs) / durationMs);

  const mark = $derived(approximateMark(scaled));
  /** Cast time is the whole fight's total under every window, never just this one's. */
  const castTimeMark = $derived(wholeFightMark(approximate));

  function castTimeText(ms: number): string {
    return ms === 0 ? 'instant' : clockDuration(ms);
  }

  const icons = new SpellIcons();
  const columns: readonly StatColumn[] = [
    { id: 'casts', label: 'Casts', sortable: true },
    { id: 'perMin', label: 'Per min', sortable: true },
    { id: 'failed', label: 'Failed', sortable: true },
  ];
  function figuresOf(row: CastRow): StatFigure[] {
    return [
      { value: `${mark}${row.succeeded}` },
      { value: perMinute(row.succeeded), testid: 'cast-per-minute' },
      failedKnown(row) ? { value: `${mark}${refused(row)}` } : { value: '–', muted: true },
    ];
  }
  const SEPARATOR = ' · ';
</script>

<Section title="Casts" meta={`Every caster, ${clockDuration(durationMs)}`} testid="cast-section" panel flush>
  {#snippet head()}{#if ordered.length > 0}<CopyCsv lines={csvLines} inHead />{/if}{/snippet}
  <div>
    {#if ordered.length === 0}
      <EmptyState inset message="No casts in this window." testid="table-empty" />
    {:else}
      <StatTable {columns} nameHead="Spell" barHead="When" {sort} {onsort} label="Casts" testid="cast-table">
        {#each ordered as row (`${row.guid}-${row.spell_id}`)}
          {@const caster = casterName(row)}
          {@const via = viaOf(row)}
          {@const colour = classColorVar(classOf.get(ownerOf(row)))}
          <StatRow
            figures={figuresOf(row)}
            phone={`${row.succeeded} ${row.succeeded === 1 ? 'cast' : 'casts'}`}
            testid={`cast-${row.guid}-${row.spell_id}`}
          >
            {#snippet icon()}<SpellIcon
                socket
                url={icons.of(row.spell_id, row.spell_name)}
                tip={row.spell_id}
              />{/snippet}
            {#snippet name()}{row.spell_name}{#if sameName.has(`${row.guid}|${row.spell_name}`)}<span
                  class="text-muted ml-1 text-[13px] font-normal tabular-nums"
                  title="Two spells share this name; this is spell id {row.spell_id}">#{row.spell_id}</span
                >{/if}{/snippet}
            {#snippet sub()}
              {#if via !== ''}<span style:color={colour} data-testid="cast-via">{via} ({caster}’s pet)</span
                >{:else}<span style:color={colour}>{caster}</span
                >{/if}{#if !scaled && (measured !== undefined || cancelled(row) > 0)}{SEPARATOR}<span
                  title="Casts started that never went off: moved, interrupted or cancelled mid-cast"
                  data-testid="cast-cancelled">{cancelled(row)} cancelled</span
                >{:else if scaled}{SEPARATOR}<span
                  title="Cancelled casts are being read from this window’s own cast lines"
                  data-testid="cast-cancelled">– cancelled</span
                >{/if}{#if row.cast_time_ms > 0}{SEPARATOR}<span
                  title="Time spent casting this spell, every cast added together, over the whole fight"
                  >casting {castTimeMark}{castTimeText(row.cast_time_ms)}</span
                >{/if}<span class="ui-st-phone">{SEPARATOR}{perMinute(row.succeeded)} a minute</span>
            {/snippet}
            {#snippet meter()}<Spark marks={row.sequence.map(position)} color={colour} />{/snippet}
          </StatRow>
        {/each}
      </StatTable>
    {/if}
  </div>
  {#if ordered.length > 0}
    <div class="text-muted flex flex-col gap-1 px-5 pt-3 pb-4 text-[13px] leading-[18px]">
      {#if rhythm !== null}
        <p
          data-testid="cast-rhythm"
          title="Casts of every spell in this scope, and the longest stretch between two of them. A gap counts time spent dead or out of range; over the night it is on the night's clock."
        >
          <span class="tabular-nums">{rhythm.casts}</span>
          casts{#if cancelledTotal > 0 && !scaled}{SEPARATOR}<span class="tabular-nums">{cancelledTotal}</span
            > cancelled{/if}{SEPARATOR}<span class="tabular-nums">{perMinute(rhythm.casts)}</span> a minute{SEPARATOR}longest
          gap <span class="tabular-nums">{clockDuration(rhythm.gap.to - rhythm.gap.from)}</span> at
          <span class="tabular-nums">{clockDuration(rhythm.gap.from)}</span>
        </p>
      {/if}
      {#if scaled}
        <p data-testid="cast-approximate-note">
          {#if measureError === ''}
            <span data-testid="cast-measuring">Reading this window’s casts from the fight’s events…</span>
          {:else}
            <span class="text-wipe" role="alert">{measureError}</span>
          {/if}
          Until then Casts and Failed are marked {mark}, the whole fight’s count scaled by the window’s share
          of the sequence, and cancelled casts read a dash. The ticks are this window’s own.
        </p>
      {:else if approximate}
        <p class="text-kill" data-testid="cast-measured-note">
          Casts, Failed and cancelled casts are this window’s own counts, read from the fight’s cast lines.
        </p>
      {/if}
      {#if approximate}<p data-testid="cast-time-note">
          Casting time is marked {castTimeMark} because it is the whole fight's accumulated casting time for that
          spell, even inside a shorter window.
        </p>{/if}
    </div>
  {/if}
</Section>
