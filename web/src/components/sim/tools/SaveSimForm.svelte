<!-- web/src/components/sim/tools/SaveSimForm.svelte -->
<!-- Design 8 / contract 10.6's browser-result save ("saved sims of every kind are public at
     /sim/<id>"), in the one form every bulk tool page shares: closed, open (a title to
     confirm), then saved (the link, copyable, and a new tab). Extracted out of
     ComboResults.svelte (Task 16 built it for Top Gear/talents/drops) so a weights result
     gets the identical flow rather than a second, drifting copy of the same six pieces of
     state -- Top Gear, talents, drops and weights all save through this one component now. -->
<script lang="ts">
  import { simCopy } from '../../../lib/sim/copy';
  import { button } from '../../../lib/ui/button';
  import { busyCopy } from '../../../lib/ui/copy';
  import Button from '../../ui/Button.svelte';
  import CopyLinkField from '../../ui/CopyLinkField.svelte';
  import Field from '../../ui/Field.svelte';

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
</script>

<div class="flex flex-wrap items-end gap-3" data-testid="sim-save">
  {#if savedUrl !== null}
    <div class="min-w-0 flex-1 md:max-w-[420px]">
      <CopyLinkField
        url={savedUrl}
        label={simCopy.savedLinkLabel}
        shareTitle={titleFor}
        labelVisible={false}
        testid="sim-save"
      >
        {#snippet note()}
          <a
            class={button('secondary', 'sm')}
            href={savedUrl}
            target="_blank"
            rel="noopener"
            data-testid="sim-open-new-tab">{simCopy.openInNewTab}</a
          >
        {/snippet}
      </CopyLinkField>
    </div>
  {:else if saveOpen}
    <div class="w-[260px]">
      <Field
        id="sim-save-title"
        label={simCopy.saveTitleLabel}
        testid="sim-save-title"
        bind:value={saveTitle}
      />
    </div>
    <Button
      size="sm"
      busy={saving}
      busyLabel={busyCopy.saving}
      onclick={() => void confirmSave()}
      testid="sim-save-confirm">{simCopy.saveAction}</Button
    >
    <Button tone="ghost" size="sm" testid="sim-save-cancel" onclick={cancelSave}>{simCopy.cancel}</Button>
    {#if saveFailed}
      <span role="alert" class="text-wipe text-[13px]" data-testid="sim-save-error">
        {simCopy.saveFailed}
      </span>
    {/if}
  {:else}
    <Button size="sm" testid="sim-save-open" disabled={!canSave} onclick={openSaveForm}
      >{simCopy.saveThisSim}</Button
    >
    <!-- Task 7: the disabled reason, said plainly rather than in a title= that never fires
         on touch -- the same pattern Task 3 used for a disabled Run. -->
    {#if !canSave}
      <p class="text-muted text-[13px]" data-testid="sim-save-disabled-note">{simCopy.saveAbortedDisabled}</p>
    {/if}
  {/if}
</div>
