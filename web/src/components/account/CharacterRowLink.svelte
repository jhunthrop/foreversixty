<!-- web/src/components/account/CharacterRowLink.svelte -->
<!-- The account Characters list's per-row action (spec 2026-09-22 §3.4): built from
     `MeCharacter.build` alone, no `sim-input` fetch -- unlike `CharacterHandoffLinks`
     (still used by the hero band and the character page, one row each, where a single
     fetch is fine), a list of N characters must not make N fetches just to draw its
     rows. There is no fetch-free "Open in planner" for a stored character (the planner
     only ever restores a `'code'`/`'addon'` pointer -- `current-character-planner.ts`'s
     own header comment -- never an `'armory'` one by ref), so this offers only the
     simulator link; the planner's own "Open in planner" link, once the character is
     loaded there, covers the rest. -->
<script lang="ts">
  import { simArmoryHref } from '../../lib/handoff-links';
  import { handoffCopy } from '../../lib/handoff-copy';
  import type { MeCharacter } from '../../lib/account/api';

  let { character }: { character: MeCharacter } = $props();

  const LINK = 'inline-flex min-h-11 items-center text-[13px] font-semibold md:min-h-0';
</script>

{#if character.build !== undefined}
  <a class={LINK} href={simArmoryHref(character.key)} data-testid="character-open-sim">
    {handoffCopy.openInSimulator}
  </a>
{:else}
  <span class="text-muted text-[13px]" data-testid="character-needs-addon">
    {handoffCopy.needsExportLead}
    <a class="text-text inline-flex min-h-11 items-center underline md:min-h-0" href={handoffCopy.pasteHref}>
      {handoffCopy.needsExportPasteLink}
    </a>
  </span>
{/if}
