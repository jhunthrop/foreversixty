<!-- web/src/components/sim/RequestDrawer.svelte -->
<!-- Design 8. The exact JSON the run will send, editable, with the engine's own Validate
     behind it and its errors beside the fields they name.

     Two exits, because they are two different things: Apply rebuilds the page from the
     request and can only carry what the page has controls for; Run sends the text exactly
     as typed, which is what makes this the escape hatch rather than a second settings bar.

     A pure render, per the lane's standing convention: `checkRequest` (request-json.ts)
     owns the parse-then-ask-the-engine question and its three-way answer; this component
     only seeds a textarea, calls that function on a click, and renders whichever `status`
     comes back.

     Drawn as a disclosure (web layout spec 2026-11-05 §3A.2): a native details element so the
     summary keeps its own testid, the body open on the ground, the field recipe on the text and
     ui/Button for the actions. No box. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import type { RequestValidation } from '../../lib/sim/engine';
  import { checkRequest, formatRequest, type RequestCheck } from '../../lib/sim/request-json';
  import type { SimRequest } from '../../lib/sim/types';
  import { FIELD_FOCUS, fieldEdge } from '../../lib/ui/field';
  import { uiCopy } from '../../lib/ui/copy';
  import { CopyState } from '../../lib/ui/copy-state.svelte';
  import Button from '../ui/Button.svelte';
  import Disclosure from '../ui/Disclosure.svelte';

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
    // The field below holds it too: a browser that refuses the clipboard is not an error.
    await clip.copy(url);
  }

  const clip = new CopyState();
  $effect(() => () => clip.dispose());

  /** The field recipe (lib/ui/field.ts) on the JSON and the share link: the page ground, one edge. */
  const FIELD = `rounded-field bg-bg text-strong w-full border ${fieldEdge(false)} ${FIELD_FOCUS}`;
</script>

<div class="mx-[18px] md:mx-0">
  <Disclosure
    summary={simCopy.requestDrawerSummary}
    meta={simCopy.requestDrawerHint}
    testid="sim-request-drawer"
    summaryTestid="sim-request-summary"
  >
    <div class="flex flex-col gap-3">
      <p class="text-small text-muted">{simCopy.requestNote}</p>
      <label class="sr-only" for="sim-request-json">{simCopy.requestDrawer}</label>
      <textarea
        id="sim-request-json"
        class={`${FIELD} h-72 p-3 font-mono text-[12px]`}
        spellcheck="false"
        bind:value={text}
        oninput={() => (edited = true)}
        data-testid="sim-request-json"></textarea>

      {#if check !== null && (check.status === 'parse-error' || check.status === 'validate-error')}
        <p role="alert" class="text-strong text-[14px] leading-5" data-testid="sim-request-errors">
          {check.message}
        </p>
      {:else if check !== null && check.status === 'invalid'}
        <ul class="flex flex-col gap-1 text-[14px] leading-5" role="alert" data-testid="sim-request-errors">
          {#each check.errors as row, index (`${row.field}-${index}`)}
            <li data-testid={`sim-request-error-${row.field}`}>
              <span class="text-strong font-mono">{row.field}</span>
              <span class="text-muted ml-2">{row.message}</span>
            </li>
          {/each}
        </ul>
      {:else if check !== null && check.status === 'valid'}
        <p class="text-muted text-[14px] leading-5" data-testid="sim-request-valid">{simCopy.requestValid}</p>
      {/if}

      <div class="flex flex-wrap items-center gap-2">
        <Button size="sm" {disabled} onclick={() => void apply()} testid="sim-request-apply"
          >{simCopy.requestApply}</Button
        >
        <Button size="sm" {disabled} onclick={() => void run()} testid="sim-request-run"
          >{simCopy.requestRun}</Button
        >
        {#if canShare}
          <Button size="sm" onclick={() => void share()} testid="sim-request-share"
            >{clip.copied() ? uiCopy.copied : simCopy.requestShare}</Button
          >
        {/if}
        <Button size="sm" onclick={reset} testid="sim-request-reset">{simCopy.requestReset}</Button>
      </div>
      <p class="text-small text-muted">{simCopy.requestApplyNote}</p>

      {#if canShare && shareError !== ''}
        <p role="alert" class="text-strong text-[14px] leading-5" data-testid="sim-request-share-error">
          {shareError}
        </p>
      {:else if canShare && shared !== ''}
        <label class="sr-only" for="sim-request-share-link">{simCopy.requestShare}</label>
        <input
          id="sim-request-share-link"
          type="text"
          readonly
          value={shared}
          class={`${FIELD} h-11 px-3 text-[14px] md:h-10`}
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
  </Disclosure>
</div>
