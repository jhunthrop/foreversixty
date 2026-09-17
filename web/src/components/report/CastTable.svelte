<!-- web/src/components/report/CastTable.svelte -->
<!-- Casts: counts, cast time, and the sequence as a timeline of ticks, which is the only
     way to see a rotation's shape without opening the events file.

     Three different truths sit in one row. The sequence is exact under any window --
     window.ts's scopeCastRow filters the millisecond offsets themselves, rather than
     scaling anything. Cast and Failed are not: the engine keeps only a whole-fight count,
     so scopeCastRow scales both by the window's share of the sequence, and they carry the
     `~` mark ActorRow uses for its per-ability and per-target splits. Cast time is a third
     kind of lie: logs/engine/summary/deaths.go accumulates it across every successful
     cast in the whole fight and scopeCastRow never touches it, so under any window it is
     silently the whole fight's total -- the `†` mark SummaryTab's Active column uses. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import {
    approximateAriaLabel,
    approximateMark,
    approximateTitle,
    classColorVar,
    formatDuration,
    wholeFightAriaLabel,
    wholeFightMark,
    wholeFightTitle,
  } from '../../lib/report/format';
  import type { CastRow } from '../../lib/report/types';
  import { castKey, type CastCounts } from '../../lib/report/exact';
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
  /** Cancelled casts over the rows on screen: the total the rhythm line adds up. */
  const cancelledTotal = $derived(ordered.reduce((sum, row) => sum + cancelled(row), 0));
  const failedKnown = (row: CastRow): boolean => recorders.size === 0 || recorders.has(row.guid);
  const ordered = $derived(
    [...shown].sort((a, b) => b.succeeded - a.succeeded || a.spell_name.localeCompare(b.spell_name)),
  );
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
  const pct = (ms: number): number => (durationMs === 0 ? 0 : ((ms - startMs) / durationMs) * 100);

  const mark = $derived(approximateMark(scaled));
  const title = $derived(approximateTitle(scaled));
  /** Cast time is the whole fight's total under every window, never just this one's. */
  const castTimeMark = wholeFightMark(true);
  const castTimeTitle = wholeFightTitle(true);

  function castTimeText(ms: number): string {
    return ms === 0 ? 'instant' : formatDuration(ms);
  }
</script>

