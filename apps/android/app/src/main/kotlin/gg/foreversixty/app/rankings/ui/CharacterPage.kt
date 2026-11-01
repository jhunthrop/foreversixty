package gg.foreversixty.app.rankings.ui

import androidx.compose.foundation.Canvas
import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Share
import androidx.compose.material.icons.filled.Star
import androidx.compose.material.icons.materialIcon
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.alpha
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.Path
import androidx.compose.ui.graphics.SolidColor
import androidx.compose.ui.graphics.vector.ImageVector
import androidx.compose.ui.graphics.vector.addPathNodes
import androidx.compose.ui.graphics.StrokeCap
import androidx.compose.ui.graphics.drawscope.Stroke
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.em
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.MicroLabel
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.Tag
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.primaryCard
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.items.ItemNumbers
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.model.CharacterFight
import gg.foreversixty.app.model.CharacterProfile
import gg.foreversixty.app.model.Item
import gg.foreversixty.app.paperdoll.Paperdoll
import gg.foreversixty.app.paperdoll.PaperdollFrame
import gg.foreversixty.app.paperdoll.PaperdollSlot
import gg.foreversixty.app.rankings.logic.CharacterCard
import gg.foreversixty.app.rankings.logic.CharacterCopy
import gg.foreversixty.app.rankings.logic.CharacterPhase
import gg.foreversixty.app.rankings.logic.CharacterRoute
import gg.foreversixty.app.rankings.logic.RankingsViewer
import gg.foreversixty.app.rankings.logic.ScoreBar
import gg.foreversixty.app.rankings.logic.WatchState
import gg.foreversixty.app.reports.logic.ParseTier
import gg.foreversixty.app.reports.ui.Chevron
import gg.foreversixty.app.reports.ui.ErrorRegion
import gg.foreversixty.app.reports.ui.OutcomeDot
import gg.foreversixty.app.reports.ui.ReportsBar
import gg.foreversixty.app.reports.ui.SkeletonBar
import gg.foreversixty.app.reports.ui.SkeletonDot
import gg.foreversixty.app.reports.ui.isLargeText
import gg.foreversixty.app.reports.ui.parseColor
import java.time.Instant
import java.time.ZoneId
import kotlin.math.roundToInt

/** What a character page draws, apart from its store so screenshots drive every state. */
data class CharacterScreenState(
    val route: CharacterRoute,
    val phase: CharacterPhase,
    val watch: WatchState?,
    val viewer: RankingsViewer,
    /** The character's class items file, by id, for the gear's icons and quality. */
    val items: Map<Int, Item> = emptyMap(),
    /** The data build the icons come from. */
    val build: String? = null,
    /** The `reports` switch: a recent fight opens its report. */
    val reportsOn: Boolean = false,
)

data class CharacterActions(
    val back: () -> Unit = {},
    val retry: () -> Unit = {},
    val toggleWatch: () -> Unit = {},
    val share: () -> Unit = {},
    val openHow: () -> Unit = {},
    val openSocket: (PaperdollSlot) -> Unit = {},
    val openFight: (CharacterFight) -> Unit = {},
    val openHistory: () -> Unit = {},
    val searchAgain: () -> Unit = {},
)

private object CharacterLayout {
    val crest = 64.dp
    val bar = 8.dp
    val barLabel = 104.dp
    val sparkWidth = 120.dp
    val sparkHeight = 40.dp
    val scoreSize = 36.sp
}

