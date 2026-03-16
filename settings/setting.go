package setting

import "wailstemplate/application/db/entity"

const Version = "v1.0.0"

const AppName = "方向图文档发布系统"
const GitRepo = "VitePressSimple"
const GitAuthor = "supershunan"
const WindowZipContainStr = "windows_"
const MacosZipContainStr = "mac_"

// EntityAutoMigrateList 自动迁移的实体列表
// var EntityAutoMigrateList = []any{new(entity.User)}
var EntityAutoMigrateList = []any{new(entity.Article)}
