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
    onshare,
    canShare = true,
  }: {
    /** The request the page would send now. Null while there is no character. */
    request: SimRequest | null;
    disabled: boolean;
    onvalidate: (json: string) => Promise<RequestValidation>;
    onapply: (request: SimRequest) => void;
    onrun: (request: SimRequest) => void;
    /** Returns the share URL, or null when the request is past the URL budget. Only called
     *  (and the Share affordance only rendered) when `canShare` is true. */
    onshare?: (request: SimRequest) => string | null;
    /**
     * Whether this page can turn a request into a shareable link at all. Design 8's Share
     * assumes the reader can round-trip `?req=` back into this same page -- true for `/sim`
     * (the only reader today), so this defaults to `true` and that call site is unaffected.
     * The four bulk/weights tool pages have no such reader yet, so they pass `false` and the
     * whole Share affordance (button, result, error) disappears rather than the page
     * inventing a message for a feature it does not have.
     */
    canShare?: boolean;
  } = $props();

  // Tracks the page's own request until the player's first edit, then stops: the drawer
  // mounts as soon as a character loads (H, final whole-branch review, finding 2), so
  // seeding once, at mount, carried whatever the *arrival* request happened to be --
  // change the fight style after opening and the textarea, and a Share or Apply from it,
  // went stale with nothing on screen saying so. Re-seeding on every `request` change
  // instead -- right up until `edited` flips true -- fixes that without reintroducing the
  // original problem seed-once solved: once the player has typed, the textarea is theirs
  // and this effect stops touching it. `edited` is set only from the textarea's own real
  // `input` event, never from this effect's own write to `text`, so a programmatic
  // re-seed can never be mistaken for one.
  let text = $state('');
  let edited = $state(false);
  let check = $state<RequestCheck | null>(null);
  let shared = $state('');
  let shareError = $state('');

  $effect(() => {
    if (!edited && request !== null) text = formatRequest(request);
  });

  /** Back to tracking the page's own current request, and clears whatever the last check
   *  said. */
  function reset(): void {
    edited = false;
    if (request !== null) text = formatRequest(request);
    check = null;
    shared = '';
    shareError = '';
  }

  async function verify(): Promise<SimRequest | null> {
    shared = '';
    shareError = '';
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

  async function share(): Promise<void> {
    if (onshare === undefined) return;
    const ok = await verify();
    if (ok === null) return;
    const url = onshare(ok);
    if (url === null) {
      shareError = simCopy.requestShareTooLong;
      return;
    }
    shared = url;
    try {
      await navigator.clipboard.writeText(url);
    } catch {
      /* the field below holds it; a browser that refuses the clipboard is not an error */
    }
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
    <label class="sr-only" for="sim-request-json">{simCopy.requestDrawer}</label>
    <textarea
      id="sim-request-json"
      class="border-line-warm rounded-control bg-card-top text-text h-72 w-full border p-3 font-mono text-[12px]"
      spellcheck="false"
      bind:value={text}
      oninput={() => (edited = true)}
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
      {#if canShare}
        <button type="button" class={button} onclick={() => void share()} data-testid="sim-request-share">
          {simCopy.requestShare}
        </button>
      {/if}
      <button type="button" class={button} onclick={reset} data-testid="sim-request-reset">
        {simCopy.requestReset}
      </button>
    </div>
    <p class="text-muted text-[12px]">{simCopy.requestApplyNote}</p>

    {#if canShare && shareError !== ''}
      <p role="alert" class="text-strong text-[13px]" data-testid="sim-request-share-error">{shareError}</p>
    {:else if canShare && shared !== ''}
      <label class="sr-only" for="sim-request-share-link">{simCopy.requestShare}</label>
      <input
        id="sim-request-share-link"
        type="text"
        readonly
        value={shared}
        class="border-line-warm rounded-control bg-raised text-text h-11 w-full border px-3 text-[13px]"
        data-testid="sim-request-share-link"
        onclick={(event) => event.currentTarget.select()}
      />
    {/if}

    <!-- The buff list as plain text, so a test (and a person) can read what is actually
         going to the engine without parsing the textarea. -->
    <p class="text-muted font-mono text-[11px]" data-testid="sim-request-buffs">
      {(request?.character.buffs ?? []).join(' ')}
    </p>
  </div>
</details>
