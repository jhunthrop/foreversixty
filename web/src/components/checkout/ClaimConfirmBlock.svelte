<!-- web/src/components/checkout/ClaimConfirmBlock.svelte -->
<!-- Spec 2026-10-27-subscribe-and-join §5.3: the in-place confirm under a claim opener, a selected
     Row (the gold ground and bar, as the account's seat offer): the confirm sentence, then in its
     foot the commit (the primary button, busy label + aria-busy) and Cancel (secondary).
     Focus moves to Cancel on open, to the refusal line when one shows; Escape closes. The
     decisions (words, outcome, refusal) are claim-flow.ts; the shell owns which block is open
     and what happens after, so this draws and reports. -->
<script lang="ts">
  import { checkoutGuildCopy as copy } from '../../lib/billing/checkout-guild-copy';
  import {
    confirmWords,
    performClaim,
    type ClaimGuild,
    type ClaimIo,
    type ClaimKind,
  } from '../../lib/billing/claim-flow';
  import { button } from '../../lib/ui/button';
  import { BUSY_CLASS } from '../../lib/ui/busy';
  import Row from '../ui/Row.svelte';

  let {
    kind,
    guild,
    io,
    error,
    frozen = false,
    onDone,
    onRefused,
    onCancel,
  }: {
    kind: ClaimKind;
    guild: ClaimGuild;
    io: ClaimIo;
    /** The §9.1 refusal line to show, or null. */
    error: string | null;
    /** A re-read is in flight: the commit waits. */
    frozen?: boolean;
    onDone: (line: string) => void;
    onRefused: (code: string, rereads: boolean) => void;
    onCancel: () => void;
  } = $props();

  let busy = $state(false);
  let cancelEl = $state<HTMLButtonElement | undefined>(undefined);
  let errorEl = $state<HTMLParagraphElement | undefined>(undefined);

  const words = $derived(confirmWords(kind, guild));

  $effect(() => {
    cancelEl?.focus();
  });

  $effect(() => {
    if (error !== null) errorEl?.focus();
  });

  async function commit(): Promise<void> {
    busy = true;
    const outcome = await performClaim(kind, guild, io);
    busy = false;
    if (outcome.ok) onDone(outcome.line);
    else onRefused(outcome.code, outcome.rereads);
  }

  function onKeydown(event: KeyboardEvent): void {
    if (event.key !== 'Escape' || busy) return;
    event.stopPropagation();
    onCancel();
  }
</script>

<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
<div
  role="group"
  aria-label={words.sentence}
  onkeydown={onKeydown}
  data-testid={`guild-${guild.id}-claim-confirm`}
>
  <Row title={words.sentence} selected>
    {#snippet titleContent()}<span class="text-[14px] leading-5">{words.sentence}</span>{/snippet}
    {#snippet foot()}
      <button
        type="button"
        class={`${button('primary')} w-full ${busy ? BUSY_CLASS : ''}`}
        disabled={busy || frozen}
        aria-busy={busy || frozen}
        onclick={() => void commit()}
        data-testid={`guild-${guild.id}-claim-commit`}
      >
        {busy ? words.busy : words.commit}
      </button>
      <button
        type="button"
        bind:this={cancelEl}
        class={`${button('secondary')} w-full`}
        disabled={busy}
        onclick={onCancel}
        data-testid={`guild-${guild.id}-claim-cancel`}
      >
        {copy.confirm.cancel}
      </button>
      {#if error !== null}
        <p
          bind:this={errorEl}
          tabindex="-1"
          role="alert"
          class="text-wipe m-0 text-[13px] outline-none"
          data-testid={`guild-${guild.id}-claim-error`}
        >
          {error}
        </p>
      {/if}
    {/snippet}
  </Row>
</div>
