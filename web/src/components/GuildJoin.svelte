<!-- web/src/components/GuildJoin.svelte -->
<!-- /guild/invite/<token> (spec 2026-11-05 §3.4): one card, centred. Which guild, which of the
     viewer's characters joins, one `Join`. The token is looked up (GET /v1/guild-invites/{token},
     public: it names the guild without joining). Signed out, the card carries the sign-in button;
     an unknown or rotated token says so plainly with no button; a member of the guild already is
     told so. Joining calls POST /v1/guilds/invite/{token}/accept and shows the joined state: the
     guild's plan line, then `Open the guild`. The inviter is not named until the link records who
     made it (§4.1, held). -->
<script lang="ts">
  import { GUILD_BANDLESS_TOP } from '../lib/guild/raid-layout';
  import { battlenetStartUrl, fetchMeOnce, forgetSession, type Me } from '../lib/account/api';
  import { guildHref, guildInviteHref, rulesetLabel } from '../lib/characters';
  import {
    GuildApiError,
    acceptInvite,
    peekInvite,
    type InviteAcceptResult,
    type InvitePeekResult,
  } from '../lib/guild/api';
  import { guildJoinCopy } from '../lib/guild/copy';
  import { joinedPlanLine } from '../lib/guild/join-plan';
  import { joiningRow } from '../lib/guild/join-view';
  import { GUILD_LOADING } from '../lib/guild/layout';
  import { guildMarkSrc } from '../lib/guild/mark';
  import { classColorVar } from '../lib/report/format';
  import { BUSY_CLASS } from '../lib/ui/busy';
  import { button } from '../lib/ui/button';
  import ClassCrestRing from './character/ClassCrestRing.svelte';
  import FactionCrest from './character/FactionCrest.svelte';
  import GuildJoinedBlock from './guild/GuildJoinedBlock.svelte';
  import GuildStatus from './GuildStatus.svelte';
  import EmptyState from './ui/EmptyState.svelte';
  import SettingsRow from './ui/SettingsRow.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  const CARD = 'card-surface mx-auto flex w-full max-w-[560px] flex-col gap-4 rounded-panel p-5';

  let { token }: { token: string } = $props();

  let me = $state<Me | null>(null);
  let peek = $state<InvitePeekResult | null>(null);
  let status = $state<'loading' | 'ready' | 'invalid' | 'failed'>('loading');
  let joined = $state<InviteAcceptResult | null>(null);
  let error = $state('');
  let busy = $state(false);
  let planLine = $state<string | null>(null);

  async function load(): Promise<void> {
    status = 'loading';
    try {
      // The lookup answers a 404 for a bad token: that is a state, not a failure.
      const [session, result] = await Promise.all([
        fetchMeOnce(),
        peekInvite(token).catch((thrown: unknown) => {
          if (thrown instanceof GuildApiError && thrown.status === 404) return null;
          throw thrown;
        }),
      ]);
      me = session;
      peek = result;
      status = result === null ? 'invalid' : 'ready';
    } catch {
      status = 'failed';
    }
  }

  $effect(() => {
    void load();
  });

  const guild = $derived(peek?.guild ?? null);
  const signedIn = $derived(me !== null);
  const alreadyIn = $derived(guild !== null && (me?.guilds.some((g) => g.id === guild.id) ?? false));
  const realm = $derived(
    guild === null ? '' : `${rulesetLabel(guild.ruleset)} ${guild.region.toUpperCase()}`,
  );
  const joining = $derived(guild === null ? null : joiningRow(peek?.joins_as, guild.name));
  /** A member already sees row 3 only for a character the guild's roster names. */
  const showJoining = $derived(
    joined === null && signedIn && (!alreadyIn || peek?.joins_as?.on_roster === true),
  );

  /** The guild's plan sentence from a fresh `/v1/me`; null when that read fails (the line is left out). */
  async function readPlanLine(guildId: number, guildName: string): Promise<string | null> {
    forgetSession();
    try {
      const fresh = await fetchMeOnce();
      const entry = fresh?.guilds.find((g) => g.id === guildId);
      return entry === undefined ? null : joinedPlanLine(guildName, entry.covered);
    } catch {
      return null;
    }
  }

  async function onJoin(): Promise<void> {
    busy = true;
    error = '';
    try {
      joined = await acceptInvite(token);
      busy = false;
      planLine = await readPlanLine(joined.guild.id, joined.guild.name);
    } catch (thrown) {
      error = thrown instanceof GuildApiError ? thrown.message : guildJoinCopy.failed;
      busy = false;
    }
  }
