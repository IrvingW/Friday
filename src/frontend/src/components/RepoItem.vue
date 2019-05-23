<template>
  <el-container>
    <el-main>
      <el-card :body-style="{ padding: '10px' }" shadow="hover">
        <el-row class="button-row">
          <el-col :span="3">
            <el-image
              style="width: 100px; height: 100px"
              :src="imageSrc" fit="cover">
            </el-image>
          </el-col>
          <el-col :span="14" style="margin-left: 20px">
            <el-row>
                <span class="title-name">{{name}}</span>
            </el-row>
            <el-row style="margin: 10px">
                <span class="author-text">由{{owner}}于{{update_time}}最后更新</span>
            </el-row>
            <el-row style="margin-left: 10px; margin-top: 30px">
                <span class="introduce-text">{{introduce}}</span>
            </el-row>
            <el-row style="margin-top: 20px">
                <el-tag v-for="tag in tags" :key="tag.name" type="info" style="margin-left:15px">{{tag}}</el-tag>
            </el-row>
          </el-col>
          <el-col :span="3">
              <el-button type="primary" @click="showDetail">查看详情</el-button>
          </el-col>
          <el-col :span="1">
              <el-button type="success" @click="createTask">快速构建</el-button>
          </el-col>
        </el-row>
      </el-card>
    </el-main>
  </el-container>
</template>
<script>
  export default {
	name:'repo_item',
	data(){
      return{
          imageSrc: '/api/picture/show?picture_name=',
          tags: []
      }
    },
    props:["name", "introduce", "picture", "update_time", "id", "owner", "labels"],
    methods: {
      showDetail() {
        this.$router.push({path: '/friday/home/repo_view', query: {repo_id: this.id}})
      }
    },
    created: function() {
      this.imageSrc += this.picture
      this.update_time = this.update_time.split("T")[0]
      var list = this.labels.split(",")
      list.forEach(element => {
        this.tags.push(element)
      });
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
    font-size: 26px;
}
.introduce-text {
    font-size: 18px;
}
.author-text {
    font-style: italic;
    font-size: 14px;
}
</style>