/** F11.4 to F11.7: one character's page. */
@Composable
fun CharacterContent(state: CharacterScreenState, specs: SpecCatalog, now: Instant, actions: CharacterActions = CharacterActions(), zone: ZoneId = ZoneId.systemDefault()) {
    val copy = CharacterCopy(LocalResources.current, zone)
    val loaded = state.phase as? CharacterPhase.Loaded
    val offline = loaded?.offlineSince
    Column(Modifier.fillMaxWidth().padding(bottom = 24.dp)) {
        val title = loaded?.profile?.character?.name ?: state.route.name
        ReportsBar(title, stringResource(R.string.rank_title), actions.back) {
            if (loaded != null) {
                state.watch?.let { WatchStar(it, title, enabled = offline == null, actions.toggleWatch) }
                IconButton(onClick = actions.share) { Icon(Icons.Filled.Share, contentDescription = stringResource(R.string.char_share), tint = FsColor.strong) }
            }
        }
        offline?.let { OfflineStrip(stringResource(R.string.offline_strip, OfflineStamp.format(it, now, zone))) }
        Column(Modifier.padding(horizontal = 16.dp).padding(top = 8.dp), verticalArrangement = Arrangement.spacedBy(22.dp)) {
            when (val phase = state.phase) {
                CharacterPhase.Loading -> Skeleton()
                CharacterPhase.Failed -> ErrorRegion(stringResource(R.string.char_error), actions.retry, inset = false)
                CharacterPhase.NotFound -> NotFound(state.route.name, actions)
                is CharacterPhase.Loaded -> Profile(phase.profile, state, specs, now, copy, actions)
            }
        }
    }
}

/** The Watch star (spec §7): outline muted, filled gold; 50% and disabled offline. */
@Composable
private fun WatchStar(watch: WatchState, name: String, enabled: Boolean, toggle: () -> Unit) {
    val label = stringResource(if (watch.watching) R.string.char_unwatchA11y else R.string.char_watchA11y, name)
    IconButton(onClick = toggle, enabled = enabled, modifier = Modifier.alpha(if (enabled) 1f else DISABLED_ALPHA)) {
        Icon(
            if (watch.watching) Icons.Filled.Star else StarOutline,
            contentDescription = label,
            tint = if (watch.watching) FsColor.gold else FsColor.muted,
        )
    }
}

@Composable
private fun Profile(profile: CharacterProfile, state: CharacterScreenState, specs: SpecCatalog, now: Instant, copy: CharacterCopy, actions: CharacterActions) {
    Head(profile, state, CharacterCard.specName(profile, specs), copy)
    ScoreCard(profile, copy, actions)
    if (profile.best.isNotEmpty()) Section(stringResource(R.string.char_bests)) { profile.best.forEachIndexed { i, f -> Divided(i) { BestRow(profile, f, now, copy) } } }
    if (profile.history.isNotEmpty()) {
        Section(stringResource(R.string.char_recent)) {
            profile.history.forEachIndexed { i, f -> Divided(i) { RecentRow(f, now, copy, state.reportsOn, actions) } }
        }
    }
    GearSection(profile, state, actions)
    if (profile.buildsSeen.isNotEmpty()) {
        Section(stringResource(R.string.char_builds)) {
            profile.buildsSeen.forEachIndexed { i, b ->
                Divided(i) {
                    Column(Modifier.fillMaxWidth().heightIn(min = 56.dp).padding(horizontal = 16.dp, vertical = 10.dp)) {
                        Text(b.talentSplit, style = FsType.mono, color = FsColor.strong)
                        Text(copy.buildMeta(b.spec, b.firstSeen), style = FsType.meta, color = FsColor.muted)
                    }
                }
            }
        }
    }
    if (CharacterCard.showsParseHistory(profile, state.viewer)) {
        Row(
            Modifier.fillMaxWidth().quietCard().clickable(role = Role.Button, onClick = actions.openHistory).heightIn(min = 56.dp).padding(horizontal = 16.dp),
            verticalAlignment = Alignment.CenterVertically,
        ) {
            Text(stringResource(R.string.char_parseHistory), style = FsType.bodyStrong, color = FsColor.strong, modifier = Modifier.weight(1f))
            Icon(painterResource(R.drawable.ic_open_external), contentDescription = null, tint = FsColor.muted, modifier = Modifier.size(18.dp))
        }
    }
}

