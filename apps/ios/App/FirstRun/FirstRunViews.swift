import Auth
import DesignSystem
import Models
import SwiftUI

/// F8 step 2 (mock ios-first-pick-main): the account's characters at level 10 and up, the highest
/// preselected and tagged Main, and one button that continues with the selection.
struct PickMainView: View {
    let characters: [AccountCharacter]
    let selection: String?
    let saving: Bool
    let failed: Bool
    let onSelect: (String) -> Void
    let onContinue: () -> Void

    var body: some View {
        ScrollView {
            VStack(alignment: .leading, spacing: 0) {
                PageHeader(title: Copy.firstPickMainTitle) {
                    SmallText(Copy.firstPickMainMeta)
                }
                SmallText(Copy.firstPickMainLead).padding(.top, FSSpace.space2)
                FSSection(Copy.firstPickMainSection) {
                    VStack(spacing: 0) {
                        ForEach(characters, id: \.key) { row($0) }
                    }
                }
                .padding(.top, FirstRunLayout.sectionGap)
                VStack(spacing: FSSpace.space3) {
                    if let chosen {
                        FSButton(Copy.firstPickMainContinue(name: chosen.name), kind: .primary, action: onContinue)
                            .disabled(saving)
                    }
                    if failed {
                        SmallText(Copy.firstPickMainFailed, color: FSColor.wipe)
                    }
                }
                .padding(.top, FSSpace.space5)
            }
            .padding(.top, FSSpace.space2)
            .padding(.horizontal, FSSpace.gutterPhone)
            .padding(.bottom, FSSpace.space6)
            .parityContent()
        }
        .background(FSColor.bg)
    }

    private var chosen: AccountCharacter? { characters.first { $0.key == selection } }

    private func row(_ character: AccountCharacter) -> some View {
        let isSelected = character.key == selection
        return Button { onSelect(character.key) } label: {
            FSRow(
                title: character.name,
                titleColor: ClassColor.of(character.className),
                sub: AccountCopy.characterLine(character),
                selected: isSelected,
                inset: isSelected ? FSSpace.rowPadX : 0,
                trailPlacement: .below,
                lead: { ClassCrest(className: character.className, size: FSSize.crest40, flat: true) },
                trail: {
                    if isSelected { FSTag(Copy.menuMain, tone: .gold) }
                }
            )
        }
        .buttonStyle(RowPressStyle())
    }
}

/// F8 step 3 (mock ios-first-permission): the guild by name, what the notification does, and the
/// ask. The system prompt follows Turn on notifications.
struct FirstPermissionView: View {
    let guild: String
    let onTurnOn: () -> Void
    let onNotNow: () -> Void

    var body: some View {
        ScrollView {
            VStack(spacing: 0) {
                VStack(spacing: 0) {
                    if let seal = BrandImages.seal {
                        Image(uiImage: seal)
                            .resizable()
                            .frame(width: FSSize.crest84, height: FSSize.crest84)
                            .accessibilityHidden(true)
                            .padding(.top, FSSpace.space10)
                            .padding(.bottom, FSSpace.space2)
                    }
                    SmallText(Copy.firstPermissionEyebrow)
                        .multilineTextAlignment(.center)
                        // `.center .eyebrow { margin-bottom: -12px }`: the guild title sits just under it.
                        .padding(.bottom, -FirstRunLayout.eyebrowPull)
                    PageHeader(title: guild, alignment: .center, balanced: true)
                    SmallText(Copy.firstPermissionBody)
                        .multilineTextAlignment(.center)
                        .frame(maxWidth: FirstRunLayout.permissionMeasure)
                        .padding(.top, FSSpace.space2)
                }
                VStack(spacing: FSSpace.space3) {
                    FSButton(Copy.firstPermissionTurnOn, kind: .primary, action: onTurnOn)
                    FSButton(Copy.firstPermissionNotNow, action: onNotNow)
                }
                .padding(.top, FSSpace.space5)
            }
            .padding(.top, FSSpace.space2)
            .padding(.horizontal, FSSpace.gutterPhone)
            .parityContent()
        }
        .background(FSColor.bg)
    }
}

enum FirstRunLayout {
    /// `.ap section { margin-top: 28px }`.
    static let sectionGap: CGFloat = FSSpace.space7
    /// `.center .eyebrow { margin-bottom: -12px }`.
    static let eyebrowPull: CGFloat = 12
    /// `.center .t-small { max-width: 300px }`.
    static let permissionMeasure: CGFloat = 300
}

/// The first run over the tabs: one full-screen step at a time, driven by the model.
struct FirstRunHost: ViewModifier {
    let session: SessionStore
    @Bindable var model: FirstRunModel
    /// Off while the sign-in switch is off: there is no account to pick a main for.
    let isOn: Bool

    func body(content: Content) -> some View {
        content
            .fullScreenCover(item: Binding(get: { isOn ? model.step : nil }, set: { if $0 == nil { model.notNow() } })) { step in
                Group {
                    switch step {
                    case .pickMain:
                        PickMainView(
                            characters: session.phase.me.map(FirstRunModel.candidates) ?? [],
                            selection: model.selection,
                            saving: model.saving,
                            failed: model.failed,
                            onSelect: { model.selection = $0 },
                            onContinue: { Task { await model.continueWithMain(session.phase.me) } }
                        )
                        .interactiveDismissDisabled()
                    case let .permission(guild):
                        FirstPermissionView(
                            guild: guild,
                            onTurnOn: { Task { await model.turnOnNotifications() } },
                            onNotNow: { model.notNow() }
                        )
                    }
                }
                .tint(FSColor.gold)
            }
            .onChange(of: session.phase.me, initial: true) { _, me in model.update(me: me) }
    }
}
