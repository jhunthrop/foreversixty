<!-- web/src/components/guild/GuildInviteLink.svelte -->
<!-- The invite-link block shared by the Settings tab/page and the officer strip: Rotate asks
     first ("the old link stops working"), then the new link is a read-only field holding the
     full URL with a Copy button ("Copied" for two seconds) and, where the browser has it, a
     Share button. The API stores only a hash of the token, so the link is shown once and
     the block says so, with the day the current link was made underneath. -->
<script lang="ts">
  import { guildInviteHref } from '../../lib/characters';
  import { GuildApiError, rotateInvite } from '../../lib/guild/api';
  import { guildSettingsCopy } from '../../lib/guild/copy';
  import { formatMadeOn, inviteUrl } from '../../lib/guild/invite-link';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { BUSY_CLASS } from '../../lib/ui/busy';

  let {
    guildId,
    rotatedAt = null,
    disabled = false,
    testidPrefix = 'guild-invite',
    onRotated = undefined,
  }: {
    guildId: number;
    /** When the current link was made (the settings read), for the muted line underneath. */
    rotatedAt?: string | null;
    disabled?: boolean;
    testidPrefix?: string;
    onRotated?: (rotatedAt: string) => void;
  } = $props();

  let confirming = $state(false);
  let busy = $state(false);
  let error = $state('');
  let url = $state<string | null>(null);
  let madeAt = $state<string | null>(null);
  let copied = $state(false);
  let copyError = $state(false);
  let copiedTimer: ReturnType<typeof setTimeout> | undefined;

  const made = $derived(madeAt ?? rotatedAt);
  const madeOn = $derived(made === null ? '' : formatMadeOn(made));
  const canShare = $derived(typeof navigator !== 'undefined' && typeof navigator.share === 'function');

  async function onRotate(): Promise<void> {
    busy = true;
    error = '';
    try {
      const result = await rotateInvite(guildId);
      url = inviteUrl(guildInviteHref(result.token));
      madeAt = result.rotated_at;
      copied = false;
      copyError = false;
      confirming = false;
      onRotated?.(result.rotated_at);
    } catch (thrown) {
      error = thrown instanceof GuildApiError ? thrown.message : guildSettingsCopy.actionFailed;
    } finally {
      busy = false;
    }
  }

  async function onCopy(): Promise<void> {
    if (url === null) return;
    try {
      await navigator.clipboard.writeText(url);
      copyError = false;
      copied = true;
      clearTimeout(copiedTimer);
      copiedTimer = setTimeout(() => (copied = false), 2000);
    } catch {
      copied = false;
      copyError = true;
    }
  }

  async function onShare(): Promise<void> {
    if (url === null) return;
    try {
      await navigator.share({ url });
    } catch {
      // Dismissing the share sheet rejects; nothing to say.
    }
  }
</script>

<div class="flex w-full flex-col gap-3" data-testid={`${testidPrefix}-block`}>
  {#if !confirming}
    <button
      class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4 ${busy ? BUSY_CLASS : ''}`}
      onclick={() => (confirming = true)}
      disabled={busy || disabled}
      aria-busy={busy}
      data-testid={`${testidPrefix}-rotate`}
    >
      {guildSettingsCopy.rotateButton}
    </button>
  {:else}
    <div
      class="border-line-soft flex flex-col gap-3 border p-4"
      data-testid={`${testidPrefix}-rotate-confirm`}
    >
      <p class="text-[13px]">{guildSettingsCopy.rotateConfirmLine}</p>
      <div class="flex gap-3">
        <button
          class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-3 ${busy ? BUSY_CLASS : ''}`}
          onclick={() => void onRotate()}
          disabled={busy}
          aria-busy={busy}
          data-testid={`${testidPrefix}-rotate-confirm-button`}
        >
          {guildSettingsCopy.rotateConfirmButton}
        </button>
        <button
          class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-3`}
          onclick={() => (confirming = false)}
          disabled={busy}
        >
          {guildSettingsCopy.rotateCancel}
        </button>
      </div>
    </div>
  {/if}
  {#if url !== null}
    <div class="flex flex-col gap-2" data-testid={`${testidPrefix}-token`}>
      <label class="flex flex-col gap-1 text-[13px]">
        <span class="label text-muted">{guildSettingsCopy.linkLabel}</span>
        <input
          class="border-line-warm bg-raised rounded-control text-text h-11 w-full min-w-0 px-3 font-mono text-[13px]"
          type="text"
          readonly
          value={url}
          onfocus={(event) => (event.currentTarget as HTMLInputElement).select()}
          data-testid={`${testidPrefix}-url`}
        />
      </label>
      <div class="flex flex-wrap items-center gap-3">
        <button
          class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4`}
          onclick={() => void onCopy()}
          data-testid={`${testidPrefix}-copy`}
        >
          {copied ? guildSettingsCopy.copied : guildSettingsCopy.copyButton}
        </button>
        {#if canShare}
          <button
            class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-4`}
            onclick={() => void onShare()}
            data-testid={`${testidPrefix}-share`}
          >
            {guildSettingsCopy.shareButton}
          </button>
        {/if}
        <span class="text-muted text-[13px]">{guildSettingsCopy.tokenShownOnce}</span>
      </div>
      {#if copyError}
        <p class="text-[13px]" role="alert" data-testid={`${testidPrefix}-copy-error`}>
          {guildSettingsCopy.copyFailed}
        </p>
      {/if}
    </div>
  {/if}
  {#if madeOn !== ''}
    <p class="text-muted text-[12px]" data-testid={`${testidPrefix}-made`}>
      {guildSettingsCopy.currentLinkMade(madeOn)}
    </p>
  {/if}
  {#if error !== ''}
    <p class="text-[13px]" role="alert" data-testid={`${testidPrefix}-error`}>{error}</p>
  {/if}
</div>
