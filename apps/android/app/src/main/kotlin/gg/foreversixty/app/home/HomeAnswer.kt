package gg.foreversixty.app.home

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ColumnScope
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.KeyboardArrowRight
import androidx.compose.material.icons.filled.Refresh
import androidx.compose.material.icons.filled.Star
import androidx.compose.material.icons.filled.PlayArrow
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.items.ItemIcon
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemQuality
import java.util.Locale

/** What a tap on Home asks the screen to do. */
sealed interface HomeIntent {
    data class OpenItem(val look: ItemLook, val sourceLine: String?, val wornId: Int?) : HomeIntent
    data object OpenGear : HomeIntent
    data class Action(val kind: HomeAction.Kind) : HomeIntent
    data object Retry : HomeIntent
}

/** From font scale 1.5 the gain moves under the name (spec §10, XL.1). */
@Composable
private fun isLargeText(): Boolean = LocalDensity.current.fontScale >= LARGE_TEXT_SCALE

/**
 * Home's content under the bar, pure over its inputs so the screenshot tests draw every state
 * (spec F2.1 to F2.5). Scrolling is the caller's.
 */
@Composable
fun HomeAnswer(
    phase: HomePhase,
    presenter: HomePresenter,
    build: String,
    offlineStrip: String?,
    send: (HomeIntent) -> Unit = {},
) {
    Column(Modifier.fillMaxWidth()) {
        offlineStrip?.let { OfflineStrip(it) }
        when (phase) {
            HomePhase.Loading -> HomeSkeleton()
            is HomePhase.Loaded -> Answer(presenter.model(phase.home), build, send)
            is HomePhase.Failed -> Failed(presenter, phase.cached?.let(presenter::model), send)
        }
    }
}

@Composable
private fun Answer(model: HomeAnswerModel, build: String, send: (HomeIntent) -> Unit) {
    Hero(model, send)
    Column(Modifier.padding(start = 16.dp, end = 16.dp, bottom = 24.dp)) {
        val body = model.body
        if (body is HomeAnswerBody.Upgrades) {
            Spacer(Modifier.height(16.dp))
            UpgradesCard(body, model.bandHead, build, send)
        }
        model.already?.let {
            Text(it, style = HomeType.meta, color = FsColor.muted, modifier = Modifier.padding(top = 12.dp))
        }
        Spacer(Modifier.height(22.dp))
        NextRows(model.nextTitle, model.next)
    }
}

@Composable
private fun Failed(presenter: HomePresenter, cached: HomeAnswerModel?, send: (HomeIntent) -> Unit) {
    HeroColumn {
        Text(presenter.heroMeta, style = HomeType.meta, color = FsColor.muted)
        Spacer(Modifier.height(14.dp))
        MicroLabel(stringResource(R.string.home_answer_label), color = FsColor.nav)
        Text(
            stringResource(R.string.home_answer_failed),
            style = HomeType.lead,
            color = FsColor.wipe,
            modifier = Modifier.padding(vertical = 10.dp).semantics { liveRegion = LiveRegionMode.Assertive },
        )
        SecondaryButton(stringResource(R.string.common_tryAgain), onClick = { send(HomeIntent.Retry) }, withRetryGlyph = true)
    }
    Column(Modifier.padding(start = 16.dp, end = 16.dp, top = 22.dp, bottom = 24.dp)) {
        NextRows(
            cached?.nextTitle ?: stringResource(R.string.home_next_title, presenter.character.name),
            cached?.next ?: presenter.nextRowsWithoutAnswer(),
        )
    }
}

@Composable
private fun HeroColumn(content: @Composable ColumnScope.() -> Unit) {
    Column(
        Modifier.fillMaxWidth().heroBand().padding(start = 16.dp, end = 16.dp, top = 14.dp, bottom = 18.dp),
        content = content,
    )
}

/** The hero band: who, the label, the two figures, the note and the sync line. */
@Composable
private fun Hero(model: HomeAnswerModel, send: (HomeIntent) -> Unit) {
    HeroColumn {
        Text(model.heroMeta, style = HomeType.meta, color = FsColor.muted)
        Spacer(Modifier.height(14.dp))
        MicroLabel(stringResource(R.string.home_answer_label), color = FsColor.nav)
        Spacer(Modifier.height(8.dp))
        when (val body = model.body) {
            is HomeAnswerBody.Upgrades -> Figures(model)
            is HomeAnswerBody.AllMatch -> Text(body.text, style = HomeType.lead, color = FsColor.strong)
            is HomeAnswerBody.Message -> Message(body, send)
        }
        model.note?.let { Text(it, style = HomeType.meta, color = FsColor.muted, modifier = Modifier.padding(top = 6.dp)) }
        model.synced?.let {
            Row(Modifier.padding(top = 10.dp), verticalAlignment = Alignment.CenterVertically) {
                Icon(Icons.Filled.Refresh, contentDescription = null, tint = FsColor.nav, modifier = Modifier.size(14.dp))
                Spacer(Modifier.width(6.dp))
                Text(it, style = HomeType.meta, color = FsColor.nav)
            }
        }
    }
}

