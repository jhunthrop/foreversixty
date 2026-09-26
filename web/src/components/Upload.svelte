<!-- web/src/components/Upload.svelte -->
<!-- Whole-file upload for a night that was already logged. The report page fills in as the
     parse job closes each fight, so this hands off to /reports/<id> the moment the API
     accepts the upload rather than waiting for a parse it cannot see. -->
<script lang="ts">
  import { fetchMeOnce } from '../lib/account/api';
  import { formatAmount } from '../lib/report/format';
  import {
    MAX_UPLOAD_BYTES,
    UPLOAD_CANCELLED,
    UPLOAD_TOO_LARGE,
    completeUpload,
    createUpload,
    uploadParts,
  } from '../lib/upload/multipart';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import SignInPrompt from './SignInPrompt.svelte';
  import { BUSY_CLASS } from '../lib/ui/busy';

  const VISIBILITIES = [
    { id: 'public', label: 'Public', note: 'Listed, ranked, anyone can open it.' },
    { id: 'unlisted', label: 'Unlisted', note: 'Ranked, but only people with the link can open it.' },
    { id: 'private', label: 'Private', note: 'Only you. Never ranked.' },
    { id: 'guild', label: 'Guild', note: 'Your guild page only. Ranked.' },
  ] as const;

  /**
   * POST /v1/uploads needs a session, so a signed-out visitor is told before they pick a
   * file rather than by a 401 after it. `unknown` keeps the form live while the session is
   * still being asked for: most visitors who get this far are signed in.
   */
  let session = $state<'unknown' | 'in' | 'out'>('unknown');
  $effect(() => {
    fetchMeOnce()
      .then((me) => (session = me === null ? 'out' : 'in'))
      .catch(() => (session = 'out'));
  });

  let file = $state<File | null>(null);
  let title = $state('');
  let visibility = $state<string>('public');
  let phase = $state<'idle' | 'uploading' | 'finishing' | 'failed'>('idle');
  let uploaded = $state(0);
  let error = $state('');
  /** Non-null only while an upload is in flight; `cancel` is the only thing that fires it. */
  let inFlight: AbortController | null = null;

  const percent = $derived(file === null || uploaded === 0 ? 0 : Math.round((uploaded / file.size) * 100));
  const tooLarge = $derived(file !== null && file.size > MAX_UPLOAD_BYTES);
  const busy = $derived(phase === 'uploading' || phase === 'finishing');
  const locked = $derived(busy || session === 'out');
  const visibilityNote = $derived(VISIBILITIES.find((option) => option.id === visibility)?.note ?? '');

  const FIELD = 'border-line-warm bg-bg rounded-control text-text h-11 border px-3 text-[15px]';
  const CHOICE =
    'border-line-warm rounded-control text-nav inline-flex min-h-11 cursor-pointer items-center border px-3 text-[12px] font-bold tracking-[0.06em] uppercase md:min-h-9';

  function pick(event: Event): void {
    const input = event.currentTarget as HTMLInputElement;
    file = input.files?.[0] ?? null;
    uploaded = 0;
    error = tooLarge ? UPLOAD_TOO_LARGE : '';
    phase = 'idle';
  }

  function onDrop(event: DragEvent): void {
    event.preventDefault();
    file = event.dataTransfer?.files?.[0] ?? null;
    uploaded = 0;
    error = tooLarge ? UPLOAD_TOO_LARGE : '';
  }

  /**
   * Cancels the part in flight, not just the queue: multipart.ts hands the signal to the
   * XMLHttpRequest, so a 64 MiB part that is halfway up stops there rather than finishing
   * first. Nothing is uploaded afterwards, and the multipart upload is left for the API to
   * clean up -- the browser has no endpoint to abort it with.
   */
  function cancel(): void {
    inFlight?.abort();
  }

  async function start(): Promise<void> {
    if (file === null || tooLarge) return;
    const chosen = file;
    const controller = new AbortController();
    inFlight = controller;
    phase = 'uploading';
    error = '';
    uploaded = 0;
    try {
      const created = await createUpload(chosen);
      // `restart` is how an upload that outlives its own signed urls -- an hour, and a
      // 4 GiB night on a home uplink is longer -- recovers instead of throwing away every
      // part it already sent. It is a fresh upload, not a re-signing of this one, so the
      // id to complete is whichever one uploadParts actually finished against.
      const { uploadId, etags } = await uploadParts(
        chosen,
        created,
        (progress) => {
          uploaded = progress.uploadedBytes;
        },
        { signal: controller.signal, restart: () => createUpload(chosen) },
      );
      phase = 'finishing';
      const reportId = await completeUpload(uploadId, etags, { title: title.trim(), visibility });
      window.location.assign(`/reports/${reportId}`);
    } catch (thrown) {
      const message = thrown instanceof Error ? thrown.message : UPLOAD_TOO_LARGE;
      // Cancelling is not a failure to report back: the visitor asked for it and already
      // knows. The form goes back to where it was, with the same file still chosen.
      if (message === UPLOAD_CANCELLED) {
        phase = 'idle';
        uploaded = 0;
        return;
      }
      phase = 'failed';
      error = message;
    } finally {
      if (inFlight === controller) inFlight = null;
    }
  }
