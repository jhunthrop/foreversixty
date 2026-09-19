<!-- web/src/components/sim/SourceSwitcher.svelte -->
<!-- The four ways in. Order is deliberate: the addon export is first because it needs no
     account and produces the most accurate character, and "your characters" is last because
     it is the one that needs a sign-in and, until Blizzard ships a Forever profile API, is
     really the addon export under a different name -- which its own copy says out loud
     rather than letting a player assume their Armory gear is being read.
     The signed-in card body is a placeholder: Task 18 replaces it with the landing state's
     character list once that lane exists. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';

  let {
    busy,
    message,
    signedIn,
    onaddon,
    onbuild,
    onfight,
    onsignin,
  }: {
    busy: boolean;
    message: string | null;
    signedIn: boolean;
    onaddon: (code: string) => void;
    onbuild: (id: string) => void;
    onfight: (ref: string) => void;
    onsignin: () => void;
  } = $props();

  let addonCode = $state('');
  let buildRef = $state('');
  let fightRef = $state('');

  const card = 'border-line bg-card-top rounded-panel flex flex-col gap-3 border p-4';
  const field =
    'border-line-warm rounded-control bg-raised text-text min-h-11 w-full border px-3 py-2 text-[14px]';
  const action = 'border-line-warm-strong rounded-control text-strong label min-h-11 self-start border px-4';

  /** Accepts a full link or a bare id: people paste whichever is in their clipboard. */
  export function lastSegment(value: string): string {
    const withoutQuery = value.trim().split('?')[0].replace(/\/+$/, '');
    return withoutQuery.split('/').at(-1) ?? '';
  }

  /** A report link carries the fight in ?fight=; a pasted ref already has it after a colon. */
  export function fightRefOf(value: string): string {
    const trimmed = value.trim();
    if (trimmed.includes(':')) return trimmed;
    const [path, query = ''] = trimmed.split('?');
    const id = path.replace(/\/+$/, '').split('/').at(-1) ?? '';
    const fight = new URLSearchParams(query).get('fight') ?? '1';
    return `${id}:${fight}`;
  }
</script>

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-sources">
  <div class="grid grid-cols-1 gap-3 md:grid-cols-2">
    <div class={card}>
      <h2 class="section-title text-[15px]">{simCopy.sourceAddonTitle}</h2>
      <p class="text-muted text-[13px]">{simCopy.sourceAddonBody}</p>
      <label class="sr-only" for="sim-addon">{simCopy.sourceAddonTitle}</label>
      <textarea
        id="sim-addon"
        rows="2"
        class={field}
        placeholder="FS1:…"
        bind:value={addonCode}
        disabled={busy}
        data-testid="sim-addon-input"></textarea>
      <button
        type="button"
        class={action}
        disabled={busy}
        onclick={() => onaddon(addonCode)}
        data-testid="sim-addon-load"
      >
        Load
      </button>
    </div>

    <div class={card}>
      <h2 class="section-title text-[15px]">{simCopy.sourceBuildTitle}</h2>
      <p class="text-muted text-[13px]">{simCopy.sourceBuildBody}</p>
      <label class="sr-only" for="sim-build">{simCopy.sourceBuildTitle}</label>
      <input
        id="sim-build"
        class={field}
        placeholder="foreversixty.gg/b/…"
        bind:value={buildRef}
        disabled={busy}
        data-testid="sim-build-input"
      />
      <button
        type="button"
        class={action}
        disabled={busy}
        onclick={() => onbuild(lastSegment(buildRef))}
        data-testid="sim-build-load">Load</button
      >
    </div>

    <div class={card}>
      <h2 class="section-title text-[15px]">{simCopy.sourceFightTitle}</h2>
      <p class="text-muted text-[13px]">{simCopy.sourceFightBody}</p>
      <label class="sr-only" for="sim-fight">{simCopy.sourceFightTitle}</label>
      <input
        id="sim-fight"
        class={field}
        placeholder="foreversixty.gg/reports/…?fight=2"
        bind:value={fightRef}
        disabled={busy}
        data-testid="sim-fight-input"
      />
      <button
        type="button"
        class={action}
        disabled={busy}
        onclick={() => onfight(fightRefOf(fightRef))}
        data-testid="sim-fight-load">Load</button
      >
    </div>

    <div class={card} data-testid="sim-account-card">
      <h2 class="section-title text-[15px]">{simCopy.sourceAccountTitle}</h2>
      {#if signedIn}
        <p class="text-muted text-[13px]" data-testid="sim-account-placeholder">
          {simCopy.sourceAccountPlaceholder}
        </p>
      {:else}
        <p class="text-muted text-[13px]">{simCopy.armorySignIn}</p>
        <button type="button" class={action} onclick={onsignin} data-testid="sim-signin">
          Sign in with Battle.net
        </button>
      {/if}
      <p class="text-muted text-[12px]" data-testid="sim-armory-note">{simCopy.armoryNotYet}</p>
    </div>
  </div>

  {#if message}
    <p role="alert" class="text-strong text-[13px]" data-testid="sim-source-message">{message}</p>
  {/if}
</section>
