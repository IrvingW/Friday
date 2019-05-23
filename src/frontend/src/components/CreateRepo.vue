<template>
<el-container>
  <el-container>
    <el-header>
      <!-- Header content -->
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/friday/home/repo_list' }">镜像仓库</el-breadcrumb-item>
        <el-breadcrumb-item :to="{ path: '/friday/home/create_repo' }">创建仓库</el-breadcrumb-item>
        <el-breadcrumb-item> </el-breadcrumb-item>
      </el-breadcrumb>
    </el-header>
    <el-main>
      <!-- Main content -->
      <el-card :body-style="{ padding: '30px' }" shadow="hover">
        <el-row :gutter="10" style="margin-top: 30px">
          <el-col :span="2" :offset="1">
            <span>上传封面:</span>
          </el-col>
          <el-col :span="12">
            <el-upload
              class="avatar-uploader"
              action="/api/picture/save"
              :show-file-list="false"
              :on-success="changeImageUrl">
              <img v-if="imageUrl" :src="imageUrl" class="avatar">
              <i v-else class="el-icon-plus avatar-uploader-icon"></i>
            </el-upload>
          </el-col>
        </el-row>
        <el-row style="margin-top: 40px">
            <el-col :span="20" :offset="1">
                <el-form :model="repoInfo" status-icon :rules="rules" ref="repoInfo" label-width="100px">
                  <el-form-item label="仓库名" prop="repo_name">
                    <el-input v-model="repoInfo.repo_name" autocomplete="off" placeholder="请输入镜像仓库名">
                        <i slot="prefix" class="el-icon-s-home"></i>
                    </el-input>
                  </el-form-item>
                  <el-form-item label="简介" prop="introduce">
                    <el-input v-model="repoInfo.introduce" autocomplete="off" placeholder="请输入对仓库的介绍">
                        <i slot="prefix" class="el-icon-s-comment"></i>
                    </el-input>
                  </el-form-item>
                  <el-form-item label="标签" prop="labels">
                      <el-input type="textarea"
                       :autosize="{ minRows: 2, maxRows: 4}"
                       placeholder="请输入标签，以逗号分隔"
                       v-model="repoInfo.labels">
                        <i slot="prefix" class="el-icon-price-tag"></i>
                     </el-input>
                  </el-form-item>
                  <el-form-item label="描述" prop="description">
                    <mavon-editor v-model="repoInfo.description" :scrollStyle="true"></mavon-editor>
                  </el-form-item>
                  <el-form-item class="button-row">
                    <el-button type="success" @click="submitForm('repoInfo')">提交</el-button>
                    <el-button @click="resetForm('repoInfo')">重置</el-button>
                  </el-form-item>
                </el-form>
            </el-col>
        </el-row>
      </el-card>
    </el-main>
  </el-container>
</el-container>

</template>
<script>
  export default {
	name:'create_repo',
	  data(){
		return{
          repoInfo: {
            repo_name: "",
            introduce: "",
            description: "",
            labels: "",
            picture: ""
          },
          rules: {
            repo_name: [
              { required: true, message: '镜像仓库名不能为空', trigger: 'blur' }
            ],
            introduce: [
              { required: true, message: '镜像仓库简介不能为空', trigger: 'blur' }
            ],
            labels: [
              { required: true, message: '镜像仓库标签不能为空', trigger: 'blur' }
            ],
          },
          imageUrl: "",
          createRepoUrl: "/api/image/repo/create"
		}
    },
    methods: {
      changeImageUrl(response, file) {
		if (response.Rtn == 0){
          this.repoInfo.picture = response.PictureName
          this.imageUrl = "/api/picture/show?picture_name=" + response.PictureName
		  this.$message({
		    message: "文件上传成功",
            type: 'success'
		  });
	    }else{
		  this.$message.error({
            message: response.Msg
	       })
	    }
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.$http.post(this.createRepoUrl, this.repoInfo,
              {emulateJSON: true, withCredentials: true}).then((response) => {
				if (response.status != 200){
				  this.$message.error({
					message: '请求发送失败'
			      })
				}
				var body = response.body
				if (body.Rtn == 0){
			      this.$message({
                    message: '镜像仓库创建成功',
                    type: 'success'
				  });
				  this.$router.push({path: "/friday/home/repo_list"})
				}else if (body.Rtn < 0){
				  this.$message.error({
                    message: body.Msg
                  });
				}else {
				  this.$message({
					message: body.Msg,
					type: 'warning'
				  })
				}
              })
          } else {
              this.$message.error({
                  message: '请正确填写表单'
              })
            return false;
          }
        });
      },
    }
  }
</script>
<style type="text/css" scoped>
  .avatar-uploader {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    width: 190px;
  }
  .avatar-uploader:hover {
    border-color: #409EFF;
  }
  .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
</style>