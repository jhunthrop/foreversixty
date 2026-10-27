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
  import { BillingApiError, startCheckout } from '../lib/billing/api';
  import { checkoutCopy } from '../lib/billing/copy';
  import { safeNextPath } from '../lib/account/safe-next';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { isGuildPlan, parseCheckoutParams, type CheckoutParams } from '../lib/billing/checkout-params';
  import CheckoutGuildChooser from './checkout/CheckoutGuildChooser.svelte';

  const INVALID_LINK_MESSAGE = 'Something is wrong with that link. Start again from the premium page.';

  let invalid = $state(false);
  let message = $state(checkoutCopy.redirecting);
  let showRetry = $state(false);
  let request = $state<CheckoutParams | null>(null);
  let chooserGuilds = $state<MeGuild[] | null>(null);
  let chooserFailed = $state(false);
  let choosing = $state(false);

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
      forgetSession();
      const me = await fetchMeOnce();
      if (isEntitled(me, parsed)) {
        message = checkoutCopy.confirmed;
        return;
      }
      if (attempt < POLL_ATTEMPTS - 1) await new Promise((resolve) => setTimeout(resolve, POLL_INTERVAL_MS));
    }
    message = checkoutCopy.confirmationSlow;
  }

  function showChooser(me: Me): void {
    choosing = true;
    chooserFailed = false;
    chooserGuilds = me.guilds;
  }

  async function reloadChooser(): Promise<void> {
    chooserFailed = false;
    chooserGuilds = null;
    forgetSession();
    try {
      const me = await fetchMeOnce();
      if (me === null) throw new Error('signed out');
      chooserGuilds = me.guilds;
    } catch {
      chooserFailed = true;
    }
  }

  function chooseGuild(guildId: number): void {
    if (request === null) return;
    choosing = false;
    request = { ...request, guildId };
    message = checkoutCopy.redirecting;
    void launch(request);
  }

  async function launch(parsed: CheckoutParams): Promise<void> {
    const me = await fetchMeOnce();
    if (me === null) {
      window.location.assign(`/login?next=${encodeURIComponent(safeNextPath(currentUrl(), '/premium'))}`);
      return;
    }
    if (parsed.status === 'success') {
      void pollForEntitlement(parsed);
      return;
    }
    if (isGuildPlan(parsed.plan) && parsed.guildId === undefined) {
      showChooser(me);
      return;
    }
    try {
      const result = await startCheckout({
        plan: parsed.plan,
        interval: parsed.interval,
        guildId: parsed.guildId,
        comebackId: parsed.comebackId,
      });
      window.location.assign(result.checkout_url);
    } catch (error) {
      showRetry = false;
      if (error instanceof BillingApiError) {
        if (error.status === 503) {
          message = checkoutCopy.notOpenYet;
        } else if (error.status === 409) {
          message = checkoutCopy.guildAlreadyOnPlan;
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
    const parsed = parseCheckoutParams(window.location.search);
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
    {#if choosing}
      <CheckoutGuildChooser
        guilds={chooserGuilds}
        failed={chooserFailed}
        onChoose={chooseGuild}
        onRetry={() => void reloadChooser()}
      />
    {:else}
      <p class="min-h-[21px] text-[14px]" role="alert" data-testid="checkout-message">{message}</p>
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
