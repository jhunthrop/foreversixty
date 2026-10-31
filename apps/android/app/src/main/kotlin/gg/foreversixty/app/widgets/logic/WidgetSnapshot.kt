package gg.foreversixty.app.widgets.logic

import gg.foreversixty.app.model.Widget
import gg.foreversixty.app.raids.logic.SignupState
import java.time.Instant
import kotlinx.serialization.Serializable

/**
 * The one snapshot every live surface draws from (spec 2026-10-31-mobile-wave3 §5.6, W3-C12):
 * `GET /v1/me/widget` as last fetched, and when. The widgets, the ongoing raid notification and
 * the quick-settings tile all read this through [WidgetFigures], so their figures cannot disagree.
 * No snapshot means signed out.
 */
@Serializable
data class WidgetSnapshot(val widget: Widget, val fetchedAt: String) {
    val fetchedInstant: Instant get() = Instant.parse(fetchedAt)
}

/**
 * A press on a widget's In, Late or Out that the server has not answered yet ([pending]), or one it
 * refused ([failedUntil]: the button says `Not saved` until then and the snapshot's answer stands).
 */
@Serializable
data class SignupOverlay(val raidId: Long, val state: String, val failedUntil: String? = null) {
    val signupState: SignupState? get() = SignupState.of(state)

    fun isFailed(now: Instant): Boolean = failedUntil?.let { now < Instant.parse(it) } == true

    val isPending: Boolean get() = failedUntil == null
}
