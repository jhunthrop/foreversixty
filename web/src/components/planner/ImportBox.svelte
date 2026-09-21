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
  }: {
    talents: TalentIndex;
    activeBuild: string;
    onimport: (build: { classSlug: string; raceSlug: string; order: number[]; gear: Gear }) => void;
  } = $props();

  let code = $state('');
  let outcome = $state<ImportOutcome | null>(null);

  const notes = $derived(outcome !== null && outcome.ok ? outcome.notes : []);
  const failure = $derived(outcome !== null && !outcome.ok ? outcome.message : null);

  function submit(): void {
    const result = importFromAddon(code.trim(), talents, activeBuild);
    outcome = result;
    if (result.ok) {
      onimport({
        classSlug: result.classSlug,
        raceSlug: result.raceSlug,
        order: result.order,
        gear: result.gear,
      });
    }
  }
</script>

<section class="border-line bg-raised rounded-panel flex flex-col gap-3 border p-4" data-testid="import-box">
  <h2 class="section-title text-[15px]">{addonCopy.importTitle}</h2>
  <textarea
    class="border-line rounded-control bg-card-top min-h-11 w-full border px-2 py-1 font-mono text-[13px]"
    rows="2"
    placeholder={addonCopy.importPlaceholder}
    aria-label={addonCopy.importTitle}
    bind:value={code}
    data-testid="import-code"></textarea>
  <button
    type="button"
    class={SECONDARY_BUTTON}
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
    <a href="/addon" class="{rowLink} text-nav" data-testid="import-get-addon"
      >{currentCharacterCopy.getTheAddon}</a
    >
  </p>
</section>
