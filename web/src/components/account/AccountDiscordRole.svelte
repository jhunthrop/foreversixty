<!-- web/src/components/account/AccountDiscordRole.svelte -->
<!-- The Discord section's Premium role row (premium spec 2026-11-08 section 7.6): what the role link
     does and the button that starts it (Discord Linked Roles); an account that already has Premium
     reads that linking turns the role on (ruling A4). Shown to every signed-in account while the deployment has Discord
     on, Premium or not: the role follows Premium from then on. The callback's outcome is read
     from `?discord_role=` once; `start` (back from sign-in) goes straight on to the start route. -->
<script lang="ts">
  import { smallButton } from '../../lib/account/small-button';
  import { onMount } from 'svelte';
  import type { Me } from '../../lib/account/api';
  import { discordRoleCopy } from '../../lib/account/discord-role-copy';
  import {
    discordRoleReturn,
    discordRoleStartUrl,
    takeDiscordRoleFlag,
  } from '../../lib/account/discord-role';
  import { premiumSource } from '../../lib/account/plan-line';
  import Row from '../ui/Row.svelte';
  import DiscordConnectLink from './discord/DiscordConnectLink.svelte';

  let { me, now = new Date() }: { me: Me; now?: Date } = $props();

  const lead = $derived(premiumSource(me, now) === 'none' ? discordRoleCopy.lead : discordRoleCopy.leadOwned);

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
  <div data-testid="account-discord-role">
    <Row title={discordRoleCopy.label} sub={lead} open wrapTrail>
      {#snippet trail()}
        <DiscordConnectLink
          label={discordRoleCopy.link}
          class={smallButton()}
          testid="discord-role-start"
          href={discordRoleStartUrl()}
          onStuck={() => (line = { message: discordRoleCopy.stuck, role: 'status' })}
        />
      {/snippet}
    </Row>
    {#if line !== null}
      <p
        class="pt-2 text-[13px] {line.role === 'alert' ? 'text-wipe' : 'text-text'}"
        role={line.role}
        data-testid="discord-role-status"
      >
        {line.message}
      </p>
    {/if}
  </div>
{/if}
