<!-- web/src/components/Upload.svelte -->
<!-- Whole-file upload for a night that was already logged. The report page fills in as the
     parse job closes each fight, so this hands off to /reports/<id> the moment the API
     accepts the upload rather than waiting for a parse it cannot see.

     Logs spec (2026-10-11) section 2.4: before a file is chosen the panel is the explanation and
     the drop zone, nothing else; Title, who can open it and the outline Upload button appear only
     once a file is chosen. The page's header action calls `openPicker()` to open the system file
     picker from its own click. Signed out, the panel says in plain words that uploading needs an
     account; it carries no sign-in control. -->
<script lang="ts">
  import { fetchMeOnce } from '../lib/account/api';
  import { formatAmount } from '../lib/report/format';
  import { markJustUploaded } from '../lib/reports/just-uploaded';
  import {
    MAX_UPLOAD_BYTES,
    UPLOAD_CANCELLED,
    UPLOAD_TOO_LARGE,
    completeUpload,
    createUpload,
    uploadParts,
  } from '../lib/upload/multipart';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import { checkLog, logWarningCopy, type LogCheck } from '../lib/upload/log-check';
  import { BUSY_CLASS } from '../lib/ui/busy';

  /** The one line a player ever sees on a failed upload (logs build review, fix round 1,
   *  2026-10-04): multipart.ts's own thrown messages carry real diagnostic detail -- a
   *  retry count, "R2 answered 500", a doubled "did not finish" wrapper -- none of which
   *  belongs on screen. That detail is still logged (see the catch block below), just not
   *  rendered. The file stays chosen either way (only `phase`/`error` reset on failure), so
   *  "Try again" really does mean one click, not re-picking the file. */
  const UPLOAD_FAILED_LINE = 'The upload did not finish. Your file is untouched; try again.';
  const SIGNED_OUT_LINE = 'Uploading needs a signed-in account; the report is filed under it.';

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
  let fileInput = $state<HTMLInputElement | null>(null);
  let dragging = $state(false);
  /** What reading the chosen file's header and boss markers found; null until it has been read.
   *  Advisory: it never blocks the Upload button. */
  let logCheck = $state<LogCheck | null>(null);
  let checkToken = 0;

  /** Reads the chosen file in the background (its first 64 KB, then a stream for boss fights);
   *  a newer choice makes an older read's answer stale, so it is dropped. */
  function inspect(chosen: File | null): void {
    const token = ++checkToken;
    logCheck = null;
    if (chosen === null) return;
    checkLog(chosen)
      .then((result) => {
        if (token === checkToken) logCheck = result;
      })
      .catch(() => {
        // An unreadable file gets no advice; the upload says what it needs to.
      });
  }

  /** The header action's handler: opens the system file picker from the click that called it. */
  export function openPicker(): void {
    fileInput?.click();
  }

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
    inspect(file);
    uploaded = 0;
    error = tooLarge ? UPLOAD_TOO_LARGE : '';
    phase = 'idle';
  }

  function onDrop(event: DragEvent): void {
    event.preventDefault();
    dragging = false;
    file = event.dataTransfer?.files?.[0] ?? null;
    inspect(file);
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
      markJustUploaded(reportId);
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
      // One honest, generic line on screen; multipart.ts's own technical detail (part
      // number, retry count, the storage backend's status) goes to the console instead,
      // never the page -- see UPLOAD_FAILED_LINE's own comment above.
      error = UPLOAD_FAILED_LINE;
      console.error('Upload failed:', thrown);
    } finally {
      if (inFlight === controller) inFlight = null;
    }
  }
</script>

