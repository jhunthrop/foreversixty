<!-- web/src/components/ui/Overlay.svelte -->
<!-- The overlay surface (web layout spec 2026-11-05 §3A.2): tooltips, sheets and menus share the S1
     ground, the 1 px --color-line edge, the plate radius and the one overlay shadow. Positioning
     stays with the caller (Floating UI, the tooltip host); this draws the box only. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    kind = 'tooltip',
    role,
    label,
    testid,
    children,
    foot,
  }: {
    kind?: 'tooltip' | 'sheet' | 'menu';
    role?: string;
    label?: string;
    testid?: string;
    children: Snippet;
    /** A muted note under a hairline at the overlay's foot (a checklist's explanation). */
    foot?: Snippet;
  } = $props();
</script>

<div
  class={[
    'ui-overlay bg-raised border-line shadow-overlay min-w-0 border',
    kind === 'sheet' ? 'rounded-t-plate' : 'rounded-plate',
  ]}
  data-kind={kind}
  role={role ?? (kind === 'menu' ? 'menu' : kind === 'tooltip' ? 'tooltip' : 'dialog')}
  aria-label={label}
  data-testid={testid}
>
  {@render children()}
  {#if foot}<div class="border-line-soft text-muted border-t px-[14px] pt-2 pb-3 text-[12px] leading-4">
      {@render foot()}
    </div>{/if}
</div>
