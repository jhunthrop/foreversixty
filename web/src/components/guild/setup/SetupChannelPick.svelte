<!-- web/src/components/guild/setup/SetupChannelPick.svelte -->
<!-- One of the bot's server channels and Save (guild onboarding spec §3): the raid channel on the
     Discord step, the recruiting channel on Recruiting. Reads the server's channels on mount. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import type { GuildBotChannel } from '../../../lib/guild/api';
  import { fetchBotChannels } from '../../../lib/guild/discord-bot-api';
  import { setupCopy } from '../../../lib/guild/setup-copy';
  import type { SelectOption } from '../../../lib/ui/listbox-model';
  import Button from '../../ui/Button.svelte';
  import Select from '../../ui/Select.svelte';

  let {
    guildId,
    label,
    current = null,
    busy = false,
    onSave,
    testid,
  }: {
    guildId: number;
    label: string;
    /** The channel saved now, picked first. */
    current?: string | null;
    busy?: boolean;
    onSave: (channelId: string) => void;
    testid: string;
  } = $props();

  let channels = $state<GuildBotChannel[] | null>(null);
  let failed = $state(false);
  let value = $state('');

  const options = $derived<SelectOption[]>(
    (channels ?? []).map((c) => ({ value: c.id, label: `#${c.name}` })),
  );

  onMount(() => {
    fetchBotChannels(guildId)
      .then((list) => {
        channels = list;
        value = list.find((c) => c.name === current)?.id ?? list[0]?.id ?? '';
      })
      .catch(() => (failed = true));
  });
</script>

<div class="flex flex-wrap items-center gap-3" data-testid={testid}>
  {#if failed || (channels !== null && channels.length === 0)}
    <span class="text-muted text-[13px]">{setupCopy.discord.noChannels}</span>
  {:else}
    <span class="text-[14px]">{label}</span>
    <div class="w-56">
      <Select
        bind:value
        {options}
        {label}
        size="compact"
        loading={channels === null}
        disabled={busy}
        testid={`${testid}-select`}
      />
    </div>
    <Button size="sm" disabled={busy || value === ''} onclick={() => onSave(value)} testid={`${testid}-save`}
      >{setupCopy.discord.save}</Button
    >
  {/if}
</div>
