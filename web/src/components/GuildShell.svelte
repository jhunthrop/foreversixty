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
    guildHref,
    isGuildLandingPath,
    parseGuildClaimPath,
    parseGuildInviteToken,
    parseGuildPath,
    parseGuildRaidPath,
    parseGuildSettingsPath,
  } from '../lib/characters';
  import { createLazyComponent, type LazyLoadState } from '../lib/report/lazy-component.svelte';
  import { onCurrentCharacterSwitch } from '../lib/current-character-events';
  import Guild from './Guild.svelte';
  import GuildLanding from './GuildLanding.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';
  import { GUILD_LOADING } from '../lib/guild/layout';
  import { GUILD_BANDLESS_TOP, GUILD_VIEW_MIN, RAID_BAND_MIN } from '../lib/guild/raid-layout';
  import { guildClaimCopy, guildSettingsCopy } from '../lib/guild/copy';
  import { raidCopy } from '../lib/guild/raid-copy';

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

  // Every guild view follows the header selector (character-selector/follow-page.ts): your guild
  // moves to the new character's, a guild you looked up stays. The work loads on the first choice.
  $effect(() =>
    onCurrentCharacterSwitch((change) => {
      void import('../lib/character-selector/follow-page-run').then(({ followGuildPage }) =>
        followGuildPage(change),
      );
    }),
  );

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
      <!-- The view's own wrapper and h1 (the guild's name is not known yet), so the swap moves nothing. -->
      <div class="flex flex-col gap-4 {GUILD_BANDLESS_TOP}">
        <a
          class="text-nav w-fit text-[13px]"
          href={guildHref(claimPath.region, claimPath.ruleset, claimPath.slug)}
        >
          ← {guildClaimCopy.backToGuild('')}
        </a>
        <h1 class="heading-page">{guildClaimCopy.heading('')}</h1>
        {@render chunkPending(claimLazy, 'guild-claim-chunk', GUILD_LOADING.claim.minHeight)}
      </div>
    {/if}
  {:else if settingsPath !== null}
    {#if settingsLazy.current}
      <settingsLazy.current path={settingsPath} />
    {:else}
      <div class="flex flex-col gap-6 {GUILD_BANDLESS_TOP}">
        <h1 class="heading-page">{guildSettingsCopy.heading}</h1>
        {@render chunkPending(settingsLazy, 'guild-settings-chunk', GUILD_LOADING.settings.minHeight)}
      </div>
    {/if}
  {:else if inviteToken !== null}
    {#if joinLazy.current}
      <joinLazy.current token={inviteToken} />
    {:else}
      <div class={GUILD_BANDLESS_TOP}>
        {@render chunkPending(joinLazy, 'guild-join-chunk', GUILD_LOADING.join.minHeight)}
      </div>
    {/if}
  {:else if raidPath !== null}
    {#if raidLazy.current}
      <raidLazy.current path={raidPath} />
    {:else}
      <!-- The band's ready height while the chunk loads (raids spec 4, Page states); the chunk's own
         loading view then names it Loading the raid. -->
      <div class="{RAID_BAND_MIN} {GUILD_BANDLESS_TOP}">
        <h1 class="sr-only">{raidCopy.page.eyebrow}</h1>
        {@render chunkPending(raidLazy, 'raid-chunk', '')}
      </div>
    {/if}
  {:else}
    <Guild path={guildPath} />
  {/if}
</div>
