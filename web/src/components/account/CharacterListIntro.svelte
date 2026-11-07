<!-- web/src/components/account/CharacterListIntro.svelte -->
<!-- The Battle.net line under /account's Characters Plate (premium spec 2026-11-08 section 7.2): where
     the characters come from, with the addon and paste links. Static copy, so the loading state draws
     it too (AccountLoading): the ready page's largest text is already painted in the first frame. -->
<script lang="ts">
  import { characterListCopy } from '../../lib/account/character-list-copy';
  import { LINK } from '../../lib/ui/primitives';
</script>

<p class="text-small text-muted max-w-[80ch] pt-3">
  {characterListCopy.introBattlenetLine}
  <a class={LINK} href="/setup">{characterListCopy.introInstallAddonLink}</a>
  {characterListCopy.introAddonTail}
  <a class={LINK} href={characterListCopy.pasteHref}>{characterListCopy.introPasteLink}</a>
  {characterListCopy.introPasteTail}
</p>
