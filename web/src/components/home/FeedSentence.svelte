<!-- web/src/components/home/FeedSentence.svelte -->
<!-- One feed sentence from its segments (GET /v1/feed): names in their class colour, items in their
     quality colour, a sim's change green up and red down, the rest plain or strong. The API writes
     the words; this only styles them. -->
<script lang="ts">
  import ItemName from '../ui/ItemName.svelte';
  import type { FeedSegment } from '../../lib/home/feed';
  import { TONE_TEXT, titleColor, type ClassSlug } from '../../lib/ui/primitives';

  let { segments }: { segments: readonly FeedSegment[] } = $props();
</script>

{#each segments as segment, index (index)}{#if segment.style === 'name'}<b
      class="font-bold"
      style:color={segment.class ? titleColor({ classSlug: segment.class as ClassSlug }) : undefined}
      >{segment.text}</b
    >{:else if segment.style === 'item'}<ItemName
      name={segment.text}
      quality={segment.quality ?? 1}
      id={segment.item_id}
      focusable={false}
      class="font-bold"
    />{:else if segment.style === 'strong'}<b class="text-strong font-bold">{segment.text}</b
    >{:else if segment.style === 'gain' || segment.style === 'loss'}<b
      class="font-bold tabular-nums {TONE_TEXT[segment.style === 'gain' ? 'gain' : 'wipe']}">{segment.text}</b
    >{:else}{segment.text}{/if}{/each}
