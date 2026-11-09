<!-- web/src/components/home/HomeFacts.svelte -->
<!-- One line of facts on a home panel (density pass 2026-11-09): each a strong figure and its words,
     as the plaque's role counts read ("3 tanks  6 healers"). Nothing drawn with no facts. -->
<script lang="ts">
  import type { HomeFact } from '../../lib/home/home-facts';

  let { facts, testid }: { facts: readonly HomeFact[]; testid?: string } = $props();
</script>

{#if facts.length > 0}
  <p class="text-text flex flex-wrap gap-x-[18px] gap-y-[6px] text-[14px] leading-5" data-testid={testid}>
    {#each facts as fact (fact.text)}<span><b class="text-strong">{fact.n}</b> {fact.text}</span>{/each}
  </p>
{/if}
