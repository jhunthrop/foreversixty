<!-- web/src/components/CurrentCharacterBar.svelte -->
<!-- The current-character chip for a page that has no character state of its own (/setup's
     step card): it reads the stored pointer, follows every write to it, adds the account
     guild line for an account character, and hands Forget to the chip. Switching character is
     the header selector's job alone (nav spec 2026-10-09); the simulator's "Jump to" spine bar
     that used to live here as a second mode was retired on 2026-10-31. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import { fetchMeOnce } from '../lib/account/api';
  import {
    clearCurrent,
    onCurrentCharacterChange,
    readCurrent,
    type CurrentCharacter,
  } from '../lib/current-character';
  import { guildRankLabel } from '../lib/characters';
  import CurrentCharacterChip from './CurrentCharacterChip.svelte';

  let {
    hasOwnPasteBox = false,
  }: {
    /** The page has its own paste box, so the chip says nothing while no character is stored. */
    hasOwnPasteBox?: boolean;
  } = $props();

  let current = $state<CurrentCharacter | null>(null);
  let guildLine = $state('');

  $effect(() => {
    const read = (): void => {
      const next = readCurrent();
      current = next;
      guildLine = '';
      if (next === null || next.source !== 'armory') return;
      const key = next.ref;
      void fetchMeOnce()
        .then((me) => {
          if (untrack(() => current)?.ref !== key) return;
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
    return onCurrentCharacterChange(read);
  });

  function forget(): void {
    clearCurrent();
    current = null;
    guildLine = '';
  }
</script>

<CurrentCharacterChip {current} {hasOwnPasteBox} {guildLine} onforget={forget} />
