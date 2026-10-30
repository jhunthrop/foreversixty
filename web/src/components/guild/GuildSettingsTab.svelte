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
  import type { GuildHome, GuildViewerRole } from '../../lib/guild/api';
  import { guildHomeCopy } from '../../lib/guild/copy';
  import { button } from '../../lib/ui/button';
  import GuildCrestSettings from './GuildCrestSettings.svelte';
  import GuildSettings from '../GuildSettings.svelte';

  let {
    path,
    home,
    role,
    isClaimantAccount,
    contestBusy,
    contestError,
    showContestConfirm,
    onContestStart,
    onContestCancel,
    onContestConfirm,
    onCrestChanged,
    officerTools,
    onPlanRefused,
  }: {
    path: CharacterPath;
    home: GuildHome;
    role: GuildViewerRole;
    isClaimantAccount: boolean;
    contestBusy: boolean;
    contestError: string;
    showContestConfirm: boolean;
    onContestStart: () => void;
    onContestCancel: () => void;
    onContestConfirm: () => void;
    onCrestChanged: () => void;
    /** Officer tools gate §3.1, for the Discord section's read-only state (raids spec 3.8). */
    officerTools?: boolean;
    onPlanRefused?: () => void;
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

<!-- Guild crest round (docs/contracts/2026-10-05-guild-crest-api.md): "officer only" -- this tab
     already only ever mounts for `role === 'officer' | 'moderator'` (Guild.svelte's own tab-content
     gate), so the narrower check here is what actually excludes a moderator, who gets no crest row in
     the web lane even though the API itself lets one remove a crest. The row sits at the foot of the
     Records card. -->
{#snippet crestRow()}
  <GuildCrestSettings guildId={home.guild.id} guild={home.guild} {onCrestChanged} />
{/snippet}

<section class="flex flex-col gap-6" data-testid="guild-settings-tab">
  <GuildSettings
    embedded
    {path}
    {officerTools}
    {onPlanRefused}
    crest={role === 'officer' ? crestRow : undefined}
  />

  {#if canContestHere}
    <div
      class="border-line-soft flex flex-col gap-3 border-t pt-6"
      data-testid="guild-settings-claim-contest"
    >
      <button
        class="{button('secondary', 'sm')} w-fit"
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
              class={button('secondary', 'sm')}
              onclick={onContestConfirm}
              disabled={contestBusy}
              data-testid="guild-home-contest-confirm-button"
            >
              {guildHomeCopy.contestConfirmButton}
            </button>
            <button class={button('ghost', 'sm')} onclick={onContestCancel} disabled={contestBusy}>
              {guildHomeCopy.cancel}
            </button>
          </div>
          {#if contestError !== ''}<p class="text-[13px]" role="alert">{contestError}</p>{/if}
        </div>
      {/if}
    </div>
  {/if}
</section>
