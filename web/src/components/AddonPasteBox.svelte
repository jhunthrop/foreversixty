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
  import { decodeFS1 } from '../lib/planner/fs1';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { CURRENT_CHARACTER_CHANGED, writeCurrent } from '../lib/current-character';
  import { plannerCodeHref, simCodeHref } from '../lib/handoff-links';
  import AddonPasteSave from './AddonPasteSave.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  let code = $state('');
  let error = $state<string | null>(null);
  let loaded = $state<string | null>(null);
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
    // The export is now the site's current character, so the planner and every simulator
    // tab open on it without a second paste. The label is the class alone: the spec needs
    // the talent file, which the tool that opens it loads and then writes a fuller label.
    const classSlug = result.build.classSlug;
    writeCurrent({
      source: 'addon',
      ref: trimmed,
      label: classSlug.charAt(0).toUpperCase() + classSlug.slice(1),
      classSlug,
      savedAt: new Date().toISOString(),
    });
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
  }
</script>

<section
  id="paste"
  class="border-line bg-raised rounded-panel flex flex-col gap-3 border p-4"
  data-testid="addon-paste-box"
>
  <h2 class="section-title text-[15px]">{addonCopy.pasteTitle}</h2>
  <textarea
    class="border-line rounded-control bg-card-top min-h-11 w-full border px-2 py-1 font-mono text-[13px]"
    rows="2"
    placeholder={addonCopy.pastePlaceholder}
    aria-label={addonCopy.pasteTitle}
    bind:value={code}
    data-testid="addon-paste-code"></textarea>
  <button
    type="button"
    class={SECONDARY_BUTTON_FIXED}
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
        href={simCodeHref(loaded)}
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
        <AddonPasteSave {signedIn} code={loaded} />
      {/key}
    {/if}
  {/if}
</section>
