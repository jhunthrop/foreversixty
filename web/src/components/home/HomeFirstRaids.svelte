<!-- web/src/components/home/HomeFirstRaids.svelte -->
<!-- The raid plaque with no guild raid to show (signed out, no guild, or a guild with none planned):
     the first raids opening on the cave art, its day, the clock in days and hours and the three
     raids; once open, "Raids open now" and the raids alone. `foot` carries the guild line (Find a
     guild, or the guild with no raid planned). -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { FIRST_RAIDS } from '../../lib/first-raids';
  import { minuteClock } from '../../lib/guild/minute-clock.svelte';
  import { longDay, plaqueClock, plaqueClockLabel, plaqueCopy } from '../../lib/home/plaque';
  import { REGION } from '../../lib/ui/regions';
  import BigFigures from '../ui/BigFigures.svelte';
  import Frame from '../ui/Frame.svelte';
  import ItemIcon from '../ui/ItemIcon.svelte';
  import HomePlaqueTitle from './HomePlaqueTitle.svelte';

  let {
    raidsOpen,
    build,
    foot,
  }: {
    /** The opening instant while ahead; null once raids are open. */
    raidsOpen: string | null;
    build: string;
    foot?: Snippet;
  } = $props();

  const clock = minuteClock(() => undefined);
  const parts = $derived(raidsOpen === null ? [] : plaqueClock(new Date(raidsOpen), clock.now));
  const open = $derived(parts.length === 0);
</script>

<Frame ornate art="cave" region={REGION.countdowns} testid="home-raids-countdown">
  <div class="home-first-box">
    <div class="home-first">
      <div class="min-w-0">
        <HomePlaqueTitle
          title={open ? plaqueCopy.raidsOpen : plaqueCopy.firstRaids}
          sub={open || raidsOpen === null ? undefined : longDay(raidsOpen)}
        />
        {#if !open}
          <div class="mt-4">
            <BigFigures groups={[parts]} timer={plaqueClockLabel(parts)} testid="home-raids-clock" />
          </div>
        {/if}
        {#if foot}<div class="home-first-foot">{@render foot()}</div>{/if}
      </div>
      <ul class="home-first-raids" aria-label="First raids">
        {#each FIRST_RAIDS as raid (raid.name)}
          <li><ItemIcon icon={raid.icon} {build} size={56} eager /><span>{raid.name}</span></li>
        {/each}
      </ul>
    </div>
  </div>
</Frame>

<style>
  .home-first {
    display: grid;
    gap: 20px 34px;
    align-items: center;
  }
  .home-first-foot {
    margin-top: 18px;
    font-size: 15px;
    line-height: 22px;
    color: var(--color-text);
  }
  .home-first-raids {
    display: flex;
    align-items: flex-start;
    gap: 18px;
    margin: 0;
    padding: 0;
    list-style: none;
  }
  .home-first-raids li {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 8px;
    width: 96px;
    font-weight: 600;
    font-size: 14px;
    line-height: 18px;
    text-align: center;
    color: var(--color-strong);
  }
  /* Side by side where the plaque is wide (it leads a signed-in page); stacked in the signed-out
     rail, whatever the viewport. */
  .home-first-box {
    container-type: inline-size;
  }
  @container (min-width: 640px) {
    .home-first {
      grid-template-columns: minmax(0, 1fr) auto;
    }
  }
  @media (max-width: 767px) {
    .home-first-raids li {
      width: auto;
      flex: 1 1 0;
    }
  }
</style>
