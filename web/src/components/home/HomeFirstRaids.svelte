<!-- web/src/components/home/HomeFirstRaids.svelte -->
<!-- The raid plaque with no guild raid to show (signed out, no guild, or a guild with none planned):
     the first raids opening on the cave art, its day and the clock in days and hours; `foot` (the
     guild line: Find a guild, or the guild with no raid planned and its facts) beside the clock; and
     the three raids as a list, each with its bosses and the drops the viewer's class can use
     (lib/home/raid-facts.ts). Once open, "Raids open now" and the raids. Where the plaque is wide
     (it leads a signed-in page) the three sit in a row; in the signed-out rail they stack. -->
<script lang="ts">
  import ZoneIcon from '../ui/ZoneIcon.svelte';
  import type { Snippet } from 'svelte';
  import { minuteClock } from '../../lib/guild/minute-clock.svelte';
  import { longDay, plaqueClock, plaqueClockLabel, plaqueCopy } from '../../lib/home/plaque';
  import { raidLines, type RaidFact } from '../../lib/home/raid-facts';
  import { REGION } from '../../lib/ui/regions';
  import BigFigures from '../ui/BigFigures.svelte';
  import Frame from '../ui/Frame.svelte';
  import HomePlaqueTitle from './HomePlaqueTitle.svelte';

  let {
    raidsOpen,
    raids,
    viewer,
    build,
    foot,
  }: {
    /** The opening instant while ahead; null once raids are open. */
    raidsOpen: string | null;
    raids: readonly RaidFact[];
    /** The viewer's class: each raid counts the drops it can use. */
    viewer?: { slug: string; name: string };
    build: string;
    foot?: Snippet;
  } = $props();

  const RAID_ICON_PX = 44;
  const clock = minuteClock(() => undefined);
  const lines = $derived(raidLines(raids, viewer));
  const parts = $derived(raidsOpen === null ? [] : plaqueClock(new Date(raidsOpen), clock.now));
  const open = $derived(parts.length === 0);
</script>

<Frame ornate art="cave" region={REGION.countdowns} testid="home-raids-countdown">
  <div class="home-first-box">
    <div class="home-first">
      <div class="home-first-lead min-w-0">
        <HomePlaqueTitle
          title={open ? plaqueCopy.raidsOpen : plaqueCopy.firstRaids}
          sub={open || raidsOpen === null ? undefined : longDay(raidsOpen)}
        />
        {#if !open}
          <div class="mt-4">
            <BigFigures groups={[parts]} timer={plaqueClockLabel(parts)} testid="home-raids-clock" />
          </div>
        {/if}
      </div>
      {#if foot}<div class="home-first-foot min-w-0" data-testid="home-raids-guild">
          {@render foot()}
        </div>{/if}
      <ul class="home-first-raids" aria-label={plaqueCopy.firstRaids}>
        {#each raids as raid, index (raid.name)}
          {@const line = lines[index]}
          <li data-testid="home-raid-zone">
            <ZoneIcon zone={raid.name} {build} size={RAID_ICON_PX} eager />
            <span class="min-w-0">
              <span class="home-first-raid-name">{raid.name}</span>
              {#if line}<span class="home-first-raid-line">{line}</span>{/if}
            </span>
          </li>
        {/each}
      </ul>
    </div>
  </div>
</Frame>

<style>
  .home-first {
    display: grid;
    gap: 20px 40px;
    align-items: center;
  }
  .home-first-foot {
    font-size: 15px;
    line-height: 22px;
    color: var(--color-text);
  }
  .home-first-raids {
    display: flex;
    flex-direction: column;
    gap: 12px;
    margin: 0;
    padding: 0 16px 0 0;
    list-style: none;
  }
  .home-first-raids li {
    display: grid;
    grid-template-columns: auto minmax(0, 1fr);
    align-items: center;
    gap: 12px;
  }
  .home-first-raid-name {
    display: block;
    font-weight: 700;
    font-size: 15px;
    line-height: 20px;
    color: var(--color-strong);
  }
  .home-first-raid-line {
    display: block;
    font-size: 13px;
    line-height: 18px;
    color: var(--color-muted);
  }
  .home-first-box {
    container-type: inline-size;
  }
  /* The signed-in plaque: the clock, the guild beside it, the raids at the right. */
  @container (min-width: 640px) {
    .home-first {
      grid-template-columns: auto minmax(0, 1fr);
    }
    .home-first-foot {
      align-self: end;
    }
    .home-first-raids {
      grid-column: 1 / -1;
      flex-direction: row;
      flex-wrap: wrap;
      gap: 12px 32px;
    }
  }
  @container (min-width: 1100px) {
    .home-first {
      grid-template-columns: auto minmax(0, 1fr) auto;
    }
    .home-first-raids {
      grid-column: auto;
      flex-direction: column;
    }
  }
</style>
