<template>
  <a-dropdown class="mx-1" arrow>
    <icon-park
      class="cursor-pointer select-none"
      :size="storeLayout.editorToolIconSize"
      fill="black"
      strokeLinejoin="bevel"
      theme="outline"
      type="terminal"
    />
    <template #overlay>
      <a-menu>
        <!--        终端配置 菜单-->
        <a-menu-item @click="isShowShellEdit = true">
          <a href="javascript:;">{{ lang("pageIndex.shell.config") }}</a>
        </a-menu-item>
        <!--        运行日志 菜单-->
        <a-menu-item @click="isShowShellLog = true">
          <a href="javascript:;">{{ lang("pageIndex.shell.runLog") }}</a>
        </a-menu-item>
        <hr />
        <!--列表 菜单-->
        <div v-for="(item, index) in cmdList" :key="index">
          <a-menu-item @click="createAndRunCmd(index)">
            <a href="javascript:;">{{ item.menuLabel }}</a>
          </a-menu-item>
          <hr v-if="[2].includes(index)" />
        </div>
      </a-menu>
    </template>
  </a-dropdown>

  <!--  弹出层：编辑shell配置-->
  <a-modal
    :style="StyleNoDrag"
    @ok="saveConfig"
    ok-text="Save"
    width="40%"
    :cancel-button-props="{ style: 'display:none' }"
    :ok-button-props="{ class: 'bg-blue-500' }"
    v-model:open="isShowShellEdit"
    :title="lang('pageIndex.shell.config')"
  >
    <!--    运行的的起始位置-->
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.shellBaseDir"
      :tooltip="lang('pageIndex.shell.tips.runPath')"
      :label="lang('pageIndex.shell.labels.runPath')"
    ></sim-input>
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.gitBaseDir"
      :tooltip="lang('pageIndex.shell.tips.gitPath')"
      :label="lang('pageIndex.shell.labels.gitPath')"
    ></sim-input>
    <!--    文档dev-->
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.cmdDocsDev"
      :tooltip="lang('pageIndex.shell.tips.docsDev')"
      :label="lang('pageIndex.shell.labels.docsDev')"
    ></sim-input>
    <!--    文档build-->
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.cmdDocsBuild"
      :tooltip="lang('pageIndex.shell.tips.docsBuild')"
      :label="lang('pageIndex.shell.labels.docsBuild')"
    ></sim-input>
    <!--    npm install-->
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.cmdNpmInstall"
      :tooltip="lang('pageIndex.shell.default') + defaultVpSimple.cmdNpmInstall"
      label="npm install"
    ></sim-input>
    <!--    git pull-->
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.cmdGitPull"
      :tooltip="lang('pageIndex.shell.default') + defaultVpSimple.cmdGitPull"
      label="git pull"
    ></sim-input>
    <!--    git add-->
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.cmdGitAdd"
      :tooltip="lang('pageIndex.shell.default') + defaultVpSimple.cmdGitAdd"
      label="git add"
    ></sim-input>
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.cmdGitCommit"
      :tooltip="lang('pageIndex.shell.default') + defaultVpSimple.cmdGitCommit"
      label="git commit"
    ></sim-input>
    <!--    git push-->
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.cmdGitPush"
      :tooltip="lang('pageIndex.shell.default') + defaultVpSimple.cmdGitPush"
      label="git push"
    ></sim-input>
    <!--    git init-->
    <sim-input
      :is-full-width="true"
      v-model="storeShell.vpsimpleConfig.cmdGitInit"
      :tooltip="lang('pageIndex.shell.default') + defaultVpSimple.cmdGitInit"
      label="git init"
    ></sim-input>
  </a-modal>

  <!--  弹出层运行日志-->
  <a-modal
    :style="StyleNoDrag"
    width="50%"
    :cancel-button-props="{ style: 'display:none' }"
    :ok-button-props="{ class: 'bg-blue-500', style: 'display:none' }"
    v-model:open="isShowShellLog"
    :title="lang('pageIndex.shell.runLog')"
  >
    <!--    顶部终端列表tabs-->
    <a-radio-group v-model:value="currShellIndex" button-style="solid">
      <a-radio-button
        v-for="(value, key) in storeShell.messages"
        :key="key"
        :value="Number(key)"
        >shell{{ key }}
      </a-radio-button>
    </a-radio-group>

    <!--    黑色背景的终端日志-->
    <div
      v-show="Object.keys(storeShell.messages).length > 0"
      style="height: 45vh; overflow-y: auto"
      class="px-5 pt-5 pb-10 mt-4 text-white bg-black"
    >
      <pre v-html="storeShell.messages[currShellIndex]"></pre>
    </div>
    <!--    底部按钮操作组-->
    <div
      v-show="
        Object.keys(storeShell.messages).length > 0 &&
        storeIndex.systemType != SystemWindows
      "
      class="flex justify-end items-center mt-5"
    >
      <!--          <a-button class="mx-2 bg-orange-200">移除当前终端</a-button>-->
      <a-button
        @click="storeShell.stopCmd(currShellIndex)"
        class="mx-2 bg-red-200"
        >中断
      </a-button>
    </div>

    <!--    空终端-->
    <div v-show="Object.keys(storeShell.messages).length == 0">
      <a-empty :description="lang('pageIndex.shell.noRunTerminal')"></a-empty>
    </div>
  </a-modal>

  <!--  Git Commit 提交描述弹窗-->
  <a-modal
    :style="StyleNoDrag"
    v-model:open="isShowCommitModal"
    :title="lang('pageIndex.shell.commitMessageTitle')"
    :ok-button-props="{ class: 'bg-blue-500' }"
    :ok-text="lang('common.confirm')"
    :cancel-text="lang('common.cancel')"
    @ok="confirmCommit"
  >
    <a-textarea
      v-model:value="commitMessage"
      :placeholder="lang('pageIndex.shell.commitMessagePlaceholder')"
      :rows="4"
      class="mt-2"
    />
  </a-modal>