{#if ordered.length === 0}
  <p class="text-muted text-[14px]" data-testid="table-empty">No casts in this window.</p>
{:else}
  <div class="flex flex-col" data-testid="cast-table">
    {#if rhythm !== null}
      <p
        class="text-[13px]"
        data-testid="cast-rhythm"
        title="Casts of every spell in this scope, and the longest stretch between two of them. A gap counts time spent dead or out of range; over the night it is on the night's clock."
      >
        <span class="tabular font-mono">{rhythm.casts}</span> casts{#if cancelledTotal > 0 && !scaled}
          · <span class="tabular font-mono">{cancelledTotal}</span> cancelled{/if} ·
        <span class="tabular font-mono">{perMinute(rhythm.casts)}</span> a minute · longest gap
        <span class="tabular font-mono">{formatDuration(rhythm.gap.to - rhythm.gap.from)}</span> at
        <span class="tabular font-mono">{formatDuration(rhythm.gap.from)}</span>
      </p>
    {/if}
    <div
      class="text-muted label hidden grid-cols-[minmax(150px,1.3fr)_minmax(170px,1.5fr)_64px_64px_64px_72px_80px_minmax(0,2fr)] gap-x-3 px-2 pb-1 md:grid"
    >
      <span>Caster</span>
      <span>Spell</span>
      <span class="text-right">Cast</span>
      <span class="text-right" title="Successful casts per minute of this window">Per min</span>
      <span
        class="text-right"
        title="Presses the game refused: out of range, no target, not enough mana. Only the player who wrote this log has them recorded"
        >Failed</span
      >
      <span
        class="text-right"
        title="Casts started that never went off: moved, interrupted or cancelled mid-cast">Cancelled</span
      >
      <span
        class="text-right"
        title="Time spent casting this spell, every cast added together; instants show none"
        >Casting, total</span
      >
      <span>Sequence</span>
    </div>
    <ul class="flex flex-col">
      {#each ordered as row (`${row.guid}-${row.spell_id}`)}
        {@const caster = casterName(row)}
        {@const via = viaOf(row)}
        <li
          class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] md:grid-cols-[minmax(150px,1.3fr)_minmax(170px,1.5fr)_64px_64px_64px_72px_80px_minmax(0,2fr)]"
          data-testid={`cast-${row.guid}-${row.spell_id}`}
        >
          <span
            class="truncate font-semibold"
            style={`color: ${classColorVar(classOf.get(ownerOf(row)))}`}
            title={caster}
          >
            {caster}
          </span>
          <span class="truncate" title={row.spell_name}
            >{row.spell_name}{#if via !== ''}
              <span
                class="text-muted ml-1 text-[11px]"
                title="Cast by this pet or guardian, counted on its owner's row"
                data-testid="cast-via">· via {via}</span
              >{/if}{#if sameName.has(`${row.guid}|${row.spell_name}`)}
              <span
                class="text-muted ml-1 font-mono text-[11px]"
                title="Two spells share this name; this is spell id {row.spell_id}">#{row.spell_id}</span
              >{/if}</span
          >
          <!-- Below `md` the heading row is hidden and these three are a card's middle
               lines, so each says what it is. The word goes into the accessible name as
               well as onto the screen: an aria-label replaces an element's text outright,
               so a visible word alone would reach a sighted reader and nobody else -- and
               these cells are divs, with no column header for a screen reader to associate
               them with at any width. -->
          <span
            class="tabular text-right font-mono"
            {title}
            aria-label={approximateAriaLabel(scaled, `${row.succeeded} cast`)}
          >
            {mark}{row.succeeded}<span class="label font-body ml-1.5 md:hidden">cast</span>
          </span>
          <span class="tabular text-muted text-right font-mono text-[13px]" data-testid="cast-per-minute">
            {perMinute(row.succeeded)}<span class="label font-body ml-1.5 md:hidden">a minute</span>
          </span>
          <span
            class="tabular text-muted text-right font-mono"
            {title}
            aria-label={approximateAriaLabel(
              scaled,
              failedKnown(row) ? `${refused(row)} failed` : 'failed casts not in this log for this player',
            )}
          >
            {#if failedKnown(row)}{mark}{refused(row)}{:else}<span
                title="Only the player whose client wrote this log has failed casts in it; this player's are not recorded"
                >—</span
              >{/if}<span class="label font-body ml-1.5 md:hidden">failed</span>
          </span>
          <!-- The summary keeps whole-fight cast starts with no instants, so until the window's
               own cast lines are read it cannot say which casts inside it were cut short; a
               prorated count named the wrong spells, and a dash names nothing wrongly. -->
          <span
            class="tabular text-muted text-right font-mono"
            title={scaled
              ? 'Cancelled casts are being read from this window’s own cast lines'
              : 'Casts started that never went off: moved, interrupted or cancelled mid-cast; instants have none'}
            aria-label={scaled
              ? 'cancelled casts not yet measured for this window'
              : `${cancelled(row)} cancelled`}
            data-testid="cast-cancelled"
          >
            {scaled ? '—' : cancelled(row)}<span class="label font-body ml-1.5 md:hidden">cancelled</span>
          </span>
          <span
            class="tabular text-muted text-right font-mono text-[13px]"
            title={castTimeTitle}
            aria-label={wholeFightAriaLabel(true, `${castTimeText(row.cast_time_ms)} cast time`)}
          >
            {castTimeMark}{castTimeText(row.cast_time_ms)}<span class="label font-body ml-1.5 md:hidden"
              >casting, total</span
            >
          </span>
          <span class="bg-line-soft relative col-span-2 block h-[6px] w-full md:col-span-1">
            {#each row.sequence as at, i (`${at}-${i}`)}
              <span
                class="bg-gold absolute top-0 h-full w-[2px]"
                style={`left: ${pct(at)}%`}
                title={formatDuration(at)}
              ></span>
            {/each}
          </span>
        </li>
      {/each}
    </ul>
    <CopyCsv lines={csvLines} />
  </div>
  {#if scaled}
    <p class="text-muted text-[12px]" data-testid="cast-approximate-note">
      {#if measureError === ''}
        <span data-testid="cast-measuring">Reading this window’s casts from the fight’s events…</span>
      {:else}
        <span class="text-wipe" role="alert">{measureError}</span>
      {/if}
      Until then Cast and Failed are marked {mark}, the whole fight’s count scaled by the window’s share of
      the sequence, and Cancelled reads a dash. The sequence above is this window’s own ticks.
    </p>
  {:else if approximate}
    <p class="text-kill text-[12px]" data-testid="cast-measured-note">
      Cast, Failed and Cancelled are this window’s own counts, read from the fight’s cast lines.
    </p>
  {/if}
  <p class="text-muted text-[12px]" data-testid="cast-time-note">
    Cast time is marked {castTimeMark} because it is the whole fight's accumulated casting time for that spell,
    even inside a shorter window. Failed and cancelled casts are only known from the client that wrote this log;
    other players' rows read a dash where the log is silent.
  </p>
{/if}
