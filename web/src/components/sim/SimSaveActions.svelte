<!-- web/src/components/sim/SimSaveActions.svelte -->
<!-- The live result's actions in the head's answer (spec 2026-11-06-sim §4.5): "Save this sim"
     (signed out: "Sign in to save"), which opens its inline title form in place; once saved the
     Button becomes "Copy link" and the answer says "Saved"; then "Find upgrades in Top Gear ›".
     The page never navigates away from the result it saved. A stopped run has nothing to save. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { simPageCopy as copy } from '../../lib/sim/page-copy';
  import type { SimResult } from '../../lib/sim/types';
  import { busyCopy } from '../../lib/ui/copy';
  import Button from '../ui/Button.svelte';
  import CopyButton from '../ui/CopyButton.svelte';
  import Field from '../ui/Field.svelte';
  import FormRow from '../ui/FormRow.svelte';

  let {
    result,
    reportTitle,
    signedIn,
    topGearHref,
    onsave,
    onsignin,
    onsaved,
  }: {
    result: SimResult;
    reportTitle: string;
    signedIn: boolean;
    topGearHref: string;
    /** Returns the new sim's id, or null on failure. */
    onsave: (title: string) => Promise<string | null>;
    onsignin: () => void;
    /** The saved link, for the answer's "Saved" stamp. */
    onsaved?: (url: string | null) => void;
  } = $props();

  let formOpen = $state(false);
  let title = $state('');
  let saving = $state(false);
  let failed = $state(false);
  let savedUrl = $state<string | null>(null);

  // A new result is a new thing to save: the form, the link and the copy state start over.
  $effect(() => {
    void result;
    formOpen = false;
    failed = false;
    savedUrl = null;
    onsaved?.(null);
  });

  const canSave = $derived(result.aborted !== true);

  async function confirm(event: SubmitEvent): Promise<void> {
    event.preventDefault();
    saving = true;
    failed = false;
    const id = await onsave(title);
    saving = false;
    if (id === null) {
      failed = true;
      return;
    }
    formOpen = false;
    savedUrl = `${window.location.origin}/sim/${id}`;
    onsaved?.(savedUrl);
  }
</script>

{#if formOpen}
  <FormRow label={copy.answer.saveThis} onsubmit={(event) => void confirm(event)} testid="sim-save">
    <Field
      id="sim-save-title"
      label={simCopy.saveTitleLabel}
      bind:value={title}
      maxlength={120}
      testid="sim-save-title"
    />
    {#snippet submit()}
      <span class="flex gap-2">
        <Button
          tone="primary"
          size="sm"
          type="submit"
          busy={saving}
          busyLabel={busyCopy.saving}
          testid="sim-save-confirm">{simCopy.saveAction}</Button
        >
        <Button size="sm" onclick={() => (formOpen = false)} testid="sim-save-cancel">{simCopy.cancel}</Button
        >
      </span>
    {/snippet}
  </FormRow>
  {#if failed}<span class="text-wipe text-[13px]" role="alert" data-testid="sim-save-error"
      >{simCopy.saveFailed}</span
    >{/if}
{:else if savedUrl !== null}
  <CopyButton text={savedUrl} label={copy.answer.copyLink} testid="sim-save-copy" />
  <a class="sr-only" href={savedUrl} data-testid="sim-save-link">{savedUrl}</a>
{:else if !signedIn}
  <Button size="sm" onclick={onsignin} testid="sim-save-open">{copy.answer.signInToSave}</Button>
{:else}
  <Button
    size="sm"
    disabled={!canSave}
    onclick={() => {
      title = reportTitle;
      failed = false;
      formOpen = true;
    }}
    testid="sim-save-open">{copy.answer.saveThis}</Button
  >
{/if}
<a class="text-gold text-[14px] font-semibold" href={topGearHref} data-testid="sim-next-action"
  >{copy.answer.topGear}</a
>
