package gg.foreversixty.app.search.ui

import androidx.activity.compose.BackHandler
import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.LazyListScope
import androidx.compose.foundation.selection.selectable
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.text.BasicTextField
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Close
import androidx.compose.material.icons.filled.Search
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.remember
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.focus.FocusRequester
import androidx.compose.ui.focus.focusRequester
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.SolidColor
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.input.ImeAction
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ScaledType
import gg.foreversixty.app.designsystem.parityContent
import gg.foreversixty.app.flows.FlowButton
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.FlowTag
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.model.CharacterSearchRow
import gg.foreversixty.app.search.logic.CharacterPhase
import gg.foreversixty.app.search.logic.IndexPhase
import gg.foreversixty.app.search.logic.SearchCopy
import gg.foreversixty.app.search.logic.SearchScope
import gg.foreversixty.app.search.logic.SearchSection
import gg.foreversixty.app.search.logic.SearchTarget
import gg.foreversixty.app.search.logic.WebHandoff

/**
 * The For group's three links (spec F12.1): the character's list, their spec's guide (null when
 * the spec is unknown) and Rankings.
 */
data class ForCharacter(val name: String, val guideLabel: String?)

/** Everything Search draws, pure so the tests set it directly. */
data class SearchScreenState(
    val query: String,
    val scope: SearchScope,
    val indexPhase: IndexPhase,
    val sections: List<SearchSection>,
    val characterPhase: CharacterPhase,
    val recents: List<SearchTarget>,
    val forCharacter: ForCharacter?,
    val build: String,
    /** The signed-in account's character keys: their rows carry `You`. */
    val ownKeys: Set<String> = emptySet(),
    /** Lane F15: the account's watched characters (`GET /v1/me/watching`), the Watching group. */
    val watching: List<CharacterSearchRow> = emptyList(),
    /** Lane F15: F11 is on, so characters and Rankings open in the app. */
    val rankingsInApp: Boolean = false,
    /** Lane F15: F9 is on, so guides open in the app (no external-link mark). */
    val guidesInApp: Boolean = false,
) {
    fun opensWebsite(target: SearchTarget): Boolean = WebHandoff.opensWebsite(target, rankingsInApp, guidesInApp)
}

/** What the screen's controls do; the route wires them. */
data class SearchActions(
    val onQuery: (String) -> Unit = {},
    val onScope: (SearchScope) -> Unit = {},
    val onBack: () -> Unit = {},
    val onOpen: (SearchTarget) -> Unit = {},
    val onClearRecent: () -> Unit = {},
    val onRetryIndex: () -> Unit = {},
    val onOpenBis: () -> Unit = {},
    val onOpenGuide: () -> Unit = {},
    val onOpenRankings: () -> Unit = {},
)

/**
 * Search (mocks `search`, `search-results`): the field with the keyboard up, then the hint and the
 * For, Watching and Recent groups, or the results by group. A group's Show all narrows to its scope;
 * Back from a narrowed scope returns to All.
 */
@Composable
fun SearchScreen(state: SearchScreenState, actions: SearchActions, focusField: Boolean = true) {
    val copy = SearchCopy(LocalResources.current)
    BackHandler(enabled = state.scope != SearchScope.All) { actions.onScope(SearchScope.All) }
    Column(Modifier.fillMaxSize().background(FsColor.bg).parityContent()) {
        SearchField(state.query, actions, focusField)
        LazyColumn(
            Modifier.fillMaxSize().weight(1f),
            contentPadding = PaddingValues(start = FsSpace.space4, end = FsSpace.space4, bottom = FsSpace.space6),
        ) {
            if (state.query.isBlank()) emptyState(state, copy, actions) else results(state, copy, actions)
        }
    }
}

