<!-- web/src/components/home/HomePlaqueTitle.svelte -->
<!-- The raid plaque's title (round 5: the zone in Marcellus at 30, the plaque's own scale) and its
     one line under it: the guild, day and time, or the first raids' day. -->
<script lang="ts">
  let { title, sub }: { title: string; sub?: string } = $props();
</script>

<h2 class="home-plaque-title">{title}</h2>
{#if sub}<p class="home-plaque-sub">{sub}</p>{/if}

<style>
  .home-plaque-title {
    font-size: 30px;
    line-height: 34px;
  }
  .home-plaque-sub {
    margin-top: 4px;
    font-size: 15px;
    line-height: 22px;
    color: var(--color-text);
  }
  @media (max-width: 767px) {
    .home-plaque-title {
      font-size: 25px;
      line-height: 30px;
    }
  }
</style>
