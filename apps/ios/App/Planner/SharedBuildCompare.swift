import Auth
import DesignSystem
import Models
import SwiftUI

/// Compare with mine (spec F13.5, lane F15): F6's compare with the shared build in the guide's
/// place and your talent export on the other side. Pure; one compare draws both (TalentsPresenter).
enum SharedBuildCompare {
    /// The character whose export the compare reads: the current one when it is of the build's
    /// class and has an export, else your first character of that class with one. Nil hides the
    /// button.
    static func yourCharacter(classSlug: String, current: AccountCharacter?, characters: [AccountCharacter]) -> AccountCharacter? {
        guard !classSlug.isEmpty else { return nil }
        let candidates = [current].compactMap(\.self) + characters
        return candidates.first { $0.className?.lowercased() == classSlug && hasExport($0) }
    }

    /// An export whose points are all 0 holds no ranks (as F6 reads it).
    static func hasExport(_ character: AccountCharacter) -> Bool {
        character.build?.talents?.points.contains { $0 > 0 } ?? false
    }

    /// The compare on the shared build's own trees. An export made on a build whose trees are
    /// laid out differently reads as stale (F6's older-export sentence).
    static func presenter(content: SharedBuildContent, file: TalentsFile, character: AccountCharacter) -> TalentsPresenter {
        let index = PlannerIndex(file)
        let shared = PlannerRules.treeRanks(index, ranks: PlannerRules.ranks(content.summary.build.pointOrder))
        let trees = character.build?.talents?.trees ?? []
        let sameTrees = trees.count == file.trees.count
        let export: TalentExport = sameTrees ? .ranks(trees.map { TalentRules.ranks(ofTree: $0) }) : .stale
        return TalentsPresenter(
            file: file, guide: .shared(treeRanks: shared), export: export, level: character.level,
            specName: character.spec ?? ""
        )
    }
}

/// The compare screen pushed from the viewer: F6's TalentsView with the shared-build words.
struct SharedBuildCompareScreen: View {
    let presenter: TalentsPresenter
    let characterName: String
    @State private var shown: TalentBuildKind = .yours
    @State private var tree: Int
    @State private var sheet: TalentSheetModel?

    init(presenter: TalentsPresenter, characterName: String) {
        self.presenter = presenter
        self.characterName = characterName
        _tree = State(initialValue: presenter.openingTree)
    }

    var body: some View {
        TalentsView(
            phase: .loaded(presenter, offlineSince: nil), build: presenter.file.build, offlineStamp: nil,
            shown: $shown, tree: $tree, selectedID: sheet?.id,
            openTalent: { sheet = presenter.sheet(shown, talentID: $0) },
            compareFooter: Copy.planCmpOn(name: characterName)
        )
        .navigationTitle(Copy.planCmpTitle)
        .navigationBarTitleDisplayMode(.inline)
        .toolbarBackground(FSColor.bg, for: .navigationBar)
        .sheet(item: $sheet) { TalentRankSheet(model: $0, build: presenter.file.build) }
    }
}
