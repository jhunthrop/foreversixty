<!-- web/src/components/guild/GuildBillingSection.svelte -->
<!-- The Billing section of guild settings (launch pricing spec 3.2): it reads the guild billing view
     itself, reserves its ready height with Skeleton lines while that loads, and draws the blocks in
     order: plan line, size line, chooser, the seats and Who pays mount points, Records. -->
<script lang="ts">
  import { tick } from 'svelte';
  import { invalidateGuild } from '../../lib/guild/api';
  import {
    INITIAL_UPGRADE_STATE,
    focusIsPending,
    createUpgradeFlow,
    type UpgradeFlowState,
  } from '../../lib/guild/upgrade-flow';
  import { getGuildBilling, previewUpgrade, upgrade } from '../../lib/billing/guild-billing-api';
  import type { GuildBillingSettings } from '../../lib/billing/contract';
  import { guildBillingCopy as copy } from '../../lib/billing/guild-billing-copy';
  import { settingsWords } from '../../lib/guild/settings-view';
  import {
    billingState,
    canBill,
    movesFromPlanLine,
    offeredToReader,
    planSeatCount,
    showSizeLine,
    showsReadOnlyLine,
    showsChooser,
  } from '../../lib/guild/billing-state';
  import {
    GUILD_BILLING_LOADING,
    GUILD_BILLING_PHONE_MAX_PX,
    guildBillingReserveKey,
  } from '../../lib/guild/layout';
  import type { CharacterPath } from '../../lib/characters';
  import {
    chooserPreselect,
    parseBillingPreselect,
    shouldOpenUpgrade,
    showsBillingHandoff,
    type BillingPreselect,
  } from '../../lib/guild/billing-handoff';
  import { readReserve, writeReserve } from '../../lib/layout-reserve';
  import { minuteClock } from '../../lib/guild/minute-clock.svelte';
  import Card from '../ui/Card.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import BillingHandoff from './BillingHandoff.svelte';
  import GuildPlanChooser from './billing/GuildPlanChooser.svelte';
  import GuildPlanLine from './billing/GuildPlanLine.svelte';
  import GuildSeats from './billing/GuildSeats.svelte';
  import GuildSizeLine from './billing/GuildSizeLine.svelte';
  import GuildTransferOffer from './billing/GuildTransferOffer.svelte';
  import GuildUpgradeConfirm from './billing/GuildUpgradeConfirm.svelte';
  import GuildWhoPays from './billing/GuildWhoPays.svelte';

  let {
    guildId,
    path,
    guildName,
    now: fixedNow,
  }: { guildId: number; path: CharacterPath; guildName: string; now?: Date } = $props();

  /** `?plan&interval&upgrade` off the settings link; read once, on the client. */
  const preselect: BillingPreselect =
    typeof window === 'undefined'
      ? { tier: null, interval: null, upgrade: false }
      : parseBillingPreselect(window.location.search);
  let upgradeOpened = false;

  /** The minute tick: a seat lock that expires while the page is open turns over without a reload. */
  const clock = minuteClock(() => fixedNow);
  const now = $derived(clock.now);

  let view = $state<GuildBillingSettings | null>(null);
  let failed = $state(false);
  let flow = $state<UpgradeFlowState>(INITIAL_UPGRADE_STATE);
  let sectionEl = $state<HTMLElement | null>(null);
  let statusEl = $state<HTMLElement | null>(null);

  const viewState = $derived(billingState(view, now));

  /** The reserve is per guild and per layout: a phone and a desktop render differ by ~250 px. */
  const reserveKey = (): string =>
    guildBillingReserveKey(
      guildId,
      typeof window !== 'undefined' && window.innerWidth <= GUILD_BILLING_PHONE_MAX_PX,
    );
  const storedReserve = typeof window === 'undefined' ? null : readReserve(reserveKey());
  let readyHeight = $state(0);
  /** Only the height a read first settles at is kept: an open confirm or picker is not the page. */
  let reserveWritten = false;
  $effect(() => {
    if (reserveWritten || view === null || readyHeight <= 0) return;
    reserveWritten = true;
    writeReserve(reserveKey(), readyHeight);
  });
  const payer = $derived(view !== null && canBill(view));
  /** The lower blocks (size line, chooser, seats, the read-only line) draw nothing for a plain paid plan. */
  const hasLowerBlocks = $derived(
    view !== null &&
      (showSizeLine(view) ||
        flow.confirming ||
        (showsChooser(viewState) && payer) ||
        view.seats.length > 0 ||
        showsReadOnlyLine(view, payer)),
  );

  /** A read that finishes after a newer one began is ignored, so it cannot overwrite the newer view. */
  let loadSequence = 0;

  async function load(): Promise<boolean> {
    const sequence = ++loadSequence;
    failed = false;
    try {
      const fresh = await getGuildBilling(guildId);
      if (sequence === loadSequence) view = fresh;
      return true;
    } catch {
      if (sequence === loadSequence) failed = true;
      return false;
    }
  }

  const reload = async (): Promise<void> => {
    await load();
  };

  $effect(() => {
    void load();
  });

  const upgradeFlow = createUpgradeFlow({
    previewUpgrade: () => previewUpgrade(guildId),
    upgrade: (prorationDate) => upgrade(guildId, prorationDate),
    invalidate: () => invalidateGuild(guildId),
    reload: load,
    adopt: (fresh) => (view = fresh),
    onState: (next) => (flow = next),
    openTransferTo: () => (view?.transfer?.state === 'offered' ? view.transfer.to_name : null),
  });
  $effect(() => () => upgradeFlow.dispose());

  /** `upgrade=1`: open the Move to Raid guild confirm once, when the read draws a Move button. */
  $effect(() => {
    if (view === null || !shouldOpenUpgrade(preselect.upgrade, view, viewState, upgradeOpened)) return;
    upgradeOpened = true;
    void upgradeFlow.open();
  });

  let lastFocusHandled = 0;

  /** After the confirm closes onto a reloaded plan: the plan line's button, else the status line;
   *  when the reload failed, the load error's retry button. Once per request. */
  $effect(() => {
    if (!focusIsPending(flow.focusRequest, lastFocusHandled)) return;
    lastFocusHandled = flow.focusRequest;
    const target = flow.focusTarget;
    void tick().then(() => {
      const found =
        target === 'failure'
          ? sectionEl?.querySelector<HTMLElement>('[data-testid="guild-billing-load-error"] button')
          : sectionEl?.querySelector<HTMLElement>('[data-testid="guild-move-raid-guild"]');
      (found ?? statusEl)?.focus();
    });
  });
