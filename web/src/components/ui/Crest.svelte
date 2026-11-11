<!-- web/src/components/ui/Crest.svelte -->
<!-- The one class crest (component audit 2026-11-10 cluster 2; owner: "one crest language"): the
     circular ringed HD crest, no glow, in the class's ring; an unknown class draws the LX seal in a
     muted ring. styles/crest.css owns the look; ui/Crest.astro is the static twin. `cls` takes a
     class name or slug. With `class`, `size` sets only the intrinsic width and height and the
     caller's classes size it (a size per container width). -->
<script lang="ts">
  import { CREST_CLASS, crestImage } from '../../lib/class-crest';

  let {
    cls,
    size,
    eager = false,
    class: extra = '',
    testid,
  }: {
    cls: string | null | undefined;
    size: number;
    /** Above the fold or swapped in for a pre-painted crest: load at once and decode with the
     *  paint (never fetchpriority high, see crest.css). */
    eager?: boolean;
    class?: string;
    testid?: string;
  } = $props();

  const image = $derived(crestImage(cls, size));
  const box = $derived(extra === '' ? `width: ${size}px; height: ${size}px;` : '');
</script>

<img
  src={image.src}
  alt=""
  width={size}
  height={size}
  loading={eager ? 'eager' : 'lazy'}
  fetchpriority={eager ? undefined : 'low'}
  decoding={eager ? 'sync' : 'async'}
  class={[CREST_CLASS, !image.known && 'crest-none', extra]}
  style={`--c: ${image.ring}; ${box}`}
  data-testid={testid === undefined ? undefined : image.known ? testid : `${testid}-fallback`}
/>
