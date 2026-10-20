<!-- web/src/components/character/CharacterHeaderSignedOut.svelte -->
<!-- The one shared signed-out band (character header spec 2026-10-11 section 3): a single muted
     line between hairlines, 44 px, no button and no link. The nav character selector is the
     only sign-in control, so a page that mounts the header adds no sign-in prose of its own. -->
<script lang="ts">
  import { characterHeaderCopy } from '../../lib/character-header/copy';

  /** Set for a signed-in account with no characters: its name, so the band says who is signed in. */
  let { account }: { account?: string } = $props();
</script>

{#if account === undefined}
  <p class="ch-out" data-testid="character-header-signed-out">{characterHeaderCopy.signedOut}</p>
{:else}
  <p class="ch-out" data-testid="character-header-no-characters">
    {characterHeaderCopy.noCharacters(account)}
    <a class="ch-out-link" href={characterHeaderCopy.getSetUpHref}>{characterHeaderCopy.getSetUp}</a>
  </p>
{/if}

<style>
  .ch-out-link {
    margin-inline-start: 8px;
    color: var(--color-strong);
    text-decoration: underline;
    text-underline-offset: 3px;
  }
  .ch-out {
    display: flex;
    align-items: center;
    box-sizing: border-box;
    min-height: 44px;
    margin: 0;
    font-size: 13px;
    color: var(--color-muted);
    border-block: 1px solid var(--color-line-soft);
  }
</style>
