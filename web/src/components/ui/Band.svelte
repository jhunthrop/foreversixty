<!-- web/src/components/ui/Band.svelte -->
<!-- A section band (round 5, ui/band.css): full bleed, the owner's painted blue band under its veil
     (or home's camp-road `hero`), the content back in the page column, their black rule under it.
     PageHeader and ControlBar already sit on it; a page reaches for Band only for a band of its own
     under the title band (a strip of controls, a plaque). `y` is the band's distance from the top of
     the page in CSS, so the image continues from the band above it. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { BandArt } from '../../lib/ui/primitives';

  let {
    art = 'band',
    y = 'var(--site-top-h, 80px)',
    testid,
    children,
  }: {
    art?: BandArt;
    /** The band's top in the page, a CSS length: the image is drawn offset by it. */
    y?: string;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

<div class="ui-band ui-band-end full-bleed" data-art={art} style:--band-y={y} data-testid={testid}>
  <div class="page-column">{@render children()}</div>
</div>
