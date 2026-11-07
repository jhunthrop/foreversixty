<!-- web/src/components/RaidsCountdown.svelte -->
<!-- The first raids' countdown panel (home B spec 2026-11-07 §1): `First raids open · Wed 9 Dec`
     with the live dot, the clock to the opening and the three raids' icons; once open, `Raids open
     now` and the raids with no clock. Home's countdowns and the guild's Before raids open share it. -->
<script lang="ts">
  import { FIRST_RAIDS } from '../lib/first-raids';
  import { dayText } from '../lib/home/feed';
  import { homeBCopy } from '../lib/home/home-b-view';
  import Countdown from './ui/Countdown.svelte';
  import ItemIcon from './ui/ItemIcon.svelte';

  let {
    raidsOpen,
    build,
    testid,
  }: {
    /** The opening instant while ahead; null once raids are open. */
    raidsOpen: string | null;
    build: string;
    testid?: string;
  } = $props();
</script>

<Countdown
  eyebrow={raidsOpen === null ? homeBCopy.raidsOpen : homeBCopy.raidsEyebrow(dayText(raidsOpen))}
  target={raidsOpen}
  live
  tone="raid"
  {testid}
>
  <ul class="raids-countdown-list" aria-label="First raids">
    {#each FIRST_RAIDS as first (first.name)}
      <li><ItemIcon icon={first.icon} {build} size={40} eager /><span>{first.name}</span></li>
    {/each}
  </ul>
</Countdown>

<style>
  .raids-countdown-list {
    display: flex;
    justify-content: flex-end;
    gap: 18px;
    margin: 0;
    padding: 0;
    list-style: none;
  }
  .raids-countdown-list li {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 6px;
    width: 88px;
    font-size: 13px;
    line-height: 16px;
    font-weight: 600;
    text-align: center;
    color: var(--color-nav);
  }
</style>
