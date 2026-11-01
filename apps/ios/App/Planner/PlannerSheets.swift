import DesignSystem
import DesignTokens
import Models
import SwiftUI
import UIKit

/// The Race chip's sheet: the class's legal races (spec F13.1).
struct PlannerRaceSheet: View {
    let races: [RaceInfo]
    let selected: String?
    let pick: (String) -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            Text(Copy.planRaceSheet).font(ScaledType.sec).foregroundStyle(FSColor.strong)
                .textCase(.uppercase)
                .padding(EdgeInsets(top: 20, leading: 16, bottom: 8, trailing: 16))
                .accessibilityAddTraits(.isHeader)
            ScrollView {
                VStack(spacing: 0) {
                    ForEach(races, id: \.id) { race in
                        Button { pick(race.slug) } label: {
                            HStack {
                                Text(race.name).font(ScaledType.bodyStrong).foregroundStyle(FSColor.strong)
                                Spacer()
                                if race.slug == selected {
                                    Image(systemName: "checkmark").foregroundStyle(FSColor.gold).accessibilityHidden(true)
                                }
                            }
                            .padding(.horizontal, 16)
                            .frame(minHeight: 52)
                            .contentShape(.rect)
                        }
                        .buttonStyle(.plain)
                        .accessibilityAddTraits(race.slug == selected ? .isSelected : [])
                        FSColor.lineSoft.frame(height: 1).padding(.leading, 16)
                    }
                }
            }
        }
        .background(FSColor.bg)
        .presentationDetents([.medium, .large])
    }
}

/// The Level chip's sheet: a 10 to 60 stepper that never goes below the points spent (F13.1).
struct PlannerLevelSheet: View {
    let level: Int
    let minimum: Int
    let set: (Int) -> Void

    var body: some View {
        VStack(alignment: .leading, spacing: 14) {
            Text(Copy.planLevelSheet).font(ScaledType.sec).foregroundStyle(FSColor.strong)
                .textCase(.uppercase)
                .accessibilityAddTraits(.isHeader)
            Stepper {
                Text(Copy.planLevel(n: level)).font(ScaledType.monoDay).foregroundStyle(FSColor.strong)
            } onIncrement: {
                set(min(PlannerRules.maxLevel, level + 1))
            } onDecrement: {
                set(max(minimum, level - 1))
            }
            Text(Copy.planLevelSheetLine(n: PlannerRules.pointsAvailable(atLevel: level)))
                .font(ScaledType.meta).foregroundStyle(FSColor.muted)
        }
        .padding(16)
        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .topLeading)
        .background(FSColor.bg)
        .presentationDetents([.height(220), .medium])
    }
}

/// F13.3 Share this build: the web's confirm step, verbatim, then Share anyway or Cancel.
struct PlannerShareSheet: View {
    @Binding var title: String
    let signedIn: Bool
    let state: PlannerShareState
    let proceed: () -> Void
    let cancel: () -> Void

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 14) {
                Text(Copy.planShareTitle).font(ScaledType.sec).foregroundStyle(FSColor.strong)
                    .textCase(.uppercase)
                    .accessibilityAddTraits(.isHeader)
                Text(Copy.planShareIntro).font(ScaledType.body).foregroundStyle(FSColor.text)
                    .fixedSize(horizontal: false, vertical: true)
                VStack(alignment: .leading, spacing: 6) {
                    ForEach(fields, id: \.self) { field in
                        Text("· " + field).font(ScaledType.body).foregroundStyle(FSColor.text)
                            .fixedSize(horizontal: false, vertical: true)
                    }
                }
                titleField
                if case let .failed(message) = state {
                    Text(message).font(ScaledType.meta).foregroundStyle(FSColor.wipe)
                        .fixedSize(horizontal: false, vertical: true)
                }
                Button(Copy.planShareProceed, action: proceed)
                    .buttonStyle(BrandButtonStyle(.primary, height: 48))
                    .disabled(state == .sending)
                Button(Copy.commonCancel, action: cancel)
                    .buttonStyle(BrandButtonStyle(.quiet, height: 48))
                if signedIn {
                    Text(Copy.planShareSignedIn).font(ScaledType.meta).foregroundStyle(FSColor.muted)
                        .fixedSize(horizontal: false, vertical: true)
                }
            }
            .padding(EdgeInsets(top: 22, leading: 16, bottom: 16, trailing: 16))
        }
        .background(FSColor.bg)
    }

    private var trimmed: String { title.trimmingCharacters(in: .whitespacesAndNewlines) }

    private var fields: [String] {
        [Copy.planShareFieldClassRace, Copy.planShareFieldTalents] + (trimmed.isEmpty ? [] : [Copy.planShareFieldTitle(title: trimmed)])
    }

    private var titleField: some View {
        VStack(alignment: .leading, spacing: 6) {
            MicroLabel(Copy.planShareTitleField)
            HStack {
                TextField("", text: limited)
                    .font(ScaledType.body)
                    .foregroundStyle(FSColor.strong)
                    .accessibilityLabel(Copy.planShareTitleField)
                Text(Copy.planShareTitleCount(n: title.count, max: PlannerEditorModel.titleLimit))
                    .font(ScaledType.monoMeta).foregroundStyle(FSColor.muted)
            }
            .padding(.horizontal, 12)
            .frame(minHeight: 48)
            .background(FSColor.raised, in: .rect(cornerRadius: FSRadius.control))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(FSColor.gold, lineWidth: 1))
        }
    }

    /// The field never holds more than 60 characters.
    private var limited: Binding<String> {
        Binding(get: { title }, set: { title = String($0.prefix(PlannerEditorModel.titleLimit)) })
    }
}

/// The system share sheet over the build's link (spec F13.4).
struct SystemShareSheet: UIViewControllerRepresentable {
    let url: URL

    func makeUIViewController(context _: Context) -> UIActivityViewController {
        UIActivityViewController(activityItems: [url], applicationActivities: nil)
    }

    func updateUIViewController(_: UIActivityViewController, context _: Context) {}
}
