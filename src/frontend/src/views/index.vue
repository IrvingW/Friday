<template>
  <el-container>
    <el-header>
      <NavMenu></NavMenu>
    </el-header>
    <el-container>
      <el-aside width='300px'>
          <Sidebar></Sidebar>
      </el-aside>
      <el-container>
        <el-main class="main">
          <Content></Content>
        </el-main>
      </el-container>

    </el-container>
  </el-container>
</template>
<script type="text/javascript">
  import NavMenu from '@/components/NavMenu'
  import Content from '@/components/Content'
  import Sidebar from '@/components/Sidebar'
  export default {
    name: 'manage',
    components: {NavMenu,Content,Sidebar}
  }
</script>
<style type="text/css" scoped>
  .el-header{
    padding: 0px;
    margin-bottom: 10px;
  }

</style>