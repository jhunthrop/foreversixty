<!-- web/src/components/account/AccountDeleteDialog.svelte -->
<!-- The account deletion confirm (DELETE /v1/me; mock ws_prims `Dialog`, web-account-delete): the
     overlay sheet over a scrim, a 560 px box centred from 768, a full-width bottom sheet below. It
     says what is deleted and what is kept, then asks for the BattleTag (or "delete") typed out; the
     danger button stays off until what is typed matches, by the server's own rule. Focus goes to
     the title on open and stays inside; Escape, the scrim and Cancel close it (not while the call
     runs), and the owner gives focus back to the control that opened it. -->
<script lang="ts">
  import { trapTab } from '../../lib/ui/focus-trap';
  import { tick } from 'svelte';
  import { accountPageCopy } from '../../lib/account/account-page-copy';
  import { confirmMatches, confirmWordFor, lineParts } from '../../lib/account/delete-confirm';
  import { button } from '../../lib/ui/button';
  import { FIELD_CONTROL, FIELD_FOCUS, FIELD_HEIGHT, fieldEdge } from '../../lib/ui/field';
  import { LINK } from '../../lib/ui/primitives';
  import { REGION } from '../../lib/ui/regions';
  import Button from '../ui/Button.svelte';
  import Overlay from '../ui/Overlay.svelte';
  import Spinner from '../ui/Spinner.svelte';

  let {
    battletag,
    busy,
    mismatch,
    refusal,
    onconfirm,
    onclose,
  }: {
    battletag: string | null | undefined;
    /** The call is running: the button shows the spinner and nothing closes the dialog. */
    busy: boolean;
    /** The server refused what was typed (400): the line under the field. */
    mismatch: boolean;
    /** The API's own message for any other refusal (409, 429, 503), shown verbatim. */
    refusal: string;
    onconfirm: (typed: string) => void;
    onclose: () => void;
  } = $props();

  const copy = accountPageCopy.deleteDialog;
  const ids = {
    title: 'account-delete-title',
    field: 'account-delete-confirm',
    error: 'account-delete-confirm-error',
  };

  let typed = $state('');
  let panel = $state<HTMLElement | undefined>(undefined);
  let heading = $state<HTMLElement | undefined>(undefined);
  let field = $state<HTMLInputElement | undefined>(undefined);
  let alert = $state<HTMLElement | undefined>(undefined);

  const word = $derived(confirmWordFor(battletag));
  const labelParts = $derived(copy.confirmLabel(word).split(word));
  const ready = $derived(confirmMatches(typed, battletag));

  $effect(() => {
    void tick().then(() => heading?.focus());
  });

  // A refused confirm sends focus back to the field, so its error line is read with it.
  $effect(() => {
    if (mismatch) void tick().then(() => field?.focus());
  });

  // A refusal sits under the field, which can be below the fold of a phone sheet's body: scroll it in.
  $effect(() => {
    if (refusal !== '') void tick().then(() => alert?.scrollIntoView({ block: 'nearest' }));
  });

  function close(): void {
    if (!busy) onclose();
  }

  function onkeydown(event: KeyboardEvent): void {
    if (event.key === 'Escape') {
      event.preventDefault();
      close();
      return;
    }
    if (panel !== undefined) trapTab(event, panel, heading);
  }

  function submit(event: SubmitEvent): void {
    event.preventDefault();
    if (ready && !busy) onconfirm(typed);
  }
</script>

<div class="account-delete-scrim" aria-hidden="true" onclick={close}></div>
<div
  bind:this={panel}
  class="account-delete-dialog"
  role="dialog"
  aria-modal="true"
  aria-labelledby={ids.title}
  tabindex="-1"
  data-region={REGION.dialog}
  data-testid="account-delete-dialog"
  {onkeydown}
