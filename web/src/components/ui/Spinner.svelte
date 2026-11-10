<!-- web/src/components/ui/Spinner.svelte -->
<!-- The 14 px ring a busy button shows in place of its glyph (polish spec 4.2, Loading). Still
     under reduced motion: the open ring alone says "working". -->
<script lang="ts">
  import Icon from './Icon.svelte';
  import { ICON_SPINNER } from '../../lib/ui/ui-icons';
</script>

<Icon icon={ICON_SPINNER} size={14} class="animate-spin motion-reduce:animate-none" testid="spinner" />
