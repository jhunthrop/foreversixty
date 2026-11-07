<!-- web/src/components/account/AccountOut.svelte -->
<!-- /account signed out (premium spec 2026-11-08 section 7A, -out): the form page with the reason
     and Sign in with Battle.net. Drawn by the page once /v1/me says nobody is signed in, and by the
     static loading state when the pre-paint session hint says so, so a signed-out visit never moves. -->
<script lang="ts">
  import { battlenetStartUrl } from '../../lib/account/api';
  import { accountPageCopy } from '../../lib/account/account-page-copy';
  import { accountSignInCopy } from '../../lib/account/signin-copy';
  import { myReportsCopy } from '../../lib/reports/my-reports-copy';
  import { button } from '../../lib/ui/button';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Plate from '../ui/Plate.svelte';

  const SIGN_IN_NEXT = '/account?signed_in=1';
</script>

<Page layout="form" testid="account-form">
  {#snippet header()}<PageHeader title={accountPageCopy.title} testid="account-head" />{/snippet}
  <Plate label="Sign in">
    <div class="flex flex-col items-start gap-3 px-5 pt-4 pb-5" data-testid="account-signin">
      <p class="text-[15px] leading-6">{accountSignInCopy.reason}</p>
      <a class={button('primary')} href={battlenetStartUrl(SIGN_IN_NEXT)} data-testid="battlenet"
        >{myReportsCopy.signInAction}</a
      >
    </div>
  </Plate>
</Page>
