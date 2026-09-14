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
  import { formatDuration, formatPercent } from '../../lib/report/format';
  import type { AuraTrack } from '../../lib/report/types';

  let {
    tracks,
    durationMs,
    kind,
  }: { tracks: AuraTrack[]; durationMs: number; kind: 'BUFF' | 'DEBUFF' } = $props();

  const rows = $derived(
    tracks
      .filter((track) => track.type === kind)
      .sort((a, b) => b.uptime_ms - a.uptime_ms || a.name.localeCompare(b.name)),
  );

  const pct = (ms: number): number => (durationMs === 0 ? 0 : (ms / durationMs) * 100);
</script>

{#if rows.length === 0}
  <p class="text-muted text-[14px]" data-testid="table-empty">
    No {kind === 'BUFF' ? 'buffs' : 'debuffs'} in this window.
  </p>
{:else}
  <div class="flex flex-col" data-testid="aura-table">
    <div class="text-muted label hidden grid-cols-[minmax(120px,1.2fr)_minmax(120px,1.2fr)_minmax(0,3fr)_80px_64px] gap-x-3 px-2 pb-1 md:grid">
      <span>Aura</span>
      <span>On</span>
      <span>Uptime</span>
      <span class="text-right">Total</span>
      <span class="text-right">Applied</span>
    </div>
    <ul class="flex flex-col">
      {#each rows as track (`${track.target_guid}-${track.spell_id}`)}
        <li
          class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] md:grid-cols-[minmax(120px,1.2fr)_minmax(120px,1.2fr)_minmax(0,3fr)_80px_64px]"
          data-testid={`aura-${track.spell_id}-${track.target_guid}`}
        >
          <span class="truncate font-semibold">{track.name}</span>
          <span class="text-muted truncate text-[13px]">{splitUnitName(track.target_name).name}</span>
          <span class="bg-line-soft relative col-span-2 block h-[6px] w-full md:col-span-1">
            {#each track.segments as segment (segment.start_ms)}
              <span
                class="bg-gold absolute top-0 h-full"
                style={`left: ${pct(segment.start_ms)}%; width: ${Math.max(pct(segment.end_ms - segment.start_ms), 0.4)}%; opacity: ${Math.min(0.4 + segment.stacks * 0.2, 1)}`}
                title={`${formatDuration(segment.start_ms)} to ${formatDuration(segment.end_ms)}${segment.stacks > 1 ? ` · ${segment.stacks} stacks` : ''}`}
              ></span>
            {/each}
          </span>
          <!-- The heading row above is `hidden` below `md`, so each figure carries the word
               it was filed under. The testid stays on the number alone: it is the figure
               report-tables.spec.ts pins, not the word beside it. -->
          <span class="font-mono tabular text-right">
            <span data-testid="aura-uptime">{formatPercent(pct(track.uptime_ms))}</span><span
              class="label font-body text-muted ml-1.5 md:hidden">uptime</span
            >
          </span>
          <span class="text-muted font-mono tabular text-right text-[13px]"
            >{track.applications}<span class="label font-body ml-1.5 md:hidden">applied</span></span
          >
        </li>
      {/each}
    </ul>
  </div>
{/if}
