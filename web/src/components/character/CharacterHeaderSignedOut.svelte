<!-- web/src/components/character/CharacterHeaderSignedOut.svelte -->
<!-- The one shared signed-out band (character header spec 2026-10-11 section 3): a single muted
     line between hairlines, 44 px, no button and no link. The nav character selector is the
     only sign-in control, so a page that mounts the header adds no sign-in prose of its own. -->
<script lang="ts">
  import { characterHeaderCopy } from '../../lib/character-header/copy';
</script>

<p class="ch-out" data-testid="character-header-signed-out">{characterHeaderCopy.signedOut}</p>

<style>
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
