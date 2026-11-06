<!-- web/src/components/ui/ItemIcon.svelte -->
<!-- An item's icon (web layout spec 2026-11-05 §4.3): the HD icon in a 1 px frame of its quality
     hue, 5 px radius, width and height always set (no layout shift). With no icon it draws the
     dark socket, never a broken image; `unknown` tints it red (an id missing from our data). A
     badge disc sits on the lower right: gold ▲ for an upgrade, green ✓ for the worn pick. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { rarityBorderColorFor } from '../../lib/planner/items';
  import { dataUrl } from '../../lib/planner/load';
  import Glyph from './Glyph.svelte';

  let {
    icon,
    quality = 1,
    build,
    size = 40,
    badge,
    unknown = false,
    eager = false,
    deferred = false,
    testid,
  }: {
    /** The icon's file stem (`inv_helmet_04`); absent, the empty socket. */
    icon?: string;
    quality?: number;
    build: string;
    size?: 28 | 32 | 36 | 40 | 56;
    badge?: 'upgrade' | 'worn';
    unknown?: boolean;
    /** Above the fold: load at once instead of lazily. */
    eager?: boolean;
    /** The image loads once the island has mounted, not with the server HTML: a page whose grid of
     *  icons sits near the fold keeps them off its first paint's critical path (the frame and its
     *  quality border still render at once, so nothing shifts). */
    deferred?: boolean;
    testid?: string;
  } = $props();

  let mounted = $state(false);
  onMount(() => {
    mounted = true;
  });

  const BADGE = {
    upgrade: { glyph: 'upgrade', title: 'Upgrade', tone: 'bg-gold text-bg' },
    worn: { glyph: 'check', title: 'You wear the pick', tone: 'bg-kill text-bg' },
  } as const;
</script>

<span
  class="ui-item-icon relative inline-flex shrink-0"
  style:width="{size}px"
  style:height="{size}px"
  data-testid={testid}
>
  {#if icon && deferred && !mounted}
    <span class="h-full w-full rounded-[5px] border" style:border-color={rarityBorderColorFor(quality)}></span>
  {:else if icon}
    <img
      src={dataUrl(build, `icons/${icon}.webp`)}
      alt=""
      width={size}
      height={size}
      loading={eager ? 'eager' : 'lazy'}
      decoding="async"
      class="h-full w-full rounded-[5px] border object-cover"
      style:border-color={rarityBorderColorFor(quality)}
    />
  {:else}
    <span
      class={[
        'h-full w-full rounded-[5px] border',
        unknown ? 'border-wipe/50 bg-wipe/10' : 'border-line-soft border-dashed',
      ]}
      aria-hidden="true"
    ></span>
  {/if}
  {#if badge}
    <span
      class="absolute -right-1 -bottom-1 flex h-[18px] w-[18px] items-center justify-center rounded-full {BADGE[
        badge
      ].tone}"
      title={BADGE[badge].title}
      data-badge={badge}><Glyph name={BADGE[badge].glyph} size={11} /></span
    >
  {/if}
</span>
