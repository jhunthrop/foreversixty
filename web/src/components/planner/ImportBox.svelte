<!-- web/src/components/planner/ImportBox.svelte -->
<!-- The paste box. One textarea, one button, and every note the import produced shown
     together: the order is always approximated, and sometimes a talent could not be
     placed or the export is a build behind. -->
<script lang="ts">
  import { addonCopy } from '../../lib/addon/copy';
  import { importFromAddon, type ImportOutcome } from '../../lib/addon/import';
  import { currentCharacterCopy } from '../../lib/current-character-copy';
  import type { TalentIndex } from '../../lib/planner/rules';
  import { SECONDARY_BUTTON } from '../../lib/planner/styles';
  import type { Gear } from '../../lib/planner/types';
  import { rowLink } from '../../lib/report/format';

  let {
    talents,
    activeBuild,
    onimport,
    onwrongclass = undefined,
    phone = false,
    class: className = '',
  }: {
    talents: TalentIndex;
    activeBuild: string;
    /**
     * `pastedCode` (Task 10) is the export string as pasted, trimmed the same way `submit`
     * decodes it -- the current-character bridge's own 'addon' pointer is written from
     * this exact string, never a re-encoding of the parsed build, so a caller reading it
     * back later decodes byte-for-byte what the player pasted.
     */
    onimport: (
      build: { classSlug: string; raceSlug: string; order: number[]; gear: Gear; characterName?: string },
      pastedCode: string,
    ) => void;
    /**
     * An export for a class other than the loaded one. When the mounting page can switch
     * class it takes the code here and imports once that class's talents arrive, instead
     * of the box telling the player to switch by hand; without this callback the box shows
     * the wrong-class message as before (Top Gear's inline box has one class).
     */
    onwrongclass?: (classSlug: string, pastedCode: string) => void;
    /**
     * Below md the standalone planner folds this behind a native, closed-by-default
     * disclosure (design loop, planner round): a visitor reaches Gear without scrolling
     * past a paste box most never touch. False for every other mount -- Top Gear's inline
     * "add a build" embeds this same component in its own layout and never grows this
     * collapse of its own.
     */
    phone?: boolean;
    /** Merged onto the root element, so the caller's own responsive grid can place this
     *  panel without either side knowing about the other's layout. */
    class?: string;
  } = $props();

  let code = $state('');
  let outcome = $state<ImportOutcome | null>(null);

  const notes = $derived(outcome !== null && outcome.ok ? outcome.notes : []);
  const failure = $derived(outcome !== null && !outcome.ok ? outcome.message : null);

  function submit(): void {
    const trimmed = code.trim();
    const result = importFromAddon(trimmed, talents, activeBuild);
    if (!result.ok && result.wrongClass !== undefined && onwrongclass !== undefined) {
      outcome = { ok: false, message: addonCopy.importSwitchingClass(result.wrongClass) };
      onwrongclass(result.wrongClass, trimmed);
      return;
    }
    outcome = result;
    if (result.ok) {
      onimport(
        {
          classSlug: result.classSlug,
          raceSlug: result.raceSlug,
          order: result.order,
          gear: result.gear,
          ...(result.characterName === undefined ? {} : { characterName: result.characterName }),
        },
        trimmed,
      );
    }
  }
</script>

<svelte:element
  this={phone ? 'details' : 'section'}
  class="border-line bg-raised rounded-panel flex flex-col gap-3 border p-4 {className}"
  data-testid="import-box"
>
  {#if phone}
    <summary class="label text-nav flex min-h-11 cursor-pointer list-none items-center justify-between gap-3">
      {addonCopy.importTitle}
      <span class="when-closed text-muted" aria-hidden="true">Show</span>
      <span class="when-open text-muted" aria-hidden="true">Hide</span>
    </summary>
  {:else}
    <h2 class="section-title text-[15px]">{addonCopy.importTitle}</h2>
  {/if}
  <textarea
    class="border-line rounded-control bg-card-top min-h-11 w-full border px-2 py-1 font-mono text-[13px]"
    rows="2"
    placeholder={addonCopy.importPlaceholder}
    aria-label={addonCopy.importTitle}
    bind:value={code}
    data-testid="import-code"></textarea>
  <button
    type="button"
    class="{SECONDARY_BUTTON} border-line-warm-strong text-strong self-start px-4"
    disabled={code.trim() === ''}
    onclick={submit}
    data-testid="import-submit"
  >
    {addonCopy.importAction}
  </button>
  {#if failure !== null}
    <!-- `text-strong`, not a red tint: no such colour exists in the design token set for a
         plain form error, and SharePanel's own failure panel (right below Share) reads the
         same way -- a bordered, `role="alert"` block in the body colour, not a hue of its
         own. -->
    <p class="text-strong text-[13px]" data-testid="import-error" role="alert">{failure}</p>
  {/if}
  {#each notes as note (note)}
    <p class="text-muted text-[13px]" data-testid="import-note">{note}</p>
  {/each}
  <p class="text-muted text-[13px]">
    <a href="/setup" class="{rowLink} text-nav" data-testid="import-get-addon"
      >{currentCharacterCopy.getTheAddon}</a
    >
  </p>
</svelte:element>
