package gg.foreversixty.app.raids.flows

import gg.foreversixty.app.designsystem.screenGround
import android.content.ClipData
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.ClipEntry
import androidx.compose.ui.platform.LocalClipboard
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.flows.FlowBackBar
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.raids.logic.RaidClock
import java.time.Instant
import kotlinx.coroutines.launch

/** The applicant on the Raid tab's stack: loads the application, then the page; Decline closes it. */
@Composable
fun ApplicantRoute(model: ApplicantModel, clock: RaidClock, guildName: String, now: () -> Instant, back: () -> Unit) {
    val res = LocalResources.current
    val clipboard = LocalClipboard.current
    val scope = rememberCoroutineScope()
    LaunchedEffect(model) { model.load() }
    Column(Modifier.fillMaxSize().screenGround()) {
        FlowBackBar(stringResource(R.string.tab_guild), back)
        when (val phase = model.phase) {
            ApplicantModel.Phase.Loading -> Box(Modifier.fillMaxSize(), contentAlignment = Alignment.Center) {
                CircularProgressIndicator(color = FsColor.gold)
            }
            ApplicantModel.Phase.Failed -> FlowPage { FlowNote(stringResource(R.string.raid_applicant_failed)) }
            is ApplicantModel.Phase.Loaded -> ApplicantScreen(
                ApplicantPresenter(res, clock, now()).content(phase.application, guildName, model.accepted),
                ApplicantActions(
                    accept = { scope.launch { model.accept() } },
                    decline = { scope.launch { if (model.decline()) back() } },
                    copy = { text -> scope.launch { clipboard.setClipEntry(ClipEntry(ClipData.newPlainText(text, text))) } },
                ),
                model.busy,
                model.error,
            )
        }
    }
}
