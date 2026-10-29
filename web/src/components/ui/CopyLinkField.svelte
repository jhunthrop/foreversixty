<!-- web/src/components/ui/CopyLinkField.svelte -->
<!-- A read-only field holding a full link, with a Copy button ("Copied" for two seconds), a
     Share button where the browser has one, and a failure line when the clipboard refuses.
     Select-all on focus, so a failed copy is one long-press away from done.

     The defaults are GuildInviteLink's render (capture 17 holds it pixel-identical). The checkout
     card and the claim page hand-off (spec 2026-10-27-subscribe-and-join §5.5) keep the label for
     screen readers only and, below 768, put the buttons full width; on a card the field takes the
     page ground so it reads as a field against the card's. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { createCopyFeedback, type CopyFeedbackState } from '../../lib/ui/copy-feedback';
  import { uiCopy } from '../../lib/ui/copy';

  let {
    url,
    label,
    shareTitle,
    shareText = undefined,
    testid,
    note = undefined,
    labelVisible = true,
    onCard = false,
    phoneFullWidth = false,
  }: {
    url: string;
    /** The field label: visible, or for screen readers only (`labelVisible` false). */
    label: string;
    shareTitle: string;
    /** Left out, the share sheet carries only the link and its title. */
    shareText?: string;
    /** Prefix for the field's test ids: `-token`, `-url`, `-copy`, `-share`, `-copy-error`. */
    testid: string;
    /** Optional muted line at the end of the button row. */
    note?: Snippet;
    /** False keeps the label for screen readers only. */
    labelVisible?: boolean;
    /** Inside a `card`: the field takes the page ground, so it stands out from the card's. */
    onCard?: boolean;
    /** Below 768 the buttons fill the row (two columns when Share exists). */
    phoneFullWidth?: boolean;
  } = $props();

  let feedback = $state<CopyFeedbackState>('idle');

  const canShare = $derived(typeof navigator !== 'undefined' && typeof navigator.share === 'function');
  const copier = createCopyFeedback(
    (text) => navigator.clipboard.writeText(text),
    (next) => (feedback = next),
  );

  const fieldGround = $derived(onCard ? 'bg-bg' : 'bg-raised');
  const buttonRow = $derived(
    phoneFullWidth
      ? `grid gap-3 md:flex md:flex-wrap md:items-center ${canShare ? 'grid-cols-2' : 'grid-cols-1'}`
      : 'flex flex-wrap items-center gap-3',
  );
  const buttonWidth = $derived(phoneFullWidth ? 'w-full justify-center md:w-fit' : 'w-fit');

  $effect(() => copier.dispose);

  async function onShare(): Promise<void> {
    try {
      await navigator.share(
        shareText === undefined ? { url, title: shareTitle } : { url, title: shareTitle, text: shareText },
      );
    } catch {
      // Dismissing the share sheet rejects; nothing to say.
    }
  }
</script>

<div class="flex flex-col gap-2" data-testid={`${testid}-token`}>
  <label class="flex flex-col gap-1 text-[13px]">
    <span class={labelVisible ? 'quiet-label' : 'sr-only'}>{label}</span>
    <input
      class={`border-line-warm ${fieldGround} rounded-control text-text h-11 w-full min-w-0 px-3 font-mono text-[13px]`}
      type="text"
      readonly
      value={url}
      onfocus={(event) => (event.currentTarget as HTMLInputElement).select()}
      data-testid={`${testid}-url`}
    />
  </label>
  <div class={buttonRow}>
    <button
      class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong ${buttonWidth} px-4`}
      onclick={() => void copier.copy(url)}
      data-testid={`${testid}-copy`}
    >
      {feedback === 'copied' ? uiCopy.copiedLink : uiCopy.copyLink}
    </button>
    {#if canShare}
      <button
        class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text ${buttonWidth} px-4`}
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