/** The two figures at one size: `11 upgrades` and `+76.2 DPS`. */
@Composable
private fun Figures(model: HomeAnswerModel) {
    val spoken = listOfNotNull(model.count, model.gainFigure?.let { "$it ${model.gainUnit.orEmpty()}" }).joinToString(", ")
    FlowRow(
        horizontalArrangement = Arrangement.spacedBy(22.dp),
        modifier = Modifier.clearAndSetSemantics { contentDescription = spoken },
    ) {
        model.count?.let { count ->
            val number = count.substringBefore(' ')
            Row(verticalAlignment = Alignment.Bottom) {
                Text(number, style = HomeType.stat, color = FsColor.strong)
                Spacer(Modifier.width(6.dp))
                Text(count.substringAfter(' '), style = HomeType.body, color = FsColor.muted, modifier = Modifier.padding(bottom = 4.dp))
            }
        }
        model.gainFigure?.let { figure ->
            Row(verticalAlignment = Alignment.Bottom) {
                Text(figure, style = HomeType.stat, color = FsColor.strong)
                Spacer(Modifier.width(6.dp))
                Text(model.gainUnit.orEmpty(), style = HomeType.body, color = FsColor.muted, modifier = Modifier.padding(bottom = 4.dp))
            }
        }
    }
}

@Composable
private fun Message(body: HomeAnswerBody.Message, send: (HomeIntent) -> Unit) {
    Text(body.text, style = HomeType.lead, color = FsColor.strong)
    if (body.buttons.isNotEmpty()) {
        FlowRow(
            horizontalArrangement = Arrangement.spacedBy(10.dp),
            verticalArrangement = Arrangement.spacedBy(10.dp),
            modifier = Modifier.padding(top = 14.dp),
        ) {
            body.buttons.forEach { button ->
                if (button.primary) {
                    PrimaryButton(button.title) { send(HomeIntent.Action(button.kind)) }
                } else {
                    SecondaryButton(button.title, onClick = { send(HomeIntent.Action(button.kind)) })
                }
            }
        }
    }
    body.note?.let { Text(it, style = HomeType.meta, color = FsColor.muted, modifier = Modifier.padding(top = 14.dp)) }
}

/** The Your upgrades card: head row, five rows in the server's order, the footer to Your gear. */
@Composable
private fun UpgradesCard(body: HomeAnswerBody.Upgrades, bandHead: String?, build: String, send: (HomeIntent) -> Unit) {
    val hint = stringResource(R.string.items_a11y_hint)
    Column(Modifier.fillMaxWidth().primaryCard()) {
        Row(Modifier.fillMaxWidth().padding(horizontal = 16.dp, vertical = 12.dp), verticalAlignment = Alignment.CenterVertically) {
            MicroLabel(stringResource(R.string.home_upgrades_title))
            Spacer(Modifier.weight(1f))
            bandHead?.let { Text(it, style = HomeType.meta, color = FsColor.muted) }
        }
        body.rows.forEach { row ->
            HorizontalDivider(color = FsColor.lineSoft)
            UpgradeRow(
                row, build,
                Modifier.clickable(onClickLabel = hint, role = Role.Button) {
                    send(HomeIntent.OpenItem(row.pick, row.sourceLine, row.wornId))
                },
            )
        }
        HorizontalDivider(color = FsColor.lineSoft)
        Row(
            Modifier.fillMaxWidth().heightIn(min = 52.dp).clickable(role = Role.Button) { send(HomeIntent.OpenGear) }
                .padding(horizontal = 16.dp),
            verticalAlignment = Alignment.CenterVertically,
        ) {
            Text(body.footer, style = HomeType.bodyStrong, color = FsColor.gold, modifier = Modifier.weight(1f))
            Icon(Icons.AutoMirrored.Filled.KeyboardArrowRight, contentDescription = null, tint = FsColor.gold)
        }
    }
}

