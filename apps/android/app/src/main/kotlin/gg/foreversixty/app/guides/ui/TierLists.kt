package gg.foreversixty.app.guides.ui

import android.content.res.Resources
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ColumnScope
import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.IntrinsicSize
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.layout.widthIn
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Share
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Size
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.pluralStringResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.common.FactionMark
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.flows.FlowChoice
import gg.foreversixty.app.flows.FlowChoiceGroup
import gg.foreversixty.app.flows.FlowControlBar
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowLabel
import gg.foreversixty.app.flows.FlowLink
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.FlowSkeleton
import gg.foreversixty.app.flows.FlowTag
import gg.foreversixty.app.guides.logic.GuideDates
import gg.foreversixty.app.guides.logic.GuidesCopy
import gg.foreversixty.app.guides.logic.Loadable
import gg.foreversixty.app.guides.logic.TierAnswer
import gg.foreversixty.app.guides.logic.TierChoice
import gg.foreversixty.app.guides.logic.TierGroup
import gg.foreversixty.app.guides.logic.TierRowModel
import gg.foreversixty.app.guides.logic.TierTone
import gg.foreversixty.app.guides.logic.Tiers
import gg.foreversixty.app.model.GuideIndex
import gg.foreversixty.app.model.TiersFile
import gg.foreversixty.app.reports.ui.ErrorRegion
import gg.foreversixty.app.reports.ui.isLargeText
import java.time.Instant

data class TierListState(
    val tiers: Loadable<TiersFile>,
    val choice: TierChoice,
    /** The current character's spec id; null with no character (no answer, no `Your spec`). */
    val yourSpec: String?,
    /** The guides index, for the rows' spec and class names. */
    val index: GuideIndex?,
    /** Signed in: no `Sign in` line under the caption. */
    val signedIn: Boolean,
)

/**
 * The page's links. A row's BiS, Guide and Planner open in the app; Sim opens the ranking sim on the
 * site. A null action draws no link.
 */
class TierListActions(
    val choose: (TierChoice) -> Unit = {},
    val open: (String) -> Unit = {},
    val openBis: ((classSlug: String, spec: String) -> Unit)? = null,
    val openPlanner: ((String) -> Unit)? = null,
    /** The ranking sim of a spec and faction (`/sim?code=…`); the tank list has none. */
    val openSim: ((spec: String, faction: String) -> Unit)? = null,
    /** `How we check the sim ›` (`/sim/specs`). */
    val openChecks: (() -> Unit)? = null,
    val retry: () -> Unit = {},
    /** Shares the role's tier list page; null draws no share button. */
    val share: ((TierChoice) -> Unit)? = null,
    val signIn: (() -> Unit)? = null,
)

/**
 * The tier list (mock `tiers`, `tiers-you`; the site's /tiers at phone width): the head with your spec's
 * place as its answer, the role and faction bar, the caption, the plate of letter groups (rank, crest,
 * spec, race, the bar of the top spec's share, ties, BiS · Guide · Planner · Sim) and How to read this.
 */
@Composable
fun TierListContent(state: TierListState, now: Instant, actions: TierListActions) {
    val res = LocalResources.current
    val copy = GuidesCopy(res)
    val file = (state.tiers as? Loadable.Loaded)?.value
    val rows = file?.let { Tiers.rows(Tiers.list(it, state.choice), state.index, state.yourSpec) }.orEmpty()
    FlowPage(top = 0.dp) {
        (state.tiers as? Loadable.Loaded)?.offlineSince?.let { OfflineStrip(stringResource(R.string.offline_strip, OfflineStamp.format(it, now))) }
        val facts = file?.let { listOf(stringResource(R.string.tiers_facts, it.presetLabel.orEmpty(), GuideDates.short(it.updated.orEmpty()))) }.orEmpty()
        val answer = Tiers.answer(rows)?.let { Answer(it, state.choice, actions) }
        FlowHead(stringResource(R.string.tiers_title), facts, answer = answer)
        Controls(state.choice, file, copy, actions)
        Caption(state.choice.role, if (state.signedIn) null else actions.signIn)
        when (val tiers = state.tiers) {
            Loadable.Loading -> TierSkeleton()
            Loadable.NotCached -> FlowNote(stringResource(R.string.guides_listNotCached), Modifier.padding(top = FsSpace.space5))
            Loadable.Failed -> ErrorRegion(stringResource(R.string.tiers_error), actions.retry, inset = false)
            is Loadable.Loaded -> {
                Plate(rows, state.choice, tiers.value.build, actions)
                HowToRead(state.choice, tiers.value.presetLabel.orEmpty(), copy, actions.openChecks)
            }
        }
    }
}

