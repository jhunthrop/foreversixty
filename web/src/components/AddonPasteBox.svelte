<!-- web/src/components/AddonPasteBox.svelte -->
<!-- /addon's own paste box: prove a pasted export decodes with the same decoder the
     planner uses (planner/fs1.ts, read-only import — this component never writes to that
     module) and hand back the two places it goes next. Unlike the planner's own
     ImportBox.svelte, this page never builds a talent draft: it only proves the string is
     readable and links on with the raw code, so it needs no TalentIndex and no active
     build to reconcile against. -->
<script lang="ts">
  import { addonCopy } from '../lib/addon/copy';
  import { decodeFS1 } from '../lib/planner/fs1';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { plannerCodeHref, simCodeHref } from '../lib/handoff-links';

  let code = $state('');
  let error = $state<string | null>(null);
  let loaded = $state<string | null>(null);

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
  }
</script>

<section
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
        class="text-gold text-[13px] font-semibold"
        href={plannerCodeHref(loaded)}
        data-testid="addon-paste-planner"
      >
        {addonCopy.pasteOpenPlanner}
      </a>
      <a class="text-gold text-[13px] font-semibold" href={simCodeHref(loaded)} data-testid="addon-paste-sim">
        {addonCopy.pasteOpenSim}
      </a>
    </div>
  {/if}
</section>
