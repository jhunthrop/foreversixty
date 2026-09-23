<!-- web/src/components/SessionNav.svelte -->
<!-- The header's "Sign in" / signed-in-name link, on every page. Split out of
     Account.svelte's `nav` mode (see this task's brief in the plan for why): shares only
     `fetchMeOnce`, nothing else. The invisible placeholder below is load-bearing -- see
     web/tests/e2e/sim-tools-phone.spec.ts's `load()` helper, which waits for it to swap to
     a real `<a>` before measuring hit targets, and web/src/components/Account.svelte's own
     copy of this same trick for the CLS story (a narrower placeholder lets the header's
     `1fr` brand column stretch, wrapping the wordmark to two lines and back on hydration). -->
<script lang="ts">
  import { fetchMeOnce, type Me } from '../lib/account/api';
  import { createQueryState } from '../lib/data/query.svelte';
  import { guildHref } from '../lib/characters';
  import { API_BASE_URL } from '../lib/planner/config';

  // One `/v1/me` read, shared with every other island through the client cache
  // (web/src/lib/data/query.ts) -- see HomeAccountPanel.svelte and Account.svelte's own
  // copies of this same call. `createQueryState`'s $effect does the one-load-on-mount work
  // this component used to do by hand.
  const session = createQueryState<Me | null>(`${API_BASE_URL}/v1/me`, () => fetchMeOnce(), {
    scope: 'private',
    ttlMs: 10 * 60 * 1000,
  });

  const signedIn = $derived(session.data !== null);
  const displayName = $derived(session.data?.user.battletag ?? session.data?.user.email ?? 'Your account');
</script>

<div class="flex items-center gap-3 text-[13px]" data-testid="session-nav">
  {#if session.status === 'loading'}
    <span class="invisible inline-flex min-h-11 items-center px-2 md:min-h-0 md:px-0" aria-hidden="true">
      Sign in
    </span>
  {:else if signedIn}
    <a href="/account" class="text-nav hover:text-strong inline-flex min-h-11 items-center md:min-h-0">
      {displayName}
    </a>
    {#if session.data !== null && session.data.guilds.length > 0}
      <!-- The API's GET /v1/me now orders guilds by most-recently-active membership
           (spec section 2.6's ORDER BY change), so [0] is the right one with no further
           sorting here. -->
      <a
        href={guildHref(
          session.data.guilds[0].region,
          session.data.guilds[0].ruleset,
          session.data.guilds[0].name,
        )}
        class="text-nav hover:text-strong inline-flex min-h-11 items-center md:min-h-0"
        data-testid="session-my-guild"
      >
        My guild
      </a>
    {/if}
  {:else}
    <a
      href="/login"
      class="text-nav hover:text-strong inline-flex min-h-11 items-center px-2 md:min-h-0 md:px-0"
    >
      Sign in
    </a>
  {/if}
</div>
