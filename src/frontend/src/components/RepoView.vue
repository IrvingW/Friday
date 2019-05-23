<template>
<el-container>
  <el-container>
    <el-header>
      <!-- Header content -->
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/friday/home/repo_list' }">镜像仓库</el-breadcrumb-item>
        <el-breadcrumb-item :to="{ path: '/friday/home/repo_view' }">仓库信息</el-breadcrumb-item>
        <el-breadcrumb-item> </el-breadcrumb-item>
      </el-breadcrumb>
    </el-header>
    <el-main height="">
      <el-row>
        <el-col :span="18">
          <el-card :body-style="{ padding: '20px' }" shadow="hover" style="margin-bottom: 40px">
            <el-row :gutter="40" style="margin-top: 15px; margin-left:15px; margin-bottom:15px">
              <el-col :span="4">
                <el-image
                  style="width: 150px; height: 150px;"
                  :src="imageSrc" fit="cover">
                </el-image>
              </el-col>
              <el-col :span="14">
                <el-row style="margin-top: 10px">
                  <span class="title">{{repoInfo.name}}</span>
                </el-row>
                <el-row style="margin-top: 10px">
                  <span class="intro">简介: {{repoInfo.introduce}}</span>
                </el-row>
                <el-row style="margin-top: 10px">
                  <span class="intro">所有者: {{repoInfo.owner.user_name}}</span>
                </el-row>
              </el-col>
              <el-col :span="2">
                <el-button type="primary" @click="uploadImage" style="margin-top: 20px">上传镜像</el-button>
              </el-col>
            </el-row>
          </el-card>
        </el-col>
      </el-row>
      <!-- Main content -->
      <el-card :body-style="{ padding: '40px' }" shadow="hover">
        <el-tabs active-name="description">
          <el-tab-pane label="说明" name="description">
            <mavon-editor
              class="md"
              :value="repoInfo.description"
              :subfield="false"
              :defaultOpen="defaultOpen"
              :toolbarsFlag ="false"
              :editable="false"
              :scrollStyle="true"></mavon-editor>
          </el-tab-pane>
          <el-tab-pane label="镜像列表" name="tags">
            <el-row>
              <el-col :span="18">
                <el-table max-height="380px"
                  :data="imageList" stripe style="width: 100%">
                  <el-table-column
                    prop="tag"
                    label="标签"
                    width="180">
                  </el-table-column>
                  <el-table-column
                    prop="create_time"
                    label="上传时间"
                    width="180">
                  </el-table-column>
                  <el-table-column
                    prop="description"
                    label="描述"
                    width="180">
                  </el-table-column>
                  <el-table-column
                    prop="download"
                    label="下载量"
                    width="180">
                  </el-table-column>
                  <el-table-column label="操作" width="180">
                    <el-button type="success" slot-scope="scope"
                      @click="createTask(scope.$index, scope.row)">创建任务</el-button>
                  </el-table-column>
                </el-table>
              </el-col>
            </el-row>
          </el-tab-pane>
        </el-tabs>
      </el-card>
    </el-main>
  </el-container>
</el-container>

</template>
<script>
  export default {
	name:'repo_view',
	  data(){
		  return{
        repoId: -1,
        repoInfo: {
        },
        defaultOpen: "preview",
        imageSrc: '/api/picture/show?picture=',
        imageList: [
        ],
        repoInfoUrl: '/api/image/repo/info?repo_id='
		  }
    },
    methods: {
      uploadImage() {
        this.$router.push({path: '/friday/home/upload_image', query: {repo_id: this.repoId}})
      },
      createTask(index, row) {
        console.log(row.id)
      }
    },
    created: function () {
      this.repoId = this.$route.query.repo_id
      this.$http.get(this.repoInfoUrl + this.repoId, {withCredentials: true}).then((response) => {
			  if (response.status != 200){
				  this.$message.error({
					  message: '请求发送失败，请联系管理员'
				  })
			  }else{
				  var body = response.body
				  if (body.Rtn == 0){
            this.repoInfo = body.Repo
            this.imageSrc += this.repoInfo.picture
            this.imageList = body.Repo.images
            this.imageList.forEach(element => {
              element.create_time = element.create_time.split("T")[0]
            });
				  }else{
					  this.$message.error({
                 message: body.Msg
						  })
				  }
			  }
      })
    }
  }
</script>
<style type="text/css" scoped>
.title {
    font-weight: bold;
    font-family: "Helvetica Neue";
    font-size: 25px;
}

.intro {
    font-style: italic;
    font-size: 16px;
}
</style>