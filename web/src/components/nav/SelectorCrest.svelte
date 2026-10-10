<!-- web/src/components/nav/SelectorCrest.svelte -->
<!-- The circular ringed HD class crest (tenet 14's "one crest language": never a square class
     icon, letter square or Blizzard avatar) at the selector's sizes. A character with no
     known class draws the neutral ringed disc instead of a broken `/crests/.webp` request,
     the same guard ClassCrestRing.svelte gives its callers. `lazy` is for the list's rows,
     which are not on the page-load path; the bar's own crest loads at once. -->
<script lang="ts">
  import { classCrestSrc } from '../../lib/class-crest';
  import { classColorVar } from '../../lib/report/format';

  let {
    classSlug,
    size,
    cssClass = 'csel-crest',
    lazy = false,
  }: { classSlug: string; size: number; cssClass?: string; lazy?: boolean } = $props();
</script>

{#if classSlug === ''}
  <span class={cssClass} style:--c="var(--color-line)" aria-hidden="true"></span>
{:else}
  <img
    class={cssClass}
    src={classCrestSrc(classSlug)}
    alt=""
    width={size}
    height={size}
    loading={lazy ? 'lazy' : undefined}
    decoding="async"
    style:--c={classColorVar(classSlug)}
  />
{/if}
