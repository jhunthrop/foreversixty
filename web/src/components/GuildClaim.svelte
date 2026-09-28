<!-- web/src/components/GuildClaim.svelte -->
<!-- /guild/<region>/<ruleset>/<name>/claim (spec section 4.4): current claim state, a
     claim button for an officer/leader-rank viewer, a confirm button for a second officer
     when one is pending, a release button for the current claimant. Reads the guild's
     public page for its name/id and its settings for claimed_by/claim_pending -- the spec
     defines no separate "claim status" read, and GuildSettingsData already carries every
     field this view needs (Task 1's api.ts). GET .../settings is verified-officer/leader-
     or-moderator only (spec section 2.6), so for everyone else the claim STATE comes from
     the member home (GET .../home, any member) -- lib/guild/claim-view.ts is the one rule
     -- and the page always says who can act and what to do next. It used to render only
     its heading to a plain member (found live, 2026-09-28). -->
<script lang="ts">
  import { fetchMeOnce, type Me } from '../lib/account/api';
  import { characterSlug, guildHref, type CharacterPath } from '../lib/characters';
  import { SETUP_NAV_ITEM } from '../lib/nav';
  import {
    claimGuild,
    confirmClaim,
    contestClaim,
    fetchGuildHome,
    fetchGuildSettings,
    releaseClaim,
    type GuildHome,
    type GuildSettingsData,
  } from '../lib/guild/api';
  import { claimView } from '../lib/guild/claim-view';
  import { guildClaimCopy, guildHomeCopy } from '../lib/guild/copy';
  import { GUILD_LOADING } from '../lib/guild/layout';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { fetchGuild } from '../lib/rankings/api';
  import GuildStatus from './GuildStatus.svelte';
  import SignInPrompt from './SignInPrompt.svelte';
  import { BUSY_CLASS } from '../lib/ui/busy';

  let { path }: { path: CharacterPath } = $props();

  let me = $state<Me | null>(null);
  let guildId = $state<number | null>(null);
  let guildName = $state('');
  let settings = $state<GuildSettingsData | null>(null);
  let home = $state<GuildHome | null>(null);
  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let error = $state('');
  let busy = $state(false);
  let showContestConfirm = $state(false);

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
      // Settings answers an officer; the home answers any member. Each is refused with a
      // 403 for everyone else, which is a state here, not a failure.
      const id = guildPage.guild.id;
      const isMember = membershipIn(session) !== null;
      [settings, home] = await Promise.all([
        fetchGuildSettings(id).catch(() => null),
        isMember ? fetchGuildHome(id).catch(() => null) : Promise.resolve(null),
      ]);
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

  const myBattletag = $derived(me?.user.battletag ?? null);

  /**
   * Same identity rule as Guild.svelte's own membership match (region + ruleset +
   * slugified name, not region + ruleset alone, which many guilds share): the viewer's
   * rank in THIS guild, read off `me.guilds`. Claim is an officer/leader-rank action
   * (spec section 4.4); without this, any signed-in account -- a stranger, a plain
   * member -- saw a live Claim/Confirm button that only failed once clicked (a 403 from
   * the API), and `guildClaimCopy.notEligible` existed but was never shown.
   */
  function membershipIn(session: Me | null) {
    return (
      session?.guilds.find(
        (g) => g.region === path.region && g.ruleset === path.ruleset && characterSlug(g.name) === path.slug,
      ) ?? null
    );
  }
  const membership = $derived(membershipIn(me));
  const view = $derived(claimView({ me, membership, settings, home }));
  const eligible = $derived(view.viewer === 'eligible');
  const claimState = $derived(view.claim?.state ?? null);

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
      await claimGuild(guildId);
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

  const onContest = (): void =>
    void run(async () => {
      if (guildId === null) return;
      await contestClaim(guildId);
      showContestConfirm = false;
      settings = await fetchGuildSettings(guildId);
    });
</script>

<div class="reveal flex flex-col gap-4" data-testid="guild-claim">
  <a class="text-nav w-fit text-[13px]" href={guildHref(path.region, path.ruleset, guildName || path.slug)}>
    ← {guildClaimCopy.backToGuild(guildName)}
  </a>
  <h1 class="section-title text-[18px]">{guildClaimCopy.heading(guildName)}</h1>
  {#if status === 'loading' || status === 'failed'}
    <GuildStatus
      status={status === 'loading' ? 'loading' : 'failed'}
      {error}
      onRetry={() => void load()}
      lines={GUILD_LOADING.claim.lines}
      minHeight={GUILD_LOADING.claim.minHeight}
      testid="guild-claim"
    />
  {:else}
    {#snippet contestOffer()}
      {#if membership !== null && claimState !== null && claimState !== 'contested'}
        {#if !showContestConfirm}
          <button
            class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-3 ${busy ? BUSY_CLASS : ''}`}
            onclick={() => (showContestConfirm = true)}
            disabled={busy}
            aria-busy={busy}
            data-testid="guild-claim-contest-button"
          >
            {guildHomeCopy.contestButton}
          </button>
        {:else}
          <div
            class="border-line-soft flex flex-col gap-3 border p-4"
            data-testid="guild-claim-contest-confirm"
          >
            <ul class="flex flex-col gap-1 text-[13px]">
              {#each guildHomeCopy.contestRules as rule (rule)}
                <li>{rule}</li>
              {/each}
            </ul>
            <div class="flex gap-3">
              <button
                class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-3 ${busy ? BUSY_CLASS : ''}`}
                onclick={onContest}
                disabled={busy}
                aria-busy={busy}
                data-testid="guild-claim-contest-confirm-button"
              >
                {guildHomeCopy.contestConfirmButton}
              </button>
              <button
                class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-3 ${busy ? BUSY_CLASS : ''}`}
                onclick={() => {
                  showContestConfirm = false;
                  error = '';
                }}
                disabled={busy}
                aria-busy={busy}
              >
                {guildHomeCopy.cancel}
              </button>
            </div>
          </div>
        {/if}
      {/if}
    {/snippet}

    {#if claimState === 'contested'}
      <p class="text-[14px]" data-testid="guild-claim-contested-notice">{guildClaimCopy.contested}</p>
    {/if}

    {#if settings !== null && settings.claimed_by !== null}
      <!-- An officer's read: who holds the claim, and release for the holder. -->
      <p class="text-[14px]" data-testid="guild-claim-state">
        {settings.claimed_by.battletag === myBattletag
          ? guildClaimCopy.claimedByYou
          : guildClaimCopy.claimedBySomeoneElse(settings.claimed_by.battletag)}
      </p>
      {#if settings.claimed_by.battletag === myBattletag}
        <button
          class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-4 ${busy ? BUSY_CLASS : ''}`}
          onclick={onRelease}
          disabled={busy}
          aria-busy={busy}
          data-testid="guild-claim-release"
        >
          {guildClaimCopy.releaseButton}
        </button>
      {/if}
      {@render contestOffer()}
    {:else if settings !== null && settings.claim_pending !== null}
      <p class="text-[14px]" data-testid="guild-claim-state">
        {guildClaimCopy.pending(settings.claim_pending.expires_at)}
      </p>
      {#if eligible}
        <button
          class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4 ${busy ? BUSY_CLASS : ''}`}
          onclick={onConfirm}
          disabled={busy}
          aria-busy={busy}
          data-testid="guild-claim-confirm"
        >
          {guildClaimCopy.confirmButton}
        </button>
      {:else}
        <p class="text-[14px]" data-testid="guild-claim-not-eligible">{guildClaimCopy.notEligible}</p>
      {/if}
      {@render contestOffer()}
    {:else if view.viewer === 'signed-out'}
      <SignInPrompt line={guildClaimCopy.signInLine} testid="guild-claim-signin" />
    {:else if view.viewer === 'stranger'}
      <p class="text-[14px]" data-testid="guild-claim-not-member">
        {guildClaimCopy.notAMember}
        <a class="text-nav" href={SETUP_NAV_ITEM.href}>{SETUP_NAV_ITEM.label}</a>
      </p>
    {:else}
      <!-- A member's read (the home's state, never who) or an officer of an unclaimed guild. -->
      {#if claimState === 'claimed'}
        <p class="text-[14px]" data-testid="guild-claim-state">{guildClaimCopy.claimedByAnOfficer}</p>
      {:else if claimState === 'pending'}
        <p class="text-[14px]" data-testid="guild-claim-state">{guildClaimCopy.pending()}</p>
      {:else if claimState === 'unclaimed'}
        <p class="text-[14px]" data-testid="guild-claim-state">{guildClaimCopy.unclaimed}</p>
      {/if}
      {#if eligible && (claimState === 'unclaimed' || claimState === null)}
        <button
          class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4 ${busy ? BUSY_CLASS : ''}`}
          onclick={onClaim}
          disabled={busy}
          aria-busy={busy}
          data-testid="guild-claim-button"
        >
          {guildClaimCopy.claimButton}
        </button>
      {:else if !eligible}
        <p class="text-[14px]" data-testid="guild-claim-not-eligible">
          {guildClaimCopy.memberCannotClaim}
          <a class="text-nav" href={SETUP_NAV_ITEM.href}>{SETUP_NAV_ITEM.label}</a>
        </p>
      {/if}
      {@render contestOffer()}
    {/if}

    <section class="flex flex-col gap-2" data-testid="guild-claim-rules">
      <h2 class="section-title text-[16px]">{guildClaimCopy.rulesHeading}</h2>
      <ul class="text-muted flex flex-col gap-1 text-[13px]">
        {#each guildClaimCopy.rules as rule (rule)}
          <li>{rule}</li>
        {/each}
      </ul>
    </section>
    <div class="min-h-[21px]">
      {#if error !== ''}<p class="text-[14px]" role="alert" data-testid="guild-claim-action-error">
          {error}
        </p>{/if}
    </div>
  {/if}
</div>