</template>
<script setup lang="ts">
import { IconPark } from "@icon-park/vue-next/es/all";
import { useLayoutStore } from "@/store/layout";
import { onMounted, ref } from "vue";

import SimInput from "@/components/simInput.vue";
import { useVpconfigStore } from "@/store/vpconfig";
import { IsEmptyValue } from "@/utils/utils";
import { useIndexStore } from "@/store";
import { defaultVpSimple } from "@/configs/defaultVpSimple";

import { useShellStore } from "@/store/shell";
import { ToastError } from "@/utils/Toast";
import { SystemWindows } from "@/constant/enums/system";
import { StyleNoDrag } from "@/configs/cnts";
import { lang } from "../../utils/language";

const storeIndex = useIndexStore();
const storeLayout = useLayoutStore();
const storeConfig = ref(useVpconfigStore());
const storeShell = useShellStore();
const currShellIndex = ref(0);
const isShowShellEdit = ref(false);
const isShowShellLog = ref(false);
const isShowCommitModal = ref(false);
const commitMessage = ref("");
const pendingCommitIndex = ref(-1);

/** cmdList 中 git commit 的索引 */
const GIT_COMMIT_INDEX = 5;

interface CmdItem {
  menuLabel: string;

  cmd: string;
  isAlone: boolean;
  isSystemShell: boolean;
}

const cmdList = ref<CmdItem[]>([]);

// const vpsimpleConfig = ref<VPSimpleConfig>({} as VPSimpleConfig);
onMounted(async () => {
  await storeIndex.loadTreeData();
  storeShell.loadVpSimpleConfig();
  setCmdList();
});

