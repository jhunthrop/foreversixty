<!-- web/src/components/account/AccountPremiumBlock.svelte -->
<!-- The account page's Premium rows, under the plan line: what the plan gives this account as
     a key/value panel (muted label left, value in a fixed right column). Reads only what
     GET /v1/me says; an older shape with no `entitlements` renders the free state. A live
     value links to where the benefit lives; a coming one is muted plain text. -->
<script lang="ts">
  import type { EntitlementsView, MeCharacter } from '../../lib/account/api';
  import { mainCharacter } from '../../lib/account/main-character';
  import { premiumBlockCopy } from '../../lib/billing/copy';
  import { characterHref } from '../../lib/characters';

  const SIM_HREF = '/sim';

  let {
    entitlements,
    characters = [],
    mainKey,
    serverRunsThisMonth,
  }: {
    entitlements?: EntitlementsView;
    characters?: readonly MeCharacter[];
    mainKey?: string;
    serverRunsThisMonth?: number;
  } = $props();

  const main = $derived(mainCharacter(characters, mainKey));
  const supporterHref = $derived(main === null ? null : characterHref(main.region, main.ruleset, main.name));

  const rows = $derived([
    {
      key: 'serverSims',
      copy: {
        ...premiumBlockCopy.rows.serverSims,
        on: premiumBlockCopy.rows.serverSims.on(serverRunsThisMonth),
      },
      live: entitlements?.server_sims === true,
      href: SIM_HREF,
      coming: false,
    },
    { key: 'history', copy: premiumBlockCopy.rows.history, live: false, href: null, coming: true },
    {
      key: 'notifications',
      copy: premiumBlockCopy.rows.notifications,
      live: false,
      href: null,
      coming: true,
    },
    {
      key: 'supporter',
      copy: premiumBlockCopy.rows.supporter,
      live: entitlements?.supporter_mark === true,
      href: supporterHref,
      coming: false,
    },
  ]);
</script>

<ul class="m-0 flex list-none flex-col p-0 text-[14px]" data-testid="account-premium-block">
  {#each rows as row (row.key)}
    {@const value = row.live ? row.copy.on : row.copy.off}
    <li
      class="grid min-h-7 grid-cols-[10rem_minmax(0,1fr)] items-baseline gap-x-3"
      data-testid={`account-premium-${row.key}`}
    >
      <span class="text-muted">{row.copy.label}</span>
      {#if row.live && row.href !== null}
        <a class="text-text underline" href={row.href} data-testid={`account-premium-${row.key}-link`}
          >{value}</a
        >
      {:else}
        <span class={row.coming ? 'text-muted' : 'text-text'}>{value}</span>
      {/if}
    </li>
  {/each}
</ul>
