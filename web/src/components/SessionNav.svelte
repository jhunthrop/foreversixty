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

  let me = $state<Me | null>(null);
  let status = $state<'loading' | 'ready' | 'failed'>('loading');

  const signedIn = $derived(me !== null);
  const displayName = $derived(me?.user.battletag ?? me?.user.email ?? 'Your account');

  async function load(): Promise<void> {
    status = 'loading';
    try {
      me = await fetchMeOnce();
      status = 'ready';
    } catch {
      status = 'failed';
    }
  }

  // One load on mount. $effect rather than onMount so this behaves identically whether
  // Astro hydrates it or a future caller mounts it by hand (Account.svelte's own comment
  // gives the same reason for the same choice).
  $effect(() => {
    void load();
  });
</script>

<div class="flex items-center gap-3 text-[13px]" data-testid="session-nav">
  {#if status === 'loading'}
    <span class="invisible inline-flex min-h-11 items-center px-2 md:min-h-0 md:px-0" aria-hidden="true">
      Sign in
    </span>
  {:else if signedIn}
    <a href="/account" class="text-nav hover:text-strong inline-flex min-h-11 items-center md:min-h-0">
      {displayName}
    </a>
  {:else}
    <a
      href="/login"
      class="text-nav hover:text-strong inline-flex min-h-11 items-center px-2 md:min-h-0 md:px-0"
    >
      Sign in
    </a>
  {/if}
</div>
