<!-- web/src/components/guild/setup/SetupChannelPick.svelte -->
<!-- One of the bot's server channels and Save (guild onboarding spec §3): the raid channel on the
     Discord step, the recruiting channel on Recruiting. Reads the server's channels on mount. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import type { GuildBotChannel } from '../../../lib/guild/api';
  import { fetchBotChannels } from '../../../lib/guild/discord-bot-api';
  import { guildBotCopy } from '../../../lib/guild/discord-bot-copy';
  import { setupCopy } from '../../../lib/guild/setup-copy';
  import type { ButtonTone } from '../../../lib/ui/button';
  import type { SelectOption } from '../../../lib/ui/listbox-model';
  import Button from '../../ui/Button.svelte';
  import Select from '../../ui/Select.svelte';

  let {
    guildId,
    label,
    current = null,
    prefer,
    tone = 'secondary',
    busy = false,
    onSave,
    testid,
  }: {
    guildId: number;
    label: string;
    /** The channel saved now, picked first. */
    current?: string | null;
    /** With nothing saved, the first channel whose name matches is picked; none matching, the
     *  officer picks (no channel is guessed for them). Without it the first channel is picked. */
    prefer?: RegExp;
    tone?: ButtonTone;
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
        const saved = list.find((c) => c.name === current);
        const guess = prefer === undefined ? list[0] : list.find((c) => prefer.test(c.name));
        value = (saved ?? guess)?.id ?? '';
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
        placeholder={guildBotCopy.picker.placeholder}
        size="compact"
        loading={channels === null}
        disabled={busy}
        testid={`${testid}-select`}
      />
    </div>
    <Button
      size="sm"
      {tone}
      disabled={busy || value === ''}
      onclick={() => onSave(value)}
      testid={`${testid}-save`}>{setupCopy.discord.save}</Button
    >
  {/if}
</div>
