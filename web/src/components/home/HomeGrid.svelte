<!-- web/src/components/home/HomeGrid.svelte -->
<!-- Home's body (round 5, design/mocks/directions/pages.py HOME_CSS): the raid plaque leads full
     width, then the feed with the rail beside it (Your upgrades, or How to start), then the full
     width block (Guild sims, or Pick your class). Below 1100 one column, the rail before the feed,
     as the phone mock orders it. A grid inside the page column, never a container of its own. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    lead,
    rail,
    foot,
    testid,
    children,
  }: { lead: Snippet; rail?: Snippet; foot?: Snippet; testid?: string; children: Snippet } = $props();
</script>

<div class="home-grid" data-testid={testid}>
  <div class="home-grid-lead">{@render lead()}</div>
  <div class="home-grid-feed">{@render children()}</div>
  {#if rail}<div class="home-grid-rail">{@render rail()}</div>{/if}
  {#if foot}<div class="home-grid-foot">{@render foot()}</div>{/if}
</div>

<style>
  .home-grid {
    display: grid;
    grid-template-columns: minmax(0, 1fr);
    grid-template-areas: 'lead' 'rail' 'feed' 'foot';
    align-items: start;
    gap: var(--gap);
    padding: 28px var(--page-inset) 0;
  }
  .home-grid > div {
    min-width: 0;
  }
  .home-grid-lead {
    grid-area: lead;
  }
  .home-grid-feed {
    grid-area: feed;
  }
  .home-grid-rail {
    grid-area: rail;
    display: flex;
    flex-direction: column;
    gap: var(--gap);
  }
  .home-grid-foot {
    grid-area: foot;
  }
  @media (min-width: 1100px) {
    .home-grid {
      /* the mock's rail: five item rows with their gains, unwrapped */
      grid-template-columns: minmax(0, 1fr) 400px;
      grid-template-areas: 'lead lead' 'feed rail' 'foot foot';
    }
  }
  @media (max-width: 767px) {
    .home-grid {
      padding-top: 20px;
    }
  }
</style>
