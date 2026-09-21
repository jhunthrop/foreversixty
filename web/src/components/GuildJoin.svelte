<!-- web/src/components/GuildJoin.svelte -->
<!-- /guild/invite/<token> (spec section 4.4): a minimal landing page. Signed out, it shows
     the guild's name (not fetchable here per plan ruling 6 -- no unauthenticated
     token-peek endpoint exists -- so the sign-in-first line names no guild) and the sign-in
     prompt; signed in, a "Join as a member" button that calls POST
     /v1/guilds/invite/{token}/accept and redirects to the guild's own page. -->
<script lang="ts">
  import { fetchMeOnce, type Me } from '../lib/account/api';
  import { guildHref } from '../lib/characters';
  import { GuildApiError, acceptInvite, type InviteAcceptResult } from '../lib/guild/api';
  import { guildJoinCopy } from '../lib/guild/copy';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import SignInPrompt from './SignInPrompt.svelte';

  let { token }: { token: string } = $props();

  let me = $state<Me | null>(null);
  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let joined = $state<InviteAcceptResult | null>(null);
  let error = $state('');
  let busy = $state(false);

  async function load(): Promise<void> {
    status = 'loading';
    try {
      me = await fetchMeOnce();
      status = 'ready';
    } catch {
      status = 'failed';
    }
  }

  $effect(() => {
    void load();
  });

  const signedIn = $derived(me !== null);

  async function onJoin(): Promise<void> {
    busy = true;
    error = '';
    try {
      joined = await acceptInvite(token);
      window.location.assign(guildHref(joined.guild.region, joined.guild.ruleset, joined.guild.name));
    } catch (thrown) {
      error = thrown instanceof GuildApiError ? thrown.message : guildJoinCopy.failed;
      busy = false;
    }
  }
</script>

<div class="flex flex-col gap-4" data-testid="guild-join">
  {#if status === 'loading'}
    <p class="text-muted text-[14px]">{guildJoinCopy.loading}</p>
  {:else if status === 'failed'}
    <p class="text-[14px]" role="alert" data-testid="guild-join-error">{guildJoinCopy.failed}</p>
  {:else if !signedIn}
    <SignInPrompt line={guildJoinCopy.signInLine} testid="guild-join-signin" />
  {:else}
    <button
      class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4"
      onclick={() => void onJoin()}
      disabled={busy}
      data-testid="guild-join-button"
    >
      {guildJoinCopy.joinButton}
    </button>
    <div class="min-h-[21px]">
      {#if error !== ''}<p class="text-[14px]" role="alert" data-testid="guild-join-action-error">
          {error}
        </p>{/if}
    </div>
  {/if}
</div>