/** The answer: `Your spec · DPS`, the crest and name over its tier and gap, `2nd of 20`, and its three links. */
private fun Answer(answer: TierAnswer, choice: TierChoice, actions: TierListActions): @Composable ColumnScope.() -> Unit = {
    val res = LocalResources.current
    val model = answer.row
    val row = model.row
    FlowLabel(stringResource(R.string.tiers_yourSpec, GuidesCopy(res).role(choice.role)))
    Row(Modifier.padding(top = ANSWER_GAP), verticalAlignment = Alignment.CenterVertically) {
        ClassCrest(row.classSlug, ANSWER_CREST, Modifier.parityItem(ParityItemKind.Image))
        Spacer(Modifier.width(ANSWER_ROW_GAP))
        Column(Modifier.weight(1f)) {
            Text(row.name, style = FsType.itemLg, color = classColor(row.classSlug))
            Text(answerLine(res, model, answer), style = FsType.small, color = FsColor.muted)
        }
        Spacer(Modifier.width(FsSpace.space2))
        Text(
            buildAnnotatedString {
                append(ordinal(res, row.rank))
                withStyle(SpanStyle(fontSize = 13.sp, fontWeight = FontWeight.SemiBold, color = FsColor.muted)) {
                    append(" " + res.getString(R.string.tiers_answer_of, answer.count))
                }
            },
            style = FsType.figureBig,
            color = FsColor.gold,
        )
    }
    FlowRow(Modifier.padding(top = ANSWER_GAP), horizontalArrangement = Arrangement.spacedBy(FsSpace.space4)) {
        actions.openBis?.let { bis -> FlowLink(stringResource(R.string.tiers_link_yourBis)) { bis(row.classSlug, row.spec) } }
        FlowLink(stringResource(R.string.tiers_link_yourGuide)) { actions.open(row.spec) }
        actions.openSim?.takeIf { choice.role != TierChoice.TANK }?.let { sim ->
            FlowLink(stringResource(R.string.tiers_link_yourSim)) { sim(row.spec, choice.faction) }
        }
    }
}

/** `Tier A · 7.3% behind Fury Warrior`, or the top spec's line. */
private fun answerLine(res: Resources, model: TierRowModel, answer: TierAnswer): String {
    val row = model.row
    if (row.rank == 1) return res.getString(R.string.tiers_answer_top)
    val gap = row.gapText.trimStart(MINUS, '-', '+')
    return row.tier?.let { res.getString(R.string.tiers_answer_tier, it, gap, answer.top.name) }
        ?: res.getString(R.string.tiers_answer_behind, gap, answer.top.name)
}

/** `1st`, `2nd`, `3rd`, `4th`, `11th`, `22nd`. */
private fun ordinal(res: Resources, n: Int): String {
    val teens = n % HUNDRED in TEENS
    val key = when {
        teens -> R.string.tiers_ordinal_th
        n % TEN == 1 -> R.string.tiers_ordinal_st
        n % TEN == 2 -> R.string.tiers_ordinal_nd
        n % TEN == 3 -> R.string.tiers_ordinal_rd
        else -> R.string.tiers_ordinal_th
    }
    return res.getString(key, n)
}

