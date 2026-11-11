<!-- web/src/components/account/AccountPremiumBlock.svelte -->
<!-- /account's Plan section (premium spec 2026-11-08 section 7.3; `#plan`, where `See your plan`
     lands): the plan's name as a Figure with the Founder tag and, at the right, Renews / Ends / Until
     over the date; the plan's lines (or the ended seat's sentence and its way back, or the failed
     payment alert); a pending seat offer; the four benefits as slim rows (a live value links to
     where it lives, an off one is muted text); and the billing foot. Reads only what GET /v1/me
     says; an older shape with no `entitlements` draws the free state. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { PREMIUM_LINKS } from '../../lib/billing/premium-links';
  import type { EntitlementsView, MeCharacter } from '../../lib/account/api';
  import { accountPageCopy } from '../../lib/account/account-page-copy';
  import { mainCharacter } from '../../lib/account/main-character';
  import { premiumRowCopy } from '../../lib/account/notifications-copy';
  import type { AccountPlanLine } from '../../lib/account/plan-line';
  import { accountPlanCopy } from '../../lib/billing/account-plan-copy';
  import { billingBlockCopy, premiumBlockCopy } from '../../lib/billing/copy';
  import { premiumSourceCopy } from '../../lib/billing/premium-source-copy';
  import { characterHref } from '../../lib/characters';
  import { button } from '../../lib/ui/button';
  import { LINK } from '../../lib/ui/primitives';
  import Figure from '../ui/Figure.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import Button from '../ui/Button.svelte';
  import { checkoutGuildBusy } from '../../lib/billing/checkout-guild-errors';
  import Tag from '../ui/Tag.svelte';

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
    error = '',
    onBilling,
    notifications = { value: premiumRowCopy.off, linked: false },
    offer,
  }: {
    entitlements?: EntitlementsView;
    characters?: readonly MeCharacter[];
    mainKey?: string;
    serverRunsThisMonth?: number;
    /** `accountPlanLine(me, now)`. Omitted, the section draws the free head and rows only. */
    plan?: AccountPlanLine;
    /** The free state's foot: See plans. */
    seePlans?: boolean;
    /** The own subscription's last payment failed: the alert takes line 2's place, Manage billing fills. */
    pastDue?: boolean;
    /** The billing portal is opening. */
    busy?: boolean;
    /** The billing portal did not open. */
    error?: string;
    onBilling?: () => void;
    /** `premiumRowValue(me)` and whether it links to #discord; `null` hides the row (notifications off). */
    notifications?: { value: string; linked: boolean } | null;
    /** A pending seat offer, drawn above the benefits. */
    offer?: Snippet;
  } = $props();

  const founderTail = ` ${accountPlanCopy.founder.title}`;
  const main = $derived(mainCharacter(characters, mainKey));
  const mainHref = $derived(main === null ? null : characterHref(main.region, main.ruleset, main.name));
  const rows = $derived.by(() => {
    const r = premiumBlockCopy.rows;
    const history = entitlements?.history === true;
    return [
      {
        key: 'serverSims',
        label: r.serverSims.label,
        value: entitlements?.server_sims === true ? r.serverSims.on(serverRunsThisMonth) : r.serverSims.off,
        href: entitlements?.server_sims === true ? SIM_HREF : null,
      },
      {
        key: 'history',
        label: r.history.label,
        value: history ? (mainHref === null ? r.history.on : r.history.onWithMain) : r.history.off,
        href: history && mainHref !== null ? `${mainHref}#history` : null,
      },
      ...(notifications === null
        ? []
        : [
            {
              key: 'notifications',
              label: r.notifications.label,
              value: notifications.value,
              href: notifications.linked ? NOTIFICATIONS_HREF : null,
            },
          ]),
      {
        key: 'supporter',
        label: r.supporter.label,
        value: entitlements?.supporter_mark === true ? r.supporter.on : r.supporter.off,
        href: entitlements?.supporter_mark === true ? mainHref : null,
      },
    ];
  });

  const showCancelNote = $derived(
    plan?.billingButton === 'manage' && plan.head.when?.label === accountPlanCopy.head.renews,
  );
  const showBilling = $derived(plan?.billingButton != null && onBilling !== undefined);
  const showSeePlans = $derived(seePlans && plan?.ended == null);
