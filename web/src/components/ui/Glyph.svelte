<!-- web/src/components/ui/Glyph.svelte -->
<!-- One UI glyph from lib/ui/glyphs.ts: 24 px grid, 2 px stroke, currentColor, decorative
     (aria-hidden): the control that holds it carries the name. -->
<script lang="ts">
  import { FILLED_GLYPHS, GLYPHS, type GlyphName } from '../../lib/ui/glyphs';

  let {
    name,
    size = 16,
    class: className = '',
  }: { name: GlyphName; size?: number; class?: string } = $props();

  const filled = $derived(FILLED_GLYPHS.has(name));
</script>

<svg
  class={['shrink-0', className]}
  width={size}
  height={size}
  viewBox="0 0 24 24"
  fill={filled ? 'currentColor' : 'none'}
  stroke={filled ? 'none' : 'currentColor'}
  stroke-width="2"
  stroke-linecap="round"
  stroke-linejoin="round"
  aria-hidden="true"
  data-glyph={name}
>
  <!-- eslint-disable-next-line svelte/no-at-html-tags -- a constant from glyphs.ts, never user text -->
  {@html GLYPHS[name]}
</svg>
