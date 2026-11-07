<!-- web/src/components/sim/ActionIcon.svelte -->
<!-- An action's icon in a fixed box (spec 2026-11-04-sim-results 8): the skeleton fill while the
     build's tables load, then the icon. A key the build has no icon for, and an image that fails,
     draw nothing: never an empty box or a missing-icon glyph (spec 15, R1). -->
<script lang="ts">
  let {
    url,
    loaded,
    size,
    label = '',
  }: {
    url: string | null;
    /** True once the tables have answered: a null `url` then means "no icon". */
    loaded: boolean;
    size: number;
    /** The name, for a row that shows only this icon. Empty beside a printed name. */
    label?: string;
  } = $props();

  let failedUrl = $state<string | null>(null);
  const shown = $derived(url !== null && url !== failedUrl);
</script>

{#if shown}
  <img
    src={url}
    width={size}
    height={size}
    alt={label}
    loading="lazy"
    decoding="async"
    class="inline-block shrink-0 rounded-[4px]"
    style:width="{size}px"
    style:height="{size}px"
    onerror={() => (failedUrl = url)}
    data-testid="action-icon"
  />
{:else if !loaded}
  <span
    class="skeleton-block inline-block shrink-0 rounded-[4px]"
    style:width="{size}px"
    style:height="{size}px"
    aria-hidden="true"
  ></span>
{/if}
