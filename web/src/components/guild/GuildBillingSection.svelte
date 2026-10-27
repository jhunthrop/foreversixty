<!-- web/src/components/guild/GuildBillingSection.svelte -->
<!-- The Billing section of guild settings (launch pricing spec 3.2): it reads the guild billing view
     itself, reserves its ready height with Skeleton lines while that loads, and draws the blocks in
     order: plan line, size line, chooser, the seats and Who pays mount points, Records. -->
<script lang="ts">
  import { BillingApiError } from '../../lib/billing/api';
  import { getGuildBilling, previewUpgrade, upgrade } from '../../lib/billing/guild-billing-api';
  import type { GuildBillingSettings, UpgradePreview } from '../../lib/billing/contract';
  import { guildBillingCopy as copy } from '../../lib/billing/guild-billing-copy';
  import {
    billingState,
    canBill,
    movesFromPlanLine,
    showSizeLine,
    showsChooser,
  } from '../../lib/guild/billing-state';
  import { guildRosterModelCopy } from '../../lib/guild/copy';
  import { GUILD_BILLING_LOADING } from '../../lib/guild/layout';
  import { minuteClock } from '../../lib/guild/minute-clock.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import GuildPlanChooser from './billing/GuildPlanChooser.svelte';
  import GuildPlanLine from './billing/GuildPlanLine.svelte';
  import GuildRecords from './billing/GuildRecords.svelte';
  import GuildSeats from './billing/GuildSeats.svelte';
  import GuildSizeLine from './billing/GuildSizeLine.svelte';
  import GuildTransferOffer from './billing/GuildTransferOffer.svelte';
  import GuildUpgradeConfirm from './billing/GuildUpgradeConfirm.svelte';
  import GuildWhoPays from './billing/GuildWhoPays.svelte';

  let { guildId, now: fixedNow }: { guildId: number; now?: Date } = $props();

  /** The minute tick: a seat lock that expires while the page is open turns over without a reload. */
  const clock = minuteClock(() => fixedNow);
  const now = $derived(clock.now);

  let view = $state<GuildBillingSettings | null>(null);
  let failed = $state(false);
  let confirming = $state(false);
  let preview = $state<UpgradePreview | null>(null);
  let committing = $state(false);
  let upgradeError = $state('');
  let moved = $state(false);

  const viewState = $derived(billingState(view, now));
  const payer = $derived(view !== null && canBill(view));

  async function load(): Promise<void> {
    failed = false;
    try {
      view = await getGuildBilling(guildId);
    } catch {
      failed = true;
    }
  }

  $effect(() => {
    void load();
  });

  let previewToken = 0;

  async function openConfirm(): Promise<void> {
    const token = ++previewToken;
    confirming = true;
    preview = null;
    upgradeError = '';
    try {
      const fresh = await previewUpgrade(guildId);
      if (token === previewToken) preview = fresh;
    } catch {
      if (token === previewToken) upgradeError = copy.upgrade.failed;
    }
  }

  function closeConfirm(): void {
    previewToken += 1;
    confirming = false;
    upgradeError = '';
  }

  /** A failure that is not a clean decline may have charged: show the plan as it now stands. */
  async function failCommit(thrown: unknown): Promise<void> {
    if (thrown instanceof BillingApiError && thrown.code === 'preview_expired') {
      await openConfirm();
      upgradeError = copy.upgrade.previewExpired;
      return;
    }
    if (thrown instanceof BillingApiError && thrown.status === 409) {
      upgradeError = copy.upgrade.failed;
      return;
    }
    confirming = false;
    await load();
  }

  async function commit(): Promise<void> {
    if (preview === null) return;
    committing = true;
    upgradeError = '';
    try {
      view = await upgrade(guildId, preview.proration_date);
      confirming = false;
      moved = true;
    } catch (thrown) {
      await failCommit(thrown);
    } finally {
      committing = false;
    }
  }
</script>

<section id="billing" class="flex scroll-mt-20 flex-col gap-5" data-testid="guild-settings-billing">
  <h2 class="section-title text-[18px]">Billing</h2>
  {#if failed}
    <LoadError message={copy.loadFailed} onRetry={() => void load()} testid="guild-billing-load-error" />
  {:else if view === null}
    <div class={GUILD_BILLING_LOADING.minHeight}>
      <Skeleton lines={GUILD_BILLING_LOADING.lines} testid="guild-billing-skeleton" />
    </div>
  {:else}
    {#if view.billing !== null && view.transfer?.to_you === true && view.transfer.state === 'offered'}
      <GuildTransferOffer {guildId} plan={view.billing} transfer={view.transfer} onChanged={load} />
    {/if}
    {#if view.billing !== null && viewState.kind !== 'none'}
      <GuildPlanLine
        {guildId}
        plan={view.billing}
        lineState={viewState}
        {payer}
        showMove={movesFromPlanLine(view, viewState)}
        onMove={() => void openConfirm()}
      />
    {/if}
    {#if moved}
      <p class="text-muted text-[13px]" role="status" data-testid="guild-upgrade-done">{copy.upgrade.done}</p>
    {/if}
    {#if showSizeLine(view)}
      <GuildSizeLine count={view.verified_count} {payer} onMove={() => void openConfirm()} />
    {/if}
    {#if confirming}
      <GuildUpgradeConfirm
        {preview}
        {committing}
        error={upgradeError}
        onCommit={() => void commit()}
        onKeep={closeConfirm}
      />
    {/if}
    {#if showsChooser(viewState) && payer}
      <GuildPlanChooser
        {guildId}
        verifiedCount={view.verified_count}
        trialUsed={view.trial_used}
        lapsed={viewState.kind === 'lapsed'}
        {now}
      />
    {/if}
    <div data-testid="guild-billing-seats-slot">
      <GuildSeats {guildId} seats={view.seats} manage={view.can.includes('seats')} {now} onChanged={load} />
    </div>
    <div data-testid="guild-billing-whopays-slot">
      {#if view.billing !== null && view.can.includes('transfer')}
        <GuildWhoPays
          {guildId}
          plan={view.billing}
          transfer={view.transfer}
          candidates={view.transfer_candidates}
          onChanged={load}
        />
      {/if}
    </div>
    <GuildRecords {guildId} />
    {#if !payer}
      <p class="text-muted text-[13px]" data-testid="guild-settings-master-only">
        {guildRosterModelCopy.guildMasterOnly}
      </p>
    {/if}
  {/if}
</section>
