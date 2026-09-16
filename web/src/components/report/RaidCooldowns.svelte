<!-- web/src/components/report/RaidCooldowns.svelte -->
<!-- The raid cooldowns used this fight, with when: the first thing a raid leader reads
     on the Buffs tab. One lane per cooldown, a tick per use, the caster named on hover,
     a time axis under the lanes and every death drawn across them, so "Ardent Defender
     at 48s, tank dead at 1:04" is one picture. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { formatDuration } from '../../lib/report/format';
  import { isRaidCooldown } from '../../lib/report/raid-cooldowns';
  import type { AuraTrack, CastRow, Death, PullMark } from '../../lib/report/types';
  import type { TimeWindow } from '../../lib/report/window';

  let {
    tracks,
    casts = [],
    pulls = [],
    window: timeWindow,
    deaths = [],
    names,
  }: {
    tracks: AuraTrack[];
    /** The fight's casts, for a cooldown that leaves no aura behind: Revival heals and is gone. */
    casts?: CastRow[];
    /** Over the night: where each pull sits, drawn as bands so a use can be read against its pull. */
    pulls?: PullMark[];
    window: TimeWindow;
    deaths?: Death[];
    names: Map<string, string>;
  } = $props();

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
    // A cooldown that is a cast and no aura (Revival, Invoke Yu'lon) is drawn from its casts,
    // as an instant, unless an aura of the same name already told the story.
    for (const row of casts) {
      if (!isRaidCooldown(row.spell_name) || byName.has(row.spell_name)) continue;
      const uses = byName.get(row.spell_name) ?? [];
      for (const at of row.sequence) uses.push({ at, end: at, source: row.name, target: 'the raid' });
      byName.set(row.spell_name, uses);
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
  /** What the pointer is over: the pull and the instant, so a band never needs a tooltip. */
  let hovered = $state('');
  function readAt(event: PointerEvent): void {
    const lane = (event.target as HTMLElement).closest<HTMLElement>('[data-lane]');
    if (lane === null) return;
    const bounds = lane.getBoundingClientRect();
    const at = timeWindow.startMs + ((event.clientX - bounds.left) / Math.max(bounds.width, 1)) * span;
    const pull = pulls.find((entry) => entry.start_ms <= at && at < entry.end_ms);
    hovered = `${pull === undefined ? '' : `${pull.label} · `}${formatDuration(at)}`;
  }
  /** Per cooldown, the pulls it was used in, and the ones it was not. */
  function pullsUsed(uses: Use[]): { used: number; missed: string[] } {
    const missed = pulls
      .filter((pull) => !uses.some((use) => use.at >= pull.start_ms && use.at < pull.end_ms))
      .map((pull) => pull.label);
    return { used: pulls.length - missed.length, missed };
  }
  /** Where an instant sits across the lane, 0..100, from the window's start. */
  const pct = (ms: number): number => ((ms - timeWindow.startMs) / span) * 100;

  /** Axis ticks at the coarsest of these steps that keeps them under about eight. */
  const axis = $derived.by(() => {
    const steps = [10_000, 30_000, 60_000, 120_000, 300_000, 600_000, 1_800_000];
    // Five at most: a phone lane is 250px wide and six labels there ran into one another.
    const step = steps.find((candidate) => span / candidate <= 5) ?? steps[steps.length - 1];
    const first = Math.ceil(timeWindow.startMs / step) * step;
    const ticks: number[] = [];
    for (let at = first; at <= timeWindow.endMs; at += step) ticks.push(at);
    return ticks;
  });

  const shownPulls = $derived(
    pulls.filter((pull) => pull.end_ms > timeWindow.startMs && pull.start_ms < timeWindow.endMs),
  );
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
    {#if pulls.length > 0}
      <p class="text-muted label h-5 truncate" data-testid="raid-cooldowns-readout">
        {#if hovered !== ''}<span class="text-strong normal-case">{hovered}</span>{:else}hover a lane for the
          pull · alternate shading is one pull each · a red top edge is a wipe{/if}
      </p>
    {/if}
    <ul class="flex flex-col">
      {#each rows as row (row.name)}
        {@const across = pullsUsed(row.uses)}
        <li
          class="border-line-soft grid min-h-9 grid-cols-[minmax(0,1fr)_56px] items-center gap-x-3 gap-y-1 border-b py-1 text-[13px] md:grid-cols-[minmax(110px,160px)_56px_minmax(0,1fr)]"
          data-testid={`raid-cooldown-${row.name}`}
        >
          <span class="truncate font-semibold">{row.name}</span>
          {#if pulls.length > 0}
            <span
              class="tabular text-right font-mono text-[12px]"
              class:text-wipe={across.used < pulls.length}
              class:text-muted={across.used === pulls.length}
              title={across.missed.length === 0
                ? 'Used on every pull'
                : `Not used on: ${across.missed.join(', ')}`}
              data-testid="raid-cooldown-pulls">{across.used}/{pulls.length}</span
            >
          {:else}
            <span class="text-muted tabular text-right font-mono text-[12px]" title="Times used"
              >{row.uses.length}</span
            >
          {/if}
          <span
            class="bg-line-soft relative col-span-2 block h-6 w-full touch-none md:col-span-1 md:h-[14px]"
            data-lane
            onpointermove={readAt}
            onpointerdown={readAt}
          >
            {#each shownPulls as pull, i (pull.start_ms)}
              {#if i % 2 === 1}
                <span
                  class="absolute top-0 h-full"
                  style={`left: ${Math.max(pct(pull.start_ms), 0)}%; width: ${Math.min(pct(pull.end_ms), 100) - Math.max(pct(pull.start_ms), 0)}%; background: var(--color-text); opacity: .08`}
                  aria-hidden="true"
                ></span>
              {/if}
              {#if !pull.kill}
                <span
                  class="bg-wipe absolute top-0 h-[2px]"
                  style={`left: ${Math.max(pct(pull.start_ms), 0)}%; width: ${Math.min(pct(pull.end_ms), 100) - Math.max(pct(pull.start_ms), 0)}%`}
                  aria-hidden="true"
                ></span>
              {/if}
            {/each}
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
        class="text-muted tabular grid grid-cols-[minmax(0,1fr)] gap-3 py-1 font-mono text-[11px] md:grid-cols-[minmax(110px,160px)_56px_minmax(0,1fr)]"
        data-testid="raid-cooldowns-axis"
      >
        <span class="relative block h-4 md:col-start-3">
          {#each axis as at (at)}
            <span class="absolute top-0 -translate-x-1/2" style={`left: ${pct(at)}%`}
              >{formatDuration(at)}</span
            >
          {/each}
          <!-- Over a whole night the daggers would smear; the lines through the lanes remain. -->
          {#each shownDeaths.length <= 12 ? shownDeaths : [] as death (`${death.guid}-${death.at_ms}`)}
            <span
              class="text-death absolute top-0 hidden -translate-x-1/2 md:inline"
              style={`left: ${pct(death.at_ms)}%`}
              title={`${splitUnitName(death.name).name} died at ${formatDuration(death.at_ms)}`}>†</span
            >
          {/each}
        </span>
      </li>
    </ul>
  </section>
{/if}
