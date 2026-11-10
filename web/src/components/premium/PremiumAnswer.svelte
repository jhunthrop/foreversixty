<!-- web/src/components/premium/PremiumAnswer.svelte -->
<!-- /premium's head answer, signed in (premium spec 2026-11-08 section 4.1), compact: `Plan` and the
     plan's name as the figure, then one line with its date and price or source, the Founder mark
     and `See your plan ›`. -->
<script lang="ts">
  import { accountPageCopy } from '../../lib/account/account-page-copy';
  import { accountPlanCopy } from '../../lib/billing/account-plan-copy';
  import type { PremiumAnswerView } from '../../lib/billing/premium-page';
  import { premiumSourceCopy } from '../../lib/billing/premium-source-copy';
  import { LINK } from '../../lib/ui/primitives';
  import Figure from '../ui/Figure.svelte';
  import Tag from '../ui/Tag.svelte';

  let { view }: { view: PremiumAnswerView } = $props();
</script>

<div class="premium-answer" data-testid="premium-answer">
  <Figure value={view.name} label={accountPageCopy.planLabel} size="sm" align="end" />
  <span class="premium-answer-line text-small text-muted">
    {#if view.sub !== ''}<span>{view.sub}</span>{/if}
    {#if view.founder}<Tag>{accountPlanCopy.founder.pill}</Tag>{/if}
    <a class={LINK} href="/account#plan" data-testid="premium-answer-link"
      >{premiumSourceCopy.buttons.seePlan} ›</a
    >
  </span>
</div>

<style>
  .premium-answer {
    display: flex;
    flex-direction: column;
    gap: 2px;
  }
  /* One line under the figure: the date or source, the Founder mark and the way to the plan. */
  .premium-answer-line {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    gap: 4px 10px;
  }
  @media (min-width: 900px) {
    .premium-answer {
      align-items: flex-end;
    }
    .premium-answer-line {
      justify-content: flex-end;
    }
  }
</style>