>
  <Overlay kind="sheet" role="none" class="account-delete-box">
    <form class="contents" onsubmit={submit} novalidate>
      <div class="flex flex-col gap-1 px-5 pt-5 pb-1">
        <h2
          bind:this={heading}
          id={ids.title}
          class="text-strong m-0 text-[20px] leading-6 font-bold outline-none"
          tabindex="-1"
        >
          {copy.title}
        </h2>
        <p class="m-0 text-[15px] leading-[22px]">{copy.intro}</p>
      </div>
      <div class="account-delete-body">
        <h3 class="account-delete-group text-wipe">{copy.deletedLabel}</h3>
        <ul class="account-delete-list" data-testid="account-delete-deleted">
          {#each copy.deleted as line (line.text)}
            {@const parts = lineParts(line)}
            <li>
              {parts.before}{#if parts.link}<a class={LINK} href={parts.link.href}>{parts.link.label}</a
                >{/if}{parts.after}
            </li>
          {/each}
        </ul>
        <h3 class="account-delete-group text-muted">{copy.keptLabel}</h3>
        <ul class="account-delete-list" data-testid="account-delete-kept">
          {#each copy.kept as line (line.text)}
            <li>{line.text}</li>
          {/each}
        </ul>
        <div class="mt-5 flex flex-col gap-[6px]">
          <label for={ids.field} class="text-[15px] leading-[22px]"
            >{labelParts[0]}<b class="text-strong">{word}</b>{labelParts[1]}</label
          >
          <input
            bind:this={field}
            bind:value={typed}
            id={ids.field}
            type="text"
            autocomplete="off"
            spellcheck="false"
            autocapitalize="off"
            disabled={busy}
            class={[FIELD_CONTROL, FIELD_HEIGHT, FIELD_FOCUS, fieldEdge(mismatch)]}
            aria-invalid={mismatch ? 'true' : undefined}
            aria-describedby={mismatch ? ids.error : undefined}
            data-testid="account-delete-confirm"
          />
          {#if mismatch}<span
              id={ids.error}
              class="text-wipe text-[13px] leading-5"
              data-testid="account-delete-mismatch">{copy.mismatch(word)}</span
            >{/if}
        </div>
        {#if refusal !== ''}<p
            bind:this={alert}
            class="text-wipe m-0 mt-3 text-[14px] leading-5"
            role="alert"
            data-testid="account-delete-refusal"
          >
            {refusal}
          </p>{/if}
      </div>
      <div class="account-delete-actions ui-overlay-rule-t">
        <!-- A busy button keeps its look (aria-busy), the spinner in front of the busy word. -->
        <button
          type="submit"
          class={button('danger')}
          disabled={!ready || busy}
          aria-busy={busy ? 'true' : undefined}
          data-testid="account-delete-submit"
          >{#if busy}<Spinner />{copy.busy}{:else}{copy.confirm}{/if}</button
        >
        <Button onclick={close} disabled={busy} testid="account-delete-cancel">{copy.cancel}</Button>
      </div>
    </form>
  </Overlay>
</div>

<style>
  .account-delete-scrim {
    position: fixed;
    inset: 0;
    z-index: 60;
    background: color-mix(in srgb, var(--color-bg) 80%, transparent);
  }
  .account-delete-dialog {
    position: fixed;
    z-index: 61;
    right: 0;
    bottom: 0;
    left: 0;
    display: flex;
    flex-direction: column;
    max-height: 88vh;
    outline: none;
  }
  .account-delete-dialog > :global(.account-delete-box) {
    display: flex;
    flex: 1;
    flex-direction: column;
    min-height: 0;
  }
  .account-delete-body {
    display: flex;
    flex: 1;
    flex-direction: column;
    min-height: 0;
    overflow: auto;
    overscroll-behavior: contain;
    padding: 8px 20px 20px;
  }
  .account-delete-group {
    margin: 16px 0 6px;
    font-size: 12px;
    line-height: 16px;
    font-weight: 700;
    letter-spacing: 0;
  }
  .account-delete-list {
    display: flex;
    flex-direction: column;
    gap: 6px;
    margin: 0;
    padding-left: 20px;
    list-style: disc;
    font-size: 15px;
    line-height: 22px;
  }
  .account-delete-list li::marker {
    color: var(--color-muted);
  }
  .account-delete-actions {
    display: flex;
    gap: 12px;
    padding: 16px 20px calc(16px + env(safe-area-inset-bottom));
  }
  @media (max-width: 767px) {
    .account-delete-actions > :global(*) {
      flex: 1;
    }
  }
  /* From 768 the sheet is a box centred in the viewport, so it takes the overlay's full edge and
     radius back (the menu kind's box); below, the bottom sheet's open bottom edge stays. */
  @media (min-width: 768px) {
    .account-delete-dialog {
      top: 50%;
      right: auto;
      bottom: auto;
      left: 50%;
      width: 560px;
      max-width: calc(100% - 48px);
      max-height: calc(100vh - 64px);
      transform: translate(-50%, -50%);
    }
    .account-delete-dialog > :global(.account-delete-box) {
      border-bottom-width: 1px;
      border-radius: var(--radius-plate);
    }
    .account-delete-actions {
      padding-bottom: 16px;
    }
  }
</style>