</script>

<Card
  id="billing"
  class="scroll-mt-20"
  glyph="crown"
  label={view?.billing ? settingsWords.guildPlan : undefined}
  title={view?.billing ? copy.chooser.tierName(view.billing.tier) : settingsWords.guildPlan}
  body="none"
  testid="guild-settings-billing"
>
  <div bind:this={sectionEl}>
    <div
      bind:this={statusEl}
      role="status"
      tabindex="-1"
      class="text-muted border-line-soft border-t px-5 py-3 text-[13px] empty:hidden"
    >
      {#if flow.notice === 'unconfirmed'}
        <p data-testid="guild-upgrade-unconfirmed">{copy.upgrade.unconfirmed}</p>
      {:else if flow.notice === 'unconfirmed-stale'}
        <p data-testid="guild-upgrade-unconfirmed">{copy.upgrade.unconfirmedStale}</p>
      {:else if flow.notice === 'moved'}
        <p data-testid="guild-upgrade-done">{copy.upgrade.done}</p>
      {/if}
    </div>
    {#if failed}
      <div class="border-line-soft border-t px-5 py-4">
        <LoadError message={copy.loadFailed} onRetry={() => void load()} testid="guild-billing-load-error" />
      </div>
    {:else if view === null}
      <div
        class="border-line-soft border-t px-5 py-4 {GUILD_BILLING_LOADING.minHeight}"
        style={storedReserve === null ? undefined : `min-height: ${storedReserve}px`}
        data-testid="guild-billing-reserve"
      >
        <Skeleton lines={GUILD_BILLING_LOADING.lines} testid="guild-billing-skeleton" />
      </div>
    {:else}
      <!-- The ready blocks, measured once settled so the next visit reserves this height. -->
      <div class="flex flex-col" bind:clientHeight={readyHeight} data-testid="guild-billing-ready">
        {#if view.billing !== null && view.transfer !== null && offeredToReader(view)}
          <div class="border-line-soft border-t px-5 py-4">
            <GuildTransferOffer {guildId} plan={view.billing} transfer={view.transfer} onChanged={reload} />
          </div>
        {/if}
        {#if view.billing !== null && viewState.kind !== 'none'}
          <GuildPlanLine
            {guildId}
            plan={view.billing}
            lineState={viewState}
            {payer}
            showMove={movesFromPlanLine(view, viewState)}
            expanded={flow.confirming}
            onMove={() => void upgradeFlow.open()}
          />
        {/if}
        <div data-testid="guild-billing-whopays-slot">
          {#if view.billing !== null && view.can.includes('transfer')}
            <GuildWhoPays
              {guildId}
              plan={view.billing}
              transfer={view.transfer}
              candidates={view.transfer_candidates}
              onChanged={reload}
            />
          {/if}
        </div>
        <div class="border-line-soft flex flex-col gap-5 border-t px-5 py-4 {hasLowerBlocks ? '' : 'hidden'}">
          {#if showSizeLine(view)}
            <GuildSizeLine
              count={view.verified_count}
              {payer}
              expanded={flow.confirming}
              onMove={() => void upgradeFlow.open()}
            />
          {/if}
          {#if flow.confirming}
            <GuildUpgradeConfirm
              plan={view.billing}
              preview={flow.preview}
              committing={flow.committing}
              error={flow.error}
              onCommit={() => void upgradeFlow.commit()}
              onKeep={upgradeFlow.close}
            />
          {/if}
          {#if showsChooser(viewState) && payer}
            <GuildPlanChooser
              {guildId}
              verifiedCount={view.verified_count}
              trialUsed={view.trial_used}
              lapsed={viewState.kind === 'lapsed'}
              {now}
              {...chooserPreselect(preselect)}
            />
          {/if}
          <div data-testid="guild-billing-seats-slot">
            <GuildSeats
              {guildId}
              seats={view.seats}
              total={planSeatCount(view)}
              manage={view.can.includes('seats')}
              {now}
              onChanged={reload}
            />
          </div>
          {#if showsReadOnlyLine(view, payer)}
            <p class="text-muted text-[13px]" data-testid="guild-settings-master-only">
              {copy.readOnly}
            </p>
            {#if showsBillingHandoff(view, viewState)}
              <BillingHandoff {path} {guildName} verifiedCount={view.verified_count} />
            {/if}
          {/if}
        </div>
      </div>
    {/if}
  </div>
</Card>
