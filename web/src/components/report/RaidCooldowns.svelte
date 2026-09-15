<!-- web/src/components/report/RaidCooldowns.svelte -->
<!-- The raid cooldowns used this fight, with when: the first thing a raid leader reads
     on the Buffs tab. One lane per cooldown, a tick per use, the caster named on hover,
     a time axis under the lanes and every death drawn across them, so "Ardent Defender
     at 48s, tank dead at 1:04" is one picture. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { formatDuration } from '../../lib/report/format';
  import { isRaidCooldown } from '../../lib/report/raid-cooldowns';
  import type { AuraTrack, Death } from '../../lib/report/types';
  import type { TimeWindow } from '../../lib/report/window';

  let {
    tracks,
    window: timeWindow,
    deaths = [],
    names,
  }: { tracks: AuraTrack[]; window: TimeWindow; deaths?: Death[]; names: Map<string, string> } = $props();

  interface Use {
    at: number;
    end: number;
    source: string;
    target: string;
  }

  /** One row per cooldown, its uses in order, with who cast it and on whom. */
  const rows = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const byName = new Map<string, Use[]>();
    for (const track of tracks) {
      if (track.type !== 'BUFF' || !isRaidCooldown(track.name)) continue;
      const uses = byName.get(track.name) ?? [];
      for (const segment of track.segments) {
        uses.push({
          at: segment.start_ms,
          end: segment.end_ms,
          source:
            names.get(segment.source_guid ?? '') ??
            track.appliers.map((guid) => names.get(guid) ?? '').find(Boolean) ??
            '',
          target: track.target_name,
        });
      }
      byName.set(track.name, uses);
    }
    return (
      [...byName.entries()]
        .map(([name, uses]) => ({ name, uses: dedupe(uses).sort((a, b) => a.at - b.at) }))
        // A track with no segment in this window has no use to draw.
        .filter((row) => row.uses.length > 0)
        .sort((a, b) => a.uses[0].at - b.uses[0].at)
    );
  });

  const span = $derived(Math.max(timeWindow.endMs - timeWindow.startMs, 1));
  /** Where an instant sits across the lane, 0..100, from the window's start. */
  const pct = (ms: number): number => ((ms - timeWindow.startMs) / span) * 100;

  /** Axis ticks every 10s, 30s or minute, whichever keeps them under about eight. */
  const axis = $derived.by(() => {
    const step = span > 240_000 ? 60_000 : span > 80_000 ? 30_000 : 10_000;
    const first = Math.ceil(timeWindow.startMs / step) * step;
    const ticks: number[] = [];
    for (let at = first; at <= timeWindow.endMs; at += step) ticks.push(at);
    return ticks;
  });

  const shownDeaths = $derived(
    deaths.filter((death) => death.at_ms >= timeWindow.startMs && death.at_ms <= timeWindow.endMs),
  );

  /** A raid-wide buff lands on everyone in the same second; that is one use, on the raid. */
  function dedupe(uses: Use[]): Use[] {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const seen = new Map<string, Use>();
    for (const use of uses) {
      const key = `${use.source}|${Math.round(use.at / 1000)}`;
      const found = seen.get(key);
      if (found === undefined) seen.set(key, { ...use });
      else if (found.target !== use.target) found.target = 'the raid';
    }
    return [...seen.values()];
  }
</script>

{#if rows.length > 0}
  <section
    class="border-line rounded-panel bg-raised flex flex-col gap-2 border p-3"
    data-testid="raid-cooldowns"
  >
    <h2 class="label text-muted flex flex-wrap items-center gap-x-3">
      Raid cooldowns
      {#if shownDeaths.length > 0}
        <span class="text-death tracking-normal normal-case" data-testid="raid-cooldowns-deaths"
          >| {shownDeaths.length} {shownDeaths.length === 1 ? 'death' : 'deaths'} marked</span
        >
      {/if}
    </h2>
    <ul class="flex flex-col">
      {#each rows as row (row.name)}
        <li
          class="border-line-soft grid min-h-9 grid-cols-[minmax(110px,160px)_40px_minmax(0,1fr)] items-center gap-3 border-b py-1 text-[13px]"
          data-testid={`raid-cooldown-${row.name}`}
        >
          <span class="truncate font-semibold">{row.name}</span>
          <span class="text-muted tabular text-right font-mono text-[12px]" title="Times used"
            >{row.uses.length}</span
          >
          <span class="bg-line-soft relative block h-[14px] w-full">
            {#each row.uses as use, i (`${use.at}-${i}`)}
              <span
                class="bg-gold absolute top-0 h-full opacity-80"
                style={`left: ${pct(use.at)}%; width: ${Math.max(pct(use.end) - pct(use.at), 0.4)}%`}
                title={`${formatDuration(use.at)}${use.source ? ` · ${splitUnitName(use.source).name}` : ''} on ${splitUnitName(use.target).name}`}
              ></span>
            {/each}
            {#each shownDeaths as death (`${death.guid}-${death.at_ms}`)}
              <span
                class="border-death absolute top-0 h-full border-l border-dashed"
                style={`left: ${pct(death.at_ms)}%`}
                title={`${splitUnitName(death.name).name} died at ${formatDuration(death.at_ms)}`}
                aria-hidden="true"
              ></span>
            {/each}
          </span>
        </li>
      {/each}
      <li
        class="text-muted tabular grid grid-cols-[minmax(110px,160px)_40px_minmax(0,1fr)] gap-3 py-1 font-mono text-[11px]"
        data-testid="raid-cooldowns-axis"
      >
        <span class="relative col-start-3 block h-4">
          {#each axis as at (at)}
            <span class="absolute top-0 -translate-x-1/2" style={`left: ${pct(at)}%`}
              >{formatDuration(at)}</span
            >
          {/each}
          {#each shownDeaths as death (`${death.guid}-${death.at_ms}`)}
            <span
              class="text-death absolute top-0 -translate-x-1/2"
              style={`left: ${pct(death.at_ms)}%`}
              title={`${splitUnitName(death.name).name} died at ${formatDuration(death.at_ms)}`}>†</span
            >
          {/each}
        </span>
      </li>
    </ul>
  </section>
{/if}
