import { defaultVpSimple } from "@/configs/defaultVpSimple";

export const EnShell = {
  //终端配置
  config: "Terminal Config",
  //运行日志
  runLog: "Run Log",
  //默认值
  default: "Default",
  //暂无运行中的终端
  noRunTerminal: "No Running Terminal",
  //起始目录不能为空
  runPathEmpty: "Run Path Can Not Be Empty",
  //"cmd不能为空"
  cmdEmpty: "Cmd Can Not Be Empty",
  //git commit 弹窗
  commitMessageTitle: "Git Commit",
  commitMessagePlaceholder: "Enter commit message",
  commitMessageEmpty: "Please enter a commit message",
  // Upload to server
  uploadToServer: "Upload to Server",
  uploadModalTitle: "Upload to Server",
  uploadHost: "Server Address",
  uploadPort: "SSH Port",
  uploadUsername: "Username",
  uploadPassword: "Password",
  uploadLocalPath: "Local Path",
  uploadRemotePath: "Remote Path",
  uploadSelectLocal: "Select Local Path",
  uploadSelectLocalTitle: "Select file or directory to upload",
  uploadSuccess: "Upload Successful",
  uploadFailed: "Upload Failed",
  uploadFormEmpty: "Please fill in server address, username and password",
  uploadLocalPathEmpty: "Please select local path",
  uploadRemotePathEmpty: "Please fill in remote path",
  uploadPrivateKeyEmpty: "Please select SSH private key file",
  uploadAuthMode: "Auth Mode",
  uploadAuthPassword: "Password",
  uploadAuthKey: "SSH Key",
  uploadPrivateKey: "Private Key",
  uploadPrivateKeyPlaceholder: "Select id_rsa or id_ed25519 etc.",
  uploadSelectKey: "Select Key",
  uploadSelectKeyTitle: "Select SSH private key file (e.g. id_rsa)",
  uploadKeyPassphrase: "Key Passphrase",
  uploadKeyPassphrasePlaceholder: "If key is encrypted (optional)",
  uploadHostPlaceholder: "e.g. 192.168.1.100 or example.com",
  uploadUsernamePlaceholder: "SSH username",
  uploadPasswordPlaceholder: "SSH password",
  uploadLocalPathPlaceholder: "Click button to select",
  uploadRemotePathPlaceholder: "e.g. /var/www/html",
  errorWailsNotAvailable: "Please run in desktop app (wails dev or built executable)",

  docsBuild: "Docs Build Command",
  //运行目录
  labels: {
    //默认目录
    runPath: "Default Run Path",
    //git运行的目录
    gitPath: "Git Run Path",
    //文档dev命令
    docsDev: "Docs Dev Command",
    //文档build命令
    docsBuild: "Docs Build Command",
  },
  tips: {
    //运行目录的路径，默认值是项目根目录
    runPath: "Run Path, Default Value Is Project Root Path",
    //git运行的目录，一般为vitepress项目根目录，但有时vitepres是做其它项目的子项目，所以需要指定
    gitPath:
      "Git Run Path, Default Value Is Vitepress Project Root Path, But Sometimes Vitepress Is A Sub Project Of Other Projects, So You Need To Specify",
    //文档dev命令，默认值是vitepress项目根目录下的dev命令
    docsDev: "Docs Dev Command, Default Value Is " + defaultVpSimple.cmdDocsDev,
    //文档build命令，默认值是vitepress项目根目录下的build命令
    docsBuild:
      "Docs Build Command, Default Value Is" + defaultVpSimple.cmdDocsBuild,
  },
};