<!-- The panel and its title belong to the page (LogsLanding.svelte); this is the form inside it. -->
<div class="flex flex-col gap-3" data-testid="upload">
  <p class="text-[14px]">
    A whole <code class="font-mono">WoWCombatLog.txt</code>. Your report appears as soon as the parse starts,
    usually within a minute or two; the fights fill in as they are read.
  </p>

  <p class="text-muted text-[13px]">
    The game writes it to its <code class="font-mono">Logs</code> folder, beside
    <code class="font-mono">Interface</code> and <code class="font-mono">WTF</code>. Up to 4 GB. Type
    <code class="font-mono">/combatlog</code> again to stop logging first, so the file is whole.
  </p>

  <!-- Signed out the panel is these two paragraphs and one plain sentence: nobody picks a file,
       titles it and chooses who can see it before finding out the upload was never going to
       start, and the page carries no sign-in button. -->
  {#if session === 'out'}
    <p class="text-[14px]" data-testid="upload-signed-out">{SIGNED_OUT_LINE}</p>
  {:else}
    <div
      class="rounded-panel flex min-h-[150px] flex-col items-center justify-center gap-2 border border-dashed p-4 text-center text-[13px]"
      class:border-gold={dragging}
      class:border-line-warm-strong={!dragging}
      style:background={dragging ? 'color-mix(in srgb, var(--color-gold) 6%, transparent)' : undefined}
      ondrop={onDrop}
      ondragover={(event) => {
        event.preventDefault();
        dragging = true;
      }}
      ondragleave={() => (dragging = false)}
      role="group"
      aria-label="Choose a combat log"
      data-testid="upload-zone"
      data-dragging={dragging}
    >
      <input
        id="upload-file"
        bind:this={fileInput}
        class="peer sr-only"
        type="file"
        accept=".txt,text/plain"
        onchange={pick}
        oncancel={() => fileInput?.focus()}
        disabled={locked}
        data-testid="upload-file"
      />
      {#if file === null}
        <p class="text-muted">Drop <code class="text-text font-mono">WoWCombatLog.txt</code> here, or</p>
      {:else}
        <p class="text-text font-mono break-all">{file.name}</p>
        <p class="text-muted numeric" data-testid="upload-size">
          {formatAmount(file.size)} byte{file.size === 1 ? '' : 's'}
        </p>
      {/if}
      <label
        for="upload-file"
        class="text-gold peer-focus-visible:outline-gold inline-flex min-h-11 cursor-pointer items-center text-[12px] font-bold tracking-[0.06em] uppercase peer-focus-visible:outline-2 peer-disabled:cursor-default peer-disabled:opacity-50"
      >
        {file === null ? 'choose a file' : 'Choose another file'}
      </label>
    </div>

    {#if file !== null}
      <div class="flex flex-col gap-1">
        <label class="quiet-label" for="upload-title">Title</label>
        <input
          id="upload-title"
          class={FIELD}
          type="text"
          maxlength="60"
          placeholder="Optional"
          bind:value={title}
          disabled={locked}
        />
        <p class="text-muted text-[13px]">Optional. Without one the report is named after its zone.</p>
      </div>

      <fieldset class="flex flex-col gap-2" disabled={locked}>
        <legend class="quiet-label mb-2">Who can open it</legend>
        <div class="flex flex-wrap gap-2">
          {#each VISIBILITIES as option (option.id)}
            <label
              class="{CHOICE} has-[:checked]:border-gold-deep has-[:checked]:text-gold has-[:focus-visible]:outline-gold has-[:focus-visible]:outline-2"
            >
              <input
                type="radio"
                name="visibility"
                value={option.id}
                bind:group={visibility}
                class="sr-only"
              />
              {option.label}
            </label>
          {/each}
        </div>
        <p class="text-muted text-[13px]" data-testid="upload-visibility-note">{visibilityNote}</p>
      </fieldset>

      {#if logCheck !== null && (logCheck.advancedOff || logCheck.noEncounters)}
        <!-- Advice, never a gate: the Upload button below stays live. -->
        <div class="flex flex-col gap-1" data-testid="upload-warnings">
          {#if logCheck.advancedOff}
            <p class="text-[13px]" data-testid="upload-warning-advanced">{logWarningCopy.advancedOff}</p>
          {/if}
          {#if logCheck.noEncounters}
            <p class="text-[13px]" data-testid="upload-warning-encounters">{logWarningCopy.noEncounters}</p>
          {/if}
        </div>
      {/if}

      <div class="flex flex-wrap items-center gap-3">
        <button
          class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong w-fit px-4 disabled:opacity-50 ${busy ? BUSY_CLASS : ''}`}
          onclick={() => void start()}
          disabled={tooLarge || locked}
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
      <p class="text-muted numeric text-[13px]">
        {phase === 'finishing' ? 'Finishing' : `${percent}%`}
      </p>
    </div>
  {/if}

  {#if error !== ''}
    <p class="text-[14px]" role="alert" data-testid="upload-error">{error}</p>
  {/if}
</div>
