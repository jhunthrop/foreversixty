<!-- web/src/components/GuildJoin.svelte -->
<!-- /guild/invite/<token> (spec section 4.4): the invite card. The token is looked up
     (GET /v1/guild-invites/{token}, public: it names the guild without joining) so the page
     can say which guild invited the visitor and what joining does before the button. Signed
     out, the same card carries "Sign in with Battle.net to join"; an unknown or rotated token
     says so plainly with no button; a member of the guild already is told so. Joining calls
     POST /v1/guilds/invite/{token}/accept and goes to the guild's own page. -->
<script lang="ts">
  import { GUILD_BANDLESS_TOP } from '../lib/guild/raid-layout';
  import { battlenetStartUrl, fetchMeOnce, forgetSession, type Me } from '../lib/account/api';
  import { guildHref, guildInviteHref, rulesetLabel } from '../lib/characters';
  import {
    GuildApiError,
    acceptInvite,
    peekInvite,
    type GuildSummary,
    type InviteAcceptResult,
  } from '../lib/guild/api';
  import { guildJoinCopy } from '../lib/guild/copy';
  import { joinedPlanLine } from '../lib/guild/join-plan';
  import { GUILD_LOADING } from '../lib/guild/layout';
  import { guildMarkSrc } from '../lib/guild/mark';
  import { PRIMARY_BUTTON_FIXED, SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import FactionCrest from './character/FactionCrest.svelte';
  import GuildStatus from './GuildStatus.svelte';
  import { BUSY_CLASS } from '../lib/ui/busy';

  let { token }: { token: string } = $props();

  let me = $state<Me | null>(null);
  let guild = $state<GuildSummary | null>(null);
  let status = $state<'loading' | 'ready' | 'invalid' | 'failed'>('loading');
  let joined = $state<InviteAcceptResult | null>(null);
  let error = $state('');
  let busy = $state(false);
  let planLine = $state<string | null>(null);

  async function load(): Promise<void> {
    status = 'loading';
    try {
      // The lookup answers a 404 for a bad token: that is a state, not a failure.
      const [session, peek] = await Promise.all([
        fetchMeOnce(),
        peekInvite(token).catch((thrown: unknown) => {
          if (thrown instanceof GuildApiError && thrown.status === 404) return null;
          throw thrown;
        }),
      ]);
      me = session;
      guild = peek === null ? null : peek.guild;
      status = peek === null ? 'invalid' : 'ready';
    } catch {
      status = 'failed';
    }
  }

  $effect(() => {
    void load();
  });

  const signedIn = $derived(me !== null);
  const alreadyIn = $derived(guild !== null && (me?.guilds.some((g) => g.id === guild!.id) ?? false));
  const realm = $derived(
    guild === null ? '' : `${rulesetLabel(guild.ruleset)} ${guild.region.toUpperCase()}`,
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
      planLine = await readPlanLine(joined.guild.id, joined.guild.name);
    } catch (thrown) {
      error = thrown instanceof GuildApiError ? thrown.message : guildJoinCopy.failed;
      busy = false;
    }
  }
</script>

<div class="reveal {GUILD_BANDLESS_TOP} flex flex-col gap-4" data-testid="guild-join">
  {#if status === 'loading' || status === 'failed'}
    <GuildStatus
      status={status === 'loading' ? 'loading' : 'failed'}
      error={guildJoinCopy.failed}
      onRetry={() => void load()}
      lines={GUILD_LOADING.join.lines}
      minHeight={GUILD_LOADING.join.minHeight}
      testid="guild-join"
    />
  {:else if status === 'invalid' || guild === null}
    <p class="max-w-[480px] text-[14px]" data-testid="guild-join-invalid">{guildJoinCopy.invalid}</p>
  {:else}
    <div
      class="border-line-soft bg-raised flex w-full max-w-[480px] flex-col gap-4 rounded-md border p-5"
      data-testid="guild-join-card"
    >
      <div class="flex items-center gap-3">
        <FactionCrest faction={guild.faction} src={guildMarkSrc(guild)} testid="guild-join-crest" />
        <h1 class="section-title text-[18px]" data-testid="guild-join-name">{guild.name}</h1>
      </div>
      <p class="text-[14px]" data-testid="guild-join-invited">
        {guildJoinCopy.invitedYou(guild.name, realm)}
      </p>
      <p class="text-muted text-[13px]" data-testid="guild-join-consequence">{guildJoinCopy.consequence}</p>
      {#if joined !== null}
        <p class="text-[14px]" role="status" data-testid="guild-join-joined">{guildJoinCopy.joined}</p>
        {#if planLine !== null}
          <p class="text-muted text-[13px]" data-testid="guild-join-plan">{planLine}</p>
        {/if}
        <a
          class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4"
          href={guildHref(joined.guild.region, joined.guild.ruleset, joined.guild.name)}
          data-testid="guild-join-open"
        >
          {guildJoinCopy.openGuild}
        </a>
      {:else if !signedIn}
        <a
          class="{PRIMARY_BUTTON_FIXED} w-fit px-4"
          href={battlenetStartUrl(guildInviteHref(token))}
          data-testid="guild-join-signin"
        >
          {guildJoinCopy.signInButton}
        </a>
        <a class="text-nav w-fit text-[13px]" href="/login" data-testid="guild-join-signin-email">
          Use an email link instead.
        </a>
      {:else if alreadyIn}
        <p class="text-[14px]" data-testid="guild-join-already">{guildJoinCopy.alreadyIn(guild.name)}</p>
        <a
          class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4"
          href={guildHref(guild.region, guild.ruleset, guild.name)}
          data-testid="guild-join-open"
        >
          {guildJoinCopy.openGuild}
        </a>
      {:else}
        <div class="flex flex-wrap items-center gap-4">
          <button
            class={`${PRIMARY_BUTTON_FIXED} w-fit px-4 ${busy ? BUSY_CLASS : ''}`}
            onclick={() => void onJoin()}
            disabled={busy}
            aria-busy={busy}
            data-testid="guild-join-button"
          >
            {guildJoinCopy.joinButton(guild.name)}
          </button>
          <a
            class="text-nav text-[13px]"
            href={guildHref(guild.region, guild.ruleset, guild.name)}
            data-testid="guild-join-not-now"
          >
            {guildJoinCopy.notNow}
          </a>
        </div>
        <div class="min-h-[21px]">
          {#if error !== ''}<p class="text-[14px]" role="alert" data-testid="guild-join-action-error">
              {error}
            </p>{/if}
        </div>
      {/if}
    </div>
  {/if}
</div>
