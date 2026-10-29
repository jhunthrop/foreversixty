<!-- web/src/components/TierStrip.svelte -->
<!-- The five Forever Score tiers as one strip: each segment a bar in the tier's exact item
     quality colour, the tier's name (rare and epic in their lightened text variants) and its
     range. Used on the profile's score card and on /ratings (where a visually hidden table
     carries the same facts for assistive tech).

     From sm: five aligned columns, every one the same padding and bar height; `current` is
     marked by an outline (which takes no room, so nothing shifts) and the other bars dim. The
     words never dim, so they keep their contrast. Below sm (fix round 2 item 8): one 12px strip
     of five segments, and under it only the current tier's name and range -- or, with no
     current tier, the five names small under their segments. -->
<script lang="ts">
  import { scoreCopy } from '../lib/score/copy';
  import { tierStrip } from '../lib/score/score';

  let { current = null }: { current?: string | null } = $props();

  const segments = tierStrip();
  const here = $derived(segments.find((segment) => segment.tier === current) ?? null);
  const phoneLabel = $derived(
    here === null
      ? `${scoreCopy.tierStripLabel}: ${segments.map((s) => `${s.label} ${s.range}`).join(', ')}`
      : `${scoreCopy.tierStripLabel}: ${here.label} ${here.range}, ${scoreCopy.currentTier}`,
  );
</script>

<div
  class="flex flex-col gap-1.5 sm:hidden"
  role="img"
  aria-label={phoneLabel}
  data-testid="tier-strip-phone"
>
  <div class="flex h-3 gap-[2px] overflow-hidden rounded-[3px]">
    {#each segments as segment (segment.tier)}
      <span
        class={`flex-1 ${here !== null && segment.tier !== here.tier ? 'opacity-40' : ''}`}
        style={`background-color: ${segment.color}`}
      ></span>
    {/each}
  </div>
  {#if here !== null}
    <p class="text-[13px]">
      <span class="font-semibold" style={`color: ${here.textColor}`}>{here.label}</span>
      <span class="text-muted numeric text-[12px]">{here.range}</span>
    </p>
  {:else}
    <div class="grid grid-cols-5 gap-[2px] text-center">
      {#each segments as segment (segment.tier)}
        <span class="text-[10px] font-semibold" style={`color: ${segment.textColor}`}>{segment.label}</span>
      {/each}
    </div>
  {/if}
</div>

<ol class="hidden grid-cols-5 gap-1 sm:grid" aria-label={scoreCopy.tierStripLabel} data-testid="tier-strip">
  {#each segments as segment (segment.tier)}
    {@const isHere = segment.tier === current}
    <li
      class={`rounded-control flex min-w-0 flex-col gap-1 p-1.5 ${isHere ? 'bg-card-top outline-gold/70 outline outline-1' : ''}`}
      aria-current={isHere ? 'true' : undefined}
      data-testid={`tier-strip-${segment.tier}`}
    >
      <span
        class={`block h-2 rounded-[2px] ${current !== null && !isHere ? 'opacity-40' : ''}`}
        style={`background-color: ${segment.color}`}
        aria-hidden="true"
      ></span>
      <span class="truncate text-[12px] font-semibold" style={`color: ${segment.textColor}`}
        >{segment.label}</span
      >
      <span class="text-muted numeric text-[11px] whitespace-nowrap">{segment.range}</span>
      {#if isHere}<span class="sr-only">{scoreCopy.currentTier}</span>{/if}
    </li>
  {/each}
</ol>
