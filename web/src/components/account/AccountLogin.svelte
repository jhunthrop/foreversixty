<!-- web/src/components/account/AccountLogin.svelte -->
<!-- /login's island, inside the page's one Plate (pages/login.astro): the reason, the Battle.net
     button (to the visitor's own `?next=`) and what Battle.net shares; or, already signed in, who
     is signed in and the way back to the account. Reads the one shared `/v1/me`. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { ACCOUNT_FAILED, battlenetStartUrl, fetchMeOnce, type Me } from '../../lib/account/api';
  import { safeNextPath } from '../../lib/account/safe-next';
  import { accountSignInCopy } from '../../lib/account/signin-copy';
  import { createQueryState } from '../../lib/data/query.svelte';
  import { API_BASE_URL } from '../../lib/planner/config';
  import { myReportsCopy } from '../../lib/reports/my-reports-copy';
  import Button from '../ui/Button.svelte';

  let { next = '/logs' }: { next?: string } = $props();

  const session = createQueryState<Me | null>(`${API_BASE_URL}/v1/me`, () => fetchMeOnce(), {
    scope: 'private',
    ttlMs: 10 * 60 * 1000,
  });
  const me = $derived(session.data);
  const displayName = $derived(me?.user.battletag ?? me?.user.email ?? accountSignInCopy.goAccount);

  // The static build never sees a visitor's query string, so the real `?next=` is read in the
  // browser, on the visitor's own request (until then, the caller's fallback).
  let resolvedNext = $state<string | null>(null);
  onMount(() => {
    resolvedNext = safeNextPath(new URLSearchParams(window.location.search).get('next'), next);
  });
</script>

<div class="flex flex-col items-start gap-3 px-5 pt-4 pb-1" data-testid="login">
  {#if me !== null}
    <p class="text-[15px] leading-6">{accountSignInCopy.signedInAs(displayName)}</p>
    <Button tone="primary" href="/account" testid="login-go-account">{accountSignInCopy.goAccount}</Button>
  {:else}
    <p class="max-w-[72ch] text-[15px] leading-6">{accountSignInCopy.reason}</p>
    <Button tone="primary" href={battlenetStartUrl(resolvedNext ?? next)} testid="battlenet"
      >{myReportsCopy.signInAction}</Button
    >
    <p class="text-muted text-[13px] leading-5">{accountSignInCopy.privacy}</p>
  {/if}
  <!-- The error line's height, reserved so the session landing never shoves the page. -->
  <p class="-mt-3 min-h-5 text-[14px] leading-5" role="status">
    {#if session.status === 'failed'}{ACCOUNT_FAILED}{/if}
  </p>
</div>
