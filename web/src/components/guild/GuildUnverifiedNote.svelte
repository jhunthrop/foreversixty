<!-- web/src/components/guild/GuildUnverifiedNote.svelte -->
<!-- The line under the guild band for a member who is not verified yet (spec 2026-10-27 §6.3).
     With an approver (or an older API that does not say) it is today's sentence; with none it says
     nobody can approve until the guild master claims, and links the member to the claim page to
     send on. -->
<script lang="ts">
  import { guildHomeCopy } from '../../lib/guild/copy';

  let {
    guildName,
    hasApprover,
    claimHref,
  }: {
    guildName: string;
    /** `home.viewer.verify.has_approver`; undefined (an older API) reads as an approver exists. */
    hasApprover: boolean | undefined;
    claimHref: string;
  } = $props();
</script>

<p class="text-muted text-[13px]" data-testid="guild-home-not-verified-note">
  {#if hasApprover === false}
    {guildHomeCopy.notVerifiedNoApprover(guildName)}
    <a class="text-gold underline" href={claimHref} data-testid="guild-home-send-claim-link"
      >{guildHomeCopy.sendGuildMasterLink}</a
    >
  {:else}
    {guildHomeCopy.notVerifiedNote}
  {/if}
</p>
