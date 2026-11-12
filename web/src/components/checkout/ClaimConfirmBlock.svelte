<!-- web/src/components/checkout/ClaimConfirmBlock.svelte -->
<!-- Spec 2026-10-27-subscribe-and-join §5.3: the in-place confirm under a claim opener, a selected
     Row (the gold ground and bar, as the account's seat offer): the confirm sentence, then in its
     foot ui/ConfirmButtons: the commit (primary, spinner and busy word) and Cancel (secondary),
     which takes focus on open and Escape; the refusal line takes focus when one shows. The
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
  import ConfirmButtons from '../ui/ConfirmButtons.svelte';
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
  let errorEl = $state<HTMLParagraphElement | undefined>(undefined);

  const words = $derived(confirmWords(kind, guild));

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
</script>

<div role="group" aria-label={words.sentence} data-testid={`guild-${guild.id}-claim-confirm`}>
  <Row title={words.sentence} selected>
    {#snippet titleContent()}<span class="text-[14px] leading-5">{words.sentence}</span>{/snippet}
    {#snippet foot()}
      <ConfirmButtons
        tone="primary"
        size="md"
        width="w-full"
        confirm={words.commit}
        cancel={copy.confirm.cancel}
        {busy}
        busyLabel={words.busy}
        disabled={frozen}
        onconfirm={() => void commit()}
        oncancel={onCancel}
        testid={`guild-${guild.id}-claim`}
        confirmTestid={`guild-${guild.id}-claim-commit`}
      />
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
