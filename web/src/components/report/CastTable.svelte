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

  let {
    rows,
    durationMs,
    startMs = 0,
    classOf = new Map<string, string>(),
    approximate = false,
  }: {
    rows: CastRow[];
    durationMs: number;
    startMs?: number;
    classOf?: Map<string, string>;
    /** True when the window is brushed, so Cast and Failed are a scaled share. */
    approximate?: boolean;
  } = $props();

  const ordered = $derived(
    [...rows].sort((a, b) => b.succeeded - a.succeeded || a.spell_name.localeCompare(b.spell_name)),
  );
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
    const casters = new Set(rows.map((row) => row.guid));
    if (casters.size !== 1) return null;
    const ticks = rows.flatMap((row) => row.sequence).sort((a, b) => a - b);
    if (ticks.length < 2) return null;
    let gap = { from: startMs, to: ticks[0] };
    for (let i = 1; i < ticks.length; i += 1) {
      if (ticks[i] - ticks[i - 1] > gap.to - gap.from) gap = { from: ticks[i - 1], to: ticks[i] };
    }
    const end = startMs + durationMs;
    if (end - ticks[ticks.length - 1] > gap.to - gap.from) gap = { from: ticks[ticks.length - 1], to: end };
    return { casts: ticks.length, gap };
  });
  const pct = (ms: number): number => (durationMs === 0 ? 0 : ((ms - startMs) / durationMs) * 100);

  const mark = $derived(approximateMark(approximate));
  const title = $derived(approximateTitle(approximate));
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
      <p class="text-[13px]" data-testid="cast-rhythm">
        <span class="tabular font-mono">{rhythm.casts}</span> casts ·
        <span class="tabular font-mono">{perMinute(rhythm.casts)}</span> a minute · longest gap
        <span class="tabular font-mono">{formatDuration(rhythm.gap.to - rhythm.gap.from)}</span> at
        <span class="tabular font-mono">{formatDuration(rhythm.gap.from)}</span>
      </p>
    {/if}
    <div
      class="text-muted label hidden grid-cols-[minmax(120px,1.2fr)_minmax(120px,1.2fr)_64px_64px_64px_80px_minmax(0,3fr)] gap-x-3 px-2 pb-1 md:grid"
    >
      <span>Caster</span>
      <span>Spell</span>
      <span class="text-right">Cast</span>
      <span class="text-right" title="Successful casts per minute of this window">Per min</span>
      <span class="text-right">Failed</span>
      <span class="text-right">Cast time</span>
      <span>Sequence</span>
    </div>
    <ul class="flex flex-col">
      {#each ordered as row (`${row.guid}-${row.spell_id}`)}
        <li
          class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] md:grid-cols-[minmax(120px,1.2fr)_minmax(120px,1.2fr)_64px_64px_64px_80px_minmax(0,3fr)]"
          data-testid={`cast-${row.guid}-${row.spell_id}`}
        >
          <span class="truncate font-semibold" style={`color: ${classColorVar(classOf.get(row.guid))}`}>
            {splitUnitName(row.name).name}
          </span>
          <span class="truncate"
            >{row.spell_name}{#if sameName.has(`${row.guid}|${row.spell_name}`)}
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
            aria-label={approximateAriaLabel(approximate, `${row.succeeded} cast`)}
          >
            {mark}{row.succeeded}<span class="label font-body ml-1.5 md:hidden">cast</span>
          </span>
          <span class="tabular text-muted text-right font-mono text-[13px]" data-testid="cast-per-minute">
            {perMinute(row.succeeded)}<span class="label font-body ml-1.5 md:hidden">a minute</span>
          </span>
          <span
            class="tabular text-muted text-right font-mono"
            {title}
            aria-label={approximateAriaLabel(approximate, `${row.failed} failed`)}
          >
            {mark}{row.failed}<span class="label font-body ml-1.5 md:hidden">failed</span>
          </span>
          <span
            class="tabular text-muted text-right font-mono text-[13px]"
            title={castTimeTitle}
            aria-label={wholeFightAriaLabel(true, `${castTimeText(row.cast_time_ms)} cast time`)}
          >
            {castTimeMark}{castTimeText(row.cast_time_ms)}<span class="label font-body ml-1.5 md:hidden"
              >cast time</span
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
  </div>
  {#if approximate}
    <p class="text-muted text-[12px]" data-testid="cast-approximate-note">
      Cast and Failed are marked {mark} because the summary keeps only a whole-fight count for each: this window's
      figure is that count scaled by the window's share of the sequence, not measured directly. The sequence above
      is this window's own ticks.
    </p>
  {/if}
  <p class="text-muted text-[12px]" data-testid="cast-time-note">
    Cast time is marked {castTimeMark} because it is the whole fight's accumulated casting time for that spell,
    even inside a shorter window.
  </p>
{/if}
