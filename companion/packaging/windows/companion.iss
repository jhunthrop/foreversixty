; companion/packaging/windows/companion.iss
; The Windows setup program, foreversixty-companion-setup.exe, built by Inno Setup 6 in
; companion.yml's Windows build and signed there like the companion itself:
;
;   iscc /DAppVersion=0.2.45 /DSourceExe=dist\foreversixty-companion_windows_amd64.exe /Odist companion.iss
;
; One screen, then the companion: it installs for this Windows account alone, under
; %LOCALAPPDATA%\Programs\Forever Sixty, with no admin rights and no UAC prompt. The one page
; asks two things, "Start with Windows" (on) and a desktop shortcut (off), and its button
; installs. Setup then opens the companion, which goes straight into its own first run
; (sign in with Battle.net, find the _forever_ folder). There is no welcome, licence, folder,
; Start menu folder, ready or finish page.
;
; The folder is the user's own, so the companion's self-update (internal/updater) swaps its
; exe in place there with no rights it does not have. Installing over an older copy is an
; upgrade: the running companion is stopped first and the answers given last time are kept.
; The uninstaller stops the companion, removes its files, shortcuts and the Start with
; Windows entry, and asks before deleting the settings, logs and unsent reports in
; %APPDATA%\ForeverSixty (internal/paths).

#ifndef AppVersion
  #define AppVersion "0.0.0"
#endif
#ifndef SourceExe
  #define SourceExe "..\..\dist\foreversixty-companion_windows_amd64.exe"
#endif

#define AppName "Forever Sixty"
#define ExeName "foreversixty-companion.exe"
; The startup entry's name under HKCU\...\Run.
#define RunValue "Forever Sixty"
; The device token's Credential Manager entry: internal/secret's Service and Account,
; joined as go-keyring names a Windows credential.
#define TokenCredential "ForeverSixty Companion:device-token"

[Setup]
; The id Apps & Features keys this install by; never change it, or an upgrade installs twice.
AppId={{FAC9C192-8284-4AEC-867A-1BF33AD93B7B}
AppName={#AppName}
AppVersion={#AppVersion}
AppVerName={#AppName} {#AppVersion}
AppPublisher=Commish LLC
AppPublisherURL=https://foreversixty.gg
AppSupportURL=https://foreversixty.gg/setup
AppCopyright=Commish LLC
VersionInfoVersion={#AppVersion}
VersionInfoCompany=Commish LLC
VersionInfoProductName=Forever Sixty companion
VersionInfoDescription=Forever Sixty companion setup
PrivilegesRequired=lowest
DefaultDirName={localappdata}\Programs\{#AppName}
UsePreviousAppDir=no
DisableWelcomePage=yes
DisableDirPage=yes
DisableProgramGroupPage=yes
DisableReadyPage=yes
DisableFinishedPage=yes
ArchitecturesAllowed=x64compatible
ArchitecturesInstallIn64BitMode=x64compatible
WizardStyle=modern
WizardSizePercent=100
SetupIconFile=..\..\internal\icon\assets\seal.ico
UninstallDisplayName={#AppName}
UninstallDisplayIcon={app}\{#ExeName}
OutputBaseFilename=foreversixty-companion-setup
Compression=lzma2/max
SolidCompression=yes
; The companion is stopped in PrepareToInstall below; this catches anything else holding the exe.
CloseApplications=force
RestartApplications=no

[Messages]
SelectTasksLabel2=
WizardSelectTasks=Install {#AppName}
SelectTasksDesc=The companion for your Windows account. No admin needed.

[Tasks]
Name: startup; Description: "Start with Windows"
Name: desktopicon; Description: "Desktop shortcut"; Flags: unchecked

[Files]
Source: "{#SourceExe}"; DestDir: "{app}"; DestName: "{#ExeName}"; Flags: ignoreversion

[Icons]
Name: "{userprograms}\{#AppName}"; Filename: "{app}\{#ExeName}"
Name: "{userdesktop}\{#AppName}"; Filename: "{app}\{#ExeName}"; Tasks: desktopicon

[Registry]
Root: HKCU; Subkey: "Software\Microsoft\Windows\CurrentVersion\Run"; ValueType: string; ValueName: "{#RunValue}"; ValueData: """{app}\{#ExeName}"""; Tasks: startup; Flags: uninsdeletevalue
; Unticked on an upgrade: the old entry goes.
Root: HKCU; Subkey: "Software\Microsoft\Windows\CurrentVersion\Run"; ValueType: none; ValueName: "{#RunValue}"; Tasks: not startup; Flags: deletevalue

[Run]
; Straight into the companion's own first run. Skipped for a silent install.
Filename: "{app}\{#ExeName}"; Flags: nowait skipifsilent

[UninstallDelete]
; What the self-update leaves beside the exe (internal/updater: applyBinary).
Type: files; Name: "{app}\{#ExeName}.previous"
Type: files; Name: "{app}\{#ExeName}.update-*"
Type: dirifempty; Name: "{app}"

[Code]
const
  StopWaitMs = 3000;

{ Stops every companion this account runs: asked to close first, then ended. }
procedure StopCompanion();
var
  Code: Integer;
begin
  Exec(ExpandConstant('{sys}\taskkill.exe'), '/IM {#ExeName}', '', SW_HIDE, ewWaitUntilTerminated, Code);
  Sleep(StopWaitMs);
  Exec(ExpandConstant('{sys}\taskkill.exe'), '/F /IM {#ExeName}', '', SW_HIDE, ewWaitUntilTerminated, Code);
end;

function PrepareToInstall(var NeedsRestart: Boolean): String;
begin
  StopCompanion();
  Result := '';
end;

function InitializeUninstall(): Boolean;
begin
  StopCompanion();
  Result := True;
end;

procedure CurUninstallStepChanged(CurUninstallStep: TUninstallStep);
var
  Data: String;
  Code: Integer;
begin
  if CurUninstallStep <> usPostUninstall then
    exit;
  Data := ExpandConstant('{userappdata}\ForeverSixty');
  if UninstallSilent or not DirExists(Data) then
    exit;
  if MsgBox('Also delete the companion''s settings, logs and any fights still waiting to upload? ' +
            'Your reports on foreversixty.gg stay either way.',
            mbConfirmation, MB_YESNO or MB_DEFBUTTON2) = IDYES then
  begin
    DelTree(Data, True, True, True);
    Exec(ExpandConstant('{sys}\cmdkey.exe'), '/delete:"{#TokenCredential}"', '', SW_HIDE, ewWaitUntilTerminated, Code);
  end;
end;
