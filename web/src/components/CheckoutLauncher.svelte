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
  import { fetchMeOnce, forgetSession, type Me } from '../lib/account/api';
  import { BillingApiError, openPortal, startCheckout } from '../lib/billing/api';
  import { runPremiumCheckout } from '../lib/billing/premium-owned-flow';
  import { checkoutCopy } from '../lib/billing/copy';
  import { checkoutGuildCopy, type GuildTier } from '../lib/billing/checkout-guild-copy';
  import { replacePlanInUrl } from '../lib/billing/checkout-plan-switch';
  import { founderOpen } from '../lib/billing/plans';
  import { safeNextPath } from '../lib/account/safe-next';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { isGuildPlan, parseCheckoutParams, type CheckoutParams } from '../lib/billing/checkout-params';
  import CheckoutBand from './checkout/CheckoutBand.svelte';
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
  let showRetry = $state(false);
  let request = $state<CheckoutParams | null>(null);
  /** The guild step shows (spec §5): a guild plan with no guild id. */
  let stepShown = $state(false);
  /** Null while the step's /v1/me read is in flight. */
  let stepMe = $state<Me | null>(null);
  let stepFailed = $state(false);
  /** The account that already has Premium: the panel replaces the redirect. */
  let ownedMe = $state<Me | null>(null);
  let portalBusy = $state(false);
  let portalFailed = $state(false);
  const ownedNow = new Date();

  const ALREADY_ON_PLAN_CODE = 'conflict';
  const NOT_FOUND = 404;

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

  async function pollForEntitlement(parsed: CheckoutParams): Promise<void> {
    message = checkoutCopy.processing;
    for (let attempt = 0; attempt < POLL_ATTEMPTS; attempt += 1) {
      // Bypass fetchMeOnce's per-page cache: each attempt must be a fresh request, or the
      // loop would just re-read the same stale snapshot fetched before the webhook ran.
      io.forgetSession();
      const me = await io.fetchMe();
      if (isEntitled(me, parsed)) {
        message = checkoutCopy.confirmed;
        return;
      }
      if (attempt < POLL_ATTEMPTS - 1) await new Promise((resolve) => setTimeout(resolve, POLL_INTERVAL_MS));
    }
    message = checkoutCopy.confirmationSlow;
  }

  function redirectToLogin(): void {
    window.location.assign(`/login?next=${encodeURIComponent(safeNextPath(currentUrl(), '/premium'))}`);
  }

  function showGuildStep(me: Me): void {
    stepShown = true;
    stepFailed = false;
    stepMe = me;
  }

  async function reloadGuildStep(): Promise<void> {
    stepFailed = false;
    stepMe = null;
    io.forgetSession();
    try {
      const me = await io.fetchMe();
      if (me === null) {
        redirectToLogin();
        return;
      }
      stepMe = me;
    } catch {
      stepFailed = true;
    }
  }

  /** Spec §5.3: Switch to Guild / Raid guild rewrites the URL's plan in place; the band follows. */
  function switchPlan(to: GuildTier): void {
    if (request === null) return;
    replacePlanInUrl(to);
    request = { ...request, plan: to };
  }

  function showOwned(me: Me): void {
    ownedMe = me;
    showRetry = false;
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
    if (parsed.status === 'success') {
      void pollForEntitlement(parsed);
      return;
    }
    if (isGuildPlan(parsed.plan) && parsed.guildId === undefined) {
      showGuildStep(me);
      return;
    }
    const checkoutRequest = {
      plan: parsed.plan,
      interval: parsed.interval,
      guildId: parsed.guildId,
      comebackId: parsed.comebackId,
    };
    try {
      if (parsed.plan === 'premium') {
        const outcome = await runPremiumCheckout(io, me, checkoutRequest, new Date());
        if (outcome.kind === 'owned') showOwned(outcome.me);
        else if (outcome.kind === 'login') redirectToLogin();
        else window.location.assign(outcome.url);
        return;
      }
      const result = await io.startCheckout(checkoutRequest);
      window.location.assign(result.checkout_url);
    } catch (error) {
      showRetry = false;
      if (error instanceof BillingApiError) {
        if (error.status === 503) {
          message = checkoutCopy.notOpenYet;
        } else if (error.status === NOT_FOUND) {
          message = checkoutCopy.comebackGone;
        } else if (error.status === 409) {
          // Only the already-on-plan refusal has this sentence; any other 409 (for example
          // no_transfer_offer) states its own reason.
          message =
            error.code === ALREADY_ON_PLAN_CODE && isGuildPlan(parsed.plan)
              ? checkoutCopy.guildAlreadyOnPlan
              : error.message;
        } else if (error.status === 403) {
          // The API's own sentence, verbatim (RunControl.svelte's own convention for an
          // engine's words: "never paraphrased, never hidden"). A 403 here covers two
          // distinct, honest refusals -- not a verified officer/leader of this guild, or a
          // verified officer whose guild's claim is currently contested and frozen -- and
          // the server already states which one applies; inventing a single generic
          // sentence here would flatten that distinction. Retrying does not help either
          // way, so showRetry stays false.
          message = error.message;
        } else {
          message = checkoutCopy.genericFailure;
          showRetry = true;
        }
      } else {
        message = checkoutCopy.genericFailure;
        showRetry = true;
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
    void launch(parsed);
  });
</script>

<CheckoutBand plan={request?.plan ?? null} interval={request?.interval ?? DEFAULT_INTERVAL} {founder} />

{#if stepShown && request !== null && request.plan !== 'premium'}
  <CheckoutGuildStep
    me={stepMe}
    failed={stepFailed}
    plan={request.plan}
    interval={request.interval}
    {founder}
    selfUrl={currentUrl()}
    onRetry={() => void reloadGuildStep()}
    onSwitch={switchPlan}
  />
{:else}
  <!-- The narrow body under the band: the redirect and refusal messages, the invalid-link
       line, and the plan=premium slot (Task 12's PremiumOwnedPanel, spec §4.2). -->
  <div
    class="mx-auto flex w-full max-w-[640px] flex-col gap-3 px-[18px] pt-10 pb-16 md:px-0"
    data-testid="checkout-body"
  >
    {#if invalid}
      <p class="min-h-[21px] text-[14px]" role="alert" data-testid="checkout-message">
        {checkoutGuildCopy.invalidLink.text}
        <a href={PREMIUM_HREF} class="underline">{checkoutGuildCopy.invalidLink.button}</a>
      </p>
    {:else if ownedMe !== null}
      <PremiumOwnedPanel
        me={ownedMe}
        now={ownedNow}
        busy={portalBusy}
        onPortal={() => void openBillingPortal()}
      />
      {#if portalFailed}
        <p class="text-[14px]" role="alert" data-testid="checkout-portal-error">
          {checkoutCopy.genericFailure}
        </p>
      {/if}
    {:else}
      <p class="min-h-[21px] text-[14px]" role="alert" data-testid="checkout-message">
        {message}
      </p>
    {/if}
    {#if !invalid && showRetry && request !== null}
      <button
        type="button"
        class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4`}
        onclick={() => void launch(request!)}
      >
        {checkoutCopy.retry}
      </button>
    {/if}
  </div>
{/if}
