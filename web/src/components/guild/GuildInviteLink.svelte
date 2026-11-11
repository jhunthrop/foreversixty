<!-- web/src/components/guild/GuildInviteLink.svelte -->
<!-- The invite row of the Members and access card (spec 2026-11-05 §3.3, §4.1). The API keeps the
     link sealed and hands it back to a verified officer (`url`), so a later visit shows the field,
     `Made {date} by {Name}.` and Make a new link. Where the API could not keep it (`url` null), the
     link shows only right after it is made and the row says so, never offering a "show current
     link" it cannot honour. A first link needs no confirm, nothing breaks; a new one asks first,
     because the old one stops working. The URL lives in props and state only, never in storage. -->
<script lang="ts">
  import { tick } from 'svelte';
  import { guildInviteHref } from '../../lib/characters';
  import { rotateInvite, type ClaimHolder, type GuildInviteView } from '../../lib/guild/api';
  import { guildSettingsCopy as copy } from '../../lib/guild/copy';
  import { formatMadeOn, inviteUrl } from '../../lib/guild/invite-link';
  import { BUSY_CLASS } from '../../lib/ui/busy';
  import { button } from '../../lib/ui/button';
  import { CONFIRM_BLOCK } from '../../lib/guild/billing-styles';
  import { COPIED_LONG_MS } from '../../lib/ui/copy-feedback';
  import CopyLinkField from '../ui/CopyLinkField.svelte';
  import CharacterName from '../ui/CharacterName.svelte';

  const TESTID = 'guild-invite';

  let {
    guildId,
    rotatedAt = null,
    keptUrl = null,
    madeBy = null,
    disabled = false,
    onRotated = undefined,
  }: {
    guildId: number;
    /** When the current link was made (the settings read); null when none exists. */
    rotatedAt?: string | null;
    /** The current link's site path when the API kept it; null when it cannot be shown again. */
    keptUrl?: string | null;
    /** Who made the current link, by their character on the roster. */
    madeBy?: ClaimHolder | null;
    disabled?: boolean;
    onRotated?: (invite: GuildInviteView) => void;
  } = $props();

  let confirming = $state(false);
  let busy = $state(false);
  let failed = $state(false);
  /** The link made on this page, until the settings read hands back the kept one. */
  let madePath = $state<string | null>(null);
  let root = $state<HTMLElement | null>(null);
  let cancelButton = $state<HTMLButtonElement | null>(null);
  let newLinkButton = $state<HTMLButtonElement | null>(null);

  const madeOn = $derived(rotatedAt === null ? '' : formatMadeOn(rotatedAt));
  const kept = $derived(keptUrl !== null);
  const url = $derived.by(() => {
    const path = keptUrl ?? madePath;
    return path === null ? null : inviteUrl(path);
  });

  async function make(): Promise<void> {
    busy = true;
    failed = false;
    try {
      const result = await rotateInvite(guildId);
      madePath = guildInviteHref(result.token);
      confirming = false;
      onRotated?.(result.invite);
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
  <div class="flex flex-col items-start gap-3 pb-3">
    {#if confirming}
      <!-- svelte-ignore a11y_no_static_element_interactions -->
      <div class="{CONFIRM_BLOCK} w-full" onkeydown={onConfirmKey} data-testid={`${TESTID}-rotate-confirm`}>
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
    {:else if url !== null}
      <div class="w-full">
        {#key url}
          <CopyLinkField
            {url}
            label={copy.linkLabel}
            shareTitle={copy.inviteShareTitle}
            testid={TESTID}
            copiedMs={COPIED_LONG_MS}
          />
        {/key}
      </div>
      {#if kept}
        <p
          class="text-muted inline-flex flex-wrap items-center gap-1 text-[13px]"
          data-testid={`${TESTID}-made`}
        >
          {#if madeBy !== null}
            {copy.linkMadeBy(madeOn)}
            <span><CharacterName name={madeBy.name} cls={madeBy.class} />.</span>
          {:else}
            {copy.linkMade(madeOn)}
          {/if}
        </p>
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
        <p class="text-muted text-[13px]" data-testid={`${TESTID}-shown-once`}>{copy.tokenShownOnce}</p>
      {/if}
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
