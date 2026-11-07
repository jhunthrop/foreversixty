<!-- web/src/components/Account.svelte -->
<!-- The account's two pages over the one shared `/v1/me` read:
       login    /login: the Battle.net button, or the way back to the account once signed in
       account  /account: the whole page (account/AccountPage.svelte, premium spec 2026-11-08 §7) -->
<script lang="ts">
  import AccountLogin from './account/AccountLogin.svelte';
  import AccountPage from './account/AccountPage.svelte';

  let { mode, next = '/logs' }: { mode: 'login' | 'account'; next?: string } = $props();
</script>

{#if mode === 'account'}
  <AccountPage />
{:else}
  <AccountLogin {next} />
{/if}
