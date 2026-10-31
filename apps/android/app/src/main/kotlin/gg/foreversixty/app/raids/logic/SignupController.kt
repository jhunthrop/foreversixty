package gg.foreversixty.app.raids.logic

import android.content.res.Resources
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateMapOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.R
import gg.foreversixty.app.model.RaidSpecOption
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.model.RaidViewerCharacter
import gg.foreversixty.app.model.SignupRequest
import gg.foreversixty.app.network.ApiResult
import java.time.Instant
import kotlin.time.Duration
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch

/** The 20 dp status line under a character's segments. */
sealed interface StatusLine {
    val text: String

    /** `Saved.` in kill green, for 4 s. */
    data class Saved(override val text: String) : StatusLine

    /** A refusal or failure in wipe red. */
    data class Error(override val text: String) : StatusLine

    /** A neutral line: `{name} moved to Bench.` */
    data class Info(override val text: String) : StatusLine
}

/** One character block's drawn state. */
data class SignupBlock(
    /** The checked segment: optimistic while a save is out. */
    val shown: SignupState? = null,
    /** A press waiting for "Which spec tonight?" (the dashed outline). */
    val pending: SignupState? = null,
    val busy: Boolean = false,
    val status: StatusLine? = null,
    val noteBusy: Boolean = false,
    val noteFailed: Boolean = false,
)

/** The spec sheet's question: a press waiting for a spec, or Change spec. */
sealed interface SpecAsk {
    val characterKey: String

    data class Ask(override val characterKey: String, val state: SignupState) : SpecAsk
    data class Change(override val characterKey: String) : SpecAsk
}

/**
 * The raid page's and the Raids tab card's sign-up, the officer tools and the reminder ask. One raid
 * view is the truth; a press draws at once, the PUT's raid replaces the view, a failure restores the
 * previous answer (spec F4.3, F4.11).
 */
