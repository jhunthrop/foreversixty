<!-- web/src/components/checkout/CheckoutSteps.svelte -->
<!-- The checkout's one Plate (premium spec 2026-11-08 section 6): three Steps, Sign in, Pay with
     Stripe, {Plan} on, marked from the launcher's phase. Back from Stripe a paid plan reads
     `Stripe took the payment.`, a guild in its free trial `Card saved. First charge {date}.`
     (ruling K1). A failure marks its row and holds the row's one action. -->
<script lang="ts" module>
  export type CheckoutPhase = 'redirecting' | 'processing' | 'confirmed' | 'slow' | 'error' | 'refused';
</script>

<script lang="ts">
  import { checkoutCopy } from '../../lib/billing/copy';
  import Button from '../ui/Button.svelte';
  import Plate from '../ui/Plate.svelte';
  import Steps, { type StepsMark } from '../ui/Steps.svelte';

  let {
    phase,
    message,
    planName,
    signedInAs,
    paidLine = checkoutCopy.paid,
    busy = false,
    onRetry,
    onCheckAgain,
  }: {
    phase: CheckoutPhase;
    /** The phase's sentence: the launcher's message (a refusal carries the API's own words). */
    message: string;
    planName: string;
    /** The account's BattleTag, when /v1/me has answered. */
    signedInAs: string | null;
    /** Row 2 after the return: `Stripe took the payment.` or the trial's card line. */
    paidLine?: string;
    busy?: boolean;
    onRetry?: () => void;
    onCheckAgain?: () => void;
  } = $props();

  const steps = checkoutCopy.steps;
  const back = $derived(phase === 'processing' || phase === 'confirmed' || phase === 'slow');
  const items = $derived([
    { key: 'signin', n: 1, mark: 'ok' as StepsMark, title: steps.signIn },
    {
      key: 'pay',
      n: 2,
      mark: (back ? 'ok' : phase === 'redirecting' ? 'busy' : 'miss') as StepsMark,
      title: steps.pay,
    },
    {
      key: 'on',
      n: 3,
      mark: (phase === 'confirmed'
        ? 'ok'
        : phase === 'processing'
          ? 'busy'
          : phase === 'slow'
            ? 'warn'
            : 'todo') as StepsMark,
      title: steps.on(planName),
    },
  ]);
</script>

<Plate label="Checkout" testid="checkout-plate">
  <Steps {items} testid="checkout-steps">
    {#snippet sub(item)}
      {#if item.key === 'signin'}
        {#if signedInAs !== null}{checkoutCopy.signedInAs(signedInAs)}{/if}
      {:else if item.key === 'pay'}
        {#if back}
          <span data-testid="checkout-paid">{paidLine}</span>
        {:else if phase === 'redirecting'}
          <span class="text-text" role="status" data-testid="checkout-message">{message}</span>
          <span class="block">{checkoutCopy.stripeLine}</span>
        {:else}
          <span class="text-wipe" role="alert" data-testid="checkout-message">{message}</span>
        {/if}
      {:else if back}
        <span
          class={phase === 'slow' ? 'text-warn' : 'text-text'}
          role="status"
          data-testid="checkout-message">{message}</span
        >
      {/if}
    {/snippet}
    {#snippet action(item)}
      {#if item.key === 'pay' && phase === 'error' && onRetry}
        <Button tone="primary" disabled={busy} onclick={onRetry} testid="checkout-retry"
          >{checkoutCopy.retry}</Button
        >
      {:else if item.key === 'on' && phase === 'slow' && onCheckAgain}
        <Button tone="primary" disabled={busy} onclick={onCheckAgain} testid="checkout-check-again"
          >{checkoutCopy.checkAgain}</Button
        >
      {/if}
    {/snippet}
  </Steps>
</Plate>
