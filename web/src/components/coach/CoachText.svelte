<!-- web/src/components/coach/CoachText.svelte -->
<!-- Follow-up: spell tokens carry a native title only; the site has no spell tooltip component (the
     planner's is talent-specific), so a shared one is a separate piece of work.
     A line of coach prose with its spell and item names as tokens (spec 3.1, tenets 2 and 3):
     a spell is its 16 px icon and name, an item opens the shared item tooltip. A name the build's
     dictionary does not hold stays plain text. -->
<script lang="ts">
  import { tokenizeNames } from '../../lib/coach/name-tokens';
  import type { CoachNames } from '../../lib/coach/names';
  import ItemHover from '../ItemHover.svelte';
  import SpellIcon from './SpellIcon.svelte';

  let {
    text,
    names,
    build,
    classSlug,
  }: { text: string; names: CoachNames; build: string; classSlug: string } = $props();

  const parts = $derived(tokenizeNames(text, names.dictionary));
</script>

{#each parts as part, index (index)}
  {#if !('entry' in part)}{part.text}{:else if part.entry.kind === 'spell'}<span
      class="inline-flex items-center gap-1 align-bottom"
      data-testid="coach-name-spell"
      ><SpellIcon url={names.iconUrl(part.entry.id)} size={16} name={part.text} />{part.text}</span
    >{:else}<ItemHover itemId={part.entry.id} {classSlug} {build} />{/if}
{/each}
