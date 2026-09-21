<!-- web/src/components/GuildShell.svelte -->
<!-- Every /guild/* path is one built asset (guild.html, worker.ts's SHELL_ROUTES), so this
     is the one place that picks which of the four guild views to render, by reading the
     path -- passed as a prop so this is testable with svelte/server's render (which never
     sees `window`) and falling back to window.location.pathname at runtime, the same
     pattern Guild.svelte's own `resolved` already uses. -->
<script lang="ts">
  import {
    parseGuildClaimPath,
    parseGuildInviteToken,
    parseGuildPath,
    parseGuildSettingsPath,
  } from '../lib/characters';
  import Guild from './Guild.svelte';
  import GuildClaim from './GuildClaim.svelte';
  import GuildJoin from './GuildJoin.svelte';
  import GuildSettings from './GuildSettings.svelte';

  let { path = null }: { path?: string | null } = $props();

  const resolvedPath = $derived(path ?? (typeof window === 'undefined' ? '' : window.location.pathname));

  const claimPath = $derived(parseGuildClaimPath(resolvedPath));
  const settingsPath = $derived(parseGuildSettingsPath(resolvedPath));
  const inviteToken = $derived(parseGuildInviteToken(resolvedPath));
  const guildPath = $derived(parseGuildPath(resolvedPath));
</script>

{#if claimPath !== null}
  <GuildClaim path={claimPath} />
{:else if settingsPath !== null}
  <GuildSettings path={settingsPath} />
{:else if inviteToken !== null}
  <GuildJoin token={inviteToken} />
{:else}
  <Guild path={guildPath} />
{/if}
