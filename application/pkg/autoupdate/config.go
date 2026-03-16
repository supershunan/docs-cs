package autoupdate

// UpdateCfg 更新配置的默认值
// 这里配置为使用你自己的 GitHub 仓库：supershunan/vitepress-simple
var UpdateCfg = UpdateConfig{
	AppName:                     "方向图文档发布系统", // 程序名称（最终 exe 名，不带 .exe）
	CurrVersion:                 "v1.0.0",         // 当前版本，对应 GitHub Release 的 tag
	GitType:                     GitTypeGithub,    // 使用 GitHub 作为更新源
	GitOwner:                    "supershunan",    // 你的 GitHub 用户名
	GitRepo:                     "vitepress-simple", // 仓库名
	WindowReleaseNameContainStr: ".exe",           // 发行版名称包含该字符串则认为：windows 系统
	MacReleaseNameContainStr:    ".zip",           // 发行版名称包含该字符串则认为：macOS 系统
}

// SetUpdateConfig 设置当前程序更新信息 建议在检查更新前进行设置【外部调用】
func SetUpdateConfig(config UpdateConfig) {
	// 程序名 & 版本号
	UpdateCfg.AppName = ifThenStr(config.AppName != "", config.AppName, "方向图文档发布系统")
	UpdateCfg.CurrVersion = ifThenStr(config.CurrVersion != "", config.CurrVersion, "v1.0.0")
	// 更新匹配字符串
	UpdateCfg.WindowReleaseNameContainStr = ifThenStr(
		config.WindowReleaseNameContainStr != "",
		config.WindowReleaseNameContainStr,
		".exe",
	)
	UpdateCfg.MacReleaseNameContainStr = ifThenStr(
		config.MacReleaseNameContainStr != "",
		config.MacReleaseNameContainStr,
		".zip",
	)
	// git 仓库信息
	UpdateCfg.GitType = GitTypeEnum(ifThenInt(config.GitType != 0, int(config.GitType), GitTypeGithub))
	UpdateCfg.GitOwner = ifThenStr(config.GitOwner != "", config.GitOwner, "supershunan")
	UpdateCfg.GitRepo = ifThenStr(config.GitRepo != "", config.GitRepo, "vitepress-simple")
}

// GitTypeEnum git源
type GitTypeEnum int

const (
	GitTypeGitee  = iota
	GitTypeGithub = iota
)

type UpdateConfig struct {
	AppName                     string      //程序名称
	CurrVersion                 string      //当前版本
	GitType                     GitTypeEnum //git源
	GitOwner                    string      //仓库拥有者
	GitRepo                     string      //仓库名
	WindowReleaseNameContainStr string      //发行版名称包含该字符串则认为：windows系统
	MacReleaseNameContainStr    string      //发行版名称包含该字符串则认为：苹果系统
}
