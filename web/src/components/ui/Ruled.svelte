<!-- web/src/components/ui/Ruled.svelte -->
<!-- A list of rows that lay out their own grid (the guild roster's blocks, the readiness board):
     each child row gets the Row primitive's --color-line-soft hairline, on its top edge (`edge`
     top, the first row included) or its bottom edge. No box. The rule is unlayered and of zero
     specificity, so the row bands (global.css `.row-band-*` with `.row-rule`) still step it back
     beside a lit row. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    edge = 'top',
    id,
    testid,
    children,
  }: {
    edge?: 'top' | 'bottom';
    id?: string;
    testid?: string;
    children: Snippet;
  } = $props();
</script>

<div class="ui-ruled flex flex-col" data-edge={edge} {id} data-testid={testid}>{@render children()}</div>

<style>
  :global(:where(.ui-ruled[data-edge='top']) > *) {
    border-top: 1px solid var(--color-line-soft);
  }
  :global(:where(.ui-ruled[data-edge='bottom']) > *) {
    border-bottom: 1px solid var(--color-line-soft);
  }
</style>