/** One upgrade row: the pick's icon, name, slot and source, what it replaces, and the gain. */
@Composable
fun UpgradeRow(row: HomeUpgradeRow, build: String, modifier: Modifier = Modifier) {
    val gainBelow = isLargeText()
    Row(
        modifier.fillMaxWidth().heightIn(min = 72.dp).padding(horizontal = 16.dp, vertical = 12.dp)
            .semantics(mergeDescendants = true) {},
        horizontalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        ItemIcon(row.pick.icon, row.pick.quality, build, 44.dp)
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
            Text(row.pick.name ?: stringResource(R.string.items_notInData), style = HomeType.body, color = ItemQuality(row.pick.quality).text)
            Text(row.meta, style = HomeType.meta, color = FsColor.muted)
            Text(row.replaces, style = HomeType.meta, color = FsColor.muted)
            when (val gain = row.gain) {
                is HomeUpgradeRow.Gain.Words -> Text(gain.text, style = HomeType.meta, color = FsColor.muted)
                is HomeUpgradeRow.Gain.Measured -> if (gainBelow) {
                    Text("${gain.figure} ${gain.unit}", style = HomeType.mono, color = FsColor.strong)
                }
            }
        }
        val gain = row.gain
        if (gain is HomeUpgradeRow.Gain.Measured && !gainBelow) {
            Column(horizontalAlignment = Alignment.End) {
                Text(gain.figure, style = HomeType.mono, color = FsColor.strong)
                Text(gain.unit.uppercase(Locale.ROOT), style = HomeType.micro, color = FsColor.muted)
            }
        }
    }
}

/** "Next for {name}": Talents and Sims (spec F2.1 item 4). */
@Composable
private fun NextRows(title: String, rows: List<HomeNextRow>) {
    MicroLabel(title)
    Spacer(Modifier.height(8.dp))
    Column(Modifier.fillMaxWidth().quietCard()) {
        rows.forEachIndexed { index, row ->
            if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
            NextRow(row)
        }
    }
}

@Composable
private fun NextRow(row: HomeNextRow) {
    Row(
        Modifier.fillMaxWidth().heightIn(min = 56.dp).padding(horizontal = 16.dp, vertical = 10.dp)
            .semantics(mergeDescendants = true) {},
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        Box(
            Modifier.size(36.dp).background(FsColor.gold.copy(alpha = 0.12f), RoundedCornerShape(FsRadius.pill)),
            contentAlignment = Alignment.Center,
        ) {
            Icon(
                if (row.kind == HomeNextRow.Kind.Talents) Icons.Filled.Star else Icons.Filled.PlayArrow,
                contentDescription = null,
                tint = FsColor.goldHover,
                modifier = Modifier.size(18.dp),
            )
        }
        Column(Modifier.weight(1f)) {
            Text(row.title, style = HomeType.bodyStrong, color = FsColor.strong)
            Text(row.line, style = HomeType.meta, color = FsColor.muted)
        }
        when (row.tag) {
            HomeNextRow.Tag.Matches -> Tag(stringResource(R.string.home_next_matches), FsColor.kill)
            HomeNextRow.Tag.Differs -> Tag(stringResource(R.string.home_next_differs), FsColor.gold)
            null -> Unit
        }
    }
}

/** F2.2: the band and the card as skeletons at their real heights, five rows, and the line. */
@Composable
fun HomeSkeleton() {
    val loading = stringResource(R.string.home_answer_loading)
    Column(Modifier.fillMaxWidth().clearAndSetSemantics { contentDescription = loading; liveRegion = LiveRegionMode.Polite }) {
        HeroColumn {
            Bar(220, 13)
            Spacer(Modifier.height(14.dp))
            Bar(160, 11)
            Spacer(Modifier.height(10.dp))
            Bar(200, 30)
            Spacer(Modifier.height(10.dp))
            Bar(260, 13)
            Spacer(Modifier.height(10.dp))
            Bar(180, 13)
        }
        Column(Modifier.padding(start = 16.dp, end = 16.dp, top = 16.dp).fillMaxWidth().quietCard().padding(vertical = 8.dp)) {
            repeat(HomePresenter.VISIBLE_ROWS) {
                Row(Modifier.height(72.dp).padding(horizontal = 16.dp, vertical = 12.dp), horizontalArrangement = Arrangement.spacedBy(12.dp)) {
                    Box(Modifier.size(44.dp).background(FsColor.cardTop, RoundedCornerShape(FsRadius.pill)))
                    Column(verticalArrangement = Arrangement.spacedBy(6.dp)) {
                        Bar(180, 14)
                        Bar(140, 11)
                        Bar(200, 11)
                    }
                }
            }
        }
        Text(loading, style = HomeType.meta, color = FsColor.muted, modifier = Modifier.padding(16.dp))
    }
}

@Composable
private fun Bar(width: Int, height: Int) {
    Box(Modifier.width(width.dp).height(height.dp).background(FsColor.cardTop, RoundedCornerShape(4.dp)))
}
