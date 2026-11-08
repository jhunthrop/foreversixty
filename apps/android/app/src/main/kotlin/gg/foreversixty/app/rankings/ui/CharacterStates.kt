package gg.foreversixty.app.rankings.ui

import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.text.BasicTextField
import androidx.compose.foundation.text.KeyboardActions
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.SolidColor
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.input.ImeAction
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
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
import gg.foreversixty.app.flows.FlowButton
import gg.foreversixty.app.flows.FlowGap
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowNotice
import gg.foreversixty.app.flows.FlowPlate
import gg.foreversixty.app.flows.FlowPlateRetry
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowRule
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.model.FeedRow
import gg.foreversixty.app.rankings.logic.CharacterRoute
import gg.foreversixty.app.rankings.logic.RankingsCopy
import java.time.Instant

/** C18: the character's own feed rows, its crest, the sentence, and when and which group on the right. */
@Composable
fun CharacterActivity(rows: List<FeedRow>, now: Instant, actions: CharacterActions) {
    if (rows.isEmpty()) return
    val copy = RankingsCopy(LocalResources.current)
    FlowSection(stringResource(R.string.shared_character_activity), top = SECTION_GAP) {
        FlowGap(FsSpace.space3)
        rows.forEachIndexed { i, row ->
            ActivityRow(row, copy.age(row.at, now)) { actions.openActivity(row) }
            if (i < rows.lastIndex) FlowRule()
        }
    }
}

@Composable
private fun ActivityRow(row: FeedRow, age: String, open: () -> Unit) {
    Row(
        Modifier.fillMaxWidth().parityItem(ParityItemKind.Row).clickable(role = Role.Button, onClick = open).padding(vertical = FsSpace.space3),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        ClassCrest(row.crest, CREST, Modifier.parityItem(ParityItemKind.Image))
        Spacer(Modifier.width(ACTIVITY_GAP))
        Text(
            buildAnnotatedString {
                row.segments.forEach { segment ->
                    when (segment.style) {
                        NAME -> withStyle(SpanStyle(fontWeight = FontWeight.Bold, color = classColor(segment.className))) { append(segment.text) }
                        ITEM -> withStyle(SpanStyle(fontWeight = FontWeight.Bold, color = ItemQuality(segment.quality).text)) { append(segment.text) }
                        null -> append(segment.text)
                        else -> withStyle(SpanStyle(fontWeight = FontWeight.Bold, color = FsColor.strong)) { append(segment.text) }
                    }
                }
            },
            style = ACTIVITY,
            color = FsColor.text,
            modifier = Modifier.weight(1f),
        )
        Spacer(Modifier.width(ACTIVITY_GAP))
        Column(horizontalAlignment = Alignment.End) {
            Text(age, style = FsType.small, color = FsColor.muted)
            tagLabel(row.tag)?.let { Text(stringResource(it), style = TAG, color = FsColor.muted) }
        }
    }
}

private fun tagLabel(tag: String): Int? = when (tag) {
    "your_guild" -> R.string.shared_feed_tag_yourGuild
    "your_spec" -> R.string.shared_feed_tag_yourSpec
    "logs" -> R.string.shared_feed_tag_logs
    "rankings" -> R.string.shared_feed_tag_rankings
    "best_in_slot" -> R.string.shared_feed_tag_bestInSlot
    "guides" -> R.string.shared_feed_tag_guides
    "patch" -> R.string.shared_feed_tag_patch
    else -> null
}

/** F11.7 (C10): the name, the sentence, a look-up field, and how to add the character. */
@Composable
fun CharacterNotFound(route: CharacterRoute, signedIn: Boolean, actions: CharacterActions) {
    FlowHead(route.name, emptyList())
    Box(Modifier.fillMaxWidth().parityRegion(ParityRegion.SUMMARY)) {
        val where = RankingsCopy(LocalResources.current)
        FlowNotice(stringResource(R.string.char_page_notFound, route.name, route.region.uppercase(), where.rulesetLabel(route.ruleset)), null, {})
    }
    FlowSection(stringResource(R.string.char_page_lookup), top = FsSpace.space5) {
        FlowGap(FsSpace.space3)
        LookUp(route.name, actions.lookUp)
    }
    FlowSection(stringResource(R.string.char_page_add), top = SECTION_GAP) {
        FlowGap(FsSpace.space3)
        if (!signedIn) FlowRow(stringResource(R.string.char_page_signIn), stringResource(R.string.char_page_signInSub), onClick = actions.signIn)
        FlowRow(stringResource(R.string.shared_cta_getAddonAndCompanion), stringResource(R.string.char_page_addonSub), last = true, onClick = actions.getAddon)
    }
}

@Composable
private fun LookUp(initial: String, lookUp: (String) -> Unit) {
    var name by rememberSaveable { mutableStateOf(initial) }
    val label = stringResource(R.string.char_page_lookupName)
    Row(verticalAlignment = Alignment.Bottom) {
        Column(Modifier.weight(1f)) {
            Text(label.uppercase(), style = LABEL, color = FsColor.muted)
            Spacer(Modifier.height(FIELD_GAP))
            val shape = RoundedCornerShape(FsRadius.field)
            BasicTextField(
                name,
                { name = it },
                Modifier.fillMaxWidth().height(FIELD).parityItem(ParityItemKind.Field).border(1.dp, FsColor.line, shape).semantics { contentDescription = label },
                textStyle = FsType.body.copy(color = FsColor.strong),
                singleLine = true,
                cursorBrush = SolidColor(FsColor.gold),
                keyboardOptions = KeyboardOptions(imeAction = ImeAction.Search),
                keyboardActions = KeyboardActions(onSearch = { lookUp(name.trim()) }),
                decorationBox = { field -> Box(Modifier.padding(horizontal = FsSpace.space3), contentAlignment = Alignment.CenterStart) { field() } },
            )
        }
        Spacer(Modifier.width(FsSpace.space3))
        FlowButton(stringResource(R.string.char_page_lookupSubmit), { lookUp(name.trim()) }, primary = true, small = false)
    }
}

/** The load failed (C9): the name, the sentence and Try again in a plate. */
@Composable
fun CharacterError(name: String, retry: () -> Unit) {
    FlowHead(name, emptyList())
    FlowGap(FsSpace.space5)
    FlowPlate { FlowPlateRetry(stringResource(R.string.char_page_failed), stringResource(R.string.common_tryAgain), retry) }
}

private const val NAME = "name"
private const val ITEM = "item"
private val CREST = 40.dp
private val ACTIVITY_GAP = 14.dp
private val FIELD = 40.dp
private val FIELD_GAP = 6.dp
private val ACTIVITY = FsType.body
private val TAG = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp, fontWeight = FontWeight.Bold)