/** Crest 64, the class-coloured name, the descriptor and where; `You` and the supporter mark. */
@Composable
private fun Head(profile: CharacterProfile, state: CharacterScreenState, specName: String?, copy: CharacterCopy) {
    val c = profile.character
    Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
        Row(horizontalArrangement = Arrangement.spacedBy(14.dp), verticalAlignment = Alignment.CenterVertically) {
            ClassCrest(c.className?.lowercase(), CharacterLayout.crest)
            Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
                Text(c.name.uppercase(), style = FsType.h2, color = classColor(c.className), modifier = Modifier.semantics { heading() })
                val tags = CharacterCard.isYours(profile, state.viewer) || profile.supporter
                if (tags) {
                    Row(horizontalArrangement = Arrangement.spacedBy(6.dp)) {
                        if (CharacterCard.isYours(profile, state.viewer)) Tag(stringResource(R.string.char_you), FsColor.gold)
                        if (profile.supporter) Tag(stringResource(R.string.char_supporter), FsColor.rarityEpicText)
                    }
                }
                Text(copy.descriptor(profile, specName), style = FsType.body, color = FsColor.muted)
                Text(copy.where(profile), style = FsType.body, color = FsColor.muted)
            }
        }
        if (state.watch?.failed == true) {
            Text(
                stringResource(R.string.char_watchFailed, c.name), style = FsType.meta, color = FsColor.wipe,
                modifier = Modifier.semantics { liveRegion = LiveRegionMode.Polite },
            )
        }
    }
}

/** The Forever Score (spec F11.4); a character with no rated fight shows `No logs yet.`, never a zero. */
@Composable
private fun ScoreCard(profile: CharacterProfile, copy: CharacterCopy, actions: CharacterActions) {
    val score = profile.score
    if (score == null) {
        Column(Modifier.fillMaxWidth().quietCard().padding(16.dp), verticalArrangement = Arrangement.spacedBy(6.dp)) {
            Text(stringResource(R.string.char_noLogs), style = FsType.lead, color = FsColor.strong)
            Text(stringResource(R.string.char_noScore, profile.character.name), style = FsType.meta, color = FsColor.muted)
        }
        return
    }
    val label = copy.scoreLabel(profile).orEmpty()
    Column(Modifier.fillMaxWidth().primaryCard().padding(16.dp), verticalArrangement = Arrangement.spacedBy(10.dp)) {
        Row(Modifier.clearAndSetSemantics { contentDescription = label }, verticalAlignment = Alignment.Top) {
            Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(4.dp)) {
                MicroLabel(stringResource(R.string.char_score))
                Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(10.dp)) {
                    Text(copy.score(score), style = FsType.mono.copy(fontSize = CharacterLayout.scoreSize, fontWeight = FontWeight.Bold, letterSpacing = SCORE_TRACKING), color = FsColor.strong)
                    Tag(copy.tier(profile.tier), ItemQuality(CharacterCard.tierQuality(profile.tier)).text)
                }
                copy.asOf(profile.asOf)?.let { Text(it, style = FsType.meta, color = FsColor.muted) }
            }
            if (!isLargeText()) Sparkline(CharacterCard.sparkline(profile.scoreHistory), ItemQuality(CharacterCard.tierQuality(profile.tier)).text)
        }
        Column(Modifier.clearAndSetSemantics {}, verticalArrangement = Arrangement.spacedBy(8.dp)) {
            CharacterCard.bars(profile.components, profile.componentCaps).forEach { ComponentBar(it) }
        }
        Text(
            stringResource(R.string.char_how), style = FsType.bodyStrong, color = FsColor.gold,
            modifier = Modifier.heightIn(min = 48.dp).clickable(role = Role.Button, onClick = actions.openHow).padding(vertical = 12.dp),
        )
    }
}

@Composable
private fun ComponentBar(bar: ScoreBar) {
    val name = stringResource(bar.component.label)
    val track = @Composable { modifier: Modifier ->
        Box(modifier.height(CharacterLayout.bar).background(FsColor.lineSoft, CircleShape)) {
            Box(Modifier.fillMaxWidth(bar.fraction).height(CharacterLayout.bar).background(FsColor.goldHover, CircleShape))
        }
    }
    if (isLargeText()) {
        Column(verticalArrangement = Arrangement.spacedBy(4.dp)) {
            Row { Text(name, style = FsType.body, color = FsColor.muted, modifier = Modifier.weight(1f)); Text(bar.value.toString(), style = FsType.mono, color = FsColor.strong) }
            track(Modifier.fillMaxWidth())
        }
    } else {
        Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(12.dp)) {
            Text(name, style = FsType.body, color = FsColor.muted, modifier = Modifier.width(CharacterLayout.barLabel))
            track(Modifier.weight(1f))
            Text(bar.value.toString(), style = FsType.mono, color = FsColor.strong, modifier = Modifier.width(36.dp), textAlign = TextAlign.End)
        }
    }
}

