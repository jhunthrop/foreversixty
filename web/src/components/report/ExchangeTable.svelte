<!-- web/src/components/report/ExchangeTable.svelte -->
<!-- Interrupts and Dispels are the same shape: someone did something to someone else's
     spell. One component, two tabs.

     summary.ExchangeRow carries no timestamp, so nothing in it can be cut to a brushed
     window: scopeSummary (window.ts) passes `interrupts` and `dispels` straight through,
     untouched. These counts are the whole fight's under every window, and this table says
     so outright rather than letting an unchanged count after a brush read as "nothing
     happened in this window" -- a mark and a note a reader could miss beat a raid leader
     concluding the wrong thing from a number that never moved. The `†` mark and the note
     below are unconditional: unlike CastTable's approximate marks, nothing here depends
     on whether the window happens to be whole. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { wholeFightAriaLabel, wholeFightMark, wholeFightTitle } from '../../lib/report/format';
  import type { AuraTrack, CastRow, ExchangeRow } from '../../lib/report/types';
  import CopyCsv from './CopyCsv.svelte';

  let {
    rows,
    emptyText,
    casts = [],
    everyone = rows,
    auras = [],
    players = new Set<string>(),
  }: {
    rows: ExchangeRow[];
    emptyText: string;
    /** The fight's cast rows, for the "went through" count on the Interrupts tab. */
    casts?: CastRow[];
    /**
     * Every source's rows, unscoped: what went through or ran its course is counted
     * against the whole group's stops, whoever the table is scoped to.
     */
    everyone?: ExchangeRow[];
    /** The fight's aura tracks, for the "ran its course" count on the Dispels tab. */
    auras?: AuraTrack[];
    players?: ReadonlySet<string>;
  } = $props();

  interface Uncured {
    spell_id: number;
    name: string;
    applied: number;
    dispelled: number;
    /** Dispels by the sources this table is scoped to, when that is not everyone. */
    own: number;
  }

  /**
   * For every debuff somebody dispelled at least once: how many times it landed on a
   * player against how many times it was dispelled. The rest ran its full course.
   */
  const uncured = $derived.by<Uncured[]>(() => {
    if (auras.length === 0) return [];
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const bySpell = new Map<number, Uncured>();
    for (const row of everyone) {
      if (row.kind !== 'dispel') continue;
      const found = bySpell.get(row.extra_spell_id) ?? {
        spell_id: row.extra_spell_id,
        name: row.extra_spell_name,
        applied: 0,
        dispelled: 0,
        own: 0,
      };
      found.dispelled += row.count;
      bySpell.set(row.extra_spell_id, found);
    }
    for (const row of rows) {
      const found = bySpell.get(row.extra_spell_id);
      if (found !== undefined && row.kind === 'dispel') found.own += row.count;
    }
    for (const track of auras) {
      if (track.type !== 'DEBUFF' || !players.has(track.target_guid)) continue;
      const found = bySpell.get(track.spell_id);
      if (found !== undefined) found.applied += track.applications;
    }
    return [...bySpell.values()]
      .map((entry) => ({ ...entry, applied: Math.max(entry.applied, entry.dispelled) }))
      .filter((entry) => entry.applied > entry.dispelled)
      .sort((a, b) => b.applied - b.dispelled - (a.applied - a.dispelled));
  });

  interface Missed {
    spell_id: number;
    name: string;
    cast: number;
    stopped: number;
    /** Stops by the sources this table is scoped to, when that is not everyone. */
    own: number;
  }

  /**
   * For every enemy spell somebody interrupted at least once: how many times an enemy
   * cast it against how many times it was stopped. A spell nobody ever interrupted is not
   * listed, because the summary cannot tell an uninterruptible cast from an ignored one.
   */
  const missed = $derived.by<Missed[]>(() => {
    if (casts.length === 0) return [];
    // A plain Map: built once inside the derived, never read reactively by key.
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const stopped = new Map<number, Missed>();
    for (const row of everyone) {
      if (row.kind !== 'interrupt') continue;
      const found = stopped.get(row.extra_spell_id) ?? {
        spell_id: row.extra_spell_id,
        name: row.extra_spell_name,
        cast: 0,
        stopped: 0,
        own: 0,
      };
      found.stopped += row.count;
      stopped.set(row.extra_spell_id, found);
    }
    for (const row of rows) {
      const found = stopped.get(row.extra_spell_id);
      if (found !== undefined && row.kind === 'interrupt') found.own += row.count;
    }
    for (const cast of casts) {
      if (players.has(cast.guid)) continue;
      const found = stopped.get(cast.spell_id);
      if (found !== undefined) found.cast += Math.max(cast.started, cast.succeeded);
    }
    return [...stopped.values()]
      .map((entry) => ({ ...entry, cast: Math.max(entry.cast, entry.stopped) }))
      .filter((entry) => entry.cast > entry.stopped)
      .sort((a, b) => b.cast - b.stopped - (a.cast - a.stopped));
  });

  /** Whether the table shows some sources only, so the group's counts name the scope's part. */
  const scopedToSome = $derived(everyone !== rows && everyone.length !== rows.length);
  const ordered = $derived(
    [...rows].sort((a, b) => b.count - a.count || a.source_name.localeCompare(b.source_name)),
  );
  /** The table as lines: who did it, with what, on whom, which spell, how often. */
  function csvLines(): string[][] {
    const credit = [
      ['By', 'With', 'On', 'Spell', 'Spell id', 'Count'],
      ...ordered.map((row) => [
        splitUnitName(row.source_name).name,
        row.spell_name,
        splitUnitName(row.target_name).name,
        row.extra_spell_name,
        String(row.extra_spell_id),
        String(row.count),
      ]),
    ];
    // The denominators under the table, as a second block: what went through or ran its
    // course is the half a roster is built from.
    const through = missed.map((entry) => [
      'Went through',
      entry.name,
      String(entry.spell_id),
      String(entry.cast),
      String(entry.stopped),
      ...(scopedToSome ? [String(entry.own)] : []),
      String(entry.cast - entry.stopped),
    ]);
    const course = uncured.map((entry) => [
      'Ran their course',
      entry.name,
      String(entry.spell_id),
      String(entry.applied),
      String(entry.dispelled),
      ...(scopedToSome ? [String(entry.own)] : []),
      String(entry.applied - entry.dispelled),
    ]);
    if (through.length === 0 && course.length === 0) return credit;
    return [
      ...credit,
      [],
      [
        'Block',
        'Spell',
        'Spell id',
        'Cast or landed',
        'Stopped or dispelled',
        ...(scopedToSome ? ['In this scope'] : []),
        'Went through or ran their course',
      ],
      ...through,
      ...course,
    ];
  }
  const mark = wholeFightMark(true);
  const title = wholeFightTitle(true);

  /** One line per source, most first: who did the interrupting or dispelling tonight. */
  const bySource = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const totals = new Map<string, { name: string; count: number }>();
    for (const row of rows) {
      const found = totals.get(row.source_guid);
      if (found === undefined) totals.set(row.source_guid, { name: row.source_name, count: row.count });
      else found.count += row.count;
    }
    return [...totals.values()].sort((a, b) => b.count - a.count || a.name.localeCompare(b.name));
  });
