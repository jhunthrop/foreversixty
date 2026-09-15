<!-- web/src/components/report/AuraTable.svelte -->
<!-- Buffs and Debuffs. The uptime bar is the aura's own segments drawn to scale, not a
     percentage bar: where the buff dropped is the thing a raid leader is looking for, and
     a single filled bar hides exactly that.

     Every figure here is exact under any window, not scaled: window.ts's
     scopeAuraTrack clips each segment to the window's own bounds and recomputes
     `uptime_ms` and `applications` from those clipped segments, rather than taking the
     whole fight's numbers and scaling them by a ratio the way the per-ability and
     per-target splits are. So this table carries neither the `~` nor the `†` mark. The
     one field on AuraTrack that IS whole-fight -- `max_stacks` -- is not rendered here. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { formatDuration, formatPercent, schoolName, schoolToken } from '../../lib/report/format';
  import type { AuraTrack } from '../../lib/report/types';

  let {
    tracks,
    durationMs,
    kind,
    bossNames = new Set<string>(),
    names = new Map<string, string>(),
  }: {
    tracks: AuraTrack[];
    durationMs: number;
    kind: 'BUFF' | 'DEBUFF';
    /** GUID to unit name, so a debuff can say who applied it. */
    names?: ReadonlyMap<string, string>;
    /** The encounter bosses' unit names, for the per-spell line across the night. */
    bossNames?: ReadonlySet<string>;
  } = $props();

  /**
   * Over the night, one line per debuff across every boss: uptime on the bosses over the
   * time the bosses were up. Only for folded tracks (those carrying their own time), and
   * only for spells on a boss, since adds come and go.
   */
  const bySpell = $derived.by(() => {
    if (kind === 'DEBUFF' && bossNames.size === 0) return [];
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const totals = new Map<string, { name: string; uptime: number; time: number; targets: number }>();
    for (const track of rows) {
      if (track.time_ms === undefined || (kind === 'DEBUFF' && !bossNames.has(track.target_name))) continue;
      const found = totals.get(track.name) ?? { name: track.name, uptime: 0, time: 0, targets: 0 };
      found.uptime += track.uptime_ms;
      found.time += track.time_ms;
      found.targets += 1;
      totals.set(track.name, found);
    }
    return [...totals.values()]
      .filter((entry) => entry.targets > 1)
      .map((entry) => ({ ...entry, pct: entry.time === 0 ? 0 : (entry.uptime / entry.time) * 100 }))
      .sort((a, b) => b.pct - a.pct);
  });

  const rows = $derived(
    tracks
      .filter((track) => track.type === kind)
      .sort((a, b) => shareOf(b) - shareOf(a) || a.name.localeCompare(b.name)),
  );

  const pct = (ms: number): number => (durationMs === 0 ? 0 : (ms / durationMs) * 100);
  /** Uptime over the time the track's target was in: the figure the row shows. */
  const shareOf = (track: AuraTrack): number =>
    track.time_ms === undefined ? pct(track.uptime_ms) : (track.uptime_ms / track.time_ms) * 100;

  /** Names that two different spells share on one target: shown with their spell id. */
  const ambiguous = $derived.by(() => {
    // Plain collections: built once inside the derived and never read reactively by key.
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const seen = new Map<string, Set<number>>();
    for (const track of rows) {
      const key = `${track.target_guid}|${track.name}`;
      // eslint-disable-next-line svelte/prefer-svelte-reactivity
      const ids = seen.get(key) ?? new Set<number>();
      ids.add(track.spell_id);
      seen.set(key, ids);
    }
    return new Set([...seen.entries()].filter(([, ids]) => ids.size > 1).map(([key]) => key));
  });
</script>

