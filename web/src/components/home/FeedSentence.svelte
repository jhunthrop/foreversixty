<!-- web/src/components/home/FeedSentence.svelte -->
<!-- One feed sentence from its segments (GET /v1/feed): names in their class colour, items in their
     quality colour, the rest plain or strong. The API writes the words; this only styles them. -->
<script lang="ts">
  import type { FeedSegment } from '../../lib/home/feed';
  import { rarityClassFor } from '../../lib/planner/items';
  import { titleColor, type ClassSlug } from '../../lib/ui/primitives';

  let { segments }: { segments: readonly FeedSegment[] } = $props();
</script>

{#each segments as segment, index (index)}{#if segment.style === 'name'}<b
      class="font-bold"
      style:color={segment.class ? titleColor({ classSlug: segment.class as ClassSlug }) : undefined}
      >{segment.text}</b
    >{:else if segment.style === 'item'}<b class="font-bold {rarityClassFor(segment.quality ?? 1)}"
      >{segment.text}</b
    >{:else if segment.style === 'strong'}<b class="text-strong font-bold">{segment.text}</b
    >{:else}{segment.text}{/if}{/each}