</script>

{#if ordered.length === 0}
  <p class="text-muted text-[14px]" data-testid="table-empty">{emptyText}</p>
{:else}
  <p class="flex flex-wrap gap-x-4 gap-y-1 text-[13px]" data-testid="exchange-by-source">
    {#each bySource as entry (entry.name)}
      <span
        ><span class="font-semibold">{splitUnitName(entry.name).name}</span>
        <span class="tabular font-mono">{mark}{entry.count}</span></span
      >
    {/each}
  </p>
  <div class="flex flex-col" data-testid="exchange-table">
    <div
      class="text-muted label hidden grid-cols-[minmax(120px,1fr)_minmax(120px,1fr)_minmax(120px,1fr)_minmax(120px,1fr)_64px] gap-x-3 px-2 pb-1 md:grid"
    >
      <span>By</span>
      <span>With</span>
      <span>On</span>
      <span>Spell</span>
      <span class="text-right">Count</span>
    </div>
    <ul class="flex flex-col">
      {#each ordered as row (`${row.source_guid}-${row.target_guid}-${row.spell_id}-${row.extra_spell_id}`)}
        <li
          class="border-line-soft grid min-h-11 grid-cols-1 items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] md:grid-cols-[minmax(120px,1fr)_minmax(120px,1fr)_minmax(120px,1fr)_minmax(120px,1fr)_64px]"
          id={`exchange-${row.source_guid}-${row.spell_id}`}
        >
          <span class="truncate font-semibold">{splitUnitName(row.source_name).name}</span>
          <span class="truncate">{row.spell_name}</span>
          <span class="text-muted truncate text-[13px]">{splitUnitName(row.target_name).name}</span>
          <span class="text-muted truncate text-[13px]">{row.extra_spell_name}</span>
          <span
            class="tabular text-right font-mono"
            {title}
            aria-label={wholeFightAriaLabel(true, String(row.count))}
          >
            {mark}{row.count}
          </span>
        </li>
      {/each}
    </ul>
    <CopyCsv lines={csvLines} />
  </div>
  {#if missed.length > 0}
    <div class="flex flex-col gap-1" data-testid="interrupts-missed">
      <h2 class="label text-muted">Went through</h2>
      <ul class="flex flex-col">
        {#each missed as entry (entry.spell_id)}
          <li
            class="border-line-soft flex min-h-11 flex-wrap items-center gap-x-3 border-b px-2 py-2 text-[14px]"
          >
            <span class="font-semibold">{entry.name}</span>
            <span class="text-muted text-[13px]">
              cast <span class="tabular font-mono">{entry.cast}</span> · stopped
              <span class="tabular font-mono">{entry.stopped}</span>{#if scopedToSome}
                <span class="text-muted"
                  >&nbsp;(<span class="tabular font-mono">{entry.own}</span> in this scope)</span
                >{/if} ·
              <span class="text-wipe tabular font-mono">{entry.cast - entry.stopped}</span> went through
            </span>
          </li>
        {/each}
      </ul>
    </div>
  {/if}
  {#if uncured.length > 0}
    <div class="flex flex-col gap-1" data-testid="dispels-uncured">
      <h2 class="label text-muted">Ran their course</h2>
      <ul class="flex flex-col">
        {#each uncured as entry (entry.spell_id)}
          <li
            class="border-line-soft flex min-h-11 flex-wrap items-center gap-x-3 border-b px-2 py-2 text-[14px]"
          >
            <span class="font-semibold">{entry.name}</span>
            <span class="text-muted text-[13px]">
              landed on the raid <span class="tabular font-mono">{entry.applied}</span> times · dispelled
              <span class="tabular font-mono">{entry.dispelled}</span>{#if scopedToSome}
                <span class="text-muted"
                  >&nbsp;(<span class="tabular font-mono">{entry.own}</span> in this scope)</span
                >{/if} ·
              <span class="text-wipe tabular font-mono">{entry.applied - entry.dispelled}</span> ran their course
            </span>
          </li>
        {/each}
      </ul>
    </div>
  {/if}
  <p class="text-muted text-[12px]" data-testid="exchange-wholefight-note">
    Count is marked {mark}: interrupts and dispels are the whole fight's totals, because the summary keeps no
    timestamp for them and a brushed window cannot cut them down.
  </p>
{/if}
