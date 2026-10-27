<!-- web/src/components/guild/billing/GuildSizeLine.svelte -->
<!-- The soft size line (spec 3.2.3): a Guild-tier guild over 25 verified members. A gold left rule;
     payers get the one Move to Raid guild button beside it. Nothing switches off. -->
<script lang="ts">
  import { guildBillingCopy as copy } from '../../../lib/billing/guild-billing-copy';
  import { BILLING_MOVE_BUTTON, OPENER_PRESSED } from '../../../lib/guild/billing-styles';

  let {
    count,
    payer,
    expanded = false,
    onMove,
  }: {
    count: number;
    payer: boolean;
    /** The confirm block it opens is showing: the opener draws pressed. */
    expanded?: boolean;
    onMove: () => void;
  } = $props();
</script>

<div
  class="flex flex-col gap-2 md:flex-row md:items-center md:justify-between md:gap-4"
  data-testid="guild-size-line"
>
  <p class="border-gold border-l-2 pl-3.5 text-[14px]">{copy.size.line(count)}</p>
  {#if payer}
    <button
      type="button"
      class="{BILLING_MOVE_BUTTON} {expanded ? OPENER_PRESSED : ''}"
      aria-expanded={expanded}
      onclick={onMove}
      data-testid="guild-move-raid-guild"
    >
      {copy.size.move}
    </button>
  {/if}
</div>
