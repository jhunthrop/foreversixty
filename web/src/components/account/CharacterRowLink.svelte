<!-- web/src/components/account/CharacterRowLink.svelte -->
<!-- The account Characters list's per-row link (design/specs/2026-11-04-account-characters.md
     section 3): "Open in simulator" for a character with a build, built from
     `MeCharacter.build` alone (no `sim-input` fetch: a list of N characters must not make N
     fetches just to draw its rows), else the short "Paste an export" link. Its row says
     "No gear yet"; the current-character card explains why once. -->
<script lang="ts">
  import { simArmoryHref } from '../../lib/handoff-links';
  import { handoffCopy } from '../../lib/handoff-copy';
  import type { MeCharacter } from '../../lib/account/api';
  import { characterListCopy } from '../../lib/account/character-list-copy';

  let { character }: { character: MeCharacter } = $props();

  const LINK = 'inline-flex min-h-11 items-center text-[13px] font-semibold whitespace-nowrap md:min-h-0';
</script>

{#if character.build !== undefined}
  <a class={LINK} href={simArmoryHref(character.key)} data-testid="character-open-sim">
    {handoffCopy.openInSimulator}
  </a>
{:else}
  <a class={LINK} href={handoffCopy.pasteHref} data-testid="character-paste-export">
    {characterListCopy.pasteAnExport}
  </a>
{/if}
