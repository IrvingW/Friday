<template>
<el-container>
  <el-container>
    <el-header>
      <!-- Header content -->
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/friday/home/image_view' }">镜像信息</el-breadcrumb-item>
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
              <el-col :span="4">
                <el-row style="margin-top: 10px">
                  <span class="title">image name</span>
                </el-row>
                <el-row style="margin-top: 10px">
                  <span class="intro">introduce</span>
                </el-row>
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
              :value="description"
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
                  :data="tagList" stripe style="width: 100%">
                  <el-table-column
                    prop="name"
                    label="标签"
                    width="180">
                  </el-table-column>
                  <el-table-column
                    prop="upload_time"
                    label="上传时间"
                    width="180">
                  </el-table-column>
                  <el-table-column
                    prop="download"
                    label="下载量"
                    width="180">
                  </el-table-column>
                  <el-table-column label="操作" width="180">
                    <el-button type="success"
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
	name:'image_view',
	  data(){
		  return{
        description: "## 哈哈",
        defaultOpen: "preview",
        imageSrc: 'https://cube.elemecdn.com/6/94/4d3ea53c084bad6931a56d5158a48jpeg.jpeg',
        tagList: [
          {name: "latest", upload_time: "2019-5-12", download: 10},
        ]
		  }
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
    font-size: 18px;
}
</style>