</script>

<Section title={accountPageCopy.planLabel} anchor="plan" testid="account-guilds-plan" panel>
  <div class="flex flex-col gap-3">
    <div class="flex min-h-6 flex-wrap items-end justify-between gap-3" data-testid="account-plan-head">
      <span class="inline-flex flex-wrap items-center gap-2">
        <Figure value={plan?.head.name ?? accountPlanCopy.head.names.free} size="sm" />
        {#if plan?.founder}<Tag title={accountPlanCopy.founder.title} testid="account-founder"
            >{accountPlanCopy.founder.pill}<span class="sr-only">{founderTail}</span></Tag
          >{/if}
      </span>
      {#if plan?.head.when != null}
        <span data-testid="account-plan-when">
          <Figure value={plan.head.when.date} label={plan.head.when.label} size="sm" align="end" />
        </span>
      {/if}
    </div>
    {#if plan !== undefined}
      {#if plan.ended !== null}
        <p class="max-w-[72ch] text-[15px] leading-6" data-testid="account-seat-ended">
          {plan.ended.text}<a class={LINK} href={plan.ended.href}>{plan.ended.action}</a>
        </p>
      {:else}
        {#if plan.head.extra !== null}
          <p class="text-small text-muted" data-testid="account-plan-extra">{plan.head.extra}</p>
        {/if}
        {#if showCancelNote}
          <p class="text-small text-muted" data-testid="account-cancel-note">
            {accountPlanCopy.head.cancelAnytime}
          </p>
        {/if}
      {/if}
      {#if pastDue}
        <p class="flex flex-wrap items-center gap-2" role="alert" data-testid="account-past-due">
          <Tag tone="danger">{billingBlockCopy.pastDueTag}</Tag>
          <span class="text-strong">{billingBlockCopy.pastDueBanner}</span>
        </p>
      {:else if plan.line2 !== null}
        <p class="text-small text-muted" data-testid="account-plan-line2">{plan.line2}</p>
      {/if}
    {/if}
    {@render offer?.()}
    <div class="account-premium-rows" data-testid="account-premium-block">
      {#each rows as row (row.key)}
        <Row title={row.label} density="slim" open wrapTrail testid={`account-premium-${row.key}`}>
          {#snippet trail()}
            {#if row.href !== null}
              <a class="{LINK} text-[15px]" href={row.href} data-testid={`account-premium-${row.key}-link`}
                >{row.value} ›</a
              >
            {:else}
              <span class="text-small text-muted">{row.value}</span>
            {/if}
          {/snippet}
        </Row>
      {/each}
    </div>
    {#if showBilling || showSeePlans}
      <div class="flex flex-wrap items-center gap-x-4 gap-y-[10px]">
        {#if showBilling}
          <Button
            tone={pastDue ? 'primary' : 'secondary'}
            onclick={onBilling}
            {busy}
            busyLabel={checkoutGuildBusy.portal}
            testid="account-billing-button"
          >
            {plan?.billingButton === 'manage'
              ? premiumSourceCopy.buttons.manage
              : premiumSourceCopy.buttons.history}
          </Button>
        {/if}
        {#if showSeePlans}
          <a class={button('primary')} href={PREMIUM_LINKS.card} data-testid="account-see-plans"
            >{billingBlockCopy.seePlans}</a
          >
        {/if}
      </div>
    {/if}
    {#if error !== ''}<p class="text-wipe text-[13px]" role="alert" data-testid="account-billing-error">
        {error}
      </p>{/if}
  </div>
</Section>

<style>
  /* In the rail (from 1024) each value sits under its label, as the phone stacks them: the rail is too
     narrow for a label and a long link side by side. */
  @media (min-width: 1024px) {
    .account-premium-rows :global(.ui-row > :last-child) {
      grid-column: 1 / -1;
      justify-self: start;
    }
  }
</style>
