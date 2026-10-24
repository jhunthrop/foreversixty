<!-- web/src/components/account/AccountPremiumBlock.svelte -->
<!-- The account page's Premium block: what the plan gives this account, one row each. Reads
     only what GET /v1/me says; an older shape with no `entitlements` renders the free state. -->
<script lang="ts">
  import type { EntitlementsView } from '../../lib/account/api';
  import { premiumBlockCopy } from '../../lib/billing/copy';

  let { entitlements }: { entitlements?: EntitlementsView } = $props();

  const rows = $derived([
    { key: 'serverSims', copy: premiumBlockCopy.rows.serverSims, live: entitlements?.server_sims === true },
    { key: 'history', copy: premiumBlockCopy.rows.history, live: false },
    { key: 'notifications', copy: premiumBlockCopy.rows.notifications, live: false },
    { key: 'supporter', copy: premiumBlockCopy.rows.supporter, live: entitlements?.supporter_mark === true },
  ]);
</script>

<div class="flex flex-col gap-2" data-testid="account-premium-block">
  <h3 class="label text-muted">{premiumBlockCopy.heading}</h3>
  <ul class="flex flex-col gap-1 text-[14px]">
    {#each rows as row (row.key)}
      <li data-testid={`account-premium-${row.key}`}>
        <span class="text-muted">{row.copy.label}</span> · {row.live ? row.copy.on : row.copy.off}
      </li>
    {/each}
  </ul>
</div>