/** The bar: `DPS · Alliance`, opening Role (with each list's count) and Faction (with its mark). */
@Composable
private fun Controls(choice: TierChoice, file: TiersFile?, copy: GuidesCopy, actions: TierListActions) {
    val counts = file?.let { Tiers.counts(it, choice.faction) }.orEmpty()
    val roles = FlowChoiceGroup(
        stringResource(R.string.tiers_role),
        TierChoice.roles.map { FlowChoice(copy.role(it), counts[it]?.toString()) },
        TierChoice.roles.indexOf(choice.role),
    ) { actions.choose(choice.copy(role = TierChoice.roles[it])) }
    val factions = FlowChoiceGroup(
        stringResource(R.string.tiers_faction),
        TierChoice.factions.map { faction -> FlowChoice(copy.faction(faction), mark = { FactionMark(faction, FACTION_MARK) }) },
        TierChoice.factions.indexOf(choice.faction),
    ) { actions.choose(choice.copy(faction = TierChoice.factions[it])) }
    val share = actions.share?.let { share ->
        @Composable {
            IconButton(onClick = { share(choice) }) { Icon(Icons.Filled.Share, contentDescription = stringResource(R.string.share_action), tint = FsColor.strong) }
        }
    }
    FlowControlBar(
        stringResource(R.string.tiers_summary, copy.role(choice.role), copy.faction(choice.faction)),
        listOf(roles, factions),
        inset = BAR_INSET,
        end = share,
    )
}

/** The caption (`.notice`): what the number is, and signed out, `Sign in to see your spec marked.` on its own line. */
@OptIn(ExperimentalLayoutApi::class)
@Composable
private fun Caption(role: String, signIn: (() -> Unit)?) {
    val text = when (role) {
        TierChoice.HEALER -> stringResource(R.string.tiers_caption_healer)
        TierChoice.TANK -> stringResource(R.string.tiers_caption_tank)
        else -> stringResource(R.string.tiers_caption_dps, percent(Tiers.TIE_PERCENT))
    }
    FlowRow(
        Modifier.fillMaxWidth().padding(top = CAPTION_TOP, bottom = FsSpace.space1),
        horizontalArrangement = Arrangement.spacedBy(CAPTION_GAP),
        verticalArrangement = Arrangement.spacedBy(6.dp),
    ) {
        Text(text, style = CAPTION, color = FsColor.text)
        if (signIn != null) {
            Row {
                Text(stringResource(R.string.tiers_signIn), style = CAPTION.copy(fontWeight = FontWeight.SemiBold), color = FsColor.gold, modifier = Modifier.clickable(role = Role.Button, onClick = signIn))
                Text(" " + stringResource(R.string.tiers_signInLine), style = CAPTION, color = FsColor.text)
            }
        }
    }
}

/** The plate (`.plate` in `.tiers-plate`): letter groups, or the plain list for tanks and healers. */
@Composable
private fun Plate(rows: List<TierRowModel>, choice: TierChoice, build: String, actions: TierListActions) {
    val groups = Tiers.groups(rows)
    Column(
        Modifier
            .padding(top = PLATE_TOP)
            .fillMaxWidth()
            .parityRegion(ParityRegion.PLATE)
            .background(FsColor.raised, PLATE_SHAPE)
            .border(1.dp, FsColor.line, PLATE_SHAPE)
            .padding(1.dp),
    ) {
        if (groups.isEmpty()) {
            rows.forEachIndexed { i, row ->
                if (i > 0) HorizontalDivider(color = FsColor.lineSoft)
                TierRowView(row, rows.size, choice, build, actions)
            }
        } else {
            groups.forEachIndexed { i, group ->
                if (i > 0) HorizontalDivider(color = FsColor.line)
                Group(group, rows.size, choice, build, actions)
            }
        }
    }
}

