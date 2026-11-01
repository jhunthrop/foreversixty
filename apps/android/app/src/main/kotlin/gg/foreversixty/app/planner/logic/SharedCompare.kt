package gg.foreversixty.app.planner.logic

import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.talents.CompareAgainst
import gg.foreversixty.app.talents.TalentExport
import gg.foreversixty.app.talents.TalentGuideBuild
import gg.foreversixty.app.talents.TalentRules
import gg.foreversixty.app.talents.TalentsPresenter
import java.util.Locale

/**
 * F13.5's `Compare with mine` (lane F15 follow-up): F6's compare, with the shared build as the
 * other side and your export of that class as yours. One compare, its words chosen by
 * [CompareAgainst.Shared]. iOS: SharedCompare.swift.
 */
object SharedCompare {
    /**
     * The export to compare: the current character's when it is of the build's class, else the
     * account's first character of that class with one; null hides the button.
     */
    fun exportFor(classSlug: String, current: AccountCharacter?, characters: List<AccountCharacter>): AccountCharacter? {
        val candidates = listOfNotNull(current) + characters.filter { it.key != current?.key }
        return candidates.firstOrNull { it.className?.lowercase(Locale.ROOT) == classSlug && hasExport(it) }
    }

    /** An export whose points are all 0 holds no ranks (as F6 reads it). */
    private fun hasExport(character: AccountCharacter): Boolean =
        character.build?.talents?.let { talents -> talents.trees.isNotEmpty() && talents.points.any { it > 0 } } == true

    /** The compare: drawn on the shared build's own trees, its ranks the other side, the export yours. */
    fun presenter(view: SharedBuildView, mine: AccountCharacter): TalentsPresenter {
        val trees = TalentRules.trees(view.file)
        // Per tree in the file's talent order, as F6 reads ranks (the rank sheet's `rankSheet` does the same).
        val shared = trees.map { tree -> tree.talents.map { view.ranks[it.id] ?: 0 } }
        val yours = mine.build?.talents?.trees.orEmpty().map(TalentRules::ranksOfTree)
        return TalentsPresenter(
            file = view.file,
            guide = TalentGuideBuild.Ranks(shared),
            export = TalentExport.Ranks(yours),
            level = mine.level,
            specName = "",
            against = CompareAgainst.Shared(mine.name),
        )
    }
}
