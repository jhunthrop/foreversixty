<!-- web/src/components/sim/RequestDrawer.svelte -->
<!-- Design 8. The exact JSON the run will send, editable, with the engine's own Validate
     behind it and its errors beside the fields they name.

     Two exits, because they are two different things: Apply rebuilds the page from the
     request and can only carry what the page has controls for; Run sends the text exactly
     as typed, which is what makes this the escape hatch rather than a second settings bar.

     A pure render, per the lane's standing convention: `checkRequest` (request-json.ts)
     owns the parse-then-ask-the-engine question and its three-way answer; this component
     only seeds a textarea, calls that function on a click, and renders whichever `status`
     comes back. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import type { RequestValidation } from '../../lib/sim/engine';
  import { checkRequest, formatRequest, type RequestCheck } from '../../lib/sim/request-json';
  import type { SimRequest } from '../../lib/sim/types';

  let {
    request,
    disabled,
    onvalidate,
    onapply,
    onrun,
  }: {
    /** The request the page would send now. Null while there is no character. */
    request: SimRequest | null;
    disabled: boolean;
    onvalidate: (json: string) => Promise<RequestValidation>;
    onapply: (request: SimRequest) => void;
    onrun: (request: SimRequest) => void;
  } = $props();

  // Seeded from the page's own request and then owned by the player: re-seeding it on
  // every settings change would throw away what they were typing.
  let text = $state('');
  let seeded = $state(false);
  let check = $state<RequestCheck | null>(null);

  $effect(() => {
    if (!seeded && request !== null) {
      text = formatRequest(request);
      seeded = true;
    }
  });

  /** Back to the page's own current request, and clears whatever the last check said. */
  function reset(): void {
    if (request !== null) text = formatRequest(request);
    check = null;
  }

  async function verify(): Promise<SimRequest | null> {
    const result = await checkRequest(text, onvalidate);
    check = result;
    return result.status === 'valid' ? result.request : null;
  }

  async function apply(): Promise<void> {
    const ok = await verify();
    if (ok !== null) onapply(ok);
  }

  async function run(): Promise<void> {
    const ok = await verify();
    if (ok !== null) onrun(ok);
  }

  const button =
    'border-line-warm rounded-control text-nav label min-h-11 border px-4 disabled:opacity-50 md:min-h-9';
</script>

<details
  class="border-line bg-raised rounded-panel mx-[18px] border md:mx-0"
  data-testid="sim-request-drawer"
>
  <summary class="label text-nav flex min-h-11 cursor-pointer items-center px-4 md:min-h-9">
    {simCopy.requestDrawer}
  </summary>
  <div class="flex flex-col gap-3 p-4 pt-0">
    <p class="text-muted text-[12px]">{simCopy.requestNote}</p>
    <textarea
      class="border-line-warm rounded-control bg-card-top text-text h-72 w-full border p-3 font-mono text-[12px]"
      spellcheck="false"
      bind:value={text}
      data-testid="sim-request-json"></textarea>

    {#if check !== null && (check.status === 'parse-error' || check.status === 'validate-error')}
      <p role="alert" class="text-strong text-[13px]" data-testid="sim-request-errors">{check.message}</p>
    {:else if check !== null && check.status === 'invalid'}
      <ul class="flex flex-col gap-1 text-[13px]" role="alert" data-testid="sim-request-errors">
        {#each check.errors as row, index (`${row.field}-${index}`)}
          <li data-testid={`sim-request-error-${row.field}`}>
            <span class="text-strong font-mono">{row.field}</span>
            <span class="text-muted ml-2">{row.message}</span>
          </li>
        {/each}
      </ul>
    {:else if check !== null && check.status === 'valid'}
      <p class="text-muted text-[13px]" data-testid="sim-request-valid">{simCopy.requestValid}</p>
    {/if}

    <div class="flex flex-wrap items-center gap-3">
      <button
        type="button"
        class={button}
        {disabled}
        onclick={() => void apply()}
        data-testid="sim-request-apply"
      >
        {simCopy.requestApply}
      </button>
      <button
        type="button"
        class={button}
        {disabled}
        onclick={() => void run()}
        data-testid="sim-request-run"
      >
        {simCopy.requestRun}
      </button>
      <button type="button" class={button} onclick={reset} data-testid="sim-request-reset">
        {simCopy.requestReset}
      </button>
    </div>
    <p class="text-muted text-[12px]">{simCopy.requestApplyNote}</p>

    <!-- The buff list as plain text, so a test (and a person) can read what is actually
         going to the engine without parsing the textarea. -->
    <p class="text-muted font-mono text-[11px]" data-testid="sim-request-buffs">
      {(request?.character.buffs ?? []).join(' ')}
    </p>
  </div>
</details>
