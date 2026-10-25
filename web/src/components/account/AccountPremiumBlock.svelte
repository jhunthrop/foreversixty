<!-- web/src/components/account/AccountPremiumBlock.svelte -->
<!-- The account page's Premium rows, under the plan line: what the plan gives this account as
     a key/value panel (muted label left, value in a fixed right column). Reads only what
     GET /v1/me says; an older shape with no `entitlements` renders the free state. A live
     value links to where the benefit lives; one that has not shipped, or is off for the plan, is muted plain text. -->
<script lang="ts">
  import type { EntitlementsView, MeCharacter } from '../../lib/account/api';
  import { mainCharacter } from '../../lib/account/main-character';
  import { premiumRowCopy } from '../../lib/account/notifications-copy';
  import { premiumBlockCopy } from '../../lib/billing/copy';
  import { characterHref } from '../../lib/characters';

  const SIM_HREF = '/sim';
  const NOTIFICATIONS_HREF = '#discord';

  let {
    entitlements,
    characters = [],
    mainKey,
    serverRunsThisMonth,
    notifications = { value: premiumRowCopy.off, linked: false },
  }: {
    entitlements?: EntitlementsView;
    characters?: readonly MeCharacter[];
    mainKey?: string;
    serverRunsThisMonth?: number;
    /**
     * The Discord row: `premiumRowValue(me)` and whether it links to #discord. `null` hides the
     * row (`premiumRowVisible(me)` is false: the deployment has notifications off); `undefined`
     * keeps the muted `Off · Premium` row. `Account.svelte` is the only caller.
     */
    notifications?: { value: string; linked: boolean } | null;
  } = $props();

  const main = $derived(mainCharacter(characters, mainKey));
  const supporterHref = $derived(main === null ? null : characterHref(main.region, main.ruleset, main.name));

  const historyHref = $derived(supporterHref === null ? null : `${supporterHref}#history`);
  const historyCopy = $derived({
    ...premiumBlockCopy.rows.history,
    on: historyHref === null ? premiumBlockCopy.rows.history.on : premiumBlockCopy.rows.history.onWithMain,
  });

  const notificationsRow = $derived(
    notifications === null
      ? []
      : [
          {
            key: 'notifications',
            copy: {
              label: premiumBlockCopy.rows.notifications.label,
              on: notifications.value,
              off: notifications.value,
            },
            live: notifications.linked,
            href: NOTIFICATIONS_HREF,
            muted: !notifications.linked,
          },
        ],
  );

  const rows = $derived([
    {
      key: 'serverSims',
      copy: {
        ...premiumBlockCopy.rows.serverSims,
        on: premiumBlockCopy.rows.serverSims.on(serverRunsThisMonth),
      },
      live: entitlements?.server_sims === true,
      href: SIM_HREF,
      muted: false,
    },
    {
      key: 'history',
      copy: historyCopy,
      live: entitlements?.history === true,
      href: historyHref,
      muted: entitlements?.history !== true,
    },
    ...notificationsRow,
    {
      key: 'supporter',
      copy: premiumBlockCopy.rows.supporter,
      live: entitlements?.supporter_mark === true,
      href: supporterHref,
      muted: false,
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
        <span class={row.muted ? 'text-muted' : 'text-text'}>{value}</span>
      {/if}
    </li>
  {/each}
</ul>
