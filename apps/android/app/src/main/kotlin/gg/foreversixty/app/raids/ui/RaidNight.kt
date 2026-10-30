package gg.foreversixty.app.raids.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.ArrowBack
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.material3.TopAppBarDefaults
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.shadow
import androidx.compose.ui.platform.LocalUriHandler
import androidx.compose.ui.res.pluralStringResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.model.RaidNightBest
import gg.foreversixty.app.model.RaidNightLoot
import gg.foreversixty.app.model.RaidNightSummary
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidLinks
import gg.foreversixty.app.raids.logic.RaidNightModel
import gg.foreversixty.app.raids.logic.roleLabel
import java.time.Instant
import java.util.Locale
import kotlinx.coroutines.launch

/** F4.22 Your raid night: kills, wipes and bests, loot awarded to you, personal bests, coach notes. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun RaidNightScreen(model: RaidNightModel, clock: RaidClock, now: () -> Instant, back: () -> Unit) {
    val scope = rememberCoroutineScope()
    LaunchedEffect(model) { model.load() }
    Column(Modifier.fillMaxSize().background(FsColor.bg)) {
        TopAppBar(
            title = { Text(stringResource(R.string.night_title).uppercase(), style = RaidType.sec, color = FsColor.strong) },
            navigationIcon = {
                IconButton(onClick = back) { Icon(Icons.AutoMirrored.Filled.ArrowBack, contentDescription = null, tint = FsColor.strong) }
            },
            colors = TopAppBarDefaults.topAppBarColors(containerColor = FsColor.bg),
        )
        Column(Modifier.fillMaxSize().verticalScroll(rememberScrollState())) {
            model.offlineSince?.let { OfflineStrip(stringResource(R.string.offline_strip, clock.stripTime(it, now()))) }
            Column(verticalArrangement = Arrangement.spacedBy(22.dp), modifier = Modifier.padding(16.dp)) {
                when (val state = model.state) {
                    RaidNightModel.State.Loading -> Column(verticalArrangement = Arrangement.spacedBy(14.dp), modifier = Modifier.clearAndSetSemantics {}) {
                        Box(Modifier.width(220.dp).height(13.dp).background(FsColor.lineSoft, RoundedCornerShape(4.dp)))
                        Box(Modifier.fillMaxWidth().height(110.dp).background(FsColor.lineSoft, RoundedCornerShape(FsRadius.panel)))
                        Box(Modifier.fillMaxWidth().height(160.dp).background(FsColor.lineSoft, RoundedCornerShape(FsRadius.panel)))
                    }
                    RaidNightModel.State.Failed -> ErrorRegion(stringResource(R.string.night_failed)) { scope.launch { model.load() } }
                    is RaidNightModel.State.Loaded -> NightContent(state.night, clock)
                }
            }
        }
    }
}

@Composable
private fun NightContent(night: RaidNightSummary, clock: RaidClock) {
    val uri = LocalUriHandler.current
    val startsAt = Instant.parse(night.startsAt)
    Text(
        stringResource(R.string.night_meta, night.zoneName, clock.weekday(startsAt), clock.dayMonth(startsAt), night.guildName),
        style = RaidType.meta,
        color = FsColor.muted,
    )
    Row(horizontalArrangement = Arrangement.spacedBy(16.dp), verticalAlignment = Alignment.CenterVertically, modifier = Modifier.primaryCard()) {
        Figure(night.kills, pluralStringResource(R.plurals.night_kill, night.kills, night.kills))
        Figure(night.wipes, pluralStringResource(R.plurals.night_wipes, night.wipes, night.wipes))
        Figure(night.newBests, pluralStringResource(R.plurals.night_best, night.newBests, night.newBests))
        Column(horizontalAlignment = Alignment.End, modifier = Modifier.weight(1f)) {
            Text(night.character.name, style = RaidType.name, color = RaidColors.classColor(night.character.className))
            Text(
                stringResource(R.string.raids_who, night.character.specName, stringResource(roleLabel(night.character.role))),
                style = RaidType.meta,
                color = FsColor.muted,
            )
        }
    }
    if (night.loot.isNotEmpty()) Section(stringResource(R.string.night_loot)) { night.loot.forEachIndexed { i, loot -> if (i > 0) Divider(); LootRow(loot) } }
    if (night.bests.isNotEmpty()) Section(stringResource(R.string.night_bests)) { night.bests.forEachIndexed { i, best -> if (i > 0) Divider(); BestRow(best) } }
    night.coachNotes?.takeIf { it.isNotEmpty() }?.let { notes ->
        Section(stringResource(R.string.night_coach)) {
            notes.forEachIndexed { index, note ->
                if (index > 0) Divider()
                Row(horizontalArrangement = Arrangement.spacedBy(10.dp), modifier = Modifier.padding(14.dp)) {
                    Text("${index + 1}", style = RaidType.mono, color = FsColor.gold)
                    Text(note, style = RaidType.body, color = FsColor.text)
                }
            }
        }
    }
    GoldLink(stringResource(R.string.night_report), { uri.openUri(RaidLinks.report(night.reportId)) })
}

@Composable
private fun Divider() = HorizontalDivider(color = FsColor.lineSoft)

/** The one figure: the number in `stat`, its word in meta (the plural carries the count). */
@Composable
private fun Figure(value: Int, words: String) {
    Column(Modifier.clearAndSetSemantics { contentDescription = words }) {
        Text("$value", style = RaidType.stat, color = FsColor.strong)
        Text(words.removePrefix("$value ").trim(), style = RaidType.meta, color = FsColor.muted)
    }
}

