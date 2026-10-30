<!-- web/src/components/guild/GuildDiscordChannelPick.svelte -->
<!-- The bot's channel picker (bot spec §3.1, state 1): reached on the install return and from
     `Change channel`. The list is fetched when this opens (loading, empty, error, ready); the
     current channel, else the first raid channel, is preselected. An empty `Save channel` press
     says `Choose a channel.`; a 422 `channel_unusable` refetches, clears the select and says so.
     The remove question replaces the actions row, as everywhere else in the section. -->
<script lang="ts">
  import { onMount, tick } from 'svelte';
  import {
    GuildApiError,
    type GuildBotChannel,
    type GuildChannelUnusable,
    type GuildSettingsDiscord,
  } from '../../lib/guild/api';
  import { guildBotCopy, shortServerName } from '../../lib/guild/discord-bot-copy';
  import { fetchBotChannels, saveBotChannel } from '../../lib/guild/discord-bot-api';
  import { groupChannels, preselectAfterUnusable, preselectChannel } from '../../lib/guild/discord-bot-state';
  import type { DiscordStatus } from '../../lib/guild/discord-settings';
  import { LINK_BUTTON } from '../../lib/account/discord-styles';
  import { button } from '../../lib/ui/button';
  import Select from '../ui/Select.svelte';
  import DiscordStatusLine from '../account/discord/DiscordStatusLine.svelte';
  import GuildBotRemoveQuestion from './GuildBotRemoveQuestion.svelte';

  const OUTLINE = button('secondary');
  const CHANNEL_UNUSABLE = 'channel_unusable';
  const PLAN_REQUIRED = 'guild_plan_required';
  const SELECT_ID = 'guild-bot-channel';
  const EMPTY_PRESS_ID = 'guild-bot-channel-empty';

  type ListState = 'loading' | 'ready' | 'empty' | 'error';

  let {
    guildId,
    server,
    currentChannelId,
    changing,
    webhook,
    gated,
    frozen,
    status,
    confirming,
    removing,
    focusSelect,
    unusable,
    onSaving,
    onStatus,
    onSaved,
    onCancel,
    onAskRemove,
    onRemove,
    onCancelRemove,
  }: {
    guildId: number;
    server: string;
    /** The saved channel on `Change channel`; null on the first pick. */
    currentChannelId: string | null;
    changing: boolean;
    webhook: boolean;
    /** Officer-read-only or frozen: the select and `Save channel` are off. */
    gated: boolean;
    frozen: boolean;
    status: DiscordStatus | null;
    confirming: boolean;
    removing: boolean;
    /** Move focus to the select once the list is ready (the install return, `Change channel`). */
    focusSelect: boolean;
    /** Why the saved channel cannot be used, when the picker was opened from that notice. */
    unusable: GuildChannelUnusable | null;
    onSaving: (saving: boolean) => void;
    onStatus: (status: DiscordStatus | null, focusLine?: boolean) => void;
    onSaved: (discord: GuildSettingsDiscord) => void;
    onCancel: () => void;
    onAskRemove: () => void;
    onRemove: () => void;
    onCancelRemove: () => void;
  } = $props();

  let listState = $state<ListState>('loading');
  let channels = $state<GuildBotChannel[]>([]);
  let selected = $state('');
  let emptyPress = $state(false);
  let saving = $state(false);

  $effect(() => {
    onSaving(saving);
    return () => onSaving(false);
  });

  const channelOptions = $derived([
    { value: '', label: guildBotCopy.picker.placeholder },
    ...groupChannels(channels).flatMap((group) =>
      group.channels.map((channel) => ({
        value: channel.id,
        label: guildBotCopy.installed.channel(channel.name),
        group: group.category ?? undefined,
      })),
    ),
  ]);
  const loadingOptions = [{ value: '', label: guildBotCopy.picker.loading }];
  const serverName = $derived(shortServerName(server));
  const saveOff = $derived(gated || saving || listState === 'loading');

  function initialChoice(): string | null {
    if (unusable === null) return preselectChannel(channels, currentChannelId);
    return preselectAfterUnusable(channels, currentChannelId, unusable.reason);
  }

  async function load(preselect: boolean): Promise<void> {
    listState = 'loading';
    try {
      channels = await fetchBotChannels(guildId);
    } catch {
      listState = 'error';
      return;
    }
    selected = preselect ? (initialChoice() ?? '') : '';
    listState = channels.length === 0 ? 'empty' : 'ready';
    if (listState === 'ready' && focusSelect) {
      await tick();
      document.getElementById(SELECT_ID)?.focus();
    }
  }

  onMount(() => void load(true));

  function choose(value: string): void {
    selected = value;
    emptyPress = false;
  }

  async function onSaveFailed(error: unknown): Promise<void> {
    if (error instanceof GuildApiError && error.code === PLAN_REQUIRED)
      return onStatus({ ok: false, message: guildBotCopy.planEnded, planEnded: true }, true);
    if (error instanceof GuildApiError && error.code === CHANNEL_UNUSABLE) {
      await load(false);
      return onStatus({ ok: false, message: guildBotCopy.picker.race }, true);
    }
    onStatus({ ok: false, message: guildBotCopy.picker.saveFailed }, true);
  }

  async function save(): Promise<void> {
    if (selected === '') {
      emptyPress = true;
      return;
    }
    saving = true;
    onStatus(null);
    try {
      onSaved(await saveBotChannel(guildId, selected));
    } catch (error) {
      await onSaveFailed(error);
    } finally {
      saving = false;
    }
  }
