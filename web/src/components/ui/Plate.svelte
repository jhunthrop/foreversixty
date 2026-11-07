<!-- web/src/components/ui/Plate.svelte -->
<!-- The one framed object on a page (web layout spec 2026-11-05 P1, §3A.2): S1 ground, one 1 px
     --color-line edge, the 6 px plate radius, no gradient and no shadow. It holds the page's
     primary object (the doll, the tier list, the result) and never a Section (§3A.2 rules).
     `flush` drops the inner padding for content that runs its own rows edge to edge. -->
<script lang="ts">
  import { PLATE_SURFACE } from '../../lib/ui/primitives';
  import type { Snippet } from 'svelte';
  import { REGION } from '../../lib/ui/regions';

  let {
    flush = true,
    label,
    testid,
    children,
  }: {
    flush?: boolean;
    /** The region's accessible name, when the plate is a landmark of its own. */
    label?: string;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

<div
  class={PLATE_SURFACE}
  class:p-4={!flush}
  role={label ? 'region' : undefined}
  aria-label={label}
  data-testid={testid}
  data-region={REGION.plate}
>
  {@render children()}
</div>
