package gg.foreversixty.app.widgets.live

import gg.foreversixty.app.model.LiveActivityRequest
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.widgets.data.LiveKind
import gg.foreversixty.app.widgets.data.WidgetService
import gg.foreversixty.app.widgets.data.WidgetStore

/**
 * Keeps the server told which live surfaces this phone shows (W3-C13): POST when a surface starts,
 * DELETE when it ends. The token is this phone's FCM registration token.
 *
 * Today [pushToken] answers null: FCM is not wired into the app yet (lane F7, see the manifest),
 * and the server has no live.* sender yet either (lane A2). So nothing is registered, and every
 * ongoing notification is drawn and updated by the app while it runs, from the snapshot and the
 * Sims poll. When F7 lands, pass its token provider here; nothing else changes.
 */
class LiveRegistrations(
    private val service: WidgetService,
    private val store: WidgetStore,
    private val pushToken: suspend () -> String? = { null },
) {
    suspend fun start(kind: LiveKind, objectId: String) {
        val key = key(kind, objectId)
        val current = store.registrations()
        if (key in current) return
        val token = pushToken() ?: return
        val request = LiveActivityRequest(kind = kind.wire, platform = PLATFORM, pushToken = token, objectId = objectId)
        val answer = service.register(request) as? ApiResult.Success ?: return
        // One surface of a kind at a time: a new raid's countdown replaces the last one's.
        val (previous, others) = current.entries.partition { it.key.startsWith(kind.wire + SEPARATOR) }
        previous.forEach { service.unregister(it.value) }
        store.setRegistrations(others.associate { it.key to it.value } + (key to answer.data.id))
    }

    suspend fun end(kind: LiveKind) {
        val current = store.registrations()
        val ending = current.filterKeys { it.startsWith(kind.wire + SEPARATOR) }
        if (ending.isEmpty()) return
        ending.values.forEach { service.unregister(it) }
        store.setRegistrations(current - ending.keys)
    }

    suspend fun endAll() {
        LiveKind.entries.forEach { end(it) }
    }

    companion object {
        private const val PLATFORM = "android"
        private const val SEPARATOR = ":"

        fun key(kind: LiveKind, objectId: String) = kind.wire + SEPARATOR + objectId
    }
}
