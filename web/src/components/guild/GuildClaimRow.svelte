<!-- web/src/components/guild/GuildClaimRow.svelte -->
<!-- The band's claim row (spec 2026-11-05 §3.1), under the raid line on every tab until a claim
     exists. The guild master gets the outline button to the claim page (the page's one fill belongs
     to the Before raids open card); an officer copies the claim page's link for them; a pending
     claim opens. Wraps below 768, the control full width. -->
<script lang="ts">
  import type { ClaimRowKind } from '../../lib/guild/claim-row';
  import { guildHomeCopy } from '../../lib/guild/copy';
  import { inviteUrl } from '../../lib/guild/invite-link';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { createCopyFeedback, type CopyFeedbackState } from '../../lib/ui/copy-feedback';
  import ActionLink from '../ui/ActionLink.svelte';
  import CopyLinkField from '../ui/CopyLinkField.svelte';

  let {
    kind,
    guildName,
    claimHref,
  }: {
    kind: ClaimRowKind;
    guildName: string;
    claimHref: string;
  } = $props();

  const words = guildHomeCopy.claimRow;
  let feedback = $state<CopyFeedbackState>('idle');
  const copier = createCopyFeedback((next) => (feedback = next));
  $effect(() => copier.dispose);

  const sentence = $derived(
    kind === 'guild-master'
      ? words.guildMaster(guildName)
      : kind === 'officer'
        ? words.officer(guildName)
        : words.pending(guildName),
  );
  const claimUrl = $derived(inviteUrl(claimHref));

  function onCopy(event: MouseEvent): void {
    event.preventDefault();
    void copier.copy(claimUrl);
  }
</script>

<div
  class="flex min-h-11 flex-col gap-3 md:flex-row md:items-center md:justify-between md:gap-4"
  data-testid="guild-claim-row"
  data-kind={kind}
>
  <p class="text-text text-[14px]" data-testid="guild-claim-row-sentence">{sentence}</p>
  {#if kind === 'guild-master'}
    <a
      class="{SECONDARY_BUTTON_FIXED} border-gold text-gold w-full shrink-0 justify-center px-4 md:w-fit"
      href={claimHref}
      data-testid="guild-claim-link">{words.claimButton(guildName)}</a
    >
  {:else if kind === 'officer'}
    <ActionLink
      href={claimHref}
      class="shrink-0"
      onclick={onCopy}
      data-testid="guild-claim-copy-link"
      aria-live="polite">{feedback === 'copied' ? words.copied : words.copyLink}</ActionLink
    >
  {:else}
    <ActionLink href={claimHref} class="shrink-0" data-testid="guild-confirm-claim-link"
      >{words.open}</ActionLink
    >
  {/if}
</div>
{#if kind === 'officer' && feedback === 'failed'}
  <CopyLinkField
    url={claimUrl}
    label={words.copyLink}
    shareTitle={words.copyLink}
    testid="guild-claim-link-field"
    autoSelect
  />
{/if}
