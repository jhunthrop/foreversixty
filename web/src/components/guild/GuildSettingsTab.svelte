<!-- web/src/components/guild/GuildSettingsTab.svelte -->
<!-- Guild control-centre spec §4.G: Settings absorbs GuildSettings.svelte's full surface,
     restyled into the tab strip rather than a separate route. Every field and copy string
     is GuildSettings.svelte's own, unchanged (tenet 10 -- this tab is a restyle, not a
     rewrite of settings logic).

     Fix round 1 (2026-10-04, item 2): "Contest this claim" moved here from the header
     band, under GuildSettings's own claim-state block, and now excludes the claimant's
     own account -- `isClaimantAccount` is computed once in Guild.svelte (it needs the
     signed-in account's battletag, which this tab does not otherwise fetch) and passed
     down, along with the existing contest state/handlers, unchanged from how they worked
     in the header. -->
<script lang="ts">
  import type { CharacterPath } from '../../lib/characters';
  import type { GuildHome } from '../../lib/guild/api';
  import { guildHomeCopy } from '../../lib/guild/copy';
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import GuildSettings from '../GuildSettings.svelte';

  let {
    path,
    home,
    isClaimantAccount,
    contestBusy,
    contestError,
    showContestConfirm,
    onContestStart,
    onContestCancel,
    onContestConfirm,
  }: {
    path: CharacterPath;
    home: GuildHome;
    isClaimantAccount: boolean;
    contestBusy: boolean;
    contestError: string;
    showContestConfirm: boolean;
    onContestStart: () => void;
    onContestCancel: () => void;
    onContestConfirm: () => void;
  } = $props();

  /** Spec item 2: a verified officer of a different account than the one holding the
   *  claim may contest it; the claimant's own officers never see the control at all --
   *  they already have the officer strip (claim badge, approvals, invite) in the band.
   *  Contest is only ever offered for a settled `claimed` guild here -- an `unclaimed`
   *  guild has nothing to contest, and a `pending` claim's own battletag is not carried
   *  on `GuildHome.claim` to check against, so that case is left to the dedicated claim
   *  page's own confirm/expire flow, unchanged. */
  const canContestHere = $derived(!isClaimantAccount && home.claim.state === 'claimed');
</script>

<section class="flex flex-col gap-6" data-testid="guild-settings-tab">
  <GuildSettings {path} />

  {#if canContestHere}
    <div
      class="border-line-soft flex flex-col gap-3 border-t pt-6"
      data-testid="guild-settings-claim-contest"
    >
      <button
        class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-3"
        onclick={onContestStart}
        disabled={contestBusy}
        data-testid="guild-home-contest-button"
      >
        {guildHomeCopy.contestButton}
      </button>

      {#if !showContestConfirm && contestError !== ''}
        <p class="text-[13px]" role="alert" data-testid="guild-home-contest-error">{contestError}</p>
      {/if}

      {#if showContestConfirm}
        <div class="border-line-soft flex flex-col gap-3 border p-4" data-testid="guild-home-contest-confirm">
          <ul class="flex flex-col gap-1 text-[13px]">
            {#each guildHomeCopy.contestRules as rule (rule)}
              <li>{rule}</li>
            {/each}
          </ul>
          <div class="flex gap-3">
            <button
              class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-3"
              onclick={onContestConfirm}
              disabled={contestBusy}
              data-testid="guild-home-contest-confirm-button"
            >
              {guildHomeCopy.contestConfirmButton}
            </button>
            <button
              class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text px-3"
              onclick={onContestCancel}
              disabled={contestBusy}
            >
              {guildHomeCopy.cancel}
            </button>
          </div>
          {#if contestError !== ''}<p class="text-[13px]" role="alert">{contestError}</p>{/if}
        </div>
      {/if}
    </div>
  {/if}
</section>
