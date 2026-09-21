<!-- web/src/components/CurrentCharacterBar.svelte -->
<!-- The current-character chip for a page that has no character state of its own: the
     account page, a character page, the addon page. CurrentCharacterChip is a pure render of
     one prop, and on /planner and /sim* the page that mounts it owns the pointer; here
     nothing else does, so this reads it once on mount and forgets it on request.
     `refresh` lets a sibling on the same page (the addon page's paste box) say "I just wrote
     a new pointer" without the two sharing a store. -->
<script lang="ts">
  import {
    CURRENT_CHARACTER_CHANGED,
    clearCurrent,
    readCurrent,
    type CurrentCharacter,
  } from '../lib/current-character';
  import CurrentCharacterChip from './CurrentCharacterChip.svelte';

  let { hasOwnPasteBox = false }: { hasOwnPasteBox?: boolean } = $props();

  let current = $state<CurrentCharacter | null>(null);

  $effect(() => {
    const read = (): void => {
      current = readCurrent();
    };
    read();
    window.addEventListener(CURRENT_CHARACTER_CHANGED, read);
    return () => window.removeEventListener(CURRENT_CHARACTER_CHANGED, read);
  });

  function forget(): void {
    clearCurrent();
    current = null;
  }
</script>

<CurrentCharacterChip {current} {hasOwnPasteBox} onforget={forget} />
