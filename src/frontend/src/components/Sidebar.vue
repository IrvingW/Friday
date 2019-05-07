<template>
  <div id="aside">
    <el-menu

      class="el-menu-vertical-demo"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
      :router="isActive"
    >

      <el-menu-item index="/uploadfile">
        <i class="el-icon-upload2"></i>
        <span slot="title">本地上传</span>
      </el-menu-item>
      <el-menu-item index="/remotedownload">
        <i class="el-icon-upload"></i>
        <span slot="title">远端上传</span>
      </el-menu-item>
      <el-menu-item index="/filesearch">
        <i class="el-icon-message"></i>
        <span slot="title">文件查看</span>
      </el-menu-item>
      <!--<el-menu-item index="/metadata">-->
        <!--<i class="el-icon-tickets"></i>-->
        <!--<span slot="title">元数据管理</span>-->
      <!--</el-menu-item>-->
      <!--<el-menu-item index="/conceptLayer">-->
        <!--<i class="el-icon-tickets"></i>-->
        <!--<span slot="title">概念层业务元数据管理</span>-->
      <!--</el-menu-item>-->
      <!--<el-menu-item index="/logicLayer">-->
        <!--<i class="el-icon-tickets"></i>-->
        <!--<span slot="title">逻辑层业务元数据管理</span>-->
      <!--</el-menu-item>-->
      <!--<el-menu-item index="/physicsLayer">-->
        <!--<i class="el-icon-tickets"></i>-->
        <!--<span slot="title">物理层元数据管理</span>-->
      <!--</el-menu-item>-->
      <el-menu-item index="/handledata">
        <i class="el-icon-printer"></i>
        <span slot="title">处理入库</span>
      </el-menu-item>
      <el-menu-item index="/inquire">
        <i class="el-icon-view"></i>
        <span slot="title">查询数据</span>
      </el-menu-item>


    </el-menu>
  </div>
</template>
<script type="text/javascript">

  export default {
    name: 'sidebar',
    data() {
      return {
        isActive: true,

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