</script>

<p class="text-[14px]" data-testid="bot-picker-server">
  <span class="text-muted">{changing ? guildBotCopy.picker.serverLead : guildBotCopy.picker.added}</span>
  <span class="text-strong font-semibold">{serverName}</span>
</p>

{#if listState === 'empty'}
  <p class="text-[14px]" data-testid="bot-channels-empty">{guildBotCopy.picker.empty(server)}</p>
{:else if listState === 'error'}
  <div class="flex flex-col gap-2">
    <span class="quiet-label">{guildBotCopy.picker.label}</span>
    <p class="text-wipe text-[13px]" role="alert" data-testid="bot-channels-error">
      {guildBotCopy.picker.loadFailed}
    </p>
  </div>
{:else}
  <div class="flex w-full flex-col gap-2 md:w-[360px]">
    <label class="quiet-label" for={SELECT_ID}>{guildBotCopy.picker.label}</label>
    <Select
      id={SELECT_ID}
      size="compact"
      value={selected}
      options={listState === 'loading' ? loadingOptions : channelOptions}
      testid="bot-channel"
      busy={listState === 'loading'}
      disabled={gated || saving || listState === 'loading'}
      invalid={emptyPress}
      describedBy={emptyPress ? EMPTY_PRESS_ID : undefined}
      onchange={(event) => choose((event.currentTarget as HTMLSelectElement).value)}
    />
    {#if emptyPress}
      <p id={EMPTY_PRESS_ID} class="text-wipe text-[13px]" data-testid="bot-channel-empty-press">
        {guildBotCopy.picker.emptyPress}
      </p>
    {:else}
      <p class="text-muted text-[13px]">{guildBotCopy.picker.helper}</p>
    {/if}
  </div>
{/if}

<div class="flex flex-col">
  {#if confirming}
    <GuildBotRemoveQuestion
      {server}
      {webhook}
      busy={removing}
      disabled={frozen}
      {onRemove}
      onCancel={onCancelRemove}
    />
  {:else}
    <div class="flex flex-wrap items-center gap-3">
      {#if listState === 'error'}
        <button class={OUTLINE} onclick={() => void load(true)} data-testid="bot-channels-retry">
          {guildBotCopy.picker.retry}
        </button>
      {:else if listState === 'empty'}
        <button class={OUTLINE} onclick={() => void load(true)} data-testid="bot-channels-reload">
          {guildBotCopy.picker.reload}
        </button>
      {:else}
        <button
          class="{button('primary')} {saveOff ? 'cursor-not-allowed opacity-50' : ''}"
          disabled={saveOff}
          aria-busy={saving}
          onclick={() => void save()}
          data-testid="bot-channel-save"
        >
          {saving ? guildBotCopy.picker.saving : guildBotCopy.picker.save}
        </button>
      {/if}
      {#if changing}
        <button class={LINK_BUTTON} disabled={saving} onclick={onCancel} data-testid="bot-channel-cancel">
          {guildBotCopy.picker.cancel}
        </button>
      {:else}
        <button
          class="{LINK_BUTTON} {frozen ? 'cursor-not-allowed opacity-50' : ''}"
          disabled={saving || frozen}
          onclick={onAskRemove}
          data-testid="bot-remove-bot"
        >
          {guildBotCopy.picker.removeBot}
        </button>
      {/if}
    </div>
  {/if}
  <div id="guild-discord-status" tabindex="-1" class="outline-none"><DiscordStatusLine {status} /></div>
</div>
