<!-- web/src/components/ui/ActionLink.svelte -->
<!-- A standalone action as a link (polish spec 4.2): gold, a trailing chevron, 44 px tall, no
     underline. A real anchor with an `href`; an action with no destination is a button row,
     not this. Extra attributes (`data-testid`, `target`) pass through. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { HTMLAnchorAttributes } from 'svelte/elements';
  import { ACTION_LINK, ACTION_LINK_CHEVRON } from '../../lib/ui/link';
  import Icon from './Icon.svelte';
  import { ICON_CHEVRON_RIGHT } from '../../lib/ui/ui-icons';

  let {
    href,
    children,
    class: className = '',
    ...rest
  }: HTMLAnchorAttributes & { href: string; children: Snippet } = $props();
</script>

<a {href} class="{ACTION_LINK} {className}" {...rest}>
  {@render children()}<Icon icon={ICON_CHEVRON_RIGHT} size={16} class={ACTION_LINK_CHEVRON} />
</a>
