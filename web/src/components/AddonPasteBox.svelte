<!-- web/src/components/AddonPasteBox.svelte -->
<!-- /addon's own paste box: prove a pasted export decodes with the same decoder the
     planner uses (planner/fs1.ts, read-only import — this component never writes to that
     module) and hand back the two places it goes next. Unlike the planner's own
     ImportBox.svelte, this page never builds a talent draft: it only proves the string is
     readable and links on with the raw code, so it needs no TalentIndex and no active
     build to reconcile against. -->
<script lang="ts">
  import { fetchMeOnce } from '../lib/account/api';
  import { addonCopy } from '../lib/addon/copy';
  import { ADDON_PASTE_STATUS_MIN_H } from '../lib/addon/paste-layout';
  import { decodeFS1, type FS1Build } from '../lib/planner/fs1';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { recordPastedExport } from '../lib/character-selector/paste';
  import { plannerCodeHref, simExportHref } from '../lib/handoff-links';
  import AddonPasteSave from './AddonPasteSave.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  /** False where the mounting page already titles the box (the /setup disclosure's own summary). */
  let { heading = true }: { heading?: boolean } = $props();

  let code = $state('');
  let error = $state<string | null>(null);
  let loaded = $state<string | null>(null);
  // The `who=` section of the loaded code, when it carries one: the save form starts from it.
  let loadedCharacter = $state<FS1Build['character']>(undefined);
  // null while fetchMeOnce is still resolving -- the save/hint block renders nothing until
  // it is known, rather than flashing the signed-out hint first (same idiom as Account.svelte's
  // own status: 'loading' | 'ready' | 'failed', simplified to the one fact this needs).
  let signedIn = $state<boolean | null>(null);

  $effect(() => {
    void fetchMeOnce()
      .then((me) => {
        signedIn = me !== null;
      })
      .catch(() => {
        signedIn = false;
      });
  });

  function submit(): void {
    const trimmed = code.trim();
    const result = decodeFS1(trimmed);
    if (!result.ok) {
      error = result.message;
      loaded = null;
      return;
    }
    error = null;
    loaded = trimmed;
    loadedCharacter = result.build.character;
    // An export that names a character is the player choosing it: it becomes the current
    // character (the account's own, when it holds one by that name). One that names nobody
    // is looked at and changes nothing -- a class word is never made the pointer.
    void recordPastedExport(trimmed);
  }
</script>

<section
  id="paste"
  class="border-line bg-raised rounded-panel flex flex-col gap-3 border p-4"
  data-testid="addon-paste-box"
>
  {#if heading}<h2 class="section-title text-[15px]">{addonCopy.pasteTitle}</h2>{/if}
  <textarea
    class="border-line rounded-control bg-card-top min-h-11 w-full border px-2 py-1 font-mono text-[13px]"
    rows="2"
    placeholder={addonCopy.pastePlaceholder}
    aria-label={addonCopy.pasteTitle}
    bind:value={code}
    data-testid="addon-paste-code"></textarea>
  <button
    type="button"
    class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong self-start px-4"
    disabled={code.trim() === ''}
    onclick={submit}
    data-testid="addon-paste-submit"
  >
    {addonCopy.pasteAction}
  </button>
  {#if error !== null}
    <p class="text-strong text-[13px]" data-testid="addon-paste-error" role="alert">{error}</p>
  {/if}
  {#if loaded !== null}
    <!-- The ready state of this box: a pasted export that decoded, so the two places it
         goes next and the save block below it. `.reveal` goes on this wrapper rather than
         the <section>, which is present from first paint and never fades (design
         2026-09-22 spec section 1.4). It repeats the section's `flex flex-col gap-3` so
         the two children keep the spacing they had as the section's own children. -->
    <div class="reveal flex flex-col gap-3">
      <div class="flex flex-wrap gap-3">
        <a
          class="text-gold inline-flex min-h-11 items-center text-[13px] font-semibold md:min-h-0"
          href={plannerCodeHref(loaded)}
          data-testid="addon-paste-planner"
        >
          {addonCopy.pasteOpenPlanner}
        </a>
        <a
          class="text-gold inline-flex min-h-11 items-center text-[13px] font-semibold md:min-h-0"
          href={simExportHref(loaded)}
          data-testid="addon-paste-sim"
        >
          {addonCopy.pasteOpenSim}
        </a>
      </div>
      {#if signedIn === null}
        <!-- One row, not the ui default of three: ADDON_PASTE_STATUS_MIN_H reserves the
           shorter (sign-in hint) branch's real height, and Skeleton's own row content
           would otherwise dominate that min-height once more than one row is stacked
           (two h-4 rows plus their gap already exceed 19.5px on their own), silently
           reopening the shrink this constant exists to prevent. -->
        <Skeleton lines={1} minHeight={ADDON_PASTE_STATUS_MIN_H} testid="addon-paste-status-skeleton" />
      {:else}
        {#key loaded}
          <AddonPasteSave {signedIn} code={loaded} character={loadedCharacter} />
        {/key}
      {/if}
    </div>
  {/if}
</section>
