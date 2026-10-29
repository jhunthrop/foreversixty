<!-- web/src/components/account/AccountPremiumBlock.svelte -->
<!-- The account page's Plan card (polish spec 4.3): a head (crown tile, `Plan` over the plan name,
     the Founder tag, the renew or end date at the right), the pending seat offer, four settings
     rows, and a foot with the billing button or See plans. Reads only what GET /v1/me says; an
     older shape with no `entitlements` renders the free state. A live value links to where the
     benefit lives; one that has not shipped, or is off for the plan, is a muted static row. -->
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
  import { INLINE_LINK } from '../../lib/guild/field-styles';
  import { button } from '../../lib/ui/button';
  import type { GlyphName } from '../../lib/ui/glyphs';
  import CardHead from '../ui/CardHead.svelte';
  import Glyph from '../ui/Glyph.svelte';
  import SettingsRow from '../ui/SettingsRow.svelte';
  import Spinner from '../ui/Spinner.svelte';
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
    onBilling,
    notifications = { value: premiumRowCopy.off, linked: false },
    offer,
  }: {
    entitlements?: EntitlementsView;
    characters?: readonly MeCharacter[];
    mainKey?: string;
    serverRunsThisMonth?: number;
    /** The head's source: `accountPlanLine(me, now)`. Omitted, the card draws the free head and rows only. */
    plan?: AccountPlanLine;
    /** The free state links to the plans page. */
    seePlans?: boolean;
    /** The own subscription's last payment failed: the alert takes line 2's place (spec 4.3). */
    pastDue?: boolean;
    /** Disables the billing button while a request is running. */
    busy?: boolean;
    /** Opens the billing portal; the button under the rows (`plan.billingButton`) calls it. */
    onBilling?: () => void;
    /**
     * The Discord row: `premiumRowValue(me)` and whether it links to #discord. `null` hides the
     * row (`premiumRowVisible(me)` is false: the deployment has notifications off); `undefined`
     * keeps the muted `Premium only` row. `Account.svelte` is the only caller.
     */
    notifications?: { value: string; linked: boolean } | null;
    /** A pending seat offer, drawn between the head and the rows. */
    offer?: Snippet;
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
            glyph: 'bell' as GlyphName,
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
      glyph: 'server' as GlyphName,
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
      glyph: 'chart' as GlyphName,
      copy: historyCopy,
      live: entitlements?.history === true,
      href: historyHref,
      muted: entitlements?.history !== true,
    },
    ...notificationsRow,
    {
      key: 'supporter',
      glyph: 'star' as GlyphName,
      copy: premiumBlockCopy.rows.supporter,
      live: entitlements?.supporter_mark === true,
      href: supporterHref,
      muted: false,
    },
  ]);

  const billingLabel = $derived(
    plan?.billingButton === 'manage' ? premiumSourceCopy.buttons.manage : premiumSourceCopy.buttons.history,
  );
  const showCancelNote = $derived(
    plan?.billingButton === 'manage' && plan.head.when?.label === accountPlanCopy.head.renews,
  );
  const showBilling = $derived(plan?.billingButton != null && onBilling !== undefined);
  const showSeePlans = $derived(seePlans && plan?.ended == null);
</script>

<section class="card-surface overflow-hidden" data-testid="account-guilds-plan">
  <CardHead
    glyph="crown"
    label={accountPageCopy.planLabel}
    title={plan?.head.name ?? accountPlanCopy.head.names.free}
    testid="account-plan-head"
  >
    {#snippet titleSnippet()}
      {#if plan?.founder}
        <Tag title={accountPlanCopy.founder.title} testid="account-founder"
          >{accountPlanCopy.founder.pill}<span class="sr-only">{founderTail}</span></Tag
        >
      {/if}
    {/snippet}
    {#snippet aside()}
      {#if plan?.head.when != null}
        <span
          class="text-muted flex flex-col items-end gap-0.5 max-[479px]:items-start"
          data-testid="account-plan-when"
        >
          <span>{plan.head.when.label}</span>
          <span class="text-strong font-semibold">{plan.head.when.date}</span>
        </span>
      {/if}
    {/snippet}
  </CardHead>
  {#if plan !== undefined}
    <div class="-mt-2 flex flex-col gap-1 px-5 pb-4" data-testid="account-plan-headline">
      {#if plan.ended !== null}
        <!-- Spec 3.3: the ended sentence takes the plan line's place, with the gold come-back link. -->
        <p class="text-[14px]" data-testid="account-seat-ended">
          {plan.ended.text}<a class={INLINE_LINK} href={plan.ended.href}>{plan.ended.action}</a>
        </p>
      {:else}
        <!-- The whole plan sentence, for a reader that skips the head's pieces. -->
        <p class="sr-only" data-testid="account-billing-row">{plan.line}</p>
        {#if plan.head.extra !== null}
          <p class="text-muted text-[13px]" data-testid="account-plan-extra">{plan.head.extra}</p>
        {/if}
        {#if showCancelNote}
          <p class="text-muted text-[13px]" data-testid="account-cancel-note">
            {accountPlanCopy.head.cancelAnytime}
          </p>
        {/if}
      {/if}
      {#if pastDue}
        <p class="flex flex-wrap items-center gap-2 text-[13px]" role="alert" data-testid="account-past-due">
          <Tag tone="danger">{billingBlockCopy.pastDueTag}</Tag>
          <span class="text-strong">{billingBlockCopy.pastDueBanner}</span>
        </p>
      {:else if plan.line2 !== null}
        <p class="text-muted text-[13px]" data-testid="account-plan-line2">{plan.line2}</p>
      {/if}
    </div>
  {/if}
  {@render offer?.()}
  <ul class="m-0 flex list-none flex-col p-0" data-testid="account-premium-block">
    {#each rows as row (row.key)}
      {@const value = row.live ? row.copy.on : row.copy.off}
      {@const linked = row.live && row.href !== null}
      <li data-testid={`account-premium-${row.key}`}>
        <SettingsRow
          label={row.copy.label}
          glyph={row.glyph}
          {value}
          muted={!linked && row.muted}
          href={linked ? (row.href ?? undefined) : undefined}
          testid={linked ? `account-premium-${row.key}-link` : undefined}
        />
      </li>
    {/each}
  </ul>
  {#if showBilling || showSeePlans}
    <div class="border-line-soft flex flex-wrap gap-3 border-t px-5 py-4">
      {#if showBilling}
        <button
          class={button('secondary')}
          onclick={onBilling}
          disabled={busy}
          aria-busy={busy ? 'true' : undefined}
          data-testid="account-billing-button"
        >
          {#if busy}<Spinner />{:else}<Glyph name="card" />{/if}{billingLabel}
        </button>
      {/if}
      {#if showSeePlans}
        <a class={button('primary')} href={PREMIUM_LINKS.card} data-testid="account-see-plans"
          >{billingBlockCopy.seePlans}</a
        >
      {/if}
    </div>
  {/if}
</section>
