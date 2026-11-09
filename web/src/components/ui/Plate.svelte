<!-- web/src/components/ui/Plate.svelte -->
<!-- The page's primary object on its panel (round 5, ui/panel.css): the card brown-black, the black
     hairline inside the warm ring, the 10 px radius. It holds the page's primary object (the doll,
     the tier list, the result) and never a Section. `frame` marks a featured panel: `framed` the
     cream inset line, `lead` the line and the corner flourishes (the page's one lead panel; two or
     three framed panels a page at most). `flush` drops the inner padding for content that runs its
     own rows edge to edge. -->
<script lang="ts">
  import { PLATE_SURFACE, frameClass, type PanelFrame } from '../../lib/ui/primitives';
  import type { Snippet } from 'svelte';
  import { REGION } from '../../lib/ui/regions';
  import FrameCorners from './FrameCorners.svelte';

  let {
    flush = true,
    frame = 'plain',
    label,
    testid,
    children,
  }: {
    flush?: boolean;
    frame?: PanelFrame;
    /** The region's accessible name, when the plate is a landmark of its own. */
    label?: string;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

<div
  class="{PLATE_SURFACE} {frameClass(frame)}"
  class:p-4={!flush && frame === 'plain'}
  class:ui-framed-pad={!flush && frame !== 'plain'}
  role={label ? 'region' : undefined}
  aria-label={label}
  data-testid={testid}
  data-region={REGION.plate}
>
  {#if frame === 'lead'}<FrameCorners />{/if}
  {@render children()}
</div>

<style>
  .ui-framed-pad {
    padding: 22px 24px;
  }
</style>