/** A letter's group (`.rg`): the letter, its band and count on a 42 line with the tier's left edge, then its rows. */
@Composable
private fun Group(group: TierGroup, count: Int, choice: TierChoice, build: String, actions: TierListActions) {
    val tone = toneColor(group.tone)
    Row(
        Modifier
            .fillMaxWidth()
            .height(IntrinsicSize.Min)
            .heightIn(min = GROUP_LINE)
            .semantics(mergeDescendants = true) {},
        verticalAlignment = Alignment.CenterVertically,
    ) {
        Box(Modifier.width(GROUP_EDGE).fillMaxHeight().background(tone))
        Text(group.letter, style = LETTER, color = tone, modifier = Modifier.padding(start = GROUP_LETTER_INSET))
        Text(band(group), style = BAND, color = FsColor.text, modifier = Modifier.padding(start = FsSpace.space3).weight(1f))
        Text(pluralStringResource(R.plurals.tiers_groupCount, group.rows.size, group.rows.size), style = BAND, color = FsColor.muted, modifier = Modifier.padding(end = ROW_INSET))
    }
    group.rows.forEach { row ->
        HorizontalDivider(color = FsColor.lineSoft)
        TierRowView(row, count, choice, build, actions)
    }
}

@Composable
private fun band(group: TierGroup): String {
    val start = percent(group.start)
    val end = group.end?.let(::percent)
    return when {
        end == null -> stringResource(R.string.tiers_band_beyond, start)
        group.start == 0 -> stringResource(R.string.tiers_band_top, end)
        else -> stringResource(R.string.tiers_band_between, group.start.toString(), end)
    }
}

/**
 * One row (`.rr`): rank, the 36 crest, the spec in its class colour over `Class · Race`, the figure and
 * gap at the right; under them the bar, the tie line and the links. Your row has the gold edge and wash.
 */
@Composable
private fun TierRowView(model: TierRowModel, count: Int, choice: TierChoice, build: String, actions: TierListActions) {
    val res = LocalResources.current
    val row = model.row
    val label = GuidesCopy(res).tierLabel(model, count)
    // The largest text: the figures move under the name, so neither is cut.
    val stacked = isLargeText()
    Column(
        Modifier
            .fillMaxWidth()
            .parityItem(ParityItemKind.Row)
            .then(if (model.isYours) Modifier.yoursGround() else Modifier)
            .padding(start = ROW_INSET, end = ROW_INSET, top = ROW_TOP, bottom = ROW_BOTTOM),
    ) {
        Row(Modifier.semantics(mergeDescendants = true) { contentDescription = label }, verticalAlignment = Alignment.CenterVertically) {
            Text(row.rank.toString(), style = RANK, color = FsColor.muted, textAlign = TextAlign.End, modifier = Modifier.width(RANK_WIDTH))
            Spacer(Modifier.width(FsSpace.space3))
            ClassCrest(row.classSlug, ROW_CREST, Modifier.parityItem(ParityItemKind.Image))
            Spacer(Modifier.width(FsSpace.space2))
            Column(Modifier.weight(1f)) {
                Row(horizontalArrangement = Arrangement.spacedBy(FsSpace.space2), verticalAlignment = Alignment.CenterVertically) {
                    Text(model.specName, style = FsType.figureSm, color = classColor(row.classSlug))
                    if (model.isYours) FlowTag(stringResource(R.string.tiers_yours), FsColor.gold)
                }
                Text("${model.className} · ${row.race}", style = SUB, color = FsColor.muted)
            }
            if (!stacked) Figures(model, choice, Alignment.End)
        }
        Column(Modifier.padding(start = ROW_BODY_INDENT)) {
            if (stacked) Figures(model, choice, Alignment.Start)
            Meter(row.fraction.toFloat(), classColor(row.classSlug))
            if (model.ties.isNotEmpty()) {
                Text(stringResource(R.string.tiers_tieWith, names(res, model.ties)), style = TIE, color = FsColor.strong, modifier = Modifier.padding(top = FsSpace.space1))
            }
            Links(model, choice, actions)
        }
    }
}

@Composable
private fun Figures(model: TierRowModel, choice: TierChoice, align: Alignment.Horizontal) {
    val row = model.row
    val top = row.rank == 1
    Column(Modifier.widthIn(min = FIGURE_WIDTH), horizontalAlignment = align) {
        Text(
            buildAnnotatedString {
                append(row.metricText)
                withStyle(SpanStyle(fontSize = 12.sp, fontWeight = FontWeight.SemiBold, color = FsColor.muted)) { append(" ${row.unit}") }
            },
            style = FsType.figureSm,
            color = FsColor.strong,
        )
        val gap = if (top && choice.role != TierChoice.TANK) stringResource(R.string.tiers_top) else row.gapText
        Text(gap, style = if (top) GAP.copy(fontWeight = FontWeight.Bold) else GAP, color = if (top) FsColor.strong else FsColor.text)
    }
}

