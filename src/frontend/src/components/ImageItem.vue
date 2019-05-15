<template>
  <el-container>
    <el-main>
      <el-card :body-style="{ padding: '10px' }" shadow="hover">
        <el-row class="button-row">
          <el-col :span="3">
            <el-image
              style="width: 100px; height: 100px"
              :src="src" fit="cover">
              <div slot="error" class="image-slot">
                <i class="el-icon-picture-outline"></i>
              </div>
            </el-image>
          </el-col>
          <el-col :span="14" style="margin-left: 20px">
            <el-row>
                <span class="title-name">ImageName</span>
            </el-row>
            <el-row style="margin: 10px">
                <span class="author-text">由XX于XXXX最后更新</span>
            </el-row>
            <el-row style="margin-left: 10px; margin-top: 30px">
                <span class="introduce-text">简介</span>
            </el-row>
            <el-row style="margin-top: 20px">
                <el-tag v-for="tag in tags" :key="tag.name" type="info">{{tag.name}}</el-tag>
            </el-row>
          </el-col>
          <el-col :span="3">
              <el-button type="primary" @click="showDetail">查看详情</el-button>
          </el-col>
          <el-col :span="1">
              <el-button type="success" @click="createTask">构建任务</el-button>
          </el-col>
        </el-row>
      </el-card>
    </el-main>
  </el-container>
</template>
<script>
  export default {
	name:'image_item',
	data(){
      return{
          src: 'https://cube.elemecdn.com/6/94/4d3ea53c084bad6931a56d5158a48jpeg.jpeg',
          tags: [
              {name: '数据库'}
          ]
      }
    },
    props:{
        image_src: {
            type: String
        },
        name: {
            type: String
        },
        introduce: {
            type: String
        },
        image_id: {
        }
    }
  }
</script>
<style type="text/css" scoped>
.button-row {
    margin: 20px;
}
.title-name {
    font-weight: bold;
    font-family: "Helvetica Neue";
    font-size: 20px;
}
.introduce-text {
    font-size: 16px;
}
.author-text {
    font-style: italic;
    font-size: 14px;
}
</style>