<!-- web/src/components/ui/ItemName.svelte -->
<!-- An item's name inline in a line of text (component audit 2026-11-10 cluster 9; tenets 2 and 9):
     the name in its quality colour and, given the item's `id`, the hook the page's shared item
     tooltip (BisTooltipHost) opens from on hover and, unless `focusable` is off, on keyboard focus,
     so no item on the site is ever just a coloured word. `focusable` is off only where the name
     already sits inside a link or another tab stop (a feed row), which would nest them. Inside an
     ItemHover, which owns its own tooltip, the name takes no `id`. -->
<script lang="ts">
  import { rarityClassFor } from '../../lib/planner/items';

  let {
    name,
    quality,
    id,
    focusable = true,
    class: extra = '',
    testid,
  }: {
    name: string;
    quality: number;
    /** The item, for the shared tooltip; absent, the name is colour alone (an ItemHover's child). */
    id?: number;
    focusable?: boolean;
    /** The caller's weight, size and truncation. */
    class?: string;
    testid?: string;
  } = $props();

  const stop = $derived(id !== undefined && focusable);
</script>

<!-- svelte-ignore a11y_no_noninteractive_tabindex -->
<span
  class={[
    rarityClassFor(quality),
    stop && 'focus-visible:outline-gold rounded-[2px] hover:underline focus-visible:outline-2',
    extra,
  ]}
  data-bis-item={id}
  tabindex={stop ? 0 : undefined}
  data-testid={testid}>{name}</span
>
