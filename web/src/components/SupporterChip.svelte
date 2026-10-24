<!-- web/src/components/SupporterChip.svelte -->
<!-- The Supporter mark: one gold text chip beside a name, the same dim warm border on every
     surface (set inline so no inherited colour can brighten it). Public surfaces show this or
     nothing; it is never drawn on a crest and never links anywhere. -->
<script lang="ts">
  import { premiumBlockCopy } from '../lib/billing/copy';
</script>

<span
  class="label text-gold rounded-control shrink-0 border px-2 py-0.5"
  style:border-color="var(--color-line-warm)"
  title={premiumBlockCopy.supporterTooltip}
  data-testid="supporter-chip">{premiumBlockCopy.supporterChip}</span
>
