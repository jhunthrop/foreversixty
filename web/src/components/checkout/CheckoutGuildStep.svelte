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
    refusalLine,
    type CardNotice,
    type ClaimIo,
    type ClaimKind,
  } from '../../lib/billing/claim-flow';
  import { readyCount, type Arrival, type CardOutcome, type PayBusy } from '../../lib/billing/card-checkout';
  import { checkoutGuildCopy as copy, type GuildTier } from '../../lib/billing/checkout-guild-copy';
  import { GUILD_CARD_MIN, GUILD_LEAD_RESERVE } from '../../lib/billing/checkout-layout';
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
    onCheckAgain: () => void;
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
  /** Set by a claim that re-reads /v1/me: focus moves once that card's step changes. */
  let awaitedStep: { id: number; from: SubscribeStep | null } | null = null;

  const stepOfGuild = (id: number): SubscribeStep | null =>
    guilds.find((guild) => guild.id === id)?.subscribe?.step ?? null;

  function focusSoon(find: () => HTMLElement | null): void {
    void tick().then(() => find()?.focus());
  }

  function closeClaim(returnFocusTo: number | null): void {
    openClaim = null;
    if (returnFocusTo !== null) focusSoon(() => document.getElementById(`guild-${returnFocusTo}-claim-open`));
  }

  function onClaimDone(id: number, line: string): void {
    awaitedStep = { id, from: stepOfGuild(id) };
    notice = { id, value: { text: line, role: 'status' } };
    closeClaim(null);
    onCheckAgain();
  }

  function onClaimRefused(guild: MeGuild, kind: ClaimKind, code: string, rereads: boolean): void {
    openClaim = { id: guild.id, kind, error: refusalLine(kind, code, null, guild.name) };
    if (!rereads) return;
    awaitedStep = { id: guild.id, from: stepOfGuild(guild.id) };
    onCheckAgain();
  }

  $effect(() => {
    if (awaitedStep === null) return;
    const { id, from } = awaitedStep;
    const now = stepOfGuild(id);
    if (now === from) return;
    awaitedStep = null;
    // The card moved on: the block closes and a refusal stands as an alert.
    if (openClaim?.id === id) {
      const name = guilds.find((guild) => guild.id === id)?.name ?? '';
      if (openClaim.error !== null) {
        notice = { id, value: { text: refusalLine(openClaim.kind, 'conflict', now, name), role: 'alert' } };
      }
      openClaim = null;
    }
    focusSoon(
      () =>
        document.querySelector<HTMLElement>(`#guild-${id} ${STEP_ACTION_SELECTOR}`) ??
        document.getElementById(`guild-${id}-notice`),
    );
  });

  function claimControls(guild: MeGuild, kind: ClaimKind | null) {
    if (kind === null) return undefined;
    const open = openClaim?.id === guild.id;
    return {
      open,
      locked: openClaim !== null && !open,
      error: open ? (openClaim?.error ?? null) : null,
      io: CLAIM_IO,
      onOpen: () => {
        notice = null;
        openClaim = { id: guild.id, kind, error: null };
      },
      onCancel: () => closeClaim(guild.id),
      onDone: (line: string) => onClaimDone(guild.id, line),
      onRefused: (code: string, rereads: boolean) => onClaimRefused(guild, kind, code, rereads),
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
    if (outcome.rereads) {
      awaitedStep = { id: guild.id, from: stepOfGuild(guild.id) };
      onCheckAgain();
      return;
    }
    focusSoon(() => document.getElementById(`guild-${guild.id}-notice`));
  }

  async function pay(guild: MeGuild, kind: Exclude<PayBusy, null>): Promise<void> {
    notice = null;
    retryId = null;
    busy = { id: guild.id, kind };
    const outcome = await (kind === 'checkout' ? onCheckout(guild) : onPortal(guild));
    // A redirect leaves the card busy until the page unloads.
    if (outcome.kind === 'refused') onPayRefused(guild, outcome);
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

<div class="px-[18px] pt-5 pb-16 md:px-12 md:pt-8" data-testid="checkout-guild-step">
  <div class="mx-auto flex w-full max-w-[880px] flex-col">
    {#if failed}
      <LoadError message={copy.loadError.title} {onRetry} testid="checkout-guild-error" />
    {:else if me === null || (empty === null && !crestsReady)}
      <p class={`m-0 mb-4 ${GUILD_LEAD_RESERVE}`} aria-hidden="true"></p>
      <ul class="m-0 flex list-none flex-col gap-3 p-0" aria-busy="true" data-testid="checkout-guild-loading">
        <li class="sr-only" role="status">{copy.loading}</li>
        {#each Array.from({ length: skeletonCards }, (_, index) => index) as index (index)}
          <li class={`skeleton-block rounded-panel ${GUILD_CARD_MIN}`} aria-hidden="true"></li>
        {/each}
      </ul>
    {:else if empty !== null}
      <div
        class={`${CARD} mx-auto flex w-full max-w-[640px] flex-col gap-2 p-5`}
        data-testid={empty === copy.emptyA ? 'checkout-guild-empty-a' : 'checkout-guild-empty-b'}
      >
        <p class="text-strong m-0 text-[15px]">{empty.title}</p>
        <p class="text-text m-0 text-[14px]">{empty.body}</p>
        <a href={setupHref} class={`${PRIMARY_BUTTON_FIXED} mt-2 w-fit px-4`}>{empty.button}</a>
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

    {#if me !== null && !failed && empty !== copy.emptyA}
      <div class="text-muted mt-6 flex flex-col gap-1 text-[13px]" data-testid="checkout-guild-foot">
        <p class="m-0">
          {copy.foot.missingGuild}
          <a href={setupHref} class={INLINE_LINK}>{copy.foot.getSetUp}</a>
        </p>
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
