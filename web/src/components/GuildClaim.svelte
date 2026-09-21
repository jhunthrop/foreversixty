<!-- web/src/components/GuildClaim.svelte -->
<!-- /guild/<region>/<ruleset>/<name>/claim (spec section 4.4): current claim state, a
     claim button for an officer/leader-rank viewer, a confirm button for a second officer
     when one is pending, a release button for the current claimant. Reads the guild's
     public page for its name/id and its settings for claimed_by/claim_pending -- the spec
     defines no separate "claim status" read, and GuildSettingsData already carries every
     field this view needs (Task 1's api.ts). GET .../settings is verified-officer/leader-
     or-moderator only (spec section 2.6): a visitor who cannot read it -- signed out, an
     unverified officer, a stranger -- simply sees no claim-state section below the
     heading, the same "fail toward nothing extra" rule Guild.svelte's own member-only
     section already follows for a gated read that most visitors are expected to be
     refused. -->
<script lang="ts">
  import { fetchMeOnce, type Me } from '../lib/account/api';
  import { characterSlug, type CharacterPath } from '../lib/characters';
  import {
    claimGuild,
    confirmClaim,
    fetchGuildSettings,
    releaseClaim,
    type GuildSettingsData,
  } from '../lib/guild/api';
  import { guildClaimCopy } from '../lib/guild/copy';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { fetchGuild } from '../lib/rankings/api';
  import SignInPrompt from './SignInPrompt.svelte';

  let { path }: { path: CharacterPath } = $props();

  let me = $state<Me | null>(null);
  let guildId = $state<number | null>(null);
  let guildName = $state('');
  let settings = $state<GuildSettingsData | null>(null);
  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let error = $state('');
  let busy = $state(false);
  let justClaimedExpiry = $state<string | undefined>(undefined);

  /**
   * `fetchMeOnce()` is awaited directly rather than `.catch()`-guarded to null:
   * `fetchMe`'s own 401/403 handling already resolves it to null for a genuinely
   * signed-out visitor, so anything it throws past that is a real failure -- a network
   * error, a 500 -- and swallowing that into "signed out" would show a signed-in officer
   * the sign-in prompt, or "claimed by someone else" instead of "claimed by you", over an
   * error that has nothing to do with their session. Unlike Guild.svelte's member-only
   * home section, `me` here is not an optional bonus read: it decides which of the four
   * claim states the visitor sees, so a real failure belongs in this page's own failed
   * state -- the same rule Account.svelte's `load()` follows by not guarding its own
   * `fetchMeOnce()` call either.
   */
  async function load(): Promise<void> {
    status = 'loading';
    error = '';
    try {
      const [session, guildPage] = await Promise.all([fetchMeOnce(), fetchGuild(path)]);
      me = session;
      guildId = guildPage.guild.id;
      guildName = guildPage.guild.name;
      settings = await fetchGuildSettings(guildPage.guild.id).catch(() => null);
      status = 'ready';
    } catch {
      status = 'failed';
      error = guildClaimCopy.failed;
    }
  }

  // One load on mount, the same `$effect` (rather than `onMount`) Account.svelte uses so
  // this works identically hydrated by Astro or mounted by hand.
  $effect(() => {
    void load();
  });

  const signedIn = $derived(me !== null);
  const myBattletag = $derived(me?.user.battletag ?? null);

  /**
   * Same identity rule as Guild.svelte's own membership match (region + ruleset +
   * slugified name, not region + ruleset alone, which many guilds share): the viewer's
   * rank in THIS guild, read off `me.guilds`. Claim is an officer/leader-rank action
   * (spec section 4.4); without this, any signed-in account -- a stranger, a plain
   * member -- saw a live Claim/Confirm button that only failed once clicked (a 403 from
   * the API), and `guildClaimCopy.notEligible` existed but was never shown.
   */
  const membership = $derived(
    me?.guilds.find(
      (g) => g.region === path.region && g.ruleset === path.ruleset && characterSlug(g.name) === path.slug,
    ) ?? null,
  );
  const eligible = $derived(
    membership !== null && (membership.rank === 'officer' || membership.rank === 'leader'),
  );

  async function run(action: () => Promise<void>): Promise<void> {
    busy = true;
    error = '';
    try {
      await action();
    } catch (thrown) {
      error = thrown instanceof Error ? thrown.message : guildClaimCopy.failed;
    } finally {
      busy = false;
    }
  }

  const onClaim = (): void =>
    void run(async () => {
      if (guildId === null) return;
      const result = await claimGuild(guildId);
      justClaimedExpiry = result.expires_at;
      settings = await fetchGuildSettings(guildId);
    });

  const onConfirm = (): void =>
    void run(async () => {
      if (guildId === null) return;
      await confirmClaim(guildId);
      settings = await fetchGuildSettings(guildId);
    });

  const onRelease = (): void =>
    void run(async () => {
      if (guildId === null) return;
      await releaseClaim(guildId);
      settings = await fetchGuildSettings(guildId);
    });
</script>

<div class="flex flex-col gap-4" data-testid="guild-claim">
  <h1 class="section-title text-[18px]">{guildClaimCopy.heading(guildName)}</h1>
  {#if status === 'loading'}
    <p class="text-muted text-[14px]">{guildClaimCopy.loading}</p>
  {:else if status === 'failed'}
    <p class="text-[14px]" role="alert" data-testid="guild-claim-error">{error}</p>
  {:else if settings !== null}
    {#if settings.claimed_by !== null}
      <p class="text-[14px]" data-testid="guild-claim-state">
        {settings.claimed_by.battletag === myBattletag
          ? guildClaimCopy.claimedByYou
          : guildClaimCopy.claimedBySomeoneElse(settings.claimed_by.battletag)}
      </p>
      {#if settings.claimed_by.battletag === myBattletag}
        <button
          class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-4"
          onclick={onRelease}
          disabled={busy}
          data-testid="guild-claim-release"
        >
          {guildClaimCopy.releaseButton}
        </button>
      {/if}
    {:else if settings.claim_pending}
      <p class="text-[14px]" data-testid="guild-claim-state">{guildClaimCopy.pending(justClaimedExpiry)}</p>
      {#if signedIn && eligible}
        <button
          class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4"
          onclick={onConfirm}
          disabled={busy}
          data-testid="guild-claim-confirm"
        >
          {guildClaimCopy.confirmButton}
        </button>
      {:else if signedIn}
        <p class="text-[14px]" data-testid="guild-claim-not-eligible">{guildClaimCopy.notEligible}</p>
      {/if}
    {:else}
      <p class="text-[14px]" data-testid="guild-claim-state">{guildClaimCopy.unclaimed}</p>
      {#if !signedIn}
        <SignInPrompt line={guildClaimCopy.signInLine} testid="guild-claim-signin" />
      {:else if eligible}
        <button
          class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4"
          onclick={onClaim}
          disabled={busy}
          data-testid="guild-claim-button"
        >
          {guildClaimCopy.claimButton}
        </button>
      {:else}
        <p class="text-[14px]" data-testid="guild-claim-not-eligible">{guildClaimCopy.notEligible}</p>
      {/if}
    {/if}
    <div class="min-h-[21px]">
      {#if error !== ''}<p class="text-[14px]" role="alert" data-testid="guild-claim-action-error">
          {error}
        </p>{/if}
    </div>
  {/if}
</div>
