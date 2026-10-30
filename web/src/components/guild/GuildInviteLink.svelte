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
  import CopyLinkField from '../ui/CopyLinkField.svelte';
  import { BUSY_CLASS } from '../../lib/ui/busy';
  import { button } from '../../lib/ui/button';
  import SettingsRow from '../ui/SettingsRow.svelte';

  let {
    guildId,
    rotatedAt = null,
    disabled = false,
    testidPrefix = 'guild-invite',
    onRotated = undefined,
    row = false,
  }: {
    guildId: number;
    /** When the current link was made (the settings read), for the muted line underneath. */
    rotatedAt?: string | null;
    disabled?: boolean;
    testidPrefix?: string;
    onRotated?: (rotatedAt: string) => void;
    /** Drawn as a settings row (the Members and access card); the default is the stand-alone block. */
    row?: boolean;
  } = $props();

  let confirming = $state(false);
  let busy = $state(false);
  let error = $state('');
  let url = $state<string | null>(null);
  let madeAt = $state<string | null>(null);

  const made = $derived(madeAt ?? rotatedAt);
  const madeOn = $derived(made === null ? '' : formatMadeOn(made));

  async function onRotate(): Promise<void> {
    busy = true;
    error = '';
    try {
      const result = await rotateInvite(guildId);
      url = inviteUrl(guildInviteHref(result.token));
      madeAt = result.rotated_at;
      confirming = false;
      onRotated?.(result.rotated_at);
    } catch (thrown) {
      error = thrown instanceof GuildApiError ? thrown.message : guildSettingsCopy.actionFailed;
    } finally {
      busy = false;
    }
  }
</script>

{#snippet rotate(cls: string)}
  <button
    class="{cls} {busy ? BUSY_CLASS : ''}"
    onclick={() => (confirming = true)}
    disabled={busy || disabled}
    aria-busy={busy}
    data-testid={`${testidPrefix}-rotate`}
  >
    {guildSettingsCopy.rotateButton}
  </button>
{/snippet}

{#snippet details()}
  {#if confirming}
    <div
      class="border-line-soft flex flex-col gap-3 border p-4"
      data-testid={`${testidPrefix}-rotate-confirm`}
    >
      <p class="text-[13px]">{guildSettingsCopy.rotateConfirmLine}</p>
      <div class="flex gap-3">
        <button
          class="{button('secondary', 'sm')} {busy ? BUSY_CLASS : ''}"
          onclick={() => void onRotate()}
          disabled={busy}
          aria-busy={busy}
          data-testid={`${testidPrefix}-rotate-confirm-button`}
        >
          {guildSettingsCopy.rotateConfirmButton}
        </button>
        <button class={button('ghost', 'sm')} onclick={() => (confirming = false)} disabled={busy}>
          {guildSettingsCopy.rotateCancel}
        </button>
      </div>
    </div>
  {/if}
  {#if url !== null}
    {#key url}
      <CopyLinkField
        {url}
        label={guildSettingsCopy.linkLabel}
        shareTitle={guildSettingsCopy.inviteShareTitle}
        shareText={guildSettingsCopy.inviteShareText}
        testid={testidPrefix}
      >
        {#snippet note()}
          <span class="text-muted text-[13px]">{guildSettingsCopy.tokenShownOnce}</span>
        {/snippet}
      </CopyLinkField>
    {/key}
  {/if}
  {#if madeOn !== ''}
    <p class="text-muted text-[12px]" data-testid={`${testidPrefix}-made`}>
      {guildSettingsCopy.currentLinkMade(madeOn)}
    </p>
  {/if}
  {#if error !== ''}
    <p class="text-[13px]" role="alert" data-testid={`${testidPrefix}-error`}>{error}</p>
  {/if}
{/snippet}

{#if row}
  <div data-testid={`${testidPrefix}-block`}>
    <SettingsRow label={guildSettingsCopy.inviteHeading} sub={guildSettingsCopy.inviteHelp} controlBelow>
      {#snippet control()}
        {#if !confirming}{@render rotate(button('secondary', 'sm'))}{/if}
      {/snippet}
    </SettingsRow>
    {#if confirming || url !== null || madeOn !== '' || error !== ''}
      <div class="flex flex-col gap-3 px-5 pb-4">{@render details()}</div>
    {/if}
  </div>
{:else}
  <div class="flex w-full flex-col gap-3" data-testid={`${testidPrefix}-block`}>
    {#if !confirming}{@render rotate(`${button('secondary')} w-fit`)}{/if}
    {@render details()}
  </div>
{/if}
