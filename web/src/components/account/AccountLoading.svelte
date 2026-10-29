<!-- web/src/components/account/AccountLoading.svelte -->
<!-- /account's loading state: the ready page's layout with its contents missing. It takes no
     props: every height comes from CSS (global.css's "account loading reserve" block) keyed on
     the hints Base.astro sets on <html> before paint. -->
<script lang="ts">
  import AccountTitle from './AccountTitle.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
</script>

<!-- The loading state draws the ready page's layout with its contents missing (spec
       2026-10-25-discord-notifications 9.1): the header row, the 8/4 grid, and the rail's
       blocks at their floors. global.css's "account loading reserve" block sizes them
       from the hints Base.astro sets before paint. `.session-reserve` collapses the
       skeletons when `<html>` carries no data-session hint, i.e. the visitor has no
       session cookie and this can only resolve to the sign-in prompt, so that page ends
       under its one button instead of reserving panels of black. -->
<div class="account-header-reserve">
  <AccountTitle />
</div>
<div
  class="session-reserve flex flex-col gap-8 lg:grid lg:grid-cols-[minmax(0,1fr)_440px] lg:items-start lg:gap-8"
>
  <div>
    <Skeleton lines={3} minHeight="account-left-skeleton" testid="account-left-skeleton" />
  </div>
  <div class="account-rail-skeleton flex flex-col gap-8" data-testid="account-rail-skeleton">
    <Skeleton lines={3} minHeight="account-devices-skeleton" testid="account-devices-skeleton" />
    <Skeleton lines={3} minHeight="account-you-skeleton" testid="account-you-skeleton" />
    <div class="flex flex-col gap-4">
      <Skeleton lines={3} minHeight="account-plan-skeleton" testid="account-plan-skeleton" />
      <Skeleton lines={3} minHeight="account-discord-skeleton" testid="account-discord-skeleton" />
    </div>
  </div>
</div>