class SignupController(
    val guild: Long,
    raid: RaidView,
    private val deps: Dependencies,
) {
    class Dependencies(
        val service: RaidService,
        val res: Resources,
        val haptics: RaidHaptics,
        val permission: NotificationPermission,
        val askStore: ReminderAskStore,
        val now: () -> Instant,
        /** Where `Saved.`, `Copied` and the toast time out; null keeps them (tests). */
        val scope: CoroutineScope?,
        val hold: Duration,
    )

    var raid: RaidView by mutableStateOf(raid)
        private set
    private val blocks = mutableStateMapOf<String, SignupBlock>()
    val notes = mutableStateMapOf<String, String>()
    var specAsk: SpecAsk? by mutableStateOf(null)
        private set

    /** The last write or refresh failed for lack of a connection: segments disable. */
    var offline: Boolean by mutableStateOf(false)
    val rosterOverrides = mutableStateMapOf<String, SignupState?>()
    val rowErrors = mutableStateMapOf<String, String>()
    var officerBusy: Boolean by mutableStateOf(false)
        private set
    var officerError: String? by mutableStateOf(null)
        private set
    var toast: String? by mutableStateOf(null)
        private set
    var copied: Boolean by mutableStateOf(false)
        private set
    var showsReminderAsk: Boolean by mutableStateOf(false)
        private set

    /** Hears every new raid view, so the Raids tab's list follows the page. */
    var onRaidChanged: ((RaidView) -> Unit)? = null

    private val res get() = deps.res

    init {
        apply(raid)
        raid.signups.filter { it.own }.forEach { notes[it.characterKey] = it.note }
    }

    val facts: RaidFacts get() = RaidFacts(raid, deps.now())

    fun block(key: String): SignupBlock = blocks[key] ?: SignupBlock()

    fun character(key: String): RaidViewerCharacter? = raid.viewer.characters.firstOrNull { it.characterKey == key }

    /** Online, a verified character, a signable raid, and not locked unless the viewer may still write. */
    fun segmentsEnabled(key: String): Boolean {
        val character = character(key) ?: return false
        if (!character.verified || offline) return false
        val facts = facts
        return facts.isSignable && (facts.chip != RaidChip.Locked || facts.canMove)
    }

    /** The status line when no press has spoken: the raid state's sentence. */
    fun restingStatus(key: String): StatusLine? {
        block(key).status?.let { return it }
        if (offline) return StatusLine.Info(res.getString(R.string.raid_offline))
        val facts = facts
        return when (facts.chip) {
            RaidChip.Locked -> StatusLine.Info(
                res.getString(if (facts.canMove) R.string.officer_memberLocked else R.string.raid_locked),
            )
            RaidChip.Ended -> StatusLine.Info(res.getString(R.string.raid_ended))
            else -> null
        }
    }

    private fun update(key: String, change: (SignupBlock) -> SignupBlock) {
        blocks[key] = change(block(key))
    }

    // Sign-up

    suspend fun press(key: String, state: SignupState) {
        val character = character(key) ?: return
        if (!segmentsEnabled(key) || block(key).busy) return
        deps.haptics.selection()
        if (state.carriesSpec && character.raidSpec == null && (character.specOptions?.size ?: 0) >= 2) {
            update(key) { it.copy(pending = state) }
            specAsk = SpecAsk.Ask(key, state)
            return
        }
        save(key, state, null)
    }

    fun changeSpec(key: String) {
        if (segmentsEnabled(key)) specAsk = SpecAsk.Change(key)
    }

    suspend fun pickSpec(option: RaidSpecOption) {
        val ask = specAsk ?: return
        specAsk = null
        when (ask) {
            is SpecAsk.Ask -> {
                update(ask.characterKey) { it.copy(pending = null) }
                save(ask.characterKey, ask.state, option)
            }
            is SpecAsk.Change -> block(ask.characterKey).shown?.let { save(ask.characterKey, it, option) }
        }
    }

    /** Closing the sheet drops the press: nothing is saved. */
    fun closeSpecSheet() {
        (specAsk as? SpecAsk.Ask)?.let { ask -> update(ask.characterKey) { it.copy(pending = null) } }
        specAsk = null
    }

    private suspend fun save(key: String, state: SignupState, spec: RaidSpecOption?) {
        val previous = block(key).shown
        val name = character(key)?.name.orEmpty()
        update(key) { it.copy(shown = state, busy = true, status = null) }
        val request = SignupRequest(state = state.wire, note = notes[key].orEmpty(), spec = spec?.spec)
        when (val result = deps.service.signup(guild, raid.id, key, request)) {
            is ApiResult.Success -> {
                offline = false
                apply(result.data.raid)
                val kept = spec?.let { " " + res.getString(R.string.spec_kept, it.name, name) }.orEmpty()
                setStatus(key, StatusLine.Saved(res.getString(R.string.common_saved) + kept))
                result.data.moved?.let { moved ->
                    raid.viewer.characters.firstOrNull { it.name == moved && it.characterKey != key }?.let { other ->
                        update(other.characterKey) { it.copy(status = StatusLine.Info(res.getString(R.string.raid_moved, moved))) }
                    }
                }
                deps.haptics.success()
                if (state.isComing) considerReminderAsk()
            }
            is ApiResult.Failure -> {
                update(key) { it.copy(shown = previous, busy = false, status = StatusLine.Error(SignupRefusal.sentence(result.error, res))) }
                offline = SignupRefusal.isOffline(result.error)
                deps.haptics.error()
            }
        }
    }

    // Note

    fun editNote(key: String, text: String) {
        notes[key] = text.take(NOTE_LIMIT)
    }

    /** Saves the note with the current answer (Done or blur). The PUT writes note and state. */
    suspend fun saveNote(key: String) {
        val state = block(key).shown ?: return
        if (!segmentsEnabled(key)) return
        val stored = raid.signups.firstOrNull { it.characterKey == key }?.note.orEmpty()
        if (notes[key].orEmpty() == stored) return
        update(key) { it.copy(noteBusy = true, noteFailed = false) }
        val result = deps.service.signup(guild, raid.id, key, SignupRequest(state = state.wire, note = notes[key].orEmpty()))
        if (result is ApiResult.Success) apply(result.data.raid)
        update(key) { it.copy(noteBusy = false, noteFailed = result is ApiResult.Failure) }
    }

    // Reminder ask

    private fun considerReminderAsk() {
        val now = deps.now()
        if (showsReminderAsk || !deps.askStore.policy().mayAsk(now) || !deps.permission.isUndetermined()) return
        deps.askStore.recordShown(now)
        showsReminderAsk = true
    }

    /** `Turn on reminders` or `Not now`; the screen asks the system for the permission. */
    fun dismissReminderAsk() {
        showsReminderAsk = false
    }

    // Officer tools

    /** An officer's move (spec F4.20): the officer PUT for the row's character, with its role. */
    suspend fun move(row: RosterRow, to: SignupState) {
        if (!facts.canMove || offline) return
        rosterOverrides[row.characterKey] = to
        rowErrors.remove(row.characterKey)
        val request = SignupRequest(state = to.wire, role = row.role, note = row.note)
        when (val result = deps.service.signup(guild, raid.id, row.characterKey, request)) {
            is ApiResult.Success -> {
                rosterOverrides.remove(row.characterKey)
                apply(result.data.raid)
                deps.haptics.success()
            }
            is ApiResult.Failure -> {
                rosterOverrides.remove(row.characterKey)
                rowErrors[row.characterKey] = SignupRefusal.sentence(result.error, res)
                deps.haptics.error()
            }
        }
    }

    suspend fun toggleLock() {
        if (!facts.canLock || officerBusy) return
        val locking = facts.chip != RaidChip.Locked
        officerBusy = true
        officerError = null
        val result = if (locking) deps.service.lock(guild, raid.id) else deps.service.unlock(guild, raid.id)
        when (result) {
            is ApiResult.Success -> {
                apply(result.data)
                if (locking) {
                    val counts = result.data.counts
                    flash({ toast = it }, res.getString(R.string.officer_lockedToast, counts.signedIn, counts.bench ?: 0))
                }
                deps.haptics.success()
            }
            is ApiResult.Failure -> {
                officerError = SignupRefusal.sentence(result.error, res)
                deps.haptics.error()
            }
        }
        officerBusy = false
    }

    fun copyList(clock: RaidClock): String {
        copied = true
        deps.scope?.launch {
            delay(deps.hold)
            copied = false
        }
        return RosterCopy.text(raid, clock, res, deps.now())
    }

    // Refresh

    private val isWriting: Boolean get() = blocks.values.any { it.busy } || rosterOverrides.isNotEmpty()

    /** A refetch replaces the view unless a save is out. */
    suspend fun refresh() {
        if (isWriting) return
        when (val result = deps.service.raid(guild, raid.id)) {
            is ApiResult.Success -> replace(result.data)
            is ApiResult.Failure -> offline = SignupRefusal.isOffline(result.error)
        }
    }

    /** A newer view from elsewhere; skipped while a write is out. */
    fun replace(view: RaidView) {
        if (isWriting) return
        offline = false
        apply(view)
    }

    private fun apply(view: RaidView) {
        raid = view
        view.viewer.characters.forEach { c ->
            update(c.characterKey) { it.copy(shown = SignupState.of(c.state), busy = false) }
        }
        onRaidChanged?.invoke(view)
    }

    private fun setStatus(key: String, status: StatusLine) {
        update(key) { it.copy(status = status) }
        deps.scope?.launch {
            delay(deps.hold)
            if (block(key).status == status) update(key) { it.copy(status = null) }
        }
    }

    private fun flash(set: (String?) -> Unit, text: String) {
        set(text)
        deps.scope?.launch {
            delay(deps.hold)
            set(null)
        }
    }

    companion object {
        const val NOTE_LIMIT = 140
        const val NOTE_COUNTER_FROM = 120
    }
}
