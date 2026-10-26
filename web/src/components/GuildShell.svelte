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
  import { GUILD_VIEW_MIN, RAID_BAND_MIN } from '../lib/guild/raid-layout';

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

<!-- One screen tall from the first paint (the server's markup), so the footer starts below the fold
     and no view's growth moves it into or out of the viewport. The gap is main's own. -->
<div class="flex flex-col gap-[22px] md:gap-8 {GUILD_VIEW_MIN}" data-testid="guild-view">
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
      <div class="bg-raised full-bleed {RAID_BAND_MIN}">
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
</div>
