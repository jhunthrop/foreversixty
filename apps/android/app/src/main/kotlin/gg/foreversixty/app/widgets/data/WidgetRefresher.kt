package gg.foreversixty.app.widgets.data

import gg.foreversixty.app.model.LiveActivityRegistered
import gg.foreversixty.app.model.LiveActivityRequest
import gg.foreversixty.app.model.SignupRequest
import gg.foreversixty.app.model.Widget
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.logic.RaidService
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.widgets.logic.SignupOverlay
import gg.foreversixty.app.widgets.logic.WidgetSnapshot
import java.time.Duration
import java.time.Instant
import kotlinx.coroutines.delay

/** The live-surface calls (W3-C12, W3-C13). Live goes through [ApiClient]; tests pass a fake. */
interface WidgetService {
    suspend fun widget(character: String?): ApiResult<Widget>
    suspend fun register(request: LiveActivityRequest): ApiResult<LiveActivityRegistered>
    suspend fun unregister(id: Long): ApiResult<Unit>
}

class LiveWidgetService(private val client: ApiClient) : WidgetService {
    override suspend fun widget(character: String?) = client.send(ApiEndpoints.getWidget(character))
    override suspend fun register(request: LiveActivityRequest) = client.send(ApiEndpoints.registerLiveActivity(request))
    override suspend fun unregister(id: Long) = client.send(ApiEndpoints.deleteLiveActivity(id))
}

/**
 * Writes the snapshot every surface reads (spec §5.6 Refresh): on every foreground, after a sign-up
 * in the app or on a widget, from the 30-minute WorkManager job and on a `widget.refresh` push.
 * Then [onChanged] redraws the widgets, the tile and the ongoing raid notification.
 */
class WidgetRefresher(
    private val service: WidgetService,
    private val store: WidgetStore,
    private val signedIn: suspend () -> Boolean,
    private val currentCharacter: () -> String?,
    private val now: () -> Instant = Instant::now,
    private val onChanged: suspend () -> Unit = {},
) {
    /** Refreshes every character a widget shows. A network failure keeps the old snapshot (it goes stale in words). */
    suspend fun refresh() {
        if (!signedIn()) {
            clear()
            return
        }
        for (character in store.characters()) {
            val asked = if (character == WidgetStore.CURRENT) currentCharacter() else character
            when (val result = service.widget(asked)) {
                is ApiResult.Success -> store.saveSnapshot(character, WidgetSnapshot(result.data, now().toString()))
                is ApiResult.Failure -> if ((result.error as? ApiError.Server)?.isUnauthorized == true) {
                    clear()
                    return
                }
            }
        }
        onChanged()
    }

    /** Sign-out: every surface returns to `Sign in to see your next raid.` */
    suspend fun clear() {
        store.clear()
        onChanged()
    }
}

/**
 * In, Late or Out pressed on the 4 x 2 widget (spec §5.6): the press redraws at once, the existing
 * sign-up PUT answers it (keeping the stored note), then the server's snapshot redraws it. A refused
 * press redraws the previous answer and says `Not saved` on its button for 10 s.
 */
class WidgetSignup(
    private val raids: RaidService,
    private val store: WidgetStore,
    private val refresher: WidgetRefresher,
    private val now: () -> Instant = Instant::now,
    private val redraw: suspend () -> Unit = {},
    private val wait: suspend (Duration) -> Unit = { delay(it.toMillis()) },
) {
    suspend fun answer(guild: Long, raid: Long, characterKey: String, state: SignupState): Boolean {
        store.setOverlay(SignupOverlay(raid, state.wire))
        redraw()
        if (put(guild, raid, characterKey, state)) {
            store.setOverlay(null)
            refresher.refresh()
            return true
        }
        val failed = SignupOverlay(raid, state.wire, failedUntil = now().plus(NOT_SAVED_FOR).toString())
        store.setOverlay(failed)
        redraw()
        wait(NOT_SAVED_FOR)
        if (store.overlay() == failed) store.setOverlay(null)
        redraw()
        return false
    }

    private suspend fun put(guild: Long, raid: Long, characterKey: String, state: SignupState): Boolean {
        val view = (raids.raid(guild, raid) as? ApiResult.Success)?.data ?: return false
        val note = view.signups.firstOrNull { it.characterKey == characterKey }?.note.orEmpty()
        return raids.signup(guild, raid, characterKey, SignupRequest(state = state.wire, note = note)) is ApiResult.Success
    }

    companion object {
        val NOT_SAVED_FOR: Duration = Duration.ofSeconds(10)
    }
}
