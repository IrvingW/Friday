<template>
  <div id="aside">
    <el-menu
      class="el-menu-vertical-demo"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
      router=true
      :default-active="$route.path"
    >
      <el-menu-item index="/friday/home/index">
        <i class="el-icon-s-home"></i>
        <span slot="title">首页</span>
      </el-menu-item>
      <el-menu-item index="/friday/home/cluster_view">
        <i class="el-icon-menu"></i>
        <span slot="title">集群视图</span>
      </el-menu-item>
      <el-menu-item index="/friday/home/task_view">
        <i class="el-icon-s-order"></i>
        <span slot="title">任务中心</span>
      </el-menu-item>
      <el-menu-item index="/friday/home/image_view">
        <i class="el-icon-s-tools"></i>
        <span slot="title">镜像管理</span>
      </el-menu-item>
      <el-menu-item index="/friday/home/user_view">
        <i class="el-icon-user-solid"></i>
        <span slot="title">用户中心</span>
      </el-menu-item>
    </el-menu>
  </div>
</template>
<script type="text/javascript">

  export default {
    name: 'sidebar',
    data() {
      return {
      }
    },
    created() {
    },
  }
</script>
<style type="text/css" scoped>
  /*a {*/
    /*text-decoration: none*/
  /*}*/

  .aside {
    width: 100%;
    min-height: calc(100vh - 60px);
    background-color: rgb(244, 251, 254);
    overflow: hidden;
  }


</style>
