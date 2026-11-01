package gg.foreversixty.app.search.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.account.AccountCopy
import gg.foreversixty.app.bis.BisLabels
import gg.foreversixty.app.items.ItemVocabulary
import gg.foreversixty.app.model.CharacterSearchRow
import gg.foreversixty.app.model.SearchGuide
import gg.foreversixty.app.model.SearchItem
import gg.foreversixty.app.model.SearchTalent
import gg.foreversixty.app.raids.logic.GameClass
import java.util.Locale

/** The words of Search's rows (spec F12.2), all from the copy catalog. iOS: SearchCopy.swift. */
class SearchCopy(private val res: Resources) {
    /** The class's own name (`Shaman`); an unknown slug is capitalised. */
    fun className(slug: String?): String {
        slug ?: return ""
        return GameClass.of(slug)?.let { res.getString(it.label) } ?: slug.replaceFirstChar { it.titlecase(Locale.ROOT) }
    }

    /** `Elemental Shaman`: a guide's title, which is also what the field matches. */
    fun guideTitle(guide: SearchGuide): String = res.getString(R.string.bis_specChip, guide.specName, className(guide.className))

    /** `Guide · DPS`. */
    fun guideMeta(guide: SearchGuide): String = res.getString(R.string.search_guide_meta, BisLabels.role(res, guide.role))

    /** `Back · item level 52 · requires level 47`, or without the level when none is required. */
    fun itemMeta(item: SearchItem): String {
        val slot = ItemVocabulary.slotLabel(res, item.slot)
        return if (item.requiredLevel > 0) {
            res.getString(R.string.search_item_meta, slot, item.itemLevel, item.requiredLevel)
        } else {
            res.getString(R.string.search_item_metaNoLevel, slot, item.itemLevel)
        }
    }

    /** `Mage · Level 58 · US Normal`. */
    fun characterMeta(row: CharacterSearchRow): String {
        val region = row.region.uppercase(Locale.ROOT)
        val ruleset = AccountCopy.rulesetName(row.ruleset)?.let(res::getString) ?: row.ruleset
        val level = row.level
        return if (level != null) {
            res.getString(R.string.search_character_meta, className(row.className), level, region, ruleset)
        } else {
            res.getString(R.string.search_character_metaNoLevel, className(row.className), region, ruleset)
        }
    }

    /** `Elemental Shaman · Elemental · 3 ranks`. */
    fun talentMeta(talent: SearchTalent): String = res.getString(
        R.string.search_talent_meta,
        talent.specName,
        className(talent.className),
        talent.tree,
        res.getQuantityString(R.plurals.search_talent_ranks, talent.maxRank, talent.maxRank),
    )

    /** A Recent row's kind word: `Item`, `Character`, `Guide`, `Talent`. */
    fun kind(target: SearchTarget): String = res.getString(
        when (target) {
            is SearchTarget.Item -> R.string.search_recent_item
            is SearchTarget.Character -> R.string.search_recent_character
            is SearchTarget.Guide -> R.string.search_recent_guide
            is SearchTarget.Talent -> R.string.search_recent_talent
        },
    )

    /** A row's title: the item, character or talent name, or the guide's spec and class. */
    fun title(target: SearchTarget): String = when (target) {
        is SearchTarget.Item -> target.item.name
        is SearchTarget.Character -> target.character.name
        is SearchTarget.Guide -> guideTitle(target.guide)
        is SearchTarget.Talent -> target.talent.name
    }

    /** A result row's meta line. */
    fun meta(target: SearchTarget): String = when (target) {
        is SearchTarget.Item -> itemMeta(target.item)
        is SearchTarget.Character -> characterMeta(target.character)
        is SearchTarget.Guide -> guideMeta(target.guide)
        is SearchTarget.Talent -> talentMeta(target.talent)
    }

    fun groupLabel(group: SearchGroup): String = res.getString(
        when (group) {
            SearchGroup.Guides -> R.string.search_group_guides
            SearchGroup.Talents -> R.string.search_group_talents
            SearchGroup.Items -> R.string.search_group_items
            SearchGroup.Characters -> R.string.search_group_characters
        },
    )

    fun scopeLabel(scope: SearchScope): String = res.getString(
        when (scope) {
            SearchScope.All -> R.string.search_scope_all
            SearchScope.Items -> R.string.search_scope_items
            SearchScope.Characters -> R.string.search_scope_characters
            SearchScope.Guides -> R.string.search_scope_guides
            SearchScope.Talents -> R.string.search_scope_talents
        },
    )
}
