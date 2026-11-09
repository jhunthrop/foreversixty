<!-- web/src/components/ui/Icon.svelte -->
<!-- One icon from the owner's custom set (lib/ui/ui-icons.ts, generated from
     design/assets/icons/ui): a filled path in currentColor, so the caller's token colours it.
     Decorative (aria-hidden): the control that holds it carries the name. Line glyphs stay
     ui/Glyph.svelte's. -->
<script lang="ts">
  import { UI_ICONS, type UiIconName } from '../../lib/ui/ui-icons';

  let {
    name,
    size = 20,
    class: className = '',
  }: { name: UiIconName; size?: number; class?: string } = $props();

  const icon = $derived(UI_ICONS[name]);
</script>

<svg
  class={['shrink-0', className]}
  width={size}
  height={size}
  viewBox={`0 0 ${icon.size} ${icon.size}`}
  fill="currentColor"
  aria-hidden="true"
  data-icon={name}
>
  <path fill-rule="evenodd" d={icon.d} />
</svg>
