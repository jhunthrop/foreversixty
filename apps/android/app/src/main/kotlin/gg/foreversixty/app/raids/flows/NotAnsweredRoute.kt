package gg.foreversixty.app.raids.flows

import gg.foreversixty.app.designsystem.screenGround
import android.content.Intent
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.flows.FlowBackBar
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidLinks
import kotlinx.coroutines.launch

/**
 * Not answered on the Raid tab's stack: loads the raid, then the page. The mention, a Discord DM and
 * editing go to the site and Discord until the app can post a role's mention itself.
 */
@Composable
fun NotAnsweredRoute(
    model: NotAnsweredModel,
    clock: RaidClock,
    guildName: String,
    guildUrl: String?,
    build: String?,
    openWeb: (String) -> Unit,
    back: () -> Unit,
) {
    val res = LocalResources.current
    val context = LocalContext.current
    val scope = androidx.compose.runtime.rememberCoroutineScope()
    LaunchedEffect(model) { model.load() }
    val share = RaidLinks.share(model.guild, model.raidId)
    val page = guildUrl?.let { "$it/raids/${model.raidId}" } ?: share
    Column(Modifier.fillMaxSize().screenGround()) {
        FlowBackBar(stringResource(R.string.tab_raids), back) {
            context.startActivity(Intent.createChooser(Intent(Intent.ACTION_SEND).setType("text/plain").putExtra(Intent.EXTRA_TEXT, share), null))
        }
        when (val phase = model.phase) {
            NotAnsweredModel.Phase.Loading -> Box(Modifier.fillMaxSize(), contentAlignment = Alignment.Center) {
                CircularProgressIndicator(color = FsColor.gold)
            }
            NotAnsweredModel.Phase.Failed -> FlowPage { FlowNote(stringResource(R.string.raid_failed)) }
            is NotAnsweredModel.Phase.Loaded -> {
                val content = NotAnsweredPresenter(res, clock, phase.raid, guildName, model.marks.toMap()).content(model.filter)
                NotAnsweredScreen(
                    content,
                    build,
                    NotAnsweredActions(
                        select = { model.filter = it },
                        mentionRole = { content.mention?.let { m -> scope.launch { model.mention(m, role = true) } } },
                        mentionAll = { content.mention?.let { m -> scope.launch { model.mention(m, role = false) } } },
                        mark = { row, state -> model.mark(row, state) },
                        undo = { model.undo(it) },
                        message = { row -> row.dmUrl?.let(openWeb) },
                        editOnSite = { openWeb(page) },
                    ),
                    model.error,
                    model.mentionNote,
                )
            }
        }
    }
}