@Composable
private fun Section(title: String, rows: @Composable () -> Unit) {
    Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
        SectionLabel(title)
        Column(Modifier.quietCard()) { rows() }
    }
}

@Composable
private fun LootRow(loot: RaidNightLoot) {
    val quality = RaidColors.qualityBorder(loot.quality)
    Row(
        horizontalArrangement = Arrangement.spacedBy(12.dp),
        modifier = Modifier.fillMaxWidth().heightIn(min = 72.dp).padding(14.dp).semantics(mergeDescendants = true) {},
    ) {
        Box(Modifier.shadow(5.dp, RoundedCornerShape(FsRadius.pill), ambientColor = quality, spotColor = quality)) {
            GameIcon(loot.icon, 44.dp, border = quality, borderWidth = 1.5.dp)
        }
        Column(verticalArrangement = Arrangement.spacedBy(2.dp)) {
            Text(loot.name, style = RaidType.bodyStrong, color = RaidColors.qualityText(loot.quality))
            Text(stringResource(R.string.night_lootLine, loot.slot, loot.boss, loot.awardedBy), style = RaidType.meta, color = FsColor.muted)
            loot.replaces?.let {
                Text(stringResource(R.string.home_upgrades_replaces, it.name, it.itemLevel), style = RaidType.meta, color = FsColor.muted)
            }
        }
    }
}

@Composable
private fun BestRow(best: RaidNightBest) {
    Row(Modifier.fillMaxWidth().padding(14.dp).semantics(mergeDescendants = true) {}) {
        Column(Modifier.weight(1f)) {
            Text(best.boss, style = RaidType.bodyStrong, color = FsColor.strong)
            Text(stringResource(R.string.night_bestLine, killTime(best.killMs)), style = RaidType.meta, color = FsColor.muted)
        }
        Column(horizontalAlignment = Alignment.End) {
            Text(stringResource(R.string.raids_night_dps, String.format(Locale.US, "%.1f", best.dps)), style = RaidType.mono, color = FsColor.strong)
            Text(stringResource(R.string.night_percentile, ordinal(best.percentile)), style = RaidType.meta, color = RaidColors.parse(best.percentile))
        }
    }
}

/** `4:28` from milliseconds. */
fun killTime(ms: Long): String {
    val seconds = ms / MILLIS
    return String.format(Locale.US, "%d:%02d", seconds / SIXTY, seconds % SIXTY)
}

/** The English ordinal: `91st`, `92nd`, `93rd`, `94th`, `11th`. */
fun ordinal(n: Int): String {
    val suffix = when {
        n % HUNDRED in TEENS -> "th"
        n % TEN == 1 -> "st"
        n % TEN == 2 -> "nd"
        n % TEN == 3 -> "rd"
        else -> "th"
    }
    return "$n$suffix"
}

private const val MILLIS = 1000L
private const val SIXTY = 60L
private const val TEN = 10
private const val HUNDRED = 100
private val TEENS = 11..13
