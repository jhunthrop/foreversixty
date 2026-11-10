<!-- web/src/components/ui/Icon.svelte -->
<!-- One icon from the site's one icon source (lib/ui/ui-icons.ts: the owner's custom set and the
     line family), in currentColor, so the caller's token colours it. The caller imports the one
     export it draws (ICON_SEARCH ...), so an island carries only its own icons. Ui/Icon.astro is
     the server-rendered twin. -->
<script lang="ts">
  import { iconAttributes } from '../../lib/ui/icon';
  import type { UiIconShape } from '../../lib/ui/ui-icons';

  let {
    icon,
    size = 20,
    class: className = '',
    testid,
  }: { icon: UiIconShape; size?: number; class?: string; testid?: string } = $props();

  const attributes = $derived(iconAttributes(icon, size));
</script>

<svg class={['shrink-0', className]} {...attributes} data-testid={testid}>
  <!-- eslint-disable-next-line svelte/no-at-html-tags -- a constant from ui-icons.ts, never user text -->
  {@html icon.body}
</svg>
