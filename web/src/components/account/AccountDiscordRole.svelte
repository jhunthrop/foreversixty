<!-- web/src/components/account/AccountDiscordRole.svelte -->
<!-- The You panel's Discord Premium role row: what the role link does and the button that starts
     it (Discord Linked Roles). Shown to every signed-in account while the deployment has Discord
     on, Premium or not: the role follows Premium from then on. The callback's outcome is read
     from `?discord_role=` once; `start` (back from sign-in) goes straight on to the start route. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import type { Me } from '../../lib/account/api';
  import { discordRoleCopy } from '../../lib/account/discord-role-copy';
  import {
    discordRoleReturn,
    discordRoleStartUrl,
    takeDiscordRoleFlag,
  } from '../../lib/account/discord-role';
  import { button } from '../../lib/ui/button';
  import SettingsRow from '../ui/SettingsRow.svelte';
  import DiscordConnectLink from './discord/DiscordConnectLink.svelte';

  let { me }: { me: Me } = $props();

  let line = $state<{ message: string; role: 'status' | 'alert' } | null>(null);

  onMount(() => {
    const flag = takeDiscordRoleFlag();
    if (flag === null) return;
    const outcome = discordRoleReturn(flag);
    if (outcome.navigate) {
      window.location.assign(discordRoleStartUrl());
      return;
    }
    line = { message: outcome.message, role: outcome.role };
  });
</script>

{#if me.discord_link_available === true}
  <SettingsRow label={discordRoleCopy.label} sub={discordRoleCopy.lead} testid="account-discord-role">
    {#snippet below()}
      <div class="mt-4 flex flex-col items-start gap-2">
        <DiscordConnectLink
          label={discordRoleCopy.link}
          class={button('secondary')}
          testid="discord-role-start"
          href={discordRoleStartUrl()}
          onStuck={() => (line = { message: discordRoleCopy.stuck, role: 'status' })}
        />
        <div class="min-h-[21px] text-[13px]" role={line?.role ?? 'status'} data-testid="discord-role-status">
          {#if line !== null}
            <p class={line.role === 'alert' ? 'text-wipe' : 'text-text'}>{line.message}</p>
          {/if}
        </div>
      </div>
    {/snippet}
  </SettingsRow>
{/if}
