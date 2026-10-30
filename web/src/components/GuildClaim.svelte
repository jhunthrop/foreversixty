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
  import { GUILD_BANDLESS_TOP } from '../lib/guild/raid-layout';
  import { fetchMeFresh, type Me } from '../lib/account/api';
  import { guildHref, guildSettingsHref, type CharacterPath } from '../lib/characters';
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
  import {
    claimLinkFor,
    claimPlanLine,
    parseClaimPlan,
    setupHrefFor,
    subscribeLabel,
    type ClaimPlan,
  } from '../lib/guild/claim-plan';
  import { founderOpen } from '../lib/billing/plans';
  import ClaimMemberHandoff from './claim/ClaimMemberHandoff.svelte';
  import ClaimPlanLine from './claim/ClaimPlanLine.svelte';
  import ClaimNextSteps from './claim/ClaimNextSteps.svelte';
  import { findMyGuild } from '../lib/guild/my-guild';
  import { guildClaimCopy, guildHomeCopy } from '../lib/guild/copy';
  import { formatMadeOn } from '../lib/guild/invite-link';
  import { GUILD_LOADING } from '../lib/guild/layout';
  import { button } from '../lib/ui/button';
  import Card from './ui/Card.svelte';
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
  let showClaimConfirm = $state(false);
  let showReleaseConfirm = $state(false);
  /** The plan this claim link was made for, and the page's own path and query (read on load, client side). */
  let claimPlan = $state<ClaimPlan | null>(null);
  let claimPath = $state('');
  let claimOrigin = $state('');

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
    claimPlan = parseClaimPlan(window.location.search);
    claimPath = `${window.location.pathname}${window.location.search}`;
    claimOrigin = window.location.origin;
    try {
      const [session, guildPage] = await Promise.all([
        fetchMeFresh((fresh) => (me = fresh)),
        fetchGuild(path),
      ]);
      me = session;
      guildId = guildPage.guild.id;
      guildName = guildPage.guild.name;
      // Settings answers an officer; the home answers any member. Each is refused with a
      // 403 for everyone else, which is a state here, not a failure.
      const id = guildPage.guild.id;
      const isMember = findMyGuild(session, path) !== null;
      [settings, home] = await Promise.all([
        fetchGuildSettings(id, undefined, (fresh) => (settings = fresh)).catch(() => null),
        isMember
          ? fetchGuildHome(id, undefined, undefined, (fresh) => (home = fresh)).catch(() => null)
          : Promise.resolve(null),
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
  const membership = $derived(findMyGuild(me, path));
  const view = $derived(claimView({ me, membership, settings, home }));
  const eligible = $derived(view.viewer === 'eligible');
  const claimState = $derived(view.claim?.state ?? null);
  /** The signed-in account is the one holding the claim: the page is then about the claim it holds. */
  const holdsClaim = $derived(
    settings !== null && settings.claimed_by !== null && settings.claimed_by.battletag === myBattletag,
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
      await claimGuild(guildId);
      showClaimConfirm = false;
      settings = await fetchGuildSettings(guildId);
      me = await fetchMeFresh();
    });

  const onConfirm = (): void =>
    void run(async () => {
      if (guildId === null) return;
      await confirmClaim(guildId);
      settings = await fetchGuildSettings(guildId);
      me = await fetchMeFresh();
    });

  /** The `ready` button's label once the viewer holds the claim of a guild the link named a plan for. */
  const subscribe = $derived.by(() => {
    const subscribeFacts = membership?.subscribe;
    if (claimPlan === null || guildId === null || subscribeFacts?.step !== 'ready') return null;
    return {
      id: guildId,
      label: subscribeLabel(claimPlan, subscribeFacts.trial_used, founderOpen(new Date())),
    };
  });

  const onRelease = (): void =>
    void run(async () => {
      if (guildId === null) return;
      await releaseClaim(guildId);
      showReleaseConfirm = false;
      settings = await fetchGuildSettings(guildId);
    });

  // The guild's own page and settings, for the next-step links once the viewer holds it.
  const guildBase = $derived(guildHref(path.region, path.ruleset, guildName || path.slug));
  const settingsBase = $derived(guildSettingsHref(path.region, path.ruleset, guildName || path.slug));

  const onContest = (): void =>
    void run(async () => {
      if (guildId === null) return;
      await contestClaim(guildId);
      showContestConfirm = false;
      settings = await fetchGuildSettings(guildId);
    });
</script>

<div class="reveal {GUILD_BANDLESS_TOP} flex flex-col gap-[22px] md:gap-8" data-testid="guild-claim">
  <a class="text-nav w-fit text-[13px]" href={guildHref(path.region, path.ruleset, guildName || path.slug)}>
    ← {guildClaimCopy.backToGuild(guildName)}
  </a>
  <h1 class="heading-page" data-testid="guild-claim-heading">
    {holdsClaim ? guildClaimCopy.headingHeld(guildName) : guildClaimCopy.heading(guildName)}
  </h1>
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
    <Card glyph="guild" title={guildClaimCopy.stateHeading} testid="guild-claim-state-card">
      {#if claimPlan !== null && !holdsClaim}
        <ClaimPlanLine line={claimPlanLine(guildName, claimPlan, founderOpen(new Date()))} />
      {/if}
      {#snippet contestOffer()}
        {#if membership !== null && claimState === 'claimed'}
          {#if !showContestConfirm}
            <button
              class={`${button('secondary')} w-fit ${busy ? BUSY_CLASS : ''}`}
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
                  class={`${button('secondary')} w-fit ${busy ? BUSY_CLASS : ''}`}
                  onclick={onContest}
                  disabled={busy}
                  aria-busy={busy}
                  data-testid="guild-claim-contest-confirm-button"
                >
                  {guildHomeCopy.contestConfirmButton}
                </button>
                <button
                  class={`${button('secondary')} w-fit ${busy ? BUSY_CLASS : ''}`}
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
        {#if settings.claimed_by.battletag === myBattletag}
          <p class="text-[16px]" data-testid="guild-claim-state">{guildClaimCopy.youHold(guildName)}</p>
          {#if claimState !== 'contested'}
            <ClaimNextSteps {subscribe} {claimPlan} {guildBase} {settingsBase} />
          {/if}
          {#if !showReleaseConfirm}
            <button
              class={`text-muted mt-2 w-fit text-[13px] underline ${busy ? BUSY_CLASS : ''}`}
              onclick={() => (showReleaseConfirm = true)}
              disabled={busy}
              aria-busy={busy}
              data-testid="guild-claim-release"
            >
              {guildClaimCopy.releaseButton}
            </button>
          {:else}
            <div
              class="border-line-soft flex flex-col gap-3 border p-4"
              data-testid="guild-claim-release-confirm"
            >
              <p class="text-[13px]">{guildClaimCopy.releaseConfirmLine(guildName)}</p>
              <div class="flex gap-3">
                <button
                  class={`${button('secondary')} w-fit ${busy ? BUSY_CLASS : ''}`}
                  onclick={onRelease}
                  disabled={busy}
                  aria-busy={busy}
                  data-testid="guild-claim-release-confirm-button"
                >
                  {guildClaimCopy.releaseConfirmButton}
                </button>
                <button
                  class={`${button('ghost')} w-fit`}
                  onclick={() => (showReleaseConfirm = false)}
                  disabled={busy}
                >
                  {guildClaimCopy.cancel}
                </button>
              </div>
            </div>
          {/if}
        {:else}
          <p class="text-[14px]" data-testid="guild-claim-state">
            {guildClaimCopy.claimedBySomeoneElse(
              settings.claimed_by.battletag,
              settings.claim.since === undefined ? undefined : formatMadeOn(settings.claim.since),
            )}
          </p>
          {@render contestOffer()}
        {/if}
      {:else if settings !== null && settings.claim_pending !== null}
        <p class="text-[14px]" data-testid="guild-claim-state">
          {guildClaimCopy.pending(settings.claim_pending.expires_at)}
        </p>
        {#if eligible}
          <button
            class={`${button('secondary')} w-fit ${busy ? BUSY_CLASS : ''}`}
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
      {:else if view.viewer === 'signed-out'}
        <SignInPrompt line={guildClaimCopy.signInLine} next={claimPath} testid="guild-claim-signin" />
      {:else if view.viewer === 'stranger'}
        <p class="text-[14px]" data-testid="guild-claim-not-member">
          {guildClaimCopy.notAMember}
          <a class="text-nav" href={setupHrefFor(SETUP_NAV_ITEM.href, claimPath)}>{SETUP_NAV_ITEM.label}</a>
        </p>
      {:else}
        <!-- A member's read (the home's state, never who) or an officer of an unclaimed guild. -->
        {#if claimState === 'claimed'}
          <p class="text-[14px]" data-testid="guild-claim-state">{guildClaimCopy.claimedByAnOfficer}</p>
        {:else if claimState === 'pending'}
          <p class="text-[14px]" data-testid="guild-claim-state">{guildClaimCopy.pending()}</p>
        {:else if claimState === 'unclaimed'}
          <p class="text-[14px]" data-testid="guild-claim-state">
            {eligible ? guildClaimCopy.unclaimedForOfficer : guildClaimCopy.unclaimed}
          </p>
        {/if}
        {#if eligible && (claimState === 'unclaimed' || claimState === null)}
          {#if !showClaimConfirm}
            <button
              class={`${button('primary')} w-fit ${busy ? BUSY_CLASS : ''}`}
              onclick={() => (showClaimConfirm = true)}
              disabled={busy}
              aria-busy={busy}
              data-testid="guild-claim-button"
            >
              {guildClaimCopy.claimButton}
            </button>
          {:else}
            <div
              class="border-line-soft flex flex-col gap-3 border p-4"
              data-testid="guild-claim-confirm-box"
            >
              <p class="text-[13px]">{guildClaimCopy.claimConfirmLine(guildName)}</p>
              <div class="flex gap-3">
                <button
                  class={`${button('secondary')} w-fit ${busy ? BUSY_CLASS : ''}`}
                  onclick={onClaim}
                  disabled={busy}
                  aria-busy={busy}
                  data-testid="guild-claim-confirm-claim"
                >
                  {guildClaimCopy.claimConfirmButton}
                </button>
                <button
                  class={`${button('ghost')} w-fit`}
                  onclick={() => (showClaimConfirm = false)}
                  disabled={busy}
                >
                  {guildClaimCopy.cancel}
                </button>
              </div>
            </div>
          {/if}
        {:else if !eligible}
          <ClaimMemberHandoff
            {guildName}
            url={claimLinkFor(path, claimPlan, claimOrigin)}
            tier={claimPlan?.plan}
          />
        {/if}
        {@render contestOffer()}
      {/if}
    </Card>

    <Card glyph="book" title={guildClaimCopy.rulesHeading} testid="guild-claim-rules">
      <ul class="text-muted flex flex-col gap-1 text-[13px]">
        {#each guildClaimCopy.rules as rule (rule)}
          <li>{rule}</li>
        {/each}
      </ul>
    </Card>
    <div class="min-h-[21px]">
      {#if error !== ''}<p class="text-[14px]" role="alert" data-testid="guild-claim-action-error">
          {error}
        </p>{/if}
    </div>
  {/if}
</div>
