<!-- web/src/components/CurrentCharacterBar.svelte -->
<!-- The current-character chip for a page that has no character state of its own: the
     account page, a character page, the addon page. CurrentCharacterChip is a pure render of
     one prop, and on /planner and /sim* the page that mounts it owns the pointer; here
     nothing else does, so this reads it once on mount and forgets it on request.
     `refresh` lets a sibling on the same page (the addon page's paste box) say "I just wrote
     a new pointer" without the two sharing a store. -->
<script lang="ts">
  import { fetchMeOnce } from '../lib/account/api';
  import {
    CURRENT_CHARACTER_CHANGED,
    clearCurrent,
    readCurrent,
    type CurrentCharacter,
  } from '../lib/current-character';
  import { guildRankLabel } from '../lib/characters';
  import CurrentCharacterChip from './CurrentCharacterChip.svelte';

  let { hasOwnPasteBox = false }: { hasOwnPasteBox?: boolean } = $props();

  let current = $state<CurrentCharacter | null>(null);
  // Spec 2026-09-22 §7.5: "Iron Vanguard · Officer", built once the matching /v1/me
  // character (by key -- see the plan's Ruling 2 for why 'armory' is the one source whose
  // `ref` is a character key) answers. Empty otherwise, which CurrentCharacterChip reads as
  // "nothing to show".
  let guildLine = $state('');

  $effect(() => {
    const read = (): void => {
      current = readCurrent();
      guildLine = '';
      if (current === null || current.source !== 'armory') return;
      const key = current.ref;
      // fetchMeOnce is the shared, deduped /v1/me promise every account island on the page
      // already uses (Account.svelte, AddonPasteBox.svelte) -- this is what makes the read
      // "no new fetch" per the spec, not a second request of its own.
      void fetchMeOnce()
        .then((me) => {
          if (current?.ref !== key) return; // a newer read landed first
          const character = me?.characters.find((c) => c.key === key);
          if (character?.guild === undefined) return;
          const rank = character.guild.rank === undefined ? '' : ` · ${guildRankLabel(character.guild.rank)}`;
          guildLine = `${character.guild.name}${rank}`;
        })
        .catch(() => {
          guildLine = '';
        });
    };
    read();
    window.addEventListener(CURRENT_CHARACTER_CHANGED, read);
    return () => window.removeEventListener(CURRENT_CHARACTER_CHANGED, read);
  });

  function forget(): void {
    clearCurrent();
    current = null;
    guildLine = '';
  }
</script>

<CurrentCharacterChip {current} {hasOwnPasteBox} {guildLine} onforget={forget} />
