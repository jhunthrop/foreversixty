<!-- web/src/components/guild/GuildInviteLink.svelte -->
<!-- The invite row of the Members and access card (spec 2026-11-05 §3.3, the no-migration interim).
     The API keeps only a fingerprint of the token, so the link can be shown only right after it is
     made: the row says so instead of offering a "show current link" it cannot honour. A first link
     needs no confirm, nothing breaks; a new one asks first, because the old one stops working. The
     URL lives in this component's state only, never in storage. -->
<script lang="ts">
  import { tick } from 'svelte';
  import { guildInviteHref } from '../../lib/characters';
  import { rotateInvite } from '../../lib/guild/api';
  import { guildSettingsCopy as copy } from '../../lib/guild/copy';
  import { formatMadeOn, inviteUrl } from '../../lib/guild/invite-link';
  import { BUSY_CLASS } from '../../lib/ui/busy';
  import { button } from '../../lib/ui/button';
  import { COPIED_LONG_MS } from '../../lib/ui/copy-feedback';
  import CopyLinkField from '../ui/CopyLinkField.svelte';
  import SettingsRow from '../ui/SettingsRow.svelte';

  const TESTID = 'guild-invite';

  let {
    guildId,
    guildName,
    rotatedAt = null,
    disabled = false,
    onRotated = undefined,
  }: {
    guildId: number;
    guildName: string;
    /** When the current link was made (the settings read); null when none exists. */
    rotatedAt?: string | null;
    disabled?: boolean;
    onRotated?: (rotatedAt: string) => void;
  } = $props();

  let confirming = $state(false);
  let busy = $state(false);
  let failed = $state(false);
  let url = $state<string | null>(null);
  let root = $state<HTMLElement | null>(null);
  let cancelButton = $state<HTMLButtonElement | null>(null);
  let newLinkButton = $state<HTMLButtonElement | null>(null);

  const madeOn = $derived(rotatedAt === null ? '' : formatMadeOn(rotatedAt));

  async function make(): Promise<void> {
    busy = true;
    failed = false;
    try {
      const result = await rotateInvite(guildId);
      url = inviteUrl(guildInviteHref(result.token));
      confirming = false;
      onRotated?.(result.rotated_at);
      await tick();
      root?.querySelector<HTMLElement>(`[data-testid="${TESTID}-copy"]`)?.focus();
    } catch {
      failed = true;
    } finally {
      busy = false;
    }
  }

  async function askFirst(): Promise<void> {
    confirming = true;
    failed = false;
    await tick();
    cancelButton?.focus();
  }

  async function cancel(): Promise<void> {
    confirming = false;
    await tick();
    newLinkButton?.focus();
  }

  function onConfirmKey(event: KeyboardEvent): void {
    if (event.key !== 'Escape') return;
    event.preventDefault();
    void cancel();
  }
</script>

<div bind:this={root} id="invite" class="flex scroll-mt-20 flex-col" data-testid={`${TESTID}-block`}>
  <SettingsRow label={copy.inviteHeading} sub={copy.inviteHelp(guildName)} />
  <div class="flex flex-col items-start gap-3 px-5 pb-4">
    {#if url !== null}
      <div class="w-full">
        {#key url}
          <CopyLinkField
            {url}
            label={copy.linkLabel}
            shareTitle={copy.inviteShareTitle}
            shareText={copy.inviteShareText}
            testid={TESTID}
            copiedMs={COPIED_LONG_MS}
          />
        {/key}
      </div>
      <p class="text-muted text-[13px]" data-testid={`${TESTID}-shown-once`}>{copy.tokenShownOnce}</p>
    {:else if confirming}
      <!-- svelte-ignore a11y_no_static_element_interactions -->
      <div
        class="border-line-soft flex w-full flex-col gap-3 border p-4"
        onkeydown={onConfirmKey}
        data-testid={`${TESTID}-rotate-confirm`}
      >
        <p class="text-[13px]">{copy.rotateConfirmLine}</p>
        <div class="flex gap-3">
          <button
            class="{button('primary', 'sm')} {busy ? BUSY_CLASS : ''}"
            onclick={() => void make()}
            disabled={busy}
            aria-busy={busy}
            data-testid={`${TESTID}-rotate-confirm-button`}
          >
            {busy ? copy.making : copy.newLinkButton}
          </button>
          <button
            bind:this={cancelButton}
            class={button('ghost', 'sm')}
            onclick={() => void cancel()}
            disabled={busy}
            data-testid={`${TESTID}-rotate-cancel`}
          >
            {copy.rotateCancel}
          </button>
        </div>
      </div>
    {:else if madeOn !== ''}
      <p class="text-muted text-[13px]" data-testid={`${TESTID}-made`}>{copy.currentLinkMade(madeOn)}</p>
      <button
        bind:this={newLinkButton}
        class={button('secondary', 'sm')}
        onclick={() => void askFirst()}
        {disabled}
        data-testid={`${TESTID}-rotate`}
      >
        {copy.newLinkButton}
      </button>
    {:else}
      <p class="text-muted text-[13px]" data-testid={`${TESTID}-none`}>{copy.noLinkYet}</p>
      <!-- Secondary: the plan card holds the page's one fill (guild round 2). -->
      <button
        class="{button('secondary', 'sm')} {busy ? BUSY_CLASS : ''}"
        onclick={() => void make()}
        disabled={busy || disabled}
        aria-busy={busy}
        data-testid={`${TESTID}-make`}
      >
        {busy ? copy.making : copy.firstLinkButton}
      </button>
    {/if}
    {#if failed}
      <p class="text-wipe text-[13px]" role="alert" data-testid={`${TESTID}-error`}>{copy.inviteFailed}</p>
    {/if}
  </div>
</div>
