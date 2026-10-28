<!-- web/src/components/ui/CopyLinkField.svelte -->
<!-- A read-only field holding a full link, with a Copy button ("Copied" for two seconds), a
     Share button where the browser has one, and a failure line when the clipboard refuses.
     Select-all on focus, so a failed copy is one long-press away from done. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { createCopyFeedback, type CopyFeedbackState } from '../../lib/ui/copy-feedback';
  import { uiCopy } from '../../lib/ui/copy';

  let {
    url,
    label,
    shareTitle,
    shareText,
    testid,
    note = undefined,
  }: {
    url: string;
    /** The visible field label. */
    label: string;
    shareTitle: string;
    shareText: string;
    /** Prefix for the field's test ids: `-token`, `-url`, `-copy`, `-share`, `-copy-error`. */
    testid: string;
    /** Optional muted line at the end of the button row. */
    note?: Snippet;
  } = $props();

  let feedback = $state<CopyFeedbackState>('idle');

  const canShare = $derived(typeof navigator !== 'undefined' && typeof navigator.share === 'function');
  const copier = createCopyFeedback(
    (text) => navigator.clipboard.writeText(text),
    (next) => (feedback = next),
  );

  $effect(() => copier.dispose);

  async function onShare(): Promise<void> {
    try {
      await navigator.share({ url, title: shareTitle, text: shareText });
    } catch {
      // Dismissing the share sheet rejects; nothing to say.
    }
  }
</script>

<div class="flex flex-col gap-2" data-testid={`${testid}-token`}>
  <label class="flex flex-col gap-1 text-[13px]">
    <span class="label text-muted">{label}</span>
    <input
      class="border-line-warm bg-raised rounded-control text-text h-11 w-full min-w-0 px-3 font-mono text-[13px]"
      type="text"
      readonly
      value={url}
      onfocus={(event) => (event.currentTarget as HTMLInputElement).select()}
      data-testid={`${testid}-url`}
    />
  </label>
  <div class="flex flex-wrap items-center gap-3">
    <button
      class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4`}
      onclick={() => void copier.copy(url)}
      data-testid={`${testid}-copy`}
    >
      {feedback === 'copied' ? uiCopy.copiedLink : uiCopy.copyLink}
    </button>
    {#if canShare}
      <button
        class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-4`}
        onclick={() => void onShare()}
        data-testid={`${testid}-share`}
      >
        {uiCopy.shareLink}
      </button>
    {/if}
    {@render note?.()}
  </div>
  {#if feedback === 'failed'}
    <p class="text-[13px]" role="alert" data-testid={`${testid}-copy-error`}>{uiCopy.copyLinkFailed}</p>
  {/if}
</div>
