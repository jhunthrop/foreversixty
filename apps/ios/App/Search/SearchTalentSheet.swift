import DesignSystem
import DesignTokens
import Models
import SwiftUI

/// A talent result's sheet (spec F12.2): the F6 rank sheet's body for the talent as the game
/// shows it unlearned (`Rank 0/5` and the first rank's text), then `Open in planner`.
struct SearchTalentSheet: View {
    let talent: SearchTalent
    let build: String
    let index: SearchIndexReader
    let openPlanner: () -> Void

    @State private var model: TalentSheetModel?
    @State private var failed = false

    var body: some View {
        Group {
            if let model {
                TalentRankContent(model: model, build: build)
            } else {
                SearchTalentPlaceholder(talent: talent, build: build, failed: failed) { Task { await load() } }
            }
        }
        .safeAreaInset(edge: .bottom) {
            SecondaryButton(title: Copy.searchTalentOpenPlanner, systemImage: "arrow.up.right", fullWidth: true, action: openPlanner)
                .padding(.horizontal, 20)
                .padding(.bottom, 12)
                .background(FSColor.raised)
        }
        .presentationDetents([.medium, .large])
        .presentationDragIndicator(.visible)
        .presentationBackground(FSColor.raised)
        .task { await load() }
    }

    private func load() async {
        failed = false
        do {
            let file = try await index.talents(className: talent.className, build: build)
            model = SearchTalentSheetModel.make(talentID: talent.id, in: file)
            failed = model == nil
        } catch {
            failed = true
        }
    }
}

/// The sheet's head while the class's talents file loads, or when it did not.
private struct SearchTalentPlaceholder: View {
    let talent: SearchTalent
    let build: String
    let failed: Bool
    let retry: () -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: 16) {
            HStack(spacing: 12) {
                SearchTalentIcon(icon: talent.icon, name: talent.name, build: build, size: 56)
                VStack(alignment: .leading, spacing: 2) {
                    Text(talent.name).font(ScaledType.name).foregroundStyle(FSColor.strong)
                    Text(Copy.searchTalentRank(rank: 0, max: talent.maxRank)).font(ScaledType.monoMeta).foregroundStyle(FSColor.muted)
                }
            }
            if failed {
                Text(Copy.itemsFailed).font(ScaledType.body).foregroundStyle(FSColor.wipe)
                SecondaryButton(title: Copy.commonTryAgain, action: retry)
            }
        }
        .frame(maxWidth: .infinity, alignment: .leading)
        .padding(.horizontal, 20)
        .padding(.top, 24)
        .frame(maxHeight: .infinity, alignment: .top)
    }
}

/// The F6 sheet model for a talent nobody has spent points in.
enum SearchTalentSheetModel {
    static func make(talentID: Int, in file: TalentsFile) -> TalentSheetModel? {
        guard let talent = file.trees.lazy.flatMap(\.talents).first(where: { $0.id == talentID }) else { return nil }
        return TalentSheetModel(
            talent: talent, rank: 0, state: .open, current: talent.ranks.first?.description, next: nil,
            requirements: [], guideLine: nil, isOver: false
        )
    }
}
