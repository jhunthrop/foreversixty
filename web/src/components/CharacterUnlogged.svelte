<!-- web/src/components/CharacterUnlogged.svelte -->
<!-- The character page for a character the API has no record of (a 404): the name from the
     address, "no logs yet" where the score would be, and how to get the character onto the
     site, since a character search lands here for anyone (Task L8.4 ruling 4). -->
<script lang="ts">
  import CharacterIdentity from './character/CharacterIdentity.svelte';
  import CharacterInGame from './CharacterInGame.svelte';
  import CharacterRatingPanel from './CharacterRatingPanel.svelte';
  import { characterHeaderCopy } from '../lib/character-profile';
  import { characterKey, type CharacterPath } from '../lib/characters';

  let { path, name, owned = false }: { path: CharacterPath; name: string; owned?: boolean } = $props();

  /** The ruleset-and-region chip, "Hardcore US" (Task L8.4 ruling 1); a logged character's is
   *  ProfileDescriptor's, with the same muted text, border token and line alignment. */
  const CHIP = 'pill text-muted border-line align-top';
</script>

<div class="reveal flex flex-col gap-[22px] md:gap-8" data-testid="character-unlogged" id="character">
  <!-- The same identity block as a logged character's, with the neutral crest (no class on
       file), so the two states read as one page. -order-3, the logged header's own (L8.4 fix
       round 3 item 1): the Forever Score card's section is -order-2 at every width, so a header
       with no order of its own sorted below it. -->
  <header class="-order-3 flex flex-col gap-1">
    <CharacterIdentity
      character={{
        key: characterKey(path.region, path.ruleset, path.slug),
        region: path.region,
        ruleset: path.ruleset,
        name,
      }}
      size="lg"
      descriptor="none"
      heading
      nameTestid="character-name"
    >
      {#snippet below()}
        <p class="text-muted text-[13px]">
          <span class={CHIP} data-testid="character-ruleset-chip"
            >{characterHeaderCopy.chip(path.ruleset, path.region)}</span
          >
        </p>
      {/snippet}
    </CharacterIdentity>
  </header>
  <CharacterRatingPanel {path} forever={null} {owned} {name} />
  <CharacterInGame forever={null} {name} {owned} />
</div>
