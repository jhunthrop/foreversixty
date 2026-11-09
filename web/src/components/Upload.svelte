<!-- web/src/components/Upload.svelte -->
<!-- Whole-file upload for a night that was already logged. The report page fills in as the
     parse job closes each fight, so this hands off to /reports/<id> the moment the API
     accepts the upload rather than waiting for a parse it cannot see.

     Logs spec (2026-11-07) section 4: the explanation and the DropZone; Title, who can open it
     (a Segmented) and the Upload button appear only once a file is chosen. Signed out the zone is
     drawn at 55% and disabled, and the sentence that uploading needs an account sits over the
     page's one primary, Sign in with Battle.net, and the sample report (decision L3). `plate`
     pads it as the page's Plate content; open in a Section it runs flush. -->
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
  import { battlenetStartUrl } from '../lib/account/api';
  import { myReportsCopy, yourReportsPageCopy } from '../lib/reports/my-reports-copy';
  import { SAMPLE_REPORT_PATH } from '../lib/report/sample';
  import { checkLog, logWarningCopy, type LogCheck } from '../lib/upload/log-check';
  import { LINK } from '../lib/ui/primitives';
  import Button from './ui/Button.svelte';
  import DropZone from './ui/DropZone.svelte';
  import Field from './ui/Field.svelte';
  import Segmented from './ui/Segmented.svelte';

  let { plate = false }: { plate?: boolean } = $props();

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
   * file rather than by a 401 after it.
   */
  let session = $state<'unknown' | 'in' | 'out'>('unknown');
  /** Until the session answers this is the signed-out form (the server markup, final for a
   *  visitor with no session hint; the page hides it from a hinted one until the answer). */
  const signedOut = $derived(session !== 'in');
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

  const percent = $derived(file === null || uploaded === 0 ? 0 : Math.round((uploaded / file.size) * 100));
  const tooLarge = $derived(file !== null && file.size > MAX_UPLOAD_BYTES);
  const busy = $derived(phase === 'uploading' || phase === 'finishing');
  const locked = $derived(busy || session === 'out');
  const visibilityNote = $derived(VISIBILITIES.find((option) => option.id === visibility)?.note ?? '');

  function pick(event: Event): void {
    const input = event.currentTarget as HTMLInputElement;
    file = input.files?.[0] ?? null;
    inspect(file);
    uploaded = 0;
    error = tooLarge ? UPLOAD_TOO_LARGE : '';
    phase = 'idle';
  }

  function onDrop(files: FileList | null): void {
    file = files?.[0] ?? null;
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

<div class={['flex flex-col gap-3', plate && 'px-5 py-[18px]', !plate && 'pt-1']} data-testid="upload">
  <p class="max-w-[72ch] text-[15px] leading-6">
    A whole <code class="font-mono text-[13px] leading-5">WoWCombatLog.txt</code>. Your report appears as soon
    as the parse starts, usually within a minute or two; the fights fill in as they are read.
  </p>

  <p class="text-muted max-w-[72ch] text-[15px] leading-6">
    The game writes it to its <code class="font-mono text-[13px] leading-5">Logs</code> folder, beside
    <code class="font-mono text-[13px] leading-5">Interface</code> and
    <code class="font-mono text-[13px] leading-5">WTF</code>. Up to 4 GB. Type
    <code class="font-mono text-[13px] leading-5">/combatlog</code> again to stop logging first, so the file is
    whole.
  </p>

  {#if signedOut}
    <DropZone off testid="upload-zone">
      <span>Drop <code class="font-mono text-[13px] leading-5">WoWCombatLog.txt</code> here, or</span>
      <span class="font-semibold">choose a file</span>
    </DropZone>
    <p class="max-w-[72ch] text-[15px] leading-6" data-testid="upload-signed-out">{SIGNED_OUT_LINE}</p>
    <div class="flex flex-wrap items-center gap-4">
      <Button tone="primary" href={battlenetStartUrl('/logs')} testid="reports-sign-in">
        {myReportsCopy.signInAction}
      </Button>
      <a class={LINK} href={SAMPLE_REPORT_PATH} data-testid="upload-sample"
        >{yourReportsPageCopy.sampleAction} ›</a
      >
    </div>
  {:else}
    <DropZone ondrop={onDrop} testid="upload-zone">
      {#if file === null}
        <span>Drop <code class="text-text font-mono">WoWCombatLog.txt</code> here, or</span>
      {:else}
        <span class="text-text font-mono break-all">{file.name}</span>
        <span class="numeric" data-testid="upload-size">
          {formatAmount(file.size)} byte{file.size === 1 ? '' : 's'}
        </span>
      {/if}
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
      <label
        for="upload-file"
        class="{LINK} peer-focus-visible:outline-gold cursor-pointer peer-focus-visible:outline-2 peer-disabled:cursor-default peer-disabled:opacity-50"
      >
        {file === null ? 'choose a file' : 'Choose another file'}
      </label>
    </DropZone>

    {#if file !== null}
      <Field
        id="upload-title"
        label="Title"
        maxlength={60}
        placeholder="Optional"
        bind:value={title}
        hint="Optional. Without one the report is named after its zone."
      />

      <div class="flex flex-col gap-2">
        <span class="text-label text-muted font-bold" id="upload-visibility-label">Who can open it</span>
        <Segmented
          options={VISIBILITIES.map((option) => ({ value: option.id, label: option.label }))}
          value={visibility}
          label="Who can open it"
          disabled={locked}
          onChange={(next) => (visibility = next)}
        />
        <p class="text-small text-muted" data-testid="upload-visibility-note">{visibilityNote}</p>
      </div>

      {#if logCheck !== null && (logCheck.advancedOff || logCheck.noEncounters)}
        <!-- Advice, never a gate: the Upload button below stays live. -->
        <div class="flex flex-col gap-1" data-testid="upload-warnings">
          {#if logCheck.advancedOff}
            <p class="text-[14px]" data-testid="upload-warning-advanced">{logWarningCopy.advancedOff}</p>
          {/if}
          {#if logCheck.noEncounters}
            <p class="text-[14px]" data-testid="upload-warning-encounters">{logWarningCopy.noEncounters}</p>
          {/if}
        </div>
      {/if}

      <div class="flex flex-wrap items-center gap-3">
        <span aria-busy={busy}>
          <Button
            tone="primary"
            onclick={() => void start()}
            disabled={tooLarge || locked}
            testid="upload-start"
          >
            Upload
          </Button>
        </span>
        {#if phase === 'uploading'}
          <Button onclick={cancel} testid="upload-cancel">Cancel</Button>
        {/if}
      </div>
    {/if}
  {/if}

  {#if busy}
    <div class="flex flex-col gap-2" data-testid="upload-progress">
      <div
        class="h-[6px] w-full overflow-hidden rounded-[3px] bg-white/10"
        role="progressbar"
        aria-valuenow={percent}
        aria-valuemin="0"
        aria-valuemax="100"
      >
        <div class="bg-gold h-full" style={`width: ${percent}%`}></div>
      </div>
      <p class="text-small text-muted numeric">
        {phase === 'finishing' ? 'Finishing' : `${percent}%`}
      </p>
    </div>
  {/if}

  {#if error !== ''}
    <p class="text-[14px]" role="alert" data-testid="upload-error">{error}</p>
  {/if}
</div>
