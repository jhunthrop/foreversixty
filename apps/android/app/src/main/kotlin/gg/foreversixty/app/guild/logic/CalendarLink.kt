package gg.foreversixty.app.guild.logic

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.auth.SecretStore
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.logic.SignupRefusal

/** The calendar link row (spec F16.4 part 2). */
sealed interface CalendarLinkState {
    data object Loading : CalendarLinkState

    /** The link this phone may copy. Never logged. */
    data class Shown(val url: String) : CalendarLinkState

    /** The API shows a link once; it was shown on another phone or browser. Only Reset. */
    data object Elsewhere : CalendarLinkState

    /** Nothing loaded and nothing stored. */
    data object Failed : CalendarLinkState
}

/** The line under the buttons. */
enum class CalendarLinkNotice { Copied, LoadFailed, ResetFailed }

/**
 * `GET /v1/me/calendar-feed` returns the url once; this keeps it in the secure store keyed per
 * account ([accountKey]) and falls back to the stored one. Writes are never queued offline.
 */
class CalendarLinkModel(
    private val service: GuildService,
    private val secrets: SecretStore,
    accountKey: String,
) {
    private val key = "$KEY_PREFIX$accountKey"

    var state: CalendarLinkState by mutableStateOf(CalendarLinkState.Loading)
        private set
    var notice: CalendarLinkNotice? by mutableStateOf(null)
        private set
    var offline: Boolean by mutableStateOf(false)
        private set
    var resetting: Boolean by mutableStateOf(false)
        private set

    suspend fun load() {
        val stored = secrets.read(key)
        when (val result = service.calendarFeed()) {
            is ApiResult.Success -> {
                offline = false
                val url = result.data.url
                if (url != null) secrets.write(key, url)
                state = (url ?: stored)?.let { CalendarLinkState.Shown(it) } ?: CalendarLinkState.Elsewhere
            }
            is ApiResult.Failure -> {
                offline = SignupRefusal.isOffline(result.error)
                state = stored?.let { CalendarLinkState.Shown(it) } ?: CalendarLinkState.Failed
                if (stored == null) notice = CalendarLinkNotice.LoadFailed
            }
        }
    }

    /** After the player confirms `Reset`: a new link, the old one revoked. */
    suspend fun reset() {
        resetting = true
        when (val result = service.resetCalendarFeed()) {
            is ApiResult.Success -> {
                val url = result.data.url
                if (url != null) {
                    secrets.write(key, url)
                    state = CalendarLinkState.Shown(url)
                    notice = null
                } else {
                    notice = CalendarLinkNotice.ResetFailed
                }
            }
            is ApiResult.Failure -> {
                offline = SignupRefusal.isOffline(result.error)
                notice = CalendarLinkNotice.ResetFailed
            }
        }
        resetting = false
    }

    /** The url to put on the clipboard, and `Link copied.` under the buttons. */
    fun copy(): String? = (state as? CalendarLinkState.Shown)?.url?.also { notice = CalendarLinkNotice.Copied }

    companion object {
        const val KEY_PREFIX = "calendar-link-"
    }
}
