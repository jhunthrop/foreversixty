<!-- web/src/components/GuildShell.svelte -->
<!-- Every /guild/* path is one built asset (guild.html, worker.ts's SHELL_ROUTES), so this
     is the one place that picks which of the guild views to render, by reading the
     path -- passed as a prop so this is testable with svelte/server's render (which never
     sees `window`) and falling back to window.location.pathname at runtime, the same
     pattern Guild.svelte's own `resolved` already uses. The guild page and the landing ship
     with this island; the raid, settings, claim and join views are route-level chunks, each
     held at its own loading height while it loads, with a retry if the chunk fails. -->
<script lang="ts">
  import {
    isGuildLandingPath,
    parseGuildClaimPath,
    parseGuildInviteToken,
    parseGuildPath,
    parseGuildRaidPath,
    parseGuildSettingsPath,
  } from '../lib/characters';
  import { createLazyComponent, type LazyLoadState } from '../lib/report/lazy-component.svelte';
  import Guild from './Guild.svelte';
  import GuildLanding from './GuildLanding.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';
  import { GUILD_LOADING } from '../lib/guild/layout';
  import { GUILD_VIEW_MIN, RAID_BAND_MIN } from '../lib/guild/raid-layout';

  // This island's entry is at its budget: each view off the guild page is its own chunk.
  const raidLazy = createLazyComponent(() => import('./GuildRaid.svelte'));
  const settingsLazy = createLazyComponent(() => import('./GuildSettings.svelte'));
  const claimLazy = createLazyComponent(() => import('./GuildClaim.svelte'));
  const joinLazy = createLazyComponent(() => import('./GuildJoin.svelte'));

  let { path = null }: { path?: string | null } = $props();

  const resolvedPath = $derived(path ?? (typeof window === 'undefined' ? '' : window.location.pathname));

  const claimPath = $derived(parseGuildClaimPath(resolvedPath));
  const settingsPath = $derived(parseGuildSettingsPath(resolvedPath));
  const inviteToken = $derived(parseGuildInviteToken(resolvedPath));
  const raidPath = $derived(parseGuildRaidPath(resolvedPath));
  const guildPath = $derived(parseGuildPath(resolvedPath));

  $effect(() => {
    if (isGuildLandingPath(resolvedPath)) return;
    if (claimPath !== null) claimLazy.load();
    else if (settingsPath !== null) settingsLazy.load();
    else if (inviteToken !== null) joinLazy.load();
    else if (raidPath !== null) raidLazy.load();
  });
</script>

<!-- A view's chunk in flight: its loading height, or the chunk's failure with a retry. -->
{#snippet chunkPending(lazy: LazyLoadState, testid: string, minHeight: string)}
  {#if lazy.error === ''}
    <Skeleton lines={0} {minHeight} testid={`${testid}-skeleton`} />
  {:else}
    <LoadError message={lazy.error} onRetry={() => lazy.load()} testid={`${testid}-error`} />
  {/if}
{/snippet}

<!-- One screen tall from the first paint (the server's markup), so the footer starts below the fold
     and no view's growth moves it into or out of the viewport. The gap is main's own. -->
<div class="flex flex-col gap-[22px] md:gap-8 {GUILD_VIEW_MIN}" data-testid="guild-view">
  {#if isGuildLandingPath(resolvedPath)}
    <GuildLanding />
  {:else if claimPath !== null}
    {#if claimLazy.current}
      <claimLazy.current path={claimPath} />
    {:else}
      {@render chunkPending(claimLazy, 'guild-claim-chunk', GUILD_LOADING.claim.minHeight)}
    {/if}
  {:else if settingsPath !== null}
    {#if settingsLazy.current}
      <settingsLazy.current path={settingsPath} />
    {:else}
      {@render chunkPending(settingsLazy, 'guild-settings-chunk', GUILD_LOADING.settings.minHeight)}
    {/if}
  {:else if inviteToken !== null}
    {#if joinLazy.current}
      <joinLazy.current token={inviteToken} />
    {:else}
      {@render chunkPending(joinLazy, 'guild-join-chunk', GUILD_LOADING.join.minHeight)}
    {/if}
  {:else if raidPath !== null}
    {#if raidLazy.current}
      <raidLazy.current path={raidPath} />
    {:else}
      <!-- The band's ready height while the chunk loads (raids spec 4, Page states); the chunk's own
         loading view then names it Loading the raid. -->
      <div class="bg-raised full-bleed {RAID_BAND_MIN}">
        {@render chunkPending(raidLazy, 'raid-chunk', '')}
      </div>
    {/if}
  {:else}
    <Guild path={guildPath} />
  {/if}
</div>
