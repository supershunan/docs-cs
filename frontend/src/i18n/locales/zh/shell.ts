import { defaultVpSimple } from "@/configs/defaultVpSimple";

export const ZhShell = {
  //终端配置
  config: "终端配置",
  //运行日志
  runLog: "运行日志",
  //默认值
  default: "默认值",
  //暂无运行中的终端
  noRunTerminal: "暂无运行中的终端",
  //起始目录不能为空
  runPathEmpty: "起始目录不能为空",
  //"cmd不能为空"
  cmdEmpty: "cmd不能为空",
  //git commit 弹窗
  commitMessageTitle: "Git 提交",
  commitMessagePlaceholder: "请输入提交描述",
  commitMessageEmpty: "请填写提交描述",
  // 上传文件到服务器
  uploadToServer: "上传文件到服务器",
  uploadModalTitle: "上传到服务器",
  uploadHost: "服务器地址",
  uploadPort: "SSH 端口",
  uploadUsername: "用户名",
  uploadPassword: "密码",
  uploadLocalPath: "本地路径",
  uploadRemotePath: "远程路径",
  uploadSelectLocal: "选择本地路径",
  uploadSelectLocalTitle: "选择要上传的文件或目录",
  uploadHostPlaceholder: "如 192.168.1.100 或 example.com",
  uploadUsernamePlaceholder: "SSH 登录用户名",
  uploadPasswordPlaceholder: "SSH 登录密码",
  uploadLocalPathPlaceholder: "点击右侧按钮选择",
  uploadRemotePathPlaceholder: "如 /var/www/html 或 /home/user/docs",
  uploadSuccess: "上传成功",
  uploadFailed: "上传失败",
  uploadFormEmpty: "请填写服务器地址、用户名和密码",
  uploadLocalPathEmpty: "请选择要上传的本地路径",
  uploadRemotePathEmpty: "请填写远程目标路径",
  uploadPrivateKeyEmpty: "请选择 SSH 私钥文件",
  errorWailsNotAvailable: "此功能需在桌面应用中运行",
  uploadAuthMode: "认证方式",
  uploadAuthPassword: "密码",
  uploadAuthKey: "SSH 密钥",
  uploadPrivateKey: "私钥文件",
  uploadPrivateKeyPlaceholder: "选择 id_rsa 或 id_ed25519 等私钥",
  uploadSelectKey: "选择私钥",
  uploadSelectKeyTitle: "选择 SSH 私钥文件（如 id_rsa）",
  uploadKeyPassphrase: "密钥密码",
  uploadKeyPassphrasePlaceholder: "若私钥已加密请输入密码（可选）",

  docsBuild: "文档build命令",
  //运行目录
  labels: {
    //默认目录
    runPath: "默认命令运行目录",
    //git运行的目录
    gitPath: "git运行的目录",
    //文档dev命令
    docsDev: "文档dev命令",
    //文档build命令
    docsBuild: "文档build命令",
    //文档serve命令
  },
  tips: {
    //运行目录的路径，默认值是项目根目录
    runPath: "运行目录的路径，默认值是项目根目录",
    //git运行的目录，一般为vitepress项目根目录，但有时vitepres是做其它项目的子项目，所以需要指定git运行的目录
    gitPath:
      "git运行的目录，一般为vitepress项目根目录，但有时vitepres是做其它项目的子项目，所以需要指定git运行的目录",
    //文档dev命令，默认值是vitepress项目根目录下的dev命令
    docsDev: "文档dev命令，默认值:" + defaultVpSimple.cmdDocsDev,
    //文档build命令，默认值是vitepress项目根目录下的build命令
    docsBuild: "文档build命令，默认值:" + defaultVpSimple.cmdDocsBuild,
  },
};