</script>

<!-- The panel and its title belong to the page (logs.astro); this is the form inside it. -->
<div class="flex flex-col gap-4 p-4 md:p-5" data-testid="upload">
  <p class="text-muted text-[14px]">
    A whole <code class="font-mono">WoWCombatLog.txt</code>. The first fight is readable within ten seconds of
    the upload finishing, and the rest appear as they are parsed.
  </p>

  <p class="text-muted text-[13px]">
    The game writes it to its <code class="font-mono">Logs</code> folder, beside
    <code class="font-mono">Interface</code> and <code class="font-mono">WTF</code>. Up to 4 GB. Type
    <code class="font-mono">/combatlog</code> again to stop logging before you upload, so the file is whole.
  </p>

  <!-- Above the form, not under it: a signed-out visitor reads this before they pick a file.
       No floor is held for it. Most visitors who get this far are signed in and never see
       it, and on a phone this panel is below the first screen either way. -->
  {#if session === 'out'}
    <SignInPrompt
      line="Sign in to upload. The report is filed under your account."
      testid="upload-signin"
      compact
    />
  {/if}

  <!-- The form itself only exists for a session: signed out, the compact prompt above is
       the whole panel, so nobody picks a file, titles it and chooses who can see it before
       finding out the upload was never going to start. -->
  {#if session !== 'out'}
    <div
      class="border-line-warm rounded-panel flex flex-col items-start gap-2 border border-dashed p-4"
      ondrop={onDrop}
      ondragover={(event) => event.preventDefault()}
      role="group"
      aria-label="Choose a combat log"
    >
      <input
        id="upload-file"
        class="peer sr-only"
        type="file"
        accept=".txt,text/plain"
        onchange={pick}
        disabled={locked}
        data-testid="upload-file"
      />
      <label
        for="upload-file"
        class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong peer-focus-visible:outline-gold cursor-pointer px-4 peer-focus-visible:outline-2 peer-disabled:cursor-default peer-disabled:opacity-50"
      >
        Choose file
      </label>
      {#if file === null}
        <p class="text-muted text-[13px]">or drag <code class="font-mono">WoWCombatLog.txt</code> here</p>
      {:else}
        <p class="text-muted tabular font-mono text-[13px] break-all" data-testid="upload-size">
          {file.name} · {formatAmount(file.size)} bytes
        </p>
      {/if}
    </div>

    <div class="flex flex-col gap-1">
      <label class="label text-muted" for="upload-title">Title</label>
      <input
        id="upload-title"
        class={FIELD}
        type="text"
        maxlength="60"
        placeholder="Molten Core, week 3"
        bind:value={title}
        disabled={locked}
      />
      <p class="text-muted text-[13px]">Optional. Without one the report is named after its zone.</p>
    </div>

    <fieldset class="flex flex-col gap-2" disabled={locked}>
      <legend class="label text-muted mb-2">Who can see it</legend>
      <div class="flex flex-wrap gap-2">
        {#each VISIBILITIES as option (option.id)}
          <label
            class="{CHOICE} has-[:checked]:border-gold-deep has-[:checked]:text-gold has-[:focus-visible]:outline-gold has-[:focus-visible]:outline-2"
          >
            <input type="radio" name="visibility" value={option.id} bind:group={visibility} class="sr-only" />
            {option.label}
          </label>
        {/each}
      </div>
      <p class="text-muted text-[13px]" data-testid="upload-visibility-note">{visibilityNote}</p>
    </fieldset>

    <div class="flex flex-wrap items-center gap-3">
      <button
        class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4 disabled:opacity-50 ${busy ? BUSY_CLASS : ''}`}
        onclick={() => void start()}
        disabled={file === null || tooLarge || locked}
        aria-busy={busy}
        data-testid="upload-start"
      >
        Upload
      </button>
      {#if phase === 'uploading'}
        <button
          class="{SECONDARY_BUTTON_FIXED} border-line-warm text-muted w-fit px-4"
          onclick={cancel}
          data-testid="upload-cancel"
        >
          Cancel
        </button>
      {/if}
    </div>
  {/if}

  {#if busy}
    <div class="flex flex-col gap-2" data-testid="upload-progress">
      <div
        class="bg-line-soft h-[6px] w-full overflow-hidden"
        role="progressbar"
        aria-valuenow={percent}
        aria-valuemin="0"
        aria-valuemax="100"
      >
        <div class="bg-gold h-full" style={`width: ${percent}%`}></div>
      </div>
      <p class="text-muted tabular font-mono text-[13px]">
        {phase === 'finishing' ? 'Finishing' : `${percent}%`}
      </p>
    </div>
  {/if}

  {#if error !== ''}
    <p class="text-[14px]" role="alert" data-testid="upload-error">{error}</p>
  {/if}
</div>