/** `BiS · Guide · Planner · Sim` (44 tall, BiS gold); the tank list has no Sim. */
@Composable
@OptIn(ExperimentalLayoutApi::class)
private fun Links(model: TierRowModel, choice: TierChoice, actions: TierListActions) {
    val row = model.row
    FlowRow(horizontalArrangement = Arrangement.spacedBy(FsSpace.space5)) {
        actions.openBis?.let { bis -> RowLink(stringResource(R.string.tiers_link_bis), FsColor.gold) { bis(row.classSlug, row.spec) } }
        RowLink(stringResource(R.string.tiers_link_guide), FsColor.nav) { actions.open(row.spec) }
        actions.openPlanner?.let { planner -> RowLink(stringResource(R.string.tiers_link_planner), FsColor.nav) { planner(row.spec) } }
        actions.openSim?.takeIf { choice.role != TierChoice.TANK }?.let { sim -> RowLink(stringResource(R.string.tiers_link_sim), FsColor.nav) { sim(row.spec, choice.faction) } }
    }
}

@Composable
private fun RowLink(text: String, color: Color, onClick: () -> Unit) {
    Box(Modifier.heightIn(min = LINK_HEIGHT).clickable(role = Role.Button, onClick = onClick), contentAlignment = Alignment.CenterStart) {
        Text(text, style = LINK, color = color)
    }
}

/** The 8 bar of the top spec's share (`.meter.cls`), in the class colour. */
@Composable
private fun Meter(fraction: Float, color: Color) {
    val shape = RoundedCornerShape(FsRadius.field)
    Box(Modifier.padding(top = FsSpace.space2).fillMaxWidth().height(METER).background(FsColor.lineSoft, shape)) {
        Box(Modifier.fillMaxHeight().fillMaxWidth(fraction.coerceIn(0f, 1f)).background(color, shape))
    }
}

/** `5%`: copy never carries the sign, so a percent goes in as an argument. */
private fun percent(n: Int): String = "$n$PERCENT"

/** `Survival`, `Survival and Arms`, `Arms, Fury and Survival`. */
private fun names(res: Resources, names: List<String>): String = when (names.size) {
    1 -> names.single()
    else -> res.getString(R.string.tiers_and, names.dropLast(1).joinToString(", "), names.last())
}

/** `S is within 5% of the top, A within 10%, …`: the letters' bands and the tie margin. */
private fun letters(res: Resources): String {
    val (s, a, b, c) = Tiers.bandEnds.map(::percent)
    return res.getString(R.string.tiers_notes_dps_letters, s, a, b, c, percent(Tiers.TIE_PERCENT))
}

/** How to read this: three labelled notes for the role, and `How we check the sim ›`. */
@Composable
private fun HowToRead(choice: TierChoice, preset: String, copy: GuidesCopy, openChecks: (() -> Unit)?) {
    val res = LocalResources.current
    val gear = res.getString(R.string.tiers_notes_gear, preset, copy.faction(choice.faction))
    val fresh = res.getString(R.string.tiers_notes_fresh)
    val notes = when (choice.role) {
        TierChoice.HEALER -> listOf(
            R.string.tiers_notes_number to res.getString(R.string.tiers_notes_healer_number, gear),
            R.string.tiers_notes_leavesOut to res.getString(R.string.tiers_notes_healer_leavesOut, fresh),
            R.string.tiers_notes_ties to res.getString(R.string.tiers_notes_healer_ties),
        )
        TierChoice.TANK -> listOf(
            R.string.tiers_notes_number to res.getString(R.string.tiers_notes_tank_number),
            R.string.tiers_notes_boss to res.getString(R.string.tiers_notes_tank_boss),
            R.string.tiers_notes_gearTitle to res.getString(R.string.tiers_notes_tank_gear, gear, fresh),
        )
        else -> listOf(
            R.string.tiers_notes_number to res.getString(R.string.tiers_notes_dps_number, gear),
            R.string.tiers_notes_leavesOut to res.getString(R.string.tiers_notes_dps_leavesOut, fresh),
            R.string.tiers_notes_letters to letters(res),
        )
    }
    FlowSection(stringResource(R.string.tiers_howToRead), action = openChecks?.let { stringResource(R.string.tiers_howWeCheck) }, onAction = openChecks ?: {}, top = SECTION_TOP) {
        Column(Modifier.padding(top = FsSpace.space3), verticalArrangement = Arrangement.spacedBy(FsSpace.space3)) {
            notes.forEach { (title, text) ->
                Column {
                    FlowLabel(stringResource(title))
                    Text(text, style = NOTE, color = FsColor.text, modifier = Modifier.padding(top = 6.dp))
                }
            }
        }
    }
}

