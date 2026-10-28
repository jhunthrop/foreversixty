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
  import { runPremiumCheckout } from '../lib/billing/premium-owned-flow';
  import { checkoutCopy } from '../lib/billing/copy';
  import { founderOpen, type PlanKey } from '../lib/billing/plans';
  import { safeNextPath } from '../lib/account/safe-next';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import {
    billableGuilds,
    isGuildPlan,
    parseCheckoutParams,
    type CheckoutParams,
  } from '../lib/billing/checkout-params';
  import { fetchGuildHome } from '../lib/guild/api';
  import { guildMarkSrc } from '../lib/guild/mark';
  import CheckoutGuildChooser from './checkout/CheckoutGuildChooser.svelte';
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

  const INVALID_LINK_MESSAGE = 'Something is wrong with that link. Start again from the premium page.';

  let invalid = $state(false);
  let message = $state(checkoutCopy.redirecting);
  let showRetry = $state(false);
  let request = $state<CheckoutParams | null>(null);
  let chooserGuilds = $state<MeGuild[] | null>(null);
  let chooserFailed = $state(false);
  /** Null until every billable guild's crest read has settled, so rows and crests arrive together. */
  let chooserCrests = $state<Record<number, string | null> | null>(null);
  let chooserSummary = $state('');
  let chooserPlan = $state<PlanKey | null>(null);
  let choosing = $state(false);
  let messageEl = $state<HTMLParagraphElement | undefined>(undefined);
  let refocusMessage = $state(false);
  /** The account that already has Premium: the panel replaces the redirect. */
  let ownedMe = $state<Me | null>(null);
  let portalBusy = $state(false);

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

  /** Best effort: a guild whose read fails keeps the plain ring. */
  async function loadCrests(guilds: readonly MeGuild[]): Promise<void> {
    const entries = await Promise.all(
      billableGuilds(guilds).map(async (guild): Promise<[number, string | null]> => {
        try {
          return [guild.id, guildMarkSrc((await fetchGuildHome(guild.id)).guild)];
        } catch {
          return [guild.id, null];
        }
      }),
    );
    chooserCrests = Object.fromEntries(entries);
  }

  function showChooser(me: Me, parsed: CheckoutParams): void {
    choosing = true;
    chooserSummary = checkoutCopy.chosenPlan(parsed.plan, parsed.interval, founderOpen(new Date()));
    chooserPlan = parsed.plan;
    chooserFailed = false;
    chooserGuilds = me.guilds;
    void loadCrests(me.guilds);
  }

  async function reloadChooser(): Promise<void> {
    chooserFailed = false;
    chooserGuilds = null;
    chooserCrests = null;
    io.forgetSession();
    try {
      const me = await io.fetchMe();
      if (me === null) {
        redirectToLogin();
        return;
      }
      chooserGuilds = me.guilds;
      void loadCrests(me.guilds);
    } catch {
      chooserFailed = true;
    }
  }

  function chooseGuild(guildId: number): void {
    if (request === null) return;
    choosing = false;
    request = { ...request, guildId };
    message = checkoutCopy.redirecting;
    refocusMessage = true;
    void launch(request);
  }

  function showOwned(me: Me): void {
    ownedMe = me;
    showRetry = false;
  }

  async function openBillingPortal(): Promise<void> {
    portalBusy = true;
    try {
      const result = await io.openPortal(undefined);
      window.location.assign(result.portal_url);
    } catch {
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
      showChooser(me, parsed);
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
    if (!refocusMessage || messageEl === undefined) return;
    messageEl.focus();
    refocusMessage = false;
  });

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

{#if invalid}
  <p class="min-h-[21px] text-[14px]" role="alert" data-testid="checkout-message">
    {INVALID_LINK_MESSAGE}
    <a href="/premium" class="underline">Premium page</a>
  </p>
{:else}
  <div class="flex flex-col gap-3">
    {#if ownedMe !== null}
      <PremiumOwnedPanel
        me={ownedMe}
        now={new Date()}
        busy={portalBusy}
        onPortal={() => void openBillingPortal()}
      />
    {:else if choosing}
      <CheckoutGuildChooser
        guilds={chooserGuilds}
        failed={chooserFailed}
        crests={chooserCrests}
        summary={chooserSummary}
        plan={chooserPlan}
        onChoose={chooseGuild}
        onRetry={() => void reloadChooser()}
      />
    {:else}
      <p
        bind:this={messageEl}
        tabindex="-1"
        class="min-h-[21px] text-[14px] outline-none"
        role="alert"
        data-testid="checkout-message"
      >
        {message}
      </p>
    {/if}
    {#if showRetry && request !== null}
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
