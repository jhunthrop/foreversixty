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
  import Tag from '../ui/Tag.svelte';
  import { button } from '../../lib/ui/button';
  import { checkoutGuildBusy } from '../../lib/billing/checkout-guild-errors';
  import Button from '../ui/Button.svelte';
  import SectionTitle from '../ui/SectionTitle.svelte';

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
  const WIDTH = 'w-full sm:w-fit';

  let titleEl = $state<HTMLElement | undefined>(undefined);
  $effect(() => titleEl?.focus());
</script>

<section class="flex flex-col gap-2" data-testid="premium-owned">
  <SectionTitle bind:element={titleEl} focusable class="m-0" title={view.title} />
  <p class="m-0 flex flex-wrap items-center gap-x-2 text-[14px]" data-testid="premium-owned-line">
    <span>{view.line}</span>
    {#if view.founder}
      <Tag title={accountPlanCopy.founder.title}
        >{accountPlanCopy.founder.pill}<span class="sr-only">{founderTail}</span></Tag
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
      <Button
        class={WIDTH}
        onclick={onPortal}
        {busy}
        busyLabel={checkoutGuildBusy.portal}
        testid="premium-owned-action">{view.action.label}</Button
      >
    {:else}
      <a class="{button('secondary')} {WIDTH}" href={view.action.href} data-testid="premium-owned-action"
        >{view.action.label}</a
      >
    {/if}
  </div>
</section>