</script>

<div class="reveal {GUILD_BANDLESS_TOP} flex flex-col gap-4" data-testid="guild-join">
  {#if status === 'loading'}
    <div class="{CARD} min-h-[380px] md:min-h-[320px]" data-testid="guild-join-loading">
      <Skeleton lines={6} testid="guild-join-skeleton" />
    </div>
  {:else if status === 'failed'}
    <GuildStatus
      status="failed"
      error={guildJoinCopy.failed}
      onRetry={() => void load()}
      lines={GUILD_LOADING.join.lines}
      minHeight={GUILD_LOADING.join.minHeight}
      testid="guild-join"
    />
  {:else if status === 'invalid' || guild === null}
    <EmptyState message={guildJoinCopy.invalid} testid="guild-join-invalid" />
  {:else}
    <div class={CARD} data-testid="guild-join-card">
      <div class="flex items-center gap-3">
        <FactionCrest faction={guild.faction} src={guildMarkSrc(guild)} small testid="guild-join-crest" />
        <div class="flex min-w-0 flex-col">
          <h1 class="heading-plan" data-testid="guild-join-name">{guild.name}</h1>
          <p class="text-muted text-[13px]" data-testid="guild-join-realm">
            {guildJoinCopy.guildLine(realm, peek?.members)}
          </p>
        </div>
      </div>
      {#if !alreadyIn || joined !== null}
        <p class="text-text text-[15px]" data-testid="guild-join-invited">
          {guildJoinCopy.invited(guild.name)}
        </p>
      {/if}
      {#if showJoining && joining !== null}
        <div class="border-line-soft -mx-5 border-b" data-testid="guild-join-as">
          <SettingsRow
            label={joining.label}
            labelColor={joining.labelClass === null ? undefined : classColorVar(joining.labelClass)}
            sub={joining.sub}
          >
            {#snippet icon()}
              <ClassCrestRing characterClass={joining.crestClass} size={32} />
            {/snippet}
          </SettingsRow>
        </div>
      {/if}
      {#if joined !== null}
        <GuildJoinedBlock
          {planLine}
          openHref={guildHref(joined.guild.region, joined.guild.ruleset, joined.guild.name)}
        />
      {:else if !signedIn}
        <p class="text-[14px]" data-testid="guild-join-signin-line">{guildJoinCopy.signInLine}</p>
        <a
          class="{button('primary', 'md')} w-full md:w-fit"
          href={battlenetStartUrl(guildInviteHref(token))}
          data-testid="guild-join-signin"
        >
          {guildJoinCopy.signInButton}
        </a>
      {:else if alreadyIn}
        <p class="text-[14px]" data-testid="guild-join-already">{guildJoinCopy.alreadyIn(guild.name)}</p>
        <a
          class="{button('primary', 'md')} w-full md:w-fit"
          href={guildHref(guild.region, guild.ruleset, guild.name)}
          data-testid="guild-join-open"
        >
          {guildJoinCopy.openGuild}
        </a>
      {:else}
        <div class="flex flex-col gap-1">
          <p class="text-muted text-[13px]" data-testid="guild-join-verified">
            {guildJoinCopy.verifiedAtOnce}
          </p>
          <p class="text-muted text-[13px]" data-testid="guild-join-consequence">
            {guildJoinCopy.consequence}
          </p>
        </div>
        <div class="flex flex-col gap-3 md:flex-row md:items-center">
          <button
            class="{button('primary', 'md')} w-full md:w-fit {busy ? BUSY_CLASS : ''}"
            onclick={() => void onJoin()}
            disabled={busy}
            aria-busy={busy}
            data-testid="guild-join-button"
          >
            {busy ? guildJoinCopy.joining : guildJoinCopy.joinButton(guild.name)}
          </button>
          <a
            class="{button('ghost', 'md')} w-full md:w-fit"
            href={guildHref(guild.region, guild.ruleset, guild.name)}
            data-testid="guild-join-not-now"
          >
            {guildJoinCopy.notNow}
          </a>
        </div>
        <div class="min-h-[21px]">
          {#if error !== ''}<p
              class="text-wipe text-[14px]"
              role="alert"
              data-testid="guild-join-action-error"
            >
              {error}
            </p>{/if}
        </div>
      {/if}
    </div>
  {/if}
</div>
