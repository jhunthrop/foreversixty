<!-- web/src/components/account/AccountLogin.svelte -->
<!-- /login's island: the Battle.net button (to the visitor's own `?next=`), or, already signed in,
     the way back to the account. Reads the one shared `/v1/me`. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { ACCOUNT_FAILED, battlenetStartUrl, fetchMeOnce, type Me } from '../../lib/account/api';
  import { safeNextPath } from '../../lib/account/safe-next';
  import { accountSignInCopy } from '../../lib/account/signin-copy';
  import { createQueryState } from '../../lib/data/query.svelte';
  import { API_BASE_URL } from '../../lib/planner/config';
  import { PRIMARY_BUTTON_FIXED } from '../../lib/planner/styles';

  let { next = '/logs' }: { next?: string } = $props();

  const session = createQueryState<Me | null>(`${API_BASE_URL}/v1/me`, () => fetchMeOnce(), {
    scope: 'private',
    ttlMs: 10 * 60 * 1000,
  });
  const me = $derived(session.data);
  const displayName = $derived(me?.user.battletag ?? me?.user.email ?? 'Your account');

  // The static build never sees a visitor's query string, so the real `?next=` is read in the
  // browser, on the visitor's own request (until then, the caller's fallback).
  let resolvedNext = $state<string | null>(null);
  onMount(() => {
    resolvedNext = safeNextPath(new URLSearchParams(window.location.search).get('next'), next);
  });
</script>

<div class="flex flex-col gap-6" data-testid="login">
  {#if me !== null}
    <p class="text-[14px]">Signed in as {displayName}.</p>
    <a class="{PRIMARY_BUTTON_FIXED} self-start px-5" href="/account" data-testid="login-go-account">
      Your account
    </a>
  {:else}
    <p class="text-muted text-[14px]">{accountSignInCopy.reason}</p>
    <div class="flex flex-col gap-3">
      <a
        class="{PRIMARY_BUTTON_FIXED} self-start px-5"
        href={battlenetStartUrl(resolvedNext ?? next)}
        data-testid="battlenet"
      >
        Sign in with Battle.net
      </a>
      <p class="text-muted text-[13px]">
        Battle.net tells the site your BattleTag and your characters. It never sees a password here.
      </p>
    </div>
  {/if}
  <!-- 21px, the error line's height, reserved so the session landing never shoves the page. -->
  <div class="min-h-[21px]">
    {#if session.status === 'failed'}<p class="text-[14px]" role="alert" data-testid="account-error">
        {ACCOUNT_FAILED}
      </p>{/if}
  </div>
</div>
