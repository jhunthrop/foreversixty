<!-- web/src/components/home/HomeHowToStart.svelte -->
<!-- How to start, the signed-out rail (round 5: the page's second featured panel, where a signed-in
     player sees Your upgrades): sign in with Battle.net, then the addon and the companion. -->
<script lang="ts">
  import { homeCopy, homeHrefs } from '../../lib/home/home-view';
  import { sectionRegion } from '../../lib/ui/regions';
  import Frame from '../ui/Frame.svelte';
  import Row from '../ui/Row.svelte';

  let { signInHref }: { signInHref: string } = $props();

  const hrefs = $derived([signInHref, homeHrefs.addon]);
</script>

<Frame region={sectionRegion(homeCopy.howToStart)} testid="home-how-to-start">
  <h2 class="text-section pb-[6px]">{homeCopy.howToStart}</h2>
  {#each homeCopy.steps as step, index (step.title)}
    <Row rank={String(index + 1)} title={step.title} href={hrefs[index]} open />
  {/each}
</Frame>
