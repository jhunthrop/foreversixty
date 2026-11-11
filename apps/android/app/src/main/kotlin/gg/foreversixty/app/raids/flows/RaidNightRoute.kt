package gg.foreversixty.app.raids.flows

import gg.foreversixty.app.designsystem.screenGround
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.runtime.Composable
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.flows.FlowBackBar
import gg.foreversixty.app.model.RaidCompTarget
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidNightModel
import gg.foreversixty.app.raids.ui.RaidNightScreen
import java.time.Instant
import kotlinx.coroutines.launch

/**
 * Raid night (spec F2) from `GET /v1/me/raid-nights/{raid}`: before and during the raid the raid page on
 * raid night (answer, who is coming, Before you pull); after it, Your raid night (the log's summary).
 */
@Composable
fun RaidNightRoute(
    model: RaidNightModel,
    clock: RaidClock,
    now: () -> Instant,
    target: RaidCompTarget?,
    build: String?,
    openReport: ((String) -> Unit)?,
    openRoster: () -> Unit,
    openWeb: (String) -> Unit,
    back: () -> Unit,
) {
    val night = (model.state as? RaidNightModel.State.Loaded)?.night
    if (night == null || night.phase == AFTER) {
        RaidNightScreen(model, clock, now, openReport, back)
        return
    }
    val res = LocalResources.current
    val scope = rememberCoroutineScope()
    Column(Modifier.fillMaxSize().screenGround()) {
        FlowBackBar(stringResource(R.string.tab_raids), back)
        RaidNightPageScreen(
            RaidNightPagePresenter(res, clock, now()).content(night, target, model.pressed, model.lateOpen),
            build,
            RaidNightPageActions(
                answer = { state -> scope.launch { model.answer(state, res = res) } },
                late = { minutes -> scope.launch { model.answer(gg.foreversixty.app.raids.logic.SignupState.Late, minutes, res) } },
                roster = openRoster,
                where = { item -> item.link?.let(openWeb) },
            ),
            model.answerError,
        )
    }
}

private const val AFTER = "after"
