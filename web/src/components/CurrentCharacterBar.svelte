<!-- web/src/components/CurrentCharacterBar.svelte -->
<!-- The current-character chip for a page that has no character state of its own: the
     account page, a character page, the addon page. CurrentCharacterChip is a pure render of
     one prop, and on /planner and /sim* the page that mounts it owns the pointer; here
     nothing else does, so this reads it once on mount and forgets it on request.
     `refresh` lets a sibling on the same page (the addon page's paste box) say "I just wrote
     a new pointer" without the two sharing a store.

     `compact` is set by Account.svelte's account mode (spec 2026-09-22 F2): the page already
     has a full Characters list a few hundred px below, so the pointer chip beside the title
     shows only when a pointer exists and renders nothing otherwise, never the "no character"
     sentence. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import { fetchMeOnce } from '../lib/account/api';
  import {
    CURRENT_CHARACTER_CHANGED,
    clearCurrent,
    readCurrent,
    type CurrentCharacter,
  } from '../lib/current-character';
  import { guildRankLabel } from '../lib/characters';
  import CurrentCharacterChip from './CurrentCharacterChip.svelte';

  let { hasOwnPasteBox = false, compact = false }: { hasOwnPasteBox?: boolean; compact?: boolean } = $props();

  let current = $state<CurrentCharacter | null>(null);
  // Spec 2026-09-22 §7.5: "Iron Vanguard · Officer", built once the matching /v1/me
  // character (by key -- see the plan's Ruling 2 for why 'armory' is the one source whose
  // `ref` is a character key) answers. Empty otherwise, which CurrentCharacterChip reads as
  // "nothing to show".
  let guildLine = $state('');

  // `read` writes `current` and must never read it back through the signal: an effect
  // that reads what it writes re-runs itself, and with a fresh object from readCurrent()
  // on every run that is an infinite loop (effect_update_depth_exceeded on /account with an
  // armory pointer stored, which starved every other island on the page). The local
  // `next` is the only thing the closure consults.
  $effect(() => {
    const read = (): void => {
      const next = readCurrent();
      current = next;
      guildLine = '';
      if (next === null || next.source !== 'armory') return;
      const key = next.ref;
      // fetchMeOnce is the shared, deduped /v1/me promise every account island on the page
      // already uses (Account.svelte, AddonPasteBox.svelte) -- this is what makes the read
      // "no new fetch" per the spec, not a second request of its own.
      void fetchMeOnce()
        .then((me) => {
          if (untrack(() => current)?.ref !== key) return; // a newer read landed first
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

<CurrentCharacterChip {current} hasOwnPasteBox={hasOwnPasteBox || compact} {guildLine} onforget={forget} />
