<template>
  <el-container>
    <el-header>
      <el-menu
        :default-active="activeIndex"
        class="el-menu-demo"
        mode="horizontal"
        :ellipsis="false"
        @select="handleSelect"
      >
        <el-menu-item index="home">首页</el-menu-item>
        <el-menu-item index="currencyexchange">兑换货币</el-menu-item>
        <el-menu-item index="news">新闻</el-menu-item>
        <el-menu-item index="login" v-if="!authStore.isAuthenticated">登录</el-menu-item>
        <el-menu-item index="register" v-if="!authStore.isAuthenticated">注册</el-menu-item>
        <el-menu-item index="logout" v-if="authStore.isAuthenticated">退出</el-menu-item>
      </el-menu>
    </el-header>

    <el-main>
      <router-view></router-view>
    </el-main>

    <el-footer>
    </el-footer>
  </el-container>
</template>


<script setup lang="ts">
import { useRouter, useRoute } from "vue-router";
import { useAuthStore } from "./store/auth";
import { ref, watch } from 'vue';

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore();
const activeIndex = ref(route.name?.toString() || 'home');

watch(route, (newRoute) => {
  activeIndex.value = newRoute.name?.toString() || 'home';
});

const handleSelect = (key: string) => {
  if ( key === 'logout') {
    authStore.logout();
    router.push({ name: 'Home' });
  } else {
    router.push({ name:  key.charAt(0).toUpperCase() +  key.slice(1) });
  }
};
</script>

<style scoped>
.el-menu-demo {
  line-height: 60px;
}
</style>