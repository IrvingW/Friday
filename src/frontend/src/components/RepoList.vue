<template>
<el-container>
  <el-container>
    <el-header height="20px">
      <!-- Header content -->
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/friday/home/repo_list' }">镜像仓库</el-breadcrumb-item>
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
              <el-button type="success" @click="labelFilter" style="margin-left: 60px">筛选</el-button>
            </div>
            <!-- card body -->
            <el-row style="margin-top: 10px">
              <el-checkbox v-model="checked1" label="数据采集" border size="medium"></el-checkbox>
              <el-checkbox v-model="checked2" label="批处理" border size="medium"></el-checkbox>
            </el-row>
            <el-row style="margin-top: 10px">
              <el-checkbox v-model="checked3" label="流处理" border size="medium"></el-checkbox>
              <el-checkbox v-model="checked4" label="服务型" border size="medium"></el-checkbox>
            </el-row>
          </el-card>
        </el-aside>
        <el-container>
          <el-main>
            <!-- Main content -->
            <el-row>
              <el-col :span="18" :offset="1">
                <p class="total">共计 {{total_cnt}} 个可用镜像仓库</p>
              </el-col>
              <el-col :span="4">
                  <el-input class="search" placeholder="搜索镜像名称" prefix-icon="el-icon-search" v-model="search_input"></el-input>
              </el-col>
            </el-row>
            <el-row v-for="repo in repoList" v-bind:key="repo.id">
              <RepoItem :name="repo.name" :introduce="repo.introduce" :picture="repo.picture"
                :update_time="repo.update_time" :owner="repo.owner.user_name" :id="repo.id"
                :labels="repo.labels.toString()"
              ></RepoItem>
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
  import RepoItem from '@/components/RepoItem'
  export default {
  name:'repo_list',
  components: {RepoItem},
	data(){
    return{
      total_cnt: 0,
      total_page: 1,
      current_page: 1,
      search_input: '',
      repoList: [],
      repoListUrl: '/api/image/repo/search?key=',
      labelListUrl: '/api/image/label/list'
    }
  },
  created: function (){
    // get repo list
    this.$http.get(this.repoListUrl, {withCredentials: true}).then((response) => {
			if (response.status != 200){
				this.$message.error({
					message: '请求发送失败，请联系管理员'
				})
			}else{
				var body = response.body
				if (body.Rtn == 0){
          this.repoList = body.RepoList
          this.total_cnt = this.repoList.length
				}else{
					this.$message.error({
               message: body.Msg
						})
				}
			}
    });

    // get label list
    this.$http.get(this.labelListUrl, {withCredentials: true}).then((response) => {
      console.log(response)
    });
  },
  methods: {
    handlePageChange(page) {
        console.log(`当前页: ${val}`);
    },
    labelListUrl() {

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