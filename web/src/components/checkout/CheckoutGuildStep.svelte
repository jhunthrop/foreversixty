<!-- web/src/components/checkout/CheckoutGuildStep.svelte -->
<!-- Spec 2026-10-27-subscribe-and-join §5.1 body and §5.7: the guild step under the band. The
     centred 880 column (gutters 48 / 18, 32 / 20 px top), the lead sentence, one card per guild
     in orderGuildCards order, and the foot. `me` null is loading: the lead slot and Skeleton
     cards hold the ready height, their count from the session's fs.meGuildCount hint. `failed`
     swaps the body for LoadError; after a retry succeeds, focus moves to the lead. Empty A (no
     characters) and Empty B (no guild) draw the set-up panel, the page's one fill.

     Each card is a GuildStepCard around a GuildStepBlock; an older /v1/me (no `subscribe`) gets
     identity rows and `Open the guild` only (§8). Rows replace the skeleton once /v1/me answered
     and every guild crest decoded or failed, capped at CREST_SETTLE_MS. The `ready` card calls
     `onSwitch` for `Switch to Guild` / `Switch to Raid guild`. -->
<script lang="ts">
  import { tick, untrack } from 'svelte';
  import type { SubscribeStep } from '../../lib/billing/contract';
  import type { Me, MeGuild } from '../../lib/account/api';
  import { guildHref } from '../../lib/characters';
  import { preloadImages } from '../../lib/ui/preload-images';
  import { CARD } from '../../lib/ui/card';
  import { premiumSource } from '../../lib/account/plan-line';
  import { readMeGuildCount } from '../../lib/account/premium-hint';
  import { claimGuild, confirmClaim } from '../../lib/guild/api';
  import {
    claimKindOf,
    claimLock,
    refusalLine,
    rereadVerdict,
    settleStaleRefusal,
    settleSuccess,
    type CardNotice,
    type ClaimIo,
    type ClaimKind,
    type ClaimSettlement,
    type RereadVerdict,
  } from '../../lib/billing/claim-flow';
  import {
    readyCount,
    restoredFromCache,
    type Arrival,
    type CardOutcome,
    type PayBusy,
  } from '../../lib/billing/card-checkout';
  import { checkoutGuildCopy as copy, type GuildTier } from '../../lib/billing/checkout-guild-copy';
  import {
    GUILD_LEAD_RESERVE,
    GUILD_SKELETON_CARD,
    releaseGuildFrameReserve,
  } from '../../lib/billing/checkout-layout';
  import { priceFor } from '../../lib/billing/plan-costs';
  import { formatPrice, type Interval } from '../../lib/billing/plans';
  import { leadSentence, orderGuildCards, type StepGuild } from '../../lib/billing/subscribe-step';
  import { INLINE_LINK } from '../../lib/guild/field-styles';
  import { PRIMARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { checkoutHref } from '../premium/model';
  import LoadError from '../ui/LoadError.svelte';
  import GuildPayAction from './GuildPayAction.svelte';
  import GuildStepBlock from './GuildStepBlock.svelte';
  import GuildStepCard from './GuildStepCard.svelte';

  let {
    me,
    failed,
    plan,
    interval,
    founder,
    selfUrl,
    onRetry,
    onCheckAgain,
    onSwitch,
    onCheckout,
    onPortal,
    arrival = undefined,
  }: {
    /** Null while /v1/me is in flight. */
    me: Me | null;
    failed: boolean;
    plan: GuildTier;
    interval: Interval;
    founder: boolean;
    /** This checkout's own path and query: the `next` of every Get set up link. */
    selfUrl: string;
    onRetry: () => void;
    /** Re-reads /v1/me without the skeleton; the `checkout_open` card's `Check again`. */
    onCheckAgain: () => Promise<boolean>;
    /** Rewrites the page's plan in place (spec §5.3); the `ready` card's switch note calls it. */
    onSwitch: (to: GuildTier) => void;
    /** Starts the guild's Stripe checkout; the launcher follows a `redirect` outcome. */
    onCheckout: (guild: MeGuild) => Promise<CardOutcome>;
    /** Opens the guild's billing portal; the launcher follows a `redirect` outcome. */
    onPortal: (guild: MeGuild) => Promise<CardOutcome>;
    /** A `guild_id` arrival: the card that goes first, or the ready guild that launches (§5.6). */
    arrival?: Arrival;
  } = $props();

  /** The longest the rows wait behind the skeleton for the guild crests (spec §5.7). */
  const CREST_SETTLE_MS = 1000;

  let leadEl = $state<HTMLParagraphElement | undefined>(undefined);
  /** Plain, not reactive: set while the error shows, so the next ready render moves focus once. */
  let retrying = false;
  /** The crest set whose decode (or failure) the reveal has waited for. */
  let settledCrests = $state<string | null>(null);

  // The step rules key on the guild page slug; an older API sends none, and only the
  // hand-off link (Task 16) reads it, never the lead or the order.
  const guilds = $derived<(MeGuild & StepGuild)[]>(
    (me?.guilds ?? []).map((guild) => ({ ...guild, slug: guild.slug ?? '' })),
  );
  const cards = $derived(orderGuildCards(guilds, arrival?.firstId));
  const readyCards = $derived(readyCount(guilds));
  const lead = $derived(leadSentence(guilds));
  const crestUrls = $derived(guilds.flatMap((guild) => (guild.crest_url ? [guild.crest_url] : [])));
  const crestKey = $derived(crestUrls.join('|'));
  // Without Image (the static render) there is nothing to wait for.
  const crestsReady = $derived(
    crestUrls.length === 0 || settledCrests === crestKey || typeof Image === 'undefined',
  );
  const viewerIsOfficer = (guild: MeGuild): boolean => guild.rank === 'officer' || guild.rank === 'leader';
  const skeletonCards = $derived(readMeGuildCount());
  const setupHref = $derived(`/setup?next=${encodeURIComponent(selfUrl)}`);
  const empty = $derived(
    me === null || me.guilds.length > 0 ? null : me.characters.length === 0 ? copy.emptyA : copy.emptyB,
  );
  const showsPremiumLine = $derived(me !== null && premiumSource(me, new Date()) === 'none');
  const premiumPrice = $derived(formatPrice(priceFor('premium', interval, founder), interval));

  /** The skeleton shows: /v1/me is in flight, or the rows wait for their crests. */
  const loading = $derived(!failed && (me === null || (empty === null && !crestsReady)));

  // Spec §7 (G2): the page's pre-paint height reserve holds only while the skeleton shows.
  $effect(() => {
    if (!loading) releaseGuildFrameReserve(document.documentElement);
  });

  $effect(() => {
    if (me === null || failed || crestsReady) return;
    const key = crestKey;
    let cancelled = false;
    void preloadImages(crestUrls, CREST_SETTLE_MS).then(() => {
      if (!cancelled) settledCrests = key;
    });
    return () => {
      cancelled = true;
    };
  });

  $effect(() => {
    if (failed) {
      retrying = true;
      return;
    }
    if (!retrying || me === null || leadEl === undefined) return;
    retrying = false;
    leadEl.focus();
  });

  const CLAIM_IO: ClaimIo = { claim: claimGuild, confirm: confirmClaim };
  const STEP_ACTION_SELECTOR = '[data-step-action]';

  /** The one open confirm (all cards), its refusal line, and the notice a finished claim left. */
  let openClaim = $state<{ id: number; kind: ClaimKind; error: string | null } | null>(null);
  let notice = $state<{ id: number; value: CardNotice } | null>(null);
  /** The card whose /v1/me re-read is in flight: its opener is disabled and aria-busy meanwhile. */
  let rereadingId = $state<number | null>(null);

  const stepOfGuild = (id: number): SubscribeStep | null =>
    guilds.find((guild) => guild.id === id)?.subscribe?.step ?? null;

  function focusSoon(find: () => HTMLElement | null): void {
    void tick().then(() => find()?.focus());
  }

  const focusNewAction = (id: number): void =>
    focusSoon(
      () =>
        document.querySelector<HTMLElement>(`#guild-${id} ${STEP_ACTION_SELECTOR}`) ??
        document.getElementById(`guild-${id}-notice`),
    );

  /**
   * Re-reads /v1/me (the launcher resolves once its read settled) and says what the card did.
   * The shell owns the focus that follows, so the launcher moves none.
   */
  async function rereadCard(id: number): Promise<RereadVerdict> {
    const before = stepOfGuild(id);
    rereadingId = id;
    const read = await onCheckAgain();
    await tick();
    rereadingId = null;
    return rereadVerdict(read, before, stepOfGuild(id));
  }

  function applySettlement(id: number, settlement: ClaimSettlement): void {
    if (settlement.closeBlock) openClaim = null;
    else if (openClaim !== null) openClaim = { ...openClaim, error: settlement.blockError };
    if (settlement.notice !== null) notice = { id, value: settlement.notice };
    if (settlement.focusNewAction) focusNewAction(id);
    else if (settlement.closeBlock) focusSoon(() => document.getElementById(`guild-${id}-claim-open`));
  }

  function closeClaim(id: number): void {
    openClaim = null;
    focusSoon(() => document.getElementById(`guild-${id}-claim-open`));
  }

  async function onClaimDone(id: number, line: string): Promise<void> {
    applySettlement(id, settleSuccess(await rereadCard(id), line));
  }

  async function onClaimRefused(
    guild: MeGuild,
    kind: ClaimKind,
    code: string,
    stale: boolean,
  ): Promise<void> {
    if (!stale) {
      openClaim = { id: guild.id, kind, error: refusalLine(kind, code, null, guild.name) };
      return;
    }
    const verdict = await rereadCard(guild.id);
    applySettlement(guild.id, settleStaleRefusal(verdict, kind, code, stepOfGuild(guild.id), guild.name));
  }

  function claimControls(guild: MeGuild, kind: ClaimKind | null) {
    if (kind === null) return undefined;
    const lock = claimLock(openClaim?.id ?? null, rereadingId, guild.id);
    return {
      ...lock,
      error: lock.open ? (openClaim?.error ?? null) : null,
      io: CLAIM_IO,
      onOpen: () => {
        notice = null;
        openClaim = { id: guild.id, kind, error: null };
      },
      onCancel: () => closeClaim(guild.id),
      onDone: (line: string) => void onClaimDone(guild.id, line),
      onRefused: (code: string, stale: boolean) => void onClaimRefused(guild, kind, code, stale),
    };
  }

  /** The card whose checkout or portal call is in flight, and the card whose failure offers Try again. */
  let busy = $state<{ id: number; kind: Exclude<PayBusy, null> } | null>(null);
  let retryId = $state<number | null>(null);

  /** Spec §9.1: a refusal stands on its card; a 409 or 403 also re-reads and moves focus when the step changes. */
  function onPayRefused(guild: MeGuild, outcome: Extract<CardOutcome, { kind: 'refused' }>): void {
    busy = null;
    retryId = outcome.retry ? guild.id : null;
    if (outcome.line !== null) notice = { id: guild.id, value: { text: outcome.line, role: 'alert' } };
    if (!outcome.rereads) {
      focusSoon(() => document.getElementById(`guild-${guild.id}-notice`));
      return;
    }
    void rereadCard(guild.id).then(() => focusNewAction(guild.id));
  }

  async function pay(guild: MeGuild, kind: Exclude<PayBusy, null>): Promise<void> {
    notice = null;
    retryId = null;
    busy = { id: guild.id, kind };
    const outcome = await (kind === 'checkout' ? onCheckout(guild) : onPortal(guild));
    // A redirect leaves the card busy until the page unloads (or is restored: onPageShow).
    if (outcome.kind === 'refused') onPayRefused(guild, outcome);
  }

  /** Back from Stripe via the bfcache: the button is usable again. `launched` stays set, so the arrival does not bounce the player straight back. */
  function onPageShow(event: PageTransitionEvent): void {
    if (restoredFromCache(event)) busy = null;
  }

  let launched = false;
  $effect(() => {
    const launchId = arrival?.launchId;
    if (launched || launchId === undefined || me === null || failed) return;
    const guild = guilds.find((entry) => entry.id === launchId);
    if (guild === undefined) return;
    launched = true;
    untrack(() => void pay(guild, 'checkout'));
  });

  let scrolled = false;
  $effect(() => {
    const firstId = arrival?.firstId;
    if (scrolled || firstId === undefined || me === null || failed || empty !== null || !crestsReady) return;
    scrolled = true;
    void tick().then(() => document.getElementById(`guild-${firstId}`)?.scrollIntoView({ block: 'start' }));
  });

  function guildPageHref(guild: MeGuild): string {
    if (guild.slug) return `/guild/${guild.region}/${guild.ruleset}/${encodeURIComponent(guild.slug)}`;
    return guildHref(guild.region, guild.ruleset, guild.name);
  }
</script>

<svelte:window onpageshow={onPageShow} />

<div class="px-[18px] pt-5 pb-16 md:px-12 md:pt-8" data-testid="checkout-guild-step">
  <!-- .measure-wide: checkout.astro sets it to this page's 880 px column. -->
  <div class="measure-wide flex flex-col">
    {#if failed}
      <LoadError message={copy.loadError.title} {onRetry} testid="checkout-guild-error" />
    {:else if loading}
      <p class={`m-0 mb-4 ${GUILD_LEAD_RESERVE}`} aria-hidden="true"></p>
      <ul class="m-0 flex list-none flex-col gap-3 p-0" aria-busy="true" data-testid="checkout-guild-loading">
        <li class="sr-only" role="status">{copy.loading}</li>
        {#each Array.from({ length: skeletonCards }, (_, index) => index) as index (index)}
          <li class={`skeleton-block rounded-panel ${GUILD_SKELETON_CARD}`} aria-hidden="true"></li>
        {/each}
      </ul>
    {:else if empty !== null}
      <!-- Left in the 880 column, where the foot sits too (row 11); full-width button on a phone. -->
      <div
        class={`${CARD} flex max-w-[640px] flex-col gap-2 p-5`}
        data-testid={empty === copy.emptyA ? 'checkout-guild-empty-a' : 'checkout-guild-empty-b'}
      >
        <p class="text-strong m-0 text-[15px]">{empty.title}</p>
        <p class="text-text m-0 text-[14px]">{empty.body}</p>
        <a href={setupHref} class={`${PRIMARY_BUTTON_FIXED} mt-2 w-full justify-center px-4 md:w-fit`}
          >{empty.button}</a
        >
      </div>
    {:else}
      <p
        bind:this={leadEl}
        tabindex="-1"
        class="text-text m-0 mb-4 max-w-[720px] text-[14px] outline-none md:text-[15px]"
        data-testid="checkout-guild-lead"
      >
        {lead}
      </p>
      <ul class="m-0 flex list-none flex-col gap-3 p-0" data-testid="checkout-guild-cards">
        {#each cards as guild (guild.id)}
          <GuildStepCard
            {guild}
            characters={me?.characters ?? []}
            mainKey={me?.main_character_key}
            standing={guild.subscribe != null}
          >
            {#snippet block()}
              <GuildStepBlock
                {guild}
                {plan}
                {interval}
                {selfUrl}
                viewerIsOfficer={viewerIsOfficer(guild)}
                guildPageHref={guildPageHref(guild)}
                {onCheckAgain}
                claim={claimControls(guild, claimKindOf(guild.subscribe?.step ?? null))}
                notice={notice?.id === guild.id ? notice.value : null}
                onTryAgain={retryId === guild.id ? () => void pay(guild, 'checkout') : undefined}
              >
                {#snippet action()}
                  <GuildPayAction
                    {guild}
                    {plan}
                    {interval}
                    {founder}
                    {readyCards}
                    busy={busy?.id === guild.id ? busy.kind : null}
                    locked={busy !== null && busy.id !== guild.id}
                    guildPageHref={guildPageHref(guild)}
                    onSubscribe={() => void pay(guild, 'checkout')}
                    onManage={() => void pay(guild, 'portal')}
                    {onSwitch}
                  />
                {/snippet}
              </GuildStepBlock>
            {/snippet}
          </GuildStepCard>
        {/each}
      </ul>
    {/if}

    <!-- An empty state's panel already offers Get set up, so the foot's missing-guild line would repeat it. -->
    {#if me !== null && !failed && (empty === null || showsPremiumLine)}
      <div class="text-muted mt-6 flex flex-col gap-1 text-[13px]" data-testid="checkout-guild-foot">
        {#if empty === null}
          <p class="m-0">
            {copy.foot.missingGuild}
            <a href={setupHref} class={INLINE_LINK}>{copy.foot.getSetUp}</a>
          </p>
        {/if}
        {#if showsPremiumLine}
          <p class="m-0">
            {copy.foot.premium(premiumPrice)}
            <a href={checkoutHref('premium', interval)} class={INLINE_LINK}>{copy.foot.getPremium}</a>
          </p>
        {/if}
      </div>
    {/if}
  </div>
</div>