/** The field (`.searchf`): 44 tall in the line, the glyph, the query or the placeholder, Clear once typed. */
@Composable
private fun SearchField(query: String, actions: SearchActions, focus: Boolean) {
    val focusRequester = remember { FocusRequester() }
    LaunchedEffect(focus) { if (focus) focusRequester.requestFocus() }
    val shape = RoundedCornerShape(FsRadius.control)
    val placeholder = stringResource(R.string.search_placeholder)
    Row(
        Modifier
            .padding(start = FsSpace.space4, end = FsSpace.space4, top = FsSpace.space2)
            .fillMaxWidth()
            .heightIn(min = FIELD_HEIGHT)
            .background(FIELD_GROUND, shape)
            .border(1.dp, FsColor.line, shape)
            .padding(start = FsSpace.space3),
        verticalAlignment = Alignment.CenterVertically,
    ) {
        Icon(Icons.Filled.Search, contentDescription = null, tint = FsColor.muted, modifier = Modifier.size(GLYPH))
        BasicTextField(
            value = query,
            onValueChange = actions.onQuery,
            singleLine = true,
            textStyle = FsType.body.copy(color = FsColor.strong),
            cursorBrush = SolidColor(FsColor.gold),
            keyboardOptions = KeyboardOptions(imeAction = ImeAction.Search),
            modifier = Modifier
                .weight(1f)
                .padding(start = FsSpace.space2, end = FsSpace.space3)
                .focusRequester(focusRequester)
                .semantics { contentDescription = placeholder },
            decorationBox = { field ->
                Box(contentAlignment = Alignment.CenterStart) {
                    if (query.isEmpty()) Text(placeholder, style = FsType.body, color = FsColor.muted, maxLines = 1)
                    field()
                }
            },
        )
        if (query.isNotEmpty()) {
            IconButton(onClick = { actions.onQuery("") }) {
                Icon(Icons.Filled.Close, contentDescription = stringResource(R.string.search_clear), tint = FsColor.muted)
            }
        }
    }
}

/** A 36 dp chip inside a 48 dp target; the selected one is gold. The item page's strip. */
@Composable
fun ScopeChip(label: String, selected: Boolean, onClick: () -> Unit) {
    Box(
        Modifier.heightIn(min = 48.dp).selectable(selected = selected, role = Role.Tab, onClick = onClick),
        contentAlignment = Alignment.Center,
    ) {
        val line = if (selected) FsColor.gold else FsColor.line
        Box(
            Modifier.heightIn(min = 36.dp)
                .background(if (selected) FsColor.gold.copy(alpha = SELECTED_WASH) else Color.Transparent, CircleShape)
                .border(1.dp, line, CircleShape)
                .padding(horizontal = CHIP_PADDING, vertical = 6.dp),
            contentAlignment = Alignment.Center,
        ) {
            Text(
                label,
                style = ScaledType.body.copy(fontSize = CHIP_TEXT, fontWeight = if (selected) FontWeight.Bold else FontWeight.Normal),
                color = if (selected) FsColor.gold else FsColor.text,
            )
        }
    }
}

/** The hint, then For {name}, Watching and Recent with its Clear. A group with nothing in it is not drawn. */
private fun LazyListScope.emptyState(state: SearchScreenState, copy: SearchCopy, actions: SearchActions) {
    item(key = "hint") { FlowNote(stringResource(R.string.search_hint), Modifier.padding(top = FsSpace.space5)) }
    state.forCharacter?.let { mine -> item(key = "for") { ForGroup(mine, state, actions) } }
    if (state.watching.isNotEmpty()) {
        item(key = "watching") {
            FlowSection(stringResource(R.string.search_group_watching)) {
                val targets = state.watching.map { SearchTarget.Character(it) }
                targets.forEachIndexed { index, target ->
                    TargetRow(target, copy.meta(target), state, last = index == targets.lastIndex) { actions.onOpen(target) }
                }
            }
        }
    }
    if (state.recents.isNotEmpty()) {
        item(key = "recent") {
            FlowSection(stringResource(R.string.search_group_recent), action = stringResource(R.string.search_recent_clear), onAction = actions.onClearRecent) {
                state.recents.forEachIndexed { index, target ->
                    TargetRow(target, copy.kind(target), state, last = index == state.recents.lastIndex) { actions.onOpen(target) }
                }
            }
        }
    }
    indexState(state.indexPhase, actions)
}

@Composable
private fun ForGroup(mine: ForCharacter, state: SearchScreenState, actions: SearchActions) {
    FlowSection(stringResource(R.string.search_group_forCharacter, mine.name)) {
        val links = buildList {
            add(QuickLink(stringResource(R.string.search_forCharacter_bis, mine.name), false, actions.onOpenBis))
            mine.guideLabel?.let { add(QuickLink(it, WebHandoff.guidesOpenWebsite(state.guidesInApp), actions.onOpenGuide)) }
            add(QuickLink(stringResource(R.string.search_forCharacter_rankings), WebHandoff.rankingsOpenWebsite(state.rankingsInApp), actions.onOpenRankings))
        }
        links.forEachIndexed { index, link ->
            FlowRow(link.label, slim = true, last = index == links.lastIndex, trail = externalMark(link.opensWebsite), onClick = link.onClick)
        }
    }
}

