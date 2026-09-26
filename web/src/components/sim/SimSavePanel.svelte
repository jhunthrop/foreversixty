<!-- web/src/components/sim/SimSavePanel.svelte -->
<!-- Extracted from SimView.svelte (2026-09-26 layout pass -- SimView had grown past this
     lane's 800-line ceiling once the Run block, hero card and example card went in, and
     this panel was the largest single self-contained piece: its own state, its own three
     inline states -- open button, the inline title form, the saved link -- and no reach
     into anything else the view owns beyond the result it is saving and the store's own
     `save` call). Disabled until there is a result, an inline title field pre-filled with
     the report title rather than a dialog, and the saved link shown in place -- the page
     never navigates away from the result it just saved (Task 17). -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import type { SimResult } from '../../lib/sim/types';
  import { BUSY_CLASS } from '../../lib/ui/busy';

  let {
    result,
    reportTitle,
    onsave,
  }: {
    result: SimResult | null;
    /** Read once, when the inline form opens -- not bound live, the same way the file this
     *  was extracted from never re-synced an open form's title against a later rename. */
    reportTitle: string;
    /** Returns the new sim's id, or null on failure. */
    onsave: (title: string) => Promise<string | null>;
  } = $props();

  let saveOpen = $state(false);
  let saveTitle = $state('');
  let saving = $state(false);
  let saveFailed = $state(false);
  let savedUrl = $state<string | null>(null);
  let savedLinkCopied = $state(false);

  $effect(() => {
    void result;
    saveOpen = false;
    saveFailed = false;
    savedUrl = null;
    savedLinkCopied = false;
  });

  // A result the run loop reports as stopped rather than finished (`sim/api`'s additive
  // `aborted`) has nothing complete to save -- the button stays disabled and says why,
  // rather than saving a partial run under a title the player chose for a real result.
  const canSave = $derived(result !== null && result.aborted !== true);

  function openSaveForm(): void {
    saveTitle = reportTitle;
    saveFailed = false;
    savedUrl = null;
    saveOpen = true;
  }

  function cancelSave(): void {
    saveOpen = false;
    saveFailed = false;
  }

  async function confirmSave(): Promise<void> {
    saving = true;
    saveFailed = false;
    const id = await onsave(saveTitle);
    saving = false;
    if (id === null) {
      saveFailed = true;
      return;
    }
    saveOpen = false;
    savedUrl = `${window.location.origin}/sim/${id}`;
  }

  async function copySavedLink(): Promise<void> {
    if (savedUrl === null) return;
    try {
      await navigator.clipboard.writeText(savedUrl);
      savedLinkCopied = true;
    } catch {
      savedLinkCopied = false;
    }
  }
</script>

<div class="mx-[18px] flex flex-wrap items-center gap-3 md:mx-0" data-testid="sim-save">
  {#if savedUrl !== null}
    <label class="sr-only" for="sim-save-link">{simCopy.savedLinkLabel}</label>
    <input
      id="sim-save-link"
      type="text"
      readonly
      value={savedUrl}
      class="border-line-warm rounded-control bg-raised text-text h-11 min-w-0 flex-1 border px-3 text-[14px] md:max-w-[420px]"
      data-testid="sim-save-link"
      onclick={(event) => event.currentTarget.select()}
    />
    <button
      type="button"
      class="border-line-warm rounded-control text-nav label min-h-11 border px-4"
      onclick={() => void copySavedLink()}
      data-testid="sim-save-copy"
    >
      {savedLinkCopied ? simCopy.copied : simCopy.copyLink}
    </button>
    <a
      class="border-line-warm rounded-control text-nav label inline-flex min-h-11 items-center border px-4"
      href={savedUrl}
      target="_blank"
      rel="noopener"
      data-testid="sim-open-new-tab"
    >
      {simCopy.openInNewTab}
    </a>
  {:else if saveOpen}
    <label class="flex flex-col gap-1">
      <span class="label text-muted">{simCopy.saveTitleLabel}</span>
      <input
        type="text"
        bind:value={saveTitle}
        class="border-line-warm rounded-control bg-raised text-text h-11 w-[260px] border px-3 text-[14px]"
        data-testid="sim-save-title"
      />
    </label>
    <button
      type="button"
      class={`border-line-warm-strong rounded-control bg-card-top text-strong label min-h-11 border px-5 disabled:opacity-50 ${saving ? BUSY_CLASS : ''}`}
      disabled={saving}
      aria-busy={saving}
      onclick={() => void confirmSave()}
      data-testid="sim-save-confirm"
    >
      {simCopy.saveAction}
    </button>
    <button
      type="button"
      class="border-line-warm rounded-control text-nav label min-h-11 border px-4"
      onclick={cancelSave}
      data-testid="sim-save-cancel"
    >
      {simCopy.cancel}
    </button>
    {#if saveFailed}
      <span role="alert" class="text-strong text-[13px]" data-testid="sim-save-error">
        {simCopy.saveFailed}
      </span>
    {/if}
  {:else}
    <button
      type="button"
      class="border-line-warm-strong rounded-control bg-card-top text-strong label min-h-11 border px-5 disabled:opacity-50"
      disabled={!canSave}
      onclick={openSaveForm}
      data-testid="sim-save-open"
    >
      {simCopy.saveThisSim}
    </button>
    <!-- Task 7: the disabled reason, said plainly (Task 3's pattern), never a title=. -->
    {#if result !== null && !canSave}
      <p class="text-muted text-[12px]" data-testid="sim-save-disabled-note">
        {simCopy.saveAbortedDisabled}
      </p>
    {/if}
  {/if}
</div>