/** `score_history[]`, oldest first, in the tier's colour. */
@Composable
private fun Sparkline(points: List<Float>, color: Color) {
    if (points.size < 2) return
    Canvas(Modifier.width(CharacterLayout.sparkWidth).height(CharacterLayout.sparkHeight)) {
        val step = size.width / (points.size - 1)
        val path = Path()
        points.forEachIndexed { i, p ->
            val at = Offset(i * step, size.height * (1f - p))
            if (i == 0) path.moveTo(at.x, at.y) else path.lineTo(at.x, at.y)
        }
        drawPath(path, color, style = Stroke(width = 2.dp.toPx(), cap = StrokeCap.Round))
    }
}

@Composable
private fun BestRow(profile: CharacterProfile, fight: CharacterFight, now: Instant, copy: CharacterCopy) {
    Row(Modifier.fillMaxWidth().heightIn(min = 64.dp).padding(horizontal = 16.dp, vertical = 10.dp), verticalAlignment = Alignment.CenterVertically) {
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
            Text(fight.encounter, style = FsType.bodyStrong, color = FsColor.strong)
            Text(copy.bestMeta(profile, fight, now), style = FsType.meta, color = FsColor.muted)
        }
        Column(horizontalAlignment = Alignment.End, verticalArrangement = Arrangement.spacedBy(2.dp)) {
            FigureWithUnit(copy.value(fight), copy.unit(fight.metric))
            fight.percentile?.let { p ->
                val n = p.roundToInt()
                Row(horizontalArrangement = Arrangement.spacedBy(4.dp), verticalAlignment = Alignment.CenterVertically) {
                    Text(n.toString(), style = FsType.mono.copy(fontWeight = FontWeight.Bold), color = parseColor(ParseTier.of(n)))
                    Text(stringResource(R.string.char_percentile), style = FsType.meta, color = FsColor.muted)
                }
            }
        }
    }
}

@Composable
private fun RecentRow(fight: CharacterFight, now: Instant, copy: CharacterCopy, reportsOn: Boolean, actions: CharacterActions) {
    val tap = if (reportsOn) Modifier.clickable(role = Role.Button) { actions.openFight(fight) } else Modifier
    Row(
        Modifier.fillMaxWidth().then(tap).heightIn(min = 64.dp).padding(horizontal = 16.dp, vertical = 10.dp),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        OutcomeDot(if (fight.kill == false) FsColor.wipe else FsColor.kill)
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
            Text(fight.encounter, style = FsType.bodyStrong, color = FsColor.strong)
            Text(copy.fightMeta(fight, now), style = FsType.meta, color = FsColor.muted)
        }
        FigureWithUnit(copy.value(fight), copy.unit(fight.metric))
        if (reportsOn) Chevron(20.dp)
    }
}

@Composable
private fun FigureWithUnit(value: String, unit: String) {
    Row(horizontalArrangement = Arrangement.spacedBy(6.dp), verticalAlignment = Alignment.CenterVertically) {
        Text(value, style = FsType.mono.copy(fontSize = 17.sp), color = FsColor.strong)
        Text(unit, style = FsType.monoMeta, color = FsColor.muted)
    }
}