/** The groups, the offline line, no results. */
private fun LazyListScope.results(state: SearchScreenState, copy: SearchCopy, actions: SearchActions) {
    if (state.characterPhase == CharacterPhase.Offline) {
        item(key = "offline") { FlowNote(stringResource(R.string.search_offline), Modifier.padding(top = FsSpace.space4)) }
    }
    if (state.sections.isEmpty() && state.indexPhase is IndexPhase.Ready) {
        item(key = "none") { NoResults(state.query.trim()) }
        return
    }
    indexState(state.indexPhase, actions)
    state.sections.forEach { section -> item(key = section.group.name) { ResultGroup(section, copy, state, actions) } }
    if (state.characterPhase == CharacterPhase.Failed && state.scope in setOf(SearchScope.All, SearchScope.Characters)) {
        item(key = "charactersFailed") {
            FlowNote(stringResource(R.string.search_charactersFailed), Modifier.padding(top = FsSpace.space4), FsColor.wipe)
        }
    }
}

@Composable
private fun ResultGroup(section: SearchSection, copy: SearchCopy, state: SearchScreenState, actions: SearchActions) {
    FlowSection(copy.groupLabel(section.group)) {
        section.rows.forEachIndexed { index, target ->
            TargetRow(target, copy.meta(target), state, last = index == section.rows.lastIndex && !section.showsAll) { actions.onOpen(target) }
        }
        if (section.showsAll) {
            FlowRow(stringResource(R.string.search_showAll, section.total), titleColor = FsColor.gold, slim = true, last = true) {
                actions.onScope(section.group.scope)
            }
        }
    }
}

/** A result, Watching or Recent row: its art, the coloured title, its line; `You` on the account's own. */
@Composable
private fun TargetRow(target: SearchTarget, sub: String, state: SearchScreenState, last: Boolean, onClick: () -> Unit) {
    val copy = SearchCopy(LocalResources.current)
    val own = (target as? SearchTarget.Character)?.character?.key in state.ownKeys
    val website = state.opensWebsite(target)
    val you = stringResource(R.string.search_character_you)
    FlowRow(
        title = copy.title(target),
        sub = sub,
        lead = leadOf(target, state.build),
        titleColor = titleColor(target),
        last = last,
        trail = if (own || website) {
            {
                Row(horizontalArrangement = Arrangement.spacedBy(FsSpace.space2), verticalAlignment = Alignment.CenterVertically) {
                    if (own) FlowTag(you, FsColor.gold, FsColor.gold)
                    if (website) ExternalGlyph()
                }
            }
        } else {
            null
        },
        onClick = onClick,
    )
}

private fun leadOf(target: SearchTarget, build: String): FlowLead = when (target) {
    is SearchTarget.Item -> FlowLead.Item(ItemLook(target.item.id, target.item.name, target.item.quality, target.item.icon), build)
    is SearchTarget.Character -> FlowLead.Crest(target.character.className.orEmpty())
    is SearchTarget.Guide -> FlowLead.Icon(target.guide.icon, build)
    is SearchTarget.Talent -> FlowLead.Icon(target.talent.icon, build)
}

private fun externalMark(opensWebsite: Boolean): (@Composable () -> Unit)? = if (opensWebsite) ({ ExternalGlyph() }) else null

@Composable
private fun ExternalGlyph() {
    Icon(painterResource(R.drawable.ic_open_external), contentDescription = null, tint = FsColor.muted, modifier = Modifier.size(GLYPH))
}

/** The index's own states: loading, or failed with Try again. */
private fun LazyListScope.indexState(phase: IndexPhase, actions: SearchActions) {
    when (phase) {
        IndexPhase.Loading -> item(key = "indexLoading") {
            FlowNote(stringResource(R.string.search_indexLoading), Modifier.padding(top = FsSpace.space4))
        }
        IndexPhase.Failed -> item(key = "indexFailed") {
            Column(Modifier.padding(top = FsSpace.space4)) {
                FlowNote(stringResource(R.string.search_indexFailed), color = FsColor.wipe)
                FlowButton(stringResource(R.string.search_tryAgain), actions.onRetryIndex, Modifier.padding(top = FsSpace.space2))
            }
        }
        is IndexPhase.Ready -> Unit
    }
}

@Composable
private fun NoResults(query: String) {
    Column(Modifier.padding(top = FsSpace.space5)) {
        Text(stringResource(R.string.search_noResults_title, query), style = FsType.item, color = FsColor.strong)
        FlowNote(stringResource(R.string.search_noResults_body))
    }
}

/** One For link: its words, whether it opens the website, and the tap. */
private data class QuickLink(val label: String, val opensWebsite: Boolean, val onClick: () -> Unit)

private val FIELD_HEIGHT = 44.dp
private val GLYPH = 18.dp
private val FIELD_GROUND = Color(0xFF0A0D14)
private const val SELECTED_WASH = 0.10f
private val CHIP_PADDING = 12.dp
private val CHIP_TEXT = 15.sp
