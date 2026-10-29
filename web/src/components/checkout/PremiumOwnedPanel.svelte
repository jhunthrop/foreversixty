<!-- web/src/components/checkout/PremiumOwnedPanel.svelte -->
<!-- "You already have Premium." (design/specs/2026-10-27-subscribe-and-join.md 4.2): the panel
     /premium/checkout shows in place of launching a second checkout. The page's band owns the
     h1, so the title is an h2; it takes focus when it replaces the redirecting message. The
     words and the button come from premiumOwnedView; this only draws them. -->
<script lang="ts">
  import type { Me } from '../../lib/account/api';
  import { accountPlanCopy } from '../../lib/billing/account-plan-copy';
  import { billingBlockCopy } from '../../lib/billing/copy';
  import { premiumOwnedView } from '../../lib/billing/premium-owned';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { CARD } from '../../lib/ui/card';

  let {
    me,
    now,
    zone,
    busy = false,
    onPortal,
  }: {
    me: Me;
    now: Date;
    zone?: string;
    busy?: boolean;
    /** Opens the billing portal; the `Manage billing` button calls it. */
    onPortal: () => void;
  } = $props();

  const view = $derived(premiumOwnedView(me, now, zone));
  const founderTail = ` ${accountPlanCopy.founder.title}`;
  const BUTTON = `${SECONDARY_BUTTON_FIXED} border-line-warm text-text w-full px-4 sm:w-fit`;

  let titleEl = $state<HTMLHeadingElement | undefined>(undefined);
  $effect(() => titleEl?.focus());
</script>

<section class={`${CARD} measure-form flex flex-col gap-2 p-5`} data-testid="premium-owned">
  <h2 bind:this={titleEl} tabindex="-1" class="heading-name text-strong m-0 text-[15px] outline-none">
    {view.title}
  </h2>
  <p class="m-0 flex flex-wrap items-center gap-x-2 text-[14px]" data-testid="premium-owned-line">
    <span>{view.line}</span>
    {#if view.founder}
      <span class="pill pill-site" title={accountPlanCopy.founder.title}
        >{accountPlanCopy.founder.pill}<span class="sr-only">{founderTail}</span></span
      >
    {/if}
  </p>
  {#if view.pastDue}
    <p class="text-strong m-0 text-[13px]" role="alert" data-testid="premium-owned-past-due">
      {billingBlockCopy.pastDueBanner}
    </p>
  {:else if view.line2 !== null}
    <p class="text-muted m-0 text-[13px]" data-testid="premium-owned-line2">{view.line2}</p>
  {/if}
  <div class="mt-3">
    {#if view.action.kind === 'portal'}
      <button class={BUTTON} onclick={onPortal} disabled={busy} data-testid="premium-owned-action">
        {view.action.label}
      </button>
    {:else}
      <a
        class="{BUTTON} inline-flex items-center justify-center"
        href={view.action.href}
        data-testid="premium-owned-action">{view.action.label}</a
      >
    {/if}
  </div>
</section>
