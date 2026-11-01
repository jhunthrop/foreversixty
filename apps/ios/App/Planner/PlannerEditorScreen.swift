import Auth
import DesignSystem
import DesignTokens
import SwiftUI

/// The editor as presented: the model, the open tree, the Race and Level sheets, the Reset alert,
/// the share confirm and the system share, and the hardware keyboard (spec F13.1, §8: arrows move
/// between talents, + and - add and remove).
struct PlannerEditorScreen: View {
    @State private var model: PlannerEditorModel
    @State private var tree = 0
    @State private var settled = false
    @State private var sheet: Sheet?
    @State private var confirmReset = false
    @State private var shareTitle = ""
    @FocusState private var focused: Bool
    @Environment(\.dismiss) private var dismiss
    @Environment(SessionStore.self) private var session: SessionStore?

    private enum Sheet: Identifiable {
        case race, level, share
        var id: Self { self }
    }

    init(launch: PlannerLaunch, dependencies: PlannerDependencies) {
        _model = State(initialValue: PlannerEditorModel(launch: launch, dependencies: dependencies))
    }

    var body: some View {
        PlannerEditorView(
            model: model, tree: $tree, build: model.state?.index.file.build ?? "",
            actions: PlannerEditorActions(
                back: { dismiss() },
                openRace: { sheet = .race },
                openLevel: { sheet = .level },
                reset: { if (model.state?.spent ?? 0) > 0 { confirmReset = true } },
                share: { sheet = .share },
                retry: { Task { await model.load() } }
            )
        )
        .focusable()
        .focused($focused)
        .focusEffectDisabled()
        .onKeyPress(keys: [.upArrow, .leftArrow]) { _ in model.moveFocus(by: -1); return .handled }
        .onKeyPress(keys: [.downArrow, .rightArrow]) { _ in model.moveFocus(by: 1); return .handled }
        .onKeyPress(characters: CharacterSet(charactersIn: "+=")) { _ in model.addFocused(); return .handled }
        .onKeyPress(characters: CharacterSet(charactersIn: "-_")) { _ in model.removeFocused(); return .handled }
        .alert(Copy.planResetConfirm(n: model.state?.spent ?? 0), isPresented: $confirmReset) {
            Button(Copy.planResetProceed, role: .destructive) { model.reset() }
            Button(Copy.commonCancel, role: .cancel) {}
        }
        .sheet(item: $sheet) { sheet in content(sheet) }
        .sheet(isPresented: systemShare) {
            if case let .shared(url) = model.share { SystemShareSheet(url: url).presentationDetents([.medium, .large]) }
        }
        .task {
            if model.state == nil { await model.load() }
            settle()
            focused = true
        }
    }

    /// Opens on the tree holding the most points, once.
    private func settle() {
        guard !settled, let state = model.state else { return }
        settled = true
        tree = TalentRules.openingTree(state.totals)
    }

    @ViewBuilder
    private func content(_ sheet: Sheet) -> some View {
        switch sheet {
        case .race:
            PlannerRaceSheet(races: model.races, selected: model.state?.raceSlug) { slug in
                model.setRace(slug)
                self.sheet = nil
            }
        case .level:
            PlannerLevelSheet(
                level: model.state?.level ?? PlannerRules.maxLevel,
                minimum: PlannerRules.minimumLevel(spent: model.state?.spent ?? 0),
                set: model.setLevel
            )
        case .share:
            PlannerShareSheet(
                title: $shareTitle, signedIn: session?.phase.isSignedIn ?? false, state: model.share,
                proceed: { Task { await share() } },
                cancel: {
                    model.dismissShare()
                    self.sheet = nil
                }
            )
            .presentationDetents([.large])
        }
    }

    private func share() async {
        await model.share(title: shareTitle)
        switch model.share {
        case .shared, .idle: sheet = nil
        case .sending, .failed: break
        }
    }

    /// The system share sheet shows once the link is back and the confirm has closed.
    private var systemShare: Binding<Bool> {
        Binding(
            get: {
                if case .shared = model.share { return sheet == nil }
                return false
            },
            set: { shown in if !shown { model.dismissShare() } }
        )
    }
}
