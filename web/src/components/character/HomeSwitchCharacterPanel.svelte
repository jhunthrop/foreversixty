<!-- web/src/components/character/HomeSwitchCharacterPanel.svelte -->
<!-- Home rebuild spec §3.B.4: the signed-in hero's right column. Reuses
     `CharacterSwitchList.svelte` -- "already exactly this list" per the spec -- rather than
     forking it: the mock's own "N upgrades" trailing stat is the same not-yet-available gap
     named in §3.B.2/§3.B.3 (no source joins a character's worn gear against a BiS list
     yet), so every non-current row's trailing slot stays exactly what CharacterSwitchList
     already draws (the "Switch" action) instead of inventing a number. Mounted by
     HomeAccountPanel.svelte into index.astro's `home-switch-character-slot`, the same
     dynamic-import-once-signed-in trick that island already uses for its own module, so a
     signed-out page never fetches this component either. -->
<script lang="ts">
  import type { Me, MeCharacter } from '../../lib/account/api';
  import CharacterSwitchList from './CharacterSwitchList.svelte';
  import { homePanelCopy } from '../../lib/home-panel-copy';

  let {
    me,
    currentKey,
    onswitch,
  }: { me: Me; currentKey: string | null; onswitch: (character: MeCharacter) => void } = $props();

  /** The hero listed first, matching the mock's own row order -- every other character
   *  follows in the account's own order. */
  const ordered = $derived.by(() => {
    const hero = me.characters.find((c) => c.key === currentKey);
    if (hero === undefined) return me.characters;
    return [hero, ...me.characters.filter((c) => c.key !== currentKey)];
  });
</script>

<div
  class="bg-raised/85 border-line rounded-panel flex min-h-[168px] min-w-0 flex-col gap-1 border px-[18px] py-[14px] [grid-area:1/1]"
  data-testid="home-switch-character-panel"
>
  <div class="mb-1 flex items-center justify-between">
    <span class="label">{homePanelCopy.switchCharacterLabel}</span>
    <a href="/account#add-character" class="text-nav text-[12px] font-bold tracking-[0.06em] uppercase">
      {homePanelCopy.addOneCharacter}
    </a>
  </div>
  <CharacterSwitchList characters={ordered} {currentKey} {onswitch} descriptor="full" />
</div>
