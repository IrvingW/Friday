<template>
  <div id="aside">
    <el-menu
      class="el-menu-vertical-demo"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
      :router="router"
      :default-active="$route.path"
    >
      <el-menu-item index="/friday/home/index">
        <i class="el-icon-s-home"></i>
        <span slot="title">首页</span>
      </el-menu-item>
      <el-submenu index="">
        <template slot="title">
          <i class="el-icon-s-shop"></i>
          <span slot="title">集群视图</span>
        </template>
        <el-menu-item index="/friday/home/cluster_view">
          <i class="el-icon-menu"></i>
          <span slot="title">概览</span>
        </el-menu-item>
        <el-menu-item index="/friday/home/cluster_list">
          <i class="el-icon-document-copy"></i>
          <span slot="title">集群列表</span>
        </el-menu-item>
        <el-menu-item index="/friday/home/add_cluster">
          <i class="el-icon-folder-add"></i>
          <span slot="title">添加集群</span>
        </el-menu-item>
        <el-menu-item index="/friday/home/node_view">
          <i class="el-icon-s-platform"></i>
          <span slot="title">节点列表</span>
        </el-menu-item>
      </el-submenu>
      <el-submenu index="">
        <template slot="title">
          <i class="el-icon-s-claim"></i>
          <span slot="title">任务中心</span>
        </template>
        <el-menu-item index="/friday/home/task_list">
          <i class="el-icon-document-copy"></i>
          <span slot="title">任务列表</span>
        </el-menu-item>
        <el-menu-item index="/friday/home/create_task">
          <i class="el-icon-edit"></i>
          <span slot="title">创建任务</span>
        </el-menu-item>
      </el-submenu>
      <el-submenu index="">
        <template slot="title">
          <i class="el-icon-s-tools"></i>
          <span slot="title">镜像管理</span>
        </template>
        <el-menu-item index="/friday/home/image_list">
          <i class="el-icon-s-cooperation"></i>
          <span slot="title">镜像仓库</span>
        </el-menu-item>
        <el-menu-item index="/friday/home/image_view">
          <i class="el-icon-info"></i>
          <span slot="title">镜像信息</span>
        </el-menu-item>
        <el-menu-item index="/friday/home/upload_image">
          <i class="el-icon-upload"></i>
          <span slot="title">上传镜像</span>
        </el-menu-item>
      </el-submenu>
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
        router: true
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
