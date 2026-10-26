<!-- web/src/components/GuildShell.svelte -->
<!-- Every /guild/* path is one built asset (guild.html, worker.ts's SHELL_ROUTES), so this
     is the one place that picks which of the four guild views to render, by reading the
     path -- passed as a prop so this is testable with svelte/server's render (which never
     sees `window`) and falling back to window.location.pathname at runtime, the same
     pattern Guild.svelte's own `resolved` already uses. -->
<script lang="ts">
  import {
    isGuildLandingPath,
    parseGuildClaimPath,
    parseGuildInviteToken,
    parseGuildPath,
    parseGuildRaidPath,
    parseGuildSettingsPath,
  } from '../lib/characters';
  import { createLazyComponent } from '../lib/report/lazy-component.svelte';
  import Guild from './Guild.svelte';
  import GuildClaim from './GuildClaim.svelte';
  import GuildJoin from './GuildJoin.svelte';
  import GuildLanding from './GuildLanding.svelte';
  import GuildSettings from './GuildSettings.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  // The raid page (raids spec 3.2) is its own chunk: this island's entry is at its budget.
  const raidLazy = createLazyComponent(() => import('./GuildRaid.svelte'));

  let { path = null }: { path?: string | null } = $props();

  const resolvedPath = $derived(path ?? (typeof window === 'undefined' ? '' : window.location.pathname));

  const claimPath = $derived(parseGuildClaimPath(resolvedPath));
  const settingsPath = $derived(parseGuildSettingsPath(resolvedPath));
  const inviteToken = $derived(parseGuildInviteToken(resolvedPath));
  const raidPath = $derived(parseGuildRaidPath(resolvedPath));
  const guildPath = $derived(parseGuildPath(resolvedPath));

  $effect(() => {
    if (raidPath !== null) raidLazy.load();
  });
</script>

{#if isGuildLandingPath(resolvedPath)}
  <GuildLanding />
{:else if claimPath !== null}
  <GuildClaim path={claimPath} />
{:else if settingsPath !== null}
  <GuildSettings path={settingsPath} />
{:else if inviteToken !== null}
  <GuildJoin token={inviteToken} />
{:else if raidPath !== null}
  {#if raidLazy.current}
    <raidLazy.current path={raidPath} />
  {:else}
    <!-- The band's ready height while the chunk loads (raids spec 4, Page states); the chunk's own
         loading view then names it Loading the raid. -->
    <div class="bg-raised full-bleed min-h-[300px] md:min-h-[232px]">
      {#if raidLazy.error === ''}
        <Skeleton lines={0} testid="raid-chunk-skeleton" />
      {:else}
        <LoadError message={raidLazy.error} onRetry={() => raidLazy.load()} testid="raid-chunk-error" />
      {/if}
    </div>
  {/if}
{:else}
  <Guild path={guildPath} />
{/if}
