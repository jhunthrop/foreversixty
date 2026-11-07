<!-- web/src/components/character/CharacterHeaderSignedOut.svelte -->
<!-- The one shared signed-out band (character header spec 2026-10-11 section 3): a single muted
     Small line, 44 px, no box, no rule and no button. The nav character selector is the only
     sign-in control, so a page that mounts the header adds no sign-in prose of its own. A
     signed-in account with no characters reads its name and one link to get set up. -->
<script lang="ts">
  import { characterHeaderCopy } from '../../lib/character-header/copy';
  import { LINK } from '../../lib/ui/primitives';

  /** Set for a signed-in account with no characters: its name, so the band says who is signed in. */
  let { account }: { account?: string } = $props();

  const LINE = 'text-small text-muted m-0 flex min-h-11 flex-wrap items-center gap-x-2';
</script>

{#if account === undefined}
  <p class={LINE} data-testid="character-header-signed-out">{characterHeaderCopy.signedOut}</p>
{:else}
  <p class={LINE} data-testid="character-header-no-characters">
    {characterHeaderCopy.noCharacters(account)}
    <a class={LINK} href={characterHeaderCopy.getSetUpHref}>{characterHeaderCopy.getSetUp} ›</a>
  </p>
{/if}
