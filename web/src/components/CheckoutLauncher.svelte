<!-- web/src/components/CheckoutLauncher.svelte -->
<!-- The one small hydrated island /premium/checkout mounts.

     This site builds statically (astro.config.mjs sets no server output/adapter, and
     wrangler.jsonc's run_worker_first does not list /premium/checkout), so the page's own
     Astro frontmatter runs once at build time and never sees a visitor's actual query
     string -- the same limitation login.astro's own `next` handling already lives with.
     That means the query string this page acts on can only be read here, client-side,
     after hydration. This component is therefore the one place that parses it: `parseCheckoutParams`
     checks every param (`plan`, `interval`, `guild_id`, `status`) against a closed
     enum/shape before anything derived from it is used or shown, exactly the way a
     server-side check would -- an invalid or malformed value falls through to `invalid`,
     which renders a fixed, honest sentence, and the raw query string is never interpolated
     into the page at all, valid or not (spec section 3's "never echo the raw string" rule).

     Once parsed: signed out goes straight to /login?next=<this exact url>; signed in posts
     /v1/billing/checkout and redirects to the returned checkout_url. A ?status=success
     return never trusts the query string for confirmation -- it re-fetches /v1/me, bypassing
     fetchMeOnce's cache with forgetSession() so every poll attempt is a genuine network
     round-trip, and shows "confirmed" only once the fetched entitlement actually says so,
     polling briefly if the webhook has not landed yet (the webhook is the source of truth,
     not the redirect). -->
<script lang="ts">
  import { fetchMeOnce, forgetSession, type Me, type MeGuild } from '../lib/account/api';
  import { BillingApiError, openPortal, startCheckout } from '../lib/billing/api';
  import {
    guildArrival,
    refusalKind,
    runCardCheckout,
    runCardPortal,
    type Arrival,
    type CardOutcome,
  } from '../lib/billing/card-checkout';
  import { runPremiumCheckout } from '../lib/billing/premium-owned-flow';
  import { checkoutCopy } from '../lib/billing/copy';
  import { checkoutGuildCopy, type GuildTier } from '../lib/billing/checkout-guild-copy';
  import { replacePlanInUrl } from '../lib/billing/checkout-plan-switch';
  import { PLANS, founderOpen, formatPrice } from '../lib/billing/plans';
  import { priceFor } from '../lib/billing/plan-costs';
  import { premiumSourceCopy } from '../lib/billing/premium-source-copy';
  import { utcDayMonthYear } from '../lib/dates';
  import { safeNextPath } from '../lib/account/safe-next';
  import { LINK } from '../lib/ui/primitives';
  import {
    isGuildPlan,
    parseCheckoutParams,
    showsGuildStep,
    type CheckoutParams,
  } from '../lib/billing/checkout-params';
  import CheckoutConfirmed from './checkout/CheckoutConfirmed.svelte';
  import CheckoutHead from './checkout/CheckoutHead.svelte';
  import CheckoutSteps, { type CheckoutPhase } from './checkout/CheckoutSteps.svelte';
  import Page from './ui/Page.svelte';
  import Plate from './ui/Plate.svelte';
  import CheckoutGuildStep from './checkout/CheckoutGuildStep.svelte';
  import PremiumOwnedPanel from './checkout/PremiumOwnedPanel.svelte';

  /** The island's outside world, injectable so tests need no network or globals. */
  interface CheckoutDeps {
    fetchMe: () => Promise<Me | null>;
    forgetSession: () => void;
    startCheckout: typeof startCheckout;
    openPortal: typeof openPortal;
  }

  const PRODUCTION_DEPS: CheckoutDeps = {
    fetchMe: () => fetchMeOnce(),
    forgetSession,
    startCheckout,
    openPortal,
  };

  let { search = undefined, deps = {} }: { search?: string; deps?: Partial<CheckoutDeps> } = $props();
  const io: CheckoutDeps = { ...PRODUCTION_DEPS, ...deps };

  const PREMIUM_HREF = '/premium';
  /** The band's interval before the query string is read; only the static render sees it. */
  const DEFAULT_INTERVAL = 'yearly';
  const founder = founderOpen(new Date());

  let invalid = $state(false);
  let message = $state(checkoutCopy.redirecting);
  let phase = $state<CheckoutPhase>('redirecting');
  /** The account's BattleTag for row 1, once /v1/me answered. */
  let battletag = $state<string | null>(null);
  /** Row 2 after the return: a paid plan, or a guild trial's saved card (ruling K1). */
  let paidLine = $state<string>(checkoutCopy.paid);
  /** Above the next steps once Premium is on (ruling K2). */
  let renewal = $state<string | null>(null);
  let polling = $state(false);
  let request = $state<CheckoutParams | null>(null);
  /** The guild step shows (spec §5): set from the URL before /v1/me answers (showsGuildStep). */
  let stepShown = $state(false);
  /** This checkout's own path and query; follows a plan switch, so every `next` returns here. */
  let selfUrl = $state('');
  /** Null while the step's /v1/me read is in flight. */
  let stepMe = $state<Me | null>(null);
  let stepFailed = $state(false);
  /** A `guild_id` arrival's card choice, set once /v1/me answered (spec §5.6). */
  let arrival = $state<Arrival | undefined>(undefined);
  /** The last `Check again` could not read /v1/me; the cards stay as they were. */
  let refreshFailed = $state(false);
  /** The account that already has Premium: the panel replaces the redirect. */
  let ownedMe = $state<Me | null>(null);
  let portalBusy = $state(false);
  let portalFailed = $state(false);
  const ownedNow = new Date();

  const POLL_INTERVAL_MS = 1500;
  const POLL_ATTEMPTS = 4;

  function currentUrl(): string {
    return window.location.pathname + window.location.search;
  }

  // The entitlement that actually confirms this specific purchase landed: for a personal
  // Premium purchase, entitlements.server_sims; for a guild purchase, that guild's own
  // billing record (MeGuild.plan), not merely "the user has an entitlements block at all" --
  // every signed-in user has one of those regardless of whether this purchase succeeded, so
  // checking for its presence alone would report "confirmed" on the very first poll no
  // matter what the webhook has done.
  function isEntitled(me: Me | null, parsed: CheckoutParams): boolean {
    if (me === null) return false;
    if (!isGuildPlan(parsed.plan)) return me.entitlements?.server_sims === true;
    return (
      parsed.guildId !== undefined &&
      me.guilds.some((guild) => guild.id === parsed.guildId && guild.plan != null)
    );
  }

  /** What the return says once the entitlement is in: a guild's trial card line, Premium's renewal. */
  function describeReturn(me: Me, parsed: CheckoutParams): void {
    if (isGuildPlan(parsed.plan)) {
      const plan = me.guilds.find((guild) => guild.id === parsed.guildId)?.plan;
      if (plan?.status === 'trialing' && plan.current_period_end) {
        paidLine = checkoutCopy.trialSaved(utcDayMonthYear(plan.current_period_end));
      }
      return;
    }
    const billing = me.entitlements?.billing;
    if (billing?.current_period_end && billing.interval) {
      const price = formatPrice(
        priceFor('premium', billing.interval, billing.founder === true),
        billing.interval,
      );
      renewal = checkoutCopy.renewalLine(price, utcDayMonthYear(billing.current_period_end));
    }
  }

  async function pollForEntitlement(parsed: CheckoutParams): Promise<void> {
    phase = 'processing';
    message = checkoutCopy.processing;
    polling = true;
    for (let attempt = 0; attempt < POLL_ATTEMPTS; attempt += 1) {
      // Bypass fetchMeOnce's per-page cache: each attempt must be a fresh request, or the
      // loop would just re-read the same stale snapshot fetched before the webhook ran.
      io.forgetSession();
      const me = await io.fetchMe();
      if (me !== null) battletag = me.user.battletag;
      if (me !== null && isEntitled(me, parsed)) {
        describeReturn(me, parsed);
        phase = 'confirmed';
        message = checkoutCopy.confirmed;
        polling = false;
        return;
      }
      if (attempt < POLL_ATTEMPTS - 1) await new Promise((resolve) => setTimeout(resolve, POLL_INTERVAL_MS));
    }
    phase = 'slow';
    message = checkoutCopy.confirmationSlow;
    polling = false;
  }

  function redirectToLogin(): void {
    window.location.assign(`/login?next=${encodeURIComponent(safeNextPath(currentUrl(), '/premium'))}`);
  }

  /** The step's /v1/me read: the skeleton until it answers, LoadError when it fails. `fresh`
   *  bypasses the per-page cache (Try again). */
  async function loadGuildStep(fresh: boolean): Promise<void> {
    stepFailed = false;
    stepMe = null;
    if (fresh) io.forgetSession();
    try {
      const me = await io.fetchMe();
      if (me === null) {
        redirectToLogin();
        return;
      }
      arrival = guildArrival(me.guilds, request?.guildId);
      stepMe = me;
    } catch {
      stepFailed = true;
    }
  }

  /**
   * `Check again`: re-reads /v1/me and redraws the cards in place, with no skeleton between. A
   * failed read says so. Resolves true when /v1/me was read; the guild step owns the focus that
   * follows (the button the player pressed may be gone, spec §5.3).
   */
  async function refreshGuildStep(): Promise<boolean> {
    refreshFailed = false;
    io.forgetSession();
    try {
      const me = await io.fetchMe();
      if (me === null) {
        redirectToLogin();
        return false;
      }
      stepMe = me;
      return true;
    } catch {
      refreshFailed = true;
      return false;
    }
  }

  /** Spec §5.3: Switch to Guild / Raid guild rewrites the URL's plan in place; the band follows. */
  function switchPlan(to: GuildTier): void {
    if (request === null) return;
    replacePlanInUrl(to);
    request = { ...request, plan: to };
    selfUrl = currentUrl();
  }

  /** Follows a card call's redirect; the shell shows a refusal on its card. */
  function followCardOutcome(outcome: CardOutcome): CardOutcome {
    if (outcome.kind === 'redirect') window.location.assign(outcome.url);
    return outcome;
  }

  async function checkoutGuild(guild: MeGuild): Promise<CardOutcome> {
    if (request === null) throw new Error('checkoutGuild before the URL was read');
    const { plan, interval } = request;
    return followCardOutcome(await runCardCheckout(io, { plan, interval, guildId: guild.id }, guild.name));
  }

  async function openGuildPortal(guild: MeGuild): Promise<CardOutcome> {
    return followCardOutcome(await runCardPortal(io, guild.id));
  }

  function showOwned(me: Me): void {
    ownedMe = me;
  }

  async function openBillingPortal(): Promise<void> {
    portalBusy = true;
    portalFailed = false;
    try {
      const result = await io.openPortal(undefined);
      window.location.assign(result.portal_url);
    } catch {
      portalFailed = true;
      portalBusy = false;
    }
  }

  async function launch(parsed: CheckoutParams): Promise<void> {
    const me = await io.fetchMe();
    if (me === null) {
      redirectToLogin();
      return;
    }
    battletag = me.user.battletag;
    if (parsed.status === 'success') {
      if (isGuildPlan(parsed.plan)) paidLine = '';
      void pollForEntitlement(parsed);
      return;
    }
    const checkoutRequest = {
      plan: parsed.plan,
      interval: parsed.interval,
      guildId: parsed.guildId,
      comebackId: parsed.comebackId,
    };
    try {
      const outcome = await runPremiumCheckout(io, me, checkoutRequest, new Date());
      if (outcome.kind === 'owned') showOwned(outcome.me);
      else if (outcome.kind === 'login') redirectToLogin();
      else window.location.assign(outcome.url);
    } catch (error) {
      const kind = refusalKind(error);
      phase = 'refused';
      if (kind === 'not_open') {
        message = checkoutCopy.notOpenYet;
      } else if (kind === 'not_found') {
        message = checkoutCopy.comebackGone;
      } else if (error instanceof BillingApiError && (kind === 'conflict' || kind === 'forbidden')) {
        // The API's own sentence, verbatim; retrying does not help, so there is no Try again.
        message = error.message;
      } else {
        message = checkoutCopy.genericFailure;
        phase = 'error';
      }
    }
  }

  $effect(() => {
    const parsed = parseCheckoutParams(search ?? window.location.search);
    if (parsed === null) {
      invalid = true;
      return;
    }
    request = parsed;
    selfUrl = currentUrl();
    if (showsGuildStep(parsed)) {
      stepShown = true;
      void loadGuildStep(false);
      return;
    }
    void launch(parsed);
  });
</script>

<CheckoutHead
  plan={request?.plan ?? null}
  interval={request?.interval ?? DEFAULT_INTERVAL}
  {founder}
  {invalid}
  showChange={phase === 'redirecting' || phase === 'error' || phase === 'refused'}
  owned={ownedMe !== null}
/>

{#if stepShown && request !== null && request.plan !== 'premium'}
  <Page layout="ledger">
    {#if refreshFailed}
      <p class="m-0 text-[14px]" role="status" data-testid="checkout-refresh-error">
        {checkoutCopy.genericFailure}
      </p>
    {/if}
    <CheckoutGuildStep
      me={stepMe}
      failed={stepFailed}
      plan={request.plan}
      interval={request.interval}
      {founder}
      {selfUrl}
      onRetry={() => void loadGuildStep(true)}
      onCheckAgain={refreshGuildStep}
      onSwitch={switchPlan}
      onCheckout={checkoutGuild}
      onPortal={openGuildPortal}
      {arrival}
    />
  </Page>
{:else}
  <Page layout="form" testid="checkout-body">
    {#if invalid}
      <p class="m-0 text-[15px] leading-6" role="alert" data-testid="checkout-message">
        {checkoutGuildCopy.invalidLink.text}
        <a href={PREMIUM_HREF} class={LINK}>{checkoutGuildCopy.invalidLink.button} ›</a>
      </p>
    {:else if ownedMe !== null}
      <Plate label={premiumSourceCopy.buttons.seePlan} frame="framed">
        <div class="checkout-owned">
          <PremiumOwnedPanel
            me={ownedMe}
            now={ownedNow}
            busy={portalBusy}
            onPortal={() => void openBillingPortal()}
          />
        </div>
      </Plate>
      {#if portalFailed}
        <p class="text-wipe m-0 pt-3 text-[14px]" role="alert" data-testid="checkout-portal-error">
          {checkoutCopy.genericFailure}
        </p>
      {/if}
    {:else if phase === 'confirmed' && request !== null && !isGuildPlan(request.plan)}
      <CheckoutConfirmed planName={PLANS[request.plan].name} {renewal} />
    {:else}
      <div>
        <CheckoutSteps
          {phase}
          {message}
          planName={PLANS[request?.plan ?? 'premium'].name}
          signedInAs={battletag}
          {paidLine}
          busy={polling}
          onRetry={request === null ? undefined : () => void launch(request!)}
          onCheckAgain={request === null ? undefined : () => void pollForEntitlement(request!)}
        />
      </div>
    {/if}
  </Page>
{/if}

<style>
  .checkout-owned {
    padding: 16px 20px 20px;
  }
</style>
