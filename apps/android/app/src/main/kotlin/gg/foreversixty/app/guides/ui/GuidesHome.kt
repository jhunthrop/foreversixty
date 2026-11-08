package gg.foreversixty.app.guides.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ColumnScope
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.flows.FlowGap
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowLabel
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.flows.FlowLink
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSkeleton
import gg.foreversixty.app.flows.FlowTag
import gg.foreversixty.app.guides.logic.GuideDates
import gg.foreversixty.app.guides.logic.Guides
import gg.foreversixty.app.guides.logic.GuidesCopy
import gg.foreversixty.app.guides.logic.Loadable
import gg.foreversixty.app.guides.logic.TierChoice
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.GuideIndex
import gg.foreversixty.app.model.GuideIndexClass
import gg.foreversixty.app.model.GuideIndexSpec
import gg.foreversixty.app.reports.ui.ErrorRegion

/** What the Guides page draws (the site's /guides at phone width). */
data class GuidesHomeState(
    val index: Loadable<GuideIndex>,
    /** Saved guides by spec id, with their bytes. */
    val saved: Map<String, Long>,
    /** The current character; null signed out (no answer). */
    val character: AccountCharacter?,
    /** The current character's guide; null when its spec has none (no answer). */
    val yours: GuideIndexSpec?,
)

class GuidesHomeActions(
    val openGuide: (String) -> Unit = {},
    val openClass: (String) -> Unit = {},
    val retry: () -> Unit = {},
)

/**
 * The Guides page (mock `guides`, `guides-you`): the head with your character's guide as its answer,
 * then one plate of every class, its crest row opening the class page, its specs (28 icons, the role
 * tag for tanks and healers, `Saved` when on the phone) opening the guide; what is saved under it.
 */
@Composable
fun GuidesHome(state: GuidesHomeState, actions: GuidesHomeActions) {
    val index = (state.index as? Loadable.Loaded)?.value
    FlowPage(top = 0.dp) {
        val facts = index?.let { Guides.updated(it) }?.let { listOf(stringResource(R.string.guides_updated, GuideDates.short(it))) }.orEmpty()
        val answer = state.character?.let { who -> state.yours?.let { spec -> index?.let { YourCharacter(who, spec, it, actions.openGuide) } } }
        FlowHead(stringResource(R.string.guides_title), facts, summary = stringResource(R.string.guides_lead), answer = answer)
        when (val phase = state.index) {
            Loadable.Loading -> GuidesSkeleton()
            Loadable.NotCached -> FlowNote(stringResource(R.string.guides_listNotCached), Modifier.padding(top = FsSpace.space5))
            Loadable.Failed -> ErrorRegion(stringResource(R.string.guides_error), actions.retry, inset = false)
            is Loadable.Loaded -> {
                ClassPlate(phase.value, state.saved, actions)
                val copy = GuidesCopy(LocalResources.current)
                FlowNote(copy.footer(Guides.footer(phase.value, state.saved)), Modifier.padding(top = FsSpace.space4))
            }
        }
    }
}

/** `Your character`: the crest, name and line, and the link to its guide. */
private fun YourCharacter(
    character: AccountCharacter,
    spec: GuideIndexSpec,
    index: GuideIndex,
    open: (String) -> Unit,
): @Composable ColumnScope.() -> Unit = {
    val className = Guides.classOf(index, spec.spec)?.className.orEmpty()
    val who = listOfNotNull(character.race, spec.specName, className).filter { it.isNotBlank() }.joinToString(" ")
    FlowLabel(stringResource(R.string.guides_yourCharacter))
    Row(Modifier.padding(top = ANSWER_GAP), verticalAlignment = Alignment.CenterVertically) {
        ClassCrest(character.className, ANSWER_CREST, Modifier.parityItem(ParityItemKind.Image))
        Spacer(Modifier.width(FsSpace.space3))
        Column {
            Text(character.name, style = FsType.itemLg, color = classColor(character.className))
            character.level?.let { Text(stringResource(R.string.guides_characterLine, it, who), style = FsType.small, color = FsColor.muted) }
        }
    }
    FlowLink(stringResource(R.string.guides_yourGuideLink, specTitle(spec.specName, className)), Modifier.padding(top = ANSWER_GAP)) { open(spec.spec) }
}

/** One plate (`.plate`) of every class group, 20 under the head. */
@Composable
private fun ClassPlate(index: GuideIndex, saved: Map<String, Long>, actions: GuidesHomeActions) {
    Column(
        Modifier
            .padding(top = FsSpace.space5)
            .fillMaxWidth()
            .parityRegion(ParityRegion.PLATE)
            .background(FsColor.raised, PLATE_SHAPE)
            .border(1.dp, FsColor.line, PLATE_SHAPE)
            .padding(1.dp),
    ) {
        index.classes.forEach { cls -> ClassGroup(cls, index.build, saved, actions) }
    }
}

/** A class (`.cgrp`): its crest row, then its specs as slim rows; 20 below the group. */
@Composable
private fun ClassGroup(cls: GuideIndexClass, build: String, saved: Map<String, Long>, actions: GuidesHomeActions) {
    val copy = GuidesCopy(LocalResources.current)
    Column(Modifier.fillMaxWidth().padding(start = FsSpace.space4, end = FsSpace.space4, bottom = FsSpace.space5)) {
        FlowRow(cls.className, lead = FlowLead.Crest(cls.classSlug), titleColor = classColor(cls.classSlug)) { actions.openClass(cls.classSlug) }
        cls.specs.forEachIndexed { i, spec ->
            val tags = listOfNotNull(
                copy.role(spec.role).takeIf { spec.role != TierChoice.DPS }?.let { it to FsColor.muted },
                stringResource(R.string.guides_savedShort).takeIf { spec.spec in saved }?.let { it to FsColor.kill },
            )
            FlowRow(
                spec.specName,
                lead = FlowLead.Icon(spec.icon, build, SPEC_ICON),
                slim = true,
                last = i == cls.specs.lastIndex,
                trail = if (tags.isEmpty()) null else ({ Tags(tags) }),
            ) { actions.openGuide(spec.spec) }
        }
    }
}

@Composable
private fun Tags(tags: List<Pair<String, androidx.compose.ui.graphics.Color>>) {
    Row(horizontalArrangement = Arrangement.spacedBy(FsSpace.space2)) { tags.forEach { (text, color) -> FlowTag(text, color) } }
}

/** Loading: the plate's first rows at their real heights. */
@Composable
private fun GuidesSkeleton() {
    Column(Modifier.padding(top = FsSpace.space5), verticalArrangement = Arrangement.spacedBy(FsSpace.space3)) {
        repeat(SKELETON_ROWS) { FlowSkeleton(SKELETON_WIDTH, SKELETON_HEIGHT) }
        FlowGap(FsSpace.space2)
        FlowNote(stringResource(R.string.guides_loading))
    }
}

/** `Elemental Shaman`. */
internal fun specTitle(spec: String, className: String) = "$spec $className"

private val PLATE_SHAPE = RoundedCornerShape(FsRadius.plate)
private val SPEC_ICON = 28.dp
private val ANSWER_CREST = 40.dp
private val ANSWER_GAP = 6.dp
private val SKELETON_WIDTH = 380.dp
private val SKELETON_HEIGHT = 44.dp
private const val SKELETON_ROWS = 6