@Composable
private fun TierSkeleton() {
    Column(Modifier.padding(top = PLATE_TOP), verticalArrangement = Arrangement.spacedBy(FsSpace.space3)) {
        repeat(SKELETON_ROWS) { FlowSkeleton(SKELETON_WIDTH, SKELETON_HEIGHT) }
        FlowNote(stringResource(R.string.tiers_loading))
    }
}

/** Your row: the selected wash and a 3 gold left edge. */
private fun Modifier.yoursGround(): Modifier = background(FsColor.selected)
    .drawBehind { drawRect(FsColor.gold, size = Size(GROUP_EDGE.toPx(), size.height)) }

/** The web's tier colours. */
private fun toneColor(tone: TierTone?): Color = when (tone) {
    TierTone.Legendary -> FsColor.rarityLegendary
    TierTone.Epic -> FsColor.rarityEpicText
    TierTone.Rare -> FsColor.rarityRareText
    TierTone.Uncommon -> FsColor.rarityUncommon
    TierTone.Muted, null -> FsColor.rarityPoor
}

private const val MINUS = '−'
private const val PERCENT = "%"
private const val TEN = 10
private const val HUNDRED = 100
private val TEENS = 11..13
private val PLATE_SHAPE = RoundedCornerShape(FsRadius.plate)
private val ANSWER_CREST = 40.dp
private val ANSWER_GAP = 6.dp
private val ANSWER_ROW_GAP = 14.dp
private val FACTION_MARK = 20.dp
private val BAR_INSET = 6.dp
private val CAPTION_TOP = 14.dp
private val CAPTION_GAP = 14.dp
private val PLATE_TOP = 32.dp
private val SECTION_TOP = 48.dp
private val GROUP_LINE = 42.dp
private val GROUP_EDGE = 3.dp
private val GROUP_LETTER_INSET = 12.dp
private val ROW_INSET = 14.dp
private val ROW_TOP = 10.dp
private val ROW_BOTTOM = 2.dp
private val RANK_WIDTH = 22.dp
private val ROW_CREST = 36.dp
private val ROW_BODY_INDENT = 78.dp
private val FIGURE_WIDTH = 78.dp
private val METER = 8.dp
private val LINK_HEIGHT = 44.dp
private val SKELETON_WIDTH = 380.dp
private val SKELETON_HEIGHT = 113.dp
private const val SKELETON_ROWS = 4
private val CAPTION = FsType.body.copy(lineHeight = 24.sp)
private val NOTE = FsType.body.copy(lineHeight = 24.sp)
private val SUB = FsType.small.copy(lineHeight = 18.sp)
private val RANK = FsType.item
private val GAP = FsType.bodyStrong.copy(lineHeight = 20.sp)
private val TIE = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.Bold)
private val LINK = FsType.bodyStrong.copy(fontSize = 14.sp, lineHeight = 20.sp)
private val LETTER = FsType.itemLg.copy(fontSize = 22.sp, lineHeight = 24.sp)
private val BAND = FsType.small.copy(lineHeight = 18.sp)
