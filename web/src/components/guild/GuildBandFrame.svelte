<!-- web/src/components/guild/GuildBandFrame.svelte -->
<!-- The guild band's frame (guild page spec 12.2.B, raids spec 3.2): the full-bleed raised band,
     the faction wash and the watermark logo clipped by the band's own edges, and the centred page
     column the content sits in. The raid page draws its band in it; the markup is Guild.svelte's
     band, so both pages read as one band. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { FACTION_BAR_COLOR, factionLogoSrc } from '../../lib/faction-mark';

  let {
    faction,
    minHeight = '',
    testid = 'guild-header-band',
    children,
  }: {
    faction: 'alliance' | 'horde' | null;
    /** Tailwind min-height classes the band holds while it loads. */
    minHeight?: string;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

<div class="bg-raised full-bleed relative overflow-hidden {minHeight}" data-testid={testid}>
  {#if faction !== null}
    <div
      class="pointer-events-none absolute inset-0 z-0"
      style={`background:linear-gradient(135deg, transparent 52%, color-mix(in srgb, ${FACTION_BAR_COLOR[faction]} 38%, transparent) 100%)`}
      aria-hidden="true"
    ></div>
  {/if}
  <div class="page-column relative flex flex-col gap-3 px-[18px] pt-4 pb-[22px] md:pt-7 md:pb-8">
    {#if faction !== null}
      <div
        class="pointer-events-none absolute top-[-22px] right-[-80px] bottom-0 z-0 w-[180px] overflow-hidden lg:top-[-38px] lg:right-[-90px] lg:w-[320px]"
        aria-hidden="true"
      >
        <img
          src={factionLogoSrc(faction)}
          alt=""
          width={320}
          height={320}
          loading="eager"
          decoding="async"
          class="h-full w-full object-contain object-top opacity-[.14] lg:opacity-[.12]"
        />
      </div>
    {/if}
    <div class="relative z-10 flex flex-col gap-3">
      {@render children()}
    </div>
  </div>
</div>
