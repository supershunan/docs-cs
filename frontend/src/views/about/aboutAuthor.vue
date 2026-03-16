<script setup lang="ts">
import { OpenURL } from "../../../wailsjs/go/system/SystemService";
import { lang } from "../../utils/language";
import { StyleNoDrag } from "@/configs/cnts";
</script>

<template>
  <div class="flex justify-start p-3 mx-20 mt-4 border">
    <div class="my-2 w-full">
      {{ lang("pageAbout.authorBlog")
      }}<span class="cursor-pointer" @click="OpenURL('http://xiaod.co/')"
        >http://xiaod.co/</span
      >
    </div>
    <div class="my-2 w-full">
      {{ lang("pageAbout.authorGithub")
      }}<span
        class="cursor-pointer"
        @click="OpenURL('https://github.com/zhangdi168')"
        >https://github.com/zhangdi168</span
      >
    </div>
    <div class="my-2 w-full">
      {{ lang("pageAbout.authorGitee")
      }}<span
        class="cursor-pointer"
        @click="OpenURL('https://gitee.com/zhangdi168')"
        >https://gitee.com/zhangdi168</span
      >
    </div>
    <div class="my-2 w-full">
      {{ lang("pageAbout.authorWx")
      }}<span
        class="cursor-pointer"
        @click="OpenURL('https://gitee.com/zhangdi168')"
        >bcxiaod</span
      >
    </div>
  </div>
</template>

<style scoped>
span {
  color: #007bff;
  text-decoration: none;
}

span:hover {
  text-decoration: underline;
}
</style>