{#if rows.length === 0}
  <p class="text-muted text-[14px]" data-testid="table-empty">
    No {kind === 'BUFF' ? 'buffs' : 'debuffs'} in this window.
  </p>
{:else}
  {#if bySpell.length > 0}
    <div class="flex flex-col gap-1" data-testid="aura-by-spell">
      <h2 class="label text-muted">
        {kind === 'DEBUFF' ? 'On the bosses, across the night' : 'Across the night, over everyone who had it'}
      </h2>
      <ul class="flex flex-col">
        {#each bySpell as entry (entry.name)}
          <li
            class="border-line-soft grid min-h-9 grid-cols-[minmax(120px,1.2fr)_minmax(0,3fr)_80px] items-center gap-x-3 border-b px-2 py-1 text-[14px]"
          >
            <span class="truncate font-semibold">{entry.name}</span>
            <span class="bg-line-soft block h-[6px] w-full"
              ><span class="bg-wipe block h-full" style={`width: ${entry.pct}%`}></span></span
            >
            <span
              class="tabular text-right font-mono"
              title={`Uptime on ${entry.targets} bosses over the time they were up`}
              >{formatPercent(entry.pct)}</span
            >
          </li>
        {/each}
      </ul>
    </div>
  {/if}
  <div class="flex flex-col" data-testid="aura-table">
    <div
      class="text-muted label hidden grid-cols-[minmax(120px,1.2fr)_minmax(120px,1.2fr)_minmax(0,3fr)_80px_64px] gap-x-3 px-2 pb-1 md:grid"
    >
      <span>Aura</span>
      <span>On</span>
      <span title="When the aura was up, drawn on the fight's timeline. A gap is where it dropped."
        >Uptime</span
      >
      <span class="text-right" title="Share of the window the aura was up">Total</span>
      <span
        class="text-right"
        title="How many times it went up when it was not already up; a refresh while up does not count"
        >Applied</span
      >
    </div>
    <ul class="flex flex-col">
      {#each rows as track (`${track.target_guid}-${track.spell_id}`)}
        <li
          class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] md:grid-cols-[minmax(120px,1.2fr)_minmax(120px,1.2fr)_minmax(0,3fr)_80px_64px]"
          data-testid={`aura-${track.spell_id}-${track.target_guid}`}
        >
          <span class="flex min-w-0 items-center gap-1">
            <span class="truncate font-semibold"
              >{track.name}{#if ambiguous.has(`${track.target_guid}|${track.name}`)}
                <span
                  class="text-muted ml-1 font-mono text-[11px]"
                  title="Two spells share this name; this is spell id {track.spell_id}"
                  >#{track.spell_id}</span
                >{/if}</span
            >
            {#if kind === 'DEBUFF' && schoolName(track.school)}
              <span
                class="shrink-0 rounded-[2px] px-1 text-[10px] font-normal tracking-[0.04em] uppercase"
                style={`background: ${schoolToken(track.school)}; color: var(--color-bg)`}
                title="The spell's school. Whether a debuff can be dispelled is not in the log; the Dispels tab shows which ones were."
                data-testid="aura-school">{schoolName(track.school)}</span
              >
            {/if}
          </span>
          <span class="text-muted truncate text-[13px]"
            >{splitUnitName(track.target_name).name}{#if kind === 'DEBUFF' && track.appliers.length > 0}
              <span class="text-[11px]" title="Who applied it">
                · from {track.appliers
                  .map((guid) => splitUnitName(names.get(guid) ?? '').name)
                  .filter(Boolean)
                  .slice(0, 2)
                  .join(', ') || 'an unnamed source'}</span
              >{/if}</span
          >
          <span class="bg-line-soft relative col-span-2 block h-[6px] w-full md:col-span-1">
            <!-- Keyed by index as well as start: a stack refreshed on the tick it was
                 applied gives two segments the same start_ms, and a bare timestamp key
                 threw on the duplicate and left the whole tab on "Loading the report." -->
            {#if track.time_ms !== undefined}
              <!-- Over the night the number is uptime over the target's own time in combat,
                   so the bar is that share, not the segments laid on the whole night. -->
              <span
                class="absolute top-0 h-full {kind === 'BUFF' ? 'bg-kill' : 'bg-wipe'}"
                style={`width: ${Math.min(shareOf(track), 100)}%`}
                title={`Up ${formatDuration(track.uptime_ms)} of the ${formatDuration(track.time_ms)} this unit was in`}
              ></span>
            {:else}
              {#each track.segments as segment, i (`${segment.start_ms}-${i}`)}
                <span
                  class="absolute top-0 h-full {kind === 'BUFF' ? 'bg-kill' : 'bg-wipe'}"
                  style={`left: ${pct(segment.start_ms)}%; width: ${Math.max(pct(segment.end_ms - segment.start_ms), 0.4)}%; opacity: ${Math.min(0.4 + segment.stacks * 0.2, 1)}`}
                  title={`${formatDuration(segment.start_ms)} to ${formatDuration(segment.end_ms)}${segment.stacks > 1 ? ` · ${segment.stacks} stacks` : ''}`}
                ></span>
              {/each}
            {/if}
          </span>
          <!-- The heading row above is `hidden` below `md`, so each figure carries the word
               it was filed under. The testid stays on the number alone: it is the figure
               report-tables.spec.ts pins, not the word beside it. -->
          <span class="tabular text-right font-mono">
            <span data-testid="aura-uptime">{formatPercent(shareOf(track))}</span><span
              class="label font-body text-muted ml-1.5 md:hidden">uptime</span
            >
          </span>
          <span class="text-muted tabular text-right font-mono text-[13px]"
            >{track.applications}<span class="label font-body ml-1.5 md:hidden">applied</span></span
          >
        </li>
      {/each}
    </ul>
  </div>
{/if}
