<!-- web/src/components/ui/Capture.svelte -->
<!-- A client screenshot at its own aspect (setup spec 2026-11-07 section 3A): a 1 px --line-soft
     frame, 4 px radius, --raised behind object-fit: contain, capped at its native width and 420 px
     high (900 for a portrait capture, or `maxHeight`), width and height set so nothing moves as it
     loads, lazy unless `eager`. The image links to its own file; a sample carries a muted Tag. -->
<script lang="ts">
  import Tag from './Tag.svelte';

  let {
    src,
    width,
    height,
    alt,
    tag,
    tall = false,
    maxHeight,
    eager = false,
    testid,
  }: {
    src: string;
    width: number;
    height: number;
    alt: string;
    /** A muted Tag under the picture (`Example`). */
    tag?: string;
    /** A portrait capture: 900 px cap instead of 420. */
    tall?: boolean;
    /** Overrides the cap (240 inside the surfaces grid). */
    maxHeight?: number;
    eager?: boolean;
    testid?: string;
  } = $props();

  const cap = $derived(maxHeight ?? (tall ? 900 : 420));
</script>

<div class="ui-capture" data-testid={testid}>
  <a
    class="ui-capture-frame"
    href={src}
    style:max-width="{width}px"
    style:max-height="{cap}px"
    style:aspect-ratio="{width} / {height}"
  >
    <img {src} {width} {height} {alt} loading={eager ? 'eager' : 'lazy'} decoding="async" />
  </a>
  {#if tag}<Tag tone="muted">{tag}</Tag>{/if}
</div>

<style>
  .ui-capture {
    display: flex;
    min-width: 0;
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }
  .ui-capture-frame {
    display: block;
    width: 100%;
    overflow: hidden;
    border: 1px solid var(--color-line-soft);
    border-radius: 4px;
    background: var(--color-raised);
  }
  .ui-capture-frame:hover {
    border-color: var(--color-line);
  }
  .ui-capture-frame:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .ui-capture-frame img {
    display: block;
    width: 100%;
    height: 100%;
    object-fit: contain;
  }
</style>
