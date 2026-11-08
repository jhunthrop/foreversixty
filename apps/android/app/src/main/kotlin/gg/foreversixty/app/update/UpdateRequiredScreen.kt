package gg.foreversixty.app.update

import androidx.compose.runtime.Composable
import androidx.compose.ui.res.stringResource
import gg.foreversixty.app.R
import gg.foreversixty.app.flows.FlowSealPage
import gg.foreversixty.app.flows.FlowStackButton

/** Play's page for this app, when the gate sent no https link of its own. */
const val PLAY_STORE_URL = "https://play.google.com/store/apps/details?id=gg.foreversixty.app"

/** The blocking screen the update gate shows (mock `update-required`): no tabs, no way past it but the store. */
@Composable
fun UpdateRequiredScreen(onOpenStore: () -> Unit) {
    FlowSealPage(stringResource(R.string.update_required_title), stringResource(R.string.update_required_body)) {
        FlowStackButton(stringResource(R.string.update_required_action), onOpenStore, primary = true)
    }
}
