<!-- web/src/components/sim/tools/SaveSimForm.svelte -->
<!-- Design 8 / contract 10.6's browser-result save ("saved sims of every kind are public at
     /sim/<id>"), in the one form every bulk tool page shares: closed, open (a title to
     confirm), then saved (the link, copyable, and a new tab). Extracted out of
     ComboResults.svelte (Task 16 built it for Top Gear/talents/drops) so a weights result
     gets the identical flow rather than a second, drifting copy of the same six pieces of
     state -- Top Gear, talents, drops and weights all save through this one component now. -->
<script lang="ts">
  import { SECONDARY_BUTTON } from '../../../lib/planner/styles';
  import { simCopy } from '../../../lib/sim/copy';

  let {
    onsave,
    titleFor,
    canSave,
  }: {
    /** `store.save`, kept a callback rather than a `store` prop -- this component only
     *  ever needs the one method. */
    onsave: (title: string) => Promise<string | null>;
    /** The title the form opens pre-filled with -- composed by the caller from its own
     *  result, never typed here. */
    titleFor: string;
    /** A stopped/aborted run has nothing finished worth naming and saving. */
    canSave: boolean;
  } = $props();

  let saveOpen = $state(false);
  let saveTitle = $state('');
  let saving = $state(false);
  let saveFailed = $state(false);
  let savedUrl = $state<string | null>(null);
  let savedLinkCopied = $state(false);

  function openSaveForm(): void {
    saveTitle = titleFor;
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

<div class="flex flex-wrap items-center gap-3" data-testid="sim-save">
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
      class="{SECONDARY_BUTTON} border-line-warm text-nav px-4"
      onclick={() => void copySavedLink()}
      data-testid="sim-save-copy"
    >
      {savedLinkCopied ? simCopy.copied : simCopy.copyLink}
    </button>
    <a
      class="{SECONDARY_BUTTON} border-line-warm text-nav px-4"
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
      class="{SECONDARY_BUTTON} border-line-warm text-nav px-5 disabled:opacity-50"
      disabled={saving}
      onclick={() => void confirmSave()}
      data-testid="sim-save-confirm"
    >
      {saving ? simCopy.savingAction : simCopy.saveAction}
    </button>
    <button
      type="button"
      class="{SECONDARY_BUTTON} border-line-warm text-nav px-4"
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
      class="{SECONDARY_BUTTON} border-line-warm text-nav px-5 disabled:opacity-50"
      disabled={!canSave}
      title={!canSave ? simCopy.saveAbortedDisabled : undefined}
      onclick={openSaveForm}
      data-testid="sim-save-open"
    >
      {simCopy.saveThisSim}
    </button>
  {/if}
</div>