/** F11.5: the Paperdoll with its caption and the gear gap, or the no-export sentence. */
@Composable
private fun GearSection(profile: CharacterProfile, state: CharacterScreenState, actions: CharacterActions) {
    Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
        MicroLabel(stringResource(R.string.char_gear))
        if (profile.gear.isEmpty()) {
            Text(
                stringResource(R.string.char_noGear, profile.character.name), style = FsType.body, color = FsColor.text,
                modifier = Modifier.fillMaxWidth().quietCard().padding(16.dp),
            )
            return@Column
        }
        val slots = Paperdoll.fromProfile(profile.gear, state.items, LocalResources.current)
        Column(Modifier.fillMaxWidth().quietCard().padding(12.dp), verticalArrangement = Arrangement.spacedBy(10.dp)) {
            PaperdollFrame(slots, state.build, actions.openSocket)
            if (slots.any { it.better }) Text(stringResource(R.string.char_goldDot), style = FsType.meta, color = FsColor.muted)
        }
        profile.gearGapTotal?.let { gap ->
            Row(Modifier.fillMaxWidth().quietCard().heightIn(min = 56.dp).padding(horizontal = 16.dp), verticalAlignment = Alignment.CenterVertically) {
                Text(stringResource(R.string.char_gap), style = FsType.bodyStrong, color = FsColor.strong, modifier = Modifier.weight(1f))
                Text(stringResource(R.string.char_gapPoints, ItemNumbers.tenths(gap)), style = FsType.mono, color = FsColor.strong)
            }
        }
    }
}

/** F11.7: the sentence, the help and `Search again`, which goes back. */
@Composable
private fun NotFound(name: String, actions: CharacterActions) {
    Column(Modifier.fillMaxWidth().quietCard().padding(16.dp), verticalArrangement = Arrangement.spacedBy(8.dp)) {
        Text(stringResource(R.string.char_notFound, name), style = FsType.lead, color = FsColor.strong)
        Text(stringResource(R.string.char_notFoundHelp), style = FsType.meta, color = FsColor.muted)
        BrandButton(stringResource(R.string.char_searchAgain), BrandButtonKind.Secondary, actions.searchAgain, height = 48.dp, fill = false)
    }
}

@Composable
private fun Section(title: String, rows: @Composable () -> Unit) {
    Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
        MicroLabel(title)
        Column(Modifier.fillMaxWidth().quietCard()) { rows() }
    }
}

@Composable
private fun Divided(index: Int, row: @Composable () -> Unit) {
    if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
    row()
}

@Composable
private fun Skeleton() {
    Column(Modifier.clearAndSetSemantics {}, verticalArrangement = Arrangement.spacedBy(22.dp)) {
        Row(horizontalArrangement = Arrangement.spacedBy(14.dp), verticalAlignment = Alignment.CenterVertically) {
            SkeletonDot(CharacterLayout.crest)
            Column(verticalArrangement = Arrangement.spacedBy(8.dp)) { SkeletonBar(160.dp, 24.dp); SkeletonBar(200.dp, 13.dp); SkeletonBar(150.dp, 13.dp) }
        }
        Column(Modifier.fillMaxWidth().primaryCard().padding(16.dp), verticalArrangement = Arrangement.spacedBy(10.dp)) {
            SkeletonBar(100.dp, 11.dp)
            SkeletonBar(140.dp, 36.dp)
            repeat(BAR_COUNT) { SkeletonBar(260.dp, 8.dp) }
        }
    }
}

private const val DISABLED_ALPHA = 0.5f

/** Mono at 36 sets the comma in a full cell; a little negative tracking keeps `1,830` one figure. */
private val SCORE_TRACKING = (-0.08).em

/** Material's `star_border` (the core icon set's outlined star is solid). */
private val StarOutline: ImageVector by lazy {
    materialIcon(name = "Outlined.StarBorder") { addPath(addPathNodes(STAR_BORDER_PATH), fill = SolidColor(Color.Black)) }
}

private const val STAR_BORDER_PATH =
    "M22,9.24l-7.19,-0.62L12,2L9.19,8.63L2,9.24l5.46,4.73L5.82,21L12,17.27L18.18,21l-1.63,-7.03L22,9.24z" +
        "M12,15.4l-3.76,2.27l1,-4.28l-3.32,-2.88l4.38,-0.38L12,6.1l1.71,4.04l4.38,0.38l-3.32,2.88l1,4.28L12,15.4z"
private const val BAR_COUNT = 5
