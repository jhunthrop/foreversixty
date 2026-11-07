<!-- web/src/components/ui/Overlay.svelte -->
<!-- The overlay surface (web layout spec 2026-11-05 §3A.2): tooltips, sheets and menus share one box,
     drawn by ui/overlay.css (`.ui-overlay` and its kinds; the tooltip kind is the in-game frame).
     Positioning stays with the caller (Floating UI, the tooltip host): `class` and `style` carry the
     caller's placement and inner layout, never a surface. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    kind = 'tooltip',
    role,
    label,
    testid,
    id,
    class: layout = '',
    style,
    children,
    foot,
  }: {
    kind?: 'tooltip' | 'sheet' | 'menu';
    role?: string;
    label?: string;
    testid?: string;
    /** The element id an anchor's aria-describedby points at (a tooltip). */
    id?: string;
    /** The caller's placement and inner layout (position, width, padding, flex). */
    class?: string;
    /** The caller's placement variables (the tooltip host's computed position). */
    style?: string;
    children: Snippet;
    /** A muted note under a hairline at the overlay's foot (a checklist's explanation, an item's source). */
    foot?: Snippet;
  } = $props();
</script>

<div
  class={['ui-overlay', layout]}
  data-kind={kind}
  role={role ?? (kind === 'menu' ? 'menu' : kind === 'tooltip' ? 'tooltip' : 'dialog')}
  aria-label={label}
  data-testid={testid}
  {id}
  {style}
>
  {@render children()}
  {#if foot}<div class="ui-overlay-rule-t text-muted px-[14px] pt-2 pb-3 text-[12px] leading-4">
      {@render foot()}
    </div>{/if}
</div>
