package gg.foreversixty.app.raids.flows

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.raids.logic.PlannedRaids
import gg.foreversixty.app.raids.logic.RaidsTabModel
import gg.foreversixty.app.raids.logic.RaidsTabState
import gg.foreversixty.app.raids.logic.SignupController
import java.time.Instant

/**
 * The next raid's sign-up on a tab root (the Raid tab's head, Home's raid card): one
 * [SignupController] per next raid of [tab], offline with the tab, and every newer raid view shared
 * both ways. A write on the card reaches the tab's list (and [onChange]); a write on the raid page
 * ([replaced]) reaches the list and the card's controller, so the card never shows an older answer.
 */
class NextRaidSignup(
    private val tab: RaidsTabModel,
    private val now: () -> Instant,
    private val onChange: (RaidView) -> Unit = {},
) {
    var controller: SignupController? by mutableStateOf(null)
        private set

    /** After the tab loads: the controller follows the next raid, a new one when the raid changed. */
    fun sync(deps: SignupController.Dependencies) {
        val content = (tab.state as? RaidsTabState.Loaded)?.content
        val raid = content?.let { PlannedRaids(it.planned, now()).next }
        val current = controller
        controller = when {
            content == null || raid == null -> null
            current != null && current.raid.id == raid.id -> current.also { it.replace(raid) }
            else -> SignupController(content.head.id, raid, deps).also { it.onRaidChanged = ::publish }
        }
        controller?.offline = content?.offlineSince != null
    }

    /** A newer view from elsewhere (the raid page's write or refetch). */
    fun replaced(view: RaidView) {
        val held = controller?.takeIf { it.raid.id == view.id }
        // The controller takes it unless a write of its own is out; taking it publishes it.
        held?.replace(view)
        if (held?.raid != view) publish(view)
    }

    private fun publish(view: RaidView) {
        tab.replace(view)
        onChange(view)
    }
}
