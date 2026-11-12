<!-- web/src/components/home/HomeHowToStart.svelte -->
<!-- How to start, the signed-out rail (round 5: the page's second featured panel, where a signed-in
     player sees Your upgrades): sign in with Battle.net, then the addon and the companion. The page's
     one sign-in (the title band carries none), on a plain panel: no inset line, no card in a card. -->
<script lang="ts">
  import { homeCopy, homeHrefs } from '../../lib/home/home-view';
  import { sectionRegion } from '../../lib/ui/regions';
  import Frame from '../ui/Frame.svelte';
  import Row from '../ui/Row.svelte';
  import SectionTitle from '../ui/SectionTitle.svelte';

  let { signInHref }: { signInHref: string } = $props();

  const hrefs = $derived([signInHref, homeHrefs.addon]);
</script>

<Frame plain region={sectionRegion(homeCopy.howToStart)} testid="home-how-to-start">
  <SectionTitle class="pb-[6px]" title={homeCopy.howToStart} />
  {#each homeCopy.steps as step, index (step.title)}
    <Row
      rank={String(index + 1)}
      title={step.title}
      href={hrefs[index]}
      open
      testid={index === 0 ? 'home-sign-in' : undefined}
    />
  {/each}
</Frame>
