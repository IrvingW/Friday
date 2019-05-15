<template>
<el-container>
  <el-container>
    <el-header height="20px">
      <!-- Header content -->
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/friday/home/task_list' }">镜像列表</el-breadcrumb-item>
        <el-breadcrumb-item> </el-breadcrumb-item>
      </el-breadcrumb>
    </el-header>
    <el-main height="">
      <el-container>
        <el-aside width="300px">
          <!-- Aside content -->
          <el-card :body-style="{ padding: '20px' }" shadow="always" style="margin-top: 30px">
            <div slot="header">
              <span>标签筛选：</span>
            </div>
            <!-- card body -->
            <div style="margin-top: 20px">
              <el-checkbox v-model="checked3" label="备选项1" border size="medium"></el-checkbox>
              <el-checkbox v-model="checked4" label="备选项2" border size="medium"></el-checkbox>
            </div>
          </el-card>
        </el-aside>
        <el-container>
          <el-main>
            <!-- Main content -->
            <el-row>
              <el-col :span="18" :offset="1">
                <p class="total">共计 {{total_cnt}} 个可用镜像</p>
              </el-col>
              <el-col :span="4">
                  <el-input class="search" placeholder="搜索镜像名称" prefix-icon="el-icon-search" v-model="search_input"></el-input>
              </el-col>
            </el-row>
            <el-row>
              <ImageItem></ImageItem>
            </el-row>
            <el-row>
              <el-col :span="12" :offset="1">
                <el-pagination background layout="prev, pager, next"
                  :total="total_page"
                  :current-page.sync="current_page"
                  @current-change="handlePageChange">
                </el-pagination>
              </el-col>
            </el-row>
          </el-main>
        </el-container>
      </el-container>
      <!-- Main content -->
    </el-main>
  </el-container>
</el-container>

</template>
<script>
  import ImageItem from '@/components/ImageItem'
  export default {
  name:'image_list',
  components: {ImageItem},
	data(){
    return{
      total_cnt: 0,
      search_input: '',
      total_page: 10,
      current_page: 2
    }
  },
  methods: {
    handlePageChange(page) {
        console.log(`当前页: ${val}`);
    }
  }
  }
</script>
<style type="text/css" scoped>
.search {
  margin-top: 15px;
}
.total {
  font-weight: bold;
  font-style: italic;
}
</style>