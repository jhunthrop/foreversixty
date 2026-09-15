<!-- web/src/components/report/RaidCooldowns.svelte -->
<!-- The raid cooldowns used this fight, with when: the first thing a raid leader reads
     on the Buffs tab. One lane per cooldown, a tick per use, the caster named on hover. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { formatDuration } from '../../lib/report/format';
  import { isRaidCooldown } from '../../lib/report/raid-cooldowns';
  import type { AuraTrack } from '../../lib/report/types';

  let { tracks, durationMs, names }: { tracks: AuraTrack[]; durationMs: number; names: Map<string, string> } =
    $props();

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
        .map(([name, uses]) => ({ name, uses: uses.sort((a, b) => a.at - b.at) }))
        // A track with no segment in this window has no use to draw.
        .filter((row) => row.uses.length > 0)
        .sort((a, b) => a.uses[0].at - b.uses[0].at)
    );
  });

  const pct = (ms: number): number => (durationMs === 0 ? 0 : (ms / durationMs) * 100);
</script>

{#if rows.length > 0}
  <section
    class="border-line rounded-panel bg-raised flex flex-col gap-2 border p-3"
    data-testid="raid-cooldowns"
  >
    <h2 class="label text-muted">Raid cooldowns</h2>
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
                style={`left: ${pct(use.at)}%; width: ${Math.max(pct(use.end - use.at), 0.4)}%`}
                title={`${formatDuration(use.at)}${use.source ? ` · ${splitUnitName(use.source).name}` : ''} on ${splitUnitName(use.target).name}`}
              ></span>
            {/each}
          </span>
        </li>
      {/each}
    </ul>
  </section>
{/if}