const setCmdList = () => {
  cmdList.value = [
    {
      menuLabel: "npm install",
      cmd: storeShell.vpsimpleConfig.cmdNpmInstall,
      isAlone: false,
      isSystemShell: false,
    },
    {
      menuLabel: "npm run docs:dev",
      cmd: storeShell.vpsimpleConfig.cmdDocsDev,
      isAlone: true,
      isSystemShell: false,
    },
    {
      menuLabel: "npm run docs:build",
      cmd: storeShell.vpsimpleConfig.cmdDocsBuild,
      isAlone: false,
      isSystemShell: false,
    },

    {
      menuLabel: "git init",
      cmd: storeShell.vpsimpleConfig.cmdGitInit,
      isAlone: false,
      isSystemShell: false,
    },
    {
      menuLabel: "git add",
      cmd: storeShell.vpsimpleConfig.cmdGitAdd,
      isAlone: false,
      isSystemShell: false,
    },
    {
      menuLabel: "git commit",
      cmd: storeShell.vpsimpleConfig.cmdGitCommit,
      isAlone: false,
      isSystemShell: false,
    },
    {
      menuLabel: "git push",
      cmd: storeShell.vpsimpleConfig.cmdGitPush,
      isAlone: false,
      isSystemShell: false,
    },

    {
      menuLabel: "git pull",
      cmd: storeShell.vpsimpleConfig.cmdGitPull,
      isAlone: false,
      isSystemShell: false,
    },
  ];
};
// 创建并运行
//cmd
//isAlone 是否单独创建一个终端(如果是则不显示日志)
//isSystemShell 是否使用系统自带的终端
const createAndRunCmd = async (index: number) => {
  const cmd = cmdList.value[index].cmd;
  if (IsEmptyValue(storeShell.vpsimpleConfig.shellBaseDir)) {
    ToastError(lang("pageIndex.shell.runPathEmpty"));
    return;
  }
  if (IsEmptyValue(cmd)) {
    ToastError(lang("pageIndex.shell.cmdEmpty"));
    return;
  }

  // git commit 先弹出输入框，确认后再执行
  if (index === GIT_COMMIT_INDEX) {
    pendingCommitIndex.value = index;
    commitMessage.value = "";
    isShowCommitModal.value = true;
    return;
  }

  await doRunCmd(index);
};

/** 实际执行命令（含 git commit 确认后的执行） */
const doRunCmd = async (index: number) => {
  const item = cmdList.value[index];
  let cmd = item.cmd;
  const isAlone = item.isAlone;
  const isSystemShell = item.isSystemShell;

  if (index === GIT_COMMIT_INDEX) {
    const msg = commitMessage.value?.trim() || "";
    if (!msg) {
      ToastError(lang("pageIndex.shell.commitMessageEmpty"));
      return;
    }
    cmd = `git commit -m "${msg.replace(/"/g, '\\"')}"`;
  }

  const cmdDir = cmd.trim().startsWith("git")
    ? storeShell.vpsimpleConfig.gitBaseDir
    : storeShell.vpsimpleConfig.shellBaseDir;

  if (isSystemShell) {
    currShellIndex.value = await storeShell.createSystemShellAndRun(
      cmdDir,
      cmd,
    );
  } else {
    currShellIndex.value = await storeShell.createShellAndRun(
      cmdDir,
      cmd,
      isAlone,
    );
  }
  if (isAlone && storeIndex.systemType == SystemWindows) {
    isShowShellLog.value = false;
  } else {
    isShowShellLog.value = true;
  }
};

const confirmCommit = async () => {
  if (pendingCommitIndex.value < 0) return;
  const msg = commitMessage.value?.trim();
  if (!msg) {
    ToastError(lang("pageIndex.shell.commitMessageEmpty"));
    return Promise.reject(); // 阻止弹窗关闭
  }
  isShowCommitModal.value = false;
  pendingCommitIndex.value = -1;
  await doRunCmd(GIT_COMMIT_INDEX);
};

const saveConfig = () => {
  storeConfig.value.configData["vpsimple"] = storeShell.vpsimpleConfig;
  storeConfig.value.saveConfig();
  setCmdList();
  isShowShellEdit.value = false;
};
</script>
<style scoped>
pre {
  white-space: pre-wrap; /* 保留空格和换行，并允许文本换行 */
  font-family: monospace; /* 使用等宽字体 */
}
</style>
