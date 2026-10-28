<!-- web/src/components/account/AccountPremiumBlock.svelte -->
<!-- The account page's Premium rows, under the plan line: what the plan gives this account as
     a key/value panel (muted label left, value in a fixed right column). Reads only what
     GET /v1/me says; an older shape with no `entitlements` renders the free state. A live
     value links to where the benefit lives; one that has not shipped, or is off for the plan, is muted plain text. -->
<script lang="ts">
  import type { EntitlementsView, MeCharacter } from '../../lib/account/api';
  import { mainCharacter } from '../../lib/account/main-character';
  import { premiumRowCopy } from '../../lib/account/notifications-copy';
  import type { AccountPlanLine } from '../../lib/account/plan-line';
  import { accountPlanCopy } from '../../lib/billing/account-plan-copy';
  import { billingBlockCopy, premiumBlockCopy } from '../../lib/billing/copy';
  import { premiumSourceCopy } from '../../lib/billing/premium-source-copy';
  import { characterHref } from '../../lib/characters';
  import { INLINE_LINK } from '../../lib/guild/field-styles';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';

  const SIM_HREF = '/sim';
  const NOTIFICATIONS_HREF = '#discord';

  let {
    entitlements,
    characters = [],
    mainKey,
    serverRunsThisMonth,
    plan,
    seePlans = false,
    pastDue = false,
    busy = false,
    onBilling,
    notifications = { value: premiumRowCopy.off, linked: false },
  }: {
    entitlements?: EntitlementsView;
    characters?: readonly MeCharacter[];
    mainKey?: string;
    serverRunsThisMonth?: number;
    /** The headline above the rows: `accountPlanLine(me, now)`. Omitted, the block draws rows only. */
    plan?: AccountPlanLine;
    /** The free state links to the plans page. */
    seePlans?: boolean;
    /** The own subscription's last payment failed: the alert takes line 2's place (spec 4.3). */
    pastDue?: boolean;
    /** Disables the billing button while a request is running. */
    busy?: boolean;
    /** Opens the billing portal; the button under the line (`plan.billingButton`) calls it. */
    onBilling?: () => void;
    /**
     * The Discord row: `premiumRowValue(me)` and whether it links to #discord. `null` hides the
     * row (`premiumRowVisible(me)` is false: the deployment has notifications off); `undefined`
     * keeps the muted `Premium only` row. `Account.svelte` is the only caller.
     */
    notifications?: { value: string; linked: boolean } | null;
  } = $props();

  const founderTail = ` ${accountPlanCopy.founder.title}`;
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

{#if plan !== undefined}
  <div class="mb-3 flex flex-col gap-1" data-testid="account-plan-headline">
    {#if plan.ended !== null}
      <!-- Spec 3.3: the ended sentence takes the plan line's place, with the gold come-back link. -->
      <p class="text-[14px]" data-testid="account-seat-ended">
        {plan.ended.text}<a class={INLINE_LINK} href={plan.ended.href}>{plan.ended.action}</a>
      </p>
    {:else}
      <p class="flex flex-wrap items-center gap-x-2 text-[14px]" data-testid="account-billing-row">
        <span>{plan.line}</span>
        {#if plan.founder}
          <span class="pill pill-site" title={accountPlanCopy.founder.title} data-testid="account-founder"
            >{accountPlanCopy.founder.pill}<span class="sr-only">{founderTail}</span></span
          >
        {/if}
        {#if seePlans}
          <a class="text-text underline" href="/premium">{billingBlockCopy.seePlans}</a>
        {/if}
      </p>
    {/if}
    {#if pastDue}
      <p class="text-strong text-[13px]" role="alert" data-testid="account-past-due">
        {billingBlockCopy.pastDueBanner}
      </p>
    {:else if plan.line2 !== null}
      <p class="text-muted text-[13px]" data-testid="account-plan-line2">{plan.line2}</p>
    {/if}
    {#if plan.billingButton !== null && onBilling !== undefined}
      <button
        class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text mt-2 w-fit px-4"
        onclick={onBilling}
        disabled={busy}
        data-testid="account-billing-button"
      >
        {plan.billingButton === 'manage'
          ? premiumSourceCopy.buttons.manage
          : premiumSourceCopy.buttons.history}
      </button>
    {/if}
  </div>
{/if}
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
