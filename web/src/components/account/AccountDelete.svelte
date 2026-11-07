<!-- web/src/components/account/AccountDelete.svelte -->
<!-- /account's last rail Section, Delete account (gap G6, closed 2026-11-07): one line and the
     danger button that opens the confirm dialog. The page runs the call and says how it went;
     closing the dialog gives focus back to the button. -->
<script lang="ts">
  import { tick } from 'svelte';
  import { accountPageCopy } from '../../lib/account/account-page-copy';
  import { smallButton } from '../../lib/account/small-button';
  import Section from '../ui/Section.svelte';
  import AccountDeleteDialog from './AccountDeleteDialog.svelte';

  let {
    battletag,
    open,
    busy,
    mismatch,
    refusal,
    onopen,
    onclose,
    onconfirm,
  }: {
    battletag: string | null | undefined;
    open: boolean;
    busy: boolean;
    mismatch: boolean;
    refusal: string;
    onopen: () => void;
    onclose: () => void;
    onconfirm: (typed: string) => void;
  } = $props();

  const copy = accountPageCopy.deleteSection;
  let trigger = $state<HTMLButtonElement | undefined>(undefined);

  function close(): void {
    onclose();
    void tick().then(() => trigger?.focus());
  }
</script>

<Section title={copy.title} anchor="delete" testid="account-delete">
  <div class="flex flex-col items-start gap-3">
    <p class="text-small text-muted">{copy.line}</p>
    <button
      bind:this={trigger}
      type="button"
      class={smallButton('danger')}
      aria-haspopup="dialog"
      onclick={onopen}
      data-testid="account-delete-open">{copy.button}</button
    >
  </div>
</Section>

{#if open}
  <AccountDeleteDialog {battletag} {busy} {mismatch} {refusal} {onconfirm} onclose={close} />
{/if}
